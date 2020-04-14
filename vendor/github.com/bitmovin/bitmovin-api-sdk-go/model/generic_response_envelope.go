package model

import "encoding/json"

type GenericResponseEnvelope struct {
	// Unique correlation id
	RequestId string `json:"requestId"`
	// Response status information
	Status ResponseStatus `json:"status"`
	// Response information
    Data json.RawMessage `json:"data"`
}