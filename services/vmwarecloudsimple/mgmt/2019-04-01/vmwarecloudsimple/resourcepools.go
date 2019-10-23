package vmwarecloudsimple

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

// ResourcePoolsClient is the description of the new service
type ResourcePoolsClient struct {
	BaseClient
}

// NewResourcePoolsClient creates an instance of the ResourcePoolsClient client.
func NewResourcePoolsClient(subscriptionID string, referer string) ResourcePoolsClient {
	return NewResourcePoolsClientWithBaseURI(DefaultBaseURI, subscriptionID, referer)
}

// NewResourcePoolsClientWithBaseURI creates an instance of the ResourcePoolsClient client.
func NewResourcePoolsClientWithBaseURI(baseURI string, subscriptionID string, referer string) ResourcePoolsClient {
	return ResourcePoolsClient{NewWithBaseURI(baseURI, subscriptionID, referer)}
}

// Get returns resource pool templates by its name
// Parameters:
// regionID - the region Id (westus, eastus)
// pcName - the private cloud name
// resourcePoolName - resource pool id (vsphereId)
func (client ResourcePoolsClient) Get(ctx context.Context, regionID string, pcName string, resourcePoolName string) (result ResourcePool, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ResourcePoolsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, regionID, pcName, resourcePoolName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.ResourcePoolsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.ResourcePoolsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.ResourcePoolsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ResourcePoolsClient) GetPreparer(ctx context.Context, regionID string, pcName string, resourcePoolName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"pcName":           autorest.Encode("path", pcName),
		"regionId":         autorest.Encode("path", regionID),
		"resourcePoolName": autorest.Encode("path", resourcePoolName),
		"subscriptionId":   autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/locations/{regionId}/privateClouds/{pcName}/resourcePools/{resourcePoolName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ResourcePoolsClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ResourcePoolsClient) GetResponder(resp *http.Response) (result ResourcePool, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List returns list of resource pools in region for private cloud
// Parameters:
// regionID - the region Id (westus, eastus)
// pcName - the private cloud name
func (client ResourcePoolsClient) List(ctx context.Context, regionID string, pcName string) (result ResourcePoolsListResponsePage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ResourcePoolsClient.List")
		defer func() {
			sc := -1
			if result.rplr.Response.Response != nil {
				sc = result.rplr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, regionID, pcName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.ResourcePoolsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.rplr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.ResourcePoolsClient", "List", resp, "Failure sending request")
		return
	}

	result.rplr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.ResourcePoolsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ResourcePoolsClient) ListPreparer(ctx context.Context, regionID string, pcName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"pcName":         autorest.Encode("path", pcName),
		"regionId":       autorest.Encode("path", regionID),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/locations/{regionId}/privateClouds/{pcName}/resourcePools", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ResourcePoolsClient) ListSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ResourcePoolsClient) ListResponder(resp *http.Response) (result ResourcePoolsListResponse, err error) {
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
func (client ResourcePoolsClient) listNextResults(ctx context.Context, lastResults ResourcePoolsListResponse) (result ResourcePoolsListResponse, err error) {
	req, err := lastResults.resourcePoolsListResponsePreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "vmwarecloudsimple.ResourcePoolsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "vmwarecloudsimple.ResourcePoolsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vmwarecloudsimple.ResourcePoolsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client ResourcePoolsClient) ListComplete(ctx context.Context, regionID string, pcName string) (result ResourcePoolsListResponseIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ResourcePoolsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, regionID, pcName)
	return
}
