package bitmovin

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/cbsinteractive/transcode-orchestrator/provider"
)

func state(bitmovin model.Status) provider.State {
	switch bitmovin {
	case model.Status_CREATED, model.Status_QUEUED:
		return provider.StateQueued
	case model.Status_RUNNING:
		return provider.StateStarted
	case model.Status_FINISHED:
		return provider.StateFinished
	case model.Status_ERROR:
		return provider.StateFailed
	case model.Status_CANCELED:
		return provider.StateCanceled
	default:
		return provider.StateUnknown
	}
}
