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
)

// InformationProtectionPoliciesClient is the API spec for Microsoft.Security (Azure Security Center) resource provider
type InformationProtectionPoliciesClient struct {
    BaseClient
}
// NewInformationProtectionPoliciesClient creates an instance of the InformationProtectionPoliciesClient client.
func NewInformationProtectionPoliciesClient(subscriptionID string, ascLocation string) InformationProtectionPoliciesClient {
    return NewInformationProtectionPoliciesClientWithBaseURI(DefaultBaseURI, subscriptionID, ascLocation)
}

// NewInformationProtectionPoliciesClientWithBaseURI creates an instance of the InformationProtectionPoliciesClient
// client using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI
// (sovereign clouds, Azure stack).
    func NewInformationProtectionPoliciesClientWithBaseURI(baseURI string, subscriptionID string, ascLocation string) InformationProtectionPoliciesClient {
        return InformationProtectionPoliciesClient{ NewWithBaseURI(baseURI, subscriptionID, ascLocation)}
    }

// CreateOrUpdate details of the information protection policy.
    // Parameters:
        // scope - scope of the query, can be subscription (/subscriptions/0b06d9ea-afe6-4779-bd59-30e5c2d9d13f) or
        // management group (/providers/Microsoft.Management/managementGroups/mgName).
        // informationProtectionPolicyName - name of the information protection policy.
func (client InformationProtectionPoliciesClient) CreateOrUpdate(ctx context.Context, scope string, informationProtectionPolicyName string) (result InformationProtectionPolicy, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/InformationProtectionPoliciesClient.CreateOrUpdate")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.CreateOrUpdatePreparer(ctx, scope, informationProtectionPolicyName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "security.InformationProtectionPoliciesClient", "CreateOrUpdate", nil , "Failure preparing request")
    return
    }

            resp, err := client.CreateOrUpdateSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "security.InformationProtectionPoliciesClient", "CreateOrUpdate", resp, "Failure sending request")
            return
            }

            result, err = client.CreateOrUpdateResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "security.InformationProtectionPoliciesClient", "CreateOrUpdate", resp, "Failure responding to request")
            }

    return
    }

    // CreateOrUpdatePreparer prepares the CreateOrUpdate request.
    func (client InformationProtectionPoliciesClient) CreateOrUpdatePreparer(ctx context.Context, scope string, informationProtectionPolicyName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "informationProtectionPolicyName": autorest.Encode("path",informationProtectionPolicyName),
            "scope": autorest.Encode("path",scope),
            }

                        const APIVersion = "2017-08-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsPut(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/{scope}/providers/Microsoft.Security/informationProtectionPolicies/{informationProtectionPolicyName}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
    // http.Response Body if it receives an error.
    func (client InformationProtectionPoliciesClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            return autorest.SendWithSender(client, req, sd...)
            }

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client InformationProtectionPoliciesClient) CreateOrUpdateResponder(resp *http.Response) (result InformationProtectionPolicy, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusCreated),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// Get details of the information protection policy.
    // Parameters:
        // scope - scope of the query, can be subscription (/subscriptions/0b06d9ea-afe6-4779-bd59-30e5c2d9d13f) or
        // management group (/providers/Microsoft.Management/managementGroups/mgName).
        // informationProtectionPolicyName - name of the information protection policy.
func (client InformationProtectionPoliciesClient) Get(ctx context.Context, scope string, informationProtectionPolicyName string) (result InformationProtectionPolicy, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/InformationProtectionPoliciesClient.Get")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.GetPreparer(ctx, scope, informationProtectionPolicyName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "security.InformationProtectionPoliciesClient", "Get", nil , "Failure preparing request")
    return
    }

            resp, err := client.GetSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "security.InformationProtectionPoliciesClient", "Get", resp, "Failure sending request")
            return
            }

            result, err = client.GetResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "security.InformationProtectionPoliciesClient", "Get", resp, "Failure responding to request")
            }

    return
    }

    // GetPreparer prepares the Get request.
    func (client InformationProtectionPoliciesClient) GetPreparer(ctx context.Context, scope string, informationProtectionPolicyName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "informationProtectionPolicyName": autorest.Encode("path",informationProtectionPolicyName),
            "scope": autorest.Encode("path",scope),
            }

                        const APIVersion = "2017-08-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/{scope}/providers/Microsoft.Security/informationProtectionPolicies/{informationProtectionPolicyName}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client InformationProtectionPoliciesClient) GetSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            return autorest.SendWithSender(client, req, sd...)
            }

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client InformationProtectionPoliciesClient) GetResponder(resp *http.Response) (result InformationProtectionPolicy, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// List information protection policies of a specific management group.
    // Parameters:
        // scope - scope of the query, can be subscription (/subscriptions/0b06d9ea-afe6-4779-bd59-30e5c2d9d13f) or
        // management group (/providers/Microsoft.Management/managementGroups/mgName).
func (client InformationProtectionPoliciesClient) List(ctx context.Context, scope string) (result InformationProtectionPolicyListPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/InformationProtectionPoliciesClient.List")
        defer func() {
            sc := -1
            if result.ippl.Response.Response != nil {
                sc = result.ippl.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
                result.fn = client.listNextResults
    req, err := client.ListPreparer(ctx, scope)
    if err != nil {
    err = autorest.NewErrorWithError(err, "security.InformationProtectionPoliciesClient", "List", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListSender(req)
            if err != nil {
            result.ippl.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "security.InformationProtectionPoliciesClient", "List", resp, "Failure sending request")
            return
            }

            result.ippl, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "security.InformationProtectionPoliciesClient", "List", resp, "Failure responding to request")
            }

    return
    }

    // ListPreparer prepares the List request.
    func (client InformationProtectionPoliciesClient) ListPreparer(ctx context.Context, scope string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "scope": autorest.Encode("path",scope),
            }

                        const APIVersion = "2017-08-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/{scope}/providers/Microsoft.Security/informationProtectionPolicies",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListSender sends the List request. The method will close the
    // http.Response Body if it receives an error.
    func (client InformationProtectionPoliciesClient) ListSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            return autorest.SendWithSender(client, req, sd...)
            }

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client InformationProtectionPoliciesClient) ListResponder(resp *http.Response) (result InformationProtectionPolicyList, err error) {
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
            func (client InformationProtectionPoliciesClient) listNextResults(ctx context.Context, lastResults InformationProtectionPolicyList) (result InformationProtectionPolicyList, err error) {
            req, err := lastResults.informationProtectionPolicyListPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "security.InformationProtectionPoliciesClient", "listNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "security.InformationProtectionPoliciesClient", "listNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "security.InformationProtectionPoliciesClient", "listNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

    // ListComplete enumerates all values, automatically crossing page boundaries as required.
    func (client InformationProtectionPoliciesClient) ListComplete(ctx context.Context, scope string) (result InformationProtectionPolicyListIterator, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/InformationProtectionPoliciesClient.List")
            defer func() {
                sc := -1
                if result.Response().Response.Response != nil {
                    sc = result.page.Response().Response.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
     }
        result.page, err = client.List(ctx, scope)
                return
        }

