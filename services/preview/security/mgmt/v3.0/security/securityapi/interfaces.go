package securityapi

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
	"github.com/Azure/azure-sdk-for-go/services/preview/security/mgmt/v3.0/security"
	"github.com/Azure/go-autorest/autorest"
)

// ComplianceResultsClientAPI contains the set of methods on the ComplianceResultsClient type.
type ComplianceResultsClientAPI interface {
	Get(ctx context.Context, resourceID string, complianceResultName string) (result security.ComplianceResult, err error)
	List(ctx context.Context, scope string) (result security.ComplianceResultListPage, err error)
	ListComplete(ctx context.Context, scope string) (result security.ComplianceResultListIterator, err error)
}

var _ ComplianceResultsClientAPI = (*security.ComplianceResultsClient)(nil)

// PricingsClientAPI contains the set of methods on the PricingsClient type.
type PricingsClientAPI interface {
	Get(ctx context.Context, pricingName string) (result security.Pricing, err error)
	List(ctx context.Context) (result security.PricingList, err error)
	Update(ctx context.Context, pricingName string, pricing security.Pricing) (result security.Pricing, err error)
}

var _ PricingsClientAPI = (*security.PricingsClient)(nil)

// AlertsClientAPI contains the set of methods on the AlertsClient type.
type AlertsClientAPI interface {
	GetResourceGroupLevelAlerts(ctx context.Context, alertName string, resourceGroupName string) (result security.Alert, err error)
	GetSubscriptionLevelAlert(ctx context.Context, alertName string) (result security.Alert, err error)
	List(ctx context.Context, filter string, selectParameter string, expand string, autoDismissRuleName string) (result security.AlertListPage, err error)
	ListComplete(ctx context.Context, filter string, selectParameter string, expand string, autoDismissRuleName string) (result security.AlertListIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string, filter string, selectParameter string, expand string, autoDismissRuleName string) (result security.AlertListPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, filter string, selectParameter string, expand string, autoDismissRuleName string) (result security.AlertListIterator, err error)
	ListResourceGroupLevelAlertsByRegion(ctx context.Context, resourceGroupName string, filter string, selectParameter string, expand string, autoDismissRuleName string) (result security.AlertListPage, err error)
	ListResourceGroupLevelAlertsByRegionComplete(ctx context.Context, resourceGroupName string, filter string, selectParameter string, expand string, autoDismissRuleName string) (result security.AlertListIterator, err error)
	ListSubscriptionLevelAlertsByRegion(ctx context.Context, filter string, selectParameter string, expand string, autoDismissRuleName string) (result security.AlertListPage, err error)
	ListSubscriptionLevelAlertsByRegionComplete(ctx context.Context, filter string, selectParameter string, expand string, autoDismissRuleName string) (result security.AlertListIterator, err error)
	UpdateResourceGroupLevelAlertStateToDismiss(ctx context.Context, alertName string, resourceGroupName string) (result autorest.Response, err error)
	UpdateResourceGroupLevelAlertStateToReactivate(ctx context.Context, alertName string, resourceGroupName string) (result autorest.Response, err error)
	UpdateSubscriptionLevelAlertStateToDismiss(ctx context.Context, alertName string) (result autorest.Response, err error)
	UpdateSubscriptionLevelAlertStateToReactivate(ctx context.Context, alertName string) (result autorest.Response, err error)
}

var _ AlertsClientAPI = (*security.AlertsClient)(nil)

// SettingsClientAPI contains the set of methods on the SettingsClient type.
type SettingsClientAPI interface {
	Get(ctx context.Context, settingName string) (result security.Setting, err error)
	List(ctx context.Context) (result security.SettingsListPage, err error)
	ListComplete(ctx context.Context) (result security.SettingsListIterator, err error)
	Update(ctx context.Context, settingName string, setting security.BasicSetting) (result security.Setting, err error)
}

var _ SettingsClientAPI = (*security.SettingsClient)(nil)

// AdvancedThreatProtectionClientAPI contains the set of methods on the AdvancedThreatProtectionClient type.
type AdvancedThreatProtectionClientAPI interface {
	Create(ctx context.Context, resourceID string, advancedThreatProtectionSetting security.AdvancedThreatProtectionSetting) (result security.AdvancedThreatProtectionSetting, err error)
	Get(ctx context.Context, resourceID string) (result security.AdvancedThreatProtectionSetting, err error)
}

