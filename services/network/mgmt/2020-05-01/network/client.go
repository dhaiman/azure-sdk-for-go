// Package network implements the Azure ARM Network service API version .
//
// Network Client
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

const (
	// DefaultBaseURI is the default URI used for the service Network
	DefaultBaseURI = "https://management.azure.com"
)

// BaseClient is the base client for Network.
type BaseClient struct {
	autorest.Client
	BaseURI        string
	SubscriptionID string
}

// New creates an instance of the BaseClient client.
func New(subscriptionID string) BaseClient {
	return NewWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWithBaseURI creates an instance of the BaseClient client using a custom endpoint.  Use this when interacting with
// an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return BaseClient{
		Client:         autorest.NewClientWithUserAgent(UserAgent()),
		BaseURI:        baseURI,
		SubscriptionID: subscriptionID,
	}
}

// CheckDNSNameAvailability checks whether a domain name in the cloudapp.azure.com zone is available for use.
// Parameters:
// location - the location of the domain name.
// domainNameLabel - the domain name to be verified. It must conform to the following regular expression:
// ^[a-z][a-z0-9-]{1,61}[a-z0-9]$.
func (client BaseClient) CheckDNSNameAvailability(ctx context.Context, location string, domainNameLabel string) (result DNSNameAvailabilityResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.CheckDNSNameAvailability")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CheckDNSNameAvailabilityPreparer(ctx, location, domainNameLabel)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.BaseClient", "CheckDNSNameAvailability", nil, "Failure preparing request")
		return
	}

	resp, err := client.CheckDNSNameAvailabilitySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.BaseClient", "CheckDNSNameAvailability", resp, "Failure sending request")
		return
	}

	result, err = client.CheckDNSNameAvailabilityResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.BaseClient", "CheckDNSNameAvailability", resp, "Failure responding to request")
	}

	return
}

