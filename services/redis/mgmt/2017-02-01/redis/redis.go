package redis

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

// Client is the REST API for Azure Redis Cache Service.
type Client struct {
	BaseClient
}

// NewClient creates an instance of the Client client.
func NewClient(subscriptionID string) Client {
	return NewClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewClientWithBaseURI creates an instance of the Client client.
func NewClientWithBaseURI(baseURI string, subscriptionID string) Client {
	return Client{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create create or replace (overwrite/recreate, with potential downtime) an existing Redis cache.
// Parameters:
// resourceGroupName - the name of the resource group.
// name - the name of the Redis cache.
// parameters - parameters supplied to the Create Redis operation.
func (client Client) Create(ctx context.Context, resourceGroupName string, name string, parameters CreateParameters) (result CreateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.Create")
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
			Constraints: []validation.Constraint{{Target: "parameters.CreateProperties", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "parameters.CreateProperties.Sku", Name: validation.Null, Rule: true,
					Chain: []validation.Constraint{{Target: "parameters.CreateProperties.Sku.Capacity", Name: validation.Null, Rule: true, Chain: nil}}},
				}}}}}); err != nil {
		return result, validation.NewError("redis.Client", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, resourceGroupName, name, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redis.Client", "Create", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redis.Client", "Create", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client Client) CreatePreparer(ctx context.Context, resourceGroupName string, name string, parameters CreateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-02-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{name}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client Client) CreateSender(req *http.Request) (future CreateFuture, err error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req, sd...)
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client Client) CreateResponder(resp *http.Response) (result ResourceType, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a Redis cache.
// Parameters:
// resourceGroupName - the name of the resource group.
// name - the name of the Redis cache.
func (client Client) Delete(ctx context.Context, resourceGroupName string, name string) (result DeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, name)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redis.Client", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redis.Client", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client Client) DeletePreparer(ctx context.Context, resourceGroupName string, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-02-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{name}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client Client) DeleteSender(req *http.Request) (future DeleteFuture, err error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req, sd...)
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client Client) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// ExportData export data from the redis cache to blobs in a container.
// Parameters:
// resourceGroupName - the name of the resource group.
// name - the name of the Redis cache.
// parameters - parameters for Redis export operation.
func (client Client) ExportData(ctx context.Context, resourceGroupName string, name string, parameters ExportRDBParameters) (result ExportDataFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.ExportData")
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
			Constraints: []validation.Constraint{{Target: "parameters.Prefix", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.Container", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("redis.Client", "ExportData", err.Error())
	}

	req, err := client.ExportDataPreparer(ctx, resourceGroupName, name, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redis.Client", "ExportData", nil, "Failure preparing request")
		return
	}

	result, err = client.ExportDataSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redis.Client", "ExportData", result.Response(), "Failure sending request")
		return
	}

	return
}

// ExportDataPreparer prepares the ExportData request.
func (client Client) ExportDataPreparer(ctx context.Context, resourceGroupName string, name string, parameters ExportRDBParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-02-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{name}/export", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ExportDataSender sends the ExportData request. The method will close the
// http.Response Body if it receives an error.
func (client Client) ExportDataSender(req *http.Request) (future ExportDataFuture, err error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req, sd...)
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// ExportDataResponder handles the response to the ExportData request. The method always
// closes the http.Response Body.
func (client Client) ExportDataResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// ForceReboot reboot specified Redis node(s). This operation requires write permission to the cache resource. There
// can be potential data loss.
// Parameters:
// resourceGroupName - the name of the resource group.
// name - the name of the Redis cache.
// parameters - specifies which Redis node(s) to reboot.
func (client Client) ForceReboot(ctx context.Context, resourceGroupName string, name string, parameters RebootParameters) (result ForceRebootResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.ForceReboot")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ForceRebootPreparer(ctx, resourceGroupName, name, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redis.Client", "ForceReboot", nil, "Failure preparing request")
		return
	}

	resp, err := client.ForceRebootSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "redis.Client", "ForceReboot", resp, "Failure sending request")
		return
	}

	result, err = client.ForceRebootResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redis.Client", "ForceReboot", resp, "Failure responding to request")
	}

	return
}

