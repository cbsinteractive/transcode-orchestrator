package bitmovin

import (
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/cbsinteractive/transcode-orchestrator/provider"
)

func status(bitmovin model.Status) provider.Status {
	switch bitmovin {
	case model.Status_CREATED, model.Status_QUEUED:
		return provider.StatusQueued
	case model.Status_RUNNING:
		return provider.StatusStarted
	case model.Status_FINISHED:
		return provider.StatusFinished
	case model.Status_ERROR:
		return provider.StatusFailed
	case model.Status_CANCELED:
		return provider.StatusCanceled
	default:
		return provider.StatusUnknown
	}
}
