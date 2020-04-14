package model

type ResponseEnvelope struct {
	// Unique correlation id (required)
	RequestId string `json:"requestId,omitempty"`
	// Response status information (required)
	Status ResponseStatus `json:"status,omitempty"`
	// Response information (required)
	Data *ResultWrapper `json:"data,omitempty"`
	// Additional endpoint specific information
	More *map[string]interface{} `json:"more,omitempty"`
}

