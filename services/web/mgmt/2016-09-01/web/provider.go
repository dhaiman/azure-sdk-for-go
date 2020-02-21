package web

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

// ProviderClient is the webSite Management Client
type ProviderClient struct {
	BaseClient
}

// NewProviderClient creates an instance of the ProviderClient client.
func NewProviderClient(subscriptionID string) ProviderClient {
	return NewProviderClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewProviderClientWithBaseURI creates an instance of the ProviderClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewProviderClientWithBaseURI(baseURI string, subscriptionID string) ProviderClient {
	return ProviderClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// GetAvailableStacks get available application frameworks and their versions
func (client ProviderClient) GetAvailableStacks(ctx context.Context, osTypeSelected string) (result ApplicationStackCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProviderClient.GetAvailableStacks")
		defer func() {
			sc := -1
			if result.asc.Response.Response != nil {
				sc = result.asc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.getAvailableStacksNextResults
	req, err := client.GetAvailableStacksPreparer(ctx, osTypeSelected)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ProviderClient", "GetAvailableStacks", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetAvailableStacksSender(req)
	if err != nil {
		result.asc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "web.ProviderClient", "GetAvailableStacks", resp, "Failure sending request")
		return
	}

	result.asc, err = client.GetAvailableStacksResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ProviderClient", "GetAvailableStacks", resp, "Failure responding to request")
	}

	return
}

// GetAvailableStacksPreparer prepares the GetAvailableStacks request.
func (client ProviderClient) GetAvailableStacksPreparer(ctx context.Context, osTypeSelected string) (*http.Request, error) {
	const APIVersion = "2016-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(string(osTypeSelected)) > 0 {
		queryParameters["osTypeSelected"] = autorest.Encode("query", osTypeSelected)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Web/availableStacks"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetAvailableStacksSender sends the GetAvailableStacks request. The method will close the
// http.Response Body if it receives an error.
func (client ProviderClient) GetAvailableStacksSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetAvailableStacksResponder handles the response to the GetAvailableStacks request. The method always
// closes the http.Response Body.
func (client ProviderClient) GetAvailableStacksResponder(resp *http.Response) (result ApplicationStackCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// getAvailableStacksNextResults retrieves the next set of results, if any.
func (client ProviderClient) getAvailableStacksNextResults(ctx context.Context, lastResults ApplicationStackCollection) (result ApplicationStackCollection, err error) {
	req, err := lastResults.applicationStackCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web.ProviderClient", "getAvailableStacksNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.GetAvailableStacksSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web.ProviderClient", "getAvailableStacksNextResults", resp, "Failure sending next results request")
	}
	result, err = client.GetAvailableStacksResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ProviderClient", "getAvailableStacksNextResults", resp, "Failure responding to next results request")
	}
	return
}

// GetAvailableStacksComplete enumerates all values, automatically crossing page boundaries as required.
func (client ProviderClient) GetAvailableStacksComplete(ctx context.Context, osTypeSelected string) (result ApplicationStackCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProviderClient.GetAvailableStacks")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.GetAvailableStacks(ctx, osTypeSelected)
	return
}

// GetAvailableStacksOnPrem get available application frameworks and their versions
func (client ProviderClient) GetAvailableStacksOnPrem(ctx context.Context, osTypeSelected string) (result ApplicationStackCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProviderClient.GetAvailableStacksOnPrem")
		defer func() {
			sc := -1
			if result.asc.Response.Response != nil {
				sc = result.asc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.getAvailableStacksOnPremNextResults
	req, err := client.GetAvailableStacksOnPremPreparer(ctx, osTypeSelected)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ProviderClient", "GetAvailableStacksOnPrem", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetAvailableStacksOnPremSender(req)
	if err != nil {
		result.asc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "web.ProviderClient", "GetAvailableStacksOnPrem", resp, "Failure sending request")
		return
	}

	result.asc, err = client.GetAvailableStacksOnPremResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ProviderClient", "GetAvailableStacksOnPrem", resp, "Failure responding to request")
	}

	return
}

// GetAvailableStacksOnPremPreparer prepares the GetAvailableStacksOnPrem request.
func (client ProviderClient) GetAvailableStacksOnPremPreparer(ctx context.Context, osTypeSelected string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(string(osTypeSelected)) > 0 {
		queryParameters["osTypeSelected"] = autorest.Encode("query", osTypeSelected)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Web/availableStacks", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetAvailableStacksOnPremSender sends the GetAvailableStacksOnPrem request. The method will close the
// http.Response Body if it receives an error.
func (client ProviderClient) GetAvailableStacksOnPremSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetAvailableStacksOnPremResponder handles the response to the GetAvailableStacksOnPrem request. The method always
// closes the http.Response Body.
func (client ProviderClient) GetAvailableStacksOnPremResponder(resp *http.Response) (result ApplicationStackCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// getAvailableStacksOnPremNextResults retrieves the next set of results, if any.
func (client ProviderClient) getAvailableStacksOnPremNextResults(ctx context.Context, lastResults ApplicationStackCollection) (result ApplicationStackCollection, err error) {
	req, err := lastResults.applicationStackCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web.ProviderClient", "getAvailableStacksOnPremNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.GetAvailableStacksOnPremSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web.ProviderClient", "getAvailableStacksOnPremNextResults", resp, "Failure sending next results request")
	}
	result, err = client.GetAvailableStacksOnPremResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ProviderClient", "getAvailableStacksOnPremNextResults", resp, "Failure responding to next results request")
	}
	return
}

// GetAvailableStacksOnPremComplete enumerates all values, automatically crossing page boundaries as required.
func (client ProviderClient) GetAvailableStacksOnPremComplete(ctx context.Context, osTypeSelected string) (result ApplicationStackCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProviderClient.GetAvailableStacksOnPrem")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.GetAvailableStacksOnPrem(ctx, osTypeSelected)
	return
}

// ListOperations gets all available operations for the Microsoft.Web resource provider. Also exposes resource metric
// definitions
func (client ProviderClient) ListOperations(ctx context.Context) (result CsmOperationCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProviderClient.ListOperations")
		defer func() {
			sc := -1
			if result.coc.Response.Response != nil {
				sc = result.coc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listOperationsNextResults
	req, err := client.ListOperationsPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ProviderClient", "ListOperations", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListOperationsSender(req)
	if err != nil {
		result.coc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "web.ProviderClient", "ListOperations", resp, "Failure sending request")
		return
	}

	result.coc, err = client.ListOperationsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ProviderClient", "ListOperations", resp, "Failure responding to request")
	}

	return
}

// ListOperationsPreparer prepares the ListOperations request.
func (client ProviderClient) ListOperationsPreparer(ctx context.Context) (*http.Request, error) {
	const APIVersion = "2016-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Web/operations"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListOperationsSender sends the ListOperations request. The method will close the
// http.Response Body if it receives an error.
func (client ProviderClient) ListOperationsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListOperationsResponder handles the response to the ListOperations request. The method always
// closes the http.Response Body.
func (client ProviderClient) ListOperationsResponder(resp *http.Response) (result CsmOperationCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listOperationsNextResults retrieves the next set of results, if any.
func (client ProviderClient) listOperationsNextResults(ctx context.Context, lastResults CsmOperationCollection) (result CsmOperationCollection, err error) {
	req, err := lastResults.csmOperationCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web.ProviderClient", "listOperationsNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListOperationsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web.ProviderClient", "listOperationsNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListOperationsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ProviderClient", "listOperationsNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListOperationsComplete enumerates all values, automatically crossing page boundaries as required.
func (client ProviderClient) ListOperationsComplete(ctx context.Context) (result CsmOperationCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProviderClient.ListOperations")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListOperations(ctx)
	return
}
