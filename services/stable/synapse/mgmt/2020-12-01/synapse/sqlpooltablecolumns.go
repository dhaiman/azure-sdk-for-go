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

// SQLPoolTableColumnsClient is the azure Synapse Analytics Management Client
type SQLPoolTableColumnsClient struct {
	BaseClient
}

// NewSQLPoolTableColumnsClient creates an instance of the SQLPoolTableColumnsClient client.
func NewSQLPoolTableColumnsClient(subscriptionID string) SQLPoolTableColumnsClient {
	return NewSQLPoolTableColumnsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewSQLPoolTableColumnsClientWithBaseURI creates an instance of the SQLPoolTableColumnsClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewSQLPoolTableColumnsClientWithBaseURI(baseURI string, subscriptionID string) SQLPoolTableColumnsClient {
	return SQLPoolTableColumnsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// ListByTableName gets columns in a given table in a SQL pool.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace
// SQLPoolName - SQL pool name
// schemaName - the name of the schema.
// tableName - the name of the table.
// filter - an OData filter expression that filters elements in the collection.
func (client SQLPoolTableColumnsClient) ListByTableName(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, schemaName string, tableName string, filter string) (result SQLPoolColumnListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SQLPoolTableColumnsClient.ListByTableName")
		defer func() {
			sc := -1
			if result.spclr.Response.Response != nil {
				sc = result.spclr.Response.Response.StatusCode
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
		return result, validation.NewError("synapse.SQLPoolTableColumnsClient", "ListByTableName", err.Error())
	}

	result.fn = client.listByTableNameNextResults
	req, err := client.ListByTableNamePreparer(ctx, resourceGroupName, workspaceName, SQLPoolName, schemaName, tableName, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.SQLPoolTableColumnsClient", "ListByTableName", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByTableNameSender(req)
	if err != nil {
		result.spclr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "synapse.SQLPoolTableColumnsClient", "ListByTableName", resp, "Failure sending request")
		return
	}

	result.spclr, err = client.ListByTableNameResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.SQLPoolTableColumnsClient", "ListByTableName", resp, "Failure responding to request")
	}
	if result.spclr.hasNextLink() && result.spclr.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListByTableNamePreparer prepares the ListByTableName request.
func (client SQLPoolTableColumnsClient) ListByTableNamePreparer(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, schemaName string, tableName string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"schemaName":        autorest.Encode("path", schemaName),
		"sqlPoolName":       autorest.Encode("path", SQLPoolName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"tableName":         autorest.Encode("path", tableName),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2020-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/schemas/{schemaName}/tables/{tableName}/columns", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByTableNameSender sends the ListByTableName request. The method will close the
// http.Response Body if it receives an error.
func (client SQLPoolTableColumnsClient) ListByTableNameSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByTableNameResponder handles the response to the ListByTableName request. The method always
// closes the http.Response Body.
func (client SQLPoolTableColumnsClient) ListByTableNameResponder(resp *http.Response) (result SQLPoolColumnListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByTableNameNextResults retrieves the next set of results, if any.
func (client SQLPoolTableColumnsClient) listByTableNameNextResults(ctx context.Context, lastResults SQLPoolColumnListResult) (result SQLPoolColumnListResult, err error) {
	req, err := lastResults.sQLPoolColumnListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "synapse.SQLPoolTableColumnsClient", "listByTableNameNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByTableNameSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "synapse.SQLPoolTableColumnsClient", "listByTableNameNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByTableNameResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.SQLPoolTableColumnsClient", "listByTableNameNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByTableNameComplete enumerates all values, automatically crossing page boundaries as required.
func (client SQLPoolTableColumnsClient) ListByTableNameComplete(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, schemaName string, tableName string, filter string) (result SQLPoolColumnListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SQLPoolTableColumnsClient.ListByTableName")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByTableName(ctx, resourceGroupName, workspaceName, SQLPoolName, schemaName, tableName, filter)
	return
}
