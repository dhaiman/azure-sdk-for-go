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
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// UsersClient is the apiManagement Client
type UsersClient struct {
	BaseClient
}

// NewUsersClient creates an instance of the UsersClient client.
func NewUsersClient(subscriptionID string) UsersClient {
	return NewUsersClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewUsersClientWithBaseURI creates an instance of the UsersClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewUsersClientWithBaseURI(baseURI string, subscriptionID string) UsersClient {
	return UsersClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or Updates a user.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// UID - user identifier. Must be unique in the current API Management service instance.
// parameters - create or update parameters.
func (client UsersClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, UID string, parameters UserCreateParameters) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UsersClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: UID,
			Constraints: []validation.Constraint{{Target: "UID", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "UID", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "UID", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.Email", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "parameters.Email", Name: validation.MaxLength, Rule: 254, Chain: nil},
					{Target: "parameters.Email", Name: validation.MinLength, Rule: 1, Chain: nil},
				}},
				{Target: "parameters.Password", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.FirstName", Name: validation.Null, Rule: true,
					Chain: []validation.Constraint{{Target: "parameters.FirstName", Name: validation.MaxLength, Rule: 100, Chain: nil},
						{Target: "parameters.FirstName", Name: validation.MinLength, Rule: 1, Chain: nil},
					}},
				{Target: "parameters.LastName", Name: validation.Null, Rule: true,
					Chain: []validation.Constraint{{Target: "parameters.LastName", Name: validation.MaxLength, Rule: 100, Chain: nil},
						{Target: "parameters.LastName", Name: validation.MinLength, Rule: 1, Chain: nil},
					}}}}}); err != nil {
		return result, validation.NewError("apimanagement.UsersClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, serviceName, UID, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.UsersClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "apimanagement.UsersClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.UsersClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client UsersClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, serviceName string, UID string, parameters UserCreateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"uid":               autorest.Encode("path", UID),
	}

	const APIVersion = "2016-10-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/users/{uid}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client UsersClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client UsersClient) CreateOrUpdateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Delete deletes specific user.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// UID - user identifier. Must be unique in the current API Management service instance.
// ifMatch - the entity state (Etag) version of the user to delete. A value of "*" can be used for If-Match to
// unconditionally apply the operation.
// deleteSubscriptions - whether to delete user's subscription or not.
func (client UsersClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, UID string, ifMatch string, deleteSubscriptions *bool) (result ErrorBodyContract, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UsersClient.Delete")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: UID,
			Constraints: []validation.Constraint{{Target: "UID", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "UID", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "UID", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.UsersClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, serviceName, UID, ifMatch, deleteSubscriptions)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.UsersClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.UsersClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.UsersClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client UsersClient) DeletePreparer(ctx context.Context, resourceGroupName string, serviceName string, UID string, ifMatch string, deleteSubscriptions *bool) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"uid":               autorest.Encode("path", UID),
	}

	const APIVersion = "2016-10-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if deleteSubscriptions != nil {
		queryParameters["deleteSubscriptions"] = autorest.Encode("query", *deleteSubscriptions)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/users/{uid}", pathParameters),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client UsersClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client UsersClient) DeleteResponder(resp *http.Response) (result ErrorBodyContract, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent, http.StatusMethodNotAllowed),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GenerateSsoURL retrieves a redirection URL containing an authentication token for signing a given user into the
// developer portal.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// UID - user identifier. Must be unique in the current API Management service instance.
func (client UsersClient) GenerateSsoURL(ctx context.Context, resourceGroupName string, serviceName string, UID string) (result GenerateSsoURLResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UsersClient.GenerateSsoURL")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: UID,
			Constraints: []validation.Constraint{{Target: "UID", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "UID", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "UID", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.UsersClient", "GenerateSsoURL", err.Error())
	}

	req, err := client.GenerateSsoURLPreparer(ctx, resourceGroupName, serviceName, UID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.UsersClient", "GenerateSsoURL", nil, "Failure preparing request")
		return
	}

	resp, err := client.GenerateSsoURLSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.UsersClient", "GenerateSsoURL", resp, "Failure sending request")
		return
	}

	result, err = client.GenerateSsoURLResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.UsersClient", "GenerateSsoURL", resp, "Failure responding to request")
	}

	return
}

