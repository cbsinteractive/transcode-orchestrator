package model

type TimeSpan struct {
	// Start offset of the time frame in milliseconds (required)
	From *int64 `json:"from,omitempty"`
	// End offset of the time frame in milliseconds (required)
	To *int64 `json:"to,omitempty"`
}

