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

// ThreatIntelligenceIndicatorMetricsClient is the API spec for Microsoft.SecurityInsights (Azure Security Insights)
// resource provider
type ThreatIntelligenceIndicatorMetricsClient struct {
    BaseClient
}
// NewThreatIntelligenceIndicatorMetricsClient creates an instance of the ThreatIntelligenceIndicatorMetricsClient
// client.
func NewThreatIntelligenceIndicatorMetricsClient(subscriptionID string) ThreatIntelligenceIndicatorMetricsClient {
    return NewThreatIntelligenceIndicatorMetricsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewThreatIntelligenceIndicatorMetricsClientWithBaseURI creates an instance of the
// ThreatIntelligenceIndicatorMetricsClient client using a custom endpoint.  Use this when interacting with an Azure
// cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
    func NewThreatIntelligenceIndicatorMetricsClientWithBaseURI(baseURI string, subscriptionID string) ThreatIntelligenceIndicatorMetricsClient {
        return ThreatIntelligenceIndicatorMetricsClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// Get get the threat intelligence metrics.
    // Parameters:
        // resourceGroupName - the name of the resource group within the user's subscription. The name is case
        // insensitive.
        // operationalInsightsResourceProvider - the namespace of workspaces resource provider-
        // Microsoft.OperationalInsights.
        // workspaceName - the name of the workspace.
        // ctiEntityKind - the threat intelligence entity kind
func (client ThreatIntelligenceIndicatorMetricsClient) Get(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, ctiEntityKind string) (result ThreatIntelligenceMetricResourceList, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ThreatIntelligenceIndicatorMetricsClient.Get")
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
        return result, validation.NewError("securityinsight.ThreatIntelligenceIndicatorMetricsClient", "Get", err.Error())
        }

        req, err := client.GetPreparer(ctx, resourceGroupName, operationalInsightsResourceProvider, workspaceName, ctiEntityKind)
    if err != nil {
    err = autorest.NewErrorWithError(err, "securityinsight.ThreatIntelligenceIndicatorMetricsClient", "Get", nil , "Failure preparing request")
    return
    }

        resp, err := client.GetSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "securityinsight.ThreatIntelligenceIndicatorMetricsClient", "Get", resp, "Failure sending request")
        return
        }

        result, err = client.GetResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "securityinsight.ThreatIntelligenceIndicatorMetricsClient", "Get", resp, "Failure responding to request")
        }

    return
}

    // GetPreparer prepares the Get request.
    func (client ThreatIntelligenceIndicatorMetricsClient) GetPreparer(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, ctiEntityKind string) (*http.Request, error) {
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
        if len(ctiEntityKind) > 0 {
        queryParameters["ctiEntityKind"] = autorest.Encode("query",ctiEntityKind)
        }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{operationalInsightsResourceProvider}/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/threatIntelligence/main/metrics",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client ThreatIntelligenceIndicatorMetricsClient) GetSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

    // GetResponder handles the response to the Get request. The method always
    // closes the http.Response Body.
    func (client ThreatIntelligenceIndicatorMetricsClient) GetResponder(resp *http.Response) (result ThreatIntelligenceMetricResourceList, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

