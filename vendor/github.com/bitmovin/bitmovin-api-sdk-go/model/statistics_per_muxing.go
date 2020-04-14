package model

type StatisticsPerMuxing struct {
	// ID of the stream (required)
	StreamId string `json:"streamId,omitempty"`
	// ID of the muxing (required)
	MuxingId string `json:"muxingId,omitempty"`
	// Multiplier for the encoded minutes. Depends on muxing type. (required)
	Multiplicator *float64 `json:"multiplicator,omitempty"`
	// Encoded bytes. (required)
	EncodedBytes *int64 `json:"encodedBytes,omitempty"`
	// Resulting minutes you will be charged for. (required)
	BillableMinutes *float64 `json:"billableMinutes,omitempty"`
	MuxingType MuxingType `json:"muxingType,omitempty"`
}

