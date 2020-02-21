package reservations

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

// QuotaClient is the client for the Quota methods of the Reservations service.
type QuotaClient struct {
	BaseClient
}

// NewQuotaClient creates an instance of the QuotaClient client.
func NewQuotaClient() QuotaClient {
	return NewQuotaClientWithBaseURI(DefaultBaseURI)
}

// NewQuotaClientWithBaseURI creates an instance of the QuotaClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewQuotaClientWithBaseURI(baseURI string) QuotaClient {
	return QuotaClient{NewWithBaseURI(baseURI)}
}

// ListStatus this API gets the current quota limit and usages for the specific resource for resource provider for the
// specified location. This response can be used to submit quotaRequests.
// Parameters:
// subscriptionID - azure subscription id.
// providerID - azure resource Provider id.
// location - azure region.
// resourceName - the Resource name for the specific resource provider, such as SKU name for Microsoft.Compute,
// pool for Microsoft.Batch.
func (client QuotaClient) ListStatus(ctx context.Context, subscriptionID string, providerID string, location string, resourceName string) (result CurrentQuotaLimitBase, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/QuotaClient.ListStatus")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListStatusPreparer(ctx, subscriptionID, providerID, location, resourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservations.QuotaClient", "ListStatus", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListStatusSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "reservations.QuotaClient", "ListStatus", resp, "Failure sending request")
		return
	}

	result, err = client.ListStatusResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservations.QuotaClient", "ListStatus", resp, "Failure responding to request")
	}

	return
}

// ListStatusPreparer prepares the ListStatus request.
func (client QuotaClient) ListStatusPreparer(ctx context.Context, subscriptionID string, providerID string, location string, resourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       autorest.Encode("path", location),
		"providerId":     autorest.Encode("path", providerID),
		"resourceName":   autorest.Encode("path", resourceName),
		"subscriptionId": autorest.Encode("path", subscriptionID),
	}

	const APIVersion = "2019-07-19-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Capacity/resourceProviders/{providerId}/locations/{location}/serviceLimits/{resourceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListStatusSender sends the ListStatus request. The method will close the
// http.Response Body if it receives an error.
func (client QuotaClient) ListStatusSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListStatusResponder handles the response to the ListStatus request. The method always
// closes the http.Response Body.
func (client QuotaClient) ListStatusResponder(resp *http.Response) (result CurrentQuotaLimitBase, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
