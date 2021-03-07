package job

import "github.com/cbsinteractive/pkg/video"

// ScanProgressive and other supported types
const (
	ScanProgressive = "progressive"
	ScanInterlaced  = "interlaced"
	ScanUnknown     = "unknown"
)

// Video transcoding parameters
type Video struct {
	Codec   string `json:"codec,omitempty"`
	Profile string `json:"profile,omitempty"`
	Level   string `json:"level,omitempty"`

	Width    int    `json:"width,omitempty"`
	Height   int    `json:"height,omitempty"`
	Scantype string `json:"scantype,omitempty"`

	FPS     float64 `json:"fps,omitempty"`
	Bitrate Bitrate `json:"bitrate"`
	Gop     Gop     `json:"gop"`

	HDR10       HDR10       `json:"hdr10"`
	DolbyVision DolbyVision `json:"dolbyVision"`
	Overlays    Overlays    `json:"overlays,omitempty"`
	Crop        video.Crop  `json:"crop"`
}

func (v *Video) On() bool {
	return v != nil && !(v.Codec == "" && v.Height == 0 && v.Width == 0)
}

type Bitrate struct {
	BPS     int    `json:"bps"`
	Control string `json:"control"`
	TwoPass bool   `json:"twopass"`
}

// Percent adjusts the bitrate by n percent
// where n is a number in the range [-100, +100]
func (b Bitrate) Percent(n int) Bitrate {
	// operate on bits to keep precision
	b.BPS = b.BPS * (100 + n) / 100
	return b
}

func (b Bitrate) Kbps() int {
	return b.BPS / 1000
}

type Gop struct {
	Unit string  `json:"unit,omitempty"`
	Size float64 `json:"size,omitempty"`
	Mode string  `json:"mode,omitempty"`
}

func (g Gop) Seconds() bool {
	return g.Unit == "seconds"
}
