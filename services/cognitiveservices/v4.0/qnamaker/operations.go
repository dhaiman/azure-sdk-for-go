package qnamaker

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

// OperationsClient is the an API for QnAMaker Service
type OperationsClient struct {
    BaseClient
}
// NewOperationsClient creates an instance of the OperationsClient client.
func NewOperationsClient(endpoint string) OperationsClient {
    return OperationsClient{ New(endpoint)}
}

// GetDetails sends the get details request.
    // Parameters:
        // operationID - operation id.
func (client OperationsClient) GetDetails(ctx context.Context, operationID string) (result Operation, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/OperationsClient.GetDetails")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.GetDetailsPreparer(ctx, operationID)
    if err != nil {
    err = autorest.NewErrorWithError(err, "qnamaker.OperationsClient", "GetDetails", nil , "Failure preparing request")
    return
    }

            resp, err := client.GetDetailsSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "qnamaker.OperationsClient", "GetDetails", resp, "Failure sending request")
            return
            }

            result, err = client.GetDetailsResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "qnamaker.OperationsClient", "GetDetails", resp, "Failure responding to request")
            }

    return
    }

    // GetDetailsPreparer prepares the GetDetails request.
    func (client OperationsClient) GetDetailsPreparer(ctx context.Context, operationID string) (*http.Request, error) {
            urlParameters := map[string]interface{} {
            "Endpoint": client.Endpoint,
            }

            pathParameters := map[string]interface{} {
            "operationId": autorest.Encode("path",operationID),
            }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithCustomBaseURL("{Endpoint}/qnamaker/v4.0", urlParameters),
    autorest.WithPathParameters("/operations/{operationId}",pathParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetDetailsSender sends the GetDetails request. The method will close the
    // http.Response Body if it receives an error.
    func (client OperationsClient) GetDetailsSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            return autorest.SendWithSender(client, req, sd...)
            }

// GetDetailsResponder handles the response to the GetDetails request. The method always
// closes the http.Response Body.
func (client OperationsClient) GetDetailsResponder(resp *http.Response) (result Operation, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

