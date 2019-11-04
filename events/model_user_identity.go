 
package events
// UserIdentity struct for UserIdentity
type UserIdentity struct {
	IdentityType IdentityType `json:"identity_type"`
	Identity string `json:"identity"`
	TimestampUnixtimeMS int64 `json:"timestamp_unixtime_ms"`
	CreatedThisBatch bool `json:"created_this_batch"`
}
