package bitmovin

import (
	"fmt"
	"path"
	"strings"

	"github.com/bitmovin/bitmovin-api-sdk-go"
	"github.com/bitmovin/bitmovin-api-sdk-go/common"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/query"
	"github.com/cbsinteractive/video-transcoding-api/config"
	"github.com/cbsinteractive/video-transcoding-api/db"
	"github.com/cbsinteractive/video-transcoding-api/db/redis"
	"github.com/cbsinteractive/video-transcoding-api/provider"
	"github.com/cbsinteractive/video-transcoding-api/provider/bitmovin/internal/configuration"
	"github.com/cbsinteractive/video-transcoding-api/provider/bitmovin/internal/container"
	"github.com/cbsinteractive/video-transcoding-api/provider/bitmovin/internal/status"
	"github.com/cbsinteractive/video-transcoding-api/provider/bitmovin/internal/storage"
	"github.com/pkg/errors"
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
	codecVP8    = "vp8"
	codecH264   = "h264"
	codecH265   = "h265"

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
	cfgStoreAAC       cfgStore = "aac"
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

	return &bitmovinProvider{
		api:         api,
		repo:        dbRepo,
		providerCfg: cfg.Bitmovin,
		cfgStores: map[cfgStore]configuration.Store{
			cfgStoreH264:      configuration.NewH264(api, dbRepo),
			cfgStoreH265:      configuration.NewH265(api, dbRepo),
			cfgStoreH264AAC:   configuration.NewH264AAC(api, dbRepo),
			cfgStoreH265AAC:   configuration.NewH265AAC(api, dbRepo),
			cfgStoreVP8Vorbis: configuration.NewVP8Vorbis(api, dbRepo),
			cfgStoreAAC:       configuration.NewAAC(api, dbRepo),
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
}

func (p *bitmovinProvider) Transcode(job *db.Job) (*provider.JobStatus, error) {
	presets := make([]db.PresetSummary, len(job.Outputs))
	for idx, output := range job.Outputs {
		summary, err := p.repo.GetPresetSummary(output.Preset.Name)
		if err != nil {
			return nil, err
		}

		presets[idx] = summary
	}

	inputID, mediaPath, err := p.inputFrom(job)
	if err != nil {
		return nil, err
	}

	outputID, destPath, err := p.outputFrom(job)
	if err != nil {
		return nil, err
	}

	inputStream := model.StreamInput{
		InputId:       inputID,
		InputPath:     mediaPath,
		SelectionMode: model.StreamSelectionMode_AUTO,
	}

	vidInputStreams := []model.StreamInput{inputStream}
	audInputStreams := []model.StreamInput{inputStream}

	var generatingHLS bool
	for _, preset := range presets {
		if preset.Container == containerHLS {
			generatingHLS = true
			break
		}
	}

	var manifestID, manifestMasterPath, manifestMasterFilename string
	if generatingHLS {
		manifestMasterPath = path.Dir(path.Join(destPath, job.StreamingParams.PlaylistFileName))
		manifestMasterFilename = path.Base(job.StreamingParams.PlaylistFileName)

		hlsManifest, err := p.api.Encoding.Manifests.Hls.Create(model.HlsManifest{
			ManifestName: manifestMasterFilename,
			Outputs:      []model.EncodingOutput{storage.EncodingOutputFrom(outputID, manifestMasterPath)},
		})
		if err != nil {
			return nil, errors.Wrap(err, "creating master manifest")
		}

		manifestID = hlsManifest.Id
	}

	encCustomData := make(map[string]map[string]interface{})
	if manifestID != "" {
		encCustomData[container.CustomDataKeyManifest] = map[string]interface{}{
			container.CustomDataKeyManifestID: manifestID,
		}
	}

	encodingCloudRegion, err := p.encodingCloudRegionFrom(job)
	if err != nil {
		return nil, errors.Wrap(err, "validating and setting encoding cloud region")
	}

	enc, err := p.api.Encoding.Encodings.Create(model.Encoding{
		Name:           "encoding",
		CustomData:     &encCustomData,
		CloudRegion:    encodingCloudRegion,
		EncoderVersion: p.providerCfg.EncodingVersion,
	})
	if err != nil {
		return nil, errors.Wrap(err, "creating encoding")
	}

	audMuxingStreams := map[string]model.MuxingStream{}
	for idx, o := range job.Outputs {
		preset := presets[idx]

		if audCfgID := preset.AudioConfigID; audCfgID != "" && audMuxingStreams[audCfgID].StreamId == "" {
			audStream, err := p.api.Encoding.Encodings.Streams.Create(enc.Id, model.Stream{
				CodecConfigId: audCfgID,
				InputStreams:  audInputStreams,
			})
			if err != nil {
				return nil, errors.Wrap(err, "adding audio stream to the encoding")
			}

			audMuxingStreams[audCfgID] = model.MuxingStream{StreamId: audStream.Id}
		}

		var vidMuxingStream model.MuxingStream
		if vidCfgID := preset.VideoConfigID; vidCfgID != "" {
			vidStream, err := p.api.Encoding.Encodings.Streams.Create(enc.Id, model.Stream{
				CodecConfigId: vidCfgID,
				InputStreams:  vidInputStreams,
			})
			if err != nil {
				return nil, errors.Wrap(err, "adding video stream to the encoding")
			}

			vidMuxingStream = model.MuxingStream{StreamId: vidStream.Id}
		}

		contnrSvcs, err := p.containerServicesFrom(preset.Container, model.CodecConfigType(preset.VideoCodec))
		if err != nil {
			return nil, err
		}

		if err = contnrSvcs.assembler.Assemble(container.AssemblerCfg{
			EncID:              enc.Id,
			OutputID:           outputID,
			DestPath:           destPath,
			OutputFilename:     o.FileName,
			AudCfgID:           preset.AudioConfigID,
			VidCfgID:           preset.VideoConfigID,
			AudMuxingStream:    audMuxingStreams[preset.AudioConfigID],
			VidMuxingStream:    vidMuxingStream,
			ManifestID:         manifestID,
			ManifestMasterPath: manifestMasterPath,
			SegDuration:        job.StreamingParams.SegmentDuration,
		}); err != nil {
			return nil, err
		}
	}

	var vodHLSManifests []model.ManifestResource
	if generatingHLS && manifestID != "" {
		vodHLSManifests = []model.ManifestResource{{ManifestId: manifestID}}
	}

	encResp, err := p.api.Encoding.Encodings.Start(enc.Id, model.StartEncodingRequest{VodHlsManifests: vodHLSManifests})
	if err != nil {
		return nil, errors.Wrap(err, "starting encoding job")
	}

	return &provider.JobStatus{
		ProviderName:  Name,
		ProviderJobID: encResp.Id,
		Status:        provider.StatusQueued,
	}, nil
}

func (p *bitmovinProvider) inputFrom(job *db.Job) (inputID string, srcPath string, err error) {
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

func (p *bitmovinProvider) outputFrom(job *db.Job) (inputID string, destPath string, err error) {
	destPath, err = storage.PathFrom(job.SourceMedia)
	if err != nil {
		return "", "", err
	}

	if alias := job.ExecutionEnv.OutputAlias; alias != "" {
		return alias, destPath, nil
	}

	outputID, err := storage.NewOutput(p.destinationForJob(job), storage.OutputAPI{
		S3:  p.api.Encoding.Outputs.S3,
		GCS: p.api.Encoding.Outputs.Gcs,
	}, p.providerCfg)
	if err != nil {
		return "", destPath, err
	}
	destPath = path.Join(destPath, job.ID)

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

func (p *bitmovinProvider) JobStatus(job *db.Job) (*provider.JobStatus, error) {
	task, err := p.api.Encoding.Encodings.Status(job.ProviderJobID)
	if err != nil {
		return nil, errors.Wrap(err, "retrieving encoding status")
	}

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
		s, err = status.EnrichSourceInfo(p.api, s)
		if err != nil {
			return nil, errors.Wrap(err, "enriching status with source info")
		}

		// TODO: it would be better to know which containers to include in this fetch
		// rather than iterating over all supported containers
		for _, svcs := range p.containerSvcs {
			s, err = svcs.statusEnricher.Enrich(s)
			if err != nil {
				return nil, err
			}
		}
	}

	return &s, nil
}

func (p *bitmovinProvider) CancelJob(id string) error {
	_, err := p.api.Encoding.Encodings.Stop(id)

	return err
}

func (p *bitmovinProvider) CreatePreset(preset db.Preset) (string, error) {
	svc, err := p.cfgServiceFrom(preset.Video.Codec, preset.Audio.Codec)
	if err != nil {
		return "", err
	}

	return svc.Create(preset)
}

// DeletePreset loops over registered cfg services and attempts to delete them
func (p *bitmovinProvider) DeletePreset(presetName string) error {
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
func (p *bitmovinProvider) GetPreset(presetName string) (interface{}, error) {
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
func (bitmovinProvider) Capabilities() provider.Capabilities {
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
	case codecAAC:
		return p.cfgStores[cfgStoreAAC], nil
	}

	return nil, fmt.Errorf("the pair of vcodec: %q and acodec: %q is not yet supported", vcodec, acodec)
}
