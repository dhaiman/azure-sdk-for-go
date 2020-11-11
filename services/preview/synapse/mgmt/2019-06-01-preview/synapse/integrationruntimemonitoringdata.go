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

// IntegrationRuntimeMonitoringDataClient is the azure Synapse Analytics Management Client
type IntegrationRuntimeMonitoringDataClient struct {
    BaseClient
}
// NewIntegrationRuntimeMonitoringDataClient creates an instance of the IntegrationRuntimeMonitoringDataClient client.
func NewIntegrationRuntimeMonitoringDataClient(subscriptionID string) IntegrationRuntimeMonitoringDataClient {
    return NewIntegrationRuntimeMonitoringDataClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewIntegrationRuntimeMonitoringDataClientWithBaseURI creates an instance of the
// IntegrationRuntimeMonitoringDataClient client using a custom endpoint.  Use this when interacting with an Azure
// cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
    func NewIntegrationRuntimeMonitoringDataClientWithBaseURI(baseURI string, subscriptionID string) IntegrationRuntimeMonitoringDataClient {
        return IntegrationRuntimeMonitoringDataClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// List get monitoring data for an integration runtime
    // Parameters:
        // resourceGroupName - the name of the resource group. The name is case insensitive.
        // workspaceName - the name of the workspace.
        // integrationRuntimeName - integration runtime name
func (client IntegrationRuntimeMonitoringDataClient) List(ctx context.Context, resourceGroupName string, workspaceName string, integrationRuntimeName string) (result IntegrationRuntimeMonitoringData, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/IntegrationRuntimeMonitoringDataClient.List")
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
        return result, validation.NewError("synapse.IntegrationRuntimeMonitoringDataClient", "List", err.Error())
        }

        req, err := client.ListPreparer(ctx, resourceGroupName, workspaceName, integrationRuntimeName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "synapse.IntegrationRuntimeMonitoringDataClient", "List", nil , "Failure preparing request")
    return
    }

        resp, err := client.ListSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "synapse.IntegrationRuntimeMonitoringDataClient", "List", resp, "Failure sending request")
        return
        }

        result, err = client.ListResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "synapse.IntegrationRuntimeMonitoringDataClient", "List", resp, "Failure responding to request")
        }

    return
}

    // ListPreparer prepares the List request.
    func (client IntegrationRuntimeMonitoringDataClient) ListPreparer(ctx context.Context, resourceGroupName string, workspaceName string, integrationRuntimeName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "integrationRuntimeName": autorest.Encode("path",integrationRuntimeName),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        "workspaceName": autorest.Encode("path",workspaceName),
        }

            const APIVersion = "2019-06-01-preview"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsPost(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/integrationRuntimes/{integrationRuntimeName}/monitoringData",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListSender sends the List request. The method will close the
    // http.Response Body if it receives an error.
    func (client IntegrationRuntimeMonitoringDataClient) ListSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

    // ListResponder handles the response to the List request. The method always
    // closes the http.Response Body.
    func (client IntegrationRuntimeMonitoringDataClient) ListResponder(resp *http.Response) (result IntegrationRuntimeMonitoringData, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

