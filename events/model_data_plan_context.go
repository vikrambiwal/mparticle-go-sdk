package events

// DataPlanContext struct
type DataPlanContext struct {
	PlanID      string `json:"plan_id"`
	PlanVersion int64  `json:"plan_version,omitempty"`
}
