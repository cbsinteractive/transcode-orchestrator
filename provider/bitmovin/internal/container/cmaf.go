package container

import (
	"fmt"
	"path"
	"path/filepath"

	"github.com/bitmovin/bitmovin-api-sdk-go"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/cbsinteractive/transcode-orchestrator/provider"
	"github.com/cbsinteractive/transcode-orchestrator/provider/bitmovin/internal/storage"
	"github.com/cbsinteractive/transcode-orchestrator/provider/bitmovin/internal/types"
	"github.com/pkg/errors"
)

// CMAFAssembler is an assembler that creates HLS outputs based on a cfg
type CMAFAssembler struct {
	api CMAFContainerAPI
}

// NewCMAFAssembler creates and returns an CMAFAssembler
func NewCMAFAssembler(api CMAFContainerAPI) *CMAFAssembler {
	return &CMAFAssembler{api: api}
}

// Assemble creates HLS outputs
func (a *CMAFAssembler) Assemble(cfg AssemblerCfg) error {
	if cfg.AudCfgID != "" {
		audCMAFMuxing, err := a.api.CMAFMuxing.Create(cfg.EncID, model.CmafMuxing{
			SegmentLength: floatToPtr(float64(cfg.SegDuration)),
			SegmentNaming: "seg_%number%.m4a",
			Streams:       []model.MuxingStream{cfg.AudMuxingStream},
			Outputs: []model.EncodingOutput{
				storage.EncodingOutputFrom(cfg.OutputID, path.Join(cfg.ManifestMasterPath, cfg.AudCfgID)),
			},
		})
		if err != nil {
			return errors.Wrap(err, "creating audio cmaf muxing")
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
			MuxingId:        audCMAFMuxing.Id,
		})
		if err != nil {
			return errors.Wrap(err, "creating audio media")
		}
	}

	if cfg.VidCfgID != "" {
		vidCMAFMuxing, err := a.api.CMAFMuxing.Create(cfg.EncID, model.CmafMuxing{
			SegmentLength: floatToPtr(float64(cfg.SegDuration)),
			SegmentNaming: "seg_%number%.m4v",
			Streams:       []model.MuxingStream{cfg.VidMuxingStream},
			Outputs: []model.EncodingOutput{
				storage.EncodingOutputFrom(cfg.OutputID, path.Join(cfg.ManifestMasterPath, cfg.VidCfgID)),
			},
		})
		if err != nil {
			return errors.Wrap(err, "creating video cmaf muxing")
		}

		vidSegLoc, err := filepath.Rel(path.Dir(path.Join(cfg.DestPath, cfg.OutputFilename)), path.Join(cfg.ManifestMasterPath, cfg.VidCfgID))
		if err != nil {
			return errors.Wrap(err, "constructing video segment location")
		}

		_, err = a.api.HLSStreams.Create(cfg.ManifestID, model.StreamInfo{
			Audio:       cfg.AudCfgID,
			Uri:         fmt.Sprintf("%s.m3u8", cfg.VidCfgID),
			SegmentPath: vidSegLoc,
			EncodingId:  cfg.EncID,
			StreamId:    cfg.VidMuxingStream.StreamId,
			MuxingId:    vidCMAFMuxing.Id,
		})
		if err != nil {
			return errors.Wrap(err, "creating video stream info")
		}
	}

	return nil
}

// CMAFStatusEnricher is responsible for adding output HLS info to a job status
type CMAFStatusEnricher struct {
	api *bitmovin.BitmovinApi
}

// NewHLSStatusEnricher creates and returns an CMAFStatusEnricher
func NewCMAFStatusEnricher(api *bitmovin.BitmovinApi) *CMAFStatusEnricher {
	return &CMAFStatusEnricher{api: api}
}

// Enrich populates information about the CMAF output if it exists
func (e *CMAFStatusEnricher) Enrich(s provider.JobStatus) (provider.JobStatus, error) {
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
