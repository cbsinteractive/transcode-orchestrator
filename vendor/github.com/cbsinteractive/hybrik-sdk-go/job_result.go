package hybrik

import "time"

type JobResultResponse struct {
	Errors []JobResultError `json:"errors"`
	Job    JobResultSummary `json:"job"`
	Tasks  []TaskResult     `json:"tasks"`
}

type JobResultSummary struct {
	ID                 int       `json:"id"`
	JobClass           int       `json:"job_class"`
	IsAPIJob           int       `json:"is_api_job"`
	Priority           int       `json:"priority"`
	CreationTime       time.Time `json:"creation_time"`
	ExpirationTime     time.Time `json:"expiration_time"`
	HintActiveTimeSec  int       `json:"hint_active_time_sec"`
	HintMachineTimeSec int       `json:"hint_machine_time_sec"`
	LastTimesUpdated   time.Time `json:"last_times_updated"`
	SubscriptionKey    string    `json:"subscription_key"`
	Flags              int       `json:"flags"`
	Status             string    `json:"status"`
	RenderStatus       string    `json:"render_status"`
	TaskCount          int       `json:"task_count"`
	Progress           int       `json:"progress"`
	Name               string    `json:"name"`
	FirstStarted       time.Time `json:"first_started"`
	LastCompleted      time.Time `json:"last_completed"`
}

type JobResultError struct {
	TaskInstanceID   int         `json:"task_instance_id"`
	RetryNr          int         `json:"retry_nr"`
	MachineFetcherID interface{} `json:"machine_fetcher_id"`
	ResultDefine     string      `json:"result_define"`
	Message          string      `json:"message"`
	Details          interface{} `json:"details"`
	Diagnostic       interface{} `json:"diagnostic"`
	RecoverableError int         `json:"recoverable_error"`
	Assigned         interface{} `json:"assigned"`
	ResultCommitted  time.Time   `json:"result_committed"`
}

type ExecutionResultDetails struct {
	MachineID int `json:"machine_id"`
	ServiceID int `json:"service_id"`
	TaskID    int `json:"task_id"`
}

type LocationResult struct {
	Path            string `json:"path"`
	StorageProvider string `json:"storage_provider"`
}

type AssetVersionsResult struct {
	Location        LocationResult         `json:"location"`
	AssetComponents []AssetComponentResult `json:"asset_components"`
	VersionUID      string                 `json:"version_uid"`
}
type AssetResultPayload struct {
	AssetVersions []AssetVersionsResult `json:"asset_versions"`
	Kind          string                `json:"kind"`
}

type ResultPayload struct {
	Kind    string             `json:"kind"`
	Payload AssetResultPayload `json:"payload"`
}

type DocumentResult struct {
	ResultPayload ResultPayload `json:"result_payload"`
	Connector     string        `json:"connector"`
}

type Params struct {
	Location    LocationResult `json:"location"`
	FilePattern string         `json:"file_pattern"`
}
type MezzQCResultPayload struct {
	Module string `json:"module"`
	Params Params `json:"params"`
}

type ResultConfig struct {
	UID     string             `json:"uid"`
	Kind    string             `json:"kind"`
	Payload AssetResultPayload `json:"payload"`
	Name    string             `json:"name"`
}

type PreprocessingResult struct {
	Task TaskTags `json:"task"`
}

type TranscodeTaskResult struct {
	Name string `json:"name"`
}

type TranscodeTaskContainerResult struct {
	Kind string `json:"kind"`
}

type TranscodeTaskVideoResult struct {
	Width          int    `json:"width"`
	Height         int    `json:"height"`
	BitrateMode    string `json:"bitrate_mode"`
	MinBitrateKb   int    `json:"min_bitrate_kb"`
	BitrateKb      int    `json:"bitrate_kb"`
	MaxBitrateKb   int    `json:"max_bitrate_kb"`
	Preset         string `json:"preset"`
	Codec          string `json:"codec"`
	Profile        string `json:"profile"`
	MinGopFrames   int    `json:"min_gop_frames"`
	MaxGopFrames   int    `json:"max_gop_frames"`
	ExactGopFrames int    `json:"exact_gop_frames"`
	InterlaceMode  string `json:"interlace_mode"`
	X265Options    string `json:"x265_options"`
	Vtag           string `json:"vtag"`
	FfmpegArgs     string `json:"ffmpeg_args"`
}

type TranscodeTaskTargetResult struct {
	FilePattern   string                       `json:"file_pattern"`
	Container     TranscodeTaskContainerResult `json:"container"`
	Video         TranscodeTaskVideoResult     `json:"video"`
	ExistingFiles string                       `json:"existing_files"`
}

type SegmentedRenderingResult struct {
	DurationSec int `json:"duration_sec"`
}

type SourcePipeline struct {
	SegmentedRendering SegmentedRenderingResult `json:"segmented_rendering"`
}

type OptionsPipelineResult struct {
	EncoderVersion string `json:"encoder_version"`
}

type TranscodeTaskOptionsResult struct {
	Pipeline       OptionsPipelineResult `json:"pipeline"`
	SourceReadMode string                `json:"source_read_mode"`
}

type TranscodeTaskResultPayload struct {
	Location       LocationResult              `json:"location"`
	Targets        []TranscodeTaskTargetResult `json:"targets"`
	SourcePipeline SourcePipeline              `json:"source_pipeline"`
	Options        TranscodeTaskOptionsResult  `json:"options"`
}

