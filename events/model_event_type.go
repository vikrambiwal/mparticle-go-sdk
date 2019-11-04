package events

// EventType the model 'EventType'
type EventType string

// List of EventType
const (
	SessionStartEventType               EventType = "session_start"
	SessionEndEventType                 EventType = "session_end"
	ScreenViewEventType                 EventType = "screen_view"
	CustomEventEventType                EventType = "custom_event"
	CrashReportEventType                EventType = "crash_report"
	OptOutEventType                     EventType = "opt_out"
	FirstRunEventType                   EventType = "first_run"
	PreAttributionEventType             EventType = "pre_attribution"
	PushRegistrationEventType           EventType = "push_registration"
	ApplicationStateTransitionEventType EventType = "application_state_transition"
	PushMessageEventType                EventType = "push_message"
	NetworkPerformanceEventType         EventType = "network_performance"
	BreadcrumbEventType                 EventType = "breadcrumb"
	ProfileEventType                    EventType = "profile"
	PushReactionEventType               EventType = "push_reaction"
	CommerceEventType                   EventType = "commerce_event"
	UserAttributeChangeEventType        EventType = "user_attribute_change"
	UserIdentityChangeEventType         EventType = "user_identity_change"
	UninstallEventType                  EventType = "uninstall"
)
