package model
// LevelH264 : Specified set of constraints that indicate a degree of required decoder performance for a profile, see: https://en.wikipedia.org/wiki/H.264/MPEG-4_AVC#Levels
type LevelH264 string

// List of LevelH264
const (
	LevelH264_L1 LevelH264 = "1"
	LevelH264_L1b LevelH264 = "1b"
	LevelH264_L1_1 LevelH264 = "1.1"
	LevelH264_L1_2 LevelH264 = "1.2"
	LevelH264_L1_3 LevelH264 = "1.3"
	LevelH264_L2 LevelH264 = "2"
	LevelH264_L2_1 LevelH264 = "2.1"
	LevelH264_L2_2 LevelH264 = "2.2"
	LevelH264_L3 LevelH264 = "3"
	LevelH264_L3_1 LevelH264 = "3.1"
	LevelH264_L3_2 LevelH264 = "3.2"
	LevelH264_L4 LevelH264 = "4"
	LevelH264_L4_1 LevelH264 = "4.1"
	LevelH264_L4_2 LevelH264 = "4.2"
	LevelH264_L5 LevelH264 = "5"
	LevelH264_L5_1 LevelH264 = "5.1"
	LevelH264_L5_2 LevelH264 = "5.2"
)