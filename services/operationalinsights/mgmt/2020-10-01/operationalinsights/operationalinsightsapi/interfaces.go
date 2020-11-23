package operationalinsightsapi

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
	"github.com/Azure/azure-sdk-for-go/services/operationalinsights/mgmt/2020-10-01/operationalinsights"
)

// ClustersClientAPI contains the set of methods on the ClustersClient type.
type ClustersClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, parameters operationalinsights.Cluster) (result operationalinsights.ClustersCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, clusterName string) (result operationalinsights.ClustersDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, clusterName string) (result operationalinsights.Cluster, err error)
	List(ctx context.Context) (result operationalinsights.ClusterListResultPage, err error)
	ListComplete(ctx context.Context) (result operationalinsights.ClusterListResultIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result operationalinsights.ClusterListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result operationalinsights.ClusterListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, clusterName string, parameters operationalinsights.ClusterPatch) (result operationalinsights.Cluster, err error)
}

var _ ClustersClientAPI = (*operationalinsights.ClustersClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result operationalinsights.OperationListResultPage, err error)
	ListComplete(ctx context.Context) (result operationalinsights.OperationListResultIterator, err error)
}

var _ OperationsClientAPI = (*operationalinsights.OperationsClient)(nil)

// WorkspacesClientAPI contains the set of methods on the WorkspacesClient type.
type WorkspacesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, parameters operationalinsights.Workspace) (result operationalinsights.WorkspacesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string, force *bool) (result operationalinsights.WorkspacesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string) (result operationalinsights.Workspace, err error)
	List(ctx context.Context) (result operationalinsights.WorkspaceListResult, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result operationalinsights.WorkspaceListResult, err error)
	Update(ctx context.Context, resourceGroupName string, workspaceName string, parameters operationalinsights.WorkspacePatch) (result operationalinsights.Workspace, err error)
}

var _ WorkspacesClientAPI = (*operationalinsights.WorkspacesClient)(nil)

// DeletedWorkspacesClientAPI contains the set of methods on the DeletedWorkspacesClient type.
type DeletedWorkspacesClientAPI interface {
	List(ctx context.Context) (result operationalinsights.WorkspaceListResult, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result operationalinsights.WorkspaceListResult, err error)
}

var _ DeletedWorkspacesClientAPI = (*operationalinsights.DeletedWorkspacesClient)(nil)
