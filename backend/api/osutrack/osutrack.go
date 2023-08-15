package osutrack

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"project/backend/api/model"
	"time"
)

const (
	baseURL = "https://osutrack-api.ameo.dev"
)

// Client is a client for the osu!track API
type Client struct {
	httpClient *http.Client
}

// NewClient creates a new osu!track API client
func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{},
	}
}

// UpdateUser updates a user's stats for a specific mode
func (c *Client) UpdateUser(userID int, mode int) (*model.UserUpdate, error) {
	u := fmt.Sprintf("%s/update?user=%d&mode=%d", baseURL, userID, mode)

	req, err := http.NewRequest("POST", u, nil)
	if err != nil {
		return nil, err
	}

	var resp model.UserUpdate
	if err := c.doRequest(req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// GetStatsHistory gets a user's stats history for a mode between two dates
func (c *Client) GetStatsHistory(userID int, mode int, from, to time.Time) ([]model.UserStats, error) {
	u := fmt.Sprintf("%s/stats_history?user=%d&mode=%d", baseURL, userID, mode)
	v := url.Values{}
	if !from.IsZero() {
		v.Set("from", from.Format("2006-01-02"))
	}
	if !to.IsZero() {
		v.Set("to", to.Format("2006-01-02"))
	}
	u += "&" + v.Encode()

	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var resp []model.UserStats
	if err := c.doRequest(req, &resp); err != nil {
		return nil, err
	}

	return resp, nil
}

// GetHiscores gets a user's hiscores for a mode between two dates
func (c *Client) GetHiscores(userID int, mode int, from, to time.Time) ([]model.UserHiscore, error) {
	u := fmt.Sprintf("%s/hiscores?user=%d&mode=%d", baseURL, userID, mode)
	v := url.Values{}
	if !from.IsZero() {
		v.Set("from", from.Format("2006-01-02"))
	}
	if !to.IsZero() {
		v.Set("to", to.Format("2006-01-02"))
	}
	u += "&" + v.Encode()

	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var resp []model.UserHiscore
	if err := c.doRequest(req, &resp); err != nil {
		return nil, err
	}

	return resp, nil
}

// GetPeak gets a user's peak rank and accuracy for a mode
func (c *Client) GetPeak(userID int, mode int) (*model.Peak, error) {
	u := fmt.Sprintf("%s/peak?user=%d&mode=%d", baseURL, userID, mode)

	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var resp []model.Peak
	if err := c.doRequest(req, &resp); err != nil {
		return nil, err
	}

	if len(resp) > 0 {
		return &resp[0], nil
	}

	return nil, nil
}

// GetBestPlays gets the best plays by pp for a mode between two dates
func (c *Client) GetBestPlays(mode int, from, to time.Time, limit int) ([]model.TopPlay, error) {
	u := fmt.Sprintf("%s/bestplays?mode=%d&limit=%d", baseURL, mode, limit)
	v := url.Values{}
	if !from.IsZero() {
		v.Set("from", from.Format("2006-01-02"))
	}
	if !to.IsZero() {
		v.Set("to", to.Format("2006-01-02"))
	}
	u += "&" + v.Encode()

	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var resp []model.TopPlay
	if err := c.doRequest(req, &resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) doRequest(req *http.Request, v interface{}) error {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	return nil
}