var _ AdvancedThreatProtectionClientAPI = (*security.AdvancedThreatProtectionClient)(nil)

// DevicesClientAPI contains the set of methods on the DevicesClient type.
type DevicesClientAPI interface {
	List(ctx context.Context, resourceGroupName string, solutionName string, limit *int32, skipToken string) (result security.DeviceListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, solutionName string, limit *int32, skipToken string) (result security.DeviceListIterator, err error)
}

var _ DevicesClientAPI = (*security.DevicesClient)(nil)

// DeviceClientAPI contains the set of methods on the DeviceClient type.
type DeviceClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, solutionName string, deviceID string) (result security.Device, err error)
	Update(ctx context.Context, resourceGroupName string, solutionName string, deviceID string, deviceData security.Device) (result security.Device, err error)
}

var _ DeviceClientAPI = (*security.DeviceClient)(nil)

// DeviceSecurityGroupsClientAPI contains the set of methods on the DeviceSecurityGroupsClient type.
type DeviceSecurityGroupsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceID string, deviceSecurityGroupName string, deviceSecurityGroup security.DeviceSecurityGroup) (result security.DeviceSecurityGroup, err error)
	Delete(ctx context.Context, resourceID string, deviceSecurityGroupName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceID string, deviceSecurityGroupName string) (result security.DeviceSecurityGroup, err error)
	List(ctx context.Context, resourceID string) (result security.DeviceSecurityGroupListPage, err error)
	ListComplete(ctx context.Context, resourceID string) (result security.DeviceSecurityGroupListIterator, err error)
}

var _ DeviceSecurityGroupsClientAPI = (*security.DeviceSecurityGroupsClient)(nil)

// IotSecuritySolutionClientAPI contains the set of methods on the IotSecuritySolutionClient type.
type IotSecuritySolutionClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, solutionName string, iotSecuritySolutionData security.IoTSecuritySolutionModel) (result security.IoTSecuritySolutionModel, err error)
	Delete(ctx context.Context, resourceGroupName string, solutionName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, solutionName string) (result security.IoTSecuritySolutionModel, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string, filter string) (result security.IoTSecuritySolutionsListPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, filter string) (result security.IoTSecuritySolutionsListIterator, err error)
	ListBySubscription(ctx context.Context, filter string) (result security.IoTSecuritySolutionsListPage, err error)
	ListBySubscriptionComplete(ctx context.Context, filter string) (result security.IoTSecuritySolutionsListIterator, err error)
	Update(ctx context.Context, resourceGroupName string, solutionName string, updateIotSecuritySolutionData security.UpdateIotSecuritySolutionData) (result security.IoTSecuritySolutionModel, err error)
}

var _ IotSecuritySolutionClientAPI = (*security.IotSecuritySolutionClient)(nil)

// IotSecuritySolutionAnalyticsClientAPI contains the set of methods on the IotSecuritySolutionAnalyticsClient type.
type IotSecuritySolutionAnalyticsClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, solutionName string) (result security.IoTSecuritySolutionAnalyticsModel, err error)
	List(ctx context.Context, resourceGroupName string, solutionName string) (result security.IoTSecuritySolutionAnalyticsModelList, err error)
}

var _ IotSecuritySolutionAnalyticsClientAPI = (*security.IotSecuritySolutionAnalyticsClient)(nil)

// IotSecuritySolutionsAnalyticsAggregatedAlertClientAPI contains the set of methods on the IotSecuritySolutionsAnalyticsAggregatedAlertClient type.
type IotSecuritySolutionsAnalyticsAggregatedAlertClientAPI interface {
	Dismiss(ctx context.Context, resourceGroupName string, solutionName string, aggregatedAlertName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, solutionName string, aggregatedAlertName string) (result security.IoTSecurityAggregatedAlert, err error)
	List(ctx context.Context, resourceGroupName string, solutionName string, top *int32) (result security.IoTSecurityAggregatedAlertListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, solutionName string, top *int32) (result security.IoTSecurityAggregatedAlertListIterator, err error)
}

