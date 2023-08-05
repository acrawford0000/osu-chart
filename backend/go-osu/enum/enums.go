package enum

import (
	"encoding/json"
	"strings"
)

// Mods is a bitwise enum of mods used in a score.
//
// Mods may appear complicated to use for a beginner programmer. Fear not!
// This is how hard they can get for creation of a mod combination:
//
//	myModCombination := enum.ModHardRock | enum.ModDoubleTime | enum.ModHidden | enum.ModSpunOut
//
// As for checking that an existing mod combination is enabled:
//
//	if modCombination&enum.ModHardRock != 0 {
//	    // HardRock is enabled
//	}
//
// To learn more about bitwise operators, have a look at it on wikipedia:
// https://en.wikipedia.org/wiki/Bitwise_operation#Bitwise_operators
//
// Credit for the Mods enum goes to thehowl on github: https://github.com/thehowl/go-osuapi/blob/master/mods.go
type Mods int

const ModNomod Mods = 0
const (
	ModNoFail Mods = 1 << iota
	ModEasy
	ModTouchDevice
	ModHidden
	ModHardRock
	ModSuddenDeath
	ModDoubleTime
	ModRelax
	ModHalfTime
	ModNightcore
	ModFlashlight
	ModAutoplay
	ModSpunOut
	ModAutopilot
	ModPerfect
	ModKey4
	ModKey5
	ModKey6
	ModKey7
	ModKey8
	ModFadeIn
	ModRandom
	ModCinema
	ModTarget
	ModKey9
	ModKeyCoop
	ModKey1
	ModKey3
	ModKey2
	ModScoreV2
	ModFreeModAllowed    = ModNoFail | ModEasy | ModHidden | ModHardRock | ModSuddenDeath | ModFlashlight | ModFadeIn | ModRelax | ModAutopilot | ModSpunOut | ModKeyMod
	ModKeyMod            = ModKey1 | ModKey2 | ModKey3 | ModKey4 | ModKey5 | ModKey6 | ModKey7 | ModKey8 | ModKey9 | ModKeyCoop
	ModScoreIncreaseMods = ModHidden | ModHardRock | ModDoubleTime | ModNightcore | ModFlashlight | ModFadeIn
)

var modsString = [...]string{
	"NF",
	"EZ",
	"NV",
	"HD",
	"HR",
	"SD",
	"DT",
	"RX",
	"HT",
	"NC",
	"FL",
	"AU", // Auto.
	"SO",
	"AP", // Autopilot.
	"PF",
	"K4",
	"K5",
	"K6",
	"K7",
	"K8",
	"FI",
	"RN",
	"CN",
	"TR",
	"K9",
	"KC",
	"K1",
	"K3",
	"K2",
	"V2",
}

// ParseMods parse a string with mods in the format "HDHRDT"
func ParseMods(mods string) (m Mods) {
	if len(mods) == 0 || len(mods)%2 != 0 {
		return
	}
	modsSl := make([]string, len(mods)/2)
	for n, modPart := range mods {
		modsSl[n/2] += string(modPart)
	}
	for _, mod := range modsSl {
		for index, availableMod := range modsString {
			if availableMod == mod {
				m |= 1 << uint(index)
				break
			}
		}
	}
	return
}

func (m Mods) String() (s string) {
	if m == 0 {
		s = "NM"
	}
	for i := 0; i < len(modsString); i++ {
		if activated := 1&m == 1; activated {
			s += modsString[i]
		}
		m >>= 1
	}
	return
}

func (m Mods) UnmarshalJSON(bytes []byte) (err error) {
	var mods []string
	err = json.Unmarshal(bytes, &mods)
	if err != nil {
		return
	}

	m = parseModArray(mods)

	return
}

func parseModArray(mods []string) Mods {
	return ParseMods(strings.Join(mods, ""))
}

// Sort is a type for specifying the sort method for various requests.
type Sort string

const (
	SortDescID Sort = "id_desc"
	SortAscID  Sort = "id_asc"
)

func (s Sort) String() string {
	return string(s)
}

// GrantType is used for auth.
type GrantType string

const GrantTypeClientCredentials GrantType = "client_credentials"

func (g GrantType) String() string {
	return string(g)
}

// Scope is used for auth.
type Scope string

const ScopePublic Scope = "public"

func (s Scope) String() string {
	return string(s)
}

