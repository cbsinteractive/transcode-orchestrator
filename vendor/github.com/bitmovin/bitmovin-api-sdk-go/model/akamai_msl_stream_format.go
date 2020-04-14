package model
type AkamaiMslStreamFormat string

// List of AkamaiMslStreamFormat
const (
	AkamaiMslStreamFormat_DASH AkamaiMslStreamFormat = "DASH"
	AkamaiMslStreamFormat_HLS AkamaiMslStreamFormat = "HLS"
	AkamaiMslStreamFormat_CMAF AkamaiMslStreamFormat = "CMAF"
)