package model

type BroadcastTsAudioInputStreamConfiguration struct {
	// The UUID of the audio stream to which this configuration belongs to. This has to be an ID of an audio stream that has been added to the current muxing. (required)
	StreamId string `json:"streamId,omitempty"`
	// An integer value. Packet Identifier (PID) for this stream.
	PacketIdentifier *int32 `json:"packetIdentifier,omitempty"`
	// Start stream with initial discontinuity indicator set to one. If true, set the discontinuity indicator in the first packet for this PID.
	StartWithDiscontinuityIndicator *bool `json:"startWithDiscontinuityIndicator,omitempty"`
	// Align access units to PES packets. If true, align access units to PES packet headers. Uses adaptation field stuffing to position an access unit at the beginning of each PES packet.
	AlignPes *bool `json:"alignPes,omitempty"`
	SetRaiOnAu RaiUnit `json:"setRaiOnAu,omitempty"`
	// Use ATSC buffer model for AC-3. If true, use the ATSC version of the T-STD buffer model is used. This parameter applies to AC-3 streams only.
	UseATSCBufferModel *bool `json:"useATSCBufferModel,omitempty"`
	// Language of the audio stream. Specified according to the ISO 639-2 alpha code for the language descriptor.
	Language string `json:"language,omitempty"`
}