// CheckDNSNameAvailabilityPreparer prepares the CheckDNSNameAvailability request.
func (client BaseClient) CheckDNSNameAvailabilityPreparer(ctx context.Context, location string, domainNameLabel string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       autorest.Encode("path", location),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-05-01"
	queryParameters := map[string]interface{}{
		"api-version":     APIVersion,
		"domainNameLabel": autorest.Encode("query", domainNameLabel),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Network/locations/{location}/CheckDnsNameAvailability", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CheckDNSNameAvailabilitySender sends the CheckDNSNameAvailability request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) CheckDNSNameAvailabilitySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CheckDNSNameAvailabilityResponder handles the response to the CheckDNSNameAvailability request. The method always
// closes the http.Response Body.
func (client BaseClient) CheckDNSNameAvailabilityResponder(resp *http.Response) (result DNSNameAvailabilityResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// DeleteBastionShareableLink deletes the Bastion Shareable Links for all the VMs specified in the request.
// Parameters:
// resourceGroupName - the name of the resource group.
// bastionHostName - the name of the Bastion Host.
// bslRequest - post request for all the Bastion Shareable Link endpoints.
func (client BaseClient) DeleteBastionShareableLink(ctx context.Context, resourceGroupName string, bastionHostName string, bslRequest BastionShareableLinkListRequest) (result DeleteBastionShareableLinkFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.DeleteBastionShareableLink")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeleteBastionShareableLinkPreparer(ctx, resourceGroupName, bastionHostName, bslRequest)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.BaseClient", "DeleteBastionShareableLink", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteBastionShareableLinkSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.BaseClient", "DeleteBastionShareableLink", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeleteBastionShareableLinkPreparer prepares the DeleteBastionShareableLink request.
func (client BaseClient) DeleteBastionShareableLinkPreparer(ctx context.Context, resourceGroupName string, bastionHostName string, bslRequest BastionShareableLinkListRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"bastionHostName":   autorest.Encode("path", bastionHostName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/bastionHosts/{bastionHostName}/deleteShareableLinks", pathParameters),
		autorest.WithJSON(bslRequest),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteBastionShareableLinkSender sends the DeleteBastionShareableLink request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) DeleteBastionShareableLinkSender(req *http.Request) (future DeleteBastionShareableLinkFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// DeleteBastionShareableLinkResponder handles the response to the DeleteBastionShareableLink request. The method always
// closes the http.Response Body.
func (client BaseClient) DeleteBastionShareableLinkResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// DisconnectActiveSessions returns the list of currently active sessions on the Bastion.
// Parameters:
// resourceGroupName - the name of the resource group.
// bastionHostName - the name of the Bastion Host.
// sessionIds - the list of sessionids to disconnect.
func (client BaseClient) DisconnectActiveSessions(ctx context.Context, resourceGroupName string, bastionHostName string, sessionIds SessionIds) (result BastionSessionDeleteResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.DisconnectActiveSessions")
		defer func() {
			sc := -1
			if result.bsdr.Response.Response != nil {
				sc = result.bsdr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.disconnectActiveSessionsNextResults
	req, err := client.DisconnectActiveSessionsPreparer(ctx, resourceGroupName, bastionHostName, sessionIds)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.BaseClient", "DisconnectActiveSessions", nil, "Failure preparing request")
		return
	}

	resp, err := client.DisconnectActiveSessionsSender(req)
	if err != nil {
		result.bsdr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.BaseClient", "DisconnectActiveSessions", resp, "Failure sending request")
		return
	}

	result.bsdr, err = client.DisconnectActiveSessionsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.BaseClient", "DisconnectActiveSessions", resp, "Failure responding to request")
	}

	return
}

// DisconnectActiveSessionsPreparer prepares the DisconnectActiveSessions request.
func (client BaseClient) DisconnectActiveSessionsPreparer(ctx context.Context, resourceGroupName string, bastionHostName string, sessionIds SessionIds) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"bastionHostName":   autorest.Encode("path", bastionHostName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/bastionHosts/{bastionHostName}/disconnectActiveSessions", pathParameters),
		autorest.WithJSON(sessionIds),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DisconnectActiveSessionsSender sends the DisconnectActiveSessions request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) DisconnectActiveSessionsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DisconnectActiveSessionsResponder handles the response to the DisconnectActiveSessions request. The method always
// closes the http.Response Body.
func (client BaseClient) DisconnectActiveSessionsResponder(resp *http.Response) (result BastionSessionDeleteResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// disconnectActiveSessionsNextResults retrieves the next set of results, if any.
func (client BaseClient) disconnectActiveSessionsNextResults(ctx context.Context, lastResults BastionSessionDeleteResult) (result BastionSessionDeleteResult, err error) {
	req, err := lastResults.bastionSessionDeleteResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.BaseClient", "disconnectActiveSessionsNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.DisconnectActiveSessionsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.BaseClient", "disconnectActiveSessionsNextResults", resp, "Failure sending next results request")
	}
	result, err = client.DisconnectActiveSessionsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.BaseClient", "disconnectActiveSessionsNextResults", resp, "Failure responding to next results request")
	}
	return
}

// DisconnectActiveSessionsComplete enumerates all values, automatically crossing page boundaries as required.
func (client BaseClient) DisconnectActiveSessionsComplete(ctx context.Context, resourceGroupName string, bastionHostName string, sessionIds SessionIds) (result BastionSessionDeleteResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.DisconnectActiveSessions")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.DisconnectActiveSessions(ctx, resourceGroupName, bastionHostName, sessionIds)
	return
}

// Generatevirtualwanvpnserverconfigurationvpnprofile generates a unique VPN profile for P2S clients for VirtualWan and
// associated VpnServerConfiguration combination in the specified resource group.
// Parameters:
// resourceGroupName - the resource group name.
// virtualWANName - the name of the VirtualWAN whose associated VpnServerConfigurations is needed.
// vpnClientParams - parameters supplied to the generate VirtualWan VPN profile generation operation.
func (client BaseClient) Generatevirtualwanvpnserverconfigurationvpnprofile(ctx context.Context, resourceGroupName string, virtualWANName string, vpnClientParams VirtualWanVpnProfileParameters) (result GeneratevirtualwanvpnserverconfigurationvpnprofileFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.Generatevirtualwanvpnserverconfigurationvpnprofile")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GeneratevirtualwanvpnserverconfigurationvpnprofilePreparer(ctx, resourceGroupName, virtualWANName, vpnClientParams)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.BaseClient", "Generatevirtualwanvpnserverconfigurationvpnprofile", nil, "Failure preparing request")
		return
	}

	result, err = client.GeneratevirtualwanvpnserverconfigurationvpnprofileSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.BaseClient", "Generatevirtualwanvpnserverconfigurationvpnprofile", result.Response(), "Failure sending request")
		return
	}

	return
}

// GeneratevirtualwanvpnserverconfigurationvpnprofilePreparer prepares the Generatevirtualwanvpnserverconfigurationvpnprofile request.
func (client BaseClient) GeneratevirtualwanvpnserverconfigurationvpnprofilePreparer(ctx context.Context, resourceGroupName string, virtualWANName string, vpnClientParams VirtualWanVpnProfileParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"virtualWANName":    autorest.Encode("path", virtualWANName),
	}

	const APIVersion = "2020-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualWans/{virtualWANName}/GenerateVpnProfile", pathParameters),
		autorest.WithJSON(vpnClientParams),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GeneratevirtualwanvpnserverconfigurationvpnprofileSender sends the Generatevirtualwanvpnserverconfigurationvpnprofile request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) GeneratevirtualwanvpnserverconfigurationvpnprofileSender(req *http.Request) (future GeneratevirtualwanvpnserverconfigurationvpnprofileFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// GeneratevirtualwanvpnserverconfigurationvpnprofileResponder handles the response to the Generatevirtualwanvpnserverconfigurationvpnprofile request. The method always
// closes the http.Response Body.
func (client BaseClient) GeneratevirtualwanvpnserverconfigurationvpnprofileResponder(resp *http.Response) (result VpnProfileResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetActiveSessions returns the list of currently active sessions on the Bastion.
// Parameters:
// resourceGroupName - the name of the resource group.
// bastionHostName - the name of the Bastion Host.
func (client BaseClient) GetActiveSessions(ctx context.Context, resourceGroupName string, bastionHostName string) (result GetActiveSessionsFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.GetActiveSessions")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetActiveSessionsPreparer(ctx, resourceGroupName, bastionHostName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.BaseClient", "GetActiveSessions", nil, "Failure preparing request")
		return
	}

	result, err = client.GetActiveSessionsSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.BaseClient", "GetActiveSessions", result.Response(), "Failure sending request")
		return
	}

	return
}

