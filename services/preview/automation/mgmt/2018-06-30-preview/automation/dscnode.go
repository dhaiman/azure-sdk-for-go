package automation

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

// DscNodeClient is the automation Client
type DscNodeClient struct {
	BaseClient
}

// NewDscNodeClient creates an instance of the DscNodeClient client.
func NewDscNodeClient(subscriptionID string) DscNodeClient {
	return NewDscNodeClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewDscNodeClientWithBaseURI creates an instance of the DscNodeClient client.
func NewDscNodeClientWithBaseURI(baseURI string, subscriptionID string) DscNodeClient {
	return DscNodeClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Delete delete the dsc node identified by node id.
// Parameters:
// resourceGroupName - name of an Azure Resource group.
// automationAccountName - the name of the automation account.
// nodeID - the node id.
func (client DscNodeClient) Delete(ctx context.Context, resourceGroupName string, automationAccountName string, nodeID string) (result DscNode, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DscNodeClient.Delete")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("automation.DscNodeClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, automationAccountName, nodeID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.DscNodeClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.DscNodeClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.DscNodeClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client DscNodeClient) DeletePreparer(ctx context.Context, resourceGroupName string, automationAccountName string, nodeID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"nodeId":                autorest.Encode("path", nodeID),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-01-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/nodes/{nodeId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client DscNodeClient) DeleteSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client DscNodeClient) DeleteResponder(resp *http.Response) (result DscNode, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get retrieve the dsc node identified by node id.
// Parameters:
// resourceGroupName - name of an Azure Resource group.
// automationAccountName - the name of the automation account.
// nodeID - the node id.
func (client DscNodeClient) Get(ctx context.Context, resourceGroupName string, automationAccountName string, nodeID string) (result DscNode, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DscNodeClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("automation.DscNodeClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, automationAccountName, nodeID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.DscNodeClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.DscNodeClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.DscNodeClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client DscNodeClient) GetPreparer(ctx context.Context, resourceGroupName string, automationAccountName string, nodeID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"nodeId":                autorest.Encode("path", nodeID),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-01-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/nodes/{nodeId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client DscNodeClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client DscNodeClient) GetResponder(resp *http.Response) (result DscNode, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByAutomationAccount retrieve a list of dsc nodes.
// Parameters:
// resourceGroupName - name of an Azure Resource group.
// automationAccountName - the name of the automation account.
// filter - the filter to apply on the operation.
// skip - the number of rows to skip.
// top - the number of rows to take.
// inlinecount - return total rows.
func (client DscNodeClient) ListByAutomationAccount(ctx context.Context, resourceGroupName string, automationAccountName string, filter string, skip *int32, top *int32, inlinecount string) (result DscNodeListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DscNodeClient.ListByAutomationAccount")
		defer func() {
			sc := -1
			if result.dnlr.Response.Response != nil {
				sc = result.dnlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("automation.DscNodeClient", "ListByAutomationAccount", err.Error())
	}

	result.fn = client.listByAutomationAccountNextResults
	req, err := client.ListByAutomationAccountPreparer(ctx, resourceGroupName, automationAccountName, filter, skip, top, inlinecount)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.DscNodeClient", "ListByAutomationAccount", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByAutomationAccountSender(req)
	if err != nil {
		result.dnlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.DscNodeClient", "ListByAutomationAccount", resp, "Failure sending request")
		return
	}

	result.dnlr, err = client.ListByAutomationAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.DscNodeClient", "ListByAutomationAccount", resp, "Failure responding to request")
	}

	return
}

// ListByAutomationAccountPreparer prepares the ListByAutomationAccount request.
func (client DscNodeClient) ListByAutomationAccountPreparer(ctx context.Context, resourceGroupName string, automationAccountName string, filter string, skip *int32, top *int32, inlinecount string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-01-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if skip != nil {
		queryParameters["$skip"] = autorest.Encode("query", *skip)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(inlinecount) > 0 {
		queryParameters["$inlinecount"] = autorest.Encode("query", inlinecount)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/nodes", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByAutomationAccountSender sends the ListByAutomationAccount request. The method will close the
// http.Response Body if it receives an error.
func (client DscNodeClient) ListByAutomationAccountSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListByAutomationAccountResponder handles the response to the ListByAutomationAccount request. The method always
// closes the http.Response Body.
func (client DscNodeClient) ListByAutomationAccountResponder(resp *http.Response) (result DscNodeListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByAutomationAccountNextResults retrieves the next set of results, if any.
func (client DscNodeClient) listByAutomationAccountNextResults(ctx context.Context, lastResults DscNodeListResult) (result DscNodeListResult, err error) {
	req, err := lastResults.dscNodeListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "automation.DscNodeClient", "listByAutomationAccountNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByAutomationAccountSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "automation.DscNodeClient", "listByAutomationAccountNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByAutomationAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.DscNodeClient", "listByAutomationAccountNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByAutomationAccountComplete enumerates all values, automatically crossing page boundaries as required.
func (client DscNodeClient) ListByAutomationAccountComplete(ctx context.Context, resourceGroupName string, automationAccountName string, filter string, skip *int32, top *int32, inlinecount string) (result DscNodeListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DscNodeClient.ListByAutomationAccount")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByAutomationAccount(ctx, resourceGroupName, automationAccountName, filter, skip, top, inlinecount)
	return
}

// Update update the dsc node.
// Parameters:
// resourceGroupName - name of an Azure Resource group.
// automationAccountName - the name of the automation account.
// nodeID - parameters supplied to the update dsc node.
// dscNodeUpdateParameters - parameters supplied to the update dsc node.
func (client DscNodeClient) Update(ctx context.Context, resourceGroupName string, automationAccountName string, nodeID string, dscNodeUpdateParameters DscNodeUpdateParameters) (result DscNode, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DscNodeClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("automation.DscNodeClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, resourceGroupName, automationAccountName, nodeID, dscNodeUpdateParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.DscNodeClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "automation.DscNodeClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automation.DscNodeClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client DscNodeClient) UpdatePreparer(ctx context.Context, resourceGroupName string, automationAccountName string, nodeID string, dscNodeUpdateParameters DscNodeUpdateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"automationAccountName": autorest.Encode("path", automationAccountName),
		"nodeId":                autorest.Encode("path", nodeID),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-01-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/nodes/{nodeId}", pathParameters),
		autorest.WithJSON(dscNodeUpdateParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client DscNodeClient) UpdateSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client DscNodeClient) UpdateResponder(resp *http.Response) (result DscNode, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
