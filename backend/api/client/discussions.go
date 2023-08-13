package client

import (
	"project/backend/api/client/opts"
	"project/backend/api/enum"
	"project/backend/api/model"
)

// GetBeatmapsetDiscussions returns a list of beatmapset discussions
func (c *OsuClient) GetBeatmapsetDiscussions(opts opts.GetBeatmapsetDiscussionsOpts) (*model.Discussions, error) {
	return handleResponse[model.Discussions](c.getReq(enum.EndpointGetBeatmapsetDiscussions, opts))
}

// GetBeatmapsetDiscussionPosts returns the posts of beatmapset discussions.
func (c *OsuClient) GetBeatmapsetDiscussionPosts(opts opts.GetBeatmapsetDiscussionPostsOpts) (*model.DiscussionPosts, error) {
	return handleResponse[model.DiscussionPosts](c.getReq(enum.EndpointGetBeatmapsetDiscussionPosts, opts))
}

// GetBeatmapsetDiscussionVotes returns the votes given to beatmapset discussions.
func (c *OsuClient) GetBeatmapsetDiscussionVotes(opts opts.GetBeatmapsetDiscussionVotesOpts) (*model.DiscussionVotes, error) {
	return handleResponse[model.DiscussionVotes](c.getReq(enum.EndpointGetBeatmapsetDiscussionVotes, opts))
}