type TranscodeElementResult struct {
	UID     string              `json:"uid"`
	Kind    string              `json:"kind"`
	Task    TranscodeTaskResult `json:"task"`
	Payload AssetResultPayload  `json:"payload"`
}

type Mp4MuxResult struct {
	Enabled     bool   `json:"enabled"`
	FilePattern string `json:"file_pattern"`
	ToolVersion string `json:"tool_version"`
}

type PostTranscodeResult struct {
	Mp4Mux Mp4MuxResult `json:"mp4_mux"`
}

type DoViResultPayload struct {
	Module        string                   `json:"module"`
	Profile       int                      `json:"profile"`
	Location      LocationResult           `json:"location"`
	Preprocessing PreprocessingResult      `json:"preprocessing"`
	Transcodes    []TranscodeElementResult `json:"transcodes"`
	PostTranscode PostTranscodeResult      `json:"post_transcode"`
}

type TaskResult struct {
	ID               int              `json:"id"`
	Priority         int              `json:"priority"`
	RetryNr          int              `json:"retry_nr"`
	RetryNrAog       int              `json:"retry_nr_aog"`
	CreationTime     time.Time        `json:"creation_time"`
	MaxRetryCountAog int              `json:"max_retry_count_aog"`
	RelatedAssetID   interface{}      `json:"related_asset_id"`
	Kind             string           `json:"kind"`
	Name             string           `json:"name"`
	RetryCount       int              `json:"retry_count"`
	UID              string           `json:"uid"`
	ElementName      string           `json:"element_name"`
	Status           string           `json:"status"`
	Assigned         time.Time        `json:"assigned"`
	Completed        time.Time        `json:"completed"`
	Documents        []DocumentResult `json:"documents"`
	FetcherID        int              `json:"fetcher_id,omitempty"`
}

type AssetComponentResult struct {
	Kind         string                         `json:"kind"`
	Name         string                         `json:"name"`
	Descriptor   AssetComponentResultDescriptor `json:"descriptor"`
	MediaAnalyze MediaAnalyzeResult             `json:"media_analyze"`
	ComponentUID string                         `json:"component_uid"`
}

type AssetComponentResultDescriptor struct {
	Size     int    `json:"size"`
	Provider string `json:"provider"`
	Checked  int64  `json:"checked"`
}

type MediaInfoAssetResult struct {
	URL                 string `json:"url"`
	Format              string `json:"format"`
	FormatProfile       string `json:"format_profile"`
	FormatCompatibility string `json:"format_compatibility"`
	Creator             string `json:"creator"`
	IsStreamable        bool   `json:"is_streamable"`
	TotalSize           int    `json:"total_size"`
	ContentType         string `json:"content_type"`
	Hash                string `json:"hash"`
	Modified            int64  `json:"modified"`
	ProbeStyle          string `json:"probe_style"`
	Requester           string `json:"requester"`
}
type MediaInfoAudioResult struct {
	NrChannels   int      `json:"nr_channels"`
	ChannelOrder string   `json:"channel_order"`
	Designators  []string `json:"designators"`
}

type MediaInfoVideoResult struct {
	BitResolution       int     `json:"bit_resolution"`
	Format              string  `json:"format"`
	BitsPerPixel        int     `json:"bits_per_pixel"`
	EncodedBitsPerPixel float64 `json:"encoded_bits_per_pixel"`
	ChromaSubsampling   string  `json:"chroma_subsampling"`
	ColorSpace          string  `json:"color_space"`
	FramerateMode       string  `json:"framerate_mode"`
	InterlaceMode       string  `json:"interlace_mode"`
	DisplayAr           float64 `json:"display_ar"`
	PixelAr             int     `json:"pixel_ar"`
	Height              int     `json:"height"`
	Width               int     `json:"width"`
	CleanApertureHeight int     `json:"clean_aperture_height"`
	CleanApertureWidth  int     `json:"clean_aperture_width"`
}

type MediaInfoResult struct {
	StreamType       string               `json:"stream_type"`
	BitPerSec        int                  `json:"bit_per_sec"`
	StreamSize       int                  `json:"stream_size,omitempty"`
	DurationOtb      int64                `json:"duration_otb"`
	FirstPtsOtb      int                  `json:"first_pts_otb"`
	FirstDtsOtb      int                  `json:"first_dts_otb"`
	ASSET            MediaInfoAssetResult `json:"ASSET,omitempty"`
	FirstDts         float64              `json:"first_dts"`
	FirstPts         float64              `json:"first_pts"`
	Duration         float64              `json:"duration"`
	ID               int                  `json:"id,omitempty"`
	Index            int                  `json:"index,omitempty"`
	SubIndex         int                  `json:"sub_index,omitempty"`
	FrameDurationOtb int                  `json:"frame_duration_otb,omitempty"`
	Codec            string               `json:"codec,omitempty"`
	CodecProfile     string               `json:"codec_profile,omitempty"`
	BitrateMode      string               `json:"bitrate_mode,omitempty"`
	AUDIO            MediaInfoAudioResult `json:"AUDIO,omitempty"`
	FrameRate        float64              `json:"frame_rate,omitempty"`
	Encoder          string               `json:"encoder,omitempty"`
	EncoderSettings  string               `json:"encoder_settings,omitempty"`
	VIDEO            MediaInfoVideoResult `json:"VIDEO,omitempty"`
}
type MediaAnalyzeResult struct {
	MediaInfo []MediaInfoResult `json:"media_info"`
}
