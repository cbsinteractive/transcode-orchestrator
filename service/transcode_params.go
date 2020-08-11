package service

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/cbsinteractive/transcode-orchestrator/db"
	"github.com/cbsinteractive/transcode-orchestrator/provider"
)

// NewTranscodeJobInputPayload makes up the parameters available for
// specifying a new transcoding job
type NewTranscodeJobInputPayload struct {
	// source media for the transcoding job.
	Source string `json:"source"`

	// SourceInfo is an optional param allowing users to add helpful information about the source content
	SourceInfo db.File `json:"sourceInfo,omitempty"`

	// list of outputs in this job
	Outputs []struct {
		FileName string `json:"fileName"`
		Preset   string `json:"preset"`
	} `json:"outputs"`

	// provider to use in this job
	Provider string `json:"provider"`

	// Name is an optional client-supplied name for the job
	Name string `json:"name,omitempty"`

	// ExecutionFeatures holds custom provider-specific execution features configuration
	ExecutionFeatures db.ExecutionFeatures `json:"executionFeatures,omitempty"`

	// ExecutionEnv contains configurations for the environment used while transcoding
	ExecutionEnv db.ExecutionEnvironment `json:"executionEnv,omitempty"`

	// SidecarAssets holds the location of any sidecar assets associated with the transcode
	SidecarAssets map[db.SidecarAssetKind]string `json:"sidecarAssets,omitempty"`

	// DestinationBasePath is an optional param allowing users to overwrite the location of outputs pre job
	DestinationBasePath string `json:"destinationBasePath,omitempty"`

	// provider Adaptive Streaming parameters
	StreamingParams db.StreamingParams `json:"streamingParams,omitempty"`

	// job labels
	Labels []string `json:"labels,omitempty"`
}

// swagger:parameters newJob
type newTranscodeJobInput struct {
	// in: body
	// required: true
	Payload NewTranscodeJobInputPayload
}

// ProviderFactory loads and validates the parameters, and then returns the
// provider factory.
func (p *newTranscodeJobInput) ProviderFactory(body io.Reader) (provider.Factory, error) {
	err := p.loadParams(body)
	if err != nil {
		return nil, err
	}
	err = p.validate()
	if err != nil {
		return nil, err
	}
	return provider.GetProviderFactory(p.Payload.Provider)
}

func (p *newTranscodeJobInput) loadParams(body io.Reader) error {
	return json.NewDecoder(body).Decode(&p.Payload)
}

func (p *newTranscodeJobInput) validate() error {
	if p.Payload.Provider == "" {
		return errors.New("missing provider from request")
	}
	if p.Payload.Source == "" {
		return errors.New("missing source media from request")
	}
	if len(p.Payload.Outputs) == 0 {
		return errors.New("missing output list from request")
	}
	return nil
}

// swagger:parameters getJob
type getTranscodeJobInput struct {
	// in: path
	// required: true
	JobID string `json:"jobId"`
}

func (p *getTranscodeJobInput) loadParams(paramsMap map[string]string) {
	p.JobID = paramsMap["jobId"]
}

// swagger:parameters cancelJob
type cancelTranscodeJobInput struct {
	getTranscodeJobInput
}
