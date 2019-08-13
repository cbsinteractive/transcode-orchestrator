package bitmovinnewsdk

import (
	"fmt"
	"path"
	"strings"

	"github.com/NYTimes/video-transcoding-api/config"
	"github.com/NYTimes/video-transcoding-api/db"
	"github.com/NYTimes/video-transcoding-api/provider"
	"github.com/NYTimes/video-transcoding-api/provider/bitmovinnewsdk/internal/configuration"
	"github.com/NYTimes/video-transcoding-api/provider/bitmovinnewsdk/internal/container"
	"github.com/NYTimes/video-transcoding-api/provider/bitmovinnewsdk/internal/status"
	"github.com/NYTimes/video-transcoding-api/provider/bitmovinnewsdk/internal/storage"
	"github.com/bitmovin/bitmovin-api-sdk-go"
	"github.com/bitmovin/bitmovin-api-sdk-go/common"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"github.com/bitmovin/bitmovin-api-sdk-go/query"
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
	Name = "bitmovin-newsdk"

	codecVorbis = "vorbis"
	codecAAC    = "aac"
	codecVP8    = "vp8"
	codecH264   = "h264"
	codecH265   = "h265"

	containerWebM mediaContainer = "webm"
	containerHLS  mediaContainer = "m3u8"
	containerMP4  mediaContainer = "mp4"
	containerMOV  mediaContainer = "mov"

	h264AAC   cfgStore = "h264aac"
	h265AAC   cfgStore = "h265aac"
	vp8Vorbis cfgStore = "vp8vorbis"
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

	return &bitmovinProvider{
		api:         api,
		providerCfg: cfg.Bitmovin,
		cfgStores: map[cfgStore]configuration.Store{
			h264AAC:   configuration.NewH264AAC(api),
			h265AAC:   configuration.NewH265AAC(api),
			vp8Vorbis: configuration.NewVP8Vorbis(api),
		},
		containerSvcs: map[mediaContainer]containerSvc{
			containerHLS: {
				assembler:      container.NewHLSAssembler(api),
				statusEnricher: container.NewHLSStatusEnricher(api),
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
}

func (p *bitmovinProvider) Transcode(job *db.Job) (*provider.JobStatus, error) {
	inputID, mediaPath, err := storage.NewInput(job.SourceMedia, storage.InputAPI{
		S3:    p.api.Encoding.Inputs.S3,
		GCS:   p.api.Encoding.Inputs.Gcs,
		HTTP:  p.api.Encoding.Inputs.Http,
		HTTPS: p.api.Encoding.Inputs.Https,
	}, p.providerCfg)
	if err != nil {
		return nil, err
	}

	outputID, destPath, err := storage.NewOutput(p.destinationForJob(job), storage.OutputAPI{
		S3:  p.api.Encoding.Outputs.S3,
		GCS: p.api.Encoding.Outputs.Gcs,
	}, p.providerCfg)
	if err != nil {
		return nil, err
	}
	destPath = path.Join(destPath, job.ID)

	inputStream := model.StreamInput{
		InputId:       inputID,
		InputPath:     mediaPath,
		SelectionMode: model.StreamSelectionMode_AUTO,
	}

	vidInputStreams := []model.StreamInput{inputStream}
	audInputStreams := []model.StreamInput{inputStream}

	var generatingHLS bool
	for _, o := range job.Outputs {
		if o.Preset.OutputOpts.Extension == containerWebM {
			break // can't be HLSAssembler
		}

		details, err := p.cfgDetailsFrom(o.Preset.ProviderMapping[Name])
		if err != nil {
			return nil, err
		}

		contnr, err := configuration.ContainerFrom(details.CustomData)
		if err != nil {
			return nil, errors.Wrap(err, "extracting container from customData")
		}
		if contnr == containerHLS {
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
			Outputs: []model.EncodingOutput{
				{
					OutputId:   outputID,
					OutputPath: manifestMasterPath,
				},
			},
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

	audMuxingStreams := make(map[string]model.MuxingStream)
	audStreams := make(map[string]*model.Stream)
	for _, o := range job.Outputs {
		presetID := o.Preset.ProviderMapping[Name]
		details, err := p.cfgDetailsFrom(presetID)
		if err != nil {
			return nil, err
		}

		audCfgID, err := configuration.AudCfgIDFrom(details.CustomData)
		if err != nil {
			return nil, err
		}

		_, audCfgExists := audMuxingStreams[audCfgID]

		if !audCfgExists {
			audStream, err := p.api.Encoding.Encodings.Streams.Create(enc.Id, model.Stream{
				CodecConfigId: audCfgID,
				InputStreams:  audInputStreams,
			})
			if err != nil {
				return nil, errors.Wrap(err, "adding audio stream to the encoding")
			}

			audMuxingStreams[audCfgID] = model.MuxingStream{StreamId: audStream.Id}
			audStreams[audCfgID] = audStream
		}

		vidStream, err := p.api.Encoding.Encodings.Streams.Create(enc.Id, model.Stream{
			CodecConfigId: presetID,
			InputStreams:  vidInputStreams,
		})
		if err != nil {
			return nil, errors.Wrap(err, "adding video stream to the encoding")
		}

		vidMuxingStream := model.MuxingStream{StreamId: vidStream.Id}

		mediaContainer, err := configuration.ContainerFrom(details.CustomData)
		if err != nil {
			return nil, err
		}

		contnrSvcs, ok := p.containerSvcs[mediaContainer]
		if !ok {
			return nil, fmt.Errorf("unknown container format %q", mediaContainer)
		}

		if err = contnrSvcs.assembler.Assemble(container.AssemblerCfg{
			EncID:              enc.Id,
			OutputID:           outputID,
			DestPath:           destPath,
			OutputFilename:     o.FileName,
			AudCfgID:           audCfgID,
			VidCfgID:           presetID,
			AudStreamID:        audStreams[audCfgID].Id,
			VidStreamID:        vidStream.Id,
			AudMuxingStream:    audMuxingStreams[audCfgID],
			VidMuxingStream:    vidMuxingStream,
			ManifestID:         manifestID,
			ManifestMasterPath: manifestMasterPath,
			SkipAudioCreation:  audCfgExists,
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
func (p *bitmovinProvider) DeletePreset(presetID string) error {
	for _, svc := range p.cfgStores {
		found, err := svc.Delete(presetID)
		if found {
			return err
		}
	}

	return errors.New("preset not found")
}

// GetPreset searches for a preset from the registered cfg services
func (p *bitmovinProvider) GetPreset(presetID string) (interface{}, error) {
	return p.cfgDetailsFrom(presetID)
}

func (p *bitmovinProvider) cfgDetailsFrom(presetID string) (configuration.Details, error) {
	for _, svc := range p.cfgStores {
		found, preset, err := svc.Get(presetID)
		if found {
			return preset, err
		}
	}

	return configuration.Details{}, errors.New("preset not found")
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
	if vcodec == codecH264 && acodec == codecAAC {
		return p.cfgStores[h264AAC], nil
	} else if vcodec == codecH265 && acodec == codecAAC {
		return p.cfgStores[h265AAC], nil
	} else if vcodec == codecVP8 && acodec == codecVorbis {
		return p.cfgStores[vp8Vorbis], nil
	}

	return nil, fmt.Errorf("the pair of vcodec: %q and acodec: %q is not yet supported", vcodec, acodec)
}
