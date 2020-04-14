package model

type SrtStatisticWindow struct {
	Congestion *int64 `json:"congestion,omitempty"`
	Flight *int64 `json:"flight,omitempty"`
	Flow *int64 `json:"flow,omitempty"`
}

