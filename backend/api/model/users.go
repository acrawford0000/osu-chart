package model

import "project/backend/api/enum"

type MyUserStruct struct {
	ID          int      `json:"id"`
	Username    string   `json:"username"`
	AvatarURL   string   `json:"avatar_url"`
	CountryCode string   `json:"country_code"`
	Country     Country  `json:"country"`
	CoverURL    string   `json:"cover_url"`
	Playstyle   []string `json:"playstyle"`
}

type UserCompact struct {
	ID                       int                    `json:"id"`
	Username                 string                 `json:"username"`
	AvatarURL                string                 `json:"avatar_url"`
	CountryCode              string                 `json:"country_code"`
	DefaultGroup             string                 `json:"default_group"`
	IsActive                 bool                   `json:"is_active"`
	IsBot                    bool                   `json:"is_bot"`
	IsDeleted                bool                   `json:"is_deleted"`
	IsOnline                 bool                   `json:"is_online"`
	IsSupporter              bool                   `json:"is_supporter"`
	LastVisit                *DateTime              `json:"last_visit"`
	FriendsOnlyPM            bool                   `json:"pm_friends_only"`
	ProfileColor             *string                `json:"profile_colour"`
	AccountHistory           []UserAccountHistory   `json:"account_history"`
	TournamentBanner         *TournamentBanner      `json:"active_tournament_banner"`
	Badges                   []UserBadge            `json:"badges"`
	Groups                   []UserGroups           `json:"groups"`
	BeatmapPlaycountsCount   int                    `json:"beatmap_playcounts_count"`
	FavoriteBeatmapsetCount  int                    `json:"favorite_beatmapset_count"`
	GraveyardBeatmapsetCount int                    `json:"graveyard_beatmapset_count"`
	LovedBeatmapsetCount     int                    `json:"loved_beatmapset_count"`
	PendingBeatmapsetCount   int                    `json:"pending_beatmapset_count"`
	RankedBeatmapsetCount    int                    `json:"ranked_beatmapset_count"`
	ReplaysWatchedCount      int                    `json:"replays_watched_count"`
	ScoresBestCount          int                    `json:"scores_best_count"`
	ScoresFirstCount         int                    `json:"scores_first_count"`
	ScoresRecentCount        int                    `json:"scores_recent_count"`
	FollowerCount            int                    `json:"follower_count"`
	Country                  Country                `json:"country"`
	Cover                    Cover                  `json:"cover"`
	MonthlyPlaycounts        []UserMonthlyPlaycount `json:"monthly_playcounts"`
	PreviousUsernames        []string               `json:"previous_usernames"`
	RankHistory              RankHistory            `json:"rank_history"`
	SupportLevel             int                    `json:"support_level"`
	Achievements             []UserAchievements     `json:"user_achievements"`
	Statistics               *UserStatistics        `json:"statistics"`
}

type User struct {
	UserCompact
	CoverURL     string        `json:"cover_url"`
	Discord      *string       `json:"discord"`
	HasSupported bool          `json:"has_supported"`
	Interests    *string       `json:"interests"`
	JoinDate     DateTime      `json:"join_date"`
	Kudosu       Kudosu        `json:"kudosu"`
	Location     *string       `json:"location"`
	MaxBlocks    *int          `json:"max_blocks"`
	MaxFriends   *int          `json:"max_friends"`
	Occupation   *string       `json:"occupation"`
	Playmode     enum.GameMode `json:"playmode"`
	Playstyle    []string      `json:"playstyle"`
	PostCount    int           `json:"post_count"`
	Title        *string       `json:"title"`
	TitleURL     *string       `json:"title_url"`
	Twitter      *string       `json:"twitter"`
	Website      *string       `json:"website"`
}

type Users struct {
	Users []UserCompact `json:"users"`
}

type Kudosu struct {
	Available int `json:"available"`
	Total     int `json:"total"`
}

type KudosuHistory struct {
	ID        int               `json:"id"`
	Action    enum.KudosuAction `json:"action"`
	Amount    int               `json:"amount"`
	Model     string            `json:"model"`
	CreatedAt DateTime          `json:"created_at"`
	Giver     *KudosuGiver      `json:"giver"`
	Post      KudosuPost        `json:"post"`
}

type KudosuGiver struct {
	URL      string `json:"url"`
	Username string `json:"username"`
}

type KudosuPost struct {
	URL   *string `json:"url"`
	Title string  `json:"title"`
}

type UserStatistics struct {
	GradeCounts            GradeCounts `json:"grade_counts"`
	Level                  Level       `json:"level"`
	HitAccuracy            float32     `json:"hit_accuracy"`
	IsRanked               bool        `json:"is_ranked"`
	MaxCombo               int         `json:"maximum_combo"`
	Playcount              int         `json:"play_count"`
	Playtime               int         `json:"play_time"`
	PP                     float32     `json:"pp"`
	GlobalRank             *int        `json:"global_rank"`
	RankedScore            int         `json:"ranked_score"`
	ReplaysWatchedByOthers int         `json:"replays_watched_by_others"`
	TotalHits              int         `json:"total_hits"`
	TotalScore             int         `json:"total_score"`
	User                   UserCompact `json:"user"`
}

type Level struct {
	Current  int `json:"current"`
	Progress int `json:"progress"`
}

type GradeCounts struct {
	A   int `json:"a"`
	S   int `json:"s"`
	SH  int `json:"sh"`
	SS  int `json:"ss"`
	SSH int `json:"ssh"`
}

type UserAchievements struct {
	AchievedAt    DateTime `json:"achieved_at"`
	AchievementID int      `json:"achievement_id"`
}

