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

// P2sVpnGatewaysClient is the network Client
type P2sVpnGatewaysClient struct {
	BaseClient
}

// NewP2sVpnGatewaysClient creates an instance of the P2sVpnGatewaysClient client.
func NewP2sVpnGatewaysClient(subscriptionID string) P2sVpnGatewaysClient {
	return NewP2sVpnGatewaysClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewP2sVpnGatewaysClientWithBaseURI creates an instance of the P2sVpnGatewaysClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewP2sVpnGatewaysClientWithBaseURI(baseURI string, subscriptionID string) P2sVpnGatewaysClient {
	return P2sVpnGatewaysClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates a virtual wan p2s vpn gateway if it doesn't exist else updates the existing gateway.
// Parameters:
// resourceGroupName - the resource group name of the P2SVpnGateway.
// gatewayName - the name of the gateway.
// p2SVpnGatewayParameters - parameters supplied to create or Update a virtual wan p2s vpn gateway.
func (client P2sVpnGatewaysClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, gatewayName string, p2SVpnGatewayParameters P2SVpnGateway) (result P2sVpnGatewaysCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/P2sVpnGatewaysClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, gatewayName, p2SVpnGatewayParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.P2sVpnGatewaysClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.P2sVpnGatewaysClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client P2sVpnGatewaysClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, gatewayName string, p2SVpnGatewayParameters P2SVpnGateway) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"gatewayName":       autorest.Encode("path", gatewayName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	p2SVpnGatewayParameters.Etag = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/p2svpnGateways/{gatewayName}", pathParameters),
		autorest.WithJSON(p2SVpnGatewayParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client P2sVpnGatewaysClient) CreateOrUpdateSender(req *http.Request) (future P2sVpnGatewaysCreateOrUpdateFuture, err error) {
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
func (client P2sVpnGatewaysClient) CreateOrUpdateResponder(resp *http.Response) (result P2SVpnGateway, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a virtual wan p2s vpn gateway.
// Parameters:
// resourceGroupName - the resource group name of the P2SVpnGateway.
// gatewayName - the name of the gateway.
func (client P2sVpnGatewaysClient) Delete(ctx context.Context, resourceGroupName string, gatewayName string) (result P2sVpnGatewaysDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/P2sVpnGatewaysClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, gatewayName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.P2sVpnGatewaysClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.P2sVpnGatewaysClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client P2sVpnGatewaysClient) DeletePreparer(ctx context.Context, resourceGroupName string, gatewayName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"gatewayName":       autorest.Encode("path", gatewayName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/p2svpnGateways/{gatewayName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client P2sVpnGatewaysClient) DeleteSender(req *http.Request) (future P2sVpnGatewaysDeleteFuture, err error) {
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
func (client P2sVpnGatewaysClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// GenerateVpnProfile generates VPN profile for P2S client of the P2SVpnGateway in the specified resource group.
// Parameters:
// resourceGroupName - the name of the resource group.
// gatewayName - the name of the P2SVpnGateway.
// parameters - parameters supplied to the generate P2SVpnGateway VPN client package operation.
func (client P2sVpnGatewaysClient) GenerateVpnProfile(ctx context.Context, resourceGroupName string, gatewayName string, parameters P2SVpnProfileParameters) (result P2sVpnGatewaysGenerateVpnProfileFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/P2sVpnGatewaysClient.GenerateVpnProfile")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GenerateVpnProfilePreparer(ctx, resourceGroupName, gatewayName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.P2sVpnGatewaysClient", "GenerateVpnProfile", nil, "Failure preparing request")
		return
	}

	result, err = client.GenerateVpnProfileSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.P2sVpnGatewaysClient", "GenerateVpnProfile", result.Response(), "Failure sending request")
		return
	}

	return
}

// GenerateVpnProfilePreparer prepares the GenerateVpnProfile request.
func (client P2sVpnGatewaysClient) GenerateVpnProfilePreparer(ctx context.Context, resourceGroupName string, gatewayName string, parameters P2SVpnProfileParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"gatewayName":       autorest.Encode("path", gatewayName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/p2svpnGateways/{gatewayName}/generatevpnprofile", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GenerateVpnProfileSender sends the GenerateVpnProfile request. The method will close the
// http.Response Body if it receives an error.
func (client P2sVpnGatewaysClient) GenerateVpnProfileSender(req *http.Request) (future P2sVpnGatewaysGenerateVpnProfileFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// GenerateVpnProfileResponder handles the response to the GenerateVpnProfile request. The method always
// closes the http.Response Body.
func (client P2sVpnGatewaysClient) GenerateVpnProfileResponder(resp *http.Response) (result VpnProfileResponse, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get retrieves the details of a virtual wan p2s vpn gateway.
// Parameters:
// resourceGroupName - the resource group name of the P2SVpnGateway.
// gatewayName - the name of the gateway.
func (client P2sVpnGatewaysClient) Get(ctx context.Context, resourceGroupName string, gatewayName string) (result P2SVpnGateway, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/P2sVpnGatewaysClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, gatewayName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.P2sVpnGatewaysClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.P2sVpnGatewaysClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.P2sVpnGatewaysClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client P2sVpnGatewaysClient) GetPreparer(ctx context.Context, resourceGroupName string, gatewayName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"gatewayName":       autorest.Encode("path", gatewayName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/p2svpnGateways/{gatewayName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client P2sVpnGatewaysClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client P2sVpnGatewaysClient) GetResponder(resp *http.Response) (result P2SVpnGateway, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetP2sVpnConnectionHealth gets the connection health of P2S clients of the virtual wan P2SVpnGateway in the
// specified resource group.
// Parameters:
// resourceGroupName - the name of the resource group.
// gatewayName - the name of the P2SVpnGateway.
func (client P2sVpnGatewaysClient) GetP2sVpnConnectionHealth(ctx context.Context, resourceGroupName string, gatewayName string) (result P2sVpnGatewaysGetP2sVpnConnectionHealthFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/P2sVpnGatewaysClient.GetP2sVpnConnectionHealth")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetP2sVpnConnectionHealthPreparer(ctx, resourceGroupName, gatewayName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.P2sVpnGatewaysClient", "GetP2sVpnConnectionHealth", nil, "Failure preparing request")
		return
	}

	result, err = client.GetP2sVpnConnectionHealthSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.P2sVpnGatewaysClient", "GetP2sVpnConnectionHealth", result.Response(), "Failure sending request")
		return
	}

	return
}

// GetP2sVpnConnectionHealthPreparer prepares the GetP2sVpnConnectionHealth request.
func (client P2sVpnGatewaysClient) GetP2sVpnConnectionHealthPreparer(ctx context.Context, resourceGroupName string, gatewayName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"gatewayName":       autorest.Encode("path", gatewayName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/p2svpnGateways/{gatewayName}/getP2sVpnConnectionHealth", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetP2sVpnConnectionHealthSender sends the GetP2sVpnConnectionHealth request. The method will close the
// http.Response Body if it receives an error.
func (client P2sVpnGatewaysClient) GetP2sVpnConnectionHealthSender(req *http.Request) (future P2sVpnGatewaysGetP2sVpnConnectionHealthFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// GetP2sVpnConnectionHealthResponder handles the response to the GetP2sVpnConnectionHealth request. The method always
// closes the http.Response Body.
func (client P2sVpnGatewaysClient) GetP2sVpnConnectionHealthResponder(resp *http.Response) (result P2SVpnGateway, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists all the P2SVpnGateways in a subscription.
func (client P2sVpnGatewaysClient) List(ctx context.Context) (result ListP2SVpnGatewaysResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/P2sVpnGatewaysClient.List")
		defer func() {
			sc := -1
			if result.lpvgr.Response.Response != nil {
				sc = result.lpvgr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.P2sVpnGatewaysClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.lpvgr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.P2sVpnGatewaysClient", "List", resp, "Failure sending request")
		return
	}

	result.lpvgr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.P2sVpnGatewaysClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client P2sVpnGatewaysClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Network/p2svpnGateways", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client P2sVpnGatewaysClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client P2sVpnGatewaysClient) ListResponder(resp *http.Response) (result ListP2SVpnGatewaysResult, err error) {
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
func (client P2sVpnGatewaysClient) listNextResults(ctx context.Context, lastResults ListP2SVpnGatewaysResult) (result ListP2SVpnGatewaysResult, err error) {
	req, err := lastResults.listP2SVpnGatewaysResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.P2sVpnGatewaysClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.P2sVpnGatewaysClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.P2sVpnGatewaysClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client P2sVpnGatewaysClient) ListComplete(ctx context.Context) (result ListP2SVpnGatewaysResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/P2sVpnGatewaysClient.List")
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

// ListByResourceGroup lists all the P2SVpnGateways in a resource group.
// Parameters:
// resourceGroupName - the resource group name of the P2SVpnGateway.
func (client P2sVpnGatewaysClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (result ListP2SVpnGatewaysResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/P2sVpnGatewaysClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.lpvgr.Response.Response != nil {
				sc = result.lpvgr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.P2sVpnGatewaysClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.lpvgr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.P2sVpnGatewaysClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.lpvgr, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.P2sVpnGatewaysClient", "ListByResourceGroup", resp, "Failure responding to request")
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client P2sVpnGatewaysClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/p2svpnGateways", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client P2sVpnGatewaysClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client P2sVpnGatewaysClient) ListByResourceGroupResponder(resp *http.Response) (result ListP2SVpnGatewaysResult, err error) {
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
func (client P2sVpnGatewaysClient) listByResourceGroupNextResults(ctx context.Context, lastResults ListP2SVpnGatewaysResult) (result ListP2SVpnGatewaysResult, err error) {
	req, err := lastResults.listP2SVpnGatewaysResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.P2sVpnGatewaysClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.P2sVpnGatewaysClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.P2sVpnGatewaysClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client P2sVpnGatewaysClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result ListP2SVpnGatewaysResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/P2sVpnGatewaysClient.ListByResourceGroup")
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

// UpdateTags updates virtual wan p2s vpn gateway tags.
// Parameters:
// resourceGroupName - the resource group name of the P2SVpnGateway.
// gatewayName - the name of the gateway.
// p2SVpnGatewayParameters - parameters supplied to update a virtual wan p2s vpn gateway tags.
func (client P2sVpnGatewaysClient) UpdateTags(ctx context.Context, resourceGroupName string, gatewayName string, p2SVpnGatewayParameters TagsObject) (result P2sVpnGatewaysUpdateTagsFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/P2sVpnGatewaysClient.UpdateTags")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdateTagsPreparer(ctx, resourceGroupName, gatewayName, p2SVpnGatewayParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.P2sVpnGatewaysClient", "UpdateTags", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateTagsSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.P2sVpnGatewaysClient", "UpdateTags", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdateTagsPreparer prepares the UpdateTags request.
func (client P2sVpnGatewaysClient) UpdateTagsPreparer(ctx context.Context, resourceGroupName string, gatewayName string, p2SVpnGatewayParameters TagsObject) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"gatewayName":       autorest.Encode("path", gatewayName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/p2svpnGateways/{gatewayName}", pathParameters),
		autorest.WithJSON(p2SVpnGatewayParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateTagsSender sends the UpdateTags request. The method will close the
// http.Response Body if it receives an error.
func (client P2sVpnGatewaysClient) UpdateTagsSender(req *http.Request) (future P2sVpnGatewaysUpdateTagsFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// UpdateTagsResponder handles the response to the UpdateTags request. The method always
// closes the http.Response Body.
func (client P2sVpnGatewaysClient) UpdateTagsResponder(resp *http.Response) (result P2SVpnGateway, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
