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

// ReplicationEventsClient is the client for the ReplicationEvents methods of the Siterecovery service.
type ReplicationEventsClient struct {
	BaseClient
}

// NewReplicationEventsClient creates an instance of the ReplicationEventsClient client.
func NewReplicationEventsClient(subscriptionID string, resourceGroupName string, resourceName string) ReplicationEventsClient {
	return NewReplicationEventsClientWithBaseURI(DefaultBaseURI, subscriptionID, resourceGroupName, resourceName)
}

// NewReplicationEventsClientWithBaseURI creates an instance of the ReplicationEventsClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewReplicationEventsClientWithBaseURI(baseURI string, subscriptionID string, resourceGroupName string, resourceName string) ReplicationEventsClient {
	return ReplicationEventsClient{NewWithBaseURI(baseURI, subscriptionID, resourceGroupName, resourceName)}
}

// Get the operation to get the details of an Azure Site recovery event.
// Parameters:
// eventName - the name of the Azure Site Recovery event.
func (client ReplicationEventsClient) Get(ctx context.Context, eventName string) (result Event, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReplicationEventsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, eventName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationEventsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationEventsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationEventsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ReplicationEventsClient) GetPreparer(ctx context.Context, eventName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"eventName":         autorest.Encode("path", eventName),
		"resourceGroupName": autorest.Encode("path", client.ResourceGroupName),
		"resourceName":      autorest.Encode("path", client.ResourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-07-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationEvents/{eventName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationEventsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ReplicationEventsClient) GetResponder(resp *http.Response) (result Event, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets the list of Azure Site Recovery events for the vault.
// Parameters:
// filter - oData filter options.
func (client ReplicationEventsClient) List(ctx context.Context, filter string) (result EventCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReplicationEventsClient.List")
		defer func() {
			sc := -1
			if result.ec.Response.Response != nil {
				sc = result.ec.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationEventsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.ec.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationEventsClient", "List", resp, "Failure sending request")
		return
	}

	result.ec, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationEventsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ReplicationEventsClient) ListPreparer(ctx context.Context, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", client.ResourceGroupName),
		"resourceName":      autorest.Encode("path", client.ResourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-07-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationEvents", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationEventsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ReplicationEventsClient) ListResponder(resp *http.Response) (result EventCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client ReplicationEventsClient) listNextResults(ctx context.Context, lastResults EventCollection) (result EventCollection, err error) {
	req, err := lastResults.eventCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "siterecovery.ReplicationEventsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "siterecovery.ReplicationEventsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.ReplicationEventsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client ReplicationEventsClient) ListComplete(ctx context.Context, filter string) (result EventCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReplicationEventsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, filter)
	return
}