type Cover struct {
	CustomURL string `json:"custom_url"`
	URL       string `json:"url"`
}

type Country struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type RankHistory struct {
	Mode string `json:"mode"`
	Data []int  `json:"data"`
}

type UserMonthlyPlaycount struct {
	StartDate DateTime `json:"start_date"`
	Count     int      `json:"count"`
}

type UserGroups struct {
	Playmodes *[]string `json:"playmodes"`
}

type UserBadge struct {
	URL         string   `json:"url"`
	ImageURL    string   `json:"image_url"`
	Description string   `json:"description"`
	AwardedAt   DateTime `json:"awarded_at"`
}

type TournamentBanner struct {
	ID           int    `json:"id"`
	TournamentID int    `json:"tournament_id"`
	Image        string `json:"image"`
}

type UserAccountHistory struct {
	ID          int      `json:"id"`
	Description *string  `json:"description"`
	Length      int      `json:"length"`
	Timestamp   DateTime `json:"timestamp"`
	Type        string   `json:"type"`
}

type Score struct {
	Mods        enum.Mods       `json:"mods"`
	ID          int64           `json:"id"`
	BestID      *int            `json:"best_id"`
	UserID      int             `json:"user_id"`
	Accuracy    float32         `json:"accuracy"`
	Score       int64           `json:"score"`
	MaxCombo    int             `json:"max_combo"`
	IsPerfect   bool            `json:"perfect"`
	IsPass      bool            `json:"passed"`
	Statistics  ScoreStatistics `json:"statistics"`
	PP          float32         `json:"pp"`
	Rank        enum.Grade      `json:"rank"`
	CreatedAt   *DateTime       `json:"created_at"`
	Mode        enum.GameMode   `json:"mode"`
	ModeInt     int             `json:"mode_int"`
	HasReplay   bool            `json:"replay"`
	Beatmap     *Beatmap        `json:"beatmap"`
	Beatmapset  *Beatmapset     `json:"beatmapset"`
	RankCountry *int            `json:"rank_country"`
	RankGlobal  *int            `json:"rank_global"`
	Weight      *Weight         `json:"weight"`
	User        *UserCompact    `json:"user"`
}

type Scores struct {
	Scores []Score `json:"scores"`
}

type Weight struct {
	Percentage float32 `json:"percentage"`
	PP         float32 `json:"pp"`
}

type ScoreStatistics struct {
	Count50   int `json:"count_50"`
	Count100  int `json:"count_100"`
	Count300  int `json:"count_300"`
	CountGeki int `json:"count_geki"`
	CountKatu int `json:"count_katu"`
	CountMiss int `json:"count_miss"`
}

type Event struct {
	ID        int      `json:"id"`
	CreatedAt DateTime `json:"created_at"`
	Type      string   `json:"type"`
	*AchievementEvent
	*BeatmapPlaycountEvent
	*BeatmapsetApproveEvent
	*BeatmapsetDeleteEvent
	*BeatmapsetUpdateEvent
	*RankEvent
	*RankLostEvent
	*UserActivityEvent
}

type AchievementEvent struct {
	ID              int           `json:"id"`
	AchievementName string        `json:"name"`
	Description     string        `json:"description"`
	Grouping        string        `json:"grouping"`
	Ordering        int           `json:"ordering"`
	Mode            enum.GameMode `json:"mode"`
	Slug            string        `json:"slug"`
	User            EventUser     `json:"user"`
	Instructions    string        `json:"instructions"`
}

type BeatmapPlaycountEvent struct {
	Count   int          `json:"count"`
	Beatmap EventBeatmap `json:"beatmap"`
}

type BeatmapsetApproveEvent struct {
	Approval   string       `json:"approval"`
	Beatmapset EventBeatmap `json:"beatmapset"`
	User       EventUser    `json:"user"`
}

type BeatmapsetDeleteEvent struct {
	Beatmapset EventBeatmap `json:"beatmapset"`
}

// BeatmapsetUpdateEvent is also used for revive and upload events.
type BeatmapsetUpdateEvent struct {
	Beatmapset EventBeatmap `json:"beatmapset"`
	User       EventUser    `json:"user"`
}

type RankEvent struct {
	ScoreRank string        `json:"score_rank"`
	Rank      int           `json:"rank"`
	Mode      enum.GameMode `json:"mode"`
	Beatmap   EventBeatmap  `json:"beatmap"`
	User      EventUser     `json:"user"`
}

type RankLostEvent struct {
	Mode       enum.GameMode `json:"mode"`
	Beatmapset EventBeatmap  `json:"beatmapset"`
	User       EventUser     `json:"user"`
}

// UserActivityEvent is used for all supporter events (first, again, and gift) and username change event.
type UserActivityEvent struct {
	User EventUser `json:"user"`
}

// EventBeatmap is also used for beatmapset events.
type EventBeatmap struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

type EventUser struct {
	Username         string  `json:"username"`
	URL              string  `json:"url"`
	PreviousUsername *string `json:"previousUsername"`
}

// Group is not returned by any endpoint. However, it might be in the future since a user default group identifier is returned by UserCompact.
type Group struct {
	ID             int     `json:"id"`
	Name           string  `json:"name"`
	ShortName      string  `json:"short_name"`
	HasListing     bool    `json:"has_listing"`
	HasPlaymodes   bool    `json:"has_playmodes"`
	Identifier     string  `json:"identifier"`
	IsProbationary bool    `json:"is_probationary"`
	Color          *string `json:"colour"`
}
