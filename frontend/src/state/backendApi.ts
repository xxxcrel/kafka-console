/**
 * Copyright 2022 Redpanda Data, Inc.
 *
 * Use of this software is governed by the Business Source License
 * included in the file https://github.com/redpanda-data/redpanda/blob/dev/licenses/bsl.md
 *
 * As of the Change Date specified in that file, in accordance with
 * the Business Source License, use of this software will be governed
 * by the Apache License, Version 2.0
 */

/*eslint block-scoped-var: "error"*/

import {comparer, computed, observable, transaction} from 'mobx';
import {config as appConfig} from '../config';
import {decodeBase64} from '../utils/utils';
import {
  AlterPartitionAssignments,
  CreateSchemaRegistrySchema,
  CreateTopic,
  DeleteConsumerGroup,
  DeleteConsumerGroupOffsets,
  DeleteSchemaRegistrySubject,
  DeleteSchemaRegistrySubjectConfig,
  DeleteSchemaRegistrySubjectVersion,
  DeleteTopic,
  DeleteTopicRecords,
  DescribeQuotas,
  EditConsumerGroupOffsets,
  EditTopicConfig,
  GetBrokerConfig,
  GetBrokersWithLogDirs,
  GetClusterInfo,
  GetConsoleInfo,
  GetConsumerGroupsOverview,
  GetEndpointCompatibility,
  GetKafkaAuthorizerInfo,
  GetKafkaInfo,
  GetSchemaRegistryConfig,
  GetSchemaRegistryInfo,
  GetSchemaRegistryMode,
  GetSchemaRegistrySchemaReferencedBy,
  GetSchemaRegistrySchemaTypes,
  GetSchemaRegistrySubjectDetails,
  GetSchemaRegistrySubjects,
  GetSchemaUsagesByID,
  GetTopicConfigs,
  GetTopicDetails,
  GetTopicDocumentation,
  GetTopicsConfigs,
  GetTopicsOverview,
  IncrementalAlterConfigs,
  ListOffsets,
  ListPartitionReassignments,
  ListTopicConsumers,
  ProducePlainRecords,
  PutSchemaRegistryConfig,
  ValidateSchemaRegistrySchema
} from "../../wailsjs/go/main/App";
import {clusterstatus, kconsole, kgo, kmsg, sr} from "../../wailsjs/go/models";
import BrokerConfigEntry = kconsole.BrokerConfigEntry;
import ClusterInfo = kconsole.ClusterInfo;
import BrokerWithLogDirs = kconsole.BrokerWithLogDirs;
import TopicSummary = kconsole.TopicSummary;
import TopicPartitionDetails = kconsole.TopicPartitionDetails;
import ConsumerGroupOverview = kconsole.ConsumerGroupOverview;
import DeleteRecordsRequestTopic = kmsg.DeleteRecordsRequestTopic;
import TopicConfig = kconsole.TopicConfig;
import TopicDocumentation = kconsole.TopicDocumentation;
import TopicConsumerGroup = kconsole.TopicConsumerGroup;
import EndpointCompatibility = kconsole.EndpointCompatibility;
import CompressionCodec = kgo.CompressionCodec;
import ProduceRecordsResponse = kconsole.ProduceRecordsResponse;
import CreateTopicResponse = kconsole.CreateTopicResponse;
import CreateTopicsRequestTopic = kmsg.CreateTopicsRequestTopic;
import QuotaResponse = kconsole.QuotaResponse;
import TopicConfigEntry = kconsole.TopicConfigEntry;
import ACLResource = kconsole.ACLResource;
import ACLOverview = kconsole.ACLOverview;
import SchemaRegistrySubject = kconsole.SchemaRegistrySubject;
import SchemaRegistrySubjectDetails = kconsole.SchemaRegistrySubjectDetails;
import SchemaReference = kconsole.SchemaReference;
import SchemaVersion = kconsole.SchemaVersion;
import SchemaRegistryConfig = kconsole.SchemaRegistryConfig;
import SchemaType = sr.SchemaType;
import Schema = sr.Schema;
import SchemaRegistrySchemaValidation = kconsole.SchemaRegistrySchemaValidation;
import CreateSchemaResponse = kconsole.CreateSchemaResponse;
import SchemaRegistryDeleteSubjectResponse = kconsole.SchemaRegistryDeleteSubjectResponse;
import SchemaRegistryDeleteSubjectVersionResponse = kconsole.SchemaRegistryDeleteSubjectVersionResponse;
import AlterPartitionReassignmentsResponse = kconsole.AlterPartitionReassignmentsResponse;
import AlterPartitionAssignmentsRequestTopic = kmsg.AlterPartitionAssignmentsRequestTopic;
import CompatibilityLevel = sr.CompatibilityLevel;
import PartitionReassignments = kconsole.PartitionReassignments;
import SetCompatibility = sr.SetCompatibility;
import IncrementalAlterConfigsRequestResource = kmsg.IncrementalAlterConfigsRequestResource;
import IncrementalAlterConfigsResourceResponse = kconsole.IncrementalAlterConfigsResourceResponse;
import ConfigResourceType = kmsg.ConfigResourceType;
import IncrementalAlterConfigOp = kmsg.IncrementalAlterConfigOp;
import IncrementalAlterConfigsRequestResourceConfig = kmsg.IncrementalAlterConfigsRequestResourceConfig;
import KafkaInfo = clusterstatus.KafkaInfo;
import SchemaRegistryInfo = clusterstatus.SchemaRegistryInfo;
import {PublishRecordsRequest, TopicPermissions} from "./restInterfaces";

export const REST_CACHE_DURATION_SEC = 20;

export interface UserData {
  displayName: string;
  avatarUrl: string;

  canViewConsoleUsers: boolean;
  canListAcls: boolean;
  canListQuotas: boolean;
  canReassignPartitions: boolean;
  canPatchConfigs: boolean;
  canCreateRoles: boolean;
  canManageUsers: boolean;
  canViewPermissionsList: boolean;

