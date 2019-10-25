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
    "github.com/Azure/go-autorest/autorest"
    "github.com/Azure/go-autorest/autorest/azure"
    "net/http"
    "context"
    "github.com/Azure/go-autorest/tracing"
    "github.com/Azure/go-autorest/autorest/validation"
)

// ElasticPoolsClient is the the Azure SQL Database management API provides a RESTful set of web services that interact
// with Azure SQL Database services to manage your databases. The API enables you to create, retrieve, update, and
// delete databases.
type ElasticPoolsClient struct {
    BaseClient
}
// NewElasticPoolsClient creates an instance of the ElasticPoolsClient client.
func NewElasticPoolsClient(subscriptionID string) ElasticPoolsClient {
    return NewElasticPoolsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewElasticPoolsClientWithBaseURI creates an instance of the ElasticPoolsClient client.
    func NewElasticPoolsClientWithBaseURI(baseURI string, subscriptionID string) ElasticPoolsClient {
        return ElasticPoolsClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// CreateOrUpdate creates or updates an elastic pool.
    // Parameters:
        // resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
        // from the Azure Resource Manager API or the portal.
        // serverName - the name of the server.
        // elasticPoolName - the name of the elastic pool.
        // parameters - the elastic pool parameters.
func (client ElasticPoolsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string, parameters ElasticPool) (result ElasticPoolsCreateOrUpdateFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ElasticPoolsClient.CreateOrUpdate")
        defer func() {
            sc := -1
            if result.Response() != nil {
                sc = result.Response().StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: parameters,
             Constraints: []validation.Constraint{	{Target: "parameters.Sku", Name: validation.Null, Rule: false ,
            Chain: []validation.Constraint{	{Target: "parameters.Sku.Name", Name: validation.Null, Rule: true, Chain: nil },
            }}}}}); err != nil {
            return result, validation.NewError("sql.ElasticPoolsClient", "CreateOrUpdate", err.Error())
            }

                req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, serverName, elasticPoolName, parameters)
    if err != nil {
    err = autorest.NewErrorWithError(err, "sql.ElasticPoolsClient", "CreateOrUpdate", nil , "Failure preparing request")
    return
    }

            result, err = client.CreateOrUpdateSender(req)
            if err != nil {
            err = autorest.NewErrorWithError(err, "sql.ElasticPoolsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
            return
            }

    return
    }

    // CreateOrUpdatePreparer prepares the CreateOrUpdate request.
    func (client ElasticPoolsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string, parameters ElasticPool) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "elasticPoolName": autorest.Encode("path",elasticPoolName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "serverName": autorest.Encode("path",serverName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2017-10-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

                    parameters.Kind = nil
    preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPut(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/elasticPools/{elasticPoolName}",pathParameters),
    autorest.WithJSON(parameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
    // http.Response Body if it receives an error.
    func (client ElasticPoolsClient) CreateOrUpdateSender(req *http.Request) (future ElasticPoolsCreateOrUpdateFuture, err error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            var resp *http.Response
            resp, err = autorest.SendWithSender(client, req, sd...)
            if err != nil {
            return
            }
            future.Future, err = azure.NewFutureFromResponse(resp)
            return
            }

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ElasticPoolsClient) CreateOrUpdateResponder(resp *http.Response) (result ElasticPool, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusCreated,http.StatusAccepted),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// Delete deletes an elastic pool.
    // Parameters:
        // resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
        // from the Azure Resource Manager API or the portal.
        // serverName - the name of the server.
        // elasticPoolName - the name of the elastic pool.
func (client ElasticPoolsClient) Delete(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string) (result ElasticPoolsDeleteFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ElasticPoolsClient.Delete")
        defer func() {
            sc := -1
            if result.Response() != nil {
                sc = result.Response().StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.DeletePreparer(ctx, resourceGroupName, serverName, elasticPoolName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "sql.ElasticPoolsClient", "Delete", nil , "Failure preparing request")
    return
    }

            result, err = client.DeleteSender(req)
            if err != nil {
            err = autorest.NewErrorWithError(err, "sql.ElasticPoolsClient", "Delete", result.Response(), "Failure sending request")
            return
            }

    return
    }

    // DeletePreparer prepares the Delete request.
    func (client ElasticPoolsClient) DeletePreparer(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "elasticPoolName": autorest.Encode("path",elasticPoolName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "serverName": autorest.Encode("path",serverName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2017-10-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsDelete(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/elasticPools/{elasticPoolName}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // DeleteSender sends the Delete request. The method will close the
    // http.Response Body if it receives an error.
    func (client ElasticPoolsClient) DeleteSender(req *http.Request) (future ElasticPoolsDeleteFuture, err error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            var resp *http.Response
            resp, err = autorest.SendWithSender(client, req, sd...)
            if err != nil {
            return
            }
            future.Future, err = azure.NewFutureFromResponse(resp)
            return
            }

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ElasticPoolsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusAccepted,http.StatusNoContent),
    autorest.ByClosing())
    result.Response = resp
        return
    }

// Failover failovers an elastic pool.
    // Parameters:
        // resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
        // from the Azure Resource Manager API or the portal.
        // serverName - the name of the server.
        // elasticPoolName - the name of the elastic pool to failover.
func (client ElasticPoolsClient) Failover(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string) (result ElasticPoolsFailoverFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ElasticPoolsClient.Failover")
        defer func() {
            sc := -1
            if result.Response() != nil {
                sc = result.Response().StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.FailoverPreparer(ctx, resourceGroupName, serverName, elasticPoolName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "sql.ElasticPoolsClient", "Failover", nil , "Failure preparing request")
    return
    }

            result, err = client.FailoverSender(req)
            if err != nil {
            err = autorest.NewErrorWithError(err, "sql.ElasticPoolsClient", "Failover", result.Response(), "Failure sending request")
            return
            }

    return
    }

    // FailoverPreparer prepares the Failover request.
    func (client ElasticPoolsClient) FailoverPreparer(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "elasticPoolName": autorest.Encode("path",elasticPoolName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "serverName": autorest.Encode("path",serverName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2018-06-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsPost(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/elasticPools/{elasticPoolName}/failover",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // FailoverSender sends the Failover request. The method will close the
    // http.Response Body if it receives an error.
    func (client ElasticPoolsClient) FailoverSender(req *http.Request) (future ElasticPoolsFailoverFuture, err error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            var resp *http.Response
            resp, err = autorest.SendWithSender(client, req, sd...)
            if err != nil {
            return
            }
            future.Future, err = azure.NewFutureFromResponse(resp)
            return
            }

// FailoverResponder handles the response to the Failover request. The method always
// closes the http.Response Body.
func (client ElasticPoolsClient) FailoverResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusAccepted),
    autorest.ByClosing())
    result.Response = resp
        return
    }

// Get gets an elastic pool.
    // Parameters:
        // resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
        // from the Azure Resource Manager API or the portal.
        // serverName - the name of the server.
        // elasticPoolName - the name of the elastic pool.
func (client ElasticPoolsClient) Get(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string) (result ElasticPool, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ElasticPoolsClient.Get")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.GetPreparer(ctx, resourceGroupName, serverName, elasticPoolName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "sql.ElasticPoolsClient", "Get", nil , "Failure preparing request")
    return
    }

            resp, err := client.GetSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "sql.ElasticPoolsClient", "Get", resp, "Failure sending request")
            return
            }

            result, err = client.GetResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "sql.ElasticPoolsClient", "Get", resp, "Failure responding to request")
            }

    return
    }

    // GetPreparer prepares the Get request.
    func (client ElasticPoolsClient) GetPreparer(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "elasticPoolName": autorest.Encode("path",elasticPoolName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "serverName": autorest.Encode("path",serverName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2017-10-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/elasticPools/{elasticPoolName}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client ElasticPoolsClient) GetSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ElasticPoolsClient) GetResponder(resp *http.Response) (result ElasticPool, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// ListByServer gets all elastic pools in a server.
    // Parameters:
        // resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
        // from the Azure Resource Manager API or the portal.
        // serverName - the name of the server.
        // skip - the number of elements in the collection to skip.
func (client ElasticPoolsClient) ListByServer(ctx context.Context, resourceGroupName string, serverName string, skip *int32) (result ElasticPoolListResultPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ElasticPoolsClient.ListByServer")
        defer func() {
            sc := -1
            if result.eplr.Response.Response != nil {
                sc = result.eplr.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
                result.fn = client.listByServerNextResults
    req, err := client.ListByServerPreparer(ctx, resourceGroupName, serverName, skip)
    if err != nil {
    err = autorest.NewErrorWithError(err, "sql.ElasticPoolsClient", "ListByServer", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListByServerSender(req)
            if err != nil {
            result.eplr.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "sql.ElasticPoolsClient", "ListByServer", resp, "Failure sending request")
            return
            }

            result.eplr, err = client.ListByServerResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "sql.ElasticPoolsClient", "ListByServer", resp, "Failure responding to request")
            }

    return
    }

    // ListByServerPreparer prepares the ListByServer request.
    func (client ElasticPoolsClient) ListByServerPreparer(ctx context.Context, resourceGroupName string, serverName string, skip *int32) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "serverName": autorest.Encode("path",serverName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2017-10-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }
            if skip != nil {
            queryParameters["$skip"] = autorest.Encode("query",*skip)
            }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/elasticPools",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListByServerSender sends the ListByServer request. The method will close the
    // http.Response Body if it receives an error.
    func (client ElasticPoolsClient) ListByServerSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// ListByServerResponder handles the response to the ListByServer request. The method always
// closes the http.Response Body.
func (client ElasticPoolsClient) ListByServerResponder(resp *http.Response) (result ElasticPoolListResult, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

            // listByServerNextResults retrieves the next set of results, if any.
            func (client ElasticPoolsClient) listByServerNextResults(ctx context.Context, lastResults ElasticPoolListResult) (result ElasticPoolListResult, err error) {
            req, err := lastResults.elasticPoolListResultPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "sql.ElasticPoolsClient", "listByServerNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListByServerSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "sql.ElasticPoolsClient", "listByServerNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListByServerResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "sql.ElasticPoolsClient", "listByServerNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

    // ListByServerComplete enumerates all values, automatically crossing page boundaries as required.
    func (client ElasticPoolsClient) ListByServerComplete(ctx context.Context, resourceGroupName string, serverName string, skip *int32) (result ElasticPoolListResultIterator, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/ElasticPoolsClient.ListByServer")
            defer func() {
                sc := -1
                if result.Response().Response.Response != nil {
                    sc = result.page.Response().Response.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
     }
        result.page, err = client.ListByServer(ctx, resourceGroupName, serverName, skip)
                return
        }

// ListMetricDefinitions returns elastic pool metric definitions.
    // Parameters:
        // resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
        // from the Azure Resource Manager API or the portal.
        // serverName - the name of the server.
        // elasticPoolName - the name of the elastic pool.
func (client ElasticPoolsClient) ListMetricDefinitions(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string) (result MetricDefinitionListResult, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ElasticPoolsClient.ListMetricDefinitions")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.ListMetricDefinitionsPreparer(ctx, resourceGroupName, serverName, elasticPoolName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "sql.ElasticPoolsClient", "ListMetricDefinitions", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListMetricDefinitionsSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "sql.ElasticPoolsClient", "ListMetricDefinitions", resp, "Failure sending request")
            return
            }

            result, err = client.ListMetricDefinitionsResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "sql.ElasticPoolsClient", "ListMetricDefinitions", resp, "Failure responding to request")
            }

    return
    }

    // ListMetricDefinitionsPreparer prepares the ListMetricDefinitions request.
    func (client ElasticPoolsClient) ListMetricDefinitionsPreparer(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "elasticPoolName": autorest.Encode("path",elasticPoolName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "serverName": autorest.Encode("path",serverName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2014-04-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/elasticPools/{elasticPoolName}/metricDefinitions",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListMetricDefinitionsSender sends the ListMetricDefinitions request. The method will close the
    // http.Response Body if it receives an error.
    func (client ElasticPoolsClient) ListMetricDefinitionsSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// ListMetricDefinitionsResponder handles the response to the ListMetricDefinitions request. The method always
// closes the http.Response Body.
func (client ElasticPoolsClient) ListMetricDefinitionsResponder(resp *http.Response) (result MetricDefinitionListResult, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// ListMetrics returns elastic pool  metrics.
    // Parameters:
        // resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
        // from the Azure Resource Manager API or the portal.
        // serverName - the name of the server.
        // elasticPoolName - the name of the elastic pool.
        // filter - an OData filter expression that describes a subset of metrics to return.
func (client ElasticPoolsClient) ListMetrics(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string, filter string) (result MetricListResult, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ElasticPoolsClient.ListMetrics")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.ListMetricsPreparer(ctx, resourceGroupName, serverName, elasticPoolName, filter)
    if err != nil {
    err = autorest.NewErrorWithError(err, "sql.ElasticPoolsClient", "ListMetrics", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListMetricsSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "sql.ElasticPoolsClient", "ListMetrics", resp, "Failure sending request")
            return
            }

            result, err = client.ListMetricsResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "sql.ElasticPoolsClient", "ListMetrics", resp, "Failure responding to request")
            }

    return
    }

    // ListMetricsPreparer prepares the ListMetrics request.
    func (client ElasticPoolsClient) ListMetricsPreparer(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string, filter string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "elasticPoolName": autorest.Encode("path",elasticPoolName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "serverName": autorest.Encode("path",serverName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2014-04-01"
        queryParameters := map[string]interface{} {
        "$filter": autorest.Encode("query",filter),
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/elasticPools/{elasticPoolName}/metrics",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListMetricsSender sends the ListMetrics request. The method will close the
    // http.Response Body if it receives an error.
    func (client ElasticPoolsClient) ListMetricsSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// ListMetricsResponder handles the response to the ListMetrics request. The method always
// closes the http.Response Body.
func (client ElasticPoolsClient) ListMetricsResponder(resp *http.Response) (result MetricListResult, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// Update updates an elastic pool.
    // Parameters:
        // resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
        // from the Azure Resource Manager API or the portal.
        // serverName - the name of the server.
        // elasticPoolName - the name of the elastic pool.
        // parameters - the elastic pool update parameters.
func (client ElasticPoolsClient) Update(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string, parameters ElasticPoolUpdate) (result ElasticPoolsUpdateFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ElasticPoolsClient.Update")
        defer func() {
            sc := -1
            if result.Response() != nil {
                sc = result.Response().StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.UpdatePreparer(ctx, resourceGroupName, serverName, elasticPoolName, parameters)
    if err != nil {
    err = autorest.NewErrorWithError(err, "sql.ElasticPoolsClient", "Update", nil , "Failure preparing request")
    return
    }

            result, err = client.UpdateSender(req)
            if err != nil {
            err = autorest.NewErrorWithError(err, "sql.ElasticPoolsClient", "Update", result.Response(), "Failure sending request")
            return
            }

    return
    }

    // UpdatePreparer prepares the Update request.
    func (client ElasticPoolsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string, parameters ElasticPoolUpdate) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "elasticPoolName": autorest.Encode("path",elasticPoolName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "serverName": autorest.Encode("path",serverName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2017-10-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPatch(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/elasticPools/{elasticPoolName}",pathParameters),
    autorest.WithJSON(parameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // UpdateSender sends the Update request. The method will close the
    // http.Response Body if it receives an error.
    func (client ElasticPoolsClient) UpdateSender(req *http.Request) (future ElasticPoolsUpdateFuture, err error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            var resp *http.Response
            resp, err = autorest.SendWithSender(client, req, sd...)
            if err != nil {
            return
            }
            future.Future, err = azure.NewFutureFromResponse(resp)
            return
            }

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client ElasticPoolsClient) UpdateResponder(resp *http.Response) (result ElasticPool, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusAccepted),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

