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
    "github.com/Azure/go-autorest/autorest/validation"
)

// ScheduledQueryRulesClient is the monitor Management Client
type ScheduledQueryRulesClient struct {
    BaseClient
}
// NewScheduledQueryRulesClient creates an instance of the ScheduledQueryRulesClient client.
func NewScheduledQueryRulesClient(subscriptionID string) ScheduledQueryRulesClient {
    return NewScheduledQueryRulesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewScheduledQueryRulesClientWithBaseURI creates an instance of the ScheduledQueryRulesClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
    func NewScheduledQueryRulesClientWithBaseURI(baseURI string, subscriptionID string) ScheduledQueryRulesClient {
        return ScheduledQueryRulesClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// CreateOrUpdate creates or updates an log search rule.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // ruleName - the name of the rule.
        // parameters - the parameters of the rule to create or update.
func (client ScheduledQueryRulesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, ruleName string, parameters LogSearchRuleResource) (result LogSearchRuleResource, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ScheduledQueryRulesClient.CreateOrUpdate")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: parameters,
             Constraints: []validation.Constraint{	{Target: "parameters.LogSearchRule", Name: validation.Null, Rule: true ,
            Chain: []validation.Constraint{	{Target: "parameters.LogSearchRule.Source", Name: validation.Null, Rule: true ,
            Chain: []validation.Constraint{	{Target: "parameters.LogSearchRule.Source.DataSourceID", Name: validation.Null, Rule: true, Chain: nil },
            }},
            	{Target: "parameters.LogSearchRule.Schedule", Name: validation.Null, Rule: false ,
            Chain: []validation.Constraint{	{Target: "parameters.LogSearchRule.Schedule.FrequencyInMinutes", Name: validation.Null, Rule: true, Chain: nil },
            	{Target: "parameters.LogSearchRule.Schedule.TimeWindowInMinutes", Name: validation.Null, Rule: true, Chain: nil },
            }},
            	{Target: "parameters.LogSearchRule.Action", Name: validation.Null, Rule: true, Chain: nil },
            }}}}}); err != nil {
            return result, validation.NewError("insights.ScheduledQueryRulesClient", "CreateOrUpdate", err.Error())
            }

                req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, ruleName, parameters)
    if err != nil {
    err = autorest.NewErrorWithError(err, "insights.ScheduledQueryRulesClient", "CreateOrUpdate", nil , "Failure preparing request")
    return
    }

            resp, err := client.CreateOrUpdateSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "insights.ScheduledQueryRulesClient", "CreateOrUpdate", resp, "Failure sending request")
            return
            }

            result, err = client.CreateOrUpdateResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "insights.ScheduledQueryRulesClient", "CreateOrUpdate", resp, "Failure responding to request")
            }

    return
    }

    // CreateOrUpdatePreparer prepares the CreateOrUpdate request.
    func (client ScheduledQueryRulesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, ruleName string, parameters LogSearchRuleResource) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "ruleName": autorest.Encode("path",ruleName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2018-04-16"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPut(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/microsoft.insights/scheduledQueryRules/{ruleName}",pathParameters),
    autorest.WithJSON(parameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
    // http.Response Body if it receives an error.
    func (client ScheduledQueryRulesClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ScheduledQueryRulesClient) CreateOrUpdateResponder(resp *http.Response) (result LogSearchRuleResource, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusCreated),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// Delete deletes a Log Search rule
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // ruleName - the name of the rule.
func (client ScheduledQueryRulesClient) Delete(ctx context.Context, resourceGroupName string, ruleName string) (result autorest.Response, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ScheduledQueryRulesClient.Delete")
        defer func() {
            sc := -1
            if result.Response != nil {
                sc = result.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.DeletePreparer(ctx, resourceGroupName, ruleName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "insights.ScheduledQueryRulesClient", "Delete", nil , "Failure preparing request")
    return
    }

            resp, err := client.DeleteSender(req)
            if err != nil {
            result.Response = resp
            err = autorest.NewErrorWithError(err, "insights.ScheduledQueryRulesClient", "Delete", resp, "Failure sending request")
            return
            }

            result, err = client.DeleteResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "insights.ScheduledQueryRulesClient", "Delete", resp, "Failure responding to request")
            }

    return
    }

    // DeletePreparer prepares the Delete request.
    func (client ScheduledQueryRulesClient) DeletePreparer(ctx context.Context, resourceGroupName string, ruleName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "ruleName": autorest.Encode("path",ruleName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2018-04-16"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsDelete(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/microsoft.insights/scheduledQueryRules/{ruleName}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // DeleteSender sends the Delete request. The method will close the
    // http.Response Body if it receives an error.
    func (client ScheduledQueryRulesClient) DeleteSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ScheduledQueryRulesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusNoContent),
    autorest.ByClosing())
    result.Response = resp
        return
    }

