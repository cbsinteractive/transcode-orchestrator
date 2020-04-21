package container

import (
	"fmt"
	"path"

	"github.com/bitmovin/bitmovin-api-sdk-go"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/cbsinteractive/transcode-orchestrator/provider"
	"github.com/cbsinteractive/transcode-orchestrator/provider/bitmovin/internal/storage"
	"github.com/cbsinteractive/transcode-orchestrator/provider/bitmovin/internal/types"
	"github.com/pkg/errors"
)

const (
	// CustomDataKeyManifest is used as the base key to store the manifestID in an encoding
	CustomDataKeyManifest = "manifest"
	// CustomDataKeyManifestID is the key used to store the manifestID in an encoding
	CustomDataKeyManifestID = "id"
)

// HLSAssembler is an assembler that creates HLS outputs based on a cfg
type HLSAssembler struct {
	api HLSContainerAPI
}

// NewHLSAssembler creates and returns an HLSAssembler
func NewHLSAssembler(api HLSContainerAPI) *HLSAssembler {
	return &HLSAssembler{api: api}
}

// Assemble creates HLS outputs
func (a *HLSAssembler) Assemble(cfg AssemblerCfg) error {
	if cfg.AudCfgID != "" {
		audTSMuxing, err := a.api.TSMuxing.Create(cfg.EncID, model.TsMuxing{
			SegmentLength: floatToPtr(float64(cfg.SegDuration)),
			SegmentNaming: "seg_%number%.ts",
			Streams:       []model.MuxingStream{cfg.AudMuxingStream},
			Outputs: []model.EncodingOutput{
				storage.EncodingOutputFrom(cfg.OutputID, path.Join(cfg.ManifestMasterPath, cfg.AudCfgID)),
			},
		})
		if err != nil {
			return errors.Wrap(err, "creating audio ts muxing")
		}

		_, err = a.api.HLSAudioMedia.Create(cfg.ManifestID, model.AudioMediaInfo{
			Uri:             cfg.AudCfgID + ".m3u8",
			GroupId:         cfg.AudCfgID,
			Language:        "en",
			Name:            cfg.AudCfgID,
			IsDefault:       boolToPtr(false),
			Autoselect:      boolToPtr(false),
			Forced:          boolToPtr(false),
			SegmentPath:     cfg.AudCfgID,
			Characteristics: []string{"public.accessibility.describes-video"},
			EncodingId:      cfg.EncID,
			StreamId:        cfg.AudMuxingStream.StreamId,
			MuxingId:        audTSMuxing.Id,
		})
		if err != nil {
			return errors.Wrap(err, "creating audio media")
		}
	}

	if cfg.VidCfgID != "" {
		vidTSMuxing, err := a.api.TSMuxing.Create(cfg.EncID, model.TsMuxing{
			SegmentLength: floatToPtr(float64(cfg.SegDuration)),
			SegmentNaming: "seg_%number%.ts",
			Streams:       []model.MuxingStream{cfg.VidMuxingStream},
			Outputs: []model.EncodingOutput{
				storage.EncodingOutputFrom(cfg.OutputID, path.Join(cfg.ManifestMasterPath, cfg.VidCfgID)),
			},
		})
		if err != nil {
			return errors.Wrap(err, "creating video ts muxing")
		}

		_, err = a.api.HLSStreams.Create(cfg.ManifestID, model.StreamInfo{
			Audio:       cfg.AudCfgID,
			Uri:         fmt.Sprintf("%s.m3u8", cfg.VidCfgID),
			SegmentPath: cfg.VidCfgID,
			EncodingId:  cfg.EncID,
			StreamId:    cfg.VidMuxingStream.StreamId,
			MuxingId:    vidTSMuxing.Id,
		})
		if err != nil {
			return errors.Wrap(err, "creating video stream info")
		}
	}

	return nil
}

// HLSStatusEnricher is responsible for adding output HLS info to a job status
type HLSStatusEnricher struct {
	api *bitmovin.BitmovinApi
}

// NewHLSStatusEnricher creates and returns an HLSStatusEnricher
func NewHLSStatusEnricher(api *bitmovin.BitmovinApi) *HLSStatusEnricher {
	return &HLSStatusEnricher{api: api}
}

// Enrich populates information about the HLS output if it exists
func (e *HLSStatusEnricher) Enrich(s provider.JobStatus) (provider.JobStatus, error) {
	data, err := e.api.Encoding.Encodings.Customdata.Get(s.ProviderJobID)
	if err != nil {
		return s, errors.Wrap(err, "retrieving the encoding from the Bitmovin API")
	}

	manifestID, err := types.CustomDataStringValAtKeys(data.CustomData, CustomDataKeyManifest, CustomDataKeyManifestID)
	if err == nil && manifestID != "" {
		s.ProviderStatus["manifestStatus"] = model.Status_FINISHED
	}

	return s, nil
}

func floatToPtr(f float64) *float64 {
	return &f
}

func boolToPtr(b bool) *bool {
	return &b
}
