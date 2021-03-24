package transcoding

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/cbsinteractive/transcode-orchestrator/av"
)

type (
	Status = av.Status
	Job    = av.Job
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
func (c *Client) Create(ctx context.Context, job av.Job) (r Status, err error) {
	return r, c.do(ctx, "POST", "/job", job, &r)
}

// Status for the job id
func (c *Client) Status(ctx context.Context, id string) (r Status, err error) {
	return r, c.do(ctx, "GET", "/job/"+id, nil, &r)
}

// Cancel a job
func (c *Client) Cancel(ctx context.Context, id string) (r Status, err error) {
	return r, c.do(ctx, "DELETE", "/job/"+id, nil, &r)
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
