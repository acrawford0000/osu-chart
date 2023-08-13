package client

import (
	"project/backend/api/client/opts"
	"project/backend/api/enum"
	"project/backend/api/model"
)

// LookupBeatmap looks up beatmaps depending on user input. See opts.GetBeatmapLookupOpts for more info.
func (c *OsuClient) LookupBeatmap(opts opts.GetBeatmapLookupOpts) (*model.Beatmap, error) {
	return handleResponse[model.Beatmap](c.getReq(enum.EndpointBeatmapLookup, opts))
}

// GetUserBeatmapScore returns user's top score on specified beatmap.
func (c *OsuClient) GetUserBeatmapScore(opts opts.GetUserBeatmapScoreOpts) (*model.BeatmapUserScore, error) {
	return handleResponse[model.BeatmapUserScore](c.getReq(enum.EndpointGetUserBeatmapScore, opts))
}

// GetUserBeatmapScores returns user's scores on specified beatmap.
func (c *OsuClient) GetUserBeatmapScores(opts opts.GetUserBeatmapScoresOpts) ([]model.Score, error) {
	userScores, err := handleResponse[model.Scores](c.getReq(enum.EndpointGetUserBeatmapScores, opts))
	if err != nil {
		return nil, err
	}

	return userScores.Scores, nil
}

// GetBeatmapScores returns the top scores for a beatmap.
func (c *OsuClient) GetBeatmapScores(opts opts.GetBeatmapScoresOpts) ([]model.Score, error) {
	beatmapScores, err := handleResponse[model.BeatmapScores](c.getReq(enum.EndpointGetBeatmapScores, opts))
	if err != nil {
		return nil, err
	}

	return beatmapScores.Scores, nil
}

// GetBeatmaps returns list of beatmaps.
func (c *OsuClient) GetBeatmaps(opts opts.GetBeatmapsOpts) ([]model.BeatmapCompact, error) {
	beatmaps, err := handleResponse[model.Beatmaps](c.getReq(enum.EndpointGetBeatmaps, opts))
	if err != nil {
		return nil, err
	}

	return beatmaps.Beatmaps, nil
}

// GetBeatmap returns beatmap data for the specified beatmap ID.
func (c *OsuClient) GetBeatmap(opts opts.GetBeatmapOpts) (*model.Beatmap, error) {
	return handleResponse[model.Beatmap](c.getReq(enum.EndpointGetBeatmap, opts))
}

// GetBeatmapAttributes returns difficulty attributes of beatmap with specific mode and mods combination.
func (c *OsuClient) GetBeatmapAttributes(opts opts.GetBeatmapAttributesOpts) (*model.BeatmapDifficultyAttributes, error) {
	attr, err := handleResponse[model.BeatmapDifficultyAttributesContainer](c.postReq(enum.EndpointGetBeatmapAttributes, opts))
	if err != nil {
		return nil, err
	}

	return &attr.Attributes, nil
}
