// +build go1.9

// Copyright 2019 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package signalr

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/signalr/mgmt/2018-10-01/signalr"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type FeatureFlags = original.FeatureFlags

const (
	EnableConnectivityLogs FeatureFlags = original.EnableConnectivityLogs
	ServiceMode            FeatureFlags = original.ServiceMode
)

type KeyType = original.KeyType

const (
	Primary   KeyType = original.Primary
	Secondary KeyType = original.Secondary
)

type ProvisioningState = original.ProvisioningState

const (
	Canceled  ProvisioningState = original.Canceled
	Creating  ProvisioningState = original.Creating
	Deleting  ProvisioningState = original.Deleting
	Failed    ProvisioningState = original.Failed
	Moving    ProvisioningState = original.Moving
	Running   ProvisioningState = original.Running
	Succeeded ProvisioningState = original.Succeeded
	Unknown   ProvisioningState = original.Unknown
	Updating  ProvisioningState = original.Updating
)

type SkuTier = original.SkuTier

const (
	Basic    SkuTier = original.Basic
	Free     SkuTier = original.Free
	Premium  SkuTier = original.Premium
	Standard SkuTier = original.Standard
)

type BaseClient = original.BaseClient
type Client = original.Client
type CorsSettings = original.CorsSettings
type CreateOrUpdateFuture = original.CreateOrUpdateFuture
type CreateOrUpdateProperties = original.CreateOrUpdateProperties
type CreateParameters = original.CreateParameters
type DeleteFuture = original.DeleteFuture
type Dimension = original.Dimension
type ErrorResponse = original.ErrorResponse
type ErrorResponseBody = original.ErrorResponseBody
type Feature = original.Feature
type Keys = original.Keys
type LogSpecification = original.LogSpecification
type MetricSpecification = original.MetricSpecification
type NameAvailability = original.NameAvailability
type NameAvailabilityParameters = original.NameAvailabilityParameters
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationList = original.OperationList
type OperationListIterator = original.OperationListIterator
type OperationListPage = original.OperationListPage
type OperationProperties = original.OperationProperties
type OperationsClient = original.OperationsClient
type Properties = original.Properties
type RegenerateKeyFuture = original.RegenerateKeyFuture
type RegenerateKeyParameters = original.RegenerateKeyParameters
type Resource = original.Resource
type ResourceList = original.ResourceList
type ResourceListIterator = original.ResourceListIterator
type ResourceListPage = original.ResourceListPage
type ResourceSku = original.ResourceSku
type ResourceType = original.ResourceType
type RestartFuture = original.RestartFuture
type ServiceSpecification = original.ServiceSpecification
type TrackedResource = original.TrackedResource
type UpdateFuture = original.UpdateFuture
type UpdateParameters = original.UpdateParameters
type Usage = original.Usage
type UsageList = original.UsageList
type UsageListIterator = original.UsageListIterator
type UsageListPage = original.UsageListPage
type UsageName = original.UsageName
type UsagesClient = original.UsagesClient

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewClient(subscriptionID string) Client {
	return original.NewClient(subscriptionID)
}
func NewClientWithBaseURI(baseURI string, subscriptionID string) Client {
	return original.NewClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationListIterator(page OperationListPage) OperationListIterator {
	return original.NewOperationListIterator(page)
}
func NewOperationListPage(getNextPage func(context.Context, OperationList) (OperationList, error)) OperationListPage {
	return original.NewOperationListPage(getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewResourceListIterator(page ResourceListPage) ResourceListIterator {
	return original.NewResourceListIterator(page)
}
func NewResourceListPage(getNextPage func(context.Context, ResourceList) (ResourceList, error)) ResourceListPage {
	return original.NewResourceListPage(getNextPage)
}
func NewUsageListIterator(page UsageListPage) UsageListIterator {
	return original.NewUsageListIterator(page)
}
func NewUsageListPage(getNextPage func(context.Context, UsageList) (UsageList, error)) UsageListPage {
	return original.NewUsageListPage(getNextPage)
}
func NewUsagesClient(subscriptionID string) UsagesClient {
	return original.NewUsagesClient(subscriptionID)
}
func NewUsagesClientWithBaseURI(baseURI string, subscriptionID string) UsagesClient {
	return original.NewUsagesClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleFeatureFlagsValues() []FeatureFlags {
	return original.PossibleFeatureFlagsValues()
}
func PossibleKeyTypeValues() []KeyType {
	return original.PossibleKeyTypeValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossibleSkuTierValues() []SkuTier {
	return original.PossibleSkuTierValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
