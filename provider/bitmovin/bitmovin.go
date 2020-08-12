package bitmovin

import (
	"context"
	"fmt"

	"net/url"
	"path"
	"sort"
	"strings"
	"sync"

	"github.com/bitmovin/bitmovin-api-sdk-go"
	"github.com/bitmovin/bitmovin-api-sdk-go/common"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/query"
	"github.com/cbsinteractive/transcode-orchestrator/config"
	"github.com/cbsinteractive/transcode-orchestrator/db"
	"github.com/cbsinteractive/transcode-orchestrator/db/redis"
	"github.com/cbsinteractive/transcode-orchestrator/provider"
	"github.com/cbsinteractive/transcode-orchestrator/provider/bitmovin/internal/configuration"
	"github.com/cbsinteractive/transcode-orchestrator/provider/bitmovin/internal/container"
	"github.com/cbsinteractive/transcode-orchestrator/provider/bitmovin/internal/status"
	"github.com/cbsinteractive/transcode-orchestrator/provider/bitmovin/internal/storage"
	"github.com/pkg/errors"
	"github.com/zsiec/pkg/tracing"
)

type containerSvc struct {
	assembler      container.Assembler
	statusEnricher container.StatusEnricher
}

type cfgStore string

type mediaContainer = string

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

	containerWebM    mediaContainer = "webm"
	containerHLS     mediaContainer = "m3u8"
	containerMP4     mediaContainer = "mp4"
	containerMOV     mediaContainer = "mov"
	containerCMAFHLS mediaContainer = "cmafhls"

	cfgStoreH264AAC   cfgStore = "h264aac"
	cfgStoreH265AAC   cfgStore = "h265aac"
	cfgStoreVP8Vorbis cfgStore = "vp8vorbis"
	cfgStoreH264      cfgStore = "h264"
	cfgStoreH265      cfgStore = "h265"
	cfgStoreAV1       cfgStore = "av1"
	cfgStoreAAC       cfgStore = "aac"
	cfgStoreOpus      cfgStore = "opus"

	filterVideoDeinterlace = iota + 1
)

func init() {
	_ = provider.Register(Name, bitmovinFactory)
}

var cloudRegions = map[model.CloudRegion]struct{}{
	model.CloudRegion_AWS_US_EAST_1: {}, model.CloudRegion_AWS_US_EAST_2: {}, model.CloudRegion_AWS_US_WEST_1: {},
	model.CloudRegion_AWS_US_WEST_2: {}, model.CloudRegion_AWS_EU_WEST_1: {}, model.CloudRegion_AWS_EU_CENTRAL_1: {},
	model.CloudRegion_AWS_AP_SOUTHEAST_1: {}, model.CloudRegion_AWS_AP_SOUTHEAST_2: {}, model.CloudRegion_AWS_AP_NORTHEAST_1: {},
	model.CloudRegion_AWS_AP_NORTHEAST_2: {}, model.CloudRegion_AWS_AP_SOUTH_1: {}, model.CloudRegion_AWS_SA_EAST_1: {},
	model.CloudRegion_AWS_EU_WEST_2: {}, model.CloudRegion_AWS_EU_WEST_3: {}, model.CloudRegion_AWS_CA_CENTRAL_1: {},
	model.CloudRegion_GOOGLE_US_CENTRAL_1: {}, model.CloudRegion_GOOGLE_US_EAST_1: {}, model.CloudRegion_GOOGLE_ASIA_EAST_1: {},
	model.CloudRegion_GOOGLE_EUROPE_WEST_1: {}, model.CloudRegion_GOOGLE_US_WEST_1: {}, model.CloudRegion_GOOGLE_ASIA_EAST_2: {},
	model.CloudRegion_GOOGLE_ASIA_NORTHEAST_1: {}, model.CloudRegion_GOOGLE_ASIA_SOUTH_1: {}, model.CloudRegion_GOOGLE_ASIA_SOUTHEAST_1: {},
	model.CloudRegion_GOOGLE_AUSTRALIA_SOUTHEAST_1: {}, model.CloudRegion_GOOGLE_EUROPE_NORTH_1: {}, model.CloudRegion_GOOGLE_EUROPE_WEST_2: {},
	model.CloudRegion_GOOGLE_EUROPE_WEST_4: {}, model.CloudRegion_GOOGLE_NORTHAMERICA_NORTHEAST_1: {}, model.CloudRegion_GOOGLE_SOUTHAMERICA_EAST_1: {},
	model.CloudRegion_GOOGLE_US_EAST_4: {}, model.CloudRegion_GOOGLE_US_WEST_2: {}, model.CloudRegion_AZURE_EUROPE_WEST: {},
	model.CloudRegion_AZURE_US_WEST2: {}, model.CloudRegion_AZURE_US_EAST: {}, model.CloudRegion_AZURE_AUSTRALIA_SOUTHEAST: {},
	model.CloudRegion_NORTH_AMERICA: {}, model.CloudRegion_SOUTH_AMERICA: {}, model.CloudRegion_EUROPE: {},
	model.CloudRegion_AFRICA: {}, model.CloudRegion_ASIA: {}, model.CloudRegion_AUSTRALIA: {},
	model.CloudRegion_AWS: {}, model.CloudRegion_GOOGLE: {}, model.CloudRegion_KUBERNETES: {},
	model.CloudRegion_EXTERNAL: {}, model.CloudRegion_AUTO: {},
}

