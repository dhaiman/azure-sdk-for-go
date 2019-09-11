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

package securityinsight

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/preview/securityinsight/mgmt/2017-08-01-preview/securityinsight"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AggregationsKind = original.AggregationsKind

const (
	AggregationsKindCasesAggregation AggregationsKind = original.AggregationsKindCasesAggregation
)

type AlertRuleKind = original.AlertRuleKind

const (
	Filter    AlertRuleKind = original.Filter
	Fusion    AlertRuleKind = original.Fusion
	Scheduled AlertRuleKind = original.Scheduled
)

type AlertSeverity = original.AlertSeverity

const (
	High          AlertSeverity = original.High
	Informational AlertSeverity = original.Informational
	Low           AlertSeverity = original.Low
	Medium        AlertSeverity = original.Medium
)

type AlertStatus = original.AlertStatus

const (
	AlertStatusDismissed  AlertStatus = original.AlertStatusDismissed
	AlertStatusInProgress AlertStatus = original.AlertStatusInProgress
	AlertStatusNew        AlertStatus = original.AlertStatusNew
	AlertStatusResolved   AlertStatus = original.AlertStatusResolved
	AlertStatusUnknown    AlertStatus = original.AlertStatusUnknown
)

type AttackTactic = original.AttackTactic

const (
	Collection          AttackTactic = original.Collection
	CommandAndControl   AttackTactic = original.CommandAndControl
	CredentialAccess    AttackTactic = original.CredentialAccess
	DefenseEvasion      AttackTactic = original.DefenseEvasion
	Discovery           AttackTactic = original.Discovery
	Execution           AttackTactic = original.Execution
	Exfiltration        AttackTactic = original.Exfiltration
	InitialAccess       AttackTactic = original.InitialAccess
	LateralMovement     AttackTactic = original.LateralMovement
	Persistence         AttackTactic = original.Persistence
	PrivilegeEscalation AttackTactic = original.PrivilegeEscalation
)

type CaseSeverity = original.CaseSeverity

const (
	CaseSeverityCritical      CaseSeverity = original.CaseSeverityCritical
	CaseSeverityHigh          CaseSeverity = original.CaseSeverityHigh
	CaseSeverityInformational CaseSeverity = original.CaseSeverityInformational
	CaseSeverityLow           CaseSeverity = original.CaseSeverityLow
	CaseSeverityMedium        CaseSeverity = original.CaseSeverityMedium
)

type CaseStatus = original.CaseStatus

const (
	CaseStatusClosed     CaseStatus = original.CaseStatusClosed
	CaseStatusDraft      CaseStatus = original.CaseStatusDraft
	CaseStatusInProgress CaseStatus = original.CaseStatusInProgress
	CaseStatusNew        CaseStatus = original.CaseStatusNew
)

type CloseReason = original.CloseReason

const (
	Dismissed     CloseReason = original.Dismissed
	FalsePositive CloseReason = original.FalsePositive
	Other         CloseReason = original.Other
	Resolved      CloseReason = original.Resolved
	TruePositive  CloseReason = original.TruePositive
)

type ConfidenceLevel = original.ConfidenceLevel

const (
	ConfidenceLevelHigh    ConfidenceLevel = original.ConfidenceLevelHigh
	ConfidenceLevelLow     ConfidenceLevel = original.ConfidenceLevelLow
	ConfidenceLevelUnknown ConfidenceLevel = original.ConfidenceLevelUnknown
)

type ConfidenceScoreStatus = original.ConfidenceScoreStatus

const (
	Final         ConfidenceScoreStatus = original.Final
	InProcess     ConfidenceScoreStatus = original.InProcess
	NotApplicable ConfidenceScoreStatus = original.NotApplicable
	NotFinal      ConfidenceScoreStatus = original.NotFinal
)

type DataConnectorKind = original.DataConnectorKind

const (
	DataConnectorKindAmazonWebServicesCloudTrail               DataConnectorKind = original.DataConnectorKindAmazonWebServicesCloudTrail
	DataConnectorKindAzureActiveDirectory                      DataConnectorKind = original.DataConnectorKindAzureActiveDirectory
	DataConnectorKindAzureAdvancedThreatProtection             DataConnectorKind = original.DataConnectorKindAzureAdvancedThreatProtection
	DataConnectorKindAzureSecurityCenter                       DataConnectorKind = original.DataConnectorKindAzureSecurityCenter
	DataConnectorKindMicrosoftCloudAppSecurity                 DataConnectorKind = original.DataConnectorKindMicrosoftCloudAppSecurity
	DataConnectorKindMicrosoftDefenderAdvancedThreatProtection DataConnectorKind = original.DataConnectorKindMicrosoftDefenderAdvancedThreatProtection
	DataConnectorKindOffice365                                 DataConnectorKind = original.DataConnectorKindOffice365
	DataConnectorKindThreatIntelligence                        DataConnectorKind = original.DataConnectorKindThreatIntelligence
)

