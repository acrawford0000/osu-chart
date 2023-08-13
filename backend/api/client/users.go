package client

import (
	"project/backend/api/client/opts"
	"project/backend/api/enum"
	"project/backend/api/model"
)

// GetUserKudosu returns kudosu history of specified user.
func (c *OsuClient) GetUserKudosu(opts opts.GetUserKudosuOpts) ([]model.KudosuHistory, error) {
	history, err := handleResponse[[]model.KudosuHistory](c.getReq(enum.EndpointGetUserKudosu, opts))
	if err != nil {
		return nil, err
	}

	return *history, nil
}

// GetUserScores returns the scores of specified user.
func (c *OsuClient) GetUserScores(opts opts.GetUserScoresOpts) ([]model.Score, error) {
	scores, err := handleResponse[[]model.Score](c.getReq(enum.EndpointGetUserScores, opts))
	if err != nil {
		return nil, err
	}

	return *scores, nil
}

// GetUserBeatmaps returns the beatmaps of specified user.
func (c *OsuClient) GetUserBeatmaps(opts opts.GetUserBeatmapsOpts) ([]model.Beatmapset, error) {
	beatmapsets, err := handleResponse[[]model.Beatmapset](c.getReq(enum.EndpointGetUserBeatmaps, opts))
	if err != nil {
		return nil, err
	}

	return *beatmapsets, nil
}

// GetUserRecentActivity returns recent activity of specified user.
func (c *OsuClient) GetUserRecentActivity(opts opts.GetUserRecentActivityOpts) ([]model.Event, error) {
	events, err := handleResponse[[]model.Event](c.getReq(enum.EndpointGetUserRecentActivity, opts))
	if err != nil {
		return nil, err
	}

	return *events, nil
}

// GetUser returns the details of specified user.
func (c *OsuClient) GetUser(opts opts.GetUserOpts) (*model.User, error) {
	return handleResponse[model.User](c.getReq(enum.EndpointGetUser, opts))
}
