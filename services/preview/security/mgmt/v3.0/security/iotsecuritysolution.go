package security

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

// IotSecuritySolutionClient is the API spec for Microsoft.Security (Azure Security Center) resource provider
type IotSecuritySolutionClient struct {
    BaseClient
}
// NewIotSecuritySolutionClient creates an instance of the IotSecuritySolutionClient client.
func NewIotSecuritySolutionClient(subscriptionID string, ascLocation string) IotSecuritySolutionClient {
    return NewIotSecuritySolutionClientWithBaseURI(DefaultBaseURI, subscriptionID, ascLocation)
}

// NewIotSecuritySolutionClientWithBaseURI creates an instance of the IotSecuritySolutionClient client.
    func NewIotSecuritySolutionClientWithBaseURI(baseURI string, subscriptionID string, ascLocation string) IotSecuritySolutionClient {
        return IotSecuritySolutionClient{ NewWithBaseURI(baseURI, subscriptionID, ascLocation)}
    }

// Create use this method to create or update yours IoT Security solution
    // Parameters:
        // resourceGroupName - the name of the resource group within the user's subscription. The name is case
        // insensitive.
        // solutionName - the name of the IoT Security solution.
        // iotSecuritySolutionData - the security solution data
func (client IotSecuritySolutionClient) Create(ctx context.Context, resourceGroupName string, solutionName string, iotSecuritySolutionData IoTSecuritySolutionModel) (result IoTSecuritySolutionModel, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/IotSecuritySolutionClient.Create")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: client.SubscriptionID,
             Constraints: []validation.Constraint{	{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil }}},
            { TargetValue: resourceGroupName,
             Constraints: []validation.Constraint{	{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil },
            	{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil },
            	{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil }}},
            { TargetValue: iotSecuritySolutionData,
             Constraints: []validation.Constraint{	{Target: "iotSecuritySolutionData.IoTSecuritySolutionProperties", Name: validation.Null, Rule: false ,
            Chain: []validation.Constraint{	{Target: "iotSecuritySolutionData.IoTSecuritySolutionProperties.Workspace", Name: validation.Null, Rule: true, Chain: nil },
            	{Target: "iotSecuritySolutionData.IoTSecuritySolutionProperties.DisplayName", Name: validation.Null, Rule: true, Chain: nil },
            	{Target: "iotSecuritySolutionData.IoTSecuritySolutionProperties.IotHubs", Name: validation.Null, Rule: true, Chain: nil },
            	{Target: "iotSecuritySolutionData.IoTSecuritySolutionProperties.UserDefinedResources", Name: validation.Null, Rule: false ,
            Chain: []validation.Constraint{	{Target: "iotSecuritySolutionData.IoTSecuritySolutionProperties.UserDefinedResources.Query", Name: validation.Null, Rule: true, Chain: nil },
            	{Target: "iotSecuritySolutionData.IoTSecuritySolutionProperties.UserDefinedResources.QuerySubscriptions", Name: validation.Null, Rule: true, Chain: nil },
            }},
            }}}}}); err != nil {
            return result, validation.NewError("security.IotSecuritySolutionClient", "Create", err.Error())
            }

                req, err := client.CreatePreparer(ctx, resourceGroupName, solutionName, iotSecuritySolutionData)
    if err != nil {
    err = autorest.NewErrorWithError(err, "security.IotSecuritySolutionClient", "Create", nil , "Failure preparing request")
    return
    }

            resp, err := client.CreateSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "security.IotSecuritySolutionClient", "Create", resp, "Failure sending request")
            return
            }

            result, err = client.CreateResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "security.IotSecuritySolutionClient", "Create", resp, "Failure responding to request")
            }

    return
    }

    // CreatePreparer prepares the Create request.
    func (client IotSecuritySolutionClient) CreatePreparer(ctx context.Context, resourceGroupName string, solutionName string, iotSecuritySolutionData IoTSecuritySolutionModel) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "solutionName": autorest.Encode("path",solutionName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2019-08-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

                    iotSecuritySolutionData.ID = nil
                iotSecuritySolutionData.Name = nil
                iotSecuritySolutionData.Type = nil
    preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPut(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/iotSecuritySolutions/{solutionName}",pathParameters),
    autorest.WithJSON(iotSecuritySolutionData),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CreateSender sends the Create request. The method will close the
    // http.Response Body if it receives an error.
    func (client IotSecuritySolutionClient) CreateSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client IotSecuritySolutionClient) CreateResponder(resp *http.Response) (result IoTSecuritySolutionModel, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusCreated),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// Delete use this method to delete yours IoT Security solution
    // Parameters:
        // resourceGroupName - the name of the resource group within the user's subscription. The name is case
        // insensitive.
        // solutionName - the name of the IoT Security solution.
func (client IotSecuritySolutionClient) Delete(ctx context.Context, resourceGroupName string, solutionName string) (result autorest.Response, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/IotSecuritySolutionClient.Delete")
        defer func() {
            sc := -1
            if result.Response != nil {
                sc = result.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: client.SubscriptionID,
             Constraints: []validation.Constraint{	{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil }}},
            { TargetValue: resourceGroupName,
             Constraints: []validation.Constraint{	{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil },
            	{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil },
            	{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("security.IotSecuritySolutionClient", "Delete", err.Error())
            }

                req, err := client.DeletePreparer(ctx, resourceGroupName, solutionName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "security.IotSecuritySolutionClient", "Delete", nil , "Failure preparing request")
    return
    }

            resp, err := client.DeleteSender(req)
            if err != nil {
            result.Response = resp
            err = autorest.NewErrorWithError(err, "security.IotSecuritySolutionClient", "Delete", resp, "Failure sending request")
            return
            }

            result, err = client.DeleteResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "security.IotSecuritySolutionClient", "Delete", resp, "Failure responding to request")
            }

    return
    }

    // DeletePreparer prepares the Delete request.
    func (client IotSecuritySolutionClient) DeletePreparer(ctx context.Context, resourceGroupName string, solutionName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "solutionName": autorest.Encode("path",solutionName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2019-08-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsDelete(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/iotSecuritySolutions/{solutionName}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // DeleteSender sends the Delete request. The method will close the
    // http.Response Body if it receives an error.
    func (client IotSecuritySolutionClient) DeleteSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client IotSecuritySolutionClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusNoContent),
    autorest.ByClosing())
    result.Response = resp
        return
    }

// Get user this method to get details of a specific IoT Security solution based on solution name
    // Parameters:
        // resourceGroupName - the name of the resource group within the user's subscription. The name is case
        // insensitive.
        // solutionName - the name of the IoT Security solution.
func (client IotSecuritySolutionClient) Get(ctx context.Context, resourceGroupName string, solutionName string) (result IoTSecuritySolutionModel, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/IotSecuritySolutionClient.Get")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: client.SubscriptionID,
             Constraints: []validation.Constraint{	{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil }}},
            { TargetValue: resourceGroupName,
             Constraints: []validation.Constraint{	{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil },
            	{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil },
            	{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("security.IotSecuritySolutionClient", "Get", err.Error())
            }

                req, err := client.GetPreparer(ctx, resourceGroupName, solutionName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "security.IotSecuritySolutionClient", "Get", nil , "Failure preparing request")
    return
    }

            resp, err := client.GetSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "security.IotSecuritySolutionClient", "Get", resp, "Failure sending request")
            return
            }

            result, err = client.GetResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "security.IotSecuritySolutionClient", "Get", resp, "Failure responding to request")
            }

    return
    }

    // GetPreparer prepares the Get request.
    func (client IotSecuritySolutionClient) GetPreparer(ctx context.Context, resourceGroupName string, solutionName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "solutionName": autorest.Encode("path",solutionName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2019-08-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/iotSecuritySolutions/{solutionName}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client IotSecuritySolutionClient) GetSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client IotSecuritySolutionClient) GetResponder(resp *http.Response) (result IoTSecuritySolutionModel, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// ListByResourceGroup use this method to get the list IoT Security solutions organized by resource group.
    // Parameters:
        // resourceGroupName - the name of the resource group within the user's subscription. The name is case
        // insensitive.
        // filter - filter the IoT Security solution with OData syntax. Supports filtering by iotHubs.
func (client IotSecuritySolutionClient) ListByResourceGroup(ctx context.Context, resourceGroupName string, filter string) (result IoTSecuritySolutionsListPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/IotSecuritySolutionClient.ListByResourceGroup")
        defer func() {
            sc := -1
            if result.itssl.Response.Response != nil {
                sc = result.itssl.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: client.SubscriptionID,
             Constraints: []validation.Constraint{	{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil }}},
            { TargetValue: resourceGroupName,
             Constraints: []validation.Constraint{	{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil },
            	{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil },
            	{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("security.IotSecuritySolutionClient", "ListByResourceGroup", err.Error())
            }

                        result.fn = client.listByResourceGroupNextResults
    req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName, filter)
    if err != nil {
    err = autorest.NewErrorWithError(err, "security.IotSecuritySolutionClient", "ListByResourceGroup", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListByResourceGroupSender(req)
            if err != nil {
            result.itssl.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "security.IotSecuritySolutionClient", "ListByResourceGroup", resp, "Failure sending request")
            return
            }

            result.itssl, err = client.ListByResourceGroupResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "security.IotSecuritySolutionClient", "ListByResourceGroup", resp, "Failure responding to request")
            }

    return
    }

    // ListByResourceGroupPreparer prepares the ListByResourceGroup request.
    func (client IotSecuritySolutionClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string, filter string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2019-08-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }
            if len(filter) > 0 {
            queryParameters["$filter"] = autorest.Encode("query",filter)
            }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/iotSecuritySolutions",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
    // http.Response Body if it receives an error.
    func (client IotSecuritySolutionClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client IotSecuritySolutionClient) ListByResourceGroupResponder(resp *http.Response) (result IoTSecuritySolutionsList, err error) {
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
            func (client IotSecuritySolutionClient) listByResourceGroupNextResults(ctx context.Context, lastResults IoTSecuritySolutionsList) (result IoTSecuritySolutionsList, err error) {
            req, err := lastResults.ioTSecuritySolutionsListPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "security.IotSecuritySolutionClient", "listByResourceGroupNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListByResourceGroupSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "security.IotSecuritySolutionClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListByResourceGroupResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "security.IotSecuritySolutionClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

    // ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
    func (client IotSecuritySolutionClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, filter string) (result IoTSecuritySolutionsListIterator, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/IotSecuritySolutionClient.ListByResourceGroup")
            defer func() {
                sc := -1
                if result.Response().Response.Response != nil {
                    sc = result.page.Response().Response.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
     }
        result.page, err = client.ListByResourceGroup(ctx, resourceGroupName, filter)
                return
        }

// ListBySubscription use this method to get the list of IoT Security solutions by subscription.
    // Parameters:
        // filter - filter the IoT Security solution with OData syntax. Supports filtering by iotHubs.
func (client IotSecuritySolutionClient) ListBySubscription(ctx context.Context, filter string) (result IoTSecuritySolutionsListPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/IotSecuritySolutionClient.ListBySubscription")
        defer func() {
            sc := -1
            if result.itssl.Response.Response != nil {
                sc = result.itssl.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: client.SubscriptionID,
             Constraints: []validation.Constraint{	{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("security.IotSecuritySolutionClient", "ListBySubscription", err.Error())
            }

                        result.fn = client.listBySubscriptionNextResults
    req, err := client.ListBySubscriptionPreparer(ctx, filter)
    if err != nil {
    err = autorest.NewErrorWithError(err, "security.IotSecuritySolutionClient", "ListBySubscription", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListBySubscriptionSender(req)
            if err != nil {
            result.itssl.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "security.IotSecuritySolutionClient", "ListBySubscription", resp, "Failure sending request")
            return
            }

            result.itssl, err = client.ListBySubscriptionResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "security.IotSecuritySolutionClient", "ListBySubscription", resp, "Failure responding to request")
            }

    return
    }

    // ListBySubscriptionPreparer prepares the ListBySubscription request.
    func (client IotSecuritySolutionClient) ListBySubscriptionPreparer(ctx context.Context, filter string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2019-08-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }
            if len(filter) > 0 {
            queryParameters["$filter"] = autorest.Encode("query",filter)
            }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Security/iotSecuritySolutions",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListBySubscriptionSender sends the ListBySubscription request. The method will close the
    // http.Response Body if it receives an error.
    func (client IotSecuritySolutionClient) ListBySubscriptionSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// ListBySubscriptionResponder handles the response to the ListBySubscription request. The method always
// closes the http.Response Body.
func (client IotSecuritySolutionClient) ListBySubscriptionResponder(resp *http.Response) (result IoTSecuritySolutionsList, err error) {
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
            func (client IotSecuritySolutionClient) listBySubscriptionNextResults(ctx context.Context, lastResults IoTSecuritySolutionsList) (result IoTSecuritySolutionsList, err error) {
            req, err := lastResults.ioTSecuritySolutionsListPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "security.IotSecuritySolutionClient", "listBySubscriptionNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListBySubscriptionSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "security.IotSecuritySolutionClient", "listBySubscriptionNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListBySubscriptionResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "security.IotSecuritySolutionClient", "listBySubscriptionNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

    // ListBySubscriptionComplete enumerates all values, automatically crossing page boundaries as required.
    func (client IotSecuritySolutionClient) ListBySubscriptionComplete(ctx context.Context, filter string) (result IoTSecuritySolutionsListIterator, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/IotSecuritySolutionClient.ListBySubscription")
            defer func() {
                sc := -1
                if result.Response().Response.Response != nil {
                    sc = result.page.Response().Response.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
     }
        result.page, err = client.ListBySubscription(ctx, filter)
                return
        }

// Update use this method to update existing IoT Security solution tags or user defined resources. To update other
// fields use the CreateOrUpdate method.
    // Parameters:
        // resourceGroupName - the name of the resource group within the user's subscription. The name is case
        // insensitive.
        // solutionName - the name of the IoT Security solution.
        // updateIotSecuritySolutionData - the security solution data
func (client IotSecuritySolutionClient) Update(ctx context.Context, resourceGroupName string, solutionName string, updateIotSecuritySolutionData UpdateIotSecuritySolutionData) (result IoTSecuritySolutionModel, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/IotSecuritySolutionClient.Update")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: client.SubscriptionID,
             Constraints: []validation.Constraint{	{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil }}},
            { TargetValue: resourceGroupName,
             Constraints: []validation.Constraint{	{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil },
            	{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil },
            	{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("security.IotSecuritySolutionClient", "Update", err.Error())
            }

                req, err := client.UpdatePreparer(ctx, resourceGroupName, solutionName, updateIotSecuritySolutionData)
    if err != nil {
    err = autorest.NewErrorWithError(err, "security.IotSecuritySolutionClient", "Update", nil , "Failure preparing request")
    return
    }

            resp, err := client.UpdateSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "security.IotSecuritySolutionClient", "Update", resp, "Failure sending request")
            return
            }

            result, err = client.UpdateResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "security.IotSecuritySolutionClient", "Update", resp, "Failure responding to request")
            }

    return
    }

    // UpdatePreparer prepares the Update request.
    func (client IotSecuritySolutionClient) UpdatePreparer(ctx context.Context, resourceGroupName string, solutionName string, updateIotSecuritySolutionData UpdateIotSecuritySolutionData) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "solutionName": autorest.Encode("path",solutionName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2019-08-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPatch(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/iotSecuritySolutions/{solutionName}",pathParameters),
    autorest.WithJSON(updateIotSecuritySolutionData),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // UpdateSender sends the Update request. The method will close the
    // http.Response Body if it receives an error.
    func (client IotSecuritySolutionClient) UpdateSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client IotSecuritySolutionClient) UpdateResponder(resp *http.Response) (result IoTSecuritySolutionModel, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

