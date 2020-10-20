package policyapi

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
	"github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2020-03-01-preview/policy"
	"github.com/Azure/go-autorest/autorest"
)

// AssignmentsClientAPI contains the set of methods on the AssignmentsClient type.
type AssignmentsClientAPI interface {
	Create(ctx context.Context, scope string, policyAssignmentName string, parameters policy.Assignment) (result policy.Assignment, err error)
	CreateByID(ctx context.Context, policyAssignmentID string, parameters policy.Assignment) (result policy.Assignment, err error)
	Delete(ctx context.Context, scope string, policyAssignmentName string) (result policy.Assignment, err error)
	DeleteByID(ctx context.Context, policyAssignmentID string) (result policy.Assignment, err error)
	Get(ctx context.Context, scope string, policyAssignmentName string) (result policy.Assignment, err error)
	GetByID(ctx context.Context, policyAssignmentID string) (result policy.Assignment, err error)
	List(ctx context.Context, subscriptionID string, filter string, top *int32) (result policy.AssignmentListResultPage, err error)
	ListComplete(ctx context.Context, subscriptionID string, filter string, top *int32) (result policy.AssignmentListResultIterator, err error)
	ListForManagementGroup(ctx context.Context, managementGroupID string, filter string, top *int32) (result policy.AssignmentListResultPage, err error)
	ListForManagementGroupComplete(ctx context.Context, managementGroupID string, filter string, top *int32) (result policy.AssignmentListResultIterator, err error)
	ListForResource(ctx context.Context, resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, subscriptionID string, filter string, top *int32) (result policy.AssignmentListResultPage, err error)
	ListForResourceComplete(ctx context.Context, resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, subscriptionID string, filter string, top *int32) (result policy.AssignmentListResultIterator, err error)
	ListForResourceGroup(ctx context.Context, resourceGroupName string, subscriptionID string, filter string, top *int32) (result policy.AssignmentListResultPage, err error)
	ListForResourceGroupComplete(ctx context.Context, resourceGroupName string, subscriptionID string, filter string, top *int32) (result policy.AssignmentListResultIterator, err error)
}

var _ AssignmentsClientAPI = (*policy.AssignmentsClient)(nil)

// DefinitionsClientAPI contains the set of methods on the DefinitionsClient type.
type DefinitionsClientAPI interface {
	CreateOrUpdate(ctx context.Context, policyDefinitionName string, parameters policy.Definition, subscriptionID string) (result policy.Definition, err error)
	CreateOrUpdateAtManagementGroup(ctx context.Context, policyDefinitionName string, parameters policy.Definition, managementGroupID string) (result policy.Definition, err error)
	Delete(ctx context.Context, policyDefinitionName string, subscriptionID string) (result autorest.Response, err error)
	DeleteAtManagementGroup(ctx context.Context, policyDefinitionName string, managementGroupID string) (result autorest.Response, err error)
	Get(ctx context.Context, policyDefinitionName string, subscriptionID string) (result policy.Definition, err error)
	GetAtManagementGroup(ctx context.Context, policyDefinitionName string, managementGroupID string) (result policy.Definition, err error)
	GetBuiltIn(ctx context.Context, policyDefinitionName string) (result policy.Definition, err error)
	List(ctx context.Context, subscriptionID string, filter string, top *int32) (result policy.DefinitionListResultPage, err error)
	ListComplete(ctx context.Context, subscriptionID string, filter string, top *int32) (result policy.DefinitionListResultIterator, err error)
	ListBuiltIn(ctx context.Context, filter string, top *int32) (result policy.DefinitionListResultPage, err error)
	ListBuiltInComplete(ctx context.Context, filter string, top *int32) (result policy.DefinitionListResultIterator, err error)
	ListByManagementGroup(ctx context.Context, managementGroupID string, filter string, top *int32) (result policy.DefinitionListResultPage, err error)
	ListByManagementGroupComplete(ctx context.Context, managementGroupID string, filter string, top *int32) (result policy.DefinitionListResultIterator, err error)
}

