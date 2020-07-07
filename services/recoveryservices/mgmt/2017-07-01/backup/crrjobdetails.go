package backup

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

// CrrJobDetailsClient is the open API 2.0 Specs for Azure RecoveryServices Backup service
type CrrJobDetailsClient struct {
	BaseClient
}

// NewCrrJobDetailsClient creates an instance of the CrrJobDetailsClient client.
func NewCrrJobDetailsClient(subscriptionID string) CrrJobDetailsClient {
	return NewCrrJobDetailsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewCrrJobDetailsClientWithBaseURI creates an instance of the CrrJobDetailsClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewCrrJobDetailsClientWithBaseURI(baseURI string, subscriptionID string) CrrJobDetailsClient {
	return CrrJobDetailsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get sends the get request.
// Parameters:
// azureRegion - azure region to hit Api
func (client CrrJobDetailsClient) Get(ctx context.Context, azureRegion string) (result JobResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CrrJobDetailsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, azureRegion)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backup.CrrJobDetailsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "backup.CrrJobDetailsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backup.CrrJobDetailsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client CrrJobDetailsClient) GetPreparer(ctx context.Context, azureRegion string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"azureRegion":    autorest.Encode("path", azureRegion),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-12-20"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/providers/Microsoft.RecoveryServices/locations/{azureRegion}/backupCrrJob", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client CrrJobDetailsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client CrrJobDetailsClient) GetResponder(resp *http.Response) (result JobResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
