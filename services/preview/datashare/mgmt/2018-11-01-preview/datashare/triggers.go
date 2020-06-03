package datashare

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

// TriggersClient is the creates a Microsoft.DataShare management client.
type TriggersClient struct {
	BaseClient
}

// NewTriggersClient creates an instance of the TriggersClient client.
func NewTriggersClient(subscriptionID string) TriggersClient {
	return NewTriggersClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewTriggersClientWithBaseURI creates an instance of the TriggersClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewTriggersClientWithBaseURI(baseURI string, subscriptionID string) TriggersClient {
	return TriggersClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create create a Trigger
// Parameters:
// resourceGroupName - the resource group name.
// accountName - the name of the share account.
// shareSubscriptionName - the name of the share subscription which will hold the data set sink.
// triggerName - the name of the trigger.
// trigger - trigger details.
func (client TriggersClient) Create(ctx context.Context, resourceGroupName string, accountName string, shareSubscriptionName string, triggerName string, trigger BasicTrigger) (result TriggersCreateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TriggersClient.Create")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreatePreparer(ctx, resourceGroupName, accountName, shareSubscriptionName, triggerName, trigger)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datashare.TriggersClient", "Create", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datashare.TriggersClient", "Create", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client TriggersClient) CreatePreparer(ctx context.Context, resourceGroupName string, accountName string, shareSubscriptionName string, triggerName string, trigger BasicTrigger) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":           autorest.Encode("path", accountName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"shareSubscriptionName": autorest.Encode("path", shareSubscriptionName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
		"triggerName":           autorest.Encode("path", triggerName),
	}

	const APIVersion = "2018-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shareSubscriptions/{shareSubscriptionName}/triggers/{triggerName}", pathParameters),
		autorest.WithJSON(trigger),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client TriggersClient) CreateSender(req *http.Request) (future TriggersCreateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client TriggersClient) CreateResponder(resp *http.Response) (result TriggerModel, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete a Trigger in a shareSubscription
// Parameters:
// resourceGroupName - the resource group name.
// accountName - the name of the share account.
// shareSubscriptionName - the name of the shareSubscription.
// triggerName - the name of the trigger.
func (client TriggersClient) Delete(ctx context.Context, resourceGroupName string, accountName string, shareSubscriptionName string, triggerName string) (result TriggersDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TriggersClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, accountName, shareSubscriptionName, triggerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datashare.TriggersClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datashare.TriggersClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client TriggersClient) DeletePreparer(ctx context.Context, resourceGroupName string, accountName string, shareSubscriptionName string, triggerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":           autorest.Encode("path", accountName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"shareSubscriptionName": autorest.Encode("path", shareSubscriptionName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
		"triggerName":           autorest.Encode("path", triggerName),
	}

	const APIVersion = "2018-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shareSubscriptions/{shareSubscriptionName}/triggers/{triggerName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client TriggersClient) DeleteSender(req *http.Request) (future TriggersDeleteFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client TriggersClient) DeleteResponder(resp *http.Response) (result OperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get get a Trigger in a shareSubscription
// Parameters:
// resourceGroupName - the resource group name.
// accountName - the name of the share account.
// shareSubscriptionName - the name of the shareSubscription.
// triggerName - the name of the trigger.
func (client TriggersClient) Get(ctx context.Context, resourceGroupName string, accountName string, shareSubscriptionName string, triggerName string) (result TriggerModel, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TriggersClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, accountName, shareSubscriptionName, triggerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datashare.TriggersClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datashare.TriggersClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datashare.TriggersClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client TriggersClient) GetPreparer(ctx context.Context, resourceGroupName string, accountName string, shareSubscriptionName string, triggerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":           autorest.Encode("path", accountName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"shareSubscriptionName": autorest.Encode("path", shareSubscriptionName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
		"triggerName":           autorest.Encode("path", triggerName),
	}

	const APIVersion = "2018-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shareSubscriptions/{shareSubscriptionName}/triggers/{triggerName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client TriggersClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client TriggersClient) GetResponder(resp *http.Response) (result TriggerModel, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByShareSubscription list Triggers in a share subscription
// Parameters:
// resourceGroupName - the resource group name.
// accountName - the name of the share account.
// shareSubscriptionName - the name of the share subscription.
// skipToken - continuation token
func (client TriggersClient) ListByShareSubscription(ctx context.Context, resourceGroupName string, accountName string, shareSubscriptionName string, skipToken string) (result TriggerListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TriggersClient.ListByShareSubscription")
		defer func() {
			sc := -1
			if result.tl.Response.Response != nil {
				sc = result.tl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByShareSubscriptionNextResults
	req, err := client.ListByShareSubscriptionPreparer(ctx, resourceGroupName, accountName, shareSubscriptionName, skipToken)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datashare.TriggersClient", "ListByShareSubscription", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByShareSubscriptionSender(req)
	if err != nil {
		result.tl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datashare.TriggersClient", "ListByShareSubscription", resp, "Failure sending request")
		return
	}

	result.tl, err = client.ListByShareSubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datashare.TriggersClient", "ListByShareSubscription", resp, "Failure responding to request")
	}

	return
}

// ListByShareSubscriptionPreparer prepares the ListByShareSubscription request.
func (client TriggersClient) ListByShareSubscriptionPreparer(ctx context.Context, resourceGroupName string, accountName string, shareSubscriptionName string, skipToken string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":           autorest.Encode("path", accountName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"shareSubscriptionName": autorest.Encode("path", shareSubscriptionName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(skipToken) > 0 {
		queryParameters["$skipToken"] = autorest.Encode("query", skipToken)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shareSubscriptions/{shareSubscriptionName}/triggers", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByShareSubscriptionSender sends the ListByShareSubscription request. The method will close the
// http.Response Body if it receives an error.
func (client TriggersClient) ListByShareSubscriptionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByShareSubscriptionResponder handles the response to the ListByShareSubscription request. The method always
// closes the http.Response Body.
func (client TriggersClient) ListByShareSubscriptionResponder(resp *http.Response) (result TriggerList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByShareSubscriptionNextResults retrieves the next set of results, if any.
func (client TriggersClient) listByShareSubscriptionNextResults(ctx context.Context, lastResults TriggerList) (result TriggerList, err error) {
	req, err := lastResults.triggerListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "datashare.TriggersClient", "listByShareSubscriptionNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByShareSubscriptionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "datashare.TriggersClient", "listByShareSubscriptionNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByShareSubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datashare.TriggersClient", "listByShareSubscriptionNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByShareSubscriptionComplete enumerates all values, automatically crossing page boundaries as required.
func (client TriggersClient) ListByShareSubscriptionComplete(ctx context.Context, resourceGroupName string, accountName string, shareSubscriptionName string, skipToken string) (result TriggerListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TriggersClient.ListByShareSubscription")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByShareSubscription(ctx, resourceGroupName, accountName, shareSubscriptionName, skipToken)
	return
}
