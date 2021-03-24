package av

// State is the Job's state
type State string

const (
	StateUnknown  = State("unknown")
	StateQueued   = State("queued")
	StateStarted  = State("started")
	StateFinished = State("finished")
	StateFailed   = State("failed")
	StateCanceled = State("canceled")
)

type Provider struct {
	Name   string                 `json:"name,omitempty"`
	JobID  string                 `json:"job_id,omitempty"`
	Status map[string]interface{} `json:"status,omitempty"`
}

// Status is the representation of the status
type Status struct {
	ID     string   `json:"jobID,omitempty"`
	Labels []string `json:"labels,omitempty"`

	State    State   `json:"status,omitempty"`
	Msg      string  `json:"msg,omitempty"`
	Progress float64 `json:"progress"`

	Input  File `json:"input"`
	Output Dir  `json:"output"`

	Provider       string                 `json:"providerName,omitempty"`
	ProviderJobID  string                 `json:"providerJobId,omitempty"`
	ProviderStatus map[string]interface{} `json:"providerStatus,omitempty"`
}
