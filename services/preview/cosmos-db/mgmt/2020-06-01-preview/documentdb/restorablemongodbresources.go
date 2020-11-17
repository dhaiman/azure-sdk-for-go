package documentdb

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

// RestorableMongodbResourcesClient is the client for the RestorableMongodbResources methods of the Documentdb service.
type RestorableMongodbResourcesClient struct {
	BaseClient
}

// NewRestorableMongodbResourcesClient creates an instance of the RestorableMongodbResourcesClient client.
func NewRestorableMongodbResourcesClient(subscriptionID string) RestorableMongodbResourcesClient {
	return NewRestorableMongodbResourcesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewRestorableMongodbResourcesClientWithBaseURI creates an instance of the RestorableMongodbResourcesClient client
// using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
func NewRestorableMongodbResourcesClientWithBaseURI(baseURI string, subscriptionID string) RestorableMongodbResourcesClient {
	return RestorableMongodbResourcesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List lists all the restorable Azure Cosmos DB MongoDB resources available for a specific database account at a given
// time and location.
// Parameters:
// location - cosmos DB region, with spaces between words and each word capitalized.
// instanceID - the instanceId GUID of a restorable database account.
// restoreLocation - the location where the restorable resources are located.
// restoreTimestampInUtc - the timestamp when the restorable resources existed.
func (client RestorableMongodbResourcesClient) List(ctx context.Context, location string, instanceID string, restoreLocation string, restoreTimestampInUtc string) (result RestorableMongodbResourcesListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RestorableMongodbResourcesClient.List")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("documentdb.RestorableMongodbResourcesClient", "List", err.Error())
	}

	req, err := client.ListPreparer(ctx, location, instanceID, restoreLocation, restoreTimestampInUtc)
	if err != nil {
		err = autorest.NewErrorWithError(err, "documentdb.RestorableMongodbResourcesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "documentdb.RestorableMongodbResourcesClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "documentdb.RestorableMongodbResourcesClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client RestorableMongodbResourcesClient) ListPreparer(ctx context.Context, location string, instanceID string, restoreLocation string, restoreTimestampInUtc string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"instanceId":     autorest.Encode("path", instanceID),
		"location":       autorest.Encode("path", location),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(restoreLocation) > 0 {
		queryParameters["restoreLocation"] = autorest.Encode("query", restoreLocation)
	}
	if len(restoreTimestampInUtc) > 0 {
		queryParameters["restoreTimestampInUtc"] = autorest.Encode("query", restoreTimestampInUtc)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.DocumentDB/locations/{location}/restorableDatabaseAccounts/{instanceId}/restorableMongodbResources", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client RestorableMongodbResourcesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client RestorableMongodbResourcesClient) ListResponder(resp *http.Response) (result RestorableMongodbResourcesListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
