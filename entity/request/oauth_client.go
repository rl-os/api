package request

// CreateOAuthClient contain incoming data with user credentials
type CreateOAuthClient struct {
	Name     string `json:"name"`
	Redirect string `json:"redirect"`
}
