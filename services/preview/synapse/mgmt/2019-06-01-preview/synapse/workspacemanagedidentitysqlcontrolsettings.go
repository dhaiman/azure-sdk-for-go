package synapse

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

// WorkspaceManagedIdentitySQLControlSettingsClient is the azure Synapse Analytics Management Client
type WorkspaceManagedIdentitySQLControlSettingsClient struct {
	BaseClient
}

// NewWorkspaceManagedIdentitySQLControlSettingsClient creates an instance of the
// WorkspaceManagedIdentitySQLControlSettingsClient client.
func NewWorkspaceManagedIdentitySQLControlSettingsClient(subscriptionID string) WorkspaceManagedIdentitySQLControlSettingsClient {
	return NewWorkspaceManagedIdentitySQLControlSettingsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWorkspaceManagedIdentitySQLControlSettingsClientWithBaseURI creates an instance of the
// WorkspaceManagedIdentitySQLControlSettingsClient client using a custom endpoint.  Use this when interacting with an
// Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewWorkspaceManagedIdentitySQLControlSettingsClientWithBaseURI(baseURI string, subscriptionID string) WorkspaceManagedIdentitySQLControlSettingsClient {
	return WorkspaceManagedIdentitySQLControlSettingsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate sends the create or update request.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace
// managedIdentitySQLControlSettings - managed Identity Sql Control Settings
func (client WorkspaceManagedIdentitySQLControlSettingsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, managedIdentitySQLControlSettings ManagedIdentitySQLControlSettingsModel) (result ManagedIdentitySQLControlSettingsModel, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WorkspaceManagedIdentitySQLControlSettingsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("synapse.WorkspaceManagedIdentitySQLControlSettingsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, workspaceName, managedIdentitySQLControlSettings)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.WorkspaceManagedIdentitySQLControlSettingsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "synapse.WorkspaceManagedIdentitySQLControlSettingsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.WorkspaceManagedIdentitySQLControlSettingsClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client WorkspaceManagedIdentitySQLControlSettingsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, workspaceName string, managedIdentitySQLControlSettings ManagedIdentitySQLControlSettingsModel) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2019-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/managedIdentitySqlControlSettings/default", pathParameters),
		autorest.WithJSON(managedIdentitySQLControlSettings),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client WorkspaceManagedIdentitySQLControlSettingsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client WorkspaceManagedIdentitySQLControlSettingsClient) CreateOrUpdateResponder(resp *http.Response) (result ManagedIdentitySQLControlSettingsModel, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get sends the get request.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace
func (client WorkspaceManagedIdentitySQLControlSettingsClient) Get(ctx context.Context, resourceGroupName string, workspaceName string) (result ManagedIdentitySQLControlSettingsModel, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WorkspaceManagedIdentitySQLControlSettingsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("synapse.WorkspaceManagedIdentitySQLControlSettingsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, workspaceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.WorkspaceManagedIdentitySQLControlSettingsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "synapse.WorkspaceManagedIdentitySQLControlSettingsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.WorkspaceManagedIdentitySQLControlSettingsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client WorkspaceManagedIdentitySQLControlSettingsClient) GetPreparer(ctx context.Context, resourceGroupName string, workspaceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2019-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/managedIdentitySqlControlSettings/default", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client WorkspaceManagedIdentitySQLControlSettingsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client WorkspaceManagedIdentitySQLControlSettingsClient) GetResponder(resp *http.Response) (result ManagedIdentitySQLControlSettingsModel, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
