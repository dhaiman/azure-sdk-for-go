package mysql

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

// MaintenanceWindowsClient is the the Microsoft Azure management API provides create, read, update, and delete
// functionality for Azure MySQL resources including servers, databases, firewall rules, VNET rules, log files and
// configurations with new business model.
type MaintenanceWindowsClient struct {
	BaseClient
}

// NewMaintenanceWindowsClient creates an instance of the MaintenanceWindowsClient client.
func NewMaintenanceWindowsClient(subscriptionID string) MaintenanceWindowsClient {
	return NewMaintenanceWindowsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewMaintenanceWindowsClientWithBaseURI creates an instance of the MaintenanceWindowsClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewMaintenanceWindowsClientWithBaseURI(baseURI string, subscriptionID string) MaintenanceWindowsClient {
	return MaintenanceWindowsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates a new maintenance window.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// serverName - the name of the server.
// maintenanceWindowName - the name of the maintenance window.
// parameters - the required parameters for creating or updating a customer maintenance window.
func (client MaintenanceWindowsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, maintenanceWindowName string, parameters MaintenanceWindow) (result MaintenanceWindowsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MaintenanceWindowsClient.CreateOrUpdate")
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
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.MaintenanceWindowProperties", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "parameters.MaintenanceWindowProperties.DayOfWeek", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.MaintenanceWindowProperties.StartHour", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.MaintenanceWindowProperties.StartMinute", Name: validation.Null, Rule: true, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("mysql.MaintenanceWindowsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, serverName, maintenanceWindowName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mysql.MaintenanceWindowsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mysql.MaintenanceWindowsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client MaintenanceWindowsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, serverName string, maintenanceWindowName string, parameters MaintenanceWindow) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"maintenanceWindowName": autorest.Encode("path", maintenanceWindowName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"serverName":            autorest.Encode("path", serverName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-07-01-privatepreview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBForMySql/flexibleServers/{serverName}/maintenanceWindows/{maintenanceWindowName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client MaintenanceWindowsClient) CreateOrUpdateSender(req *http.Request) (future MaintenanceWindowsCreateOrUpdateFuture, err error) {
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
func (client MaintenanceWindowsClient) CreateOrUpdateResponder(resp *http.Response) (result MaintenanceWindow, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a maintenance window.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// serverName - the name of the server.
// maintenanceWindowName - the name of the maintenance window.
func (client MaintenanceWindowsClient) Delete(ctx context.Context, resourceGroupName string, serverName string, maintenanceWindowName string) (result MaintenanceWindowsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MaintenanceWindowsClient.Delete")
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
		return result, validation.NewError("mysql.MaintenanceWindowsClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, serverName, maintenanceWindowName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mysql.MaintenanceWindowsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mysql.MaintenanceWindowsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client MaintenanceWindowsClient) DeletePreparer(ctx context.Context, resourceGroupName string, serverName string, maintenanceWindowName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"maintenanceWindowName": autorest.Encode("path", maintenanceWindowName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"serverName":            autorest.Encode("path", serverName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-07-01-privatepreview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBForMySql/flexibleServers/{serverName}/maintenanceWindows/{maintenanceWindowName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client MaintenanceWindowsClient) DeleteSender(req *http.Request) (future MaintenanceWindowsDeleteFuture, err error) {
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
func (client MaintenanceWindowsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get the maintenance windows of a given server.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// serverName - the name of the server.
// maintenanceWindowName - the name of the maintenance window.
func (client MaintenanceWindowsClient) Get(ctx context.Context, resourceGroupName string, serverName string, maintenanceWindowName string) (result MaintenanceWindow, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MaintenanceWindowsClient.Get")
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
		return result, validation.NewError("mysql.MaintenanceWindowsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, serverName, maintenanceWindowName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mysql.MaintenanceWindowsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "mysql.MaintenanceWindowsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mysql.MaintenanceWindowsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client MaintenanceWindowsClient) GetPreparer(ctx context.Context, resourceGroupName string, serverName string, maintenanceWindowName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"maintenanceWindowName": autorest.Encode("path", maintenanceWindowName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"serverName":            autorest.Encode("path", serverName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-07-01-privatepreview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBForMySql/flexibleServers/{serverName}/maintenanceWindows/{maintenanceWindowName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client MaintenanceWindowsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client MaintenanceWindowsClient) GetResponder(resp *http.Response) (result MaintenanceWindow, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByServer list all the maintenance windows in a given server.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// serverName - the name of the server.
func (client MaintenanceWindowsClient) ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result MaintenanceWindowListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MaintenanceWindowsClient.ListByServer")
		defer func() {
			sc := -1
			if result.mwlr.Response.Response != nil {
				sc = result.mwlr.Response.Response.StatusCode
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
		return result, validation.NewError("mysql.MaintenanceWindowsClient", "ListByServer", err.Error())
	}

	result.fn = client.listByServerNextResults
	req, err := client.ListByServerPreparer(ctx, resourceGroupName, serverName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mysql.MaintenanceWindowsClient", "ListByServer", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByServerSender(req)
	if err != nil {
		result.mwlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "mysql.MaintenanceWindowsClient", "ListByServer", resp, "Failure sending request")
		return
	}

	result.mwlr, err = client.ListByServerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mysql.MaintenanceWindowsClient", "ListByServer", resp, "Failure responding to request")
	}
	if result.mwlr.hasNextLink() && result.mwlr.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListByServerPreparer prepares the ListByServer request.
func (client MaintenanceWindowsClient) ListByServerPreparer(ctx context.Context, resourceGroupName string, serverName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-07-01-privatepreview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBForMySql/flexibleServers/{serverName}/maintenanceWindows", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByServerSender sends the ListByServer request. The method will close the
// http.Response Body if it receives an error.
func (client MaintenanceWindowsClient) ListByServerSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByServerResponder handles the response to the ListByServer request. The method always
// closes the http.Response Body.
func (client MaintenanceWindowsClient) ListByServerResponder(resp *http.Response) (result MaintenanceWindowListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByServerNextResults retrieves the next set of results, if any.
func (client MaintenanceWindowsClient) listByServerNextResults(ctx context.Context, lastResults MaintenanceWindowListResult) (result MaintenanceWindowListResult, err error) {
	req, err := lastResults.maintenanceWindowListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "mysql.MaintenanceWindowsClient", "listByServerNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByServerSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "mysql.MaintenanceWindowsClient", "listByServerNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByServerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mysql.MaintenanceWindowsClient", "listByServerNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByServerComplete enumerates all values, automatically crossing page boundaries as required.
func (client MaintenanceWindowsClient) ListByServerComplete(ctx context.Context, resourceGroupName string, serverName string) (result MaintenanceWindowListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MaintenanceWindowsClient.ListByServer")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByServer(ctx, resourceGroupName, serverName)
	return
}
