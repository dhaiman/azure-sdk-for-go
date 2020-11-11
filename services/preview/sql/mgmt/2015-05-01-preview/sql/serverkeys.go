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

// ServerKeysClient is the the Azure SQL Database management API provides a RESTful set of web services that interact
// with Azure SQL Database services to manage your databases. The API enables you to create, retrieve, update, and
// delete databases.
type ServerKeysClient struct {
    BaseClient
}
// NewServerKeysClient creates an instance of the ServerKeysClient client.
func NewServerKeysClient(subscriptionID string) ServerKeysClient {
    return NewServerKeysClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewServerKeysClientWithBaseURI creates an instance of the ServerKeysClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
    func NewServerKeysClientWithBaseURI(baseURI string, subscriptionID string) ServerKeysClient {
        return ServerKeysClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// CreateOrUpdate creates or updates a server key.
    // Parameters:
        // resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
        // from the Azure Resource Manager API or the portal.
        // serverName - the name of the server.
        // keyName - the name of the server key to be operated on (updated or created). The key name is required to be
        // in the format of 'vault_key_version'. For example, if the keyId is
        // https://YourVaultName.vault.azure.net/keys/YourKeyName/01234567890123456789012345678901, then the server key
        // name should be formatted as: YourVaultName_YourKeyName_01234567890123456789012345678901
        // parameters - the requested server key resource state.
func (client ServerKeysClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, keyName string, parameters ServerKey) (result ServerKeysCreateOrUpdateFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ServerKeysClient.CreateOrUpdate")
        defer func() {
            sc := -1
        if result.Response() != nil {
        sc = result.Response().StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, serverName, keyName, parameters)
    if err != nil {
    err = autorest.NewErrorWithError(err, "sql.ServerKeysClient", "CreateOrUpdate", nil , "Failure preparing request")
    return
    }

        result, err = client.CreateOrUpdateSender(req)
        if err != nil {
        err = autorest.NewErrorWithError(err, "sql.ServerKeysClient", "CreateOrUpdate", result.Response(), "Failure sending request")
        return
        }

    return
}

    // CreateOrUpdatePreparer prepares the CreateOrUpdate request.
    func (client ServerKeysClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, serverName string, keyName string, parameters ServerKey) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "keyName": autorest.Encode("path",keyName),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "serverName": autorest.Encode("path",serverName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2015-05-01-preview"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

            parameters.Location = nil
    preparer := autorest.CreatePreparer(
autorest.AsContentType("application/json; charset=utf-8"),
autorest.AsPut(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/keys/{keyName}",pathParameters),
autorest.WithJSON(parameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
    // http.Response Body if it receives an error.
    func (client ServerKeysClient) CreateOrUpdateSender(req *http.Request) (future ServerKeysCreateOrUpdateFuture, err error) {
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
    func (client ServerKeysClient) CreateOrUpdateResponder(resp *http.Response) (result ServerKey, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusCreated,http.StatusAccepted),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// Delete deletes the server key with the given name.
    // Parameters:
        // resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
        // from the Azure Resource Manager API or the portal.
        // serverName - the name of the server.
        // keyName - the name of the server key to be deleted.
func (client ServerKeysClient) Delete(ctx context.Context, resourceGroupName string, serverName string, keyName string) (result ServerKeysDeleteFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ServerKeysClient.Delete")
        defer func() {
            sc := -1
        if result.Response() != nil {
        sc = result.Response().StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.DeletePreparer(ctx, resourceGroupName, serverName, keyName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "sql.ServerKeysClient", "Delete", nil , "Failure preparing request")
    return
    }

        result, err = client.DeleteSender(req)
        if err != nil {
        err = autorest.NewErrorWithError(err, "sql.ServerKeysClient", "Delete", result.Response(), "Failure sending request")
        return
        }

    return
}

    // DeletePreparer prepares the Delete request.
    func (client ServerKeysClient) DeletePreparer(ctx context.Context, resourceGroupName string, serverName string, keyName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "keyName": autorest.Encode("path",keyName),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "serverName": autorest.Encode("path",serverName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2015-05-01-preview"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsDelete(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/keys/{keyName}",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // DeleteSender sends the Delete request. The method will close the
    // http.Response Body if it receives an error.
    func (client ServerKeysClient) DeleteSender(req *http.Request) (future ServerKeysDeleteFuture, err error) {
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
    func (client ServerKeysClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusAccepted,http.StatusNoContent),
            autorest.ByClosing())
            result.Response = resp
            return
    }

// Get gets a server key.
    // Parameters:
        // resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
        // from the Azure Resource Manager API or the portal.
        // serverName - the name of the server.
        // keyName - the name of the server key to be retrieved.
func (client ServerKeysClient) Get(ctx context.Context, resourceGroupName string, serverName string, keyName string) (result ServerKey, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ServerKeysClient.Get")
        defer func() {
            sc := -1
        if result.Response.Response != nil {
        sc = result.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.GetPreparer(ctx, resourceGroupName, serverName, keyName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "sql.ServerKeysClient", "Get", nil , "Failure preparing request")
    return
    }

        resp, err := client.GetSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "sql.ServerKeysClient", "Get", resp, "Failure sending request")
        return
        }

        result, err = client.GetResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "sql.ServerKeysClient", "Get", resp, "Failure responding to request")
        }

    return
}

    // GetPreparer prepares the Get request.
    func (client ServerKeysClient) GetPreparer(ctx context.Context, resourceGroupName string, serverName string, keyName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "keyName": autorest.Encode("path",keyName),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "serverName": autorest.Encode("path",serverName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2015-05-01-preview"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/keys/{keyName}",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client ServerKeysClient) GetSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

    // GetResponder handles the response to the Get request. The method always
    // closes the http.Response Body.
    func (client ServerKeysClient) GetResponder(resp *http.Response) (result ServerKey, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// ListByServer gets a list of server keys.
    // Parameters:
        // resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
        // from the Azure Resource Manager API or the portal.
        // serverName - the name of the server.
func (client ServerKeysClient) ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result ServerKeyListResultPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ServerKeysClient.ListByServer")
        defer func() {
            sc := -1
        if result.sklr.Response.Response != nil {
        sc = result.sklr.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        result.fn = client.listByServerNextResults
    req, err := client.ListByServerPreparer(ctx, resourceGroupName, serverName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "sql.ServerKeysClient", "ListByServer", nil , "Failure preparing request")
    return
    }

        resp, err := client.ListByServerSender(req)
        if err != nil {
        result.sklr.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "sql.ServerKeysClient", "ListByServer", resp, "Failure sending request")
        return
        }

        result.sklr, err = client.ListByServerResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "sql.ServerKeysClient", "ListByServer", resp, "Failure responding to request")
        }
            if result.sklr.hasNextLink() && result.sklr.IsEmpty() {
            err = result.NextWithContext(ctx)
            }

    return
}

    // ListByServerPreparer prepares the ListByServer request.
    func (client ServerKeysClient) ListByServerPreparer(ctx context.Context, resourceGroupName string, serverName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "serverName": autorest.Encode("path",serverName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2015-05-01-preview"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/keys",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListByServerSender sends the ListByServer request. The method will close the
    // http.Response Body if it receives an error.
    func (client ServerKeysClient) ListByServerSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

    // ListByServerResponder handles the response to the ListByServer request. The method always
    // closes the http.Response Body.
    func (client ServerKeysClient) ListByServerResponder(resp *http.Response) (result ServerKeyListResult, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

            // listByServerNextResults retrieves the next set of results, if any.
            func (client ServerKeysClient) listByServerNextResults(ctx context.Context, lastResults ServerKeyListResult) (result ServerKeyListResult, err error) {
            req, err := lastResults.serverKeyListResultPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "sql.ServerKeysClient", "listByServerNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListByServerSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "sql.ServerKeysClient", "listByServerNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListByServerResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "sql.ServerKeysClient", "listByServerNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

            // ListByServerComplete enumerates all values, automatically crossing page boundaries as required.
            func (client ServerKeysClient) ListByServerComplete(ctx context.Context, resourceGroupName string, serverName string) (result ServerKeyListResultIterator, err error) {
            if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/ServerKeysClient.ListByServer")
            defer func() {
            sc := -1
            if result.Response().Response.Response != nil {
            sc = result.page.Response().Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
            }()
            }
                    result.page, err = client.ListByServer(ctx, resourceGroupName, serverName)
                            return
            }

