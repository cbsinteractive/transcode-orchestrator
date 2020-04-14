package model

type TsStreamConfiguration struct {
	// The UUID of the stream to which this configuration belongs to. This has to be a ID of a stream that has been added to the current muxing. (required)
	StreamId string `json:"streamId,omitempty"`
	// An integer value. MPEG Transport Stream Packet Identifier (PID) for this stream. (required)
	Pid *int32 `json:"pid,omitempty"`
}

