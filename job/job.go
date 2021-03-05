package job

import (
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/cbsinteractive/pkg/timecode"
	"github.com/cbsinteractive/pkg/video"
	"github.com/gofrs/uuid"
)

// Job is a transcoding job
type Job struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt time.Time
	Labels    []string

	Provider      string `json:"provider"`
	ProviderJobID string

	Input  File
	Output Dir

	Streaming Streaming

	Features Features
	Env      Env

	ExtraFiles map[string]string
}

func (j *Job) Asset(sidecar string) *File {
	loc := j.ExtraFiles[sidecar]
	if loc == "" {
		return nil
	}
	return &File{Name: loc}
}

// State is the state of a transcoding job.
type State string

const (
	StateUnknown  = State("unknown")
	StateQueued   = State("queued")
	StateStarted  = State("started")
	StateFinished = State("finished")
	StateFailed   = State("failed")
	StateCanceled = State("canceled")
)

type Provider struct {
	Name   string                 `json:"name,omitempty"`
	JobID  string                 `json:"job_id,omitempty"`
	Status map[string]interface{} `json:"status,omitempty"`
}

// Status is the representation of the status
type Status struct {
	ID     string   `json:"jobID,omitempty"`
	Labels []string `json:"labels,omitempty"`

	State    State   `json:"status,omitempty"`
	Msg      string  `json:"msg,omitempty"`
	Progress float64 `json:"progress"`

	Input  File `json:"input"`
	Output Dir  `json:"output"`

	ProviderName   string                 `json:"providerName,omitempty"`
	ProviderJobID  string                 `json:"providerJobId,omitempty"`
	ProviderStatus map[string]interface{} `json:"providerStatus,omitempty"`
}

// Dir is a named directory of files
type Dir struct {
	Path string `json:"path,omitempty"`
	File []File `json:"files,omitempty"`
}

func (d *Dir) Len() int {
	return len(d.File)
}

func (d *Dir) Add(f ...File) {
	d.File = append(d.File, f...)
}

func (d Dir) Location() url.URL {
	u, _ := url.Parse(d.Path)
	if u == nil {
		return url.URL{}
	}
	return *u
}

func (j *Job) Dir() File {
	return File{Name: j.Location("")}
}
func (j Job) Location(file string) string {
	u := j.Output.Location()
	u.Path = path.Join(u.Path, j.rootFolder(), file)
	return u.String()
}
func (j *Job) Abs(f File) File {
	f.Name = j.Location(f.Name)
	return f
}
func (j Job) rootFolder() string {
	if j.Name != "" {
		if _, err := uuid.FromString(j.Name); err == nil {
			return j.Name
		}
	}
	return j.ID
}

const DolbyVisionMetadata = "dolbyVisionMetadata"

// Env contains configurations for the environment used while transcoding
type Env struct {
	Cloud       string
	Region      string
	InputAlias  string
	OutputAlias string

	Tags map[string]string
}

// TagTranscodeDefault runs any default transcodes
// TagDolbyVisionTranscode runs Dolby Vision transcodes
// TagDolbyVisionPreprocess runs Dolby Vision pre-processing
// TagDolbyVisionMezzQC runs QC check on the mezzanine
const (
	TagTranscodeDefault      = "transcodeDefault"
	TagDolbyVisionTranscode  = "doViTranscode"
	TagDolbyVisionPreprocess = "doViPreprocess"
	TagDolbyVisionMezzQC     = "doViMezzQC"
	TagDolbyVisionMetadata   = "dolbyVisionMetadata" // inconsistent
)

// Streaming configures Adaptive Streaming jobs
type Streaming struct {
	SegmentDuration  uint
	Protocol         string
	PlaylistFileName string
}

// ScanProgressive and other supported types
const (
	ScanProgressive = "progressive"
	ScanInterlaced  = "interlaced"
	ScanUnknown     = "unknown"
)

//ChannelLayout describes layout of an audio channel
type ChannelLayout string

const (
	ChannelLayoutCenter        ChannelLayout = "C"
	ChannelLayoutLeft          ChannelLayout = "L"
	ChannelLayoutRight         ChannelLayout = "R"
	ChannelLayoutLeftSurround  ChannelLayout = "Ls"
	ChannelLayoutRightSurround ChannelLayout = "Rs"
	ChannelLayoutLeftBack      ChannelLayout = "Lb"
	ChannelLayoutRightBack     ChannelLayout = "Rb"
	ChannelLayoutLeftTotal     ChannelLayout = "Lt"
	ChannelLayoutRightTotal    ChannelLayout = "Rt"
	ChannelLayoutLFE           ChannelLayout = "LFE"
)

// AudioChannel describes the position and attributes of a
// single channel of audio inside a container
type AudioChannel struct {
	TrackIdx, ChannelIdx int
	Layout               string
}

//AudioDownmix holds source and output channels for providers
//to handle downmixing
type Downmix struct {
	Src []AudioChannel
	Dst []AudioChannel
}

// Features is a map whose key is a custom feature name and value is a json string
// representing the corresponding custom feature definition
type Features map[string]interface{}

// File
type File struct {
	Size      int64         `json:"size,omitempty"`
	Duration  time.Duration `json:"dur,omitempty"`
	Name      string        `json:"name,omitempty"`
	Container string        `json:"container,omitempty"`
	Video     Video         `json:"video,omitempty"`
	Audio     Audio         `json:"audio,omitempty"`

	Splice                  timecode.Splice `json:"splice,omitempty"`
	Downmix                 *Downmix
	ExplicitKeyframeOffsets []float64
}

func (f File) Join(name string) File {
	u := f.URL()
	u.Path = path.Join(u.Path, name)
	f.Name = u.String()
	return f
}

func (f File) URL() url.URL {
	u, _ := url.Parse(f.Name)
	if u == nil {
		return url.URL{}
	}
	return *u
}
func (f File) Provider() string {
	return f.URL().Scheme
}
func (f File) Type() string {
	return strings.TrimPrefix(path.Ext(f.URL().Path), ".")
}
func (f File) Dir() string {
	if f.Type() == "" {
		return f.Name
	}
	u := f.URL()
	u.Path = path.Dir(u.Path)
	return u.String()
}
func (f File) Base() string {
	return path.Base(f.URL().Path)
}

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

// Audio defines audio transcoding parameters
type Audio struct {
	Codec     string `json:"codec,omitempty"`
	Bitrate   int    `json:"bitrate,omitempty"`
	Normalize bool   `json:"normalize,omitempty"`
	Discrete  bool   `json:"discrete,omitempty"`
}

func (a *Audio) On() bool {
	return a != nil && *a != (Audio{})
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

//Overlays defines all the overlay settings for a Video preset
type Overlays struct {
	Images         []Image   `json:"images,omitempty"`
	TimecodeBurnin *Timecode `json:"timecodeBurnin,omitempty"`
}

//Image defines the image overlay settings
type Image struct {
	URL string `json:"url"`
}

// Timecode settings
type Timecode struct {
	FontSize int    `json:"fontSize,omitempty"`
	Position int    `json:"position,omitempty"`
	Prefix   string `json:"prefix,omitempty"`
}

// HDR10 configurations and metadata
type HDR10 struct {
	Enabled       bool   `json:"enabled"`
	MaxCLL        int    `json:"maxCLL,omitempty"`
	MaxFALL       int    `json:"maxFALL,omitempty"`
	MasterDisplay string `json:"masterDisplay,omitempty"`
}

// DolbyVision settings
type DolbyVision struct {
	Enabled bool `json:"enabled"`
}
