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

// IntegrationRuntimeAuthKeysClient is the azure Synapse Analytics Management Client
type IntegrationRuntimeAuthKeysClient struct {
	BaseClient
}

// NewIntegrationRuntimeAuthKeysClient creates an instance of the IntegrationRuntimeAuthKeysClient client.
func NewIntegrationRuntimeAuthKeysClient(subscriptionID string) IntegrationRuntimeAuthKeysClient {
	return NewIntegrationRuntimeAuthKeysClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewIntegrationRuntimeAuthKeysClientWithBaseURI creates an instance of the IntegrationRuntimeAuthKeysClient client
// using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
func NewIntegrationRuntimeAuthKeysClientWithBaseURI(baseURI string, subscriptionID string) IntegrationRuntimeAuthKeysClient {
	return IntegrationRuntimeAuthKeysClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List list authentication keys in an integration runtime
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace.
// integrationRuntimeName - integration runtime name
func (client IntegrationRuntimeAuthKeysClient) List(ctx context.Context, resourceGroupName string, workspaceName string, integrationRuntimeName string) (result IntegrationRuntimeAuthKeys, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IntegrationRuntimeAuthKeysClient.List")
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
		return result, validation.NewError("synapse.IntegrationRuntimeAuthKeysClient", "List", err.Error())
	}

	req, err := client.ListPreparer(ctx, resourceGroupName, workspaceName, integrationRuntimeName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.IntegrationRuntimeAuthKeysClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "synapse.IntegrationRuntimeAuthKeysClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.IntegrationRuntimeAuthKeysClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client IntegrationRuntimeAuthKeysClient) ListPreparer(ctx context.Context, resourceGroupName string, workspaceName string, integrationRuntimeName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"integrationRuntimeName": autorest.Encode("path", integrationRuntimeName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
		"workspaceName":          autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2019-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/integrationRuntimes/{integrationRuntimeName}/listAuthKeys", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client IntegrationRuntimeAuthKeysClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client IntegrationRuntimeAuthKeysClient) ListResponder(resp *http.Response) (result IntegrationRuntimeAuthKeys, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Regenerate regenerate the authentication key for an integration runtime
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace.
// integrationRuntimeName - integration runtime name
// regenerateKeyParameters - the parameters for regenerating integration runtime authentication key.
func (client IntegrationRuntimeAuthKeysClient) Regenerate(ctx context.Context, resourceGroupName string, workspaceName string, integrationRuntimeName string, regenerateKeyParameters IntegrationRuntimeRegenerateKeyParameters) (result IntegrationRuntimeAuthKeys, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IntegrationRuntimeAuthKeysClient.Regenerate")
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
		return result, validation.NewError("synapse.IntegrationRuntimeAuthKeysClient", "Regenerate", err.Error())
	}

	req, err := client.RegeneratePreparer(ctx, resourceGroupName, workspaceName, integrationRuntimeName, regenerateKeyParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.IntegrationRuntimeAuthKeysClient", "Regenerate", nil, "Failure preparing request")
		return
	}

	resp, err := client.RegenerateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "synapse.IntegrationRuntimeAuthKeysClient", "Regenerate", resp, "Failure sending request")
		return
	}

	result, err = client.RegenerateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.IntegrationRuntimeAuthKeysClient", "Regenerate", resp, "Failure responding to request")
	}

	return
}

// RegeneratePreparer prepares the Regenerate request.
func (client IntegrationRuntimeAuthKeysClient) RegeneratePreparer(ctx context.Context, resourceGroupName string, workspaceName string, integrationRuntimeName string, regenerateKeyParameters IntegrationRuntimeRegenerateKeyParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"integrationRuntimeName": autorest.Encode("path", integrationRuntimeName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
		"workspaceName":          autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2019-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/integrationRuntimes/{integrationRuntimeName}/regenerateAuthKey", pathParameters),
		autorest.WithJSON(regenerateKeyParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RegenerateSender sends the Regenerate request. The method will close the
// http.Response Body if it receives an error.
func (client IntegrationRuntimeAuthKeysClient) RegenerateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// RegenerateResponder handles the response to the Regenerate request. The method always
// closes the http.Response Body.
func (client IntegrationRuntimeAuthKeysClient) RegenerateResponder(resp *http.Response) (result IntegrationRuntimeAuthKeys, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
