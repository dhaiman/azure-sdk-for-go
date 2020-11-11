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
    "github.com/Azure/go-autorest/autorest"
    "github.com/Azure/go-autorest/autorest/azure"
    "net/http"
    "context"
    "github.com/Azure/go-autorest/tracing"
    "github.com/Azure/go-autorest/autorest/validation"
)

// WorkspaceManagedSQLServerUsagesClient is the azure Synapse Analytics Management Client
type WorkspaceManagedSQLServerUsagesClient struct {
    BaseClient
}
// NewWorkspaceManagedSQLServerUsagesClient creates an instance of the WorkspaceManagedSQLServerUsagesClient client.
func NewWorkspaceManagedSQLServerUsagesClient(subscriptionID string) WorkspaceManagedSQLServerUsagesClient {
    return NewWorkspaceManagedSQLServerUsagesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWorkspaceManagedSQLServerUsagesClientWithBaseURI creates an instance of the WorkspaceManagedSQLServerUsagesClient
// client using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI
// (sovereign clouds, Azure stack).
    func NewWorkspaceManagedSQLServerUsagesClientWithBaseURI(baseURI string, subscriptionID string) WorkspaceManagedSQLServerUsagesClient {
        return WorkspaceManagedSQLServerUsagesClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// List get list of server usages metric for workspace managed sql server.
    // Parameters:
        // resourceGroupName - the name of the resource group. The name is case insensitive.
        // workspaceName - the name of the workspace
func (client WorkspaceManagedSQLServerUsagesClient) List(ctx context.Context, resourceGroupName string, workspaceName string) (result ServerUsageListResultPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/WorkspaceManagedSQLServerUsagesClient.List")
        defer func() {
            sc := -1
        if result.sulr.Response.Response != nil {
        sc = result.sulr.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        if err := validation.Validate([]validation.Validation{
        { TargetValue: client.SubscriptionID,
         Constraints: []validation.Constraint{	{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil }}},
        { TargetValue: resourceGroupName,
         Constraints: []validation.Constraint{	{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil },
        	{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil },
        	{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil }}}}); err != nil {
        return result, validation.NewError("synapse.WorkspaceManagedSQLServerUsagesClient", "List", err.Error())
        }

            result.fn = client.listNextResults
    req, err := client.ListPreparer(ctx, resourceGroupName, workspaceName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "synapse.WorkspaceManagedSQLServerUsagesClient", "List", nil , "Failure preparing request")
    return
    }

        resp, err := client.ListSender(req)
        if err != nil {
        result.sulr.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "synapse.WorkspaceManagedSQLServerUsagesClient", "List", resp, "Failure sending request")
        return
        }

        result.sulr, err = client.ListResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "synapse.WorkspaceManagedSQLServerUsagesClient", "List", resp, "Failure responding to request")
        }
            if result.sulr.hasNextLink() && result.sulr.IsEmpty() {
            err = result.NextWithContext(ctx)
            }

    return
}

    // ListPreparer prepares the List request.
    func (client WorkspaceManagedSQLServerUsagesClient) ListPreparer(ctx context.Context, resourceGroupName string, workspaceName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        "workspaceName": autorest.Encode("path",workspaceName),
        }

            const APIVersion = "2019-06-01-preview"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlUsages",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListSender sends the List request. The method will close the
    // http.Response Body if it receives an error.
    func (client WorkspaceManagedSQLServerUsagesClient) ListSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

    // ListResponder handles the response to the List request. The method always
    // closes the http.Response Body.
    func (client WorkspaceManagedSQLServerUsagesClient) ListResponder(resp *http.Response) (result ServerUsageListResult, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

            // listNextResults retrieves the next set of results, if any.
            func (client WorkspaceManagedSQLServerUsagesClient) listNextResults(ctx context.Context, lastResults ServerUsageListResult) (result ServerUsageListResult, err error) {
            req, err := lastResults.serverUsageListResultPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "synapse.WorkspaceManagedSQLServerUsagesClient", "listNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "synapse.WorkspaceManagedSQLServerUsagesClient", "listNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "synapse.WorkspaceManagedSQLServerUsagesClient", "listNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

            // ListComplete enumerates all values, automatically crossing page boundaries as required.
            func (client WorkspaceManagedSQLServerUsagesClient) ListComplete(ctx context.Context, resourceGroupName string, workspaceName string) (result ServerUsageListResultIterator, err error) {
            if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/WorkspaceManagedSQLServerUsagesClient.List")
            defer func() {
            sc := -1
            if result.Response().Response.Response != nil {
            sc = result.page.Response().Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
            }()
            }
                    result.page, err = client.List(ctx, resourceGroupName, workspaceName)
                            return
            }

