package model

import (
	"time"
)

// UserStats represents a user stats history entry
type UserStats struct {
	Count300    int       `json:"count300"`
	Count100    int       `json:"count100"`
	Count50     int       `json:"count50"`
	PlayCount   int       `json:"playcount"`
	RankedScore string    `json:"ranked_score"`
	TotalScore  string    `json:"total_score"`
	PPRank      int       `json:"pp_rank"`
	Level       float64   `json:"level"`
	PPRaw       float64   `json:"pp_raw"`
	Accuracy    float64   `json:"accuracy"`
	CountRankSS int       `json:"count_rank_ss"`
	CountRankS  int       `json:"count_rank_s"`
	CountRankA  int       `json:"count_rank_a"`
	Timestamp   time.Time `json:"timestamp"`
}
