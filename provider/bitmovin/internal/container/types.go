package container

import "github.com/bitmovin/bitmovin-api-sdk-go/model"

// HLSContainerAPI holds underlying api interfaces for HLS outputs
type HLSContainerAPI struct {
	HLSAudioMedia HLSAudioMediaAPI
	TSMuxing      TSMuxingAPI
	HLSStreams    HLSStreamsAPI
}

// HLSAudioMediaAPI contains methods for managing HLS Media Audio objects
type HLSAudioMediaAPI interface {
	Create(string, model.AudioMediaInfo) (*model.AudioMediaInfo, error)
}

// TSMuxingAPI contains methods for managing TS muxing objects
type TSMuxingAPI interface {
	Create(string, model.TsMuxing) (*model.TsMuxing, error)
}

// HLSStreamsAPI contains methods for managing HLS stream objects
type HLSStreamsAPI interface {
	Create(string, model.StreamInfo) (*model.StreamInfo, error)
}

// CMAFContainerAPI holds underlying api interfaces for CMAF outputs
type CMAFContainerAPI struct {
	HLSAudioMedia HLSAudioMediaAPI
	CMAFMuxing    CMAFMuxingAPI
	HLSStreams    HLSStreamsAPI
}

// CMAFMuxingAPI contains methods for managing CMAF muxing objects
type CMAFMuxingAPI interface {
	Create(string, model.CmafMuxing) (*model.CmafMuxing, error)
}

func int64Value(i *int64) int64 {
	if i == nil {
		return 0
	}

	return *i
}

func dimensionToInt64(i *int32) int64 {
	if i == nil {
		return 0
	}

	return int64(*i)
}