var _ IotSecuritySolutionsAnalyticsAggregatedAlertClientAPI = (*security.IotSecuritySolutionsAnalyticsAggregatedAlertClient)(nil)

// IotSecuritySolutionsAnalyticsRecommendationClientAPI contains the set of methods on the IotSecuritySolutionsAnalyticsRecommendationClient type.
type IotSecuritySolutionsAnalyticsRecommendationClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, solutionName string, aggregatedRecommendationName string) (result security.IoTSecurityAggregatedRecommendation, err error)
	List(ctx context.Context, resourceGroupName string, solutionName string, top *int32) (result security.IoTSecurityAggregatedRecommendationListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, solutionName string, top *int32) (result security.IoTSecurityAggregatedRecommendationListIterator, err error)
}

var _ IotSecuritySolutionsAnalyticsRecommendationClientAPI = (*security.IotSecuritySolutionsAnalyticsRecommendationClient)(nil)

// IotAlertTypesClientAPI contains the set of methods on the IotAlertTypesClient type.
type IotAlertTypesClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, solutionName string, iotAlertTypeName string) (result security.IotAlertType, err error)
	List(ctx context.Context, resourceGroupName string, solutionName string) (result security.IotAlertTypeList, err error)
}

var _ IotAlertTypesClientAPI = (*security.IotAlertTypesClient)(nil)

// IotAlertsClientAPI contains the set of methods on the IotAlertsClient type.
type IotAlertsClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, solutionName string, iotAlertID string) (result security.IotAlert, err error)
	List(ctx context.Context, resourceGroupName string, solutionName string, minStartTimeUtc string, maxStartTimeUtc string, alertType string, limit *int32, skipToken string) (result security.IotAlertListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, solutionName string, minStartTimeUtc string, maxStartTimeUtc string, alertType string, limit *int32, skipToken string) (result security.IotAlertListIterator, err error)
}

var _ IotAlertsClientAPI = (*security.IotAlertsClient)(nil)

// IotRecommendationTypesClientAPI contains the set of methods on the IotRecommendationTypesClient type.
type IotRecommendationTypesClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, solutionName string, iotRecommendationTypeName string) (result security.IotRecommendationType, err error)
	List(ctx context.Context, resourceGroupName string, solutionName string) (result security.IotRecommendationTypeList, err error)
}

var _ IotRecommendationTypesClientAPI = (*security.IotRecommendationTypesClient)(nil)

// IotRecommendationsClientAPI contains the set of methods on the IotRecommendationsClient type.
type IotRecommendationsClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, solutionName string, iotRecommendationID string) (result security.IotRecommendation, err error)
	List(ctx context.Context, resourceGroupName string, solutionName string, recommendationType string, limit *int32, skipToken string) (result security.IotRecommendationListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, solutionName string, recommendationType string, limit *int32, skipToken string) (result security.IotRecommendationListIterator, err error)
}

var _ IotRecommendationsClientAPI = (*security.IotRecommendationsClient)(nil)

// LocationsClientAPI contains the set of methods on the LocationsClient type.
type LocationsClientAPI interface {
	Get(ctx context.Context) (result security.AscLocation, err error)
	List(ctx context.Context) (result security.AscLocationListPage, err error)
	ListComplete(ctx context.Context) (result security.AscLocationListIterator, err error)
}

var _ LocationsClientAPI = (*security.LocationsClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result security.OperationListPage, err error)
	ListComplete(ctx context.Context) (result security.OperationListIterator, err error)
}

var _ OperationsClientAPI = (*security.OperationsClient)(nil)

// TasksClientAPI contains the set of methods on the TasksClient type.
type TasksClientAPI interface {
	GetResourceGroupLevelTask(ctx context.Context, resourceGroupName string, taskName string) (result security.Task, err error)
	GetSubscriptionLevelTask(ctx context.Context, taskName string) (result security.Task, err error)
	List(ctx context.Context, filter string) (result security.TaskListPage, err error)
	ListComplete(ctx context.Context, filter string) (result security.TaskListIterator, err error)
	ListByHomeRegion(ctx context.Context, filter string) (result security.TaskListPage, err error)
	ListByHomeRegionComplete(ctx context.Context, filter string) (result security.TaskListIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string, filter string) (result security.TaskListPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, filter string) (result security.TaskListIterator, err error)
	UpdateResourceGroupLevelTaskState(ctx context.Context, resourceGroupName string, taskName string, taskUpdateActionType string) (result autorest.Response, err error)
	UpdateSubscriptionLevelTaskState(ctx context.Context, taskName string, taskUpdateActionType string) (result autorest.Response, err error)
}

