package events

// CrashReportEventData struct
type CrashReportEventData struct {
	TimestampUnixtimeMS         int64                        `json:"timestamp_unixtime_ms,omitempty"`
	EventID                     int64                        `json:"event_id,omitempty"`
	SourceMessageID             string                       `json:"source_message_id,omitempty"`
	SessionID                   int64                        `json:"session_id,omitempty"`
	SessionUUID                 string                       `json:"session_uuid,omitempty"`
	SessionStartUnixtimeMS      int64                        `json:"session_start_unixtime_ms,omitempty"`
	EventStartUnixtimeMS        int64                        `json:"event_start_unixtime_ms,omitempty"`
	CustomAttributes            map[string]string            `json:"custom_attributes,omitempty"`
	Location                    *GeoLocation                 `json:"location,omitempty"`
	DeviceCurrentState          *DeviceCurrentState          `json:"device_current_state,omitempty"`
	IsGoalDefined               bool                         `json:"is_goal_defined,omitempty"`
	LifetimeValueChange         bool                         `json:"lifetime_value_change,omitempty"`
	LifetimeValueAttributeName  string                       `json:"lifetime_value_attribute_name,omitempty"`
	DataConnectionType          string                       `json:"data_connection_type,omitempty"`
	EventNum                    int64                        `json:"event_num,omitempty"`
	ViewController              string                       `json:"view_controller,omitempty"`
	IsMainThread                bool                         `json:"is_main_thread,omitempty"`
	CanonicalName               string                       `json:"canonical_name,omitempty"`
	EventSystemNotificationInfo *EventSystemNotificationInfo `json:"event_system_notification_info,omitempty"`
	Breadcrumbs                 []string                     `json:"breadcrumbs,omitempty"`
	ClassName                   string                       `json:"class_name,omitempty"`
	Severity                    string                       `json:"severity,omitempty"`
	Message                     string                       `json:"message,omitempty"`
	StackTrace                  string                       `json:"stack_trace,omitempty"`
	ExceptionHandled            bool                         `json:"exception_handled,omitempty"`
	TopmostContext              string                       `json:"topmost_context,omitempty"`
	PlCrashReportFileBase64     string                       `json:"pl_crash_report_file_base64,omitempty"`
	IosImageBaseAddress         int64                        `json:"ios_image_base_address,omitempty"`
	IosImageSize                int64                        `json:"ios_image_size,omitempty"`
	SessionNumber               int64                        `json:"session_number,omitempty"`
}
