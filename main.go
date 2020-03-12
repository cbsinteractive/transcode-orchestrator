package main

import (
	"io/ioutil"
	"log"
	"net"

	"github.com/NYTimes/gizmo/server"
	"github.com/aws/aws-xray-sdk-go/awsplugins/ec2"
	"github.com/aws/aws-xray-sdk-go/awsplugins/ecs"
	"github.com/aws/aws-xray-sdk-go/xray"
	"github.com/cbsinteractive/video-transcoding-api/config"
	_ "github.com/cbsinteractive/video-transcoding-api/provider/bitmovin"
	_ "github.com/cbsinteractive/video-transcoding-api/provider/hybrik"
	_ "github.com/cbsinteractive/video-transcoding-api/provider/mediaconvert"
	"github.com/cbsinteractive/video-transcoding-api/service"
	"github.com/google/gops/agent"
)

func main() {
	agent.Listen(agent.Options{})
	defer agent.Close()
	cfg := config.LoadConfig()
	server.Init("video-transcoding-api", cfg.Server)
	server.Log.Out = ioutil.Discard

	logger, err := cfg.Log.Logger()
	if err != nil {
		log.Fatal(err)
	}

	var emitter xray.Emitter
	if cfg.EnableXray {
		emitter, err = xray.NewDefaultEmitter(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 2000})
		if err != nil {
			logger.Fatalf("creating xray emitter: %v", err)
		}

		if cfg.EnableXrayAWSPlugins {
			ec2.Init()
			ecs.Init()
		}
	} else {
		emitter = &NoopEmitter{}
	}

	err = xray.Configure(xray.Config{
		ContextMissingStrategy: ctxMissingStrategy{},
		Emitter:                emitter,
	})
	if err != nil {
		logger.Fatalf("configuring xray: %v", err)
	}

	service, err := service.NewTranscodingService(cfg, logger)
	if err != nil {
		logger.Fatal("unable to initialize service: ", err)
	}
	err = server.Register(service)
	if err != nil {
		logger.Fatal("unable to register service: ", err)
	}
	err = server.Run()
	if err != nil {
		logger.Fatal("server encountered a fatal error: ", err)
	}
}

type NoopEmitter struct{}

func (te *NoopEmitter) Emit(*xray.Segment)                     {}
func (te *NoopEmitter) RefreshEmitterWithAddress(*net.UDPAddr) {}

type ctxMissingStrategy struct{}

func (s ctxMissingStrategy) ContextMissing(interface{}) {}