var _ TasksClientAPI = (*security.TasksClient)(nil)

// AutoProvisioningSettingsClientAPI contains the set of methods on the AutoProvisioningSettingsClient type.
type AutoProvisioningSettingsClientAPI interface {
	Create(ctx context.Context, settingName string, setting security.AutoProvisioningSetting) (result security.AutoProvisioningSetting, err error)
	Get(ctx context.Context, settingName string) (result security.AutoProvisioningSetting, err error)
	List(ctx context.Context) (result security.AutoProvisioningSettingListPage, err error)
	ListComplete(ctx context.Context) (result security.AutoProvisioningSettingListIterator, err error)
}

var _ AutoProvisioningSettingsClientAPI = (*security.AutoProvisioningSettingsClient)(nil)

// CompliancesClientAPI contains the set of methods on the CompliancesClient type.
type CompliancesClientAPI interface {
	Get(ctx context.Context, scope string, complianceName string) (result security.Compliance, err error)
	List(ctx context.Context, scope string) (result security.ComplianceListPage, err error)
	ListComplete(ctx context.Context, scope string) (result security.ComplianceListIterator, err error)
}

var _ CompliancesClientAPI = (*security.CompliancesClient)(nil)

// InformationProtectionPoliciesClientAPI contains the set of methods on the InformationProtectionPoliciesClient type.
type InformationProtectionPoliciesClientAPI interface {
	CreateOrUpdate(ctx context.Context, scope string, informationProtectionPolicyName string, informationProtectionPolicy security.InformationProtectionPolicy) (result security.InformationProtectionPolicy, err error)
	Get(ctx context.Context, scope string, informationProtectionPolicyName string) (result security.InformationProtectionPolicy, err error)
	List(ctx context.Context, scope string) (result security.InformationProtectionPolicyListPage, err error)
	ListComplete(ctx context.Context, scope string) (result security.InformationProtectionPolicyListIterator, err error)
}

var _ InformationProtectionPoliciesClientAPI = (*security.InformationProtectionPoliciesClient)(nil)

// ContactsClientAPI contains the set of methods on the ContactsClient type.
type ContactsClientAPI interface {
	Create(ctx context.Context, securityContactName string, securityContact security.Contact) (result security.Contact, err error)
	Delete(ctx context.Context, securityContactName string) (result autorest.Response, err error)
	Get(ctx context.Context, securityContactName string) (result security.Contact, err error)
	List(ctx context.Context) (result security.ContactListPage, err error)
	ListComplete(ctx context.Context) (result security.ContactListIterator, err error)
	Update(ctx context.Context, securityContactName string, securityContact security.Contact) (result security.Contact, err error)
}

var _ ContactsClientAPI = (*security.ContactsClient)(nil)

// WorkspaceSettingsClientAPI contains the set of methods on the WorkspaceSettingsClient type.
type WorkspaceSettingsClientAPI interface {
	Create(ctx context.Context, workspaceSettingName string, workspaceSetting security.WorkspaceSetting) (result security.WorkspaceSetting, err error)
	Delete(ctx context.Context, workspaceSettingName string) (result autorest.Response, err error)
	Get(ctx context.Context, workspaceSettingName string) (result security.WorkspaceSetting, err error)
	List(ctx context.Context) (result security.WorkspaceSettingListPage, err error)
	ListComplete(ctx context.Context) (result security.WorkspaceSettingListIterator, err error)
	Update(ctx context.Context, workspaceSettingName string, workspaceSetting security.WorkspaceSetting) (result security.WorkspaceSetting, err error)
}

var _ WorkspaceSettingsClientAPI = (*security.WorkspaceSettingsClient)(nil)

// RegulatoryComplianceStandardsClientAPI contains the set of methods on the RegulatoryComplianceStandardsClient type.
type RegulatoryComplianceStandardsClientAPI interface {
	Get(ctx context.Context, regulatoryComplianceStandardName string) (result security.RegulatoryComplianceStandard, err error)
	List(ctx context.Context, filter string) (result security.RegulatoryComplianceStandardListPage, err error)
	ListComplete(ctx context.Context, filter string) (result security.RegulatoryComplianceStandardListIterator, err error)
}

