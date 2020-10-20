package batch

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

// PrivateLinkResourceClient is the client for the PrivateLinkResource methods of the Batch service.
type PrivateLinkResourceClient struct {
	BaseClient
}

// NewPrivateLinkResourceClient creates an instance of the PrivateLinkResourceClient client.
func NewPrivateLinkResourceClient(subscriptionID string) PrivateLinkResourceClient {
	return NewPrivateLinkResourceClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewPrivateLinkResourceClientWithBaseURI creates an instance of the PrivateLinkResourceClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewPrivateLinkResourceClientWithBaseURI(baseURI string, subscriptionID string) PrivateLinkResourceClient {
	return PrivateLinkResourceClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get gets information about the specified private link resource.
// Parameters:
// resourceGroupName - the name of the resource group that contains the Batch account.
// accountName - the name of the Batch account.
// privateLinkResourceName - the private link resource name. This must be unique within the account.
func (client PrivateLinkResourceClient) Get(ctx context.Context, resourceGroupName string, accountName string, privateLinkResourceName string) (result PrivateLinkResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateLinkResourceClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]+$`, Chain: nil}}},
		{TargetValue: privateLinkResourceName,
			Constraints: []validation.Constraint{{Target: "privateLinkResourceName", Name: validation.MaxLength, Rule: 101, Chain: nil},
				{Target: "privateLinkResourceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "privateLinkResourceName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_-]+\.?[a-fA-F0-9-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("batch.PrivateLinkResourceClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, accountName, privateLinkResourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.PrivateLinkResourceClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "batch.PrivateLinkResourceClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.PrivateLinkResourceClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client PrivateLinkResourceClient) GetPreparer(ctx context.Context, resourceGroupName string, accountName string, privateLinkResourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":             autorest.Encode("path", accountName),
		"privateLinkResourceName": autorest.Encode("path", privateLinkResourceName),
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/privateLinkResources/{privateLinkResourceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateLinkResourceClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client PrivateLinkResourceClient) GetResponder(resp *http.Response) (result PrivateLinkResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByBatchAccount lists all of the private link resources in the specified account.
// Parameters:
// resourceGroupName - the name of the resource group that contains the Batch account.
// accountName - the name of the Batch account.
// maxresults - the maximum number of items to return in the response.
func (client PrivateLinkResourceClient) ListByBatchAccount(ctx context.Context, resourceGroupName string, accountName string, maxresults *int32) (result ListPrivateLinkResourcesResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateLinkResourceClient.ListByBatchAccount")
		defer func() {
			sc := -1
			if result.lplrr.Response.Response != nil {
				sc = result.lplrr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("batch.PrivateLinkResourceClient", "ListByBatchAccount", err.Error())
	}

	result.fn = client.listByBatchAccountNextResults
	req, err := client.ListByBatchAccountPreparer(ctx, resourceGroupName, accountName, maxresults)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.PrivateLinkResourceClient", "ListByBatchAccount", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByBatchAccountSender(req)
	if err != nil {
		result.lplrr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "batch.PrivateLinkResourceClient", "ListByBatchAccount", resp, "Failure sending request")
		return
	}

	result.lplrr, err = client.ListByBatchAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.PrivateLinkResourceClient", "ListByBatchAccount", resp, "Failure responding to request")
	}
	if result.lplrr.hasNextLink() && result.lplrr.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListByBatchAccountPreparer prepares the ListByBatchAccount request.
func (client PrivateLinkResourceClient) ListByBatchAccountPreparer(ctx context.Context, resourceGroupName string, accountName string, maxresults *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if maxresults != nil {
		queryParameters["maxresults"] = autorest.Encode("query", *maxresults)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/privateLinkResources", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByBatchAccountSender sends the ListByBatchAccount request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateLinkResourceClient) ListByBatchAccountSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByBatchAccountResponder handles the response to the ListByBatchAccount request. The method always
// closes the http.Response Body.
func (client PrivateLinkResourceClient) ListByBatchAccountResponder(resp *http.Response) (result ListPrivateLinkResourcesResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByBatchAccountNextResults retrieves the next set of results, if any.
func (client PrivateLinkResourceClient) listByBatchAccountNextResults(ctx context.Context, lastResults ListPrivateLinkResourcesResult) (result ListPrivateLinkResourcesResult, err error) {
	req, err := lastResults.listPrivateLinkResourcesResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "batch.PrivateLinkResourceClient", "listByBatchAccountNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByBatchAccountSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "batch.PrivateLinkResourceClient", "listByBatchAccountNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByBatchAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.PrivateLinkResourceClient", "listByBatchAccountNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByBatchAccountComplete enumerates all values, automatically crossing page boundaries as required.
func (client PrivateLinkResourceClient) ListByBatchAccountComplete(ctx context.Context, resourceGroupName string, accountName string, maxresults *int32) (result ListPrivateLinkResourcesResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateLinkResourceClient.ListByBatchAccount")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByBatchAccount(ctx, resourceGroupName, accountName, maxresults)
	return
}
