package regionmove

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

// MoveResourcesClient is the a first party Azure service orchestrating the move of Azure resources from one Azure
// region to another or between zones within a region.
type MoveResourcesClient struct {
	BaseClient
}

// NewMoveResourcesClient creates an instance of the MoveResourcesClient client.
func NewMoveResourcesClient(subscriptionID string) MoveResourcesClient {
	return NewMoveResourcesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewMoveResourcesClientWithBaseURI creates an instance of the MoveResourcesClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewMoveResourcesClientWithBaseURI(baseURI string, subscriptionID string) MoveResourcesClient {
	return MoveResourcesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create sends the create request.
func (client MoveResourcesClient) Create(ctx context.Context, resourceGroupName string, moveCollectionName string, moveResourceName string, body *MoveResource) (result MoveResourcesCreateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MoveResourcesClient.Create")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: body,
			Constraints: []validation.Constraint{{Target: "body", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "body.Properties", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "body.Properties.SourceID", Name: validation.Null, Rule: true, Chain: nil},
						{Target: "body.Properties.ResourceSettings", Name: validation.Null, Rule: false,
							Chain: []validation.Constraint{{Target: "body.Properties.ResourceSettings.TargetResourceName", Name: validation.Null, Rule: true, Chain: nil}}},
					}},
				}}}}}); err != nil {
		return result, validation.NewError("regionmove.MoveResourcesClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, resourceGroupName, moveCollectionName, moveResourceName, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "regionmove.MoveResourcesClient", "Create", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "regionmove.MoveResourcesClient", "Create", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client MoveResourcesClient) CreatePreparer(ctx context.Context, resourceGroupName string, moveCollectionName string, moveResourceName string, body *MoveResource) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"moveCollectionName": autorest.Encode("path", moveCollectionName),
		"moveResourceName":   autorest.Encode("path", moveResourceName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-10-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	body.ID = nil
	body.Name = nil
	body.Type = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/moveCollections/{moveCollectionName}/moveResources/{moveResourceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if body != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(body))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client MoveResourcesClient) CreateSender(req *http.Request) (future MoveResourcesCreateFuture, err error) {
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
func (client MoveResourcesClient) CreateResponder(resp *http.Response) (result MoveResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete sends the delete request.
func (client MoveResourcesClient) Delete(ctx context.Context, resourceGroupName string, moveCollectionName string, moveResourceName string) (result MoveResourcesDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MoveResourcesClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, moveCollectionName, moveResourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "regionmove.MoveResourcesClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "regionmove.MoveResourcesClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client MoveResourcesClient) DeletePreparer(ctx context.Context, resourceGroupName string, moveCollectionName string, moveResourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"moveCollectionName": autorest.Encode("path", moveCollectionName),
		"moveResourceName":   autorest.Encode("path", moveResourceName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-10-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/moveCollections/{moveCollectionName}/moveResources/{moveResourceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client MoveResourcesClient) DeleteSender(req *http.Request) (future MoveResourcesDeleteFuture, err error) {
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
func (client MoveResourcesClient) DeleteResponder(resp *http.Response) (result OperationStatus, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get sends the get request.
func (client MoveResourcesClient) Get(ctx context.Context, resourceGroupName string, moveCollectionName string, moveResourceName string) (result MoveResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MoveResourcesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, moveCollectionName, moveResourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "regionmove.MoveResourcesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "regionmove.MoveResourcesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "regionmove.MoveResourcesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client MoveResourcesClient) GetPreparer(ctx context.Context, resourceGroupName string, moveCollectionName string, moveResourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"moveCollectionName": autorest.Encode("path", moveCollectionName),
		"moveResourceName":   autorest.Encode("path", moveResourceName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-10-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/moveCollections/{moveCollectionName}/moveResources/{moveResourceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client MoveResourcesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client MoveResourcesClient) GetResponder(resp *http.Response) (result MoveResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List sends the list request.
// Parameters:
// filter - the filter to apply on the operation. For example, you can use $filter=Properties/ProvisioningState
// eq 'Succeeded'.
func (client MoveResourcesClient) List(ctx context.Context, resourceGroupName string, moveCollectionName string, filter string) (result MoveResourceCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MoveResourcesClient.List")
		defer func() {
			sc := -1
			if result.mrc.Response.Response != nil {
				sc = result.mrc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, moveCollectionName, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "regionmove.MoveResourcesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.mrc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "regionmove.MoveResourcesClient", "List", resp, "Failure sending request")
		return
	}

	result.mrc, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "regionmove.MoveResourcesClient", "List", resp, "Failure responding to request")
	}
	if result.mrc.hasNextLink() && result.mrc.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListPreparer prepares the List request.
func (client MoveResourcesClient) ListPreparer(ctx context.Context, resourceGroupName string, moveCollectionName string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"moveCollectionName": autorest.Encode("path", moveCollectionName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-10-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/moveCollections/{moveCollectionName}/moveResources", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client MoveResourcesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client MoveResourcesClient) ListResponder(resp *http.Response) (result MoveResourceCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client MoveResourcesClient) listNextResults(ctx context.Context, lastResults MoveResourceCollection) (result MoveResourceCollection, err error) {
	req, err := lastResults.moveResourceCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "regionmove.MoveResourcesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "regionmove.MoveResourcesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "regionmove.MoveResourcesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client MoveResourcesClient) ListComplete(ctx context.Context, resourceGroupName string, moveCollectionName string, filter string) (result MoveResourceCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MoveResourcesClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, moveCollectionName, filter)
	return
}
