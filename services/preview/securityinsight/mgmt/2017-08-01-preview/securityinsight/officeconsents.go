package securityinsight

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

// OfficeConsentsClient is the API spec for Microsoft.SecurityInsights (Azure Security Insights) resource provider
type OfficeConsentsClient struct {
    BaseClient
}
// NewOfficeConsentsClient creates an instance of the OfficeConsentsClient client.
func NewOfficeConsentsClient(subscriptionID string) OfficeConsentsClient {
    return NewOfficeConsentsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewOfficeConsentsClientWithBaseURI creates an instance of the OfficeConsentsClient client.
    func NewOfficeConsentsClientWithBaseURI(baseURI string, subscriptionID string) OfficeConsentsClient {
        return OfficeConsentsClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// Delete delete the office365 consent.
    // Parameters:
        // resourceGroupName - the name of the resource group within the user's subscription. The name is case
        // insensitive.
        // operationalInsightsResourceProvider - the namespace of workspaces resource provider-
        // Microsoft.OperationalInsights.
        // workspaceName - the name of the workspace.
        // consentID - consent ID
func (client OfficeConsentsClient) Delete(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, consentID string) (result autorest.Response, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/OfficeConsentsClient.Delete")
        defer func() {
            sc := -1
            if result.Response != nil {
                sc = result.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: client.SubscriptionID,
             Constraints: []validation.Constraint{	{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil }}},
            { TargetValue: resourceGroupName,
             Constraints: []validation.Constraint{	{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil },
            	{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil },
            	{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil }}},
            { TargetValue: workspaceName,
             Constraints: []validation.Constraint{	{Target: "workspaceName", Name: validation.MaxLength, Rule: 90, Chain: nil },
            	{Target: "workspaceName", Name: validation.MinLength, Rule: 1, Chain: nil }}}}); err != nil {
            return result, validation.NewError("securityinsight.OfficeConsentsClient", "Delete", err.Error())
            }

                req, err := client.DeletePreparer(ctx, resourceGroupName, operationalInsightsResourceProvider, workspaceName, consentID)
    if err != nil {
    err = autorest.NewErrorWithError(err, "securityinsight.OfficeConsentsClient", "Delete", nil , "Failure preparing request")
    return
    }

            resp, err := client.DeleteSender(req)
            if err != nil {
            result.Response = resp
            err = autorest.NewErrorWithError(err, "securityinsight.OfficeConsentsClient", "Delete", resp, "Failure sending request")
            return
            }

            result, err = client.DeleteResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "securityinsight.OfficeConsentsClient", "Delete", resp, "Failure responding to request")
            }

    return
    }

    // DeletePreparer prepares the Delete request.
    func (client OfficeConsentsClient) DeletePreparer(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, consentID string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "consentId": autorest.Encode("path",consentID),
            "operationalInsightsResourceProvider": autorest.Encode("path",operationalInsightsResourceProvider),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            "workspaceName": autorest.Encode("path",workspaceName),
            }

                        const APIVersion = "2019-01-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsDelete(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{operationalInsightsResourceProvider}/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/officeConsents/{consentId}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // DeleteSender sends the Delete request. The method will close the
    // http.Response Body if it receives an error.
    func (client OfficeConsentsClient) DeleteSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client OfficeConsentsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusNoContent),
    autorest.ByClosing())
    result.Response = resp
        return
    }

// Get gets an office365 consent.
    // Parameters:
        // resourceGroupName - the name of the resource group within the user's subscription. The name is case
        // insensitive.
        // operationalInsightsResourceProvider - the namespace of workspaces resource provider-
        // Microsoft.OperationalInsights.
        // workspaceName - the name of the workspace.
        // consentID - consent ID
