// Package postgresql implements the Azure ARM Postgresql service API version 2020-02-14-privatepreview.
//
// The Microsoft Azure management API provides create, read, update, and delete functionality for Azure PostgreSQL
// resources including servers, databases, firewall rules, VNET rules, security alert policies, log files and
// configurations with new business model.
package postgresql

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

const (
	// DefaultBaseURI is the default URI used for the service Postgresql
	DefaultBaseURI = "https://management.azure.com"
)

// BaseClient is the base client for Postgresql.
type BaseClient struct {
	autorest.Client
	BaseURI        string
	SubscriptionID string
}

// New creates an instance of the BaseClient client.
func New(subscriptionID string) BaseClient {
	return NewWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWithBaseURI creates an instance of the BaseClient client using a custom endpoint.  Use this when interacting with
// an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return BaseClient{
		Client:         autorest.NewClientWithUserAgent(UserAgent()),
		BaseURI:        baseURI,
		SubscriptionID: subscriptionID,
	}
}

// LocationBasedCapabilities get capabilities at specified location in a given subscription.
// Parameters:
// locationName - the name of the location.
func (client BaseClient) LocationBasedCapabilities(ctx context.Context, locationName string) (result CapabilitiesListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.LocationBasedCapabilities")
		defer func() {
			sc := -1
			if result.clr.Response.Response != nil {
				sc = result.clr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("postgresql.BaseClient", "LocationBasedCapabilities", err.Error())
	}

	result.fn = client.locationBasedCapabilitiesNextResults
	req, err := client.LocationBasedCapabilitiesPreparer(ctx, locationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "postgresql.BaseClient", "LocationBasedCapabilities", nil, "Failure preparing request")
		return
	}

	resp, err := client.LocationBasedCapabilitiesSender(req)
	if err != nil {
		result.clr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "postgresql.BaseClient", "LocationBasedCapabilities", resp, "Failure sending request")
		return
	}

	result.clr, err = client.LocationBasedCapabilitiesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "postgresql.BaseClient", "LocationBasedCapabilities", resp, "Failure responding to request")
	}
	if result.clr.hasNextLink() && result.clr.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// LocationBasedCapabilitiesPreparer prepares the LocationBasedCapabilities request.
func (client BaseClient) LocationBasedCapabilitiesPreparer(ctx context.Context, locationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"locationName":   autorest.Encode("path", locationName),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-02-14-privatepreview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.DBForPostgreSql/locations/{locationName}/capabilities", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// LocationBasedCapabilitiesSender sends the LocationBasedCapabilities request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) LocationBasedCapabilitiesSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// LocationBasedCapabilitiesResponder handles the response to the LocationBasedCapabilities request. The method always
// closes the http.Response Body.
func (client BaseClient) LocationBasedCapabilitiesResponder(resp *http.Response) (result CapabilitiesListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// locationBasedCapabilitiesNextResults retrieves the next set of results, if any.
func (client BaseClient) locationBasedCapabilitiesNextResults(ctx context.Context, lastResults CapabilitiesListResult) (result CapabilitiesListResult, err error) {
	req, err := lastResults.capabilitiesListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "postgresql.BaseClient", "locationBasedCapabilitiesNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.LocationBasedCapabilitiesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "postgresql.BaseClient", "locationBasedCapabilitiesNextResults", resp, "Failure sending next results request")
	}
	result, err = client.LocationBasedCapabilitiesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "postgresql.BaseClient", "locationBasedCapabilitiesNextResults", resp, "Failure responding to next results request")
	}
	return
}

// LocationBasedCapabilitiesComplete enumerates all values, automatically crossing page boundaries as required.
func (client BaseClient) LocationBasedCapabilitiesComplete(ctx context.Context, locationName string) (result CapabilitiesListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.LocationBasedCapabilities")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.LocationBasedCapabilities(ctx, locationName)
	return
}
