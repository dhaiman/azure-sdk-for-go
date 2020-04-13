package kubernetesconfigurationapi

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
	"github.com/Azure/azure-sdk-for-go/services/preview/kubernetesconfiguration/mgmt/2019-11-01-preview/kubernetesconfiguration"
)

// SourceControlConfigurationsClientAPI contains the set of methods on the SourceControlConfigurationsClient type.
type SourceControlConfigurationsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, sourceControlConfigurationName string, APIVersion string, sourceControlConfiguration kubernetesconfiguration.SourceControlConfiguration) (result kubernetesconfiguration.SourceControlConfiguration, err error)
	Delete(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, sourceControlConfigurationName string, APIVersion string) (result kubernetesconfiguration.SourceControlConfigurationsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, sourceControlConfigurationName string, APIVersion string) (result kubernetesconfiguration.SourceControlConfiguration, err error)
	List(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, APIVersion string) (result kubernetesconfiguration.SourceControlConfigurationListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, clusterRp string, clusterResourceName string, clusterName string, APIVersion string) (result kubernetesconfiguration.SourceControlConfigurationListIterator, err error)
}

var _ SourceControlConfigurationsClientAPI = (*kubernetesconfiguration.SourceControlConfigurationsClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context, APIVersion string) (result kubernetesconfiguration.ResourceProviderOperationListPage, err error)
	ListComplete(ctx context.Context, APIVersion string) (result kubernetesconfiguration.ResourceProviderOperationListIterator, err error)
}

var _ OperationsClientAPI = (*kubernetesconfiguration.OperationsClient)(nil)
