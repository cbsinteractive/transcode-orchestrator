package transcoding

import (
	"context"
	"net/http"
	"net/url"
	"time"
)

// Client holds video-transcoding-api configuration and exposes
// methods for interacting with the transcoding service
type Client interface {
	// Jobs
	CreateJob(ctx context.Context, job CreateJobRequest) (CreateJobResponse, error)
	DescribeJob(ctx context.Context, jobID JobID) (JobStatusResponse, error)
	CancelJob(ctx context.Context, jobID JobID) (CancelJobResponse, error)

	// Presets
	CreatePreset(ctx context.Context, preset CreatePresetRequest) (CreatePresetResponse, error)
	DeletePreset(ctx context.Context, name PresetName) (DeletePresetResponse, error)

	// Providers
	AllProviders(ctx context.Context) (ProviderNames, error)
	GetProvider(ctx context.Context, name ProviderName) (ProviderDescription, error)
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
func (c *DefaultClient) CreateJob(ctx context.Context, job CreateJobRequest) (CreateJobResponse, error) {
	c.ensure()

	var jobResponse CreateJobResponse
	err := c.postResource(ctx, job, &jobResponse, "/jobs")
	if err != nil {
		return CreateJobResponse{}, err
	}

	return jobResponse, nil
}

// CancelJob will stop the execution of work in given provider
func (c *DefaultClient) CancelJob(ctx context.Context, jobID JobID) (CancelJobResponse, error) {
	c.ensure()

	var cancelResp CancelJobResponse
	err := c.postResource(ctx, nil, &cancelResp, "/jobs/"+string(jobID)+"/cancel")
	if err != nil {
		return CancelJobResponse{}, err
	}

	return cancelResp, nil
}

// DescribeJob returns details about a single job
func (c *DefaultClient) DescribeJob(ctx context.Context, jobID JobID) (JobStatusResponse, error) {
	c.ensure()

	var describeResp JobStatusResponse
	err := c.getResource(ctx, &describeResp, "/jobs/"+string(jobID))
	if err != nil {
		return JobStatusResponse{}, err
	}

	return describeResp, nil
}

// CreatePreset attempts to create a new preset based on the request definition
func (c *DefaultClient) CreatePreset(ctx context.Context, preset CreatePresetRequest) (CreatePresetResponse, error) {
	c.ensure()

	var presetResponse CreatePresetResponse
	err := c.postResource(ctx, preset, &presetResponse, "/presets")
	if err != nil {
		return CreatePresetResponse{}, err
	}

	return presetResponse, nil
}

// DeletePreset removes the preset from all providers
func (c *DefaultClient) DeletePreset(ctx context.Context, name PresetName) (DeletePresetResponse, error) {
	c.ensure()

	var deleteResponse DeletePresetResponse
	err := c.removeResource(ctx, &deleteResponse, "/presets/"+string(name))
	if err != nil {
		return DeletePresetResponse{}, err
	}

	return deleteResponse, nil
}

// AllProviders returns all configured providers
func (c *DefaultClient) AllProviders(ctx context.Context) (ProviderNames, error) {
	c.ensure()

	providerNames := ProviderNames{}
	err := c.getResource(ctx, &providerNames, "/providers")
	if err != nil {
		return providerNames, err
	}

	return providerNames, nil
}

// GetProvider returns information on a specific provider
func (c *DefaultClient) GetProvider(ctx context.Context, name ProviderName) (ProviderDescription, error) {
	c.ensure()

	var resp ProviderDescription
	err := c.getResource(ctx, &resp, "/providers/"+string(name))
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
