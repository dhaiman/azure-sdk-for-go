package apimanagement

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

// SignInSettingsClient is the apiManagement Client
type SignInSettingsClient struct {
    BaseClient
}
// NewSignInSettingsClient creates an instance of the SignInSettingsClient client.
func NewSignInSettingsClient(subscriptionID string) SignInSettingsClient {
    return NewSignInSettingsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewSignInSettingsClientWithBaseURI creates an instance of the SignInSettingsClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
    func NewSignInSettingsClientWithBaseURI(baseURI string, subscriptionID string) SignInSettingsClient {
        return SignInSettingsClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// CreateOrUpdate create or Update Sign-In settings.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // serviceName - the name of the API Management service.
        // parameters - create or update parameters.
func (client SignInSettingsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, parameters PortalSigninSettings) (result PortalSigninSettings, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/SignInSettingsClient.CreateOrUpdate")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: serviceName,
             Constraints: []validation.Constraint{	{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil },
            	{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil },
            	{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("apimanagement.SignInSettingsClient", "CreateOrUpdate", err.Error())
            }

                req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, serviceName, parameters)
    if err != nil {
    err = autorest.NewErrorWithError(err, "apimanagement.SignInSettingsClient", "CreateOrUpdate", nil , "Failure preparing request")
    return
    }

            resp, err := client.CreateOrUpdateSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "apimanagement.SignInSettingsClient", "CreateOrUpdate", resp, "Failure sending request")
            return
            }

            result, err = client.CreateOrUpdateResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "apimanagement.SignInSettingsClient", "CreateOrUpdate", resp, "Failure responding to request")
            }

    return
    }

    // CreateOrUpdatePreparer prepares the CreateOrUpdate request.
    func (client SignInSettingsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, serviceName string, parameters PortalSigninSettings) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "serviceName": autorest.Encode("path",serviceName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2017-03-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPut(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/portalsettings/signin",pathParameters),
    autorest.WithJSON(parameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
    // http.Response Body if it receives an error.
    func (client SignInSettingsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client SignInSettingsClient) CreateOrUpdateResponder(resp *http.Response) (result PortalSigninSettings, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusNoContent),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// Get get Sign-In settings.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // serviceName - the name of the API Management service.
func (client SignInSettingsClient) Get(ctx context.Context, resourceGroupName string, serviceName string) (result PortalSigninSettings, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/SignInSettingsClient.Get")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: serviceName,
             Constraints: []validation.Constraint{	{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil },
            	{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil },
            	{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("apimanagement.SignInSettingsClient", "Get", err.Error())
            }

                req, err := client.GetPreparer(ctx, resourceGroupName, serviceName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "apimanagement.SignInSettingsClient", "Get", nil , "Failure preparing request")
    return
    }

            resp, err := client.GetSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "apimanagement.SignInSettingsClient", "Get", resp, "Failure sending request")
            return
            }

            result, err = client.GetResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "apimanagement.SignInSettingsClient", "Get", resp, "Failure responding to request")
            }

    return
    }

    // GetPreparer prepares the Get request.
    func (client SignInSettingsClient) GetPreparer(ctx context.Context, resourceGroupName string, serviceName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "serviceName": autorest.Encode("path",serviceName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2017-03-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/portalsettings/signin",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client SignInSettingsClient) GetSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client SignInSettingsClient) GetResponder(resp *http.Response) (result PortalSigninSettings, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// Head gets the entity state (Etag) version of the SignInSettings.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // serviceName - the name of the API Management service.
func (client SignInSettingsClient) Head(ctx context.Context, resourceGroupName string, serviceName string) (result autorest.Response, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/SignInSettingsClient.Head")
        defer func() {
            sc := -1
            if result.Response != nil {
                sc = result.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: serviceName,
             Constraints: []validation.Constraint{	{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil },
            	{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil },
            	{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("apimanagement.SignInSettingsClient", "Head", err.Error())
            }

                req, err := client.HeadPreparer(ctx, resourceGroupName, serviceName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "apimanagement.SignInSettingsClient", "Head", nil , "Failure preparing request")
    return
    }

            resp, err := client.HeadSender(req)
            if err != nil {
            result.Response = resp
            err = autorest.NewErrorWithError(err, "apimanagement.SignInSettingsClient", "Head", resp, "Failure sending request")
            return
            }

            result, err = client.HeadResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "apimanagement.SignInSettingsClient", "Head", resp, "Failure responding to request")
            }

    return
    }

    // HeadPreparer prepares the Head request.
    func (client SignInSettingsClient) HeadPreparer(ctx context.Context, resourceGroupName string, serviceName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "serviceName": autorest.Encode("path",serviceName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2017-03-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsHead(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/portalsettings/signin",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // HeadSender sends the Head request. The method will close the
    // http.Response Body if it receives an error.
    func (client SignInSettingsClient) HeadSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

// HeadResponder handles the response to the Head request. The method always
// closes the http.Response Body.
func (client SignInSettingsClient) HeadResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByClosing())
    result.Response = resp
        return
    }

// Update update Sign-In settings.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // serviceName - the name of the API Management service.
        // parameters - update Sign-In settings.
        // ifMatch - the entity state (Etag) version of the property to update. A value of "*" can be used for If-Match
        // to unconditionally apply the operation.
func (client SignInSettingsClient) Update(ctx context.Context, resourceGroupName string, serviceName string, parameters PortalSigninSettings, ifMatch string) (result autorest.Response, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/SignInSettingsClient.Update")
        defer func() {
            sc := -1
            if result.Response != nil {
                sc = result.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: serviceName,
             Constraints: []validation.Constraint{	{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil },
            	{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil },
            	{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("apimanagement.SignInSettingsClient", "Update", err.Error())
            }

                req, err := client.UpdatePreparer(ctx, resourceGroupName, serviceName, parameters, ifMatch)
    if err != nil {
    err = autorest.NewErrorWithError(err, "apimanagement.SignInSettingsClient", "Update", nil , "Failure preparing request")
    return
    }

            resp, err := client.UpdateSender(req)
            if err != nil {
            result.Response = resp
            err = autorest.NewErrorWithError(err, "apimanagement.SignInSettingsClient", "Update", resp, "Failure sending request")
            return
            }

            result, err = client.UpdateResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "apimanagement.SignInSettingsClient", "Update", resp, "Failure responding to request")
            }

    return
    }

    // UpdatePreparer prepares the Update request.
    func (client SignInSettingsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, serviceName string, parameters PortalSigninSettings, ifMatch string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "serviceName": autorest.Encode("path",serviceName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2017-03-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPatch(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/portalsettings/signin",pathParameters),
    autorest.WithJSON(parameters),
    autorest.WithQueryParameters(queryParameters),
    autorest.WithHeader("If-Match", autorest.String(ifMatch)))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // UpdateSender sends the Update request. The method will close the
    // http.Response Body if it receives an error.
    func (client SignInSettingsClient) UpdateSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client SignInSettingsClient) UpdateResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusNoContent),
    autorest.ByClosing())
    result.Response = resp
        return
    }

