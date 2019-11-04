package events

// GdprConsentState struct for GdprConsentState
type GdprConsentState struct {
	Regulation          string `json:"regulation,omitempty,omitempty"`
	Document            string `json:"document,omitempty,omitempty"`
	Consented           bool   `json:"consented,omitempty,omitempty"`
	TimestampUnixtimeMS int64  `json:"timestamp_unixtime_ms,omitempty"`
	Location            string `json:"location,omitempty"`
	HardwareID          string `json:"hardware_id,omitempty"`
}
