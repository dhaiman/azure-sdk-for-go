package advisor

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

// SuppressionsClient is the REST APIs for Azure Advisor
type SuppressionsClient struct {
    BaseClient
}
// NewSuppressionsClient creates an instance of the SuppressionsClient client.
func NewSuppressionsClient(subscriptionID string) SuppressionsClient {
    return NewSuppressionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewSuppressionsClientWithBaseURI creates an instance of the SuppressionsClient client.
    func NewSuppressionsClientWithBaseURI(baseURI string, subscriptionID string) SuppressionsClient {
        return SuppressionsClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// Create enables the snoozed or dismissed attribute of a recommendation. The snoozed or dismissed attribute is
// referred to as a suppression. Use this API to create or update the snoozed or dismissed status of a recommendation.
    // Parameters:
        // resourceURI - the fully qualified Azure Resource Manager identifier of the resource to which the
        // recommendation applies.
        // recommendationID - the recommendation ID.
        // name - the name of the suppression.
        // suppressionContract - the snoozed or dismissed attribute; for example, the snooze duration.
func (client SuppressionsClient) Create(ctx context.Context, resourceURI string, recommendationID string, name string, suppressionContract SuppressionContract) (result SuppressionContract, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/SuppressionsClient.Create")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.CreatePreparer(ctx, resourceURI, recommendationID, name, suppressionContract)
    if err != nil {
    err = autorest.NewErrorWithError(err, "advisor.SuppressionsClient", "Create", nil , "Failure preparing request")
    return
    }

            resp, err := client.CreateSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "advisor.SuppressionsClient", "Create", resp, "Failure sending request")
            return
            }

            result, err = client.CreateResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "advisor.SuppressionsClient", "Create", resp, "Failure responding to request")
            }

    return
    }

    // CreatePreparer prepares the Create request.
    func (client SuppressionsClient) CreatePreparer(ctx context.Context, resourceURI string, recommendationID string, name string, suppressionContract SuppressionContract) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "name": autorest.Encode("path",name),
            "recommendationId": autorest.Encode("path",recommendationID),
            "resourceUri": autorest.Encode("path",resourceURI),
            }

                        const APIVersion = "2016-07-12-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPut(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/{resourceUri}/providers/Microsoft.Advisor/recommendations/{recommendationId}/suppressions/{name}",pathParameters),
    autorest.WithJSON(suppressionContract),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CreateSender sends the Create request. The method will close the
    // http.Response Body if it receives an error.
    func (client SuppressionsClient) CreateSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            return autorest.SendWithSender(client, req, sd...)
            }

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client SuppressionsClient) CreateResponder(resp *http.Response) (result SuppressionContract, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// Delete enables the activation of a snoozed or dismissed recommendation. The snoozed or dismissed attribute of a
// recommendation is referred to as a suppression.
    // Parameters:
        // resourceURI - the fully qualified Azure Resource Manager identifier of the resource to which the
        // recommendation applies.
        // recommendationID - the recommendation ID.
        // name - the name of the suppression.
func (client SuppressionsClient) Delete(ctx context.Context, resourceURI string, recommendationID string, name string) (result autorest.Response, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/SuppressionsClient.Delete")
        defer func() {
            sc := -1
            if result.Response != nil {
                sc = result.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.DeletePreparer(ctx, resourceURI, recommendationID, name)
    if err != nil {
    err = autorest.NewErrorWithError(err, "advisor.SuppressionsClient", "Delete", nil , "Failure preparing request")
    return
    }

            resp, err := client.DeleteSender(req)
            if err != nil {
            result.Response = resp
            err = autorest.NewErrorWithError(err, "advisor.SuppressionsClient", "Delete", resp, "Failure sending request")
            return
            }

            result, err = client.DeleteResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "advisor.SuppressionsClient", "Delete", resp, "Failure responding to request")
            }

    return
    }

    // DeletePreparer prepares the Delete request.
    func (client SuppressionsClient) DeletePreparer(ctx context.Context, resourceURI string, recommendationID string, name string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "name": autorest.Encode("path",name),
            "recommendationId": autorest.Encode("path",recommendationID),
            "resourceUri": autorest.Encode("path",resourceURI),
            }

                        const APIVersion = "2016-07-12-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsDelete(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/{resourceUri}/providers/Microsoft.Advisor/recommendations/{recommendationId}/suppressions/{name}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // DeleteSender sends the Delete request. The method will close the
    // http.Response Body if it receives an error.
    func (client SuppressionsClient) DeleteSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            return autorest.SendWithSender(client, req, sd...)
            }

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client SuppressionsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusNoContent),
    autorest.ByClosing())
    result.Response = resp
        return
    }

// Get obtains the details of a suppression.
    // Parameters:
        // resourceURI - the fully qualified Azure Resource Manager identifier of the resource to which the
        // recommendation applies.
        // recommendationID - the recommendation ID.
        // name - the name of the suppression.
func (client SuppressionsClient) Get(ctx context.Context, resourceURI string, recommendationID string, name string) (result SuppressionContract, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/SuppressionsClient.Get")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.GetPreparer(ctx, resourceURI, recommendationID, name)
    if err != nil {
    err = autorest.NewErrorWithError(err, "advisor.SuppressionsClient", "Get", nil , "Failure preparing request")
    return
    }

            resp, err := client.GetSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "advisor.SuppressionsClient", "Get", resp, "Failure sending request")
            return
            }

            result, err = client.GetResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "advisor.SuppressionsClient", "Get", resp, "Failure responding to request")
            }

    return
    }

    // GetPreparer prepares the Get request.
    func (client SuppressionsClient) GetPreparer(ctx context.Context, resourceURI string, recommendationID string, name string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "name": autorest.Encode("path",name),
            "recommendationId": autorest.Encode("path",recommendationID),
            "resourceUri": autorest.Encode("path",resourceURI),
            }

                        const APIVersion = "2016-07-12-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/{resourceUri}/providers/Microsoft.Advisor/recommendations/{recommendationId}/suppressions/{name}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client SuppressionsClient) GetSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            return autorest.SendWithSender(client, req, sd...)
            }

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client SuppressionsClient) GetResponder(resp *http.Response) (result SuppressionContract, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// List retrieves the list of snoozed or dismissed suppressions for a subscription. The snoozed or dismissed attribute
// of a recommendation is referred to as a suppression.
func (client SuppressionsClient) List(ctx context.Context) (result ListSuppressionContract, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/SuppressionsClient.List")
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
    err = autorest.NewErrorWithError(err, "advisor.SuppressionsClient", "List", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "advisor.SuppressionsClient", "List", resp, "Failure sending request")
            return
            }

            result, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "advisor.SuppressionsClient", "List", resp, "Failure responding to request")
            }

    return
    }

    // ListPreparer prepares the List request.
    func (client SuppressionsClient) ListPreparer(ctx context.Context) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2016-07-12-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Advisor/suppressions",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListSender sends the List request. The method will close the
    // http.Response Body if it receives an error.
    func (client SuppressionsClient) ListSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client SuppressionsClient) ListResponder(resp *http.Response) (result ListSuppressionContract, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result.Value),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

