package documentdbapi

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
    "context"
    "github.com/Azure/azure-sdk-for-go/services/cosmos-db/mgmt/2015-04-08/documentdb"
    "github.com/Azure/go-autorest/autorest"
)

        // DatabaseAccountsClientAPI contains the set of methods on the DatabaseAccountsClient type.
        type DatabaseAccountsClientAPI interface {
            CheckNameExists(ctx context.Context, accountName string) (result autorest.Response, err error)
            CreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, createUpdateParameters documentdb.DatabaseAccountCreateUpdateParameters) (result documentdb.DatabaseAccountsCreateOrUpdateFuture, err error)
            CreateUpdateCassandraKeyspace(ctx context.Context, resourceGroupName string, accountName string, keyspaceName string, createUpdateCassandraKeyspaceParameters documentdb.CassandraKeyspaceCreateUpdateParameters) (result documentdb.DatabaseAccountsCreateUpdateCassandraKeyspaceFuture, err error)
            CreateUpdateCassandraTable(ctx context.Context, resourceGroupName string, accountName string, keyspaceName string, tableName string, createUpdateCassandraTableParameters documentdb.CassandraTableCreateUpdateParameters) (result documentdb.DatabaseAccountsCreateUpdateCassandraTableFuture, err error)
            CreateUpdateGremlinDatabase(ctx context.Context, resourceGroupName string, accountName string, databaseName string, createUpdateGremlinDatabaseParameters documentdb.GremlinDatabaseCreateUpdateParameters) (result documentdb.DatabaseAccountsCreateUpdateGremlinDatabaseFuture, err error)
            CreateUpdateGremlinGraph(ctx context.Context, resourceGroupName string, accountName string, databaseName string, graphName string, createUpdateGremlinGraphParameters documentdb.GremlinGraphCreateUpdateParameters) (result documentdb.DatabaseAccountsCreateUpdateGremlinGraphFuture, err error)
            CreateUpdateMongoDBCollection(ctx context.Context, resourceGroupName string, accountName string, databaseName string, collectionName string, createUpdateMongoDBCollectionParameters documentdb.MongoDBCollectionCreateUpdateParameters) (result documentdb.DatabaseAccountsCreateUpdateMongoDBCollectionFuture, err error)
            CreateUpdateMongoDBDatabase(ctx context.Context, resourceGroupName string, accountName string, databaseName string, createUpdateMongoDBDatabaseParameters documentdb.MongoDBDatabaseCreateUpdateParameters) (result documentdb.DatabaseAccountsCreateUpdateMongoDBDatabaseFuture, err error)
            CreateUpdateSQLContainer(ctx context.Context, resourceGroupName string, accountName string, databaseName string, containerName string, createUpdateSQLContainerParameters documentdb.SQLContainerCreateUpdateParameters) (result documentdb.DatabaseAccountsCreateUpdateSQLContainerFuture, err error)
            CreateUpdateSQLDatabase(ctx context.Context, resourceGroupName string, accountName string, databaseName string, createUpdateSQLDatabaseParameters documentdb.SQLDatabaseCreateUpdateParameters) (result documentdb.DatabaseAccountsCreateUpdateSQLDatabaseFuture, err error)
            CreateUpdateTable(ctx context.Context, resourceGroupName string, accountName string, tableName string, createUpdateTableParameters documentdb.TableCreateUpdateParameters) (result documentdb.DatabaseAccountsCreateUpdateTableFuture, err error)
            Delete(ctx context.Context, resourceGroupName string, accountName string) (result documentdb.DatabaseAccountsDeleteFuture, err error)
            DeleteCassandraKeyspace(ctx context.Context, resourceGroupName string, accountName string, keyspaceName string) (result documentdb.DatabaseAccountsDeleteCassandraKeyspaceFuture, err error)
            DeleteCassandraTable(ctx context.Context, resourceGroupName string, accountName string, keyspaceName string, tableName string) (result documentdb.DatabaseAccountsDeleteCassandraTableFuture, err error)
            DeleteGremlinDatabase(ctx context.Context, resourceGroupName string, accountName string, databaseName string) (result documentdb.DatabaseAccountsDeleteGremlinDatabaseFuture, err error)
            DeleteGremlinGraph(ctx context.Context, resourceGroupName string, accountName string, databaseName string, graphName string) (result documentdb.DatabaseAccountsDeleteGremlinGraphFuture, err error)
            DeleteMongoDBCollection(ctx context.Context, resourceGroupName string, accountName string, databaseName string, collectionName string) (result documentdb.DatabaseAccountsDeleteMongoDBCollectionFuture, err error)
            DeleteMongoDBDatabase(ctx context.Context, resourceGroupName string, accountName string, databaseName string) (result documentdb.DatabaseAccountsDeleteMongoDBDatabaseFuture, err error)
            DeleteSQLContainer(ctx context.Context, resourceGroupName string, accountName string, databaseName string, containerName string) (result documentdb.DatabaseAccountsDeleteSQLContainerFuture, err error)
            DeleteSQLDatabase(ctx context.Context, resourceGroupName string, accountName string, databaseName string) (result documentdb.DatabaseAccountsDeleteSQLDatabaseFuture, err error)
            DeleteTable(ctx context.Context, resourceGroupName string, accountName string, tableName string) (result documentdb.DatabaseAccountsDeleteTableFuture, err error)
            FailoverPriorityChange(ctx context.Context, resourceGroupName string, accountName string, failoverParameters documentdb.FailoverPolicies) (result documentdb.DatabaseAccountsFailoverPriorityChangeFuture, err error)
            Get(ctx context.Context, resourceGroupName string, accountName string) (result documentdb.DatabaseAccount, err error)
            GetCassandraKeyspace(ctx context.Context, resourceGroupName string, accountName string, keyspaceName string) (result documentdb.CassandraKeyspace, err error)
            GetCassandraKeyspaceThroughput(ctx context.Context, resourceGroupName string, accountName string, keyspaceName string) (result documentdb.Throughput, err error)
            GetCassandraTable(ctx context.Context, resourceGroupName string, accountName string, keyspaceName string, tableName string) (result documentdb.CassandraTable, err error)
            GetCassandraTableThroughput(ctx context.Context, resourceGroupName string, accountName string, keyspaceName string, tableName string) (result documentdb.Throughput, err error)
            GetGremlinDatabase(ctx context.Context, resourceGroupName string, accountName string, databaseName string) (result documentdb.GremlinDatabase, err error)
            GetGremlinDatabaseThroughput(ctx context.Context, resourceGroupName string, accountName string, databaseName string) (result documentdb.Throughput, err error)
            GetGremlinGraph(ctx context.Context, resourceGroupName string, accountName string, databaseName string, graphName string) (result documentdb.GremlinGraph, err error)
            GetGremlinGraphThroughput(ctx context.Context, resourceGroupName string, accountName string, databaseName string, graphName string) (result documentdb.Throughput, err error)
            GetMongoDBCollection(ctx context.Context, resourceGroupName string, accountName string, databaseName string, collectionName string) (result documentdb.MongoDBCollection, err error)
            GetMongoDBCollectionThroughput(ctx context.Context, resourceGroupName string, accountName string, databaseName string, collectionName string) (result documentdb.Throughput, err error)
            GetMongoDBDatabase(ctx context.Context, resourceGroupName string, accountName string, databaseName string) (result documentdb.MongoDBDatabase, err error)
            GetMongoDBDatabaseThroughput(ctx context.Context, resourceGroupName string, accountName string, databaseName string) (result documentdb.Throughput, err error)
            GetReadOnlyKeys(ctx context.Context, resourceGroupName string, accountName string) (result documentdb.DatabaseAccountListReadOnlyKeysResult, err error)
            GetSQLContainer(ctx context.Context, resourceGroupName string, accountName string, databaseName string, containerName string) (result documentdb.SQLContainer, err error)
            GetSQLContainerThroughput(ctx context.Context, resourceGroupName string, accountName string, databaseName string, containerName string) (result documentdb.Throughput, err error)
            GetSQLDatabase(ctx context.Context, resourceGroupName string, accountName string, databaseName string) (result documentdb.SQLDatabase, err error)
            GetSQLDatabaseThroughput(ctx context.Context, resourceGroupName string, accountName string, databaseName string) (result documentdb.Throughput, err error)
            GetTable(ctx context.Context, resourceGroupName string, accountName string, tableName string) (result documentdb.Table, err error)
            GetTableThroughput(ctx context.Context, resourceGroupName string, accountName string, tableName string) (result documentdb.Throughput, err error)
            List(ctx context.Context) (result documentdb.DatabaseAccountsListResult, err error)
            ListByResourceGroup(ctx context.Context, resourceGroupName string) (result documentdb.DatabaseAccountsListResult, err error)
            ListCassandraKeyspaces(ctx context.Context, resourceGroupName string, accountName string) (result documentdb.CassandraKeyspaceListResult, err error)
            ListCassandraTables(ctx context.Context, resourceGroupName string, accountName string, keyspaceName string) (result documentdb.CassandraTableListResult, err error)
            ListConnectionStrings(ctx context.Context, resourceGroupName string, accountName string) (result documentdb.DatabaseAccountListConnectionStringsResult, err error)
            ListGremlinDatabases(ctx context.Context, resourceGroupName string, accountName string) (result documentdb.GremlinDatabaseListResult, err error)
            ListGremlinGraphs(ctx context.Context, resourceGroupName string, accountName string, databaseName string) (result documentdb.GremlinGraphListResult, err error)
            ListKeys(ctx context.Context, resourceGroupName string, accountName string) (result documentdb.DatabaseAccountListKeysResult, err error)
            ListMetricDefinitions(ctx context.Context, resourceGroupName string, accountName string) (result documentdb.MetricDefinitionsListResult, err error)
            ListMetrics(ctx context.Context, resourceGroupName string, accountName string, filter string) (result documentdb.MetricListResult, err error)
            ListMongoDBCollections(ctx context.Context, resourceGroupName string, accountName string, databaseName string) (result documentdb.MongoDBCollectionListResult, err error)
            ListMongoDBDatabases(ctx context.Context, resourceGroupName string, accountName string) (result documentdb.MongoDBDatabaseListResult, err error)
            ListReadOnlyKeys(ctx context.Context, resourceGroupName string, accountName string) (result documentdb.DatabaseAccountListReadOnlyKeysResult, err error)
            ListSQLContainers(ctx context.Context, resourceGroupName string, accountName string, databaseName string) (result documentdb.SQLContainerListResult, err error)
            ListSQLDatabases(ctx context.Context, resourceGroupName string, accountName string) (result documentdb.SQLDatabaseListResult, err error)
            ListTables(ctx context.Context, resourceGroupName string, accountName string) (result documentdb.TableListResult, err error)
            ListUsages(ctx context.Context, resourceGroupName string, accountName string, filter string) (result documentdb.UsagesResult, err error)
            OfflineRegion(ctx context.Context, resourceGroupName string, accountName string, regionParameterForOffline documentdb.RegionForOnlineOffline) (result documentdb.DatabaseAccountsOfflineRegionFuture, err error)
            OnlineRegion(ctx context.Context, resourceGroupName string, accountName string, regionParameterForOnline documentdb.RegionForOnlineOffline) (result documentdb.DatabaseAccountsOnlineRegionFuture, err error)
            Patch(ctx context.Context, resourceGroupName string, accountName string, updateParameters documentdb.DatabaseAccountPatchParameters) (result documentdb.DatabaseAccountsPatchFuture, err error)
            RegenerateKey(ctx context.Context, resourceGroupName string, accountName string, keyToRegenerate documentdb.DatabaseAccountRegenerateKeyParameters) (result documentdb.DatabaseAccountsRegenerateKeyFuture, err error)
            UpdateCassandraKeyspaceThroughput(ctx context.Context, resourceGroupName string, accountName string, keyspaceName string, updateThroughputParameters documentdb.ThroughputUpdateParameters) (result documentdb.DatabaseAccountsUpdateCassandraKeyspaceThroughputFuture, err error)
            UpdateCassandraTableThroughput(ctx context.Context, resourceGroupName string, accountName string, keyspaceName string, tableName string, updateThroughputParameters documentdb.ThroughputUpdateParameters) (result documentdb.DatabaseAccountsUpdateCassandraTableThroughputFuture, err error)
            UpdateGremlinDatabaseThroughput(ctx context.Context, resourceGroupName string, accountName string, databaseName string, updateThroughputParameters documentdb.ThroughputUpdateParameters) (result documentdb.DatabaseAccountsUpdateGremlinDatabaseThroughputFuture, err error)
            UpdateGremlinGraphThroughput(ctx context.Context, resourceGroupName string, accountName string, databaseName string, graphName string, updateThroughputParameters documentdb.ThroughputUpdateParameters) (result documentdb.DatabaseAccountsUpdateGremlinGraphThroughputFuture, err error)
            UpdateMongoDBCollectionThroughput(ctx context.Context, resourceGroupName string, accountName string, databaseName string, collectionName string, updateThroughputParameters documentdb.ThroughputUpdateParameters) (result documentdb.DatabaseAccountsUpdateMongoDBCollectionThroughputFuture, err error)
            UpdateMongoDBDatabaseThroughput(ctx context.Context, resourceGroupName string, accountName string, databaseName string, updateThroughputParameters documentdb.ThroughputUpdateParameters) (result documentdb.DatabaseAccountsUpdateMongoDBDatabaseThroughputFuture, err error)
            UpdateSQLContainerThroughput(ctx context.Context, resourceGroupName string, accountName string, databaseName string, containerName string, updateThroughputParameters documentdb.ThroughputUpdateParameters) (result documentdb.DatabaseAccountsUpdateSQLContainerThroughputFuture, err error)
            UpdateSQLDatabaseThroughput(ctx context.Context, resourceGroupName string, accountName string, databaseName string, updateThroughputParameters documentdb.ThroughputUpdateParameters) (result documentdb.DatabaseAccountsUpdateSQLDatabaseThroughputFuture, err error)
            UpdateTableThroughput(ctx context.Context, resourceGroupName string, accountName string, tableName string, updateThroughputParameters documentdb.ThroughputUpdateParameters) (result documentdb.DatabaseAccountsUpdateTableThroughputFuture, err error)
        }

        var _ DatabaseAccountsClientAPI = (*documentdb.DatabaseAccountsClient)(nil)
        // OperationsClientAPI contains the set of methods on the OperationsClient type.
        type OperationsClientAPI interface {
            List(ctx context.Context) (result documentdb.OperationListResultPage, err error)
                ListComplete(ctx context.Context) (result documentdb.OperationListResultIterator, err error)
        }

        var _ OperationsClientAPI = (*documentdb.OperationsClient)(nil)
        // DatabaseClientAPI contains the set of methods on the DatabaseClient type.
        type DatabaseClientAPI interface {
            ListMetricDefinitions(ctx context.Context, resourceGroupName string, accountName string, databaseRid string) (result documentdb.MetricDefinitionsListResult, err error)
            ListMetrics(ctx context.Context, resourceGroupName string, accountName string, databaseRid string, filter string) (result documentdb.MetricListResult, err error)
            ListUsages(ctx context.Context, resourceGroupName string, accountName string, databaseRid string, filter string) (result documentdb.UsagesResult, err error)
        }

        var _ DatabaseClientAPI = (*documentdb.DatabaseClient)(nil)
        // CollectionClientAPI contains the set of methods on the CollectionClient type.
        type CollectionClientAPI interface {
            ListMetricDefinitions(ctx context.Context, resourceGroupName string, accountName string, databaseRid string, collectionRid string) (result documentdb.MetricDefinitionsListResult, err error)
            ListMetrics(ctx context.Context, resourceGroupName string, accountName string, databaseRid string, collectionRid string, filter string) (result documentdb.MetricListResult, err error)
            ListUsages(ctx context.Context, resourceGroupName string, accountName string, databaseRid string, collectionRid string, filter string) (result documentdb.UsagesResult, err error)
        }

        var _ CollectionClientAPI = (*documentdb.CollectionClient)(nil)
        // CollectionRegionClientAPI contains the set of methods on the CollectionRegionClient type.
        type CollectionRegionClientAPI interface {
            ListMetrics(ctx context.Context, resourceGroupName string, accountName string, region string, databaseRid string, collectionRid string, filter string) (result documentdb.MetricListResult, err error)
        }

        var _ CollectionRegionClientAPI = (*documentdb.CollectionRegionClient)(nil)
        // DatabaseAccountRegionClientAPI contains the set of methods on the DatabaseAccountRegionClient type.
        type DatabaseAccountRegionClientAPI interface {
            ListMetrics(ctx context.Context, resourceGroupName string, accountName string, region string, filter string) (result documentdb.MetricListResult, err error)
        }

        var _ DatabaseAccountRegionClientAPI = (*documentdb.DatabaseAccountRegionClient)(nil)
        // PercentileSourceTargetClientAPI contains the set of methods on the PercentileSourceTargetClient type.
        type PercentileSourceTargetClientAPI interface {
            ListMetrics(ctx context.Context, resourceGroupName string, accountName string, sourceRegion string, targetRegion string, filter string) (result documentdb.PercentileMetricListResult, err error)
        }

        var _ PercentileSourceTargetClientAPI = (*documentdb.PercentileSourceTargetClient)(nil)
        // PercentileTargetClientAPI contains the set of methods on the PercentileTargetClient type.
        type PercentileTargetClientAPI interface {
            ListMetrics(ctx context.Context, resourceGroupName string, accountName string, targetRegion string, filter string) (result documentdb.PercentileMetricListResult, err error)
        }

        var _ PercentileTargetClientAPI = (*documentdb.PercentileTargetClient)(nil)
        // PercentileClientAPI contains the set of methods on the PercentileClient type.
        type PercentileClientAPI interface {
            ListMetrics(ctx context.Context, resourceGroupName string, accountName string, filter string) (result documentdb.PercentileMetricListResult, err error)
        }

        var _ PercentileClientAPI = (*documentdb.PercentileClient)(nil)
        // CollectionPartitionRegionClientAPI contains the set of methods on the CollectionPartitionRegionClient type.
        type CollectionPartitionRegionClientAPI interface {
            ListMetrics(ctx context.Context, resourceGroupName string, accountName string, region string, databaseRid string, collectionRid string, filter string) (result documentdb.PartitionMetricListResult, err error)
        }

        var _ CollectionPartitionRegionClientAPI = (*documentdb.CollectionPartitionRegionClient)(nil)
        // CollectionPartitionClientAPI contains the set of methods on the CollectionPartitionClient type.
        type CollectionPartitionClientAPI interface {
            ListMetrics(ctx context.Context, resourceGroupName string, accountName string, databaseRid string, collectionRid string, filter string) (result documentdb.PartitionMetricListResult, err error)
            ListUsages(ctx context.Context, resourceGroupName string, accountName string, databaseRid string, collectionRid string, filter string) (result documentdb.PartitionUsagesResult, err error)
        }

        var _ CollectionPartitionClientAPI = (*documentdb.CollectionPartitionClient)(nil)
        // PartitionKeyRangeIDClientAPI contains the set of methods on the PartitionKeyRangeIDClient type.
        type PartitionKeyRangeIDClientAPI interface {
            ListMetrics(ctx context.Context, resourceGroupName string, accountName string, databaseRid string, collectionRid string, partitionKeyRangeID string, filter string) (result documentdb.PartitionMetricListResult, err error)
        }

        var _ PartitionKeyRangeIDClientAPI = (*documentdb.PartitionKeyRangeIDClient)(nil)
        // PartitionKeyRangeIDRegionClientAPI contains the set of methods on the PartitionKeyRangeIDRegionClient type.
        type PartitionKeyRangeIDRegionClientAPI interface {
            ListMetrics(ctx context.Context, resourceGroupName string, accountName string, region string, databaseRid string, collectionRid string, partitionKeyRangeID string, filter string) (result documentdb.PartitionMetricListResult, err error)
        }

        var _ PartitionKeyRangeIDRegionClientAPI = (*documentdb.PartitionKeyRangeIDRegionClient)(nil)
