package client

import (
	"project/backend/go-osu/client/opts"
	"project/backend/go-osu/enum"
	"project/backend/go-osu/model"
)

// GetTopicAndPosts returns a topic and its posts.
func (c *OsuClient) GetTopicAndPosts(opts opts.GetTopicAndPostsOpts) (*model.ForumTopicAndPosts, error) {
	return handleResponse[model.ForumTopicAndPosts](c.getReq(enum.EndpointGetTopicAndPosts, opts))
}
