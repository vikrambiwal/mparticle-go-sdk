package events

// PromotionAction struct for PromotionAction
type PromotionAction struct {
	Action     string      `json:"action,omitempty"`
	Promotions []Promotion `json:"promotions,omitempty"`
}
