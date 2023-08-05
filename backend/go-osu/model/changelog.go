package model

type Build struct {
	ID               int              `json:"id"`
	CreatedAt        DateTime         `json:"created_at"`
	DisplayVersion   string           `json:"display_version"`
	UpdateStream     UpdateStream     `json:"update_stream"`
	Users            int              `json:"users"`
	Version          string           `json:"version"`
	ChangelogEntries []ChangelogEntry `json:"changelog_entries"`
}

type ChangelogEntry struct {
	ID                  int        `json:"id"`
	Category            string     `json:"category"`
	CreatedAt           DateTime   `json:"created_at"`
	GitHubPullRequestID int        `json:"github_pull_request_id"`
	GitHubURL           string     `json:"github_url"`
	IsMajor             bool       `json:"major"`
	Repository          string     `json:"repository"`
	Title               string     `json:"title"`
	Type                string     `json:"type"`
	URL                 string     `json:"url"`
	GitHubUser          GitHubUser `json:"github_user"`
	Message             string     `json:"message"`
	MessageHtml         string     `json:"message_html"`
}

type GitHubUser struct {
	ID          int    `json:"id"`
	DisplayName string `json:"display_name"`
	GitHubURL   string `json:"git_hub_url"`
	OsuUsername string `json:"osu_username"`
	UserID      int    `json:"user_id"`
	UserURL     string `json:"user_url"`
}

type UpdateStream struct {
	ID          int    `json:"id"`
	DisplayName string `json:"display_name"`
	IsFeatured  bool   `json:"is_featured"`
	Name        string `json:"name"`
}
