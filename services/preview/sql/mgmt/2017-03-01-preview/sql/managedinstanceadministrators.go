package sql

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

// ManagedInstanceAdministratorsClient is the the Azure SQL Database management API provides a RESTful set of web
// services that interact with Azure SQL Database services to manage your databases. The API enables you to create,
// retrieve, update, and delete databases.
type ManagedInstanceAdministratorsClient struct {
	BaseClient
}

// NewManagedInstanceAdministratorsClient creates an instance of the ManagedInstanceAdministratorsClient client.
func NewManagedInstanceAdministratorsClient(subscriptionID string) ManagedInstanceAdministratorsClient {
	return NewManagedInstanceAdministratorsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewManagedInstanceAdministratorsClientWithBaseURI creates an instance of the ManagedInstanceAdministratorsClient
// client using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI
// (sovereign clouds, Azure stack).
func NewManagedInstanceAdministratorsClientWithBaseURI(baseURI string, subscriptionID string) ManagedInstanceAdministratorsClient {
	return ManagedInstanceAdministratorsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a managed instance administrator.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// managedInstanceName - the name of the managed instance.
// parameters - the requested administrator parameters.
func (client ManagedInstanceAdministratorsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, managedInstanceName string, parameters ManagedInstanceAdministrator) (result ManagedInstanceAdministratorsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ManagedInstanceAdministratorsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.ManagedInstanceAdministratorProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.ManagedInstanceAdministratorProperties.AdministratorType", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.ManagedInstanceAdministratorProperties.Login", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.ManagedInstanceAdministratorProperties.Sid", Name: validation.Null, Rule: true, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("sql.ManagedInstanceAdministratorsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, managedInstanceName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedInstanceAdministratorsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedInstanceAdministratorsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ManagedInstanceAdministratorsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, managedInstanceName string, parameters ManagedInstanceAdministrator) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"administratorName":   autorest.Encode("path", "ActiveDirectory"),
		"managedInstanceName": autorest.Encode("path", managedInstanceName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/administrators/{administratorName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ManagedInstanceAdministratorsClient) CreateOrUpdateSender(req *http.Request) (future ManagedInstanceAdministratorsCreateOrUpdateFuture, err error) {
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
func (client ManagedInstanceAdministratorsClient) CreateOrUpdateResponder(resp *http.Response) (result ManagedInstanceAdministrator, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a managed instance administrator.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// managedInstanceName - the name of the managed instance.
func (client ManagedInstanceAdministratorsClient) Delete(ctx context.Context, resourceGroupName string, managedInstanceName string) (result ManagedInstanceAdministratorsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ManagedInstanceAdministratorsClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, managedInstanceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedInstanceAdministratorsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedInstanceAdministratorsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ManagedInstanceAdministratorsClient) DeletePreparer(ctx context.Context, resourceGroupName string, managedInstanceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"administratorName":   autorest.Encode("path", "ActiveDirectory"),
		"managedInstanceName": autorest.Encode("path", managedInstanceName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/administrators/{administratorName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ManagedInstanceAdministratorsClient) DeleteSender(req *http.Request) (future ManagedInstanceAdministratorsDeleteFuture, err error) {
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
func (client ManagedInstanceAdministratorsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets a managed instance administrator.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// managedInstanceName - the name of the managed instance.
func (client ManagedInstanceAdministratorsClient) Get(ctx context.Context, resourceGroupName string, managedInstanceName string) (result ManagedInstanceAdministrator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ManagedInstanceAdministratorsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, managedInstanceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedInstanceAdministratorsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.ManagedInstanceAdministratorsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedInstanceAdministratorsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ManagedInstanceAdministratorsClient) GetPreparer(ctx context.Context, resourceGroupName string, managedInstanceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"administratorName":   autorest.Encode("path", "ActiveDirectory"),
		"managedInstanceName": autorest.Encode("path", managedInstanceName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/administrators/{administratorName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ManagedInstanceAdministratorsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ManagedInstanceAdministratorsClient) GetResponder(resp *http.Response) (result ManagedInstanceAdministrator, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByInstance gets a list of managed instance administrators.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// managedInstanceName - the name of the managed instance.
func (client ManagedInstanceAdministratorsClient) ListByInstance(ctx context.Context, resourceGroupName string, managedInstanceName string) (result ManagedInstanceAdministratorListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ManagedInstanceAdministratorsClient.ListByInstance")
		defer func() {
			sc := -1
			if result.mialr.Response.Response != nil {
				sc = result.mialr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByInstanceNextResults
	req, err := client.ListByInstancePreparer(ctx, resourceGroupName, managedInstanceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedInstanceAdministratorsClient", "ListByInstance", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByInstanceSender(req)
	if err != nil {
		result.mialr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.ManagedInstanceAdministratorsClient", "ListByInstance", resp, "Failure sending request")
		return
	}

	result.mialr, err = client.ListByInstanceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedInstanceAdministratorsClient", "ListByInstance", resp, "Failure responding to request")
	}

	return
}

// ListByInstancePreparer prepares the ListByInstance request.
func (client ManagedInstanceAdministratorsClient) ListByInstancePreparer(ctx context.Context, resourceGroupName string, managedInstanceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"managedInstanceName": autorest.Encode("path", managedInstanceName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/administrators", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByInstanceSender sends the ListByInstance request. The method will close the
// http.Response Body if it receives an error.
func (client ManagedInstanceAdministratorsClient) ListByInstanceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByInstanceResponder handles the response to the ListByInstance request. The method always
// closes the http.Response Body.
func (client ManagedInstanceAdministratorsClient) ListByInstanceResponder(resp *http.Response) (result ManagedInstanceAdministratorListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByInstanceNextResults retrieves the next set of results, if any.
func (client ManagedInstanceAdministratorsClient) listByInstanceNextResults(ctx context.Context, lastResults ManagedInstanceAdministratorListResult) (result ManagedInstanceAdministratorListResult, err error) {
	req, err := lastResults.managedInstanceAdministratorListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "sql.ManagedInstanceAdministratorsClient", "listByInstanceNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByInstanceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "sql.ManagedInstanceAdministratorsClient", "listByInstanceNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByInstanceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedInstanceAdministratorsClient", "listByInstanceNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByInstanceComplete enumerates all values, automatically crossing page boundaries as required.
func (client ManagedInstanceAdministratorsClient) ListByInstanceComplete(ctx context.Context, resourceGroupName string, managedInstanceName string) (result ManagedInstanceAdministratorListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ManagedInstanceAdministratorsClient.ListByInstance")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByInstance(ctx, resourceGroupName, managedInstanceName)
	return
}
