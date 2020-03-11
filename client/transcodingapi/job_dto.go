package transcodingapi

import "time"

// JobID is our job's unique identifier
type JobID string

// CreateJobRequest contains the parameters needed to create a new transcode job
type CreateJobRequest struct {
	Source              string                      `json:"source"`
	Name                string                      `json:"name,omitempty"`
	DestinationBasePath string                      `json:"destinationBasePath,omitempty"`
	Outputs             []JobOutput                 `json:"outputs"`
	Provider            string                      `json:"provider"`
	StreamingParams     StreamingParams             `json:"streamingParams,omitempty"`
	SidecarAssets       map[SidecarAssetKind]string `json:"sidecarAssets,omitempty"`
	ExecutionFeatures   ExecutionFeatures           `json:"executionFeatures,omitempty"`
	ExecutionEnv        ExecutionEnvironment        `json:"executionEnv,omitempty"`
}

// ComputeClass represents a group of resources with similar capability
type ComputeClass = string

const (
	// ComputeClassTranscodeDefault runs any default transcodes
	ComputeClassTranscodeDefault ComputeClass = "transcodeDefault"

	// ComputeClassDolbyVisionTranscode runs Dolby Vision transcodes
	ComputeClassDolbyVisionTranscode ComputeClass = "doViTranscode"

	// ComputeClassDolbyVisionPreprocess runs Dolby Vision pre-processing
	ComputeClassDolbyVisionPreprocess ComputeClass = "doViPreprocess"
)

// ExecutionEnvironment contains configurations for the environment used while transcoding
type ExecutionEnvironment struct {
	Cloud            string                  `json:"cloud"`
	Region           string                  `json:"region"`
	ComputeTags      map[ComputeClass]string `json:"computeTags,omitempty"`
	CredentialsAlias string                  `json:"credentialsAlias,omitempty"`
	InputAlias       string                  `json:"inputAlias,omitempty"`
	OutputAlias      string                  `json:"outputAlias,omitempty"`
}

// ExecutionFeatures is a map whose key is a custom feature name and value is a json string
// representing the corresponding custom feature definition
type ExecutionFeatures map[string]interface{}

// SidecarAssetKind is the type of sidecar asset being defined
type SidecarAssetKind = string

// SidecarAssetKindDolbyVisionMetadata defines the dolby vision dynamic metadata location
const SidecarAssetKindDolbyVisionMetadata SidecarAssetKind = "dolbyVisionMetadata"

// JobOutput defines config parameters for single output in a job
type JobOutput struct {
	FileName string     `json:"fileName"`
	Preset   PresetName `json:"preset"`
}

// CreateJobResponse contains the results of new job request
type CreateJobResponse struct {
	JobID JobID `json:"jobId"`
}

// StreamingParams contains the configuration for media packaging
type StreamingParams struct {
	SegmentDuration  uint   `json:"segmentDuration"`
	Protocol         string `json:"protocol"`
	PlaylistFileName string `json:"playlistFileName,omitempty"`
}

// CancelJobRequest contains the parameters needed to cancel a transcode job
type CancelJobRequest struct {
	JobID JobID `json:"jobId"`
}

// CancelJobResponse contains the results of cancel job request
type CancelJobResponse struct {
	JobStatus
}

// JobStatusResponse contains the results of describe job request
type JobStatusResponse struct {
	JobStatus
}

// JobStatus is the representation of the status as the provider sees it.
type JobStatus struct {
	ProviderJobID  string                 `json:"providerJobId,omitempty"`
	Status         Status                 `json:"status,omitempty"`
	ProviderName   string                 `json:"providerName,omitempty"`
	StatusMessage  string                 `json:"statusMessage,omitempty"`
	Progress       float64                `json:"progress"`
	ProviderStatus map[string]interface{} `json:"providerStatus,omitempty"`
	Output         OutputFiles            `json:"output"`
	SourceInfo     SourceInfo             `json:"sourceInfo,omitempty"`
}

// SourceInfo contains information about media inputs
type SourceInfo struct {
	// Duration of the media
	Duration time.Duration `json:"duration,omitempty"`

	// Dimension of the media, in pixels
	Height int64 `json:"height,omitempty"`
	Width  int64 `json:"width,omitempty"`

	// Codec used for video medias
	VideoCodec string `json:"videoCodec,omitempty"`
}

// OutputFiles represents information about a job's outputs
type OutputFiles struct {
	Destination string       `json:"destination,omitempty"`
	Files       []OutputFile `json:"files,omitempty"`
}

// OutputFile represents an output file in a given job.
type OutputFile struct {
	Path       string `json:"path"`
	Container  string `json:"container"`
	VideoCodec string `json:"videoCodec"`
	Height     int64  `json:"height"`
	Width      int64  `json:"width"`
	FileSize   int64  `json:"fileSize"`
}

// Status is the status of a transcoding job.
type Status string
