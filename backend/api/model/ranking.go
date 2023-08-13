package model

import "project/backend/api/client/opts"

type Rankings struct {
	Beatmapsets []Beatmapset     `json:"beatmapsets"`
	Ranking     []UserStatistics `json:"ranking"`
	Spotlight   *Spotlight       `json:"spotlight"`
	Total       int              `json:"total"`
	Cursor      opts.Cursor      `json:"cursor"`
}

type Spotlight struct {
	EndDate          DateTime `json:"end_date"`
	ID               int      `json:"id"`
	ModeSpecific     bool     `json:"mode_specific"`
	ParticipantCount *int     `json:"participant_count"`
	Name             string   `json:"name"`
	StartDate        DateTime `json:"start_date"`

	// TODO: No clue what the possible types are. Weakly documented. "yearly" example was given in the docs.
	// Probably yearly, monthly, weekly? Not sure.
	Type string `json:"type"`
}

type Spotlights struct {
	Spotlights []Spotlight `json:"spotlights"`
}
