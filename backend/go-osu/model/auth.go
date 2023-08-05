package model

import (
	"encoding/json"
	"project/backend/go-osu/enum"
)

type UserCredentials struct {
	ClientID     SensitiveInt    `json:"client_id"`
	ClientSecret SensitiveString `json:"client_secret"`
	GrantType    enum.GrantType  `json:"grant_type"`
	Scope        enum.Scope      `json:"scope"`
}

type Token struct {
	TokenType   string          `json:"token_type"`
	ExpiresIn   int64           `json:"expires_in"`
	AccessToken SensitiveString `json:"access_token"`
}

// SensitiveString is a type for sensitive data of type string. When logging something is needed, instead of logging the
// actual sensitive data, it converts it to a string of asterisks.
type SensitiveString string

func (s SensitiveString) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.GetActualValue())
}

func (SensitiveString) String() string {
	return "******"
}

func (s SensitiveString) GetActualValue() string {
	return string(s)
}

// SensitiveInt is a type for sensitive data of type int. When logging something is needed, instead of logging the
// actual sensitive data, it converts it to a string of asterisks.
type SensitiveInt int

func (s SensitiveInt) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.GetActualValue())
}

func (SensitiveInt) String() string {
	return "******"
}

func (s SensitiveInt) GetActualValue() int {
	return int(s)
}
