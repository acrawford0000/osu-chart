package client

import (
	"io"
	"net/http"
	"project/backend/api/client/opts"
	"project/backend/api/enum"
	"project/backend/api/model"
	"strings"
	"testing"
	"time"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

type trackedBody struct {
	io.Reader
	closed bool
}

func (b *trackedBody) Close() error { b.closed = true; return nil }

func TestLoginRejectsErrorStatusesAndMissingTokens(t *testing.T) {
	for _, test := range []struct {
		name   string
		status int
		body   string
		want   string
	}{
		{"rejected", 401, `{"error":"invalid_client"}`, "AUTH_INVALID"},
		{"rate limited", 429, `{}`, "RATE_LIMITED"},
		{"server failure", 503, `{}`, "SERVICE_UNAVAILABLE"},
		{"no token", 200, `{"expires_in":3600}`, "AUTH_INVALID"},
		{"invalid JSON", 200, `not json`, "INVALID_RESPONSE"},
	} {
		t.Run(test.name, func(t *testing.T) {
			body := &trackedBody{Reader: strings.NewReader(test.body)}
			c := &OsuClient{httpClient: &http.Client{Transport: roundTripFunc(func(*http.Request) (*http.Response, error) {
				return &http.Response{StatusCode: test.status, Body: body}, nil
			})}}
			err := c.login()
			if err == nil || !strings.Contains(err.Error(), test.want) {
				t.Fatalf("expected %s, got %v", test.want, err)
			}
			if !body.closed {
				t.Fatal("response was not closed")
			}
			if c.token.AccessToken.GetActualValue() != "" {
				t.Fatal("failed login installed a token")
			}
		})
	}
}
func TestFailedProfileResponseClosesBody(t *testing.T) {
	body := &trackedBody{Reader: strings.NewReader(`{}`)}
	c := &OsuClient{token: model.Token{AccessToken: model.SensitiveString("test")}, loginExpiration: time.Now().Unix() + 3600, httpClient: &http.Client{Transport: roundTripFunc(func(*http.Request) (*http.Response, error) { return &http.Response{StatusCode: 404, Body: body}, nil })}}
	if _, err := c.doReq("GET", enum.EndpointGetUser, opts.GetUserOpts{Username: "demo"}); err == nil || !strings.Contains(err.Error(), "NOT_FOUND") {
		t.Fatalf("unexpected error %v", err)
	}
	if !body.closed {
		t.Fatal("error response leaked its body")
	}
}
