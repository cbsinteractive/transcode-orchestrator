package main

import (
	"io/ioutil"
	"log"

	"github.com/NYTimes/gizmo/server"
	"github.com/cbsinteractive/transcode-orchestrator/config"
	"github.com/cbsinteractive/transcode-orchestrator/db"
	_ "github.com/cbsinteractive/transcode-orchestrator/provider/bitmovin"
	_ "github.com/cbsinteractive/transcode-orchestrator/provider/flock"
	_ "github.com/cbsinteractive/transcode-orchestrator/provider/hybrik"
	_ "github.com/cbsinteractive/transcode-orchestrator/provider/mediaconvert"
	"github.com/cbsinteractive/transcode-orchestrator/service"
	"github.com/google/gops/agent"
	"github.com/zsiec/pkg/tracing"
	"github.com/zsiec/pkg/xrayutil"
)

func main() {
	agent.Listen(agent.Options{})
	defer agent.Close()
	cfg := config.LoadConfig()
	server.Init("transcode-orchestrator", cfg.Server)
	server.Log.Out = ioutil.Discard

	logger, err := cfg.Log.Logger()
	if err != nil {
		log.Fatal(err)
	}

	if cfg.EnableXray {
		cfg.Tracer = xrayutil.XrayTracer{
			EnableAWSPlugins: cfg.EnableXrayAWSPlugins,
			InfoLogFn:        logger.Infof,
		}
	} else {
		cfg.Tracer = tracing.NoopTracer{}
	}

	err = cfg.Tracer.Init()
	if err != nil {
		logger.Fatalf("initializing tracer: %v", err)
	}
	store, err := db.NewClient(nil)
	if err != nil {
		logger.Fatalf("initializing db: %v", err)
	}
	service, err := service.NewTranscodingService(cfg, logger, store)
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
