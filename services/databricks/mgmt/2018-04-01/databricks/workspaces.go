package databricks

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

// WorkspacesClient is the ARM Databricks
type WorkspacesClient struct {
    BaseClient
}
// NewWorkspacesClient creates an instance of the WorkspacesClient client.
func NewWorkspacesClient(subscriptionID string) WorkspacesClient {
    return NewWorkspacesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWorkspacesClientWithBaseURI creates an instance of the WorkspacesClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
    func NewWorkspacesClientWithBaseURI(baseURI string, subscriptionID string) WorkspacesClient {
        return WorkspacesClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// CreateOrUpdate creates a new workspace.
    // Parameters:
        // parameters - parameters supplied to the create or update a workspace.
        // resourceGroupName - the name of the resource group. The name is case insensitive.
        // workspaceName - the name of the workspace.
func (client WorkspacesClient) CreateOrUpdate(ctx context.Context, parameters Workspace, resourceGroupName string, workspaceName string) (result WorkspacesCreateOrUpdateFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/WorkspacesClient.CreateOrUpdate")
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
         Constraints: []validation.Constraint{	{Target: "parameters.WorkspaceProperties", Name: validation.Null, Rule: true ,
        Chain: []validation.Constraint{	{Target: "parameters.WorkspaceProperties.ManagedResourceGroupID", Name: validation.Null, Rule: true, Chain: nil },
        	{Target: "parameters.WorkspaceProperties.Parameters", Name: validation.Null, Rule: false ,
        Chain: []validation.Constraint{	{Target: "parameters.WorkspaceProperties.Parameters.AmlWorkspaceID", Name: validation.Null, Rule: false ,
        Chain: []validation.Constraint{	{Target: "parameters.WorkspaceProperties.Parameters.AmlWorkspaceID.Value", Name: validation.Null, Rule: true, Chain: nil },
        }},
        	{Target: "parameters.WorkspaceProperties.Parameters.CustomVirtualNetworkID", Name: validation.Null, Rule: false ,
        Chain: []validation.Constraint{	{Target: "parameters.WorkspaceProperties.Parameters.CustomVirtualNetworkID.Value", Name: validation.Null, Rule: true, Chain: nil },
        }},
        	{Target: "parameters.WorkspaceProperties.Parameters.CustomPublicSubnetName", Name: validation.Null, Rule: false ,
        Chain: []validation.Constraint{	{Target: "parameters.WorkspaceProperties.Parameters.CustomPublicSubnetName.Value", Name: validation.Null, Rule: true, Chain: nil },
        }},
        	{Target: "parameters.WorkspaceProperties.Parameters.CustomPrivateSubnetName", Name: validation.Null, Rule: false ,
        Chain: []validation.Constraint{	{Target: "parameters.WorkspaceProperties.Parameters.CustomPrivateSubnetName.Value", Name: validation.Null, Rule: true, Chain: nil },
        }},
        	{Target: "parameters.WorkspaceProperties.Parameters.EnableNoPublicIP", Name: validation.Null, Rule: false ,
        Chain: []validation.Constraint{	{Target: "parameters.WorkspaceProperties.Parameters.EnableNoPublicIP.Value", Name: validation.Null, Rule: true, Chain: nil },
        }},
        	{Target: "parameters.WorkspaceProperties.Parameters.PrepareEncryption", Name: validation.Null, Rule: false ,
        Chain: []validation.Constraint{	{Target: "parameters.WorkspaceProperties.Parameters.PrepareEncryption.Value", Name: validation.Null, Rule: true, Chain: nil },
        }},
        	{Target: "parameters.WorkspaceProperties.Parameters.RequireInfrastructureEncryption", Name: validation.Null, Rule: false ,
        Chain: []validation.Constraint{	{Target: "parameters.WorkspaceProperties.Parameters.RequireInfrastructureEncryption.Value", Name: validation.Null, Rule: true, Chain: nil },
        }},
        }},
        }},
        	{Target: "parameters.Sku", Name: validation.Null, Rule: false ,
        Chain: []validation.Constraint{	{Target: "parameters.Sku.Name", Name: validation.Null, Rule: true, Chain: nil },
        }}}},
        { TargetValue: resourceGroupName,
         Constraints: []validation.Constraint{	{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil },
        	{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil },
        	{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil }}},
        { TargetValue: workspaceName,
         Constraints: []validation.Constraint{	{Target: "workspaceName", Name: validation.MaxLength, Rule: 64, Chain: nil },
        	{Target: "workspaceName", Name: validation.MinLength, Rule: 3, Chain: nil }}}}); err != nil {
        return result, validation.NewError("databricks.WorkspacesClient", "CreateOrUpdate", err.Error())
        }

        req, err := client.CreateOrUpdatePreparer(ctx, parameters, resourceGroupName, workspaceName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "databricks.WorkspacesClient", "CreateOrUpdate", nil , "Failure preparing request")
    return
    }

        result, err = client.CreateOrUpdateSender(req)
        if err != nil {
        err = autorest.NewErrorWithError(err, "databricks.WorkspacesClient", "CreateOrUpdate", result.Response(), "Failure sending request")
        return
        }

    return
}

    // CreateOrUpdatePreparer prepares the CreateOrUpdate request.
    func (client WorkspacesClient) CreateOrUpdatePreparer(ctx context.Context, parameters Workspace, resourceGroupName string, workspaceName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        "workspaceName": autorest.Encode("path",workspaceName),
        }

            const APIVersion = "2018-04-01"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsContentType("application/json; charset=utf-8"),
