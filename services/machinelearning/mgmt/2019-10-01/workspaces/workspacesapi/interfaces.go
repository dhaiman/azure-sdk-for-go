package workspacesapi

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
	"github.com/Azure/azure-sdk-for-go/services/machinelearning/mgmt/2019-10-01/workspaces"
	"github.com/Azure/go-autorest/autorest"
)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result workspaces.OperationListResult, err error)
}

var _ OperationsClientAPI = (*workspaces.OperationsClient)(nil)

// ClientAPI contains the set of methods on the Client type.
type ClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, parameters workspaces.Workspace) (result workspaces.Workspace, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string) (result workspaces.Workspace, err error)
	List(ctx context.Context) (result workspaces.ListResultPage, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result workspaces.ListResultPage, err error)
	ListWorkspaceKeys(ctx context.Context, workspaceName string, resourceGroupName string) (result workspaces.KeysResponse, err error)
	ResyncStorageKeys(ctx context.Context, workspaceName string, resourceGroupName string) (result autorest.Response, err error)
	Update(ctx context.Context, resourceGroupName string, workspaceName string, parameters workspaces.UpdateParameters) (result workspaces.Workspace, err error)
}

var _ ClientAPI = (*workspaces.Client)(nil)
