package mediaconvert

import (
	"context"
	"fmt"
	"path"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/mediaconvert"
	"github.com/cbsinteractive/pkg/timecode"
	"github.com/cbsinteractive/transcode-orchestrator/config"
	"github.com/cbsinteractive/transcode-orchestrator/db"
	"github.com/cbsinteractive/transcode-orchestrator/db/redis"
	"github.com/cbsinteractive/transcode-orchestrator/provider"
	"github.com/pkg/errors"
)

const (
	// Name identifies the MediaConvert provider by name
	Name = "mediaconvert"

	defaultAudioSampleRate     = 48000
	defaultQueueHopTimeoutMins = 1
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

func splice2clippings(s timecode.Splice, fps float64) (ic []mediaconvert.InputClipping) {
	// NOTE(as): While this could be a helper function in the time/timecode package
	// we probably don't want the uglyness of importing the AWS API in that package
	// and having to recognize mediaconvert.InputClippings

	// NOTE(as): We need to take into account embedded timecodes. Maybe it would
	// be better to have this be a method on a timecode object or have it passed in as
	// a reference argument (object could also provide fps info)

	for _, r := range s {
		s, e := r.Timecodes(fps)
		ic = append(ic, mediaconvert.InputClipping{
			StartTimecode: &s,
			EndTimecode:   &e,
		})
	}
	return ic
}

func (p *mcProvider) Transcode(ctx context.Context, job *db.Job) (*provider.JobStatus, error) {
	outputGroups, err := p.outputGroupsFrom(ctx, job)
	if err != nil {
		return nil, fmt.Errorf("mediaconvert: output group generator: %w", err)
	}

	queue := aws.String(p.cfg.DefaultQueueARN)

	var hopDestinations []mediaconvert.HopDestination
	if preferred := p.cfg.PreferredQueueARN; p.canUsePreferredQueue(job.SourceInfo) && preferred != "" {
		queue = aws.String(preferred)
		hopDestinations = append(hopDestinations, mediaconvert.HopDestination{
			WaitMinutes: aws.Int64(defaultQueueHopTimeoutMins),
		})
	}

	var accelerationSettings *mediaconvert.AccelerationSettings
	if p.requiresAcceleration(job.SourceInfo) {
		accelerationSettings = &mediaconvert.AccelerationSettings{
			Mode: mediaconvert.AccelerationModePreferred,
		}
	}

	audioSelector := mediaconvert.AudioSelector{
		DefaultSelection: mediaconvert.AudioDefaultSelectionDefault,
	}
	if job.AudioDownmix != nil {
		if err = audioSelectorFrom(job.AudioDownmix, &audioSelector); err != nil {
			return nil, fmt.Errorf("mediaconvert: audio selectors generator: %w", err)
		}
	}

	resp, err := p.client.CreateJobRequest(&mediaconvert.CreateJobInput{
		AccelerationSettings: accelerationSettings,
		Queue:                queue,
		HopDestinations:      hopDestinations,
		Role:                 aws.String(p.cfg.Role),
		Settings: &mediaconvert.JobSettings{
			Inputs: []mediaconvert.Input{
				{
					InputClippings: splice2clippings(job.SourceSplice, 0), // TODO(as): Find FPS in job
					FileInput:      aws.String(job.SourceMedia),
					AudioSelectors: map[string]mediaconvert.AudioSelector{
						"Audio Selector 1": audioSelector,
					},
					VideoSelector: &mediaconvert.VideoSelector{
						ColorSpace: mediaconvert.ColorSpaceFollow,
					},
					TimecodeSource: mediaconvert.InputTimecodeSourceZerobased,
				},
			},
			OutputGroups: outputGroups,
			TimecodeConfig: &mediaconvert.TimecodeConfig{
				Source: mediaconvert.TimecodeSourceZerobased,
			},
		},
		Tags:              p.tagsFrom(job.Labels),
		BillingTagsSource: "JOB",
	}).Send(ctx)
	if err != nil {
		return nil, err
	}

	return &provider.JobStatus{
		ProviderName:  Name,
		ProviderJobID: aws.StringValue(resp.Job.Id),
		Status:        provider.StatusQueued,
	}, nil
}

var defaultDestinationSettings = mediaconvert.DestinationSettings{
	S3Settings: &mediaconvert.S3DestinationSettings{
		AccessControl: &mediaconvert.S3DestinationAccessControl{
			CannedAcl: mediaconvert.S3ObjectCannedAclBucketOwnerFullControl,
		},
	},
}

func (p *mcProvider) outputGroupsFrom(ctx context.Context, job *db.Job) ([]mediaconvert.OutputGroup, error) {
	outputGroups := map[mediaconvert.ContainerType][]outputCfg{}
	for _, output := range job.Outputs {
		presetName := output.Preset.Name
		presetResponse, err := p.GetPreset(ctx, presetName)
		if err != nil {
			return nil, err
		}

		localPreset, ok := presetResponse.(*db.LocalPreset)
		if !ok {
			return nil, fmt.Errorf("could not convert preset response into a db.LocalPreset")
		}

		mcOutput, err := outputFrom(localPreset.Preset, job.SourceInfo)
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

		destination := p.destinationPathFrom(job)

		switch container {
		case mediaconvert.ContainerTypeCmfc:
			mcOutputGroup.OutputGroupSettings = &mediaconvert.OutputGroupSettings{
				Type: mediaconvert.OutputGroupTypeCmafGroupSettings,
				CmafGroupSettings: &mediaconvert.CmafGroupSettings{
					Destination:            aws.String(destination),
					DestinationSettings:    &defaultDestinationSettings,
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
					DestinationSettings:    &defaultDestinationSettings,
					SegmentLength:          aws.Int64(int64(job.StreamingParams.SegmentDuration)),
					MinSegmentLength:       aws.Int64(0),
					DirectoryStructure:     mediaconvert.HlsDirectoryStructureSingleDirectory,
					ManifestDurationFormat: mediaconvert.HlsManifestDurationFormatFloatingPoint,
					OutputSelection:        mediaconvert.HlsOutputSelectionManifestsAndSegments,
					SegmentControl:         mediaconvert.HlsSegmentControlSegmentedFiles,
				},
			}
		case mediaconvert.ContainerTypeMp4, mediaconvert.ContainerTypeMov, mediaconvert.ContainerTypeWebm, mediaconvert.ContainerTypeMxf:
			mcOutputGroup.OutputGroupSettings = &mediaconvert.OutputGroupSettings{
				Type: mediaconvert.OutputGroupTypeFileGroupSettings,
				FileGroupSettings: &mediaconvert.FileGroupSettings{
					Destination:         aws.String(destination + "m"),
					DestinationSettings: &defaultDestinationSettings,
				},
			}
		default:
			return nil, fmt.Errorf("container %s is not yet supported with mediaconvert", string(container))
		}

		mcOutputGroups = append(mcOutputGroups, mcOutputGroup)
	}

	return mcOutputGroups, nil
}

func (p *mcProvider) destinationPathFrom(job *db.Job) string {
	var basePath string
	if cfgBasePath := job.DestinationBasePath; cfgBasePath != "" {
		basePath = cfgBasePath
	} else {
		basePath = p.cfg.Destination
	}
	return fmt.Sprintf("%s/%s/", strings.TrimRight(basePath, "/"), job.RootFolder())
}

func (p *mcProvider) CreatePreset(_ context.Context, preset db.Preset) (string, error) {
	err := p.repository.CreateLocalPreset(&db.LocalPreset{
		Name:   preset.Name,
		Preset: preset,
	})
	if err != nil {
		return "", err
	}

	return preset.Name, nil
}

func (p *mcProvider) GetPreset(_ context.Context, presetID string) (interface{}, error) {
	return p.repository.GetLocalPreset(presetID)
}

func (p *mcProvider) DeletePreset(ctx context.Context, presetID string) error {
	preset, err := p.GetPreset(ctx, presetID)
	if err != nil {
		return err
	}

	return p.repository.DeleteLocalPreset(preset.(*db.LocalPreset))
}

func (p *mcProvider) JobStatus(ctx context.Context, job *db.Job) (*provider.JobStatus, error) {
	jobResp, err := p.client.GetJobRequest(&mediaconvert.GetJobInput{
		Id: aws.String(job.ProviderJobID),
	}).Send(ctx)
	if err != nil {
		return &provider.JobStatus{}, errors.Wrap(err, "fetching job info with the mediaconvert API")
	}

	return p.jobStatusFrom(job.ProviderJobID, job, jobResp.Job), nil
}

func (p *mcProvider) jobStatusFrom(providerJobID string, job *db.Job, mcJob *mediaconvert.Job) *provider.JobStatus {
	status := &provider.JobStatus{
		ProviderJobID: providerJobID,
		ProviderName:  Name,
		Status:        providerStatusFrom(mcJob.Status),
		StatusMessage: statusMsgFrom(mcJob),
		Output: provider.JobOutput{
			Destination: p.destinationPathFrom(job),
		},
	}

	if status.Status == provider.StatusFinished {
		status.Progress = 100
	} else if p := mcJob.JobPercentComplete; p != nil {
		status.Progress = float64(*p)
	}

	var files []provider.OutputFile
	if settings := mcJob.Settings; settings != nil {
		for _, group := range settings.OutputGroups {
			groupDestination, err := outputGroupDestinationFrom(group)
			if err != nil {
				continue
			}
			for _, output := range group.Outputs {
				file := provider.OutputFile{}

				if modifier := output.NameModifier; modifier != nil {
					if extension, err := fileExtensionFromContainer(output.ContainerSettings); err == nil {
						file.Path = groupDestination + *modifier + extension
					} else {
						continue
					}
				} else {
					continue
				}

				if video := output.VideoDescription; video != nil {
					if height := video.Height; height != nil {
						file.Height = *height
					}

					if width := video.Width; width != nil {
						file.Width = *width
					}
				}

				if container, err := containerIdentifierFrom(output.ContainerSettings); err == nil {
					file.Container = container
				}

				files = append(files, file)
			}
		}
	}

	status.Output.Files = files

	return status
}

func outputGroupDestinationFrom(group mediaconvert.OutputGroup) (string, error) {
	if group.OutputGroupSettings == nil {
		return "", errors.New("output group contained no settings")
	}

	switch group.OutputGroupSettings.Type {
	case mediaconvert.OutputGroupTypeFileGroupSettings:
		fsSettings := group.OutputGroupSettings.FileGroupSettings
		if fsSettings == nil {
			return "", errors.New("file group settings were nil")
		}

		if fsSettings.Destination == nil {
			return "", errors.New("file group destination was nil")
		}

		return *fsSettings.Destination, nil
	default:
		return "", fmt.Errorf("output enumeration not supported for output group %q",
			group.OutputGroupSettings.Type)
	}
}

func fileExtensionFromContainer(settings *mediaconvert.ContainerSettings) (string, error) {
	if settings == nil {
		return "", errors.New("container settings were nil")
	}

	switch settings.Container {
	case mediaconvert.ContainerTypeMp4:
		return ".mp4", nil
	case mediaconvert.ContainerTypeMov:
		return ".mov", nil
	case mediaconvert.ContainerTypeWebm:
		return ".webm", nil
	default:
		return "", fmt.Errorf("could not determine extension from output container %q", settings.Container)
	}
}

func containerIdentifierFrom(settings *mediaconvert.ContainerSettings) (string, error) {
	if settings == nil {
		return "", errors.New("container settings were nil")
	}

	switch settings.Container {
	case mediaconvert.ContainerTypeMp4:
		return "mp4", nil
	case mediaconvert.ContainerTypeMov:
		return "mov", nil
	case mediaconvert.ContainerTypeWebm:
		return "webm", nil
	default:
		return "", fmt.Errorf("could not determine container identifier from output container %q", settings.Container)
	}
}

func statusMsgFrom(job *mediaconvert.Job) string {
	if job.ErrorMessage != nil {
		return *job.ErrorMessage
	}

	return string(job.CurrentPhase)
}

func (p *mcProvider) CancelJob(ctx context.Context, id string) error {
	_, err := p.client.CancelJobRequest(&mediaconvert.CancelJobInput{
		Id: aws.String(id),
	}).Send(ctx)

	return err
}

func (p *mcProvider) Healthcheck() error {
	_, err := p.client.ListJobsRequest(nil).Send(context.Background()) // TODO(as): plump context
	if err != nil {
		return errors.Wrap(err, "listing jobs")
	}
	return nil
}

func (p *mcProvider) Capabilities() provider.Capabilities {
	return provider.Capabilities{
		InputFormats:  []string{"h264", "h265", "hdr10"},
		OutputFormats: []string{"mp4", "hls", "hdr10", "cmaf", "mov"},
		Destinations:  []string{"s3"},
	}
}

func (p *mcProvider) tagsFrom(labels []string) map[string]string {
	tags := make(map[string]string)

	for _, label := range labels {
		if strings.HasPrefix(label, "bill:") {
			// in case of billing tag we need to use a Cost Allocation Tag
			// for cost to be reported correctly by Cost Explorer API
			tags["bu"] = label
		} else {
			tags[label] = "true"
		}
	}

	return tags
}

func mediaconvertFactory(cfg *config.Config) (provider.TranscodingProvider, error) {
	if cfg.MediaConvert.Endpoint == "" || cfg.MediaConvert.DefaultQueueARN == "" || cfg.MediaConvert.Role == "" {
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