// BeatmapsetStatus is a type for specifying a beatmapset status.
type BeatmapsetStatus string

const (
	BeatmapsetStatusAll            BeatmapsetStatus = "all"
	BeatmapsetStatusRanked         BeatmapsetStatus = "ranked"
	BeatmapsetStatusQualified      BeatmapsetStatus = "qualified"
	BeatmapsetStatusDisqualified   BeatmapsetStatus = "disqualified"
	BeatmapsetStatusNeverQualified BeatmapsetStatus = "never_qualified"
)

func (b BeatmapsetStatus) String() string {
	return string(b)
}

// GameMode is used for specifying the osu! game mode.
type GameMode string

const (
	GameModeStandard GameMode = "osu"
	GameModeTaiko    GameMode = "taiko"
	GameModeCtb      GameMode = "fruits"
	GameModeMania    GameMode = "mania"
)

func (g GameMode) String() string {
	return string(g)
}

// DiscussionMessageType is a type for specifying the type of messages in beatmapset discussion.
type DiscussionMessageType string

const (
	DiscussionMessageTypeFirst  DiscussionMessageType = "first"
	DiscussionMessageTypeReply  DiscussionMessageType = "reply"
	DiscussionMessageTypeSystem DiscussionMessageType = "system"
)

func (d DiscussionMessageType) String() string {
	return string(d)
}

// DiscussionScore is used for specifying the score of a message inside a beatmapset discussion.
type DiscussionScore int

const (
	DiscussionScoreDownvote DiscussionScore = iota - 1
	_                                       // skip 0 value
	DiscussionScoreUpvote
)

var discussionScoreStrings = map[DiscussionScore]string{
	DiscussionScoreUpvote:   "upvote",
	DiscussionScoreDownvote: "downvote",
}

func (d DiscussionScore) Int() int {
	return int(d)
}

func (d DiscussionScore) String() string {
	return discussionScoreStrings[d]
}

// BaseEndpoint is a type for osu API v2 base URLs.
type BaseEndpoint string

const (
	BaseEndpointOsu   BaseEndpoint = "https://osu.ppy.sh/api/v2"
	BaseEndpointOAuth BaseEndpoint = "https://osu.ppy.sh/oauth/token"
)

func (b BaseEndpoint) String() string {
	return string(b)
}

// Endpoint is a type for osu API v2 endpoints. Only endpoints with public scope are supported.
type Endpoint string

const (
	EndpointSearch                       Endpoint = "/search"
	EndpointBeatmapLookup                Endpoint = "/beatmaps/lookup"
	EndpointGetBeatmaps                  Endpoint = "/beatmaps"
	EndpointGetBeatmap                   Endpoint = "/beatmaps/%d"
	EndpointGetBeatmapScores             Endpoint = "/beatmaps/%d/scores"
	EndpointGetUserBeatmapScore          Endpoint = "/beatmaps/%d/scores/users/%d"
	EndpointGetUserBeatmapScores         Endpoint = "/beatmaps/%d/scores/users/%d/all"
	EndpointGetBeatmapAttributes         Endpoint = "/beatmaps/%d/attributes"
	EndpointGetBeatmapsetDiscussions     Endpoint = "/beatmapsets/discussions"
	EndpointGetBeatmapsetDiscussionPosts Endpoint = "/beatmapsets/discussions/posts"
	EndpointGetBeatmapsetDiscussionVotes Endpoint = "/beatmapsets/discussions/votes"
	EndpointGetTopicAndPosts             Endpoint = "/forums/topics/%d"
	EndpointGetNews                      Endpoint = "/news"
	EndpointGetNewsPost                  Endpoint = "/news/%v"
	EndpointGetRankings                  Endpoint = "/rankings/%s/%s"
	EndpointGetSpotlights                Endpoint = "/spotlights"
	EndpointGetUserKudosu                Endpoint = "/users/%d/kudosu"
	EndpointGetUserScores                Endpoint = "/users/%d/scores/%s"
	EndpointGetUserBeatmaps              Endpoint = "/users/%d/beatmapsets/%s"
	EndpointGetUserRecentActivity        Endpoint = "/users/%d/recent_activity"
	EndpointGetUser                      Endpoint = "/users/%v/%v"

	// Undocumented endpoints
	EndpointGetMatches Endpoint = "/matches"
	EndpointGetMatch   Endpoint = "/matches/%d"
	// EndpointGetSeasonalBackgrounds Endpoint = "/seasonal-backgrounds"
	// EndpointGetRoomLeaderboard     Endpoint = "/rooms/%s/leaderboard"
	// EndpointGetRoom                Endpoint = "/rooms/%s"
	// EndpointGetRooms               Endpoint = "/rooms/%s"

	// Unimplemented
	EndpointGetBeatmapsetEvents Endpoint = "/beatmapsets/events"
)

