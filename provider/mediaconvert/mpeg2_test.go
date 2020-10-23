package mediaconvert

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestGenerateMPEG2(t *testing.T) {
	s, err := mpeg2XDCAM.generate()
	t.Log(s, s.Validate(), err)

	if false{

	want := makemap(`
{
                  "InterlaceMode": "TOP_FIELD",
                  "Syntax": "DEFAULT",
                  "GopClosedCadence": 1,
                  "GopSize": 12,
                  "SlowPal": "DISABLED",
                  "SpatialAdaptiveQuantization": "ENABLED",
                  "TemporalAdaptiveQuantization": "ENABLED",
                  "Bitrate": 50000000,
                  "IntraDcPrecision": "AUTO",
                  "FramerateControl": "INITIALIZE_FROM_SOURCE",
                  "RateControlMode": "CBR",
                  "CodecProfile": "PROFILE_422",
                  "Telecine": "NONE",
                  "MinIInterval": 0,
                  "AdaptiveQuantization": "HIGH",
                  "CodecLevel": "HIGH",
                  "SceneChangeDetect": "ENABLED",
                  "QualityTuningLevel": "SINGLE_PASS",
                  "FramerateConversionAlgorithm": "DUPLICATE_DROP",
                  "GopSizeUnits": "FRAMES",
                  "ParControl": "INITIALIZE_FROM_SOURCE",
                  "NumberBFramesBetweenReferenceFrames": 2,
                  "DynamicSubGop": "STATIC"
}
`)
	have := makemap(s.Mpeg2Settings)

	if !reflect.DeepEqual(have, want) {
		t.Fatalf("not equal:\n		have: %+v\n		want: %+v\n", have, want)
	}
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
