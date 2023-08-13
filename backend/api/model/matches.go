package model

import (
	"encoding/json"
	"project/backend/api/client/opts"
	"project/backend/api/enum"
)

type MultiplayerMatches struct {
	Cursor  opts.Cursor        `json:"cursor"`
	Matches []MultiplayerMatch `json:"matches"`
}

type MultiplayerMatchDetails struct {
	Match  MultiplayerMatch        `json:"match"`
	Events []MultiplayerMatchEvent `json:"events"`
	Users  []UserCompact           `json:"users"`
}

type MultiplayerMatchEvent struct {
	Detail    MultiplayerMatchEventType
	ID        int      `json:"id"`
	Timestamp DateTime `json:"timestamp"`

	// These fields are optional and are present depending on Detail.Type value.
	UserID    *int             `json:"user_id"`
	Game      *MultiplayerGame `json:"game"`
	MatchName *string          `json:"match_name"`
}

type MultiplayerMatchEventType struct {
	Type enum.MultiplayerMatchEventType `json:"type"`
}

type MultiplayerGame struct {
	ID          int                         `json:"id"`
	StartTime   DateTime                    `json:"start_time"`
	EndTime     *DateTime                   `json:"end_time"`
	Mode        enum.GameMode               `json:"mode"`
	ScoringType enum.MultiplayerScoringType `json:"scoring_type"`
	TeamType    enum.MultiplayerTeamType    `json:"team_type"`
	Mods        enum.Mods                   `json:"mods"`
	Beatmap     *BeatmapCompact             `json:"beatmap"`
	Scores      []MultiplayerMatchScore     `json:"scores"`
}

type MultiplayerMatchScore struct {
	Accuracy   float32              `json:"accuracy"`
	MaxCombo   int                  `json:"max_combo"`
	Mods       enum.Mods            `json:"mods"`
	IsPass     bool                 `json:"pass"`
	IsPerfect  IntBool              `json:"perfect"`
	Score      int                  `json:"score"`
	Slot       int                  `json:"slot"`
	Statistics ScoreStatistics      `json:"statistics"`
	Team       enum.MultiplayerTeam `json:"team"`
	UserID     int                  `json:"user_id"`
}

type MultiplayerMatch struct {
	ID        int       `json:"id"`
	StartTime DateTime  `json:"start_time"`
	EndTime   *DateTime `json:"end_time"`
	Name      string    `json:"name"`
}

type MatchParams struct {
	Limit int       `json:"limit"`
	Sort  enum.Sort `json:"sort"`
}

// IntBool is used for when API returns 0 or 1 instead of a boolean... for whatever reason.
type IntBool bool

func (b IntBool) Bool() bool {
	return bool(b)
}

func (b IntBool) UnmarshalJSON(bytes []byte) (err error) {
	var intVal int
	err = json.Unmarshal(bytes, &intVal)
	if err != nil {
		return
	}

	b = intVal == 1
	return
}
