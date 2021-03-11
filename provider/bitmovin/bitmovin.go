package bitmovin

import (
	"context"
	"fmt"
	"time"

	"strings"
	"sync"

	"github.com/cbsinteractive/transcode-orchestrator/config"
	"github.com/cbsinteractive/transcode-orchestrator/job"
	"github.com/cbsinteractive/transcode-orchestrator/provider"
	"github.com/cbsinteractive/transcode-orchestrator/provider/bitmovin/codec"
	"github.com/cbsinteractive/transcode-orchestrator/provider/bitmovin/storage"

	"github.com/pkg/errors"
	"github.com/zsiec/pkg/tracing"

	"github.com/bitmovin/bitmovin-api-sdk-go"
	"github.com/bitmovin/bitmovin-api-sdk-go/common"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/query"
)

func init() {
	_ = provider.Register(Name, factory)
}

type (
	Preset = codec.Preset
	Job    = job.Job
	Status = job.Status
)

const (
	// Name is the name used for registering the bitmovin driver in the
	// registry of drivers.
	Name = "bitmovin"

	codecVorbis = "vorbis"
	codecAAC    = "aac"
	codecOpus   = "opus"
	codecVP8    = "vp8"
	codecH264   = "h264"
	codecH265   = "h265"
	codecAV1    = "av1"

	containerWebM = "webm"
	containerMP4  = "mp4"
	containerMOV  = "mov"
)

func factory(cfg *config.Config) (provider.Provider, error) {
	if cfg.Bitmovin.APIKey == "" {
		return nil, fmt.Errorf("%w: no api key", provider.ErrConfig)
	}

	if _, ok := cloudRegions[model.CloudRegion(cfg.Bitmovin.EncodingRegion)]; !ok {
		return nil, fmt.Errorf("%w: cloud: bad encoding region", provider.ErrConfig)
	}

	if _, ok := awsCloudRegions[model.AwsCloudRegion(cfg.Bitmovin.AWSStorageRegion)]; !ok {
		return nil, fmt.Errorf("%w: cloud: bad storage region", provider.ErrConfig)
	}

	api, err := bitmovin.NewBitmovinApi(func(apiClient *common.ApiClient) {
		apiClient.ApiKey = cfg.Bitmovin.APIKey
		apiClient.BaseUrl = cfg.Bitmovin.Endpoint
	})
	if err != nil {
		return nil, err
	}

	tracer := cfg.Tracer
	if tracer == nil {
		tracer = tracing.NoopTracer{}
	}

	return &driver{
		api:    api,
		cfg:    cfg.Bitmovin,
		tracer: tracer,
	}, nil
}

type driver struct {
	api    *bitmovin.BitmovinApi
	cfg    *config.Bitmovin
	tracer tracing.Tracer
}

