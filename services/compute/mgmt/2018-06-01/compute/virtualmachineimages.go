package compute

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

// VirtualMachineImagesClient is the compute Client
type VirtualMachineImagesClient struct {
	BaseClient
}

// NewVirtualMachineImagesClient creates an instance of the VirtualMachineImagesClient client.
func NewVirtualMachineImagesClient(subscriptionID string) VirtualMachineImagesClient {
	return NewVirtualMachineImagesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewVirtualMachineImagesClientWithBaseURI creates an instance of the VirtualMachineImagesClient client.
func NewVirtualMachineImagesClientWithBaseURI(baseURI string, subscriptionID string) VirtualMachineImagesClient {
	return VirtualMachineImagesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get gets a virtual machine image.
// Parameters:
// location - the name of a supported Azure region.
// publisherName - a valid image publisher.
// offer - a valid image publisher offer.
// skus - a valid image SKU.
// version - a valid image SKU version.
func (client VirtualMachineImagesClient) Get(ctx context.Context, location string, publisherName string, offer string, skus string, version string) (result VirtualMachineImage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualMachineImagesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, location, publisherName, offer, skus, version)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineImagesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineImagesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineImagesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client VirtualMachineImagesClient) GetPreparer(ctx context.Context, location string, publisherName string, offer string, skus string, version string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       autorest.Encode("path", location),
		"offer":          autorest.Encode("path", offer),
		"publisherName":  autorest.Encode("path", publisherName),
		"skus":           autorest.Encode("path", skus),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
		"version":        autorest.Encode("path", version),
	}

	const APIVersion = "2018-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/publishers/{publisherName}/artifacttypes/vmimage/offers/{offer}/skus/{skus}/versions/{version}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachineImagesClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client VirtualMachineImagesClient) GetResponder(resp *http.Response) (result VirtualMachineImage, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets a list of all virtual machine image versions for the specified location, publisher, offer, and SKU.
// Parameters:
// location - the name of a supported Azure region.
// publisherName - a valid image publisher.
// offer - a valid image publisher offer.
// skus - a valid image SKU.
// filter - the filter to apply on the operation.
func (client VirtualMachineImagesClient) List(ctx context.Context, location string, publisherName string, offer string, skus string, filter string, top *int32, orderby string) (result ListVirtualMachineImageResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualMachineImagesClient.List")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListPreparer(ctx, location, publisherName, offer, skus, filter, top, orderby)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineImagesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineImagesClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineImagesClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client VirtualMachineImagesClient) ListPreparer(ctx context.Context, location string, publisherName string, offer string, skus string, filter string, top *int32, orderby string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       autorest.Encode("path", location),
		"offer":          autorest.Encode("path", offer),
		"publisherName":  autorest.Encode("path", publisherName),
		"skus":           autorest.Encode("path", skus),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(orderby) > 0 {
		queryParameters["$orderby"] = autorest.Encode("query", orderby)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/publishers/{publisherName}/artifacttypes/vmimage/offers/{offer}/skus/{skus}/versions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachineImagesClient) ListSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client VirtualMachineImagesClient) ListResponder(resp *http.Response) (result ListVirtualMachineImageResource, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListOffers gets a list of virtual machine image offers for the specified location and publisher.
// Parameters:
// location - the name of a supported Azure region.
// publisherName - a valid image publisher.
func (client VirtualMachineImagesClient) ListOffers(ctx context.Context, location string, publisherName string) (result ListVirtualMachineImageResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualMachineImagesClient.ListOffers")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListOffersPreparer(ctx, location, publisherName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineImagesClient", "ListOffers", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListOffersSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineImagesClient", "ListOffers", resp, "Failure sending request")
		return
	}

	result, err = client.ListOffersResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineImagesClient", "ListOffers", resp, "Failure responding to request")
	}

	return
}

// ListOffersPreparer prepares the ListOffers request.
func (client VirtualMachineImagesClient) ListOffersPreparer(ctx context.Context, location string, publisherName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       autorest.Encode("path", location),
		"publisherName":  autorest.Encode("path", publisherName),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/publishers/{publisherName}/artifacttypes/vmimage/offers", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListOffersSender sends the ListOffers request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachineImagesClient) ListOffersSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListOffersResponder handles the response to the ListOffers request. The method always
// closes the http.Response Body.
func (client VirtualMachineImagesClient) ListOffersResponder(resp *http.Response) (result ListVirtualMachineImageResource, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListPublishers gets a list of virtual machine image publishers for the specified Azure location.
// Parameters:
// location - the name of a supported Azure region.
func (client VirtualMachineImagesClient) ListPublishers(ctx context.Context, location string) (result ListVirtualMachineImageResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualMachineImagesClient.ListPublishers")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListPublishersPreparer(ctx, location)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineImagesClient", "ListPublishers", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListPublishersSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineImagesClient", "ListPublishers", resp, "Failure sending request")
		return
	}

	result, err = client.ListPublishersResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineImagesClient", "ListPublishers", resp, "Failure responding to request")
	}

	return
}

// ListPublishersPreparer prepares the ListPublishers request.
func (client VirtualMachineImagesClient) ListPublishersPreparer(ctx context.Context, location string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       autorest.Encode("path", location),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/publishers", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListPublishersSender sends the ListPublishers request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachineImagesClient) ListPublishersSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListPublishersResponder handles the response to the ListPublishers request. The method always
// closes the http.Response Body.
func (client VirtualMachineImagesClient) ListPublishersResponder(resp *http.Response) (result ListVirtualMachineImageResource, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListSkus gets a list of virtual machine image SKUs for the specified location, publisher, and offer.
// Parameters:
// location - the name of a supported Azure region.
// publisherName - a valid image publisher.
// offer - a valid image publisher offer.
func (client VirtualMachineImagesClient) ListSkus(ctx context.Context, location string, publisherName string, offer string) (result ListVirtualMachineImageResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualMachineImagesClient.ListSkus")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListSkusPreparer(ctx, location, publisherName, offer)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineImagesClient", "ListSkus", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSkusSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineImagesClient", "ListSkus", resp, "Failure sending request")
		return
	}

	result, err = client.ListSkusResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.VirtualMachineImagesClient", "ListSkus", resp, "Failure responding to request")
	}

	return
}

// ListSkusPreparer prepares the ListSkus request.
func (client VirtualMachineImagesClient) ListSkusPreparer(ctx context.Context, location string, publisherName string, offer string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       autorest.Encode("path", location),
		"offer":          autorest.Encode("path", offer),
		"publisherName":  autorest.Encode("path", publisherName),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/publishers/{publisherName}/artifacttypes/vmimage/offers/{offer}/skus", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSkusSender sends the ListSkus request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualMachineImagesClient) ListSkusSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListSkusResponder handles the response to the ListSkus request. The method always
// closes the http.Response Body.
func (client VirtualMachineImagesClient) ListSkusResponder(resp *http.Response) (result ListVirtualMachineImageResource, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
