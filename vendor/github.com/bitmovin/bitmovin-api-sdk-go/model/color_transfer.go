package model
type ColorTransfer string

// List of ColorTransfer
const (
	ColorTransfer_UNSPECIFIED ColorTransfer = "UNSPECIFIED"
	ColorTransfer_BT709 ColorTransfer = "BT709"
	ColorTransfer_GAMMA22 ColorTransfer = "GAMMA22"
	ColorTransfer_GAMMA28 ColorTransfer = "GAMMA28"
	ColorTransfer_SMPTE170M ColorTransfer = "SMPTE170M"
	ColorTransfer_SMPTE240M ColorTransfer = "SMPTE240M"
	ColorTransfer_LINEAR ColorTransfer = "LINEAR"
	ColorTransfer_LOG ColorTransfer = "LOG"
	ColorTransfer_LOG_SQRT ColorTransfer = "LOG_SQRT"
	ColorTransfer_IEC61966_2_4 ColorTransfer = "IEC61966_2_4"
	ColorTransfer_BT1361_ECG ColorTransfer = "BT1361_ECG"
	ColorTransfer_IEC61966_2_1 ColorTransfer = "IEC61966_2_1"
	ColorTransfer_BT2020_10 ColorTransfer = "BT2020_10"
	ColorTransfer_BT2020_12 ColorTransfer = "BT2020_12"
	ColorTransfer_SMPTE2084 ColorTransfer = "SMPTE2084"
	ColorTransfer_SMPTE428 ColorTransfer = "SMPTE428"
	ColorTransfer_ARIB_STD_B67 ColorTransfer = "ARIB_STD_B67"
)