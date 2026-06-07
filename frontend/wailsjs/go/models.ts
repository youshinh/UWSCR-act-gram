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
	    KnowledgeDir: string;
	    UWSCRDocURL: string;
	    CustomBaseURL: string;
	    LocalLLMType: string;
	    LocalLLMURL: string;
	    TestTimeout: number;
	    Layers: Record<string, LayerConfig>;
	    UseUnifiedModel: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Config(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Port = source["Port"];
	        this.UWSCRPath = source["UWSCRPath"];
	        this.KnowledgeDir = source["KnowledgeDir"];
	        this.UWSCRDocURL = source["UWSCRDocURL"];
	        this.CustomBaseURL = source["CustomBaseURL"];
	        this.LocalLLMType = source["LocalLLMType"];
	        this.LocalLLMURL = source["LocalLLMURL"];
	        this.TestTimeout = source["TestTimeout"];
	        this.Layers = this.convertValues(source["Layers"], LayerConfig, true);
	        this.UseUnifiedModel = source["UseUnifiedModel"];
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
	
	export class LocalLLMConfig {
	    type: string;
	    url: string;
	
	    static createFrom(source: any = {}) {
	        return new LocalLLMConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.type = source["type"];
	        this.url = source["url"];
	    }
	}
	export class TestRunResult {
	    logs: string;
	    success: boolean;
	
	    static createFrom(source: any = {}) {
	        return new TestRunResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.logs = source["logs"];
	        this.success = source["success"];
	    }
	}

}

export namespace manual {
	
	export class ManualStep {
	    step_id: number;
	    title: string;
	    instruction: string;
	    click_x: number;
	    click_y: number;
	    uws_code: string;
	    image_path: string;
	
	    static createFrom(source: any = {}) {
	        return new ManualStep(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.step_id = source["step_id"];
	        this.title = source["title"];
	        this.instruction = source["instruction"];
	        this.click_x = source["click_x"];
	        this.click_y = source["click_y"];
	        this.uws_code = source["uws_code"];
	        this.image_path = source["image_path"];
	    }
	}

}

