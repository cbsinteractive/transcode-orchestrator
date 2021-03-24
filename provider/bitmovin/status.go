package bitmovin

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/cbsinteractive/transcode-orchestrator/av"
)

func state(bitmovin model.Status) av.State {
	switch bitmovin {
	case model.Status_CREATED, model.Status_QUEUED:
		return av.StateQueued
	case model.Status_RUNNING:
		return av.StateStarted
	case model.Status_FINISHED:
		return av.StateFinished
	case model.Status_ERROR:
		return av.StateFailed
	case model.Status_CANCELED:
		return av.StateCanceled
	default:
		return av.StateUnknown
	}
}
