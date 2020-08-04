package adhybridhealthservice

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

// ListClient is the REST APIs for Azure Active Directory Connect Health
type ListClient struct {
	BaseClient
}

// NewListClient creates an instance of the ListClient client.
func NewListClient() ListClient {
	return NewListClientWithBaseURI(DefaultBaseURI)
}

// NewListClientWithBaseURI creates an instance of the ListClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewListClientWithBaseURI(baseURI string) ListClient {
	return ListClient{NewWithBaseURI(baseURI)}
}

// IPAddressAggregatesByService gets the IP address aggregates for a given service.
// Parameters:
// serviceName - the name of the service.
// skiptoken - a continuationtoken value returned in paginated result to load different pages.
func (client ListClient) IPAddressAggregatesByService(ctx context.Context, serviceName string, skiptoken string) (result IPAddressAggregatesPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ListClient.IPAddressAggregatesByService")
		defer func() {
			sc := -1
			if result.iaa.Response.Response != nil {
				sc = result.iaa.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.iPAddressAggregatesByServiceNextResults
	req, err := client.IPAddressAggregatesByServicePreparer(ctx, serviceName, skiptoken)
	if err != nil {
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.ListClient", "IPAddressAggregatesByService", nil, "Failure preparing request")
		return
	}

	resp, err := client.IPAddressAggregatesByServiceSender(req)
	if err != nil {
		result.iaa.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.ListClient", "IPAddressAggregatesByService", resp, "Failure sending request")
		return
	}

	result.iaa, err = client.IPAddressAggregatesByServiceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.ListClient", "IPAddressAggregatesByService", resp, "Failure responding to request")
	}
	if result.iaa.hasNextLink() && result.iaa.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// IPAddressAggregatesByServicePreparer prepares the IPAddressAggregatesByService request.
func (client ListClient) IPAddressAggregatesByServicePreparer(ctx context.Context, serviceName string, skiptoken string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"serviceName": autorest.Encode("path", serviceName),
	}

	const APIVersion = "2014-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(skiptoken) > 0 {
		queryParameters["skiptoken"] = autorest.Encode("query", skiptoken)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/ipAddressAggregates", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// IPAddressAggregatesByServiceSender sends the IPAddressAggregatesByService request. The method will close the
// http.Response Body if it receives an error.
func (client ListClient) IPAddressAggregatesByServiceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// IPAddressAggregatesByServiceResponder handles the response to the IPAddressAggregatesByService request. The method always
// closes the http.Response Body.
func (client ListClient) IPAddressAggregatesByServiceResponder(resp *http.Response) (result IPAddressAggregates, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// iPAddressAggregatesByServiceNextResults retrieves the next set of results, if any.
func (client ListClient) iPAddressAggregatesByServiceNextResults(ctx context.Context, lastResults IPAddressAggregates) (result IPAddressAggregates, err error) {
	req, err := lastResults.iPAddressAggregatesPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "adhybridhealthservice.ListClient", "iPAddressAggregatesByServiceNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.IPAddressAggregatesByServiceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "adhybridhealthservice.ListClient", "iPAddressAggregatesByServiceNextResults", resp, "Failure sending next results request")
	}
	result, err = client.IPAddressAggregatesByServiceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.ListClient", "iPAddressAggregatesByServiceNextResults", resp, "Failure responding to next results request")
	}
	return
}

// IPAddressAggregatesByServiceComplete enumerates all values, automatically crossing page boundaries as required.
func (client ListClient) IPAddressAggregatesByServiceComplete(ctx context.Context, serviceName string, skiptoken string) (result IPAddressAggregatesIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ListClient.IPAddressAggregatesByService")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.IPAddressAggregatesByService(ctx, serviceName, skiptoken)
	return
}

// IPAddressAggregateSettings gets the IP address aggregate settings.
// Parameters:
// serviceName - the name of the service.
func (client ListClient) IPAddressAggregateSettings(ctx context.Context, serviceName string) (result IPAddressAggregateSetting, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ListClient.IPAddressAggregateSettings")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.IPAddressAggregateSettingsPreparer(ctx, serviceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.ListClient", "IPAddressAggregateSettings", nil, "Failure preparing request")
		return
	}

	resp, err := client.IPAddressAggregateSettingsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.ListClient", "IPAddressAggregateSettings", resp, "Failure sending request")
		return
	}

	result, err = client.IPAddressAggregateSettingsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.ListClient", "IPAddressAggregateSettings", resp, "Failure responding to request")
	}

	return
}

// IPAddressAggregateSettingsPreparer prepares the IPAddressAggregateSettings request.
func (client ListClient) IPAddressAggregateSettingsPreparer(ctx context.Context, serviceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"serviceName": autorest.Encode("path", serviceName),
	}

	const APIVersion = "2014-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/ipAddressAggregateSettings", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// IPAddressAggregateSettingsSender sends the IPAddressAggregateSettings request. The method will close the
// http.Response Body if it receives an error.
func (client ListClient) IPAddressAggregateSettingsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// IPAddressAggregateSettingsResponder handles the response to the IPAddressAggregateSettings request. The method always
// closes the http.Response Body.
func (client ListClient) IPAddressAggregateSettingsResponder(resp *http.Response) (result IPAddressAggregateSetting, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}