export namespace app {
	
	export class osuAuthCredentials {
	    client_id: string;
	    client_secret: string;
	
	    static createFrom(source: any = {}) {
	        return new osuAuthCredentials(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.client_id = source["client_id"];
	        this.client_secret = source["client_secret"];
	    }
	}

}

export namespace model {
	
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
	export class MyUserStruct {
	    id: number;
	    username: string;
	    avatar_url: string;
	    country_code: string;
	    country: Country;
	    cover_url: string;
	    playstyle: string[];
	
	    static createFrom(source: any = {}) {
	        return new MyUserStruct(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.username = source["username"];
	        this.avatar_url = source["avatar_url"];
	        this.country_code = source["country_code"];
	        this.country = this.convertValues(source["country"], Country);
	        this.cover_url = source["cover_url"];
	        this.playstyle = source["playstyle"];
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
	export class UserStats {
	    count300: number;
	    count100: number;
	    count50: number;
	    playcount: number;
	    ranked_score: string;
	    total_score: string;
	    pp_rank: number;
	    level: number;
	    pp_raw: number;
	    accuracy: number;
	    count_rank_ss: number;
	    count_rank_s: number;
	    count_rank_a: number;
	    // Go type: time
	    timestamp: any;
	
	    static createFrom(source: any = {}) {
	        return new UserStats(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.count300 = source["count300"];
	        this.count100 = source["count100"];
	        this.count50 = source["count50"];
	        this.playcount = source["playcount"];
	        this.ranked_score = source["ranked_score"];
	        this.total_score = source["total_score"];
	        this.pp_rank = source["pp_rank"];
	        this.level = source["level"];
	        this.pp_raw = source["pp_raw"];
	        this.accuracy = source["accuracy"];
	        this.count_rank_ss = source["count_rank_ss"];
	        this.count_rank_s = source["count_rank_s"];
	        this.count_rank_a = source["count_rank_a"];
	        this.timestamp = this.convertValues(source["timestamp"], null);
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

