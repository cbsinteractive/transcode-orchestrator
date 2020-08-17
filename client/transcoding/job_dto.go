package transcoding

import (
	"time"

	"github.com/cbsinteractive/pkg/timecode"
)

// ScanType is a string that represents the scan type of the content.
type ScanType string

const (
	ScanTypeUnknown     ScanType = "unknown"
	ScanTypeProgressive ScanType = "progressive"
	ScanTypeInterlaced  ScanType = "interlaced"
)

// ComputeClass represents a group of resources with similar capability
// TODO(as): This is a type alias, we shouldn't be using it.
type ComputeClass = string

const (
	ComputeClassTranscodeDefault      ComputeClass = "transcodeDefault" // any default transcodes
	ComputeClassDolbyVisionTranscode  ComputeClass = "doViTranscode"    // Dolby Vision transcodes
	ComputeClassDolbyVisionPreprocess ComputeClass = "doViPreprocess"   // Dolby Vision pre-processing
)

// SidecarAssetKindDolbyVisionMetadata defines the dolby vision dynamic metadata location
const SidecarAssetKindDolbyVisionMetadata SidecarAssetKind = "dolbyVisionMetadata"

type (
	// JobID is our job's unique identifier
	JobID string

	// Status is the status of a transcoding job.
	Status string

	// ExecutionFeatures is a map whose key is a custom feature name and value is a json string
	// representing the corresponding custom feature definition
	ExecutionFeatures map[string]interface{}

	// SidecarAssetKind is the type of sidecar asset being defined
	// TODO(as): This is a type alias, we shouldn't be using it.
	SidecarAssetKind = string
)

//AudioChannel describes and Audio Mix
type AudioChannel struct {
	TrackIdx   int
	ChannelIdx int
	Layout     string
}

//AudioDownmix holds source and output channels layouts for providers
//to handle downmixing
type AudioDownmix struct {
	SrcChannels  []AudioChannel
	DestChannels []AudioChannel
}

// File is a media file. It replaces the following objects
// SourceInfo: Duration, Height, Width, Codec
// CreateJobSourceInfo: Height, Width, FrameRate, File Size, ScanType
// SourceInfo:
type File struct {
	Path       string        `json:"path"`
	Size       int64         `json:"fileSize"`
	Container  string        `json:"container"`
	Duration   time.Duration `json:"duration,omitempty"`
	VideoCodec string        `json:"videoCodec,omitempty"`
	Width      int           `json:"width,omitempty"`
	Height     int           `json:"height,omitempty"`
	FrameRate  float64       `json:"frameRate,omitempty"`
	ScanType   ScanType      `json:"scanType,omitempty"`
}

type (
	// CreateJobRequest and similar data structures describe the requests and replies
	// of the Job management endpoints.
	CreateJobRequest struct {
		Name       string `json:"name,omitempty"`
		Source     string `json:"source"`
		SourceInfo File   `json:"sourceInfo,omitempty"`

		// Splice is a request to cut the source before processing. It falls somewhere
		// between file metadata and a request by the user to operate on the source.
		// Not every provider currently supports this feature.
		Splice timecode.Splice `json:"splice,omitempty"`

		Provider          string                      `json:"provider"`
		ExecutionFeatures ExecutionFeatures           `json:"executionFeatures,omitempty"`
		ExecutionEnv      ExecutionEnvironment        `json:"executionEnv,omitempty"`
		StreamingParams   StreamingParams             `json:"streamingParams,omitempty"`
		SidecarAssets     map[SidecarAssetKind]string `json:"sidecarAssets,omitempty"`

		DestinationBasePath string      `json:"destinationBasePath,omitempty"`
		Outputs             []JobOutput `json:"outputs"`

		AudioDownmix            AudioDownmix `json:"audioDownmix"`
		ExplicitKeyframeOffsets []float64    `json:"explicitKeyframeOffsets,omitempty"`
	}
	CreateJobResponse struct {
		JobID JobID `json:"jobId"`
	}
	CancelJobRequest struct {
		JobID JobID `json:"jobId"`
	}
	CancelJobResponse struct{ JobStatus }
)

// StreamingParams contains the configuration for media packaging
type StreamingParams struct {
	SegmentDuration  uint   `json:"segmentDuration"`
	Protocol         string `json:"protocol"`
	PlaylistFileName string `json:"playlistFileName,omitempty"`
}

// ExecutionEnvironment contains configurations for the environment used while transcoding
type ExecutionEnvironment struct {
	Cloud            string                  `json:"cloud"`
	Region           string                  `json:"region"`
	ComputeTags      map[ComputeClass]string `json:"computeTags,omitempty"`
	CredentialsAlias string                  `json:"credentialsAlias,omitempty"`
	InputAlias       string                  `json:"inputAlias,omitempty"`
	OutputAlias      string                  `json:"outputAlias,omitempty"`
}

// JobStatus is the representation of the status as the provider sees it.
type JobStatus struct {
	Progress      float64 `json:"progress"`
	Status        Status  `json:"status,omitempty"`
	StatusMessage string  `json:"statusMessage,omitempty"`

	ProviderJobID  string                 `json:"providerJobId,omitempty"`
	ProviderName   string                 `json:"providerName,omitempty"`
	ProviderStatus map[string]interface{} `json:"providerStatus,omitempty"`

	SourceInfo File `json:"sourceInfo,omitempty"`

	Output OutputFiles `json:"output"`
}

// JobOutput defines config parameters for single output in a job
type JobOutput struct {
	FileName string     `json:"fileName"`
	Preset   PresetName `json:"preset"`
}

// JobStatusResponse contains the results of describe job request
type JobStatusResponse struct {
	JobStatus
}

// OutputFiles represents information about a job's outputs
type OutputFiles struct {
	Destination string `json:"destination,omitempty"`
	Files       []File `json:"files,omitempty"`
}
