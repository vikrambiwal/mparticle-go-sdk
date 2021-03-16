package events

// ATTAuthorizationStatus model for CustomEventData
type ATTAuthorizationStatus string

// List of ATTAuthorizationStatus
const (
	AuthorizedATTAuthorizationStatus    ATTAuthorizationStatus = "authorized"
	DeniedATTAuthorizationStatus        ATTAuthorizationStatus = "denied"
	NotDeterminedATTAuthorizationStatus ATTAuthorizationStatus = "not_determined"
	RestrictedATTAuthorizationStatus    ATTAuthorizationStatus = "restricted"
)
