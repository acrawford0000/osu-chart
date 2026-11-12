package apierror

import (
	"errors"
	"fmt"
	"net"
)

// Errors cross Wails as strings. A stable prefix lets the UI distinguish failures
// without exposing server response bodies, tokens, or filesystem details.
func New(code, message string) error { return fmt.Errorf("%s: %s", code, message) }

func Network(service string, err error) error {
	var networkError net.Error
	if errors.As(err, &networkError) && networkError.Timeout() {
		return New("NETWORK_TIMEOUT", service+" did not respond in time. Please try again.")
	}
	return New("NETWORK_ERROR", "Could not reach "+service+". Check your connection and try again.")
}

func Status(service string, status int) error {
	switch status {
	case 400:
		return New("INVALID_REQUEST", "The request was rejected. Check the player name and game mode.")
	case 401, 403:
		return New("AUTH_INVALID", "The API request was rejected. Check your credentials in Settings.")
	case 404:
		return New("NOT_FOUND", "No matching player or history was found on "+service+".")
	case 429:
		return New("RATE_LIMITED", service+" is receiving too many requests. Please wait and retry.")
	default:
		return New("SERVICE_UNAVAILABLE", service+" returned an error. Please try again later.")
	}
}