// GenerateSsoURLPreparer prepares the GenerateSsoURL request.
func (client UsersClient) GenerateSsoURLPreparer(ctx context.Context, resourceGroupName string, serviceName string, UID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"uid":               autorest.Encode("path", UID),
	}

	const APIVersion = "2016-10-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/users/{uid}/generateSsoUrl", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GenerateSsoURLSender sends the GenerateSsoURL request. The method will close the
// http.Response Body if it receives an error.
func (client UsersClient) GenerateSsoURLSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GenerateSsoURLResponder handles the response to the GenerateSsoURL request. The method always
// closes the http.Response Body.
func (client UsersClient) GenerateSsoURLResponder(resp *http.Response) (result GenerateSsoURLResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get gets the details of the user specified by its identifier.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// UID - user identifier. Must be unique in the current API Management service instance.
func (client UsersClient) Get(ctx context.Context, resourceGroupName string, serviceName string, UID string) (result UserContract, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UsersClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: UID,
			Constraints: []validation.Constraint{{Target: "UID", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "UID", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "UID", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.UsersClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, serviceName, UID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.UsersClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.UsersClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.UsersClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client UsersClient) GetPreparer(ctx context.Context, resourceGroupName string, serviceName string, UID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"uid":               autorest.Encode("path", UID),
	}

	const APIVersion = "2016-10-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/users/{uid}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client UsersClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client UsersClient) GetResponder(resp *http.Response) (result UserContract, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetSharedAccessToken gets the Shared Access Authorization Token for the User.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// UID - user identifier. Must be unique in the current API Management service instance.
// parameters - create Authorization Token parameters.
func (client UsersClient) GetSharedAccessToken(ctx context.Context, resourceGroupName string, serviceName string, UID string, parameters UserTokenParameters) (result UserTokenResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UsersClient.GetSharedAccessToken")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: UID,
			Constraints: []validation.Constraint{{Target: "UID", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "UID", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "UID", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.Expiry", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.UsersClient", "GetSharedAccessToken", err.Error())
	}

	req, err := client.GetSharedAccessTokenPreparer(ctx, resourceGroupName, serviceName, UID, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.UsersClient", "GetSharedAccessToken", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSharedAccessTokenSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.UsersClient", "GetSharedAccessToken", resp, "Failure sending request")
		return
	}

	result, err = client.GetSharedAccessTokenResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.UsersClient", "GetSharedAccessToken", resp, "Failure responding to request")
	}

	return
}

// GetSharedAccessTokenPreparer prepares the GetSharedAccessToken request.
func (client UsersClient) GetSharedAccessTokenPreparer(ctx context.Context, resourceGroupName string, serviceName string, UID string, parameters UserTokenParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"uid":               autorest.Encode("path", UID),
	}

	const APIVersion = "2016-10-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/users/{uid}/token", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSharedAccessTokenSender sends the GetSharedAccessToken request. The method will close the
// http.Response Body if it receives an error.
func (client UsersClient) GetSharedAccessTokenSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetSharedAccessTokenResponder handles the response to the GetSharedAccessToken request. The method always
// closes the http.Response Body.
func (client UsersClient) GetSharedAccessTokenResponder(resp *http.Response) (result UserTokenResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByService lists a collection of registered users in the specified service instance.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// filter - | Field            | Supported operators    | Supported functions               |
// |------------------|------------------------|-----------------------------------|
// | id               | ge, le, eq, ne, gt, lt | substringof, contains, startswith, endswith |
// | firstName        | ge, le, eq, ne, gt, lt | substringof, contains, startswith, endswith |
// | lastName         | ge, le, eq, ne, gt, lt | substringof, contains, startswith, endswith |
// | email            | ge, le, eq, ne, gt, lt | substringof, contains, startswith, endswith |
// | state            | eq                     | N/A                               |
// | registrationDate | ge, le, eq, ne, gt, lt | N/A                               |
// | note             | ge, le, eq, ne, gt, lt | substringof, contains, startswith, endswith |
// top - number of records to return.
// skip - number of records to skip.
func (client UsersClient) ListByService(ctx context.Context, resourceGroupName string, serviceName string, filter string, top *int32, skip *int32) (result UserCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UsersClient.ListByService")
		defer func() {
			sc := -1
			if result.uc.Response.Response != nil {
				sc = result.uc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMinimum, Rule: int64(1), Chain: nil}}}}},
		{TargetValue: skip,
			Constraints: []validation.Constraint{{Target: "skip", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "skip", Name: validation.InclusiveMinimum, Rule: int64(0), Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("apimanagement.UsersClient", "ListByService", err.Error())
	}

	result.fn = client.listByServiceNextResults
	req, err := client.ListByServicePreparer(ctx, resourceGroupName, serviceName, filter, top, skip)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.UsersClient", "ListByService", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByServiceSender(req)
	if err != nil {
		result.uc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.UsersClient", "ListByService", resp, "Failure sending request")
		return
	}

	result.uc, err = client.ListByServiceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.UsersClient", "ListByService", resp, "Failure responding to request")
	}
	if result.uc.hasNextLink() && result.uc.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListByServicePreparer prepares the ListByService request.
func (client UsersClient) ListByServicePreparer(ctx context.Context, resourceGroupName string, serviceName string, filter string, top *int32, skip *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-10-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if skip != nil {
		queryParameters["$skip"] = autorest.Encode("query", *skip)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/users", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByServiceSender sends the ListByService request. The method will close the
// http.Response Body if it receives an error.
func (client UsersClient) ListByServiceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByServiceResponder handles the response to the ListByService request. The method always
// closes the http.Response Body.
func (client UsersClient) ListByServiceResponder(resp *http.Response) (result UserCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByServiceNextResults retrieves the next set of results, if any.
func (client UsersClient) listByServiceNextResults(ctx context.Context, lastResults UserCollection) (result UserCollection, err error) {
	req, err := lastResults.userCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "apimanagement.UsersClient", "listByServiceNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByServiceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "apimanagement.UsersClient", "listByServiceNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByServiceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.UsersClient", "listByServiceNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByServiceComplete enumerates all values, automatically crossing page boundaries as required.
func (client UsersClient) ListByServiceComplete(ctx context.Context, resourceGroupName string, serviceName string, filter string, top *int32, skip *int32) (result UserCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UsersClient.ListByService")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByService(ctx, resourceGroupName, serviceName, filter, top, skip)
	return
}

// Update updates the details of the user specified by its identifier.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// UID - user identifier. Must be unique in the current API Management service instance.
// parameters - update parameters.
// ifMatch - the entity state (Etag) version of the user to update. A value of "*" can be used for If-Match to
// unconditionally apply the operation.
func (client UsersClient) Update(ctx context.Context, resourceGroupName string, serviceName string, UID string, parameters UserUpdateParameters, ifMatch string) (result ErrorBodyContract, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UsersClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: UID,
			Constraints: []validation.Constraint{{Target: "UID", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "UID", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "UID", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.UsersClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, resourceGroupName, serviceName, UID, parameters, ifMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.UsersClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.UsersClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.UsersClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client UsersClient) UpdatePreparer(ctx context.Context, resourceGroupName string, serviceName string, UID string, parameters UserUpdateParameters, ifMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"uid":               autorest.Encode("path", UID),
	}

	const APIVersion = "2016-10-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/users/{uid}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client UsersClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client UsersClient) UpdateResponder(resp *http.Response) (result ErrorBodyContract, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent, http.StatusMethodNotAllowed),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
