package compute

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

// DisksClient is the compute Client
type DisksClient struct {
	BaseClient
}

// NewDisksClient creates an instance of the DisksClient client.
func NewDisksClient(subscriptionID string) DisksClient {
	return NewDisksClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewDisksClientWithBaseURI creates an instance of the DisksClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewDisksClientWithBaseURI(baseURI string, subscriptionID string) DisksClient {
	return DisksClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a disk.
// Parameters:
// resourceGroupName - the name of the resource group.
// diskName - the name of the managed disk that is being created. The name can't be changed after the disk is
// created. Supported characters for the name are a-z, A-Z, 0-9 and _. The maximum name length is 80
// characters.
// disk - disk object supplied in the body of the Put disk operation.
func (client DisksClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, diskName string, disk Disk) (result DisksCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DisksClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: disk,
			Constraints: []validation.Constraint{{Target: "disk.DiskProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "disk.DiskProperties.CreationData", Name: validation.Null, Rule: true,
					Chain: []validation.Constraint{{Target: "disk.DiskProperties.CreationData.ImageReference", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "disk.DiskProperties.CreationData.ImageReference.ID", Name: validation.Null, Rule: true, Chain: nil}}},
					}},
					{Target: "disk.DiskProperties.EncryptionSettingsCollection", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "disk.DiskProperties.EncryptionSettingsCollection.Enabled", Name: validation.Null, Rule: true, Chain: nil}}},
				}}}}}); err != nil {
		return result, validation.NewError("compute.DisksClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, diskName, disk)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DisksClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DisksClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client DisksClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, diskName string, disk Disk) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"diskName":          autorest.Encode("path", diskName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	disk.ManagedBy = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/disks/{diskName}", pathParameters),
		autorest.WithJSON(disk),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client DisksClient) CreateOrUpdateSender(req *http.Request) (future DisksCreateOrUpdateFuture, err error) {
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
func (client DisksClient) CreateOrUpdateResponder(resp *http.Response) (result Disk, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a disk.
// Parameters:
// resourceGroupName - the name of the resource group.
// diskName - the name of the managed disk that is being created. The name can't be changed after the disk is
// created. Supported characters for the name are a-z, A-Z, 0-9 and _. The maximum name length is 80
// characters.
func (client DisksClient) Delete(ctx context.Context, resourceGroupName string, diskName string) (result DisksDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DisksClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, diskName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DisksClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DisksClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client DisksClient) DeletePreparer(ctx context.Context, resourceGroupName string, diskName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"diskName":          autorest.Encode("path", diskName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/disks/{diskName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client DisksClient) DeleteSender(req *http.Request) (future DisksDeleteFuture, err error) {
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
func (client DisksClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets information about a disk.
// Parameters:
// resourceGroupName - the name of the resource group.
// diskName - the name of the managed disk that is being created. The name can't be changed after the disk is
// created. Supported characters for the name are a-z, A-Z, 0-9 and _. The maximum name length is 80
// characters.
func (client DisksClient) Get(ctx context.Context, resourceGroupName string, diskName string) (result Disk, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DisksClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, diskName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DisksClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.DisksClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DisksClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client DisksClient) GetPreparer(ctx context.Context, resourceGroupName string, diskName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"diskName":          autorest.Encode("path", diskName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/disks/{diskName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client DisksClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client DisksClient) GetResponder(resp *http.Response) (result Disk, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GrantAccess grants access to a disk.
// Parameters:
// resourceGroupName - the name of the resource group.
// diskName - the name of the managed disk that is being created. The name can't be changed after the disk is
// created. Supported characters for the name are a-z, A-Z, 0-9 and _. The maximum name length is 80
// characters.
// grantAccessData - access data object supplied in the body of the get disk access operation.
func (client DisksClient) GrantAccess(ctx context.Context, resourceGroupName string, diskName string, grantAccessData GrantAccessData) (result DisksGrantAccessFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DisksClient.GrantAccess")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: grantAccessData,
			Constraints: []validation.Constraint{{Target: "grantAccessData.DurationInSeconds", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("compute.DisksClient", "GrantAccess", err.Error())
	}

	req, err := client.GrantAccessPreparer(ctx, resourceGroupName, diskName, grantAccessData)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DisksClient", "GrantAccess", nil, "Failure preparing request")
		return
	}

	result, err = client.GrantAccessSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DisksClient", "GrantAccess", result.Response(), "Failure sending request")
		return
	}

	return
}

// GrantAccessPreparer prepares the GrantAccess request.
func (client DisksClient) GrantAccessPreparer(ctx context.Context, resourceGroupName string, diskName string, grantAccessData GrantAccessData) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"diskName":          autorest.Encode("path", diskName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/disks/{diskName}/beginGetAccess", pathParameters),
		autorest.WithJSON(grantAccessData),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GrantAccessSender sends the GrantAccess request. The method will close the
// http.Response Body if it receives an error.
func (client DisksClient) GrantAccessSender(req *http.Request) (future DisksGrantAccessFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// GrantAccessResponder handles the response to the GrantAccess request. The method always
// closes the http.Response Body.
func (client DisksClient) GrantAccessResponder(resp *http.Response) (result AccessURI, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists all the disks under a subscription.
func (client DisksClient) List(ctx context.Context) (result DiskListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DisksClient.List")
		defer func() {
			sc := -1
			if result.dl.Response.Response != nil {
				sc = result.dl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DisksClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.dl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.DisksClient", "List", resp, "Failure sending request")
		return
	}

	result.dl, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DisksClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client DisksClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/disks", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client DisksClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client DisksClient) ListResponder(resp *http.Response) (result DiskList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client DisksClient) listNextResults(ctx context.Context, lastResults DiskList) (result DiskList, err error) {
	req, err := lastResults.diskListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "compute.DisksClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "compute.DisksClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DisksClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client DisksClient) ListComplete(ctx context.Context) (result DiskListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DisksClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx)
	return
}

// ListByResourceGroup lists all the disks under a resource group.
// Parameters:
// resourceGroupName - the name of the resource group.
func (client DisksClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (result DiskListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DisksClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.dl.Response.Response != nil {
				sc = result.dl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DisksClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.dl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.DisksClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.dl, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DisksClient", "ListByResourceGroup", resp, "Failure responding to request")
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client DisksClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/disks", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client DisksClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client DisksClient) ListByResourceGroupResponder(resp *http.Response) (result DiskList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client DisksClient) listByResourceGroupNextResults(ctx context.Context, lastResults DiskList) (result DiskList, err error) {
	req, err := lastResults.diskListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "compute.DisksClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "compute.DisksClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DisksClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client DisksClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result DiskListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DisksClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName)
	return
}

// RevokeAccess revokes access to a disk.
// Parameters:
// resourceGroupName - the name of the resource group.
// diskName - the name of the managed disk that is being created. The name can't be changed after the disk is
// created. Supported characters for the name are a-z, A-Z, 0-9 and _. The maximum name length is 80
// characters.
func (client DisksClient) RevokeAccess(ctx context.Context, resourceGroupName string, diskName string) (result DisksRevokeAccessFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DisksClient.RevokeAccess")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.RevokeAccessPreparer(ctx, resourceGroupName, diskName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DisksClient", "RevokeAccess", nil, "Failure preparing request")
		return
	}

	result, err = client.RevokeAccessSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DisksClient", "RevokeAccess", result.Response(), "Failure sending request")
		return
	}

	return
}

// RevokeAccessPreparer prepares the RevokeAccess request.
func (client DisksClient) RevokeAccessPreparer(ctx context.Context, resourceGroupName string, diskName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"diskName":          autorest.Encode("path", diskName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/disks/{diskName}/endGetAccess", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RevokeAccessSender sends the RevokeAccess request. The method will close the
// http.Response Body if it receives an error.
func (client DisksClient) RevokeAccessSender(req *http.Request) (future DisksRevokeAccessFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// RevokeAccessResponder handles the response to the RevokeAccess request. The method always
// closes the http.Response Body.
func (client DisksClient) RevokeAccessResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Update updates (patches) a disk.
// Parameters:
// resourceGroupName - the name of the resource group.
// diskName - the name of the managed disk that is being created. The name can't be changed after the disk is
// created. Supported characters for the name are a-z, A-Z, 0-9 and _. The maximum name length is 80
// characters.
// disk - disk object supplied in the body of the Patch disk operation.
func (client DisksClient) Update(ctx context.Context, resourceGroupName string, diskName string, disk DiskUpdate) (result DisksUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DisksClient.Update")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, resourceGroupName, diskName, disk)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DisksClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.DisksClient", "Update", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client DisksClient) UpdatePreparer(ctx context.Context, resourceGroupName string, diskName string, disk DiskUpdate) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"diskName":          autorest.Encode("path", diskName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/disks/{diskName}", pathParameters),
		autorest.WithJSON(disk),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client DisksClient) UpdateSender(req *http.Request) (future DisksUpdateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client DisksClient) UpdateResponder(resp *http.Response) (result Disk, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
