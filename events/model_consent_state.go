package events

// ConsentState struct
type ConsentState struct {
	GDPR map[string]GdprConsentState `json:"gdpr,omitempty"`
}
