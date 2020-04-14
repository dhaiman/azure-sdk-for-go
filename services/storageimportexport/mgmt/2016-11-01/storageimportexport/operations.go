package storageimportexport

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

// OperationsClient is the the Storage Import/Export Resource Provider API.
type OperationsClient struct {
    BaseClient
}
// NewOperationsClient creates an instance of the OperationsClient client.
func NewOperationsClient(subscriptionID string, acceptLanguage string) OperationsClient {
    return NewOperationsClientWithBaseURI(DefaultBaseURI, subscriptionID, acceptLanguage)
}

// NewOperationsClientWithBaseURI creates an instance of the OperationsClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
    func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string, acceptLanguage string) OperationsClient {
        return OperationsClient{ NewWithBaseURI(baseURI, subscriptionID, acceptLanguage)}
    }

// List returns the list of operations supported by the import/export resource provider.
func (client OperationsClient) List(ctx context.Context) (result ListOperationsResponse, err error) {
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
    err = autorest.NewErrorWithError(err, "storageimportexport.OperationsClient", "List", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "storageimportexport.OperationsClient", "List", resp, "Failure sending request")
            return
            }

            result, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "storageimportexport.OperationsClient", "List", resp, "Failure responding to request")
            }

    return
    }

    // ListPreparer prepares the List request.
    func (client OperationsClient) ListPreparer(ctx context.Context) (*http.Request, error) {
                    const APIVersion = "2016-11-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPath("/providers/Microsoft.ImportExport/operations"),
    autorest.WithQueryParameters(queryParameters))
            if len(client.AcceptLanguage) > 0 {
            preparer = autorest.DecoratePreparer(preparer,
            autorest.WithHeader("Accept-Language",autorest.String(client.AcceptLanguage)))
            }
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListSender sends the List request. The method will close the
    // http.Response Body if it receives an error.
    func (client OperationsClient) ListSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            }

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client OperationsClient) ListResponder(resp *http.Response) (result ListOperationsResponse, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

