package client

import (
	"log"
	"os"
	"project/backend/api/client/opts"
	"project/backend/api/enum"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

var c *OsuClient

// TestMain runs before running any of the actual tests to initialize the osu! client.
// The tests are mainly testing for errors that might occur while handling the requests. E.g. unmarshalling JSONs etc.
// The goal is not to test the actual osu! API.
func TestMain(m *testing.M) {
	c, _ = NewClient(getOsuAuthCredentials())
	os.Exit(m.Run())
}

// getOsuAuthCredentials returns client ID and client secret
func getOsuAuthCredentials() (int, string) {
	clientSecret := os.Getenv("OSU_CLIENT_SECRET")
	clientID, err := strconv.Atoi(os.Getenv("OSU_CLIENT_ID"))
	if err != nil {
		log.Fatalln(err)
	}

	if clientID == 0 || clientSecret == "" {
		log.Fatalln("client credentials not set in environment")
	}

	return clientID, clientSecret
}

func Test_RequestsWithCursor(t *testing.T) {
	rankings, err := c.GetRankings(opts.GetRankingsOpts{
		Mode:        enum.GameModeStandard,
		RankingType: enum.RankingTypePerformance,
	})

	require.Nil(t, err)
	require.NotEmpty(t, rankings)
	require.NotEmpty(t, rankings.Cursor)

	nextRankings, err := c.GetRankings(opts.GetRankingsOpts{
		Mode:        enum.GameModeStandard,
		RankingType: enum.RankingTypePerformance,
		Cursor:      &rankings.Cursor,
	})

	require.NotEqual(t, rankings, nextRankings)
}

func Test_RequestsWithCursorString(t *testing.T) {
	topic, err := c.GetTopicAndPosts(opts.GetTopicAndPostsOpts{
		TopicID: 1614314,
	})

	require.Nil(t, err)
	require.NotEmpty(t, topic)
	require.NotNil(t, topic.CursorString)

	sort := enum.SortDescID
	nextTopic, err := c.GetTopicAndPosts(opts.GetTopicAndPostsOpts{
		TopicID:      1614314,
		CursorString: topic.CursorString,
		Sort:         &sort,
	})

	require.NotEqual(t, topic, nextTopic)
}

func TestOsuClient_GetBeatmaps(t *testing.T) {
	beatmaps, err := c.GetBeatmaps(opts.GetBeatmapsOpts{IDs: []int{3400798, 3736505}})

	require.Nil(t, err)
	require.NotEmpty(t, beatmaps)
}

func TestOsuClient_GetBeatmap(t *testing.T) {
	beatmap, err := c.GetBeatmap(opts.GetBeatmapOpts{BeatmapID: 3400798})

	require.Nil(t, err)
	require.NotEmpty(t, beatmap)
}

func TestOsuClient_GetBeatmapAttributes(t *testing.T) {
	beatmap, err := c.GetBeatmapAttributes(opts.GetBeatmapAttributesOpts{BeatmapID: 3795706})

	require.Nil(t, err)
	require.NotEmpty(t, beatmap)
}

func TestOsuClient_LookupBeatmap_IDOnly(t *testing.T) {
	beatmapID := 3400798
	beatmap, err := c.LookupBeatmap(opts.GetBeatmapLookupOpts{BeatmapID: &beatmapID})

	require.Nil(t, err)
	require.NotEmpty(t, beatmap)
}

func TestOsuClient_LookupBeatmap_FilenameOnly(t *testing.T) {
	// Absolutely stupid. Has to be the exact filename or else it won't work.
	filename := "Kano - Stella-rium (Moeka) [Starlight].osu"
	beatmap, err := c.LookupBeatmap(opts.GetBeatmapLookupOpts{Filename: &filename})

	require.Nil(t, err)
	require.NotEmpty(t, beatmap)
}

func TestOsuClient_LookupBeatmap_ChecksumOnly(t *testing.T) {
	// md5sum *.osu
	checksum := "0ff49cc30e2de31f6a7c263963162dbf"
	opt := opts.GetBeatmapLookupOpts{Checksum: &checksum}
	beatmap, err := c.LookupBeatmap(opt)

	require.Nil(t, err)
	require.NotEmpty(t, beatmap)
}

func TestOsuClient_LookupBeatmap_All(t *testing.T) {
	beatmapID := 785097
	checksum := "0ff49cc30e2de31f6a7c263963162dbf"
	filename := "Kano - Stella-rium (Moeka) [Starlight].osu"
	beatmap, err := c.LookupBeatmap(opts.GetBeatmapLookupOpts{
		BeatmapID: &beatmapID,
		Checksum:  &checksum,
		Filename:  &filename,
	})

	require.Nil(t, err)
	require.NotEmpty(t, beatmap)
}

func TestOsuClient_LookupBeatmap_Fail(t *testing.T) {
	badBeatmapID := 0
	beatmap, err := c.LookupBeatmap(opts.GetBeatmapLookupOpts{BeatmapID: &badBeatmapID})

	require.NotNil(t, err)
	require.Nil(t, beatmap)
}

func TestOsuClient_GetUserBeatmapScore(t *testing.T) {
	userScore, err := c.GetUserBeatmapScore(opts.GetUserBeatmapScoreOpts{
		BeatmapID: 1029562,
		UserID:    7093124,
	})

	require.Nil(t, err)
	require.NotEmpty(t, userScore)
}

func TestOsuClient_GetUserBeatmapScores(t *testing.T) {
	scores, err := c.GetUserBeatmapScores(opts.GetUserBeatmapScoresOpts{
		BeatmapID: 1029562,
		UserID:    7093124,
	})

	require.Nil(t, err)
	require.NotEmpty(t, scores)
}

func TestOsuClient_GetUserBeatmapScores_WithMode(t *testing.T) {
	taiko, std := enum.GameModeTaiko, enum.GameModeStandard
	opt := opts.GetUserBeatmapScoresOpts{
		BeatmapID: 1669074,
		UserID:    7093124,
		Mode:      &taiko,
	}

	taikoScores, err := c.GetUserBeatmapScores(opt)
	require.Nil(t, err)

	opt.Mode = &std

	stdScores, err := c.GetUserBeatmapScores(opt)
	require.Nil(t, err)

	require.NotEqual(t, taikoScores, stdScores)
}

func TestOsuClient_GetBeatmapScores(t *testing.T) {
	scores, err := c.GetBeatmapScores(opts.GetBeatmapScoresOpts{
		BeatmapID: 1029562,
	})

	require.Nil(t, err)
	require.NotEmpty(t, scores)
}

func TestOsuClient_GetBeatmapScores_WithMode(t *testing.T) {
	taiko, std := enum.GameModeTaiko, enum.GameModeStandard
	opt := opts.GetBeatmapScoresOpts{
		BeatmapID: 1029562,
		Mode:      &taiko,
	}
	taikoScores, err := c.GetBeatmapScores(opt)
	require.Nil(t, err)

	opt.Mode = &std

	stdScores, err := c.GetBeatmapScores(opt)
	require.Nil(t, err)

	require.NotEqual(t, taikoScores, stdScores)
}

func TestOsuClient_GetBeatmapsetDiscussions(t *testing.T) {
	setID, sort := 1665789, enum.SortAscID
	discussions, err := c.GetBeatmapsetDiscussions(opts.GetBeatmapsetDiscussionsOpts{
		BeatmapsetID: &setID,
		Sort:         &sort,
	})

	require.Nil(t, err)
	require.NotEmpty(t, discussions)
}

func TestOsuClient_GetBeatmapsetDiscussionPosts(t *testing.T) {
	setDiscussionID, limit, sort := 2829296, 1, enum.SortAscID
	discussionPosts, err := c.GetBeatmapsetDiscussionPosts(opts.GetBeatmapsetDiscussionPostsOpts{
		BeatmapsetDiscussionID: &setDiscussionID,
		Limit:                  &limit,
		Sort:                   &sort,
	})

	require.Nil(t, err)
	require.NotEmpty(t, discussionPosts)
}

func TestOsuClient_GetBeatmapsetDiscussionVotes(t *testing.T) {
	setDiscussionID, limit, score := 2829296, 1, enum.DiscussionScoreUpvote
	discussionVotes, err := c.GetBeatmapsetDiscussionVotes(opts.GetBeatmapsetDiscussionVotesOpts{
		BeatmapsetDiscussionID: &setDiscussionID,
		Limit:                  &limit,
		Score:                  &score,
	})

	require.Nil(t, err)
	require.NotEmpty(t, discussionVotes)
}

func TestOsuClient_GetTopicAndPosts(t *testing.T) {
	topic, err := c.GetTopicAndPosts(opts.GetTopicAndPostsOpts{
		TopicID: 1614314,
	})

	require.Nil(t, err)
	require.NotEmpty(t, topic)
}

func TestOsuClient_GetMultiplayerMatches(t *testing.T) {
	limit, sort := 4, enum.SortAscID
	matches, err := c.GetMultiplayerMatches(opts.GetMultiplayerMatchesOpts{
		Limit: &limit,
		Sort:  &sort,
	})

	require.Nil(t, err)
	require.NotEmpty(t, matches)
}

func TestOsuClient_GetMultiplayerMatch(t *testing.T) {
	after := 484460136
	match, err := c.GetMultiplayerMatch(opts.GetMultiplayerMatchOpts{MatchID: 16159321, AfterEventID: &after})

	require.Nil(t, err)
	require.NotEmpty(t, match)
}

func TestOsuClient_GetNewsListing(t *testing.T) {
	limit, year := 3, 2021
	news, err := c.GetNewsListing(opts.GetNewsListingOpts{
		Limit: &limit,
		Year:  &year,
	})

	require.Nil(t, err)
	require.NotEmpty(t, news)
}

func TestOsuClient_GetNewsPost(t *testing.T) {
	postID := 1065
	post, err := c.GetNewsPost(opts.GetNewsPostOpts{
		PostID: &postID,
	})

	require.Nil(t, err)
	require.NotEmpty(t, post)
}

func TestOsuClient_GetRankings(t *testing.T) {
	rankings, err := c.GetRankings(opts.GetRankingsOpts{
		Mode:        enum.GameModeStandard,
		RankingType: enum.RankingTypePerformance,
	})

	require.Nil(t, err)
	require.NotEmpty(t, rankings)
}

func TestOsuClient_GetSpotlights(t *testing.T) {
	spotlights, err := c.GetSpotlights()

	require.Nil(t, err)
	require.NotEmpty(t, spotlights)
}

func TestOsuClient_Search(t *testing.T) {
	searchMode, query := enum.SearchModeAll, "slay"
	res, err := c.Search(opts.SearchOpts{
		Mode:  &searchMode,
		Query: &query,
		Page:  nil,
	})

	require.Nil(t, err)
	require.NotEmpty(t, res)
}

func TestOsuClient_GetUserKudosu(t *testing.T) {
	kudosu, err := c.GetUserKudosu(opts.GetUserKudosuOpts{UserID: 6063342})

	require.Nil(t, err)
	require.NotEmpty(t, kudosu)
}

func TestOsuClient_GetUserScores(t *testing.T) {
	limit := 5
	scores, err := c.GetUserScores(opts.GetUserScoresOpts{
		UserID:    7093124,
		ScoreType: enum.ScoreTypeBest,
		Limit:     &limit,
	})

	require.Nil(t, err)
	require.NotEmpty(t, scores)
}

func TestOsuClient_GetUserBeatmaps(t *testing.T) {
	beatmaps, err := c.GetUserBeatmaps(opts.GetUserBeatmapsOpts{
		UserID:      7093124,
		BeatmapType: enum.UserBeatmapTypeGraveyard,
	})

	require.Nil(t, err)
	require.NotEmpty(t, beatmaps)
}

// TestOsuClient_GetUserRecentActivity can fail if the tested user hasn't been active or hasn't done anything notable
// in a while.
func TestOsuClient_GetUserRecentActivity(t *testing.T) {
	recentActivity, err := c.GetUserRecentActivity(opts.GetUserRecentActivityOpts{UserID: 7093124})

	require.Nil(t, err)
	require.NotEmpty(t, recentActivity)
}

func TestOsuClient_GetUser(t *testing.T) {
	key := enum.UserKeyID
	user, err := c.GetUser(opts.GetUserOpts{
		UserID: 7093124,
		Key:    &key,
	})

	require.Nil(t, err)
	require.NotEmpty(t, user)
}
