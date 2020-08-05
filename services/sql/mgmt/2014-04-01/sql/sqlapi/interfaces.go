package sqlapi

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
	"github.com/Azure/azure-sdk-for-go/services/sql/mgmt/2014-04-01/sql"
	"github.com/Azure/go-autorest/autorest"
)

// ServerAdvisorsClientAPI contains the set of methods on the ServerAdvisorsClient type.
type ServerAdvisorsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, advisorName string, parameters sql.Advisor) (result sql.Advisor, err error)
	Get(ctx context.Context, resourceGroupName string, serverName string, advisorName string) (result sql.Advisor, err error)
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result sql.AdvisorListResult, err error)
	Update(ctx context.Context, resourceGroupName string, serverName string, advisorName string, parameters sql.Advisor) (result sql.Advisor, err error)
}

var _ ServerAdvisorsClientAPI = (*sql.ServerAdvisorsClient)(nil)

// DatabaseAdvisorsClientAPI contains the set of methods on the DatabaseAdvisorsClient type.
type DatabaseAdvisorsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, databaseName string, advisorName string, parameters sql.Advisor) (result sql.Advisor, err error)
	Get(ctx context.Context, resourceGroupName string, serverName string, databaseName string, advisorName string) (result sql.Advisor, err error)
	ListByDatabase(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result sql.AdvisorListResult, err error)
}

var _ DatabaseAdvisorsClientAPI = (*sql.DatabaseAdvisorsClient)(nil)

// RecoverableDatabasesClientAPI contains the set of methods on the RecoverableDatabasesClient type.
type RecoverableDatabasesClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result sql.RecoverableDatabase, err error)
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result sql.RecoverableDatabaseListResult, err error)
}

var _ RecoverableDatabasesClientAPI = (*sql.RecoverableDatabasesClient)(nil)

// RestorableDroppedDatabasesClientAPI contains the set of methods on the RestorableDroppedDatabasesClient type.
type RestorableDroppedDatabasesClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, serverName string, restorableDroppededDatabaseID string) (result sql.RestorableDroppedDatabase, err error)
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result sql.RestorableDroppedDatabaseListResult, err error)
}

var _ RestorableDroppedDatabasesClientAPI = (*sql.RestorableDroppedDatabasesClient)(nil)

// CapabilitiesClientAPI contains the set of methods on the CapabilitiesClient type.
type CapabilitiesClientAPI interface {
	ListByLocation(ctx context.Context, locationID string) (result sql.LocationCapabilities, err error)
}

var _ CapabilitiesClientAPI = (*sql.CapabilitiesClient)(nil)

// ServersClientAPI contains the set of methods on the ServersClient type.
type ServersClientAPI interface {
	CheckNameAvailability(ctx context.Context, parameters sql.CheckNameAvailabilityRequest) (result sql.CheckNameAvailabilityResponse, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, parameters sql.Server) (result sql.Server, err error)
	Delete(ctx context.Context, resourceGroupName string, serverName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, serverName string) (result sql.Server, err error)
	List(ctx context.Context) (result sql.ServerListResult, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result sql.ServerListResult, err error)
	Update(ctx context.Context, resourceGroupName string, serverName string, parameters sql.ServerUpdate) (result sql.Server, err error)
}

var _ ServersClientAPI = (*sql.ServersClient)(nil)

// ServerConnectionPoliciesClientAPI contains the set of methods on the ServerConnectionPoliciesClient type.
type ServerConnectionPoliciesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, parameters sql.ServerConnectionPolicy) (result sql.ServerConnectionPolicy, err error)
	Get(ctx context.Context, resourceGroupName string, serverName string) (result sql.ServerConnectionPolicy, err error)
}

var _ ServerConnectionPoliciesClientAPI = (*sql.ServerConnectionPoliciesClient)(nil)