  canManageLicense: boolean;
  canViewSchemas: boolean;
  canCreateSchemas: boolean;
  canDeleteSchemas: boolean;
  canManageSchemaRegistry: boolean;
  canViewDebugBundle: boolean;

  canListTransforms: boolean;
  canCreateTransforms: boolean;
  canDeleteTransforms: boolean;
}


export interface ClusterOverview {
  kafkaAuthorizerInfo: number | null;
  kafka: KafkaInfo | null;
  console: string | null;
  schemaRegistry: SchemaRegistryInfo | null;
}
//
// BackendAPI
//
const apiStore = {
  // Data
  endpointCompatibility: null as EndpointCompatibility | null,

  clusterOverview: null as ClusterOverview | null,
  brokers: null as BrokerWithLogDirs[] | null,

  clusters: ['A', 'B', 'C'],
  clusterInfo: null as ClusterInfo | null,

  brokerConfigs: new Map<number, BrokerConfigEntry[] | string>(), // config entries, or error string

  schemaOverviewIsConfigured: undefined as boolean | undefined,

  schemaMode: undefined as string | null | undefined, // undefined = not yet known, null = got not configured response
  schemaCompatibility: undefined as CompatibilityLevel | null | undefined, // undefined = not yet known, null = got not configured response
  schemaSubjects: undefined as SchemaRegistrySubject[] | undefined,
  schemaTypes: undefined as SchemaType[] | undefined,
  schemaDetails: new Map<string, SchemaRegistrySubjectDetails>(), // subjectName => details
  schemaReferencedBy: new Map<string, Map<number, SchemaReference[]>>(), // subjectName => version => details
  schemaUsagesById: new Map<number, SchemaVersion[]>(),

  topics: null as TopicSummary[] | null,
  topicConfig: new Map<string, TopicConfig | null>(), // null = not allowed to view config of this topic
  topicDocumentation: new Map<string, TopicDocumentation>(),
  topicPermissions: new Map<string, TopicPermissions | null>(),
  topicPartitions: new Map<string, TopicPartitionDetails[] | null>(), // null = not allowed to view partitions of this config
  topicPartitionErrors: new Map<string, Array<{ id: number; partitionError: string }>>(),
  topicWatermarksErrors: new Map<string, Array<{ id: number; waterMarksError: string }>>(),
  topicConsumers: new Map<string, TopicConsumerGroup[]>(),
  topicAcls: new Map<string, ACLOverview | null>(),

  Quotas: undefined as QuotaResponse | undefined | null,

  consumerGroups: new Map<string, ConsumerGroupOverview>(),

  partitionReassignments: undefined as PartitionReassignments[] | null | undefined,

  hasDebugProcess: false as boolean,

  // undefined = we haven't checked yet
  // null = call completed, and we're not logged in
  userData: undefined as UserData | null | undefined,

  async logout() {
    await appConfig.fetch('./auth/logout');
    this.userData = null;
  },
  async refreshUserData() {
    this.userData = {
      displayName: 'Kafka Console',
      avatarUrl: '',
      canListAcls: true,
      canListQuotas: true,
      canPatchConfigs: true,
      canReassignPartitions: true,
      canCreateRoles: true,
      canViewPermissionsList: true,
      canManageLicense: true,
      canManageUsers: true,
      canCreateSchemas: true,
      canDeleteSchemas: true,
      canManageSchemaRegistry: true,
      canViewSchemas: true,
      canListTransforms: true,
      canCreateTransforms: true,
      canDeleteTransforms: true,
      canViewDebugBundle: true,
      canViewConsoleUsers: true,
    }
  },

  // Fetch errors
  errors: [] as any[],

  refreshTopics() {
    GetTopicsOverview()
      .then((topics) => {
        this.topics = topics;
      }, addError);
  },

  async refreshTopicConfig(topicName: string): Promise<void> {
    GetTopicConfigs(topicName, [])
      .then((topicConfig) => {
        if (!topicConfig) {
          this.topicConfig.delete(topicName);
          return;
        }
        if (topicConfig.error) {
          this.topicConfig.set(topicName, topicConfig);
          return;
        }
        const topicDescription = topicConfig;
        prepareSynonyms(topicDescription.configEntries);
        this.topicConfig.set(topicName, topicDescription);
      }, addError)
  },

  async getTopicOffsetsByTimestamp(topicNames: string[], timestampUnixMs: number): Promise<kconsole.TopicOffset[]> {
    return ListOffsets(topicNames, timestampUnixMs);
  },

  refreshTopicDocumentation(topicName: string) {
    GetTopicDocumentation(topicName)
      .then((documentation) => {
        documentation.text = documentation.markdown == null ? "" : decodeBase64(Buffer.from(documentation.markdown).toString("utf-8"));
        this.topicDocumentation.set(topicName, documentation);
      }, addError);
  },

  async deleteTopic(topicName: string) {
    return DeleteTopic(topicName)
      .catch(
        addError,
      );
  },

  async deleteTopicRecords(topicName: string, offset: number, partitionId?: number) {
    const partitions =
      partitionId !== undefined
        ? [{Partition: partitionId, Offset: offset}]
        : this.topicPartitions?.get(topicName)?.map((partition) => ({Partition: partition.id, Offset: offset}));

    if (!partitions || partitions.length === 0) {
      addError(new Error(`Topic ${topicName} doesn't have partitions.`));
      return;
    }

    return this.deleteTopicRecordsFromMultiplePartitionOffsetPairs(topicName, partitions);
  },

  async deleteTopicRecordsFromAllPartitionsHighWatermark(topicName: string) {
    const partitions = this.topicPartitions?.get(topicName)?.map(({waterMarkHigh, id}) => ({
      Partition: id,
      Offset: waterMarkHigh,
    }));

    if (!partitions || partitions.length === 0) {
      addError(new Error(`Topic ${topicName} doesn't have partitions.`));
      return;
    }

    return this.deleteTopicRecordsFromMultiplePartitionOffsetPairs(topicName, partitions);
  },

  async deleteTopicRecordsFromMultiplePartitionOffsetPairs(
    topicName: string,
    pairs: Array<{ Partition: number; Offset: number }>,
  ) {
    return DeleteTopicRecords(DeleteRecordsRequestTopic.createFrom({Topic: topicName, Partition: [...pairs]}))
      .catch(addError);
  },

  refreshPartitions(topics: 'all' | string[] = 'all'): Promise<void> {
    return GetTopicDetails(topics === 'all' ? [] : [...topics])
      .then((topics) => {
        if (!topics) return;
        if (topics instanceof Array) {
          transaction(() => {
              const errors: {
                topicName: string;
                partitionErrors: { partitionId: number; error: string }[];
                waterMarkErrors: { partitionId: number; error: string }[];
              }[] = [];

              for (const t of topics) {
                if (t.error != null) {
                  // kconsole.error(`refreshAllTopicPartitions: error for topic ${t.topicName}: ${t.error}`);
                  continue;
                }

                // If any partition has any errors, don't set the result for that topic
                const partitionErrors = [];
                const waterMarkErrors = [];
                for (const p of t.partitions) {
                  // topicName
                  p.topicName = t.topicName;

                  let partitionHasError = false;
                  if (p.partitionError) {
                    partitionErrors.push({partitionId: p.id, error: p.partitionError});
                    partitionHasError = true;
                  }
                  if (p.waterMarksError) {
                    waterMarkErrors.push({partitionId: p.id, error: p.waterMarksError});
                    partitionHasError = true;
                  }
                  if (partitionHasError) {
                    p.hasErrors = true;
                    continue;
                  }

                  // Add some local/cached properties to make working with the data easier
                  const validLogDirs = p.partitionLogDirs.filter((e) => !e.error && e.size >= 0);
                  const replicaSize = validLogDirs.length > 0 ? validLogDirs.max((e) => e.size) : 0;
                  p.replicaSize = replicaSize >= 0 ? replicaSize : 0;
                }

                // Set partition
                this.topicPartitions.set(t.topicName, t.partitions);

                if (partitionErrors.length === 0 && waterMarkErrors.length === 0) {
                } else {
                  errors.push({
                    topicName: t.topicName,
                    partitionErrors: partitionErrors,
                    waterMarkErrors: waterMarkErrors,
                  });
                }
              }

              // if (errors.length > 0)
              //     kconsole.error('refreshAllTopicPartitions: response had errors', errors);
            }
          )
        }
      }, addError);
  },

  refreshPartitionsForTopic(topicName: string) {
    GetTopicDetails([topicName])
      .then((topicDetails) => {
        if (topicDetails instanceof Array) {
          if (topicDetails.length != 1) {
            return;
          }
          let partitions = topicDetails[0].partitions;
          if (partitions) {
            const partitionErrors: Array<{ id: number; partitionError: string }> = [];
            const waterMarksErrors: Array<{ id: number; waterMarksError: string }> = [];

            // Add some local/cached properties to make working with the data easier
            for (const p of partitions) {
              // topicName
              p.topicName = topicName;

              if (p.partitionError) partitionErrors.push({id: p.id, partitionError: p.partitionError});
              if (p.waterMarksError) waterMarksErrors.push({id: p.id, waterMarksError: p.waterMarksError});
              if (partitionErrors.length || waterMarksErrors.length) continue;

              // replicaSize
              const validLogDirs = p.partitionLogDirs.filter((e) => (e.error == null || e.error === '') && e.size >= 0);
              const replicaSize = validLogDirs.length > 0 ? validLogDirs.max((e) => e.size) : 0;
              p.replicaSize = replicaSize >= 0 ? replicaSize : 0;
            }

            if (partitionErrors.length === 0 && waterMarksErrors.length === 0) {
              // Set partitions
              this.topicPartitionErrors.delete(topicName);
              this.topicWatermarksErrors.delete(topicName);
              this.topicPartitions.set(topicName, partitions);
            } else {
              this.topicPartitionErrors.set(topicName, partitionErrors);
              this.topicWatermarksErrors.set(topicName, waterMarksErrors);
              console.error(
                `refreshPartitionsForTopic: response has partition errors (t=${topicName} p=${partitionErrors.length}, w=${waterMarksErrors.length})`,
              );
            }
          } else {
            // Set null to indicate that we're not allowed to see the partitions
            this.topicPartitions.set(topicName, null);
            return;
          }

          let partitionErrors = 0;
          let waterMarkErrors = 0;

          // Add some local/cached properties to make working with the data easier
          for (const p of partitions) {
            // topicName
            p.topicName = topicName;

            if (p.partitionError) partitionErrors++;
            if (p.waterMarksError) waterMarkErrors++;
            if (partitionErrors || waterMarkErrors) {
              continue;
            }

            // replicaSize
            const validLogDirs = p.partitionLogDirs.filter((e) => (e.error == null || e.error === '') && e.size >= 0);
            const replicaSize = validLogDirs.length > 0 ? validLogDirs.max((e) => e.size) : 0;
            p.replicaSize = replicaSize >= 0 ? replicaSize : 0;
          }

          // Set partitions
          this.topicPartitions.set(topicName, partitions);

          if (partitionErrors > 0 || waterMarkErrors > 0)
            console.warn(
              `refreshPartitionsForTopic: response has partition errors (topic=${topicName} partitionErrors=${partitionErrors}, waterMarkErrors=${waterMarkErrors})`,
            )
        }
      }, addError);
  },

  get getTopicPartitionArray() {
    const result: string[] = [];

    this.topicPartitions.forEach((partitions, topicName) => {
      if (partitions !== null) {
        for (const partition of partitions) {
          result.push(`${topicName}/${partition.id}`);
        }
      }
    });

    return result;
  },

  refreshTopicConsumers(topicName: string) {
    ListTopicConsumers(topicName)
      .then((consumerGroups) => this.topicConsumers.set(topicName, consumerGroups), addError);
  },

  refreshQuotas() {
    DescribeQuotas()
      .then(
        (v) => (this.Quotas = v ?? null),
        addError,
      );
  },

  async refreshSupportedEndpoints() {
    const r = await GetEndpointCompatibility();
    if (!r) return null;
    this.endpointCompatibility = r;
  },

  async refreshClusterOverview() {
    const requests: Array<Promise<any>> = [
      GetKafkaAuthorizerInfo().catch((e) => {
        console.error(e);
        return null;
      }),
      GetConsoleInfo().catch((e) => {
        console.error(e);
        return null;
      }),
      GetKafkaInfo().catch((e) => {
        console.error(e);
        return null;
      }),
      GetSchemaRegistryInfo().catch((e) => {
        console.error(e);
        return null;
      })
    ]
    const responses = await Promise.all(requests);
    const [
      aclCnt,
      consoleInfo,
      kafkaInfo,
      schemaRegistryInfo
    ] = responses;

    this.clusterOverview = {
      kafkaAuthorizerInfo: aclCnt,
      console: consoleInfo,
      kafka: kafkaInfo,
      schemaRegistry: schemaRegistryInfo,
    };
  },

  refreshBrokers() {
    GetBrokersWithLogDirs()
      .then((v) => {
        this.brokers = v;
      }, addError);
  },

  refreshCluster() {
    GetClusterInfo().then((clusterInfo) => {
      if (clusterInfo != null) {
        transaction(() => {
          // add 'type' to each synonym entry
          for (const broker of clusterInfo.brokers)
            if (broker.config && !broker.config.error) prepareBrokerSynonyms(broker.config.configs ?? []);

          // don't assign if the value didn't change
          // we'd re-trigger all observers!
          // TODO: it would probably be easier to just annotate 'clusterInfo' with a structural comparer
          if (!comparer.structural(this.clusterInfo, clusterInfo)) this.clusterInfo = clusterInfo;

          for (const b of clusterInfo.brokers)
            if (b.config.error) this.brokerConfigs.set(b.brokerId, b.config.error);
            else this.brokerConfigs.set(b.brokerId, b.config.configs ?? []);
        });
      }
    }, addError);
  },

  refreshBrokerConfig(brokerId: number) {
    GetBrokerConfig(brokerId)
      .then(brokerConfigs => {
        prepareBrokerSynonyms(brokerConfigs);
        this.brokerConfigs.set(brokerId, brokerConfigs);
      })
      .catch((err) => {
        this.brokerConfigs.set(brokerId, String(err));
      });
  },

  refreshConsumerGroup(groupId: string) {
    GetConsumerGroupsOverview([groupId])
      .then((groupOverviews) => {
        addFrontendFieldsForConsumerGroup(groupOverviews[0]);
        this.consumerGroups.set(groupOverviews[0].groupId, groupOverviews[0]);
      }, addError);
  },

  refreshConsumerGroups() {
    GetConsumerGroupsOverview([])
      .then((groupOverviews) => {
        if (groupOverviews != null) {
          for (const g of groupOverviews) addFrontendFieldsForConsumerGroup(g);

          transaction(() => {
            this.consumerGroups.clear();
            for (const g of groupOverviews) this.consumerGroups.set(g.groupId, g);
          });
        }
      }, addError);
  },

  async editConsumerGroupOffsets(
    groupId: string,
    topics: kmsg.OffsetCommitRequestTopic[],
  ): Promise<kconsole.EditConsumerGroupOffsetsResponse> {
    return await EditConsumerGroupOffsets(groupId, topics);
  },

  async deleteConsumerGroupOffsets(
    groupId: string,
    topics: kmsg.OffsetDeleteRequestTopic[],
  ): Promise<Array<kconsole.DeleteConsumerGroupOffsetsResponseTopic>> {
    return DeleteConsumerGroupOffsets(groupId, topics);
  },

  async deleteConsumerGroup(groupId: string): Promise<void> {
    return DeleteConsumerGroup(groupId);
  },

  async refreshSchemaMode() {
    return GetSchemaRegistryMode()
      .then((r) => {
        this.schemaOverviewIsConfigured = true;
        this.schemaMode = r.mode;
      })
      .catch((err) => {
        this.schemaMode = 'Unknown';
        console.warn('failed to request schema mode', err);
      });
  },

  async refreshSchemaCompatibilityConfig() {
    return GetSchemaRegistryConfig("")
      .then((r) => {
        this.schemaOverviewIsConfigured = true;
        this.schemaCompatibility = r.compatibility;
      })
      .catch(addError);
  },

  refreshSchemaSubjects() {
    GetSchemaRegistrySubjects()
      .then(
        (subjects) => {
          // could also be a "not configured" response
          if (Array.isArray(subjects)) {
            this.schemaSubjects = subjects;
          }
        },
        addError,
      );
  },

  refreshSchemaTypes() {
    GetSchemaRegistrySchemaTypes()
      .then((types) => {
        // could also be a "not configured" response
        if (types.schemaTypes) {
          this.schemaTypes = types.schemaTypes;
        }
      })
      .catch((err) => {
        this.schemaTypes = undefined;
        console.warn('failed to request schema type', err);
      });
  },

  refreshSchemaDetails(subjectName: string) {
    // Always refresh all versions, otherwise we cannot know wether or not we have to refresh with 'all,
    // If we refresh with 'latest' or specific number, we'd need to keep track of what information we're missing
    return GetSchemaRegistrySubjectDetails(subjectName, "all")
      .then((details) => {
        this.schemaDetails.set(subjectName, details);
      })
      .catch(addError);
  },

  refreshSchemaReferencedBy(subjectName: string, version: number) {
    GetSchemaRegistrySchemaReferencedBy(subjectName, version)
      .then((references) => {
        const cleanedReferences = [] as SchemaReference[];
        for (const ref of references) {
          if (ref.error) {
            console.error('error in refreshSchemaReferencedBy, reference entry has error', {
              subjectName,
              version,
              error: ref.error,
              refRaw: ref,
            });
            continue;
          }
          cleanedReferences.push(ref);
        }

        let subjectVersions = this.schemaReferencedBy.get(subjectName);
        if (!subjectVersions) {
          // @ts-ignore MobX does not play nice with TypeScript 5: Type 'ObservableMap<number, SchemaReferencedByEntry[]>' is not assignable to type 'Map<number, SchemaReferencedByEntry[]>'.
          subjectVersions = observable(new Map<number, SchemaReferencedByEntry[]>());
          if (subjectVersions) {
            this.schemaReferencedBy.set(subjectName, subjectVersions);
          }
        }

        subjectVersions?.set(version, cleanedReferences);
      })
      .catch(() => {
      });
  },

  async refreshSchemaUsagesById(schemaId: number): Promise<void> {
    type SchemaNotConfiguredType = { isConfigured: false };

    function isSchemaVersionArray(r: SchemaVersion[] | SchemaNotConfiguredType): r is SchemaVersion[] {
      return Array.isArray(r);
    }

    GetSchemaUsagesByID(schemaId)
      .then(
        (r) => {
          if (isSchemaVersionArray(r)) {
            this.schemaUsagesById.set(schemaId, r);
          }
        },
        (err) => {
          if (err instanceof Error) {
            // Currently we don't get helpful status codes (502) so we have to inspect the message
            if (err.message.includes('404') && err.message.includes('not found')) {
              // Do nothing, most likely cause is that the user has entered a value into the search box that doesn't exist
              return null;
            }
          }
          throw err;
        },
      );
  },

  async setSchemaRegistryCompatibilityMode(
    mode: CompatibilityLevel,
  ): Promise<SchemaRegistryConfig> {
    return PutSchemaRegistryConfig("", SetCompatibility.createFrom({compatibility: mode}))
  },

  async setSchemaRegistrySubjectCompatibilityMode(
    subjectName: string,
    mode: CompatibilityLevel,
  ): Promise<SchemaRegistryConfig> {
    if (mode === CompatibilityLevel.NONE) {
      await DeleteSchemaRegistrySubjectConfig(subjectName);
      return new Promise((_v) => {
      });
    }
    return PutSchemaRegistryConfig(subjectName, SetCompatibility.createFrom({compatibility: mode}))
  },

  async validateSchema(
    subjectName: string,
    version: number,
    schema: Schema,
  ): Promise<SchemaRegistrySchemaValidation> {
    return ValidateSchemaRegistrySchema(subjectName, version, schema);
  },

  async createSchema(
    subjectName: string,
    schema: Schema,
  ): Promise<CreateSchemaResponse> {
    return CreateSchemaRegistrySchema(subjectName, schema);
  },

  async deleteSchemaSubject(subjectName: string, permanent: boolean): Promise<SchemaRegistryDeleteSubjectResponse> {
    return DeleteSchemaRegistrySubject(subjectName, permanent);
  },

  async deleteSchemaSubjectVersion(
    subjectName: string,
    version: number,
    permanent: boolean,
  ): Promise<SchemaRegistryDeleteSubjectVersionResponse> {
    return DeleteSchemaRegistrySubjectVersion(subjectName, version, permanent);
  },

  refreshPartitionReassignments(): Promise<void> {
    return ListPartitionReassignments().then((v) => {
      if (v === null) this.partitionReassignments = null;
      else this.partitionReassignments = v;
    }, addError);
  },

  async startPartitionReassignment(
    request: AlterPartitionAssignmentsRequestTopic[],
  ): Promise<Array<AlterPartitionReassignmentsResponse>> {
    return await AlterPartitionAssignments(request);
  },

  async setReplicationThrottleRate(brokerIds: number[], maxBytesPerSecond: number): Promise<Array<IncrementalAlterConfigsResourceResponse>> {
    maxBytesPerSecond = Math.ceil(maxBytesPerSecond);

    const configRequest: IncrementalAlterConfigsRequestResource[] = [];

    for (const b of brokerIds) {
      configRequest.push(IncrementalAlterConfigsRequestResource.createFrom({
        ResourceType: ConfigResourceType.BROKER,
        ResourceName: String(b),
        Configs: [
          {
            Name: 'leader.replication.throttled.rate',
            Op: IncrementalAlterConfigOp.SET,
            Value: String(maxBytesPerSecond)
          },
          {
            Name: 'follower.replication.throttled.rate',
            Op: IncrementalAlterConfigOp.SET,
            Value: String(maxBytesPerSecond),
          },
        ],
      }));
    }

    return await this.changeConfig(configRequest);
  },

  async setThrottledReplicas(
    topicReplicas: {
      topicName: string;
      leaderReplicas: { brokerId: number; partitionId: number }[];
      followerReplicas: { brokerId: number; partitionId: number }[];
    }[],
  ): Promise<Array<IncrementalAlterConfigsResourceResponse>> {
    const configRequest: IncrementalAlterConfigsRequestResource[] = [];

    for (const t of topicReplicas) {
      const res: IncrementalAlterConfigsRequestResource = IncrementalAlterConfigsRequestResource.createFrom({
        // Set which topics to throttle
        ResourceType: ConfigResourceType.TOPIC,
        ResourceName: t.topicName,
        Configs: [],
      });

      const leaderReplicas = t.leaderReplicas.map((e) => `${e.partitionId}:${e.brokerId}`).join(',');
      res.Configs.push(IncrementalAlterConfigsRequestResourceConfig.createFrom({
        Name: 'leader.replication.throttled.replicas',
        Op: IncrementalAlterConfigOp.SET,
        Value: leaderReplicas,
      }));
      const followerReplicas = t.followerReplicas.map((e) => `${e.partitionId}:${e.brokerId}`).join(',');
      res.Configs.push(IncrementalAlterConfigsRequestResourceConfig.createFrom({
        Name: 'follower.replication.throttled.replicas',
        Op: IncrementalAlterConfigOp.SET,
        Value: followerReplicas,
      }));

      // individual request for each topic
      configRequest.push(res);
    }

    return await this.changeConfig(configRequest);
  },

  async resetThrottledReplicas(topicNames: string[]): Promise<Array<IncrementalAlterConfigsResourceResponse>> {
    const configRequest: IncrementalAlterConfigsRequestResource[] = [];

    // reset throttled replicas for those topics
    for (const t of topicNames) {
      configRequest.push(IncrementalAlterConfigsRequestResource.createFrom({
        ResourceType: ConfigResourceType.TOPIC,
        ResourceName: t,
        Configs: [
          {Name: 'leader.replication.throttled.replicas', Op: IncrementalAlterConfigOp.DELETE},
          {Name: 'follower.replication.throttled.replicas', Op: IncrementalAlterConfigOp.DELETE},
        ],
      }));
    }

    return await this.changeConfig(configRequest);
  },

  async resetReplicationThrottleRate(brokerIds: number[]): Promise<Array<IncrementalAlterConfigsResourceResponse>> {
    const configRequest: IncrementalAlterConfigsRequestResource[] = [];

    // We currently only set replication throttle on each broker, instead of cluster-wide (same effect, but different kind of 'ConfigSource')
    // So we don't remove the cluster-wide setting, only the ones we've set (the per-broker) settings

    // remove throttle configs from all brokers (DYNAMIC_DEFAULT_BROKER_CONFIG)
    // configRequest.resources.push({
    //     resourceType: ConfigResourceType.Broker,
    //     resourceName: "", // empty = all brokers
    //     configs: [
    //         { name: 'leader.replication.throttled.rate', op: AlterConfigOperation.Delete },
    //         { name: 'follower.replication.throttled.rate', op: AlterConfigOperation.Delete },
    //     ]
    // });

    // remove throttle configs from each broker individually (DYNAMIC_BROKER_CONFIG)
    for (const b of brokerIds) {
      configRequest.push(IncrementalAlterConfigsRequestResource.createFrom({
        ResourceType: ConfigResourceType.BROKER,
        ResourceName: String(b),
        Configs: [
          {Name: 'leader.replication.throttled.rate', Op: IncrementalAlterConfigOp.DELETE},
          {Name: 'follower.replication.throttled.rate', Op: IncrementalAlterConfigOp.DELETE},
        ],
      }));
    }

    return await this.changeConfig(configRequest);
  },

  async changeConfig(request: Array<IncrementalAlterConfigsRequestResource>): Promise<Array<IncrementalAlterConfigsResourceResponse>> {
    return IncrementalAlterConfigs(request);
  },


  // PATCH /topics/{topicName}/configuration   //
  // PATCH /topics/configuration               // default config
  async changeTopicConfig(topicName: string, configs: kmsg.IncrementalAlterConfigsRequestResourceConfig[]): Promise<void> {
    return EditTopicConfig(topicName, configs);
  },

  async publishRecords(request: PublishRecordsRequest): Promise<ProduceRecordsResponse> {
    // POST "/topics-records"
    return ProducePlainRecords(request.records, request.useTransactions, [CompressionCodec.createFrom(request.compressionType)])
  },

  // // New version of "publishRecords"
  // async publishMessage(request: PublishMessageRequest): Promise<PublishMessageResponse> {
  //   // biome-ignore lint/style/noNonNullAssertion: leave as is for now due to MobX
  //   const client = appConfig.consoleClient!;
  //   if (!client) {
  //     // this shouldn't happen but better to explicitly throw
  //     throw new Error('Console client is not initialized');
  //   }
  //   const r = await client.publishMessage(request);
  //
  //   return r;
  // },

  async createTopic(request: CreateTopicsRequestTopic): Promise<CreateTopicResponse> {
    return CreateTopic(request);
  },

};