var _ RegulatoryComplianceStandardsClientAPI = (*security.RegulatoryComplianceStandardsClient)(nil)

// RegulatoryComplianceControlsClientAPI contains the set of methods on the RegulatoryComplianceControlsClient type.
type RegulatoryComplianceControlsClientAPI interface {
	Get(ctx context.Context, regulatoryComplianceStandardName string, regulatoryComplianceControlName string) (result security.RegulatoryComplianceControl, err error)
	List(ctx context.Context, regulatoryComplianceStandardName string, filter string) (result security.RegulatoryComplianceControlListPage, err error)
	ListComplete(ctx context.Context, regulatoryComplianceStandardName string, filter string) (result security.RegulatoryComplianceControlListIterator, err error)
}

var _ RegulatoryComplianceControlsClientAPI = (*security.RegulatoryComplianceControlsClient)(nil)

// RegulatoryComplianceAssessmentsClientAPI contains the set of methods on the RegulatoryComplianceAssessmentsClient type.
type RegulatoryComplianceAssessmentsClientAPI interface {
	Get(ctx context.Context, regulatoryComplianceStandardName string, regulatoryComplianceControlName string, regulatoryComplianceAssessmentName string) (result security.RegulatoryComplianceAssessment, err error)
	List(ctx context.Context, regulatoryComplianceStandardName string, regulatoryComplianceControlName string, filter string) (result security.RegulatoryComplianceAssessmentListPage, err error)
	ListComplete(ctx context.Context, regulatoryComplianceStandardName string, regulatoryComplianceControlName string, filter string) (result security.RegulatoryComplianceAssessmentListIterator, err error)
}

var _ RegulatoryComplianceAssessmentsClientAPI = (*security.RegulatoryComplianceAssessmentsClient)(nil)

// SubAssessmentsClientAPI contains the set of methods on the SubAssessmentsClient type.
type SubAssessmentsClientAPI interface {
	Get(ctx context.Context, scope string, assessmentName string, subAssessmentName string) (result security.SubAssessment, err error)
	List(ctx context.Context, scope string, assessmentName string) (result security.SubAssessmentListPage, err error)
	ListComplete(ctx context.Context, scope string, assessmentName string) (result security.SubAssessmentListIterator, err error)
	ListAll(ctx context.Context, scope string) (result security.SubAssessmentListPage, err error)
	ListAllComplete(ctx context.Context, scope string) (result security.SubAssessmentListIterator, err error)
}

var _ SubAssessmentsClientAPI = (*security.SubAssessmentsClient)(nil)

// AutomationsClientAPI contains the set of methods on the AutomationsClient type.
type AutomationsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, automationName string, automation security.Automation) (result security.Automation, err error)
	Delete(ctx context.Context, resourceGroupName string, automationName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, automationName string) (result security.Automation, err error)
	List(ctx context.Context) (result security.AutomationListPage, err error)
	ListComplete(ctx context.Context) (result security.AutomationListIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result security.AutomationListPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result security.AutomationListIterator, err error)
	Validate(ctx context.Context, resourceGroupName string, automationName string, automation security.Automation) (result security.AutomationValidationStatus, err error)
}

var _ AutomationsClientAPI = (*security.AutomationsClient)(nil)

// AlertsSuppressionRulesClientAPI contains the set of methods on the AlertsSuppressionRulesClient type.
type AlertsSuppressionRulesClientAPI interface {
	Delete(ctx context.Context, alertsSuppressionRuleName string) (result autorest.Response, err error)
	Get(ctx context.Context, alertsSuppressionRuleName string) (result security.AlertsSuppressionRule, err error)
	List(ctx context.Context, alertType string) (result security.AlertsSuppressionRulesListPage, err error)
	ListComplete(ctx context.Context, alertType string) (result security.AlertsSuppressionRulesListIterator, err error)
	Update(ctx context.Context, alertsSuppressionRuleName string, alertsSuppressionRule security.AlertsSuppressionRule) (result security.AlertsSuppressionRule, err error)
}