// ForceRebootPreparer prepares the ForceReboot request.
func (client Client) ForceRebootPreparer(ctx context.Context, resourceGroupName string, name string, parameters RebootParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-02-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{name}/forceReboot", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ForceRebootSender sends the ForceReboot request. The method will close the
// http.Response Body if it receives an error.
func (client Client) ForceRebootSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ForceRebootResponder handles the response to the ForceReboot request. The method always
// closes the http.Response Body.
func (client Client) ForceRebootResponder(resp *http.Response) (result ForceRebootResponse, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get gets a Redis cache (resource description).
// Parameters:
// resourceGroupName - the name of the resource group.
// name - the name of the Redis cache.
func (client Client) Get(ctx context.Context, resourceGroupName string, name string) (result ResourceType, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, name)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redis.Client", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "redis.Client", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redis.Client", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client Client) GetPreparer(ctx context.Context, resourceGroupName string, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-02-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{name}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client Client) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client Client) GetResponder(resp *http.Response) (result ResourceType, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ImportData import data into Redis cache.
// Parameters:
// resourceGroupName - the name of the resource group.
// name - the name of the Redis cache.
// parameters - parameters for Redis import operation.
func (client Client) ImportData(ctx context.Context, resourceGroupName string, name string, parameters ImportRDBParameters) (result ImportDataFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.ImportData")
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
			Constraints: []validation.Constraint{{Target: "parameters.Files", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("redis.Client", "ImportData", err.Error())
	}

	req, err := client.ImportDataPreparer(ctx, resourceGroupName, name, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redis.Client", "ImportData", nil, "Failure preparing request")
		return
	}

	result, err = client.ImportDataSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redis.Client", "ImportData", result.Response(), "Failure sending request")
		return
	}

	return
}

// ImportDataPreparer prepares the ImportData request.
func (client Client) ImportDataPreparer(ctx context.Context, resourceGroupName string, name string, parameters ImportRDBParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-02-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{name}/import", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ImportDataSender sends the ImportData request. The method will close the
// http.Response Body if it receives an error.
func (client Client) ImportDataSender(req *http.Request) (future ImportDataFuture, err error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req, sd...)
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// ImportDataResponder handles the response to the ImportData request. The method always
// closes the http.Response Body.
func (client Client) ImportDataResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// List gets all Redis caches in the specified subscription.
func (client Client) List(ctx context.Context) (result ListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.List")
		defer func() {
			sc := -1
			if result.lr.Response.Response != nil {
				sc = result.lr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redis.Client", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.lr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "redis.Client", "List", resp, "Failure sending request")
		return
	}

	result.lr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redis.Client", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client Client) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-02-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Cache/Redis/", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client Client) ListSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client Client) ListResponder(resp *http.Response) (result ListResult, err error) {
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
func (client Client) listNextResults(ctx context.Context, lastResults ListResult) (result ListResult, err error) {
	req, err := lastResults.listResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "redis.Client", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "redis.Client", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redis.Client", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client Client) ListComplete(ctx context.Context) (result ListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.List")
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

// ListByResourceGroup lists all Redis caches in a resource group.
// Parameters:
// resourceGroupName - the name of the resource group.
func (client Client) ListByResourceGroup(ctx context.Context, resourceGroupName string) (result ListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.lr.Response.Response != nil {
				sc = result.lr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redis.Client", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.lr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "redis.Client", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.lr, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redis.Client", "ListByResourceGroup", resp, "Failure responding to request")
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client Client) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-02-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client Client) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client Client) ListByResourceGroupResponder(resp *http.Response) (result ListResult, err error) {
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
func (client Client) listByResourceGroupNextResults(ctx context.Context, lastResults ListResult) (result ListResult, err error) {
	req, err := lastResults.listResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "redis.Client", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "redis.Client", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redis.Client", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client Client) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result ListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.ListByResourceGroup")
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

// ListKeys retrieve a Redis cache's access keys. This operation requires write permission to the cache resource.
// Parameters:
// resourceGroupName - the name of the resource group.
// name - the name of the Redis cache.
func (client Client) ListKeys(ctx context.Context, resourceGroupName string, name string) (result AccessKeys, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.ListKeys")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListKeysPreparer(ctx, resourceGroupName, name)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redis.Client", "ListKeys", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListKeysSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "redis.Client", "ListKeys", resp, "Failure sending request")
		return
	}

	result, err = client.ListKeysResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redis.Client", "ListKeys", resp, "Failure responding to request")
	}

	return
}

// ListKeysPreparer prepares the ListKeys request.
func (client Client) ListKeysPreparer(ctx context.Context, resourceGroupName string, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-02-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{name}/listKeys", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListKeysSender sends the ListKeys request. The method will close the
// http.Response Body if it receives an error.
func (client Client) ListKeysSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListKeysResponder handles the response to the ListKeys request. The method always
// closes the http.Response Body.
func (client Client) ListKeysResponder(resp *http.Response) (result AccessKeys, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// RegenerateKey regenerate Redis cache's access keys. This operation requires write permission to the cache resource.
// Parameters:
// resourceGroupName - the name of the resource group.
// name - the name of the Redis cache.
// parameters - specifies which key to regenerate.
func (client Client) RegenerateKey(ctx context.Context, resourceGroupName string, name string, parameters RegenerateKeyParameters) (result AccessKeys, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.RegenerateKey")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.RegenerateKeyPreparer(ctx, resourceGroupName, name, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redis.Client", "RegenerateKey", nil, "Failure preparing request")
		return
	}

	resp, err := client.RegenerateKeySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "redis.Client", "RegenerateKey", resp, "Failure sending request")
		return
	}

	result, err = client.RegenerateKeyResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redis.Client", "RegenerateKey", resp, "Failure responding to request")
	}

	return
}

// RegenerateKeyPreparer prepares the RegenerateKey request.
func (client Client) RegenerateKeyPreparer(ctx context.Context, resourceGroupName string, name string, parameters RegenerateKeyParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-02-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{name}/regenerateKey", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RegenerateKeySender sends the RegenerateKey request. The method will close the
// http.Response Body if it receives an error.
func (client Client) RegenerateKeySender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// RegenerateKeyResponder handles the response to the RegenerateKey request. The method always
// closes the http.Response Body.
func (client Client) RegenerateKeyResponder(resp *http.Response) (result AccessKeys, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update update an existing Redis cache.
// Parameters:
// resourceGroupName - the name of the resource group.
// name - the name of the Redis cache.
// parameters - parameters supplied to the Update Redis operation.
func (client Client) Update(ctx context.Context, resourceGroupName string, name string, parameters UpdateParameters) (result ResourceType, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, resourceGroupName, name, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redis.Client", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "redis.Client", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "redis.Client", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client Client) UpdatePreparer(ctx context.Context, resourceGroupName string, name string, parameters UpdateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-02-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/Redis/{name}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client Client) UpdateSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client Client) UpdateResponder(resp *http.Response) (result ResourceType, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
