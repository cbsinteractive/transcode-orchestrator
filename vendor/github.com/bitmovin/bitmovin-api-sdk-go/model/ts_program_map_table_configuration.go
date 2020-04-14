package model

type TsProgramMapTableConfiguration struct {
	// An integer value. Packet Identifier (PID) for the MPEG Transport Stream PMT.
	Pid *int32 `json:"pid,omitempty"`
}

