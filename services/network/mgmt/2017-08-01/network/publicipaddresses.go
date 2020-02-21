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
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// PublicIPAddressesClient is the network Client
type PublicIPAddressesClient struct {
	BaseClient
}

// NewPublicIPAddressesClient creates an instance of the PublicIPAddressesClient client.
func NewPublicIPAddressesClient(subscriptionID string) PublicIPAddressesClient {
	return NewPublicIPAddressesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewPublicIPAddressesClientWithBaseURI creates an instance of the PublicIPAddressesClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewPublicIPAddressesClientWithBaseURI(baseURI string, subscriptionID string) PublicIPAddressesClient {
	return PublicIPAddressesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a static or dynamic public IP address.
// Parameters:
// resourceGroupName - the name of the resource group.
// publicIPAddressName - the name of the public IP address.
// parameters - parameters supplied to the create or update public IP address operation.
func (client PublicIPAddressesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, publicIPAddressName string, parameters PublicIPAddress) (result PublicIPAddressesCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PublicIPAddressesClient.CreateOrUpdate")
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
			Constraints: []validation.Constraint{{Target: "parameters.PublicIPAddressPropertiesFormat", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.PublicIPAddressPropertiesFormat.IPConfiguration", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "parameters.PublicIPAddressPropertiesFormat.IPConfiguration.IPConfigurationPropertiesFormat", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.PublicIPAddressPropertiesFormat.IPConfiguration.IPConfigurationPropertiesFormat.PublicIPAddress", Name: validation.Null, Rule: false, Chain: nil}}},
					}},
				}}}}}); err != nil {
		return result, validation.NewError("network.PublicIPAddressesClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, publicIPAddressName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PublicIPAddressesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PublicIPAddressesClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client PublicIPAddressesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, publicIPAddressName string, parameters PublicIPAddress) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"publicIpAddressName": autorest.Encode("path", publicIPAddressName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPAddresses/{publicIpAddressName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client PublicIPAddressesClient) CreateOrUpdateSender(req *http.Request) (future PublicIPAddressesCreateOrUpdateFuture, err error) {
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
func (client PublicIPAddressesClient) CreateOrUpdateResponder(resp *http.Response) (result PublicIPAddress, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the specified public IP address.
// Parameters:
// resourceGroupName - the name of the resource group.
// publicIPAddressName - the name of the subnet.
func (client PublicIPAddressesClient) Delete(ctx context.Context, resourceGroupName string, publicIPAddressName string) (result PublicIPAddressesDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PublicIPAddressesClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, publicIPAddressName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PublicIPAddressesClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PublicIPAddressesClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client PublicIPAddressesClient) DeletePreparer(ctx context.Context, resourceGroupName string, publicIPAddressName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"publicIpAddressName": autorest.Encode("path", publicIPAddressName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPAddresses/{publicIpAddressName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client PublicIPAddressesClient) DeleteSender(req *http.Request) (future PublicIPAddressesDeleteFuture, err error) {
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
func (client PublicIPAddressesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the specified public IP address in a specified resource group.
// Parameters:
// resourceGroupName - the name of the resource group.
// publicIPAddressName - the name of the subnet.
// expand - expands referenced resources.
func (client PublicIPAddressesClient) Get(ctx context.Context, resourceGroupName string, publicIPAddressName string, expand string) (result PublicIPAddress, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PublicIPAddressesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, publicIPAddressName, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PublicIPAddressesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.PublicIPAddressesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PublicIPAddressesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client PublicIPAddressesClient) GetPreparer(ctx context.Context, resourceGroupName string, publicIPAddressName string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"publicIpAddressName": autorest.Encode("path", publicIPAddressName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPAddresses/{publicIpAddressName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client PublicIPAddressesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client PublicIPAddressesClient) GetResponder(resp *http.Response) (result PublicIPAddress, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetVirtualMachineScaleSetPublicIPAddress get the specified public IP address in a virtual machine scale set.
// Parameters:
// resourceGroupName - the name of the resource group.
// virtualMachineScaleSetName - the name of the virtual machine scale set.
// virtualmachineIndex - the virtual machine index.
// networkInterfaceName - the name of the network interface.
// IPConfigurationName - the name of the IP configuration.
// publicIPAddressName - the name of the public IP Address.
// expand - expands referenced resources.
func (client PublicIPAddressesClient) GetVirtualMachineScaleSetPublicIPAddress(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string, virtualmachineIndex string, networkInterfaceName string, IPConfigurationName string, publicIPAddressName string, expand string) (result PublicIPAddress, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PublicIPAddressesClient.GetVirtualMachineScaleSetPublicIPAddress")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetVirtualMachineScaleSetPublicIPAddressPreparer(ctx, resourceGroupName, virtualMachineScaleSetName, virtualmachineIndex, networkInterfaceName, IPConfigurationName, publicIPAddressName, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PublicIPAddressesClient", "GetVirtualMachineScaleSetPublicIPAddress", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetVirtualMachineScaleSetPublicIPAddressSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.PublicIPAddressesClient", "GetVirtualMachineScaleSetPublicIPAddress", resp, "Failure sending request")
		return
	}

	result, err = client.GetVirtualMachineScaleSetPublicIPAddressResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PublicIPAddressesClient", "GetVirtualMachineScaleSetPublicIPAddress", resp, "Failure responding to request")
	}

	return
}

// GetVirtualMachineScaleSetPublicIPAddressPreparer prepares the GetVirtualMachineScaleSetPublicIPAddress request.
func (client PublicIPAddressesClient) GetVirtualMachineScaleSetPublicIPAddressPreparer(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string, virtualmachineIndex string, networkInterfaceName string, IPConfigurationName string, publicIPAddressName string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"ipConfigurationName":        autorest.Encode("path", IPConfigurationName),
		"networkInterfaceName":       autorest.Encode("path", networkInterfaceName),
		"publicIpAddressName":        autorest.Encode("path", publicIPAddressName),
		"resourceGroupName":          autorest.Encode("path", resourceGroupName),
		"subscriptionId":             autorest.Encode("path", client.SubscriptionID),
		"virtualmachineIndex":        autorest.Encode("path", virtualmachineIndex),
		"virtualMachineScaleSetName": autorest.Encode("path", virtualMachineScaleSetName),
	}

	const APIVersion = "2017-03-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{virtualMachineScaleSetName}/virtualMachines/{virtualmachineIndex}/networkInterfaces/{networkInterfaceName}/ipconfigurations/{ipConfigurationName}/publicipaddresses/{publicIpAddressName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetVirtualMachineScaleSetPublicIPAddressSender sends the GetVirtualMachineScaleSetPublicIPAddress request. The method will close the
// http.Response Body if it receives an error.
func (client PublicIPAddressesClient) GetVirtualMachineScaleSetPublicIPAddressSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetVirtualMachineScaleSetPublicIPAddressResponder handles the response to the GetVirtualMachineScaleSetPublicIPAddress request. The method always
// closes the http.Response Body.
func (client PublicIPAddressesClient) GetVirtualMachineScaleSetPublicIPAddressResponder(resp *http.Response) (result PublicIPAddress, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets all public IP addresses in a resource group.
// Parameters:
// resourceGroupName - the name of the resource group.
func (client PublicIPAddressesClient) List(ctx context.Context, resourceGroupName string) (result PublicIPAddressListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PublicIPAddressesClient.List")
		defer func() {
			sc := -1
			if result.pialr.Response.Response != nil {
				sc = result.pialr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PublicIPAddressesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.pialr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.PublicIPAddressesClient", "List", resp, "Failure sending request")
		return
	}

	result.pialr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PublicIPAddressesClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client PublicIPAddressesClient) ListPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPAddresses", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client PublicIPAddressesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client PublicIPAddressesClient) ListResponder(resp *http.Response) (result PublicIPAddressListResult, err error) {
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
func (client PublicIPAddressesClient) listNextResults(ctx context.Context, lastResults PublicIPAddressListResult) (result PublicIPAddressListResult, err error) {
	req, err := lastResults.publicIPAddressListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.PublicIPAddressesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.PublicIPAddressesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PublicIPAddressesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client PublicIPAddressesClient) ListComplete(ctx context.Context, resourceGroupName string) (result PublicIPAddressListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PublicIPAddressesClient.List")
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

// ListAll gets all the public IP addresses in a subscription.
func (client PublicIPAddressesClient) ListAll(ctx context.Context) (result PublicIPAddressListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PublicIPAddressesClient.ListAll")
		defer func() {
			sc := -1
			if result.pialr.Response.Response != nil {
				sc = result.pialr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listAllNextResults
	req, err := client.ListAllPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PublicIPAddressesClient", "ListAll", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListAllSender(req)
	if err != nil {
		result.pialr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.PublicIPAddressesClient", "ListAll", resp, "Failure sending request")
		return
	}

	result.pialr, err = client.ListAllResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PublicIPAddressesClient", "ListAll", resp, "Failure responding to request")
	}

	return
}

// ListAllPreparer prepares the ListAll request.
func (client PublicIPAddressesClient) ListAllPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Network/publicIPAddresses", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListAllSender sends the ListAll request. The method will close the
// http.Response Body if it receives an error.
func (client PublicIPAddressesClient) ListAllSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListAllResponder handles the response to the ListAll request. The method always
// closes the http.Response Body.
func (client PublicIPAddressesClient) ListAllResponder(resp *http.Response) (result PublicIPAddressListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listAllNextResults retrieves the next set of results, if any.
func (client PublicIPAddressesClient) listAllNextResults(ctx context.Context, lastResults PublicIPAddressListResult) (result PublicIPAddressListResult, err error) {
	req, err := lastResults.publicIPAddressListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.PublicIPAddressesClient", "listAllNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListAllSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.PublicIPAddressesClient", "listAllNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListAllResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PublicIPAddressesClient", "listAllNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListAllComplete enumerates all values, automatically crossing page boundaries as required.
func (client PublicIPAddressesClient) ListAllComplete(ctx context.Context) (result PublicIPAddressListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PublicIPAddressesClient.ListAll")
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

// ListVirtualMachineScaleSetPublicIPAddresses gets information about all public IP addresses on a virtual machine
// scale set level.
// Parameters:
// resourceGroupName - the name of the resource group.
// virtualMachineScaleSetName - the name of the virtual machine scale set.
func (client PublicIPAddressesClient) ListVirtualMachineScaleSetPublicIPAddresses(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string) (result PublicIPAddressListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PublicIPAddressesClient.ListVirtualMachineScaleSetPublicIPAddresses")
		defer func() {
			sc := -1
			if result.pialr.Response.Response != nil {
				sc = result.pialr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listVirtualMachineScaleSetPublicIPAddressesNextResults
	req, err := client.ListVirtualMachineScaleSetPublicIPAddressesPreparer(ctx, resourceGroupName, virtualMachineScaleSetName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PublicIPAddressesClient", "ListVirtualMachineScaleSetPublicIPAddresses", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListVirtualMachineScaleSetPublicIPAddressesSender(req)
	if err != nil {
		result.pialr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.PublicIPAddressesClient", "ListVirtualMachineScaleSetPublicIPAddresses", resp, "Failure sending request")
		return
	}

	result.pialr, err = client.ListVirtualMachineScaleSetPublicIPAddressesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PublicIPAddressesClient", "ListVirtualMachineScaleSetPublicIPAddresses", resp, "Failure responding to request")
	}

	return
}

// ListVirtualMachineScaleSetPublicIPAddressesPreparer prepares the ListVirtualMachineScaleSetPublicIPAddresses request.
func (client PublicIPAddressesClient) ListVirtualMachineScaleSetPublicIPAddressesPreparer(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":          autorest.Encode("path", resourceGroupName),
		"subscriptionId":             autorest.Encode("path", client.SubscriptionID),
		"virtualMachineScaleSetName": autorest.Encode("path", virtualMachineScaleSetName),
	}

	const APIVersion = "2017-03-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{virtualMachineScaleSetName}/publicipaddresses", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListVirtualMachineScaleSetPublicIPAddressesSender sends the ListVirtualMachineScaleSetPublicIPAddresses request. The method will close the
// http.Response Body if it receives an error.
func (client PublicIPAddressesClient) ListVirtualMachineScaleSetPublicIPAddressesSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListVirtualMachineScaleSetPublicIPAddressesResponder handles the response to the ListVirtualMachineScaleSetPublicIPAddresses request. The method always
// closes the http.Response Body.
func (client PublicIPAddressesClient) ListVirtualMachineScaleSetPublicIPAddressesResponder(resp *http.Response) (result PublicIPAddressListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listVirtualMachineScaleSetPublicIPAddressesNextResults retrieves the next set of results, if any.
func (client PublicIPAddressesClient) listVirtualMachineScaleSetPublicIPAddressesNextResults(ctx context.Context, lastResults PublicIPAddressListResult) (result PublicIPAddressListResult, err error) {
	req, err := lastResults.publicIPAddressListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.PublicIPAddressesClient", "listVirtualMachineScaleSetPublicIPAddressesNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListVirtualMachineScaleSetPublicIPAddressesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.PublicIPAddressesClient", "listVirtualMachineScaleSetPublicIPAddressesNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListVirtualMachineScaleSetPublicIPAddressesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PublicIPAddressesClient", "listVirtualMachineScaleSetPublicIPAddressesNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListVirtualMachineScaleSetPublicIPAddressesComplete enumerates all values, automatically crossing page boundaries as required.
func (client PublicIPAddressesClient) ListVirtualMachineScaleSetPublicIPAddressesComplete(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string) (result PublicIPAddressListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PublicIPAddressesClient.ListVirtualMachineScaleSetPublicIPAddresses")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListVirtualMachineScaleSetPublicIPAddresses(ctx, resourceGroupName, virtualMachineScaleSetName)
	return
}

// ListVirtualMachineScaleSetVMPublicIPAddresses gets information about all public IP addresses in a virtual machine IP
// configuration in a virtual machine scale set.
// Parameters:
// resourceGroupName - the name of the resource group.
// virtualMachineScaleSetName - the name of the virtual machine scale set.
// virtualmachineIndex - the virtual machine index.
// networkInterfaceName - the network interface name.
// IPConfigurationName - the IP configuration name.
func (client PublicIPAddressesClient) ListVirtualMachineScaleSetVMPublicIPAddresses(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string, virtualmachineIndex string, networkInterfaceName string, IPConfigurationName string) (result PublicIPAddressListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PublicIPAddressesClient.ListVirtualMachineScaleSetVMPublicIPAddresses")
		defer func() {
			sc := -1
			if result.pialr.Response.Response != nil {
				sc = result.pialr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listVirtualMachineScaleSetVMPublicIPAddressesNextResults
	req, err := client.ListVirtualMachineScaleSetVMPublicIPAddressesPreparer(ctx, resourceGroupName, virtualMachineScaleSetName, virtualmachineIndex, networkInterfaceName, IPConfigurationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PublicIPAddressesClient", "ListVirtualMachineScaleSetVMPublicIPAddresses", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListVirtualMachineScaleSetVMPublicIPAddressesSender(req)
	if err != nil {
		result.pialr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.PublicIPAddressesClient", "ListVirtualMachineScaleSetVMPublicIPAddresses", resp, "Failure sending request")
		return
	}

	result.pialr, err = client.ListVirtualMachineScaleSetVMPublicIPAddressesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PublicIPAddressesClient", "ListVirtualMachineScaleSetVMPublicIPAddresses", resp, "Failure responding to request")
	}

	return
}

// ListVirtualMachineScaleSetVMPublicIPAddressesPreparer prepares the ListVirtualMachineScaleSetVMPublicIPAddresses request.
func (client PublicIPAddressesClient) ListVirtualMachineScaleSetVMPublicIPAddressesPreparer(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string, virtualmachineIndex string, networkInterfaceName string, IPConfigurationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"ipConfigurationName":        autorest.Encode("path", IPConfigurationName),
		"networkInterfaceName":       autorest.Encode("path", networkInterfaceName),
		"resourceGroupName":          autorest.Encode("path", resourceGroupName),
		"subscriptionId":             autorest.Encode("path", client.SubscriptionID),
		"virtualmachineIndex":        autorest.Encode("path", virtualmachineIndex),
		"virtualMachineScaleSetName": autorest.Encode("path", virtualMachineScaleSetName),
	}

	const APIVersion = "2017-03-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{virtualMachineScaleSetName}/virtualMachines/{virtualmachineIndex}/networkInterfaces/{networkInterfaceName}/ipconfigurations/{ipConfigurationName}/publicipaddresses", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListVirtualMachineScaleSetVMPublicIPAddressesSender sends the ListVirtualMachineScaleSetVMPublicIPAddresses request. The method will close the
// http.Response Body if it receives an error.
func (client PublicIPAddressesClient) ListVirtualMachineScaleSetVMPublicIPAddressesSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListVirtualMachineScaleSetVMPublicIPAddressesResponder handles the response to the ListVirtualMachineScaleSetVMPublicIPAddresses request. The method always
// closes the http.Response Body.
func (client PublicIPAddressesClient) ListVirtualMachineScaleSetVMPublicIPAddressesResponder(resp *http.Response) (result PublicIPAddressListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listVirtualMachineScaleSetVMPublicIPAddressesNextResults retrieves the next set of results, if any.
func (client PublicIPAddressesClient) listVirtualMachineScaleSetVMPublicIPAddressesNextResults(ctx context.Context, lastResults PublicIPAddressListResult) (result PublicIPAddressListResult, err error) {
	req, err := lastResults.publicIPAddressListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.PublicIPAddressesClient", "listVirtualMachineScaleSetVMPublicIPAddressesNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListVirtualMachineScaleSetVMPublicIPAddressesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.PublicIPAddressesClient", "listVirtualMachineScaleSetVMPublicIPAddressesNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListVirtualMachineScaleSetVMPublicIPAddressesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PublicIPAddressesClient", "listVirtualMachineScaleSetVMPublicIPAddressesNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListVirtualMachineScaleSetVMPublicIPAddressesComplete enumerates all values, automatically crossing page boundaries as required.
func (client PublicIPAddressesClient) ListVirtualMachineScaleSetVMPublicIPAddressesComplete(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string, virtualmachineIndex string, networkInterfaceName string, IPConfigurationName string) (result PublicIPAddressListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PublicIPAddressesClient.ListVirtualMachineScaleSetVMPublicIPAddresses")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListVirtualMachineScaleSetVMPublicIPAddresses(ctx, resourceGroupName, virtualMachineScaleSetName, virtualmachineIndex, networkInterfaceName, IPConfigurationName)
	return
}