export interface MessageSearchRequest {
  topicName: string;
  startOffset: number;
  startTimestamp: number;
  partitionId: number;
  maxResults: number; // should also support '-1' soon, so we can do live tailing
  filterInterpreterCode: string; // js code, base64 encoded
  includeRawPayload?: boolean;
  ignoreSizeLimit?: boolean;

  keyDeserializer?: PayloadEncoding;
  valueDeserializer?: PayloadEncoding;
}


export function createMessageSearch() {
  const messageSearch = {
    // Parameters last passed to 'startMessageSearch'
    searchRequest: null as MessageSearchRequest | null,

    // Some statistics that might be interesting to show in the UI
    elapsedMs: null as null | number, // Reported by the backend, only set once the search is done
    bytesConsumed: 0,
    totalMessagesConsumed: 0,

    // Call 'stopSearch' instead of using this directly
    abortController: null as AbortController | null,

    // Live view of messages, gets updated as new messages arrive
    messages: observable([] as TopicMessage[], {deep: false}),

    async startSearch(_searchRequest: MessageSearchRequest): Promise<TopicMessage[]> {

      this.searchRequest = searchRequest;
      this.searchPhase = 'Connecting';
      this.bytesConsumed = 0;
      this.totalMessagesConsumed = 0;
      this.messages.length = 0;
      this.elapsedMs = null;

      const messageSearchAbortController = (this.abortController = new AbortController());

      // do it
      const req = new ListMessagesRequest();
      req.topic = searchRequest.topicName;
      req.startOffset = BigInt(searchRequest.startOffset);
      req.startTimestamp = BigInt(searchRequest.startTimestamp);
      req.partitionId = searchRequest.partitionId;
      req.maxResults = searchRequest.maxResults;
      req.filterInterpreterCode = searchRequest.filterInterpreterCode;
      req.includeOriginalRawPayload = searchRequest.includeRawPayload ?? false;
      req.ignoreMaxSizeLimit = searchRequest.ignoreSizeLimit ?? false;
      req.keyDeserializer = searchRequest.keyDeserializer;
      req.valueDeserializer = searchRequest.valueDeserializer;

      // For StartOffset = Newest and any set push-down filter we need to bump the default timeout
      // from 30s to 30 minutes before ending the request gracefully.
      let timeoutMs = 30 * 1000;
      if (searchRequest.startOffset === PartitionOffsetOrigin.End || req.filterInterpreterCode != null) {
        const minuteMs = 60 * 1000;
        timeoutMs = 30 * minuteMs;
      }

      try {
        for await (const res of client.listMessages(req, {signal: messageSearchAbortController.signal, timeoutMs})) {
          if (messageSearchAbortController.signal.aborted) break;

          try {
            switch (res.controlMessage.case) {
              case 'phase':
                console.log(`phase: ${res.controlMessage.value.phase}`);
                this.searchPhase = res.controlMessage.value.phase;
                break;
              case 'progress':
                console.log(`progress: ${res.controlMessage.value.messagesConsumed}`);
                this.bytesConsumed = Number(res.controlMessage.value.bytesConsumed);
                this.totalMessagesConsumed = Number(res.controlMessage.value.messagesConsumed);
                break;
              case 'done':
                this.elapsedMs = Number(res.controlMessage.value.elapsedMs);
                this.bytesConsumed = Number(res.controlMessage.value.bytesConsumed);
                // this.MessageSearchCancelled = msg.isCancelled;
                this.searchPhase = 'Done';
                this.searchPhase = null;
                break;
              case 'error':
                // error doesn't necessarily mean the whole request is done
                console.info(`ws backend error: ${res.controlMessage.value.message}`);
                toast({
                  title: 'Backend Error',
                  description: res.controlMessage.value.message,
                  status: 'error',
                });

                break;
              case 'data': {
                // TODO I would guess we should replace the rest interface types and just utilize the generated Connect types
                // this is my hacky way of attempting to get things working by converting the Connect types
                // to the rest interface types that are hooked up to other things

                const m = {} as TopicMessage;
                m.partitionID = res.controlMessage.value.partitionId;

                m.compression = CompressionType.Unknown;
                switch (res.controlMessage.value.compression) {
                  case ProtoCompressionType.UNCOMPRESSED:
                    m.compression = CompressionType.Uncompressed;
                    break;
                  case ProtoCompressionType.GZIP:
                    m.compression = CompressionType.GZip;
                    break;
                  case ProtoCompressionType.SNAPPY:
                    m.compression = CompressionType.Snappy;
                    break;
                  case ProtoCompressionType.LZ4:
                    m.compression = CompressionType.LZ4;
                    break;
                  case ProtoCompressionType.ZSTD:
                    m.compression = CompressionType.ZStd;
                    break;
                }

                m.offset = Number(res.controlMessage.value.offset);
                m.timestamp = Number(res.controlMessage.value.timestamp);
                m.isTransactional = res.controlMessage.value.isTransactional;
                m.headers = [];
                for (const header of res.controlMessage.value.headers) {
                  m.headers.push({
                    key: header.key,
                    value: {
                      payload: JSON.stringify(new TextDecoder().decode(header.value)),
                      encoding: 'text',
                      schemaId: 0,
                      size: header.value.length,
                      isPayloadNull: header.value == null,
                    },
                  });
                }

                // key
                const key = res.controlMessage.value.key;
                const keyPayload = new TextDecoder().decode(key?.normalizedPayload);

                m.key = {} as Payload;
                m.key.rawBytes = key?.originalPayload;

                switch (key?.encoding) {
                  case PayloadEncoding.NULL:
                    m.key.encoding = 'null';
                    break;
                  case PayloadEncoding.BINARY:
                    m.key.encoding = 'binary';
                    break;
                  case PayloadEncoding.XML:
                    m.key.encoding = 'xml';
                    break;
                  case PayloadEncoding.AVRO:
                    m.key.encoding = 'avro';
                    break;
                  case PayloadEncoding.JSON:
                    m.key.encoding = 'json';
                    break;
                  case PayloadEncoding.PROTOBUF:
                    m.key.encoding = 'protobuf';
                    break;
                  case PayloadEncoding.MESSAGE_PACK:
                    m.key.encoding = 'msgpack';
                    break;
                  case PayloadEncoding.TEXT:
                    m.key.encoding = 'text';
                    break;
                  case PayloadEncoding.UTF8:
                    m.key.encoding = 'utf8WithControlChars';
                    break;
                  case PayloadEncoding.UINT:
                    m.key.encoding = 'uint';
                    break;
                  case PayloadEncoding.SMILE:
                    m.key.encoding = 'smile';
                    break;
                  case PayloadEncoding.CONSUMER_OFFSETS:
                    m.key.encoding = 'consumerOffsets';
                    break;
                  case PayloadEncoding.CBOR:
                    m.key.encoding = 'cbor';
                    break;
                  default:
                    console.log('unhandled key encoding type', {
                      encoding: key?.encoding,
                      encodingName:
                        key?.encoding != null
                          ? proto3.getEnumType(PayloadEncoding).findNumber(key.encoding)?.localName
                          : undefined,
                      message: res,
                    });
                }

                m.key.isPayloadNull = key?.encoding === PayloadEncoding.NULL;
                m.key.payload = keyPayload;
                m.key.normalizedPayload = key?.normalizedPayload;

                try {
                  m.key.payload = JSON.parse(keyPayload);
                } catch {
                }

                m.key.troubleshootReport = key?.troubleshootReport;
                m.key.schemaId = key?.schemaId ?? 0;
                m.keyJson = JSON.stringify(m.key.payload);
                m.key.size = Number(key?.payloadSize);
                m.key.isPayloadTooLarge = key?.isPayloadTooLarge;

                // kconsole.log(m.keyJson)

                // value
                const val = res.controlMessage.value.value;
                const valuePayload = new TextDecoder().decode(val?.normalizedPayload);

                m.value = {} as Payload;
                m.value.payload = valuePayload;
                m.value.normalizedPayload = val?.normalizedPayload;
                m.value.rawBytes = val?.originalPayload;

                switch (val?.encoding) {
                  case PayloadEncoding.NULL:
                    m.value.encoding = 'null';
                    break;
                  case PayloadEncoding.BINARY:
                    m.value.encoding = 'binary';
                    break;
                  case PayloadEncoding.XML:
                    m.value.encoding = 'xml';
                    break;
                  case PayloadEncoding.AVRO:
                    m.value.encoding = 'avro';
                    break;
                  case PayloadEncoding.JSON:
                    m.value.encoding = 'json';
                    break;
                  case PayloadEncoding.PROTOBUF:
                    m.value.encoding = 'protobuf';
                    break;
                  case PayloadEncoding.MESSAGE_PACK:
                    m.value.encoding = 'msgpack';
                    break;
                  case PayloadEncoding.TEXT:
                    m.value.encoding = 'text';
                    break;
                  case PayloadEncoding.UTF8:
                    m.value.encoding = 'utf8WithControlChars';
                    break;
                  case PayloadEncoding.UINT:
                    m.value.encoding = 'uint';
                    break;
                  case PayloadEncoding.SMILE:
                    m.value.encoding = 'smile';
                    break;
                  case PayloadEncoding.CONSUMER_OFFSETS:
                    m.value.encoding = 'consumerOffsets';
                    break;
                  case PayloadEncoding.CBOR:
                    m.value.encoding = 'cbor';
                    break;
                  default:
                    console.log('unhandled value encoding type', {
                      encoding: val?.encoding,
                      encodingName:
                        val?.encoding != null
                          ? proto3.getEnumType(PayloadEncoding).findNumber(val.encoding)?.localName
                          : undefined,
                      message: res,
                    });
                }

                m.value.schemaId = val?.schemaId ?? 0;
                m.value.troubleshootReport = val?.troubleshootReport;
                m.value.isPayloadNull = val?.encoding === PayloadEncoding.NULL;
                m.valueJson = valuePayload;
                m.value.isPayloadTooLarge = val?.isPayloadTooLarge;

                try {
                  m.value.payload = JSON.parse(valuePayload);
                } catch {
                }

                m.valueJson = JSON.stringify(m.value.payload);
                m.value.size = Number(val?.payloadSize);

                this.messages.push(m);
                break;
              }
            }
          } catch (e) {
            console.error('error in listMessages loop', {error: e});
          }
        }
      } catch (e) {
        this.abortController = null;
        this.searchPhase = 'Done';
        this.bytesConsumed = 0;
        this.totalMessagesConsumed = 0;
        this.searchPhase = null;
        // https://connectrpc.com/docs/web/errors
        if (messageSearchAbortController.signal.aborted) {
          // Do not throw, this is a user cancellation
        } else {
          console.error('startMessageSearchNew: error in await loop of client.listMessages', {error: e});
          throw e;
        }
      }

      // one done
      this.stopSearch();
      return this.messages;
    },

  };

  return observable(messageSearch);
}

