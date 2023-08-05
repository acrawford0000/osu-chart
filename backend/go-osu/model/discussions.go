package model

import (
	"project/backend/go-osu/client/opts"
	"project/backend/go-osu/enum"
)

type Discussions struct {
	Beatmaps            []Beatmap              `json:"beatmaps"`
	CursorString        *opts.CursorString     `json:"cursor_string"`
	Discussions         []BeatmapsetDiscussion `json:"discussions"`
	IncludedDiscussions []BeatmapsetDiscussion `json:"included_discussions"`
	ReviewsConfig       ReviewsConfig          `json:"reviews_config"`
	Users               []UserCompact          `json:"users"`
}

type DiscussionPosts struct {
	Beatmapsets  []BeatmapsetCompact        `json:"beatmapsets"`
	CursorString *opts.CursorString         `json:"cursor_string"`
	Posts        []BeatmapsetDiscussionPost `json:"posts"`
	Users        []UserCompact              `json:"users"`
}

type DiscussionVotes struct {
	CursorString *opts.CursorString         `json:"cursor_string"`
	Discussions  []BeatmapsetDiscussion     `json:"discussions"`
	Users        []UserCompact              `json:"users"`
	Votes        []BeatmapsetDiscussionVote `json:"votes"`
}

// ReviewsConfig is mostly undocumented. Used in beatmapset discussions.
type ReviewsConfig struct {
	MaxBlocks int `json:"max_blocks"`
}

type BeatmapsetDiscussion struct {
	ID                    int                             `json:"id"`
	Beatmap               BeatmapCompact                  `json:"beatmap"`
	BeatmapID             int                             `json:"beatmap_id"`
	Beatmapset            BeatmapsetCompact               `json:"beatmapset"`
	CanBeResolved         bool                            `json:"can_be_resolved"`
	CanGrantKudosu        bool                            `json:"can_grant_kudosu"`
	CreatedAt             DateTime                        `json:"created_at"`
	CurrentUserAttributes BeatmapsetDiscussionPermissions `json:"current_user_attributes"`
	DeletedAt             DateTime                        `json:"deleted_at"`
	DeletedByID           int                             `json:"deleted_by_id"`
	KudosuDenied          bool                            `json:"kudosu_denied"`
	LastPostAt            DateTime                        `json:"last_post_at"`
	MessageType           enum.MessageType                `json:"message_type"`
	ParentID              int                             `json:"parent_id"`
	Posts                 []BeatmapsetDiscussionPost      `json:"posts"`
	IsResolved            bool                            `json:"resolved"`
	StartingPost          BeatmapsetDiscussionPost        `json:"starting_post"`
	Timestamp             DateTime                        `json:"timestamp"`
	UpdatedAt             DateTime                        `json:"updated_at"`
	UserID                int                             `json:"user_id"`
}

type BeatmapsetDiscussionPost struct {
	ID                     int      `json:"id"`
	BeatmapsetDiscussionID int      `json:"beatmapset_discussion_id"`
	CreatedAt              DateTime `json:"created_at"`
	DeletedAt              DateTime `json:"deleted_at"`
	DeletedByID            int      `json:"deleted_by_id"`
	LastEditorID           int      `json:"last_editor_id"`
	Message                string   `json:"message"`
	IsSystemPost           bool     `json:"system"`
	UpdatedAt              DateTime `json:"updated_at"`
	UserID                 int      `json:"user_id"`
}

type BeatmapsetDiscussionVote struct {
	ID                     int      `json:"id"`
	BeatmapsetDiscussionID int      `json:"beatmapset_discussion_id"`
	CreatedAt              DateTime `json:"created_at"`
	Score                  int      `json:"score"`
	UpdatedAt              DateTime `json:"updated_at"`
	UserID                 int      `json:"user_id"`
}

type BeatmapsetDiscussionPermissions struct {
	CanDestroy        bool                     `json:"can_destroy"`
	CanReopen         bool                     `json:"can_reopen"`
	CanModerateKudosu bool                     `json:"can_moderate_kudosu"`
	CanResolve        bool                     `json:"can_resolve"`
	VoteScore         BeatmapsetDiscussionVote `json:"vote_score"`
}

type MessageType struct {
	Hype         int    `json:"hype"`
	MapperNote   string `json:"mapper_note"`
	IsPraise     bool   `json:"praise"`
	IsProblem    bool   `json:"problem"`
	IsReview     bool   `json:"review"`
	IsSuggestion bool   `json:"suggestion"`
}
