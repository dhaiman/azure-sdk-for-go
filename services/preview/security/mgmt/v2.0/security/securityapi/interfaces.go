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
    "github.com/Azure/azure-sdk-for-go/services/preview/security/mgmt/v2.0/security"
    "github.com/Azure/go-autorest/autorest"
)

        // PlaybookConfigurationsClientAPI contains the set of methods on the PlaybookConfigurationsClient type.
        type PlaybookConfigurationsClientAPI interface {
            CreateOrUpdate(ctx context.Context, resourceGroupName string, playbookConfigurationName string, playbookConfiguration security.PlaybookConfiguration) (result security.PlaybookConfiguration, err error)
            Delete(ctx context.Context, resourceGroupName string, playbookConfigurationName string) (result autorest.Response, err error)
            Get(ctx context.Context, resourceGroupName string, playbookConfigurationName string) (result security.PlaybookConfiguration, err error)
            List(ctx context.Context) (result security.PlaybookConfigurationListPage, err error)
            ListByResourceGroup(ctx context.Context, resourceGroupName string) (result security.PlaybookConfigurationListPage, err error)
            Patch(ctx context.Context, resourceGroupName string, playbookConfigurationName string, playbookConfiguration security.PlaybookConfiguration) (result security.PlaybookConfiguration, err error)
            Validate(ctx context.Context, resourceGroupName string, playbookConfigurationName string, playbookConfiguration security.PlaybookConfiguration) (result security.PlaybookConfigurationValidationStatus, err error)
        }

        var _ PlaybookConfigurationsClientAPI = (*security.PlaybookConfigurationsClient)(nil)
        // RegulatoryComplianceStandardsClientAPI contains the set of methods on the RegulatoryComplianceStandardsClient type.
        type RegulatoryComplianceStandardsClientAPI interface {
            Get(ctx context.Context, regulatoryComplianceStandardName string) (result security.RegulatoryComplianceStandard, err error)
            List(ctx context.Context, filter string) (result security.RegulatoryComplianceStandardListPage, err error)
        }

        var _ RegulatoryComplianceStandardsClientAPI = (*security.RegulatoryComplianceStandardsClient)(nil)
        // RegulatoryComplianceControlsClientAPI contains the set of methods on the RegulatoryComplianceControlsClient type.
        type RegulatoryComplianceControlsClientAPI interface {
            Get(ctx context.Context, regulatoryComplianceStandardName string, regulatoryComplianceControlName string) (result security.RegulatoryComplianceControl, err error)
            List(ctx context.Context, regulatoryComplianceStandardName string, filter string) (result security.RegulatoryComplianceControlListPage, err error)
        }

        var _ RegulatoryComplianceControlsClientAPI = (*security.RegulatoryComplianceControlsClient)(nil)
        // RegulatoryComplianceAssessmentsClientAPI contains the set of methods on the RegulatoryComplianceAssessmentsClient type.
        type RegulatoryComplianceAssessmentsClientAPI interface {
            Get(ctx context.Context, regulatoryComplianceStandardName string, regulatoryComplianceControlName string, regulatoryComplianceAssessmentName string) (result security.RegulatoryComplianceAssessment, err error)
            List(ctx context.Context, regulatoryComplianceStandardName string, regulatoryComplianceControlName string, filter string) (result security.RegulatoryComplianceAssessmentListPage, err error)
        }

        var _ RegulatoryComplianceAssessmentsClientAPI = (*security.RegulatoryComplianceAssessmentsClient)(nil)
        // PricingsClientAPI contains the set of methods on the PricingsClient type.
        type PricingsClientAPI interface {
            Get(ctx context.Context, pricingName string) (result security.Pricing, err error)
            List(ctx context.Context) (result security.PricingList, err error)
            Update(ctx context.Context, pricingName string, pricing security.Pricing) (result security.Pricing, err error)
        }

        var _ PricingsClientAPI = (*security.PricingsClient)(nil)
        // ContactsClientAPI contains the set of methods on the ContactsClient type.
        type ContactsClientAPI interface {
            Create(ctx context.Context, securityContactName string, securityContact security.Contact) (result security.Contact, err error)
            Delete(ctx context.Context, securityContactName string) (result autorest.Response, err error)
            Get(ctx context.Context, securityContactName string) (result security.Contact, err error)
            List(ctx context.Context) (result security.ContactListPage, err error)
            Update(ctx context.Context, securityContactName string, securityContact security.Contact) (result security.Contact, err error)
        }

        var _ ContactsClientAPI = (*security.ContactsClient)(nil)
        // WorkspaceSettingsClientAPI contains the set of methods on the WorkspaceSettingsClient type.
        type WorkspaceSettingsClientAPI interface {
            Create(ctx context.Context, workspaceSettingName string, workspaceSetting security.WorkspaceSetting) (result security.WorkspaceSetting, err error)
            Delete(ctx context.Context, workspaceSettingName string) (result autorest.Response, err error)
            Get(ctx context.Context, workspaceSettingName string) (result security.WorkspaceSetting, err error)
            List(ctx context.Context) (result security.WorkspaceSettingListPage, err error)
            Update(ctx context.Context, workspaceSettingName string, workspaceSetting security.WorkspaceSetting) (result security.WorkspaceSetting, err error)
        }

        var _ WorkspaceSettingsClientAPI = (*security.WorkspaceSettingsClient)(nil)
        // AutoProvisioningSettingsClientAPI contains the set of methods on the AutoProvisioningSettingsClient type.
        type AutoProvisioningSettingsClientAPI interface {
            Create(ctx context.Context, settingName string, setting security.AutoProvisioningSetting) (result security.AutoProvisioningSetting, err error)
            Get(ctx context.Context, settingName string) (result security.AutoProvisioningSetting, err error)
            List(ctx context.Context) (result security.AutoProvisioningSettingListPage, err error)
        }

        var _ AutoProvisioningSettingsClientAPI = (*security.AutoProvisioningSettingsClient)(nil)
        // CompliancesClientAPI contains the set of methods on the CompliancesClient type.
        type CompliancesClientAPI interface {
            Get(ctx context.Context, scope string, complianceName string) (result security.Compliance, err error)
            List(ctx context.Context, scope string) (result security.ComplianceListPage, err error)
        }

        var _ CompliancesClientAPI = (*security.CompliancesClient)(nil)
        // AdvancedThreatProtectionClientAPI contains the set of methods on the AdvancedThreatProtectionClient type.
        type AdvancedThreatProtectionClientAPI interface {
            Create(ctx context.Context, resourceID string, advancedThreatProtectionSetting security.AdvancedThreatProtectionSetting) (result security.AdvancedThreatProtectionSetting, err error)
            Get(ctx context.Context, resourceID string) (result security.AdvancedThreatProtectionSetting, err error)
        }

        var _ AdvancedThreatProtectionClientAPI = (*security.AdvancedThreatProtectionClient)(nil)
        // DeviceSecurityGroupsClientAPI contains the set of methods on the DeviceSecurityGroupsClient type.
        type DeviceSecurityGroupsClientAPI interface {
            CreateOrUpdate(ctx context.Context, resourceID string, deviceSecurityGroupName string, deviceSecurityGroup security.DeviceSecurityGroup) (result security.DeviceSecurityGroup, err error)
            Delete(ctx context.Context, resourceID string, deviceSecurityGroupName string) (result autorest.Response, err error)
            Get(ctx context.Context, resourceID string, deviceSecurityGroupName string) (result security.DeviceSecurityGroup, err error)
            List(ctx context.Context, resourceID string) (result security.DeviceSecurityGroupListPage, err error)
        }

        var _ DeviceSecurityGroupsClientAPI = (*security.DeviceSecurityGroupsClient)(nil)
        // SettingsClientAPI contains the set of methods on the SettingsClient type.
        type SettingsClientAPI interface {
            Get(ctx context.Context, settingName string) (result security.Setting, err error)
            List(ctx context.Context) (result security.SettingsListPage, err error)
            Update(ctx context.Context, settingName string, setting security.Setting) (result security.Setting, err error)
        }

        var _ SettingsClientAPI = (*security.SettingsClient)(nil)
        // InformationProtectionPoliciesClientAPI contains the set of methods on the InformationProtectionPoliciesClient type.
        type InformationProtectionPoliciesClientAPI interface {
            CreateOrUpdate(ctx context.Context, scope string, informationProtectionPolicyName string) (result security.InformationProtectionPolicy, err error)
            Get(ctx context.Context, scope string, informationProtectionPolicyName string) (result security.InformationProtectionPolicy, err error)
            List(ctx context.Context, scope string) (result security.InformationProtectionPolicyListPage, err error)
        }

        var _ InformationProtectionPoliciesClientAPI = (*security.InformationProtectionPoliciesClient)(nil)
        // OperationsClientAPI contains the set of methods on the OperationsClient type.
        type OperationsClientAPI interface {
            List(ctx context.Context) (result security.OperationListPage, err error)
        }

        var _ OperationsClientAPI = (*security.OperationsClient)(nil)
        // LocationsClientAPI contains the set of methods on the LocationsClient type.
        type LocationsClientAPI interface {
            Get(ctx context.Context) (result security.AscLocation, err error)
            List(ctx context.Context) (result security.AscLocationListPage, err error)
        }

        var _ LocationsClientAPI = (*security.LocationsClient)(nil)
        // TasksClientAPI contains the set of methods on the TasksClient type.
        type TasksClientAPI interface {
            GetResourceGroupLevelTask(ctx context.Context, resourceGroupName string, taskName string) (result security.Task, err error)
            GetSubscriptionLevelTask(ctx context.Context, taskName string) (result security.Task, err error)
            List(ctx context.Context, filter string) (result security.TaskListPage, err error)
            ListByHomeRegion(ctx context.Context, filter string) (result security.TaskListPage, err error)
            ListByResourceGroup(ctx context.Context, resourceGroupName string, filter string) (result security.TaskListPage, err error)
            UpdateResourceGroupLevelTaskState(ctx context.Context, resourceGroupName string, taskName string, taskUpdateActionType string) (result autorest.Response, err error)
            UpdateSubscriptionLevelTaskState(ctx context.Context, taskName string, taskUpdateActionType string) (result autorest.Response, err error)
        }

        var _ TasksClientAPI = (*security.TasksClient)(nil)
        // AlertsClientAPI contains the set of methods on the AlertsClient type.
        type AlertsClientAPI interface {
            GetResourceGroupLevelAlerts(ctx context.Context, alertName string, resourceGroupName string) (result security.Alert, err error)
            GetSubscriptionLevelAlert(ctx context.Context, alertName string) (result security.Alert, err error)
            List(ctx context.Context, filter string, selectParameter string, expand string) (result security.AlertListPage, err error)
            ListByResourceGroup(ctx context.Context, resourceGroupName string, filter string, selectParameter string, expand string) (result security.AlertListPage, err error)
            ListResourceGroupLevelAlertsByRegion(ctx context.Context, resourceGroupName string, filter string, selectParameter string, expand string) (result security.AlertListPage, err error)
            ListSubscriptionLevelAlertsByRegion(ctx context.Context, filter string, selectParameter string, expand string) (result security.AlertListPage, err error)
            UpdateResourceGroupLevelAlertState(ctx context.Context, alertName string, alertUpdateActionType string, resourceGroupName string) (result autorest.Response, err error)
            UpdateSubscriptionLevelAlertState(ctx context.Context, alertName string, alertUpdateActionType string) (result autorest.Response, err error)
        }

        var _ AlertsClientAPI = (*security.AlertsClient)(nil)
        // DiscoveredSecuritySolutionsClientAPI contains the set of methods on the DiscoveredSecuritySolutionsClient type.
        type DiscoveredSecuritySolutionsClientAPI interface {
            Get(ctx context.Context, resourceGroupName string, discoveredSecuritySolutionName string) (result security.DiscoveredSecuritySolution, err error)
            List(ctx context.Context) (result security.DiscoveredSecuritySolutionListPage, err error)
            ListByHomeRegion(ctx context.Context) (result security.DiscoveredSecuritySolutionListPage, err error)
        }

        var _ DiscoveredSecuritySolutionsClientAPI = (*security.DiscoveredSecuritySolutionsClient)(nil)
        // JitNetworkAccessPoliciesClientAPI contains the set of methods on the JitNetworkAccessPoliciesClient type.
        type JitNetworkAccessPoliciesClientAPI interface {
            CreateOrUpdate(ctx context.Context, resourceGroupName string, jitNetworkAccessPolicyName string, body security.JitNetworkAccessPolicy) (result security.JitNetworkAccessPolicy, err error)
            Delete(ctx context.Context, resourceGroupName string, jitNetworkAccessPolicyName string) (result autorest.Response, err error)
            Get(ctx context.Context, resourceGroupName string, jitNetworkAccessPolicyName string) (result security.JitNetworkAccessPolicy, err error)
            Initiate(ctx context.Context, resourceGroupName string, jitNetworkAccessPolicyName string, body security.JitNetworkAccessPolicyInitiateRequest) (result security.JitNetworkAccessRequest, err error)
            List(ctx context.Context) (result security.JitNetworkAccessPoliciesListPage, err error)
            ListByRegion(ctx context.Context) (result security.JitNetworkAccessPoliciesListPage, err error)
            ListByResourceGroup(ctx context.Context, resourceGroupName string) (result security.JitNetworkAccessPoliciesListPage, err error)
            ListByResourceGroupAndRegion(ctx context.Context, resourceGroupName string) (result security.JitNetworkAccessPoliciesListPage, err error)
        }

        var _ JitNetworkAccessPoliciesClientAPI = (*security.JitNetworkAccessPoliciesClient)(nil)
        // ExternalSecuritySolutionsClientAPI contains the set of methods on the ExternalSecuritySolutionsClient type.
        type ExternalSecuritySolutionsClientAPI interface {
            Get(ctx context.Context, resourceGroupName string, externalSecuritySolutionsName string) (result security.ExternalSecuritySolutionModel, err error)
            List(ctx context.Context) (result security.ExternalSecuritySolutionListPage, err error)
            ListByHomeRegion(ctx context.Context) (result security.ExternalSecuritySolutionListPage, err error)
        }

        var _ ExternalSecuritySolutionsClientAPI = (*security.ExternalSecuritySolutionsClient)(nil)
        // TopologyClientAPI contains the set of methods on the TopologyClient type.
        type TopologyClientAPI interface {
            Get(ctx context.Context, resourceGroupName string, topologyResourceName string) (result security.TopologyResource, err error)
            List(ctx context.Context) (result security.TopologyListPage, err error)
            ListByHomeRegion(ctx context.Context) (result security.TopologyListPage, err error)
        }

        var _ TopologyClientAPI = (*security.TopologyClient)(nil)
        // AllowedConnectionsClientAPI contains the set of methods on the AllowedConnectionsClient type.
        type AllowedConnectionsClientAPI interface {
            Get(ctx context.Context, resourceGroupName string, connectionType security.ConnectionType) (result security.AllowedConnectionsResource, err error)
            List(ctx context.Context) (result security.AllowedConnectionsListPage, err error)
            ListByHomeRegion(ctx context.Context) (result security.AllowedConnectionsListPage, err error)
        }

        var _ AllowedConnectionsClientAPI = (*security.AllowedConnectionsClient)(nil)
        // AdaptiveNetworkHardeningsClientAPI contains the set of methods on the AdaptiveNetworkHardeningsClient type.
        type AdaptiveNetworkHardeningsClientAPI interface {
            Enforce(ctx context.Context, resourceGroupName string, resourceNamespace string, resourceType string, resourceName string, adaptiveNetworkHardeningResourceName string, body security.AdaptiveNetworkHardeningEnforceRequest) (result security.AdaptiveNetworkHardeningsEnforceFuture, err error)
            Get(ctx context.Context, resourceGroupName string, resourceNamespace string, resourceType string, resourceName string, adaptiveNetworkHardeningResourceName string) (result security.AdaptiveNetworkHardening, err error)
            ListByExtendedResource(ctx context.Context, resourceGroupName string, resourceNamespace string, resourceType string, resourceName string) (result security.AdaptiveNetworkHardeningsListPage, err error)
        }

        var _ AdaptiveNetworkHardeningsClientAPI = (*security.AdaptiveNetworkHardeningsClient)(nil)
