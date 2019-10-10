package frontdoor

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

// FrontendEndpointsClient is the frontDoor Client
type FrontendEndpointsClient struct {
    BaseClient
}
// NewFrontendEndpointsClient creates an instance of the FrontendEndpointsClient client.
func NewFrontendEndpointsClient(subscriptionID string) FrontendEndpointsClient {
    return NewFrontendEndpointsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewFrontendEndpointsClientWithBaseURI creates an instance of the FrontendEndpointsClient client.
    func NewFrontendEndpointsClientWithBaseURI(baseURI string, subscriptionID string) FrontendEndpointsClient {
        return FrontendEndpointsClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// DisableHTTPS disables a frontendEndpoint for HTTPS traffic
    // Parameters:
        // resourceGroupName - name of the Resource group within the Azure subscription.
        // frontDoorName - name of the Front Door which is globally unique.
        // frontendEndpointName - name of the Frontend endpoint which is unique within the Front Door.
func (client FrontendEndpointsClient) DisableHTTPS(ctx context.Context, resourceGroupName string, frontDoorName string, frontendEndpointName string) (result FrontendEndpointsDisableHTTPSFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/FrontendEndpointsClient.DisableHTTPS")
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
             Constraints: []validation.Constraint{	{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 80, Chain: nil },
            	{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil },
            	{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_\-\(\)\.]*[^\.]$`, Chain: nil }}},
            { TargetValue: frontDoorName,
             Constraints: []validation.Constraint{	{Target: "frontDoorName", Name: validation.MaxLength, Rule: 64, Chain: nil },
            	{Target: "frontDoorName", Name: validation.MinLength, Rule: 5, Chain: nil },
            	{Target: "frontDoorName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]+([-a-zA-Z0-9]?[a-zA-Z0-9])*$`, Chain: nil }}},
            { TargetValue: frontendEndpointName,
             Constraints: []validation.Constraint{	{Target: "frontendEndpointName", Name: validation.MaxLength, Rule: 255, Chain: nil },
            	{Target: "frontendEndpointName", Name: validation.MinLength, Rule: 1, Chain: nil },
            	{Target: "frontendEndpointName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]+(-*[a-zA-Z0-9])*$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("frontdoor.FrontendEndpointsClient", "DisableHTTPS", err.Error())
            }

                req, err := client.DisableHTTPSPreparer(ctx, resourceGroupName, frontDoorName, frontendEndpointName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "frontdoor.FrontendEndpointsClient", "DisableHTTPS", nil , "Failure preparing request")
    return
    }

            result, err = client.DisableHTTPSSender(req)
            if err != nil {
            err = autorest.NewErrorWithError(err, "frontdoor.FrontendEndpointsClient", "DisableHTTPS", result.Response(), "Failure sending request")
            return
            }

    return
    }

    // DisableHTTPSPreparer prepares the DisableHTTPS request.
    func (client FrontendEndpointsClient) DisableHTTPSPreparer(ctx context.Context, resourceGroupName string, frontDoorName string, frontendEndpointName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "frontDoorName": autorest.Encode("path",frontDoorName),
            "frontendEndpointName": autorest.Encode("path",frontendEndpointName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2019-05-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsPost(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/frontDoors/{frontDoorName}/frontendEndpoints/{frontendEndpointName}/disableHttps",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // DisableHTTPSSender sends the DisableHTTPS request. The method will close the
    // http.Response Body if it receives an error.
    func (client FrontendEndpointsClient) DisableHTTPSSender(req *http.Request) (future FrontendEndpointsDisableHTTPSFuture, err error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            var resp *http.Response
            resp, err = autorest.SendWithSender(client, req, sd...)
            if err != nil {
            return
            }
            future.Future, err = azure.NewFutureFromResponse(resp)
            return
            }

// DisableHTTPSResponder handles the response to the DisableHTTPS request. The method always
// closes the http.Response Body.
func (client FrontendEndpointsClient) DisableHTTPSResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusAccepted),
    autorest.ByClosing())
    result.Response = resp
        return
    }

// EnableHTTPS enables a frontendEndpoint for HTTPS traffic
    // Parameters:
        // resourceGroupName - name of the Resource group within the Azure subscription.
        // frontDoorName - name of the Front Door which is globally unique.
        // frontendEndpointName - name of the Frontend endpoint which is unique within the Front Door.
        // customHTTPSConfiguration - the configuration specifying how to enable HTTPS
func (client FrontendEndpointsClient) EnableHTTPS(ctx context.Context, resourceGroupName string, frontDoorName string, frontendEndpointName string, customHTTPSConfiguration CustomHTTPSConfiguration) (result FrontendEndpointsEnableHTTPSFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/FrontendEndpointsClient.EnableHTTPS")
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
             Constraints: []validation.Constraint{	{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 80, Chain: nil },
            	{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil },
            	{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_\-\(\)\.]*[^\.]$`, Chain: nil }}},
            { TargetValue: frontDoorName,
             Constraints: []validation.Constraint{	{Target: "frontDoorName", Name: validation.MaxLength, Rule: 64, Chain: nil },
            	{Target: "frontDoorName", Name: validation.MinLength, Rule: 5, Chain: nil },
            	{Target: "frontDoorName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]+([-a-zA-Z0-9]?[a-zA-Z0-9])*$`, Chain: nil }}},
            { TargetValue: frontendEndpointName,
             Constraints: []validation.Constraint{	{Target: "frontendEndpointName", Name: validation.MaxLength, Rule: 255, Chain: nil },
            	{Target: "frontendEndpointName", Name: validation.MinLength, Rule: 1, Chain: nil },
            	{Target: "frontendEndpointName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]+(-*[a-zA-Z0-9])*$`, Chain: nil }}},
            { TargetValue: customHTTPSConfiguration,
             Constraints: []validation.Constraint{	{Target: "customHTTPSConfiguration.ProtocolType", Name: validation.Null, Rule: true, Chain: nil }}}}); err != nil {
            return result, validation.NewError("frontdoor.FrontendEndpointsClient", "EnableHTTPS", err.Error())
            }

                req, err := client.EnableHTTPSPreparer(ctx, resourceGroupName, frontDoorName, frontendEndpointName, customHTTPSConfiguration)
    if err != nil {
    err = autorest.NewErrorWithError(err, "frontdoor.FrontendEndpointsClient", "EnableHTTPS", nil , "Failure preparing request")
    return
    }

            result, err = client.EnableHTTPSSender(req)
            if err != nil {
            err = autorest.NewErrorWithError(err, "frontdoor.FrontendEndpointsClient", "EnableHTTPS", result.Response(), "Failure sending request")
            return
            }

    return
    }

    // EnableHTTPSPreparer prepares the EnableHTTPS request.
    func (client FrontendEndpointsClient) EnableHTTPSPreparer(ctx context.Context, resourceGroupName string, frontDoorName string, frontendEndpointName string, customHTTPSConfiguration CustomHTTPSConfiguration) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "frontDoorName": autorest.Encode("path",frontDoorName),
            "frontendEndpointName": autorest.Encode("path",frontendEndpointName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2019-05-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPost(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/frontDoors/{frontDoorName}/frontendEndpoints/{frontendEndpointName}/enableHttps",pathParameters),
    autorest.WithJSON(customHTTPSConfiguration),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // EnableHTTPSSender sends the EnableHTTPS request. The method will close the
    // http.Response Body if it receives an error.
    func (client FrontendEndpointsClient) EnableHTTPSSender(req *http.Request) (future FrontendEndpointsEnableHTTPSFuture, err error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            var resp *http.Response
            resp, err = autorest.SendWithSender(client, req, sd...)
            if err != nil {
            return
            }
            future.Future, err = azure.NewFutureFromResponse(resp)
            return
            }

// EnableHTTPSResponder handles the response to the EnableHTTPS request. The method always
// closes the http.Response Body.
func (client FrontendEndpointsClient) EnableHTTPSResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusAccepted),
    autorest.ByClosing())
    result.Response = resp
        return
    }

