package model

import (
	"project/backend/api/client/opts"
	"project/backend/api/enum"
)

type ForumTopicAndPosts struct {
	Posts        []ForumPost        `json:"posts"`
	Topic        ForumTopic         `json:"topic"`
	CursorString *opts.CursorString `json:"cursor_string"`
	// Parameters used in the request.
	Search opts.GetTopicAndPostsOpts `json:"search"`
}

type ForumTopic struct {
	CreatedAt   DateTime            `json:"created_at"`
	DeletedAt   *DateTime           `json:"deleted_at"`
	UpdatedAt   DateTime            `json:"updated_at"`
	FirstPostID int                 `json:"first_post_id"`
	ForumID     int                 `json:"forum_id"`
	ID          int                 `json:"id"`
	IsLocked    bool                `json:"is_locked"`
	LastPostID  int                 `json:"last_post_id"`
	Poll        *Poll               `json:"poll"`
	PostCount   int                 `json:"post_count"`
	Title       string              `json:"title"`
	Type        enum.ForumTopicType `json:"type"`
	UserID      int                 `json:"user_id"`
}

type ForumPost struct {
	CreatedAt  DateTime       `json:"created_at"`
	DeletedAt  *DateTime      `json:"deleted_at"`
	EditedAt   *DateTime      `json:"edited_at"`
	EditedByID *int           `json:"edited_by_id"`
	ForumID    int            `json:"forum_id"`
	ID         int            `json:"id"`
	TopicID    int            `json:"topic_id"`
	UserID     int            `json:"user_id"`
	Body       *ForumPostBody `json:"body"`
}

type ForumPostBody struct {
	HTML string `json:"html"`
	Raw  string `json:"raw"`
}

type Poll struct {
	AllowVoteChange       bool         `json:"allow_vote_change"`
	EndedAt               *DateTime    `json:"ended_at"`
	HideIncompleteResults bool         `json:"hide_incomplete_results"`
	LastVoteAt            *DateTime    `json:"last_vote_at"`
	MaxVotes              int          `json:"max_votes"`
	Options               []PollOption `json:"options"`
	StartedAt             DateTime     `json:"started_at"`
	Title                 PollText     `json:"title"`
	TotalVoteCount        int          `json:"total_vote_count"`
}

type PollOption struct {
	ID        int      `json:"id"`
	Text      PollText `json:"text"`
	VoteCount *int     `json:"vote_count"`
}

type PollText struct {
	BBCode string `json:"bbcode"`
	HTML   string `json:"html"`
}