type DataTypeState = original.DataTypeState

const (
	Disabled DataTypeState = original.Disabled
	Enabled  DataTypeState = original.Enabled
)

type DataTypeStatus = original.DataTypeStatus

const (
	Exist    DataTypeStatus = original.Exist
	NotExist DataTypeStatus = original.NotExist
)

type ElevationToken = original.ElevationToken

const (
	Default ElevationToken = original.Default
	Full    ElevationToken = original.Full
	Limited ElevationToken = original.Limited
)

type EntityKind = original.EntityKind

const (
	EntityKindAccount          EntityKind = original.EntityKindAccount
	EntityKindAzureResource    EntityKind = original.EntityKindAzureResource
	EntityKindBookmark         EntityKind = original.EntityKindBookmark
	EntityKindCloudApplication EntityKind = original.EntityKindCloudApplication
	EntityKindDNSResolution    EntityKind = original.EntityKindDNSResolution
	EntityKindFile             EntityKind = original.EntityKindFile
	EntityKindFileHash         EntityKind = original.EntityKindFileHash
	EntityKindHost             EntityKind = original.EntityKindHost
	EntityKindIP               EntityKind = original.EntityKindIP
	EntityKindMalware          EntityKind = original.EntityKindMalware
	EntityKindProcess          EntityKind = original.EntityKindProcess
	EntityKindRegistryKey      EntityKind = original.EntityKindRegistryKey
	EntityKindRegistryValue    EntityKind = original.EntityKindRegistryValue
	EntityKindSecurityAlert    EntityKind = original.EntityKindSecurityAlert
	EntityKindSecurityGroup    EntityKind = original.EntityKindSecurityGroup
	EntityKindURL              EntityKind = original.EntityKindURL
)

type EntityType = original.EntityType

const (
	EntityTypeAccount          EntityType = original.EntityTypeAccount
	EntityTypeAzureResource    EntityType = original.EntityTypeAzureResource
	EntityTypeCloudApplication EntityType = original.EntityTypeCloudApplication
	EntityTypeDNS              EntityType = original.EntityTypeDNS
	EntityTypeFile             EntityType = original.EntityTypeFile
	EntityTypeFileHash         EntityType = original.EntityTypeFileHash
	EntityTypeHost             EntityType = original.EntityTypeHost
	EntityTypeHuntingBookmark  EntityType = original.EntityTypeHuntingBookmark
	EntityTypeIP               EntityType = original.EntityTypeIP
	EntityTypeMalware          EntityType = original.EntityTypeMalware
	EntityTypeProcess          EntityType = original.EntityTypeProcess
	EntityTypeRegistryKey      EntityType = original.EntityTypeRegistryKey
	EntityTypeRegistryValue    EntityType = original.EntityTypeRegistryValue
	EntityTypeSecurityAlert    EntityType = original.EntityTypeSecurityAlert
	EntityTypeSecurityGroup    EntityType = original.EntityTypeSecurityGroup
	EntityTypeURL              EntityType = original.EntityTypeURL
)

type FileHashAlgorithm = original.FileHashAlgorithm

const (
	MD5      FileHashAlgorithm = original.MD5
	SHA1     FileHashAlgorithm = original.SHA1
	SHA256   FileHashAlgorithm = original.SHA256
	SHA256AC FileHashAlgorithm = original.SHA256AC
	Unknown  FileHashAlgorithm = original.Unknown
)

type KillChainIntent = original.KillChainIntent

const (
	KillChainIntentCollection          KillChainIntent = original.KillChainIntentCollection
	KillChainIntentCommandAndControl   KillChainIntent = original.KillChainIntentCommandAndControl
	KillChainIntentCredentialAccess    KillChainIntent = original.KillChainIntentCredentialAccess
	KillChainIntentDefenseEvasion      KillChainIntent = original.KillChainIntentDefenseEvasion
	KillChainIntentDiscovery           KillChainIntent = original.KillChainIntentDiscovery
	KillChainIntentExecution           KillChainIntent = original.KillChainIntentExecution
	KillChainIntentExfiltration        KillChainIntent = original.KillChainIntentExfiltration
	KillChainIntentExploitation        KillChainIntent = original.KillChainIntentExploitation
	KillChainIntentImpact              KillChainIntent = original.KillChainIntentImpact
	KillChainIntentLateralMovement     KillChainIntent = original.KillChainIntentLateralMovement
	KillChainIntentPersistence         KillChainIntent = original.KillChainIntentPersistence
	KillChainIntentPrivilegeEscalation KillChainIntent = original.KillChainIntentPrivilegeEscalation
	KillChainIntentProbing             KillChainIntent = original.KillChainIntentProbing
	KillChainIntentUnknown             KillChainIntent = original.KillChainIntentUnknown
)

type Kind = original.Kind

