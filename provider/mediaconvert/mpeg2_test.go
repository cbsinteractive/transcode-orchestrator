package mediaconvert

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/cbsinteractive/transcode-orchestrator/db"
)

func TestGenerateMPEG2(t *testing.T) {
	p := db.Preset{}
	p.Video.Codec = "hd422"
	p.Video.GopSize = 6000
	p.Video.Bitrate = 6000
	s, err := mpeg2XDCAM.generate(p)
	if err != nil {
		t.Fatalf("%v", err)
	}
	if *s.Mpeg2Settings.GopSize != 6000 || *s.Mpeg2Settings.Bitrate != 6000 {
		t.Fatalf("bad config: %+v", s)
	}

	p.Video.Profile = "hd4444444"
	s, err = mpeg2XDCAM.generate(p)
	if !errors.Is(err, ErrProfileUnsupported) {
		t.Fatalf("have %v want %v", err, ErrProfileUnsupported)
	}
}

func makemap(v interface{}) (m map[string]interface{}) {
	data := []byte{}
	switch v := v.(type) {
	case string:
		data = []byte(v)
	case []byte:
		data = v
	case interface{}:
		data, _ = json.Marshal(v)
	}
	json.Unmarshal(data, &m)
	return m
}
