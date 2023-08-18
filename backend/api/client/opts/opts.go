package opts

import (
	"fmt"
	"project/backend/api/enum"
)

// Cursor is used in some osu! requests and responses to get the next set of results.
// The values of the cursor should be provided to next request of the same endpoint to get the next set of results.
// If there are no more results available, a cursor with a value of null is returned: "cursor": null.
// Note that sort option should also be specified for it to work.
type Cursor map[string]interface{}

// CursorString is a string value included in some API responses containing the parameter to get the next set of results.
// Its value will be null (or not defined) if there are no more results available.
// Note that all parameters used in previous request also need to be passed.
type CursorString string

type Opts interface {
	Valid() error
}

// GetUserOpts is a struct that contains fields for a successful user request.
// UserID or Username have to be specified. Mode is optional.
// Key is used to specify the lookup method. If empty, the API will first perform an ID lookup followed by username
// lookup if not found.
// For more information: https://osu.ppy.sh/docs/index.html#get-user
type GetUserOpts struct {
	UserID   int            `param:"path" index:"0"`
	Username string         `param:"path" index:"0"`
	Mode     *enum.GameMode `param:"path" index:"1"`
	Key      *enum.UserKey  `param:"query,key"`
}

func (g GetUserOpts) Valid() error {
	if g.UserID == 0 && g.Username == "" {
		return fmt.Errorf("both user ID and username are empty")
	}

	return nil
}