// Get gets a Frontend endpoint with the specified name within the specified Front Door.
    // Parameters:
        // resourceGroupName - name of the Resource group within the Azure subscription.
        // frontDoorName - name of the Front Door which is globally unique.
        // frontendEndpointName - name of the Frontend endpoint which is unique within the Front Door.
func (client FrontendEndpointsClient) Get(ctx context.Context, resourceGroupName string, frontDoorName string, frontendEndpointName string) (result FrontendEndpoint, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/FrontendEndpointsClient.Get")
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
             Constraints: []validation.Constraint{	{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 80, Chain: nil },
            	{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil },
            	{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_\-\(\)\.]*[^\.]$`, Chain: nil }}},
            { TargetValue: frontDoorName,
             Constraints: []validation.Constraint{	{Target: "frontDoorName", Name: validation.MaxLength, Rule: 64, Chain: nil },
            	{Target: "frontDoorName", Name: validation.MinLength, Rule: 5, Chain: nil },
            	{Target: "frontDoorName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]+([-a-zA-Z0-9]?[a-zA-Z0-9])*$`, Chain: nil }}},
            { TargetValue: frontendEndpointName,
             Constraints: []validation.Constraint{	{Target: "frontendEndpointName", Name: validation.MaxLength, Rule: 255, Chain: nil },
            	{Target: "frontendEndpointName", Name: validation.MinLength, Rule: 1, Chain: nil },
            	{Target: "frontendEndpointName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]+(-*[a-zA-Z0-9])*$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("frontdoor.FrontendEndpointsClient", "Get", err.Error())
            }

                req, err := client.GetPreparer(ctx, resourceGroupName, frontDoorName, frontendEndpointName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "frontdoor.FrontendEndpointsClient", "Get", nil , "Failure preparing request")
    return
    }

            resp, err := client.GetSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "frontdoor.FrontendEndpointsClient", "Get", resp, "Failure sending request")
            return
            }

            result, err = client.GetResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "frontdoor.FrontendEndpointsClient", "Get", resp, "Failure responding to request")
            }

    return
    }

    // GetPreparer prepares the Get request.
    func (client FrontendEndpointsClient) GetPreparer(ctx context.Context, resourceGroupName string, frontDoorName string, frontendEndpointName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "frontDoorName": autorest.Encode("path",frontDoorName),
            "frontendEndpointName": autorest.Encode("path",frontendEndpointName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2019-05-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/frontDoors/{frontDoorName}/frontendEndpoints/{frontendEndpointName}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client FrontendEndpointsClient) GetSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client FrontendEndpointsClient) GetResponder(resp *http.Response) (result FrontendEndpoint, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// ListByFrontDoor lists all of the frontend endpoints within a Front Door.
    // Parameters:
        // resourceGroupName - name of the Resource group within the Azure subscription.
        // frontDoorName - name of the Front Door which is globally unique.
func (client FrontendEndpointsClient) ListByFrontDoor(ctx context.Context, resourceGroupName string, frontDoorName string) (result FrontendEndpointsListResultPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/FrontendEndpointsClient.ListByFrontDoor")
        defer func() {
            sc := -1
            if result.felr.Response.Response != nil {
                sc = result.felr.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: resourceGroupName,
             Constraints: []validation.Constraint{	{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 80, Chain: nil },
            	{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil },
            	{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_\-\(\)\.]*[^\.]$`, Chain: nil }}},
            { TargetValue: frontDoorName,
             Constraints: []validation.Constraint{	{Target: "frontDoorName", Name: validation.MaxLength, Rule: 64, Chain: nil },
            	{Target: "frontDoorName", Name: validation.MinLength, Rule: 5, Chain: nil },
            	{Target: "frontDoorName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]+([-a-zA-Z0-9]?[a-zA-Z0-9])*$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("frontdoor.FrontendEndpointsClient", "ListByFrontDoor", err.Error())
            }

                        result.fn = client.listByFrontDoorNextResults
    req, err := client.ListByFrontDoorPreparer(ctx, resourceGroupName, frontDoorName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "frontdoor.FrontendEndpointsClient", "ListByFrontDoor", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListByFrontDoorSender(req)
            if err != nil {
            result.felr.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "frontdoor.FrontendEndpointsClient", "ListByFrontDoor", resp, "Failure sending request")
            return
            }

            result.felr, err = client.ListByFrontDoorResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "frontdoor.FrontendEndpointsClient", "ListByFrontDoor", resp, "Failure responding to request")
            }

    return
    }

    // ListByFrontDoorPreparer prepares the ListByFrontDoor request.
    func (client FrontendEndpointsClient) ListByFrontDoorPreparer(ctx context.Context, resourceGroupName string, frontDoorName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "frontDoorName": autorest.Encode("path",frontDoorName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2019-05-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/frontDoors/{frontDoorName}/frontendEndpoints",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListByFrontDoorSender sends the ListByFrontDoor request. The method will close the
    // http.Response Body if it receives an error.
    func (client FrontendEndpointsClient) ListByFrontDoorSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// ListByFrontDoorResponder handles the response to the ListByFrontDoor request. The method always
// closes the http.Response Body.
func (client FrontendEndpointsClient) ListByFrontDoorResponder(resp *http.Response) (result FrontendEndpointsListResult, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

            // listByFrontDoorNextResults retrieves the next set of results, if any.
            func (client FrontendEndpointsClient) listByFrontDoorNextResults(ctx context.Context, lastResults FrontendEndpointsListResult) (result FrontendEndpointsListResult, err error) {
            req, err := lastResults.frontendEndpointsListResultPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "frontdoor.FrontendEndpointsClient", "listByFrontDoorNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListByFrontDoorSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "frontdoor.FrontendEndpointsClient", "listByFrontDoorNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListByFrontDoorResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "frontdoor.FrontendEndpointsClient", "listByFrontDoorNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

    // ListByFrontDoorComplete enumerates all values, automatically crossing page boundaries as required.
    func (client FrontendEndpointsClient) ListByFrontDoorComplete(ctx context.Context, resourceGroupName string, frontDoorName string) (result FrontendEndpointsListResultIterator, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/FrontendEndpointsClient.ListByFrontDoor")
            defer func() {
                sc := -1
                if result.Response().Response.Response != nil {
                    sc = result.page.Response().Response.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
     }
        result.page, err = client.ListByFrontDoor(ctx, resourceGroupName, frontDoorName)
                return
        }

