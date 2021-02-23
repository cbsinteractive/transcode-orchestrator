package bitmovin

import (
	"path"

	"github.com/bitmovin/bitmovin-api-sdk-go"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/query"
	"github.com/bitmovin/bitmovin-api-sdk-go/serialization"
	"github.com/cbsinteractive/transcode-orchestrator/job"
	"github.com/cbsinteractive/transcode-orchestrator/provider/bitmovin/storage"
)

var containers = map[string]interface {
	Assemble(*bitmovin.BitmovinApi, AssemblerCfg) error
	Enrich(*bitmovin.BitmovinApi, job.Status) (job.Status, error)
}{
	"webm": &WEBM{},
	"mp4":  &MP4{},
	"mov":  &MOV{},
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

type MOV struct{}
type MP4 struct{}
type WEBM struct{}

func (a AssemblerCfg) Streams() (s []model.MuxingStream) {
	empty := model.MuxingStream{}
	if a.VidMuxingStream != empty {
		s = append(s, a.VidMuxingStream)
	}
	if a.AudMuxingStream != empty {
		s = append(s, a.AudMuxingStream)
	}
	return s
}
func (a AssemblerCfg) Filename() string { return path.Base(a.OutputFilename) }
func (a AssemblerCfg) Outputs() []model.EncodingOutput {
	return []model.EncodingOutput{
		storage.EncodingOutputFrom(a.OutputID, path.Dir(path.Join(a.DestPath, a.OutputFilename))),
	}
}

func (a *MOV) Assemble(api *bitmovin.BitmovinApi, cfg AssemblerCfg) error {
	_, err := api.Encoding.Encodings.Muxings.ProgressiveMov.Create(cfg.EncID, model.ProgressiveMovMuxing{
		Filename:             cfg.Filename(),
		Streams:              cfg.Streams(),
		StreamConditionsMode: model.StreamConditionsMode_DROP_STREAM,
		Outputs:              cfg.Outputs(),
	})
	return err
}
func (a *MP4) Assemble(api *bitmovin.BitmovinApi, cfg AssemblerCfg) error {
	_, err := api.Encoding.Encodings.Muxings.Mp4.Create(cfg.EncID, model.Mp4Muxing{
		Filename:             cfg.Filename(),
		Streams:              cfg.Streams(),
		StreamConditionsMode: model.StreamConditionsMode_DROP_STREAM,
		Outputs:              cfg.Outputs(),
	})
	return err
}
func (a *WEBM) Assemble(api *bitmovin.BitmovinApi, cfg AssemblerCfg) error {
	_, err := api.Encoding.Encodings.Muxings.ProgressiveWebm.Create(cfg.EncID, model.ProgressiveWebmMuxing{
		Filename:             cfg.Filename(),
		Streams:              cfg.Streams(),
		StreamConditionsMode: model.StreamConditionsMode_DROP_STREAM,
		Outputs:              cfg.Outputs(),
	})
	return err
}

// Enrich populates information about MOV outputs if they exist
func (e *MOV) Enrich(api *bitmovin.BitmovinApi, s job.Status) (job.Status, error) {
	get := api.Encoding.Encodings.Muxings.ProgressiveMov.Information.Get
	mux, err := ListMuxing(api, s.ProviderJobID)
	if err != nil {
		return s, err
	}
	for _, mux := range mux {
		info, err := get(s.ProviderJobID, mux.Id)
		if err != nil {
			return s, nil
		}
		s.Output.Files = append(s.Output.Files, mux.Output(s, infoMOV(*info)))
	}
	return s, nil
}

// Enrich populates information about MP4 outputs if they exist
func (e *MP4) Enrich(api *bitmovin.BitmovinApi, s job.Status) (job.Status, error) {
	get := api.Encoding.Encodings.Muxings.Mp4.Information.Get
	mux, err := ListMuxing(api, s.ProviderJobID)
	if err != nil {
		return s, err
	}
	for _, mux := range mux {
		info, err := get(s.ProviderJobID, mux.Id)
		if err != nil {
			return s, nil
		}
		s.Output.Files = append(s.Output.Files, mux.Output(s, infoMP4(*info)))
	}
	return s, nil
}

// Enrich populates information about ProgressiveWebM outputs if they exist
func (e *WEBM) Enrich(api *bitmovin.BitmovinApi, s job.Status) (job.Status, error) {
	get := api.Encoding.Encodings.Muxings.ProgressiveWebm.Information.Get
	mux, err := ListMuxing(api, s.ProviderJobID)
	if err != nil {
		return s, err
	}
	for _, mux := range mux {
		info, err := get(s.ProviderJobID, mux.Id)
		if err != nil {
			return s, nil
		}
		s.Output.Files = append(s.Output.Files, mux.Output(s, infoWEBM(*info)))
	}
	return s, nil
}

type Muxing struct {
	Id       string `json:"id,omitempty"`
	Filename string `json:"filename,omitempty"`
	Type     string `json:"type,omitempty"`
}

func ListMuxing(api *bitmovin.BitmovinApi, jobID string) ([]Muxing, error) {
	list := api.Encoding.Encodings.Muxings.List

	total := 1
	var items []model.Muxing
	for len(items) < total {
		resp, err := list(jobID, func(params *query.MuxingListQueryParams) {
			params.Offset = int32(len(items))
			params.Limit = 100
		})
		if err != nil {
			return nil, err
		}
		if n := resp.TotalCount; n != nil {
			total = int(*n)
		}
		items = append(items, resp.Items...)
	}

	mux := make([]Muxing, len(items))
	for i := range items {
		serialization.Decode(items[i], &mux[i])
	}
	return mux, nil
}

func (m Muxing) Output(s job.Status, t track) job.File {
	var (
		height, width int64
		codec         string
	)
	video := t.Video()
	if len(video) > 0 {
		height, width = int64(*video[0].FrameHeight), int64(*video[0].FrameWidth)
		codec = video[0].Codec
	}
	return job.File{
		Path:       s.Output.Destination + m.Filename,
		Container:  m.Type,
		Size:       t.Filesize(),
		VideoCodec: codec,
		Width:      width,
		Height:     height,
	}
}

type track interface {
	Filesize() int64
	Video() []model.MuxingInformationVideoTrack
}

type infoWEBM model.ProgressiveWebmMuxingInformation
type infoMP4 model.Mp4MuxingInformation
type infoMOV model.ProgressiveMovMuxingInformation

func (i infoWEBM) Filesize() int64 { return *i.FileSize }
func (i infoMP4) Filesize() int64  { return *i.FileSize }
func (i infoMOV) Filesize() int64  { return *i.FileSize }

func (i infoWEBM) Video() []model.MuxingInformationVideoTrack { return i.VideoTracks }
func (i infoMP4) Video() []model.MuxingInformationVideoTrack  { return i.VideoTracks }
func (i infoMOV) Video() []model.MuxingInformationVideoTrack  { return i.VideoTracks }
