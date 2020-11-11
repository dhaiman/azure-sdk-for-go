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

// SecureScoresClient is the API spec for Microsoft.Security (Azure Security Center) resource provider
type SecureScoresClient struct {
    BaseClient
}
// NewSecureScoresClient creates an instance of the SecureScoresClient client.
func NewSecureScoresClient(subscriptionID string, ascLocation string) SecureScoresClient {
    return NewSecureScoresClientWithBaseURI(DefaultBaseURI, subscriptionID, ascLocation)
}

// NewSecureScoresClientWithBaseURI creates an instance of the SecureScoresClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
    func NewSecureScoresClientWithBaseURI(baseURI string, subscriptionID string, ascLocation string) SecureScoresClient {
        return SecureScoresClient{ NewWithBaseURI(baseURI, subscriptionID, ascLocation)}
    }

// Get get secure score for a specific Security Center initiative within your current scope. For the ASC Default
// initiative, use 'ascScore'.
    // Parameters:
        // secureScoreName - the initiative name. For the ASC Default initiative, use 'ascScore' as in the sample
        // request below.
func (client SecureScoresClient) Get(ctx context.Context, secureScoreName string) (result SecureScoreItem, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/SecureScoresClient.Get")
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
        return result, validation.NewError("security.SecureScoresClient", "Get", err.Error())
        }

        req, err := client.GetPreparer(ctx, secureScoreName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "security.SecureScoresClient", "Get", nil , "Failure preparing request")
    return
    }

        resp, err := client.GetSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "security.SecureScoresClient", "Get", resp, "Failure sending request")
        return
        }

        result, err = client.GetResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "security.SecureScoresClient", "Get", resp, "Failure responding to request")
        }

    return
}

    // GetPreparer prepares the Get request.
    func (client SecureScoresClient) GetPreparer(ctx context.Context, secureScoreName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "secureScoreName": autorest.Encode("path",secureScoreName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2020-01-01-preview"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Security/secureScores/{secureScoreName}",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client SecureScoresClient) GetSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

    // GetResponder handles the response to the Get request. The method always
    // closes the http.Response Body.
    func (client SecureScoresClient) GetResponder(resp *http.Response) (result SecureScoreItem, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// List list secure scores for all your Security Center initiatives within your current scope.
func (client SecureScoresClient) List(ctx context.Context) (result SecureScoresListPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/SecureScoresClient.List")
        defer func() {
            sc := -1
        if result.ssl.Response.Response != nil {
        sc = result.ssl.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        if err := validation.Validate([]validation.Validation{
        { TargetValue: client.SubscriptionID,
         Constraints: []validation.Constraint{	{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil }}}}); err != nil {
        return result, validation.NewError("security.SecureScoresClient", "List", err.Error())
        }

            result.fn = client.listNextResults
    req, err := client.ListPreparer(ctx)
    if err != nil {
    err = autorest.NewErrorWithError(err, "security.SecureScoresClient", "List", nil , "Failure preparing request")
    return
    }

        resp, err := client.ListSender(req)
        if err != nil {
        result.ssl.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "security.SecureScoresClient", "List", resp, "Failure sending request")
        return
        }

        result.ssl, err = client.ListResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "security.SecureScoresClient", "List", resp, "Failure responding to request")
        }
            if result.ssl.hasNextLink() && result.ssl.IsEmpty() {
            err = result.NextWithContext(ctx)
            }

    return
}

    // ListPreparer prepares the List request.
    func (client SecureScoresClient) ListPreparer(ctx context.Context) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2020-01-01-preview"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Security/secureScores",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListSender sends the List request. The method will close the
    // http.Response Body if it receives an error.
    func (client SecureScoresClient) ListSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

    // ListResponder handles the response to the List request. The method always
    // closes the http.Response Body.
    func (client SecureScoresClient) ListResponder(resp *http.Response) (result SecureScoresList, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

            // listNextResults retrieves the next set of results, if any.
            func (client SecureScoresClient) listNextResults(ctx context.Context, lastResults SecureScoresList) (result SecureScoresList, err error) {
            req, err := lastResults.secureScoresListPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "security.SecureScoresClient", "listNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "security.SecureScoresClient", "listNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "security.SecureScoresClient", "listNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

            // ListComplete enumerates all values, automatically crossing page boundaries as required.
            func (client SecureScoresClient) ListComplete(ctx context.Context) (result SecureScoresListIterator, err error) {
            if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/SecureScoresClient.List")
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

