package events

// ApplicationInformation struct
type ApplicationInformation struct {
	ApplicationName          string                       `json:"application_name,omitempty"`
	ApplicationVersion       string                       `json:"application_version,omitempty"`
	ApplicationBuildNumber   string                       `json:"application_build_number,omitempty"`
	InstallReferrer          string                       `json:"install_referrer,omitempty"`
	Package                  string                       `json:"package,omitempty"`
	Os                       string                       `json:"os,omitempty"`
	AppleSearchAdsAttributes map[string]map[string]string `json:"apple_search_ads_attributes,omitempty"`
}
