package osutrack

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"project/backend/api/model"
)

func GetStatsUpdates(user int, mode int) ([]*model.UserStats, error) {
	// Create the URL for the Update User API endpoint
	updateURL := fmt.Sprintf("https://osutrack-api.ameo.dev/update?user=%d&mode=%d", user, mode)

	// Make the HTTP POST request to trigger an update on the server side
	_, err := http.Post(updateURL, "application/json", nil)
	if err != nil {
		return nil, err
	}

	// Create the URL for the API endpoint
	url := fmt.Sprintf("https://osutrack-api.ameo.dev/stats_history?user=%d&mode=%d", user, mode)

	// Make the HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON response into a slice of StatsUpdate structs
	var statsUpdates []*model.UserStats
	err = json.Unmarshal(body, &statsUpdates)
	if err != nil {
		return nil, err
	}

	return statsUpdates, nil
}
