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

// AuthorizationServerClient is the apiManagement Client
type AuthorizationServerClient struct {
	BaseClient
}

// NewAuthorizationServerClient creates an instance of the AuthorizationServerClient client.
func NewAuthorizationServerClient(subscriptionID string) AuthorizationServerClient {
	return NewAuthorizationServerClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewAuthorizationServerClientWithBaseURI creates an instance of the AuthorizationServerClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewAuthorizationServerClientWithBaseURI(baseURI string, subscriptionID string) AuthorizationServerClient {
	return AuthorizationServerClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates new authorization server or updates an existing authorization server.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// authsid - identifier of the authorization server.
// parameters - create or update parameters.
// ifMatch - eTag of the Entity. Not required when creating an entity, but required when updating an entity.
func (client AuthorizationServerClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, authsid string, parameters AuthorizationServerContract, ifMatch string) (result AuthorizationServerContract, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AuthorizationServerClient.CreateOrUpdate")
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
		{TargetValue: authsid,
			Constraints: []validation.Constraint{{Target: "authsid", Name: validation.MaxLength, Rule: 80, Chain: nil},
				{Target: "authsid", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "authsid", Name: validation.Pattern, Rule: `(^[\w]+$)|(^[\w][\w\-]+[\w]$)`, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.AuthorizationServerContractProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.AuthorizationServerContractProperties.DisplayName", Name: validation.Null, Rule: true,
					Chain: []validation.Constraint{{Target: "parameters.AuthorizationServerContractProperties.DisplayName", Name: validation.MaxLength, Rule: 50, Chain: nil},
						{Target: "parameters.AuthorizationServerContractProperties.DisplayName", Name: validation.MinLength, Rule: 1, Chain: nil},
					}},
					{Target: "parameters.AuthorizationServerContractProperties.ClientRegistrationEndpoint", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.AuthorizationServerContractProperties.AuthorizationEndpoint", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.AuthorizationServerContractProperties.GrantTypes", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.AuthorizationServerContractProperties.ClientID", Name: validation.Null, Rule: true, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("apimanagement.AuthorizationServerClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, serviceName, authsid, parameters, ifMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.AuthorizationServerClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.AuthorizationServerClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.AuthorizationServerClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client AuthorizationServerClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, serviceName string, authsid string, parameters AuthorizationServerContract, ifMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"authsid":           autorest.Encode("path", authsid),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/authorizationServers/{authsid}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	if len(ifMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client AuthorizationServerClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client AuthorizationServerClient) CreateOrUpdateResponder(resp *http.Response) (result AuthorizationServerContract, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes specific authorization server instance.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// authsid - identifier of the authorization server.
// ifMatch - eTag of the Entity. ETag should match the current entity state from the header response of the GET
// request or it should be * for unconditional update.
func (client AuthorizationServerClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, authsid string, ifMatch string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AuthorizationServerClient.Delete")
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
		{TargetValue: authsid,
			Constraints: []validation.Constraint{{Target: "authsid", Name: validation.MaxLength, Rule: 80, Chain: nil},
				{Target: "authsid", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "authsid", Name: validation.Pattern, Rule: `(^[\w]+$)|(^[\w][\w\-]+[\w]$)`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.AuthorizationServerClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, serviceName, authsid, ifMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.AuthorizationServerClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "apimanagement.AuthorizationServerClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.AuthorizationServerClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client AuthorizationServerClient) DeletePreparer(ctx context.Context, resourceGroupName string, serviceName string, authsid string, ifMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"authsid":           autorest.Encode("path", authsid),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/authorizationServers/{authsid}", pathParameters),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client AuthorizationServerClient) DeleteSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client AuthorizationServerClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the details of the authorization server specified by its identifier.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// authsid - identifier of the authorization server.
func (client AuthorizationServerClient) Get(ctx context.Context, resourceGroupName string, serviceName string, authsid string) (result AuthorizationServerContract, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AuthorizationServerClient.Get")
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
		{TargetValue: authsid,
			Constraints: []validation.Constraint{{Target: "authsid", Name: validation.MaxLength, Rule: 80, Chain: nil},
				{Target: "authsid", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "authsid", Name: validation.Pattern, Rule: `(^[\w]+$)|(^[\w][\w\-]+[\w]$)`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.AuthorizationServerClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, serviceName, authsid)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.AuthorizationServerClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.AuthorizationServerClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.AuthorizationServerClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client AuthorizationServerClient) GetPreparer(ctx context.Context, resourceGroupName string, serviceName string, authsid string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"authsid":           autorest.Encode("path", authsid),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/authorizationServers/{authsid}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client AuthorizationServerClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client AuthorizationServerClient) GetResponder(resp *http.Response) (result AuthorizationServerContract, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetEntityTag gets the entity state (Etag) version of the authorizationServer specified by its identifier.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// authsid - identifier of the authorization server.
func (client AuthorizationServerClient) GetEntityTag(ctx context.Context, resourceGroupName string, serviceName string, authsid string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AuthorizationServerClient.GetEntityTag")
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
		{TargetValue: authsid,
			Constraints: []validation.Constraint{{Target: "authsid", Name: validation.MaxLength, Rule: 80, Chain: nil},
				{Target: "authsid", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "authsid", Name: validation.Pattern, Rule: `(^[\w]+$)|(^[\w][\w\-]+[\w]$)`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.AuthorizationServerClient", "GetEntityTag", err.Error())
	}

	req, err := client.GetEntityTagPreparer(ctx, resourceGroupName, serviceName, authsid)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.AuthorizationServerClient", "GetEntityTag", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetEntityTagSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "apimanagement.AuthorizationServerClient", "GetEntityTag", resp, "Failure sending request")
		return
	}

	result, err = client.GetEntityTagResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.AuthorizationServerClient", "GetEntityTag", resp, "Failure responding to request")
	}

	return
}

// GetEntityTagPreparer prepares the GetEntityTag request.
func (client AuthorizationServerClient) GetEntityTagPreparer(ctx context.Context, resourceGroupName string, serviceName string, authsid string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"authsid":           autorest.Encode("path", authsid),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsHead(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/authorizationServers/{authsid}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetEntityTagSender sends the GetEntityTag request. The method will close the
// http.Response Body if it receives an error.
func (client AuthorizationServerClient) GetEntityTagSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetEntityTagResponder handles the response to the GetEntityTag request. The method always
// closes the http.Response Body.
func (client AuthorizationServerClient) GetEntityTagResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// ListByService lists a collection of authorization servers defined within a service instance.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// filter - | Field | Supported operators    | Supported functions                         |
// |-------|------------------------|---------------------------------------------|
// | id    | ge, le, eq, ne, gt, lt | substringof, contains, startswith, endswith |
// | name  | ge, le, eq, ne, gt, lt | substringof, contains, startswith, endswith |
// top - number of records to return.
// skip - number of records to skip.
func (client AuthorizationServerClient) ListByService(ctx context.Context, resourceGroupName string, serviceName string, filter string, top *int32, skip *int32) (result AuthorizationServerCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AuthorizationServerClient.ListByService")
		defer func() {
			sc := -1
			if result.asc.Response.Response != nil {
				sc = result.asc.Response.Response.StatusCode
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
		return result, validation.NewError("apimanagement.AuthorizationServerClient", "ListByService", err.Error())
	}

	result.fn = client.listByServiceNextResults
	req, err := client.ListByServicePreparer(ctx, resourceGroupName, serviceName, filter, top, skip)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.AuthorizationServerClient", "ListByService", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByServiceSender(req)
	if err != nil {
		result.asc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.AuthorizationServerClient", "ListByService", resp, "Failure sending request")
		return
	}

	result.asc, err = client.ListByServiceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.AuthorizationServerClient", "ListByService", resp, "Failure responding to request")
	}

	return
}

// ListByServicePreparer prepares the ListByService request.
func (client AuthorizationServerClient) ListByServicePreparer(ctx context.Context, resourceGroupName string, serviceName string, filter string, top *int32, skip *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-01-01"
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
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/authorizationServers", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByServiceSender sends the ListByService request. The method will close the
// http.Response Body if it receives an error.
func (client AuthorizationServerClient) ListByServiceSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListByServiceResponder handles the response to the ListByService request. The method always
// closes the http.Response Body.
func (client AuthorizationServerClient) ListByServiceResponder(resp *http.Response) (result AuthorizationServerCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByServiceNextResults retrieves the next set of results, if any.
func (client AuthorizationServerClient) listByServiceNextResults(ctx context.Context, lastResults AuthorizationServerCollection) (result AuthorizationServerCollection, err error) {
	req, err := lastResults.authorizationServerCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "apimanagement.AuthorizationServerClient", "listByServiceNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByServiceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "apimanagement.AuthorizationServerClient", "listByServiceNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByServiceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.AuthorizationServerClient", "listByServiceNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByServiceComplete enumerates all values, automatically crossing page boundaries as required.
func (client AuthorizationServerClient) ListByServiceComplete(ctx context.Context, resourceGroupName string, serviceName string, filter string, top *int32, skip *int32) (result AuthorizationServerCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AuthorizationServerClient.ListByService")
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

// Update updates the details of the authorization server specified by its identifier.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// authsid - identifier of the authorization server.
// parameters - oAuth2 Server settings Update parameters.
// ifMatch - eTag of the Entity. ETag should match the current entity state from the header response of the GET
// request or it should be * for unconditional update.
func (client AuthorizationServerClient) Update(ctx context.Context, resourceGroupName string, serviceName string, authsid string, parameters AuthorizationServerUpdateContract, ifMatch string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AuthorizationServerClient.Update")
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
		{TargetValue: authsid,
			Constraints: []validation.Constraint{{Target: "authsid", Name: validation.MaxLength, Rule: 80, Chain: nil},
				{Target: "authsid", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "authsid", Name: validation.Pattern, Rule: `(^[\w]+$)|(^[\w][\w\-]+[\w]$)`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.AuthorizationServerClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, resourceGroupName, serviceName, authsid, parameters, ifMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.AuthorizationServerClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "apimanagement.AuthorizationServerClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.AuthorizationServerClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client AuthorizationServerClient) UpdatePreparer(ctx context.Context, resourceGroupName string, serviceName string, authsid string, parameters AuthorizationServerUpdateContract, ifMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"authsid":           autorest.Encode("path", authsid),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/authorizationServers/{authsid}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client AuthorizationServerClient) UpdateSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client AuthorizationServerClient) UpdateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}
