package model

import (
	"time"
)

// UserUpdate represents the response from /update
type UserUpdate struct {
	Username    string        `json:"username"`
	Mode        int           `json:"mode"`
	PlayCount   int           `json:"playcount"`
	PPRank      int           `json:"pp_rank"`
	PPRaw       float64       `json:"pp_raw"`
	Accuracy    float64       `json:"accuracy"`
	TotalScore  int64         `json:"total_score"`
	RankedScore int64         `json:"ranked_score"`
	Count300    int           `json:"count300"`
	Count100    int           `json:"count100"`
	Count50     int           `json:"count50"`
	Level       float64       `json:"level"`
	CountRankA  int           `json:"count_rank_a"`
	CountRankS  int           `json:"count_rank_s"`
	CountRankSS int           `json:"count_rank_ss"`
	LevelUp     bool          `json:"levelup"`
	First       bool          `json:"first"`
	Exists      bool          `json:"exists"`
	NewHiscores []UserHiscore `json:"newhs"`
}

// UserStats represents a user stats history entry
type UserStats struct {
	Count300    int       `json:"count300"`
	Count100    int       `json:"count100"`
	Count50     int       `json:"count50"`
	PlayCount   int       `json:"playcount"`
	RankedScore int64     `json:"ranked_score"`
	TotalScore  int64     `json:"total_score"`
	PPRank      int       `json:"pp_rank"`
	Level       float64   `json:"level"`
	PPRaw       float64   `json:"pp_raw"`
	Accuracy    float64   `json:"accuracy"`
	CountRankSS int       `json:"count_rank_ss"`
	CountRankS  int       `json:"count_rank_s"`
	CountRankA  int       `json:"count_rank_a"`
	Timestamp   time.Time `json:"timestamp"`
}

// UserHiscore represents a user hiscore
type UserHiscore struct {
	BeatmapID  int       `json:"beatmap_id"`
	Score      int64     `json:"score"`
	PP         float64   `json:"pp"`
	Mods       int       `json:"mods"`
	Rank       string    `json:"rank"`
	ScoreTime  time.Time `json:"score_time"`
	UpdateTime time.Time `json:"update_time"`
}

// Peak contains a user's peak rank and accuracy
type Peak struct {
	BestGlobalRank int     `json:"best_global_rank"`
	BestAccuracy   float64 `json:"best_accuracy"`
}

// TopPlay represents a top play by pp
type TopPlay struct {
	User       int64     `json:"user"`
	BeatmapID  int       `json:"beatmap_id"`
	Score      int64     `json:"score"`
	PP         float64   `json:"pp"`
	Mods       int       `json:"mods"`
	Rank       string    `json:"rank"`
	ScoreTime  time.Time `json:"score_time"`
	UpdateTime time.Time `json:"update_time"`
}
