package transcodingapi

import (
	"net/http"
	"net/url"
	"time"
)

// Client holds video-transcoding-api configuration and exposes
// methods for interacting with the transcoding service
type Client interface {
	// Jobs
	CreateJob(job CreateJobRequest) (CreateJobResponse, error)
	DescribeJob(jobID JobID) (JobStatusResponse, error)
	CancelJob(jobID JobID) (CancelJobResponse, error)

	// Presets
	CreatePreset(preset CreatePresetRequest) (CreatePresetResponse, error)
	DeletePreset(name PresetName) (DeletePresetResponse, error)

	// Providers
	AllProviders() (ProviderNames, error)
	GetProvider(name ProviderName) (ProviderDescription, error)
}

const (
	defaultTimeout = 30 * time.Second
	defaultBaseURL = "http://localhost:8080"
)

type DefaultClient struct {
	BaseURL *url.URL
	Client  *http.Client
}

// CreateJob creates a new transcode job based on the request definition
func (c *DefaultClient) CreateJob(job CreateJobRequest) (CreateJobResponse, error) {
	c.ensure()

	var jobResponse CreateJobResponse
	err := c.postResource(job, &jobResponse, "/jobs")
	if err != nil {
		return CreateJobResponse{}, err
	}

	return jobResponse, nil
}

// CancelJob will stop the execution of work in given provider
func (c *DefaultClient) CancelJob(jobID JobID) (CancelJobResponse, error) {
	c.ensure()

	var cancelResp CancelJobResponse
	err := c.postResource(nil, &cancelResp, "/jobs/"+string(jobID)+"/cancel")
	if err != nil {
		return CancelJobResponse{}, err
	}

	return cancelResp, nil
}

// DescribeJob returns details about a single job
func (c *DefaultClient) DescribeJob(jobID JobID) (JobStatusResponse, error) {
	c.ensure()

	var describeResp JobStatusResponse
	err := c.getResource(&describeResp, "/jobs/"+string(jobID))
	if err != nil {
		return JobStatusResponse{}, err
	}

	return describeResp, nil
}

// CreatePreset attempts to create a new preset based on the request definition
func (c *DefaultClient) CreatePreset(preset CreatePresetRequest) (CreatePresetResponse, error) {
	c.ensure()

	var presetResponse CreatePresetResponse
	err := c.postResource(preset, &presetResponse, "/presets")
	if err != nil {
		return CreatePresetResponse{}, err
	}

	return presetResponse, nil
}

// DeletePreset removes the preset from all providers
func (c *DefaultClient) DeletePreset(name PresetName) (DeletePresetResponse, error) {
	c.ensure()

	var deleteResponse DeletePresetResponse
	err := c.removeResource(&deleteResponse, "/presets/"+string(name))
	if err != nil {
		return DeletePresetResponse{}, err
	}

	return deleteResponse, nil
}

// AllProviders returns all configured providers
func (c *DefaultClient) AllProviders() (ProviderNames, error) {
	c.ensure()

	providerNames := ProviderNames{}
	err := c.getResource(&providerNames, "/providers")
	if err != nil {
		return providerNames, err
	}

	return providerNames, nil
}

// GetProvider returns information on a specific provider
func (c *DefaultClient) GetProvider(name ProviderName) (ProviderDescription, error) {
	c.ensure()

	var resp ProviderDescription
	err := c.getResource(&resp, "/providers/"+string(name))
	if err != nil {
		return ProviderDescription{}, err
	}

	return resp, nil
}

func (c *DefaultClient) ensure() {
	if c.Client == nil {
		c.Client = &http.Client{Timeout: defaultTimeout}
	}

	if c.BaseURL == nil {
		c.BaseURL = urlMust(url.Parse(defaultBaseURL))
	}
}

func urlMust(u *url.URL, _ error) *url.URL { return u }
