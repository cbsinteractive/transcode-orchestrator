package bitmovin

import (
	"context"
	"fmt"
	"time"

	"net/url"
	"path"
	"strings"
	"sync"

	"github.com/cbsinteractive/transcode-orchestrator/config"
	"github.com/cbsinteractive/transcode-orchestrator/db"
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
	_ = provider.Register(Name, bitmovinFactory)
}

const (
	// Name is the name used for registering the bitmovin provider in the
	// registry of providers.
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

var errBitmovinInvalidConfig = provider.InvalidConfigError("Invalid configuration")

func bitmovinFactory(cfg *config.Config) (provider.Provider, error) {
	if cfg.Bitmovin.APIKey == "" {
		return nil, errBitmovinInvalidConfig
	}

	if _, ok := cloudRegions[model.CloudRegion(cfg.Bitmovin.EncodingRegion)]; !ok {
		return nil, errBitmovinInvalidConfig
	}

	if _, ok := awsCloudRegions[model.AwsCloudRegion(cfg.Bitmovin.AWSStorageRegion)]; !ok {
		return nil, errBitmovinInvalidConfig
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

	return &bitmovinProvider{
		api:         api,
		providerCfg: cfg.Bitmovin,
		tracer:      tracer,
	}, nil
}

type bitmovinProvider struct {
	api         *bitmovin.BitmovinApi
	providerCfg *config.Bitmovin
	tracer      tracing.Tracer
}

func (p *bitmovinProvider) Transcode(ctx context.Context, job *db.Job) (*provider.JobStatus, error) {
	presets := make([]db.PresetSummary, len(job.Outputs))
	for i, output := range job.Outputs {
		if err := p.createPreset(ctx, output.Preset, &presets[i]); err != nil {
			return nil, fmt.Errorf("output[%d]: preset: %w", i, err)
		}
	}

	inputID, mediaPath, err := p.inputFrom(ctx, job)
	if err != nil {
		return nil, err
	}

	outputID, destPath, err := p.outputFrom(ctx, job)
	if err != nil {
		return nil, err
	}

	encCustomData := make(map[string]map[string]interface{})

	infrastructureSettings, encodingCloudRegion, err := p.encodingInfrastructureFrom(job)
	if err != nil {
		return nil, errors.Wrap(err, "validating and setting encoding infrastructure")
	}

	jobName := job.ID
	if name := job.Name; name != "" {
		jobName = name
	}

	subSeg := p.tracer.BeginSubsegment(ctx, "bitmovin-create-encoding")
	enc, err := p.api.Encoding.Encodings.Create(model.Encoding{
		Name:           jobName,
		CustomData:     &encCustomData,
		CloudRegion:    encodingCloudRegion,
		EncoderVersion: p.providerCfg.EncodingVersion,
		Infrastructure: infrastructureSettings,
		Labels:         job.Labels,
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
			InputPath:     mediaPath,
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

	inputID, err = p.splice(ctx, enc.Id, inputID, job.SourceSplice)
	if err != nil {
		return nil, fmt.Errorf("splice: %w", err)
	}

	var wg sync.WaitGroup
	errorc := make(chan error)

	subSeg = p.tracer.BeginSubsegment(ctx, "bitmovin-create-outputs")
	for idx, o := range job.Outputs {
		wg.Add(1)
		go p.createOutput(outputCfg{
			preset:         presets[idx],
			encodingID:     enc.Id,
			audioIn:        inputID,
			videoIn:        inputID,
			outputID:       outputID,
			outputFilename: o.FileName,
			destPath:       destPath,
			job:            job,
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

	if o := job.ExplicitKeyframeOffsets; len(o) > 0 {
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

	return &provider.JobStatus{
		ProviderName:  Name,
		ProviderJobID: encResp.Id,
		Status:        provider.StatusQueued,
	}, nil
}

func (p *bitmovinProvider) JobStatus(ctx context.Context, job *db.Job) (*provider.JobStatus, error) {
	subSeg := p.tracer.BeginSubsegment(ctx, "bitmovin-create-get-encoding-status")
	task, err := p.api.Encoding.Encodings.Status(job.ProviderJobID)
	if err != nil {
		subSeg.Close(err)
		return nil, errors.Wrap(err, "retrieving encoding status")
	}
	subSeg.Close(nil)

	var progress float64
	if task.Progress != nil {
		progress = float64(*task.Progress)
	}

	s := provider.JobStatus{
		ProviderName:  Name,
		ProviderJobID: job.ProviderJobID,
		Status:        status(task.Status),
		Progress:      progress,
		ProviderStatus: map[string]interface{}{
			"messages":       task.Messages,
			"originalStatus": task.Status,
		},
		Output: provider.JobOutput{
			Destination: strings.TrimRight(p.destinationForJob(job), "/") + "/" + job.RootFolder() + "/",
		},
	}

	if s.Status == provider.StatusFinished {
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

func (p *bitmovinProvider) CancelJob(ctx context.Context, id string) (err error) {
	defer p.trace(ctx, "bitmovin-delete-job", &err)()
	_, err = p.api.Encoding.Encodings.Stop(id)
	return err
}

type outputCfg struct {
	preset             db.PresetSummary
	encodingID         string
	videoIn, audioIn   string
	outputID           string
	destPath           string
	outputFilename     string
	manifestID         string
	manifestMasterPath string
	job                *db.Job
}

func (p *bitmovinProvider) createOutput(cfg outputCfg, wg *sync.WaitGroup, errorc chan error) {
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
		SegDuration:        cfg.job.StreamingParams.SegmentDuration,
	}); err != nil {
		errorc <- err
		return
	}
}

func (p *bitmovinProvider) inputFrom(ctx context.Context, job *db.Job) (inputID string, srcPath string, err error) {
	defer p.trace(ctx, "bitmovin-create-input", &err)()

	srcPath, err = storage.PathFrom(job.SourceMedia)
	if err != nil {
		return "", "", err
	}

	if alias := job.ExecutionEnv.InputAlias; alias != "" {
		return alias, srcPath, nil
	}

	inputID, err = storage.NewInput(job.SourceMedia, storage.InputAPI{
		S3:    p.api.Encoding.Inputs.S3,
		GCS:   p.api.Encoding.Inputs.Gcs,
		HTTP:  p.api.Encoding.Inputs.Http,
		HTTPS: p.api.Encoding.Inputs.Https,
	}, p.providerCfg)
	if err != nil {
		return "", srcPath, err
	}

	return inputID, srcPath, nil
}

func (p *bitmovinProvider) outputFrom(ctx context.Context, job *db.Job) (inputID string, destPath string, err error) {
	defer p.trace(ctx, "bitmovin-create-output", &err)()

	destBasePath := p.destinationForJob(job)
	destURL, err := url.Parse(destBasePath)
	if err != nil {
		return "", "", err
	}

	destPath = path.Join(destURL.Path, job.RootFolder())

	if alias := job.ExecutionEnv.OutputAlias; alias != "" {
		return alias, destPath, nil
	}

	outputID, err := storage.NewOutput(destBasePath, storage.OutputAPI{
		S3:  p.api.Encoding.Outputs.S3,
		GCS: p.api.Encoding.Outputs.Gcs,
	}, p.providerCfg)
	if err != nil {
		return "", destPath, err
	}

	return outputID, destPath, nil
}

func (p *bitmovinProvider) encodingCloudRegionFrom(job *db.Job) (model.CloudRegion, error) {
	if cloud, region := job.ExecutionEnv.Cloud, job.ExecutionEnv.Region; cloud+region != "" {
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

	return model.CloudRegion(p.providerCfg.EncodingRegion), nil
}

func (p *bitmovinProvider) encodingInfrastructureFrom(job *db.Job) (*model.InfrastructureSettings, model.CloudRegion, error) {
	encodingCloudRegion, err := p.encodingCloudRegionFrom(job)
	if err != nil {
		return nil, encodingCloudRegion, errors.Wrap(err, "validating and setting encoding cloud region")
	}

	if tag, found := job.ExecutionEnv.ComputeTags[db.ComputeClassTranscodeDefault]; found {
		return &model.InfrastructureSettings{
			InfrastructureId: tag,
			CloudRegion:      encodingCloudRegion,
		}, model.CloudRegion_EXTERNAL, nil
	}

	return nil, encodingCloudRegion, nil
}

func (p *bitmovinProvider) createExplicitKeyframes(encodingID string, offsets []float64) error {
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

func (p *bitmovinProvider) createPreset(_ context.Context, preset db.Preset, summary *db.PresetSummary) error {
	vc, _ := codec.New(preset.Video.Codec, preset)
	ac, _ := codec.New(preset.Audio.Codec, preset)
	c := []codec.Codec{}
	if preset.Video.Codec != "" {
		c = append(c, vc)
	}
	if preset.Audio.Codec != "" {
		c = append(c, ac)
	}
	for _, c := range c {
		c.Create(p.api)
		if c.Err() != nil {
			return c.Err()
		}
		*summary = codec.Summary(c, preset, *summary)
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

		if c := preset.Video.Crop; !c.Empty() {
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

		if overlays := preset.Video.Overlays; overlays != nil && overlays.Images != nil {
			for _, image := range overlays.Images {
				watermark, err := p.api.Encoding.Filters.Watermark.Create(model.WatermarkFilter{
					Name:  "imageOverlay",
					Right: bitmovin.Int32Ptr(0),
					Top:   bitmovin.Int32Ptr(0),
					Unit:  model.PositionUnit_PERCENTS,
					Image: image.URL,
				})
				if err != nil {
					return fmt.Errorf("creating watermark filter: %w", err)
				}

				summary.VideoFilters = append(summary.VideoFilters, watermark.Id)
			}
		}
	}

	return nil
}

func (p *bitmovinProvider) enrichStreams(s provider.JobStatus) (provider.JobStatus, error) {
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
		width, height int64
	)
	if len(streamInput.VideoStreams) > 0 {
		vidStreamInput := streamInput.VideoStreams[0]
		vidCodec = vidStreamInput.Codec
		width = int64(int32Value(vidStreamInput.Width))
		height = int64(int32Value(vidStreamInput.Height))
	}

	s.SourceInfo = provider.SourceInfo{
		Duration:   time.Duration(floatValue(streamInput.Duration) * float64(time.Second)),
		Width:      width,
		Height:     height,
		VideoCodec: vidCodec,
	}

	return s, nil

}

func (p *bitmovinProvider) destinationForJob(job *db.Job) string {
	if path := job.DestinationBasePath; path != "" {
		return path
	}
	return p.providerCfg.Destination
}

// Healthcheck returns an error if a call to List Encodings with a limit of one
// returns an error
func (p *bitmovinProvider) Healthcheck() error {
	_, err := p.api.Encoding.Encodings.List(func(params *query.EncodingListQueryParams) {
		params.Limit = 1
	})
	if err != nil {
		return errors.Wrap(err, "bitmovin service unavailable")
	}

	return nil
}

// Capabilities describes the capabilities of the provider.
func (p *bitmovinProvider) Capabilities() provider.Capabilities {
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
