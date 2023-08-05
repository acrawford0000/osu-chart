package client

import (
	"project/backend/go-osu/client/opts"
	"project/backend/go-osu/enum"
	"project/backend/go-osu/model"
)

// GetMultiplayerMatches returns a list of currently active multiplayer matches.
func (c *OsuClient) GetMultiplayerMatches(opts opts.GetMultiplayerMatchesOpts) (*model.MultiplayerMatches, error) {
	return handleResponse[model.MultiplayerMatches](c.getReq(enum.EndpointGetMatches, opts))
}

// GetMultiplayerMatch returns a multiplayer match with its history and players.
func (c *OsuClient) GetMultiplayerMatch(opts opts.GetMultiplayerMatchOpts) (*model.MultiplayerMatchDetails, error) {
	return handleResponse[model.MultiplayerMatchDetails](c.getReq(enum.EndpointGetMatch, opts))
}
