package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/cbsinteractive/transcode-orchestrator/config"
	"github.com/cbsinteractive/transcode-orchestrator/db"
	"github.com/cbsinteractive/transcode-orchestrator/service"

	_ "github.com/cbsinteractive/transcode-orchestrator/provider/bitmovin"
	_ "github.com/cbsinteractive/transcode-orchestrator/provider/flock"
	_ "github.com/cbsinteractive/transcode-orchestrator/provider/hybrik"
	_ "github.com/cbsinteractive/transcode-orchestrator/provider/mediaconvert"
)

var addr = flag.String("addr", ":"+os.Getenv("HTTP_PORT"), "http listen address")

func main() {
	flag.Parse()

	cfg := config.LoadConfig()
	store, err := db.NewClient(nil)
	if err != nil {
		log.Fatalf("initializing db: %v", err)
	}
	srv := service.Server{
		Config: cfg,
		DB:     store,
	}
	log.Println(http.ListenAndServe(*addr, srv))
}