func (client OfficeConsentsClient) Get(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, consentID string) (result OfficeConsent, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/OfficeConsentsClient.Get")
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
             Constraints: []validation.Constraint{	{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil }}},
            { TargetValue: resourceGroupName,
             Constraints: []validation.Constraint{	{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil },
            	{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil },
            	{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil }}},
            { TargetValue: workspaceName,
             Constraints: []validation.Constraint{	{Target: "workspaceName", Name: validation.MaxLength, Rule: 90, Chain: nil },
            	{Target: "workspaceName", Name: validation.MinLength, Rule: 1, Chain: nil }}}}); err != nil {
            return result, validation.NewError("securityinsight.OfficeConsentsClient", "Get", err.Error())
            }

                req, err := client.GetPreparer(ctx, resourceGroupName, operationalInsightsResourceProvider, workspaceName, consentID)
    if err != nil {
    err = autorest.NewErrorWithError(err, "securityinsight.OfficeConsentsClient", "Get", nil , "Failure preparing request")
    return
    }

            resp, err := client.GetSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "securityinsight.OfficeConsentsClient", "Get", resp, "Failure sending request")
            return
            }

            result, err = client.GetResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "securityinsight.OfficeConsentsClient", "Get", resp, "Failure responding to request")
            }

    return
    }

    // GetPreparer prepares the Get request.
    func (client OfficeConsentsClient) GetPreparer(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, consentID string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "consentId": autorest.Encode("path",consentID),
            "operationalInsightsResourceProvider": autorest.Encode("path",operationalInsightsResourceProvider),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            "workspaceName": autorest.Encode("path",workspaceName),
            }

                        const APIVersion = "2019-01-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{operationalInsightsResourceProvider}/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/officeConsents/{consentId}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client OfficeConsentsClient) GetSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client OfficeConsentsClient) GetResponder(resp *http.Response) (result OfficeConsent, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// List gets all office365 consents.
    // Parameters:
        // resourceGroupName - the name of the resource group within the user's subscription. The name is case
        // insensitive.
        // operationalInsightsResourceProvider - the namespace of workspaces resource provider-
        // Microsoft.OperationalInsights.
        // workspaceName - the name of the workspace.
func (client OfficeConsentsClient) List(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string) (result OfficeConsentListPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/OfficeConsentsClient.List")
        defer func() {
            sc := -1
            if result.ocl.Response.Response != nil {
                sc = result.ocl.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: client.SubscriptionID,
             Constraints: []validation.Constraint{	{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil }}},
            { TargetValue: resourceGroupName,
             Constraints: []validation.Constraint{	{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil },
            	{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil },
            	{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil }}},
            { TargetValue: workspaceName,
             Constraints: []validation.Constraint{	{Target: "workspaceName", Name: validation.MaxLength, Rule: 90, Chain: nil },
            	{Target: "workspaceName", Name: validation.MinLength, Rule: 1, Chain: nil }}}}); err != nil {
            return result, validation.NewError("securityinsight.OfficeConsentsClient", "List", err.Error())
            }

                        result.fn = client.listNextResults
    req, err := client.ListPreparer(ctx, resourceGroupName, operationalInsightsResourceProvider, workspaceName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "securityinsight.OfficeConsentsClient", "List", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListSender(req)
            if err != nil {
            result.ocl.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "securityinsight.OfficeConsentsClient", "List", resp, "Failure sending request")
            return
            }

            result.ocl, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "securityinsight.OfficeConsentsClient", "List", resp, "Failure responding to request")
            }

    return
    }

    // ListPreparer prepares the List request.
    func (client OfficeConsentsClient) ListPreparer(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "operationalInsightsResourceProvider": autorest.Encode("path",operationalInsightsResourceProvider),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            "workspaceName": autorest.Encode("path",workspaceName),
            }

                        const APIVersion = "2019-01-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{operationalInsightsResourceProvider}/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/officeConsents",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListSender sends the List request. The method will close the
    // http.Response Body if it receives an error.
    func (client OfficeConsentsClient) ListSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client OfficeConsentsClient) ListResponder(resp *http.Response) (result OfficeConsentList, err error) {
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
            func (client OfficeConsentsClient) listNextResults(ctx context.Context, lastResults OfficeConsentList) (result OfficeConsentList, err error) {
            req, err := lastResults.officeConsentListPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "securityinsight.OfficeConsentsClient", "listNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "securityinsight.OfficeConsentsClient", "listNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "securityinsight.OfficeConsentsClient", "listNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

    // ListComplete enumerates all values, automatically crossing page boundaries as required.
    func (client OfficeConsentsClient) ListComplete(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string) (result OfficeConsentListIterator, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/OfficeConsentsClient.List")
            defer func() {
                sc := -1
                if result.Response().Response.Response != nil {
                    sc = result.page.Response().Response.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
     }
        result.page, err = client.List(ctx, resourceGroupName, operationalInsightsResourceProvider, workspaceName)
                return
        }

