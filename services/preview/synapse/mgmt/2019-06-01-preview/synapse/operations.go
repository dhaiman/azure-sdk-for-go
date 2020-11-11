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

// OperationsClient is the azure Synapse Analytics Management Client
type OperationsClient struct {
    BaseClient
}
// NewOperationsClient creates an instance of the OperationsClient client.
func NewOperationsClient(subscriptionID string) OperationsClient {
    return NewOperationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewOperationsClientWithBaseURI creates an instance of the OperationsClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
    func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
        return OperationsClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// CheckNameAvailability check whether a workspace name is available
    // Parameters:
        // request - the check request
func (client OperationsClient) CheckNameAvailability(ctx context.Context, request CheckNameAvailabilityRequest) (result CheckNameAvailabilityResponse, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/OperationsClient.CheckNameAvailability")
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
         Constraints: []validation.Constraint{	{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil }}}}); err != nil {
        return result, validation.NewError("synapse.OperationsClient", "CheckNameAvailability", err.Error())
        }

        req, err := client.CheckNameAvailabilityPreparer(ctx, request)
    if err != nil {
    err = autorest.NewErrorWithError(err, "synapse.OperationsClient", "CheckNameAvailability", nil , "Failure preparing request")
    return
    }

        resp, err := client.CheckNameAvailabilitySender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "synapse.OperationsClient", "CheckNameAvailability", resp, "Failure sending request")
        return
        }

        result, err = client.CheckNameAvailabilityResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "synapse.OperationsClient", "CheckNameAvailability", resp, "Failure responding to request")
        }

    return
}

    // CheckNameAvailabilityPreparer prepares the CheckNameAvailability request.
    func (client OperationsClient) CheckNameAvailabilityPreparer(ctx context.Context, request CheckNameAvailabilityRequest) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2019-06-01-preview"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsContentType("application/json; charset=utf-8"),
