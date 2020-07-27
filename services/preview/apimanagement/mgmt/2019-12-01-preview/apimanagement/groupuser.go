package apimanagement

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

// GroupUserClient is the apiManagement Client
type GroupUserClient struct {
	BaseClient
}

// NewGroupUserClient creates an instance of the GroupUserClient client.
func NewGroupUserClient(subscriptionID string) GroupUserClient {
	return NewGroupUserClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewGroupUserClientWithBaseURI creates an instance of the GroupUserClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewGroupUserClientWithBaseURI(baseURI string, subscriptionID string) GroupUserClient {
	return GroupUserClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CheckEntityExists checks that user entity specified by identifier is associated with the group entity.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// groupID - group identifier. Must be unique in the current API Management service instance.
// userID - user identifier. Must be unique in the current API Management service instance.
func (client GroupUserClient) CheckEntityExists(ctx context.Context, resourceGroupName string, serviceName string, groupID string, userID string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/GroupUserClient.CheckEntityExists")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: groupID,
			Constraints: []validation.Constraint{{Target: "groupID", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "groupID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: userID,
			Constraints: []validation.Constraint{{Target: "userID", Name: validation.MaxLength, Rule: 80, Chain: nil},
				{Target: "userID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.GroupUserClient", "CheckEntityExists", err.Error())
	}

	req, err := client.CheckEntityExistsPreparer(ctx, resourceGroupName, serviceName, groupID, userID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.GroupUserClient", "CheckEntityExists", nil, "Failure preparing request")
		return
	}

	resp, err := client.CheckEntityExistsSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "apimanagement.GroupUserClient", "CheckEntityExists", resp, "Failure sending request")
		return
	}

	result, err = client.CheckEntityExistsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.GroupUserClient", "CheckEntityExists", resp, "Failure responding to request")
	}

	return
}

// CheckEntityExistsPreparer prepares the CheckEntityExists request.
func (client GroupUserClient) CheckEntityExistsPreparer(ctx context.Context, resourceGroupName string, serviceName string, groupID string, userID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"groupId":           autorest.Encode("path", groupID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"userId":            autorest.Encode("path", userID),
	}

	const APIVersion = "2019-12-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsHead(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/groups/{groupId}/users/{userId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CheckEntityExistsSender sends the CheckEntityExists request. The method will close the
// http.Response Body if it receives an error.
func (client GroupUserClient) CheckEntityExistsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CheckEntityExistsResponder handles the response to the CheckEntityExists request. The method always
// closes the http.Response Body.
func (client GroupUserClient) CheckEntityExistsResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent, http.StatusNotFound),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Create add existing user to existing group
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// groupID - group identifier. Must be unique in the current API Management service instance.
// userID - user identifier. Must be unique in the current API Management service instance.
func (client GroupUserClient) Create(ctx context.Context, resourceGroupName string, serviceName string, groupID string, userID string) (result UserContract, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/GroupUserClient.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: groupID,
			Constraints: []validation.Constraint{{Target: "groupID", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "groupID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: userID,
			Constraints: []validation.Constraint{{Target: "userID", Name: validation.MaxLength, Rule: 80, Chain: nil},
				{Target: "userID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.GroupUserClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, resourceGroupName, serviceName, groupID, userID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.GroupUserClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.GroupUserClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.GroupUserClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client GroupUserClient) CreatePreparer(ctx context.Context, resourceGroupName string, serviceName string, groupID string, userID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"groupId":           autorest.Encode("path", groupID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"userId":            autorest.Encode("path", userID),
	}

	const APIVersion = "2019-12-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/groups/{groupId}/users/{userId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client GroupUserClient) CreateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client GroupUserClient) CreateResponder(resp *http.Response) (result UserContract, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete remove existing user from existing group.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// groupID - group identifier. Must be unique in the current API Management service instance.
// userID - user identifier. Must be unique in the current API Management service instance.
func (client GroupUserClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, groupID string, userID string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/GroupUserClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: groupID,
			Constraints: []validation.Constraint{{Target: "groupID", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "groupID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: userID,
			Constraints: []validation.Constraint{{Target: "userID", Name: validation.MaxLength, Rule: 80, Chain: nil},
				{Target: "userID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.GroupUserClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, serviceName, groupID, userID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.GroupUserClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "apimanagement.GroupUserClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.GroupUserClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client GroupUserClient) DeletePreparer(ctx context.Context, resourceGroupName string, serviceName string, groupID string, userID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"groupId":           autorest.Encode("path", groupID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"userId":            autorest.Encode("path", userID),
	}

	const APIVersion = "2019-12-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/groups/{groupId}/users/{userId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client GroupUserClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client GroupUserClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// List lists a collection of user entities associated with the group.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// groupID - group identifier. Must be unique in the current API Management service instance.
// filter - |   Field     |     Usage     |     Supported operators     |     Supported functions
// |</br>|-------------|-------------|-------------|-------------|</br>| name | filter | ge, le, eq, ne, gt, lt
// | substringof, contains, startswith, endswith | </br>| firstName | filter | ge, le, eq, ne, gt, lt |
// substringof, contains, startswith, endswith | </br>| lastName | filter | ge, le, eq, ne, gt, lt |
// substringof, contains, startswith, endswith | </br>| email | filter | ge, le, eq, ne, gt, lt | substringof,
// contains, startswith, endswith | </br>| registrationDate | filter | ge, le, eq, ne, gt, lt |     | </br>|
// note | filter | ge, le, eq, ne, gt, lt | substringof, contains, startswith, endswith | </br>
// top - number of records to return.
// skip - number of records to skip.
func (client GroupUserClient) List(ctx context.Context, resourceGroupName string, serviceName string, groupID string, filter string, top *int32, skip *int32) (result UserCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/GroupUserClient.List")
		defer func() {
			sc := -1
			if result.uc.Response.Response != nil {
				sc = result.uc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: groupID,
			Constraints: []validation.Constraint{{Target: "groupID", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "groupID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMinimum, Rule: int64(1), Chain: nil}}}}},
		{TargetValue: skip,
			Constraints: []validation.Constraint{{Target: "skip", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "skip", Name: validation.InclusiveMinimum, Rule: int64(0), Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("apimanagement.GroupUserClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, serviceName, groupID, filter, top, skip)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.GroupUserClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.uc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.GroupUserClient", "List", resp, "Failure sending request")
		return
	}

	result.uc, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.GroupUserClient", "List", resp, "Failure responding to request")
	}
	if result.uc.hasNextLink() && result.uc.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListPreparer prepares the List request.
func (client GroupUserClient) ListPreparer(ctx context.Context, resourceGroupName string, serviceName string, groupID string, filter string, top *int32, skip *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"groupId":           autorest.Encode("path", groupID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-12-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if skip != nil {
		queryParameters["$skip"] = autorest.Encode("query", *skip)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/groups/{groupId}/users", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client GroupUserClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client GroupUserClient) ListResponder(resp *http.Response) (result UserCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client GroupUserClient) listNextResults(ctx context.Context, lastResults UserCollection) (result UserCollection, err error) {
	req, err := lastResults.userCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "apimanagement.GroupUserClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "apimanagement.GroupUserClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.GroupUserClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client GroupUserClient) ListComplete(ctx context.Context, resourceGroupName string, serviceName string, groupID string, filter string, top *int32, skip *int32) (result UserCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/GroupUserClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, serviceName, groupID, filter, top, skip)
	return
}
