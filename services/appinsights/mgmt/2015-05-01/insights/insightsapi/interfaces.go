package insightsapi

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
	"github.com/Azure/azure-sdk-for-go/services/appinsights/mgmt/2015-05-01/insights"
	"github.com/Azure/go-autorest/autorest"
)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result insights.OperationListResultPage, err error)
	ListComplete(ctx context.Context) (result insights.OperationListResultIterator, err error)
}

var _ OperationsClientAPI = (*insights.OperationsClient)(nil)

// AnnotationsClientAPI contains the set of methods on the AnnotationsClient type.
type AnnotationsClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, resourceName string, annotationProperties insights.Annotation) (result insights.ListAnnotation, err error)
	Delete(ctx context.Context, resourceGroupName string, resourceName string, annotationID string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, resourceName string, annotationID string) (result insights.ListAnnotation, err error)
	List(ctx context.Context, resourceGroupName string, resourceName string, start string, end string) (result insights.AnnotationsListResult, err error)
}

var _ AnnotationsClientAPI = (*insights.AnnotationsClient)(nil)

// APIKeysClientAPI contains the set of methods on the APIKeysClient type.
type APIKeysClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, resourceName string, APIKeyProperties insights.APIKeyRequest) (result insights.ApplicationInsightsComponentAPIKey, err error)
	Delete(ctx context.Context, resourceGroupName string, resourceName string, keyID string) (result insights.ApplicationInsightsComponentAPIKey, err error)
	Get(ctx context.Context, resourceGroupName string, resourceName string, keyID string) (result insights.ApplicationInsightsComponentAPIKey, err error)
	List(ctx context.Context, resourceGroupName string, resourceName string) (result insights.ApplicationInsightsComponentAPIKeyListResult, err error)
}

var _ APIKeysClientAPI = (*insights.APIKeysClient)(nil)

// ExportConfigurationsClientAPI contains the set of methods on the ExportConfigurationsClient type.
type ExportConfigurationsClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, resourceName string, exportProperties insights.ApplicationInsightsComponentExportRequest) (result insights.ListApplicationInsightsComponentExportConfiguration, err error)
	Delete(ctx context.Context, resourceGroupName string, resourceName string, exportID string) (result insights.ApplicationInsightsComponentExportConfiguration, err error)
	Get(ctx context.Context, resourceGroupName string, resourceName string, exportID string) (result insights.ApplicationInsightsComponentExportConfiguration, err error)
	List(ctx context.Context, resourceGroupName string, resourceName string) (result insights.ListApplicationInsightsComponentExportConfiguration, err error)
	Update(ctx context.Context, resourceGroupName string, resourceName string, exportID string, exportProperties insights.ApplicationInsightsComponentExportRequest) (result insights.ApplicationInsightsComponentExportConfiguration, err error)
}

var _ ExportConfigurationsClientAPI = (*insights.ExportConfigurationsClient)(nil)

// ComponentCurrentBillingFeaturesClientAPI contains the set of methods on the ComponentCurrentBillingFeaturesClient type.
type ComponentCurrentBillingFeaturesClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, resourceName string) (result insights.ApplicationInsightsComponentBillingFeatures, err error)
	Update(ctx context.Context, resourceGroupName string, resourceName string, billingFeaturesProperties insights.ApplicationInsightsComponentBillingFeatures) (result insights.ApplicationInsightsComponentBillingFeatures, err error)
}

var _ ComponentCurrentBillingFeaturesClientAPI = (*insights.ComponentCurrentBillingFeaturesClient)(nil)

// ComponentQuotaStatusClientAPI contains the set of methods on the ComponentQuotaStatusClient type.
type ComponentQuotaStatusClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, resourceName string) (result insights.ApplicationInsightsComponentQuotaStatus, err error)
}

var _ ComponentQuotaStatusClientAPI = (*insights.ComponentQuotaStatusClient)(nil)

// ComponentFeatureCapabilitiesClientAPI contains the set of methods on the ComponentFeatureCapabilitiesClient type.
type ComponentFeatureCapabilitiesClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, resourceName string) (result insights.ApplicationInsightsComponentFeatureCapabilities, err error)
}

var _ ComponentFeatureCapabilitiesClientAPI = (*insights.ComponentFeatureCapabilitiesClient)(nil)

