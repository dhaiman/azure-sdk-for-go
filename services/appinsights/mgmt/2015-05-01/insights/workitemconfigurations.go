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

// WorkItemConfigurationsClient is the composite Swagger for Application Insights Management Client
type WorkItemConfigurationsClient struct {
    BaseClient
}
// NewWorkItemConfigurationsClient creates an instance of the WorkItemConfigurationsClient client.
func NewWorkItemConfigurationsClient(subscriptionID string) WorkItemConfigurationsClient {
    return NewWorkItemConfigurationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWorkItemConfigurationsClientWithBaseURI creates an instance of the WorkItemConfigurationsClient client.
    func NewWorkItemConfigurationsClientWithBaseURI(baseURI string, subscriptionID string) WorkItemConfigurationsClient {
        return WorkItemConfigurationsClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// Create create a work item configuration for an Application Insights component.
    // Parameters:
        // resourceGroupName - the name of the resource group. The name is case insensitive.
        // resourceName - the name of the Application Insights component resource.
        // workItemConfigurationProperties - properties that need to be specified to create a work item configuration
        // of a Application Insights component.
func (client WorkItemConfigurationsClient) Create(ctx context.Context, resourceGroupName string, resourceName string, workItemConfigurationProperties WorkItemCreateConfiguration) (result WorkItemConfiguration, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/WorkItemConfigurationsClient.Create")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: resourceGroupName,
             Constraints: []validation.Constraint{	{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil },
            	{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil },
            	{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil }}},
            { TargetValue: client.SubscriptionID,
             Constraints: []validation.Constraint{	{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil }}}}); err != nil {
            return result, validation.NewError("insights.WorkItemConfigurationsClient", "Create", err.Error())
            }

                req, err := client.CreatePreparer(ctx, resourceGroupName, resourceName, workItemConfigurationProperties)
    if err != nil {
    err = autorest.NewErrorWithError(err, "insights.WorkItemConfigurationsClient", "Create", nil , "Failure preparing request")
    return
    }

            resp, err := client.CreateSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "insights.WorkItemConfigurationsClient", "Create", resp, "Failure sending request")
            return
            }

            result, err = client.CreateResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "insights.WorkItemConfigurationsClient", "Create", resp, "Failure responding to request")
            }

    return
    }

    // CreatePreparer prepares the Create request.
    func (client WorkItemConfigurationsClient) CreatePreparer(ctx context.Context, resourceGroupName string, resourceName string, workItemConfigurationProperties WorkItemCreateConfiguration) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "resourceName": autorest.Encode("path",resourceName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2015-05-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPost(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/WorkItemConfigs",pathParameters),
    autorest.WithJSON(workItemConfigurationProperties),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CreateSender sends the Create request. The method will close the
    // http.Response Body if it receives an error.
    func (client WorkItemConfigurationsClient) CreateSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client WorkItemConfigurationsClient) CreateResponder(resp *http.Response) (result WorkItemConfiguration, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// Delete delete a work item configuration of an Application Insights component.
    // Parameters:
        // resourceGroupName - the name of the resource group. The name is case insensitive.
        // resourceName - the name of the Application Insights component resource.
        // workItemConfigID - the unique work item configuration Id. This can be either friendly name of connector as
        // defined in connector configuration
func (client WorkItemConfigurationsClient) Delete(ctx context.Context, resourceGroupName string, resourceName string, workItemConfigID string) (result autorest.Response, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/WorkItemConfigurationsClient.Delete")
        defer func() {
            sc := -1
            if result.Response != nil {
                sc = result.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: resourceGroupName,
             Constraints: []validation.Constraint{	{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil },
            	{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil },
            	{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil }}},
            { TargetValue: client.SubscriptionID,
             Constraints: []validation.Constraint{	{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil }}}}); err != nil {
            return result, validation.NewError("insights.WorkItemConfigurationsClient", "Delete", err.Error())
            }

                req, err := client.DeletePreparer(ctx, resourceGroupName, resourceName, workItemConfigID)
    if err != nil {
    err = autorest.NewErrorWithError(err, "insights.WorkItemConfigurationsClient", "Delete", nil , "Failure preparing request")
    return
    }

            resp, err := client.DeleteSender(req)
            if err != nil {
            result.Response = resp
            err = autorest.NewErrorWithError(err, "insights.WorkItemConfigurationsClient", "Delete", resp, "Failure sending request")
            return
            }

            result, err = client.DeleteResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "insights.WorkItemConfigurationsClient", "Delete", resp, "Failure responding to request")
            }

    return
    }

    // DeletePreparer prepares the Delete request.
    func (client WorkItemConfigurationsClient) DeletePreparer(ctx context.Context, resourceGroupName string, resourceName string, workItemConfigID string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "resourceName": autorest.Encode("path",resourceName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            "workItemConfigId": autorest.Encode("path",workItemConfigID),
            }

                        const APIVersion = "2015-05-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsDelete(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/WorkItemConfigs/{workItemConfigId}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // DeleteSender sends the Delete request. The method will close the
    // http.Response Body if it receives an error.
    func (client WorkItemConfigurationsClient) DeleteSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client WorkItemConfigurationsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByClosing())
    result.Response = resp
        return
    }

