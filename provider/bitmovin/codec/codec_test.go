package codec

import (
	"testing"

	"github.com/cbsinteractive/transcode-orchestrator/job"
)

func TestCodec(t *testing.T) {
	p := job.Preset{Name: "test", Video: job.Video{Profile: "HIGH"}}
	c := CodecH264{}
	c.set(p)
	if c.Err() != nil {
		t.Fatalf("unexpected error: %v", c.Err())
	}
	if c.cfg.Profile != "HIGH" {
		t.Logf("%#v", c)
		t.Fatalf("h264: default HIGH profile not applied")
	}
}
