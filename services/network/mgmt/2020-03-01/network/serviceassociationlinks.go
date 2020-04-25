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

// ServiceAssociationLinksClient is the network Client
type ServiceAssociationLinksClient struct {
	BaseClient
}

// NewServiceAssociationLinksClient creates an instance of the ServiceAssociationLinksClient client.
func NewServiceAssociationLinksClient(subscriptionID string) ServiceAssociationLinksClient {
	return NewServiceAssociationLinksClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewServiceAssociationLinksClientWithBaseURI creates an instance of the ServiceAssociationLinksClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewServiceAssociationLinksClientWithBaseURI(baseURI string, subscriptionID string) ServiceAssociationLinksClient {
	return ServiceAssociationLinksClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List gets a list of service association links for a subnet.
// Parameters:
// resourceGroupName - the name of the resource group.
// virtualNetworkName - the name of the virtual network.
// subnetName - the name of the subnet.
func (client ServiceAssociationLinksClient) List(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string) (result ServiceAssociationLinksListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServiceAssociationLinksClient.List")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListPreparer(ctx, resourceGroupName, virtualNetworkName, subnetName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ServiceAssociationLinksClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.ServiceAssociationLinksClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ServiceAssociationLinksClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ServiceAssociationLinksClient) ListPreparer(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subnetName":         autorest.Encode("path", subnetName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
		"virtualNetworkName": autorest.Encode("path", virtualNetworkName),
	}

	const APIVersion = "2020-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}/ServiceAssociationLinks", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ServiceAssociationLinksClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ServiceAssociationLinksClient) ListResponder(resp *http.Response) (result ServiceAssociationLinksListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
