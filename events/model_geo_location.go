 
package events
// GeoLocation struct for GeoLocation
type GeoLocation struct {
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Accuracy float64 `json:"accuracy,omitempty"`
}
