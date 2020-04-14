package model

type SrtStatisticSend struct {
	Bytes *int64 `json:"bytes,omitempty"`
	BytesDropped *int64 `json:"bytesDropped,omitempty"`
	MbitRate *float64 `json:"mbitRate,omitempty"`
	Packets *int64 `json:"packets,omitempty"`
	PacketsDropped *int64 `json:"packetsDropped,omitempty"`
	PacketsLost *int64 `json:"packetsLost,omitempty"`
	PacketsRetransmitted *int64 `json:"packetsRetransmitted,omitempty"`
}