var _ AlertsSuppressionRulesClientAPI = (*security.AlertsSuppressionRulesClient)(nil)

// ServerVulnerabilityAssessmentClientAPI contains the set of methods on the ServerVulnerabilityAssessmentClient type.
type ServerVulnerabilityAssessmentClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceNamespace string, resourceType string, resourceName string) (result security.ServerVulnerabilityAssessment, err error)
	Delete(ctx context.Context, resourceGroupName string, resourceNamespace string, resourceType string, resourceName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, resourceNamespace string, resourceType string, resourceName string) (result security.ServerVulnerabilityAssessment, err error)
	ListByExtendedResource(ctx context.Context, resourceGroupName string, resourceNamespace string, resourceType string, resourceName string) (result security.ServerVulnerabilityAssessmentsList, err error)
}

var _ ServerVulnerabilityAssessmentClientAPI = (*security.ServerVulnerabilityAssessmentClient)(nil)

// AssessmentsMetadataClientAPI contains the set of methods on the AssessmentsMetadataClient type.
type AssessmentsMetadataClientAPI interface {
	CreateInSubscription(ctx context.Context, assessmentMetadataName string, assessmentMetadata security.AssessmentMetadata) (result security.AssessmentMetadata, err error)
	DeleteInSubscription(ctx context.Context, assessmentMetadataName string) (result autorest.Response, err error)
	Get(ctx context.Context, assessmentMetadataName string) (result security.AssessmentMetadata, err error)
	GetInSubscription(ctx context.Context, assessmentMetadataName string) (result security.AssessmentMetadata, err error)
	List(ctx context.Context) (result security.AssessmentMetadataListPage, err error)
	ListComplete(ctx context.Context) (result security.AssessmentMetadataListIterator, err error)
	ListBySubscription(ctx context.Context) (result security.AssessmentMetadataListPage, err error)
	ListBySubscriptionComplete(ctx context.Context) (result security.AssessmentMetadataListIterator, err error)
}

var _ AssessmentsMetadataClientAPI = (*security.AssessmentsMetadataClient)(nil)

// AssessmentsClientAPI contains the set of methods on the AssessmentsClient type.
type AssessmentsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceID string, assessmentName string, assessment security.Assessment) (result security.Assessment, err error)
	Delete(ctx context.Context, resourceID string, assessmentName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceID string, assessmentName string, expand security.ExpandEnum) (result security.Assessment, err error)
	List(ctx context.Context, scope string) (result security.AssessmentListPage, err error)
	ListComplete(ctx context.Context, scope string) (result security.AssessmentListIterator, err error)
}

var _ AssessmentsClientAPI = (*security.AssessmentsClient)(nil)

// AdaptiveApplicationControlsClientAPI contains the set of methods on the AdaptiveApplicationControlsClient type.
type AdaptiveApplicationControlsClientAPI interface {
	Delete(ctx context.Context, groupName string) (result autorest.Response, err error)
	Get(ctx context.Context, groupName string) (result security.AdaptiveApplicationControlGroup, err error)
	List(ctx context.Context, includePathRecommendations *bool, summary *bool) (result security.AdaptiveApplicationControlGroups, err error)
	Put(ctx context.Context, groupName string, body security.AdaptiveApplicationControlGroup) (result security.AdaptiveApplicationControlGroup, err error)
}

var _ AdaptiveApplicationControlsClientAPI = (*security.AdaptiveApplicationControlsClient)(nil)

// AdaptiveNetworkHardeningsClientAPI contains the set of methods on the AdaptiveNetworkHardeningsClient type.
type AdaptiveNetworkHardeningsClientAPI interface {
	Enforce(ctx context.Context, resourceGroupName string, resourceNamespace string, resourceType string, resourceName string, adaptiveNetworkHardeningResourceName string, body security.AdaptiveNetworkHardeningEnforceRequest) (result security.AdaptiveNetworkHardeningsEnforceFuture, err error)
	Get(ctx context.Context, resourceGroupName string, resourceNamespace string, resourceType string, resourceName string, adaptiveNetworkHardeningResourceName string) (result security.AdaptiveNetworkHardening, err error)
	ListByExtendedResource(ctx context.Context, resourceGroupName string, resourceNamespace string, resourceType string, resourceName string) (result security.AdaptiveNetworkHardeningsListPage, err error)
	ListByExtendedResourceComplete(ctx context.Context, resourceGroupName string, resourceNamespace string, resourceType string, resourceName string) (result security.AdaptiveNetworkHardeningsListIterator, err error)
}