func (p *driver) Create(ctx context.Context, j *Job) (*Status, error) {
	presets := make([]Preset, len(j.Output.File))
	for i, f := range j.Output.File {
		if err := p.createPreset(ctx, f, &presets[i]); err != nil {
			return nil, fmt.Errorf("output[%d]: preset: %w", i, err)
		}
	}

	inputPath := j.Input.URL().Path
	inputID, err := p.inputFrom(ctx, j)
	if err != nil {
		return nil, err
	}

	destPath := p.path(*j)
	outputID, err := p.outputFrom(ctx, j)
	if err != nil {
		return nil, err
	}

	encCustomData := make(map[string]map[string]interface{})

	infrastructureSettings, encodingCloudRegion, err := p.encodingInfrastructureFrom(j)
	if err != nil {
		return nil, errors.Wrap(err, "validating and setting encoding infrastructure")
	}

	jobName := j.ID
	if name := j.Name; name != "" {
		jobName = name
	}

	subSeg := p.tracer.BeginSubsegment(ctx, "bitmovin-create-encoding")
	enc, err := p.api.Encoding.Encodings.Create(model.Encoding{
		Name:           jobName,
		CustomData:     &encCustomData,
		CloudRegion:    encodingCloudRegion,
		EncoderVersion: p.cfg.EncodingVersion,
		Infrastructure: infrastructureSettings,
		Labels:         j.Labels,
	})
	if err != nil {
		subSeg.Close(err)
		return nil, errors.Wrap(err, "creating encoding")
	}
	subSeg.Close(nil)

	subSeg = p.tracer.BeginSubsegment(ctx, "bitmovin-create-ingest")
	inputID, err = func(inputID string) (string, error) {
		istream, err := p.api.Encoding.Encodings.InputStreams.Ingest.Create(enc.Id, model.IngestInputStream{
			InputId:       inputID,
			InputPath:     inputPath,
			SelectionMode: model.StreamSelectionMode_AUTO,
		})
		if err != nil {
			return inputID, err
		}
		return istream.Id, err
	}(inputID)
	subSeg.Close(err)
	if err != nil {
		return nil, fmt.Errorf("ingest: %v", err)
	}

	inputID, err = p.splice(ctx, enc.Id, inputID, j.Input.Splice)
	if err != nil {
		return nil, fmt.Errorf("splice: %w", err)
	}

	var wg sync.WaitGroup
	errorc := make(chan error)

	subSeg = p.tracer.BeginSubsegment(ctx, "bitmovin-create-outputs")
	for idx, o := range j.Output.File {
		wg.Add(1)
		go p.createOutput(outputCfg{
			preset:         presets[idx],
			encodingID:     enc.Id,
			audioIn:        inputID,
			videoIn:        inputID,
			outputID:       outputID,
			outputFilename: o.Name,
			destPath:       destPath,
			job:            j,
		}, &wg, errorc)
	}

	go func() {
		wg.Wait()
		close(errorc)
	}()

	for err := range errorc {
		subSeg.Close(err)
		if err != nil {
			return nil, err
		}
	}
	subSeg.Close(nil)

	if o := j.Input.ExplicitKeyframeOffsets; len(o) > 0 {
		subSeg = p.tracer.BeginSubsegment(ctx, "bitmovin-create-keyframes")
		if err = p.createExplicitKeyframes(enc.Id, o); err != nil {
			subSeg.Close(err)
			return nil, fmt.Errorf("creating keyframes: %w", err)
		}
		subSeg.Close(nil)
	}

	subSeg = p.tracer.BeginSubsegment(ctx, "bitmovin-start-encoding")
	encResp, err := p.api.Encoding.Encodings.Start(enc.Id, model.StartEncodingRequest{})
	if err != nil {
		subSeg.Close(err)
		return nil, errors.Wrap(err, "starting encoding job")
	}
	subSeg.Close(nil)

	return &Status{
		Provider:      Name,
		ProviderJobID: encResp.Id,
		State:         job.StateQueued,
	}, nil
}

func (p *driver) Status(ctx context.Context, j *Job) (*Status, error) {
	subSeg := p.tracer.BeginSubsegment(ctx, "bitmovin-create-get-encoding-status")
	task, err := p.api.Encoding.Encodings.Status(j.ProviderJobID)
	if err != nil {
		subSeg.Close(err)
		return nil, errors.Wrap(err, "retrieving encoding status")
	}
	subSeg.Close(nil)

	var progress float64
	if task.Progress != nil {
		progress = float64(*task.Progress)
	}

	s := Status{
		Provider:      Name,
		ProviderJobID: j.ProviderJobID,
		State:         state(task.Status),
		Progress:      progress,
		ProviderStatus: map[string]interface{}{
			"messages":       task.Messages,
			"originalStatus": task.Status,
		},
		Output: job.Dir{
			Path: j.Location(""),
		},
	}

	if s.State == job.StateFinished {
		subSeg := p.tracer.BeginSubsegment(ctx, "bitmovin-get-output-info")
		s, err = p.enrichStreams(s)
		if err != nil {
			subSeg.Close(err)
			return nil, errors.Wrap(err, "enriching status with source info")
		}

		// TODO: it would be better to know which containers to include in this fetch
		// rather than iterating over all supported containers
		for _, c := range containers {
			s, err = c.Enrich(p.api, s)
			if err != nil {
				subSeg.Close(err)
				return nil, err
			}
		}
		subSeg.Close(nil)
	}

	return &s, nil
}

func (p *driver) Cancel(ctx context.Context, id string) (err error) {
	defer p.trace(ctx, "bitmovin-delete-job", &err)()
	_, err = p.api.Encoding.Encodings.Stop(id)
	return err
}

type outputCfg struct {
	preset             Preset
	encodingID         string
	videoIn, audioIn   string
	outputID           string
	destPath           string
	outputFilename     string
	manifestID         string
	manifestMasterPath string
	job                *Job
}

