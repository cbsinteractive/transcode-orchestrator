package model

type BroadcastTsTransportConfiguration struct {
	// Output rate in bps. The value zero implies to use minimal rate. The minimal rate leaves approximately 15kbps of null packets in the stream.
	Muxrate *float64 `json:"muxrate,omitempty"`
	// Stop mux on errors. If true, implies halt multiplexing when any error is encountered. If false, errors are ignored and multiplexing continues. Note that the recovery from an error will usually result in an illegal transport stream and artifacts on a decoder.
	StopOnError *bool `json:"stopOnError,omitempty"`
	// If true, prevents adaptation fields with length field equal to zero in video, i.e., zero-length AF. Please note that this condition can only occur when pesAlign for the input stream is set to true.
	PreventEmptyAdaptionFieldsInVideo *bool `json:"preventEmptyAdaptionFieldsInVideo,omitempty"`
	// Program Association Table (PAT) repetition rate per second. Number of PATs per second.
	PatRepetitionRatePerSec *float64 `json:"patRepetitionRatePerSec,omitempty"`
	// Program Map Table (PMT) repetition rate per second. Number of PMTs for each program per second.
	PmtRepetitionRatePerSec *float64 `json:"pmtRepetitionRatePerSec,omitempty"`
	// When false, the output stream is created at a constant bit rate. When true, the output rate is allowed to vary from a maximum rate set by the muxrate parameter down to the minimum required to carry the stream. Default: false
	VariableMuxRate *bool `json:"variableMuxRate,omitempty"`
	// Sets the presentation time stamp value for the first video frame. The timestamp is specified in the timescale of 90000. Default: 0
	InitialPresentationTimeStamp *float64 `json:"initialPresentationTimeStamp,omitempty"`
	// Sets the Program Clock Reference value at the beginning of the first packet for the transport stream. The PCR is specified in the timescale of 90000. Default: 0
	InitialProgramClockReference *float64 `json:"initialProgramClockReference,omitempty"`
}