var _ DefinitionsClientAPI = (*policy.DefinitionsClient)(nil)

// SetDefinitionsClientAPI contains the set of methods on the SetDefinitionsClient type.
type SetDefinitionsClientAPI interface {
	CreateOrUpdate(ctx context.Context, policySetDefinitionName string, parameters policy.SetDefinition, subscriptionID string) (result policy.SetDefinition, err error)
	CreateOrUpdateAtManagementGroup(ctx context.Context, policySetDefinitionName string, parameters policy.SetDefinition, managementGroupID string) (result policy.SetDefinition, err error)
	Delete(ctx context.Context, policySetDefinitionName string, subscriptionID string) (result autorest.Response, err error)
	DeleteAtManagementGroup(ctx context.Context, policySetDefinitionName string, managementGroupID string) (result autorest.Response, err error)
	Get(ctx context.Context, policySetDefinitionName string, subscriptionID string) (result policy.SetDefinition, err error)
	GetAtManagementGroup(ctx context.Context, policySetDefinitionName string, managementGroupID string) (result policy.SetDefinition, err error)
	GetBuiltIn(ctx context.Context, policySetDefinitionName string) (result policy.SetDefinition, err error)
	List(ctx context.Context, subscriptionID string, filter string, top *int32) (result policy.SetDefinitionListResultPage, err error)
	ListComplete(ctx context.Context, subscriptionID string, filter string, top *int32) (result policy.SetDefinitionListResultIterator, err error)
	ListBuiltIn(ctx context.Context, filter string, top *int32) (result policy.SetDefinitionListResultPage, err error)
	ListBuiltInComplete(ctx context.Context, filter string, top *int32) (result policy.SetDefinitionListResultIterator, err error)
	ListByManagementGroup(ctx context.Context, managementGroupID string, filter string, top *int32) (result policy.SetDefinitionListResultPage, err error)
	ListByManagementGroupComplete(ctx context.Context, managementGroupID string, filter string, top *int32) (result policy.SetDefinitionListResultIterator, err error)
}

var _ SetDefinitionsClientAPI = (*policy.SetDefinitionsClient)(nil)

// ExemptionsClientAPI contains the set of methods on the ExemptionsClient type.
type ExemptionsClientAPI interface {
	CreateOrUpdate(ctx context.Context, parameters policy.Exemption, scope string, policyExemptionName string) (result policy.Exemption, err error)
	Delete(ctx context.Context, scope string, policyExemptionName string) (result autorest.Response, err error)
	Get(ctx context.Context, scope string, policyExemptionName string) (result policy.Exemption, err error)
	List(ctx context.Context, subscriptionID string, filter string) (result policy.ExemptionListResultPage, err error)
	ListComplete(ctx context.Context, subscriptionID string, filter string) (result policy.ExemptionListResultIterator, err error)
	ListForManagementGroup(ctx context.Context, managementGroupID string, filter string) (result policy.ExemptionListResultPage, err error)
	ListForManagementGroupComplete(ctx context.Context, managementGroupID string, filter string) (result policy.ExemptionListResultIterator, err error)
	ListForResource(ctx context.Context, subscriptionID string, resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, filter string) (result policy.ExemptionListResultPage, err error)
	ListForResourceComplete(ctx context.Context, subscriptionID string, resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, filter string) (result policy.ExemptionListResultIterator, err error)
	ListForResourceGroup(ctx context.Context, subscriptionID string, resourceGroupName string, filter string) (result policy.ExemptionListResultPage, err error)
	ListForResourceGroupComplete(ctx context.Context, subscriptionID string, resourceGroupName string, filter string) (result policy.ExemptionListResultIterator, err error)
}

var _ ExemptionsClientAPI = (*policy.ExemptionsClient)(nil)
