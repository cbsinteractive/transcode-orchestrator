package transcoding

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/cbsinteractive/transcode-orchestrator/job"
)

type (
	Status = job.Status
	Job    = job.Job
)

const (
	defaultTimeout = 30 * time.Second
	defaultURL     = "http://localhost:8080"
)

type Client struct {
	Base   *url.URL
	Client *http.Client
}

// Create a job
func (c *Client) Create(ctx context.Context, job Job) (r Status, err error) {
	c.ensure()
	return r, c.postResource(ctx, job, &r, "/jobs")
}

// Cancel a job
func (c *Client) Cancel(ctx context.Context, id string) (r Status, err error) {
	c.ensure()
	return r, c.postResource(ctx, nil, &r, "/jobs/"+id+"/cancel")
}

// DescribeJob returns details about a single job
func (c *Client) Status(ctx context.Context, id string) (r Status, err error) {
	c.ensure()
	return r, c.postResource(ctx, nil, &r, "/jobs/"+id)
}

func (c *Client) ensure() {
	if c.Client == nil {
		c.Client = &http.Client{Timeout: defaultTimeout}
	}
	if c.Base == nil {
		c.Base = urlMust(url.Parse(defaultURL))
	}
}

func urlMust(u *url.URL, _ error) *url.URL { return u }
