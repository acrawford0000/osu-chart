package client

import (
	"project/backend/api/client/opts"
	"project/backend/api/enum"
	"project/backend/api/model"
)

// GetRankings returns the current ranking for the specified type and game mode from the osu! API.
func (c *OsuClient) GetRankings(opts opts.GetRankingsOpts) (*model.Rankings, error) {
	return handleResponse[model.Rankings](c.getReq(enum.EndpointGetRankings, opts))
}

// GetSpotlights returns a slice of spotlights from the osu! API.
// The method does not require any opts due to it not taking any parameters.
func (c *OsuClient) GetSpotlights() ([]model.Spotlight, error) {
	container, err := handleResponse[model.Spotlights](c.getReq(enum.EndpointGetSpotlights, nil))
	if err != nil {
		return nil, err
	}

	return container.Spotlights, nil
}
