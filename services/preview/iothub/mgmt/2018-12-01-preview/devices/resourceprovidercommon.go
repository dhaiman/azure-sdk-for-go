package devices

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

// ResourceProviderCommonClient is the use this API to manage the IoT hubs in your Azure subscription.
type ResourceProviderCommonClient struct {
	BaseClient
}

// NewResourceProviderCommonClient creates an instance of the ResourceProviderCommonClient client.
func NewResourceProviderCommonClient(subscriptionID string) ResourceProviderCommonClient {
	return NewResourceProviderCommonClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewResourceProviderCommonClientWithBaseURI creates an instance of the ResourceProviderCommonClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewResourceProviderCommonClientWithBaseURI(baseURI string, subscriptionID string) ResourceProviderCommonClient {
	return ResourceProviderCommonClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// GetSubscriptionQuota get the number of free and paid iot hubs in the subscription
func (client ResourceProviderCommonClient) GetSubscriptionQuota(ctx context.Context) (result UserSubscriptionQuotaListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ResourceProviderCommonClient.GetSubscriptionQuota")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetSubscriptionQuotaPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devices.ResourceProviderCommonClient", "GetSubscriptionQuota", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSubscriptionQuotaSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "devices.ResourceProviderCommonClient", "GetSubscriptionQuota", resp, "Failure sending request")
		return
	}

	result, err = client.GetSubscriptionQuotaResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devices.ResourceProviderCommonClient", "GetSubscriptionQuota", resp, "Failure responding to request")
	}

	return
}

// GetSubscriptionQuotaPreparer prepares the GetSubscriptionQuota request.
func (client ResourceProviderCommonClient) GetSubscriptionQuotaPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-12-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Devices/usages", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSubscriptionQuotaSender sends the GetSubscriptionQuota request. The method will close the
// http.Response Body if it receives an error.
func (client ResourceProviderCommonClient) GetSubscriptionQuotaSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetSubscriptionQuotaResponder handles the response to the GetSubscriptionQuota request. The method always
// closes the http.Response Body.
func (client ResourceProviderCommonClient) GetSubscriptionQuotaResponder(resp *http.Response) (result UserSubscriptionQuotaListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
