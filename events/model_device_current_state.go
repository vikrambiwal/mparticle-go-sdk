package events

// DeviceCurrentState struct
type DeviceCurrentState struct {
	TimeSinceStartMS                int64   `json:"time_since_start_ms,omitempty"`
	BatteryLevel                    float64 `json:"battery_level,omitempty"`
	DataConnectionType              string  `json:"data_connection_type,omitempty"`
	DataConnectionTypeDetail        string  `json:"data_connection_type_detail,omitempty"`
	GpsState                        bool    `json:"gps_state,omitempty"`
	TotalSystemMemoryUsageBytes     int64   `json:"total_system_memory_usage_bytes,omitempty"`
	DiskSpaceFreeBytes              int64   `json:"disk_space_free_bytes,omitempty"`
	ExternalDiskSpaceFreeBytes      int64   `json:"external_disk_space_free_bytes,omitempty"`
	CPU                             string  `json:"cpu,omitempty"`
	SystemMemoryAvailableBytes      float64 `json:"system_memory_available_bytes,omitempty"`
	SystemMemoryLow                 bool    `json:"system_memory_low,omitempty"`
	SystemMemoryThresholdBytes      float64 `json:"system_memory_threshold_bytes,omitempty"`
	ApplicationMemoryAvailableBytes float64 `json:"application_memory_available_bytes,omitempty"`
	ApplicationMemoryMaxBytes       float64 `json:"application_memory_max_bytes,omitempty"`
	ApplicationMemoryTotalBytes     float64 `json:"application_memory_total_bytes,omitempty"`
	DeviceOrientation               string  `json:"device_orientation,omitempty"`
	StatusBarOrientation            string  `json:"status_bar_orientation,omitempty"`
}