autorest.AsPut(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Databricks/workspaces/{workspaceName}",pathParameters),
autorest.WithJSON(parameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
    // http.Response Body if it receives an error.
    func (client WorkspacesClient) CreateOrUpdateSender(req *http.Request) (future WorkspacesCreateOrUpdateFuture, err error) {
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
    func (client WorkspacesClient) CreateOrUpdateResponder(resp *http.Response) (result Workspace, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusCreated),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// Delete deletes the workspace.
    // Parameters:
        // resourceGroupName - the name of the resource group. The name is case insensitive.
        // workspaceName - the name of the workspace.
func (client WorkspacesClient) Delete(ctx context.Context, resourceGroupName string, workspaceName string) (result WorkspacesDeleteFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/WorkspacesClient.Delete")
        defer func() {
            sc := -1
        if result.Response() != nil {
        sc = result.Response().StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        if err := validation.Validate([]validation.Validation{
        { TargetValue: resourceGroupName,
         Constraints: []validation.Constraint{	{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil },
        	{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil },
        	{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil }}},
        { TargetValue: workspaceName,
         Constraints: []validation.Constraint{	{Target: "workspaceName", Name: validation.MaxLength, Rule: 64, Chain: nil },
        	{Target: "workspaceName", Name: validation.MinLength, Rule: 3, Chain: nil }}}}); err != nil {
        return result, validation.NewError("databricks.WorkspacesClient", "Delete", err.Error())
        }

        req, err := client.DeletePreparer(ctx, resourceGroupName, workspaceName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "databricks.WorkspacesClient", "Delete", nil , "Failure preparing request")
    return
    }

        result, err = client.DeleteSender(req)
        if err != nil {
        err = autorest.NewErrorWithError(err, "databricks.WorkspacesClient", "Delete", result.Response(), "Failure sending request")
        return
        }

    return
}

    // DeletePreparer prepares the Delete request.
    func (client WorkspacesClient) DeletePreparer(ctx context.Context, resourceGroupName string, workspaceName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        "workspaceName": autorest.Encode("path",workspaceName),
        }

            const APIVersion = "2018-04-01"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsDelete(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Databricks/workspaces/{workspaceName}",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // DeleteSender sends the Delete request. The method will close the
    // http.Response Body if it receives an error.
    func (client WorkspacesClient) DeleteSender(req *http.Request) (future WorkspacesDeleteFuture, err error) {
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
    func (client WorkspacesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusAccepted,http.StatusNoContent),
            autorest.ByClosing())
            result.Response = resp
            return
    }

// Get gets the workspace.
    // Parameters:
        // resourceGroupName - the name of the resource group. The name is case insensitive.
        // workspaceName - the name of the workspace.
func (client WorkspacesClient) Get(ctx context.Context, resourceGroupName string, workspaceName string) (result Workspace, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/WorkspacesClient.Get")
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
        { TargetValue: workspaceName,
         Constraints: []validation.Constraint{	{Target: "workspaceName", Name: validation.MaxLength, Rule: 64, Chain: nil },
        	{Target: "workspaceName", Name: validation.MinLength, Rule: 3, Chain: nil }}}}); err != nil {
        return result, validation.NewError("databricks.WorkspacesClient", "Get", err.Error())
        }

        req, err := client.GetPreparer(ctx, resourceGroupName, workspaceName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "databricks.WorkspacesClient", "Get", nil , "Failure preparing request")
    return
    }

        resp, err := client.GetSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "databricks.WorkspacesClient", "Get", resp, "Failure sending request")
        return
        }

        result, err = client.GetResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "databricks.WorkspacesClient", "Get", resp, "Failure responding to request")
        }

    return
}

    // GetPreparer prepares the Get request.
    func (client WorkspacesClient) GetPreparer(ctx context.Context, resourceGroupName string, workspaceName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        "workspaceName": autorest.Encode("path",workspaceName),
        }

            const APIVersion = "2018-04-01"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Databricks/workspaces/{workspaceName}",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client WorkspacesClient) GetSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

    // GetResponder handles the response to the Get request. The method always
    // closes the http.Response Body.
    func (client WorkspacesClient) GetResponder(resp *http.Response) (result Workspace, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// ListByResourceGroup gets all the workspaces within a resource group.
    // Parameters:
        // resourceGroupName - the name of the resource group. The name is case insensitive.
func (client WorkspacesClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (result WorkspaceListResultPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/WorkspacesClient.ListByResourceGroup")
        defer func() {
            sc := -1
        if result.wlr.Response.Response != nil {
        sc = result.wlr.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        if err := validation.Validate([]validation.Validation{
        { TargetValue: resourceGroupName,
         Constraints: []validation.Constraint{	{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil },
        	{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil },
        	{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil }}}}); err != nil {
        return result, validation.NewError("databricks.WorkspacesClient", "ListByResourceGroup", err.Error())
        }

            result.fn = client.listByResourceGroupNextResults
    req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "databricks.WorkspacesClient", "ListByResourceGroup", nil , "Failure preparing request")
    return
    }

        resp, err := client.ListByResourceGroupSender(req)
        if err != nil {
        result.wlr.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "databricks.WorkspacesClient", "ListByResourceGroup", resp, "Failure sending request")
        return
        }

        result.wlr, err = client.ListByResourceGroupResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "databricks.WorkspacesClient", "ListByResourceGroup", resp, "Failure responding to request")
        }
            if result.wlr.hasNextLink() && result.wlr.IsEmpty() {
            err = result.NextWithContext(ctx)
            }

    return
}

    // ListByResourceGroupPreparer prepares the ListByResourceGroup request.
    func (client WorkspacesClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2018-04-01"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Databricks/workspaces",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
    // http.Response Body if it receives an error.
    func (client WorkspacesClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

    // ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
    // closes the http.Response Body.
    func (client WorkspacesClient) ListByResourceGroupResponder(resp *http.Response) (result WorkspaceListResult, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

            // listByResourceGroupNextResults retrieves the next set of results, if any.
            func (client WorkspacesClient) listByResourceGroupNextResults(ctx context.Context, lastResults WorkspaceListResult) (result WorkspaceListResult, err error) {
            req, err := lastResults.workspaceListResultPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "databricks.WorkspacesClient", "listByResourceGroupNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListByResourceGroupSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "databricks.WorkspacesClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListByResourceGroupResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "databricks.WorkspacesClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

            // ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
            func (client WorkspacesClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result WorkspaceListResultIterator, err error) {
            if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/WorkspacesClient.ListByResourceGroup")
            defer func() {
            sc := -1
            if result.Response().Response.Response != nil {
            sc = result.page.Response().Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
            }()
            }
                    result.page, err = client.ListByResourceGroup(ctx, resourceGroupName)
                            return
            }

// ListBySubscription gets all the workspaces within a subscription.
func (client WorkspacesClient) ListBySubscription(ctx context.Context) (result WorkspaceListResultPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/WorkspacesClient.ListBySubscription")
        defer func() {
            sc := -1
        if result.wlr.Response.Response != nil {
        sc = result.wlr.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        result.fn = client.listBySubscriptionNextResults
    req, err := client.ListBySubscriptionPreparer(ctx)
    if err != nil {
    err = autorest.NewErrorWithError(err, "databricks.WorkspacesClient", "ListBySubscription", nil , "Failure preparing request")
    return
    }

        resp, err := client.ListBySubscriptionSender(req)
        if err != nil {
        result.wlr.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "databricks.WorkspacesClient", "ListBySubscription", resp, "Failure sending request")
        return
        }

        result.wlr, err = client.ListBySubscriptionResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "databricks.WorkspacesClient", "ListBySubscription", resp, "Failure responding to request")
        }
            if result.wlr.hasNextLink() && result.wlr.IsEmpty() {
            err = result.NextWithContext(ctx)
            }

    return
}

    // ListBySubscriptionPreparer prepares the ListBySubscription request.
    func (client WorkspacesClient) ListBySubscriptionPreparer(ctx context.Context) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2018-04-01"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Databricks/workspaces",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListBySubscriptionSender sends the ListBySubscription request. The method will close the
    // http.Response Body if it receives an error.
    func (client WorkspacesClient) ListBySubscriptionSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

    // ListBySubscriptionResponder handles the response to the ListBySubscription request. The method always
    // closes the http.Response Body.
    func (client WorkspacesClient) ListBySubscriptionResponder(resp *http.Response) (result WorkspaceListResult, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

            // listBySubscriptionNextResults retrieves the next set of results, if any.
            func (client WorkspacesClient) listBySubscriptionNextResults(ctx context.Context, lastResults WorkspaceListResult) (result WorkspaceListResult, err error) {
            req, err := lastResults.workspaceListResultPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "databricks.WorkspacesClient", "listBySubscriptionNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListBySubscriptionSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "databricks.WorkspacesClient", "listBySubscriptionNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListBySubscriptionResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "databricks.WorkspacesClient", "listBySubscriptionNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

            // ListBySubscriptionComplete enumerates all values, automatically crossing page boundaries as required.
            func (client WorkspacesClient) ListBySubscriptionComplete(ctx context.Context) (result WorkspaceListResultIterator, err error) {
            if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/WorkspacesClient.ListBySubscription")
            defer func() {
            sc := -1
            if result.Response().Response.Response != nil {
            sc = result.page.Response().Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
            }()
            }
                    result.page, err = client.ListBySubscription(ctx)
                            return
            }

// Update updates a workspace.
    // Parameters:
        // parameters - the update to the workspace.
        // resourceGroupName - the name of the resource group. The name is case insensitive.
        // workspaceName - the name of the workspace.
func (client WorkspacesClient) Update(ctx context.Context, parameters WorkspaceUpdate, resourceGroupName string, workspaceName string) (result WorkspacesUpdateFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/WorkspacesClient.Update")
        defer func() {
            sc := -1
        if result.Response() != nil {
        sc = result.Response().StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        if err := validation.Validate([]validation.Validation{
        { TargetValue: resourceGroupName,
         Constraints: []validation.Constraint{	{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil },
        	{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil },
        	{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil }}},
        { TargetValue: workspaceName,
         Constraints: []validation.Constraint{	{Target: "workspaceName", Name: validation.MaxLength, Rule: 64, Chain: nil },
        	{Target: "workspaceName", Name: validation.MinLength, Rule: 3, Chain: nil }}}}); err != nil {
        return result, validation.NewError("databricks.WorkspacesClient", "Update", err.Error())
        }

        req, err := client.UpdatePreparer(ctx, parameters, resourceGroupName, workspaceName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "databricks.WorkspacesClient", "Update", nil , "Failure preparing request")
    return
    }

        result, err = client.UpdateSender(req)
        if err != nil {
        err = autorest.NewErrorWithError(err, "databricks.WorkspacesClient", "Update", result.Response(), "Failure sending request")
        return
        }

    return
}

    // UpdatePreparer prepares the Update request.
    func (client WorkspacesClient) UpdatePreparer(ctx context.Context, parameters WorkspaceUpdate, resourceGroupName string, workspaceName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        "workspaceName": autorest.Encode("path",workspaceName),
        }

            const APIVersion = "2018-04-01"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsContentType("application/json; charset=utf-8"),
autorest.AsPatch(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Databricks/workspaces/{workspaceName}",pathParameters),
autorest.WithJSON(parameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // UpdateSender sends the Update request. The method will close the
    // http.Response Body if it receives an error.
    func (client WorkspacesClient) UpdateSender(req *http.Request) (future WorkspacesUpdateFuture, err error) {
            var resp *http.Response
            resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
            if err != nil {
            return
            }
            future.Future, err = azure.NewFutureFromResponse(resp)
            return
            }

    // UpdateResponder handles the response to the Update request. The method always
    // closes the http.Response Body.
    func (client WorkspacesClient) UpdateResponder(resp *http.Response) (result Workspace, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusAccepted),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

