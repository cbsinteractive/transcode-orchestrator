package codec

import (
	"testing"

	"github.com/cbsinteractive/transcode-orchestrator/db"
)

func TestCodec(t *testing.T) {
	p := db.Preset{Name: "test", Video: db.Video{Profile: "HIGH"}}
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
