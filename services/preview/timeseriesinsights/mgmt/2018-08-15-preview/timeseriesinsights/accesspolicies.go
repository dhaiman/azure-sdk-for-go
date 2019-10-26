package timeseriesinsights

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

// AccessPoliciesClient is the time Series Insights client
type AccessPoliciesClient struct {
    BaseClient
}
// NewAccessPoliciesClient creates an instance of the AccessPoliciesClient client.
func NewAccessPoliciesClient(subscriptionID string) AccessPoliciesClient {
    return NewAccessPoliciesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewAccessPoliciesClientWithBaseURI creates an instance of the AccessPoliciesClient client.
    func NewAccessPoliciesClientWithBaseURI(baseURI string, subscriptionID string) AccessPoliciesClient {
        return AccessPoliciesClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// CreateOrUpdate create or update an access policy in the specified environment.
    // Parameters:
        // resourceGroupName - name of an Azure Resource group.
        // environmentName - the name of the Time Series Insights environment associated with the specified resource
        // group.
        // accessPolicyName - name of the access policy.
        // parameters - parameters for creating an access policy.
func (client AccessPoliciesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, environmentName string, accessPolicyName string, parameters AccessPolicyCreateOrUpdateParameters) (result AccessPolicyResource, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/AccessPoliciesClient.CreateOrUpdate")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: accessPolicyName,
             Constraints: []validation.Constraint{	{Target: "accessPolicyName", Name: validation.MaxLength, Rule: 90, Chain: nil },
            	{Target: "accessPolicyName", Name: validation.MinLength, Rule: 1, Chain: nil },
            	{Target: "accessPolicyName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil }}},
            { TargetValue: parameters,
             Constraints: []validation.Constraint{	{Target: "parameters.AccessPolicyResourceProperties", Name: validation.Null, Rule: true, Chain: nil }}}}); err != nil {
            return result, validation.NewError("timeseriesinsights.AccessPoliciesClient", "CreateOrUpdate", err.Error())
            }

                req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, environmentName, accessPolicyName, parameters)
    if err != nil {
    err = autorest.NewErrorWithError(err, "timeseriesinsights.AccessPoliciesClient", "CreateOrUpdate", nil , "Failure preparing request")
    return
    }

            resp, err := client.CreateOrUpdateSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "timeseriesinsights.AccessPoliciesClient", "CreateOrUpdate", resp, "Failure sending request")
            return
            }

            result, err = client.CreateOrUpdateResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "timeseriesinsights.AccessPoliciesClient", "CreateOrUpdate", resp, "Failure responding to request")
            }

    return
    }

    // CreateOrUpdatePreparer prepares the CreateOrUpdate request.
    func (client AccessPoliciesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, environmentName string, accessPolicyName string, parameters AccessPolicyCreateOrUpdateParameters) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "accessPolicyName": autorest.Encode("path",accessPolicyName),
            "environmentName": autorest.Encode("path",environmentName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2018-08-15-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPut(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/accessPolicies/{accessPolicyName}",pathParameters),
    autorest.WithJSON(parameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
    // http.Response Body if it receives an error.
    func (client AccessPoliciesClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client AccessPoliciesClient) CreateOrUpdateResponder(resp *http.Response) (result AccessPolicyResource, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusCreated),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// Delete deletes the access policy with the specified name in the specified subscription, resource group, and
// environment
    // Parameters:
        // resourceGroupName - name of an Azure Resource group.
        // environmentName - the name of the Time Series Insights environment associated with the specified resource
        // group.
        // accessPolicyName - the name of the Time Series Insights access policy associated with the specified
        // environment.
func (client AccessPoliciesClient) Delete(ctx context.Context, resourceGroupName string, environmentName string, accessPolicyName string) (result autorest.Response, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/AccessPoliciesClient.Delete")
        defer func() {
            sc := -1
            if result.Response != nil {
                sc = result.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.DeletePreparer(ctx, resourceGroupName, environmentName, accessPolicyName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "timeseriesinsights.AccessPoliciesClient", "Delete", nil , "Failure preparing request")
    return
    }

            resp, err := client.DeleteSender(req)
            if err != nil {
            result.Response = resp
            err = autorest.NewErrorWithError(err, "timeseriesinsights.AccessPoliciesClient", "Delete", resp, "Failure sending request")
            return
            }

            result, err = client.DeleteResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "timeseriesinsights.AccessPoliciesClient", "Delete", resp, "Failure responding to request")
            }

    return
    }

    // DeletePreparer prepares the Delete request.
    func (client AccessPoliciesClient) DeletePreparer(ctx context.Context, resourceGroupName string, environmentName string, accessPolicyName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "accessPolicyName": autorest.Encode("path",accessPolicyName),
            "environmentName": autorest.Encode("path",environmentName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2018-08-15-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsDelete(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/accessPolicies/{accessPolicyName}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // DeleteSender sends the Delete request. The method will close the
    // http.Response Body if it receives an error.
    func (client AccessPoliciesClient) DeleteSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client AccessPoliciesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusNoContent),
    autorest.ByClosing())
    result.Response = resp
        return
    }

// Get gets the access policy with the specified name in the specified environment.
    // Parameters:
        // resourceGroupName - name of an Azure Resource group.
        // environmentName - the name of the Time Series Insights environment associated with the specified resource
        // group.
        // accessPolicyName - the name of the Time Series Insights access policy associated with the specified
        // environment.
func (client AccessPoliciesClient) Get(ctx context.Context, resourceGroupName string, environmentName string, accessPolicyName string) (result AccessPolicyResource, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/AccessPoliciesClient.Get")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.GetPreparer(ctx, resourceGroupName, environmentName, accessPolicyName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "timeseriesinsights.AccessPoliciesClient", "Get", nil , "Failure preparing request")
    return
    }

            resp, err := client.GetSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "timeseriesinsights.AccessPoliciesClient", "Get", resp, "Failure sending request")
            return
            }

            result, err = client.GetResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "timeseriesinsights.AccessPoliciesClient", "Get", resp, "Failure responding to request")
            }

    return
    }

    // GetPreparer prepares the Get request.
    func (client AccessPoliciesClient) GetPreparer(ctx context.Context, resourceGroupName string, environmentName string, accessPolicyName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "accessPolicyName": autorest.Encode("path",accessPolicyName),
            "environmentName": autorest.Encode("path",environmentName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2018-08-15-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/accessPolicies/{accessPolicyName}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client AccessPoliciesClient) GetSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client AccessPoliciesClient) GetResponder(resp *http.Response) (result AccessPolicyResource, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// ListByEnvironment lists all the available access policies associated with the environment.
    // Parameters:
        // resourceGroupName - name of an Azure Resource group.
        // environmentName - the name of the Time Series Insights environment associated with the specified resource
        // group.
func (client AccessPoliciesClient) ListByEnvironment(ctx context.Context, resourceGroupName string, environmentName string) (result AccessPolicyListResponse, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/AccessPoliciesClient.ListByEnvironment")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.ListByEnvironmentPreparer(ctx, resourceGroupName, environmentName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "timeseriesinsights.AccessPoliciesClient", "ListByEnvironment", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListByEnvironmentSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "timeseriesinsights.AccessPoliciesClient", "ListByEnvironment", resp, "Failure sending request")
            return
            }

            result, err = client.ListByEnvironmentResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "timeseriesinsights.AccessPoliciesClient", "ListByEnvironment", resp, "Failure responding to request")
            }

    return
    }

    // ListByEnvironmentPreparer prepares the ListByEnvironment request.
    func (client AccessPoliciesClient) ListByEnvironmentPreparer(ctx context.Context, resourceGroupName string, environmentName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "environmentName": autorest.Encode("path",environmentName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2018-08-15-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/accessPolicies",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListByEnvironmentSender sends the ListByEnvironment request. The method will close the
    // http.Response Body if it receives an error.
    func (client AccessPoliciesClient) ListByEnvironmentSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// ListByEnvironmentResponder handles the response to the ListByEnvironment request. The method always
// closes the http.Response Body.
func (client AccessPoliciesClient) ListByEnvironmentResponder(resp *http.Response) (result AccessPolicyListResponse, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// Update updates the access policy with the specified name in the specified subscription, resource group, and
// environment.
    // Parameters:
        // resourceGroupName - name of an Azure Resource group.
        // environmentName - the name of the Time Series Insights environment associated with the specified resource
        // group.
        // accessPolicyName - the name of the Time Series Insights access policy associated with the specified
        // environment.
        // accessPolicyUpdateParameters - request object that contains the updated information for the access policy.
func (client AccessPoliciesClient) Update(ctx context.Context, resourceGroupName string, environmentName string, accessPolicyName string, accessPolicyUpdateParameters AccessPolicyUpdateParameters) (result AccessPolicyResource, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/AccessPoliciesClient.Update")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.UpdatePreparer(ctx, resourceGroupName, environmentName, accessPolicyName, accessPolicyUpdateParameters)
    if err != nil {
    err = autorest.NewErrorWithError(err, "timeseriesinsights.AccessPoliciesClient", "Update", nil , "Failure preparing request")
    return
    }

            resp, err := client.UpdateSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "timeseriesinsights.AccessPoliciesClient", "Update", resp, "Failure sending request")
            return
            }

            result, err = client.UpdateResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "timeseriesinsights.AccessPoliciesClient", "Update", resp, "Failure responding to request")
            }

    return
    }

    // UpdatePreparer prepares the Update request.
    func (client AccessPoliciesClient) UpdatePreparer(ctx context.Context, resourceGroupName string, environmentName string, accessPolicyName string, accessPolicyUpdateParameters AccessPolicyUpdateParameters) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "accessPolicyName": autorest.Encode("path",accessPolicyName),
            "environmentName": autorest.Encode("path",environmentName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2018-08-15-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPatch(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/accessPolicies/{accessPolicyName}",pathParameters),
    autorest.WithJSON(accessPolicyUpdateParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // UpdateSender sends the Update request. The method will close the
    // http.Response Body if it receives an error.
    func (client AccessPoliciesClient) UpdateSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client AccessPoliciesClient) UpdateResponder(resp *http.Response) (result AccessPolicyResource, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