const (
	KindAggregations     Kind = original.KindAggregations
	KindCasesAggregation Kind = original.KindCasesAggregation
)

type KindBasicAlertRule = original.KindBasicAlertRule

const (
	KindAlertRule KindBasicAlertRule = original.KindAlertRule
	KindScheduled KindBasicAlertRule = original.KindScheduled
)

type KindBasicAlertRuleTemplate = original.KindBasicAlertRuleTemplate

const (
	KindBasicAlertRuleTemplateKindAlertRuleTemplate KindBasicAlertRuleTemplate = original.KindBasicAlertRuleTemplateKindAlertRuleTemplate
	KindBasicAlertRuleTemplateKindFilter            KindBasicAlertRuleTemplate = original.KindBasicAlertRuleTemplateKindFilter
	KindBasicAlertRuleTemplateKindFusion            KindBasicAlertRuleTemplate = original.KindBasicAlertRuleTemplateKindFusion
	KindBasicAlertRuleTemplateKindScheduled         KindBasicAlertRuleTemplate = original.KindBasicAlertRuleTemplateKindScheduled
)

type KindBasicDataConnector = original.KindBasicDataConnector

const (
	KindAmazonWebServicesCloudTrail               KindBasicDataConnector = original.KindAmazonWebServicesCloudTrail
	KindAzureActiveDirectory                      KindBasicDataConnector = original.KindAzureActiveDirectory
	KindAzureAdvancedThreatProtection             KindBasicDataConnector = original.KindAzureAdvancedThreatProtection
	KindAzureSecurityCenter                       KindBasicDataConnector = original.KindAzureSecurityCenter
	KindDataConnector                             KindBasicDataConnector = original.KindDataConnector
	KindMicrosoftCloudAppSecurity                 KindBasicDataConnector = original.KindMicrosoftCloudAppSecurity
	KindMicrosoftDefenderAdvancedThreatProtection KindBasicDataConnector = original.KindMicrosoftDefenderAdvancedThreatProtection
	KindOffice365                                 KindBasicDataConnector = original.KindOffice365
	KindThreatIntelligence                        KindBasicDataConnector = original.KindThreatIntelligence
)

type KindBasicEntity = original.KindBasicEntity

const (
	KindAccount          KindBasicEntity = original.KindAccount
	KindAzureResource    KindBasicEntity = original.KindAzureResource
	KindCloudApplication KindBasicEntity = original.KindCloudApplication
	KindDNSResolution    KindBasicEntity = original.KindDNSResolution
	KindEntity           KindBasicEntity = original.KindEntity
	KindFile             KindBasicEntity = original.KindFile
	KindFileHash         KindBasicEntity = original.KindFileHash
	KindHost             KindBasicEntity = original.KindHost
	KindIP               KindBasicEntity = original.KindIP
	KindMalware          KindBasicEntity = original.KindMalware
	KindProcess          KindBasicEntity = original.KindProcess
	KindRegistryKey      KindBasicEntity = original.KindRegistryKey
	KindRegistryValue    KindBasicEntity = original.KindRegistryValue
	KindSecurityAlert    KindBasicEntity = original.KindSecurityAlert
	KindSecurityGroup    KindBasicEntity = original.KindSecurityGroup
	KindURL              KindBasicEntity = original.KindURL
)

type KindBasicSettings = original.KindBasicSettings

const (
	KindSettings       KindBasicSettings = original.KindSettings
	KindToggleSettings KindBasicSettings = original.KindToggleSettings
	KindUebaSettings   KindBasicSettings = original.KindUebaSettings
)

type LicenseStatus = original.LicenseStatus

const (
	LicenseStatusDisabled LicenseStatus = original.LicenseStatusDisabled
	LicenseStatusEnabled  LicenseStatus = original.LicenseStatusEnabled
)

type OSFamily = original.OSFamily

const (
	Android OSFamily = original.Android
	IOS     OSFamily = original.IOS
	Linux   OSFamily = original.Linux
	Windows OSFamily = original.Windows
)

type RegistryHive = original.RegistryHive

const (
	HKEYA                        RegistryHive = original.HKEYA
	HKEYCLASSESROOT              RegistryHive = original.HKEYCLASSESROOT
	HKEYCURRENTCONFIG            RegistryHive = original.HKEYCURRENTCONFIG
	HKEYCURRENTUSER              RegistryHive = original.HKEYCURRENTUSER
	HKEYCURRENTUSERLOCALSETTINGS RegistryHive = original.HKEYCURRENTUSERLOCALSETTINGS
	HKEYLOCALMACHINE             RegistryHive = original.HKEYLOCALMACHINE
	HKEYPERFORMANCEDATA          RegistryHive = original.HKEYPERFORMANCEDATA
	HKEYPERFORMANCENLSTEXT       RegistryHive = original.HKEYPERFORMANCENLSTEXT
	HKEYPERFORMANCETEXT          RegistryHive = original.HKEYPERFORMANCETEXT
	HKEYUSERS                    RegistryHive = original.HKEYUSERS
)

