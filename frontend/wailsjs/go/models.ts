export namespace model {
	
	export class Kudosu {
	    available: number;
	    total: number;
	
	    static createFrom(source: any = {}) {
	        return new Kudosu(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.available = source["available"];
	        this.total = source["total"];
	    }
	}
	export class UserCompact {
	    id: number;
	    username: string;
	    avatar_url: string;
	    country_code: string;
	    default_group: string;
	    is_active: boolean;
	    is_bot: boolean;
	    is_deleted: boolean;
	    is_online: boolean;
	    is_supporter: boolean;
	    // Go type: DateTime
	    last_visit?: any;
	    pm_friends_only: boolean;
	    profile_colour?: string;
	    account_history: UserAccountHistory[];
	    // Go type: TournamentBanner
	    active_tournament_banner?: any;
	    badges: UserBadge[];
	    groups: UserGroups[];
	    beatmap_playcounts_count: number;
	    favorite_beatmapset_count: number;
	    graveyard_beatmapset_count: number;
	    loved_beatmapset_count: number;
	    pending_beatmapset_count: number;
	    ranked_beatmapset_count: number;
	    replays_watched_count: number;
	    scores_best_count: number;
	    scores_first_count: number;
	    scores_recent_count: number;
	    follower_count: number;
	    // Go type: Country
	    country: any;
	    // Go type: Cover
	    cover: any;
	    monthly_playcounts: UserMonthlyPlaycount[];
	    previous_usernames: string[];
	    // Go type: RankHistory
	    rank_history: any;
	    support_level: number;
	    user_achievements: UserAchievements[];
	    // Go type: UserStatistics
	    statistics?: any;
	
	    static createFrom(source: any = {}) {
	        return new UserCompact(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.username = source["username"];
	        this.avatar_url = source["avatar_url"];
	        this.country_code = source["country_code"];
	        this.default_group = source["default_group"];
	        this.is_active = source["is_active"];
	        this.is_bot = source["is_bot"];
	        this.is_deleted = source["is_deleted"];
	        this.is_online = source["is_online"];
	        this.is_supporter = source["is_supporter"];
	        this.last_visit = this.convertValues(source["last_visit"], null);
	        this.pm_friends_only = source["pm_friends_only"];
	        this.profile_colour = source["profile_colour"];
	        this.account_history = this.convertValues(source["account_history"], UserAccountHistory);
	        this.active_tournament_banner = this.convertValues(source["active_tournament_banner"], null);
	        this.badges = this.convertValues(source["badges"], UserBadge);
	        this.groups = this.convertValues(source["groups"], UserGroups);
	        this.beatmap_playcounts_count = source["beatmap_playcounts_count"];
	        this.favorite_beatmapset_count = source["favorite_beatmapset_count"];
	        this.graveyard_beatmapset_count = source["graveyard_beatmapset_count"];
	        this.loved_beatmapset_count = source["loved_beatmapset_count"];
	        this.pending_beatmapset_count = source["pending_beatmapset_count"];
	        this.ranked_beatmapset_count = source["ranked_beatmapset_count"];
	        this.replays_watched_count = source["replays_watched_count"];
	        this.scores_best_count = source["scores_best_count"];
	        this.scores_first_count = source["scores_first_count"];
	        this.scores_recent_count = source["scores_recent_count"];
	        this.follower_count = source["follower_count"];
	        this.country = this.convertValues(source["country"], null);
	        this.cover = this.convertValues(source["cover"], null);
	        this.monthly_playcounts = this.convertValues(source["monthly_playcounts"], UserMonthlyPlaycount);
	        this.previous_usernames = source["previous_usernames"];
	        this.rank_history = this.convertValues(source["rank_history"], null);
	        this.support_level = source["support_level"];
	        this.user_achievements = this.convertValues(source["user_achievements"], UserAchievements);
	        this.statistics = this.convertValues(source["statistics"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Level {
	    current: number;
	    progress: number;
	
	    static createFrom(source: any = {}) {
	        return new Level(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.current = source["current"];
	        this.progress = source["progress"];
	    }
	}
	export class GradeCounts {
	    a: number;
	    s: number;
	    sh: number;
	    ss: number;
	    ssh: number;
	
	    static createFrom(source: any = {}) {
	        return new GradeCounts(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.a = source["a"];
	        this.s = source["s"];
	        this.sh = source["sh"];
	        this.ss = source["ss"];
	        this.ssh = source["ssh"];
	    }
	}
	export class UserStatistics {
	    // Go type: GradeCounts
	    grade_counts: any;
	    // Go type: Level
	    level: any;
	    hit_accuracy: number;
	    is_ranked: boolean;
	    maximum_combo: number;
	    play_count: number;
	    play_time: number;
	    pp: number;
	    global_rank?: number;
	    ranked_score: number;
	    replays_watched_by_others: number;
	    total_hits: number;
	    total_score: number;
	    // Go type: UserCompact
	    user: any;
	
	    static createFrom(source: any = {}) {
	        return new UserStatistics(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.grade_counts = this.convertValues(source["grade_counts"], null);
	        this.level = this.convertValues(source["level"], null);
	        this.hit_accuracy = source["hit_accuracy"];
	        this.is_ranked = source["is_ranked"];
	        this.maximum_combo = source["maximum_combo"];
	        this.play_count = source["play_count"];
	        this.play_time = source["play_time"];
	        this.pp = source["pp"];
	        this.global_rank = source["global_rank"];
	        this.ranked_score = source["ranked_score"];
	        this.replays_watched_by_others = source["replays_watched_by_others"];
	        this.total_hits = source["total_hits"];
	        this.total_score = source["total_score"];
	        this.user = this.convertValues(source["user"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class UserAchievements {
	    // Go type: DateTime
	    achieved_at: any;
	    achievement_id: number;
	
	    static createFrom(source: any = {}) {
	        return new UserAchievements(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.achieved_at = this.convertValues(source["achieved_at"], null);
	        this.achievement_id = source["achievement_id"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class RankHistory {
	    mode: string;
	    data: number[];
	
	    static createFrom(source: any = {}) {
	        return new RankHistory(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.mode = source["mode"];
	        this.data = source["data"];
	    }
	}
	export class UserMonthlyPlaycount {
	    // Go type: DateTime
	    start_date: any;
	    count: number;
	
	    static createFrom(source: any = {}) {
	        return new UserMonthlyPlaycount(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.start_date = this.convertValues(source["start_date"], null);
	        this.count = source["count"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Cover {
	    custom_url: string;
	    url: string;
	
	    static createFrom(source: any = {}) {
	        return new Cover(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.custom_url = source["custom_url"];
	        this.url = source["url"];
	    }
	}
	export class Country {
	    code: string;
	    name: string;
	
	    static createFrom(source: any = {}) {
	        return new Country(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.code = source["code"];
	        this.name = source["name"];
	    }
	}
	export class UserGroups {
	    playmodes?: string[];
	
	    static createFrom(source: any = {}) {
	        return new UserGroups(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.playmodes = source["playmodes"];
	    }
	}
	export class UserBadge {
	    url: string;
	    image_url: string;
	    description: string;
	    // Go type: DateTime
	    awarded_at: any;
	
	    static createFrom(source: any = {}) {
	        return new UserBadge(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.url = source["url"];
	        this.image_url = source["image_url"];
	        this.description = source["description"];
	        this.awarded_at = this.convertValues(source["awarded_at"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class TournamentBanner {
	    id: number;
	    tournament_id: number;
	    image: string;
	
	    static createFrom(source: any = {}) {
	        return new TournamentBanner(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.tournament_id = source["tournament_id"];
	        this.image = source["image"];
	    }
	}
	export class UserAccountHistory {
	    id: number;
	    description?: string;
	    length: number;
	    // Go type: DateTime
	    timestamp: any;
	    type: string;
	
	    static createFrom(source: any = {}) {
	        return new UserAccountHistory(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.description = source["description"];
	        this.length = source["length"];
	        this.timestamp = this.convertValues(source["timestamp"], null);
	        this.type = source["type"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class User {
	    id: number;
	    username: string;
	    avatar_url: string;
	    country_code: string;
	    default_group: string;
	    is_active: boolean;
	    is_bot: boolean;
	    is_deleted: boolean;
	    is_online: boolean;
	    is_supporter: boolean;
	    // Go type: DateTime
	    last_visit?: any;
	    pm_friends_only: boolean;
	    profile_colour?: string;
	    account_history: UserAccountHistory[];
	    // Go type: TournamentBanner
	    active_tournament_banner?: any;
	    badges: UserBadge[];
	    groups: UserGroups[];
	    beatmap_playcounts_count: number;
	    favorite_beatmapset_count: number;
	    graveyard_beatmapset_count: number;
	    loved_beatmapset_count: number;
	    pending_beatmapset_count: number;
	    ranked_beatmapset_count: number;
	    replays_watched_count: number;
	    scores_best_count: number;
	    scores_first_count: number;
	    scores_recent_count: number;
	    follower_count: number;
	    // Go type: Country
	    country: any;
	    // Go type: Cover
	    cover: any;
	    monthly_playcounts: UserMonthlyPlaycount[];
	    previous_usernames: string[];
	    // Go type: RankHistory
	    rank_history: any;
	    support_level: number;
	    user_achievements: UserAchievements[];
	    // Go type: UserStatistics
	    statistics?: any;
	    cover_url: string;
	    discord?: string;
	    has_supported: boolean;
	    interests?: string;
	    // Go type: DateTime
	    join_date: any;
	    kudosu: Kudosu;
	    location?: string;
	    max_blocks?: number;
	    max_friends?: number;
	    occupation?: string;
	    playmode: string;
	    playstyle: string[];
	    post_count: number;
	    title?: string;
	    title_url?: string;
	    twitter?: string;
	    website?: string;
	
	    static createFrom(source: any = {}) {
	        return new User(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.username = source["username"];
	        this.avatar_url = source["avatar_url"];
	        this.country_code = source["country_code"];
	        this.default_group = source["default_group"];
	        this.is_active = source["is_active"];
	        this.is_bot = source["is_bot"];
	        this.is_deleted = source["is_deleted"];
	        this.is_online = source["is_online"];
	        this.is_supporter = source["is_supporter"];
	        this.last_visit = this.convertValues(source["last_visit"], null);
	        this.pm_friends_only = source["pm_friends_only"];
	        this.profile_colour = source["profile_colour"];
	        this.account_history = this.convertValues(source["account_history"], UserAccountHistory);
	        this.active_tournament_banner = this.convertValues(source["active_tournament_banner"], null);
	        this.badges = this.convertValues(source["badges"], UserBadge);
	        this.groups = this.convertValues(source["groups"], UserGroups);
	        this.beatmap_playcounts_count = source["beatmap_playcounts_count"];
	        this.favorite_beatmapset_count = source["favorite_beatmapset_count"];
	        this.graveyard_beatmapset_count = source["graveyard_beatmapset_count"];
	        this.loved_beatmapset_count = source["loved_beatmapset_count"];
	        this.pending_beatmapset_count = source["pending_beatmapset_count"];
	        this.ranked_beatmapset_count = source["ranked_beatmapset_count"];
	        this.replays_watched_count = source["replays_watched_count"];
	        this.scores_best_count = source["scores_best_count"];
	        this.scores_first_count = source["scores_first_count"];
	        this.scores_recent_count = source["scores_recent_count"];
	        this.follower_count = source["follower_count"];
	        this.country = this.convertValues(source["country"], null);
	        this.cover = this.convertValues(source["cover"], null);
	        this.monthly_playcounts = this.convertValues(source["monthly_playcounts"], UserMonthlyPlaycount);
	        this.previous_usernames = source["previous_usernames"];
	        this.rank_history = this.convertValues(source["rank_history"], null);
	        this.support_level = source["support_level"];
	        this.user_achievements = this.convertValues(source["user_achievements"], UserAchievements);
	        this.statistics = this.convertValues(source["statistics"], null);
	        this.cover_url = source["cover_url"];
	        this.discord = source["discord"];
	        this.has_supported = source["has_supported"];
	        this.interests = source["interests"];
	        this.join_date = this.convertValues(source["join_date"], null);
	        this.kudosu = this.convertValues(source["kudosu"], Kudosu);
	        this.location = source["location"];
	        this.max_blocks = source["max_blocks"];
	        this.max_friends = source["max_friends"];
	        this.occupation = source["occupation"];
	        this.playmode = source["playmode"];
	        this.playstyle = source["playstyle"];
	        this.post_count = source["post_count"];
	        this.title = source["title"];
	        this.title_url = source["title_url"];
	        this.twitter = source["twitter"];
	        this.website = source["website"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