func (p *driver) createOutput(cfg outputCfg, wg *sync.WaitGroup, errorc chan error) {
	defer wg.Done()
	var audioMuxingStream, videoMuxingStream model.MuxingStream

	if audCfgID := cfg.preset.AudioConfigID; audCfgID != "" {
		audStream, err := p.api.Encoding.Encodings.Streams.Create(cfg.encodingID, model.Stream{
			CodecConfigId: audCfgID,
			InputStreams:  []model.StreamInput{{InputStreamId: cfg.audioIn}},
		})
		if err != nil {
			errorc <- errors.Wrap(err, "adding audio stream to the encoding")
			return
		}

		audioMuxingStream = model.MuxingStream{StreamId: audStream.Id}
	}

	if vidCfgID := cfg.preset.VideoConfigID; vidCfgID != "" {
		vidStream, err := p.api.Encoding.Encodings.Streams.Create(cfg.encodingID, model.Stream{
			CodecConfigId: vidCfgID,
			InputStreams:  []model.StreamInput{{InputStreamId: cfg.videoIn}},
		})
		if err != nil {
			errorc <- errors.Wrap(err, "adding video stream to the encoding")
			return
		}

		if videoFilters := cfg.preset.VideoFilters; videoFilters != nil {
			for i, filter := range videoFilters {
				_, err = p.api.Encoding.Encodings.Streams.Filters.Create(cfg.encodingID, vidStream.Id, []model.StreamFilter{
					{Id: filter, Position: bitmovin.Int32Ptr(int32(i))},
				})
				if err != nil {
					errorc <- fmt.Errorf("adding filter %s to video stream: %w", filter, err)
					return
				}
			}
		}

		videoMuxingStream = model.MuxingStream{StreamId: vidStream.Id}
	}

	container := containers[strings.ToLower(cfg.preset.Container)]
	if container == nil {
		errorc <- fmt.Errorf("unknown container format %q", cfg.preset.Container)
		return
	}

	if err := container.Assemble(p.api, AssemblerCfg{
		EncID:              cfg.encodingID,
		OutputID:           cfg.outputID,
		DestPath:           cfg.destPath,
		OutputFilename:     cfg.outputFilename,
		AudCfgID:           cfg.preset.AudioConfigID,
		VidCfgID:           cfg.preset.VideoConfigID,
		AudMuxingStream:    audioMuxingStream,
		VidMuxingStream:    videoMuxingStream,
		ManifestID:         cfg.manifestID,
		ManifestMasterPath: cfg.manifestMasterPath,
	}); err != nil {
		errorc <- err
		return
	}
}

func (p *driver) inputFrom(ctx context.Context, job *Job) (inputID string, err error) {
	defer p.trace(ctx, "bitmovin-create-input", &err)()

	if alias := job.Env.InputAlias; alias != "" {
		return alias, nil
	}
	return storage.NewInput(job.Input.Name, storage.InputAPI{
		S3:    p.api.Encoding.Inputs.S3,
		GCS:   p.api.Encoding.Inputs.Gcs,
		HTTP:  p.api.Encoding.Inputs.Http,
		HTTPS: p.api.Encoding.Inputs.Https,
	}, p.cfg)
}

func (p *driver) outputFrom(ctx context.Context, job *Job) (id string, err error) {
	defer p.trace(ctx, "bitmovin-create-output", &err)()

	if alias := job.Env.OutputAlias; alias != "" {
		return alias, nil
	}

	return storage.NewOutput(p.path(*job), storage.OutputAPI{
		S3:  p.api.Encoding.Outputs.S3,
		GCS: p.api.Encoding.Outputs.Gcs,
	}, p.cfg)
}

func (p *driver) encodingCloudRegionFrom(j *Job) (model.CloudRegion, error) {
	if cloud, region := j.Env.Cloud, j.Env.Region; cloud+region != "" {
		regions, found := regionByCloud[cloud]
		if !found {
			return "", fmt.Errorf("unsupported cloud %q", cloud)
		}

		bitmovinRegion, found := regions[region]
		if !found {
			return "", fmt.Errorf("region %q is not supported with cloud %q", region, cloud)
		}

		return bitmovinRegion, nil
	}

	return model.CloudRegion(p.cfg.EncodingRegion), nil
}

func (p *driver) encodingInfrastructureFrom(j *Job) (*model.InfrastructureSettings, model.CloudRegion, error) {
	encodingCloudRegion, err := p.encodingCloudRegionFrom(j)
	if err != nil {
		return nil, encodingCloudRegion, errors.Wrap(err, "validating and setting encoding cloud region")
	}

	if tag, found := j.Env.Tags[job.TagTranscodeDefault]; found {
		return &model.InfrastructureSettings{
			InfrastructureId: tag,
			CloudRegion:      encodingCloudRegion,
		}, model.CloudRegion_EXTERNAL, nil
	}

	return nil, encodingCloudRegion, nil
}

func (p *driver) createExplicitKeyframes(encodingID string, offsets []float64) error {
	if len(offsets) == 0 {
		return nil
	}

	type work struct {
		offset float64
		err    error
	}

	workc := make(chan work, len(offsets))

	for _, o := range offsets {
		w := work{offset: o}
		go func() {
			_, w.err = p.api.Encoding.Encodings.Keyframes.Create(encodingID, model.Keyframe{Time: &w.offset})
			workc <- w
		}()
	}

	for i := 0; i < cap(workc); i++ {
		w := <-workc
		if w.err != nil {
			return w.err
		}
	}

	return nil
}

