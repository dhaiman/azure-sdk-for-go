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
)

// DiagnosticSettingsClient is the monitor Management Client
type DiagnosticSettingsClient struct {
    BaseClient
}
// NewDiagnosticSettingsClient creates an instance of the DiagnosticSettingsClient client.
func NewDiagnosticSettingsClient(subscriptionID string) DiagnosticSettingsClient {
    return NewDiagnosticSettingsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewDiagnosticSettingsClientWithBaseURI creates an instance of the DiagnosticSettingsClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
    func NewDiagnosticSettingsClientWithBaseURI(baseURI string, subscriptionID string) DiagnosticSettingsClient {
        return DiagnosticSettingsClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// CreateOrUpdate creates or updates diagnostic settings for the specified resource.
    // Parameters:
        // resourceURI - the identifier of the resource.
        // parameters - parameters supplied to the operation.
        // name - the name of the diagnostic setting.
func (client DiagnosticSettingsClient) CreateOrUpdate(ctx context.Context, resourceURI string, parameters DiagnosticSettingsResource, name string) (result DiagnosticSettingsResource, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/DiagnosticSettingsClient.CreateOrUpdate")
        defer func() {
            sc := -1
        if result.Response.Response != nil {
        sc = result.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.CreateOrUpdatePreparer(ctx, resourceURI, parameters, name)
    if err != nil {
    err = autorest.NewErrorWithError(err, "insights.DiagnosticSettingsClient", "CreateOrUpdate", nil , "Failure preparing request")
    return
    }

        resp, err := client.CreateOrUpdateSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "insights.DiagnosticSettingsClient", "CreateOrUpdate", resp, "Failure sending request")
        return
        }

        result, err = client.CreateOrUpdateResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "insights.DiagnosticSettingsClient", "CreateOrUpdate", resp, "Failure responding to request")
        }

    return
}

    // CreateOrUpdatePreparer prepares the CreateOrUpdate request.
    func (client DiagnosticSettingsClient) CreateOrUpdatePreparer(ctx context.Context, resourceURI string, parameters DiagnosticSettingsResource, name string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "name": autorest.Encode("path",name),
        "resourceUri": resourceURI,
        }

            const APIVersion = "2017-05-01-preview"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsContentType("application/json; charset=utf-8"),
autorest.AsPut(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/{resourceUri}/providers/microsoft.insights/diagnosticSettings/{name}",pathParameters),
autorest.WithJSON(parameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
    // http.Response Body if it receives an error.
    func (client DiagnosticSettingsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            }

    // CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
    // closes the http.Response Body.
    func (client DiagnosticSettingsClient) CreateOrUpdateResponder(resp *http.Response) (result DiagnosticSettingsResource, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// Delete deletes existing diagnostic settings for the specified resource.
    // Parameters:
        // resourceURI - the identifier of the resource.
        // name - the name of the diagnostic setting.
func (client DiagnosticSettingsClient) Delete(ctx context.Context, resourceURI string, name string) (result autorest.Response, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/DiagnosticSettingsClient.Delete")
        defer func() {
            sc := -1
        if result.Response != nil {
        sc = result.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.DeletePreparer(ctx, resourceURI, name)
    if err != nil {
    err = autorest.NewErrorWithError(err, "insights.DiagnosticSettingsClient", "Delete", nil , "Failure preparing request")
    return
    }

        resp, err := client.DeleteSender(req)
        if err != nil {
        result.Response = resp
        err = autorest.NewErrorWithError(err, "insights.DiagnosticSettingsClient", "Delete", resp, "Failure sending request")
        return
        }

        result, err = client.DeleteResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "insights.DiagnosticSettingsClient", "Delete", resp, "Failure responding to request")
        }

    return
}

    // DeletePreparer prepares the Delete request.
    func (client DiagnosticSettingsClient) DeletePreparer(ctx context.Context, resourceURI string, name string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "name": autorest.Encode("path",name),
        "resourceUri": resourceURI,
        }

            const APIVersion = "2017-05-01-preview"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsDelete(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/{resourceUri}/providers/microsoft.insights/diagnosticSettings/{name}",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // DeleteSender sends the Delete request. The method will close the
    // http.Response Body if it receives an error.
    func (client DiagnosticSettingsClient) DeleteSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            }

    // DeleteResponder handles the response to the Delete request. The method always
    // closes the http.Response Body.
    func (client DiagnosticSettingsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusNoContent),
            autorest.ByClosing())
            result.Response = resp
            return
    }

// Get gets the active diagnostic settings for the specified resource.
    // Parameters:
        // resourceURI - the identifier of the resource.
        // name - the name of the diagnostic setting.
func (client DiagnosticSettingsClient) Get(ctx context.Context, resourceURI string, name string) (result DiagnosticSettingsResource, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/DiagnosticSettingsClient.Get")
        defer func() {
            sc := -1
        if result.Response.Response != nil {
        sc = result.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.GetPreparer(ctx, resourceURI, name)
    if err != nil {
    err = autorest.NewErrorWithError(err, "insights.DiagnosticSettingsClient", "Get", nil , "Failure preparing request")
    return
    }

        resp, err := client.GetSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "insights.DiagnosticSettingsClient", "Get", resp, "Failure sending request")
        return
        }

        result, err = client.GetResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "insights.DiagnosticSettingsClient", "Get", resp, "Failure responding to request")
        }

    return
}

    // GetPreparer prepares the Get request.
    func (client DiagnosticSettingsClient) GetPreparer(ctx context.Context, resourceURI string, name string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "name": autorest.Encode("path",name),
        "resourceUri": resourceURI,
        }

            const APIVersion = "2017-05-01-preview"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/{resourceUri}/providers/microsoft.insights/diagnosticSettings/{name}",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client DiagnosticSettingsClient) GetSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            }

    // GetResponder handles the response to the Get request. The method always
    // closes the http.Response Body.
    func (client DiagnosticSettingsClient) GetResponder(resp *http.Response) (result DiagnosticSettingsResource, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// List gets the active diagnostic settings list for the specified resource.
    // Parameters:
        // resourceURI - the identifier of the resource.
func (client DiagnosticSettingsClient) List(ctx context.Context, resourceURI string) (result DiagnosticSettingsResourceCollection, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/DiagnosticSettingsClient.List")
        defer func() {
            sc := -1
        if result.Response.Response != nil {
        sc = result.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.ListPreparer(ctx, resourceURI)
    if err != nil {
    err = autorest.NewErrorWithError(err, "insights.DiagnosticSettingsClient", "List", nil , "Failure preparing request")
    return
    }

        resp, err := client.ListSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "insights.DiagnosticSettingsClient", "List", resp, "Failure sending request")
        return
        }

        result, err = client.ListResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "insights.DiagnosticSettingsClient", "List", resp, "Failure responding to request")
        }

    return
}

    // ListPreparer prepares the List request.
    func (client DiagnosticSettingsClient) ListPreparer(ctx context.Context, resourceURI string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "resourceUri": resourceURI,
        }

            const APIVersion = "2017-05-01-preview"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/{resourceUri}/providers/microsoft.insights/diagnosticSettings",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListSender sends the List request. The method will close the
    // http.Response Body if it receives an error.
    func (client DiagnosticSettingsClient) ListSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            }

    // ListResponder handles the response to the List request. The method always
    // closes the http.Response Body.
    func (client DiagnosticSettingsClient) ListResponder(resp *http.Response) (result DiagnosticSettingsResourceCollection, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

