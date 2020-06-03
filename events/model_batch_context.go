package events

// BatchContext struct
type BatchContext struct {
	DataPlan *DataPlanContext `json:"data_plan,omitempty"`
	Batching *BatchingContext `json:"batching,omitempty"`
}
