package opts

import (
	"fmt"
	"project/backend/api/enum"
	"reflect"
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

// GetUserKudosuOpts is a struct that contains fields for a successful user kudosu request.
// UserID is required. Limit and Offset are optional.
// For more information: https://osu.ppy.sh/docs/index.html#get-user-kudosu
type GetUserKudosuOpts struct {
	UserID int  `param:"path" index:"0"`
	Limit  *int `param:"query,limit"`
	Offset *int `param:"query,offset"`
}

func (g GetUserKudosuOpts) Valid() error {
	if g.UserID == 0 {
		return fmt.Errorf("both user ID")
	}

	return nil
}

// GetUserScoresOpts is a struct that contains fields for a successful user scores request.
// UserID and ScoreType are required. IncludeFails, Mode, Limit, and Offset are optional.
// For more information: https://osu.ppy.sh/docs/index.html#get-user-scores
type GetUserScoresOpts struct {
	UserID       int            `param:"path" index:"0"`
	ScoreType    enum.ScoreType `param:"path" index:"1"`
	IncludeFails bool           `param:"query,include_fails"`
	Mode         *enum.GameMode `param:"query,mode"`
	Limit        *int           `param:"query,limit"`
	Offset       *int           `param:"query,offset"`
}

func (g GetUserScoresOpts) Valid() error {
	if g.UserID == 0 || g.ScoreType == "" {
		return fmt.Errorf("both user ID and score type are empty")
	}

	return nil
}

// GetUserBeatmapsOpts is a struct that contains fields for a successful user beatmaps request.
// UserID and BeatmapType are required. Limit and Offset are optional.
// For more information: https://osu.ppy.sh/docs/index.html#get-user-beatmaps
type GetUserBeatmapsOpts struct {
	UserID      int                  `param:"path" index:"0"`
	BeatmapType enum.UserBeatmapType `param:"path" index:"1"`
	Limit       *int                 `param:"query,limit"`
	Offset      *int                 `param:"query,offset"`
}

func (g GetUserBeatmapsOpts) Valid() error {
	if g.UserID == 0 || g.BeatmapType == "" {
		return fmt.Errorf("user ID or beatmap type are empty")
	}

	return nil
}

// GetUserRecentActivityOpts is a struct that contains fields for a successful user recent activity request.
// UserID is required. Limit and Offset are optional.
// For more information: https://osu.ppy.sh/docs/index.html#get-user-recent-activity
type GetUserRecentActivityOpts struct {
	UserID int  `param:"path" index:"0"`
	Limit  *int `param:"query,limit"`
	Offset *int `param:"query,offset"`
}

func (g GetUserRecentActivityOpts) Valid() error {
	if g.UserID == 0 {
		return fmt.Errorf("user ID is empty")
	}

	return nil
}

// GetBeatmapLookupOpts is a struct that contains fields for a successful beatmap lookup request.
// All fields are optional but least one field has to be included.
// For more information: https://osu.ppy.sh/docs/index.html#lookup-beatmap
type GetBeatmapLookupOpts struct {
	BeatmapID *int    `param:"query,id"`
	Checksum  *string `param:"query,checksum"`
	Filename  *string `param:"query,filename"`
}

func (g GetBeatmapLookupOpts) Valid() error {
	if g.BeatmapID == nil && g.Checksum == nil && g.Filename == nil {
		return fmt.Errorf("at least one field has to be set")
	}

	return nil
}

// GetUserBeatmapScoreOpts is a struct that contains fields for a successful user beatmap score request.
// UserID and BeatmapID are required. Mode is optional.
// For more information: https://osu.ppy.sh/docs/index.html#get-a-user-beatmap-score
type GetUserBeatmapScoreOpts struct {
	BeatmapID int            `param:"path" index:"0"`
	UserID    int            `param:"path" index:"1"`
	Mode      *enum.GameMode `param:"query,mode"`
	// Mods are not implemented in the osu! API itself yet but the functionality is coming later afaik :clueless:
	// Mods      *enum.Mods     `param:"query,mods"`
}

func (g GetUserBeatmapScoreOpts) Valid() error {
	if g.UserID == 0 || g.BeatmapID == 0 {
		return fmt.Errorf("user ID or beatmap ID are empty")
	}

	return nil
}

// GetUserBeatmapScoresOpts is a struct that contains fields for a successful user beatmap scores request.
// UserID and BeatmapID are required. Mode is optional.
// For more information: https://osu.ppy.sh/docs/index.html#get-a-user-beatmap-scores
type GetUserBeatmapScoresOpts struct {
	BeatmapID int            `param:"path" index:"0"`
	UserID    int            `param:"path" index:"1"`
	Mode      *enum.GameMode `param:"query,mode"`
}

func (g GetUserBeatmapScoresOpts) Valid() error {
	if g.UserID == 0 && g.BeatmapID == 0 {
		return fmt.Errorf("user ID or beatmap ID are empty")
	}

	return nil
}

// GetBeatmapScoresOpts is a struct that contains fields for a successful beatmap scores request.
// BeatmapID is required. Mode is optional.
// For more information: https://osu.ppy.sh/docs/index.html#get-beatmap-scores
type GetBeatmapScoresOpts struct {
	BeatmapID int            `param:"path" index:"0"`
	Mode      *enum.GameMode `param:"query,mode"`
	// Type and Mods are not implemented in the osu! API itself yet but the functionality is coming later afaik :clueless:
	// Type      *enum.ScoreType `param:"query,type"`
	// Mods      *enum.Mods      `param:"query,mods"`
}

func (g GetBeatmapScoresOpts) Valid() error {
	if g.BeatmapID == 0 {
		return fmt.Errorf("beatmap ID is empty")
	}

	return nil
}

// GetBeatmapsOpts is a struct that contains fields for a successful beatmaps request.
// IDs is optional.
// For more information: https://osu.ppy.sh/docs/index.html#get-beatmaps
type GetBeatmapsOpts struct {
	IDs []int `param:"query,ids[]"`
}

// Valid validation is not needed for GetBeatmapsOpts since IDs are optional and the request will work without them just fine.
func (g GetBeatmapsOpts) Valid() error {
	return nil
}

// GetBeatmapOpts is a struct that contains fields for a successful beatmap request.
// BeatmapID is required.
// For more information: https://osu.ppy.sh/docs/index.html#get-beatmap
type GetBeatmapOpts struct {
	BeatmapID int `param:"path" index:"0"`
}

func (g GetBeatmapOpts) Valid() error {
	if g.BeatmapID == 0 {
		return fmt.Errorf("beatmap ID is empty")
	}

	return nil
}

// GetBeatmapAttributesOpts is a struct that contains fields for a successful beatmap attributes request.
// BeatmapID is required. Mods, and Mode are optional.
// For more information: https://osu.ppy.sh/docs/index.html#get-beatmap-attributes
type GetBeatmapAttributesOpts struct {
	BeatmapID int            `param:"path" index:"0"`
	Mods      *enum.Mods     `param:"body,mods"`
	Mode      *enum.GameMode `param:"body,ruleset"`
}

func (g GetBeatmapAttributesOpts) Valid() error {
	if g.BeatmapID == 0 {
		return fmt.Errorf("beatmap ID is empty")
	}

	return nil
}

// GetBeatmapsetDiscussionsOpts is a struct that contains fields for a successful beatmapset discussions request.
// All fields are optional. However, at least one has to be specified. WithDelete is currently unimplemented in the osu! API.
// For more information: https://osu.ppy.sh/docs/index.html#get-beatmapset-discussions
type GetBeatmapsetDiscussionsOpts struct {
	BeatmapID        *int                   `param:"query,beatmap_id"`
	BeatmapsetID     *int                   `param:"query,beatmapset_id"`
	BeatmapsetStatus *enum.BeatmapsetStatus `param:"query,beatmapset_status"`
	Limit            *int                   `param:"query,limit"`
	MessageTypes     []enum.MessageType     `param:"query,message_types[]"`
	OnlyUnresolved   *bool                  `param:"query,only_unresolved"`
	Page             *int                   `param:"query,page"`
	Sort             *enum.Sort             `param:"query,sort"`
	UserID           *int                   `param:"query,user"`
	WithDeleted      *string                `param:"query,with_deleted"`
}

func (g GetBeatmapsetDiscussionsOpts) Valid() error {
	if reflect.DeepEqual(g, GetBeatmapsetDiscussionsOpts{}) {
		return fmt.Errorf("none of the fields are set")
	}

	return nil
}

// GetBeatmapsetDiscussionPostsOpts is a struct that contains fields for a successful beatmapset discussion post request.
// All fields are optional. However, at least one has to be specified. WithDelete is currently unimplemented in the osu! API.
// For more information: https://osu.ppy.sh/docs/index.html#get-beatmapset-discussion-posts
type GetBeatmapsetDiscussionPostsOpts struct {
	BeatmapsetDiscussionID *int                         `param:"query,beatmapset_discussion_id"`
	Limit                  *int                         `param:"query,limit"`
	Types                  []enum.DiscussionMessageType `param:"query,types[]"`
	Page                   *int                         `param:"query,page"`
	Sort                   *enum.Sort                   `param:"query,sort"`
	UserID                 *int                         `param:"query,user"`
	WithDeleted            *string                      `param:"query,with_deleted"`
}

func (g GetBeatmapsetDiscussionPostsOpts) Valid() error {
	if reflect.DeepEqual(g, GetBeatmapsetDiscussionPostsOpts{}) {
		return fmt.Errorf("none of the fields are set")
	}

	return nil
}

// GetBeatmapsetDiscussionVotesOpts is a struct that contains fields for a successful beatmapset discussion votes request.
// All fields are optional, however, at least one has to be specified. WithDelete is currently unimplemented in the osu! API.
// For more information: https://osu.ppy.sh/docs/index.html#get-beatmapset-discussion-votes
type GetBeatmapsetDiscussionVotesOpts struct {
	BeatmapsetDiscussionID *int                  `param:"query,beatmapset_discussion_id"`
	Limit                  *int                  `param:"query,limit"`
	ReceiverUserID         *int                  `param:"query,receiver"`
	Score                  *enum.DiscussionScore `param:"query,score"`
	Page                   *int                  `param:"query,page"`
	Sort                   *enum.Sort            `param:"query,sort"`
	UserID                 *int                  `param:"query,user"`
	WithDeleted            *string               `param:"query,with_deleted"`
}

func (g GetBeatmapsetDiscussionVotesOpts) Valid() error {
	if reflect.DeepEqual(g, GetBeatmapsetDiscussionVotesOpts{}) {
		return fmt.Errorf("none of the fields are set")
	}

	return nil
}

// GetTopicAndPostsOpts is a struct that contains fields for a successful topic and posts request.
// TopicID is required. Limit, Sort, Start, and End are optional.
// If Sort is set, it is required to use Start and End. Start and End are ignored if CursorString is set.
//
// This struct is also returned by the endpoint itself, excluding the cursor string.
// CursorString is returned separately as a field.
// For more information: https://osu.ppy.sh/docs/index.html#get-topic-and-posts
type GetTopicAndPostsOpts struct {
	TopicID      int           `param:"path" index:"0" json:"topic"`
	Limit        *int          `param:"query,limit" json:"limit"`
	Sort         *enum.Sort    `param:"query,sort" json:"sort"`
	Start        *string       `param:"query,start" json:"start"`
	End          *string       `param:"query,end" json:"end"`
	CursorString *CursorString `param:"query,cursor_string" json:"cursor_string"`
}

func (g GetTopicAndPostsOpts) Valid() error {
	if g.TopicID == 0 {
		return fmt.Errorf("topic ID is empty")
	}

	if g.Sort != nil && g.CursorString == nil && (g.Start == nil || g.End == nil) {
		return fmt.Errorf("start and end should be set when using sort. alternatively, you can set a cursor")
	}

	return nil
}

// SearchOpts is a struct that contains fields for a successful search request.
// All fields are optional. However, at least one has to be specified.
// For more information: https://osu.ppy.sh/docs/index.html#search
type SearchOpts struct {
	Mode  *enum.SearchMode `param:"query,mode"`
	Query *string          `param:"query,query"`
	Page  *int             `param:"query,page"`
}

func (s SearchOpts) Valid() error {
	if reflect.DeepEqual(s, SearchOpts{}) {
		return fmt.Errorf("none of the fields are set")
	}

	return nil
}

// GetNewsListingOpts is a struct that contains fields for a successful news listing request.
// All fields are optional. However, at least one has to be specified.
// For more information: https://osu.ppy.sh/docs/index.html#get-news-listing
type GetNewsListingOpts struct {
	Limit  *int    `param:"query,limit"`
	Year   *int    `param:"query,year"`
	Cursor *string `param:"query,cursor"`
}

func (g GetNewsListingOpts) Valid() error {
	if reflect.DeepEqual(g, GetNewsListingOpts{}) {
		return fmt.Errorf("none of the fields are set")
	}

	return nil
}

// GetNewsPostOpts is a struct that contains fields for a successful news post request.
// At least one of the fields has to be specified. If PostID is set, Key should be set to enum.NewsPostKeyID.
// For more information: https://osu.ppy.sh/docs/index.html#get-news-post
type GetNewsPostOpts struct {
	PostID   *int              `param:"path" index:"0"`
	PostSlug *string           `param:"path" index:"0"`
	Key      *enum.NewsPostKey `param:"query,key"`
}

func (g GetNewsPostOpts) Valid() error {
	if reflect.DeepEqual(g, GetNewsPostOpts{}) {
		return fmt.Errorf("none of the fields are set")
	}

	return nil
}

// GetRankingsOpts is a struct that contains fields for a successful ranking request.
// Mode and RankingType are required. CountryCode and Variant are optional. SpotlightID has to be set if RankingType is Spotlight.
// For more information: https://osu.ppy.sh/docs/index.html#get-ranking
type GetRankingsOpts struct {
	Mode        enum.GameMode    `param:"path" index:"0"`
	RankingType enum.RankingType `param:"path" index:"1"`
	CountryCode *string          `param:"query,country"`
	SpotlightID *int             `param:"query,spotlight"`
	Variant     *string          `param:"query,variant"`
	Cursor      *Cursor          `param:"query,cursor"`
}

func (g GetRankingsOpts) Valid() error {
	if g.Mode == "" || g.RankingType == "" {
		return fmt.Errorf("mode or ranking type are empty")
	}

	if g.RankingType == enum.RankingTypeSpotlight && g.SpotlightID == nil {
		return fmt.Errorf("spotlight ID has to be set if ranking type is spotlight")
	}

	return nil
}

// GetMultiplayerMatchesOpts is a struct that contains fields for a successful multiplayer matches request.
// All fields are optional.
// Information is available only in a form of source code since it is undocumented.
// https://github.com/ppy/osu-web/blob/master/app/Http/Controllers/MatchesController.php#L20
type GetMultiplayerMatchesOpts struct {
	Limit *int       `param:"query,limit"`
	Sort  *enum.Sort `param:"query,sort"`

	CursorString *Cursor `param:"query,cursor"`
}

func (g GetMultiplayerMatchesOpts) Valid() error {
	return nil
}

// GetMultiplayerMatchOpts is a struct that contains fields for a successful multiplayer match request.
// MatchID is required. AfterEventID, BeforeEventID, and Limit are optional.
// AfterEventID and BeforeEventID are used to get the match state containing only events within the given ID constraints.
// Information is available only in a form of source code since it is undocumented.
// https://github.com/ppy/osu-web/blob/master/app/Http/Controllers/MatchesController.php#L39
type GetMultiplayerMatchOpts struct {
	MatchID       int  `param:"path" index:"0"`
	AfterEventID  *int `param:"query,after"`
	BeforeEventID *int `param:"query,before"`
	Limit         *int `param:"query,limit"`
}

func (g GetMultiplayerMatchOpts) Valid() error {
	if g.MatchID == 0 {
		return fmt.Errorf("match ID is not specified")
	}

	return nil
}
