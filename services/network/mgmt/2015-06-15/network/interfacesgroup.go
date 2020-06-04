package network

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

// InterfacesClient is the network Client
type InterfacesClient struct {
	BaseClient
}

// NewInterfacesClient creates an instance of the InterfacesClient client.
func NewInterfacesClient(subscriptionID string) InterfacesClient {
	return NewInterfacesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewInterfacesClientWithBaseURI creates an instance of the InterfacesClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewInterfacesClientWithBaseURI(baseURI string, subscriptionID string) InterfacesClient {
	return InterfacesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a network interface.
// Parameters:
// resourceGroupName - the name of the resource group.
// networkInterfaceName - the name of the network interface.
// parameters - parameters supplied to the create or update network interface operation.
func (client InterfacesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, networkInterfaceName string, parameters Interface) (result InterfacesCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InterfacesClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, networkInterfaceName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.InterfacesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.InterfacesClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client InterfacesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, networkInterfaceName string, parameters Interface) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"networkInterfaceName": autorest.Encode("path", networkInterfaceName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-06-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client InterfacesClient) CreateOrUpdateSender(req *http.Request) (future InterfacesCreateOrUpdateFuture, err error) {
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
func (client InterfacesClient) CreateOrUpdateResponder(resp *http.Response) (result Interface, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the specified network interface.
// Parameters:
// resourceGroupName - the name of the resource group.
// networkInterfaceName - the name of the network interface.
func (client InterfacesClient) Delete(ctx context.Context, resourceGroupName string, networkInterfaceName string) (result InterfacesDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InterfacesClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, networkInterfaceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.InterfacesClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.InterfacesClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client InterfacesClient) DeletePreparer(ctx context.Context, resourceGroupName string, networkInterfaceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"networkInterfaceName": autorest.Encode("path", networkInterfaceName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-06-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client InterfacesClient) DeleteSender(req *http.Request) (future InterfacesDeleteFuture, err error) {
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
func (client InterfacesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets information about the specified network interface.
// Parameters:
// resourceGroupName - the name of the resource group.
// networkInterfaceName - the name of the network interface.
// expand - expands referenced resources.
func (client InterfacesClient) Get(ctx context.Context, resourceGroupName string, networkInterfaceName string, expand string) (result Interface, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InterfacesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, networkInterfaceName, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.InterfacesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.InterfacesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.InterfacesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client InterfacesClient) GetPreparer(ctx context.Context, resourceGroupName string, networkInterfaceName string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"networkInterfaceName": autorest.Encode("path", networkInterfaceName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-06-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client InterfacesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client InterfacesClient) GetResponder(resp *http.Response) (result Interface, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetVirtualMachineScaleSetNetworkInterface get the specified network interface in a virtual machine scale set.
// Parameters:
// resourceGroupName - the name of the resource group.
// virtualMachineScaleSetName - the name of the virtual machine scale set.
// virtualmachineIndex - the virtual machine index.
// networkInterfaceName - the name of the network interface.
// expand - expands referenced resources.
func (client InterfacesClient) GetVirtualMachineScaleSetNetworkInterface(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string, virtualmachineIndex string, networkInterfaceName string, expand string) (result Interface, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InterfacesClient.GetVirtualMachineScaleSetNetworkInterface")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetVirtualMachineScaleSetNetworkInterfacePreparer(ctx, resourceGroupName, virtualMachineScaleSetName, virtualmachineIndex, networkInterfaceName, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.InterfacesClient", "GetVirtualMachineScaleSetNetworkInterface", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetVirtualMachineScaleSetNetworkInterfaceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.InterfacesClient", "GetVirtualMachineScaleSetNetworkInterface", resp, "Failure sending request")
		return
	}

	result, err = client.GetVirtualMachineScaleSetNetworkInterfaceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.InterfacesClient", "GetVirtualMachineScaleSetNetworkInterface", resp, "Failure responding to request")
	}

	return
}

// GetVirtualMachineScaleSetNetworkInterfacePreparer prepares the GetVirtualMachineScaleSetNetworkInterface request.
func (client InterfacesClient) GetVirtualMachineScaleSetNetworkInterfacePreparer(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string, virtualmachineIndex string, networkInterfaceName string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"networkInterfaceName":       autorest.Encode("path", networkInterfaceName),
		"resourceGroupName":          autorest.Encode("path", resourceGroupName),
		"subscriptionId":             autorest.Encode("path", client.SubscriptionID),
		"virtualmachineIndex":        autorest.Encode("path", virtualmachineIndex),
		"virtualMachineScaleSetName": autorest.Encode("path", virtualMachineScaleSetName),
	}

	const APIVersion = "2015-06-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.Compute/virtualMachineScaleSets/{virtualMachineScaleSetName}/virtualMachines/{virtualmachineIndex}/networkInterfaces/{networkInterfaceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetVirtualMachineScaleSetNetworkInterfaceSender sends the GetVirtualMachineScaleSetNetworkInterface request. The method will close the
// http.Response Body if it receives an error.
func (client InterfacesClient) GetVirtualMachineScaleSetNetworkInterfaceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetVirtualMachineScaleSetNetworkInterfaceResponder handles the response to the GetVirtualMachineScaleSetNetworkInterface request. The method always
// closes the http.Response Body.
func (client InterfacesClient) GetVirtualMachineScaleSetNetworkInterfaceResponder(resp *http.Response) (result Interface, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets all network interfaces in a resource group.
// Parameters:
// resourceGroupName - the name of the resource group.
func (client InterfacesClient) List(ctx context.Context, resourceGroupName string) (result InterfaceListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InterfacesClient.List")
		defer func() {
			sc := -1
			if result.ilr.Response.Response != nil {
				sc = result.ilr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.InterfacesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.ilr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.InterfacesClient", "List", resp, "Failure sending request")
		return
	}

	result.ilr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.InterfacesClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client InterfacesClient) ListPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-06-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client InterfacesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client InterfacesClient) ListResponder(resp *http.Response) (result InterfaceListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client InterfacesClient) listNextResults(ctx context.Context, lastResults InterfaceListResult) (result InterfaceListResult, err error) {
	req, err := lastResults.interfaceListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.InterfacesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.InterfacesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.InterfacesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client InterfacesClient) ListComplete(ctx context.Context, resourceGroupName string) (result InterfaceListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InterfacesClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName)
	return
}

// ListAll gets all network interfaces in a subscription.
func (client InterfacesClient) ListAll(ctx context.Context) (result InterfaceListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InterfacesClient.ListAll")
		defer func() {
			sc := -1
			if result.ilr.Response.Response != nil {
				sc = result.ilr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listAllNextResults
	req, err := client.ListAllPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.InterfacesClient", "ListAll", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListAllSender(req)
	if err != nil {
		result.ilr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.InterfacesClient", "ListAll", resp, "Failure sending request")
		return
	}

	result.ilr, err = client.ListAllResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.InterfacesClient", "ListAll", resp, "Failure responding to request")
	}

	return
}

// ListAllPreparer prepares the ListAll request.
func (client InterfacesClient) ListAllPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-06-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Network/networkInterfaces", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListAllSender sends the ListAll request. The method will close the
// http.Response Body if it receives an error.
func (client InterfacesClient) ListAllSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListAllResponder handles the response to the ListAll request. The method always
// closes the http.Response Body.
func (client InterfacesClient) ListAllResponder(resp *http.Response) (result InterfaceListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listAllNextResults retrieves the next set of results, if any.
func (client InterfacesClient) listAllNextResults(ctx context.Context, lastResults InterfaceListResult) (result InterfaceListResult, err error) {
	req, err := lastResults.interfaceListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.InterfacesClient", "listAllNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListAllSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.InterfacesClient", "listAllNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListAllResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.InterfacesClient", "listAllNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListAllComplete enumerates all values, automatically crossing page boundaries as required.
func (client InterfacesClient) ListAllComplete(ctx context.Context) (result InterfaceListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InterfacesClient.ListAll")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListAll(ctx)
	return
}

// ListVirtualMachineScaleSetNetworkInterfaces gets all network interfaces in a virtual machine scale set.
// Parameters:
// resourceGroupName - the name of the resource group.
// virtualMachineScaleSetName - the name of the virtual machine scale set.
func (client InterfacesClient) ListVirtualMachineScaleSetNetworkInterfaces(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string) (result InterfaceListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InterfacesClient.ListVirtualMachineScaleSetNetworkInterfaces")
		defer func() {
			sc := -1
			if result.ilr.Response.Response != nil {
				sc = result.ilr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listVirtualMachineScaleSetNetworkInterfacesNextResults
	req, err := client.ListVirtualMachineScaleSetNetworkInterfacesPreparer(ctx, resourceGroupName, virtualMachineScaleSetName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.InterfacesClient", "ListVirtualMachineScaleSetNetworkInterfaces", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListVirtualMachineScaleSetNetworkInterfacesSender(req)
	if err != nil {
		result.ilr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.InterfacesClient", "ListVirtualMachineScaleSetNetworkInterfaces", resp, "Failure sending request")
		return
	}

	result.ilr, err = client.ListVirtualMachineScaleSetNetworkInterfacesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.InterfacesClient", "ListVirtualMachineScaleSetNetworkInterfaces", resp, "Failure responding to request")
	}

	return
}

// ListVirtualMachineScaleSetNetworkInterfacesPreparer prepares the ListVirtualMachineScaleSetNetworkInterfaces request.
func (client InterfacesClient) ListVirtualMachineScaleSetNetworkInterfacesPreparer(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":          autorest.Encode("path", resourceGroupName),
		"subscriptionId":             autorest.Encode("path", client.SubscriptionID),
		"virtualMachineScaleSetName": autorest.Encode("path", virtualMachineScaleSetName),
	}

	const APIVersion = "2015-06-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.Compute/virtualMachineScaleSets/{virtualMachineScaleSetName}/networkInterfaces", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListVirtualMachineScaleSetNetworkInterfacesSender sends the ListVirtualMachineScaleSetNetworkInterfaces request. The method will close the
// http.Response Body if it receives an error.
func (client InterfacesClient) ListVirtualMachineScaleSetNetworkInterfacesSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListVirtualMachineScaleSetNetworkInterfacesResponder handles the response to the ListVirtualMachineScaleSetNetworkInterfaces request. The method always
// closes the http.Response Body.
func (client InterfacesClient) ListVirtualMachineScaleSetNetworkInterfacesResponder(resp *http.Response) (result InterfaceListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listVirtualMachineScaleSetNetworkInterfacesNextResults retrieves the next set of results, if any.
func (client InterfacesClient) listVirtualMachineScaleSetNetworkInterfacesNextResults(ctx context.Context, lastResults InterfaceListResult) (result InterfaceListResult, err error) {
	req, err := lastResults.interfaceListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.InterfacesClient", "listVirtualMachineScaleSetNetworkInterfacesNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListVirtualMachineScaleSetNetworkInterfacesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.InterfacesClient", "listVirtualMachineScaleSetNetworkInterfacesNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListVirtualMachineScaleSetNetworkInterfacesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.InterfacesClient", "listVirtualMachineScaleSetNetworkInterfacesNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListVirtualMachineScaleSetNetworkInterfacesComplete enumerates all values, automatically crossing page boundaries as required.
func (client InterfacesClient) ListVirtualMachineScaleSetNetworkInterfacesComplete(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string) (result InterfaceListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InterfacesClient.ListVirtualMachineScaleSetNetworkInterfaces")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListVirtualMachineScaleSetNetworkInterfaces(ctx, resourceGroupName, virtualMachineScaleSetName)
	return
}

// ListVirtualMachineScaleSetVMNetworkInterfaces gets information about all network interfaces in a virtual machine in
// a virtual machine scale set.
// Parameters:
// resourceGroupName - the name of the resource group.
// virtualMachineScaleSetName - the name of the virtual machine scale set.
// virtualmachineIndex - the virtual machine index.
func (client InterfacesClient) ListVirtualMachineScaleSetVMNetworkInterfaces(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string, virtualmachineIndex string) (result InterfaceListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InterfacesClient.ListVirtualMachineScaleSetVMNetworkInterfaces")
		defer func() {
			sc := -1
			if result.ilr.Response.Response != nil {
				sc = result.ilr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listVirtualMachineScaleSetVMNetworkInterfacesNextResults
	req, err := client.ListVirtualMachineScaleSetVMNetworkInterfacesPreparer(ctx, resourceGroupName, virtualMachineScaleSetName, virtualmachineIndex)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.InterfacesClient", "ListVirtualMachineScaleSetVMNetworkInterfaces", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListVirtualMachineScaleSetVMNetworkInterfacesSender(req)
	if err != nil {
		result.ilr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.InterfacesClient", "ListVirtualMachineScaleSetVMNetworkInterfaces", resp, "Failure sending request")
		return
	}

	result.ilr, err = client.ListVirtualMachineScaleSetVMNetworkInterfacesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.InterfacesClient", "ListVirtualMachineScaleSetVMNetworkInterfaces", resp, "Failure responding to request")
	}

	return
}

// ListVirtualMachineScaleSetVMNetworkInterfacesPreparer prepares the ListVirtualMachineScaleSetVMNetworkInterfaces request.
func (client InterfacesClient) ListVirtualMachineScaleSetVMNetworkInterfacesPreparer(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string, virtualmachineIndex string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":          autorest.Encode("path", resourceGroupName),
		"subscriptionId":             autorest.Encode("path", client.SubscriptionID),
		"virtualmachineIndex":        autorest.Encode("path", virtualmachineIndex),
		"virtualMachineScaleSetName": autorest.Encode("path", virtualMachineScaleSetName),
	}

	const APIVersion = "2015-06-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.Compute/virtualMachineScaleSets/{virtualMachineScaleSetName}/virtualMachines/{virtualmachineIndex}/networkInterfaces", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListVirtualMachineScaleSetVMNetworkInterfacesSender sends the ListVirtualMachineScaleSetVMNetworkInterfaces request. The method will close the
// http.Response Body if it receives an error.
func (client InterfacesClient) ListVirtualMachineScaleSetVMNetworkInterfacesSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListVirtualMachineScaleSetVMNetworkInterfacesResponder handles the response to the ListVirtualMachineScaleSetVMNetworkInterfaces request. The method always
// closes the http.Response Body.
func (client InterfacesClient) ListVirtualMachineScaleSetVMNetworkInterfacesResponder(resp *http.Response) (result InterfaceListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listVirtualMachineScaleSetVMNetworkInterfacesNextResults retrieves the next set of results, if any.
func (client InterfacesClient) listVirtualMachineScaleSetVMNetworkInterfacesNextResults(ctx context.Context, lastResults InterfaceListResult) (result InterfaceListResult, err error) {
	req, err := lastResults.interfaceListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.InterfacesClient", "listVirtualMachineScaleSetVMNetworkInterfacesNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListVirtualMachineScaleSetVMNetworkInterfacesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.InterfacesClient", "listVirtualMachineScaleSetVMNetworkInterfacesNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListVirtualMachineScaleSetVMNetworkInterfacesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.InterfacesClient", "listVirtualMachineScaleSetVMNetworkInterfacesNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListVirtualMachineScaleSetVMNetworkInterfacesComplete enumerates all values, automatically crossing page boundaries as required.
func (client InterfacesClient) ListVirtualMachineScaleSetVMNetworkInterfacesComplete(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string, virtualmachineIndex string) (result InterfaceListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InterfacesClient.ListVirtualMachineScaleSetVMNetworkInterfaces")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListVirtualMachineScaleSetVMNetworkInterfaces(ctx, resourceGroupName, virtualMachineScaleSetName, virtualmachineIndex)
	return
}
