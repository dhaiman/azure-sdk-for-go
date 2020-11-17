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

// RestorableMongodbCollectionsClient is the client for the RestorableMongodbCollections methods of the Documentdb
// service.
type RestorableMongodbCollectionsClient struct {
	BaseClient
}

// NewRestorableMongodbCollectionsClient creates an instance of the RestorableMongodbCollectionsClient client.
func NewRestorableMongodbCollectionsClient(subscriptionID string) RestorableMongodbCollectionsClient {
	return NewRestorableMongodbCollectionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewRestorableMongodbCollectionsClientWithBaseURI creates an instance of the RestorableMongodbCollectionsClient
// client using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI
// (sovereign clouds, Azure stack).
func NewRestorableMongodbCollectionsClientWithBaseURI(baseURI string, subscriptionID string) RestorableMongodbCollectionsClient {
	return RestorableMongodbCollectionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get get the properties of a restorable MongoDB Collection.
// Parameters:
// location - cosmos DB region, with spaces between words and each word capitalized.
// instanceID - the instanceId GUID of a restorable database account.
// restorableMongodbCollectionID - the GUID of the restorable MongoDB collection.
func (client RestorableMongodbCollectionsClient) Get(ctx context.Context, location string, instanceID string, restorableMongodbCollectionID string) (result RestorableMongodbCollectionGetResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RestorableMongodbCollectionsClient.Get")
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
		return result, validation.NewError("documentdb.RestorableMongodbCollectionsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, location, instanceID, restorableMongodbCollectionID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "documentdb.RestorableMongodbCollectionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "documentdb.RestorableMongodbCollectionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "documentdb.RestorableMongodbCollectionsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client RestorableMongodbCollectionsClient) GetPreparer(ctx context.Context, location string, instanceID string, restorableMongodbCollectionID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"instanceId":                    autorest.Encode("path", instanceID),
		"location":                      autorest.Encode("path", location),
		"restorableMongodbCollectionId": autorest.Encode("path", restorableMongodbCollectionID),
		"subscriptionId":                autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.DocumentDB/locations/{location}/restorableDatabaseAccounts/{instanceId}/restorableMongodbCollections/{restorableMongodbCollectionId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client RestorableMongodbCollectionsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client RestorableMongodbCollectionsClient) GetResponder(resp *http.Response) (result RestorableMongodbCollectionGetResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists all the restorable Azure Cosmos DB MongoDB collection available for a specific database.
// Parameters:
// location - cosmos DB region, with spaces between words and each word capitalized.
// instanceID - the instanceId GUID of a restorable database account.
// restorableMongodbDatabaseRid - the resource id of the restorable Mongo database.
func (client RestorableMongodbCollectionsClient) List(ctx context.Context, location string, instanceID string, restorableMongodbDatabaseRid string) (result RestorableMongodbCollectionsListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RestorableMongodbCollectionsClient.List")
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
		return result, validation.NewError("documentdb.RestorableMongodbCollectionsClient", "List", err.Error())
	}

	req, err := client.ListPreparer(ctx, location, instanceID, restorableMongodbDatabaseRid)
	if err != nil {
		err = autorest.NewErrorWithError(err, "documentdb.RestorableMongodbCollectionsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "documentdb.RestorableMongodbCollectionsClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "documentdb.RestorableMongodbCollectionsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client RestorableMongodbCollectionsClient) ListPreparer(ctx context.Context, location string, instanceID string, restorableMongodbDatabaseRid string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"instanceId":     autorest.Encode("path", instanceID),
		"location":       autorest.Encode("path", location),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(restorableMongodbDatabaseRid) > 0 {
		queryParameters["restorableMongodbDatabaseRid"] = autorest.Encode("query", restorableMongodbDatabaseRid)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.DocumentDB/locations/{location}/restorableDatabaseAccounts/{instanceId}/restorableMongodbCollections", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client RestorableMongodbCollectionsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client RestorableMongodbCollectionsClient) ListResponder(resp *http.Response) (result RestorableMongodbCollectionsListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
