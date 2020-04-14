package model
// H264NalHrd : Signal hypothetical reference decoder (HRD) information (requires bufsize to be set)
type H264NalHrd string

// List of H264NalHrd
const (
	H264NalHrd_NONE H264NalHrd = "NONE"
	H264NalHrd_VBR H264NalHrd = "VBR"
	H264NalHrd_CBR H264NalHrd = "CBR"
)