// Get gets an Log Search rule
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // ruleName - the name of the rule.
func (client ScheduledQueryRulesClient) Get(ctx context.Context, resourceGroupName string, ruleName string) (result LogSearchRuleResource, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ScheduledQueryRulesClient.Get")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.GetPreparer(ctx, resourceGroupName, ruleName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "insights.ScheduledQueryRulesClient", "Get", nil , "Failure preparing request")
    return
    }

            resp, err := client.GetSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "insights.ScheduledQueryRulesClient", "Get", resp, "Failure sending request")
            return
            }

            result, err = client.GetResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "insights.ScheduledQueryRulesClient", "Get", resp, "Failure responding to request")
            }

    return
    }

    // GetPreparer prepares the Get request.
    func (client ScheduledQueryRulesClient) GetPreparer(ctx context.Context, resourceGroupName string, ruleName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "ruleName": autorest.Encode("path",ruleName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2018-04-16"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/microsoft.insights/scheduledQueryRules/{ruleName}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client ScheduledQueryRulesClient) GetSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ScheduledQueryRulesClient) GetResponder(resp *http.Response) (result LogSearchRuleResource, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// ListByResourceGroup list the Log Search rules within a resource group.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // filter - the filter to apply on the operation. For more information please see
        // https://msdn.microsoft.com/en-us/library/azure/dn931934.aspx
func (client ScheduledQueryRulesClient) ListByResourceGroup(ctx context.Context, resourceGroupName string, filter string) (result LogSearchRuleResourceCollection, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ScheduledQueryRulesClient.ListByResourceGroup")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName, filter)
    if err != nil {
    err = autorest.NewErrorWithError(err, "insights.ScheduledQueryRulesClient", "ListByResourceGroup", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListByResourceGroupSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "insights.ScheduledQueryRulesClient", "ListByResourceGroup", resp, "Failure sending request")
            return
            }

            result, err = client.ListByResourceGroupResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "insights.ScheduledQueryRulesClient", "ListByResourceGroup", resp, "Failure responding to request")
            }

    return
    }

    // ListByResourceGroupPreparer prepares the ListByResourceGroup request.
    func (client ScheduledQueryRulesClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string, filter string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2018-04-16"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }
            if len(filter) > 0 {
            queryParameters["$filter"] = autorest.Encode("query",filter)
            }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/microsoft.insights/scheduledQueryRules",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
    // http.Response Body if it receives an error.
    func (client ScheduledQueryRulesClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client ScheduledQueryRulesClient) ListByResourceGroupResponder(resp *http.Response) (result LogSearchRuleResourceCollection, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// ListBySubscription list the Log Search rules within a subscription group.
    // Parameters:
        // filter - the filter to apply on the operation. For more information please see
        // https://msdn.microsoft.com/en-us/library/azure/dn931934.aspx
func (client ScheduledQueryRulesClient) ListBySubscription(ctx context.Context, filter string) (result LogSearchRuleResourceCollection, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ScheduledQueryRulesClient.ListBySubscription")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.ListBySubscriptionPreparer(ctx, filter)
    if err != nil {
    err = autorest.NewErrorWithError(err, "insights.ScheduledQueryRulesClient", "ListBySubscription", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListBySubscriptionSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "insights.ScheduledQueryRulesClient", "ListBySubscription", resp, "Failure sending request")
            return
            }

            result, err = client.ListBySubscriptionResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "insights.ScheduledQueryRulesClient", "ListBySubscription", resp, "Failure responding to request")
            }

    return
    }

    // ListBySubscriptionPreparer prepares the ListBySubscription request.
    func (client ScheduledQueryRulesClient) ListBySubscriptionPreparer(ctx context.Context, filter string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2018-04-16"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }
            if len(filter) > 0 {
            queryParameters["$filter"] = autorest.Encode("query",filter)
            }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/microsoft.insights/scheduledQueryRules",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListBySubscriptionSender sends the ListBySubscription request. The method will close the
    // http.Response Body if it receives an error.
    func (client ScheduledQueryRulesClient) ListBySubscriptionSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

// ListBySubscriptionResponder handles the response to the ListBySubscription request. The method always
// closes the http.Response Body.
func (client ScheduledQueryRulesClient) ListBySubscriptionResponder(resp *http.Response) (result LogSearchRuleResourceCollection, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// Update update log search Rule.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // ruleName - the name of the rule.
        // parameters - the parameters of the rule to update.
func (client ScheduledQueryRulesClient) Update(ctx context.Context, resourceGroupName string, ruleName string, parameters LogSearchRuleResourcePatch) (result LogSearchRuleResource, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ScheduledQueryRulesClient.Update")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.UpdatePreparer(ctx, resourceGroupName, ruleName, parameters)
    if err != nil {
    err = autorest.NewErrorWithError(err, "insights.ScheduledQueryRulesClient", "Update", nil , "Failure preparing request")
    return
    }

            resp, err := client.UpdateSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "insights.ScheduledQueryRulesClient", "Update", resp, "Failure sending request")
            return
            }

            result, err = client.UpdateResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "insights.ScheduledQueryRulesClient", "Update", resp, "Failure responding to request")
            }

    return
    }

    // UpdatePreparer prepares the Update request.
    func (client ScheduledQueryRulesClient) UpdatePreparer(ctx context.Context, resourceGroupName string, ruleName string, parameters LogSearchRuleResourcePatch) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "ruleName": autorest.Encode("path",ruleName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2018-04-16"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPatch(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/microsoft.insights/scheduledQueryRules/{ruleName}",pathParameters),
    autorest.WithJSON(parameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // UpdateSender sends the Update request. The method will close the
    // http.Response Body if it receives an error.
    func (client ScheduledQueryRulesClient) UpdateSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client ScheduledQueryRulesClient) UpdateResponder(resp *http.Response) (result LogSearchRuleResource, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