type RegistryValueKind = original.RegistryValueKind

const (
	RegistryValueKindBinary       RegistryValueKind = original.RegistryValueKindBinary
	RegistryValueKindDWord        RegistryValueKind = original.RegistryValueKindDWord
	RegistryValueKindExpandString RegistryValueKind = original.RegistryValueKindExpandString
	RegistryValueKindMultiString  RegistryValueKind = original.RegistryValueKindMultiString
	RegistryValueKindNone         RegistryValueKind = original.RegistryValueKindNone
	RegistryValueKindQWord        RegistryValueKind = original.RegistryValueKindQWord
	RegistryValueKindString       RegistryValueKind = original.RegistryValueKindString
	RegistryValueKindUnknown      RegistryValueKind = original.RegistryValueKindUnknown
)

type SettingKind = original.SettingKind

const (
	SettingKindToggleSettings SettingKind = original.SettingKindToggleSettings
	SettingKindUebaSettings   SettingKind = original.SettingKindUebaSettings
)

type StatusInMcas = original.StatusInMcas

const (
	StatusInMcasDisabled StatusInMcas = original.StatusInMcasDisabled
	StatusInMcasEnabled  StatusInMcas = original.StatusInMcasEnabled
)

type TemplateStatus = original.TemplateStatus

const (
	Available    TemplateStatus = original.Available
	Installed    TemplateStatus = original.Installed
	NotAvailable TemplateStatus = original.NotAvailable
)

type TriggerOperator = original.TriggerOperator

const (
	Equal       TriggerOperator = original.Equal
	GreaterThan TriggerOperator = original.GreaterThan
	LessThan    TriggerOperator = original.LessThan
	NotEqual    TriggerOperator = original.NotEqual
)

