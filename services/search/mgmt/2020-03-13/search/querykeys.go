package search

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
	"github.com/Azure/go-autorest/tracing"
	"github.com/satori/go.uuid"
	"net/http"
)

// QueryKeysClient is the client that can be used to manage Azure Cognitive Search services and API keys.
type QueryKeysClient struct {
	BaseClient
}

// NewQueryKeysClient creates an instance of the QueryKeysClient client.
func NewQueryKeysClient(subscriptionID string) QueryKeysClient {
	return NewQueryKeysClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewQueryKeysClientWithBaseURI creates an instance of the QueryKeysClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewQueryKeysClientWithBaseURI(baseURI string, subscriptionID string) QueryKeysClient {
	return QueryKeysClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create generates a new query key for the specified Search service. You can create up to 50 query keys per service.
// Parameters:
// resourceGroupName - the name of the resource group within the current subscription. You can obtain this
// value from the Azure Resource Manager API or the portal.
// searchServiceName - the name of the Azure Cognitive Search service associated with the specified resource
// group.
// name - the name of the new query API key.
// clientRequestID - a client-generated GUID value that identifies this request. If specified, this will be
// included in response information as a way to track the request.
func (client QueryKeysClient) Create(ctx context.Context, resourceGroupName string, searchServiceName string, name string, clientRequestID *uuid.UUID) (result QueryKey, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/QueryKeysClient.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreatePreparer(ctx, resourceGroupName, searchServiceName, name, clientRequestID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "search.QueryKeysClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "search.QueryKeysClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "search.QueryKeysClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client QueryKeysClient) CreatePreparer(ctx context.Context, resourceGroupName string, searchServiceName string, name string, clientRequestID *uuid.UUID) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"searchServiceName": autorest.Encode("path", searchServiceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-03-13"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices/{searchServiceName}/createQueryKey/{name}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if clientRequestID != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("x-ms-client-request-id", autorest.String(clientRequestID)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client QueryKeysClient) CreateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client QueryKeysClient) CreateResponder(resp *http.Response) (result QueryKey, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the specified query key. Unlike admin keys, query keys are not regenerated. The process for
// regenerating a query key is to delete and then recreate it.
// Parameters:
// resourceGroupName - the name of the resource group within the current subscription. You can obtain this
// value from the Azure Resource Manager API or the portal.
// searchServiceName - the name of the Azure Cognitive Search service associated with the specified resource
// group.
// key - the query key to be deleted. Query keys are identified by value, not by name.
// clientRequestID - a client-generated GUID value that identifies this request. If specified, this will be
// included in response information as a way to track the request.
func (client QueryKeysClient) Delete(ctx context.Context, resourceGroupName string, searchServiceName string, key string, clientRequestID *uuid.UUID) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/QueryKeysClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, searchServiceName, key, clientRequestID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "search.QueryKeysClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "search.QueryKeysClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "search.QueryKeysClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client QueryKeysClient) DeletePreparer(ctx context.Context, resourceGroupName string, searchServiceName string, key string, clientRequestID *uuid.UUID) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"key":               autorest.Encode("path", key),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"searchServiceName": autorest.Encode("path", searchServiceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-03-13"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices/{searchServiceName}/deleteQueryKey/{key}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if clientRequestID != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("x-ms-client-request-id", autorest.String(clientRequestID)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client QueryKeysClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client QueryKeysClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent, http.StatusNotFound),
		autorest.ByClosing())
	result.Response = resp
	return
}

// ListBySearchService returns the list of query API keys for the given Azure Cognitive Search service.
// Parameters:
// resourceGroupName - the name of the resource group within the current subscription. You can obtain this
// value from the Azure Resource Manager API or the portal.
// searchServiceName - the name of the Azure Cognitive Search service associated with the specified resource
// group.
// clientRequestID - a client-generated GUID value that identifies this request. If specified, this will be
// included in response information as a way to track the request.
func (client QueryKeysClient) ListBySearchService(ctx context.Context, resourceGroupName string, searchServiceName string, clientRequestID *uuid.UUID) (result ListQueryKeysResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/QueryKeysClient.ListBySearchService")
		defer func() {
			sc := -1
			if result.lqkr.Response.Response != nil {
				sc = result.lqkr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listBySearchServiceNextResults
	req, err := client.ListBySearchServicePreparer(ctx, resourceGroupName, searchServiceName, clientRequestID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "search.QueryKeysClient", "ListBySearchService", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBySearchServiceSender(req)
	if err != nil {
		result.lqkr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "search.QueryKeysClient", "ListBySearchService", resp, "Failure sending request")
		return
	}

	result.lqkr, err = client.ListBySearchServiceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "search.QueryKeysClient", "ListBySearchService", resp, "Failure responding to request")
	}
	if result.lqkr.hasNextLink() && result.lqkr.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListBySearchServicePreparer prepares the ListBySearchService request.
func (client QueryKeysClient) ListBySearchServicePreparer(ctx context.Context, resourceGroupName string, searchServiceName string, clientRequestID *uuid.UUID) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"searchServiceName": autorest.Encode("path", searchServiceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-03-13"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices/{searchServiceName}/listQueryKeys", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if clientRequestID != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("x-ms-client-request-id", autorest.String(clientRequestID)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListBySearchServiceSender sends the ListBySearchService request. The method will close the
// http.Response Body if it receives an error.
func (client QueryKeysClient) ListBySearchServiceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListBySearchServiceResponder handles the response to the ListBySearchService request. The method always
// closes the http.Response Body.
func (client QueryKeysClient) ListBySearchServiceResponder(resp *http.Response) (result ListQueryKeysResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listBySearchServiceNextResults retrieves the next set of results, if any.
func (client QueryKeysClient) listBySearchServiceNextResults(ctx context.Context, lastResults ListQueryKeysResult) (result ListQueryKeysResult, err error) {
	req, err := lastResults.listQueryKeysResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "search.QueryKeysClient", "listBySearchServiceNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListBySearchServiceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "search.QueryKeysClient", "listBySearchServiceNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListBySearchServiceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "search.QueryKeysClient", "listBySearchServiceNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListBySearchServiceComplete enumerates all values, automatically crossing page boundaries as required.
func (client QueryKeysClient) ListBySearchServiceComplete(ctx context.Context, resourceGroupName string, searchServiceName string, clientRequestID *uuid.UUID) (result ListQueryKeysResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/QueryKeysClient.ListBySearchService")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListBySearchService(ctx, resourceGroupName, searchServiceName, clientRequestID)
	return
}
