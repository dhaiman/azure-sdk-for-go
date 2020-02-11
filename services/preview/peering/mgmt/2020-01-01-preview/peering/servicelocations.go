package peering

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

// ServiceLocationsClient is the peering Client
type ServiceLocationsClient struct {
    BaseClient
}
// NewServiceLocationsClient creates an instance of the ServiceLocationsClient client.
func NewServiceLocationsClient(subscriptionID string) ServiceLocationsClient {
    return NewServiceLocationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewServiceLocationsClientWithBaseURI creates an instance of the ServiceLocationsClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
    func NewServiceLocationsClientWithBaseURI(baseURI string, subscriptionID string) ServiceLocationsClient {
        return ServiceLocationsClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// List lists all of the available locations for peering service.
    // Parameters:
        // country - the country of interest, in which the locations are to be present.
func (client ServiceLocationsClient) List(ctx context.Context, country string) (result ServiceLocationListResultPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ServiceLocationsClient.List")
        defer func() {
            sc := -1
            if result.sllr.Response.Response != nil {
                sc = result.sllr.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
                result.fn = client.listNextResults
    req, err := client.ListPreparer(ctx, country)
    if err != nil {
    err = autorest.NewErrorWithError(err, "peering.ServiceLocationsClient", "List", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListSender(req)
            if err != nil {
            result.sllr.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "peering.ServiceLocationsClient", "List", resp, "Failure sending request")
            return
            }

            result.sllr, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "peering.ServiceLocationsClient", "List", resp, "Failure responding to request")
            }

    return
    }

    // ListPreparer prepares the List request.
    func (client ServiceLocationsClient) ListPreparer(ctx context.Context, country string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2020-01-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }
            if len(country) > 0 {
            queryParameters["country"] = autorest.Encode("query",country)
            }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Peering/peeringServiceLocations",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListSender sends the List request. The method will close the
    // http.Response Body if it receives an error.
    func (client ServiceLocationsClient) ListSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ServiceLocationsClient) ListResponder(resp *http.Response) (result ServiceLocationListResult, err error) {
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
            func (client ServiceLocationsClient) listNextResults(ctx context.Context, lastResults ServiceLocationListResult) (result ServiceLocationListResult, err error) {
            req, err := lastResults.serviceLocationListResultPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "peering.ServiceLocationsClient", "listNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "peering.ServiceLocationsClient", "listNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "peering.ServiceLocationsClient", "listNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

    // ListComplete enumerates all values, automatically crossing page boundaries as required.
    func (client ServiceLocationsClient) ListComplete(ctx context.Context, country string) (result ServiceLocationListResultIterator, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/ServiceLocationsClient.List")
            defer func() {
                sc := -1
                if result.Response().Response.Response != nil {
                    sc = result.page.Response().Response.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
     }
        result.page, err = client.List(ctx, country)
                return
        }

