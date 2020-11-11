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

// AlertRuleIncidentsClient is the monitor Management Client
type AlertRuleIncidentsClient struct {
    BaseClient
}
// NewAlertRuleIncidentsClient creates an instance of the AlertRuleIncidentsClient client.
func NewAlertRuleIncidentsClient(subscriptionID string) AlertRuleIncidentsClient {
    return NewAlertRuleIncidentsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewAlertRuleIncidentsClientWithBaseURI creates an instance of the AlertRuleIncidentsClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
    func NewAlertRuleIncidentsClientWithBaseURI(baseURI string, subscriptionID string) AlertRuleIncidentsClient {
        return AlertRuleIncidentsClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// Get gets an incident associated to an alert rule
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // ruleName - the name of the rule.
        // incidentName - the name of the incident to retrieve.
func (client AlertRuleIncidentsClient) Get(ctx context.Context, resourceGroupName string, ruleName string, incidentName string) (result Incident, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/AlertRuleIncidentsClient.Get")
        defer func() {
            sc := -1
        if result.Response.Response != nil {
        sc = result.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.GetPreparer(ctx, resourceGroupName, ruleName, incidentName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "insights.AlertRuleIncidentsClient", "Get", nil , "Failure preparing request")
    return
    }

        resp, err := client.GetSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "insights.AlertRuleIncidentsClient", "Get", resp, "Failure sending request")
        return
        }

        result, err = client.GetResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "insights.AlertRuleIncidentsClient", "Get", resp, "Failure responding to request")
        }

    return
}

    // GetPreparer prepares the Get request.
    func (client AlertRuleIncidentsClient) GetPreparer(ctx context.Context, resourceGroupName string, ruleName string, incidentName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "incidentName": autorest.Encode("path",incidentName),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "ruleName": autorest.Encode("path",ruleName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2016-03-01"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/microsoft.insights/alertrules/{ruleName}/incidents/{incidentName}",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client AlertRuleIncidentsClient) GetSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

    // GetResponder handles the response to the Get request. The method always
    // closes the http.Response Body.
    func (client AlertRuleIncidentsClient) GetResponder(resp *http.Response) (result Incident, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// ListByAlertRule gets a list of incidents associated to an alert rule
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // ruleName - the name of the rule.
func (client AlertRuleIncidentsClient) ListByAlertRule(ctx context.Context, resourceGroupName string, ruleName string) (result IncidentListResult, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/AlertRuleIncidentsClient.ListByAlertRule")
        defer func() {
            sc := -1
        if result.Response.Response != nil {
        sc = result.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.ListByAlertRulePreparer(ctx, resourceGroupName, ruleName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "insights.AlertRuleIncidentsClient", "ListByAlertRule", nil , "Failure preparing request")
    return
    }

        resp, err := client.ListByAlertRuleSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "insights.AlertRuleIncidentsClient", "ListByAlertRule", resp, "Failure sending request")
        return
        }

        result, err = client.ListByAlertRuleResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "insights.AlertRuleIncidentsClient", "ListByAlertRule", resp, "Failure responding to request")
        }

    return
}

    // ListByAlertRulePreparer prepares the ListByAlertRule request.
    func (client AlertRuleIncidentsClient) ListByAlertRulePreparer(ctx context.Context, resourceGroupName string, ruleName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "ruleName": autorest.Encode("path",ruleName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2016-03-01"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/microsoft.insights/alertrules/{ruleName}/incidents",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListByAlertRuleSender sends the ListByAlertRule request. The method will close the
    // http.Response Body if it receives an error.
    func (client AlertRuleIncidentsClient) ListByAlertRuleSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

    // ListByAlertRuleResponder handles the response to the ListByAlertRule request. The method always
    // closes the http.Response Body.
    func (client AlertRuleIncidentsClient) ListByAlertRuleResponder(resp *http.Response) (result IncidentListResult, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