// DatabasesClientAPI contains the set of methods on the DatabasesClient type.
type DatabasesClientAPI interface {
	CreateImportOperation(ctx context.Context, resourceGroupName string, serverName string, databaseName string, parameters sql.ImportExtensionRequest) (result sql.DatabasesCreateImportOperationFuture, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, databaseName string, parameters sql.Database) (result sql.DatabasesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result autorest.Response, err error)
	Export(ctx context.Context, resourceGroupName string, serverName string, databaseName string, parameters sql.ExportRequest) (result sql.DatabasesExportFuture, err error)
	Get(ctx context.Context, resourceGroupName string, serverName string, databaseName string, expand string) (result sql.Database, err error)
	GetByElasticPool(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string, databaseName string) (result sql.Database, err error)
	GetByRecommendedElasticPool(ctx context.Context, resourceGroupName string, serverName string, recommendedElasticPoolName string, databaseName string) (result sql.Database, err error)
	Import(ctx context.Context, resourceGroupName string, serverName string, parameters sql.ImportRequest) (result sql.DatabasesImportFuture, err error)
	ListByElasticPool(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string) (result sql.DatabaseListResult, err error)
	ListByRecommendedElasticPool(ctx context.Context, resourceGroupName string, serverName string, recommendedElasticPoolName string) (result sql.DatabaseListResult, err error)
	ListByServer(ctx context.Context, resourceGroupName string, serverName string, expand string, filter string) (result sql.DatabaseListResult, err error)
	ListMetricDefinitions(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result sql.MetricDefinitionListResult, err error)
	ListMetrics(ctx context.Context, resourceGroupName string, serverName string, databaseName string, filter string) (result sql.MetricListResult, err error)
	Pause(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result sql.DatabasesPauseFuture, err error)
	Resume(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result sql.DatabasesResumeFuture, err error)
	Update(ctx context.Context, resourceGroupName string, serverName string, databaseName string, parameters sql.DatabaseUpdate) (result sql.DatabasesUpdateFuture, err error)
}

var _ DatabasesClientAPI = (*sql.DatabasesClient)(nil)

// DatabaseThreatDetectionPoliciesClientAPI contains the set of methods on the DatabaseThreatDetectionPoliciesClient type.
type DatabaseThreatDetectionPoliciesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, databaseName string, parameters sql.DatabaseSecurityAlertPolicy) (result sql.DatabaseSecurityAlertPolicy, err error)
	Get(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result sql.DatabaseSecurityAlertPolicy, err error)
}

var _ DatabaseThreatDetectionPoliciesClientAPI = (*sql.DatabaseThreatDetectionPoliciesClient)(nil)

// DataMaskingPoliciesClientAPI contains the set of methods on the DataMaskingPoliciesClient type.
type DataMaskingPoliciesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, databaseName string, parameters sql.DataMaskingPolicy) (result sql.DataMaskingPolicy, err error)
	Get(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result sql.DataMaskingPolicy, err error)
}

var _ DataMaskingPoliciesClientAPI = (*sql.DataMaskingPoliciesClient)(nil)

// DataMaskingRulesClientAPI contains the set of methods on the DataMaskingRulesClient type.
type DataMaskingRulesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, databaseName string, dataMaskingRuleName string, parameters sql.DataMaskingRule) (result sql.DataMaskingRule, err error)
	ListByDatabase(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result sql.DataMaskingRuleListResult, err error)
}

var _ DataMaskingRulesClientAPI = (*sql.DataMaskingRulesClient)(nil)

// TransparentDataEncryptionConfigurationsClientAPI contains the set of methods on the TransparentDataEncryptionConfigurationsClient type.
type TransparentDataEncryptionConfigurationsClientAPI interface {
	ListByDatabase(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result sql.TransparentDataEncryptionListResult, err error)
}

var _ TransparentDataEncryptionConfigurationsClientAPI = (*sql.TransparentDataEncryptionConfigurationsClient)(nil)

// ExtensionsClientAPI contains the set of methods on the ExtensionsClient type.
type ExtensionsClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result autorest.Response, err error)
	ListByDatabase(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result sql.ExtensionListResult, err error)
}

var _ ExtensionsClientAPI = (*sql.ExtensionsClient)(nil)

