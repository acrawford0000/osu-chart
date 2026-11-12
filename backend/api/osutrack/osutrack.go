package osutrack

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"project/backend/api/apierror"
	"project/backend/api/model"
	"time"
)

var historyClient = &http.Client{Timeout: 20 * time.Second}

func GetStatsUpdates(user int, mode int) ([]*model.UserStats, error) {
	return getStatsUpdates(historyClient, "https://osutrack-api.ameo.dev", user, mode)
}

func getStatsUpdates(httpClient *http.Client, baseURL string, user, mode int) ([]*model.UserStats, error) {
	updateURL := fmt.Sprintf("%s/update?user=%d&mode=%d", baseURL, user, mode)
	response, err := httpClient.Post(updateURL, "application/json", nil)
	if err != nil {
		return nil, apierror.Network("osu!track", err)
	}
	// Drain and close the update response before requesting history, including errors.
	_, readErr := io.Copy(io.Discard, io.LimitReader(response.Body, 1<<20))
	response.Body.Close()
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return nil, apierror.Status("osu!track", response.StatusCode)
	}
	if readErr != nil {
		return nil, apierror.Network("osu!track", readErr)
	}

	url := fmt.Sprintf("%s/stats_history?user=%d&mode=%d", baseURL, user, mode)
	response, err = httpClient.Get(url)
	if err != nil {
		return nil, apierror.Network("osu!track", err)
	}
	defer response.Body.Close()
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return nil, apierror.Status("osu!track", response.StatusCode)
	}
	body, err := io.ReadAll(io.LimitReader(response.Body, 32<<20))
	if err != nil {
		return nil, apierror.Network("osu!track", err)
	}
	var stats []*model.UserStats
	if err = json.Unmarshal(body, &stats); err != nil {
		return nil, apierror.New("INVALID_RESPONSE", "osu!track returned unreadable history. Please retry.")
	}
	if stats == nil {
		stats = []*model.UserStats{}
	}
	return stats, nil
}
