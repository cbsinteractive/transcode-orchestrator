package model

type StreamDetails struct {
	Id string `json:"id,omitempty"`
	Codec string `json:"codec,omitempty"`
	Duration *int32 `json:"duration,omitempty"`
	Position *int32 `json:"position,omitempty"`
}

