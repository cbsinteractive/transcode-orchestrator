package model

type ResponseError struct {
	// Unique correlation id (required)
	RequestId string `json:"requestId,omitempty"`
	// Response status information (required)
	Status ResponseStatus `json:"status,omitempty"`
	// Response information (required)
	Data *ResponseErrorData `json:"data,omitempty"`
}

