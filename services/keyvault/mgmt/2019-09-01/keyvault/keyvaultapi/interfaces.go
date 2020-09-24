package keyvaultapi

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
	"github.com/Azure/azure-sdk-for-go/services/keyvault/mgmt/2019-09-01/keyvault"
	"github.com/Azure/go-autorest/autorest"
)

// VaultsClientAPI contains the set of methods on the VaultsClient type.
type VaultsClientAPI interface {
	CheckNameAvailability(ctx context.Context, vaultName keyvault.VaultCheckNameAvailabilityParameters) (result keyvault.CheckNameAvailabilityResult, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, vaultName string, parameters keyvault.VaultCreateOrUpdateParameters) (result keyvault.VaultsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, vaultName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, vaultName string) (result keyvault.Vault, err error)
	GetDeleted(ctx context.Context, vaultName string, location string) (result keyvault.DeletedVault, err error)
	List(ctx context.Context, top *int32) (result keyvault.ResourceListResultPage, err error)
	ListComplete(ctx context.Context, top *int32) (result keyvault.ResourceListResultIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string, top *int32) (result keyvault.VaultListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, top *int32) (result keyvault.VaultListResultIterator, err error)
	ListBySubscription(ctx context.Context, top *int32) (result keyvault.VaultListResultPage, err error)
	ListBySubscriptionComplete(ctx context.Context, top *int32) (result keyvault.VaultListResultIterator, err error)
	ListDeleted(ctx context.Context) (result keyvault.DeletedVaultListResultPage, err error)
	ListDeletedComplete(ctx context.Context) (result keyvault.DeletedVaultListResultIterator, err error)
	PurgeDeleted(ctx context.Context, vaultName string, location string) (result keyvault.VaultsPurgeDeletedFuture, err error)
	Update(ctx context.Context, resourceGroupName string, vaultName string, parameters keyvault.VaultPatchParameters) (result keyvault.Vault, err error)
	UpdateAccessPolicy(ctx context.Context, resourceGroupName string, vaultName string, operationKind keyvault.AccessPolicyUpdateKind, parameters keyvault.VaultAccessPolicyParameters) (result keyvault.VaultAccessPolicyParameters, err error)
}

var _ VaultsClientAPI = (*keyvault.VaultsClient)(nil)

// PrivateEndpointConnectionsClientAPI contains the set of methods on the PrivateEndpointConnectionsClient type.
type PrivateEndpointConnectionsClientAPI interface {
	Delete(ctx context.Context, resourceGroupName string, vaultName string, privateEndpointConnectionName string) (result keyvault.PrivateEndpointConnectionsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, vaultName string, privateEndpointConnectionName string) (result keyvault.PrivateEndpointConnection, err error)
	Put(ctx context.Context, resourceGroupName string, vaultName string, privateEndpointConnectionName string, properties keyvault.PrivateEndpointConnection) (result keyvault.PrivateEndpointConnection, err error)
}

var _ PrivateEndpointConnectionsClientAPI = (*keyvault.PrivateEndpointConnectionsClient)(nil)

// PrivateLinkResourcesClientAPI contains the set of methods on the PrivateLinkResourcesClient type.
type PrivateLinkResourcesClientAPI interface {
	ListByVault(ctx context.Context, resourceGroupName string, vaultName string) (result keyvault.PrivateLinkResourceListResult, err error)
}

var _ PrivateLinkResourcesClientAPI = (*keyvault.PrivateLinkResourcesClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result keyvault.OperationListResultPage, err error)
	ListComplete(ctx context.Context) (result keyvault.OperationListResultIterator, err error)
}

var _ OperationsClientAPI = (*keyvault.OperationsClient)(nil)

// SecretsClientAPI contains the set of methods on the SecretsClient type.
type SecretsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, vaultName string, secretName string, parameters keyvault.SecretCreateOrUpdateParameters) (result keyvault.Secret, err error)
	Get(ctx context.Context, resourceGroupName string, vaultName string, secretName string) (result keyvault.Secret, err error)
	List(ctx context.Context, resourceGroupName string, vaultName string, top *int32) (result keyvault.SecretListResultPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, vaultName string, top *int32) (result keyvault.SecretListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, vaultName string, secretName string, parameters keyvault.SecretPatchParameters) (result keyvault.Secret, err error)
}

var _ SecretsClientAPI = (*keyvault.SecretsClient)(nil)