// GetActiveSessionsPreparer prepares the GetActiveSessions request.
func (client BaseClient) GetActiveSessionsPreparer(ctx context.Context, resourceGroupName string, bastionHostName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"bastionHostName":   autorest.Encode("path", bastionHostName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/bastionHosts/{bastionHostName}/getActiveSessions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetActiveSessionsSender sends the GetActiveSessions request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) GetActiveSessionsSender(req *http.Request) (future GetActiveSessionsFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// GetActiveSessionsResponder handles the response to the GetActiveSessions request. The method always
// closes the http.Response Body.
func (client BaseClient) GetActiveSessionsResponder(resp *http.Response) (result BastionActiveSessionListResultPage, err error) {
	result.baslr, err = client.getActiveSessionsResponder(resp)
	result.fn = client.getActiveSessionsNextResults
	return
}

func (client BaseClient) getActiveSessionsResponder(resp *http.Response) (result BastionActiveSessionListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// getActiveSessionsNextResults retrieves the next set of results, if any.
func (client BaseClient) getActiveSessionsNextResults(ctx context.Context, lastResults BastionActiveSessionListResult) (result BastionActiveSessionListResult, err error) {
	req, err := lastResults.bastionActiveSessionListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.BaseClient", "getActiveSessionsNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	var resp *http.Response
	resp, err = client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.BaseClient", "getActiveSessionsNextResults", resp, "Failure sending next results request")
	}
	return client.getActiveSessionsResponder(resp)
}

// GetActiveSessionsComplete enumerates all values, automatically crossing page boundaries as required.
func (client BaseClient) GetActiveSessionsComplete(ctx context.Context, resourceGroupName string, bastionHostName string) (result GetActiveSessionsAllFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.GetActiveSessions")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	var future GetActiveSessionsFuture
	future, err = client.GetActiveSessions(ctx, resourceGroupName, bastionHostName)
	result.Future = future.Future
	return
}

// GetBastionShareableLink return the Bastion Shareable Links for all the VMs specified in the request.
// Parameters:
// resourceGroupName - the name of the resource group.
// bastionHostName - the name of the Bastion Host.
// bslRequest - post request for all the Bastion Shareable Link endpoints.
func (client BaseClient) GetBastionShareableLink(ctx context.Context, resourceGroupName string, bastionHostName string, bslRequest BastionShareableLinkListRequest) (result BastionShareableLinkListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.GetBastionShareableLink")
		defer func() {
			sc := -1
			if result.bsllr.Response.Response != nil {
				sc = result.bsllr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.getBastionShareableLinkNextResults
	req, err := client.GetBastionShareableLinkPreparer(ctx, resourceGroupName, bastionHostName, bslRequest)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.BaseClient", "GetBastionShareableLink", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetBastionShareableLinkSender(req)
	if err != nil {
		result.bsllr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.BaseClient", "GetBastionShareableLink", resp, "Failure sending request")
		return
	}

	result.bsllr, err = client.GetBastionShareableLinkResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.BaseClient", "GetBastionShareableLink", resp, "Failure responding to request")
	}

	return
}

// GetBastionShareableLinkPreparer prepares the GetBastionShareableLink request.
func (client BaseClient) GetBastionShareableLinkPreparer(ctx context.Context, resourceGroupName string, bastionHostName string, bslRequest BastionShareableLinkListRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"bastionHostName":   autorest.Encode("path", bastionHostName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/bastionHosts/{bastionHostName}/getShareableLinks", pathParameters),
		autorest.WithJSON(bslRequest),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetBastionShareableLinkSender sends the GetBastionShareableLink request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) GetBastionShareableLinkSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetBastionShareableLinkResponder handles the response to the GetBastionShareableLink request. The method always
// closes the http.Response Body.
func (client BaseClient) GetBastionShareableLinkResponder(resp *http.Response) (result BastionShareableLinkListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// getBastionShareableLinkNextResults retrieves the next set of results, if any.
func (client BaseClient) getBastionShareableLinkNextResults(ctx context.Context, lastResults BastionShareableLinkListResult) (result BastionShareableLinkListResult, err error) {
	req, err := lastResults.bastionShareableLinkListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.BaseClient", "getBastionShareableLinkNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.GetBastionShareableLinkSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.BaseClient", "getBastionShareableLinkNextResults", resp, "Failure sending next results request")
	}
	result, err = client.GetBastionShareableLinkResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.BaseClient", "getBastionShareableLinkNextResults", resp, "Failure responding to next results request")
	}
	return
}

// GetBastionShareableLinkComplete enumerates all values, automatically crossing page boundaries as required.
func (client BaseClient) GetBastionShareableLinkComplete(ctx context.Context, resourceGroupName string, bastionHostName string, bslRequest BastionShareableLinkListRequest) (result BastionShareableLinkListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.GetBastionShareableLink")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.GetBastionShareableLink(ctx, resourceGroupName, bastionHostName, bslRequest)
	return
}

// PutBastionShareableLink creates a Bastion Shareable Links for all the VMs specified in the request.
// Parameters:
// resourceGroupName - the name of the resource group.
// bastionHostName - the name of the Bastion Host.
// bslRequest - post request for all the Bastion Shareable Link endpoints.
func (client BaseClient) PutBastionShareableLink(ctx context.Context, resourceGroupName string, bastionHostName string, bslRequest BastionShareableLinkListRequest) (result PutBastionShareableLinkFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.PutBastionShareableLink")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.PutBastionShareableLinkPreparer(ctx, resourceGroupName, bastionHostName, bslRequest)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.BaseClient", "PutBastionShareableLink", nil, "Failure preparing request")
		return
	}

	result, err = client.PutBastionShareableLinkSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.BaseClient", "PutBastionShareableLink", result.Response(), "Failure sending request")
		return
	}

	return
}

