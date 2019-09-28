package main

import (
	"io/ioutil"
	"log"

	"github.com/NYTimes/gizmo/server"
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
