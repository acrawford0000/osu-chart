package model

import "project/backend/api/enum"

type BeatmapCompact struct {
	ID          int               `json:"id"`
	SR          float32           `json:"difficulty_rating"`
	Status      enum.RankedStatus `json:"status"`
	TotalLength int               `json:"total_length"`
	CreatorID   int               `json:"user_id"`
	Version     string            `json:"version"`
	MaxCombo    int               `json:"max_combo"` // TODO: Fix this. Zero value after unmarshalling.
	FailTimes   *FailTimes        `json:"failtimes"`
	Gamemode    enum.GameMode     `json:"mode"`
	Checksum    string            `json:"checksum"`
	Beatmapset  Beatmapset        `json:"beatmapset"`
}

type Beatmaps struct {
	Beatmaps []BeatmapCompact `json:"beatmaps"`
}

type FailTimes struct {
	Exit []int `json:"exit"`
	Fail []int `json:"fail"`
}

type Beatmap struct {
	BeatmapCompact
	Accuracy      float32           `json:"accuracy"`
	AR            float32           `json:"ar"`
	BeatmapsetID  int               `json:"beatmapset_id"`
	BPM           float32           `json:"bpm"`
	IsConvert     bool              `json:"convert"`
	CountCircles  int               `json:"count_circles"`
	CountSliders  int               `json:"count_sliders"`
	CountSpinners int               `json:"count_spinners"`
	CS            float32           `json:"cs"`
	DeletedAt     DateTime          `json:"deleted_at"`
	Drain         float32           `json:"drain"`
	HitLength     int               `json:"hit_len gth"`
	IsScoreable   bool              `json:"is_scoreable"`
	LastUpdated   DateTime          `json:"last_updated"`
	Passcount     int               `json:"passcount"`
	Playcount     int               `json:"playcount"`
	RankedStatus  enum.RankedStatus `json:"ranked"` // TODO: Fix this. Always pending for some reason.
	URL           string            `json:"url"`
}

type BeatmapDifficultyAttributes struct {
	MaxCombo             int     `json:"max_combo"`
	StarRating           float64 `json:"star_rating"`
	ApproachRate         float64 `json:"approach_rate"`
	GreatHitWindow       float64 `json:"great_hit_window"`
	AimDifficulty        float64 `json:"aim_difficulty"`
	FlashlightDifficulty float64 `json:"flashlight_difficulty"`
	OverallDifficulty    float64 `json:"overall_difficulty"`
	SliderFactor         float64 `json:"slider_factor"`
	SpeedDifficulty      float64 `json:"speed_difficulty"`
	StaminaDifficulty    float64 `json:"stamina_difficulty"`
	RhythmDifficulty     float64 `json:"rhythm_difficulty"`
	ColourDifficulty     float64 `json:"colour_difficulty"`
	ScoreMultiplier      float64 `json:"score_multiplier"`
}

type BeatmapDifficultyAttributesContainer struct {
	Attributes BeatmapDifficultyAttributes `json:"attributes"`
}

type BeatmapPlaycount struct {
	BeatmapID  int               `json:"beatmap_id"`
	Beatmap    BeatmapCompact    `json:"beatmap"`
	Beatmapset BeatmapsetCompact `json:"beatmapset"`
	Count      int               `json:"count"`
}

type BeatmapScores struct {
	Scores    []Score            `json:"scores"`
	UserScore []BeatmapUserScore `json:"userScore"`
}

type BeatmapUserScore struct {
	Position int   `json:"position"`
	Score    Score `json:"score"`
}

type BeatmapsetCompact struct {
	ID             int       `json:"id"`
	Artist         string    `json:"artist"`
	ArtistUnicode  string    `json:"artist_unicode"`
	Covers         Covers    `json:"covers"`
	Creator        string    `json:"creator"`
	FavouriteCount int       `json:"favourite_count"`
	IsNSFW         bool      `json:"nsfw"`
	Playcount      int       `json:"play_count"`
	PreviewURL     string    `json:"preview_url"`
	Source         string    `json:"source"`
	Status         string    `json:"status"`
	Title          string    `json:"title"`
	TitleUnicode   string    `json:"title_unicode"`
	UserID         int       `json:"user_id"`
	HasVideo       bool      `json:"video"`
	Beatmaps       []Beatmap `json:"beatmaps"`
	User           User      `json:"user"`
	Genre          string    `json:"genre"`
	Description    string    `json:"description"`
}

type Covers struct {
	Cover       string `json:"cover"`
	Cover2x     string `json:"cover@2x"`
	Card        string `json:"card"`
	Card2x      string `json:"card@2x"`
	List        string `json:"list"`
	List2x      string `json:"list@2x"`
	SlimCover   string `json:"slimcover"`
	SlimCover2x string `json:"slimcover@2x"`
}

type Beatmapset struct {
	BeatmapsetCompact
	Availability      Availability      `json:"availability"`
	BPM               float32           `json:"bpm"`
	CanBeHyped        bool              `json:"can_be_hyped"`
	DiscussionEnabled bool              `json:"discussion_enabled"`
	DiscussionLocked  bool              `json:"discussion_locked"`
	Hype              Hype              `json:"hype"`
	IsScoreable       bool              `json:"is_scoreable"`
	LastUpdated       DateTime          `json:"last_updated"`
	LegacyThreadUrl   string            `json:"legacy_thread_url"`
	Nominations       Nominations       `json:"nominations"`
	RankedStatus      enum.RankedStatus `json:"ranked"`
	RankedDate        DateTime          `json:"ranked_date"`
	HasStoryboard     bool              `json:"storyboard"`
	SubmittedDate     DateTime          `json:"submitted_date"`
	Tags              string            `json:"tags"`
}

type Nominations struct {
	Current  int `json:"current"`
	Required int `json:"required"`
}

type Hype struct {
	Current  int `json:"current"`
	Required int `json:"required"`
}

type Availability struct {
	DownloadDisabled bool   `json:"download_disabled"`
	MoreInformation  string `json:"more_information"`
}
