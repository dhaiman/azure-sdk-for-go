package siterecovery

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

// ReplicationLogicalNetworksClient is the client for the ReplicationLogicalNetworks methods of the Siterecovery
// service.
type ReplicationLogicalNetworksClient struct {
	BaseClient
}

// NewReplicationLogicalNetworksClient creates an instance of the ReplicationLogicalNetworksClient client.
func NewReplicationLogicalNetworksClient(subscriptionID string, resourceGroupName string, resourceName string) ReplicationLogicalNetworksClient {
	return NewReplicationLogicalNetworksClientWithBaseURI(DefaultBaseURI, subscriptionID, resourceGroupName, resourceName)
}

// NewReplicationLogicalNetworksClientWithBaseURI creates an instance of the ReplicationLogicalNetworksClient client
// using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
func NewReplicationLogicalNetworksClientWithBaseURI(baseURI string, subscriptionID string, resourceGroupName string, resourceName string) ReplicationLogicalNetworksClient {
	return ReplicationLogicalNetworksClient{NewWithBaseURI(baseURI, subscriptionID, resourceGroupName, resourceName)}
}

// Get gets the details of a logical network.
// Parameters:
// fabricName - server Id.
// logicalNetworkName - logical network name.
func (client ReplicationLogicalNetworksClient) Get(ctx context.Context, fabricName string, logicalNetworkName string) (result LogicalNetwork, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReplicationLogicalNetworksClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, fabricName, logicalNetworkName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationLogicalNetworksClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationLogicalNetworksClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationLogicalNetworksClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ReplicationLogicalNetworksClient) GetPreparer(ctx context.Context, fabricName string, logicalNetworkName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fabricName":         autorest.Encode("path", fabricName),
		"logicalNetworkName": autorest.Encode("path", logicalNetworkName),
		"resourceGroupName":  autorest.Encode("path", client.ResourceGroupName),
		"resourceName":       autorest.Encode("path", client.ResourceName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-01-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationLogicalNetworks/{logicalNetworkName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationLogicalNetworksClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ReplicationLogicalNetworksClient) GetResponder(resp *http.Response) (result LogicalNetwork, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByReplicationFabrics lists all the logical networks of the Azure Site Recovery fabric
// Parameters:
// fabricName - server Id.
func (client ReplicationLogicalNetworksClient) ListByReplicationFabrics(ctx context.Context, fabricName string) (result LogicalNetworkCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReplicationLogicalNetworksClient.ListByReplicationFabrics")
		defer func() {
			sc := -1
			if result.lnc.Response.Response != nil {
				sc = result.lnc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByReplicationFabricsNextResults
	req, err := client.ListByReplicationFabricsPreparer(ctx, fabricName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationLogicalNetworksClient", "ListByReplicationFabrics", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByReplicationFabricsSender(req)
	if err != nil {
		result.lnc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationLogicalNetworksClient", "ListByReplicationFabrics", resp, "Failure sending request")
		return
	}

	result.lnc, err = client.ListByReplicationFabricsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationLogicalNetworksClient", "ListByReplicationFabrics", resp, "Failure responding to request")
	}

	return
}

// ListByReplicationFabricsPreparer prepares the ListByReplicationFabrics request.
func (client ReplicationLogicalNetworksClient) ListByReplicationFabricsPreparer(ctx context.Context, fabricName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fabricName":        autorest.Encode("path", fabricName),
		"resourceGroupName": autorest.Encode("path", client.ResourceGroupName),
		"resourceName":      autorest.Encode("path", client.ResourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-01-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationLogicalNetworks", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByReplicationFabricsSender sends the ListByReplicationFabrics request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationLogicalNetworksClient) ListByReplicationFabricsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByReplicationFabricsResponder handles the response to the ListByReplicationFabrics request. The method always
// closes the http.Response Body.
func (client ReplicationLogicalNetworksClient) ListByReplicationFabricsResponder(resp *http.Response) (result LogicalNetworkCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByReplicationFabricsNextResults retrieves the next set of results, if any.
func (client ReplicationLogicalNetworksClient) listByReplicationFabricsNextResults(ctx context.Context, lastResults LogicalNetworkCollection) (result LogicalNetworkCollection, err error) {
	req, err := lastResults.logicalNetworkCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "siterecovery.ReplicationLogicalNetworksClient", "listByReplicationFabricsNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByReplicationFabricsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "siterecovery.ReplicationLogicalNetworksClient", "listByReplicationFabricsNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByReplicationFabricsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationLogicalNetworksClient", "listByReplicationFabricsNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByReplicationFabricsComplete enumerates all values, automatically crossing page boundaries as required.
func (client ReplicationLogicalNetworksClient) ListByReplicationFabricsComplete(ctx context.Context, fabricName string) (result LogicalNetworkCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReplicationLogicalNetworksClient.ListByReplicationFabrics")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByReplicationFabrics(ctx, fabricName)
	return
}
