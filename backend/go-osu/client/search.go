package client

import (
	"project/backend/go-osu/client/opts"
	"project/backend/go-osu/enum"
	"project/backend/go-osu/model"
)

// Search searches users and wiki pages.
func (c *OsuClient) Search(opts opts.SearchOpts) (*model.SearchResults, error) {
	return handleResponse[model.SearchResults](c.getReq(enum.EndpointSearch, opts))
}
