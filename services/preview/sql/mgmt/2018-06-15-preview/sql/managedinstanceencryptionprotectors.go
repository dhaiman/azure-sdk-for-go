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

// ManagedInstanceEncryptionProtectorsClient is the the Azure SQL Database management API provides a RESTful set of web
// services that interact with Azure SQL Database services to manage your databases. The API enables you to create,
// retrieve, update, and delete databases.
type ManagedInstanceEncryptionProtectorsClient struct {
    BaseClient
}
// NewManagedInstanceEncryptionProtectorsClient creates an instance of the ManagedInstanceEncryptionProtectorsClient
// client.
func NewManagedInstanceEncryptionProtectorsClient(subscriptionID string) ManagedInstanceEncryptionProtectorsClient {
    return NewManagedInstanceEncryptionProtectorsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewManagedInstanceEncryptionProtectorsClientWithBaseURI creates an instance of the
// ManagedInstanceEncryptionProtectorsClient client.
    func NewManagedInstanceEncryptionProtectorsClientWithBaseURI(baseURI string, subscriptionID string) ManagedInstanceEncryptionProtectorsClient {
        return ManagedInstanceEncryptionProtectorsClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// CreateOrUpdate updates an existing encryption protector.
    // Parameters:
        // resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
        // from the Azure Resource Manager API or the portal.
        // managedInstanceName - the name of the managed instance.
        // parameters - the requested encryption protector resource state.
func (client ManagedInstanceEncryptionProtectorsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, managedInstanceName string, parameters ManagedInstanceEncryptionProtector) (result ManagedInstanceEncryptionProtectorsCreateOrUpdateFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ManagedInstanceEncryptionProtectorsClient.CreateOrUpdate")
        defer func() {
            sc := -1
            if result.Response() != nil {
                sc = result.Response().StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, managedInstanceName, parameters)
    if err != nil {
    err = autorest.NewErrorWithError(err, "sql.ManagedInstanceEncryptionProtectorsClient", "CreateOrUpdate", nil , "Failure preparing request")
    return
    }

            result, err = client.CreateOrUpdateSender(req)
            if err != nil {
            err = autorest.NewErrorWithError(err, "sql.ManagedInstanceEncryptionProtectorsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
            return
            }

    return
    }

    // CreateOrUpdatePreparer prepares the CreateOrUpdate request.
    func (client ManagedInstanceEncryptionProtectorsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, managedInstanceName string, parameters ManagedInstanceEncryptionProtector) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "encryptionProtectorName": autorest.Encode("path", "current"),
            "managedInstanceName": autorest.Encode("path",managedInstanceName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
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
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/encryptionProtector/{encryptionProtectorName}",pathParameters),
    autorest.WithJSON(parameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
    // http.Response Body if it receives an error.
    func (client ManagedInstanceEncryptionProtectorsClient) CreateOrUpdateSender(req *http.Request) (future ManagedInstanceEncryptionProtectorsCreateOrUpdateFuture, err error) {
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
func (client ManagedInstanceEncryptionProtectorsClient) CreateOrUpdateResponder(resp *http.Response) (result ManagedInstanceEncryptionProtector, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusAccepted),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// Get gets a managed instance encryption protector.
    // Parameters:
        // resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
        // from the Azure Resource Manager API or the portal.
        // managedInstanceName - the name of the managed instance.
func (client ManagedInstanceEncryptionProtectorsClient) Get(ctx context.Context, resourceGroupName string, managedInstanceName string) (result ManagedInstanceEncryptionProtector, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ManagedInstanceEncryptionProtectorsClient.Get")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.GetPreparer(ctx, resourceGroupName, managedInstanceName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "sql.ManagedInstanceEncryptionProtectorsClient", "Get", nil , "Failure preparing request")
    return
    }

            resp, err := client.GetSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "sql.ManagedInstanceEncryptionProtectorsClient", "Get", resp, "Failure sending request")
            return
            }

            result, err = client.GetResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "sql.ManagedInstanceEncryptionProtectorsClient", "Get", resp, "Failure responding to request")
            }

    return
    }

    // GetPreparer prepares the Get request.
    func (client ManagedInstanceEncryptionProtectorsClient) GetPreparer(ctx context.Context, resourceGroupName string, managedInstanceName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "encryptionProtectorName": autorest.Encode("path", "current"),
            "managedInstanceName": autorest.Encode("path",managedInstanceName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2017-10-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/encryptionProtector/{encryptionProtectorName}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client ManagedInstanceEncryptionProtectorsClient) GetSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ManagedInstanceEncryptionProtectorsClient) GetResponder(resp *http.Response) (result ManagedInstanceEncryptionProtector, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// ListByInstance gets a list of managed instance encryption protectors
    // Parameters:
        // resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
        // from the Azure Resource Manager API or the portal.
        // managedInstanceName - the name of the managed instance.
func (client ManagedInstanceEncryptionProtectorsClient) ListByInstance(ctx context.Context, resourceGroupName string, managedInstanceName string) (result ManagedInstanceEncryptionProtectorListResultPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ManagedInstanceEncryptionProtectorsClient.ListByInstance")
        defer func() {
            sc := -1
            if result.mieplr.Response.Response != nil {
                sc = result.mieplr.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
                result.fn = client.listByInstanceNextResults
    req, err := client.ListByInstancePreparer(ctx, resourceGroupName, managedInstanceName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "sql.ManagedInstanceEncryptionProtectorsClient", "ListByInstance", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListByInstanceSender(req)
            if err != nil {
            result.mieplr.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "sql.ManagedInstanceEncryptionProtectorsClient", "ListByInstance", resp, "Failure sending request")
            return
            }

            result.mieplr, err = client.ListByInstanceResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "sql.ManagedInstanceEncryptionProtectorsClient", "ListByInstance", resp, "Failure responding to request")
            }

    return
    }

    // ListByInstancePreparer prepares the ListByInstance request.
    func (client ManagedInstanceEncryptionProtectorsClient) ListByInstancePreparer(ctx context.Context, resourceGroupName string, managedInstanceName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "managedInstanceName": autorest.Encode("path",managedInstanceName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2017-10-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/encryptionProtector",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListByInstanceSender sends the ListByInstance request. The method will close the
    // http.Response Body if it receives an error.
    func (client ManagedInstanceEncryptionProtectorsClient) ListByInstanceSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// ListByInstanceResponder handles the response to the ListByInstance request. The method always
// closes the http.Response Body.
func (client ManagedInstanceEncryptionProtectorsClient) ListByInstanceResponder(resp *http.Response) (result ManagedInstanceEncryptionProtectorListResult, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

            // listByInstanceNextResults retrieves the next set of results, if any.
            func (client ManagedInstanceEncryptionProtectorsClient) listByInstanceNextResults(ctx context.Context, lastResults ManagedInstanceEncryptionProtectorListResult) (result ManagedInstanceEncryptionProtectorListResult, err error) {
            req, err := lastResults.managedInstanceEncryptionProtectorListResultPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "sql.ManagedInstanceEncryptionProtectorsClient", "listByInstanceNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListByInstanceSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "sql.ManagedInstanceEncryptionProtectorsClient", "listByInstanceNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListByInstanceResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "sql.ManagedInstanceEncryptionProtectorsClient", "listByInstanceNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

    // ListByInstanceComplete enumerates all values, automatically crossing page boundaries as required.
    func (client ManagedInstanceEncryptionProtectorsClient) ListByInstanceComplete(ctx context.Context, resourceGroupName string, managedInstanceName string) (result ManagedInstanceEncryptionProtectorListResultIterator, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/ManagedInstanceEncryptionProtectorsClient.ListByInstance")
            defer func() {
                sc := -1
                if result.Response().Response.Response != nil {
                    sc = result.page.Response().Response.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
     }
        result.page, err = client.ListByInstance(ctx, resourceGroupName, managedInstanceName)
                return
        }

// Revalidate revalidates an existing encryption protector.
    // Parameters:
        // resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
        // from the Azure Resource Manager API or the portal.
        // managedInstanceName - the name of the managed instance.
func (client ManagedInstanceEncryptionProtectorsClient) Revalidate(ctx context.Context, resourceGroupName string, managedInstanceName string) (result ManagedInstanceEncryptionProtectorsRevalidateFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ManagedInstanceEncryptionProtectorsClient.Revalidate")
        defer func() {
            sc := -1
            if result.Response() != nil {
                sc = result.Response().StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.RevalidatePreparer(ctx, resourceGroupName, managedInstanceName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "sql.ManagedInstanceEncryptionProtectorsClient", "Revalidate", nil , "Failure preparing request")
    return
    }

            result, err = client.RevalidateSender(req)
            if err != nil {
            err = autorest.NewErrorWithError(err, "sql.ManagedInstanceEncryptionProtectorsClient", "Revalidate", result.Response(), "Failure sending request")
            return
            }

    return
    }

    // RevalidatePreparer prepares the Revalidate request.
    func (client ManagedInstanceEncryptionProtectorsClient) RevalidatePreparer(ctx context.Context, resourceGroupName string, managedInstanceName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "encryptionProtectorName": autorest.Encode("path", "current"),
            "managedInstanceName": autorest.Encode("path",managedInstanceName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2017-10-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsPost(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/encryptionProtector/{encryptionProtectorName}/revalidate",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // RevalidateSender sends the Revalidate request. The method will close the
    // http.Response Body if it receives an error.
    func (client ManagedInstanceEncryptionProtectorsClient) RevalidateSender(req *http.Request) (future ManagedInstanceEncryptionProtectorsRevalidateFuture, err error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            var resp *http.Response
            resp, err = autorest.SendWithSender(client, req, sd...)
            if err != nil {
            return
            }
            future.Future, err = azure.NewFutureFromResponse(resp)
            return
            }

// RevalidateResponder handles the response to the Revalidate request. The method always
// closes the http.Response Body.
func (client ManagedInstanceEncryptionProtectorsClient) RevalidateResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusAccepted),
    autorest.ByClosing())
    result.Response = resp
        return
    }