// DisasterRecoveryConfigurationsClientAPI contains the set of methods on the DisasterRecoveryConfigurationsClient type.
type DisasterRecoveryConfigurationsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, disasterRecoveryConfigurationName string) (result sql.DisasterRecoveryConfigurationsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, serverName string, disasterRecoveryConfigurationName string) (result sql.DisasterRecoveryConfigurationsDeleteFuture, err error)
	Failover(ctx context.Context, resourceGroupName string, serverName string, disasterRecoveryConfigurationName string) (result sql.DisasterRecoveryConfigurationsFailoverFuture, err error)
	FailoverAllowDataLoss(ctx context.Context, resourceGroupName string, serverName string, disasterRecoveryConfigurationName string) (result sql.DisasterRecoveryConfigurationsFailoverAllowDataLossFuture, err error)
	Get(ctx context.Context, resourceGroupName string, serverName string, disasterRecoveryConfigurationName string) (result sql.DisasterRecoveryConfiguration, err error)
	List(ctx context.Context, resourceGroupName string, serverName string) (result sql.DisasterRecoveryConfigurationListResult, err error)
}

var _ DisasterRecoveryConfigurationsClientAPI = (*sql.DisasterRecoveryConfigurationsClient)(nil)

// ElasticPoolsClientAPI contains the set of methods on the ElasticPoolsClient type.
type ElasticPoolsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string, parameters sql.ElasticPool) (result sql.ElasticPoolsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string) (result sql.ElasticPool, err error)
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result sql.ElasticPoolListResult, err error)
	ListMetricDefinitions(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string) (result sql.MetricDefinitionListResult, err error)
	ListMetrics(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string, filter string) (result sql.MetricListResult, err error)
	Update(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string, parameters sql.ElasticPoolUpdate) (result sql.ElasticPoolsUpdateFuture, err error)
}

var _ ElasticPoolsClientAPI = (*sql.ElasticPoolsClient)(nil)

// FirewallRulesClientAPI contains the set of methods on the FirewallRulesClient type.
type FirewallRulesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, firewallRuleName string, parameters sql.FirewallRule) (result sql.FirewallRule, err error)
	Delete(ctx context.Context, resourceGroupName string, serverName string, firewallRuleName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, serverName string, firewallRuleName string) (result sql.FirewallRule, err error)
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result sql.FirewallRuleListResult, err error)
}

var _ FirewallRulesClientAPI = (*sql.FirewallRulesClient)(nil)

// GeoBackupPoliciesClientAPI contains the set of methods on the GeoBackupPoliciesClient type.
type GeoBackupPoliciesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, databaseName string, parameters sql.GeoBackupPolicy) (result sql.GeoBackupPolicy, err error)
	Get(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result sql.GeoBackupPolicy, err error)
	ListByDatabase(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result sql.GeoBackupPolicyListResult, err error)
}

var _ GeoBackupPoliciesClientAPI = (*sql.GeoBackupPoliciesClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result sql.OperationListResult, err error)
}

var _ OperationsClientAPI = (*sql.OperationsClient)(nil)

// QueriesClientAPI contains the set of methods on the QueriesClient type.
type QueriesClientAPI interface {
	ListByDatabase(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result sql.TopQueriesListResult, err error)
}

var _ QueriesClientAPI = (*sql.QueriesClient)(nil)

// QueryStatisticsClientAPI contains the set of methods on the QueryStatisticsClient type.
type QueryStatisticsClientAPI interface {
	ListByQuery(ctx context.Context, resourceGroupName string, serverName string, databaseName string, queryID string) (result sql.QueryStatisticListResult, err error)
}

var _ QueryStatisticsClientAPI = (*sql.QueryStatisticsClient)(nil)

// QueryTextsClientAPI contains the set of methods on the QueryTextsClient type.
type QueryTextsClientAPI interface {
	ListByQuery(ctx context.Context, resourceGroupName string, serverName string, databaseName string, queryID string) (result sql.QueryTextListResult, err error)
}

