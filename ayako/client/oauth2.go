package client

import (
	"github.com/pkg/errors"
	"gopkg.in/h2non/gentleman.v2/plugins/multipart"
)

type Oauth2API struct {
	*OsuAPI
}

type Oauth2Token struct {
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (b *Oauth2API) TokenRenew(scope string, refreshToken string) (*Oauth2Token, error) {
	json := Oauth2Token{}

	form := multipart.FormData{
		Data: map[string]multipart.Values{
			"refresh_token": []string{refreshToken},
			"scope":         []string{scope},
			"grant_type":    []string{"refresh_token"},
			"client_id":     []string{b.clientId},
			"client_secret": []string{b.clientSecret},
		},
	}

	req := b.client.
		Request().
		Method("POST").
		Path("/oauth/token").
		Form(form)

	res, err := req.Send()
	if err != nil {
		return nil, err
	}
	if !res.Ok {
		return nil, errors.Wrap(err, res.RawResponse.Status)
	}

	if err := res.JSON(&json); err != nil {
		return nil, err
	}

	return &json, nil
}