func (p *driver) createPreset(_ context.Context, f job.File, summary *Preset) error {
	vc, _ := codec.New(f.Video.Codec, f)
	ac, _ := codec.New(f.Audio.Codec, f)
	c := []codec.Codec{}
	if f.Video.Codec != "" {
		c = append(c, vc)
	}
	if f.Audio.Codec != "" {
		c = append(c, ac)
	}
	for _, c := range c {
		c.Create(p.api)
		if c.Err() != nil {
			return c.Err()
		}
		*summary = codec.Summary(c, f, *summary)
	}

	if summary.HasVideo() {
		deInterlace, err := p.api.Encoding.Filters.Deinterlace.Create(model.DeinterlaceFilter{
			Name:       "deinterlace",
			AutoEnable: model.DeinterlaceAutoEnable_META_DATA_AND_CONTENT_BASED,
		})
		if err != nil {
			return fmt.Errorf("creating deinterlace filter: %w", err)
		}

		summary.VideoFilters = append(summary.VideoFilters, deInterlace.Id)

		if c := f.Video.Crop; !c.Empty() {
			f, err := p.api.Encoding.Filters.Crop.Create(model.CropFilter{
				Left:   bitmovin.Int32Ptr(int32(c.Left)),
				Right:  bitmovin.Int32Ptr(int32(c.Right)),
				Top:    bitmovin.Int32Ptr(int32(c.Top)),
				Bottom: bitmovin.Int32Ptr(int32(c.Bottom)),
			})
			if err != nil {
				return fmt.Errorf("creating crop filter: %w", err)
			}

			summary.VideoFilters = append(summary.VideoFilters, f.Id)
		}

		for _, img := range f.Video.Overlays.Images {
			watermark, err := p.api.Encoding.Filters.Watermark.Create(model.WatermarkFilter{
				Name:  "imageOverlay",
				Right: bitmovin.Int32Ptr(0),
				Top:   bitmovin.Int32Ptr(0),
				Unit:  model.PositionUnit_PERCENTS,
				Image: img.URL,
			})
			if err != nil {
				return fmt.Errorf("creating watermark filter: %w", err)
			}
			summary.VideoFilters = append(summary.VideoFilters, watermark.Id)
		}
	}

	return nil
}

func (p *driver) enrichStreams(s Status) (Status, error) {
	inStreams, err := p.api.Encoding.Encodings.Streams.List(s.ProviderJobID, func(params *query.StreamListQueryParams) {
		params.Limit = 1
		params.Offset = 0
	})
	if err != nil {
		return s, errors.Wrap(err, "retrieving input streams from the Bitmovin API")
	}
	if len(inStreams.Items) == 0 {
		return s, fmt.Errorf("no streams found for encodingID %s", s.ProviderJobID)
	}

	inStream := inStreams.Items[0]
	streamInput, err := p.api.Encoding.Encodings.Streams.Input.Get(s.ProviderJobID, inStream.Id)
	if err != nil {
		return s, errors.Wrap(err, "retrieving stream input details from the Bitmovin API")
	}

	var (
		vidCodec      string
		width, height int
	)
	if len(streamInput.VideoStreams) > 0 {
		vidStreamInput := streamInput.VideoStreams[0]
		vidCodec = vidStreamInput.Codec
		width = int(int32Value(vidStreamInput.Width))
		height = int(int32Value(vidStreamInput.Height))
	}

	s.Input = job.File{
		Duration: time.Duration(floatValue(streamInput.Duration) * float64(time.Second)),
		Video:    job.Video{Width: width, Height: height, Codec: vidCodec},
	}

	return s, nil

}

func (p *driver) path(j Job) string {
	if j.Output.Path == "" {
		j.Output.Path = p.cfg.Destination
	}
	return j.Location("")
}

// Healthcheck returns an error if a call to List Encodings with a limit of one
// returns an error
func (p *driver) Healthcheck() error {
	_, err := p.api.Encoding.Encodings.List(func(params *query.EncodingListQueryParams) {
		params.Limit = 1
	})
	if err != nil {
		return errors.Wrap(err, "bitmovin service unavailable")
	}

	return nil
}

// Capabilities describes the capabilities of the driver.
func (p *driver) Capabilities() provider.Capabilities {
	return provider.Capabilities{
		InputFormats:  []string{"prores", "h264"},
		OutputFormats: []string{containerMP4, containerMOV, containerWebM},
		Destinations:  []string{"s3", "gcs"},
	}
}

func floatValue(f *float64) float64 {
	if f == nil {
		return 0
	}
	return *f
}

func int32Value(i *int32) int32 {
	if i == nil {
		return 0
	}
	return *i
}