var _ QueryTextsClientAPI = (*sql.QueryTextsClient)(nil)

// RecommendedElasticPoolsClientAPI contains the set of methods on the RecommendedElasticPoolsClient type.
type RecommendedElasticPoolsClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, serverName string, recommendedElasticPoolName string) (result sql.RecommendedElasticPool, err error)
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result sql.RecommendedElasticPoolListResult, err error)
	ListMetrics(ctx context.Context, resourceGroupName string, serverName string, recommendedElasticPoolName string) (result sql.RecommendedElasticPoolListMetricsResult, err error)
}

var _ RecommendedElasticPoolsClientAPI = (*sql.RecommendedElasticPoolsClient)(nil)

// ReplicationLinksClientAPI contains the set of methods on the ReplicationLinksClient type.
type ReplicationLinksClientAPI interface {
	Delete(ctx context.Context, resourceGroupName string, serverName string, databaseName string, linkID string) (result autorest.Response, err error)
	Failover(ctx context.Context, resourceGroupName string, serverName string, databaseName string, linkID string) (result sql.ReplicationLinksFailoverFuture, err error)
	FailoverAllowDataLoss(ctx context.Context, resourceGroupName string, serverName string, databaseName string, linkID string) (result sql.ReplicationLinksFailoverAllowDataLossFuture, err error)
	Get(ctx context.Context, resourceGroupName string, serverName string, databaseName string, linkID string) (result sql.ReplicationLink, err error)
	ListByDatabase(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result sql.ReplicationLinkListResult, err error)
	Unlink(ctx context.Context, resourceGroupName string, serverName string, databaseName string, linkID string, parameters sql.UnlinkParameters) (result sql.ReplicationLinksUnlinkFuture, err error)
}

var _ ReplicationLinksClientAPI = (*sql.ReplicationLinksClient)(nil)

// RestorePointsClientAPI contains the set of methods on the RestorePointsClient type.
type RestorePointsClientAPI interface {
	ListByDatabase(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result sql.RestorePointListResult, err error)
}

var _ RestorePointsClientAPI = (*sql.RestorePointsClient)(nil)

// ServerAzureADAdministratorsClientAPI contains the set of methods on the ServerAzureADAdministratorsClient type.
type ServerAzureADAdministratorsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, properties sql.ServerAzureADAdministrator) (result sql.ServerAzureADAdministratorsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, serverName string) (result sql.ServerAzureADAdministratorsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, serverName string) (result sql.ServerAzureADAdministrator, err error)
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result sql.ServerAdministratorListResult, err error)
}

var _ ServerAzureADAdministratorsClientAPI = (*sql.ServerAzureADAdministratorsClient)(nil)

// ServerCommunicationLinksClientAPI contains the set of methods on the ServerCommunicationLinksClient type.
type ServerCommunicationLinksClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, communicationLinkName string, parameters sql.ServerCommunicationLink) (result sql.ServerCommunicationLinksCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, serverName string, communicationLinkName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, serverName string, communicationLinkName string) (result sql.ServerCommunicationLink, err error)
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result sql.ServerCommunicationLinkListResult, err error)
}

var _ ServerCommunicationLinksClientAPI = (*sql.ServerCommunicationLinksClient)(nil)

// ServiceObjectivesClientAPI contains the set of methods on the ServiceObjectivesClient type.
type ServiceObjectivesClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, serverName string, serviceObjectiveName string) (result sql.ServiceObjective, err error)
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result sql.ServiceObjectiveListResult, err error)
}

var _ ServiceObjectivesClientAPI = (*sql.ServiceObjectivesClient)(nil)

// ElasticPoolActivitiesClientAPI contains the set of methods on the ElasticPoolActivitiesClient type.
type ElasticPoolActivitiesClientAPI interface {
	ListByElasticPool(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string) (result sql.ElasticPoolActivityListResult, err error)
}

var _ ElasticPoolActivitiesClientAPI = (*sql.ElasticPoolActivitiesClient)(nil)

