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

// InvitationsClient is the creates a Microsoft.DataShare management client.
type InvitationsClient struct {
	BaseClient
}

// NewInvitationsClient creates an instance of the InvitationsClient client.
func NewInvitationsClient(subscriptionID string) InvitationsClient {
	return NewInvitationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewInvitationsClientWithBaseURI creates an instance of the InvitationsClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewInvitationsClientWithBaseURI(baseURI string, subscriptionID string) InvitationsClient {
	return InvitationsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create create an invitation
// Parameters:
// resourceGroupName - the resource group name.
// accountName - the name of the share account.
// shareName - the name of the share to send the invitation for.
// invitationName - the name of the invitation.
// invitation - invitation details.
func (client InvitationsClient) Create(ctx context.Context, resourceGroupName string, accountName string, shareName string, invitationName string, invitation Invitation) (result Invitation, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InvitationsClient.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreatePreparer(ctx, resourceGroupName, accountName, shareName, invitationName, invitation)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datashare.InvitationsClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datashare.InvitationsClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datashare.InvitationsClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client InvitationsClient) CreatePreparer(ctx context.Context, resourceGroupName string, accountName string, shareName string, invitationName string, invitation Invitation) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"invitationName":    autorest.Encode("path", invitationName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"shareName":         autorest.Encode("path", shareName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shares/{shareName}/invitations/{invitationName}", pathParameters),
		autorest.WithJSON(invitation),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client InvitationsClient) CreateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client InvitationsClient) CreateResponder(resp *http.Response) (result Invitation, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete an invitation in a share
// Parameters:
// resourceGroupName - the resource group name.
// accountName - the name of the share account.
// shareName - the name of the share.
// invitationName - the name of the invitation.
func (client InvitationsClient) Delete(ctx context.Context, resourceGroupName string, accountName string, shareName string, invitationName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InvitationsClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, accountName, shareName, invitationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datashare.InvitationsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "datashare.InvitationsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datashare.InvitationsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client InvitationsClient) DeletePreparer(ctx context.Context, resourceGroupName string, accountName string, shareName string, invitationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"invitationName":    autorest.Encode("path", invitationName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"shareName":         autorest.Encode("path", shareName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shares/{shareName}/invitations/{invitationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client InvitationsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client InvitationsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get an invitation in a share
// Parameters:
// resourceGroupName - the resource group name.
// accountName - the name of the share account.
// shareName - the name of the share.
// invitationName - the name of the invitation.
func (client InvitationsClient) Get(ctx context.Context, resourceGroupName string, accountName string, shareName string, invitationName string) (result Invitation, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InvitationsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, accountName, shareName, invitationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datashare.InvitationsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datashare.InvitationsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datashare.InvitationsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client InvitationsClient) GetPreparer(ctx context.Context, resourceGroupName string, accountName string, shareName string, invitationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"invitationName":    autorest.Encode("path", invitationName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"shareName":         autorest.Encode("path", shareName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shares/{shareName}/invitations/{invitationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client InvitationsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client InvitationsClient) GetResponder(resp *http.Response) (result Invitation, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByShare list invitations in a share
// Parameters:
// resourceGroupName - the resource group name.
// accountName - the name of the share account.
// shareName - the name of the share.
// skipToken - the continuation token
func (client InvitationsClient) ListByShare(ctx context.Context, resourceGroupName string, accountName string, shareName string, skipToken string) (result InvitationListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InvitationsClient.ListByShare")
		defer func() {
			sc := -1
			if result.il.Response.Response != nil {
				sc = result.il.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByShareNextResults
	req, err := client.ListBySharePreparer(ctx, resourceGroupName, accountName, shareName, skipToken)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datashare.InvitationsClient", "ListByShare", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByShareSender(req)
	if err != nil {
		result.il.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datashare.InvitationsClient", "ListByShare", resp, "Failure sending request")
		return
	}

	result.il, err = client.ListByShareResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datashare.InvitationsClient", "ListByShare", resp, "Failure responding to request")
	}

	return
}

// ListBySharePreparer prepares the ListByShare request.
func (client InvitationsClient) ListBySharePreparer(ctx context.Context, resourceGroupName string, accountName string, shareName string, skipToken string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"shareName":         autorest.Encode("path", shareName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(skipToken) > 0 {
		queryParameters["$skipToken"] = autorest.Encode("query", skipToken)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shares/{shareName}/invitations", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByShareSender sends the ListByShare request. The method will close the
// http.Response Body if it receives an error.
func (client InvitationsClient) ListByShareSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByShareResponder handles the response to the ListByShare request. The method always
// closes the http.Response Body.
func (client InvitationsClient) ListByShareResponder(resp *http.Response) (result InvitationList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByShareNextResults retrieves the next set of results, if any.
func (client InvitationsClient) listByShareNextResults(ctx context.Context, lastResults InvitationList) (result InvitationList, err error) {
	req, err := lastResults.invitationListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "datashare.InvitationsClient", "listByShareNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByShareSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "datashare.InvitationsClient", "listByShareNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByShareResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datashare.InvitationsClient", "listByShareNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByShareComplete enumerates all values, automatically crossing page boundaries as required.
func (client InvitationsClient) ListByShareComplete(ctx context.Context, resourceGroupName string, accountName string, shareName string, skipToken string) (result InvitationListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InvitationsClient.ListByShare")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByShare(ctx, resourceGroupName, accountName, shareName, skipToken)
	return
}
