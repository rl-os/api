package request

// CreateOauthToken contain incoming data with user credentials
type CreateOauthToken struct {
	GrantType    string `json:"grant_type" form:"grant_type" validate:"required"`
	ClientID     uint   `json:"client_id" form:"client_id" validate:"required"`
	ClientSecret string `json:"client_secret" form:"client_secret" validate:"required"`
	Scope        string `json:"scope" form:"scope" validate:"required"`

	// grant_type == password
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`

	// grant_type == refresh_token
	RefreshToken string `json:"refresh_token" form:"refresh_token"`
}
