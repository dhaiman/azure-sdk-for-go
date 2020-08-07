package authorization

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
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// RoleDefinitionsClient is the client for the RoleDefinitions methods of the Authorization service.
type RoleDefinitionsClient struct {
	BaseClient
}

// NewRoleDefinitionsClient creates an instance of the RoleDefinitionsClient client.
func NewRoleDefinitionsClient(subscriptionID string) RoleDefinitionsClient {
	return NewRoleDefinitionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewRoleDefinitionsClientWithBaseURI creates an instance of the RoleDefinitionsClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewRoleDefinitionsClientWithBaseURI(baseURI string, subscriptionID string) RoleDefinitionsClient {
	return RoleDefinitionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a role definition.
// Parameters:
// scope - the scope of the role definition.
// roleDefinitionID - the ID of the role definition.
// roleDefinition - the values for the role definition.
func (client RoleDefinitionsClient) CreateOrUpdate(ctx context.Context, scope string, roleDefinitionID string, roleDefinition RoleDefinition) (result RoleDefinition, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RoleDefinitionsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, scope, roleDefinitionID, roleDefinition)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.RoleDefinitionsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "authorization.RoleDefinitionsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.RoleDefinitionsClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client RoleDefinitionsClient) CreateOrUpdatePreparer(ctx context.Context, scope string, roleDefinitionID string, roleDefinition RoleDefinition) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"roleDefinitionId": autorest.Encode("path", roleDefinitionID),
		"scope":            scope,
	}

	const APIVersion = "2018-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	roleDefinition.ID = nil
	roleDefinition.Name = nil
	roleDefinition.Type = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.Authorization/roleDefinitions/{roleDefinitionId}", pathParameters),
		autorest.WithJSON(roleDefinition),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client RoleDefinitionsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client RoleDefinitionsClient) CreateOrUpdateResponder(resp *http.Response) (result RoleDefinition, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a role definition.
// Parameters:
// scope - the scope of the role definition.
// roleDefinitionID - the ID of the role definition to delete.
func (client RoleDefinitionsClient) Delete(ctx context.Context, scope string, roleDefinitionID string) (result RoleDefinition, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RoleDefinitionsClient.Delete")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, scope, roleDefinitionID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.RoleDefinitionsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "authorization.RoleDefinitionsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.RoleDefinitionsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client RoleDefinitionsClient) DeletePreparer(ctx context.Context, scope string, roleDefinitionID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"roleDefinitionId": autorest.Encode("path", roleDefinitionID),
		"scope":            scope,
	}

	const APIVersion = "2018-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.Authorization/roleDefinitions/{roleDefinitionId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client RoleDefinitionsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client RoleDefinitionsClient) DeleteResponder(resp *http.Response) (result RoleDefinition, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get get role definition by name (GUID).
// Parameters:
// scope - the scope of the role definition.
// roleDefinitionID - the ID of the role definition.
func (client RoleDefinitionsClient) Get(ctx context.Context, scope string, roleDefinitionID string) (result RoleDefinition, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RoleDefinitionsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, scope, roleDefinitionID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.RoleDefinitionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "authorization.RoleDefinitionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.RoleDefinitionsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client RoleDefinitionsClient) GetPreparer(ctx context.Context, scope string, roleDefinitionID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"roleDefinitionId": autorest.Encode("path", roleDefinitionID),
		"scope":            scope,
	}

	const APIVersion = "2018-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.Authorization/roleDefinitions/{roleDefinitionId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client RoleDefinitionsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client RoleDefinitionsClient) GetResponder(resp *http.Response) (result RoleDefinition, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetByID gets a role definition by ID.
// Parameters:
// roleID - the fully qualified role definition ID. Use the format,
// /subscriptions/{guid}/providers/Microsoft.Authorization/roleDefinitions/{roleDefinitionId} for subscription
// level role definitions, or /providers/Microsoft.Authorization/roleDefinitions/{roleDefinitionId} for tenant
// level role definitions.
func (client RoleDefinitionsClient) GetByID(ctx context.Context, roleID string) (result RoleDefinition, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RoleDefinitionsClient.GetByID")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetByIDPreparer(ctx, roleID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.RoleDefinitionsClient", "GetByID", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetByIDSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "authorization.RoleDefinitionsClient", "GetByID", resp, "Failure sending request")
		return
	}

	result, err = client.GetByIDResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.RoleDefinitionsClient", "GetByID", resp, "Failure responding to request")
	}

	return
}

// GetByIDPreparer prepares the GetByID request.
func (client RoleDefinitionsClient) GetByIDPreparer(ctx context.Context, roleID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"roleId": roleID,
	}

	const APIVersion = "2018-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{roleId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetByIDSender sends the GetByID request. The method will close the
// http.Response Body if it receives an error.
func (client RoleDefinitionsClient) GetByIDSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetByIDResponder handles the response to the GetByID request. The method always
// closes the http.Response Body.
func (client RoleDefinitionsClient) GetByIDResponder(resp *http.Response) (result RoleDefinition, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List get all role definitions that are applicable at scope and above.
// Parameters:
// scope - the scope of the role definition.
// filter - the filter to apply on the operation. Use atScopeAndBelow filter to search below the given scope as
// well.
func (client RoleDefinitionsClient) List(ctx context.Context, scope string, filter string) (result RoleDefinitionListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RoleDefinitionsClient.List")
		defer func() {
			sc := -1
			if result.rdlr.Response.Response != nil {
				sc = result.rdlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, scope, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.RoleDefinitionsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.rdlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "authorization.RoleDefinitionsClient", "List", resp, "Failure sending request")
		return
	}

	result.rdlr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.RoleDefinitionsClient", "List", resp, "Failure responding to request")
	}
	if result.rdlr.hasNextLink() && result.rdlr.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListPreparer prepares the List request.
func (client RoleDefinitionsClient) ListPreparer(ctx context.Context, scope string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"scope": scope,
	}

	const APIVersion = "2018-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.Authorization/roleDefinitions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client RoleDefinitionsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client RoleDefinitionsClient) ListResponder(resp *http.Response) (result RoleDefinitionListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client RoleDefinitionsClient) listNextResults(ctx context.Context, lastResults RoleDefinitionListResult) (result RoleDefinitionListResult, err error) {
	req, err := lastResults.roleDefinitionListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "authorization.RoleDefinitionsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "authorization.RoleDefinitionsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.RoleDefinitionsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client RoleDefinitionsClient) ListComplete(ctx context.Context, scope string, filter string) (result RoleDefinitionListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RoleDefinitionsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, scope, filter)
	return
}