type AADDataConnector = original.AADDataConnector
type AADDataConnectorProperties = original.AADDataConnectorProperties
type AATPDataConnector = original.AATPDataConnector
type AATPDataConnectorProperties = original.AATPDataConnectorProperties
type ASCDataConnector = original.ASCDataConnector
type ASCDataConnectorProperties = original.ASCDataConnectorProperties
type AccountEntity = original.AccountEntity
type AccountEntityProperties = original.AccountEntityProperties
type ActionRequest = original.ActionRequest
type ActionRequestProperties = original.ActionRequestProperties
type ActionResponse = original.ActionResponse
type ActionResponseProperties = original.ActionResponseProperties
type ActionsClient = original.ActionsClient
type ActionsList = original.ActionsList
type ActionsListIterator = original.ActionsListIterator
type ActionsListPage = original.ActionsListPage
type Aggregations = original.Aggregations
type AggregationsKind1 = original.AggregationsKind1
type AggregationsModel = original.AggregationsModel
type AlertRule = original.AlertRule
type AlertRuleKind1 = original.AlertRuleKind1
type AlertRuleModel = original.AlertRuleModel
type AlertRuleTemplate = original.AlertRuleTemplate
type AlertRuleTemplateModel = original.AlertRuleTemplateModel
type AlertRuleTemplatesClient = original.AlertRuleTemplatesClient
type AlertRuleTemplatesList = original.AlertRuleTemplatesList
type AlertRuleTemplatesListIterator = original.AlertRuleTemplatesListIterator
type AlertRuleTemplatesListPage = original.AlertRuleTemplatesListPage
type AlertRulesClient = original.AlertRulesClient
type AlertRulesList = original.AlertRulesList
type AlertRulesListIterator = original.AlertRulesListIterator
type AlertRulesListPage = original.AlertRulesListPage
type AlertsDataTypeOfDataConnector = original.AlertsDataTypeOfDataConnector
type AlertsDataTypeOfDataConnectorAlerts = original.AlertsDataTypeOfDataConnectorAlerts
type AwsCloudTrailDataConnector = original.AwsCloudTrailDataConnector
type AwsCloudTrailDataConnectorDataTypes = original.AwsCloudTrailDataConnectorDataTypes
type AwsCloudTrailDataConnectorDataTypesLogs = original.AwsCloudTrailDataConnectorDataTypesLogs
type AwsCloudTrailDataConnectorProperties = original.AwsCloudTrailDataConnectorProperties
type AzureResourceEntity = original.AzureResourceEntity
type AzureResourceEntityProperties = original.AzureResourceEntityProperties
type BaseAlertRuleTemplateProperties = original.BaseAlertRuleTemplateProperties
type BaseClient = original.BaseClient
type BasicAggregations = original.BasicAggregations
type BasicAlertRule = original.BasicAlertRule
type BasicAlertRuleTemplate = original.BasicAlertRuleTemplate
type BasicDataConnector = original.BasicDataConnector
type BasicEntity = original.BasicEntity
type BasicSettings = original.BasicSettings
type Bookmark = original.Bookmark
type BookmarkList = original.BookmarkList
type BookmarkListIterator = original.BookmarkListIterator
type BookmarkListPage = original.BookmarkListPage
type BookmarkProperties = original.BookmarkProperties
type BookmarksClient = original.BookmarksClient
type Case = original.Case
type CaseComment = original.CaseComment
type CaseCommentList = original.CaseCommentList
type CaseCommentListIterator = original.CaseCommentListIterator
type CaseCommentListPage = original.CaseCommentListPage
type CaseCommentProperties = original.CaseCommentProperties
type CaseCommentsClient = original.CaseCommentsClient
type CaseList = original.CaseList
type CaseListIterator = original.CaseListIterator
type CaseListPage = original.CaseListPage
type CaseProperties = original.CaseProperties
type CasesAggregation = original.CasesAggregation
type CasesAggregationBySeverityProperties = original.CasesAggregationBySeverityProperties
type CasesAggregationByStatusProperties = original.CasesAggregationByStatusProperties
type CasesAggregationProperties = original.CasesAggregationProperties
type CasesAggregationsClient = original.CasesAggregationsClient
type CasesClient = original.CasesClient
type CloudApplicationEntity = original.CloudApplicationEntity
type CloudApplicationEntityProperties = original.CloudApplicationEntityProperties
type CloudError = original.CloudError
type CloudErrorBody = original.CloudErrorBody
type CommentsClient = original.CommentsClient
type DNSEntity = original.DNSEntity
type DNSEntityProperties = original.DNSEntityProperties
type DataConnector = original.DataConnector
type DataConnectorDataTypeCommon = original.DataConnectorDataTypeCommon
type DataConnectorKind1 = original.DataConnectorKind1
type DataConnectorList = original.DataConnectorList
type DataConnectorListIterator = original.DataConnectorListIterator
type DataConnectorListPage = original.DataConnectorListPage
type DataConnectorModel = original.DataConnectorModel
type DataConnectorStatus = original.DataConnectorStatus
type DataConnectorTenantID = original.DataConnectorTenantID
type DataConnectorWithAlertsProperties = original.DataConnectorWithAlertsProperties
type DataConnectorsClient = original.DataConnectorsClient
type EntitiesClient = original.EntitiesClient
type Entity = original.Entity
type EntityCommonProperties = original.EntityCommonProperties
type EntityExpandParameters = original.EntityExpandParameters
type EntityExpandResponse = original.EntityExpandResponse
type EntityExpandResponseValue = original.EntityExpandResponseValue
type EntityKind1 = original.EntityKind1
type EntityList = original.EntityList
type EntityListIterator = original.EntityListIterator
type EntityListPage = original.EntityListPage
type EntityModel = original.EntityModel
type EntityQueriesClient = original.EntityQueriesClient
type EntityQuery = original.EntityQuery
type EntityQueryList = original.EntityQueryList
type EntityQueryListIterator = original.EntityQueryListIterator
type EntityQueryListPage = original.EntityQueryListPage
type EntityQueryProperties = original.EntityQueryProperties
type ExpansionResultAggregation = original.ExpansionResultAggregation
type ExpansionResultsMetadata = original.ExpansionResultsMetadata
type FileEntity = original.FileEntity
type FileEntityProperties = original.FileEntityProperties
type FileHashEntity = original.FileHashEntity
type FileHashEntityProperties = original.FileHashEntityProperties
type FilterAlertRuleTemplate = original.FilterAlertRuleTemplate
type FilterAlertRuleTemplateProperties = original.FilterAlertRuleTemplateProperties
type FilterAlertRuleTemplatePropertiesModel = original.FilterAlertRuleTemplatePropertiesModel
type FusionAlertRuleTemplate = original.FusionAlertRuleTemplate
type FusionAlertRuleTemplateProperties = original.FusionAlertRuleTemplateProperties
type FusionAlertRuleTemplatePropertiesModel = original.FusionAlertRuleTemplatePropertiesModel
type GeoLocation = original.GeoLocation
type HostEntity = original.HostEntity
type HostEntityProperties = original.HostEntityProperties
type IPEntity = original.IPEntity
type IPEntityProperties = original.IPEntityProperties
type MCASDataConnector = original.MCASDataConnector
type MCASDataConnectorDataTypes = original.MCASDataConnectorDataTypes
type MCASDataConnectorDataTypesDiscoveryLogs = original.MCASDataConnectorDataTypesDiscoveryLogs
type MCASDataConnectorProperties = original.MCASDataConnectorProperties
type MDATPDataConnector = original.MDATPDataConnector
type MDATPDataConnectorProperties = original.MDATPDataConnectorProperties
type MalwareEntity = original.MalwareEntity
type MalwareEntityProperties = original.MalwareEntityProperties
type OfficeConsent = original.OfficeConsent
type OfficeConsentList = original.OfficeConsentList
type OfficeConsentListIterator = original.OfficeConsentListIterator
type OfficeConsentListPage = original.OfficeConsentListPage
type OfficeConsentProperties = original.OfficeConsentProperties
type OfficeConsentsClient = original.OfficeConsentsClient
type OfficeDataConnector = original.OfficeDataConnector
type OfficeDataConnectorDataTypes = original.OfficeDataConnectorDataTypes
type OfficeDataConnectorDataTypesExchange = original.OfficeDataConnectorDataTypesExchange
type OfficeDataConnectorDataTypesSharePoint = original.OfficeDataConnectorDataTypesSharePoint
type OfficeDataConnectorProperties = original.OfficeDataConnectorProperties
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationsClient = original.OperationsClient
type OperationsList = original.OperationsList
type OperationsListIterator = original.OperationsListIterator
type OperationsListPage = original.OperationsListPage
type ProcessEntity = original.ProcessEntity
type ProcessEntityProperties = original.ProcessEntityProperties
type ProductSettingsClient = original.ProductSettingsClient
type RegistryKeyEntity = original.RegistryKeyEntity
type RegistryKeyEntityProperties = original.RegistryKeyEntityProperties
type RegistryValueEntity = original.RegistryValueEntity
type RegistryValueEntityProperties = original.RegistryValueEntityProperties
type Resource = original.Resource
type ScheduledAlertRule = original.ScheduledAlertRule
type ScheduledAlertRuleProperties = original.ScheduledAlertRuleProperties
type ScheduledAlertRuleTemplate = original.ScheduledAlertRuleTemplate
type ScheduledAlertRuleTemplateProperties = original.ScheduledAlertRuleTemplateProperties
type ScheduledAlertRuleTemplatePropertiesModel = original.ScheduledAlertRuleTemplatePropertiesModel
type SecurityAlert = original.SecurityAlert
type SecurityAlertProperties = original.SecurityAlertProperties
type SecurityAlertPropertiesConfidenceReasonsItem = original.SecurityAlertPropertiesConfidenceReasonsItem
type SecurityGroupEntity = original.SecurityGroupEntity
type SecurityGroupEntityProperties = original.SecurityGroupEntityProperties
type Settings = original.Settings
type SettingsKind = original.SettingsKind
type SettingsModel = original.SettingsModel
type TIDataConnector = original.TIDataConnector
type TIDataConnectorDataTypes = original.TIDataConnectorDataTypes
type TIDataConnectorDataTypesIndicators = original.TIDataConnectorDataTypesIndicators
type TIDataConnectorProperties = original.TIDataConnectorProperties
type ThreatIntelligence = original.ThreatIntelligence
type ToggleSettings = original.ToggleSettings
type ToggleSettingsProperties = original.ToggleSettingsProperties
type URLEntity = original.URLEntity
type URLEntityProperties = original.URLEntityProperties
type UebaSettings = original.UebaSettings
type UebaSettingsProperties = original.UebaSettingsProperties
type UserInfo = original.UserInfo

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewActionsClient(subscriptionID string) ActionsClient {
	return original.NewActionsClient(subscriptionID)
}
func NewActionsClientWithBaseURI(baseURI string, subscriptionID string) ActionsClient {
	return original.NewActionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewActionsListIterator(page ActionsListPage) ActionsListIterator {
	return original.NewActionsListIterator(page)
}
func NewActionsListPage(getNextPage func(context.Context, ActionsList) (ActionsList, error)) ActionsListPage {
	return original.NewActionsListPage(getNextPage)
}
func NewAlertRuleTemplatesClient(subscriptionID string) AlertRuleTemplatesClient {
	return original.NewAlertRuleTemplatesClient(subscriptionID)
}
func NewAlertRuleTemplatesClientWithBaseURI(baseURI string, subscriptionID string) AlertRuleTemplatesClient {
	return original.NewAlertRuleTemplatesClientWithBaseURI(baseURI, subscriptionID)
}
func NewAlertRuleTemplatesListIterator(page AlertRuleTemplatesListPage) AlertRuleTemplatesListIterator {
	return original.NewAlertRuleTemplatesListIterator(page)
}
func NewAlertRuleTemplatesListPage(getNextPage func(context.Context, AlertRuleTemplatesList) (AlertRuleTemplatesList, error)) AlertRuleTemplatesListPage {
	return original.NewAlertRuleTemplatesListPage(getNextPage)
}
func NewAlertRulesClient(subscriptionID string) AlertRulesClient {
	return original.NewAlertRulesClient(subscriptionID)
}
func NewAlertRulesClientWithBaseURI(baseURI string, subscriptionID string) AlertRulesClient {
	return original.NewAlertRulesClientWithBaseURI(baseURI, subscriptionID)
}
func NewAlertRulesListIterator(page AlertRulesListPage) AlertRulesListIterator {
	return original.NewAlertRulesListIterator(page)
}
func NewAlertRulesListPage(getNextPage func(context.Context, AlertRulesList) (AlertRulesList, error)) AlertRulesListPage {
	return original.NewAlertRulesListPage(getNextPage)
}
func NewBookmarkListIterator(page BookmarkListPage) BookmarkListIterator {
	return original.NewBookmarkListIterator(page)
}
func NewBookmarkListPage(getNextPage func(context.Context, BookmarkList) (BookmarkList, error)) BookmarkListPage {
	return original.NewBookmarkListPage(getNextPage)
}
func NewBookmarksClient(subscriptionID string) BookmarksClient {
	return original.NewBookmarksClient(subscriptionID)
}
func NewBookmarksClientWithBaseURI(baseURI string, subscriptionID string) BookmarksClient {
	return original.NewBookmarksClientWithBaseURI(baseURI, subscriptionID)
}
func NewCaseCommentListIterator(page CaseCommentListPage) CaseCommentListIterator {
	return original.NewCaseCommentListIterator(page)
}
func NewCaseCommentListPage(getNextPage func(context.Context, CaseCommentList) (CaseCommentList, error)) CaseCommentListPage {
	return original.NewCaseCommentListPage(getNextPage)
}
func NewCaseCommentsClient(subscriptionID string) CaseCommentsClient {
	return original.NewCaseCommentsClient(subscriptionID)
}
func NewCaseCommentsClientWithBaseURI(baseURI string, subscriptionID string) CaseCommentsClient {
	return original.NewCaseCommentsClientWithBaseURI(baseURI, subscriptionID)
}
func NewCaseListIterator(page CaseListPage) CaseListIterator {
	return original.NewCaseListIterator(page)
}
func NewCaseListPage(getNextPage func(context.Context, CaseList) (CaseList, error)) CaseListPage {
	return original.NewCaseListPage(getNextPage)
}
func NewCasesAggregationsClient(subscriptionID string) CasesAggregationsClient {
	return original.NewCasesAggregationsClient(subscriptionID)
}
func NewCasesAggregationsClientWithBaseURI(baseURI string, subscriptionID string) CasesAggregationsClient {
	return original.NewCasesAggregationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewCasesClient(subscriptionID string) CasesClient {
	return original.NewCasesClient(subscriptionID)
}
func NewCasesClientWithBaseURI(baseURI string, subscriptionID string) CasesClient {
	return original.NewCasesClientWithBaseURI(baseURI, subscriptionID)
}
func NewCommentsClient(subscriptionID string) CommentsClient {
	return original.NewCommentsClient(subscriptionID)
}
func NewCommentsClientWithBaseURI(baseURI string, subscriptionID string) CommentsClient {
	return original.NewCommentsClientWithBaseURI(baseURI, subscriptionID)
}
func NewDataConnectorListIterator(page DataConnectorListPage) DataConnectorListIterator {
	return original.NewDataConnectorListIterator(page)
}
func NewDataConnectorListPage(getNextPage func(context.Context, DataConnectorList) (DataConnectorList, error)) DataConnectorListPage {
	return original.NewDataConnectorListPage(getNextPage)
}
func NewDataConnectorsClient(subscriptionID string) DataConnectorsClient {
	return original.NewDataConnectorsClient(subscriptionID)
}
func NewDataConnectorsClientWithBaseURI(baseURI string, subscriptionID string) DataConnectorsClient {
	return original.NewDataConnectorsClientWithBaseURI(baseURI, subscriptionID)
}
func NewEntitiesClient(subscriptionID string) EntitiesClient {
	return original.NewEntitiesClient(subscriptionID)
}
func NewEntitiesClientWithBaseURI(baseURI string, subscriptionID string) EntitiesClient {
	return original.NewEntitiesClientWithBaseURI(baseURI, subscriptionID)
}
func NewEntityListIterator(page EntityListPage) EntityListIterator {
	return original.NewEntityListIterator(page)
}
func NewEntityListPage(getNextPage func(context.Context, EntityList) (EntityList, error)) EntityListPage {
	return original.NewEntityListPage(getNextPage)
}
func NewEntityQueriesClient(subscriptionID string) EntityQueriesClient {
	return original.NewEntityQueriesClient(subscriptionID)
}
func NewEntityQueriesClientWithBaseURI(baseURI string, subscriptionID string) EntityQueriesClient {
	return original.NewEntityQueriesClientWithBaseURI(baseURI, subscriptionID)
}
func NewEntityQueryListIterator(page EntityQueryListPage) EntityQueryListIterator {
	return original.NewEntityQueryListIterator(page)
}
func NewEntityQueryListPage(getNextPage func(context.Context, EntityQueryList) (EntityQueryList, error)) EntityQueryListPage {
	return original.NewEntityQueryListPage(getNextPage)
}
func NewOfficeConsentListIterator(page OfficeConsentListPage) OfficeConsentListIterator {
	return original.NewOfficeConsentListIterator(page)
}
func NewOfficeConsentListPage(getNextPage func(context.Context, OfficeConsentList) (OfficeConsentList, error)) OfficeConsentListPage {
	return original.NewOfficeConsentListPage(getNextPage)
}
func NewOfficeConsentsClient(subscriptionID string) OfficeConsentsClient {
	return original.NewOfficeConsentsClient(subscriptionID)
}
func NewOfficeConsentsClientWithBaseURI(baseURI string, subscriptionID string) OfficeConsentsClient {
	return original.NewOfficeConsentsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsListIterator(page OperationsListPage) OperationsListIterator {
	return original.NewOperationsListIterator(page)
}
func NewOperationsListPage(getNextPage func(context.Context, OperationsList) (OperationsList, error)) OperationsListPage {
	return original.NewOperationsListPage(getNextPage)
}
func NewProductSettingsClient(subscriptionID string) ProductSettingsClient {
	return original.NewProductSettingsClient(subscriptionID)
}
func NewProductSettingsClientWithBaseURI(baseURI string, subscriptionID string) ProductSettingsClient {
	return original.NewProductSettingsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleAggregationsKindValues() []AggregationsKind {
	return original.PossibleAggregationsKindValues()
}
func PossibleAlertRuleKindValues() []AlertRuleKind {
	return original.PossibleAlertRuleKindValues()
}
func PossibleAlertSeverityValues() []AlertSeverity {
	return original.PossibleAlertSeverityValues()
}
func PossibleAlertStatusValues() []AlertStatus {
	return original.PossibleAlertStatusValues()
}
func PossibleAttackTacticValues() []AttackTactic {
	return original.PossibleAttackTacticValues()
}
func PossibleCaseSeverityValues() []CaseSeverity {
	return original.PossibleCaseSeverityValues()
}
func PossibleCaseStatusValues() []CaseStatus {
	return original.PossibleCaseStatusValues()
}
func PossibleCloseReasonValues() []CloseReason {
	return original.PossibleCloseReasonValues()
}
func PossibleConfidenceLevelValues() []ConfidenceLevel {
	return original.PossibleConfidenceLevelValues()
}
func PossibleConfidenceScoreStatusValues() []ConfidenceScoreStatus {
	return original.PossibleConfidenceScoreStatusValues()
}
func PossibleDataConnectorKindValues() []DataConnectorKind {
	return original.PossibleDataConnectorKindValues()
}
func PossibleDataTypeStateValues() []DataTypeState {
	return original.PossibleDataTypeStateValues()
}
func PossibleDataTypeStatusValues() []DataTypeStatus {
	return original.PossibleDataTypeStatusValues()
}
func PossibleElevationTokenValues() []ElevationToken {
	return original.PossibleElevationTokenValues()
}
func PossibleEntityKindValues() []EntityKind {
	return original.PossibleEntityKindValues()
}
func PossibleEntityTypeValues() []EntityType {
	return original.PossibleEntityTypeValues()
}
func PossibleFileHashAlgorithmValues() []FileHashAlgorithm {
	return original.PossibleFileHashAlgorithmValues()
}
func PossibleKillChainIntentValues() []KillChainIntent {
	return original.PossibleKillChainIntentValues()
}
func PossibleKindBasicAlertRuleTemplateValues() []KindBasicAlertRuleTemplate {
	return original.PossibleKindBasicAlertRuleTemplateValues()
}
func PossibleKindBasicAlertRuleValues() []KindBasicAlertRule {
	return original.PossibleKindBasicAlertRuleValues()
}
func PossibleKindBasicDataConnectorValues() []KindBasicDataConnector {
	return original.PossibleKindBasicDataConnectorValues()
}
func PossibleKindBasicEntityValues() []KindBasicEntity {
	return original.PossibleKindBasicEntityValues()
}
func PossibleKindBasicSettingsValues() []KindBasicSettings {
	return original.PossibleKindBasicSettingsValues()
}
func PossibleKindValues() []Kind {
	return original.PossibleKindValues()
}
func PossibleLicenseStatusValues() []LicenseStatus {
	return original.PossibleLicenseStatusValues()
}
func PossibleOSFamilyValues() []OSFamily {
	return original.PossibleOSFamilyValues()
}
func PossibleRegistryHiveValues() []RegistryHive {
	return original.PossibleRegistryHiveValues()
}
func PossibleRegistryValueKindValues() []RegistryValueKind {
	return original.PossibleRegistryValueKindValues()
}
func PossibleSettingKindValues() []SettingKind {
	return original.PossibleSettingKindValues()
}
func PossibleStatusInMcasValues() []StatusInMcas {
	return original.PossibleStatusInMcasValues()
}
func PossibleTemplateStatusValues() []TemplateStatus {
	return original.PossibleTemplateStatusValues()
}
func PossibleTriggerOperatorValues() []TriggerOperator {
	return original.PossibleTriggerOperatorValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