var regionByCloud = map[string]map[string]model.CloudRegion{
	provider.CloudAWS: {
		provider.AWSRegionUSEast1: model.CloudRegion_AWS_US_EAST_1,
		provider.AWSRegionUSEast2: model.CloudRegion_AWS_US_EAST_2,
		provider.AWSRegionUSWest1: model.CloudRegion_AWS_US_WEST_1,
		provider.AWSRegionUSWest2: model.CloudRegion_AWS_US_WEST_2,
	},
	provider.CloudGCP: {
		provider.GCPRegionUSEast1:    model.CloudRegion_GOOGLE_US_EAST_1,
		provider.GCPRegionUSEast4:    model.CloudRegion_GOOGLE_US_EAST_4,
		provider.GCPRegionUSWest1:    model.CloudRegion_GOOGLE_US_WEST_1,
		provider.GCPRegionUSWest2:    model.CloudRegion_GOOGLE_US_WEST_2,
		provider.GCPRegionUSCentral1: model.CloudRegion_GOOGLE_US_CENTRAL_1,
	},
}

var awsCloudRegions = map[model.AwsCloudRegion]struct{}{
	model.AwsCloudRegion_US_EAST_1: {}, model.AwsCloudRegion_US_EAST_2: {}, model.AwsCloudRegion_US_WEST_1: {},
	model.AwsCloudRegion_US_WEST_2: {}, model.AwsCloudRegion_EU_WEST_1: {}, model.AwsCloudRegion_EU_CENTRAL_1: {},
	model.AwsCloudRegion_AP_SOUTHEAST_1: {}, model.AwsCloudRegion_AP_SOUTHEAST_2: {}, model.AwsCloudRegion_AP_NORTHEAST_1: {},
	model.AwsCloudRegion_AP_NORTHEAST_2: {}, model.AwsCloudRegion_AP_SOUTH_1: {}, model.AwsCloudRegion_SA_EAST_1: {},
	model.AwsCloudRegion_EU_WEST_2: {}, model.AwsCloudRegion_EU_WEST_3: {}, model.AwsCloudRegion_CA_CENTRAL_1: {},
}

var errBitmovinInvalidConfig = provider.InvalidConfigError("Invalid configuration")

