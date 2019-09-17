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

// PeeringPoliciesClient is the the Microsoft Azure Managed Network management API provides a RESTful set of web
// services that interact with Microsoft Azure Networks service to programmatically view, control, change, and monitor
// your entire Azure network centrally and with ease.
type PeeringPoliciesClient struct {
    BaseClient
}
// NewPeeringPoliciesClient creates an instance of the PeeringPoliciesClient client.
func NewPeeringPoliciesClient(subscriptionID string) PeeringPoliciesClient {
    return NewPeeringPoliciesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewPeeringPoliciesClientWithBaseURI creates an instance of the PeeringPoliciesClient client.
    func NewPeeringPoliciesClientWithBaseURI(baseURI string, subscriptionID string) PeeringPoliciesClient {
        return PeeringPoliciesClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// CreateOrUpdate the Put ManagedNetworkPeeringPolicies operation creates/updates a new Managed Network Peering Policy
    // Parameters:
        // managedNetworkPolicy - parameters supplied to create/update a Managed Network Peering Policy
        // resourceGroupName - the name of the resource group.
        // managedNetworkName - the name of the Managed Network.
        // managedNetworkPeeringPolicyName - the name of the Managed Network Peering Policy.
func (client PeeringPoliciesClient) CreateOrUpdate(ctx context.Context, managedNetworkPolicy PeeringPolicy, resourceGroupName string, managedNetworkName string, managedNetworkPeeringPolicyName string) (result PeeringPoliciesCreateOrUpdateFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/PeeringPoliciesClient.CreateOrUpdate")
        defer func() {
            sc := -1
            if result.Response() != nil {
                sc = result.Response().StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.CreateOrUpdatePreparer(ctx, managedNetworkPolicy, resourceGroupName, managedNetworkName, managedNetworkPeeringPolicyName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "managednetwork.PeeringPoliciesClient", "CreateOrUpdate", nil , "Failure preparing request")
    return
    }

            result, err = client.CreateOrUpdateSender(req)
            if err != nil {
            err = autorest.NewErrorWithError(err, "managednetwork.PeeringPoliciesClient", "CreateOrUpdate", result.Response(), "Failure sending request")
            return
            }

    return
    }

    // CreateOrUpdatePreparer prepares the CreateOrUpdate request.
    func (client PeeringPoliciesClient) CreateOrUpdatePreparer(ctx context.Context, managedNetworkPolicy PeeringPolicy, resourceGroupName string, managedNetworkName string, managedNetworkPeeringPolicyName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "managedNetworkName": autorest.Encode("path",managedNetworkName),
            "managedNetworkPeeringPolicyName": autorest.Encode("path",managedNetworkPeeringPolicyName),
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
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetwork/managedNetworks/{managedNetworkName}/managedNetworkPeeringPolicies/{managedNetworkPeeringPolicyName}",pathParameters),
    autorest.WithJSON(managedNetworkPolicy),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
    // http.Response Body if it receives an error.
    func (client PeeringPoliciesClient) CreateOrUpdateSender(req *http.Request) (future PeeringPoliciesCreateOrUpdateFuture, err error) {
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
func (client PeeringPoliciesClient) CreateOrUpdateResponder(resp *http.Response) (result PeeringPolicy, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusCreated),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// Delete the Delete ManagedNetworkPeeringPolicies operation deletes a Managed Network Peering Policy, specified by the
// resource group, Managed Network name, and peering policy name
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // managedNetworkName - the name of the Managed Network.
        // managedNetworkPeeringPolicyName - the name of the Managed Network Peering Policy.
func (client PeeringPoliciesClient) Delete(ctx context.Context, resourceGroupName string, managedNetworkName string, managedNetworkPeeringPolicyName string) (result PeeringPoliciesDeleteFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/PeeringPoliciesClient.Delete")
        defer func() {
            sc := -1
            if result.Response() != nil {
                sc = result.Response().StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.DeletePreparer(ctx, resourceGroupName, managedNetworkName, managedNetworkPeeringPolicyName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "managednetwork.PeeringPoliciesClient", "Delete", nil , "Failure preparing request")
    return
    }

            result, err = client.DeleteSender(req)
            if err != nil {
            err = autorest.NewErrorWithError(err, "managednetwork.PeeringPoliciesClient", "Delete", result.Response(), "Failure sending request")
            return
            }

    return
    }

    // DeletePreparer prepares the Delete request.
    func (client PeeringPoliciesClient) DeletePreparer(ctx context.Context, resourceGroupName string, managedNetworkName string, managedNetworkPeeringPolicyName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "managedNetworkName": autorest.Encode("path",managedNetworkName),
            "managedNetworkPeeringPolicyName": autorest.Encode("path",managedNetworkPeeringPolicyName),
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
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetwork/managedNetworks/{managedNetworkName}/managedNetworkPeeringPolicies/{managedNetworkPeeringPolicyName}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // DeleteSender sends the Delete request. The method will close the
    // http.Response Body if it receives an error.
    func (client PeeringPoliciesClient) DeleteSender(req *http.Request) (future PeeringPoliciesDeleteFuture, err error) {
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
func (client PeeringPoliciesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusAccepted,http.StatusNoContent),
    autorest.ByClosing())
    result.Response = resp
        return
    }

// Get the Get ManagedNetworkPeeringPolicies operation gets a Managed Network Peering Policy resource, specified by the
// resource group, Managed Network name, and peering policy name
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // managedNetworkName - the name of the Managed Network.
        // managedNetworkPeeringPolicyName - the name of the Managed Network Peering Policy.
func (client PeeringPoliciesClient) Get(ctx context.Context, resourceGroupName string, managedNetworkName string, managedNetworkPeeringPolicyName string) (result PeeringPolicy, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/PeeringPoliciesClient.Get")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.GetPreparer(ctx, resourceGroupName, managedNetworkName, managedNetworkPeeringPolicyName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "managednetwork.PeeringPoliciesClient", "Get", nil , "Failure preparing request")
    return
    }

            resp, err := client.GetSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "managednetwork.PeeringPoliciesClient", "Get", resp, "Failure sending request")
            return
            }

            result, err = client.GetResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "managednetwork.PeeringPoliciesClient", "Get", resp, "Failure responding to request")
            }

    return
    }

    // GetPreparer prepares the Get request.
    func (client PeeringPoliciesClient) GetPreparer(ctx context.Context, resourceGroupName string, managedNetworkName string, managedNetworkPeeringPolicyName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "managedNetworkName": autorest.Encode("path",managedNetworkName),
            "managedNetworkPeeringPolicyName": autorest.Encode("path",managedNetworkPeeringPolicyName),
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
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetwork/managedNetworks/{managedNetworkName}/managedNetworkPeeringPolicies/{managedNetworkPeeringPolicyName}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client PeeringPoliciesClient) GetSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client PeeringPoliciesClient) GetResponder(resp *http.Response) (result PeeringPolicy, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// ListByManagedNetwork the ListByManagedNetwork PeeringPolicies operation retrieves all the Managed Network Peering
// Policies in a specified Managed Network, in a paginated format.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // managedNetworkName - the name of the Managed Network.
        // top - may be used to limit the number of results in a page for list queries.
        // skiptoken - skiptoken is only used if a previous operation returned a partial result. If a previous response
        // contains a nextLink element, the value of the nextLink element will include a skiptoken parameter that
        // specifies a starting point to use for subsequent calls.
func (client PeeringPoliciesClient) ListByManagedNetwork(ctx context.Context, resourceGroupName string, managedNetworkName string, top *int32, skiptoken string) (result PeeringPolicyListResultPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/PeeringPoliciesClient.ListByManagedNetwork")
        defer func() {
            sc := -1
            if result.pplr.Response.Response != nil {
                sc = result.pplr.Response.Response.StatusCode
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
            return result, validation.NewError("managednetwork.PeeringPoliciesClient", "ListByManagedNetwork", err.Error())
            }

                        result.fn = client.listByManagedNetworkNextResults
    req, err := client.ListByManagedNetworkPreparer(ctx, resourceGroupName, managedNetworkName, top, skiptoken)
    if err != nil {
    err = autorest.NewErrorWithError(err, "managednetwork.PeeringPoliciesClient", "ListByManagedNetwork", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListByManagedNetworkSender(req)
            if err != nil {
            result.pplr.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "managednetwork.PeeringPoliciesClient", "ListByManagedNetwork", resp, "Failure sending request")
            return
            }

            result.pplr, err = client.ListByManagedNetworkResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "managednetwork.PeeringPoliciesClient", "ListByManagedNetwork", resp, "Failure responding to request")
            }

    return
    }

    // ListByManagedNetworkPreparer prepares the ListByManagedNetwork request.
    func (client PeeringPoliciesClient) ListByManagedNetworkPreparer(ctx context.Context, resourceGroupName string, managedNetworkName string, top *int32, skiptoken string) (*http.Request, error) {
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
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetwork/managedNetworks/{managedNetworkName}/managedNetworkPeeringPolicies",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListByManagedNetworkSender sends the ListByManagedNetwork request. The method will close the
    // http.Response Body if it receives an error.
    func (client PeeringPoliciesClient) ListByManagedNetworkSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// ListByManagedNetworkResponder handles the response to the ListByManagedNetwork request. The method always
// closes the http.Response Body.
func (client PeeringPoliciesClient) ListByManagedNetworkResponder(resp *http.Response) (result PeeringPolicyListResult, err error) {
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
            func (client PeeringPoliciesClient) listByManagedNetworkNextResults(ctx context.Context, lastResults PeeringPolicyListResult) (result PeeringPolicyListResult, err error) {
            req, err := lastResults.peeringPolicyListResultPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "managednetwork.PeeringPoliciesClient", "listByManagedNetworkNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListByManagedNetworkSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "managednetwork.PeeringPoliciesClient", "listByManagedNetworkNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListByManagedNetworkResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "managednetwork.PeeringPoliciesClient", "listByManagedNetworkNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

    // ListByManagedNetworkComplete enumerates all values, automatically crossing page boundaries as required.
    func (client PeeringPoliciesClient) ListByManagedNetworkComplete(ctx context.Context, resourceGroupName string, managedNetworkName string, top *int32, skiptoken string) (result PeeringPolicyListResultIterator, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/PeeringPoliciesClient.ListByManagedNetwork")
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