var _ AdaptiveNetworkHardeningsClientAPI = (*security.AdaptiveNetworkHardeningsClient)(nil)

// AllowedConnectionsClientAPI contains the set of methods on the AllowedConnectionsClient type.
type AllowedConnectionsClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, connectionType security.ConnectionType) (result security.AllowedConnectionsResource, err error)
	List(ctx context.Context) (result security.AllowedConnectionsListPage, err error)
	ListComplete(ctx context.Context) (result security.AllowedConnectionsListIterator, err error)
	ListByHomeRegion(ctx context.Context) (result security.AllowedConnectionsListPage, err error)
	ListByHomeRegionComplete(ctx context.Context) (result security.AllowedConnectionsListIterator, err error)
}

var _ AllowedConnectionsClientAPI = (*security.AllowedConnectionsClient)(nil)

// TopologyClientAPI contains the set of methods on the TopologyClient type.
type TopologyClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, topologyResourceName string) (result security.TopologyResource, err error)
	List(ctx context.Context) (result security.TopologyListPage, err error)
	ListComplete(ctx context.Context) (result security.TopologyListIterator, err error)
	ListByHomeRegion(ctx context.Context) (result security.TopologyListPage, err error)
	ListByHomeRegionComplete(ctx context.Context) (result security.TopologyListIterator, err error)
}

var _ TopologyClientAPI = (*security.TopologyClient)(nil)

// JitNetworkAccessPoliciesClientAPI contains the set of methods on the JitNetworkAccessPoliciesClient type.
type JitNetworkAccessPoliciesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, jitNetworkAccessPolicyName string, body security.JitNetworkAccessPolicy) (result security.JitNetworkAccessPolicy, err error)
	Delete(ctx context.Context, resourceGroupName string, jitNetworkAccessPolicyName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, jitNetworkAccessPolicyName string) (result security.JitNetworkAccessPolicy, err error)
	Initiate(ctx context.Context, resourceGroupName string, jitNetworkAccessPolicyName string, body security.JitNetworkAccessPolicyInitiateRequest) (result security.JitNetworkAccessRequest, err error)
	List(ctx context.Context) (result security.JitNetworkAccessPoliciesListPage, err error)
	ListComplete(ctx context.Context) (result security.JitNetworkAccessPoliciesListIterator, err error)
	ListByRegion(ctx context.Context) (result security.JitNetworkAccessPoliciesListPage, err error)
	ListByRegionComplete(ctx context.Context) (result security.JitNetworkAccessPoliciesListIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result security.JitNetworkAccessPoliciesListPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result security.JitNetworkAccessPoliciesListIterator, err error)
	ListByResourceGroupAndRegion(ctx context.Context, resourceGroupName string) (result security.JitNetworkAccessPoliciesListPage, err error)
	ListByResourceGroupAndRegionComplete(ctx context.Context, resourceGroupName string) (result security.JitNetworkAccessPoliciesListIterator, err error)
}

var _ JitNetworkAccessPoliciesClientAPI = (*security.JitNetworkAccessPoliciesClient)(nil)

// DiscoveredSecuritySolutionsClientAPI contains the set of methods on the DiscoveredSecuritySolutionsClient type.
type DiscoveredSecuritySolutionsClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, discoveredSecuritySolutionName string) (result security.DiscoveredSecuritySolution, err error)
	List(ctx context.Context) (result security.DiscoveredSecuritySolutionListPage, err error)
	ListComplete(ctx context.Context) (result security.DiscoveredSecuritySolutionListIterator, err error)
	ListByHomeRegion(ctx context.Context) (result security.DiscoveredSecuritySolutionListPage, err error)
	ListByHomeRegionComplete(ctx context.Context) (result security.DiscoveredSecuritySolutionListIterator, err error)
}

var _ DiscoveredSecuritySolutionsClientAPI = (*security.DiscoveredSecuritySolutionsClient)(nil)

