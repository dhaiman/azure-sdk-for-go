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

package advisor

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/advisor/mgmt/2017-04-19/advisor"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type Category = original.Category

const (
	Cost                  Category = original.Cost
	HighAvailability      Category = original.HighAvailability
	OperationalExcellence Category = original.OperationalExcellence
	Performance           Category = original.Performance
	Security              Category = original.Security
)

type Impact = original.Impact

const (
	High   Impact = original.High
	Low    Impact = original.Low
	Medium Impact = original.Medium
)

type Risk = original.Risk

const (
	Error   Risk = original.Error
	None    Risk = original.None
	Warning Risk = original.Warning
)

type Scenario = original.Scenario

const (
	Alerts Scenario = original.Alerts
)

type ARMErrorResponseBody = original.ARMErrorResponseBody
type BaseClient = original.BaseClient
type ConfigData = original.ConfigData
type ConfigDataProperties = original.ConfigDataProperties
type ConfigurationListResult = original.ConfigurationListResult
type ConfigurationListResultIterator = original.ConfigurationListResultIterator
type ConfigurationListResultPage = original.ConfigurationListResultPage
type ConfigurationsClient = original.ConfigurationsClient
type MetadataEntity = original.MetadataEntity
type MetadataEntityListResult = original.MetadataEntityListResult
type MetadataEntityListResultIterator = original.MetadataEntityListResultIterator
type MetadataEntityListResultPage = original.MetadataEntityListResultPage
type MetadataEntityProperties = original.MetadataEntityProperties
type MetadataSupportedValueDetail = original.MetadataSupportedValueDetail
type OperationDisplayInfo = original.OperationDisplayInfo
type OperationEntity = original.OperationEntity
type OperationEntityListResult = original.OperationEntityListResult
type OperationEntityListResultIterator = original.OperationEntityListResultIterator
type OperationEntityListResultPage = original.OperationEntityListResultPage
type OperationsClient = original.OperationsClient
type RecommendationMetadataClient = original.RecommendationMetadataClient
type RecommendationProperties = original.RecommendationProperties
type RecommendationsClient = original.RecommendationsClient
type Resource = original.Resource
type ResourceRecommendationBase = original.ResourceRecommendationBase
type ResourceRecommendationBaseListResult = original.ResourceRecommendationBaseListResult
type ResourceRecommendationBaseListResultIterator = original.ResourceRecommendationBaseListResultIterator
type ResourceRecommendationBaseListResultPage = original.ResourceRecommendationBaseListResultPage
type SetObject = original.SetObject
type ShortDescription = original.ShortDescription
type SuppressionContract = original.SuppressionContract
type SuppressionContractListResult = original.SuppressionContractListResult
type SuppressionContractListResultIterator = original.SuppressionContractListResultIterator
type SuppressionContractListResultPage = original.SuppressionContractListResultPage
type SuppressionProperties = original.SuppressionProperties
type SuppressionsClient = original.SuppressionsClient

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewConfigurationListResultIterator(page ConfigurationListResultPage) ConfigurationListResultIterator {
	return original.NewConfigurationListResultIterator(page)
}
func NewConfigurationListResultPage(getNextPage func(context.Context, ConfigurationListResult) (ConfigurationListResult, error)) ConfigurationListResultPage {
	return original.NewConfigurationListResultPage(getNextPage)
}
func NewConfigurationsClient(subscriptionID string) ConfigurationsClient {
	return original.NewConfigurationsClient(subscriptionID)
}
func NewConfigurationsClientWithBaseURI(baseURI string, subscriptionID string) ConfigurationsClient {
	return original.NewConfigurationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewMetadataEntityListResultIterator(page MetadataEntityListResultPage) MetadataEntityListResultIterator {
	return original.NewMetadataEntityListResultIterator(page)
}
func NewMetadataEntityListResultPage(getNextPage func(context.Context, MetadataEntityListResult) (MetadataEntityListResult, error)) MetadataEntityListResultPage {
	return original.NewMetadataEntityListResultPage(getNextPage)
}
func NewOperationEntityListResultIterator(page OperationEntityListResultPage) OperationEntityListResultIterator {
	return original.NewOperationEntityListResultIterator(page)
}
func NewOperationEntityListResultPage(getNextPage func(context.Context, OperationEntityListResult) (OperationEntityListResult, error)) OperationEntityListResultPage {
	return original.NewOperationEntityListResultPage(getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewRecommendationMetadataClient(subscriptionID string) RecommendationMetadataClient {
	return original.NewRecommendationMetadataClient(subscriptionID)
}
func NewRecommendationMetadataClientWithBaseURI(baseURI string, subscriptionID string) RecommendationMetadataClient {
	return original.NewRecommendationMetadataClientWithBaseURI(baseURI, subscriptionID)
}
func NewRecommendationsClient(subscriptionID string) RecommendationsClient {
	return original.NewRecommendationsClient(subscriptionID)
}
func NewRecommendationsClientWithBaseURI(baseURI string, subscriptionID string) RecommendationsClient {
	return original.NewRecommendationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewResourceRecommendationBaseListResultIterator(page ResourceRecommendationBaseListResultPage) ResourceRecommendationBaseListResultIterator {
	return original.NewResourceRecommendationBaseListResultIterator(page)
}
func NewResourceRecommendationBaseListResultPage(getNextPage func(context.Context, ResourceRecommendationBaseListResult) (ResourceRecommendationBaseListResult, error)) ResourceRecommendationBaseListResultPage {
	return original.NewResourceRecommendationBaseListResultPage(getNextPage)
}
func NewSuppressionContractListResultIterator(page SuppressionContractListResultPage) SuppressionContractListResultIterator {
	return original.NewSuppressionContractListResultIterator(page)
}
func NewSuppressionContractListResultPage(getNextPage func(context.Context, SuppressionContractListResult) (SuppressionContractListResult, error)) SuppressionContractListResultPage {
	return original.NewSuppressionContractListResultPage(getNextPage)
}
func NewSuppressionsClient(subscriptionID string) SuppressionsClient {
	return original.NewSuppressionsClient(subscriptionID)
}
func NewSuppressionsClientWithBaseURI(baseURI string, subscriptionID string) SuppressionsClient {
	return original.NewSuppressionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleCategoryValues() []Category {
	return original.PossibleCategoryValues()
}
func PossibleImpactValues() []Impact {
	return original.PossibleImpactValues()
}
func PossibleRiskValues() []Risk {
	return original.PossibleRiskValues()
}
func PossibleScenarioValues() []Scenario {
	return original.PossibleScenarioValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
