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

// ManagedInstancesClient is the the Azure SQL Database management API provides a RESTful set of web services that
// interact with Azure SQL Database services to manage your databases. The API enables you to create, retrieve, update,
// and delete databases.
type ManagedInstancesClient struct {
	BaseClient
}

// NewManagedInstancesClient creates an instance of the ManagedInstancesClient client.
func NewManagedInstancesClient(subscriptionID string) ManagedInstancesClient {
	return NewManagedInstancesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewManagedInstancesClientWithBaseURI creates an instance of the ManagedInstancesClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewManagedInstancesClientWithBaseURI(baseURI string, subscriptionID string) ManagedInstancesClient {
	return ManagedInstancesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a managed instance.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// managedInstanceName - the name of the managed instance.
// parameters - the requested managed instance resource state.
func (client ManagedInstancesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, managedInstanceName string, parameters ManagedInstance) (result ManagedInstancesCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ManagedInstancesClient.CreateOrUpdate")
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
			Constraints: []validation.Constraint{{Target: "parameters.Sku", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.Sku.Name", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("sql.ManagedInstancesClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, managedInstanceName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedInstancesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedInstancesClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ManagedInstancesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, managedInstanceName string, parameters ManagedInstance) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"managedInstanceName": autorest.Encode("path", managedInstanceName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ManagedInstancesClient) CreateOrUpdateSender(req *http.Request) (future ManagedInstancesCreateOrUpdateFuture, err error) {
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
func (client ManagedInstancesClient) CreateOrUpdateResponder(resp *http.Response) (result ManagedInstance, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a managed instance.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// managedInstanceName - the name of the managed instance.
func (client ManagedInstancesClient) Delete(ctx context.Context, resourceGroupName string, managedInstanceName string) (result ManagedInstancesDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ManagedInstancesClient.Delete")
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
		err = autorest.NewErrorWithError(err, "sql.ManagedInstancesClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedInstancesClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ManagedInstancesClient) DeletePreparer(ctx context.Context, resourceGroupName string, managedInstanceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"managedInstanceName": autorest.Encode("path", managedInstanceName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ManagedInstancesClient) DeleteSender(req *http.Request) (future ManagedInstancesDeleteFuture, err error) {
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
func (client ManagedInstancesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets a managed instance.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// managedInstanceName - the name of the managed instance.
func (client ManagedInstancesClient) Get(ctx context.Context, resourceGroupName string, managedInstanceName string) (result ManagedInstance, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ManagedInstancesClient.Get")
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
		err = autorest.NewErrorWithError(err, "sql.ManagedInstancesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.ManagedInstancesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedInstancesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ManagedInstancesClient) GetPreparer(ctx context.Context, resourceGroupName string, managedInstanceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"managedInstanceName": autorest.Encode("path", managedInstanceName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ManagedInstancesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ManagedInstancesClient) GetResponder(resp *http.Response) (result ManagedInstance, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets a list of all managed instances in the subscription.
func (client ManagedInstancesClient) List(ctx context.Context) (result ManagedInstanceListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ManagedInstancesClient.List")
		defer func() {
			sc := -1
			if result.milr.Response.Response != nil {
				sc = result.milr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedInstancesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.milr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.ManagedInstancesClient", "List", resp, "Failure sending request")
		return
	}

	result.milr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedInstancesClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ManagedInstancesClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Sql/managedInstances", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ManagedInstancesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ManagedInstancesClient) ListResponder(resp *http.Response) (result ManagedInstanceListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client ManagedInstancesClient) listNextResults(ctx context.Context, lastResults ManagedInstanceListResult) (result ManagedInstanceListResult, err error) {
	req, err := lastResults.managedInstanceListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "sql.ManagedInstancesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "sql.ManagedInstancesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedInstancesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client ManagedInstancesClient) ListComplete(ctx context.Context) (result ManagedInstanceListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ManagedInstancesClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx)
	return
}

// ListByInstancePool gets a list of all managed instances in an instance pool.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// instancePoolName - the instance pool name.
func (client ManagedInstancesClient) ListByInstancePool(ctx context.Context, resourceGroupName string, instancePoolName string) (result ManagedInstanceListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ManagedInstancesClient.ListByInstancePool")
		defer func() {
			sc := -1
			if result.milr.Response.Response != nil {
				sc = result.milr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByInstancePoolNextResults
	req, err := client.ListByInstancePoolPreparer(ctx, resourceGroupName, instancePoolName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedInstancesClient", "ListByInstancePool", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByInstancePoolSender(req)
	if err != nil {
		result.milr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.ManagedInstancesClient", "ListByInstancePool", resp, "Failure sending request")
		return
	}

	result.milr, err = client.ListByInstancePoolResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedInstancesClient", "ListByInstancePool", resp, "Failure responding to request")
	}

	return
}

// ListByInstancePoolPreparer prepares the ListByInstancePool request.
func (client ManagedInstancesClient) ListByInstancePoolPreparer(ctx context.Context, resourceGroupName string, instancePoolName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"instancePoolName":  autorest.Encode("path", instancePoolName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/instancePools/{instancePoolName}/managedInstances", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByInstancePoolSender sends the ListByInstancePool request. The method will close the
// http.Response Body if it receives an error.
func (client ManagedInstancesClient) ListByInstancePoolSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByInstancePoolResponder handles the response to the ListByInstancePool request. The method always
// closes the http.Response Body.
func (client ManagedInstancesClient) ListByInstancePoolResponder(resp *http.Response) (result ManagedInstanceListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByInstancePoolNextResults retrieves the next set of results, if any.
func (client ManagedInstancesClient) listByInstancePoolNextResults(ctx context.Context, lastResults ManagedInstanceListResult) (result ManagedInstanceListResult, err error) {
	req, err := lastResults.managedInstanceListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "sql.ManagedInstancesClient", "listByInstancePoolNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByInstancePoolSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "sql.ManagedInstancesClient", "listByInstancePoolNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByInstancePoolResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedInstancesClient", "listByInstancePoolNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByInstancePoolComplete enumerates all values, automatically crossing page boundaries as required.
func (client ManagedInstancesClient) ListByInstancePoolComplete(ctx context.Context, resourceGroupName string, instancePoolName string) (result ManagedInstanceListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ManagedInstancesClient.ListByInstancePool")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByInstancePool(ctx, resourceGroupName, instancePoolName)
	return
}

// ListByResourceGroup gets a list of managed instances in a resource group.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
func (client ManagedInstancesClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (result ManagedInstanceListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ManagedInstancesClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.milr.Response.Response != nil {
				sc = result.milr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedInstancesClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.milr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.ManagedInstancesClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.milr, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedInstancesClient", "ListByResourceGroup", resp, "Failure responding to request")
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client ManagedInstancesClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client ManagedInstancesClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client ManagedInstancesClient) ListByResourceGroupResponder(resp *http.Response) (result ManagedInstanceListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client ManagedInstancesClient) listByResourceGroupNextResults(ctx context.Context, lastResults ManagedInstanceListResult) (result ManagedInstanceListResult, err error) {
	req, err := lastResults.managedInstanceListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "sql.ManagedInstancesClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "sql.ManagedInstancesClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedInstancesClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client ManagedInstancesClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result ManagedInstanceListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ManagedInstancesClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName)
	return
}

// Update updates a managed instance.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// managedInstanceName - the name of the managed instance.
// parameters - the requested managed instance resource state.
func (client ManagedInstancesClient) Update(ctx context.Context, resourceGroupName string, managedInstanceName string, parameters ManagedInstanceUpdate) (result ManagedInstancesUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ManagedInstancesClient.Update")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, resourceGroupName, managedInstanceName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedInstancesClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedInstancesClient", "Update", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client ManagedInstancesClient) UpdatePreparer(ctx context.Context, resourceGroupName string, managedInstanceName string, parameters ManagedInstanceUpdate) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"managedInstanceName": autorest.Encode("path", managedInstanceName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client ManagedInstancesClient) UpdateSender(req *http.Request) (future ManagedInstancesUpdateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client ManagedInstancesClient) UpdateResponder(resp *http.Response) (result ManagedInstance, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
