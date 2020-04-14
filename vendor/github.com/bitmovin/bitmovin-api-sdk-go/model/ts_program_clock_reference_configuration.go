package model

type TsProgramClockReferenceConfiguration struct {
	// An integer value. Packet Identifier (PID) for the MPEG Transport Stream PCR. This should generally point to the video stream PID. If it is not explicitly set it will point to the video stream PID if exists, otherwise to the audio stream PID.
	Pid *int32 `json:"pid,omitempty"`
	// An integer value. Nominal time between MPEG Transport Stream PCRs in milliseconds.
	Interval *int32 `json:"interval,omitempty"`
}

