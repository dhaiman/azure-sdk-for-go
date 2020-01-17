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

// APIIssuCommentClient is the apiManagement Client
type APIIssuCommentClient struct {
	BaseClient
}

// NewAPIIssuCommentClient creates an instance of the APIIssuCommentClient client.
func NewAPIIssuCommentClient(subscriptionID string) APIIssuCommentClient {
	return NewAPIIssuCommentClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewAPIIssuCommentClientWithBaseURI creates an instance of the APIIssuCommentClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewAPIIssuCommentClientWithBaseURI(baseURI string, subscriptionID string) APIIssuCommentClient {
	return APIIssuCommentClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Head gets the entity state (Etag) version of the issue Comment for an API specified by its identifier.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// apiid - API identifier. Must be unique in the current API Management service instance.
// issueID - issue identifier. Must be unique in the current API Management service instance.
// commentID - comment identifier within an Issue. Must be unique in the current Issue.
func (client APIIssuCommentClient) Head(ctx context.Context, resourceGroupName string, serviceName string, apiid string, issueID string, commentID string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/APIIssuCommentClient.Head")
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
			Constraints: []validation.Constraint{{Target: "apiid", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "apiid", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "apiid", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}},
		{TargetValue: issueID,
			Constraints: []validation.Constraint{{Target: "issueID", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "issueID", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "issueID", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}},
		{TargetValue: commentID,
			Constraints: []validation.Constraint{{Target: "commentID", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "commentID", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "commentID", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.APIIssuCommentClient", "Head", err.Error())
	}

	req, err := client.HeadPreparer(ctx, resourceGroupName, serviceName, apiid, issueID, commentID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.APIIssuCommentClient", "Head", nil, "Failure preparing request")
		return
	}

	resp, err := client.HeadSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "apimanagement.APIIssuCommentClient", "Head", resp, "Failure sending request")
		return
	}

	result, err = client.HeadResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.APIIssuCommentClient", "Head", resp, "Failure responding to request")
	}

	return
}

// HeadPreparer prepares the Head request.
func (client APIIssuCommentClient) HeadPreparer(ctx context.Context, resourceGroupName string, serviceName string, apiid string, issueID string, commentID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"apiId":             autorest.Encode("path", apiid),
		"commentId":         autorest.Encode("path", commentID),
		"issueId":           autorest.Encode("path", issueID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01"
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

// HeadSender sends the Head request. The method will close the
// http.Response Body if it receives an error.
func (client APIIssuCommentClient) HeadSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// HeadResponder handles the response to the Head request. The method always
// closes the http.Response Body.
func (client APIIssuCommentClient) HeadResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}
