package featuresapi

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
    "github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2019-07-01/features"
    "github.com/Azure/go-autorest/autorest"
)

        // OperationsClientAPI contains the set of methods on the OperationsClient type.
        type OperationsClientAPI interface {
            List(ctx context.Context) (result features.OperationListResultPage, err error)
                ListComplete(ctx context.Context) (result features.OperationListResultIterator, err error)
        }

        var _ OperationsClientAPI = (*features.OperationsClient)(nil)
        // DeploymentsClientAPI contains the set of methods on the DeploymentsClient type.
        type DeploymentsClientAPI interface {
            CalculateTemplateHash(ctx context.Context, templateParameter interface{}) (result features.TemplateHashResult, err error)
            Cancel(ctx context.Context, resourceGroupName string, deploymentName string) (result autorest.Response, err error)
            CancelAtManagementGroupScope(ctx context.Context, groupID string, deploymentName string) (result autorest.Response, err error)
            CancelAtScope(ctx context.Context, scope string, deploymentName string) (result autorest.Response, err error)
            CancelAtSubscriptionScope(ctx context.Context, deploymentName string) (result autorest.Response, err error)
            CancelAtTenantScope(ctx context.Context, deploymentName string) (result autorest.Response, err error)
            CheckExistence(ctx context.Context, resourceGroupName string, deploymentName string) (result autorest.Response, err error)
            CheckExistenceAtManagementGroupScope(ctx context.Context, groupID string, deploymentName string) (result autorest.Response, err error)
            CheckExistenceAtScope(ctx context.Context, scope string, deploymentName string) (result autorest.Response, err error)
            CheckExistenceAtSubscriptionScope(ctx context.Context, deploymentName string) (result autorest.Response, err error)
            CheckExistenceAtTenantScope(ctx context.Context, deploymentName string) (result autorest.Response, err error)
            CreateOrUpdate(ctx context.Context, resourceGroupName string, deploymentName string, parameters features.Deployment) (result features.DeploymentsCreateOrUpdateFuture, err error)
            CreateOrUpdateAtManagementGroupScope(ctx context.Context, groupID string, deploymentName string, parameters features.Deployment) (result features.DeploymentsCreateOrUpdateAtManagementGroupScopeFuture, err error)
            CreateOrUpdateAtScope(ctx context.Context, scope string, deploymentName string, parameters features.Deployment) (result features.DeploymentsCreateOrUpdateAtScopeFuture, err error)
            CreateOrUpdateAtSubscriptionScope(ctx context.Context, deploymentName string, parameters features.Deployment) (result features.DeploymentsCreateOrUpdateAtSubscriptionScopeFuture, err error)
            CreateOrUpdateAtTenantScope(ctx context.Context, deploymentName string, parameters features.Deployment) (result features.DeploymentsCreateOrUpdateAtTenantScopeFuture, err error)
            Delete(ctx context.Context, resourceGroupName string, deploymentName string) (result features.DeploymentsDeleteFuture, err error)
            DeleteAtManagementGroupScope(ctx context.Context, groupID string, deploymentName string) (result features.DeploymentsDeleteAtManagementGroupScopeFuture, err error)
            DeleteAtScope(ctx context.Context, scope string, deploymentName string) (result features.DeploymentsDeleteAtScopeFuture, err error)
            DeleteAtSubscriptionScope(ctx context.Context, deploymentName string) (result features.DeploymentsDeleteAtSubscriptionScopeFuture, err error)
            DeleteAtTenantScope(ctx context.Context, deploymentName string) (result features.DeploymentsDeleteAtTenantScopeFuture, err error)
            ExportTemplate(ctx context.Context, resourceGroupName string, deploymentName string) (result features.DeploymentExportResult, err error)
            ExportTemplateAtManagementGroupScope(ctx context.Context, groupID string, deploymentName string) (result features.DeploymentExportResult, err error)
            ExportTemplateAtScope(ctx context.Context, scope string, deploymentName string) (result features.DeploymentExportResult, err error)
            ExportTemplateAtSubscriptionScope(ctx context.Context, deploymentName string) (result features.DeploymentExportResult, err error)
            ExportTemplateAtTenantScope(ctx context.Context, deploymentName string) (result features.DeploymentExportResult, err error)
            Get(ctx context.Context, resourceGroupName string, deploymentName string) (result features.DeploymentExtended, err error)
            GetAtManagementGroupScope(ctx context.Context, groupID string, deploymentName string) (result features.DeploymentExtended, err error)
            GetAtScope(ctx context.Context, scope string, deploymentName string) (result features.DeploymentExtended, err error)
            GetAtSubscriptionScope(ctx context.Context, deploymentName string) (result features.DeploymentExtended, err error)
            GetAtTenantScope(ctx context.Context, deploymentName string) (result features.DeploymentExtended, err error)
            ListAtManagementGroupScope(ctx context.Context, groupID string, filter string, top *int32) (result features.DeploymentListResultPage, err error)
                ListAtManagementGroupScopeComplete(ctx context.Context, groupID string, filter string, top *int32) (result features.DeploymentListResultIterator, err error)
            ListAtScope(ctx context.Context, scope string, filter string, top *int32) (result features.DeploymentListResultPage, err error)
                ListAtScopeComplete(ctx context.Context, scope string, filter string, top *int32) (result features.DeploymentListResultIterator, err error)
            ListAtSubscriptionScope(ctx context.Context, filter string, top *int32) (result features.DeploymentListResultPage, err error)
                ListAtSubscriptionScopeComplete(ctx context.Context, filter string, top *int32) (result features.DeploymentListResultIterator, err error)
            ListAtTenantScope(ctx context.Context, filter string, top *int32) (result features.DeploymentListResultPage, err error)
                ListAtTenantScopeComplete(ctx context.Context, filter string, top *int32) (result features.DeploymentListResultIterator, err error)
            ListByResourceGroup(ctx context.Context, resourceGroupName string, filter string, top *int32) (result features.DeploymentListResultPage, err error)
                ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, filter string, top *int32) (result features.DeploymentListResultIterator, err error)
            Validate(ctx context.Context, resourceGroupName string, deploymentName string, parameters features.Deployment) (result features.DeploymentValidateResult, err error)
            ValidateAtManagementGroupScope(ctx context.Context, groupID string, deploymentName string, parameters features.Deployment) (result features.DeploymentValidateResult, err error)
            ValidateAtScope(ctx context.Context, scope string, deploymentName string, parameters features.Deployment) (result features.DeploymentValidateResult, err error)
            ValidateAtSubscriptionScope(ctx context.Context, deploymentName string, parameters features.Deployment) (result features.DeploymentValidateResult, err error)
            ValidateAtTenantScope(ctx context.Context, deploymentName string, parameters features.Deployment) (result features.DeploymentValidateResult, err error)
            WhatIf(ctx context.Context, resourceGroupName string, deploymentName string, parameters features.DeploymentWhatIf) (result features.DeploymentsWhatIfFuture, err error)
            WhatIfAtSubscriptionScope(ctx context.Context, deploymentName string, parameters features.DeploymentWhatIf) (result features.DeploymentsWhatIfAtSubscriptionScopeFuture, err error)
        }

        var _ DeploymentsClientAPI = (*features.DeploymentsClient)(nil)
        // ProvidersClientAPI contains the set of methods on the ProvidersClient type.
        type ProvidersClientAPI interface {
            Get(ctx context.Context, resourceProviderNamespace string, expand string) (result features.Provider, err error)
            GetAtTenantScope(ctx context.Context, resourceProviderNamespace string, expand string) (result features.Provider, err error)
            List(ctx context.Context, top *int32, expand string) (result features.ProviderListResultPage, err error)
                ListComplete(ctx context.Context, top *int32, expand string) (result features.ProviderListResultIterator, err error)
            ListAtTenantScope(ctx context.Context, top *int32, expand string) (result features.ProviderListResultPage, err error)
                ListAtTenantScopeComplete(ctx context.Context, top *int32, expand string) (result features.ProviderListResultIterator, err error)
            Register(ctx context.Context, resourceProviderNamespace string) (result features.Provider, err error)
            Unregister(ctx context.Context, resourceProviderNamespace string) (result features.Provider, err error)
        }

        var _ ProvidersClientAPI = (*features.ProvidersClient)(nil)
        // ResourcesClientAPI contains the set of methods on the ResourcesClient type.
        type ResourcesClientAPI interface {
            CheckExistence(ctx context.Context, resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, APIVersion string) (result autorest.Response, err error)
            CheckExistenceByID(ctx context.Context, resourceID string, APIVersion string) (result autorest.Response, err error)
            CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, APIVersion string, parameters features.GenericResource) (result features.ResourcesCreateOrUpdateFuture, err error)
            CreateOrUpdateByID(ctx context.Context, resourceID string, APIVersion string, parameters features.GenericResource) (result features.ResourcesCreateOrUpdateByIDFuture, err error)
            Delete(ctx context.Context, resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, APIVersion string) (result features.ResourcesDeleteFuture, err error)
            DeleteByID(ctx context.Context, resourceID string, APIVersion string) (result features.ResourcesDeleteByIDFuture, err error)
            Get(ctx context.Context, resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, APIVersion string) (result features.GenericResource, err error)
            GetByID(ctx context.Context, resourceID string, APIVersion string) (result features.GenericResource, err error)
            List(ctx context.Context, filter string, expand string, top *int32) (result features.ResourceListResultPage, err error)
                ListComplete(ctx context.Context, filter string, expand string, top *int32) (result features.ResourceListResultIterator, err error)
            ListByResourceGroup(ctx context.Context, resourceGroupName string, filter string, expand string, top *int32) (result features.ResourceListResultPage, err error)
                ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, filter string, expand string, top *int32) (result features.ResourceListResultIterator, err error)
            MoveResources(ctx context.Context, sourceResourceGroupName string, parameters features.ResourcesMoveInfo) (result features.ResourcesMoveResourcesFuture, err error)
            Update(ctx context.Context, resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, APIVersion string, parameters features.GenericResource) (result features.ResourcesUpdateFuture, err error)
            UpdateByID(ctx context.Context, resourceID string, APIVersion string, parameters features.GenericResource) (result features.ResourcesUpdateByIDFuture, err error)
            ValidateMoveResources(ctx context.Context, sourceResourceGroupName string, parameters features.ResourcesMoveInfo) (result features.ResourcesValidateMoveResourcesFuture, err error)
        }

        var _ ResourcesClientAPI = (*features.ResourcesClient)(nil)
        // ResourceGroupsClientAPI contains the set of methods on the ResourceGroupsClient type.
        type ResourceGroupsClientAPI interface {
            CheckExistence(ctx context.Context, resourceGroupName string) (result autorest.Response, err error)
            CreateOrUpdate(ctx context.Context, resourceGroupName string, parameters features.ResourceGroup) (result features.ResourceGroup, err error)
            Delete(ctx context.Context, resourceGroupName string) (result features.ResourceGroupsDeleteFuture, err error)
            ExportTemplate(ctx context.Context, resourceGroupName string, parameters features.ExportTemplateRequest) (result features.ResourceGroupExportResult, err error)
            Get(ctx context.Context, resourceGroupName string) (result features.ResourceGroup, err error)
            List(ctx context.Context, filter string, top *int32) (result features.ResourceGroupListResultPage, err error)
                ListComplete(ctx context.Context, filter string, top *int32) (result features.ResourceGroupListResultIterator, err error)
            Update(ctx context.Context, resourceGroupName string, parameters features.ResourceGroupPatchable) (result features.ResourceGroup, err error)
        }

        var _ ResourceGroupsClientAPI = (*features.ResourceGroupsClient)(nil)
        // TagsClientAPI contains the set of methods on the TagsClient type.
        type TagsClientAPI interface {
            CreateOrUpdate(ctx context.Context, tagName string) (result features.TagDetails, err error)
            CreateOrUpdateValue(ctx context.Context, tagName string, tagValue string) (result features.TagValue, err error)
            Delete(ctx context.Context, tagName string) (result autorest.Response, err error)
            DeleteValue(ctx context.Context, tagName string, tagValue string) (result autorest.Response, err error)
            List(ctx context.Context) (result features.TagsListResultPage, err error)
                ListComplete(ctx context.Context) (result features.TagsListResultIterator, err error)
        }

        var _ TagsClientAPI = (*features.TagsClient)(nil)
        // DeploymentOperationsClientAPI contains the set of methods on the DeploymentOperationsClient type.
        type DeploymentOperationsClientAPI interface {
            Get(ctx context.Context, resourceGroupName string, deploymentName string, operationID string) (result features.DeploymentOperation, err error)
            GetAtManagementGroupScope(ctx context.Context, groupID string, deploymentName string, operationID string) (result features.DeploymentOperation, err error)
            GetAtScope(ctx context.Context, scope string, deploymentName string, operationID string) (result features.DeploymentOperation, err error)
            GetAtSubscriptionScope(ctx context.Context, deploymentName string, operationID string) (result features.DeploymentOperation, err error)
            GetAtTenantScope(ctx context.Context, deploymentName string, operationID string) (result features.DeploymentOperation, err error)
            List(ctx context.Context, resourceGroupName string, deploymentName string, top *int32) (result features.DeploymentOperationsListResultPage, err error)
                ListComplete(ctx context.Context, resourceGroupName string, deploymentName string, top *int32) (result features.DeploymentOperationsListResultIterator, err error)
            ListAtManagementGroupScope(ctx context.Context, groupID string, deploymentName string, top *int32) (result features.DeploymentOperationsListResultPage, err error)
                ListAtManagementGroupScopeComplete(ctx context.Context, groupID string, deploymentName string, top *int32) (result features.DeploymentOperationsListResultIterator, err error)
            ListAtScope(ctx context.Context, scope string, deploymentName string, top *int32) (result features.DeploymentOperationsListResultPage, err error)
                ListAtScopeComplete(ctx context.Context, scope string, deploymentName string, top *int32) (result features.DeploymentOperationsListResultIterator, err error)
            ListAtSubscriptionScope(ctx context.Context, deploymentName string, top *int32) (result features.DeploymentOperationsListResultPage, err error)
                ListAtSubscriptionScopeComplete(ctx context.Context, deploymentName string, top *int32) (result features.DeploymentOperationsListResultIterator, err error)
            ListAtTenantScope(ctx context.Context, deploymentName string, top *int32) (result features.DeploymentOperationsListResultPage, err error)
                ListAtTenantScopeComplete(ctx context.Context, deploymentName string, top *int32) (result features.DeploymentOperationsListResultIterator, err error)
        }

        var _ DeploymentOperationsClientAPI = (*features.DeploymentOperationsClient)(nil)
