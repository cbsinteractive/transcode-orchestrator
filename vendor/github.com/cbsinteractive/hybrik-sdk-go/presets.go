package hybrik

import (
	"bytes"
	"encoding/json"
	"fmt"
)

const (
	failed                 = "false"
	EncoderVersion4_10bit  = "hybrik_4.0_10bit"
	FFMPEGVersion3_4_10bit = "hybrik_3.4_10bit"
)

type presetResponse struct {
	Success string `json:"success"`
	Msg     string `json:"message"`
}

// ErrCreatePreset occurs when there is a problem creating a preset
type ErrCreatePreset struct {
	Msg string
}

func (e ErrCreatePreset) Error() string {
	return fmt.Sprintf("unable to create preset, error: %s", e.Msg)
}

// ErrGetPreset occurs when there is a problem obtaining a preset
type ErrGetPreset struct {
	Msg string
}

func (e ErrGetPreset) Error() string {
	return fmt.Sprintf("unable to get preset, error: %s", e.Msg)
}

// GetPreset return details of a given presetID
func (c *Client) GetPreset(presetID string) (Preset, error) {

	result, err := c.client.CallAPI("GET", fmt.Sprintf("/presets/%s", presetID), nil, nil)
	if err != nil {
		return Preset{}, err
	}

	var preset Preset
	err = json.Unmarshal([]byte(result), &preset)
	if err != nil {
		return Preset{}, err
	}

	if preset.Name == "" {
		var pr presetResponse
		err = json.Unmarshal([]byte(result), &pr)
		if err != nil {
			return Preset{}, err
		}
		if pr.Success == failed {
			return Preset{}, ErrGetPreset{Msg: pr.Msg}
		}
	}

	return preset, nil
}

// CreatePreset creates a new preset
func (c *Client) CreatePreset(preset Preset) (Preset, error) {
	body, err := json.Marshal(preset)
	if err != nil {
		return Preset{}, err
	}

	resp, err := c.client.CallAPI("POST", "/presets", nil, bytes.NewReader(body))
	if err != nil {
		return Preset{}, err
	}

	var pr presetResponse
	err = json.Unmarshal([]byte(resp), &pr)
	if err != nil {
		return Preset{}, err
	}

	if pr.Success == failed {
		return Preset{}, ErrCreatePreset{Msg: pr.Msg}
	}

	return preset, nil
}

// DeletePreset removes a preset based on its presetID
func (c *Client) DeletePreset(presetID string) error {
	_, err := c.client.CallAPI("DELETE", fmt.Sprintf("/presets/%s", presetID), nil, nil)

	return err
}

// PresetList represents the response returned by
// a query for the list of jobs
type PresetList []Preset

// Preset represents a transcoding preset
type Preset struct {
	Key         string        `json:"key"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	UserData    string        `json:"user_data,omitempty"`
	Kind        string        `json:"kind"`
	Path        string        `json:"path"`
	Payload     PresetPayload `json:"payload"`
}

type PresetPayload struct {
	Targets []PresetTarget        `json:"targets"`
	Options *TranscodeTaskOptions `json:"options,omitempty"`
}

type TranscodeTaskOptions struct {
	Pipeline       *PipelineOptions `json:"pipeline,omitempty"`
	SourceReadMode string           `json:"source_read_mode,omitempty"`
}

type PipelineOptions struct {
	EncoderVersion string `json:"encoder_version,omitempty"`
	FFMPEGVersion  string `json:"ffmpeg_version,omitempty"`
}

// PresetTarget holds configuration options for presets
type PresetTarget struct {
	FilePattern   string             `json:"file_pattern"`
	Container     TranscodeContainer `json:"container"`
	NumPasses     int                `json:"nr_of_passes,omitempty"`
	Video         VideoTarget        `json:"video,omitempty"`
	Audio         []AudioTarget      `json:"audio,omitempty"`
	ExistingFiles string             `json:"existing_files,omitempty"`
	UID           string             `json:"uid,omitempty"`
}

// VideoTarget holds configuration options for video outputs
type VideoTarget struct {
	Width             *int           `json:"width,omitempty"`
	Height            *int           `json:"height,omitempty"`
	BitrateMode       string         `json:"bitrate_mode,omitempty"`
	MinBitrateKb      int            `json:"min_bitrate_kb,omitempty"`
	BitrateKb         int            `json:"bitrate_kb,omitempty"`
	MaxBitrateKb      int            `json:"max_bitrate_kb,omitempty"`
	VbvBufferSizeKb   int            `json:"vbv_buffer_size_kb,omitempty"`
	Preset            string         `json:"preset,omitempty"`
	FrameRate         string         `json:"frame_rate,omitempty"`
	Codec             string         `json:"codec,omitempty"`
	Profile           string         `json:"profile,omitempty"`
	Level             string         `json:"level,omitempty"`
	Tune              string         `json:"tune,omitempty"`
	MinGOPFrames      int            `json:"min_gop_frames,omitempty"`
	MaxGOPFrames      int            `json:"max_gop_frames,omitempty"`
	ExactGOPFrames    int            `json:"exact_gop_frames,omitempty"`
	MinKeyFrame       int            `json:"min_keyframe_interval_sec,omitempty"`
	MaxKeyFrame       int            `json:"max_keyframe_interval_sec,omitempty"`
	ExactKeyFrame     int            `json:"exact_keyframe_interval_sec,omitempty"`
	UseClosedGOP      bool           `json:"use_closed_gop,omitempty"`
	InterlaceMode     string         `json:"interlace_mode,omitempty"`
	ChromaFormat      string         `json:"chroma_format,omitempty"`
	ColorPrimaries    string         `json:"color_primaries,omitempty"`
	ColorMatrix       string         `json:"color_matrix,omitempty"`
	UseSceneDetection bool           `json:"use_scene_detection,omitempty"`
	ColorTRC          string         `json:"color_trc,omitempty"`
	X265Options       string         `json:"x265_options,omitempty"`
	VTag              string         `json:"vtag,omitempty"`
	HDR10             *HDR10Settings `json:"hdr10,omitempty"`
	FFMPEGArgs        string         `json:"ffmpeg_args,omitempty"`
}

// HDR10Settings holds configuration information for the HDR color data.
type HDR10Settings struct {
	// Source holds the location of the HDR10 metadata.
	// can be one of: config, source_metadata, source_document, media, metadata_file, or none
	Source string `json:"source"`

	// MasterDisplay is an encoded string containing mastering display brightness.
	MasterDisplay string `json:"master_display"`

	// MaxCLL is the maximum Content Light Level (CLL) for the file.
	MaxCLL int `json:"max_cll"`

	// MaxFALL is the maximum Frame Average Light Level (FALL) for the file
	MaxFALL int `json:"max_fall"`
}

// AudioTarget holds configuration options for audio outputs
type AudioTarget struct {
	Codec      string              `json:"codec,omitempty"`
	Channels   int                 `json:"channels,omitempty"`
	SampleRate int                 `json:"sample_rate,omitempty"`
	SampleSize int                 `json:"sample_size,omitempty"`
	BitrateKb  int                 `json:"bitrate_kb,omitempty"`
	Source     []AudioTargetSource `json:"source,omitempty"`
}

// AudioTargetSource .
type AudioTargetSource struct {
	TrackNum int `json:"track"`
}
