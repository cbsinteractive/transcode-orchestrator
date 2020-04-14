package model
type CodecConfigType string

// List of CodecConfigType
const (
	CodecConfigType_AAC CodecConfigType = "AAC"
	CodecConfigType_HE_AAC_V1 CodecConfigType = "HE_AAC_V1"
	CodecConfigType_HE_AAC_V2 CodecConfigType = "HE_AAC_V2"
	CodecConfigType_H264 CodecConfigType = "H264"
	CodecConfigType_H265 CodecConfigType = "H265"
	CodecConfigType_VP9 CodecConfigType = "VP9"
	CodecConfigType_VP8 CodecConfigType = "VP8"
	CodecConfigType_MP2 CodecConfigType = "MP2"
	CodecConfigType_MP3 CodecConfigType = "MP3"
	CodecConfigType_AC3 CodecConfigType = "AC3"
	CodecConfigType_EAC3 CodecConfigType = "EAC3"
	CodecConfigType_OPUS CodecConfigType = "OPUS"
	CodecConfigType_VORBIS CodecConfigType = "VORBIS"
	CodecConfigType_MJPEG CodecConfigType = "MJPEG"
	CodecConfigType_AV1 CodecConfigType = "AV1"
)