package maintenanceapi

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
	"github.com/Azure/azure-sdk-for-go/services/preview/maintenance/mgmt/2018-06-01-preview/maintenance"
)

// ApplyUpdatesClientAPI contains the set of methods on the ApplyUpdatesClient type.
type ApplyUpdatesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, providerName string, resourceType string, resourceName string) (result maintenance.ApplyUpdate, err error)
	CreateOrUpdateParent(ctx context.Context, resourceGroupName string, resourceParentType string, resourceParentName string, providerName string, resourceType string, resourceName string) (result maintenance.ApplyUpdate, err error)
	Get(ctx context.Context, resourceGroupName string, providerName string, resourceType string, resourceName string, applyUpdateName string) (result maintenance.ApplyUpdate, err error)
	GetParent(ctx context.Context, resourceGroupName string, resourceParentType string, resourceParentName string, providerName string, resourceType string, resourceName string, applyUpdateName string) (result maintenance.ApplyUpdate, err error)
}

var _ ApplyUpdatesClientAPI = (*maintenance.ApplyUpdatesClient)(nil)

// ConfigurationAssignmentsClientAPI contains the set of methods on the ConfigurationAssignmentsClient type.
type ConfigurationAssignmentsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, providerName string, resourceType string, resourceName string, configurationAssignmentName string, configurationAssignment maintenance.ConfigurationAssignment) (result maintenance.ConfigurationAssignment, err error)
	CreateOrUpdateParent(ctx context.Context, resourceGroupName string, providerName string, resourceParentType string, resourceParentName string, resourceType string, resourceName string, configurationAssignmentName string, configurationAssignment maintenance.ConfigurationAssignment) (result maintenance.ConfigurationAssignment, err error)
	Delete(ctx context.Context, resourceGroupName string, providerName string, resourceType string, resourceName string, configurationAssignmentName string) (result maintenance.ConfigurationAssignment, err error)
	DeleteParent(ctx context.Context, resourceGroupName string, providerName string, resourceParentType string, resourceParentName string, resourceType string, resourceName string, configurationAssignmentName string) (result maintenance.ConfigurationAssignment, err error)
}

var _ ConfigurationAssignmentsClientAPI = (*maintenance.ConfigurationAssignmentsClient)(nil)

// ConfigurationsClientAPI contains the set of methods on the ConfigurationsClient type.
type ConfigurationsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, configuration maintenance.Configuration) (result maintenance.Configuration, err error)
	Delete(ctx context.Context, resourceGroupName string, resourceName string) (result maintenance.Configuration, err error)
	Get(ctx context.Context, resourceGroupName string, resourceName string) (result maintenance.Configuration, err error)
	List(ctx context.Context) (result maintenance.ListMaintenanceConfigurationsResult, err error)
	UpdateMethod(ctx context.Context, resourceGroupName string, resourceName string, configuration maintenance.Configuration) (result maintenance.Configuration, err error)
}

var _ ConfigurationsClientAPI = (*maintenance.ConfigurationsClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result maintenance.OperationsListResult, err error)
}

var _ OperationsClientAPI = (*maintenance.OperationsClient)(nil)

// UpdatesClientAPI contains the set of methods on the UpdatesClient type.
type UpdatesClientAPI interface {
	List(ctx context.Context, resourceGroupName string, providerName string, resourceType string, resourceName string) (result maintenance.ListUpdatesResult, err error)
	ListParent(ctx context.Context, resourceGroupName string, providerName string, resourceParentType string, resourceParentName string, resourceType string, resourceName string) (result maintenance.ListUpdatesResult, err error)
}

var _ UpdatesClientAPI = (*maintenance.UpdatesClient)(nil)
