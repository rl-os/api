package services

// OAuthToken model
type OAuthToken struct {
	Id           int    `db:"id"`
	UserId       int    `db:"user_id"`
	AccessToken  string `db:"access_token"`
	RefreshToken string `db:"refresh_token"`
	Scopes       string `db:"scopes"`
}

// CreateOAuthToken and return model with valid access_token and refresh_token
func CreateOAuthToken(userId int, clientId string, secretId string, scopes string) (OAuthToken, err error) {
	// todo: verify client_id and client_secret
	return nil, nil
}

// RevokeOAuthToken and return error if not successfully
func RevokeOAuthToken() (err error) {
	return nil
}

// ValidateOAuthToken and return OAuthToken with full information
func ValidateOAuthToken(accessToken string) (OAuthToken, err error) {
	return nil, nil
}

// RefreshOAuthToken create new access_token and return it
func RefreshOAuthToken(refreshToken string) (OAuthToken, err error) {
	return nil, nil
}
