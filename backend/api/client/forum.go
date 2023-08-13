package client

import (
	"project/backend/api/client/opts"
	"project/backend/api/enum"
	"project/backend/api/model"
)

// GetTopicAndPosts returns a topic and its posts.
func (c *OsuClient) GetTopicAndPosts(opts opts.GetTopicAndPostsOpts) (*model.ForumTopicAndPosts, error) {
	return handleResponse[model.ForumTopicAndPosts](c.getReq(enum.EndpointGetTopicAndPosts, opts))
}