// SolutionsReferenceDataClientAPI contains the set of methods on the SolutionsReferenceDataClient type.
type SolutionsReferenceDataClientAPI interface {
	List(ctx context.Context) (result security.SolutionsReferenceDataList, err error)
	ListByHomeRegion(ctx context.Context) (result security.SolutionsReferenceDataList, err error)
}

var _ SolutionsReferenceDataClientAPI = (*security.SolutionsReferenceDataClient)(nil)

// ExternalSecuritySolutionsClientAPI contains the set of methods on the ExternalSecuritySolutionsClient type.
type ExternalSecuritySolutionsClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, externalSecuritySolutionsName string) (result security.ExternalSecuritySolutionModel, err error)
	List(ctx context.Context) (result security.ExternalSecuritySolutionListPage, err error)
	ListComplete(ctx context.Context) (result security.ExternalSecuritySolutionListIterator, err error)
	ListByHomeRegion(ctx context.Context) (result security.ExternalSecuritySolutionListPage, err error)
	ListByHomeRegionComplete(ctx context.Context) (result security.ExternalSecuritySolutionListIterator, err error)
}

var _ ExternalSecuritySolutionsClientAPI = (*security.ExternalSecuritySolutionsClient)(nil)

// SecureScoresClientAPI contains the set of methods on the SecureScoresClient type.
type SecureScoresClientAPI interface {
	Get(ctx context.Context, secureScoreName string) (result security.SecureScoreItem, err error)
	List(ctx context.Context) (result security.SecureScoresListPage, err error)
	ListComplete(ctx context.Context) (result security.SecureScoresListIterator, err error)
}

var _ SecureScoresClientAPI = (*security.SecureScoresClient)(nil)

// SecureScoreControlsClientAPI contains the set of methods on the SecureScoreControlsClient type.
type SecureScoreControlsClientAPI interface {
	List(ctx context.Context, expand security.ExpandControlsEnum) (result security.SecureScoreControlListPage, err error)
	ListComplete(ctx context.Context, expand security.ExpandControlsEnum) (result security.SecureScoreControlListIterator, err error)
	ListBySecureScore(ctx context.Context, secureScoreName string, expand security.ExpandControlsEnum) (result security.SecureScoreControlListPage, err error)
	ListBySecureScoreComplete(ctx context.Context, secureScoreName string, expand security.ExpandControlsEnum) (result security.SecureScoreControlListIterator, err error)
}

var _ SecureScoreControlsClientAPI = (*security.SecureScoreControlsClient)(nil)

// SecureScoreControlDefinitionsClientAPI contains the set of methods on the SecureScoreControlDefinitionsClient type.
type SecureScoreControlDefinitionsClientAPI interface {
	List(ctx context.Context) (result security.SecureScoreControlDefinitionListPage, err error)
	ListComplete(ctx context.Context) (result security.SecureScoreControlDefinitionListIterator, err error)
	ListBySubscription(ctx context.Context) (result security.SecureScoreControlDefinitionListPage, err error)
	ListBySubscriptionComplete(ctx context.Context) (result security.SecureScoreControlDefinitionListIterator, err error)
}

var _ SecureScoreControlDefinitionsClientAPI = (*security.SecureScoreControlDefinitionsClient)(nil)

// SolutionsClientAPI contains the set of methods on the SolutionsClient type.
type SolutionsClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, securitySolutionName string) (result security.Solution, err error)
	List(ctx context.Context) (result security.SolutionListPage, err error)
	ListComplete(ctx context.Context) (result security.SolutionListIterator, err error)
}

var _ SolutionsClientAPI = (*security.SolutionsClient)(nil)

// ConnectorsClientAPI contains the set of methods on the ConnectorsClient type.
type ConnectorsClientAPI interface {
	CreateOrUpdate(ctx context.Context, connectorName string, connectorSetting security.ConnectorSetting) (result security.ConnectorSetting, err error)
	Delete(ctx context.Context, connectorName string) (result autorest.Response, err error)
	Get(ctx context.Context, connectorName string) (result security.ConnectorSetting, err error)
	List(ctx context.Context) (result security.ConnectorSettingListPage, err error)
	ListComplete(ctx context.Context) (result security.ConnectorSettingListIterator, err error)
}

var _ ConnectorsClientAPI = (*security.ConnectorsClient)(nil)
