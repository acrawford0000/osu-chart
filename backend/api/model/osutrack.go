package model

import (
	"time"
)

// Model of the osutrack-api response
type stats struct {
	Count300    int       `json:"count300"`
	Count100    int       `json:"count100"`
	Count50     int       `json:"count50"`
	Playcount   int       `json:"playcount"`
	RankedScore string    `json:"ranked_score"`
	TotalScore  string    `json:"total_score"`
	PpRank      int       `json:"pp_rank"`
	Level       float64   `json:"level"`
	PpRaw       float64   `json:"pp_raw"`
	Accuracy    float64   `json:"accuracy"`
	CountRankSS int       `json:"count_rank_ss"`
	CountRankS  int       `json:"count_rank_s"`
	CountRankA  int       `json:"count_rank_a"`
	Timestamp   time.Time `json:"timestamp"`
}
