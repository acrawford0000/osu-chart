package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"project/backend/go-osu/enum"
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
		return fmt.Errorf("unable to retrieve Token from POST request: %v", err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("unable to read response body from POST request: %v", err)
	}

	err = json.Unmarshal(body, &c.token)
	if err != nil {
		return fmt.Errorf("unable to unmarshal Token JSON string to Token struct: %v", err)
	}

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
