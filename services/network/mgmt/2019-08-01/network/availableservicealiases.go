package network

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
	"net/http"
)

// AvailableServiceAliasesClient is the network Client
type AvailableServiceAliasesClient struct {
	BaseClient
}

// NewAvailableServiceAliasesClient creates an instance of the AvailableServiceAliasesClient client.
func NewAvailableServiceAliasesClient(subscriptionID string) AvailableServiceAliasesClient {
	return NewAvailableServiceAliasesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewAvailableServiceAliasesClientWithBaseURI creates an instance of the AvailableServiceAliasesClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewAvailableServiceAliasesClientWithBaseURI(baseURI string, subscriptionID string) AvailableServiceAliasesClient {
	return AvailableServiceAliasesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List gets all available service aliases for this subscription in this region.
// Parameters:
// location - the location.
func (client AvailableServiceAliasesClient) List(ctx context.Context, location string) (result AvailableServiceAliasesResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AvailableServiceAliasesClient.List")
		defer func() {
			sc := -1
			if result.asar.Response.Response != nil {
				sc = result.asar.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, location)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.AvailableServiceAliasesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.asar.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.AvailableServiceAliasesClient", "List", resp, "Failure sending request")
		return
	}

	result.asar, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.AvailableServiceAliasesClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client AvailableServiceAliasesClient) ListPreparer(ctx context.Context, location string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       autorest.Encode("path", location),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Network/locations/{location}/availableServiceAliases", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client AvailableServiceAliasesClient) ListSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client AvailableServiceAliasesClient) ListResponder(resp *http.Response) (result AvailableServiceAliasesResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client AvailableServiceAliasesClient) listNextResults(ctx context.Context, lastResults AvailableServiceAliasesResult) (result AvailableServiceAliasesResult, err error) {
	req, err := lastResults.availableServiceAliasesResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.AvailableServiceAliasesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.AvailableServiceAliasesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.AvailableServiceAliasesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client AvailableServiceAliasesClient) ListComplete(ctx context.Context, location string) (result AvailableServiceAliasesResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AvailableServiceAliasesClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, location)
	return
}

// ListByResourceGroup gets all available service aliases for this resource group in this region.
// Parameters:
// resourceGroupName - the name of the resource group.
// location - the location.
func (client AvailableServiceAliasesClient) ListByResourceGroup(ctx context.Context, resourceGroupName string, location string) (result AvailableServiceAliasesResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AvailableServiceAliasesClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.asar.Response.Response != nil {
				sc = result.asar.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName, location)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.AvailableServiceAliasesClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.asar.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.AvailableServiceAliasesClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.asar, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.AvailableServiceAliasesClient", "ListByResourceGroup", resp, "Failure responding to request")
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client AvailableServiceAliasesClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string, location string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":          autorest.Encode("path", location),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/locations/{location}/availableServiceAliases", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client AvailableServiceAliasesClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client AvailableServiceAliasesClient) ListByResourceGroupResponder(resp *http.Response) (result AvailableServiceAliasesResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client AvailableServiceAliasesClient) listByResourceGroupNextResults(ctx context.Context, lastResults AvailableServiceAliasesResult) (result AvailableServiceAliasesResult, err error) {
	req, err := lastResults.availableServiceAliasesResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.AvailableServiceAliasesClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.AvailableServiceAliasesClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.AvailableServiceAliasesClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client AvailableServiceAliasesClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, location string) (result AvailableServiceAliasesResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AvailableServiceAliasesClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName, location)
	return
}
