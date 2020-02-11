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

// RegulatoryComplianceControlsClient is the API spec for Microsoft.Security (Azure Security Center) resource provider
type RegulatoryComplianceControlsClient struct {
    BaseClient
}
// NewRegulatoryComplianceControlsClient creates an instance of the RegulatoryComplianceControlsClient client.
func NewRegulatoryComplianceControlsClient(subscriptionID string, ascLocation string) RegulatoryComplianceControlsClient {
    return NewRegulatoryComplianceControlsClientWithBaseURI(DefaultBaseURI, subscriptionID, ascLocation)
}

// NewRegulatoryComplianceControlsClientWithBaseURI creates an instance of the RegulatoryComplianceControlsClient
// client using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI
// (sovereign clouds, Azure stack).
    func NewRegulatoryComplianceControlsClientWithBaseURI(baseURI string, subscriptionID string, ascLocation string) RegulatoryComplianceControlsClient {
        return RegulatoryComplianceControlsClient{ NewWithBaseURI(baseURI, subscriptionID, ascLocation)}
    }

// Get selected regulatory compliance control details and state
    // Parameters:
        // regulatoryComplianceStandardName - name of the regulatory compliance standard object
        // regulatoryComplianceControlName - name of the regulatory compliance control object
func (client RegulatoryComplianceControlsClient) Get(ctx context.Context, regulatoryComplianceStandardName string, regulatoryComplianceControlName string) (result RegulatoryComplianceControl, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/RegulatoryComplianceControlsClient.Get")
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
            return result, validation.NewError("security.RegulatoryComplianceControlsClient", "Get", err.Error())
            }

                req, err := client.GetPreparer(ctx, regulatoryComplianceStandardName, regulatoryComplianceControlName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "security.RegulatoryComplianceControlsClient", "Get", nil , "Failure preparing request")
    return
    }

            resp, err := client.GetSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "security.RegulatoryComplianceControlsClient", "Get", resp, "Failure sending request")
            return
            }

            result, err = client.GetResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "security.RegulatoryComplianceControlsClient", "Get", resp, "Failure responding to request")
            }

    return
    }

    // GetPreparer prepares the Get request.
    func (client RegulatoryComplianceControlsClient) GetPreparer(ctx context.Context, regulatoryComplianceStandardName string, regulatoryComplianceControlName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "regulatoryComplianceControlName": autorest.Encode("path",regulatoryComplianceControlName),
            "regulatoryComplianceStandardName": autorest.Encode("path",regulatoryComplianceStandardName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2019-01-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Security/regulatoryComplianceStandards/{regulatoryComplianceStandardName}/regulatoryComplianceControls/{regulatoryComplianceControlName}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client RegulatoryComplianceControlsClient) GetSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client RegulatoryComplianceControlsClient) GetResponder(resp *http.Response) (result RegulatoryComplianceControl, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// List all supported regulatory compliance controls details and state for selected standard
    // Parameters:
        // regulatoryComplianceStandardName - name of the regulatory compliance standard object
        // filter - oData filter. Optional.
func (client RegulatoryComplianceControlsClient) List(ctx context.Context, regulatoryComplianceStandardName string, filter string) (result RegulatoryComplianceControlListPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/RegulatoryComplianceControlsClient.List")
        defer func() {
            sc := -1
            if result.rccl.Response.Response != nil {
                sc = result.rccl.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: client.SubscriptionID,
             Constraints: []validation.Constraint{	{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("security.RegulatoryComplianceControlsClient", "List", err.Error())
            }

                        result.fn = client.listNextResults
    req, err := client.ListPreparer(ctx, regulatoryComplianceStandardName, filter)
    if err != nil {
    err = autorest.NewErrorWithError(err, "security.RegulatoryComplianceControlsClient", "List", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListSender(req)
            if err != nil {
            result.rccl.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "security.RegulatoryComplianceControlsClient", "List", resp, "Failure sending request")
            return
            }

            result.rccl, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "security.RegulatoryComplianceControlsClient", "List", resp, "Failure responding to request")
            }

    return
    }

    // ListPreparer prepares the List request.
    func (client RegulatoryComplianceControlsClient) ListPreparer(ctx context.Context, regulatoryComplianceStandardName string, filter string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "regulatoryComplianceStandardName": autorest.Encode("path",regulatoryComplianceStandardName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2019-01-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }
            if len(filter) > 0 {
            queryParameters["$filter"] = autorest.Encode("query",filter)
            }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Security/regulatoryComplianceStandards/{regulatoryComplianceStandardName}/regulatoryComplianceControls",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListSender sends the List request. The method will close the
    // http.Response Body if it receives an error.
    func (client RegulatoryComplianceControlsClient) ListSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client RegulatoryComplianceControlsClient) ListResponder(resp *http.Response) (result RegulatoryComplianceControlList, err error) {
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
            func (client RegulatoryComplianceControlsClient) listNextResults(ctx context.Context, lastResults RegulatoryComplianceControlList) (result RegulatoryComplianceControlList, err error) {
            req, err := lastResults.regulatoryComplianceControlListPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "security.RegulatoryComplianceControlsClient", "listNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "security.RegulatoryComplianceControlsClient", "listNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "security.RegulatoryComplianceControlsClient", "listNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

    // ListComplete enumerates all values, automatically crossing page boundaries as required.
    func (client RegulatoryComplianceControlsClient) ListComplete(ctx context.Context, regulatoryComplianceStandardName string, filter string) (result RegulatoryComplianceControlListIterator, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/RegulatoryComplianceControlsClient.List")
            defer func() {
                sc := -1
                if result.Response().Response.Response != nil {
                    sc = result.page.Response().Response.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
     }
        result.page, err = client.List(ctx, regulatoryComplianceStandardName, filter)
                return
        }

