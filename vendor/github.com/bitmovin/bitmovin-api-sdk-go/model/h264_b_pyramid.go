package model
// H264BPyramid : Keep some B-frames as references
type H264BPyramid string

// List of H264BPyramid
const (
	H264BPyramid_NONE H264BPyramid = "NONE"
	H264BPyramid_STRICT H264BPyramid = "STRICT"
	H264BPyramid_NORMAL H264BPyramid = "NORMAL"
)