package bitmovin

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/cbsinteractive/transcode-orchestrator/job"
)

func state(bitmovin model.Status) job.State {
	switch bitmovin {
	case model.Status_CREATED, model.Status_QUEUED:
		return job.StateQueued
	case model.Status_RUNNING:
		return job.StateStarted
	case model.Status_FINISHED:
		return job.StateFinished
	case model.Status_ERROR:
		return job.StateFailed
	case model.Status_CANCELED:
		return job.StateCanceled
	default:
		return job.StateUnknown
	}
}
