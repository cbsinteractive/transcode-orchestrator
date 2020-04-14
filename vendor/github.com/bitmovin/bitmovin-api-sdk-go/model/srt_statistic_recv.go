package model

type SrtStatisticRecv struct {
	Bytes *int64 `json:"bytes,omitempty"`
	BytesDropped *int64 `json:"bytesDropped,omitempty"`
	BytesLost *int64 `json:"bytesLost,omitempty"`
	MbitRate *float64 `json:"mbitRate,omitempty"`
	Packets *int64 `json:"packets,omitempty"`
	PacketsBelated *int64 `json:"packetsBelated,omitempty"`
	PacketsDropped *int64 `json:"packetsDropped,omitempty"`
	PacketsLost *int64 `json:"packetsLost,omitempty"`
	PacketsRetransmitted *int64 `json:"packetsRetransmitted,omitempty"`
}

