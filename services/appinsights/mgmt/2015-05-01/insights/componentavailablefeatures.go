package insights

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

// ComponentAvailableFeaturesClient is the composite Swagger for Application Insights Management Client
type ComponentAvailableFeaturesClient struct {
    BaseClient
}
// NewComponentAvailableFeaturesClient creates an instance of the ComponentAvailableFeaturesClient client.
func NewComponentAvailableFeaturesClient(subscriptionID string) ComponentAvailableFeaturesClient {
    return NewComponentAvailableFeaturesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewComponentAvailableFeaturesClientWithBaseURI creates an instance of the ComponentAvailableFeaturesClient client.
    func NewComponentAvailableFeaturesClientWithBaseURI(baseURI string, subscriptionID string) ComponentAvailableFeaturesClient {
        return ComponentAvailableFeaturesClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// Get returns all available features of the application insights component.
    // Parameters:
        // resourceGroupName - the name of the resource group. The name is case insensitive.
        // resourceName - the name of the Application Insights component resource.
func (client ComponentAvailableFeaturesClient) Get(ctx context.Context, resourceGroupName string, resourceName string) (result ApplicationInsightsComponentAvailableFeatures, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ComponentAvailableFeaturesClient.Get")
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
            { TargetValue: client.SubscriptionID,
             Constraints: []validation.Constraint{	{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil }}}}); err != nil {
            return result, validation.NewError("insights.ComponentAvailableFeaturesClient", "Get", err.Error())
            }

                req, err := client.GetPreparer(ctx, resourceGroupName, resourceName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "insights.ComponentAvailableFeaturesClient", "Get", nil , "Failure preparing request")
    return
    }

            resp, err := client.GetSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "insights.ComponentAvailableFeaturesClient", "Get", resp, "Failure sending request")
            return
            }

            result, err = client.GetResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "insights.ComponentAvailableFeaturesClient", "Get", resp, "Failure responding to request")
            }

    return
    }

    // GetPreparer prepares the Get request.
    func (client ComponentAvailableFeaturesClient) GetPreparer(ctx context.Context, resourceGroupName string, resourceName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "resourceName": autorest.Encode("path",resourceName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2015-05-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/getavailablebillingfeatures",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client ComponentAvailableFeaturesClient) GetSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ComponentAvailableFeaturesClient) GetResponder(resp *http.Response) (result ApplicationInsightsComponentAvailableFeatures, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

