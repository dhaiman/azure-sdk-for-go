package attestationapi

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
	"github.com/Azure/azure-sdk-for-go/services/attestation/mgmt/2018-09-01/attestation"
	"github.com/Azure/go-autorest/autorest"
)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result attestation.OperationList, err error)
}

var _ OperationsClientAPI = (*attestation.OperationsClient)(nil)

// ProvidersClientAPI contains the set of methods on the ProvidersClient type.
type ProvidersClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, providerName string, creationParams attestation.ServiceCreationParams) (result attestation.Provider, err error)
	Delete(ctx context.Context, resourceGroupName string, providerName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, providerName string) (result attestation.Provider, err error)
	GetDefaultByLocation(ctx context.Context, location string) (result attestation.Provider, err error)
	List(ctx context.Context) (result attestation.ProviderListResult, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result attestation.ProviderListResult, err error)
	ListDefault(ctx context.Context) (result attestation.ProviderListResult, err error)
	Update(ctx context.Context, resourceGroupName string, providerName string, updateParams attestation.ServicePatchParams) (result attestation.Provider, err error)
}

var _ ProvidersClientAPI = (*attestation.ProvidersClient)(nil)
