package managednetwork

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

// GroupsClient is the the Microsoft Azure Managed Network management API provides a RESTful set of web services that
// interact with Microsoft Azure Networks service to programmatically view, control, change, and monitor your entire
// Azure network centrally and with ease.
type GroupsClient struct {
    BaseClient
}
// NewGroupsClient creates an instance of the GroupsClient client.
func NewGroupsClient(subscriptionID string) GroupsClient {
    return NewGroupsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewGroupsClientWithBaseURI creates an instance of the GroupsClient client.
    func NewGroupsClientWithBaseURI(baseURI string, subscriptionID string) GroupsClient {
        return GroupsClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// CreateOrUpdate the Put ManagedNetworkGroups operation creates or updates a Managed Network Group resource
    // Parameters:
        // managedNetworkGroup - parameters supplied to the create/update a Managed Network Group resource
        // resourceGroupName - the name of the resource group.
        // managedNetworkName - the name of the Managed Network.
        // managedNetworkGroupName - the name of the Managed Network Group.
func (client GroupsClient) CreateOrUpdate(ctx context.Context, managedNetworkGroup Group, resourceGroupName string, managedNetworkName string, managedNetworkGroupName string) (result GroupsCreateOrUpdateFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/GroupsClient.CreateOrUpdate")
        defer func() {
            sc := -1
            if result.Response() != nil {
                sc = result.Response().StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.CreateOrUpdatePreparer(ctx, managedNetworkGroup, resourceGroupName, managedNetworkName, managedNetworkGroupName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "managednetwork.GroupsClient", "CreateOrUpdate", nil , "Failure preparing request")
    return
    }

            result, err = client.CreateOrUpdateSender(req)
            if err != nil {
            err = autorest.NewErrorWithError(err, "managednetwork.GroupsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
            return
            }

    return
    }

    // CreateOrUpdatePreparer prepares the CreateOrUpdate request.
    func (client GroupsClient) CreateOrUpdatePreparer(ctx context.Context, managedNetworkGroup Group, resourceGroupName string, managedNetworkName string, managedNetworkGroupName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "managedNetworkGroupName": autorest.Encode("path",managedNetworkGroupName),
            "managedNetworkName": autorest.Encode("path",managedNetworkName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2019-06-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPut(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetwork/managedNetworks/{managedNetworkName}/managedNetworkGroups/{managedNetworkGroupName}",pathParameters),
    autorest.WithJSON(managedNetworkGroup),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
    // http.Response Body if it receives an error.
    func (client GroupsClient) CreateOrUpdateSender(req *http.Request) (future GroupsCreateOrUpdateFuture, err error) {
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
func (client GroupsClient) CreateOrUpdateResponder(resp *http.Response) (result Group, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusCreated),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// Delete the Delete ManagedNetworkGroups operation deletes a Managed Network Group specified by the resource group,
// Managed Network name, and group name
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // managedNetworkName - the name of the Managed Network.
        // managedNetworkGroupName - the name of the Managed Network Group.
func (client GroupsClient) Delete(ctx context.Context, resourceGroupName string, managedNetworkName string, managedNetworkGroupName string) (result GroupsDeleteFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/GroupsClient.Delete")
        defer func() {
            sc := -1
            if result.Response() != nil {
                sc = result.Response().StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.DeletePreparer(ctx, resourceGroupName, managedNetworkName, managedNetworkGroupName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "managednetwork.GroupsClient", "Delete", nil , "Failure preparing request")
    return
    }

            result, err = client.DeleteSender(req)
            if err != nil {
            err = autorest.NewErrorWithError(err, "managednetwork.GroupsClient", "Delete", result.Response(), "Failure sending request")
            return
            }

    return
    }

    // DeletePreparer prepares the Delete request.
    func (client GroupsClient) DeletePreparer(ctx context.Context, resourceGroupName string, managedNetworkName string, managedNetworkGroupName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "managedNetworkGroupName": autorest.Encode("path",managedNetworkGroupName),
            "managedNetworkName": autorest.Encode("path",managedNetworkName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2019-06-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsDelete(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetwork/managedNetworks/{managedNetworkName}/managedNetworkGroups/{managedNetworkGroupName}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // DeleteSender sends the Delete request. The method will close the
    // http.Response Body if it receives an error.
    func (client GroupsClient) DeleteSender(req *http.Request) (future GroupsDeleteFuture, err error) {
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
func (client GroupsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusAccepted,http.StatusNoContent),
    autorest.ByClosing())
    result.Response = resp
        return
    }

// Get the Get ManagedNetworkGroups operation gets a Managed Network Group specified by the resource group, Managed
// Network name, and group name
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // managedNetworkName - the name of the Managed Network.
        // managedNetworkGroupName - the name of the Managed Network Group.
func (client GroupsClient) Get(ctx context.Context, resourceGroupName string, managedNetworkName string, managedNetworkGroupName string) (result Group, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/GroupsClient.Get")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.GetPreparer(ctx, resourceGroupName, managedNetworkName, managedNetworkGroupName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "managednetwork.GroupsClient", "Get", nil , "Failure preparing request")
    return
    }

            resp, err := client.GetSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "managednetwork.GroupsClient", "Get", resp, "Failure sending request")
            return
            }

            result, err = client.GetResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "managednetwork.GroupsClient", "Get", resp, "Failure responding to request")
            }

    return
    }

    // GetPreparer prepares the Get request.
    func (client GroupsClient) GetPreparer(ctx context.Context, resourceGroupName string, managedNetworkName string, managedNetworkGroupName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "managedNetworkGroupName": autorest.Encode("path",managedNetworkGroupName),
            "managedNetworkName": autorest.Encode("path",managedNetworkName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2019-06-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetwork/managedNetworks/{managedNetworkName}/managedNetworkGroups/{managedNetworkGroupName}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client GroupsClient) GetSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client GroupsClient) GetResponder(resp *http.Response) (result Group, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// ListByManagedNetwork the ListByManagedNetwork ManagedNetworkGroup operation retrieves all the Managed Network Groups
// in a specified Managed Networks in a paginated format.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // managedNetworkName - the name of the Managed Network.
        // top - may be used to limit the number of results in a page for list queries.
        // skiptoken - skiptoken is only used if a previous operation returned a partial result. If a previous response
        // contains a nextLink element, the value of the nextLink element will include a skiptoken parameter that
        // specifies a starting point to use for subsequent calls.
func (client GroupsClient) ListByManagedNetwork(ctx context.Context, resourceGroupName string, managedNetworkName string, top *int32, skiptoken string) (result GroupListResultPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/GroupsClient.ListByManagedNetwork")
        defer func() {
            sc := -1
            if result.glr.Response.Response != nil {
                sc = result.glr.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: top,
             Constraints: []validation.Constraint{	{Target: "top", Name: validation.Null, Rule: false ,
            Chain: []validation.Constraint{	{Target: "top", Name: validation.InclusiveMaximum, Rule: int64(20), Chain: nil },
            	{Target: "top", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil },
            }}}}}); err != nil {
            return result, validation.NewError("managednetwork.GroupsClient", "ListByManagedNetwork", err.Error())
            }

                        result.fn = client.listByManagedNetworkNextResults
    req, err := client.ListByManagedNetworkPreparer(ctx, resourceGroupName, managedNetworkName, top, skiptoken)
    if err != nil {
    err = autorest.NewErrorWithError(err, "managednetwork.GroupsClient", "ListByManagedNetwork", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListByManagedNetworkSender(req)
            if err != nil {
            result.glr.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "managednetwork.GroupsClient", "ListByManagedNetwork", resp, "Failure sending request")
            return
            }

            result.glr, err = client.ListByManagedNetworkResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "managednetwork.GroupsClient", "ListByManagedNetwork", resp, "Failure responding to request")
            }

    return
    }

    // ListByManagedNetworkPreparer prepares the ListByManagedNetwork request.
    func (client GroupsClient) ListByManagedNetworkPreparer(ctx context.Context, resourceGroupName string, managedNetworkName string, top *int32, skiptoken string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "managedNetworkName": autorest.Encode("path",managedNetworkName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2019-06-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }
            if top != nil {
            queryParameters["$top"] = autorest.Encode("query",*top)
            }
            if len(skiptoken) > 0 {
            queryParameters["$skiptoken"] = autorest.Encode("query",skiptoken)
            }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetwork/managedNetworks/{managedNetworkName}/managedNetworkGroups",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListByManagedNetworkSender sends the ListByManagedNetwork request. The method will close the
    // http.Response Body if it receives an error.
    func (client GroupsClient) ListByManagedNetworkSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// ListByManagedNetworkResponder handles the response to the ListByManagedNetwork request. The method always
// closes the http.Response Body.
func (client GroupsClient) ListByManagedNetworkResponder(resp *http.Response) (result GroupListResult, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

            // listByManagedNetworkNextResults retrieves the next set of results, if any.
            func (client GroupsClient) listByManagedNetworkNextResults(ctx context.Context, lastResults GroupListResult) (result GroupListResult, err error) {
            req, err := lastResults.groupListResultPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "managednetwork.GroupsClient", "listByManagedNetworkNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListByManagedNetworkSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "managednetwork.GroupsClient", "listByManagedNetworkNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListByManagedNetworkResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "managednetwork.GroupsClient", "listByManagedNetworkNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

    // ListByManagedNetworkComplete enumerates all values, automatically crossing page boundaries as required.
    func (client GroupsClient) ListByManagedNetworkComplete(ctx context.Context, resourceGroupName string, managedNetworkName string, top *int32, skiptoken string) (result GroupListResultIterator, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/GroupsClient.ListByManagedNetwork")
            defer func() {
                sc := -1
                if result.Response().Response.Response != nil {
                    sc = result.page.Response().Response.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
     }
        result.page, err = client.ListByManagedNetwork(ctx, resourceGroupName, managedNetworkName, top, skiptoken)
                return
        }

