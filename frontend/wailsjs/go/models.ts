export namespace clusterstatus {
	
	export class ComponentStatus {
	    status: number;
	    statusReason: string;
	
	    static createFrom(source: any = {}) {
	        return new ComponentStatus(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.status = source["status"];
	        this.statusReason = source["statusReason"];
	    }
	}
	export class KafkaBroker {
	    brokerId: number;
	    host: string;
	    rackId?: string;
	
	    static createFrom(source: any = {}) {
	        return new KafkaBroker(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.brokerId = source["brokerId"];
	        this.host = source["host"];
	        this.rackId = source["rackId"];
	    }
	}
	export class KafkaConnectCluster {
	    name?: string;
	    status?: ComponentStatus;
	    host?: string;
	    version?: string;
	    installed_plugins_count?: number;
	
	    static createFrom(source: any = {}) {
	        return new KafkaConnectCluster(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.status = this.convertValues(source["status"], ComponentStatus);
	        this.host = source["host"];
	        this.version = source["version"];
	        this.installed_plugins_count = source["installed_plugins_count"];
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
	export class KafkaInfo {
	    status?: ComponentStatus;
	    version: string;
	    distribution?: number;
	    brokersOnline?: number;
	    brokersExpected?: number;
	    topicsCount?: number;
	    partitionsCount?: number;
	    replicasCount?: number;
	    controllerId?: number;
	    brokers?: KafkaBroker[];
	    clusterId?: string;
	
	    static createFrom(source: any = {}) {
	        return new KafkaInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.status = this.convertValues(source["status"], ComponentStatus);
	        this.version = source["version"];
	        this.distribution = source["distribution"];
	        this.brokersOnline = source["brokersOnline"];
	        this.brokersExpected = source["brokersExpected"];
	        this.topicsCount = source["topicsCount"];
	        this.partitionsCount = source["partitionsCount"];
	        this.replicasCount = source["replicasCount"];
	        this.controllerId = source["controllerId"];
	        this.brokers = this.convertValues(source["brokers"], KafkaBroker);
	        this.clusterId = source["clusterId"];
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
	export class SchemaRegistryInfo {
	    status?: ComponentStatus;
	    registered_subjects_count?: number;
	
	    static createFrom(source: any = {}) {
	        return new SchemaRegistryInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.status = this.convertValues(source["status"], ComponentStatus);
	        this.registered_subjects_count = source["registered_subjects_count"];
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

export namespace console {
	
	export enum PayloadEncoding {
	    UNSPECIFIED = 0,
	    NULL = 1,
	    AVRO = 2,
	    PROTOBUF = 3,
	    PROTOBUF_SCHEMA = 4,
	    JSON = 5,
	    JSON_SCHEMA = 6,
	    XML = 7,
	    TEXT = 8,
	    UTF8 = 9,
	    MESSAGE_PACK = 10,
	    SMILE = 11,
	    BINARY = 12,
	    UINT = 13,
	    CONSUMER_OFFSETS = 14,
	    CBOR = 15,
	}
	export enum CompressionType {
	    UNSPECIFIED = 0,
	    UNCOMPRESSED = 1,
	    GZIP = 2,
	    SNAPPY = 3,
	    LZ4 = 4,
	    ZSTD = 5,
	}

}

export namespace kconsole {
	
	export enum FrontendFormat {
	    UNKNOWN = 0,
	    BOOLEAN = 1,
	    PASSWORD = 2,
	    STRING = 3,
	    SELECT = 4,
	    MULTI_SELECT = 5,
	    BYTE_SIZE = 6,
	    RATIO = 7,
	    DURATION = 8,
	    DECIMAL = 9,
	    INTEGER = 10,
	}
	export class ACLRule {
	    principal: string;
	    principalType: string;
	    principalName: string;
	    host: string;
	    operation: string;
	    permissionType: string;
	
	    static createFrom(source: any = {}) {
	        return new ACLRule(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.principal = source["principal"];
	        this.principalType = source["principalType"];
	        this.principalName = source["principalName"];
	        this.host = source["host"];
	        this.operation = source["operation"];
	        this.permissionType = source["permissionType"];
	    }
	}
	export class ACLResource {
	    resourceType: string;
	    resourceName: string;
	    resourcePatternType: string;
	    acls: ACLRule[];
	
	    static createFrom(source: any = {}) {
	        return new ACLResource(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.resourceType = source["resourceType"];
	        this.resourceName = source["resourceName"];
	        this.resourcePatternType = source["resourcePatternType"];
	        this.acls = this.convertValues(source["acls"], ACLRule);
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
	export class ACLOverview {
	    aclResources: ACLResource[];
	    isAuthorizerEnabled: boolean;
	
	    static createFrom(source: any = {}) {
	        return new ACLOverview(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.aclResources = this.convertValues(source["aclResources"], ACLResource);
	        this.isAuthorizerEnabled = source["isAuthorizerEnabled"];
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
	
	
	export class APIVersion {
	    keyId: number;
	    keyName: string;
	    maxVersion: number;
	    minVersion: number;
	
	    static createFrom(source: any = {}) {
	        return new APIVersion(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.keyId = source["keyId"];
	        this.keyName = source["keyName"];
	        this.maxVersion = source["maxVersion"];
	        this.minVersion = source["minVersion"];
	    }
	}
	export class AlterPartitionReassignmentsPartitionResponse {
	    partitionId: number;
	    errorCode: string;
	    errorMessage?: string;
	
	    static createFrom(source: any = {}) {
	        return new AlterPartitionReassignmentsPartitionResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.partitionId = source["partitionId"];
	        this.errorCode = source["errorCode"];
	        this.errorMessage = source["errorMessage"];
	    }
	}
	export class AlterPartitionReassignmentsResponse {
	    topicName: string;
	    partitions: AlterPartitionReassignmentsPartitionResponse[];
	
	    static createFrom(source: any = {}) {
	        return new AlterPartitionReassignmentsResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.topicName = source["topicName"];
	        this.partitions = this.convertValues(source["partitions"], AlterPartitionReassignmentsPartitionResponse);
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
	export class BrokerConfigSynonym {
	    name: string;
	    value?: string;
	    source: string;
	    type: string;
	
	    static createFrom(source: any = {}) {
	        return new BrokerConfigSynonym(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.value = source["value"];
	        this.source = source["source"];
	        this.type = source["type"];
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
	    lagSum: number;
	    isInUse: boolean;
	    noEditPerms: boolean;
	    noDeletePerms: boolean;
	    allowedActions: string[];
	
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
	        this.lagSum = source["lagSum"];
	        this.isInUse = source["isInUse"];
	        this.noEditPerms = source["noEditPerms"];
	        this.noDeletePerms = source["noDeletePerms"];
	        this.allowedActions = source["allowedActions"];
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
	export class CreateSchemaResponse {
	    id: number;
	
	    static createFrom(source: any = {}) {
	        return new CreateSchemaResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	    }
	}
	export class CreateTopicResponseConfig {
	    name: string;
	    value?: string;
	
	    static createFrom(source: any = {}) {
	        return new CreateTopicResponseConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.value = source["value"];
	    }
	}
	export class CreateTopicResponse {
	    topicName: string;
	    partitionCount: number;
	    replicationFactor: number;
	    configs: CreateTopicResponseConfig[];
	
	    static createFrom(source: any = {}) {
	        return new CreateTopicResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.topicName = source["topicName"];
	        this.partitionCount = source["partitionCount"];
	        this.replicationFactor = source["replicationFactor"];
	        this.configs = this.convertValues(source["configs"], CreateTopicResponseConfig);
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
	
	export class DeleteACLsResponse {
	    errorMessage: string[];
	    matchedACLs: number;
	    deletedACLs: number;
	
	    static createFrom(source: any = {}) {
	        return new DeleteACLsResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.errorMessage = source["errorMessage"];
	        this.matchedACLs = source["matchedACLs"];
	        this.deletedACLs = source["deletedACLs"];
	    }
	}
	export class DeleteConsumerGroupOffsetsResponseTopicPartition {
	    partitionID: number;
	    error?: string;
	
	    static createFrom(source: any = {}) {
	        return new DeleteConsumerGroupOffsetsResponseTopicPartition(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.partitionID = source["partitionID"];
	        this.error = source["error"];
	    }
	}
	export class DeleteConsumerGroupOffsetsResponseTopic {
	    topicName: string;
	    partitions: DeleteConsumerGroupOffsetsResponseTopicPartition[];
	
	    static createFrom(source: any = {}) {
	        return new DeleteConsumerGroupOffsetsResponseTopic(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.topicName = source["topicName"];
	        this.partitions = this.convertValues(source["partitions"], DeleteConsumerGroupOffsetsResponseTopicPartition);
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
	
	export class DeleteTopicRecordsResponsePartition {
	    partitionId: number;
	    lowWaterMark: number;
	    error?: string;
	
	    static createFrom(source: any = {}) {
	        return new DeleteTopicRecordsResponsePartition(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.partitionId = source["partitionId"];
	        this.lowWaterMark = source["lowWaterMark"];
	        this.error = source["error"];
	    }
	}
	export class DeleteTopicRecordsResponse {
	    topicName: string;
	    partitions: DeleteTopicRecordsResponsePartition[];
	
	    static createFrom(source: any = {}) {
	        return new DeleteTopicRecordsResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.topicName = source["topicName"];
	        this.partitions = this.convertValues(source["partitions"], DeleteTopicRecordsResponsePartition);
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
	
	export class EditConsumerGroupOffsetsResponseTopicPartition {
	    partitionID: number;
	    error?: string;
	
	    static createFrom(source: any = {}) {
	        return new EditConsumerGroupOffsetsResponseTopicPartition(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.partitionID = source["partitionID"];
	        this.error = source["error"];
	    }
	}
	export class EditConsumerGroupOffsetsResponseTopic {
	    topicName: string;
	    partitions: EditConsumerGroupOffsetsResponseTopicPartition[];
	
	    static createFrom(source: any = {}) {
	        return new EditConsumerGroupOffsetsResponseTopic(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.topicName = source["topicName"];
	        this.partitions = this.convertValues(source["partitions"], EditConsumerGroupOffsetsResponseTopicPartition);
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
	export class EditConsumerGroupOffsetsResponse {
	    error?: string;
	    topics: EditConsumerGroupOffsetsResponseTopic[];
	
	    static createFrom(source: any = {}) {
	        return new EditConsumerGroupOffsetsResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.error = source["error"];
	        this.topics = this.convertValues(source["topics"], EditConsumerGroupOffsetsResponseTopic);
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
	
	
	
	
	export class IncrementalAlterConfigsResourceResponse {
	    error?: string;
	    resourceName: string;
	    resourceType: number;
	
	    static createFrom(source: any = {}) {
	        return new IncrementalAlterConfigsResourceResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.error = source["error"];
	        this.resourceName = source["resourceName"];
	        this.resourceType = source["resourceType"];
	    }
	}
	export class KafkaError {
	    code: number;
	    message: string;
	    description: string;
	
	    static createFrom(source: any = {}) {
	        return new KafkaError(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.code = source["code"];
	        this.message = source["message"];
	        this.description = source["description"];
	    }
	}
	export class ListMessageRequest {
	    TopicName: string;
	    PartitionID: number;
	    StartOffset: number;
	    StartTimestamp: number;
	    MessageCount: number;
	    FilterInterpreterCode: string;
	    Troubleshoot: boolean;
	    IncludeRawPayload: boolean;
	    IgnoreMaxSizeLimit: boolean;
	    KeyDeserializer: string;
	    ValueDeserializer: string;
	
	    static createFrom(source: any = {}) {
	        return new ListMessageRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.TopicName = source["TopicName"];
	        this.PartitionID = source["PartitionID"];
	        this.StartOffset = source["StartOffset"];
	        this.StartTimestamp = source["StartTimestamp"];
	        this.MessageCount = source["MessageCount"];
	        this.FilterInterpreterCode = source["FilterInterpreterCode"];
	        this.Troubleshoot = source["Troubleshoot"];
	        this.IncludeRawPayload = source["IncludeRawPayload"];
	        this.IgnoreMaxSizeLimit = source["IgnoreMaxSizeLimit"];
	        this.KeyDeserializer = source["KeyDeserializer"];
	        this.ValueDeserializer = source["ValueDeserializer"];
	    }
	}
	export class MessageHeader {
	    key: string;
	    value: number[];
	    isValueTooLarge: boolean;
	    encoding: string;
	
	    static createFrom(source: any = {}) {
	        return new MessageHeader(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.key = source["key"];
	        this.value = source["value"];
	        this.isValueTooLarge = source["isValueTooLarge"];
	        this.encoding = source["encoding"];
	    }
	}
	export class PartitionOffset {
	    error?: string;
	    partitionId: number;
	    offset: number;
	    timestamp: number;
	
	    static createFrom(source: any = {}) {
	        return new PartitionOffset(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.error = source["error"];
	        this.partitionId = source["partitionId"];
	        this.offset = source["offset"];
	        this.timestamp = source["timestamp"];
	    }
	}
	
	export class PartitionReassignmentsPartition {
	    partitionId: number;
	    addingReplicas: number[];
	    removingReplicas: number[];
	    replicas: number[];
	
	    static createFrom(source: any = {}) {
	        return new PartitionReassignmentsPartition(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.partitionId = source["partitionId"];
	        this.addingReplicas = source["addingReplicas"];
	        this.removingReplicas = source["removingReplicas"];
	        this.replicas = source["replicas"];
	    }
	}
	export class PartitionReassignments {
	    topicName: string;
	    partitions: PartitionReassignmentsPartition[];
	
	    static createFrom(source: any = {}) {
	        return new PartitionReassignments(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.topicName = source["topicName"];
	        this.partitions = this.convertValues(source["partitions"], PartitionReassignmentsPartition);
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
	
	export class ProduceRecordResponse {
	    topicName: string;
	    partitionId: number;
	    offset: number;
	    error?: string;
	    keyTroubleshooting?: serde.TroubleshootingReport[];
	    valueTroubleshooting?: serde.TroubleshootingReport[];
	
	    static createFrom(source: any = {}) {
	        return new ProduceRecordResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.topicName = source["topicName"];
	        this.partitionId = source["partitionId"];
	        this.offset = source["offset"];
	        this.error = source["error"];
	        this.keyTroubleshooting = this.convertValues(source["keyTroubleshooting"], serde.TroubleshootingReport);
	        this.valueTroubleshooting = this.convertValues(source["valueTroubleshooting"], serde.TroubleshootingReport);
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
	export class ProduceRecordsResponse {
	    records: ProduceRecordResponse[];
	    error?: string;
	
	    static createFrom(source: any = {}) {
	        return new ProduceRecordsResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.records = this.convertValues(source["records"], ProduceRecordResponse);
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
	export class QuotaResponseSetting {
	    key: string;
	    value: number;
	
	    static createFrom(source: any = {}) {
	        return new QuotaResponseSetting(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.key = source["key"];
	        this.value = source["value"];
	    }
	}
	export class QuotaResponseItem {
	    entityType: string;
	    entityName: string;
	    settings: QuotaResponseSetting[];
	
	    static createFrom(source: any = {}) {
	        return new QuotaResponseItem(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.entityType = source["entityType"];
	        this.entityName = source["entityName"];
	        this.settings = this.convertValues(source["settings"], QuotaResponseSetting);
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
	export class QuotaResponse {
	    error?: string;
	    items: QuotaResponseItem[];
	
	    static createFrom(source: any = {}) {
	        return new QuotaResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.error = source["error"];
	        this.items = this.convertValues(source["items"], QuotaResponseItem);
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
	
	
	export class Reference {
	    name: string;
	    subject: string;
	    version: number;
	
	    static createFrom(source: any = {}) {
	        return new Reference(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.subject = source["subject"];
	        this.version = source["version"];
	    }
	}
	export class SchemaUsage {
	    subject: string;
	    version: number;
	
	    static createFrom(source: any = {}) {
	        return new SchemaUsage(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.subject = source["subject"];
	        this.version = source["version"];
	    }
	}
	export class SchemaReference {
	    schemaId: number;
	    error?: string;
	    usages: SchemaUsage[];
	
	    static createFrom(source: any = {}) {
	        return new SchemaReference(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.schemaId = source["schemaId"];
	        this.error = source["error"];
	        this.usages = this.convertValues(source["usages"], SchemaUsage);
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
	export class SchemaRegistryConfig {
	    compatibility: sr.CompatibilityLevel;
	
	    static createFrom(source: any = {}) {
	        return new SchemaRegistryConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.compatibility = source["compatibility"];
	    }
	}
	export class SchemaRegistryDeleteSubjectResponse {
	    deletedVersions: number[];
	
	    static createFrom(source: any = {}) {
	        return new SchemaRegistryDeleteSubjectResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.deletedVersions = source["deletedVersions"];
	    }
	}
	export class SchemaRegistryDeleteSubjectVersionResponse {
	    deletedVersion: number;
	
	    static createFrom(source: any = {}) {
	        return new SchemaRegistryDeleteSubjectVersionResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.deletedVersion = source["deletedVersion"];
	    }
	}
	export class SchemaRegistryMode {
	    mode: string;
	
	    static createFrom(source: any = {}) {
	        return new SchemaRegistryMode(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.mode = source["mode"];
	    }
	}
	export class SchemaRegistrySchemaTypes {
	    schemaTypes: number[];
	
	    static createFrom(source: any = {}) {
	        return new SchemaRegistrySchemaTypes(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.schemaTypes = source["schemaTypes"];
	    }
	}
	export class SchemaRegistrySchemaValidationCompatibility {
	    isCompatible: boolean;
	    error?: string;
	
	    static createFrom(source: any = {}) {
	        return new SchemaRegistrySchemaValidationCompatibility(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.isCompatible = source["isCompatible"];
	        this.error = source["error"];
	    }
	}
	export class SchemaRegistrySchemaValidation {
	    compatibility: SchemaRegistrySchemaValidationCompatibility;
	    parsingError?: string;
	    isValid: boolean;
	
	    static createFrom(source: any = {}) {
	        return new SchemaRegistrySchemaValidation(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.compatibility = this.convertValues(source["compatibility"], SchemaRegistrySchemaValidationCompatibility);
	        this.parsingError = source["parsingError"];
	        this.isValid = source["isValid"];
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
	
	export class SchemaRegistrySubject {
	    name: string;
	    isSoftDeleted: boolean;
	
	    static createFrom(source: any = {}) {
	        return new SchemaRegistrySubject(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.isSoftDeleted = source["isSoftDeleted"];
	    }
	}
	export class SchemaRegistryVersionedSchema {
	    id: number;
	    version: number;
	    isSoftDeleted: boolean;
	    type: sr.SchemaType;
	    schema: string;
	    references: Reference[];
	
	    static createFrom(source: any = {}) {
	        return new SchemaRegistryVersionedSchema(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.version = source["version"];
	        this.isSoftDeleted = source["isSoftDeleted"];
	        this.type = source["type"];
	        this.schema = source["schema"];
	        this.references = this.convertValues(source["references"], Reference);
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
	export class SchemaRegistrySubjectDetailsVersion {
	    version: number;
	    isSoftDeleted: boolean;
	
	    static createFrom(source: any = {}) {
	        return new SchemaRegistrySubjectDetailsVersion(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.version = source["version"];
	        this.isSoftDeleted = source["isSoftDeleted"];
	    }
	}
	export class SchemaRegistrySubjectDetails {
	    name: string;
	    type: sr.SchemaType;
	    compatibility?: sr.CompatibilityLevel;
	    versions: SchemaRegistrySubjectDetailsVersion[];
	    latestActiveVersion: number;
	    schemas: SchemaRegistryVersionedSchema[];
	
	    static createFrom(source: any = {}) {
	        return new SchemaRegistrySubjectDetails(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.type = source["type"];
	        this.compatibility = source["compatibility"];
	        this.versions = this.convertValues(source["versions"], SchemaRegistrySubjectDetailsVersion);
	        this.latestActiveVersion = source["latestActiveVersion"];
	        this.schemas = this.convertValues(source["schemas"], SchemaRegistryVersionedSchema);
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
	
	
	
	export class SchemaVersion {
	    subject: string;
	    version: number;
	
	    static createFrom(source: any = {}) {
	        return new SchemaVersion(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.subject = source["subject"];
	        this.version = source["version"];
	    }
	}
	export class TopicConfigSynonym {
	    name: string;
	    value?: string;
	    source: string;
	    type: kmsg.ConfigType;
	
	    static createFrom(source: any = {}) {
	        return new TopicConfigSynonym(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.value = source["value"];
	        this.source = source["source"];
	        this.type = source["type"];
	    }
	}
	export class TopicConfigEntry {
	    name: string;
	    value?: string;
	    source: string;
	    isExplicitlySet: boolean;
	    isDefaultValue: boolean;
	    isSensitive: boolean;
	    isReadOnly: boolean;
	    synonyms: TopicConfigSynonym[];
	    documentation?: string;
	    type: kmsg.ConfigType;
	    category?: string;
	    frontendFormat?: FrontendFormat;
	    enumValues?: string[];
	
	    static createFrom(source: any = {}) {
	        return new TopicConfigEntry(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.value = source["value"];
	        this.source = source["source"];
	        this.isExplicitlySet = source["isExplicitlySet"];
	        this.isDefaultValue = source["isDefaultValue"];
	        this.isSensitive = source["isSensitive"];
	        this.isReadOnly = source["isReadOnly"];
	        this.synonyms = this.convertValues(source["synonyms"], TopicConfigSynonym);
	        this.documentation = source["documentation"];
	        this.type = source["type"];
	        this.category = source["category"];
	        this.frontendFormat = source["frontendFormat"];
	        this.enumValues = source["enumValues"];
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
	export class TopicConfig {
	    topicName: string;
	    configEntries: TopicConfigEntry[];
	    error?: KafkaError;
	
	    static createFrom(source: any = {}) {
	        return new TopicConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.topicName = source["topicName"];
	        this.configEntries = this.convertValues(source["configEntries"], TopicConfigEntry);
	        this.error = this.convertValues(source["error"], KafkaError);
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
	
	
	export class TopicConsumerGroup {
	    groupId: string;
	    summedLag: number;
	
	    static createFrom(source: any = {}) {
	        return new TopicConsumerGroup(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.groupId = source["groupId"];
	        this.summedLag = source["summedLag"];
	    }
	}
	export class TopicPartitionLogDirs {
	    brokerId: number;
	    error?: string;
	    partitionId: number;
	    size: number;
	
	    static createFrom(source: any = {}) {
	        return new TopicPartitionLogDirs(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.brokerId = source["brokerId"];
	        this.error = source["error"];
	        this.partitionId = source["partitionId"];
	        this.size = source["size"];
	    }
	}
	export class TopicPartitionDetails {
	    id: number;
	    partitionError?: string;
	    replicas: number[];
	    offlineReplicas: number[];
	    inSyncReplicas: number[];
	    leader: number;
	    waterMarksError?: string;
	    waterMarkLow: number;
	    waterMarkHigh: number;
	    partitionLogDirs: TopicPartitionLogDirs[];
	    topicName: string;
	    replicaSize: number;
	    hasErrors: boolean;
	
	    static createFrom(source: any = {}) {
	        return new TopicPartitionDetails(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.partitionError = source["partitionError"];
	        this.replicas = source["replicas"];
	        this.offlineReplicas = source["offlineReplicas"];
	        this.inSyncReplicas = source["inSyncReplicas"];
	        this.leader = source["leader"];
	        this.waterMarksError = source["waterMarksError"];
	        this.waterMarkLow = source["waterMarkLow"];
	        this.waterMarkHigh = source["waterMarkHigh"];
	        this.partitionLogDirs = this.convertValues(source["partitionLogDirs"], TopicPartitionLogDirs);
	        this.topicName = source["topicName"];
	        this.replicaSize = source["replicaSize"];
	        this.hasErrors = source["hasErrors"];
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
	export class TopicDetails {
	    topicName: string;
	    error?: string;
	    partitions: TopicPartitionDetails[];
	    replicationFactor: number;
	
	    static createFrom(source: any = {}) {
	        return new TopicDetails(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.topicName = source["topicName"];
	        this.error = source["error"];
	        this.partitions = this.convertValues(source["partitions"], TopicPartitionDetails);
	        this.replicationFactor = source["replicationFactor"];
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
	export class TopicDocumentation {
	    isEnabled: boolean;
	    markdown: number[];
	    text: string;
	
	    static createFrom(source: any = {}) {
	        return new TopicDocumentation(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.isEnabled = source["isEnabled"];
	        this.markdown = source["markdown"];
	        this.text = source["text"];
	    }
	}
	export class TopicLogDirSummaryReplicaError {
	    brokerId: number;
	    error?: string;
	
	    static createFrom(source: any = {}) {
	        return new TopicLogDirSummaryReplicaError(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.brokerId = source["brokerId"];
	        this.error = source["error"];
	    }
	}
	export class TopicLogDirSummary {
	    totalSizeBytes: number;
	    replicaErrors?: TopicLogDirSummaryReplicaError[];
	    hint?: string;
	
	    static createFrom(source: any = {}) {
	        return new TopicLogDirSummary(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.totalSizeBytes = source["totalSizeBytes"];
	        this.replicaErrors = this.convertValues(source["replicaErrors"], TopicLogDirSummaryReplicaError);
	        this.hint = source["hint"];
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
	
	export class TopicMessage {
	    partitionID: number;
	    offset: number;
	    timestamp: number;
	    compression: string;
	    isTransactional: boolean;
	    headers: MessageHeader[];
	    key?: serde.RecordPayload;
	    value?: serde.RecordPayload;
	
	    static createFrom(source: any = {}) {
	        return new TopicMessage(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.partitionID = source["partitionID"];
	        this.offset = source["offset"];
	        this.timestamp = source["timestamp"];
	        this.compression = source["compression"];
	        this.isTransactional = source["isTransactional"];
	        this.headers = this.convertValues(source["headers"], MessageHeader);
	        this.key = this.convertValues(source["key"], serde.RecordPayload);
	        this.value = this.convertValues(source["value"], serde.RecordPayload);
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
	export class TopicOffset {
	    topicName: string;
	    partitions: PartitionOffset[];
	
	    static createFrom(source: any = {}) {
	        return new TopicOffset(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.topicName = source["topicName"];
	        this.partitions = this.convertValues(source["partitions"], PartitionOffset);
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
	
	
	export class TopicSummary {
	    topicName: string;
	    isInternal: boolean;
	    partitionCount: number;
	    replicationFactor: number;
	    cleanupPolicy: string;
	    documentation: string;
	    logDirSummary: TopicLogDirSummary;
	    allowedActions: string[];
	
	    static createFrom(source: any = {}) {
	        return new TopicSummary(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.topicName = source["topicName"];
	        this.isInternal = source["isInternal"];
	        this.partitionCount = source["partitionCount"];
	        this.replicationFactor = source["replicationFactor"];
	        this.cleanupPolicy = source["cleanupPolicy"];
	        this.documentation = source["documentation"];
	        this.logDirSummary = this.convertValues(source["logDirSummary"], TopicLogDirSummary);
	        this.allowedActions = source["allowedActions"];
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

export namespace kgo {
	
	export class CompressionCodec {
	
	
	    static createFrom(source: any = {}) {
	        return new CompressionCodec(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	
	    }
	}
	export class RecordAttrs {
	
	
	    static createFrom(source: any = {}) {
	        return new RecordAttrs(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	
	    }
	}
	export class RecordHeader {
	    Key: string;
	    Value: number[];
	
	    static createFrom(source: any = {}) {
	        return new RecordHeader(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Key = source["Key"];
	        this.Value = source["Value"];
	    }
	}
	export class Record {
	    Key: number[];
	    Value: number[];
	    Headers: RecordHeader[];
	    // Go type: time
	    Timestamp: any;
	    Topic: string;
	    Partition: number;
	    // Go type: RecordAttrs
	    Attrs: any;
	    ProducerEpoch: number;
	    ProducerID: number;
	    LeaderEpoch: number;
	    Offset: number;
	    Context: any;
	
	    static createFrom(source: any = {}) {
	        return new Record(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Key = source["Key"];
	        this.Value = source["Value"];
	        this.Headers = this.convertValues(source["Headers"], RecordHeader);
	        this.Timestamp = this.convertValues(source["Timestamp"], null);
	        this.Topic = source["Topic"];
	        this.Partition = source["Partition"];
	        this.Attrs = this.convertValues(source["Attrs"], null);
	        this.ProducerEpoch = source["ProducerEpoch"];
	        this.ProducerID = source["ProducerID"];
	        this.LeaderEpoch = source["LeaderEpoch"];
	        this.Offset = source["Offset"];
	        this.Context = source["Context"];
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

export namespace kmsg {
	
	export enum IncrementalAlterConfigOp {
	    SET = 0,
	    DELETE = 1,
	    APPEND = 2,
	    SUBTRACT = 3,
	}
	export enum ConfigResourceType {
	    UNKNOWN = 0,
	    TOPIC = 2,
	    BROKER = 4,
	    BROKER_LOGGER = 8,
	}
	export enum ACLOperation {
	    UNKNOWN = 0,
	    ANY = 1,
	    ALL = 2,
	    READ = 3,
	    WRITE = 4,
	    CREATE = 5,
	    DELETE = 6,
	    ALTER = 7,
	    DESCRIBE = 8,
	    CLUSTER_ACTION = 9,
	    DESCRIBE_CONFIGS = 10,
	    ALTER_CONFIGS = 11,
	    IDEMPOTENT_WRITE = 12,
	    CREATE_TOKENS = 13,
	    DESCRIBE_TOKENS = 14,
	}
	export enum ACLPermissionType {
	    UNKNOWN = 0,
	    ANY = 1,
	    DENY = 2,
	    ALLOW = 3,
	}
	export enum ACLResourceType {
	    UNKNOWN = 0,
	    ANY = 1,
	    TOPIC = 2,
	    GROUP = 3,
	    CLUSTER = 4,
	    TRANSACTIONAL_ID = 5,
	    DELEGATION_TOKEN = 6,
	    USER = 7,
	}
	export enum ACLResourcePatternType {
	    UNKNOWN = 0,
	    ANY = 1,
	    MATCH = 2,
	    LITERAL = 3,
	    PREFIXED = 4,
	}
	export enum ConfigType {
	    BOOLEAN = 1,
	    STRING = 2,
	    INT = 3,
	    SHORT = 4,
	    LONG = 5,
	    DOUBLE = 6,
	    LIST = 7,
	    CLASS = 8,
	    PASSWORD = 9,
	}
	export class Tags {
	
	
	    static createFrom(source: any = {}) {
	        return new Tags(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	
	    }
	}
	export class AlterConfigsRequestResourceConfig {
	    Name: string;
	    Value?: string;
	    // Go type: Tags
	    UnknownTags: any;
	
	    static createFrom(source: any = {}) {
	        return new AlterConfigsRequestResourceConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Name = source["Name"];
	        this.Value = source["Value"];
	        this.UnknownTags = this.convertValues(source["UnknownTags"], null);
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
	export class AlterConfigsRequestResource {
	    ResourceType: ConfigResourceType;
	    ResourceName: string;
	    Configs: AlterConfigsRequestResourceConfig[];
	    // Go type: Tags
	    UnknownTags: any;
	
	    static createFrom(source: any = {}) {
	        return new AlterConfigsRequestResource(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ResourceType = source["ResourceType"];
	        this.ResourceName = source["ResourceName"];
	        this.Configs = this.convertValues(source["Configs"], AlterConfigsRequestResourceConfig);
	        this.UnknownTags = this.convertValues(source["UnknownTags"], null);
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
	export class AlterConfigsRequest {
	    Version: number;
	    Resources: AlterConfigsRequestResource[];
	    ValidateOnly: boolean;
	    // Go type: Tags
	    UnknownTags: any;
	
	    static createFrom(source: any = {}) {
	        return new AlterConfigsRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Version = source["Version"];
	        this.Resources = this.convertValues(source["Resources"], AlterConfigsRequestResource);
	        this.ValidateOnly = source["ValidateOnly"];
	        this.UnknownTags = this.convertValues(source["UnknownTags"], null);
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
	
	
	export class AlterConfigsResponseResource {
	    ErrorCode: number;
	    ErrorMessage?: string;
	    ResourceType: ConfigResourceType;
	    ResourceName: string;
	    // Go type: Tags
	    UnknownTags: any;
	
	    static createFrom(source: any = {}) {
	        return new AlterConfigsResponseResource(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ErrorCode = source["ErrorCode"];
	        this.ErrorMessage = source["ErrorMessage"];
	        this.ResourceType = source["ResourceType"];
	        this.ResourceName = source["ResourceName"];
	        this.UnknownTags = this.convertValues(source["UnknownTags"], null);
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
	export class AlterConfigsResponse {
	    Version: number;
	    ThrottleMillis: number;
	    Resources: AlterConfigsResponseResource[];
	    // Go type: Tags
	    UnknownTags: any;
	
	    static createFrom(source: any = {}) {
	        return new AlterConfigsResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Version = source["Version"];
	        this.ThrottleMillis = source["ThrottleMillis"];
	        this.Resources = this.convertValues(source["Resources"], AlterConfigsResponseResource);
	        this.UnknownTags = this.convertValues(source["UnknownTags"], null);
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
	
	export class AlterPartitionAssignmentsRequestTopicPartition {
	    Partition: number;
	    Replicas: number[];
	    // Go type: Tags
	    UnknownTags: any;
	
	    static createFrom(source: any = {}) {
	        return new AlterPartitionAssignmentsRequestTopicPartition(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Partition = source["Partition"];
	        this.Replicas = source["Replicas"];
	        this.UnknownTags = this.convertValues(source["UnknownTags"], null);
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
	export class AlterPartitionAssignmentsRequestTopic {
	    Topic: string;
	    Partitions: AlterPartitionAssignmentsRequestTopicPartition[];
	    // Go type: Tags
	    UnknownTags: any;
	
	    static createFrom(source: any = {}) {
	        return new AlterPartitionAssignmentsRequestTopic(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Topic = source["Topic"];
	        this.Partitions = this.convertValues(source["Partitions"], AlterPartitionAssignmentsRequestTopicPartition);
	        this.UnknownTags = this.convertValues(source["UnknownTags"], null);
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
	
	export class CreateACLsRequestCreation {
	    ResourceType: ACLResourceType;
	    ResourceName: string;
	    ResourcePatternType: ACLResourcePatternType;
	    Principal: string;
	    Host: string;
	    Operation: ACLOperation;
	    PermissionType: ACLPermissionType;
	    // Go type: Tags
	    UnknownTags: any;
	
	    static createFrom(source: any = {}) {
	        return new CreateACLsRequestCreation(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ResourceType = source["ResourceType"];
	        this.ResourceName = source["ResourceName"];
	        this.ResourcePatternType = source["ResourcePatternType"];
	        this.Principal = source["Principal"];
	        this.Host = source["Host"];
	        this.Operation = source["Operation"];
	        this.PermissionType = source["PermissionType"];
	        this.UnknownTags = this.convertValues(source["UnknownTags"], null);
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
	export class CreateACLsRequest {
	    Version: number;
	    Creations: CreateACLsRequestCreation[];
	    // Go type: Tags
	    UnknownTags: any;
	
	    static createFrom(source: any = {}) {
	        return new CreateACLsRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Version = source["Version"];
	        this.Creations = this.convertValues(source["Creations"], CreateACLsRequestCreation);
	        this.UnknownTags = this.convertValues(source["UnknownTags"], null);
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
	
	export class CreateACLsResponseResult {
	    ErrorCode: number;
	    ErrorMessage?: string;
	    // Go type: Tags
	    UnknownTags: any;
	
	    static createFrom(source: any = {}) {
	        return new CreateACLsResponseResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ErrorCode = source["ErrorCode"];
	        this.ErrorMessage = source["ErrorMessage"];
	        this.UnknownTags = this.convertValues(source["UnknownTags"], null);
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
	export class CreateACLsResponse {
	    Version: number;
	    ThrottleMillis: number;
	    Results: CreateACLsResponseResult[];
	    // Go type: Tags
	    UnknownTags: any;
	
	    static createFrom(source: any = {}) {
	        return new CreateACLsResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Version = source["Version"];
	        this.ThrottleMillis = source["ThrottleMillis"];
	        this.Results = this.convertValues(source["Results"], CreateACLsResponseResult);
	        this.UnknownTags = this.convertValues(source["UnknownTags"], null);
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
	
	export class CreateTopicsRequestTopicConfig {
	    Name: string;
	    Value?: string;
	    // Go type: Tags
	    UnknownTags: any;
	
	    static createFrom(source: any = {}) {
	        return new CreateTopicsRequestTopicConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Name = source["Name"];
	        this.Value = source["Value"];
	        this.UnknownTags = this.convertValues(source["UnknownTags"], null);
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
	export class CreateTopicsRequestTopicReplicaAssignment {
	    Partition: number;
	    Replicas: number[];
	    // Go type: Tags
	    UnknownTags: any;
	
	    static createFrom(source: any = {}) {
	        return new CreateTopicsRequestTopicReplicaAssignment(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Partition = source["Partition"];
	        this.Replicas = source["Replicas"];
	        this.UnknownTags = this.convertValues(source["UnknownTags"], null);
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
	export class CreateTopicsRequestTopic {
	    Topic: string;
	    NumPartitions: number;
	    ReplicationFactor: number;
	    ReplicaAssignment: CreateTopicsRequestTopicReplicaAssignment[];
	    Configs: CreateTopicsRequestTopicConfig[];
	    // Go type: Tags
	    UnknownTags: any;
	
	    static createFrom(source: any = {}) {
	        return new CreateTopicsRequestTopic(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Topic = source["Topic"];
	        this.NumPartitions = source["NumPartitions"];
	        this.ReplicationFactor = source["ReplicationFactor"];
	        this.ReplicaAssignment = this.convertValues(source["ReplicaAssignment"], CreateTopicsRequestTopicReplicaAssignment);
	        this.Configs = this.convertValues(source["Configs"], CreateTopicsRequestTopicConfig);
	        this.UnknownTags = this.convertValues(source["UnknownTags"], null);
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
	export class CreateTopicsRequest {
	    Version: number;
	    Topics: CreateTopicsRequestTopic[];
	    TimeoutMillis: number;
	    ValidateOnly: boolean;
	    // Go type: Tags
	    UnknownTags: any;
	
	    static createFrom(source: any = {}) {
	        return new CreateTopicsRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Version = source["Version"];
	        this.Topics = this.convertValues(source["Topics"], CreateTopicsRequestTopic);
	        this.TimeoutMillis = source["TimeoutMillis"];
	        this.ValidateOnly = source["ValidateOnly"];
	        this.UnknownTags = this.convertValues(source["UnknownTags"], null);
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
	
	
	
	export class CreateTopicsResponseTopicConfig {
	    Name: string;
	    Value?: string;
	    ReadOnly: boolean;
	    Source: number;
	    IsSensitive: boolean;
	    // Go type: Tags
	    UnknownTags: any;
	
	    static createFrom(source: any = {}) {
	        return new CreateTopicsResponseTopicConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Name = source["Name"];
	        this.Value = source["Value"];
	        this.ReadOnly = source["ReadOnly"];
	        this.Source = source["Source"];
	        this.IsSensitive = source["IsSensitive"];
	        this.UnknownTags = this.convertValues(source["UnknownTags"], null);
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
	export class CreateTopicsResponseTopic {
	    Topic: string;
	    TopicID: number[];
	    ErrorCode: number;
	    ErrorMessage?: string;
	    ConfigErrorCode: number;
	    NumPartitions: number;
	    ReplicationFactor: number;
	    Configs: CreateTopicsResponseTopicConfig[];
	    // Go type: Tags
	    UnknownTags: any;
	
	    static createFrom(source: any = {}) {
	        return new CreateTopicsResponseTopic(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Topic = source["Topic"];
	        this.TopicID = source["TopicID"];
	        this.ErrorCode = source["ErrorCode"];
	        this.ErrorMessage = source["ErrorMessage"];
	        this.ConfigErrorCode = source["ConfigErrorCode"];
	        this.NumPartitions = source["NumPartitions"];
	        this.ReplicationFactor = source["ReplicationFactor"];
	        this.Configs = this.convertValues(source["Configs"], CreateTopicsResponseTopicConfig);
	        this.UnknownTags = this.convertValues(source["UnknownTags"], null);
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
	export class CreateTopicsResponse {
	    Version: number;
	    ThrottleMillis: number;
	    Topics: CreateTopicsResponseTopic[];
	    // Go type: Tags
	    UnknownTags: any;
	
	    static createFrom(source: any = {}) {
	        return new CreateTopicsResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Version = source["Version"];
	        this.ThrottleMillis = source["ThrottleMillis"];
	        this.Topics = this.convertValues(source["Topics"], CreateTopicsResponseTopic);
	        this.UnknownTags = this.convertValues(source["UnknownTags"], null);
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
	
	
	export class DeleteACLsRequestFilter {
	    ResourceType: ACLResourceType;
	    ResourceName?: string;
	    ResourcePatternType: ACLResourcePatternType;
	    Principal?: string;
	    Host?: string;
	    Operation: ACLOperation;
	    PermissionType: ACLPermissionType;
	    // Go type: Tags
	    UnknownTags: any;
	
	    static createFrom(source: any = {}) {
	        return new DeleteACLsRequestFilter(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ResourceType = source["ResourceType"];
	        this.ResourceName = source["ResourceName"];
	        this.ResourcePatternType = source["ResourcePatternType"];
	        this.Principal = source["Principal"];
	        this.Host = source["Host"];
	        this.Operation = source["Operation"];
	        this.PermissionType = source["PermissionType"];
	        this.UnknownTags = this.convertValues(source["UnknownTags"], null);
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
	export class DeleteACLsRequest {
	    Version: number;
	    Filters: DeleteACLsRequestFilter[];
	    // Go type: Tags
	    UnknownTags: any;
	
	    static createFrom(source: any = {}) {
	        return new DeleteACLsRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Version = source["Version"];
	        this.Filters = this.convertValues(source["Filters"], DeleteACLsRequestFilter);
	        this.UnknownTags = this.convertValues(source["UnknownTags"], null);
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
	
	export class DeleteACLsResponseResultMatchingACL {
	    ErrorCode: number;
	    ErrorMessage?: string;
	    ResourceType: ACLResourceType;
	    ResourceName: string;
	    ResourcePatternType: ACLResourcePatternType;
	    Principal: string;
	    Host: string;
	    Operation: ACLOperation;
	    PermissionType: ACLPermissionType;
	    // Go type: Tags
	    UnknownTags: any;
	
	    static createFrom(source: any = {}) {
	        return new DeleteACLsResponseResultMatchingACL(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ErrorCode = source["ErrorCode"];
	        this.ErrorMessage = source["ErrorMessage"];
	        this.ResourceType = source["ResourceType"];
	        this.ResourceName = source["ResourceName"];
	        this.ResourcePatternType = source["ResourcePatternType"];
	        this.Principal = source["Principal"];
	        this.Host = source["Host"];
	        this.Operation = source["Operation"];
	        this.PermissionType = source["PermissionType"];
	        this.UnknownTags = this.convertValues(source["UnknownTags"], null);
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
	export class DeleteACLsResponseResult {
	    ErrorCode: number;
	    ErrorMessage?: string;
	    MatchingACLs: DeleteACLsResponseResultMatchingACL[];
	    // Go type: Tags
	    UnknownTags: any;
	
	    static createFrom(source: any = {}) {
	        return new DeleteACLsResponseResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ErrorCode = source["ErrorCode"];
	        this.ErrorMessage = source["ErrorMessage"];
	        this.MatchingACLs = this.convertValues(source["MatchingACLs"], DeleteACLsResponseResultMatchingACL);
	        this.UnknownTags = this.convertValues(source["UnknownTags"], null);
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
	export class DeleteACLsResponse {
	    Version: number;
	    ThrottleMillis: number;
	    Results: DeleteACLsResponseResult[];
	    // Go type: Tags
	    UnknownTags: any;
	
	    static createFrom(source: any = {}) {
	        return new DeleteACLsResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Version = source["Version"];
	        this.ThrottleMillis = source["ThrottleMillis"];
	        this.Results = this.convertValues(source["Results"], DeleteACLsResponseResult);
	        this.UnknownTags = this.convertValues(source["UnknownTags"], null);
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
	
	
	export class DeleteRecordsRequestTopicPartition {
	    Partition: number;
	    Offset: number;
	    // Go type: Tags
	    UnknownTags: any;
	
	    static createFrom(source: any = {}) {
	        return new DeleteRecordsRequestTopicPartition(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Partition = source["Partition"];
	        this.Offset = source["Offset"];
	        this.UnknownTags = this.convertValues(source["UnknownTags"], null);
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
	export class DeleteRecordsRequestTopic {
	    Topic: string;
	    Partitions: DeleteRecordsRequestTopicPartition[];
	    // Go type: Tags
	    UnknownTags: any;
	
	    static createFrom(source: any = {}) {
	        return new DeleteRecordsRequestTopic(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Topic = source["Topic"];
	        this.Partitions = this.convertValues(source["Partitions"], DeleteRecordsRequestTopicPartition);
	        this.UnknownTags = this.convertValues(source["UnknownTags"], null);
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
	
	export class DeleteTopicsRequestTopic {
	    Topic?: string;
	    TopicID: number[];
	    // Go type: Tags
	    UnknownTags: any;
	
	    static createFrom(source: any = {}) {
	        return new DeleteTopicsRequestTopic(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Topic = source["Topic"];
	        this.TopicID = source["TopicID"];
	        this.UnknownTags = this.convertValues(source["UnknownTags"], null);
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
	export class DeleteTopicsRequest {
	    Version: number;
	    TopicNames: string[];
	    Topics: DeleteTopicsRequestTopic[];
	    TimeoutMillis: number;
	    // Go type: Tags
	    UnknownTags: any;
	
	    static createFrom(source: any = {}) {
	        return new DeleteTopicsRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Version = source["Version"];
	        this.TopicNames = source["TopicNames"];
	        this.Topics = this.convertValues(source["Topics"], DeleteTopicsRequestTopic);
	        this.TimeoutMillis = source["TimeoutMillis"];
	        this.UnknownTags = this.convertValues(source["UnknownTags"], null);
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
	
	export class DeleteTopicsResponseTopic {
	    Topic?: string;
	    TopicID: number[];
	    ErrorCode: number;
	    ErrorMessage?: string;
	    // Go type: Tags
	    UnknownTags: any;
	
	    static createFrom(source: any = {}) {
	        return new DeleteTopicsResponseTopic(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Topic = source["Topic"];
	        this.TopicID = source["TopicID"];
	        this.ErrorCode = source["ErrorCode"];
	        this.ErrorMessage = source["ErrorMessage"];
	        this.UnknownTags = this.convertValues(source["UnknownTags"], null);
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
	export class DeleteTopicsResponse {
	    Version: number;
	    ThrottleMillis: number;
	    Topics: DeleteTopicsResponseTopic[];
	    // Go type: Tags
	    UnknownTags: any;
	
	    static createFrom(source: any = {}) {
	        return new DeleteTopicsResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Version = source["Version"];
	        this.ThrottleMillis = source["ThrottleMillis"];
	        this.Topics = this.convertValues(source["Topics"], DeleteTopicsResponseTopic);
	        this.UnknownTags = this.convertValues(source["UnknownTags"], null);
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
	
	export class DescribeACLsRequest {
	    Version: number;
	    ResourceType: ACLResourceType;
	    ResourceName?: string;
	    ResourcePatternType: ACLResourcePatternType;
	    Principal?: string;
	    Host?: string;
	    Operation: ACLOperation;
	    PermissionType: ACLPermissionType;
	    // Go type: Tags
	    UnknownTags: any;
	
	    static createFrom(source: any = {}) {
	        return new DescribeACLsRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Version = source["Version"];
	        this.ResourceType = source["ResourceType"];
	        this.ResourceName = source["ResourceName"];
	        this.ResourcePatternType = source["ResourcePatternType"];
	        this.Principal = source["Principal"];
	        this.Host = source["Host"];
	        this.Operation = source["Operation"];
	        this.PermissionType = source["PermissionType"];
	        this.UnknownTags = this.convertValues(source["UnknownTags"], null);
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
	export class DescribeACLsResponseResourceACL {
	    Principal: string;
	    Host: string;
	    Operation: ACLOperation;
	    PermissionType: ACLPermissionType;
	    // Go type: Tags
	    UnknownTags: any;
	
	    static createFrom(source: any = {}) {
	        return new DescribeACLsResponseResourceACL(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Principal = source["Principal"];
	        this.Host = source["Host"];
	        this.Operation = source["Operation"];
	        this.PermissionType = source["PermissionType"];
	        this.UnknownTags = this.convertValues(source["UnknownTags"], null);
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
	export class DescribeACLsResponseResource {
	    ResourceType: ACLResourceType;
	    ResourceName: string;
	    ResourcePatternType: ACLResourcePatternType;
	    ACLs: DescribeACLsResponseResourceACL[];
	    // Go type: Tags
	    UnknownTags: any;
	
	    static createFrom(source: any = {}) {
	        return new DescribeACLsResponseResource(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ResourceType = source["ResourceType"];
	        this.ResourceName = source["ResourceName"];
	        this.ResourcePatternType = source["ResourcePatternType"];
	        this.ACLs = this.convertValues(source["ACLs"], DescribeACLsResponseResourceACL);
	        this.UnknownTags = this.convertValues(source["UnknownTags"], null);
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
	export class DescribeACLsResponse {
	    Version: number;
	    ThrottleMillis: number;
	    ErrorCode: number;
	    ErrorMessage?: string;
	    Resources: DescribeACLsResponseResource[];
	    // Go type: Tags
	    UnknownTags: any;
	
	    static createFrom(source: any = {}) {
	        return new DescribeACLsResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Version = source["Version"];
	        this.ThrottleMillis = source["ThrottleMillis"];
	        this.ErrorCode = source["ErrorCode"];
	        this.ErrorMessage = source["ErrorMessage"];
	        this.Resources = this.convertValues(source["Resources"], DescribeACLsResponseResource);
	        this.UnknownTags = this.convertValues(source["UnknownTags"], null);
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
	
	
	export class DescribeConfigsRequestResource {
	    ResourceType: ConfigResourceType;
	    ResourceName: string;
	    ConfigNames: string[];
	    // Go type: Tags
	    UnknownTags: any;
	
	    static createFrom(source: any = {}) {
	        return new DescribeConfigsRequestResource(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ResourceType = source["ResourceType"];
	        this.ResourceName = source["ResourceName"];
	        this.ConfigNames = source["ConfigNames"];
	        this.UnknownTags = this.convertValues(source["UnknownTags"], null);
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
	export class DescribeConfigsRequest {
	    Version: number;
	    Resources: DescribeConfigsRequestResource[];
	    IncludeSynonyms: boolean;
	    IncludeDocumentation: boolean;
	    // Go type: Tags
	    UnknownTags: any;
	
	    static createFrom(source: any = {}) {
	        return new DescribeConfigsRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Version = source["Version"];
	        this.Resources = this.convertValues(source["Resources"], DescribeConfigsRequestResource);
	        this.IncludeSynonyms = source["IncludeSynonyms"];
	        this.IncludeDocumentation = source["IncludeDocumentation"];
	        this.UnknownTags = this.convertValues(source["UnknownTags"], null);
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
	
	export class DescribeConfigsResponseResourceConfigConfigSynonym {
	    Name: string;
	    Value?: string;
	    Source: number;
	    // Go type: Tags
	    UnknownTags: any;
	
	    static createFrom(source: any = {}) {
	        return new DescribeConfigsResponseResourceConfigConfigSynonym(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Name = source["Name"];
	        this.Value = source["Value"];
	        this.Source = source["Source"];
	        this.UnknownTags = this.convertValues(source["UnknownTags"], null);
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
	export class DescribeConfigsResponseResourceConfig {
	    Name: string;
	    Value?: string;
	    ReadOnly: boolean;
	    IsDefault: boolean;
	    Source: number;
	    IsSensitive: boolean;
	    ConfigSynonyms: DescribeConfigsResponseResourceConfigConfigSynonym[];
	    ConfigType: ConfigType;
	    Documentation?: string;
	    // Go type: Tags
	    UnknownTags: any;
	
	    static createFrom(source: any = {}) {
	        return new DescribeConfigsResponseResourceConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Name = source["Name"];
	        this.Value = source["Value"];
	        this.ReadOnly = source["ReadOnly"];
	        this.IsDefault = source["IsDefault"];
	        this.Source = source["Source"];
	        this.IsSensitive = source["IsSensitive"];
	        this.ConfigSynonyms = this.convertValues(source["ConfigSynonyms"], DescribeConfigsResponseResourceConfigConfigSynonym);
	        this.ConfigType = source["ConfigType"];
	        this.Documentation = source["Documentation"];
	        this.UnknownTags = this.convertValues(source["UnknownTags"], null);
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
	export class DescribeConfigsResponseResource {
	    ErrorCode: number;
	    ErrorMessage?: string;
	    ResourceType: ConfigResourceType;
	    ResourceName: string;
	    Configs: DescribeConfigsResponseResourceConfig[];
	    // Go type: Tags
	    UnknownTags: any;
	
	    static createFrom(source: any = {}) {
	        return new DescribeConfigsResponseResource(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ErrorCode = source["ErrorCode"];
	        this.ErrorMessage = source["ErrorMessage"];
	        this.ResourceType = source["ResourceType"];
	        this.ResourceName = source["ResourceName"];
	        this.Configs = this.convertValues(source["Configs"], DescribeConfigsResponseResourceConfig);
	        this.UnknownTags = this.convertValues(source["UnknownTags"], null);
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
	export class DescribeConfigsResponse {
	    Version: number;
	    ThrottleMillis: number;
	    Resources: DescribeConfigsResponseResource[];
	    // Go type: Tags
	    UnknownTags: any;
	
	    static createFrom(source: any = {}) {
	        return new DescribeConfigsResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Version = source["Version"];
	        this.ThrottleMillis = source["ThrottleMillis"];
	        this.Resources = this.convertValues(source["Resources"], DescribeConfigsResponseResource);
	        this.UnknownTags = this.convertValues(source["UnknownTags"], null);
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
	
	
	
	export class IncrementalAlterConfigsRequestResourceConfig {
	    Name: string;
	    Op: IncrementalAlterConfigOp;
	    Value?: string;
	    // Go type: Tags
	    UnknownTags: any;
	
	    static createFrom(source: any = {}) {
	        return new IncrementalAlterConfigsRequestResourceConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Name = source["Name"];
	        this.Op = source["Op"];
	        this.Value = source["Value"];
	        this.UnknownTags = this.convertValues(source["UnknownTags"], null);
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
	export class IncrementalAlterConfigsRequestResource {
	    ResourceType: ConfigResourceType;
	    ResourceName: string;
	    Configs: IncrementalAlterConfigsRequestResourceConfig[];
	    // Go type: Tags
	    UnknownTags: any;
	
	    static createFrom(source: any = {}) {
	        return new IncrementalAlterConfigsRequestResource(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ResourceType = source["ResourceType"];
	        this.ResourceName = source["ResourceName"];
	        this.Configs = this.convertValues(source["Configs"], IncrementalAlterConfigsRequestResourceConfig);
	        this.UnknownTags = this.convertValues(source["UnknownTags"], null);
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
	export class IncrementalAlterConfigsRequest {
	    Version: number;
	    Resources: IncrementalAlterConfigsRequestResource[];
	    ValidateOnly: boolean;
	    // Go type: Tags
	    UnknownTags: any;
	
	    static createFrom(source: any = {}) {
	        return new IncrementalAlterConfigsRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Version = source["Version"];
	        this.Resources = this.convertValues(source["Resources"], IncrementalAlterConfigsRequestResource);
	        this.ValidateOnly = source["ValidateOnly"];
	        this.UnknownTags = this.convertValues(source["UnknownTags"], null);
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
	
	
	export class IncrementalAlterConfigsResponseResource {
	    ErrorCode: number;
	    ErrorMessage?: string;
	    ResourceType: ConfigResourceType;
	    ResourceName: string;
	    // Go type: Tags
	    UnknownTags: any;
	
	    static createFrom(source: any = {}) {
	        return new IncrementalAlterConfigsResponseResource(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ErrorCode = source["ErrorCode"];
	        this.ErrorMessage = source["ErrorMessage"];
	        this.ResourceType = source["ResourceType"];
	        this.ResourceName = source["ResourceName"];
	        this.UnknownTags = this.convertValues(source["UnknownTags"], null);
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
	export class IncrementalAlterConfigsResponse {
	    Version: number;
	    ThrottleMillis: number;
	    Resources: IncrementalAlterConfigsResponseResource[];
	    // Go type: Tags
	    UnknownTags: any;
	
	    static createFrom(source: any = {}) {
	        return new IncrementalAlterConfigsResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Version = source["Version"];
	        this.ThrottleMillis = source["ThrottleMillis"];
	        this.Resources = this.convertValues(source["Resources"], IncrementalAlterConfigsResponseResource);
	        this.UnknownTags = this.convertValues(source["UnknownTags"], null);
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
	
	export class MetadataRequestTopic {
	    TopicID: number[];
	    Topic?: string;
	    // Go type: Tags
	    UnknownTags: any;
	
	    static createFrom(source: any = {}) {
	        return new MetadataRequestTopic(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.TopicID = source["TopicID"];
	        this.Topic = source["Topic"];
	        this.UnknownTags = this.convertValues(source["UnknownTags"], null);
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
	export class MetadataRequest {
	    Version: number;
	    Topics: MetadataRequestTopic[];
	    AllowAutoTopicCreation: boolean;
	    IncludeClusterAuthorizedOperations: boolean;
	    IncludeTopicAuthorizedOperations: boolean;
	    // Go type: Tags
	    UnknownTags: any;
	
	    static createFrom(source: any = {}) {
	        return new MetadataRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Version = source["Version"];
	        this.Topics = this.convertValues(source["Topics"], MetadataRequestTopic);
	        this.AllowAutoTopicCreation = source["AllowAutoTopicCreation"];
	        this.IncludeClusterAuthorizedOperations = source["IncludeClusterAuthorizedOperations"];
	        this.IncludeTopicAuthorizedOperations = source["IncludeTopicAuthorizedOperations"];
	        this.UnknownTags = this.convertValues(source["UnknownTags"], null);
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
	
	export class MetadataResponseTopicPartition {
	    ErrorCode: number;
	    Partition: number;
	    Leader: number;
	    LeaderEpoch: number;
	    Replicas: number[];
	    ISR: number[];
	    OfflineReplicas: number[];
	    // Go type: Tags
	    UnknownTags: any;
	
	    static createFrom(source: any = {}) {
	        return new MetadataResponseTopicPartition(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ErrorCode = source["ErrorCode"];
	        this.Partition = source["Partition"];
	        this.Leader = source["Leader"];
	        this.LeaderEpoch = source["LeaderEpoch"];
	        this.Replicas = source["Replicas"];
	        this.ISR = source["ISR"];
	        this.OfflineReplicas = source["OfflineReplicas"];
	        this.UnknownTags = this.convertValues(source["UnknownTags"], null);
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
	export class MetadataResponseTopic {
	    ErrorCode: number;
	    Topic?: string;
	    TopicID: number[];
	    IsInternal: boolean;
	    Partitions: MetadataResponseTopicPartition[];
	    AuthorizedOperations: number;
	    // Go type: Tags
	    UnknownTags: any;
	
	    static createFrom(source: any = {}) {
	        return new MetadataResponseTopic(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ErrorCode = source["ErrorCode"];
	        this.Topic = source["Topic"];
	        this.TopicID = source["TopicID"];
	        this.IsInternal = source["IsInternal"];
	        this.Partitions = this.convertValues(source["Partitions"], MetadataResponseTopicPartition);
	        this.AuthorizedOperations = source["AuthorizedOperations"];
	        this.UnknownTags = this.convertValues(source["UnknownTags"], null);
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
	export class MetadataResponseBroker {
	    NodeID: number;
	    Host: string;
	    Port: number;
	    Rack?: string;
	    // Go type: Tags
	    UnknownTags: any;
	
	    static createFrom(source: any = {}) {
	        return new MetadataResponseBroker(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.NodeID = source["NodeID"];
	        this.Host = source["Host"];
	        this.Port = source["Port"];
	        this.Rack = source["Rack"];
	        this.UnknownTags = this.convertValues(source["UnknownTags"], null);
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
	export class MetadataResponse {
	    Version: number;
	    ThrottleMillis: number;
	    Brokers: MetadataResponseBroker[];
	    ClusterID?: string;
	    ControllerID: number;
	    Topics: MetadataResponseTopic[];
	    AuthorizedOperations: number;
	    // Go type: Tags
	    UnknownTags: any;
	
	    static createFrom(source: any = {}) {
	        return new MetadataResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Version = source["Version"];
	        this.ThrottleMillis = source["ThrottleMillis"];
	        this.Brokers = this.convertValues(source["Brokers"], MetadataResponseBroker);
	        this.ClusterID = source["ClusterID"];
	        this.ControllerID = source["ControllerID"];
	        this.Topics = this.convertValues(source["Topics"], MetadataResponseTopic);
	        this.AuthorizedOperations = source["AuthorizedOperations"];
	        this.UnknownTags = this.convertValues(source["UnknownTags"], null);
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
	
	
	
	export class OffsetCommitRequestTopicPartition {
	    Partition: number;
	    Offset: number;
	    Timestamp: number;
	    LeaderEpoch: number;
	    Metadata?: string;
	    // Go type: Tags
	    UnknownTags: any;
	
	    static createFrom(source: any = {}) {
	        return new OffsetCommitRequestTopicPartition(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Partition = source["Partition"];
	        this.Offset = source["Offset"];
	        this.Timestamp = source["Timestamp"];
	        this.LeaderEpoch = source["LeaderEpoch"];
	        this.Metadata = source["Metadata"];
	        this.UnknownTags = this.convertValues(source["UnknownTags"], null);
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
	export class OffsetCommitRequestTopic {
	    Topic: string;
	    Partitions: OffsetCommitRequestTopicPartition[];
	    // Go type: Tags
	    UnknownTags: any;
	
	    static createFrom(source: any = {}) {
	        return new OffsetCommitRequestTopic(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Topic = source["Topic"];
	        this.Partitions = this.convertValues(source["Partitions"], OffsetCommitRequestTopicPartition);
	        this.UnknownTags = this.convertValues(source["UnknownTags"], null);
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
	
	export class OffsetDeleteRequestTopicPartition {
	    Partition: number;
	
	    static createFrom(source: any = {}) {
	        return new OffsetDeleteRequestTopicPartition(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Partition = source["Partition"];
	    }
	}
	export class OffsetDeleteRequestTopic {
	    Topic: string;
	    Partitions: OffsetDeleteRequestTopicPartition[];
	
	    static createFrom(source: any = {}) {
	        return new OffsetDeleteRequestTopic(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Topic = source["Topic"];
	        this.Partitions = this.convertValues(source["Partitions"], OffsetDeleteRequestTopicPartition);
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

export namespace serde {
	
	export class TroubleshootingReport {
	    serdeName: string;
	    message: string;
	
	    static createFrom(source: any = {}) {
	        return new TroubleshootingReport(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.serdeName = source["serdeName"];
	        this.message = source["message"];
	    }
	}
	export class RecordPayload {
	    originalPayload?: number[];
	    payloadSizeBytes: number;
	    humanReadablePayload: number[];
	    isPayloadTooLarge: boolean;
	    isPayloadNull: boolean;
	    encoding: string;
	    schemaId?: number;
	    troubleshooting?: TroubleshootingReport[];
	    extraMetadata?: Record<string, string>;
	
	    static createFrom(source: any = {}) {
	        return new RecordPayload(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.originalPayload = source["originalPayload"];
	        this.payloadSizeBytes = source["payloadSizeBytes"];
	        this.humanReadablePayload = source["humanReadablePayload"];
	        this.isPayloadTooLarge = source["isPayloadTooLarge"];
	        this.isPayloadNull = source["isPayloadNull"];
	        this.encoding = source["encoding"];
	        this.schemaId = source["schemaId"];
	        this.troubleshooting = this.convertValues(source["troubleshooting"], TroubleshootingReport);
	        this.extraMetadata = source["extraMetadata"];
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
	export class RecordPayloadInput {
	    Payload: any;
	    Encoding: string;
	    Options: any[];
	
	    static createFrom(source: any = {}) {
	        return new RecordPayloadInput(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Payload = source["Payload"];
	        this.Encoding = source["Encoding"];
	        this.Options = source["Options"];
	    }
	}

}

export namespace sr {
	
	export enum SchemaType {
	    AVRO = 0,
	    PROTOBUF = 1,
	    JSON = 2,
	}
	export enum CompatibilityLevel {
	    NONE = 1,
	    BACKWROD = 2,
	    BACKWORD_TRANSITIVE = 3,
	    FORWARD = 4,
	    FORWARD_TRANSITIVE = 5,
	    FULL = 6,
	    FULL_TRANSITIVE = 7,
	}
	export enum Mode {
	    IMPORT = 0,
	    READONLY = 1,
	    READWRITE = 2,
	}
	export enum SchemaRuleMode {
	    UPGRADE = 0,
	    DOWNGRADE = 1,
	    UPDOWN = 2,
	    WRITE = 3,
	    READ = 4,
	    WRITEREAD = 5,
	}
	export enum SchemaRuleKind {
	    TRANSFORM = 0,
	    CONDITION = 1,
	}
	export class SchemaRule {
	    name: string;
	    doc?: string;
	    kind: SchemaRuleKind;
	    mode: SchemaRuleMode;
	    type: string;
	    tags: string[];
	    params?: Record<string, string>;
	    expr: string;
	    onSuccess?: string;
	    onFailure?: string;
	    disabled?: boolean;
	
	    static createFrom(source: any = {}) {
	        return new SchemaRule(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.doc = source["doc"];
	        this.kind = source["kind"];
	        this.mode = source["mode"];
	        this.type = source["type"];
	        this.tags = source["tags"];
	        this.params = source["params"];
	        this.expr = source["expr"];
	        this.onSuccess = source["onSuccess"];
	        this.onFailure = source["onFailure"];
	        this.disabled = source["disabled"];
	    }
	}
	export class SchemaRuleSet {
	    migrationRules?: SchemaRule[];
	    domainRules?: SchemaRule[];
	
	    static createFrom(source: any = {}) {
	        return new SchemaRuleSet(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.migrationRules = this.convertValues(source["migrationRules"], SchemaRule);
	        this.domainRules = this.convertValues(source["domainRules"], SchemaRule);
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
	export class SchemaMetadata {
	    tags?: Record<string, Array<string>>;
	    properties?: Record<string, string>;
	    sensitive?: string[];
	
	    static createFrom(source: any = {}) {
	        return new SchemaMetadata(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.tags = source["tags"];
	        this.properties = source["properties"];
	        this.sensitive = source["sensitive"];
	    }
	}
	export class SchemaReference {
	    name: string;
	    subject: string;
	    version: number;
	
	    static createFrom(source: any = {}) {
	        return new SchemaReference(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.subject = source["subject"];
	        this.version = source["version"];
	    }
	}
	export class Schema {
	    schema: string;
	    schemaType?: SchemaType;
	    references?: SchemaReference[];
	    metadata?: SchemaMetadata;
	    ruleSet?: SchemaRuleSet;
	
	    static createFrom(source: any = {}) {
	        return new Schema(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.schema = source["schema"];
	        this.schemaType = source["schemaType"];
	        this.references = this.convertValues(source["references"], SchemaReference);
	        this.metadata = this.convertValues(source["metadata"], SchemaMetadata);
	        this.ruleSet = this.convertValues(source["ruleSet"], SchemaRuleSet);
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
	
	
	
	
	export class SetCompatibility {
	    compatibility: CompatibilityLevel;
	    alias?: string;
	    normalize?: boolean;
	    compatibilityGroup?: string;
	    defaultMetadata?: SchemaMetadata;
	    overrideMetadata?: SchemaMetadata;
	    defaultRuleSet?: SchemaRuleSet;
	    overrideRuleSet?: SchemaRuleSet;
	
	    static createFrom(source: any = {}) {
	        return new SetCompatibility(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.compatibility = source["compatibility"];
	        this.alias = source["alias"];
	        this.normalize = source["normalize"];
	        this.compatibilityGroup = source["compatibilityGroup"];
	        this.defaultMetadata = this.convertValues(source["defaultMetadata"], SchemaMetadata);
	        this.overrideMetadata = this.convertValues(source["overrideMetadata"], SchemaMetadata);
	        this.defaultRuleSet = this.convertValues(source["defaultRuleSet"], SchemaRuleSet);
	        this.overrideRuleSet = this.convertValues(source["overrideRuleSet"], SchemaRuleSet);
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

