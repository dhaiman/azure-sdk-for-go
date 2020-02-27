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

// VirtualNetworkGatewaysClient is the network Client
type VirtualNetworkGatewaysClient struct {
	BaseClient
}

// NewVirtualNetworkGatewaysClient creates an instance of the VirtualNetworkGatewaysClient client.
func NewVirtualNetworkGatewaysClient(subscriptionID string) VirtualNetworkGatewaysClient {
	return NewVirtualNetworkGatewaysClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewVirtualNetworkGatewaysClientWithBaseURI creates an instance of the VirtualNetworkGatewaysClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewVirtualNetworkGatewaysClientWithBaseURI(baseURI string, subscriptionID string) VirtualNetworkGatewaysClient {
	return VirtualNetworkGatewaysClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a virtual network gateway in the specified resource group.
// Parameters:
// resourceGroupName - the name of the resource group.
// virtualNetworkGatewayName - the name of the virtual network gateway.
// parameters - parameters supplied to create or update virtual network gateway operation.
func (client VirtualNetworkGatewaysClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string, parameters VirtualNetworkGateway) (result VirtualNetworkGatewaysCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualNetworkGatewaysClient.CreateOrUpdate")
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
			Constraints: []validation.Constraint{{Target: "parameters.VirtualNetworkGatewayPropertiesFormat", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "parameters.VirtualNetworkGatewayPropertiesFormat.IPConfigurations", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("network.VirtualNetworkGatewaysClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, virtualNetworkGatewayName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewaysClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewaysClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client VirtualNetworkGatewaysClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string, parameters VirtualNetworkGateway) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":         autorest.Encode("path", resourceGroupName),
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
		"virtualNetworkGatewayName": autorest.Encode("path", virtualNetworkGatewayName),
	}

	const APIVersion = "2016-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkGateways/{virtualNetworkGatewayName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworkGatewaysClient) CreateOrUpdateSender(req *http.Request) (future VirtualNetworkGatewaysCreateOrUpdateFuture, err error) {
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
func (client VirtualNetworkGatewaysClient) CreateOrUpdateResponder(resp *http.Response) (result VirtualNetworkGateway, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the specified virtual network gateway.
// Parameters:
// resourceGroupName - the name of the resource group.
// virtualNetworkGatewayName - the name of the virtual network gateway.
func (client VirtualNetworkGatewaysClient) Delete(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string) (result VirtualNetworkGatewaysDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualNetworkGatewaysClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, virtualNetworkGatewayName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewaysClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewaysClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client VirtualNetworkGatewaysClient) DeletePreparer(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":         autorest.Encode("path", resourceGroupName),
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
		"virtualNetworkGatewayName": autorest.Encode("path", virtualNetworkGatewayName),
	}

	const APIVersion = "2016-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkGateways/{virtualNetworkGatewayName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworkGatewaysClient) DeleteSender(req *http.Request) (future VirtualNetworkGatewaysDeleteFuture, err error) {
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
func (client VirtualNetworkGatewaysClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Generatevpnclientpackage generates VPN client package for P2S client of the virtual network gateway in the specified
// resource group.
// Parameters:
// resourceGroupName - the name of the resource group.
// virtualNetworkGatewayName - the name of the virtual network gateway.
// parameters - parameters supplied to the generate virtual network gateway VPN client package operation.
func (client VirtualNetworkGatewaysClient) Generatevpnclientpackage(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string, parameters VpnClientParameters) (result VirtualNetworkGatewaysGeneratevpnclientpackageFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualNetworkGatewaysClient.Generatevpnclientpackage")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GeneratevpnclientpackagePreparer(ctx, resourceGroupName, virtualNetworkGatewayName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewaysClient", "Generatevpnclientpackage", nil, "Failure preparing request")
		return
	}

	result, err = client.GeneratevpnclientpackageSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewaysClient", "Generatevpnclientpackage", result.Response(), "Failure sending request")
		return
	}

	return
}

// GeneratevpnclientpackagePreparer prepares the Generatevpnclientpackage request.
func (client VirtualNetworkGatewaysClient) GeneratevpnclientpackagePreparer(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string, parameters VpnClientParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":         autorest.Encode("path", resourceGroupName),
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
		"virtualNetworkGatewayName": autorest.Encode("path", virtualNetworkGatewayName),
	}

	const APIVersion = "2016-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkGateways/{virtualNetworkGatewayName}/generatevpnclientpackage", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GeneratevpnclientpackageSender sends the Generatevpnclientpackage request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworkGatewaysClient) GeneratevpnclientpackageSender(req *http.Request) (future VirtualNetworkGatewaysGeneratevpnclientpackageFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// GeneratevpnclientpackageResponder handles the response to the Generatevpnclientpackage request. The method always
// closes the http.Response Body.
func (client VirtualNetworkGatewaysClient) GeneratevpnclientpackageResponder(resp *http.Response) (result String, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get gets the specified virtual network gateway by resource group.
// Parameters:
// resourceGroupName - the name of the resource group.
// virtualNetworkGatewayName - the name of the virtual network gateway.
func (client VirtualNetworkGatewaysClient) Get(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string) (result VirtualNetworkGateway, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualNetworkGatewaysClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, virtualNetworkGatewayName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewaysClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewaysClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewaysClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client VirtualNetworkGatewaysClient) GetPreparer(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":         autorest.Encode("path", resourceGroupName),
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
		"virtualNetworkGatewayName": autorest.Encode("path", virtualNetworkGatewayName),
	}

	const APIVersion = "2016-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkGateways/{virtualNetworkGatewayName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworkGatewaysClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client VirtualNetworkGatewaysClient) GetResponder(resp *http.Response) (result VirtualNetworkGateway, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetAdvertisedRoutes this operation retrieves a list of routes the virtual network gateway is advertising to the
// specified peer.
// Parameters:
// resourceGroupName - the name of the resource group.
// virtualNetworkGatewayName - the name of the virtual network gateway.
// peer - the IP address of the peer
func (client VirtualNetworkGatewaysClient) GetAdvertisedRoutes(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string, peer string) (result VirtualNetworkGatewaysGetAdvertisedRoutesFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualNetworkGatewaysClient.GetAdvertisedRoutes")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetAdvertisedRoutesPreparer(ctx, resourceGroupName, virtualNetworkGatewayName, peer)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewaysClient", "GetAdvertisedRoutes", nil, "Failure preparing request")
		return
	}

	result, err = client.GetAdvertisedRoutesSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewaysClient", "GetAdvertisedRoutes", result.Response(), "Failure sending request")
		return
	}

	return
}

// GetAdvertisedRoutesPreparer prepares the GetAdvertisedRoutes request.
func (client VirtualNetworkGatewaysClient) GetAdvertisedRoutesPreparer(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string, peer string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":         autorest.Encode("path", resourceGroupName),
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
		"virtualNetworkGatewayName": autorest.Encode("path", virtualNetworkGatewayName),
	}

	const APIVersion = "2016-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
		"peer":        autorest.Encode("query", peer),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkGateways/{virtualNetworkGatewayName}/getAdvertisedRoutes", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetAdvertisedRoutesSender sends the GetAdvertisedRoutes request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworkGatewaysClient) GetAdvertisedRoutesSender(req *http.Request) (future VirtualNetworkGatewaysGetAdvertisedRoutesFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// GetAdvertisedRoutesResponder handles the response to the GetAdvertisedRoutes request. The method always
// closes the http.Response Body.
func (client VirtualNetworkGatewaysClient) GetAdvertisedRoutesResponder(resp *http.Response) (result GatewayRouteListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetBgpPeerStatus the GetBgpPeerStatus operation retrieves the status of all BGP peers.
// Parameters:
// resourceGroupName - the name of the resource group.
// virtualNetworkGatewayName - the name of the virtual network gateway.
// peer - the IP address of the peer to retrieve the status of.
func (client VirtualNetworkGatewaysClient) GetBgpPeerStatus(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string, peer string) (result VirtualNetworkGatewaysGetBgpPeerStatusFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualNetworkGatewaysClient.GetBgpPeerStatus")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetBgpPeerStatusPreparer(ctx, resourceGroupName, virtualNetworkGatewayName, peer)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewaysClient", "GetBgpPeerStatus", nil, "Failure preparing request")
		return
	}

	result, err = client.GetBgpPeerStatusSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewaysClient", "GetBgpPeerStatus", result.Response(), "Failure sending request")
		return
	}

	return
}