// ComponentAvailableFeaturesClientAPI contains the set of methods on the ComponentAvailableFeaturesClient type.
type ComponentAvailableFeaturesClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, resourceName string) (result insights.ApplicationInsightsComponentAvailableFeatures, err error)
}

var _ ComponentAvailableFeaturesClientAPI = (*insights.ComponentAvailableFeaturesClient)(nil)

// ProactiveDetectionConfigurationsClientAPI contains the set of methods on the ProactiveDetectionConfigurationsClient type.
type ProactiveDetectionConfigurationsClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, resourceName string, configurationID string) (result insights.ApplicationInsightsComponentProactiveDetectionConfiguration, err error)
	List(ctx context.Context, resourceGroupName string, resourceName string) (result insights.ListApplicationInsightsComponentProactiveDetectionConfiguration, err error)
	Update(ctx context.Context, resourceGroupName string, resourceName string, configurationID string, proactiveDetectionProperties insights.ApplicationInsightsComponentProactiveDetectionConfiguration) (result insights.ApplicationInsightsComponentProactiveDetectionConfiguration, err error)
}

var _ ProactiveDetectionConfigurationsClientAPI = (*insights.ProactiveDetectionConfigurationsClient)(nil)

// ComponentsClientAPI contains the set of methods on the ComponentsClient type.
type ComponentsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, insightProperties insights.ApplicationInsightsComponent) (result insights.ApplicationInsightsComponent, err error)
	Delete(ctx context.Context, resourceGroupName string, resourceName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, resourceName string) (result insights.ApplicationInsightsComponent, err error)
	GetPurgeStatus(ctx context.Context, resourceGroupName string, resourceName string, purgeID string) (result insights.ComponentPurgeStatusResponse, err error)
	List(ctx context.Context) (result insights.ApplicationInsightsComponentListResultPage, err error)
	ListComplete(ctx context.Context) (result insights.ApplicationInsightsComponentListResultIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result insights.ApplicationInsightsComponentListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result insights.ApplicationInsightsComponentListResultIterator, err error)
	Purge(ctx context.Context, resourceGroupName string, resourceName string, body insights.ComponentPurgeBody) (result insights.ComponentPurgeResponse, err error)
	UpdateTags(ctx context.Context, resourceGroupName string, resourceName string, componentTags insights.TagsResource) (result insights.ApplicationInsightsComponent, err error)
}

var _ ComponentsClientAPI = (*insights.ComponentsClient)(nil)

// WorkItemConfigurationsClientAPI contains the set of methods on the WorkItemConfigurationsClient type.
type WorkItemConfigurationsClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, resourceName string, workItemConfigurationProperties insights.WorkItemCreateConfiguration) (result insights.WorkItemConfiguration, err error)
	Delete(ctx context.Context, resourceGroupName string, resourceName string, workItemConfigID string) (result autorest.Response, err error)
	GetDefault(ctx context.Context, resourceGroupName string, resourceName string) (result insights.WorkItemConfiguration, err error)
	GetItem(ctx context.Context, resourceGroupName string, resourceName string, workItemConfigID string) (result insights.WorkItemConfiguration, err error)
	List(ctx context.Context, resourceGroupName string, resourceName string) (result insights.WorkItemConfigurationsListResult, err error)
	UpdateItem(ctx context.Context, resourceGroupName string, resourceName string, workItemConfigID string, workItemConfigurationProperties insights.WorkItemCreateConfiguration) (result insights.WorkItemConfiguration, err error)
}

var _ WorkItemConfigurationsClientAPI = (*insights.WorkItemConfigurationsClient)(nil)

// FavoritesClientAPI contains the set of methods on the FavoritesClient type.
type FavoritesClientAPI interface {
	Add(ctx context.Context, resourceGroupName string, resourceName string, favoriteID string, favoriteProperties insights.ApplicationInsightsComponentFavorite) (result insights.ApplicationInsightsComponentFavorite, err error)
	Delete(ctx context.Context, resourceGroupName string, resourceName string, favoriteID string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, resourceName string, favoriteID string) (result insights.ApplicationInsightsComponentFavorite, err error)
	List(ctx context.Context, resourceGroupName string, resourceName string, favoriteType insights.FavoriteType, sourceType insights.FavoriteSourceType, canFetchContent *bool, tags []string) (result insights.ListApplicationInsightsComponentFavorite, err error)
	Update(ctx context.Context, resourceGroupName string, resourceName string, favoriteID string, favoriteProperties insights.ApplicationInsightsComponentFavorite) (result insights.ApplicationInsightsComponentFavorite, err error)
}

