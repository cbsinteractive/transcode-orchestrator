package model
// H264SubMe : Subpixel motion estimation and mode decision
type H264SubMe string

// List of H264SubMe
const (
	H264SubMe_FULLPEL H264SubMe = "FULLPEL"
	H264SubMe_SAD H264SubMe = "SAD"
	H264SubMe_SATD H264SubMe = "SATD"
	H264SubMe_QPEL3 H264SubMe = "QPEL3"
	H264SubMe_QPEL4 H264SubMe = "QPEL4"
	H264SubMe_QPEL5 H264SubMe = "QPEL5"
	H264SubMe_RD_IP H264SubMe = "RD_IP"
	H264SubMe_RD_ALL H264SubMe = "RD_ALL"
	H264SubMe_RD_REF_IP H264SubMe = "RD_REF_IP"
	H264SubMe_RD_REF_ALL H264SubMe = "RD_REF_ALL"
)