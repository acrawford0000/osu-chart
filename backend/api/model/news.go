package model

import (
	"project/backend/api/client/opts"
	"project/backend/api/enum"
)

type NewsListing struct {
	NewsPosts   []NewsPost  `json:"news_posts"`
	NewsSidebar NewsSidebar `json:"news_sidebar"`
	NewsSearch  NewsSearch  `json:"search"`
	Cursor      opts.Cursor `json:"cursor"`
}

type NewsPost struct {
	ID             int      `json:"id"`
	Slug           string   `json:"slug"`
	Title          string   `json:"title"`
	PublishedAt    DateTime `json:"published_at"`
	UpdatedAt      DateTime `json:"updated_at"`
	AuthorUsername string   `json:"author"`
	EditURL        string   `json:"edit_url"`
	FirstImage     *string  `json:"first_image"`
}

type NewsSidebar struct {
	CurrentYear int        `json:"current_year"`
	NewsPosts   []NewsPost `json:"news_posts"`
	Years       []int      `json:"years"`
}

type NewsSearch struct {
	Limit int       `json:"limit"`
	Sort  enum.Sort `json:"sort"`
}
