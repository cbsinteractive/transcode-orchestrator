package model

type BroadcastTsInputStreamConfiguration struct {
	// The UUID of the stream to which this configuration belongs to. This has to be a ID of a stream that has been added to the current muxing.
	StreamId string `json:"streamId,omitempty"`
	// An integer value. Packet Identifier (PID) for this stream.
	PacketIdentifier *int32 `json:"packetIdentifier,omitempty"`
	// Start stream with initial discontinuity indicator set to one. If true, set the discontinuity indicator in the first packet for this PID.
	StartWithDiscontinuityIndicator *bool `json:"startWithDiscontinuityIndicator,omitempty"`
	// Align access units to PES packets. If true, align access units to PES packet headers. Uses adaptation field stuffing to position an access unit at the beginning of each PES packet.
	AlignPes *bool `json:"alignPes,omitempty"`
	SetRaiOnAu RaiUnit `json:"setRaiOnAu,omitempty"`
}

