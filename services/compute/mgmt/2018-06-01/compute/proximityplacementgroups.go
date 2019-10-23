package compute

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

// ProximityPlacementGroupsClient is the compute Client
type ProximityPlacementGroupsClient struct {
    BaseClient
}
// NewProximityPlacementGroupsClient creates an instance of the ProximityPlacementGroupsClient client.
func NewProximityPlacementGroupsClient(subscriptionID string) ProximityPlacementGroupsClient {
    return NewProximityPlacementGroupsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewProximityPlacementGroupsClientWithBaseURI creates an instance of the ProximityPlacementGroupsClient client.
    func NewProximityPlacementGroupsClientWithBaseURI(baseURI string, subscriptionID string) ProximityPlacementGroupsClient {
        return ProximityPlacementGroupsClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// CreateOrUpdate create or update a proximity placement group.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // proximityPlacementGroupName - the name of the proximity placement group.
        // parameters - parameters supplied to the Create Proximity Placement Group operation.
func (client ProximityPlacementGroupsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, proximityPlacementGroupName string, parameters ProximityPlacementGroup) (result ProximityPlacementGroup, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ProximityPlacementGroupsClient.CreateOrUpdate")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, proximityPlacementGroupName, parameters)
    if err != nil {
    err = autorest.NewErrorWithError(err, "compute.ProximityPlacementGroupsClient", "CreateOrUpdate", nil , "Failure preparing request")
    return
    }

            resp, err := client.CreateOrUpdateSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "compute.ProximityPlacementGroupsClient", "CreateOrUpdate", resp, "Failure sending request")
            return
            }

            result, err = client.CreateOrUpdateResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "compute.ProximityPlacementGroupsClient", "CreateOrUpdate", resp, "Failure responding to request")
            }

    return
    }

    // CreateOrUpdatePreparer prepares the CreateOrUpdate request.
    func (client ProximityPlacementGroupsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, proximityPlacementGroupName string, parameters ProximityPlacementGroup) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "proximityPlacementGroupName": autorest.Encode("path",proximityPlacementGroupName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2018-06-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPut(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/proximityPlacementGroups/{proximityPlacementGroupName}",pathParameters),
    autorest.WithJSON(parameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
    // http.Response Body if it receives an error.
    func (client ProximityPlacementGroupsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ProximityPlacementGroupsClient) CreateOrUpdateResponder(resp *http.Response) (result ProximityPlacementGroup, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusCreated),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// Delete delete a proximity placement group.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // proximityPlacementGroupName - the name of the proximity placement group.
func (client ProximityPlacementGroupsClient) Delete(ctx context.Context, resourceGroupName string, proximityPlacementGroupName string) (result autorest.Response, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ProximityPlacementGroupsClient.Delete")
        defer func() {
            sc := -1
            if result.Response != nil {
                sc = result.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.DeletePreparer(ctx, resourceGroupName, proximityPlacementGroupName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "compute.ProximityPlacementGroupsClient", "Delete", nil , "Failure preparing request")
    return
    }

            resp, err := client.DeleteSender(req)
            if err != nil {
            result.Response = resp
            err = autorest.NewErrorWithError(err, "compute.ProximityPlacementGroupsClient", "Delete", resp, "Failure sending request")
            return
            }

            result, err = client.DeleteResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "compute.ProximityPlacementGroupsClient", "Delete", resp, "Failure responding to request")
            }

    return
    }

    // DeletePreparer prepares the Delete request.
    func (client ProximityPlacementGroupsClient) DeletePreparer(ctx context.Context, resourceGroupName string, proximityPlacementGroupName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "proximityPlacementGroupName": autorest.Encode("path",proximityPlacementGroupName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2018-06-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsDelete(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/proximityPlacementGroups/{proximityPlacementGroupName}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // DeleteSender sends the Delete request. The method will close the
    // http.Response Body if it receives an error.
    func (client ProximityPlacementGroupsClient) DeleteSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ProximityPlacementGroupsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByClosing())
    result.Response = resp
        return
    }

// Get retrieves information about a proximity placement group .
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // proximityPlacementGroupName - the name of the proximity placement group.
func (client ProximityPlacementGroupsClient) Get(ctx context.Context, resourceGroupName string, proximityPlacementGroupName string) (result ProximityPlacementGroup, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ProximityPlacementGroupsClient.Get")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.GetPreparer(ctx, resourceGroupName, proximityPlacementGroupName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "compute.ProximityPlacementGroupsClient", "Get", nil , "Failure preparing request")
    return
    }

            resp, err := client.GetSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "compute.ProximityPlacementGroupsClient", "Get", resp, "Failure sending request")
            return
            }

            result, err = client.GetResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "compute.ProximityPlacementGroupsClient", "Get", resp, "Failure responding to request")
            }

    return
    }

    // GetPreparer prepares the Get request.
    func (client ProximityPlacementGroupsClient) GetPreparer(ctx context.Context, resourceGroupName string, proximityPlacementGroupName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "proximityPlacementGroupName": autorest.Encode("path",proximityPlacementGroupName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2018-06-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/proximityPlacementGroups/{proximityPlacementGroupName}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client ProximityPlacementGroupsClient) GetSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ProximityPlacementGroupsClient) GetResponder(resp *http.Response) (result ProximityPlacementGroup, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// ListByResourceGroup lists all proximity placement groups in a resource group.
    // Parameters:
        // resourceGroupName - the name of the resource group.
func (client ProximityPlacementGroupsClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (result ProximityPlacementGroupListResultPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ProximityPlacementGroupsClient.ListByResourceGroup")
        defer func() {
            sc := -1
            if result.ppglr.Response.Response != nil {
                sc = result.ppglr.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
                result.fn = client.listByResourceGroupNextResults
    req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "compute.ProximityPlacementGroupsClient", "ListByResourceGroup", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListByResourceGroupSender(req)
            if err != nil {
            result.ppglr.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "compute.ProximityPlacementGroupsClient", "ListByResourceGroup", resp, "Failure sending request")
            return
            }

            result.ppglr, err = client.ListByResourceGroupResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "compute.ProximityPlacementGroupsClient", "ListByResourceGroup", resp, "Failure responding to request")
            }

    return
    }

    // ListByResourceGroupPreparer prepares the ListByResourceGroup request.
    func (client ProximityPlacementGroupsClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2018-06-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/proximityPlacementGroups",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
    // http.Response Body if it receives an error.
    func (client ProximityPlacementGroupsClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client ProximityPlacementGroupsClient) ListByResourceGroupResponder(resp *http.Response) (result ProximityPlacementGroupListResult, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

            // listByResourceGroupNextResults retrieves the next set of results, if any.
            func (client ProximityPlacementGroupsClient) listByResourceGroupNextResults(ctx context.Context, lastResults ProximityPlacementGroupListResult) (result ProximityPlacementGroupListResult, err error) {
            req, err := lastResults.proximityPlacementGroupListResultPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "compute.ProximityPlacementGroupsClient", "listByResourceGroupNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListByResourceGroupSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "compute.ProximityPlacementGroupsClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListByResourceGroupResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "compute.ProximityPlacementGroupsClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

    // ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
    func (client ProximityPlacementGroupsClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result ProximityPlacementGroupListResultIterator, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/ProximityPlacementGroupsClient.ListByResourceGroup")
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

// ListBySubscription lists all proximity placement groups in a subscription.
func (client ProximityPlacementGroupsClient) ListBySubscription(ctx context.Context) (result ProximityPlacementGroupListResultPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ProximityPlacementGroupsClient.ListBySubscription")
        defer func() {
            sc := -1
            if result.ppglr.Response.Response != nil {
                sc = result.ppglr.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
                result.fn = client.listBySubscriptionNextResults
    req, err := client.ListBySubscriptionPreparer(ctx)
    if err != nil {
    err = autorest.NewErrorWithError(err, "compute.ProximityPlacementGroupsClient", "ListBySubscription", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListBySubscriptionSender(req)
            if err != nil {
            result.ppglr.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "compute.ProximityPlacementGroupsClient", "ListBySubscription", resp, "Failure sending request")
            return
            }

            result.ppglr, err = client.ListBySubscriptionResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "compute.ProximityPlacementGroupsClient", "ListBySubscription", resp, "Failure responding to request")
            }

    return
    }

    // ListBySubscriptionPreparer prepares the ListBySubscription request.
    func (client ProximityPlacementGroupsClient) ListBySubscriptionPreparer(ctx context.Context) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2018-06-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/proximityPlacementGroups",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListBySubscriptionSender sends the ListBySubscription request. The method will close the
    // http.Response Body if it receives an error.
    func (client ProximityPlacementGroupsClient) ListBySubscriptionSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// ListBySubscriptionResponder handles the response to the ListBySubscription request. The method always
// closes the http.Response Body.
func (client ProximityPlacementGroupsClient) ListBySubscriptionResponder(resp *http.Response) (result ProximityPlacementGroupListResult, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

            // listBySubscriptionNextResults retrieves the next set of results, if any.
            func (client ProximityPlacementGroupsClient) listBySubscriptionNextResults(ctx context.Context, lastResults ProximityPlacementGroupListResult) (result ProximityPlacementGroupListResult, err error) {
            req, err := lastResults.proximityPlacementGroupListResultPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "compute.ProximityPlacementGroupsClient", "listBySubscriptionNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListBySubscriptionSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "compute.ProximityPlacementGroupsClient", "listBySubscriptionNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListBySubscriptionResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "compute.ProximityPlacementGroupsClient", "listBySubscriptionNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

    // ListBySubscriptionComplete enumerates all values, automatically crossing page boundaries as required.
    func (client ProximityPlacementGroupsClient) ListBySubscriptionComplete(ctx context.Context) (result ProximityPlacementGroupListResultIterator, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/ProximityPlacementGroupsClient.ListBySubscription")
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

// Update update a proximity placement group.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // proximityPlacementGroupName - the name of the proximity placement group.
        // parameters - parameters supplied to the Update Proximity Placement Group operation.
func (client ProximityPlacementGroupsClient) Update(ctx context.Context, resourceGroupName string, proximityPlacementGroupName string, parameters ProximityPlacementGroupUpdate) (result ProximityPlacementGroup, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ProximityPlacementGroupsClient.Update")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.UpdatePreparer(ctx, resourceGroupName, proximityPlacementGroupName, parameters)
    if err != nil {
    err = autorest.NewErrorWithError(err, "compute.ProximityPlacementGroupsClient", "Update", nil , "Failure preparing request")
    return
    }

            resp, err := client.UpdateSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "compute.ProximityPlacementGroupsClient", "Update", resp, "Failure sending request")
            return
            }

            result, err = client.UpdateResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "compute.ProximityPlacementGroupsClient", "Update", resp, "Failure responding to request")
            }

    return
    }

    // UpdatePreparer prepares the Update request.
    func (client ProximityPlacementGroupsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, proximityPlacementGroupName string, parameters ProximityPlacementGroupUpdate) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "proximityPlacementGroupName": autorest.Encode("path",proximityPlacementGroupName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2018-06-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPatch(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/proximityPlacementGroups/{proximityPlacementGroupName}",pathParameters),
    autorest.WithJSON(parameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // UpdateSender sends the Update request. The method will close the
    // http.Response Body if it receives an error.
    func (client ProximityPlacementGroupsClient) UpdateSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client ProximityPlacementGroupsClient) UpdateResponder(resp *http.Response) (result ProximityPlacementGroup, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

