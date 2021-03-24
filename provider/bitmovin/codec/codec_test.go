package codec

import (
	"testing"

	"github.com/cbsinteractive/transcode-orchestrator/av"
)

func TestCodec(t *testing.T) {
	p := av.File{Name: "test", Video: av.Video{Profile: "HIGH"}}
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
