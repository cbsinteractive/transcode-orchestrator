package transcoding

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/cbsinteractive/transcode-orchestrator/db"
)

type (
	Status = transcode.Status
	Job    = job.Job
)

const (
	defaultTimeout = 30 * time.Second
	defaultBaseURL = "http://localhost:8080"
)

type Client struct {
	Base   *url.URL
	Client *http.Client
}

// Create a job
func (c *Client) Create(ctx context.Context, job job.Job) (r Status, err error) {
	c.ensure()
	return r, c.postResource(ctx, job, &r, "/jobs")
}

// Cancel a job
func (c *Client) Cancel(ctx context.Context, jobID JobID) (Status, error) {
	c.ensure()

	var cancelResp CancelJobResponse
	err := c.postResource(ctx, nil, &cancelResp, "/jobs/"+string(jobID)+"/cancel")
	if err != nil {
		return CancelJobResponse{}, err
	}

	return cancelResp, nil
}

// DescribeJob returns details about a single job
func (c *Client) Status(ctx context.Context, jobID JobID) (Status, error) {
	c.ensure()

	var describeResp JobStatusResponse
	err := c.getResource(ctx, &describeResp, "/jobs/"+string(jobID))
	if err != nil {
		return JobStatusResponse{}, err
	}

	return describeResp, nil
}

func (c *Client) ensure() {
	if c.Client == nil {
		c.Client = &http.Client{Timeout: defaultTimeout}
	}

	if c.BaseURL == nil {
		c.BaseURL = urlMust(url.Parse(defaultBaseURL))
	}
}

func urlMust(u *url.URL, _ error) *url.URL { return u }
