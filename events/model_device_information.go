package events

// DeviceInformation struct
type DeviceInformation struct {
	Brand                          string `json:"brand,omitempty"`
	Product                        string `json:"product,omitempty"`
	Device                         string `json:"device,omitempty"`
	AndroidUUID                    string `json:"android_uuid,omitempty"`
	DeviceManufacturer             string `json:"device_manufacturer,omitempty"`
	Platform                       string `json:"platform,omitempty"`
	OsVersion                      string `json:"os_version,omitempty"`
	DeviceModel                    string `json:"device_model,omitempty"`
	ScreenHeight                   int64  `json:"screen_height,omitempty"`
	ScreenWidth                    int64  `json:"screen_width,omitempty"`
	ScreenDPI                      int64  `json:"screen_dpi,omitempty"`
	DeviceCountry                  string `json:"device_country,omitempty"`
	LocaleLanguage                 string `json:"locale_language,omitempty"`
	LocaleCountry                  string `json:"locale_country,omitempty"`
	NetworkCountry                 string `json:"network_country,omitempty"`
	NetworkCarrier                 string `json:"network_carrier,omitempty"`
	NetworkCode                    string `json:"network_code,omitempty"`
	NetworkMobileCountryCode       string `json:"network_mobile_country_code,omitempty"`
	TimezoneOffset                 int64  `json:"timezone_offset,omitempty"`
	BuildIdentifier                string `json:"build_identifier,omitempty"`
	HTTPHeaderUserAgent            string `json:"http_header_user_agent,omitempty"`
	IOSAdvertisingID               string `json:"ios_advertising_id,omitempty"`
	PushToken                      string `json:"push_token,omitempty"`
	CPUArchitecture                string `json:"cpu_architecture,omitempty"`
	IsTablet                       bool   `json:"is_tablet,omitempty"`
	PushNotificationSoundEnabled   bool   `json:"push_notification_sound_enabled,omitempty"`
	PushNotificationVibrateEnabled bool   `json:"push_notification_vibrate_enabled,omitempty"`
	RadioAccessTechnology          string `json:"radio_access_technology,omitempty"`
	SupportsTelephony              bool   `json:"supports_telephony,omitempty"`
	HasNFC                         bool   `json:"has_nfc,omitempty"`
	BluetoothEnabled               bool   `json:"bluetooth_enabled,omitempty"`
	BluetoothVersion               string `json:"bluetooth_version,omitempty"`
	IOSVendorIdentifier            string `json:"ios_idfv,omitempty"`
	AndroidAdvertisingID           string `json:"android_advertising_id,omitempty"`
	BuildVersionRelease            string `json:"build_version_release,omitempty"`
	LimitAdTracking                bool   `json:"limit_ad_tracking,omitempty"`
	AmpID                          string `json:"amp_id,omitempty"`
	IsDST                          bool   `json:"is_dst,omitempty"`
	RokuAdvertisingID              string `json:"roku_advertising_id,omitempty"`
	RokuPublisherID                string `json:"roku_publisher_id,omitempty"`
	MicrosoftAdvertisingID         string `json:"microsoft_advertising_id,omitempty"`
	MicrosoftPublisherID           string `json:"microsoft_publisher_id,omitempty"`
	FireAdvertisingID              string `json:"fire_advertising_id,omitempty"`
}
