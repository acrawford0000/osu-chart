package model

type WikiPage struct {
	AvailableLocales []string `json:"available_locales"`
	Layout           string   `json:"layout"`
	Locale           string   `json:"locale"`
	Markdown         string   `json:"markdown"`
	Path             string   `json:"path"`
	Subtitle         *string  `json:"subtitle"`
	Tags             []string `json:"tags"`
	Title            string   `json:"title"`
}

type SearchResults struct {
	User     *SearchResult[UserCompact] `json:"user"`
	WikiPage *SearchResult[WikiPage]    `json:"wiki_page"`
}

type SearchResult[T UserCompact | WikiPage] struct {
	Data  []T `json:"data"`
	Total int `json:"total"`
}
