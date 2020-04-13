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
    "github.com/Azure/go-autorest/autorest"
    "github.com/Azure/go-autorest/autorest/azure"
    "net/http"
    "context"
    "github.com/Azure/go-autorest/tracing"
    "github.com/Azure/go-autorest/autorest/validation"
)

// ExpressRouteCrossConnectionPeeringsClient is the network Client
type ExpressRouteCrossConnectionPeeringsClient struct {
    BaseClient
}
// NewExpressRouteCrossConnectionPeeringsClient creates an instance of the ExpressRouteCrossConnectionPeeringsClient
// client.
func NewExpressRouteCrossConnectionPeeringsClient(subscriptionID string) ExpressRouteCrossConnectionPeeringsClient {
    return NewExpressRouteCrossConnectionPeeringsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewExpressRouteCrossConnectionPeeringsClientWithBaseURI creates an instance of the
// ExpressRouteCrossConnectionPeeringsClient client using a custom endpoint.  Use this when interacting with an Azure
// cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
    func NewExpressRouteCrossConnectionPeeringsClientWithBaseURI(baseURI string, subscriptionID string) ExpressRouteCrossConnectionPeeringsClient {
        return ExpressRouteCrossConnectionPeeringsClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// CreateOrUpdate creates or updates a peering in the specified ExpressRouteCrossConnection.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // crossConnectionName - the name of the ExpressRouteCrossConnection.
        // peeringName - the name of the peering.
        // peeringParameters - parameters supplied to the create or update ExpressRouteCrossConnection peering
        // operation.
func (client ExpressRouteCrossConnectionPeeringsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, peeringParameters ExpressRouteCrossConnectionPeering) (result ExpressRouteCrossConnectionPeeringsCreateOrUpdateFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ExpressRouteCrossConnectionPeeringsClient.CreateOrUpdate")
        defer func() {
            sc := -1
            if result.Response() != nil {
                sc = result.Response().StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: peeringParameters,
             Constraints: []validation.Constraint{	{Target: "peeringParameters.ExpressRouteCrossConnectionPeeringProperties", Name: validation.Null, Rule: false ,
            Chain: []validation.Constraint{	{Target: "peeringParameters.ExpressRouteCrossConnectionPeeringProperties.PeerASN", Name: validation.Null, Rule: false ,
            Chain: []validation.Constraint{	{Target: "peeringParameters.ExpressRouteCrossConnectionPeeringProperties.PeerASN", Name: validation.InclusiveMaximum, Rule: int64(4294967295), Chain: nil },
            	{Target: "peeringParameters.ExpressRouteCrossConnectionPeeringProperties.PeerASN", Name: validation.InclusiveMinimum, Rule: int64(1), Chain: nil },
            }},
            }}}}}); err != nil {
            return result, validation.NewError("network.ExpressRouteCrossConnectionPeeringsClient", "CreateOrUpdate", err.Error())
            }

                req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, crossConnectionName, peeringName, peeringParameters)
    if err != nil {
    err = autorest.NewErrorWithError(err, "network.ExpressRouteCrossConnectionPeeringsClient", "CreateOrUpdate", nil , "Failure preparing request")
    return
    }

            result, err = client.CreateOrUpdateSender(req)
            if err != nil {
            err = autorest.NewErrorWithError(err, "network.ExpressRouteCrossConnectionPeeringsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
            return
            }

    return
    }

    // CreateOrUpdatePreparer prepares the CreateOrUpdate request.
    func (client ExpressRouteCrossConnectionPeeringsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, peeringParameters ExpressRouteCrossConnectionPeering) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "crossConnectionName": autorest.Encode("path",crossConnectionName),
            "peeringName": autorest.Encode("path",peeringName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2018-10-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

                    peeringParameters.Etag = nil
    preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPut(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}/peerings/{peeringName}",pathParameters),
    autorest.WithJSON(peeringParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
    // http.Response Body if it receives an error.
    func (client ExpressRouteCrossConnectionPeeringsClient) CreateOrUpdateSender(req *http.Request) (future ExpressRouteCrossConnectionPeeringsCreateOrUpdateFuture, err error) {
            var resp *http.Response
            resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
            if err != nil {
            return
            }
            future.Future, err = azure.NewFutureFromResponse(resp)
            return
            }

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ExpressRouteCrossConnectionPeeringsClient) CreateOrUpdateResponder(resp *http.Response) (result ExpressRouteCrossConnectionPeering, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusCreated),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// Delete deletes the specified peering from the ExpressRouteCrossConnection.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // crossConnectionName - the name of the ExpressRouteCrossConnection.
        // peeringName - the name of the peering.
func (client ExpressRouteCrossConnectionPeeringsClient) Delete(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string) (result ExpressRouteCrossConnectionPeeringsDeleteFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ExpressRouteCrossConnectionPeeringsClient.Delete")
        defer func() {
            sc := -1
            if result.Response() != nil {
                sc = result.Response().StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.DeletePreparer(ctx, resourceGroupName, crossConnectionName, peeringName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "network.ExpressRouteCrossConnectionPeeringsClient", "Delete", nil , "Failure preparing request")
    return
    }

            result, err = client.DeleteSender(req)
            if err != nil {
            err = autorest.NewErrorWithError(err, "network.ExpressRouteCrossConnectionPeeringsClient", "Delete", result.Response(), "Failure sending request")
            return
            }

    return
    }

    // DeletePreparer prepares the Delete request.
    func (client ExpressRouteCrossConnectionPeeringsClient) DeletePreparer(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "crossConnectionName": autorest.Encode("path",crossConnectionName),
            "peeringName": autorest.Encode("path",peeringName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2018-10-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsDelete(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}/peerings/{peeringName}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // DeleteSender sends the Delete request. The method will close the
    // http.Response Body if it receives an error.
    func (client ExpressRouteCrossConnectionPeeringsClient) DeleteSender(req *http.Request) (future ExpressRouteCrossConnectionPeeringsDeleteFuture, err error) {
            var resp *http.Response
            resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
            if err != nil {
            return
            }
            future.Future, err = azure.NewFutureFromResponse(resp)
            return
            }

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ExpressRouteCrossConnectionPeeringsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusAccepted,http.StatusNoContent),
    autorest.ByClosing())
    result.Response = resp
        return
    }

// Get gets the specified peering for the ExpressRouteCrossConnection.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // crossConnectionName - the name of the ExpressRouteCrossConnection.
        // peeringName - the name of the peering.
func (client ExpressRouteCrossConnectionPeeringsClient) Get(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string) (result ExpressRouteCrossConnectionPeering, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ExpressRouteCrossConnectionPeeringsClient.Get")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.GetPreparer(ctx, resourceGroupName, crossConnectionName, peeringName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "network.ExpressRouteCrossConnectionPeeringsClient", "Get", nil , "Failure preparing request")
    return
    }

            resp, err := client.GetSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "network.ExpressRouteCrossConnectionPeeringsClient", "Get", resp, "Failure sending request")
            return
            }

            result, err = client.GetResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "network.ExpressRouteCrossConnectionPeeringsClient", "Get", resp, "Failure responding to request")
            }

    return
    }

    // GetPreparer prepares the Get request.
    func (client ExpressRouteCrossConnectionPeeringsClient) GetPreparer(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "crossConnectionName": autorest.Encode("path",crossConnectionName),
            "peeringName": autorest.Encode("path",peeringName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2018-10-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}/peerings/{peeringName}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client ExpressRouteCrossConnectionPeeringsClient) GetSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ExpressRouteCrossConnectionPeeringsClient) GetResponder(resp *http.Response) (result ExpressRouteCrossConnectionPeering, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// List gets all peerings in a specified ExpressRouteCrossConnection.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // crossConnectionName - the name of the ExpressRouteCrossConnection.
func (client ExpressRouteCrossConnectionPeeringsClient) List(ctx context.Context, resourceGroupName string, crossConnectionName string) (result ExpressRouteCrossConnectionPeeringListPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ExpressRouteCrossConnectionPeeringsClient.List")
        defer func() {
            sc := -1
            if result.erccpl.Response.Response != nil {
                sc = result.erccpl.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
                result.fn = client.listNextResults
    req, err := client.ListPreparer(ctx, resourceGroupName, crossConnectionName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "network.ExpressRouteCrossConnectionPeeringsClient", "List", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListSender(req)
            if err != nil {
            result.erccpl.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "network.ExpressRouteCrossConnectionPeeringsClient", "List", resp, "Failure sending request")
            return
            }

            result.erccpl, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "network.ExpressRouteCrossConnectionPeeringsClient", "List", resp, "Failure responding to request")
            }

    return
    }

    // ListPreparer prepares the List request.
    func (client ExpressRouteCrossConnectionPeeringsClient) ListPreparer(ctx context.Context, resourceGroupName string, crossConnectionName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "crossConnectionName": autorest.Encode("path",crossConnectionName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2018-10-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}/peerings",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListSender sends the List request. The method will close the
    // http.Response Body if it receives an error.
    func (client ExpressRouteCrossConnectionPeeringsClient) ListSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ExpressRouteCrossConnectionPeeringsClient) ListResponder(resp *http.Response) (result ExpressRouteCrossConnectionPeeringList, err error) {
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
            func (client ExpressRouteCrossConnectionPeeringsClient) listNextResults(ctx context.Context, lastResults ExpressRouteCrossConnectionPeeringList) (result ExpressRouteCrossConnectionPeeringList, err error) {
            req, err := lastResults.expressRouteCrossConnectionPeeringListPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "network.ExpressRouteCrossConnectionPeeringsClient", "listNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "network.ExpressRouteCrossConnectionPeeringsClient", "listNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "network.ExpressRouteCrossConnectionPeeringsClient", "listNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

    // ListComplete enumerates all values, automatically crossing page boundaries as required.
    func (client ExpressRouteCrossConnectionPeeringsClient) ListComplete(ctx context.Context, resourceGroupName string, crossConnectionName string) (result ExpressRouteCrossConnectionPeeringListIterator, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/ExpressRouteCrossConnectionPeeringsClient.List")
            defer func() {
                sc := -1
                if result.Response().Response.Response != nil {
                    sc = result.page.Response().Response.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
     }
        result.page, err = client.List(ctx, resourceGroupName, crossConnectionName)
                return
        }