// GetBgpPeerStatusPreparer prepares the GetBgpPeerStatus request.
func (client VirtualNetworkGatewaysClient) GetBgpPeerStatusPreparer(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string, peer string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":         autorest.Encode("path", resourceGroupName),
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
		"virtualNetworkGatewayName": autorest.Encode("path", virtualNetworkGatewayName),
	}

	const APIVersion = "2016-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(peer) > 0 {
		queryParameters["peer"] = autorest.Encode("query", peer)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkGateways/{virtualNetworkGatewayName}/getBgpPeerStatus", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetBgpPeerStatusSender sends the GetBgpPeerStatus request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworkGatewaysClient) GetBgpPeerStatusSender(req *http.Request) (future VirtualNetworkGatewaysGetBgpPeerStatusFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// GetBgpPeerStatusResponder handles the response to the GetBgpPeerStatus request. The method always
// closes the http.Response Body.
func (client VirtualNetworkGatewaysClient) GetBgpPeerStatusResponder(resp *http.Response) (result BgpPeerStatusListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetLearnedRoutes this operation retrieves a list of routes the virtual network gateway has learned, including routes
// learned from BGP peers.
// Parameters:
// resourceGroupName - the name of the resource group.
// virtualNetworkGatewayName - the name of the virtual network gateway.
func (client VirtualNetworkGatewaysClient) GetLearnedRoutes(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string) (result VirtualNetworkGatewaysGetLearnedRoutesFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualNetworkGatewaysClient.GetLearnedRoutes")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetLearnedRoutesPreparer(ctx, resourceGroupName, virtualNetworkGatewayName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewaysClient", "GetLearnedRoutes", nil, "Failure preparing request")
		return
	}

	result, err = client.GetLearnedRoutesSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewaysClient", "GetLearnedRoutes", result.Response(), "Failure sending request")
		return
	}

	return
}