// ElasticPoolDatabaseActivitiesClientAPI contains the set of methods on the ElasticPoolDatabaseActivitiesClient type.
type ElasticPoolDatabaseActivitiesClientAPI interface {
	ListByElasticPool(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string) (result sql.ElasticPoolDatabaseActivityListResult, err error)
}

var _ ElasticPoolDatabaseActivitiesClientAPI = (*sql.ElasticPoolDatabaseActivitiesClient)(nil)

// ServiceTierAdvisorsClientAPI contains the set of methods on the ServiceTierAdvisorsClient type.
type ServiceTierAdvisorsClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, serverName string, databaseName string, serviceTierAdvisorName string) (result sql.ServiceTierAdvisor, err error)
	ListByDatabase(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result sql.ServiceTierAdvisorListResult, err error)
}

var _ ServiceTierAdvisorsClientAPI = (*sql.ServiceTierAdvisorsClient)(nil)

// TransparentDataEncryptionsClientAPI contains the set of methods on the TransparentDataEncryptionsClient type.
type TransparentDataEncryptionsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, databaseName string, parameters sql.TransparentDataEncryption) (result sql.TransparentDataEncryption, err error)
	Get(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result sql.TransparentDataEncryption, err error)
}

var _ TransparentDataEncryptionsClientAPI = (*sql.TransparentDataEncryptionsClient)(nil)

// TransparentDataEncryptionActivitiesClientAPI contains the set of methods on the TransparentDataEncryptionActivitiesClient type.
type TransparentDataEncryptionActivitiesClientAPI interface {
	ListByConfiguration(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result sql.TransparentDataEncryptionActivityListResult, err error)
}

var _ TransparentDataEncryptionActivitiesClientAPI = (*sql.TransparentDataEncryptionActivitiesClient)(nil)

// ServerTableAuditingPoliciesClientAPI contains the set of methods on the ServerTableAuditingPoliciesClient type.
type ServerTableAuditingPoliciesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, parameters sql.ServerTableAuditingPolicy) (result sql.ServerTableAuditingPolicy, err error)
	Get(ctx context.Context, resourceGroupName string, serverName string) (result sql.ServerTableAuditingPolicy, err error)
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result sql.ServerTableAuditingPolicyListResult, err error)
}

var _ ServerTableAuditingPoliciesClientAPI = (*sql.ServerTableAuditingPoliciesClient)(nil)

// DatabaseTableAuditingPoliciesClientAPI contains the set of methods on the DatabaseTableAuditingPoliciesClient type.
type DatabaseTableAuditingPoliciesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, databaseName string, parameters sql.DatabaseTableAuditingPolicy) (result sql.DatabaseTableAuditingPolicy, err error)
	Get(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result sql.DatabaseTableAuditingPolicy, err error)
	ListByDatabase(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result sql.DatabaseTableAuditingPolicyListResult, err error)
}

var _ DatabaseTableAuditingPoliciesClientAPI = (*sql.DatabaseTableAuditingPoliciesClient)(nil)

// DatabaseConnectionPoliciesClientAPI contains the set of methods on the DatabaseConnectionPoliciesClient type.
type DatabaseConnectionPoliciesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, databaseName string, parameters sql.DatabaseConnectionPolicy) (result sql.DatabaseConnectionPolicy, err error)
	Get(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result sql.DatabaseConnectionPolicy, err error)
}

var _ DatabaseConnectionPoliciesClientAPI = (*sql.DatabaseConnectionPoliciesClient)(nil)

// ServerUsagesClientAPI contains the set of methods on the ServerUsagesClient type.
type ServerUsagesClientAPI interface {
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result sql.ServerUsageListResult, err error)
}

var _ ServerUsagesClientAPI = (*sql.ServerUsagesClient)(nil)

// DatabaseUsagesClientAPI contains the set of methods on the DatabaseUsagesClient type.
type DatabaseUsagesClientAPI interface {
	ListByDatabase(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result sql.DatabaseUsageListResult, err error)
}

var _ DatabaseUsagesClientAPI = (*sql.DatabaseUsagesClient)(nil)
