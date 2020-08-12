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

// VpnSitesClient is the network Client
type VpnSitesClient struct {
	BaseClient
}

// NewVpnSitesClient creates an instance of the VpnSitesClient client.
func NewVpnSitesClient(subscriptionID string) VpnSitesClient {
	return NewVpnSitesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewVpnSitesClientWithBaseURI creates an instance of the VpnSitesClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewVpnSitesClientWithBaseURI(baseURI string, subscriptionID string) VpnSitesClient {
	return VpnSitesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates a VpnSite resource if it doesn't exist else updates the existing VpnSite.
// Parameters:
// resourceGroupName - the resource group name of the VpnSite.
// vpnSiteName - the name of the VpnSite being created or updated.
// vpnSiteParameters - parameters supplied to create or update VpnSite.
func (client VpnSitesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, vpnSiteName string, vpnSiteParameters VpnSite) (result VpnSitesCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VpnSitesClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, vpnSiteName, vpnSiteParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VpnSitesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VpnSitesClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client VpnSitesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, vpnSiteName string, vpnSiteParameters VpnSite) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vpnSiteName":       autorest.Encode("path", vpnSiteName),
	}

	const APIVersion = "2018-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	vpnSiteParameters.Etag = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnSites/{vpnSiteName}", pathParameters),
		autorest.WithJSON(vpnSiteParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client VpnSitesClient) CreateOrUpdateSender(req *http.Request) (future VpnSitesCreateOrUpdateFuture, err error) {
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
func (client VpnSitesClient) CreateOrUpdateResponder(resp *http.Response) (result VpnSite, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a VpnSite.
// Parameters:
// resourceGroupName - the resource group name of the VpnSite.
// vpnSiteName - the name of the VpnSite being deleted.
func (client VpnSitesClient) Delete(ctx context.Context, resourceGroupName string, vpnSiteName string) (result VpnSitesDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VpnSitesClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, vpnSiteName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VpnSitesClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VpnSitesClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client VpnSitesClient) DeletePreparer(ctx context.Context, resourceGroupName string, vpnSiteName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vpnSiteName":       autorest.Encode("path", vpnSiteName),
	}

	const APIVersion = "2018-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnSites/{vpnSiteName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client VpnSitesClient) DeleteSender(req *http.Request) (future VpnSitesDeleteFuture, err error) {
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
func (client VpnSitesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get retrieves the details of a VPN site.
// Parameters:
// resourceGroupName - the resource group name of the VpnSite.
// vpnSiteName - the name of the VpnSite being retrieved.
func (client VpnSitesClient) Get(ctx context.Context, resourceGroupName string, vpnSiteName string) (result VpnSite, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VpnSitesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, vpnSiteName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VpnSitesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.VpnSitesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VpnSitesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client VpnSitesClient) GetPreparer(ctx context.Context, resourceGroupName string, vpnSiteName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vpnSiteName":       autorest.Encode("path", vpnSiteName),
	}

	const APIVersion = "2018-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnSites/{vpnSiteName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client VpnSitesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client VpnSitesClient) GetResponder(resp *http.Response) (result VpnSite, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists all the VpnSites in a subscription.
func (client VpnSitesClient) List(ctx context.Context) (result ListVpnSitesResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VpnSitesClient.List")
		defer func() {
			sc := -1
			if result.lvsr.Response.Response != nil {
				sc = result.lvsr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VpnSitesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.lvsr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.VpnSitesClient", "List", resp, "Failure sending request")
		return
	}

	result.lvsr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VpnSitesClient", "List", resp, "Failure responding to request")
	}
	if result.lvsr.hasNextLink() && result.lvsr.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListPreparer prepares the List request.
func (client VpnSitesClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Network/vpnSites", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client VpnSitesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client VpnSitesClient) ListResponder(resp *http.Response) (result ListVpnSitesResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client VpnSitesClient) listNextResults(ctx context.Context, lastResults ListVpnSitesResult) (result ListVpnSitesResult, err error) {
	req, err := lastResults.listVpnSitesResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.VpnSitesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.VpnSitesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VpnSitesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client VpnSitesClient) ListComplete(ctx context.Context) (result ListVpnSitesResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VpnSitesClient.List")
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

// ListByResourceGroup lists all the vpnSites in a resource group.
// Parameters:
// resourceGroupName - the resource group name of the VpnSite.
func (client VpnSitesClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (result ListVpnSitesResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VpnSitesClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.lvsr.Response.Response != nil {
				sc = result.lvsr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VpnSitesClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.lvsr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.VpnSitesClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.lvsr, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VpnSitesClient", "ListByResourceGroup", resp, "Failure responding to request")
	}
	if result.lvsr.hasNextLink() && result.lvsr.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client VpnSitesClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnSites", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client VpnSitesClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client VpnSitesClient) ListByResourceGroupResponder(resp *http.Response) (result ListVpnSitesResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client VpnSitesClient) listByResourceGroupNextResults(ctx context.Context, lastResults ListVpnSitesResult) (result ListVpnSitesResult, err error) {
	req, err := lastResults.listVpnSitesResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.VpnSitesClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.VpnSitesClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VpnSitesClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client VpnSitesClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result ListVpnSitesResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VpnSitesClient.ListByResourceGroup")
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

// UpdateTags updates VpnSite tags.
// Parameters:
// resourceGroupName - the resource group name of the VpnSite.
// vpnSiteName - the name of the VpnSite being updated.
// vpnSiteParameters - parameters supplied to update VpnSite tags.
func (client VpnSitesClient) UpdateTags(ctx context.Context, resourceGroupName string, vpnSiteName string, vpnSiteParameters TagsObject) (result VpnSitesUpdateTagsFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VpnSitesClient.UpdateTags")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdateTagsPreparer(ctx, resourceGroupName, vpnSiteName, vpnSiteParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VpnSitesClient", "UpdateTags", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateTagsSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VpnSitesClient", "UpdateTags", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdateTagsPreparer prepares the UpdateTags request.
func (client VpnSitesClient) UpdateTagsPreparer(ctx context.Context, resourceGroupName string, vpnSiteName string, vpnSiteParameters TagsObject) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vpnSiteName":       autorest.Encode("path", vpnSiteName),
	}

	const APIVersion = "2018-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnSites/{vpnSiteName}", pathParameters),
		autorest.WithJSON(vpnSiteParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateTagsSender sends the UpdateTags request. The method will close the
// http.Response Body if it receives an error.
func (client VpnSitesClient) UpdateTagsSender(req *http.Request) (future VpnSitesUpdateTagsFuture, err error) {
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
func (client VpnSitesClient) UpdateTagsResponder(resp *http.Response) (result VpnSite, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