export type MessageSearch = ReturnType<typeof createMessageSearch>;


function addFrontendFieldsForConsumerGroup(g: ConsumerGroupOverview) {
  g.lagSum = g.topicOffsets.sum((o) => o.summedLag);

  if (g.allowedActions) {
    if (g.allowedActions.includes('all')) {
      // All perms
    } else {
      // Not all perms, set helper props
      g.noEditPerms = !g.allowedActions?.includes('editConsumerGroup');
      g.noDeletePerms = !g.allowedActions?.includes('deleteConsumerGroup');
    }
  }
  g.isInUse = g.state.toLowerCase() !== 'empty';
}

export const brokerMap = computed(
  () => {
    const brokers = api.clusterInfo?.brokers;
    if (brokers == null) return null;

    const map = new Map<number, kconsole.Broker>();
    for (const b of brokers) map.set(b.brokerId, b);

    return map;
  },
  {name: 'brokerMap', equals: comparer.structural},
);

// 1. add 'type' to each synonym, so when expanding a config entry (to view its synonyms), we can still see the type
// 2. remove redundant synonym entries (those that have the same source as the root config entry)
function prepareSynonyms(configEntries: TopicConfigEntry[]) {
  if (!Array.isArray(configEntries)) return;

  for (const e of configEntries) {
    if (e.synonyms === undefined) continue;

    // add 'type' from root object
    for (const s of e.synonyms) s.type = e.type;
  }
}

