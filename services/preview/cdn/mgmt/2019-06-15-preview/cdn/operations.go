package cdn

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

// OperationsClient is the cdn Management Client
type OperationsClient struct {
    BaseClient
}
// NewOperationsClient creates an instance of the OperationsClient client.
func NewOperationsClient(subscriptionID string, subscriptionID1 string) OperationsClient {
    return NewOperationsClientWithBaseURI(DefaultBaseURI, subscriptionID, subscriptionID1)
}

// NewOperationsClientWithBaseURI creates an instance of the OperationsClient client.
    func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string, subscriptionID1 string) OperationsClient {
        return OperationsClient{ NewWithBaseURI(baseURI, subscriptionID, subscriptionID1)}
    }

// List lists all of the available CDN REST API operations.
func (client OperationsClient) List(ctx context.Context) (result OperationsListResultPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/OperationsClient.List")
        defer func() {
            sc := -1
            if result.olr.Response.Response != nil {
                sc = result.olr.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
                result.fn = client.listNextResults
    req, err := client.ListPreparer(ctx)
    if err != nil {
    err = autorest.NewErrorWithError(err, "cdn.OperationsClient", "List", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListSender(req)
            if err != nil {
            result.olr.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "cdn.OperationsClient", "List", resp, "Failure sending request")
            return
            }

            result.olr, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "cdn.OperationsClient", "List", resp, "Failure responding to request")
            }

    return
    }

    // ListPreparer prepares the List request.
    func (client OperationsClient) ListPreparer(ctx context.Context) (*http.Request, error) {
                    const APIVersion = "2019-06-15-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPath("/providers/Microsoft.Cdn/operations"),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListSender sends the List request. The method will close the
    // http.Response Body if it receives an error.
    func (client OperationsClient) ListSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            return autorest.SendWithSender(client, req, sd...)
            }

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client OperationsClient) ListResponder(resp *http.Response) (result OperationsListResult, err error) {
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
            func (client OperationsClient) listNextResults(ctx context.Context, lastResults OperationsListResult) (result OperationsListResult, err error) {
            req, err := lastResults.operationsListResultPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "cdn.OperationsClient", "listNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "cdn.OperationsClient", "listNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "cdn.OperationsClient", "listNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

    // ListComplete enumerates all values, automatically crossing page boundaries as required.
    func (client OperationsClient) ListComplete(ctx context.Context) (result OperationsListResultIterator, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/OperationsClient.List")
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