func (e Endpoint) String() string {
	return string(e)
}

// Grade is a type for specifying the grade of a Score.
type Grade string

const (
	GradeSSH Grade = "XH"
	GradeSS  Grade = "X"
	GradeSH  Grade = "SH"
	GradeS   Grade = "S"
	GradeA   Grade = "A"
	GradeB   Grade = "B"
	GradeC   Grade = "C"
	GradeD   Grade = "D"
	GradeF   Grade = "F"
)

func (g Grade) String() string {
	return string(g)
}

// MessageType is a type for specifying a message type in a beatmapset discussion.
type MessageType string

const (
	MessageTypeSuggestion MessageType = "suggestion"
	MessageTypeProblem    MessageType = "problem"
	MessageTypeMapperNote MessageType = "mapper_note"
	MessageTypePraise     MessageType = "praise"
	MessageTypeHype       MessageType = "hype"
	MessageTypeReview     MessageType = "review"
)

func (m MessageType) String() string {
	return string(m)
}

// RankedStatus is a type for specifying a ranked status of a beatmap/beatmapset.
type RankedStatus int

const (
	RankedStatusGraveyard RankedStatus = iota - 2
	RankedStatusWIP
	RankedStatusPending
	RankedStatusRanked
	RankedStatusApproved
	RankedStatusQualified
	RankedStatusLoved
)

var rankStatusStrings = map[RankedStatus]string{
	RankedStatusGraveyard: "graveyard",
	RankedStatusWIP:       "wip",
	RankedStatusPending:   "pending",
	RankedStatusRanked:    "ranked",
	RankedStatusApproved:  "approved",
	RankedStatusQualified: "qualified",
	RankedStatusLoved:     "loved",
}

var rankStatusFromString = map[string]RankedStatus{
	"graveyard": RankedStatusGraveyard,
	"wip":       RankedStatusWIP,
	"pending":   RankedStatusPending,
	"ranked":    RankedStatusRanked,
	"approved":  RankedStatusApproved,
	"qualified": RankedStatusQualified,
	"loved":     RankedStatusLoved,
}

func (rs RankedStatus) Int() int {
	return int(rs)
}

func (rs RankedStatus) String() string {
	return rankStatusStrings[rs]
}