func bitmovinFactory(cfg *config.Config) (provider.TranscodingProvider, error) {
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

	dbRepo, err := redis.NewRepository(cfg)
	if err != nil {
		return nil, fmt.Errorf("error initializing bitmovin wrapper: %s", err)
	}

	tracer := cfg.Tracer
	if tracer == nil {
		tracer = tracing.NoopTracer{}
	}

	return &bitmovinProvider{
		api:         api,
		repo:        dbRepo,
		providerCfg: cfg.Bitmovin,
		tracer:      tracer,
		cfgStores: map[cfgStore]configuration.Store{
			cfgStoreH264:      configuration.NewH264(api, dbRepo),
			cfgStoreH265:      configuration.NewH265(api, dbRepo),
			cfgStoreAV1:       configuration.NewAV1(api, dbRepo),
			cfgStoreH264AAC:   configuration.NewH264AAC(api, dbRepo),
			cfgStoreH265AAC:   configuration.NewH265AAC(api, dbRepo),
			cfgStoreVP8Vorbis: configuration.NewVP8Vorbis(api, dbRepo),
			cfgStoreAAC:       configuration.NewAAC(api, dbRepo),
			cfgStoreOpus:      configuration.NewOpus(api, dbRepo),
		},
		containerSvcs: map[mediaContainer]containerSvc{
			containerHLS: {
				assembler: container.NewHLSAssembler(container.HLSContainerAPI{
					HLSAudioMedia: api.Encoding.Manifests.Hls.Media.Audio,
					TSMuxing:      api.Encoding.Encodings.Muxings.Ts,
					HLSStreams:    api.Encoding.Manifests.Hls.Streams,
				}),
				statusEnricher: container.NewHLSStatusEnricher(api),
			},
			containerCMAFHLS: {
				assembler: container.NewCMAFAssembler(container.CMAFContainerAPI{
					HLSAudioMedia: api.Encoding.Manifests.Hls.Media.Audio,
					CMAFMuxing:    api.Encoding.Encodings.Muxings.Cmaf,
					HLSStreams:    api.Encoding.Manifests.Hls.Streams,
				}),
				statusEnricher: container.NewCMAFStatusEnricher(api),
			},
			containerWebM: {
				assembler:      container.NewProgressiveWebMAssembler(api),
				statusEnricher: container.NewProgressiveWebMStatusEnricher(api),
			},
			containerMP4: {
				assembler:      container.NewMP4Assembler(api),
				statusEnricher: container.NewMP4StatusEnricher(api),
			},
			containerMOV: {
				assembler:      container.NewMOVAssembler(api),
				statusEnricher: container.NewMOVStatusEnricher(api),
			},
		},
	}, nil
}

type bitmovinProvider struct {
	api           *bitmovin.BitmovinApi
	providerCfg   *config.Bitmovin
	cfgStores     map[cfgStore]configuration.Store
	containerSvcs map[mediaContainer]containerSvc
	repo          db.Repository
	tracer        tracing.Tracer
	presetMutex   sync.Mutex
}