// GetDefault gets default work item configurations that exist for the application
    // Parameters:
        // resourceGroupName - the name of the resource group. The name is case insensitive.
        // resourceName - the name of the Application Insights component resource.
func (client WorkItemConfigurationsClient) GetDefault(ctx context.Context, resourceGroupName string, resourceName string) (result WorkItemConfiguration, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/WorkItemConfigurationsClient.GetDefault")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: resourceGroupName,
             Constraints: []validation.Constraint{	{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil },
            	{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil },
            	{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil }}},
            { TargetValue: client.SubscriptionID,
             Constraints: []validation.Constraint{	{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil }}}}); err != nil {
            return result, validation.NewError("insights.WorkItemConfigurationsClient", "GetDefault", err.Error())
            }

                req, err := client.GetDefaultPreparer(ctx, resourceGroupName, resourceName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "insights.WorkItemConfigurationsClient", "GetDefault", nil , "Failure preparing request")
    return
    }

            resp, err := client.GetDefaultSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "insights.WorkItemConfigurationsClient", "GetDefault", resp, "Failure sending request")
            return
            }

            result, err = client.GetDefaultResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "insights.WorkItemConfigurationsClient", "GetDefault", resp, "Failure responding to request")
            }

    return
    }

    // GetDefaultPreparer prepares the GetDefault request.
    func (client WorkItemConfigurationsClient) GetDefaultPreparer(ctx context.Context, resourceGroupName string, resourceName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "resourceName": autorest.Encode("path",resourceName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2015-05-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/DefaultWorkItemConfig",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetDefaultSender sends the GetDefault request. The method will close the
    // http.Response Body if it receives an error.
    func (client WorkItemConfigurationsClient) GetDefaultSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// GetDefaultResponder handles the response to the GetDefault request. The method always
// closes the http.Response Body.
func (client WorkItemConfigurationsClient) GetDefaultResponder(resp *http.Response) (result WorkItemConfiguration, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// GetItem gets specified work item configuration for an Application Insights component.
    // Parameters:
        // resourceGroupName - the name of the resource group. The name is case insensitive.
        // resourceName - the name of the Application Insights component resource.
        // workItemConfigID - the unique work item configuration Id. This can be either friendly name of connector as
        // defined in connector configuration
func (client WorkItemConfigurationsClient) GetItem(ctx context.Context, resourceGroupName string, resourceName string, workItemConfigID string) (result WorkItemConfiguration, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/WorkItemConfigurationsClient.GetItem")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: resourceGroupName,
             Constraints: []validation.Constraint{	{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil },
            	{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil },
            	{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil }}},
            { TargetValue: client.SubscriptionID,
             Constraints: []validation.Constraint{	{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil }}}}); err != nil {
            return result, validation.NewError("insights.WorkItemConfigurationsClient", "GetItem", err.Error())
            }

                req, err := client.GetItemPreparer(ctx, resourceGroupName, resourceName, workItemConfigID)
    if err != nil {
    err = autorest.NewErrorWithError(err, "insights.WorkItemConfigurationsClient", "GetItem", nil , "Failure preparing request")
    return
    }

            resp, err := client.GetItemSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "insights.WorkItemConfigurationsClient", "GetItem", resp, "Failure sending request")
            return
            }

            result, err = client.GetItemResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "insights.WorkItemConfigurationsClient", "GetItem", resp, "Failure responding to request")
            }

    return
    }

    // GetItemPreparer prepares the GetItem request.
    func (client WorkItemConfigurationsClient) GetItemPreparer(ctx context.Context, resourceGroupName string, resourceName string, workItemConfigID string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "resourceName": autorest.Encode("path",resourceName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            "workItemConfigId": autorest.Encode("path",workItemConfigID),
            }

                        const APIVersion = "2015-05-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/WorkItemConfigs/{workItemConfigId}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetItemSender sends the GetItem request. The method will close the
    // http.Response Body if it receives an error.
    func (client WorkItemConfigurationsClient) GetItemSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// GetItemResponder handles the response to the GetItem request. The method always
// closes the http.Response Body.
func (client WorkItemConfigurationsClient) GetItemResponder(resp *http.Response) (result WorkItemConfiguration, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// List gets the list work item configurations that exist for the application
    // Parameters:
        // resourceGroupName - the name of the resource group. The name is case insensitive.
        // resourceName - the name of the Application Insights component resource.
func (client WorkItemConfigurationsClient) List(ctx context.Context, resourceGroupName string, resourceName string) (result WorkItemConfigurationsListResult, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/WorkItemConfigurationsClient.List")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: resourceGroupName,
             Constraints: []validation.Constraint{	{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil },
            	{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil },
            	{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil }}},
            { TargetValue: client.SubscriptionID,
             Constraints: []validation.Constraint{	{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil }}}}); err != nil {
            return result, validation.NewError("insights.WorkItemConfigurationsClient", "List", err.Error())
            }

                req, err := client.ListPreparer(ctx, resourceGroupName, resourceName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "insights.WorkItemConfigurationsClient", "List", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "insights.WorkItemConfigurationsClient", "List", resp, "Failure sending request")
            return
            }

            result, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "insights.WorkItemConfigurationsClient", "List", resp, "Failure responding to request")
            }

    return
    }

    // ListPreparer prepares the List request.
    func (client WorkItemConfigurationsClient) ListPreparer(ctx context.Context, resourceGroupName string, resourceName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "resourceName": autorest.Encode("path",resourceName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2015-05-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/WorkItemConfigs",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListSender sends the List request. The method will close the
    // http.Response Body if it receives an error.
    func (client WorkItemConfigurationsClient) ListSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client WorkItemConfigurationsClient) ListResponder(resp *http.Response) (result WorkItemConfigurationsListResult, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// UpdateItem update a work item configuration for an Application Insights component.
    // Parameters:
        // resourceGroupName - the name of the resource group. The name is case insensitive.
        // resourceName - the name of the Application Insights component resource.
        // workItemConfigID - the unique work item configuration Id. This can be either friendly name of connector as
        // defined in connector configuration
        // workItemConfigurationProperties - properties that need to be specified to update a work item configuration
        // for this Application Insights component.
func (client WorkItemConfigurationsClient) UpdateItem(ctx context.Context, resourceGroupName string, resourceName string, workItemConfigID string, workItemConfigurationProperties WorkItemCreateConfiguration) (result WorkItemConfiguration, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/WorkItemConfigurationsClient.UpdateItem")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: resourceGroupName,
             Constraints: []validation.Constraint{	{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil },
            	{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil },
            	{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil }}},
            { TargetValue: client.SubscriptionID,
             Constraints: []validation.Constraint{	{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil }}}}); err != nil {
            return result, validation.NewError("insights.WorkItemConfigurationsClient", "UpdateItem", err.Error())
            }

                req, err := client.UpdateItemPreparer(ctx, resourceGroupName, resourceName, workItemConfigID, workItemConfigurationProperties)
    if err != nil {
    err = autorest.NewErrorWithError(err, "insights.WorkItemConfigurationsClient", "UpdateItem", nil , "Failure preparing request")
    return
    }

            resp, err := client.UpdateItemSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "insights.WorkItemConfigurationsClient", "UpdateItem", resp, "Failure sending request")
            return
            }

            result, err = client.UpdateItemResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "insights.WorkItemConfigurationsClient", "UpdateItem", resp, "Failure responding to request")
            }

    return
    }

    // UpdateItemPreparer prepares the UpdateItem request.
    func (client WorkItemConfigurationsClient) UpdateItemPreparer(ctx context.Context, resourceGroupName string, resourceName string, workItemConfigID string, workItemConfigurationProperties WorkItemCreateConfiguration) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "resourceName": autorest.Encode("path",resourceName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            "workItemConfigId": autorest.Encode("path",workItemConfigID),
            }

                        const APIVersion = "2015-05-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPatch(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/WorkItemConfigs/{workItemConfigId}",pathParameters),
    autorest.WithJSON(workItemConfigurationProperties),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // UpdateItemSender sends the UpdateItem request. The method will close the
    // http.Response Body if it receives an error.
    func (client WorkItemConfigurationsClient) UpdateItemSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// UpdateItemResponder handles the response to the UpdateItem request. The method always
// closes the http.Response Body.
func (client WorkItemConfigurationsClient) UpdateItemResponder(resp *http.Response) (result WorkItemConfiguration, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

