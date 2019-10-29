package netapp

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

// ResourceClient is the microsoft NetApp Azure Resource Provider specification
type ResourceClient struct {
	BaseClient
}

// NewResourceClient creates an instance of the ResourceClient client.
func NewResourceClient(subscriptionID string) ResourceClient {
	return NewResourceClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewResourceClientWithBaseURI creates an instance of the ResourceClient client.
func NewResourceClientWithBaseURI(baseURI string, subscriptionID string) ResourceClient {
	return ResourceClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CheckFilePathAvailability check if a file path is available.
// Parameters:
// body - file path availability request.
// location - the location
func (client ResourceClient) CheckFilePathAvailability(ctx context.Context, body ResourceNameAvailabilityRequest, location string) (result ResourceNameAvailability, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ResourceClient.CheckFilePathAvailability")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: body,
			Constraints: []validation.Constraint{{Target: "body.Name", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "body.ResourceGroup", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("netapp.ResourceClient", "CheckFilePathAvailability", err.Error())
	}

	req, err := client.CheckFilePathAvailabilityPreparer(ctx, body, location)
	if err != nil {
		err = autorest.NewErrorWithError(err, "netapp.ResourceClient", "CheckFilePathAvailability", nil, "Failure preparing request")
		return
	}

	resp, err := client.CheckFilePathAvailabilitySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "netapp.ResourceClient", "CheckFilePathAvailability", resp, "Failure sending request")
		return
	}

	result, err = client.CheckFilePathAvailabilityResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "netapp.ResourceClient", "CheckFilePathAvailability", resp, "Failure responding to request")
	}

	return
}

// CheckFilePathAvailabilityPreparer prepares the CheckFilePathAvailability request.
func (client ResourceClient) CheckFilePathAvailabilityPreparer(ctx context.Context, body ResourceNameAvailabilityRequest, location string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       autorest.Encode("path", location),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.NetApp/locations/{location}/checkFilePathAvailability", pathParameters),
		autorest.WithJSON(body),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CheckFilePathAvailabilitySender sends the CheckFilePathAvailability request. The method will close the
// http.Response Body if it receives an error.
func (client ResourceClient) CheckFilePathAvailabilitySender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// CheckFilePathAvailabilityResponder handles the response to the CheckFilePathAvailability request. The method always
// closes the http.Response Body.
func (client ResourceClient) CheckFilePathAvailabilityResponder(resp *http.Response) (result ResourceNameAvailability, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CheckNameAvailability check if a resource name is available.
// Parameters:
// body - name availability request.
// location - the location
func (client ResourceClient) CheckNameAvailability(ctx context.Context, body ResourceNameAvailabilityRequest, location string) (result ResourceNameAvailability, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ResourceClient.CheckNameAvailability")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: body,
			Constraints: []validation.Constraint{{Target: "body.Name", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "body.ResourceGroup", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("netapp.ResourceClient", "CheckNameAvailability", err.Error())
	}

	req, err := client.CheckNameAvailabilityPreparer(ctx, body, location)
	if err != nil {
		err = autorest.NewErrorWithError(err, "netapp.ResourceClient", "CheckNameAvailability", nil, "Failure preparing request")
		return
	}

	resp, err := client.CheckNameAvailabilitySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "netapp.ResourceClient", "CheckNameAvailability", resp, "Failure sending request")
		return
	}

	result, err = client.CheckNameAvailabilityResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "netapp.ResourceClient", "CheckNameAvailability", resp, "Failure responding to request")
	}

	return
}

// CheckNameAvailabilityPreparer prepares the CheckNameAvailability request.
func (client ResourceClient) CheckNameAvailabilityPreparer(ctx context.Context, body ResourceNameAvailabilityRequest, location string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       autorest.Encode("path", location),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.NetApp/locations/{location}/checkNameAvailability", pathParameters),
		autorest.WithJSON(body),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CheckNameAvailabilitySender sends the CheckNameAvailability request. The method will close the
// http.Response Body if it receives an error.
func (client ResourceClient) CheckNameAvailabilitySender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// CheckNameAvailabilityResponder handles the response to the CheckNameAvailability request. The method always
// closes the http.Response Body.
func (client ResourceClient) CheckNameAvailabilityResponder(resp *http.Response) (result ResourceNameAvailability, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
