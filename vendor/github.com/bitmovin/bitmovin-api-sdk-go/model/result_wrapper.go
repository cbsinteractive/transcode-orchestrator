package model

import "encoding/json"

type ResultWrapper struct {
	Result json.RawMessage `json:"result,omitempty"`
}
