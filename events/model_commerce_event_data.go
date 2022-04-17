package events

// CommerceEventData struct
type CommerceEventData struct {
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
	ProductAction               *map[string]interface{}      `json:"product_action,omitempty"`
	PromotionAction             *PromotionAction             `json:"promotion_action,omitempty"`
	ProductImpressions          []ProductImpression          `json:"product_impressions,omitempty"`
	ShoppingCart                *ShoppingCart                `json:"shopping_cart,omitempty"`
	CurrencyCode                string                       `json:"currency_code,omitempty"`
	ScreenName                  string                       `json:"screen_name,omitempty"`
	IsNonInteractive            bool                         `json:"is_non_interactive,omitempty"`
	EventName                   string                       `json:"event_name,omitempty"`
	CustomEventType             string                       `json:"custom_event_type,omitempty"`
	CustomFlags                 map[string]interface{}       `json:"custom_flags,omitempty"`
}