func (p *bitmovinProvider) Transcode(ctx context.Context, job *db.Job) (*provider.JobStatus, error) {
	presets := make([]db.PresetSummary, len(job.Outputs))
	for idx, output := range job.Outputs {
		summary, err := p.repo.GetPresetSummary(output.Preset.Name)
		if err != nil {
			return nil, err
		}

		presets[idx] = summary
	}

	inputID, mediaPath, err := p.inputFrom(ctx, job)
	if err != nil {
		return nil, err
	}

	outputID, destPath, err := p.outputFrom(ctx, job)
	if err != nil {
		return nil, err
	}

	var generatingHLS, processingVideo bool
	for _, preset := range presets {
		if preset.Container == containerHLS {
			generatingHLS = true
		}
		if preset.VideoConfigID != "" {
			processingVideo = true
		}
	}

	var manifestID, manifestMasterPath, manifestMasterFilename string
	if generatingHLS {
		manifestMasterPath = path.Dir(path.Join(destPath, job.StreamingParams.PlaylistFileName))
		manifestMasterFilename = path.Base(job.StreamingParams.PlaylistFileName)

		subSeg := p.tracer.BeginSubsegment(ctx, "bitmovin-create-hls-manifest")
		hlsManifest, err := p.api.Encoding.Manifests.Hls.Create(model.HlsManifest{
			ManifestName: manifestMasterFilename,
			Outputs:      []model.EncodingOutput{storage.EncodingOutputFrom(outputID, manifestMasterPath)},
		})
		if err != nil {
			subSeg.Close(err)
			return nil, errors.Wrap(err, "creating master manifest")
		}
		subSeg.Close(nil)

		manifestID = hlsManifest.Id
	}

	encCustomData := make(map[string]map[string]interface{})
	if manifestID != "" {
		encCustomData[container.CustomDataKeyManifest] = map[string]interface{}{
			container.CustomDataKeyManifestID: manifestID,
		}
	}

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

	subSeg = p.tracer.BeginSubsegment(ctx, "bitmovin-create-concatenated-splice")
	inputID, err = func(inputID string) (string, error) {
		if len(job.SourceSplice) == 0 {
			return inputID, nil
		}

		type work struct {
			pos        int32
			start, dur float64
			id         string
			err        error
		}
		workc := make(chan work, len(job.SourceSplice))

		// splice each range concurrently
		for i, r := range job.SourceSplice {
			w := work{
				pos:   int32(i),
				start: r[0],
				dur:   r[1] - r[0],
			}
			go func() {
				// NOTE(as): don't use the timecode "api", it seems to look for a real
				// timecode track in the source. If it doesn't find it, it just doesn't trim
				// the clip and provides no logging or errors. For this "api", it wants
				// start, duration; not start, end, and it also wants pointers
				splice, err := p.api.Encoding.Encodings.InputStreams.Trimming.TimeBased.Create(enc.Id, model.TimeBasedTrimmingInputStream{
					InputStreamId: inputID,
					Offset:        &w.start,
					Duration:      &w.dur,
				})
				if splice != nil {
					w.id = splice.Id
				}
				w.err = err
				workc <- w
			}()
		}

		// collect the results serially
		cat := []model.ConcatenationInputConfiguration{}
		for i := 0; i < cap(workc); i++ {
			w := <-workc
			if w.err != nil {
				return inputID, fmt.Errorf("trim: range#%d: %w", w.pos, err)
			}
			main := i == 0
			cat = append(cat, model.ConcatenationInputConfiguration{
				IsMain:        &main,
				InputStreamId: w.id,
				Position:      &w.pos,
			})
		}

		// although there are position markers in the struct,
		// sort it just in case, this makes the logging consistent too
		sort.Slice(cat, func(i, j int) bool {
			return *cat[i].Position < *cat[j].Position
		})
		c, err := p.api.Encoding.Encodings.InputStreams.Concatenation.Create(enc.Id, model.ConcatenationInputStream{
			Concatenation: cat,
		})
		if err != nil {
			return inputID, fmt.Errorf("concatenation: %v", err)
		}
		return c.Id, nil
	}(inputID)
	subSeg.Close(err)
	if err != nil {
		return nil, fmt.Errorf("splice: %w", err)
	}

	videoFilters := map[int]string{}
	if processingVideo {
		subSeg := p.tracer.BeginSubsegment(ctx, "bitmovin-create-deinterlace-filter")
		deInterlace, err := p.api.Encoding.Filters.Deinterlace.Create(model.DeinterlaceFilter{
			Name:       "deinterlace",
			AutoEnable: model.DeinterlaceAutoEnable_META_DATA_AND_CONTENT_BASED,
		})
		if err != nil {
			subSeg.Close(err)
			return nil, errors.Wrap(err, "creating deinterlace filter")
		}
		subSeg.Close(nil)
		videoFilters[filterVideoDeinterlace] = deInterlace.Id
	}

	var wg sync.WaitGroup
	errorc := make(chan error)

	subSeg = p.tracer.BeginSubsegment(ctx, "bitmovin-create-outputs")
	for idx, o := range job.Outputs {
		wg.Add(1)
		go p.createOutput(outputCfg{
			preset:             presets[idx],
			encodingID:         enc.Id,
			audioIn:            inputID,
			videoIn:            inputID,
			videoFilters:       videoFilters,
			outputID:           outputID,
			outputFilename:     o.FileName,
			destPath:           destPath,
			manifestID:         manifestID,
			manifestMasterPath: manifestMasterPath,
			job:                job,
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

	subSeg = p.tracer.BeginSubsegment(ctx, "bitmovin-create-splice")
	err = func() error {
		for _, r := range job.SourceSplice {
			sp, ep := r.Timecodes(0)
			splice, err := p.api.Encoding.Encodings.InputStreams.Trimming.TimecodeTrack.Create(enc.Id, model.TimecodeTrackTrimmingInputStream{
				Name:          "splice",
				StartTimeCode: sp,
				EndTimeCode:   ep,
			})
			if err != nil {
				return err
			}
			splice = splice
		}
		return nil
	}()
	subSeg.Close(err)
	if err != nil {
		return nil, fmt.Errorf("splice: %w", err)
	}

	var vodHLSManifests []model.ManifestResource
	if generatingHLS && manifestID != "" {
		vodHLSManifests = []model.ManifestResource{{ManifestId: manifestID}}
	}

	subSeg = p.tracer.BeginSubsegment(ctx, "bitmovin-start-encoding")
	encResp, err := p.api.Encoding.Encodings.Start(enc.Id, model.StartEncodingRequest{VodHlsManifests: vodHLSManifests})
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

type outputCfg struct {
	preset             db.PresetSummary
	encodingID         string
	videoIn, audioIn   string
	videoFilters       map[int]string
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

		if deinterlaceID, ok := cfg.videoFilters[filterVideoDeinterlace]; ok {
			_, err = p.api.Encoding.Encodings.Streams.Filters.Create(cfg.encodingID, vidStream.Id, []model.StreamFilter{
				{Id: deinterlaceID, Position: bitmovin.Int32Ptr(0)},
			})
			if err != nil {
				errorc <- errors.Wrap(err, "adding filter to video stream")
				return
			}
		}

		videoMuxingStream = model.MuxingStream{StreamId: vidStream.Id}
	}

	contnrSvcs, err := p.containerServicesFrom(cfg.preset.Container, model.CodecConfigType(cfg.preset.VideoCodec))
	if err != nil {
		errorc <- err
		return
	}

	if err = contnrSvcs.assembler.Assemble(container.AssemblerCfg{
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
	srcPath, err = storage.PathFrom(job.SourceMedia)
	if err != nil {
		return "", "", err
	}

	if alias := job.ExecutionEnv.InputAlias; alias != "" {
		return alias, srcPath, nil
	}

	subSeg := p.tracer.BeginSubsegment(ctx, "bitmovin-create-input")

	inputID, err = storage.NewInput(job.SourceMedia, storage.InputAPI{
		S3:    p.api.Encoding.Inputs.S3,
		GCS:   p.api.Encoding.Inputs.Gcs,
		HTTP:  p.api.Encoding.Inputs.Http,
		HTTPS: p.api.Encoding.Inputs.Https,
	}, p.providerCfg)
	if err != nil {
		subSeg.Close(err)
		return "", srcPath, err
	}
	subSeg.Close(nil)

	return inputID, srcPath, nil
}

func (p *bitmovinProvider) outputFrom(ctx context.Context, job *db.Job) (inputID string, destPath string, err error) {
	destBasePath := p.destinationForJob(job)
	destURL, err := url.Parse(destBasePath)
	if err != nil {
		return "", "", err
	}
	destPath = path.Join(destURL.Path, job.ID)

	if alias := job.ExecutionEnv.OutputAlias; alias != "" {
		return alias, destPath, nil
	}

	subSeg := p.tracer.BeginSubsegment(ctx, "bitmovin-create-output")
	defer subSeg.Close(nil)

	outputID, err := storage.NewOutput(destBasePath, storage.OutputAPI{
		S3:  p.api.Encoding.Outputs.S3,
		GCS: p.api.Encoding.Outputs.Gcs,
	}, p.providerCfg)
	if err != nil {
		return "", destPath, err
	}

	return outputID, destPath, nil
}

func (p *bitmovinProvider) containerServicesFrom(mediaContainer string, cfgType model.CodecConfigType) (containerSvc, error) {
	if cfgType == model.CodecConfigType_H265 && mediaContainer == containerHLS {
		mediaContainer = containerCMAFHLS
	}

	containerSvcs, ok := p.containerSvcs[mediaContainer]
	if !ok {
		return containerSvc{}, fmt.Errorf("unknown container format %q", mediaContainer)
	}

	return containerSvcs, nil
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
		Status:        status.ToProviderStatus(task.Status),
		Progress:      progress,
		ProviderStatus: map[string]interface{}{
			"messages":       task.Messages,
			"originalStatus": task.Status,
		},
		Output: provider.JobOutput{
			Destination: strings.TrimRight(p.destinationForJob(job), "/") + "/" + job.ID + "/",
		},
	}

	if s.Status == provider.StatusFinished {
		subSeg := p.tracer.BeginSubsegment(ctx, "bitmovin-get-output-info")
		s, err = status.EnrichSourceInfo(p.api, s)
		if err != nil {
			subSeg.Close(err)
			return nil, errors.Wrap(err, "enriching status with source info")
		}

		// TODO: it would be better to know which containers to include in this fetch
		// rather than iterating over all supported containers
		for _, svcs := range p.containerSvcs {
			s, err = svcs.statusEnricher.Enrich(s)
			if err != nil {
				subSeg.Close(err)
				return nil, err
			}
		}
		subSeg.Close(nil)
	}

	return &s, nil
}

func (p *bitmovinProvider) CancelJob(ctx context.Context, id string) error {
	subSeg := p.tracer.BeginSubsegment(ctx, "bitmovin-delete-job")
	_, err := p.api.Encoding.Encodings.Stop(id)
	subSeg.Close(err)

	return err
}

func (p *bitmovinProvider) CreatePreset(_ context.Context, preset db.Preset) (string, error) {
	p.presetMutex.Lock()
	defer p.presetMutex.Unlock()

	if existing, err := p.repo.GetPresetSummary(preset.Name); err == nil {
		return existing.Name, nil
	}

	svc, err := p.cfgServiceFrom(preset.Video.Codec, preset.Audio.Codec)
	if err != nil {
		return "", err
	}

	return svc.Create(preset)
}

// DeletePreset loops over registered cfg services and attempts to delete them
func (p *bitmovinProvider) DeletePreset(_ context.Context, presetName string) error {
	summary, err := p.repo.GetPresetSummary(presetName)
	if err != nil {
		return err
	}

	svc, err := p.cfgServiceFrom(summary.VideoCodec, summary.AudioCodec)
	if err != nil {
		return err
	}

	return svc.Delete(presetName)
}

// GetPreset searches for a preset from the registered cfg services
func (p *bitmovinProvider) GetPreset(_ context.Context, presetName string) (interface{}, error) {
	return p.repo.GetPresetSummary(presetName)
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
		OutputFormats: []string{containerMP4, containerMOV, containerHLS, containerWebM},
		Destinations:  []string{"s3", "gcs"},
	}
}

func (p *bitmovinProvider) cfgServiceFrom(vcodec, acodec string) (configuration.Store, error) {
	vcodec, acodec = strings.ToLower(vcodec), strings.ToLower(acodec)

	switch vcodec + acodec {
	case codecH264 + codecAAC:
		return p.cfgStores[cfgStoreH264AAC], nil
	case codecH265 + codecAAC:
		return p.cfgStores[cfgStoreH265AAC], nil
	case codecVP8 + codecVorbis:
		return p.cfgStores[cfgStoreVP8Vorbis], nil
	case codecH264:
		return p.cfgStores[cfgStoreH264], nil
	case codecH265:
		return p.cfgStores[cfgStoreH265], nil
	case codecAV1:
		return p.cfgStores[cfgStoreAV1], nil
	case codecAAC:
		return p.cfgStores[cfgStoreAAC], nil
	case codecOpus:
		return p.cfgStores[cfgStoreOpus], nil
	}

	return nil, fmt.Errorf("the pair of vcodec: %q and acodec: %q is not yet supported", vcodec, acodec)
}
