.PHONY: all testdeps lint test gotest build run

HTTP_ADDR ?= :8080
LOG_LEVEL ?= debug
CI_TAG ?= $(shell git describe --tags $(shell git rev-list --tags --max-count=1))

all: test

test:
	go test -count=1 -race ./...

coverage:
	go test -coverprofile=coverage.txt -covermode=atomic ./...

build:
	go build

run: build
	APP_LOG_LEVEL=$(LOG_LEVEL) ./transcode-orchestrator -addr $HTTP_ADDR
