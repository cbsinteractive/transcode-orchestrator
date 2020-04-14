package model

type ResponseErrorData struct {
	// Contains an error code as defined in https://bitmovin.com/encoding-documentation/bitmovin-api/#/introduction/api-error-codes (required)
	Code *int32 `json:"code,omitempty"`
	// General error message (required)
	Message string `json:"message,omitempty"`
	// More detailed message meant for developers (required)
	DeveloperMessage string `json:"developerMessage,omitempty"`
	// collection of links to webpages containing further information on the topic
	Links []Link `json:"links,omitempty"`
	// collection of messages containing more detailed information on the cause of the error
	Details []Message `json:"details,omitempty"`
}

