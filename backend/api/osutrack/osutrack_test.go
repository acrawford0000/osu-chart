package osutrack

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestHistoryRequestsKeepPlayerAndMode(t *testing.T) {
	var paths []string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("user") != "42" || r.URL.Query().Get("mode") != "3" {
			t.Error("wrong player or mode")
		}
		paths = append(paths, r.URL.Path)
		if r.URL.Path == "/update" {
			fmt.Fprint(w, `{}`)
		} else {
			fmt.Fprint(w, `[]`)
		}
	}))
	defer server.Close()
	result, err := getStatsUpdates(server.Client(), server.URL, 42, 3)
	if err != nil || result == nil || len(result) != 0 {
		t.Fatalf("empty history should be a successful empty slice: %v", err)
	}
	if len(paths) != 2 || paths[0] != "/update" || paths[1] != "/stats_history" {
		t.Fatal("incorrect request sequence")
	}
}
func TestHistoryHTTPAndParseFailures(t *testing.T) {
	for _, test := range []struct {
		name, path string
		status     int
		body, want string
	}{
		{"update failure", "/update", 503, `{}`, "SERVICE_UNAVAILABLE"},
		{"history rate limit", "/stats_history", 429, `{}`, "RATE_LIMITED"},
		{"bad history", "/stats_history", 200, `invalid`, "INVALID_RESPONSE"},
	} {
		t.Run(test.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path == test.path {
					w.WriteHeader(test.status)
					fmt.Fprint(w, test.body)
				} else {
					fmt.Fprint(w, `{}`)
				}
			}))
			defer server.Close()
			if _, err := getStatsUpdates(server.Client(), server.URL, 1, 0); err == nil || !strings.Contains(err.Error(), test.want) {
				t.Fatalf("unexpected error %v", err)
			}
		})
	}
}
func TestHistoryTimeoutIsActionable(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { time.Sleep(40 * time.Millisecond); fmt.Fprint(w, `{}`) }))
	defer server.Close()
	c := server.Client()
	c.Timeout = 5 * time.Millisecond
	if _, err := getStatsUpdates(c, server.URL, 1, 0); err == nil || !strings.Contains(err.Error(), "NETWORK_TIMEOUT") {
		t.Fatalf("unexpected error %v", err)
	}
}
