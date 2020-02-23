package security

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

// AdvancedThreatProtectionClient is the API spec for Microsoft.Security (Azure Security Center) resource provider
type AdvancedThreatProtectionClient struct {
	BaseClient
}

// NewAdvancedThreatProtectionClient creates an instance of the AdvancedThreatProtectionClient client.
func NewAdvancedThreatProtectionClient(subscriptionID string, ascLocation string) AdvancedThreatProtectionClient {
	return NewAdvancedThreatProtectionClientWithBaseURI(DefaultBaseURI, subscriptionID, ascLocation)
}

// NewAdvancedThreatProtectionClientWithBaseURI creates an instance of the AdvancedThreatProtectionClient client using
// a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
func NewAdvancedThreatProtectionClientWithBaseURI(baseURI string, subscriptionID string, ascLocation string) AdvancedThreatProtectionClient {
	return AdvancedThreatProtectionClient{NewWithBaseURI(baseURI, subscriptionID, ascLocation)}
}

// Create creates or updates the Advanced Threat Protection settings on a specified resource.
// Parameters:
// resourceID - the identifier of the resource.
// advancedThreatProtectionSetting - advanced Threat Protection Settings
func (client AdvancedThreatProtectionClient) Create(ctx context.Context, resourceID string, advancedThreatProtectionSetting AdvancedThreatProtectionSetting) (result AdvancedThreatProtectionSetting, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AdvancedThreatProtectionClient.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreatePreparer(ctx, resourceID, advancedThreatProtectionSetting)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.AdvancedThreatProtectionClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "security.AdvancedThreatProtectionClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.AdvancedThreatProtectionClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client AdvancedThreatProtectionClient) CreatePreparer(ctx context.Context, resourceID string, advancedThreatProtectionSetting AdvancedThreatProtectionSetting) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceId":  autorest.Encode("path", resourceID),
		"settingName": autorest.Encode("path", "current"),
	}

	const APIVersion = "2019-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{resourceId}/providers/Microsoft.Security/advancedThreatProtectionSettings/{settingName}", pathParameters),
		autorest.WithJSON(advancedThreatProtectionSetting),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client AdvancedThreatProtectionClient) CreateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client AdvancedThreatProtectionClient) CreateResponder(resp *http.Response) (result AdvancedThreatProtectionSetting, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get gets the Advanced Threat Protection settings for the specified resource.
// Parameters:
// resourceID - the identifier of the resource.
func (client AdvancedThreatProtectionClient) Get(ctx context.Context, resourceID string) (result AdvancedThreatProtectionSetting, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AdvancedThreatProtectionClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.AdvancedThreatProtectionClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "security.AdvancedThreatProtectionClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.AdvancedThreatProtectionClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client AdvancedThreatProtectionClient) GetPreparer(ctx context.Context, resourceID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceId":  autorest.Encode("path", resourceID),
		"settingName": autorest.Encode("path", "current"),
	}

	const APIVersion = "2019-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{resourceId}/providers/Microsoft.Security/advancedThreatProtectionSettings/{settingName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client AdvancedThreatProtectionClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client AdvancedThreatProtectionClient) GetResponder(resp *http.Response) (result AdvancedThreatProtectionSetting, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
