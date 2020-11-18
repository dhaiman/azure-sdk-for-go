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

// IPFirewallRulesClient is the azure Synapse Analytics Management Client
type IPFirewallRulesClient struct {
	BaseClient
}

// NewIPFirewallRulesClient creates an instance of the IPFirewallRulesClient client.
func NewIPFirewallRulesClient(subscriptionID string) IPFirewallRulesClient {
	return NewIPFirewallRulesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewIPFirewallRulesClientWithBaseURI creates an instance of the IPFirewallRulesClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewIPFirewallRulesClientWithBaseURI(baseURI string, subscriptionID string) IPFirewallRulesClient {
	return IPFirewallRulesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a firewall rule
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace
// ruleName - the IP firewall rule name
// IPFirewallRuleInfo - IP firewall rule properties
func (client IPFirewallRulesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, ruleName string, IPFirewallRuleInfo IPFirewallRuleInfo) (result IPFirewallRulesCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IPFirewallRulesClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
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
		return result, validation.NewError("synapse.IPFirewallRulesClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, workspaceName, ruleName, IPFirewallRuleInfo)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.IPFirewallRulesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.IPFirewallRulesClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client IPFirewallRulesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, workspaceName string, ruleName string, IPFirewallRuleInfo IPFirewallRuleInfo) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"ruleName":          autorest.Encode("path", ruleName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2020-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/firewallRules/{ruleName}", pathParameters),
		autorest.WithJSON(IPFirewallRuleInfo),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client IPFirewallRulesClient) CreateOrUpdateSender(req *http.Request) (future IPFirewallRulesCreateOrUpdateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client IPFirewallRulesClient) CreateOrUpdateResponder(resp *http.Response) (result IPFirewallRuleInfo, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a firewall rule
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace
// ruleName - the IP firewall rule name
func (client IPFirewallRulesClient) Delete(ctx context.Context, resourceGroupName string, workspaceName string, ruleName string) (result IPFirewallRulesDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IPFirewallRulesClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
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
		return result, validation.NewError("synapse.IPFirewallRulesClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, workspaceName, ruleName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.IPFirewallRulesClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.IPFirewallRulesClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client IPFirewallRulesClient) DeletePreparer(ctx context.Context, resourceGroupName string, workspaceName string, ruleName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"ruleName":          autorest.Encode("path", ruleName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2020-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/firewallRules/{ruleName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client IPFirewallRulesClient) DeleteSender(req *http.Request) (future IPFirewallRulesDeleteFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client IPFirewallRulesClient) DeleteResponder(resp *http.Response) (result SetObject, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get get a firewall rule
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace
// ruleName - the IP firewall rule name
func (client IPFirewallRulesClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, ruleName string) (result IPFirewallRuleInfo, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IPFirewallRulesClient.Get")
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
		return result, validation.NewError("synapse.IPFirewallRulesClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, workspaceName, ruleName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.IPFirewallRulesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "synapse.IPFirewallRulesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.IPFirewallRulesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client IPFirewallRulesClient) GetPreparer(ctx context.Context, resourceGroupName string, workspaceName string, ruleName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"ruleName":          autorest.Encode("path", ruleName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2020-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/firewallRules/{ruleName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client IPFirewallRulesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client IPFirewallRulesClient) GetResponder(resp *http.Response) (result IPFirewallRuleInfo, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByWorkspace returns a list of firewall rules
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace
func (client IPFirewallRulesClient) ListByWorkspace(ctx context.Context, resourceGroupName string, workspaceName string) (result IPFirewallRuleInfoListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IPFirewallRulesClient.ListByWorkspace")
		defer func() {
			sc := -1
			if result.ifrilr.Response.Response != nil {
				sc = result.ifrilr.Response.Response.StatusCode
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
		return result, validation.NewError("synapse.IPFirewallRulesClient", "ListByWorkspace", err.Error())
	}

	result.fn = client.listByWorkspaceNextResults
	req, err := client.ListByWorkspacePreparer(ctx, resourceGroupName, workspaceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.IPFirewallRulesClient", "ListByWorkspace", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByWorkspaceSender(req)
	if err != nil {
		result.ifrilr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "synapse.IPFirewallRulesClient", "ListByWorkspace", resp, "Failure sending request")
		return
	}

	result.ifrilr, err = client.ListByWorkspaceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.IPFirewallRulesClient", "ListByWorkspace", resp, "Failure responding to request")
	}
	if result.ifrilr.hasNextLink() && result.ifrilr.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListByWorkspacePreparer prepares the ListByWorkspace request.
func (client IPFirewallRulesClient) ListByWorkspacePreparer(ctx context.Context, resourceGroupName string, workspaceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2020-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/firewallRules", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByWorkspaceSender sends the ListByWorkspace request. The method will close the
// http.Response Body if it receives an error.
func (client IPFirewallRulesClient) ListByWorkspaceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByWorkspaceResponder handles the response to the ListByWorkspace request. The method always
// closes the http.Response Body.
func (client IPFirewallRulesClient) ListByWorkspaceResponder(resp *http.Response) (result IPFirewallRuleInfoListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByWorkspaceNextResults retrieves the next set of results, if any.
func (client IPFirewallRulesClient) listByWorkspaceNextResults(ctx context.Context, lastResults IPFirewallRuleInfoListResult) (result IPFirewallRuleInfoListResult, err error) {
	req, err := lastResults.iPFirewallRuleInfoListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "synapse.IPFirewallRulesClient", "listByWorkspaceNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByWorkspaceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "synapse.IPFirewallRulesClient", "listByWorkspaceNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByWorkspaceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.IPFirewallRulesClient", "listByWorkspaceNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByWorkspaceComplete enumerates all values, automatically crossing page boundaries as required.
func (client IPFirewallRulesClient) ListByWorkspaceComplete(ctx context.Context, resourceGroupName string, workspaceName string) (result IPFirewallRuleInfoListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IPFirewallRulesClient.ListByWorkspace")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByWorkspace(ctx, resourceGroupName, workspaceName)
	return
}

// ReplaceAll replaces firewall rules
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace
// request - replace all IP firewall rules request
func (client IPFirewallRulesClient) ReplaceAll(ctx context.Context, resourceGroupName string, workspaceName string, request ReplaceAllIPFirewallRulesRequest) (result IPFirewallRulesReplaceAllFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IPFirewallRulesClient.ReplaceAll")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
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
		return result, validation.NewError("synapse.IPFirewallRulesClient", "ReplaceAll", err.Error())
	}

	req, err := client.ReplaceAllPreparer(ctx, resourceGroupName, workspaceName, request)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.IPFirewallRulesClient", "ReplaceAll", nil, "Failure preparing request")
		return
	}

	result, err = client.ReplaceAllSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.IPFirewallRulesClient", "ReplaceAll", result.Response(), "Failure sending request")
		return
	}

	return
}

// ReplaceAllPreparer prepares the ReplaceAll request.
func (client IPFirewallRulesClient) ReplaceAllPreparer(ctx context.Context, resourceGroupName string, workspaceName string, request ReplaceAllIPFirewallRulesRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2020-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/replaceAllIpFirewallRules", pathParameters),
		autorest.WithJSON(request),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ReplaceAllSender sends the ReplaceAll request. The method will close the
// http.Response Body if it receives an error.
func (client IPFirewallRulesClient) ReplaceAllSender(req *http.Request) (future IPFirewallRulesReplaceAllFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// ReplaceAllResponder handles the response to the ReplaceAll request. The method always
// closes the http.Response Body.
func (client IPFirewallRulesClient) ReplaceAllResponder(resp *http.Response) (result ReplaceAllFirewallRulesOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
