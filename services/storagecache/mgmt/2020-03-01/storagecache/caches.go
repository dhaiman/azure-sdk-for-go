package storagecache

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
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// CachesClient is the a Storage Cache provides scalable caching service for NAS clients, serving data from either
// NFSv3 or Blob at-rest storage (referred to as "Storage Targets"). These operations allow you to manage Caches.
type CachesClient struct {
	BaseClient
}

// NewCachesClient creates an instance of the CachesClient client.
func NewCachesClient(subscriptionID string) CachesClient {
	return NewCachesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewCachesClientWithBaseURI creates an instance of the CachesClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewCachesClientWithBaseURI(baseURI string, subscriptionID string) CachesClient {
	return CachesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate create or update a Cache.
// Parameters:
// resourceGroupName - target resource group.
// cacheName - name of Cache. Length of name must be not greater than 80 and chars must be in list of
// [-0-9a-zA-Z_] char class.
// cache - object containing the user-selectable properties of the new Cache. If read-only properties are
// included, they must match the existing values of those properties.
func (client CachesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, cacheName string, cache *Cache) (result CachesCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CachesClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: cacheName,
			Constraints: []validation.Constraint{{Target: "cacheName", Name: validation.Pattern, Rule: `^[-0-9a-zA-Z_]{1,80}$`, Chain: nil}}},
		{TargetValue: cache,
			Constraints: []validation.Constraint{{Target: "cache", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "cache.CacheProperties", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "cache.CacheProperties.NetworkSettings", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "cache.CacheProperties.NetworkSettings.Mtu", Name: validation.Null, Rule: false,
							Chain: []validation.Constraint{{Target: "cache.CacheProperties.NetworkSettings.Mtu", Name: validation.InclusiveMaximum, Rule: int64(1500), Chain: nil},
								{Target: "cache.CacheProperties.NetworkSettings.Mtu", Name: validation.InclusiveMinimum, Rule: int64(576), Chain: nil},
							}},
						}},
						{Target: "cache.CacheProperties.EncryptionSettings", Name: validation.Null, Rule: false,
							Chain: []validation.Constraint{{Target: "cache.CacheProperties.EncryptionSettings.KeyEncryptionKey", Name: validation.Null, Rule: false,
								Chain: []validation.Constraint{{Target: "cache.CacheProperties.EncryptionSettings.KeyEncryptionKey.KeyURL", Name: validation.Null, Rule: true, Chain: nil},
									{Target: "cache.CacheProperties.EncryptionSettings.KeyEncryptionKey.SourceVault", Name: validation.Null, Rule: true, Chain: nil},
								}},
							}},
					}},
				}}}}}); err != nil {
		return result, validation.NewError("storagecache.CachesClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, cacheName, cache)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagecache.CachesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagecache.CachesClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client CachesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, cacheName string, cache *Cache) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"cacheName":         autorest.Encode("path", cacheName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	cache.ID = nil
	cache.Name = nil
	cache.Type = nil
	cache.SystemData = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StorageCache/caches/{cacheName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if cache != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(cache))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client CachesClient) CreateOrUpdateSender(req *http.Request) (future CachesCreateOrUpdateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client CachesClient) CreateOrUpdateResponder(resp *http.Response) (result Cache, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete schedules a Cache for deletion.
// Parameters:
// resourceGroupName - target resource group.
// cacheName - name of Cache. Length of name must be not greater than 80 and chars must be in list of
// [-0-9a-zA-Z_] char class.
func (client CachesClient) Delete(ctx context.Context, resourceGroupName string, cacheName string) (result CachesDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CachesClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: cacheName,
			Constraints: []validation.Constraint{{Target: "cacheName", Name: validation.Pattern, Rule: `^[-0-9a-zA-Z_]{1,80}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storagecache.CachesClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, cacheName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagecache.CachesClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagecache.CachesClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client CachesClient) DeletePreparer(ctx context.Context, resourceGroupName string, cacheName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"cacheName":         autorest.Encode("path", cacheName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StorageCache/caches/{cacheName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client CachesClient) DeleteSender(req *http.Request) (future CachesDeleteFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client CachesClient) DeleteResponder(resp *http.Response) (result SetObject, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Flush tells a Cache to write all dirty data to the Storage Target(s). During the flush, clients will see errors
// returned until the flush is complete.
// Parameters:
// resourceGroupName - target resource group.
// cacheName - name of Cache. Length of name must be not greater than 80 and chars must be in list of
// [-0-9a-zA-Z_] char class.
func (client CachesClient) Flush(ctx context.Context, resourceGroupName string, cacheName string) (result CachesFlushFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CachesClient.Flush")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: cacheName,
			Constraints: []validation.Constraint{{Target: "cacheName", Name: validation.Pattern, Rule: `^[-0-9a-zA-Z_]{1,80}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storagecache.CachesClient", "Flush", err.Error())
	}

	req, err := client.FlushPreparer(ctx, resourceGroupName, cacheName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagecache.CachesClient", "Flush", nil, "Failure preparing request")
		return
	}

	result, err = client.FlushSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagecache.CachesClient", "Flush", result.Response(), "Failure sending request")
		return
	}

	return
}

// FlushPreparer prepares the Flush request.
func (client CachesClient) FlushPreparer(ctx context.Context, resourceGroupName string, cacheName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"cacheName":         autorest.Encode("path", cacheName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StorageCache/caches/{cacheName}/flush", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// FlushSender sends the Flush request. The method will close the
// http.Response Body if it receives an error.
func (client CachesClient) FlushSender(req *http.Request) (future CachesFlushFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// FlushResponder handles the response to the Flush request. The method always
// closes the http.Response Body.
func (client CachesClient) FlushResponder(resp *http.Response) (result SetObject, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get returns a Cache.
// Parameters:
// resourceGroupName - target resource group.
// cacheName - name of Cache. Length of name must be not greater than 80 and chars must be in list of
// [-0-9a-zA-Z_] char class.
func (client CachesClient) Get(ctx context.Context, resourceGroupName string, cacheName string) (result Cache, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CachesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: cacheName,
			Constraints: []validation.Constraint{{Target: "cacheName", Name: validation.Pattern, Rule: `^[-0-9a-zA-Z_]{1,80}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storagecache.CachesClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, cacheName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagecache.CachesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storagecache.CachesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagecache.CachesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client CachesClient) GetPreparer(ctx context.Context, resourceGroupName string, cacheName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"cacheName":         autorest.Encode("path", cacheName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StorageCache/caches/{cacheName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client CachesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client CachesClient) GetResponder(resp *http.Response) (result Cache, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List returns all Caches the user has access to under a subscription.
func (client CachesClient) List(ctx context.Context) (result CachesListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CachesClient.List")
		defer func() {
			sc := -1
			if result.clr.Response.Response != nil {
				sc = result.clr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagecache.CachesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.clr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storagecache.CachesClient", "List", resp, "Failure sending request")
		return
	}

	result.clr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagecache.CachesClient", "List", resp, "Failure responding to request")
	}
	if result.clr.hasNextLink() && result.clr.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListPreparer prepares the List request.
func (client CachesClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.StorageCache/caches", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client CachesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client CachesClient) ListResponder(resp *http.Response) (result CachesListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client CachesClient) listNextResults(ctx context.Context, lastResults CachesListResult) (result CachesListResult, err error) {
	req, err := lastResults.cachesListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "storagecache.CachesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "storagecache.CachesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagecache.CachesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client CachesClient) ListComplete(ctx context.Context) (result CachesListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CachesClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx)
	return
}

// ListByResourceGroup returns all Caches the user has access to under a resource group.
// Parameters:
// resourceGroupName - target resource group.
func (client CachesClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (result CachesListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CachesClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.clr.Response.Response != nil {
				sc = result.clr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagecache.CachesClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.clr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storagecache.CachesClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.clr, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagecache.CachesClient", "ListByResourceGroup", resp, "Failure responding to request")
	}
	if result.clr.hasNextLink() && result.clr.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client CachesClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StorageCache/caches", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client CachesClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client CachesClient) ListByResourceGroupResponder(resp *http.Response) (result CachesListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client CachesClient) listByResourceGroupNextResults(ctx context.Context, lastResults CachesListResult) (result CachesListResult, err error) {
	req, err := lastResults.cachesListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "storagecache.CachesClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "storagecache.CachesClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagecache.CachesClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client CachesClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result CachesListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CachesClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName)
	return
}

// Start tells a Stopped state Cache to transition to Active state.
// Parameters:
// resourceGroupName - target resource group.
// cacheName - name of Cache. Length of name must be not greater than 80 and chars must be in list of
// [-0-9a-zA-Z_] char class.
func (client CachesClient) Start(ctx context.Context, resourceGroupName string, cacheName string) (result CachesStartFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CachesClient.Start")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: cacheName,
			Constraints: []validation.Constraint{{Target: "cacheName", Name: validation.Pattern, Rule: `^[-0-9a-zA-Z_]{1,80}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storagecache.CachesClient", "Start", err.Error())
	}

	req, err := client.StartPreparer(ctx, resourceGroupName, cacheName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagecache.CachesClient", "Start", nil, "Failure preparing request")
		return
	}

	result, err = client.StartSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagecache.CachesClient", "Start", result.Response(), "Failure sending request")
		return
	}

	return
}

// StartPreparer prepares the Start request.
func (client CachesClient) StartPreparer(ctx context.Context, resourceGroupName string, cacheName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"cacheName":         autorest.Encode("path", cacheName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StorageCache/caches/{cacheName}/start", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// StartSender sends the Start request. The method will close the
// http.Response Body if it receives an error.
func (client CachesClient) StartSender(req *http.Request) (future CachesStartFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// StartResponder handles the response to the Start request. The method always
// closes the http.Response Body.
func (client CachesClient) StartResponder(resp *http.Response) (result SetObject, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Stop tells an Active Cache to transition to Stopped state.
// Parameters:
// resourceGroupName - target resource group.
// cacheName - name of Cache. Length of name must be not greater than 80 and chars must be in list of
// [-0-9a-zA-Z_] char class.
func (client CachesClient) Stop(ctx context.Context, resourceGroupName string, cacheName string) (result CachesStopFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CachesClient.Stop")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: cacheName,
			Constraints: []validation.Constraint{{Target: "cacheName", Name: validation.Pattern, Rule: `^[-0-9a-zA-Z_]{1,80}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storagecache.CachesClient", "Stop", err.Error())
	}

	req, err := client.StopPreparer(ctx, resourceGroupName, cacheName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagecache.CachesClient", "Stop", nil, "Failure preparing request")
		return
	}

	result, err = client.StopSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagecache.CachesClient", "Stop", result.Response(), "Failure sending request")
		return
	}

	return
}

// StopPreparer prepares the Stop request.
func (client CachesClient) StopPreparer(ctx context.Context, resourceGroupName string, cacheName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"cacheName":         autorest.Encode("path", cacheName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StorageCache/caches/{cacheName}/stop", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// StopSender sends the Stop request. The method will close the
// http.Response Body if it receives an error.
func (client CachesClient) StopSender(req *http.Request) (future CachesStopFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// StopResponder handles the response to the Stop request. The method always
// closes the http.Response Body.
func (client CachesClient) StopResponder(resp *http.Response) (result SetObject, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update update a Cache instance.
// Parameters:
// resourceGroupName - target resource group.
// cacheName - name of Cache. Length of name must be not greater than 80 and chars must be in list of
// [-0-9a-zA-Z_] char class.
// cache - object containing the user-selectable properties of the Cache. If read-only properties are included,
// they must match the existing values of those properties.
func (client CachesClient) Update(ctx context.Context, resourceGroupName string, cacheName string, cache *Cache) (result Cache, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CachesClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: cacheName,
			Constraints: []validation.Constraint{{Target: "cacheName", Name: validation.Pattern, Rule: `^[-0-9a-zA-Z_]{1,80}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storagecache.CachesClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, resourceGroupName, cacheName, cache)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagecache.CachesClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storagecache.CachesClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagecache.CachesClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client CachesClient) UpdatePreparer(ctx context.Context, resourceGroupName string, cacheName string, cache *Cache) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"cacheName":         autorest.Encode("path", cacheName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	cache.ID = nil
	cache.Name = nil
	cache.Type = nil
	cache.SystemData = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StorageCache/caches/{cacheName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if cache != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(cache))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client CachesClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client CachesClient) UpdateResponder(resp *http.Response) (result Cache, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// UpgradeFirmware upgrade a Cache's firmware if a new version is available. Otherwise, this operation has no effect.
// Parameters:
// resourceGroupName - target resource group.
// cacheName - name of Cache. Length of name must be not greater than 80 and chars must be in list of
// [-0-9a-zA-Z_] char class.
func (client CachesClient) UpgradeFirmware(ctx context.Context, resourceGroupName string, cacheName string) (result CachesUpgradeFirmwareFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CachesClient.UpgradeFirmware")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: cacheName,
			Constraints: []validation.Constraint{{Target: "cacheName", Name: validation.Pattern, Rule: `^[-0-9a-zA-Z_]{1,80}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storagecache.CachesClient", "UpgradeFirmware", err.Error())
	}

	req, err := client.UpgradeFirmwarePreparer(ctx, resourceGroupName, cacheName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagecache.CachesClient", "UpgradeFirmware", nil, "Failure preparing request")
		return
	}

	result, err = client.UpgradeFirmwareSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagecache.CachesClient", "UpgradeFirmware", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpgradeFirmwarePreparer prepares the UpgradeFirmware request.
func (client CachesClient) UpgradeFirmwarePreparer(ctx context.Context, resourceGroupName string, cacheName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"cacheName":         autorest.Encode("path", cacheName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StorageCache/caches/{cacheName}/upgrade", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpgradeFirmwareSender sends the UpgradeFirmware request. The method will close the
// http.Response Body if it receives an error.
func (client CachesClient) UpgradeFirmwareSender(req *http.Request) (future CachesUpgradeFirmwareFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// UpgradeFirmwareResponder handles the response to the UpgradeFirmware request. The method always
// closes the http.Response Body.
func (client CachesClient) UpgradeFirmwareResponder(resp *http.Response) (result SetObject, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