// PutBastionShareableLinkPreparer prepares the PutBastionShareableLink request.
func (client BaseClient) PutBastionShareableLinkPreparer(ctx context.Context, resourceGroupName string, bastionHostName string, bslRequest BastionShareableLinkListRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"bastionHostName":   autorest.Encode("path", bastionHostName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/bastionHosts/{bastionHostName}/createShareableLinks", pathParameters),
		autorest.WithJSON(bslRequest),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PutBastionShareableLinkSender sends the PutBastionShareableLink request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) PutBastionShareableLinkSender(req *http.Request) (future PutBastionShareableLinkFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// PutBastionShareableLinkResponder handles the response to the PutBastionShareableLink request. The method always
// closes the http.Response Body.
func (client BaseClient) PutBastionShareableLinkResponder(resp *http.Response) (result BastionShareableLinkListResultPage, err error) {
	result.bsllr, err = client.putBastionShareableLinkResponder(resp)
	result.fn = client.putBastionShareableLinkNextResults
	return
}

func (client BaseClient) putBastionShareableLinkResponder(resp *http.Response) (result BastionShareableLinkListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// putBastionShareableLinkNextResults retrieves the next set of results, if any.
func (client BaseClient) putBastionShareableLinkNextResults(ctx context.Context, lastResults BastionShareableLinkListResult) (result BastionShareableLinkListResult, err error) {
	req, err := lastResults.bastionShareableLinkListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.BaseClient", "putBastionShareableLinkNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	var resp *http.Response
	resp, err = client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.BaseClient", "putBastionShareableLinkNextResults", resp, "Failure sending next results request")
	}
	return client.putBastionShareableLinkResponder(resp)
}

// PutBastionShareableLinkComplete enumerates all values, automatically crossing page boundaries as required.
func (client BaseClient) PutBastionShareableLinkComplete(ctx context.Context, resourceGroupName string, bastionHostName string, bslRequest BastionShareableLinkListRequest) (result PutBastionShareableLinkAllFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.PutBastionShareableLink")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	var future PutBastionShareableLinkFuture
	future, err = client.PutBastionShareableLink(ctx, resourceGroupName, bastionHostName, bslRequest)
	result.Future = future.Future
	return
}

// SupportedSecurityProviders gives the supported security providers for the virtual wan.
// Parameters:
// resourceGroupName - the resource group name.
// virtualWANName - the name of the VirtualWAN for which supported security providers are needed.
func (client BaseClient) SupportedSecurityProviders(ctx context.Context, resourceGroupName string, virtualWANName string) (result VirtualWanSecurityProviders, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.SupportedSecurityProviders")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.SupportedSecurityProvidersPreparer(ctx, resourceGroupName, virtualWANName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.BaseClient", "SupportedSecurityProviders", nil, "Failure preparing request")
		return
	}

	resp, err := client.SupportedSecurityProvidersSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.BaseClient", "SupportedSecurityProviders", resp, "Failure sending request")
		return
	}

	result, err = client.SupportedSecurityProvidersResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.BaseClient", "SupportedSecurityProviders", resp, "Failure responding to request")
	}

	return
}

// SupportedSecurityProvidersPreparer prepares the SupportedSecurityProviders request.
func (client BaseClient) SupportedSecurityProvidersPreparer(ctx context.Context, resourceGroupName string, virtualWANName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"virtualWANName":    autorest.Encode("path", virtualWANName),
	}

	const APIVersion = "2020-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualWans/{virtualWANName}/supportedSecurityProviders", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// SupportedSecurityProvidersSender sends the SupportedSecurityProviders request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) SupportedSecurityProvidersSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// SupportedSecurityProvidersResponder handles the response to the SupportedSecurityProviders request. The method always
// closes the http.Response Body.
func (client BaseClient) SupportedSecurityProvidersResponder(resp *http.Response) (result VirtualWanSecurityProviders, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
