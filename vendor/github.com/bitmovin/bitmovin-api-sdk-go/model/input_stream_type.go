package model
type InputStreamType string

// List of InputStreamType
const (
	InputStreamType_INGEST InputStreamType = "INGEST"
	InputStreamType_CONCATENATION InputStreamType = "CONCATENATION"
	InputStreamType_TRIMMING_TIME_BASED InputStreamType = "TRIMMING_TIME_BASED"
	InputStreamType_TRIMMING_TIME_CODE_TRACK InputStreamType = "TRIMMING_TIME_CODE_TRACK"
	InputStreamType_TRIMMING_H264_PICTURE_TIMING InputStreamType = "TRIMMING_H264_PICTURE_TIMING"
	InputStreamType_SIDECAR_DOLBY_VISION_METADATA InputStreamType = "SIDECAR_DOLBY_VISION_METADATA"
	InputStreamType_AUDIO_MIX InputStreamType = "AUDIO_MIX"
	InputStreamType_FILE InputStreamType = "FILE"
	InputStreamType_CAPTION_CEA608 InputStreamType = "CAPTION_CEA608"
	InputStreamType_CAPTION_CEA708 InputStreamType = "CAPTION_CEA708"
	InputStreamType_DVB_TELETEXT InputStreamType = "DVB_TELETEXT"
)