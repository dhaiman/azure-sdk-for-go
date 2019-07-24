package dns

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

// ResourceReferenceClient is the the DNS Management Client.
type ResourceReferenceClient struct {
	BaseClient
}

// NewResourceReferenceClient creates an instance of the ResourceReferenceClient client.
func NewResourceReferenceClient(subscriptionID string) ResourceReferenceClient {
	return NewResourceReferenceClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewResourceReferenceClientWithBaseURI creates an instance of the ResourceReferenceClient client.
func NewResourceReferenceClientWithBaseURI(baseURI string, subscriptionID string) ResourceReferenceClient {
	return ResourceReferenceClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// GetByTargetResources returns the DNS records specified by the referencing targetResourceIds.
// Parameters:
// parameters - properties for dns resource reference request.
func (client ResourceReferenceClient) GetByTargetResources(ctx context.Context, parameters ResourceReferenceRequest) (result ResourceReferenceResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ResourceReferenceClient.GetByTargetResources")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetByTargetResourcesPreparer(ctx, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dns.ResourceReferenceClient", "GetByTargetResources", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetByTargetResourcesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "dns.ResourceReferenceClient", "GetByTargetResources", resp, "Failure sending request")
		return
	}

	result, err = client.GetByTargetResourcesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dns.ResourceReferenceClient", "GetByTargetResources", resp, "Failure responding to request")
	}

	return
}

// GetByTargetResourcesPreparer prepares the GetByTargetResources request.
func (client ResourceReferenceClient) GetByTargetResourcesPreparer(ctx context.Context, parameters ResourceReferenceRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Network/getDnsResourceReference", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetByTargetResourcesSender sends the GetByTargetResources request. The method will close the
// http.Response Body if it receives an error.
func (client ResourceReferenceClient) GetByTargetResourcesSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetByTargetResourcesResponder handles the response to the GetByTargetResources request. The method always
// closes the http.Response Body.
func (client ResourceReferenceClient) GetByTargetResourcesResponder(resp *http.Response) (result ResourceReferenceResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
