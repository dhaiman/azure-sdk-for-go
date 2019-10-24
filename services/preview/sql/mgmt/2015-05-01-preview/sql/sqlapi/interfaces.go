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
    "github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql"
    "github.com/Azure/go-autorest/autorest"
)

        // BackupLongTermRetentionPoliciesClientAPI contains the set of methods on the BackupLongTermRetentionPoliciesClient type.
        type BackupLongTermRetentionPoliciesClientAPI interface {
            CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, databaseName string, parameters sql.BackupLongTermRetentionPolicy) (result sql.BackupLongTermRetentionPoliciesCreateOrUpdateFuture, err error)
            Get(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result sql.BackupLongTermRetentionPolicy, err error)
            ListByDatabase(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result sql.BackupLongTermRetentionPolicyListResult, err error)
        }

        var _ BackupLongTermRetentionPoliciesClientAPI = (*sql.BackupLongTermRetentionPoliciesClient)(nil)
        // BackupLongTermRetentionVaultsClientAPI contains the set of methods on the BackupLongTermRetentionVaultsClient type.
        type BackupLongTermRetentionVaultsClientAPI interface {
            CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, parameters sql.BackupLongTermRetentionVault) (result sql.BackupLongTermRetentionVaultsCreateOrUpdateFuture, err error)
            Get(ctx context.Context, resourceGroupName string, serverName string) (result sql.BackupLongTermRetentionVault, err error)
            ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result sql.BackupLongTermRetentionVaultListResult, err error)
        }

        var _ BackupLongTermRetentionVaultsClientAPI = (*sql.BackupLongTermRetentionVaultsClient)(nil)
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
        // RestorePointsClientAPI contains the set of methods on the RestorePointsClient type.
        type RestorePointsClientAPI interface {
            ListByDatabase(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result sql.RestorePointListResult, err error)
        }

        var _ RestorePointsClientAPI = (*sql.RestorePointsClient)(nil)
        // ServersClientAPI contains the set of methods on the ServersClient type.
        type ServersClientAPI interface {
            CheckNameAvailability(ctx context.Context, parameters sql.CheckNameAvailabilityRequest) (result sql.CheckNameAvailabilityResponse, err error)
            CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, parameters sql.Server) (result sql.ServersCreateOrUpdateFuture, err error)
            Delete(ctx context.Context, resourceGroupName string, serverName string) (result sql.ServersDeleteFuture, err error)
            Get(ctx context.Context, resourceGroupName string, serverName string) (result sql.Server, err error)
            List(ctx context.Context) (result sql.ServerListResultPage, err error)
            ListByResourceGroup(ctx context.Context, resourceGroupName string) (result sql.ServerListResultPage, err error)
            Update(ctx context.Context, resourceGroupName string, serverName string, parameters sql.ServerUpdate) (result sql.ServersUpdateFuture, err error)
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
        // ReplicationLinksClientAPI contains the set of methods on the ReplicationLinksClient type.
        type ReplicationLinksClientAPI interface {
            Delete(ctx context.Context, resourceGroupName string, serverName string, databaseName string, linkID string) (result autorest.Response, err error)
            Failover(ctx context.Context, resourceGroupName string, serverName string, databaseName string, linkID string) (result sql.ReplicationLinksFailoverFuture, err error)
            FailoverAllowDataLoss(ctx context.Context, resourceGroupName string, serverName string, databaseName string, linkID string) (result sql.ReplicationLinksFailoverAllowDataLossFuture, err error)
            Get(ctx context.Context, resourceGroupName string, serverName string, databaseName string, linkID string) (result sql.ReplicationLink, err error)
            ListByDatabase(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result sql.ReplicationLinkListResult, err error)
        }

        var _ ReplicationLinksClientAPI = (*sql.ReplicationLinksClient)(nil)
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
        // CapabilitiesClientAPI contains the set of methods on the CapabilitiesClient type.
        type CapabilitiesClientAPI interface {
            ListByLocation(ctx context.Context, locationName string) (result sql.LocationCapabilities, err error)
        }

        var _ CapabilitiesClientAPI = (*sql.CapabilitiesClient)(nil)
        // DatabaseBlobAuditingPoliciesClientAPI contains the set of methods on the DatabaseBlobAuditingPoliciesClient type.
        type DatabaseBlobAuditingPoliciesClientAPI interface {
            CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, databaseName string, parameters sql.DatabaseBlobAuditingPolicy) (result sql.DatabaseBlobAuditingPolicy, err error)
            Get(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result sql.DatabaseBlobAuditingPolicy, err error)
            ListByDatabase(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result sql.DatabaseBlobAuditingPolicyListResultPage, err error)
        }

        var _ DatabaseBlobAuditingPoliciesClientAPI = (*sql.DatabaseBlobAuditingPoliciesClient)(nil)
        // EncryptionProtectorsClientAPI contains the set of methods on the EncryptionProtectorsClient type.
        type EncryptionProtectorsClientAPI interface {
            CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, parameters sql.EncryptionProtector) (result sql.EncryptionProtectorsCreateOrUpdateFuture, err error)
            Get(ctx context.Context, resourceGroupName string, serverName string) (result sql.EncryptionProtector, err error)
            ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result sql.EncryptionProtectorListResultPage, err error)
            Revalidate(ctx context.Context, resourceGroupName string, serverName string) (result sql.EncryptionProtectorsRevalidateFuture, err error)
        }

        var _ EncryptionProtectorsClientAPI = (*sql.EncryptionProtectorsClient)(nil)
        // FailoverGroupsClientAPI contains the set of methods on the FailoverGroupsClient type.
        type FailoverGroupsClientAPI interface {
            CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, failoverGroupName string, parameters sql.FailoverGroup) (result sql.FailoverGroupsCreateOrUpdateFuture, err error)
            Delete(ctx context.Context, resourceGroupName string, serverName string, failoverGroupName string) (result sql.FailoverGroupsDeleteFuture, err error)
            Failover(ctx context.Context, resourceGroupName string, serverName string, failoverGroupName string) (result sql.FailoverGroupsFailoverFuture, err error)
            ForceFailoverAllowDataLoss(ctx context.Context, resourceGroupName string, serverName string, failoverGroupName string) (result sql.FailoverGroupsForceFailoverAllowDataLossFuture, err error)
            Get(ctx context.Context, resourceGroupName string, serverName string, failoverGroupName string) (result sql.FailoverGroup, err error)
            ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result sql.FailoverGroupListResultPage, err error)
            Update(ctx context.Context, resourceGroupName string, serverName string, failoverGroupName string, parameters sql.FailoverGroupUpdate) (result sql.FailoverGroupsUpdateFuture, err error)
        }

        var _ FailoverGroupsClientAPI = (*sql.FailoverGroupsClient)(nil)
        // ManagedInstancesClientAPI contains the set of methods on the ManagedInstancesClient type.
        type ManagedInstancesClientAPI interface {
            CreateOrUpdate(ctx context.Context, resourceGroupName string, managedInstanceName string, parameters sql.ManagedInstance) (result sql.ManagedInstancesCreateOrUpdateFuture, err error)
            Delete(ctx context.Context, resourceGroupName string, managedInstanceName string) (result sql.ManagedInstancesDeleteFuture, err error)
            Get(ctx context.Context, resourceGroupName string, managedInstanceName string) (result sql.ManagedInstance, err error)
            List(ctx context.Context) (result sql.ManagedInstanceListResultPage, err error)
            ListByResourceGroup(ctx context.Context, resourceGroupName string) (result sql.ManagedInstanceListResultPage, err error)
            Update(ctx context.Context, resourceGroupName string, managedInstanceName string, parameters sql.ManagedInstanceUpdate) (result sql.ManagedInstancesUpdateFuture, err error)
        }

        var _ ManagedInstancesClientAPI = (*sql.ManagedInstancesClient)(nil)
        // OperationsClientAPI contains the set of methods on the OperationsClient type.
        type OperationsClientAPI interface {
            List(ctx context.Context) (result sql.OperationListResultPage, err error)
        }

        var _ OperationsClientAPI = (*sql.OperationsClient)(nil)
        // ServerKeysClientAPI contains the set of methods on the ServerKeysClient type.
        type ServerKeysClientAPI interface {
            CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, keyName string, parameters sql.ServerKey) (result sql.ServerKeysCreateOrUpdateFuture, err error)
            Delete(ctx context.Context, resourceGroupName string, serverName string, keyName string) (result sql.ServerKeysDeleteFuture, err error)
            Get(ctx context.Context, resourceGroupName string, serverName string, keyName string) (result sql.ServerKey, err error)
            ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result sql.ServerKeyListResultPage, err error)
        }

        var _ ServerKeysClientAPI = (*sql.ServerKeysClient)(nil)
        // SyncAgentsClientAPI contains the set of methods on the SyncAgentsClient type.
        type SyncAgentsClientAPI interface {
            CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, syncAgentName string, parameters sql.SyncAgent) (result sql.SyncAgentsCreateOrUpdateFuture, err error)
            Delete(ctx context.Context, resourceGroupName string, serverName string, syncAgentName string) (result sql.SyncAgentsDeleteFuture, err error)
            GenerateKey(ctx context.Context, resourceGroupName string, serverName string, syncAgentName string) (result sql.SyncAgentKeyProperties, err error)
            Get(ctx context.Context, resourceGroupName string, serverName string, syncAgentName string) (result sql.SyncAgent, err error)
            ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result sql.SyncAgentListResultPage, err error)
            ListLinkedDatabases(ctx context.Context, resourceGroupName string, serverName string, syncAgentName string) (result sql.SyncAgentLinkedDatabaseListResultPage, err error)
        }

        var _ SyncAgentsClientAPI = (*sql.SyncAgentsClient)(nil)
        // SyncGroupsClientAPI contains the set of methods on the SyncGroupsClient type.
        type SyncGroupsClientAPI interface {
            CancelSync(ctx context.Context, resourceGroupName string, serverName string, databaseName string, syncGroupName string) (result autorest.Response, err error)
            CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, databaseName string, syncGroupName string, parameters sql.SyncGroup) (result sql.SyncGroupsCreateOrUpdateFuture, err error)
            Delete(ctx context.Context, resourceGroupName string, serverName string, databaseName string, syncGroupName string) (result sql.SyncGroupsDeleteFuture, err error)
            Get(ctx context.Context, resourceGroupName string, serverName string, databaseName string, syncGroupName string) (result sql.SyncGroup, err error)
            ListByDatabase(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result sql.SyncGroupListResultPage, err error)
            ListHubSchemas(ctx context.Context, resourceGroupName string, serverName string, databaseName string, syncGroupName string) (result sql.SyncFullSchemaPropertiesListResultPage, err error)
            ListLogs(ctx context.Context, resourceGroupName string, serverName string, databaseName string, syncGroupName string, startTime string, endTime string, typeParameter string, continuationToken string) (result sql.SyncGroupLogListResultPage, err error)
            ListSyncDatabaseIds(ctx context.Context, locationName string) (result sql.SyncDatabaseIDListResultPage, err error)
            RefreshHubSchema(ctx context.Context, resourceGroupName string, serverName string, databaseName string, syncGroupName string) (result sql.SyncGroupsRefreshHubSchemaFuture, err error)
            TriggerSync(ctx context.Context, resourceGroupName string, serverName string, databaseName string, syncGroupName string) (result autorest.Response, err error)
            Update(ctx context.Context, resourceGroupName string, serverName string, databaseName string, syncGroupName string, parameters sql.SyncGroup) (result sql.SyncGroupsUpdateFuture, err error)
        }

        var _ SyncGroupsClientAPI = (*sql.SyncGroupsClient)(nil)
        // SyncMembersClientAPI contains the set of methods on the SyncMembersClient type.
        type SyncMembersClientAPI interface {
            CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, databaseName string, syncGroupName string, syncMemberName string, parameters sql.SyncMember) (result sql.SyncMembersCreateOrUpdateFuture, err error)
            Delete(ctx context.Context, resourceGroupName string, serverName string, databaseName string, syncGroupName string, syncMemberName string) (result sql.SyncMembersDeleteFuture, err error)
            Get(ctx context.Context, resourceGroupName string, serverName string, databaseName string, syncGroupName string, syncMemberName string) (result sql.SyncMember, err error)
            ListBySyncGroup(ctx context.Context, resourceGroupName string, serverName string, databaseName string, syncGroupName string) (result sql.SyncMemberListResultPage, err error)
            ListMemberSchemas(ctx context.Context, resourceGroupName string, serverName string, databaseName string, syncGroupName string, syncMemberName string) (result sql.SyncFullSchemaPropertiesListResultPage, err error)
            RefreshMemberSchema(ctx context.Context, resourceGroupName string, serverName string, databaseName string, syncGroupName string, syncMemberName string) (result sql.SyncMembersRefreshMemberSchemaFuture, err error)
            Update(ctx context.Context, resourceGroupName string, serverName string, databaseName string, syncGroupName string, syncMemberName string, parameters sql.SyncMember) (result sql.SyncMembersUpdateFuture, err error)
        }

        var _ SyncMembersClientAPI = (*sql.SyncMembersClient)(nil)
        // SubscriptionUsagesClientAPI contains the set of methods on the SubscriptionUsagesClient type.
        type SubscriptionUsagesClientAPI interface {
            Get(ctx context.Context, locationName string, usageName string) (result sql.SubscriptionUsage, err error)
            ListByLocation(ctx context.Context, locationName string) (result sql.SubscriptionUsageListResultPage, err error)
        }

        var _ SubscriptionUsagesClientAPI = (*sql.SubscriptionUsagesClient)(nil)
        // VirtualClustersClientAPI contains the set of methods on the VirtualClustersClient type.
        type VirtualClustersClientAPI interface {
            Delete(ctx context.Context, resourceGroupName string, virtualClusterName string) (result sql.VirtualClustersDeleteFuture, err error)
            Get(ctx context.Context, resourceGroupName string, virtualClusterName string) (result sql.VirtualCluster, err error)
            List(ctx context.Context) (result sql.VirtualClusterListResultPage, err error)
            ListByResourceGroup(ctx context.Context, resourceGroupName string) (result sql.VirtualClusterListResultPage, err error)
            Update(ctx context.Context, resourceGroupName string, virtualClusterName string, parameters sql.VirtualClusterUpdate) (result sql.VirtualClustersUpdateFuture, err error)
        }

        var _ VirtualClustersClientAPI = (*sql.VirtualClustersClient)(nil)
        // VirtualNetworkRulesClientAPI contains the set of methods on the VirtualNetworkRulesClient type.
        type VirtualNetworkRulesClientAPI interface {
            CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, virtualNetworkRuleName string, parameters sql.VirtualNetworkRule) (result sql.VirtualNetworkRulesCreateOrUpdateFuture, err error)
            Delete(ctx context.Context, resourceGroupName string, serverName string, virtualNetworkRuleName string) (result sql.VirtualNetworkRulesDeleteFuture, err error)
            Get(ctx context.Context, resourceGroupName string, serverName string, virtualNetworkRuleName string) (result sql.VirtualNetworkRule, err error)
            ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result sql.VirtualNetworkRuleListResultPage, err error)
        }

        var _ VirtualNetworkRulesClientAPI = (*sql.VirtualNetworkRulesClient)(nil)
