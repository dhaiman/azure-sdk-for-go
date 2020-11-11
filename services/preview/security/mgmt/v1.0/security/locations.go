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

// LocationsClient is the API spec for Microsoft.Security (Azure Security Center) resource provider
type LocationsClient struct {
    BaseClient
}
// NewLocationsClient creates an instance of the LocationsClient client.
func NewLocationsClient(subscriptionID string, ascLocation string) LocationsClient {
    return NewLocationsClientWithBaseURI(DefaultBaseURI, subscriptionID, ascLocation)
}

// NewLocationsClientWithBaseURI creates an instance of the LocationsClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
    func NewLocationsClientWithBaseURI(baseURI string, subscriptionID string, ascLocation string) LocationsClient {
        return LocationsClient{ NewWithBaseURI(baseURI, subscriptionID, ascLocation)}
    }

// Get details of a specific location
func (client LocationsClient) Get(ctx context.Context) (result AscLocation, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/LocationsClient.Get")
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
         Constraints: []validation.Constraint{	{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil }}}}); err != nil {
        return result, validation.NewError("security.LocationsClient", "Get", err.Error())
        }

        req, err := client.GetPreparer(ctx)
    if err != nil {
    err = autorest.NewErrorWithError(err, "security.LocationsClient", "Get", nil , "Failure preparing request")
    return
    }

        resp, err := client.GetSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "security.LocationsClient", "Get", resp, "Failure sending request")
        return
        }

        result, err = client.GetResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "security.LocationsClient", "Get", resp, "Failure responding to request")
        }

    return
}

    // GetPreparer prepares the Get request.
    func (client LocationsClient) GetPreparer(ctx context.Context) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "ascLocation": autorest.Encode("path",client.AscLocation),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2015-06-01-preview"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Security/locations/{ascLocation}",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client LocationsClient) GetSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

    // GetResponder handles the response to the Get request. The method always
    // closes the http.Response Body.
    func (client LocationsClient) GetResponder(resp *http.Response) (result AscLocation, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// List the location of the responsible ASC of the specific subscription (home region). For each subscription there is
// only one responsible location. The location in the response should be used to read or write other resources in ASC
// according to their ID.
func (client LocationsClient) List(ctx context.Context) (result AscLocationListPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/LocationsClient.List")
        defer func() {
            sc := -1
        if result.all.Response.Response != nil {
        sc = result.all.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        if err := validation.Validate([]validation.Validation{
        { TargetValue: client.SubscriptionID,
         Constraints: []validation.Constraint{	{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil }}}}); err != nil {
        return result, validation.NewError("security.LocationsClient", "List", err.Error())
        }

            result.fn = client.listNextResults
    req, err := client.ListPreparer(ctx)
    if err != nil {
    err = autorest.NewErrorWithError(err, "security.LocationsClient", "List", nil , "Failure preparing request")
    return
    }

        resp, err := client.ListSender(req)
        if err != nil {
        result.all.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "security.LocationsClient", "List", resp, "Failure sending request")
        return
        }

        result.all, err = client.ListResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "security.LocationsClient", "List", resp, "Failure responding to request")
        }
            if result.all.hasNextLink() && result.all.IsEmpty() {
            err = result.NextWithContext(ctx)
            }

    return
}

    // ListPreparer prepares the List request.
    func (client LocationsClient) ListPreparer(ctx context.Context) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2015-06-01-preview"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Security/locations",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListSender sends the List request. The method will close the
    // http.Response Body if it receives an error.
    func (client LocationsClient) ListSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

    // ListResponder handles the response to the List request. The method always
    // closes the http.Response Body.
    func (client LocationsClient) ListResponder(resp *http.Response) (result AscLocationList, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

            // listNextResults retrieves the next set of results, if any.
            func (client LocationsClient) listNextResults(ctx context.Context, lastResults AscLocationList) (result AscLocationList, err error) {
            req, err := lastResults.ascLocationListPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "security.LocationsClient", "listNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "security.LocationsClient", "listNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "security.LocationsClient", "listNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

            // ListComplete enumerates all values, automatically crossing page boundaries as required.
            func (client LocationsClient) ListComplete(ctx context.Context) (result AscLocationListIterator, err error) {
            if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/LocationsClient.List")
            defer func() {
            sc := -1
            if result.Response().Response.Response != nil {
            sc = result.page.Response().Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
            }()
            }
                    result.page, err = client.List(ctx)
                            return
            }

