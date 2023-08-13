package client

import (
	"project/backend/api/client/opts"
	"project/backend/api/enum"
	"project/backend/api/model"
)

// Search searches users and wiki pages.
func (c *OsuClient) Search(opts opts.SearchOpts) (*model.SearchResults, error) {
	return handleResponse[model.SearchResults](c.getReq(enum.EndpointSearch, opts))
}
