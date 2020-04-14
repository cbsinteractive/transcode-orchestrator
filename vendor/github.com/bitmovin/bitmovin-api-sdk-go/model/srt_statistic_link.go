package model

type SrtStatisticLink struct {
	Bandwidth *float64 `json:"bandwidth,omitempty"`
	MaxBandwidth *float64 `json:"maxBandwidth,omitempty"`
	Rtt *float64 `json:"rtt,omitempty"`
}

