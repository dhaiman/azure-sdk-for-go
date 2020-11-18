package synapse

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

// SQLPoolOperationsClient is the azure Synapse Analytics Management Client
type SQLPoolOperationsClient struct {
	BaseClient
}

// NewSQLPoolOperationsClient creates an instance of the SQLPoolOperationsClient client.
func NewSQLPoolOperationsClient(subscriptionID string) SQLPoolOperationsClient {
	return NewSQLPoolOperationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewSQLPoolOperationsClientWithBaseURI creates an instance of the SQLPoolOperationsClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewSQLPoolOperationsClientWithBaseURI(baseURI string, subscriptionID string) SQLPoolOperationsClient {
	return SQLPoolOperationsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List gets a list of operations performed on the SQL pool.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace
// SQLPoolName - SQL pool name
func (client SQLPoolOperationsClient) List(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string) (result SQLPoolBlobAuditingPolicySQLPoolOperationListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SQLPoolOperationsClient.List")
		defer func() {
			sc := -1
			if result.spbapspolr.Response.Response != nil {
				sc = result.spbapspolr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("synapse.SQLPoolOperationsClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, workspaceName, SQLPoolName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.SQLPoolOperationsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.spbapspolr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "synapse.SQLPoolOperationsClient", "List", resp, "Failure sending request")
		return
	}

	result.spbapspolr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.SQLPoolOperationsClient", "List", resp, "Failure responding to request")
	}
	if result.spbapspolr.hasNextLink() && result.spbapspolr.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListPreparer prepares the List request.
func (client SQLPoolOperationsClient) ListPreparer(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"sqlPoolName":       autorest.Encode("path", SQLPoolName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2020-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/operations", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client SQLPoolOperationsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client SQLPoolOperationsClient) ListResponder(resp *http.Response) (result SQLPoolBlobAuditingPolicySQLPoolOperationListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client SQLPoolOperationsClient) listNextResults(ctx context.Context, lastResults SQLPoolBlobAuditingPolicySQLPoolOperationListResult) (result SQLPoolBlobAuditingPolicySQLPoolOperationListResult, err error) {
	req, err := lastResults.sQLPoolBlobAuditingPolicySQLPoolOperationListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "synapse.SQLPoolOperationsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "synapse.SQLPoolOperationsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.SQLPoolOperationsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client SQLPoolOperationsClient) ListComplete(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string) (result SQLPoolBlobAuditingPolicySQLPoolOperationListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SQLPoolOperationsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, workspaceName, SQLPoolName)
	return
}
