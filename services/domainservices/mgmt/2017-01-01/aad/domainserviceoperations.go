package aad

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

// DomainServiceOperationsClient is the the AAD Domain Services API.
type DomainServiceOperationsClient struct {
	BaseClient
}

// NewDomainServiceOperationsClient creates an instance of the DomainServiceOperationsClient client.
func NewDomainServiceOperationsClient(subscriptionID string) DomainServiceOperationsClient {
	return NewDomainServiceOperationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewDomainServiceOperationsClientWithBaseURI creates an instance of the DomainServiceOperationsClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewDomainServiceOperationsClientWithBaseURI(baseURI string, subscriptionID string) DomainServiceOperationsClient {
	return DomainServiceOperationsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List lists all the available Domain Services operations.
func (client DomainServiceOperationsClient) List(ctx context.Context) (result OperationEntityListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DomainServiceOperationsClient.List")
		defer func() {
			sc := -1
			if result.oelr.Response.Response != nil {
				sc = result.oelr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "aad.DomainServiceOperationsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.oelr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "aad.DomainServiceOperationsClient", "List", resp, "Failure sending request")
		return
	}

	result.oelr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "aad.DomainServiceOperationsClient", "List", resp, "Failure responding to request")
	}
	if result.oelr.hasNextLink() && result.oelr.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListPreparer prepares the List request.
func (client DomainServiceOperationsClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	const APIVersion = "2017-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.AAD/operations"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client DomainServiceOperationsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client DomainServiceOperationsClient) ListResponder(resp *http.Response) (result OperationEntityListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client DomainServiceOperationsClient) listNextResults(ctx context.Context, lastResults OperationEntityListResult) (result OperationEntityListResult, err error) {
	req, err := lastResults.operationEntityListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "aad.DomainServiceOperationsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "aad.DomainServiceOperationsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "aad.DomainServiceOperationsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client DomainServiceOperationsClient) ListComplete(ctx context.Context) (result OperationEntityListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DomainServiceOperationsClient.List")
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