var _ FavoritesClientAPI = (*insights.FavoritesClient)(nil)

// WebTestLocationsClientAPI contains the set of methods on the WebTestLocationsClient type.
type WebTestLocationsClientAPI interface {
	List(ctx context.Context, resourceGroupName string, resourceName string) (result insights.ApplicationInsightsWebTestLocationsListResult, err error)
}

var _ WebTestLocationsClientAPI = (*insights.WebTestLocationsClient)(nil)

// WebTestsClientAPI contains the set of methods on the WebTestsClient type.
type WebTestsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, webTestName string, webTestDefinition insights.WebTest) (result insights.WebTest, err error)
	Delete(ctx context.Context, resourceGroupName string, webTestName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, webTestName string) (result insights.WebTest, err error)
	List(ctx context.Context) (result insights.WebTestListResultPage, err error)
	ListComplete(ctx context.Context) (result insights.WebTestListResultIterator, err error)
	ListByComponent(ctx context.Context, componentName string, resourceGroupName string) (result insights.WebTestListResultPage, err error)
	ListByComponentComplete(ctx context.Context, componentName string, resourceGroupName string) (result insights.WebTestListResultIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result insights.WebTestListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result insights.WebTestListResultIterator, err error)
	UpdateTags(ctx context.Context, resourceGroupName string, webTestName string, webTestTags insights.TagsResource) (result insights.WebTest, err error)
}

var _ WebTestsClientAPI = (*insights.WebTestsClient)(nil)

// AnalyticsItemsClientAPI contains the set of methods on the AnalyticsItemsClient type.
type AnalyticsItemsClientAPI interface {
	Delete(ctx context.Context, resourceGroupName string, resourceName string, scopePath insights.ItemScopePath, ID string, name string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, resourceName string, scopePath insights.ItemScopePath, ID string, name string) (result insights.ApplicationInsightsComponentAnalyticsItem, err error)
	List(ctx context.Context, resourceGroupName string, resourceName string, scopePath insights.ItemScopePath, scope insights.ItemScope, typeParameter insights.ItemTypeParameter, includeContent *bool) (result insights.ListApplicationInsightsComponentAnalyticsItem, err error)
	Put(ctx context.Context, resourceGroupName string, resourceName string, scopePath insights.ItemScopePath, itemProperties insights.ApplicationInsightsComponentAnalyticsItem, overrideItem *bool) (result insights.ApplicationInsightsComponentAnalyticsItem, err error)
}

var _ AnalyticsItemsClientAPI = (*insights.AnalyticsItemsClient)(nil)

// WorkbooksClientAPI contains the set of methods on the WorkbooksClient type.
type WorkbooksClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, workbookProperties insights.Workbook) (result insights.Workbook, err error)
	Delete(ctx context.Context, resourceGroupName string, resourceName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, resourceName string) (result insights.Workbook, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string, category insights.CategoryType, tags []string, canFetchContent *bool) (result insights.WorkbooksListResult, err error)
	Update(ctx context.Context, resourceGroupName string, resourceName string, workbookProperties insights.Workbook) (result insights.Workbook, err error)
}

var _ WorkbooksClientAPI = (*insights.WorkbooksClient)(nil)

// MyWorkbooksClientAPI contains the set of methods on the MyWorkbooksClient type.
type MyWorkbooksClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, workbookProperties insights.MyWorkbook) (result insights.MyWorkbook, err error)
	Delete(ctx context.Context, resourceGroupName string, resourceName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, resourceName string) (result insights.MyWorkbook, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string, category insights.CategoryType, tags []string, canFetchContent *bool) (result insights.MyWorkbooksListResult, err error)
	ListBySubscription(ctx context.Context, category insights.CategoryType, tags []string, canFetchContent *bool) (result insights.MyWorkbooksListResult, err error)
	Update(ctx context.Context, resourceGroupName string, resourceName string, workbookProperties insights.MyWorkbook) (result insights.MyWorkbook, err error)
}

var _ MyWorkbooksClientAPI = (*insights.MyWorkbooksClient)(nil)
