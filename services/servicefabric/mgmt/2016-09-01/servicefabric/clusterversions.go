package servicefabric

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

// ClusterVersionsClient is the service Fabric Management Client
type ClusterVersionsClient struct {
    BaseClient
}
// NewClusterVersionsClient creates an instance of the ClusterVersionsClient client.
func NewClusterVersionsClient(subscriptionID string) ClusterVersionsClient {
    return NewClusterVersionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewClusterVersionsClientWithBaseURI creates an instance of the ClusterVersionsClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
    func NewClusterVersionsClientWithBaseURI(baseURI string, subscriptionID string) ClusterVersionsClient {
        return ClusterVersionsClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// Get get cluster code versions by environment and version
    // Parameters:
        // location - the location for the cluster code versions, this is different from cluster location
        // environment - cluster operating system, the default means all
        // clusterVersion - the cluster code version
func (client ClusterVersionsClient) Get(ctx context.Context, location string, environment string, clusterVersion string) (result ClusterCodeVersionsResult, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ClusterVersionsClient.Get")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.GetPreparer(ctx, location, environment, clusterVersion)
    if err != nil {
    err = autorest.NewErrorWithError(err, "servicefabric.ClusterVersionsClient", "Get", nil , "Failure preparing request")
    return
    }

            resp, err := client.GetSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "servicefabric.ClusterVersionsClient", "Get", resp, "Failure sending request")
            return
            }

            result, err = client.GetResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "servicefabric.ClusterVersionsClient", "Get", resp, "Failure responding to request")
            }

    return
    }

    // GetPreparer prepares the Get request.
    func (client ClusterVersionsClient) GetPreparer(ctx context.Context, location string, environment string, clusterVersion string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "clusterVersion": autorest.Encode("path",clusterVersion),
            "environment": autorest.Encode("path",environment),
            "location": autorest.Encode("path",location),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2016-09-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/locations/{location}/environments/{environment}/clusterVersions/{clusterVersion}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client ClusterVersionsClient) GetSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ClusterVersionsClient) GetResponder(resp *http.Response) (result ClusterCodeVersionsResult, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// List list cluster code versions by location
    // Parameters:
        // location - the location for the cluster code versions, this is different from cluster location
func (client ClusterVersionsClient) List(ctx context.Context, location string) (result ClusterCodeVersionsListResultPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ClusterVersionsClient.List")
        defer func() {
            sc := -1
            if result.ccvlr.Response.Response != nil {
                sc = result.ccvlr.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
                result.fn = client.listNextResults
    req, err := client.ListPreparer(ctx, location)
    if err != nil {
    err = autorest.NewErrorWithError(err, "servicefabric.ClusterVersionsClient", "List", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListSender(req)
            if err != nil {
            result.ccvlr.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "servicefabric.ClusterVersionsClient", "List", resp, "Failure sending request")
            return
            }

            result.ccvlr, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "servicefabric.ClusterVersionsClient", "List", resp, "Failure responding to request")
            }

    return
    }

    // ListPreparer prepares the List request.
    func (client ClusterVersionsClient) ListPreparer(ctx context.Context, location string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "location": autorest.Encode("path",location),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2016-09-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/locations/{location}/clusterVersions",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListSender sends the List request. The method will close the
    // http.Response Body if it receives an error.
    func (client ClusterVersionsClient) ListSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ClusterVersionsClient) ListResponder(resp *http.Response) (result ClusterCodeVersionsListResult, err error) {
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
            func (client ClusterVersionsClient) listNextResults(ctx context.Context, lastResults ClusterCodeVersionsListResult) (result ClusterCodeVersionsListResult, err error) {
            req, err := lastResults.clusterCodeVersionsListResultPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "servicefabric.ClusterVersionsClient", "listNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "servicefabric.ClusterVersionsClient", "listNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "servicefabric.ClusterVersionsClient", "listNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

    // ListComplete enumerates all values, automatically crossing page boundaries as required.
    func (client ClusterVersionsClient) ListComplete(ctx context.Context, location string) (result ClusterCodeVersionsListResultIterator, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/ClusterVersionsClient.List")
            defer func() {
                sc := -1
                if result.Response().Response.Response != nil {
                    sc = result.page.Response().Response.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
     }
        result.page, err = client.List(ctx, location)
                return
        }

// ListByEnvironment list cluster code versions by environment
    // Parameters:
        // location - the location for the cluster code versions, this is different from cluster location
        // environment - cluster operating system, the default means all
func (client ClusterVersionsClient) ListByEnvironment(ctx context.Context, location string, environment string) (result ClusterCodeVersionsListResultPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ClusterVersionsClient.ListByEnvironment")
        defer func() {
            sc := -1
            if result.ccvlr.Response.Response != nil {
                sc = result.ccvlr.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
                result.fn = client.listByEnvironmentNextResults
    req, err := client.ListByEnvironmentPreparer(ctx, location, environment)
    if err != nil {
    err = autorest.NewErrorWithError(err, "servicefabric.ClusterVersionsClient", "ListByEnvironment", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListByEnvironmentSender(req)
            if err != nil {
            result.ccvlr.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "servicefabric.ClusterVersionsClient", "ListByEnvironment", resp, "Failure sending request")
            return
            }

            result.ccvlr, err = client.ListByEnvironmentResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "servicefabric.ClusterVersionsClient", "ListByEnvironment", resp, "Failure responding to request")
            }

    return
    }

    // ListByEnvironmentPreparer prepares the ListByEnvironment request.
    func (client ClusterVersionsClient) ListByEnvironmentPreparer(ctx context.Context, location string, environment string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "environment": autorest.Encode("path",environment),
            "location": autorest.Encode("path",location),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2016-09-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/locations/{location}/environments/{environment}/clusterVersions",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListByEnvironmentSender sends the ListByEnvironment request. The method will close the
    // http.Response Body if it receives an error.
    func (client ClusterVersionsClient) ListByEnvironmentSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// ListByEnvironmentResponder handles the response to the ListByEnvironment request. The method always
// closes the http.Response Body.
func (client ClusterVersionsClient) ListByEnvironmentResponder(resp *http.Response) (result ClusterCodeVersionsListResult, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

            // listByEnvironmentNextResults retrieves the next set of results, if any.
            func (client ClusterVersionsClient) listByEnvironmentNextResults(ctx context.Context, lastResults ClusterCodeVersionsListResult) (result ClusterCodeVersionsListResult, err error) {
            req, err := lastResults.clusterCodeVersionsListResultPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "servicefabric.ClusterVersionsClient", "listByEnvironmentNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListByEnvironmentSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "servicefabric.ClusterVersionsClient", "listByEnvironmentNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListByEnvironmentResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "servicefabric.ClusterVersionsClient", "listByEnvironmentNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

    // ListByEnvironmentComplete enumerates all values, automatically crossing page boundaries as required.
    func (client ClusterVersionsClient) ListByEnvironmentComplete(ctx context.Context, location string, environment string) (result ClusterCodeVersionsListResultIterator, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/ClusterVersionsClient.ListByEnvironment")
            defer func() {
                sc := -1
                if result.Response().Response.Response != nil {
                    sc = result.page.Response().Response.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
     }
        result.page, err = client.ListByEnvironment(ctx, location, environment)
                return
        }

// ListByVersion list cluster code versions by version
    // Parameters:
        // location - the location for the cluster code versions, this is different from cluster location
        // clusterVersion - the cluster code version
func (client ClusterVersionsClient) ListByVersion(ctx context.Context, location string, clusterVersion string) (result ClusterCodeVersionsListResultPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ClusterVersionsClient.ListByVersion")
        defer func() {
            sc := -1
            if result.ccvlr.Response.Response != nil {
                sc = result.ccvlr.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
                result.fn = client.listByVersionNextResults
    req, err := client.ListByVersionPreparer(ctx, location, clusterVersion)
    if err != nil {
    err = autorest.NewErrorWithError(err, "servicefabric.ClusterVersionsClient", "ListByVersion", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListByVersionSender(req)
            if err != nil {
            result.ccvlr.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "servicefabric.ClusterVersionsClient", "ListByVersion", resp, "Failure sending request")
            return
            }

            result.ccvlr, err = client.ListByVersionResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "servicefabric.ClusterVersionsClient", "ListByVersion", resp, "Failure responding to request")
            }

    return
    }

    // ListByVersionPreparer prepares the ListByVersion request.
    func (client ClusterVersionsClient) ListByVersionPreparer(ctx context.Context, location string, clusterVersion string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "clusterVersion": autorest.Encode("path",clusterVersion),
            "location": autorest.Encode("path",location),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2016-09-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/locations/{location}/clusterVersions/{clusterVersion}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListByVersionSender sends the ListByVersion request. The method will close the
    // http.Response Body if it receives an error.
    func (client ClusterVersionsClient) ListByVersionSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// ListByVersionResponder handles the response to the ListByVersion request. The method always
// closes the http.Response Body.
func (client ClusterVersionsClient) ListByVersionResponder(resp *http.Response) (result ClusterCodeVersionsListResult, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

            // listByVersionNextResults retrieves the next set of results, if any.
            func (client ClusterVersionsClient) listByVersionNextResults(ctx context.Context, lastResults ClusterCodeVersionsListResult) (result ClusterCodeVersionsListResult, err error) {
            req, err := lastResults.clusterCodeVersionsListResultPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "servicefabric.ClusterVersionsClient", "listByVersionNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListByVersionSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "servicefabric.ClusterVersionsClient", "listByVersionNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListByVersionResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "servicefabric.ClusterVersionsClient", "listByVersionNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

    // ListByVersionComplete enumerates all values, automatically crossing page boundaries as required.
    func (client ClusterVersionsClient) ListByVersionComplete(ctx context.Context, location string, clusterVersion string) (result ClusterCodeVersionsListResultIterator, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/ClusterVersionsClient.ListByVersion")
            defer func() {
                sc := -1
                if result.Response().Response.Response != nil {
                    sc = result.page.Response().Response.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
     }
        result.page, err = client.ListByVersion(ctx, location, clusterVersion)
                return
        }

