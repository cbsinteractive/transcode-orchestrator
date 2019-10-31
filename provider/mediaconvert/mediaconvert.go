package mediaconvert

import (
	"context"
	"fmt"
	"path"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/mediaconvert"
	"github.com/cbsinteractive/video-transcoding-api/config"
	"github.com/cbsinteractive/video-transcoding-api/db"
	"github.com/cbsinteractive/video-transcoding-api/db/redis"
	"github.com/cbsinteractive/video-transcoding-api/provider"
	"github.com/pkg/errors"
)

const (
	// Name identifies the MediaConvert provider by name
	Name = "mediaconvert"

	defaultAudioSampleRate = 48000
)

func init() {
	err := provider.Register(Name, mediaconvertFactory)
	if err != nil {
		fmt.Printf("registering mediaconvert factory: %v", err)
	}
}

type mediaconvertClient interface {
	CreateJobRequest(*mediaconvert.CreateJobInput) mediaconvert.CreateJobRequest
	GetJobRequest(*mediaconvert.GetJobInput) mediaconvert.GetJobRequest
	ListJobsRequest(*mediaconvert.ListJobsInput) mediaconvert.ListJobsRequest
	CancelJobRequest(*mediaconvert.CancelJobInput) mediaconvert.CancelJobRequest
	CreatePresetRequest(*mediaconvert.CreatePresetInput) mediaconvert.CreatePresetRequest
	GetPresetRequest(*mediaconvert.GetPresetInput) mediaconvert.GetPresetRequest
	DeletePresetRequest(*mediaconvert.DeletePresetInput) mediaconvert.DeletePresetRequest
}

type mcProvider struct {
	client     mediaconvertClient
	cfg        *config.MediaConvert
	repository db.Repository
}

type outputCfg struct {
	output   mediaconvert.Output
	filename string
}

func (p *mcProvider) Transcode(job *db.Job) (*provider.JobStatus, error) {
	outputGroups, err := p.outputGroupsFrom(job)
	if err != nil {
		return nil, errors.Wrap(err, "generating Mediaconvert output groups")
	}

	createJobInput := mediaconvert.CreateJobInput{
		Queue: aws.String(p.cfg.Queue),
		Role:  aws.String(p.cfg.Role),
		Settings: &mediaconvert.JobSettings{
			Inputs: []mediaconvert.Input{
				{
					FileInput: aws.String(job.SourceMedia),
					AudioSelectors: map[string]mediaconvert.AudioSelector{
						"Audio Selector 1": {DefaultSelection: mediaconvert.AudioDefaultSelectionDefault},
					},
					VideoSelector: &mediaconvert.VideoSelector{
						ColorSpace: mediaconvert.ColorSpaceFollow,
					},
				},
			},
			OutputGroups: outputGroups,
		},
	}

	resp, err := p.client.CreateJobRequest(&createJobInput).Send(context.Background())
	if err != nil {
		return nil, err
	}
	return &provider.JobStatus{
		ProviderName:  Name,
		ProviderJobID: aws.StringValue(resp.Job.Id),
		Status:        provider.StatusQueued,
	}, nil
}

