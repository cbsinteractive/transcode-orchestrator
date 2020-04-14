package model

type Statistics struct {
	// Bytes encoded total. (required)
	BytesEncodedTotal *int64 `json:"bytesEncodedTotal,omitempty"`
	// Time in seconds encoded for all contained daily statistics. (required)
	TimeEncodedTotal *int64 `json:"timeEncodedTotal,omitempty"`
}

