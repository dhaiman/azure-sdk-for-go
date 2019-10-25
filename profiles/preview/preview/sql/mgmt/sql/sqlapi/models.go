// +build go1.9

// Copyright 2019 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package sqlapi

import original "github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2018-06-15-preview/sql/sqlapi"

type BackupLongTermRetentionPoliciesClientAPI = original.BackupLongTermRetentionPoliciesClientAPI
type BackupShortTermRetentionPoliciesClientAPI = original.BackupShortTermRetentionPoliciesClientAPI
type CapabilitiesClientAPI = original.CapabilitiesClientAPI
type DataMaskingPoliciesClientAPI = original.DataMaskingPoliciesClientAPI
type DataMaskingRulesClientAPI = original.DataMaskingRulesClientAPI
type DatabaseAutomaticTuningClientAPI = original.DatabaseAutomaticTuningClientAPI
type DatabaseBlobAuditingPoliciesClientAPI = original.DatabaseBlobAuditingPoliciesClientAPI
type DatabaseOperationsClientAPI = original.DatabaseOperationsClientAPI
type DatabaseThreatDetectionPoliciesClientAPI = original.DatabaseThreatDetectionPoliciesClientAPI
type DatabaseUsagesClientAPI = original.DatabaseUsagesClientAPI
type DatabaseVulnerabilityAssessmentRuleBaselinesClientAPI = original.DatabaseVulnerabilityAssessmentRuleBaselinesClientAPI
type DatabaseVulnerabilityAssessmentScansClientAPI = original.DatabaseVulnerabilityAssessmentScansClientAPI
type DatabaseVulnerabilityAssessmentsClientAPI = original.DatabaseVulnerabilityAssessmentsClientAPI
type DatabasesClientAPI = original.DatabasesClientAPI
type ElasticPoolActivitiesClientAPI = original.ElasticPoolActivitiesClientAPI
type ElasticPoolDatabaseActivitiesClientAPI = original.ElasticPoolDatabaseActivitiesClientAPI
type ElasticPoolOperationsClientAPI = original.ElasticPoolOperationsClientAPI
type ElasticPoolsClientAPI = original.ElasticPoolsClientAPI
type EncryptionProtectorsClientAPI = original.EncryptionProtectorsClientAPI
type ExtendedDatabaseBlobAuditingPoliciesClientAPI = original.ExtendedDatabaseBlobAuditingPoliciesClientAPI
type ExtendedServerBlobAuditingPoliciesClientAPI = original.ExtendedServerBlobAuditingPoliciesClientAPI
type FailoverGroupsClientAPI = original.FailoverGroupsClientAPI
type FirewallRulesClientAPI = original.FirewallRulesClientAPI
type GeoBackupPoliciesClientAPI = original.GeoBackupPoliciesClientAPI
type InstanceFailoverGroupsClientAPI = original.InstanceFailoverGroupsClientAPI
type InstancePoolsClientAPI = original.InstancePoolsClientAPI
type JobAgentsClientAPI = original.JobAgentsClientAPI
type JobCredentialsClientAPI = original.JobCredentialsClientAPI
type JobExecutionsClientAPI = original.JobExecutionsClientAPI
type JobStepExecutionsClientAPI = original.JobStepExecutionsClientAPI
type JobStepsClientAPI = original.JobStepsClientAPI
type JobTargetExecutionsClientAPI = original.JobTargetExecutionsClientAPI
type JobTargetGroupsClientAPI = original.JobTargetGroupsClientAPI
type JobVersionsClientAPI = original.JobVersionsClientAPI
type JobsClientAPI = original.JobsClientAPI
type LongTermRetentionBackupsClientAPI = original.LongTermRetentionBackupsClientAPI
type ManagedBackupShortTermRetentionPoliciesClientAPI = original.ManagedBackupShortTermRetentionPoliciesClientAPI
type ManagedDatabaseRestoreDetailsClientAPI = original.ManagedDatabaseRestoreDetailsClientAPI
type ManagedDatabaseSecurityAlertPoliciesClientAPI = original.ManagedDatabaseSecurityAlertPoliciesClientAPI
type ManagedDatabaseSensitivityLabelsClientAPI = original.ManagedDatabaseSensitivityLabelsClientAPI
type ManagedDatabaseVulnerabilityAssessmentRuleBaselinesClientAPI = original.ManagedDatabaseVulnerabilityAssessmentRuleBaselinesClientAPI
type ManagedDatabaseVulnerabilityAssessmentScansClientAPI = original.ManagedDatabaseVulnerabilityAssessmentScansClientAPI
type ManagedDatabaseVulnerabilityAssessmentsClientAPI = original.ManagedDatabaseVulnerabilityAssessmentsClientAPI
type ManagedDatabasesClientAPI = original.ManagedDatabasesClientAPI
type ManagedInstanceAdministratorsClientAPI = original.ManagedInstanceAdministratorsClientAPI
type ManagedInstanceEncryptionProtectorsClientAPI = original.ManagedInstanceEncryptionProtectorsClientAPI
type ManagedInstanceKeysClientAPI = original.ManagedInstanceKeysClientAPI
type ManagedInstanceTdeCertificatesClientAPI = original.ManagedInstanceTdeCertificatesClientAPI
type ManagedInstanceVulnerabilityAssessmentsClientAPI = original.ManagedInstanceVulnerabilityAssessmentsClientAPI
type ManagedInstancesClientAPI = original.ManagedInstancesClientAPI
type ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClientAPI = original.ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClientAPI
type ManagedServerSecurityAlertPoliciesClientAPI = original.ManagedServerSecurityAlertPoliciesClientAPI
type OperationsClientAPI = original.OperationsClientAPI
type PrivateEndpointConnectionsClientAPI = original.PrivateEndpointConnectionsClientAPI
type PrivateLinkResourcesClientAPI = original.PrivateLinkResourcesClientAPI
type RecommendedElasticPoolsClientAPI = original.RecommendedElasticPoolsClientAPI
type RecoverableDatabasesClientAPI = original.RecoverableDatabasesClientAPI
type RecoverableManagedDatabasesClientAPI = original.RecoverableManagedDatabasesClientAPI
type ReplicationLinksClientAPI = original.ReplicationLinksClientAPI
type RestorableDroppedDatabasesClientAPI = original.RestorableDroppedDatabasesClientAPI
type RestorableDroppedManagedDatabasesClientAPI = original.RestorableDroppedManagedDatabasesClientAPI
type RestorePointsClientAPI = original.RestorePointsClientAPI
type SensitivityLabelsClientAPI = original.SensitivityLabelsClientAPI
type ServerAutomaticTuningClientAPI = original.ServerAutomaticTuningClientAPI
type ServerAzureADAdministratorsClientAPI = original.ServerAzureADAdministratorsClientAPI
type ServerBlobAuditingPoliciesClientAPI = original.ServerBlobAuditingPoliciesClientAPI
type ServerCommunicationLinksClientAPI = original.ServerCommunicationLinksClientAPI
type ServerConnectionPoliciesClientAPI = original.ServerConnectionPoliciesClientAPI
type ServerDNSAliasesClientAPI = original.ServerDNSAliasesClientAPI
type ServerKeysClientAPI = original.ServerKeysClientAPI
type ServerSecurityAlertPoliciesClientAPI = original.ServerSecurityAlertPoliciesClientAPI
type ServerUsagesClientAPI = original.ServerUsagesClientAPI
type ServerVulnerabilityAssessmentsClientAPI = original.ServerVulnerabilityAssessmentsClientAPI
type ServersClientAPI = original.ServersClientAPI
type ServiceObjectivesClientAPI = original.ServiceObjectivesClientAPI
type ServiceTierAdvisorsClientAPI = original.ServiceTierAdvisorsClientAPI
type SubscriptionUsagesClientAPI = original.SubscriptionUsagesClientAPI
type SyncAgentsClientAPI = original.SyncAgentsClientAPI
type SyncGroupsClientAPI = original.SyncGroupsClientAPI
type SyncMembersClientAPI = original.SyncMembersClientAPI
type TdeCertificatesClientAPI = original.TdeCertificatesClientAPI
type TransparentDataEncryptionActivitiesClientAPI = original.TransparentDataEncryptionActivitiesClientAPI
type TransparentDataEncryptionsClientAPI = original.TransparentDataEncryptionsClientAPI
type UsagesClientAPI = original.UsagesClientAPI
type VirtualClustersClientAPI = original.VirtualClustersClientAPI
type VirtualNetworkRulesClientAPI = original.VirtualNetworkRulesClientAPI
type WorkloadClassifiersClientAPI = original.WorkloadClassifiersClientAPI
type WorkloadGroupsClientAPI = original.WorkloadGroupsClientAPI