func (p *mcProvider) outputGroupsFrom(job *db.Job) ([]mediaconvert.OutputGroup, error) {
	outputGroups := map[mediaconvert.ContainerType][]outputCfg{}
	for _, output := range job.Outputs {
		presetName := output.Preset.Name
		presetResponse, err := p.GetPreset(presetName)
		if err != nil {
			return nil, err
		}

		localPreset, ok := presetResponse.(*db.LocalPreset)
		if !ok {
			return nil, fmt.Errorf("could not convert preset response into a db.LocalPreset")
		}

		mcOutput, err := outputFrom(localPreset.Preset)
		if err != nil {
			return nil, fmt.Errorf("could not determine output settings from db.Preset %v: %w",
				localPreset.Preset, err)
		}

		cSettings := mcOutput.ContainerSettings
		if cSettings == nil {
			return nil, fmt.Errorf("no container was found on outout settings %+v", mcOutput)
		}

		outputGroups[cSettings.Container] = append(outputGroups[cSettings.Container], outputCfg{
			output:   mcOutput,
			filename: output.FileName,
		})
	}

	mcOutputGroups := []mediaconvert.OutputGroup{}
	for container, outputs := range outputGroups {
		mcOutputGroup := mediaconvert.OutputGroup{}

		mcOutputs := make([]mediaconvert.Output, len(outputs))
		for i, o := range outputs {
			rawExtension := path.Ext(o.filename)
			filename := strings.Replace(path.Base(o.filename), rawExtension, "", 1)
			extension := strings.Replace(rawExtension, ".", "", -1)

			mcOutputs[i] = mediaconvert.Output{
				NameModifier:      aws.String(filename),
				Extension:         aws.String(extension),
				ContainerSettings: o.output.ContainerSettings,
				AudioDescriptions: o.output.AudioDescriptions,
				VideoDescription:  o.output.VideoDescription,
			}
		}
		mcOutputGroup.Outputs = mcOutputs

		destination := destinationPathFrom(p.cfg.Destination, job.ID)

		switch container {
		case mediaconvert.ContainerTypeCmfc:
			mcOutputGroup.OutputGroupSettings = &mediaconvert.OutputGroupSettings{
				Type: mediaconvert.OutputGroupTypeCmafGroupSettings,
				CmafGroupSettings: &mediaconvert.CmafGroupSettings{
					Destination:            aws.String(destination),
					FragmentLength:         aws.Int64(int64(job.StreamingParams.SegmentDuration)),
					ManifestDurationFormat: mediaconvert.CmafManifestDurationFormatFloatingPoint,
					SegmentControl:         mediaconvert.CmafSegmentControlSegmentedFiles,
					SegmentLength:          aws.Int64(int64(job.StreamingParams.SegmentDuration)),
					WriteDashManifest:      mediaconvert.CmafWriteDASHManifestEnabled,
					WriteHlsManifest:       mediaconvert.CmafWriteHLSManifestEnabled,
				},
			}
		case mediaconvert.ContainerTypeM3u8:
			mcOutputGroup.OutputGroupSettings = &mediaconvert.OutputGroupSettings{
				Type: mediaconvert.OutputGroupTypeHlsGroupSettings,
				HlsGroupSettings: &mediaconvert.HlsGroupSettings{
					Destination:            aws.String(destination),
					SegmentLength:          aws.Int64(int64(job.StreamingParams.SegmentDuration)),
					MinSegmentLength:       aws.Int64(0),
					DirectoryStructure:     mediaconvert.HlsDirectoryStructureSingleDirectory,
					ManifestDurationFormat: mediaconvert.HlsManifestDurationFormatFloatingPoint,
					OutputSelection:        mediaconvert.HlsOutputSelectionManifestsAndSegments,
					SegmentControl:         mediaconvert.HlsSegmentControlSegmentedFiles,
				},
			}
		case mediaconvert.ContainerTypeMp4:
			mcOutputGroup.OutputGroupSettings = &mediaconvert.OutputGroupSettings{
				Type: mediaconvert.OutputGroupTypeFileGroupSettings,
				FileGroupSettings: &mediaconvert.FileGroupSettings{
					Destination: aws.String(destination),
				},
			}
		default:
			return nil, fmt.Errorf("container %s is not yet supported with mediaconvert", string(container))
		}

		mcOutputGroups = append(mcOutputGroups, mcOutputGroup)
	}

	return mcOutputGroups, nil
}

func destinationPathFrom(destBase string, jobID string) string {
	return fmt.Sprintf("%s/%s/", strings.TrimRight(destBase, "/"), jobID)
}

func (p *mcProvider) CreatePreset(preset db.Preset) (string, error) {
	err := p.repository.CreateLocalPreset(&db.LocalPreset{
		Name:   preset.Name,
		Preset: preset,
	})
	if err != nil {
		return "", err
	}

	return preset.Name, nil
}

func (p *mcProvider) GetPreset(presetID string) (interface{}, error) {
	return p.repository.GetLocalPreset(presetID)
}