autorest.AsPost(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Synapse/checkNameAvailability",pathParameters),
autorest.WithJSON(request),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CheckNameAvailabilitySender sends the CheckNameAvailability request. The method will close the
    // http.Response Body if it receives an error.
    func (client OperationsClient) CheckNameAvailabilitySender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

    // CheckNameAvailabilityResponder handles the response to the CheckNameAvailability request. The method always
    // closes the http.Response Body.
    func (client OperationsClient) CheckNameAvailabilityResponder(resp *http.Response) (result CheckNameAvailabilityResponse, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// GetAzureAsyncHeaderResult get the status of an operation
    // Parameters:
        // resourceGroupName - the name of the resource group. The name is case insensitive.
        // workspaceName - the name of the workspace
        // operationID - operation ID
func (client OperationsClient) GetAzureAsyncHeaderResult(ctx context.Context, resourceGroupName string, workspaceName string, operationID string) (result OperationResource, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/OperationsClient.GetAzureAsyncHeaderResult")
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
         Constraints: []validation.Constraint{	{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil }}},
        { TargetValue: resourceGroupName,
         Constraints: []validation.Constraint{	{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil },
        	{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil },
        	{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil }}}}); err != nil {
        return result, validation.NewError("synapse.OperationsClient", "GetAzureAsyncHeaderResult", err.Error())
        }

        req, err := client.GetAzureAsyncHeaderResultPreparer(ctx, resourceGroupName, workspaceName, operationID)
    if err != nil {
    err = autorest.NewErrorWithError(err, "synapse.OperationsClient", "GetAzureAsyncHeaderResult", nil , "Failure preparing request")
    return
    }

        resp, err := client.GetAzureAsyncHeaderResultSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "synapse.OperationsClient", "GetAzureAsyncHeaderResult", resp, "Failure sending request")
        return
        }

        result, err = client.GetAzureAsyncHeaderResultResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "synapse.OperationsClient", "GetAzureAsyncHeaderResult", resp, "Failure responding to request")
        }

    return
}

    // GetAzureAsyncHeaderResultPreparer prepares the GetAzureAsyncHeaderResult request.
    func (client OperationsClient) GetAzureAsyncHeaderResultPreparer(ctx context.Context, resourceGroupName string, workspaceName string, operationID string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "operationId": autorest.Encode("path",operationID),
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
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/operationStatuses/{operationId}",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetAzureAsyncHeaderResultSender sends the GetAzureAsyncHeaderResult request. The method will close the
    // http.Response Body if it receives an error.
    func (client OperationsClient) GetAzureAsyncHeaderResultSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

    // GetAzureAsyncHeaderResultResponder handles the response to the GetAzureAsyncHeaderResult request. The method always
    // closes the http.Response Body.
    func (client OperationsClient) GetAzureAsyncHeaderResultResponder(resp *http.Response) (result OperationResource, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusNotFound),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// GetLocationHeaderResult get the result of an operation
    // Parameters:
        // resourceGroupName - the name of the resource group. The name is case insensitive.
        // workspaceName - the name of the workspace
        // operationID - operation ID
func (client OperationsClient) GetLocationHeaderResult(ctx context.Context, resourceGroupName string, workspaceName string, operationID string) (result autorest.Response, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/OperationsClient.GetLocationHeaderResult")
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
         Constraints: []validation.Constraint{	{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil }}},
        { TargetValue: resourceGroupName,
         Constraints: []validation.Constraint{	{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil },
        	{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil },
        	{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil }}}}); err != nil {
        return result, validation.NewError("synapse.OperationsClient", "GetLocationHeaderResult", err.Error())
        }

        req, err := client.GetLocationHeaderResultPreparer(ctx, resourceGroupName, workspaceName, operationID)
    if err != nil {
    err = autorest.NewErrorWithError(err, "synapse.OperationsClient", "GetLocationHeaderResult", nil , "Failure preparing request")
    return
    }

        resp, err := client.GetLocationHeaderResultSender(req)
        if err != nil {
        result.Response = resp
        err = autorest.NewErrorWithError(err, "synapse.OperationsClient", "GetLocationHeaderResult", resp, "Failure sending request")
        return
        }

        result, err = client.GetLocationHeaderResultResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "synapse.OperationsClient", "GetLocationHeaderResult", resp, "Failure responding to request")
        }

    return
}

    // GetLocationHeaderResultPreparer prepares the GetLocationHeaderResult request.
    func (client OperationsClient) GetLocationHeaderResultPreparer(ctx context.Context, resourceGroupName string, workspaceName string, operationID string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "operationId": autorest.Encode("path",operationID),
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
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/operationResults/{operationId}",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetLocationHeaderResultSender sends the GetLocationHeaderResult request. The method will close the
    // http.Response Body if it receives an error.
    func (client OperationsClient) GetLocationHeaderResultSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

    // GetLocationHeaderResultResponder handles the response to the GetLocationHeaderResult request. The method always
    // closes the http.Response Body.
    func (client OperationsClient) GetLocationHeaderResultResponder(resp *http.Response) (result autorest.Response, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusCreated,http.StatusAccepted,http.StatusNoContent),
            autorest.ByClosing())
            result.Response = resp
            return
    }

// List get all available operations
func (client OperationsClient) List(ctx context.Context) (result ListAvailableRpOperation, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/OperationsClient.List")
        defer func() {
            sc := -1
        if result.Response.Response != nil {
        sc = result.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.ListPreparer(ctx)
    if err != nil {
    err = autorest.NewErrorWithError(err, "synapse.OperationsClient", "List", nil , "Failure preparing request")
    return
    }

        resp, err := client.ListSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "synapse.OperationsClient", "List", resp, "Failure sending request")
        return
        }

        result, err = client.ListResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "synapse.OperationsClient", "List", resp, "Failure responding to request")
        }

    return
}

    // ListPreparer prepares the List request.
    func (client OperationsClient) ListPreparer(ctx context.Context) (*http.Request, error) {
    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPath("/providers/Microsoft.Synapse/operations"))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListSender sends the List request. The method will close the
    // http.Response Body if it receives an error.
    func (client OperationsClient) ListSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            }

    // ListResponder handles the response to the List request. The method always
    // closes the http.Response Body.
    func (client OperationsClient) ListResponder(resp *http.Response) (result ListAvailableRpOperation, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result.Value),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

