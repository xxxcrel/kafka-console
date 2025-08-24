export namespace kconsole {
	
	export class BrokerConfigSynonym {
	    name: string;
	    value?: string;
	    source: string;
	
	    static createFrom(source: any = {}) {
	        return new BrokerConfigSynonym(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.value = source["value"];
	        this.source = source["source"];
	    }
	}
	export class BrokerConfigEntry {
	    name: string;
	    value?: string;
	    source: string;
	    type: string;
	    isExplicitlySet: boolean;
	    isDefaultValue: boolean;
	    isReadOnly: boolean;
	    isSensitive: boolean;
	    synonyms: BrokerConfigSynonym[];
	
	    static createFrom(source: any = {}) {
	        return new BrokerConfigEntry(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.value = source["value"];
	        this.source = source["source"];
	        this.type = source["type"];
	        this.isExplicitlySet = source["isExplicitlySet"];
	        this.isDefaultValue = source["isDefaultValue"];
	        this.isReadOnly = source["isReadOnly"];
	        this.isSensitive = source["isSensitive"];
	        this.synonyms = this.convertValues(source["synonyms"], BrokerConfigSynonym);
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
	export class BrokerConfig {
	    configs?: BrokerConfigEntry[];
	    error?: string;
	
	    static createFrom(source: any = {}) {
	        return new BrokerConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.configs = this.convertValues(source["configs"], BrokerConfigEntry);
	        this.error = source["error"];
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
	export class Broker {
	    brokerId: number;
	    logDirSize: number;
	    address: string;
	    rack?: string;
	    config: BrokerConfig;
	
	    static createFrom(source: any = {}) {
	        return new Broker(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.brokerId = source["brokerId"];
	        this.logDirSize = source["logDirSize"];
	        this.address = source["address"];
	        this.rack = source["rack"];
	        this.config = this.convertValues(source["config"], BrokerConfig);
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
	
	
	
	export class BrokerWithLogDirs {
	    brokerId: number;
	    isController: boolean;
	    address?: string;
	    rack?: string;
	    totalLogDirSizeBytes?: number;
	    totalPrimaryLogDirSizeBytes?: number;
	
	    static createFrom(source: any = {}) {
	        return new BrokerWithLogDirs(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.brokerId = source["brokerId"];
	        this.isController = source["isController"];
	        this.address = source["address"];
	        this.rack = source["rack"];
	        this.totalLogDirSizeBytes = source["totalLogDirSizeBytes"];
	        this.totalPrimaryLogDirSizeBytes = source["totalPrimaryLogDirSizeBytes"];
	    }
	}
	export class ClusterInfo {
	    controllerId: number;
	    brokers: Broker[];
	    kafkaVersion: string;
	
	    static createFrom(source: any = {}) {
	        return new ClusterInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.controllerId = source["controllerId"];
	        this.brokers = this.convertValues(source["brokers"], Broker);
	        this.kafkaVersion = source["kafkaVersion"];
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
	export class PartitionOffsets {
	    error?: string;
	    partitionId: number;
	    groupOffset: number;
	    highWaterMark: number;
	    lag: number;
	
	    static createFrom(source: any = {}) {
	        return new PartitionOffsets(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.error = source["error"];
	        this.partitionId = source["partitionId"];
	        this.groupOffset = source["groupOffset"];
	        this.highWaterMark = source["highWaterMark"];
	        this.lag = source["lag"];
	    }
	}
	export class GroupTopicOffsets {
	    topic: string;
	    summedLag: number;
	    partitionCount: number;
	    partitionsWithOffset: number;
	    partitionOffsets: PartitionOffsets[];
	
	    static createFrom(source: any = {}) {
	        return new GroupTopicOffsets(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.topic = source["topic"];
	        this.summedLag = source["summedLag"];
	        this.partitionCount = source["partitionCount"];
	        this.partitionsWithOffset = source["partitionsWithOffset"];
	        this.partitionOffsets = this.convertValues(source["partitionOffsets"], PartitionOffsets);
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
	export class GroupMemberAssignment {
	    topicName: string;
	    partitionIds: number[];
	
	    static createFrom(source: any = {}) {
	        return new GroupMemberAssignment(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.topicName = source["topicName"];
	        this.partitionIds = source["partitionIds"];
	    }
	}
	export class GroupMemberDescription {
	    id: string;
	    clientId: string;
	    clientHost: string;
	    assignments: GroupMemberAssignment[];
	
	    static createFrom(source: any = {}) {
	        return new GroupMemberDescription(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.clientId = source["clientId"];
	        this.clientHost = source["clientHost"];
	        this.assignments = this.convertValues(source["assignments"], GroupMemberAssignment);
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
	export class ConsumerGroupOverview {
	    groupId: string;
	    state: string;
	    protocolType: string;
	    protocol: string;
	    members: GroupMemberDescription[];
	    coordinatorId: number;
	    topicOffsets: GroupTopicOffsets[];
	
	    static createFrom(source: any = {}) {
	        return new ConsumerGroupOverview(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.groupId = source["groupId"];
	        this.state = source["state"];
	        this.protocolType = source["protocolType"];
	        this.protocol = source["protocol"];
	        this.members = this.convertValues(source["members"], GroupMemberDescription);
	        this.coordinatorId = source["coordinatorId"];
	        this.topicOffsets = this.convertValues(source["topicOffsets"], GroupTopicOffsets);
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
	export class EndpointCompatibilityEndpoint {
	    endpoint: string;
	    method: string;
	    isSupported: boolean;
	
	    static createFrom(source: any = {}) {
	        return new EndpointCompatibilityEndpoint(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.endpoint = source["endpoint"];
	        this.method = source["method"];
	        this.isSupported = source["isSupported"];
	    }
	}
	export class EndpointCompatibility {
	    kafkaVersion: string;
	    endpoints: EndpointCompatibilityEndpoint[];
	
	    static createFrom(source: any = {}) {
	        return new EndpointCompatibility(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.kafkaVersion = source["kafkaVersion"];
	        this.endpoints = this.convertValues(source["endpoints"], EndpointCompatibilityEndpoint);
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
	
	
	
	

}

export namespace rest {
	
	export class Error {
	    statusCode: number;
	    message: string;
	
	    static createFrom(source: any = {}) {
	        return new Error(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.statusCode = source["statusCode"];
	        this.message = source["message"];
	    }
	}

}

export namespace zapcore {
	
	export class Field {
	    Key: string;
	    Type: number;
	    Integer: number;
	    String: string;
	    Interface: any;
	
	    static createFrom(source: any = {}) {
	        return new Field(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Key = source["Key"];
	        this.Type = source["Type"];
	        this.Integer = source["Integer"];
	        this.String = source["String"];
	        this.Interface = source["Interface"];
	    }
	}

}

