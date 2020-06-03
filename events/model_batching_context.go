package events

// BatchingContext struct
type BatchingContext struct {
	Bypass       bool     `json:"bypass,omitempty"`
	PartitionKey string   `json:"partition_key,omitempty"`
	RequestIDs   []string `json:"request_ids,omitempty"`
}
