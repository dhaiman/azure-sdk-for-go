package media

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

// LocationsClient is the client for the Locations methods of the Media service.
type LocationsClient struct {
    BaseClient
}
// NewLocationsClient creates an instance of the LocationsClient client.
func NewLocationsClient(subscriptionID string) LocationsClient {
    return NewLocationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewLocationsClientWithBaseURI creates an instance of the LocationsClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
    func NewLocationsClientWithBaseURI(baseURI string, subscriptionID string) LocationsClient {
        return LocationsClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// CheckNameAvailability checks whether the Media Service resource name is available.
    // Parameters:
        // parameters - the request parameters
func (client LocationsClient) CheckNameAvailability(ctx context.Context, locationName string, parameters CheckNameAvailabilityInput) (result EntityNameAvailabilityCheckOutput, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/LocationsClient.CheckNameAvailability")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.CheckNameAvailabilityPreparer(ctx, locationName, parameters)
    if err != nil {
    err = autorest.NewErrorWithError(err, "media.LocationsClient", "CheckNameAvailability", nil , "Failure preparing request")
    return
    }

            resp, err := client.CheckNameAvailabilitySender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "media.LocationsClient", "CheckNameAvailability", resp, "Failure sending request")
            return
            }

            result, err = client.CheckNameAvailabilityResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "media.LocationsClient", "CheckNameAvailability", resp, "Failure responding to request")
            }

    return
    }

    // CheckNameAvailabilityPreparer prepares the CheckNameAvailability request.
    func (client LocationsClient) CheckNameAvailabilityPreparer(ctx context.Context, locationName string, parameters CheckNameAvailabilityInput) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "locationName": autorest.Encode("path",locationName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2018-03-30-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPost(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Media/locations/{locationName}/checkNameAvailability",pathParameters),
    autorest.WithJSON(parameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CheckNameAvailabilitySender sends the CheckNameAvailability request. The method will close the
    // http.Response Body if it receives an error.
    func (client LocationsClient) CheckNameAvailabilitySender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

// CheckNameAvailabilityResponder handles the response to the CheckNameAvailability request. The method always
// closes the http.Response Body.
func (client LocationsClient) CheckNameAvailabilityResponder(resp *http.Response) (result EntityNameAvailabilityCheckOutput, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

