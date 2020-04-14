package model

type BackupSrtInputs struct {
	// When there is no input signal present for this number of seconds, the encoder will switch to the next input
	DelayThreshold *int32 `json:"delayThreshold,omitempty"`
	SrtInputs []SrtInput `json:"srtInputs,omitempty"`
}

