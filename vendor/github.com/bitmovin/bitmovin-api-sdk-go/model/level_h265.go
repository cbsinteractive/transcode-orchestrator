package model
// LevelH265 : Specified set of constraints that indicate a degree of required decoder performance for a profile, see: https://en.wikipedia.org/wiki/High_Efficiency_Video_Coding_tiers_and_levels
type LevelH265 string

// List of LevelH265
const (
	LevelH265_L1 LevelH265 = "1"
	LevelH265_L2 LevelH265 = "2"
	LevelH265_L2_1 LevelH265 = "2.1"
	LevelH265_L3 LevelH265 = "3"
	LevelH265_L3_1 LevelH265 = "3.1"
	LevelH265_L4 LevelH265 = "4"
	LevelH265_L4_1 LevelH265 = "4.1"
	LevelH265_L5 LevelH265 = "5"
	LevelH265_L5_1 LevelH265 = "5.1"
	LevelH265_L5_2 LevelH265 = "5.2"
	LevelH265_L6 LevelH265 = "6"
	LevelH265_L6_1 LevelH265 = "6.1"
	LevelH265_L6_2 LevelH265 = "6.2"
)