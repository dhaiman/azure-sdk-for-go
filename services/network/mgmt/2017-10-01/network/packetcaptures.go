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

// PacketCapturesClient is the network Client
type PacketCapturesClient struct {
	BaseClient
}

// NewPacketCapturesClient creates an instance of the PacketCapturesClient client.
func NewPacketCapturesClient(subscriptionID string) PacketCapturesClient {
	return NewPacketCapturesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewPacketCapturesClientWithBaseURI creates an instance of the PacketCapturesClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewPacketCapturesClientWithBaseURI(baseURI string, subscriptionID string) PacketCapturesClient {
	return PacketCapturesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create create and start a packet capture on the specified VM.
// Parameters:
// resourceGroupName - the name of the resource group.
// networkWatcherName - the name of the network watcher.
// packetCaptureName - the name of the packet capture session.
// parameters - parameters that define the create packet capture operation.
func (client PacketCapturesClient) Create(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string, parameters PacketCapture) (result PacketCapturesCreateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PacketCapturesClient.Create")
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
			Constraints: []validation.Constraint{{Target: "parameters.PacketCaptureParameters", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "parameters.PacketCaptureParameters.Target", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.PacketCaptureParameters.StorageLocation", Name: validation.Null, Rule: true, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("network.PacketCapturesClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, resourceGroupName, networkWatcherName, packetCaptureName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PacketCapturesClient", "Create", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PacketCapturesClient", "Create", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client PacketCapturesClient) CreatePreparer(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string, parameters PacketCapture) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"networkWatcherName": autorest.Encode("path", networkWatcherName),
		"packetCaptureName":  autorest.Encode("path", packetCaptureName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/packetCaptures/{packetCaptureName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client PacketCapturesClient) CreateSender(req *http.Request) (future PacketCapturesCreateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client PacketCapturesClient) CreateResponder(resp *http.Response) (result PacketCaptureResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the specified packet capture session.
// Parameters:
// resourceGroupName - the name of the resource group.
// networkWatcherName - the name of the network watcher.
// packetCaptureName - the name of the packet capture session.
func (client PacketCapturesClient) Delete(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string) (result PacketCapturesDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PacketCapturesClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, networkWatcherName, packetCaptureName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PacketCapturesClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PacketCapturesClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client PacketCapturesClient) DeletePreparer(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"networkWatcherName": autorest.Encode("path", networkWatcherName),
		"packetCaptureName":  autorest.Encode("path", packetCaptureName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/packetCaptures/{packetCaptureName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client PacketCapturesClient) DeleteSender(req *http.Request) (future PacketCapturesDeleteFuture, err error) {
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
func (client PacketCapturesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets a packet capture session by name.
// Parameters:
// resourceGroupName - the name of the resource group.
// networkWatcherName - the name of the network watcher.
// packetCaptureName - the name of the packet capture session.
func (client PacketCapturesClient) Get(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string) (result PacketCaptureResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PacketCapturesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, networkWatcherName, packetCaptureName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PacketCapturesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.PacketCapturesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PacketCapturesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client PacketCapturesClient) GetPreparer(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"networkWatcherName": autorest.Encode("path", networkWatcherName),
		"packetCaptureName":  autorest.Encode("path", packetCaptureName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/packetCaptures/{packetCaptureName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client PacketCapturesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client PacketCapturesClient) GetResponder(resp *http.Response) (result PacketCaptureResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetStatus query the status of a running packet capture session.
// Parameters:
// resourceGroupName - the name of the resource group.
// networkWatcherName - the name of the Network Watcher resource.
// packetCaptureName - the name given to the packet capture session.
func (client PacketCapturesClient) GetStatus(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string) (result PacketCapturesGetStatusFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PacketCapturesClient.GetStatus")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetStatusPreparer(ctx, resourceGroupName, networkWatcherName, packetCaptureName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PacketCapturesClient", "GetStatus", nil, "Failure preparing request")
		return
	}

	result, err = client.GetStatusSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PacketCapturesClient", "GetStatus", result.Response(), "Failure sending request")
		return
	}

	return
}

// GetStatusPreparer prepares the GetStatus request.
func (client PacketCapturesClient) GetStatusPreparer(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"networkWatcherName": autorest.Encode("path", networkWatcherName),
		"packetCaptureName":  autorest.Encode("path", packetCaptureName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/packetCaptures/{packetCaptureName}/queryStatus", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetStatusSender sends the GetStatus request. The method will close the
// http.Response Body if it receives an error.
func (client PacketCapturesClient) GetStatusSender(req *http.Request) (future PacketCapturesGetStatusFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// GetStatusResponder handles the response to the GetStatus request. The method always
// closes the http.Response Body.
func (client PacketCapturesClient) GetStatusResponder(resp *http.Response) (result PacketCaptureQueryStatusResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists all packet capture sessions within the specified resource group.
// Parameters:
// resourceGroupName - the name of the resource group.
// networkWatcherName - the name of the Network Watcher resource.
func (client PacketCapturesClient) List(ctx context.Context, resourceGroupName string, networkWatcherName string) (result PacketCaptureListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PacketCapturesClient.List")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListPreparer(ctx, resourceGroupName, networkWatcherName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PacketCapturesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.PacketCapturesClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PacketCapturesClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client PacketCapturesClient) ListPreparer(ctx context.Context, resourceGroupName string, networkWatcherName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"networkWatcherName": autorest.Encode("path", networkWatcherName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/packetCaptures", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client PacketCapturesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client PacketCapturesClient) ListResponder(resp *http.Response) (result PacketCaptureListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Stop stops a specified packet capture session.
// Parameters:
// resourceGroupName - the name of the resource group.
// networkWatcherName - the name of the network watcher.
// packetCaptureName - the name of the packet capture session.
func (client PacketCapturesClient) Stop(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string) (result PacketCapturesStopFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PacketCapturesClient.Stop")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.StopPreparer(ctx, resourceGroupName, networkWatcherName, packetCaptureName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PacketCapturesClient", "Stop", nil, "Failure preparing request")
		return
	}

	result, err = client.StopSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PacketCapturesClient", "Stop", result.Response(), "Failure sending request")
		return
	}

	return
}

// StopPreparer prepares the Stop request.
func (client PacketCapturesClient) StopPreparer(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"networkWatcherName": autorest.Encode("path", networkWatcherName),
		"packetCaptureName":  autorest.Encode("path", packetCaptureName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/packetCaptures/{packetCaptureName}/stop", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// StopSender sends the Stop request. The method will close the
// http.Response Body if it receives an error.
func (client PacketCapturesClient) StopSender(req *http.Request) (future PacketCapturesStopFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// StopResponder handles the response to the Stop request. The method always
// closes the http.Response Body.
func (client PacketCapturesClient) StopResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}
