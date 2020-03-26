package client

import (
	"bytes"
	"github.com/pkg/errors"
	"mime/multipart"
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

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)

	_ = writer.WriteField("client_id", b.clientId)
	_ = writer.WriteField("client_secret", b.clientSecret)
	_ = writer.WriteField("scope", scope)
	_ = writer.WriteField("grant_type", "refresh_token")
	_ = writer.WriteField("refresh_token", refreshToken)
	if err := writer.Close(); err != nil {
		return nil, errors.Wrap(err, "setting up form data")
	}

	req := b.client.
		Request().
		Method("POST").
		Path("/oauth/token").
		SetHeader("Content-Type", "application/json").
		SetHeader("Content-Type", writer.FormDataContentType()).
		Body(payload)

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
