package kubernetesapi

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
	"github.com/Azure/azure-sdk-for-go/services/preview/kubernetes/mgmt/2019-11-01-preview/kubernetes"
	"github.com/Azure/go-autorest/autorest"
)

// ConnectedClusterClientAPI contains the set of methods on the ConnectedClusterClient type.
type ConnectedClusterClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, clusterName string, connectedCluster kubernetes.ConnectedCluster) (result kubernetes.ConnectedCluster, err error)
	Delete(ctx context.Context, resourceGroupName string, clusterName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, clusterName string) (result kubernetes.ConnectedCluster, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result kubernetes.ConnectedClusterListPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result kubernetes.ConnectedClusterListIterator, err error)
	ListBySubscription(ctx context.Context) (result kubernetes.ConnectedClusterListPage, err error)
	ListBySubscriptionComplete(ctx context.Context) (result kubernetes.ConnectedClusterListIterator, err error)
	ListCredentials(ctx context.Context, resourceGroupName string, clusterName string) (result kubernetes.ConnectedClusterAccessProfile, err error)
	Update(ctx context.Context, resourceGroupName string, clusterName string, connectedClusterPatch kubernetes.ConnectedClusterPatch) (result kubernetes.ConnectedCluster, err error)
}

var _ ConnectedClusterClientAPI = (*kubernetes.ConnectedClusterClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	Get(ctx context.Context) (result kubernetes.OperationListPage, err error)
	GetComplete(ctx context.Context) (result kubernetes.OperationListIterator, err error)
}

var _ OperationsClientAPI = (*kubernetes.OperationsClient)(nil)
