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

// APIIssueCommentClient is the apiManagement Client
type APIIssueCommentClient struct {
	BaseClient
}

// NewAPIIssueCommentClient creates an instance of the APIIssueCommentClient client.
func NewAPIIssueCommentClient(subscriptionID string) APIIssueCommentClient {
	return NewAPIIssueCommentClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewAPIIssueCommentClientWithBaseURI creates an instance of the APIIssueCommentClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewAPIIssueCommentClientWithBaseURI(baseURI string, subscriptionID string) APIIssueCommentClient {
	return APIIssueCommentClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates a new Comment for the Issue in an API or updates an existing one.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// apiid - API identifier. Must be unique in the current API Management service instance.
// issueID - issue identifier. Must be unique in the current API Management service instance.
// commentID - comment identifier within an Issue. Must be unique in the current Issue.
// parameters - create parameters.
// ifMatch - eTag of the Issue Entity. ETag should match the current entity state from the header response of
// the GET request or it should be * for unconditional update.
func (client APIIssueCommentClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, apiid string, issueID string, commentID string, parameters IssueCommentContract, ifMatch string) (result IssueCommentContract, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/APIIssueCommentClient.CreateOrUpdate")
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
		{TargetValue: apiid,
			Constraints: []validation.Constraint{{Target: "apiid", Name: validation.MaxLength, Rule: 80, Chain: nil},
				{Target: "apiid", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "apiid", Name: validation.Pattern, Rule: `(^[\w]+$)|(^[\w][\w\-]+[\w]$)`, Chain: nil}}},
		{TargetValue: issueID,
			Constraints: []validation.Constraint{{Target: "issueID", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "issueID", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "issueID", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}},
		{TargetValue: commentID,
			Constraints: []validation.Constraint{{Target: "commentID", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "commentID", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "commentID", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.IssueCommentContractProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.IssueCommentContractProperties.Text", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.IssueCommentContractProperties.UserID", Name: validation.Null, Rule: true, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("apimanagement.APIIssueCommentClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, serviceName, apiid, issueID, commentID, parameters, ifMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.APIIssueCommentClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.APIIssueCommentClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.APIIssueCommentClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client APIIssueCommentClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, serviceName string, apiid string, issueID string, commentID string, parameters IssueCommentContract, ifMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"apiId":             autorest.Encode("path", apiid),
		"commentId":         autorest.Encode("path", commentID),
		"issueId":           autorest.Encode("path", issueID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/issues/{issueId}/comments/{commentId}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	if len(ifMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client APIIssueCommentClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client APIIssueCommentClient) CreateOrUpdateResponder(resp *http.Response) (result IssueCommentContract, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the specified comment from an Issue.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// apiid - API identifier. Must be unique in the current API Management service instance.
// issueID - issue identifier. Must be unique in the current API Management service instance.
// commentID - comment identifier within an Issue. Must be unique in the current Issue.
// ifMatch - eTag of the Issue Entity. ETag should match the current entity state from the header response of
// the GET request or it should be * for unconditional update.
func (client APIIssueCommentClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, apiid string, issueID string, commentID string, ifMatch string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/APIIssueCommentClient.Delete")
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
		{TargetValue: apiid,
			Constraints: []validation.Constraint{{Target: "apiid", Name: validation.MaxLength, Rule: 80, Chain: nil},
				{Target: "apiid", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "apiid", Name: validation.Pattern, Rule: `(^[\w]+$)|(^[\w][\w\-]+[\w]$)`, Chain: nil}}},
		{TargetValue: issueID,
			Constraints: []validation.Constraint{{Target: "issueID", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "issueID", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "issueID", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}},
		{TargetValue: commentID,
			Constraints: []validation.Constraint{{Target: "commentID", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "commentID", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "commentID", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.APIIssueCommentClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, serviceName, apiid, issueID, commentID, ifMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.APIIssueCommentClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "apimanagement.APIIssueCommentClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.APIIssueCommentClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client APIIssueCommentClient) DeletePreparer(ctx context.Context, resourceGroupName string, serviceName string, apiid string, issueID string, commentID string, ifMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"apiId":             autorest.Encode("path", apiid),
		"commentId":         autorest.Encode("path", commentID),
		"issueId":           autorest.Encode("path", issueID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/issues/{issueId}/comments/{commentId}", pathParameters),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client APIIssueCommentClient) DeleteSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client APIIssueCommentClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the details of the issue Comment for an API specified by its identifier.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// apiid - API identifier. Must be unique in the current API Management service instance.
// issueID - issue identifier. Must be unique in the current API Management service instance.
// commentID - comment identifier within an Issue. Must be unique in the current Issue.
func (client APIIssueCommentClient) Get(ctx context.Context, resourceGroupName string, serviceName string, apiid string, issueID string, commentID string) (result IssueCommentContract, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/APIIssueCommentClient.Get")
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
		{TargetValue: apiid,
			Constraints: []validation.Constraint{{Target: "apiid", Name: validation.MaxLength, Rule: 80, Chain: nil},
				{Target: "apiid", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "apiid", Name: validation.Pattern, Rule: `(^[\w]+$)|(^[\w][\w\-]+[\w]$)`, Chain: nil}}},
		{TargetValue: issueID,
			Constraints: []validation.Constraint{{Target: "issueID", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "issueID", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "issueID", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}},
		{TargetValue: commentID,
			Constraints: []validation.Constraint{{Target: "commentID", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "commentID", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "commentID", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.APIIssueCommentClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, serviceName, apiid, issueID, commentID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.APIIssueCommentClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.APIIssueCommentClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.APIIssueCommentClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client APIIssueCommentClient) GetPreparer(ctx context.Context, resourceGroupName string, serviceName string, apiid string, issueID string, commentID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"apiId":             autorest.Encode("path", apiid),
		"commentId":         autorest.Encode("path", commentID),
		"issueId":           autorest.Encode("path", issueID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/issues/{issueId}/comments/{commentId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client APIIssueCommentClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client APIIssueCommentClient) GetResponder(resp *http.Response) (result IssueCommentContract, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetEntityTag gets the entity state (Etag) version of the issue Comment for an API specified by its identifier.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// apiid - API identifier. Must be unique in the current API Management service instance.
// issueID - issue identifier. Must be unique in the current API Management service instance.
// commentID - comment identifier within an Issue. Must be unique in the current Issue.
func (client APIIssueCommentClient) GetEntityTag(ctx context.Context, resourceGroupName string, serviceName string, apiid string, issueID string, commentID string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/APIIssueCommentClient.GetEntityTag")
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
		{TargetValue: apiid,
			Constraints: []validation.Constraint{{Target: "apiid", Name: validation.MaxLength, Rule: 80, Chain: nil},
				{Target: "apiid", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "apiid", Name: validation.Pattern, Rule: `(^[\w]+$)|(^[\w][\w\-]+[\w]$)`, Chain: nil}}},
		{TargetValue: issueID,
			Constraints: []validation.Constraint{{Target: "issueID", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "issueID", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "issueID", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}},
		{TargetValue: commentID,
			Constraints: []validation.Constraint{{Target: "commentID", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "commentID", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "commentID", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.APIIssueCommentClient", "GetEntityTag", err.Error())
	}

	req, err := client.GetEntityTagPreparer(ctx, resourceGroupName, serviceName, apiid, issueID, commentID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.APIIssueCommentClient", "GetEntityTag", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetEntityTagSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "apimanagement.APIIssueCommentClient", "GetEntityTag", resp, "Failure sending request")
		return
	}

	result, err = client.GetEntityTagResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.APIIssueCommentClient", "GetEntityTag", resp, "Failure responding to request")
	}

	return
}

// GetEntityTagPreparer prepares the GetEntityTag request.
func (client APIIssueCommentClient) GetEntityTagPreparer(ctx context.Context, resourceGroupName string, serviceName string, apiid string, issueID string, commentID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"apiId":             autorest.Encode("path", apiid),
		"commentId":         autorest.Encode("path", commentID),
		"issueId":           autorest.Encode("path", issueID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsHead(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/issues/{issueId}/comments/{commentId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetEntityTagSender sends the GetEntityTag request. The method will close the
// http.Response Body if it receives an error.
func (client APIIssueCommentClient) GetEntityTagSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetEntityTagResponder handles the response to the GetEntityTag request. The method always
// closes the http.Response Body.
func (client APIIssueCommentClient) GetEntityTagResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// ListByService lists all comments for the Issue associated with the specified API.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// apiid - API identifier. Must be unique in the current API Management service instance.
// issueID - issue identifier. Must be unique in the current API Management service instance.
// filter - | Field       | Supported operators    | Supported functions               |
// |-------------|------------------------|-----------------------------------|
// | id          | ge, le, eq, ne, gt, lt | substringof, startswith, endswith |
// | userId          | ge, le, eq, ne, gt, lt | substringof, startswith, endswith |
// top - number of records to return.
// skip - number of records to skip.
func (client APIIssueCommentClient) ListByService(ctx context.Context, resourceGroupName string, serviceName string, apiid string, issueID string, filter string, top *int32, skip *int32) (result IssueCommentCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/APIIssueCommentClient.ListByService")
		defer func() {
			sc := -1
			if result.icc.Response.Response != nil {
				sc = result.icc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: apiid,
			Constraints: []validation.Constraint{{Target: "apiid", Name: validation.MaxLength, Rule: 80, Chain: nil},
				{Target: "apiid", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "apiid", Name: validation.Pattern, Rule: `(^[\w]+$)|(^[\w][\w\-]+[\w]$)`, Chain: nil}}},
		{TargetValue: issueID,
			Constraints: []validation.Constraint{{Target: "issueID", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "issueID", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "issueID", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}},
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMinimum, Rule: int64(1), Chain: nil}}}}},
		{TargetValue: skip,
			Constraints: []validation.Constraint{{Target: "skip", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "skip", Name: validation.InclusiveMinimum, Rule: int64(0), Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("apimanagement.APIIssueCommentClient", "ListByService", err.Error())
	}

	result.fn = client.listByServiceNextResults
	req, err := client.ListByServicePreparer(ctx, resourceGroupName, serviceName, apiid, issueID, filter, top, skip)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.APIIssueCommentClient", "ListByService", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByServiceSender(req)
	if err != nil {
		result.icc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.APIIssueCommentClient", "ListByService", resp, "Failure sending request")
		return
	}

	result.icc, err = client.ListByServiceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.APIIssueCommentClient", "ListByService", resp, "Failure responding to request")
	}

	return
}

// ListByServicePreparer prepares the ListByService request.
func (client APIIssueCommentClient) ListByServicePreparer(ctx context.Context, resourceGroupName string, serviceName string, apiid string, issueID string, filter string, top *int32, skip *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"apiId":             autorest.Encode("path", apiid),
		"issueId":           autorest.Encode("path", issueID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-01-01"
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
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/issues/{issueId}/comments", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByServiceSender sends the ListByService request. The method will close the
// http.Response Body if it receives an error.
func (client APIIssueCommentClient) ListByServiceSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListByServiceResponder handles the response to the ListByService request. The method always
// closes the http.Response Body.
func (client APIIssueCommentClient) ListByServiceResponder(resp *http.Response) (result IssueCommentCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByServiceNextResults retrieves the next set of results, if any.
func (client APIIssueCommentClient) listByServiceNextResults(ctx context.Context, lastResults IssueCommentCollection) (result IssueCommentCollection, err error) {
	req, err := lastResults.issueCommentCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "apimanagement.APIIssueCommentClient", "listByServiceNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByServiceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "apimanagement.APIIssueCommentClient", "listByServiceNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByServiceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.APIIssueCommentClient", "listByServiceNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByServiceComplete enumerates all values, automatically crossing page boundaries as required.
func (client APIIssueCommentClient) ListByServiceComplete(ctx context.Context, resourceGroupName string, serviceName string, apiid string, issueID string, filter string, top *int32, skip *int32) (result IssueCommentCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/APIIssueCommentClient.ListByService")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByService(ctx, resourceGroupName, serviceName, apiid, issueID, filter, top, skip)
	return
}
