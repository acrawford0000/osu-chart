package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"project/backend/api/apierror"
	"project/backend/api/enum"
	"project/backend/api/model"
	"time"
)

// Uses model.UserCredentials to access osu! API v2. Assigns token to the token field in OsuClient.
func (c *OsuClient) login() error {
	tokenReqBytes, err := json.Marshal(c.userCredentials)
	if err != nil {
		return fmt.Errorf("unable to marshal TokenRequest: %v", err)
	}

	res, err := c.httpClient.Post(enum.BaseEndpointOAuth.String(), "application/json", bytes.NewBuffer(tokenReqBytes))
	if err != nil {
		return apierror.Network("osu!", err)
	}

	defer res.Body.Close()
	if res.StatusCode == 400 {
		return apierror.New("AUTH_INVALID", "osu! rejected these credentials. Check your Client ID and Client Secret in Settings.")
	}
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return apierror.Status("osu!", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("unable to read response body from POST request: %v", err)
	}

	var token model.Token
	err = json.Unmarshal(body, &token)
	if err != nil {
		return apierror.New("INVALID_RESPONSE", "osu! returned an unreadable login response. Please retry.")
	}
	if token.AccessToken.GetActualValue() == "" || token.ExpiresIn <= 0 {
		return apierror.New("AUTH_INVALID", "osu! did not issue an access token. Check your credentials in Settings.")
	}
	c.token = token

	c.loginExpiration = time.Now().Unix() + c.token.ExpiresIn

	return nil
}

func (c *OsuClient) loginIfTokenExpired() error {
	if c.loginExpiration < time.Now().Unix() {
		if err := c.login(); err != nil {
			return err
		}
	}

	return nil
}
