package hybrik

import (
	"io"
	"net/url"

	"github.com/cbsinteractive/hybrik-sdk-go"
)

type testClient struct {
	getPresetReturn hybrik.Preset
}

func (testClient) CallAPI(method string, apiPath string, params url.Values, body io.Reader) (string, error) {
	return "", nil
}

func (testClient) QueueJob(string) (string, error)                   { return "", nil }
func (testClient) StopJob(string) error                              { return nil }
func (c *testClient) GetPreset(string) (hybrik.Preset, error)        { return c.getPresetReturn, nil }
func (testClient) GetJobInfo(string) (hybrik.JobInfo, error)         { return hybrik.JobInfo{}, nil }
func (testClient) CreatePreset(hybrik.Preset) (hybrik.Preset, error) { return hybrik.Preset{}, nil }
func (testClient) DeletePreset(string) error                         { return nil }
