package model

type TsMuxingConfiguration struct {
	// An integer value. Value for program_number field in the MPEG Transport Stream Program Map Table (PMT). The value zero is reserved for the NIT PID entry in the PAT.
	ProgramNumber *int32 `json:"programNumber,omitempty"`
	Pmt *TsProgramMapTableConfiguration `json:"pmt,omitempty"`
	Pcr *TsProgramClockReferenceConfiguration `json:"pcr,omitempty"`
	VideoStreams []TsVideoStreamConfiguration `json:"videoStreams,omitempty"`
	AudioStreams []TsAudioStreamConfiguration `json:"audioStreams,omitempty"`
}

