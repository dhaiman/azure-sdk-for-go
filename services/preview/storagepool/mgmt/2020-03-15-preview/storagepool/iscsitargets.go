package storagepool

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

// IscsiTargetsClient is the client for the IscsiTargets methods of the Storagepool service.
type IscsiTargetsClient struct {
	BaseClient
}

// NewIscsiTargetsClient creates an instance of the IscsiTargetsClient client.
func NewIscsiTargetsClient(subscriptionID string) IscsiTargetsClient {
	return NewIscsiTargetsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewIscsiTargetsClientWithBaseURI creates an instance of the IscsiTargetsClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewIscsiTargetsClientWithBaseURI(baseURI string, subscriptionID string) IscsiTargetsClient {
	return IscsiTargetsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate create or Update an iSCSI target.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// diskPoolName - the name of the Disk Pool.
// iscsiTargetName - the name of the iSCSI target.
// iscsiTargetPayload - request payload for iSCSI target operations.
func (client IscsiTargetsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, diskPoolName string, iscsiTargetName string, iscsiTargetPayload IscsiTarget) (result IscsiTargetsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IscsiTargetsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]*[0-9A-Za-z]$`, Chain: nil}}},
		{TargetValue: diskPoolName,
			Constraints: []validation.Constraint{{Target: "diskPoolName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "diskPoolName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "diskPoolName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: iscsiTargetPayload,
			Constraints: []validation.Constraint{{Target: "iscsiTargetPayload.IscsiTargetProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "iscsiTargetPayload.IscsiTargetProperties.Tpgs", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "iscsiTargetPayload.IscsiTargetProperties.TargetIqn", Name: validation.Null, Rule: true, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("storagepool.IscsiTargetsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, diskPoolName, iscsiTargetName, iscsiTargetPayload)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagepool.IscsiTargetsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagepool.IscsiTargetsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client IscsiTargetsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, diskPoolName string, iscsiTargetName string, iscsiTargetPayload IscsiTarget) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"diskPoolName":      autorest.Encode("path", diskPoolName),
		"iscsiTargetName":   autorest.Encode("path", iscsiTargetName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-03-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StoragePool/diskPools/{diskPoolName}/iscsiTargets/{iscsiTargetName}", pathParameters),
		autorest.WithJSON(iscsiTargetPayload),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client IscsiTargetsClient) CreateOrUpdateSender(req *http.Request) (future IscsiTargetsCreateOrUpdateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client IscsiTargetsClient) CreateOrUpdateResponder(resp *http.Response) (result IscsiTarget, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes an iSCSI Target
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// diskPoolName - the name of the Disk Pool.
// iscsiTargetName - the name of the iSCSI target.
func (client IscsiTargetsClient) Delete(ctx context.Context, resourceGroupName string, diskPoolName string, iscsiTargetName string) (result IscsiTargetsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IscsiTargetsClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]*[0-9A-Za-z]$`, Chain: nil}}},
		{TargetValue: diskPoolName,
			Constraints: []validation.Constraint{{Target: "diskPoolName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "diskPoolName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "diskPoolName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storagepool.IscsiTargetsClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, diskPoolName, iscsiTargetName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagepool.IscsiTargetsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagepool.IscsiTargetsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client IscsiTargetsClient) DeletePreparer(ctx context.Context, resourceGroupName string, diskPoolName string, iscsiTargetName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"diskPoolName":      autorest.Encode("path", diskPoolName),
		"iscsiTargetName":   autorest.Encode("path", iscsiTargetName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-03-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StoragePool/diskPools/{diskPoolName}/iscsiTargets/{iscsiTargetName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client IscsiTargetsClient) DeleteSender(req *http.Request) (future IscsiTargetsDeleteFuture, err error) {
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
func (client IscsiTargetsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets an iSCSI Target.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// diskPoolName - the name of the Disk Pool.
// iscsiTargetName - the name of the iSCSI target.
func (client IscsiTargetsClient) Get(ctx context.Context, resourceGroupName string, diskPoolName string, iscsiTargetName string) (result IscsiTarget, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IscsiTargetsClient.Get")
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
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]*[0-9A-Za-z]$`, Chain: nil}}},
		{TargetValue: diskPoolName,
			Constraints: []validation.Constraint{{Target: "diskPoolName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "diskPoolName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "diskPoolName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storagepool.IscsiTargetsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, diskPoolName, iscsiTargetName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagepool.IscsiTargetsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storagepool.IscsiTargetsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagepool.IscsiTargetsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client IscsiTargetsClient) GetPreparer(ctx context.Context, resourceGroupName string, diskPoolName string, iscsiTargetName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"diskPoolName":      autorest.Encode("path", diskPoolName),
		"iscsiTargetName":   autorest.Encode("path", iscsiTargetName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-03-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StoragePool/diskPools/{diskPoolName}/iscsiTargets/{iscsiTargetName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client IscsiTargetsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client IscsiTargetsClient) GetResponder(resp *http.Response) (result IscsiTarget, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByDiskPool get iSCSI Targets within a Disk Pool
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// diskPoolName - the name of the Disk Pool.
func (client IscsiTargetsClient) ListByDiskPool(ctx context.Context, resourceGroupName string, diskPoolName string) (result IscsiTargetListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IscsiTargetsClient.ListByDiskPool")
		defer func() {
			sc := -1
			if result.itl.Response.Response != nil {
				sc = result.itl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]*[0-9A-Za-z]$`, Chain: nil}}},
		{TargetValue: diskPoolName,
			Constraints: []validation.Constraint{{Target: "diskPoolName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "diskPoolName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "diskPoolName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storagepool.IscsiTargetsClient", "ListByDiskPool", err.Error())
	}

	result.fn = client.listByDiskPoolNextResults
	req, err := client.ListByDiskPoolPreparer(ctx, resourceGroupName, diskPoolName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagepool.IscsiTargetsClient", "ListByDiskPool", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByDiskPoolSender(req)
	if err != nil {
		result.itl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storagepool.IscsiTargetsClient", "ListByDiskPool", resp, "Failure sending request")
		return
	}

	result.itl, err = client.ListByDiskPoolResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagepool.IscsiTargetsClient", "ListByDiskPool", resp, "Failure responding to request")
	}
	if result.itl.hasNextLink() && result.itl.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListByDiskPoolPreparer prepares the ListByDiskPool request.
func (client IscsiTargetsClient) ListByDiskPoolPreparer(ctx context.Context, resourceGroupName string, diskPoolName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"diskPoolName":      autorest.Encode("path", diskPoolName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-03-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StoragePool/diskPools/{diskPoolName}/iscsiTargets", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByDiskPoolSender sends the ListByDiskPool request. The method will close the
// http.Response Body if it receives an error.
func (client IscsiTargetsClient) ListByDiskPoolSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByDiskPoolResponder handles the response to the ListByDiskPool request. The method always
// closes the http.Response Body.
func (client IscsiTargetsClient) ListByDiskPoolResponder(resp *http.Response) (result IscsiTargetList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByDiskPoolNextResults retrieves the next set of results, if any.
func (client IscsiTargetsClient) listByDiskPoolNextResults(ctx context.Context, lastResults IscsiTargetList) (result IscsiTargetList, err error) {
	req, err := lastResults.iscsiTargetListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "storagepool.IscsiTargetsClient", "listByDiskPoolNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByDiskPoolSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "storagepool.IscsiTargetsClient", "listByDiskPoolNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByDiskPoolResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagepool.IscsiTargetsClient", "listByDiskPoolNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByDiskPoolComplete enumerates all values, automatically crossing page boundaries as required.
func (client IscsiTargetsClient) ListByDiskPoolComplete(ctx context.Context, resourceGroupName string, diskPoolName string) (result IscsiTargetListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IscsiTargetsClient.ListByDiskPool")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByDiskPool(ctx, resourceGroupName, diskPoolName)
	return
}
