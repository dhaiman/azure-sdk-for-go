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

// PreconfiguredEndpointsClient is the frontDoor Client
type PreconfiguredEndpointsClient struct {
    BaseClient
}
// NewPreconfiguredEndpointsClient creates an instance of the PreconfiguredEndpointsClient client.
func NewPreconfiguredEndpointsClient(subscriptionID string) PreconfiguredEndpointsClient {
    return NewPreconfiguredEndpointsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewPreconfiguredEndpointsClientWithBaseURI creates an instance of the PreconfiguredEndpointsClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
    func NewPreconfiguredEndpointsClientWithBaseURI(baseURI string, subscriptionID string) PreconfiguredEndpointsClient {
        return PreconfiguredEndpointsClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// List sends the list request.
    // Parameters:
        // resourceGroupName - name of the Resource group within the Azure subscription.
        // profileName - the Profile identifier associated with the Tenant and Partner
func (client PreconfiguredEndpointsClient) List(ctx context.Context, resourceGroupName string, profileName string) (result PreconfiguredEndpointListPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/PreconfiguredEndpointsClient.List")
        defer func() {
            sc := -1
            if result.pel.Response.Response != nil {
                sc = result.pel.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: resourceGroupName,
             Constraints: []validation.Constraint{	{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 80, Chain: nil },
            	{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil },
            	{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_\-\(\)\.]*[^\.]$`, Chain: nil }}},
            { TargetValue: profileName,
             Constraints: []validation.Constraint{	{Target: "profileName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_\-\(\)\.]*[^\.]$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("frontdoor.PreconfiguredEndpointsClient", "List", err.Error())
            }

                        result.fn = client.listNextResults
    req, err := client.ListPreparer(ctx, resourceGroupName, profileName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "frontdoor.PreconfiguredEndpointsClient", "List", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListSender(req)
            if err != nil {
            result.pel.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "frontdoor.PreconfiguredEndpointsClient", "List", resp, "Failure sending request")
            return
            }

            result.pel, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "frontdoor.PreconfiguredEndpointsClient", "List", resp, "Failure responding to request")
            }

    return
    }

    // ListPreparer prepares the List request.
    func (client PreconfiguredEndpointsClient) ListPreparer(ctx context.Context, resourceGroupName string, profileName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "profileName": autorest.Encode("path",profileName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2019-11-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/NetworkExperimentProfiles/{profileName}/PreconfiguredEndpoints",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListSender sends the List request. The method will close the
    // http.Response Body if it receives an error.
    func (client PreconfiguredEndpointsClient) ListSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client PreconfiguredEndpointsClient) ListResponder(resp *http.Response) (result PreconfiguredEndpointList, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

            // listNextResults retrieves the next set of results, if any.
            func (client PreconfiguredEndpointsClient) listNextResults(ctx context.Context, lastResults PreconfiguredEndpointList) (result PreconfiguredEndpointList, err error) {
            req, err := lastResults.preconfiguredEndpointListPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "frontdoor.PreconfiguredEndpointsClient", "listNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "frontdoor.PreconfiguredEndpointsClient", "listNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "frontdoor.PreconfiguredEndpointsClient", "listNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

    // ListComplete enumerates all values, automatically crossing page boundaries as required.
    func (client PreconfiguredEndpointsClient) ListComplete(ctx context.Context, resourceGroupName string, profileName string) (result PreconfiguredEndpointListIterator, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/PreconfiguredEndpointsClient.List")
            defer func() {
                sc := -1
                if result.Response().Response.Response != nil {
                    sc = result.page.Response().Response.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
     }
        result.page, err = client.List(ctx, resourceGroupName, profileName)
                return
        }