function prepareBrokerSynonyms(configEntries: BrokerConfigEntry[]) {
  if (!Array.isArray(configEntries)) return;

  for (const e of configEntries) {
    if (e.synonyms === undefined) continue;

    // add 'type' from root object
    for (const s of e.synonyms) s.type = e.type;
  }
}

function normalizeAcls(acls: ACLResource[]) {
  function upperFirst(str: string): string {
    if (!str) return str;
    const lower = str.toLowerCase();
    const first = lower[0];
    const result = first.toUpperCase() + lower.slice(1);
    return result;
  }

  const specialCaseMap = {
    TRANSACTIONAL_ID: 'TransactionalID',
  } as { [key: string]: string };

  function normalizeStringEnum<T extends string>(str: T): T {
    if (!str) return str;
    if (specialCaseMap[str]) return specialCaseMap[str] as T;

    const parts = str.split('_');
    for (let i = 0; i < parts.length; i++) {
      parts[i] = upperFirst(parts[i].toLowerCase());
    }
    const result = parts.join('');
    return result as T;
  }

  for (const e of acls) {
    e.resourceType = normalizeStringEnum(e.resourceType);
    e.resourcePatternType = normalizeStringEnum(e.resourcePatternType);

    for (const acl of e.acls) {
      acl.operation = normalizeStringEnum(acl.operation);
      acl.permissionType = normalizeStringEnum(acl.permissionType);
    }
  }
}


export async function partialTopicConfigs(
  configKeys: string[],
  topics: string[],
): Promise<Array<kconsole.TopicConfig>> {
  return GetTopicsConfigs(topics, configKeys).then(Object.values);
}

// export interface MessageSearchRequest {
//   topicName: string;
//   startOffset: number;
//   startTimestamp: number;
//   partitionId: number;
//   maxResults: number; // should also support '-1' soon, so we can do live tailing
//   filterInterpreterCode: string; // js code, base64 encoded
//   enterprise?: {
//     redpandaCloud?: {
//       accessToken: string;
//     };
//   };
//   includeRawPayload?: boolean;
//   ignoreSizeLimit?: boolean;
//
//   keyDeserializer?: PayloadEncoding;
//   valueDeserializer?: PayloadEncoding;
// }

function addError(err: Error) {
  api.errors.push(err);
}

type apiStoreType = typeof apiStore;
export const api = observable(apiStore) as apiStoreType;
