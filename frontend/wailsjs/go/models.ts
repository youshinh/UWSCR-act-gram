export namespace main {
	
	export class LayerConfig {
	    Provider: string;
	    Model: string;
	
	    static createFrom(source: any = {}) {
	        return new LayerConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Provider = source["Provider"];
	        this.Model = source["Model"];
	    }
	}
	export class Config {
	    Port: number;
	    UWSCRPath: string;
	    Layers: Record<string, LayerConfig>;
	
	    static createFrom(source: any = {}) {
	        return new Config(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Port = source["Port"];
	        this.UWSCRPath = source["UWSCRPath"];
	        this.Layers = this.convertValues(source["Layers"], LayerConfig, true);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
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
	
	export class SessionContext {
	    active_title: string;
	    screenshot_path: string;
	    screenshot_base64: string;
	
	    static createFrom(source: any = {}) {
	        return new SessionContext(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.active_title = source["active_title"];
	        this.screenshot_path = source["screenshot_path"];
	        this.screenshot_base64 = source["screenshot_base64"];
	    }
	}

}

