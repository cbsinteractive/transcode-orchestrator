package model

type BroadcastTsMuxingConfiguration struct {
	// Transport configuration details for the Broadcast TS muxing.
	Transport *BroadcastTsTransportConfiguration `json:"transport,omitempty"`
	// Program configuration details for the Broadcast TS muxing.
	Program *BroadcastTsProgramConfiguration `json:"program,omitempty"`
	VideoStreams []BroadcastTsVideoInputStreamConfiguration `json:"videoStreams,omitempty"`
	AudioStreams []BroadcastTsAudioInputStreamConfiguration `json:"audioStreams,omitempty"`
}

