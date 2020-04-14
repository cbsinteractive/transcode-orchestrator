package model
type MuxingType string

// List of MuxingType
const (
	MuxingType_FMP4 MuxingType = "FMP4"
	MuxingType_CMAF MuxingType = "CMAF"
	MuxingType_MP4 MuxingType = "MP4"
	MuxingType_TS MuxingType = "TS"
	MuxingType_WEBM MuxingType = "WEBM"
	MuxingType_MP3 MuxingType = "MP3"
	MuxingType_PROGRESSIVE_WEBM MuxingType = "PROGRESSIVE_WEBM"
	MuxingType_PROGRESSIVE_MOV MuxingType = "PROGRESSIVE_MOV"
	MuxingType_PROGRESSIVE_TS MuxingType = "PROGRESSIVE_TS"
	MuxingType_BROADCAST_TS MuxingType = "BROADCAST_TS"
	MuxingType_CHUNKED_TEXT MuxingType = "CHUNKED_TEXT"
	MuxingType_TEXT MuxingType = "TEXT"
	MuxingType_SEGMENTED_RAW MuxingType = "SEGMENTED_RAW"
)