package container

import "github.com/bitmovin/bitmovin-api-sdk-go/model"

// Assembler is responsible for creating all resources for a given container output
type Assembler interface {
	Assemble(AssemblerCfg) error
}

// AssemblerCfg holds properties any individual assembler might need when creating resources
type AssemblerCfg struct {
	EncID                            string
	OutputID                         string
	DestPath                         string
	OutputFilename                   string
	AudCfgID, VidCfgID               string
	AudMuxingStream, VidMuxingStream model.MuxingStream
	ManifestID                       string
	ManifestMasterPath               string
	SegDuration                      uint
}

func streamsFrom(cfg AssemblerCfg) []model.MuxingStream {
	streams := []model.MuxingStream{}

	if !(cfg.VidMuxingStream == model.MuxingStream{}) {
		streams = append(streams, cfg.VidMuxingStream)
	}

	if !(cfg.AudMuxingStream == model.MuxingStream{}) {
		streams = append(streams, cfg.AudMuxingStream)
	}

	return streams
}
