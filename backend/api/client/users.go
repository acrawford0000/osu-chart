package client

import (
	"project/backend/api/client/opts"
	"project/backend/api/enum"
	"project/backend/api/model"
)

// GetUser returns the details of specified user.
func (c *OsuClient) GetUser(opts opts.GetUserOpts) (*model.User, error) {
	return handleResponse[model.User](c.getReq(enum.EndpointGetUser, opts))
}

// GetMyUserData filters only the data that will be used and returned to the frontend
func GetMyUserData(user *model.User) *model.MyUserStruct {
	return &model.MyUserStruct{
		ID:          user.ID,
		Username:    user.Username,
		AvatarURL:   user.AvatarURL,
		CountryCode: user.CountryCode,
		Country:     user.Country,
		CoverURL:    user.CoverURL,
		Playstyle:   user.Playstyle,
	}
}
