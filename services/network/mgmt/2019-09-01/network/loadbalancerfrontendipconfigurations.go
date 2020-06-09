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

// LoadBalancerFrontendIPConfigurationsClient is the network Client
type LoadBalancerFrontendIPConfigurationsClient struct {
	BaseClient
}

// NewLoadBalancerFrontendIPConfigurationsClient creates an instance of the LoadBalancerFrontendIPConfigurationsClient
// client.
func NewLoadBalancerFrontendIPConfigurationsClient(subscriptionID string) LoadBalancerFrontendIPConfigurationsClient {
	return NewLoadBalancerFrontendIPConfigurationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewLoadBalancerFrontendIPConfigurationsClientWithBaseURI creates an instance of the
// LoadBalancerFrontendIPConfigurationsClient client using a custom endpoint.  Use this when interacting with an Azure
// cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewLoadBalancerFrontendIPConfigurationsClientWithBaseURI(baseURI string, subscriptionID string) LoadBalancerFrontendIPConfigurationsClient {
	return LoadBalancerFrontendIPConfigurationsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get gets load balancer frontend IP configuration.
// Parameters:
// resourceGroupName - the name of the resource group.
// loadBalancerName - the name of the load balancer.
// frontendIPConfigurationName - the name of the frontend IP configuration.
func (client LoadBalancerFrontendIPConfigurationsClient) Get(ctx context.Context, resourceGroupName string, loadBalancerName string, frontendIPConfigurationName string) (result FrontendIPConfiguration, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LoadBalancerFrontendIPConfigurationsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, loadBalancerName, frontendIPConfigurationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.LoadBalancerFrontendIPConfigurationsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.LoadBalancerFrontendIPConfigurationsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.LoadBalancerFrontendIPConfigurationsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client LoadBalancerFrontendIPConfigurationsClient) GetPreparer(ctx context.Context, resourceGroupName string, loadBalancerName string, frontendIPConfigurationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"frontendIPConfigurationName": autorest.Encode("path", frontendIPConfigurationName),
		"loadBalancerName":            autorest.Encode("path", loadBalancerName),
		"resourceGroupName":           autorest.Encode("path", resourceGroupName),
		"subscriptionId":              autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/frontendIPConfigurations/{frontendIPConfigurationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client LoadBalancerFrontendIPConfigurationsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client LoadBalancerFrontendIPConfigurationsClient) GetResponder(resp *http.Response) (result FrontendIPConfiguration, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets all the load balancer frontend IP configurations.
// Parameters:
// resourceGroupName - the name of the resource group.
// loadBalancerName - the name of the load balancer.
func (client LoadBalancerFrontendIPConfigurationsClient) List(ctx context.Context, resourceGroupName string, loadBalancerName string) (result LoadBalancerFrontendIPConfigurationListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LoadBalancerFrontendIPConfigurationsClient.List")
		defer func() {
			sc := -1
			if result.lbficlr.Response.Response != nil {
				sc = result.lbficlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, loadBalancerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.LoadBalancerFrontendIPConfigurationsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.lbficlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.LoadBalancerFrontendIPConfigurationsClient", "List", resp, "Failure sending request")
		return
	}

	result.lbficlr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.LoadBalancerFrontendIPConfigurationsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client LoadBalancerFrontendIPConfigurationsClient) ListPreparer(ctx context.Context, resourceGroupName string, loadBalancerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"loadBalancerName":  autorest.Encode("path", loadBalancerName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/frontendIPConfigurations", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client LoadBalancerFrontendIPConfigurationsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client LoadBalancerFrontendIPConfigurationsClient) ListResponder(resp *http.Response) (result LoadBalancerFrontendIPConfigurationListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client LoadBalancerFrontendIPConfigurationsClient) listNextResults(ctx context.Context, lastResults LoadBalancerFrontendIPConfigurationListResult) (result LoadBalancerFrontendIPConfigurationListResult, err error) {
	req, err := lastResults.loadBalancerFrontendIPConfigurationListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.LoadBalancerFrontendIPConfigurationsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.LoadBalancerFrontendIPConfigurationsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.LoadBalancerFrontendIPConfigurationsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client LoadBalancerFrontendIPConfigurationsClient) ListComplete(ctx context.Context, resourceGroupName string, loadBalancerName string) (result LoadBalancerFrontendIPConfigurationListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LoadBalancerFrontendIPConfigurationsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, loadBalancerName)
	return
}
