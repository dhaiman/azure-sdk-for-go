package sql

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

// QueryStatisticsClient is the the Azure SQL Database management API provides a RESTful set of web services that
// interact with Azure SQL Database services to manage your databases. The API enables you to create, retrieve, update,
// and delete databases.
type QueryStatisticsClient struct {
	BaseClient
}

// NewQueryStatisticsClient creates an instance of the QueryStatisticsClient client.
func NewQueryStatisticsClient(subscriptionID string) QueryStatisticsClient {
	return NewQueryStatisticsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewQueryStatisticsClientWithBaseURI creates an instance of the QueryStatisticsClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewQueryStatisticsClientWithBaseURI(baseURI string, subscriptionID string) QueryStatisticsClient {
	return QueryStatisticsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// ListByQuery lists a query's statistics.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
// databaseName - the name of the database.
// queryID - the id of the query
func (client QueryStatisticsClient) ListByQuery(ctx context.Context, resourceGroupName string, serverName string, databaseName string, queryID string) (result QueryStatisticListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/QueryStatisticsClient.ListByQuery")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListByQueryPreparer(ctx, resourceGroupName, serverName, databaseName, queryID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.QueryStatisticsClient", "ListByQuery", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByQuerySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.QueryStatisticsClient", "ListByQuery", resp, "Failure sending request")
		return
	}

	result, err = client.ListByQueryResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.QueryStatisticsClient", "ListByQuery", resp, "Failure responding to request")
	}

	return
}

// ListByQueryPreparer prepares the ListByQuery request.
func (client QueryStatisticsClient) ListByQueryPreparer(ctx context.Context, resourceGroupName string, serverName string, databaseName string, queryID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"databaseName":      autorest.Encode("path", databaseName),
		"queryId":           autorest.Encode("path", queryID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/topQueries/{queryId}/statistics", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByQuerySender sends the ListByQuery request. The method will close the
// http.Response Body if it receives an error.
func (client QueryStatisticsClient) ListByQuerySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByQueryResponder handles the response to the ListByQuery request. The method always
// closes the http.Response Body.
func (client QueryStatisticsClient) ListByQueryResponder(resp *http.Response) (result QueryStatisticListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
