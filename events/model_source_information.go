package events

// SourceInformation struct for SourceInformation
type SourceInformation struct {
	Channel         string `json:"channel,omitempty"`
	Partner         string `json:"partner,omitempty"`
	ReplayRequestID string `json:"replay_request_id,omitempty"`
	ReplayJobID     string `json:"replay_job_id,omitempty"`
	IsHistorical    bool   `json:"is_historical,omitempty"`
}
