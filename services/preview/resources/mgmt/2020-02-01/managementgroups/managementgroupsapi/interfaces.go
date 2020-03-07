package managementgroupsapi

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
	"github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2020-02-01/managementgroups"
	"github.com/Azure/go-autorest/autorest"
)

// BaseClientAPI contains the set of methods on the BaseClient type.
type BaseClientAPI interface {
	CheckNameAvailability(ctx context.Context, checkNameAvailabilityRequest managementgroups.CheckNameAvailabilityRequest) (result managementgroups.CheckNameAvailabilityResult, err error)
	StartTenantBackfill(ctx context.Context) (result managementgroups.TenantBackfillStatusResult, err error)
	TenantBackfillStatus(ctx context.Context) (result managementgroups.TenantBackfillStatusResult, err error)
}

var _ BaseClientAPI = (*managementgroups.BaseClient)(nil)

// ClientAPI contains the set of methods on the Client type.
type ClientAPI interface {
	CreateOrUpdate(ctx context.Context, groupID string, createManagementGroupRequest managementgroups.CreateManagementGroupRequest, cacheControl string) (result managementgroups.CreateOrUpdateFuture, err error)
	Delete(ctx context.Context, groupID string, cacheControl string) (result managementgroups.DeleteFuture, err error)
	Get(ctx context.Context, groupID string, expand string, recurse *bool, filter string, cacheControl string) (result managementgroups.ManagementGroup, err error)
	GetDescendants(ctx context.Context, groupID string) (result managementgroups.DescendantListResultPage, err error)
	GetDescendantsComplete(ctx context.Context, groupID string) (result managementgroups.DescendantListResultIterator, err error)
	List(ctx context.Context, cacheControl string) (result managementgroups.ListResultPage, err error)
	ListComplete(ctx context.Context, cacheControl string) (result managementgroups.ListResultIterator, err error)
	Update(ctx context.Context, groupID string, patchGroupRequest managementgroups.PatchManagementGroupRequest, cacheControl string) (result managementgroups.ManagementGroup, err error)
}

var _ ClientAPI = (*managementgroups.Client)(nil)

// SubscriptionsClientAPI contains the set of methods on the SubscriptionsClient type.
type SubscriptionsClientAPI interface {
	Create(ctx context.Context, groupID string, subscriptionID string, cacheControl string) (result autorest.Response, err error)
	Delete(ctx context.Context, groupID string, subscriptionID string, cacheControl string) (result autorest.Response, err error)
}

var _ SubscriptionsClientAPI = (*managementgroups.SubscriptionsClient)(nil)

// HierarchySettingsClientAPI contains the set of methods on the HierarchySettingsClient type.
type HierarchySettingsClientAPI interface {
	CreateOrUpdate(ctx context.Context, groupID string, createTenantSettingsRequest managementgroups.CreateOrUpdateSettingsRequest) (result managementgroups.HierarchySettings, err error)
	Delete(ctx context.Context, groupID string) (result autorest.Response, err error)
	Get(ctx context.Context, groupID string) (result managementgroups.HierarchySettings, err error)
	List(ctx context.Context, groupID string) (result managementgroups.HierarchySettingsList, err error)
	Update(ctx context.Context, groupID string, createTenantSettingsRequest managementgroups.CreateOrUpdateSettingsRequest) (result managementgroups.HierarchySettings, err error)
}

var _ HierarchySettingsClientAPI = (*managementgroups.HierarchySettingsClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result managementgroups.OperationListResultPage, err error)
	ListComplete(ctx context.Context) (result managementgroups.OperationListResultIterator, err error)
}

var _ OperationsClientAPI = (*managementgroups.OperationsClient)(nil)

// EntitiesClientAPI contains the set of methods on the EntitiesClient type.
type EntitiesClientAPI interface {
	List(ctx context.Context, selectParameter string, search string, filter string, view string, groupName string, cacheControl string) (result managementgroups.EntityListResultPage, err error)
	ListComplete(ctx context.Context, selectParameter string, search string, filter string, view string, groupName string, cacheControl string) (result managementgroups.EntityListResultIterator, err error)
}

var _ EntitiesClientAPI = (*managementgroups.EntitiesClient)(nil)
