package events

// CustomEventType model for CustomEventData
type CustomEventType string

// List of CustomEventType
const (
	OtherCustomEventType          CustomEventType = "other"
	LocationCustomEventType       CustomEventType = "location"
	SearchCustomEventType         CustomEventType = "search"
	TransactionCustomEventType    CustomEventType = "transaction"
	UserContentCustomEventType    CustomEventType = "user_content"
	UserPreferenceCustomEventType CustomEventType = "user_preference"
	SocialCustomEventType         CustomEventType = "social"
	MediaCustomEventType          CustomEventType = "media"
)