// GetLearnedRoutesPreparer prepares the GetLearnedRoutes request.
func (client VirtualNetworkGatewaysClient) GetLearnedRoutesPreparer(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":         autorest.Encode("path", resourceGroupName),
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
		"virtualNetworkGatewayName": autorest.Encode("path", virtualNetworkGatewayName),
	}

	const APIVersion = "2016-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkGateways/{virtualNetworkGatewayName}/getLearnedRoutes", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetLearnedRoutesSender sends the GetLearnedRoutes request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworkGatewaysClient) GetLearnedRoutesSender(req *http.Request) (future VirtualNetworkGatewaysGetLearnedRoutesFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// GetLearnedRoutesResponder handles the response to the GetLearnedRoutes request. The method always
// closes the http.Response Body.
func (client VirtualNetworkGatewaysClient) GetLearnedRoutesResponder(resp *http.Response) (result GatewayRouteListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets all virtual network gateways by resource group.
// Parameters:
// resourceGroupName - the name of the resource group.
func (client VirtualNetworkGatewaysClient) List(ctx context.Context, resourceGroupName string) (result VirtualNetworkGatewayListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualNetworkGatewaysClient.List")
		defer func() {
			sc := -1
			if result.vnglr.Response.Response != nil {
				sc = result.vnglr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewaysClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.vnglr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewaysClient", "List", resp, "Failure sending request")
		return
	}

	result.vnglr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewaysClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client VirtualNetworkGatewaysClient) ListPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkGateways", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworkGatewaysClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client VirtualNetworkGatewaysClient) ListResponder(resp *http.Response) (result VirtualNetworkGatewayListResult, err error) {
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
func (client VirtualNetworkGatewaysClient) listNextResults(ctx context.Context, lastResults VirtualNetworkGatewayListResult) (result VirtualNetworkGatewayListResult, err error) {
	req, err := lastResults.virtualNetworkGatewayListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.VirtualNetworkGatewaysClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.VirtualNetworkGatewaysClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewaysClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client VirtualNetworkGatewaysClient) ListComplete(ctx context.Context, resourceGroupName string) (result VirtualNetworkGatewayListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualNetworkGatewaysClient.List")
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

// Reset resets the primary of the virtual network gateway in the specified resource group.
// Parameters:
// resourceGroupName - the name of the resource group.
// virtualNetworkGatewayName - the name of the virtual network gateway.
// gatewayVip - virtual network gateway vip address supplied to the begin reset of the active-active feature
// enabled gateway.
func (client VirtualNetworkGatewaysClient) Reset(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string, gatewayVip string) (result VirtualNetworkGatewaysResetFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualNetworkGatewaysClient.Reset")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ResetPreparer(ctx, resourceGroupName, virtualNetworkGatewayName, gatewayVip)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewaysClient", "Reset", nil, "Failure preparing request")
		return
	}

	result, err = client.ResetSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewaysClient", "Reset", result.Response(), "Failure sending request")
		return
	}

	return
}

// ResetPreparer prepares the Reset request.
func (client VirtualNetworkGatewaysClient) ResetPreparer(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string, gatewayVip string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":         autorest.Encode("path", resourceGroupName),
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
		"virtualNetworkGatewayName": autorest.Encode("path", virtualNetworkGatewayName),
	}

	const APIVersion = "2016-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(gatewayVip) > 0 {
		queryParameters["gatewayVip"] = autorest.Encode("query", gatewayVip)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkGateways/{virtualNetworkGatewayName}/reset", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ResetSender sends the Reset request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworkGatewaysClient) ResetSender(req *http.Request) (future VirtualNetworkGatewaysResetFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// ResetResponder handles the response to the Reset request. The method always
// closes the http.Response Body.
func (client VirtualNetworkGatewaysClient) ResetResponder(resp *http.Response) (result VirtualNetworkGateway, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