func (rs *RankedStatus) UnmarshalJSON(data []byte) error {
	var temp interface{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	switch x := temp.(type) {
	case string:
		fromString := rankStatusFromString[x]
		*rs = fromString
		return nil
	case int:
		var rankStatus RankedStatus
		err := json.Unmarshal(data, &rankStatus)
		if err != nil {
			return err
		}
		*rs = rankStatus
	}

	return nil
}

// RankingType is a type for specifying the ranking type for leaderboards.
type RankingType string

const (
	RankingTypeSpotlight   RankingType = "charts"
	RankingTypeCountry     RankingType = "country"
	RankingTypePerformance RankingType = "performance"
	RankingTypeScore       RankingType = "score"
)

func (r RankingType) String() string {
	return string(r)
}

// ScoreType is a type for specifying a score type. For example, Best for top plays.
type ScoreType string

const (
	ScoreTypeRecent ScoreType = "recent"
	ScoreTypeBest   ScoreType = "best"
	ScoreTypeFirsts ScoreType = "firsts"
)

func (s ScoreType) String() string {
	return string(s)
}

// SearchMode is a type for specifying a search scope.
type SearchMode int

const (
	SearchModeAll SearchMode = iota
	SearchModeUser
	SearchModeWikiPage
)

var searchModeStrings = map[SearchMode]string{
	SearchModeAll:      "all",
	SearchModeUser:     "user",
	SearchModeWikiPage: "wiki_page",
}

func (s SearchMode) Int() int {
	return int(s)
}

func (s SearchMode) String() string {
	return searchModeStrings[s]
}

// UserBeatmapType is a type for specifying user beatmap types.
type UserBeatmapType string

const (
	UserBeatmapTypeFavorite  UserBeatmapType = "favourite"
	UserBeatmapTypeGraveyard UserBeatmapType = "graveyard"
	UserBeatmapTypeLoved     UserBeatmapType = "loved"
	UserBeatmapTypePending   UserBeatmapType = "pending"
	UserBeatmapTypeRanked    UserBeatmapType = "ranked"
)

func (u UserBeatmapType) String() string {
	return string(u)
}

type UserMostPlayedBeatmapType string

const UserBeatmapTypeMostPlayed UserMostPlayedBeatmapType = "most_played"

func (u UserMostPlayedBeatmapType) String() string {
	return string(u)
}

// KudosuAction is a type for specifying kudosu action.
type KudosuAction string

func (k KudosuAction) String() string {
	return string(k)
}

// ForumTopicType is a type for specifying the type of forum post.
type ForumTopicType string

const (
	ForumTopicTypeNormal       ForumTopicType = "normal"
	ForumTopicTypeSticky       ForumTopicType = "sticky"
	ForumTopicTypeAnnouncement ForumTopicType = "announcement"
)

func (f ForumTopicType) String() string {
	return string(f)
}

// MultiplayerScoresSort is a type for specifying the multiplayer score sort method.
type MultiplayerScoresSort string

const (
	MultiplayerScoresSortAsc  MultiplayerScoresSort = "score_asc"
	MultiplayerScoresSortDesc MultiplayerScoresSort = "score_desc"
)

func (m MultiplayerScoresSort) String() string {
	return string(m)
}

// MultiplayerMatchEventType is a type for specifying the multiplayer match event type.
type MultiplayerMatchEventType string

const (
	MultiplayerMatchEventTypeCreated     MultiplayerMatchEventType = "match-created"
	MultiplayerMatchEventTypeDisbanded   MultiplayerMatchEventType = "match-disbanded"
	MultiplayerMatchEventTypeGame        MultiplayerMatchEventType = "other"
	MultiplayerMatchEventTypeHostChanged MultiplayerMatchEventType = "host-changed"
	MultiplayerMatchEventTypeJoined      MultiplayerMatchEventType = "player-joined"
	MultiplayerMatchEventTypeKicked      MultiplayerMatchEventType = "player-kicked"
	MultiplayerMatchEventTypeLeft        MultiplayerMatchEventType = "player-left"
)

func (m MultiplayerMatchEventType) String() string {
	return string(m)
}

// MultiplayerScoringType is a type for specifying a multiplayer scoring type.
type MultiplayerScoringType string

const (
	MultiplayerScoringTypeScore    MultiplayerScoringType = "score"
	MultiplayerScoringTypeAccuracy MultiplayerScoringType = "accuracy"
	MultiplayerScoringTypeCombo    MultiplayerScoringType = "combo"
	MultiplayerScoringTypeScoreV2  MultiplayerScoringType = "scorev2"
)

func (m MultiplayerScoringType) String() string {
	return string(m)
}

// MultiplayerTeamType is a type for specifying a multiplayer team type.
type MultiplayerTeamType string

const (
	HeadToHead MultiplayerTeamType = "head-to-head"
	TagCoop    MultiplayerTeamType = "tag-coop"
	TeamVS     MultiplayerTeamType = "team-vs"
	TagTeamVS  MultiplayerTeamType = "tag-team-vs"
)

func (m MultiplayerTeamType) String() string {
	return string(m)
}

// MultiplayerTeam is a type for specifying the multiplayer team.
type MultiplayerTeam string

const (
	None MultiplayerTeam = "none"
	Blue MultiplayerTeam = "blue"
	Red  MultiplayerTeam = "red"
)

func (m MultiplayerTeam) String() string {
	return string(m)
}

// NewsPostKey is a type for specifying which lookup method is going to be used.
type NewsPostKey string

const (
	NewsPostKeyID   NewsPostKey = "id"
	NewsPostKeySlug NewsPostKey = "slug"
)

func (n NewsPostKey) String() string {
	return string(n)
}

// UserKey is a type for specifying which lookup method is going to be used.
type UserKey string

const (
	UserKeyID       UserKey = "id"
	UserKeyUsername UserKey = "username"
)

func (u UserKey) String() string {
	return string(u)
}
