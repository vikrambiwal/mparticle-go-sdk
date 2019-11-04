package events

// Promotion struct for Promotion
type Promotion struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Creative string `json:"creative,omitempty"`
	Position string `json:"position,omitempty"`
}
