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
)

// ReplicationLinksClient is the the Azure SQL Database management API provides a RESTful set of web services that
// interact with Azure SQL Database services to manage your databases. The API enables you to create, retrieve, update,
// and delete databases.
type ReplicationLinksClient struct {
    BaseClient
}
// NewReplicationLinksClient creates an instance of the ReplicationLinksClient client.
func NewReplicationLinksClient(subscriptionID string) ReplicationLinksClient {
    return NewReplicationLinksClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewReplicationLinksClientWithBaseURI creates an instance of the ReplicationLinksClient client.
    func NewReplicationLinksClientWithBaseURI(baseURI string, subscriptionID string) ReplicationLinksClient {
        return ReplicationLinksClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// Delete deletes a database replication link. Cannot be done during failover.
    // Parameters:
        // resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
        // from the Azure Resource Manager API or the portal.
        // serverName - the name of the server.
        // databaseName - the name of the database that has the replication link to be dropped.
        // linkID - the ID of the replication link to be deleted.
func (client ReplicationLinksClient) Delete(ctx context.Context, resourceGroupName string, serverName string, databaseName string, linkID string) (result autorest.Response, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ReplicationLinksClient.Delete")
        defer func() {
            sc := -1
            if result.Response != nil {
                sc = result.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.DeletePreparer(ctx, resourceGroupName, serverName, databaseName, linkID)
    if err != nil {
    err = autorest.NewErrorWithError(err, "sql.ReplicationLinksClient", "Delete", nil , "Failure preparing request")
    return
    }

            resp, err := client.DeleteSender(req)
            if err != nil {
            result.Response = resp
            err = autorest.NewErrorWithError(err, "sql.ReplicationLinksClient", "Delete", resp, "Failure sending request")
            return
            }

            result, err = client.DeleteResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "sql.ReplicationLinksClient", "Delete", resp, "Failure responding to request")
            }

    return
    }

    // DeletePreparer prepares the Delete request.
    func (client ReplicationLinksClient) DeletePreparer(ctx context.Context, resourceGroupName string, serverName string, databaseName string, linkID string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "databaseName": autorest.Encode("path",databaseName),
            "linkId": autorest.Encode("path",linkID),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "serverName": autorest.Encode("path",serverName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2014-04-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsDelete(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/replicationLinks/{linkId}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // DeleteSender sends the Delete request. The method will close the
    // http.Response Body if it receives an error.
    func (client ReplicationLinksClient) DeleteSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ReplicationLinksClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusNoContent),
    autorest.ByClosing())
    result.Response = resp
        return
    }

// Failover sets which replica database is primary by failing over from the current primary replica database.
    // Parameters:
        // resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
        // from the Azure Resource Manager API or the portal.
        // serverName - the name of the server.
        // databaseName - the name of the database that has the replication link to be failed over.
        // linkID - the ID of the replication link to be failed over.
func (client ReplicationLinksClient) Failover(ctx context.Context, resourceGroupName string, serverName string, databaseName string, linkID string) (result ReplicationLinksFailoverFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ReplicationLinksClient.Failover")
        defer func() {
            sc := -1
            if result.Response() != nil {
                sc = result.Response().StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.FailoverPreparer(ctx, resourceGroupName, serverName, databaseName, linkID)
    if err != nil {
    err = autorest.NewErrorWithError(err, "sql.ReplicationLinksClient", "Failover", nil , "Failure preparing request")
    return
    }

            result, err = client.FailoverSender(req)
            if err != nil {
            err = autorest.NewErrorWithError(err, "sql.ReplicationLinksClient", "Failover", result.Response(), "Failure sending request")
            return
            }

    return
    }

    // FailoverPreparer prepares the Failover request.
    func (client ReplicationLinksClient) FailoverPreparer(ctx context.Context, resourceGroupName string, serverName string, databaseName string, linkID string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "databaseName": autorest.Encode("path",databaseName),
            "linkId": autorest.Encode("path",linkID),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "serverName": autorest.Encode("path",serverName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2014-04-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsPost(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/replicationLinks/{linkId}/failover",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // FailoverSender sends the Failover request. The method will close the
    // http.Response Body if it receives an error.
    func (client ReplicationLinksClient) FailoverSender(req *http.Request) (future ReplicationLinksFailoverFuture, err error) {
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
func (client ReplicationLinksClient) FailoverResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusAccepted,http.StatusNoContent),
    autorest.ByClosing())
    result.Response = resp
        return
    }

// FailoverAllowDataLoss sets which replica database is primary by failing over from the current primary replica
// database. This operation might result in data loss.
    // Parameters:
        // resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
        // from the Azure Resource Manager API or the portal.
        // serverName - the name of the server.
        // databaseName - the name of the database that has the replication link to be failed over.
        // linkID - the ID of the replication link to be failed over.
func (client ReplicationLinksClient) FailoverAllowDataLoss(ctx context.Context, resourceGroupName string, serverName string, databaseName string, linkID string) (result ReplicationLinksFailoverAllowDataLossFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ReplicationLinksClient.FailoverAllowDataLoss")
        defer func() {
            sc := -1
            if result.Response() != nil {
                sc = result.Response().StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.FailoverAllowDataLossPreparer(ctx, resourceGroupName, serverName, databaseName, linkID)
    if err != nil {
    err = autorest.NewErrorWithError(err, "sql.ReplicationLinksClient", "FailoverAllowDataLoss", nil , "Failure preparing request")
    return
    }

            result, err = client.FailoverAllowDataLossSender(req)
            if err != nil {
            err = autorest.NewErrorWithError(err, "sql.ReplicationLinksClient", "FailoverAllowDataLoss", result.Response(), "Failure sending request")
            return
            }

    return
    }

    // FailoverAllowDataLossPreparer prepares the FailoverAllowDataLoss request.
    func (client ReplicationLinksClient) FailoverAllowDataLossPreparer(ctx context.Context, resourceGroupName string, serverName string, databaseName string, linkID string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "databaseName": autorest.Encode("path",databaseName),
            "linkId": autorest.Encode("path",linkID),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "serverName": autorest.Encode("path",serverName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2014-04-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsPost(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/replicationLinks/{linkId}/forceFailoverAllowDataLoss",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // FailoverAllowDataLossSender sends the FailoverAllowDataLoss request. The method will close the
    // http.Response Body if it receives an error.
    func (client ReplicationLinksClient) FailoverAllowDataLossSender(req *http.Request) (future ReplicationLinksFailoverAllowDataLossFuture, err error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            var resp *http.Response
            resp, err = autorest.SendWithSender(client, req, sd...)
            if err != nil {
            return
            }
            future.Future, err = azure.NewFutureFromResponse(resp)
            return
            }

// FailoverAllowDataLossResponder handles the response to the FailoverAllowDataLoss request. The method always
// closes the http.Response Body.
func (client ReplicationLinksClient) FailoverAllowDataLossResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusAccepted,http.StatusNoContent),
    autorest.ByClosing())
    result.Response = resp
        return
    }

// Get gets a database replication link.
    // Parameters:
        // resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
        // from the Azure Resource Manager API or the portal.
        // serverName - the name of the server.
        // databaseName - the name of the database to get the link for.
        // linkID - the replication link ID to be retrieved.
func (client ReplicationLinksClient) Get(ctx context.Context, resourceGroupName string, serverName string, databaseName string, linkID string) (result ReplicationLink, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ReplicationLinksClient.Get")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.GetPreparer(ctx, resourceGroupName, serverName, databaseName, linkID)
    if err != nil {
    err = autorest.NewErrorWithError(err, "sql.ReplicationLinksClient", "Get", nil , "Failure preparing request")
    return
    }

            resp, err := client.GetSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "sql.ReplicationLinksClient", "Get", resp, "Failure sending request")
            return
            }

            result, err = client.GetResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "sql.ReplicationLinksClient", "Get", resp, "Failure responding to request")
            }

    return
    }

    // GetPreparer prepares the Get request.
    func (client ReplicationLinksClient) GetPreparer(ctx context.Context, resourceGroupName string, serverName string, databaseName string, linkID string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "databaseName": autorest.Encode("path",databaseName),
            "linkId": autorest.Encode("path",linkID),
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
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/replicationLinks/{linkId}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client ReplicationLinksClient) GetSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ReplicationLinksClient) GetResponder(resp *http.Response) (result ReplicationLink, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// ListByDatabase lists a database's replication links.
    // Parameters:
        // resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
        // from the Azure Resource Manager API or the portal.
        // serverName - the name of the server.
        // databaseName - the name of the database to retrieve links for.
func (client ReplicationLinksClient) ListByDatabase(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result ReplicationLinkListResult, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ReplicationLinksClient.ListByDatabase")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.ListByDatabasePreparer(ctx, resourceGroupName, serverName, databaseName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "sql.ReplicationLinksClient", "ListByDatabase", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListByDatabaseSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "sql.ReplicationLinksClient", "ListByDatabase", resp, "Failure sending request")
            return
            }

            result, err = client.ListByDatabaseResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "sql.ReplicationLinksClient", "ListByDatabase", resp, "Failure responding to request")
            }

    return
    }

    // ListByDatabasePreparer prepares the ListByDatabase request.
    func (client ReplicationLinksClient) ListByDatabasePreparer(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "databaseName": autorest.Encode("path",databaseName),
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
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/replicationLinks",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListByDatabaseSender sends the ListByDatabase request. The method will close the
    // http.Response Body if it receives an error.
    func (client ReplicationLinksClient) ListByDatabaseSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// ListByDatabaseResponder handles the response to the ListByDatabase request. The method always
// closes the http.Response Body.
func (client ReplicationLinksClient) ListByDatabaseResponder(resp *http.Response) (result ReplicationLinkListResult, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