func (p *mcProvider) DeletePreset(presetID string) error {
	preset, err := p.GetPreset(presetID)
	if err != nil {
		return err
	}

	return p.repository.DeleteLocalPreset(preset.(*db.LocalPreset))
}

func (p *mcProvider) JobStatus(job *db.Job) (*provider.JobStatus, error) {
	jobResp, err := p.client.GetJobRequest(&mediaconvert.GetJobInput{
		Id: aws.String(job.ProviderJobID),
	}).Send(context.Background())
	if err != nil {
		return &provider.JobStatus{}, errors.Wrap(err, "fetching job info with the mediaconvert API")
	}

	return p.jobStatusFrom(job.ProviderJobID, job.ID, jobResp.Job), nil
}

func (p *mcProvider) jobStatusFrom(providerJobID string, jobID string, job *mediaconvert.Job) *provider.JobStatus {
	status := &provider.JobStatus{
		ProviderJobID: providerJobID,
		ProviderName:  Name,
		Status:        providerStatusFrom(job.Status),
		StatusMessage: statusMsgFrom(job),
		Output: provider.JobOutput{
			Destination: destinationPathFrom(p.cfg.Destination, jobID),
		},
	}

	if status.Status == provider.StatusFinished {
		status.Progress = 100
	} else if p := job.JobPercentComplete; p != nil {
		status.Progress = float64(*p)
	}

	var files []provider.OutputFile
	for _, groupDetails := range job.OutputGroupDetails {
		for _, outputDetails := range groupDetails.OutputDetails {
			if outputDetails.VideoDetails == nil {
				continue
			}

			file := provider.OutputFile{}

			if height := outputDetails.VideoDetails.HeightInPx; height != nil {
				file.Height = *height
			}

			if width := outputDetails.VideoDetails.WidthInPx; width != nil {
				file.Width = *width
			}

			files = append(files, file)
		}
	}
	status.Output.Files = files

	return status
}

func statusMsgFrom(job *mediaconvert.Job) string {
	if job.ErrorMessage != nil {
		return *job.ErrorMessage
	}

	return string(job.CurrentPhase)
}

func (p *mcProvider) CancelJob(id string) error {
	_, err := p.client.CancelJobRequest(&mediaconvert.CancelJobInput{
		Id: aws.String(id),
	}).Send(context.Background())

	return err
}

func (p *mcProvider) Healthcheck() error {
	_, err := p.client.ListJobsRequest(nil).Send(context.Background())
	if err != nil {
		return errors.Wrap(err, "listing jobs")
	}
	return nil
}

func (p *mcProvider) Capabilities() provider.Capabilities {
	return provider.Capabilities{
		InputFormats:  []string{"h264", "h265", "hdr10"},
		OutputFormats: []string{"mp4", "hls", "hdr10", "cmaf"},
		Destinations:  []string{"s3"},
	}
}

func mediaconvertFactory(cfg *config.Config) (provider.TranscodingProvider, error) {
	if cfg.MediaConvert.Endpoint == "" || cfg.MediaConvert.Queue == "" || cfg.MediaConvert.Role == "" {
		return nil, errors.New("incomplete MediaConvert config")
	}

	mcCfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return nil, errors.Wrap(err, "loading default aws config")
	}

	if cfg.MediaConvert.AccessKeyID+cfg.MediaConvert.SecretAccessKey != "" {
		mcCfg.Credentials = &aws.StaticCredentialsProvider{Value: aws.Credentials{
			AccessKeyID:     cfg.MediaConvert.AccessKeyID,
			SecretAccessKey: cfg.MediaConvert.SecretAccessKey,
		}}
	}

	if cfg.MediaConvert.Region != "" {
		mcCfg.Region = cfg.MediaConvert.Region
	}

	mcCfg.EndpointResolver = &aws.ResolveWithEndpoint{
		URL: cfg.MediaConvert.Endpoint,
	}

	dbRepo, err := redis.NewRepository(cfg)
	if err != nil {
		return nil, fmt.Errorf("error initializing mediaconvert wrapper: %s", err)
	}

	return &mcProvider{
		client:     mediaconvert.New(mcCfg),
		cfg:        cfg.MediaConvert,
		repository: dbRepo,
	}, nil
}
