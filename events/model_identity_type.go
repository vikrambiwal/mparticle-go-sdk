package events

// IdentityType the model 'IdentityType'
type IdentityType string

// List of IdentityType
const (
	Other                    IdentityType = "other"
	CustomerID               IdentityType = "customer_id"
	Facebook                 IdentityType = "facebook"
	Twitter                  IdentityType = "twitter"
	Google                   IdentityType = "google"
	Microsoft                IdentityType = "microsoft"
	Yahoo                    IdentityType = "yahoo"
	Email                    IdentityType = "email"
	Alias                    IdentityType = "alias"
	FacebookCustomAudienceID IdentityType = "facebook_custom_audience_id"
	OtherID2                 IdentityType = "other_id_2"
	OtherID3                 IdentityType = "other_id_3"
	OtherID4                 IdentityType = "other_id_4"
)
