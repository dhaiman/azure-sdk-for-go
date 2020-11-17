package powerplatformapi

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
	"github.com/Azure/azure-sdk-for-go/services/preview/powerplatform/mgmt/2020-10-30/powerplatform"
	"github.com/Azure/go-autorest/autorest"
)

// EnterprisePoliciesClientAPI contains the set of methods on the EnterprisePoliciesClient type.
type EnterprisePoliciesClientAPI interface {
	CreateOrUpdate(ctx context.Context, enterprisePolicyName string, resourceGroupName string, parameters powerplatform.EnterprisePolicy) (result powerplatform.EnterprisePolicy, err error)
	Delete(ctx context.Context, resourceGroupName string, enterprisePolicyName string) (result autorest.Response, err error)
	Get(ctx context.Context, enterprisePolicyName string, resourceGroupName string) (result powerplatform.EnterprisePolicy, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result powerplatform.EnterprisePolicyList, err error)
	ListBySubscription(ctx context.Context) (result powerplatform.EnterprisePolicyList, err error)
	Update(ctx context.Context, enterprisePolicyName string, resourceGroupName string, parameters powerplatform.EnterprisePolicy) (result powerplatform.EnterprisePolicy, err error)
}

var _ EnterprisePoliciesClientAPI = (*powerplatform.EnterprisePoliciesClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result powerplatform.OperationList, err error)
}

var _ OperationsClientAPI = (*powerplatform.OperationsClient)(nil)

// PrivateEndpointConnectionsClientAPI contains the set of methods on the PrivateEndpointConnectionsClient type.
type PrivateEndpointConnectionsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, enterprisePolicyName string, privateEndpointConnectionName string, parameters powerplatform.PrivateEndpointConnection) (result powerplatform.PrivateEndpointConnectionsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, enterprisePolicyName string, privateEndpointConnectionName string) (result powerplatform.PrivateEndpointConnectionsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, enterprisePolicyName string, privateEndpointConnectionName string) (result powerplatform.PrivateEndpointConnection, err error)
	ListByEnterprisePolicy(ctx context.Context, resourceGroupName string, enterprisePolicyName string) (result powerplatform.PrivateEndpointConnectionListResult, err error)
}

var _ PrivateEndpointConnectionsClientAPI = (*powerplatform.PrivateEndpointConnectionsClient)(nil)

// PrivateLinkResourcesClientAPI contains the set of methods on the PrivateLinkResourcesClient type.
type PrivateLinkResourcesClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, enterprisePolicyName string, groupName string) (result powerplatform.PrivateLinkResource, err error)
	ListByEnterprisePolicy(ctx context.Context, resourceGroupName string, enterprisePolicyName string) (result powerplatform.PrivateLinkResourceListResult, err error)
}

var _ PrivateLinkResourcesClientAPI = (*powerplatform.PrivateLinkResourcesClient)(nil)

// SubnetsClientAPI contains the set of methods on the SubnetsClient type.
type SubnetsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, enterprisePolicyName string, subnetName string) (result powerplatform.Subnet, err error)
	Get(ctx context.Context, resourceGroupName string, enterprisePolicyName string, subnetName string) (result powerplatform.Subnet, err error)
	ListByEnterprisePolicy(ctx context.Context, resourceGroupName string, enterprisePolicyName string) (result powerplatform.SubnetListResult, err error)
}

var _ SubnetsClientAPI = (*powerplatform.SubnetsClient)(nil)
