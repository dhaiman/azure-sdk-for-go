package documentdb

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
    "github.com/Azure/go-autorest/autorest"
    "github.com/Azure/go-autorest/autorest/azure"
    "net/http"
    "context"
    "github.com/Azure/go-autorest/tracing"
    "github.com/Azure/go-autorest/autorest/validation"
)

// DatabaseAccountRegionClient is the azure Cosmos DB Database Service Resource Provider REST API
type DatabaseAccountRegionClient struct {
    BaseClient
}
// NewDatabaseAccountRegionClient creates an instance of the DatabaseAccountRegionClient client.
func NewDatabaseAccountRegionClient(subscriptionID string, subscriptionID1 string) DatabaseAccountRegionClient {
    return NewDatabaseAccountRegionClientWithBaseURI(DefaultBaseURI, subscriptionID, subscriptionID1)
}

// NewDatabaseAccountRegionClientWithBaseURI creates an instance of the DatabaseAccountRegionClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
    func NewDatabaseAccountRegionClientWithBaseURI(baseURI string, subscriptionID string, subscriptionID1 string) DatabaseAccountRegionClient {
        return DatabaseAccountRegionClient{ NewWithBaseURI(baseURI, subscriptionID, subscriptionID1)}
    }

// ListMetrics retrieves the metrics determined by the given filter for the given database account and region.
    // Parameters:
        // resourceGroupName - name of an Azure resource group.
        // accountName - cosmos DB database account name.
        // region - cosmos DB region, with spaces between words and each word capitalized.
        // filter - an OData filter expression that describes a subset of metrics to return. The parameters that can be
        // filtered are name.value (name of the metric, can have an or of multiple names), startTime, endTime, and
        // timeGrain. The supported operator is eq.
func (client DatabaseAccountRegionClient) ListMetrics(ctx context.Context, resourceGroupName string, accountName string, region string, filter string) (result MetricListResult, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/DatabaseAccountRegionClient.ListMetrics")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: resourceGroupName,
             Constraints: []validation.Constraint{	{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil },
            	{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil },
            	{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil }}},
            { TargetValue: accountName,
             Constraints: []validation.Constraint{	{Target: "accountName", Name: validation.MaxLength, Rule: 50, Chain: nil },
            	{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil },
            	{Target: "accountName", Name: validation.Pattern, Rule: `^[a-z0-9]+(-[a-z0-9]+)*`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("documentdb.DatabaseAccountRegionClient", "ListMetrics", err.Error())
            }

                req, err := client.ListMetricsPreparer(ctx, resourceGroupName, accountName, region, filter)
    if err != nil {
    err = autorest.NewErrorWithError(err, "documentdb.DatabaseAccountRegionClient", "ListMetrics", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListMetricsSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "documentdb.DatabaseAccountRegionClient", "ListMetrics", resp, "Failure sending request")
            return
            }

            result, err = client.ListMetricsResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "documentdb.DatabaseAccountRegionClient", "ListMetrics", resp, "Failure responding to request")
            }

    return
    }

    // ListMetricsPreparer prepares the ListMetrics request.
    func (client DatabaseAccountRegionClient) ListMetricsPreparer(ctx context.Context, resourceGroupName string, accountName string, region string, filter string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "accountName": autorest.Encode("path",accountName),
            "region": autorest.Encode("path",region),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2019-08-01"
        queryParameters := map[string]interface{} {
        "$filter": autorest.Encode("query",filter),
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/region/{region}/metrics",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListMetricsSender sends the ListMetrics request. The method will close the
    // http.Response Body if it receives an error.
    func (client DatabaseAccountRegionClient) ListMetricsSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

// ListMetricsResponder handles the response to the ListMetrics request. The method always
// closes the http.Response Body.
func (client DatabaseAccountRegionClient) ListMetricsResponder(resp *http.Response) (result MetricListResult, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

