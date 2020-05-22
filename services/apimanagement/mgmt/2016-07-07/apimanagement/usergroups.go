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

// UserGroupsClient is the apiManagement Client
type UserGroupsClient struct {
	BaseClient
}

// NewUserGroupsClient creates an instance of the UserGroupsClient client.
func NewUserGroupsClient(subscriptionID string) UserGroupsClient {
	return NewUserGroupsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewUserGroupsClientWithBaseURI creates an instance of the UserGroupsClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewUserGroupsClientWithBaseURI(baseURI string, subscriptionID string) UserGroupsClient {
	return UserGroupsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// ListByUser lists all user groups.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// UID - user identifier. Must be unique in the current API Management service instance.
// filter - | Field       | Supported operators    | Supported functions                         |
// |-------------|------------------------|---------------------------------------------|
// | id          | ge, le, eq, ne, gt, lt | substringof, contains, startswith, endswith |
// | name        | ge, le, eq, ne, gt, lt | substringof, contains, startswith, endswith |
// | description | ge, le, eq, ne, gt, lt | substringof, contains, startswith, endswith |
// top - number of records to return.
// skip - number of records to skip.
func (client UserGroupsClient) ListByUser(ctx context.Context, resourceGroupName string, serviceName string, UID string, filter string, top *int32, skip *int32) (result GroupCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UserGroupsClient.ListByUser")
		defer func() {
			sc := -1
			if result.gc.Response.Response != nil {
				sc = result.gc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: UID,
			Constraints: []validation.Constraint{{Target: "UID", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "UID", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "UID", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}},
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMinimum, Rule: int64(1), Chain: nil}}}}},
		{TargetValue: skip,
			Constraints: []validation.Constraint{{Target: "skip", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "skip", Name: validation.InclusiveMinimum, Rule: int64(0), Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("apimanagement.UserGroupsClient", "ListByUser", err.Error())
	}

	result.fn = client.listByUserNextResults
	req, err := client.ListByUserPreparer(ctx, resourceGroupName, serviceName, UID, filter, top, skip)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.UserGroupsClient", "ListByUser", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByUserSender(req)
	if err != nil {
		result.gc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.UserGroupsClient", "ListByUser", resp, "Failure sending request")
		return
	}

	result.gc, err = client.ListByUserResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.UserGroupsClient", "ListByUser", resp, "Failure responding to request")
	}

	return
}

// ListByUserPreparer prepares the ListByUser request.
func (client UserGroupsClient) ListByUserPreparer(ctx context.Context, resourceGroupName string, serviceName string, UID string, filter string, top *int32, skip *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"uid":               autorest.Encode("path", UID),
	}

	const APIVersion = "2016-07-07"
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
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/users/{uid}/groups", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByUserSender sends the ListByUser request. The method will close the
// http.Response Body if it receives an error.
func (client UserGroupsClient) ListByUserSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByUserResponder handles the response to the ListByUser request. The method always
// closes the http.Response Body.
func (client UserGroupsClient) ListByUserResponder(resp *http.Response) (result GroupCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByUserNextResults retrieves the next set of results, if any.
func (client UserGroupsClient) listByUserNextResults(ctx context.Context, lastResults GroupCollection) (result GroupCollection, err error) {
	req, err := lastResults.groupCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "apimanagement.UserGroupsClient", "listByUserNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByUserSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "apimanagement.UserGroupsClient", "listByUserNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByUserResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.UserGroupsClient", "listByUserNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByUserComplete enumerates all values, automatically crossing page boundaries as required.
func (client UserGroupsClient) ListByUserComplete(ctx context.Context, resourceGroupName string, serviceName string, UID string, filter string, top *int32, skip *int32) (result GroupCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UserGroupsClient.ListByUser")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByUser(ctx, resourceGroupName, serviceName, UID, filter, top, skip)
	return
}
