package datacollaboration

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

// InvitationsClient is the creates a Microsoft.DataCollaboration management client.
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
// workspaceName - the name of the workspace.
// proposalName - the name of the proposal to send the invitation for.
// invitationName - the name of the invitation.
// invitation - invitation details.
func (client InvitationsClient) Create(ctx context.Context, resourceGroupName string, workspaceName string, proposalName string, invitationName string, invitation Invitation) (result Invitation, err error) {
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
	req, err := client.CreatePreparer(ctx, resourceGroupName, workspaceName, proposalName, invitationName, invitation)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datacollaboration.InvitationsClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datacollaboration.InvitationsClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datacollaboration.InvitationsClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client InvitationsClient) CreatePreparer(ctx context.Context, resourceGroupName string, workspaceName string, proposalName string, invitationName string, invitation Invitation) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"invitationName":    autorest.Encode("path", invitationName),
		"proposalName":      autorest.Encode("path", proposalName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2020-05-04-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataCollaboration/workspaces/{workspaceName}/proposals/{proposalName}/invitations/{invitationName}", pathParameters),
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

// Delete delete an invitation in a proposal
// Parameters:
// resourceGroupName - the resource group name.
// workspaceName - the name of the workspace.
// proposalName - the name of the proposal.
// invitationName - the name of the invitation.
func (client InvitationsClient) Delete(ctx context.Context, resourceGroupName string, workspaceName string, proposalName string, invitationName string) (result OperationResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InvitationsClient.Delete")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, workspaceName, proposalName, invitationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datacollaboration.InvitationsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datacollaboration.InvitationsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datacollaboration.InvitationsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client InvitationsClient) DeletePreparer(ctx context.Context, resourceGroupName string, workspaceName string, proposalName string, invitationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"invitationName":    autorest.Encode("path", invitationName),
		"proposalName":      autorest.Encode("path", proposalName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2020-05-04-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataCollaboration/workspaces/{workspaceName}/proposals/{proposalName}/invitations/{invitationName}", pathParameters),
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
func (client InvitationsClient) DeleteResponder(resp *http.Response) (result OperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get get an invitation in a proposal
// Parameters:
// resourceGroupName - the resource group name.
// workspaceName - the name of the workspace.
// proposalName - the name of the proposal.
// invitationName - the name of the invitation.
func (client InvitationsClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, proposalName string, invitationName string) (result Invitation, err error) {
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
	req, err := client.GetPreparer(ctx, resourceGroupName, workspaceName, proposalName, invitationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datacollaboration.InvitationsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datacollaboration.InvitationsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datacollaboration.InvitationsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client InvitationsClient) GetPreparer(ctx context.Context, resourceGroupName string, workspaceName string, proposalName string, invitationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"invitationName":    autorest.Encode("path", invitationName),
		"proposalName":      autorest.Encode("path", proposalName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2020-05-04-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataCollaboration/workspaces/{workspaceName}/proposals/{proposalName}/invitations/{invitationName}", pathParameters),
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

// ListByProposal list invitations in a proposal
// Parameters:
// resourceGroupName - the resource group name.
// workspaceName - the name of the workspace.
// proposalName - the name of the proposal.
// skipToken - the continuation token
func (client InvitationsClient) ListByProposal(ctx context.Context, resourceGroupName string, workspaceName string, proposalName string, skipToken string) (result InvitationListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InvitationsClient.ListByProposal")
		defer func() {
			sc := -1
			if result.il.Response.Response != nil {
				sc = result.il.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByProposalNextResults
	req, err := client.ListByProposalPreparer(ctx, resourceGroupName, workspaceName, proposalName, skipToken)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datacollaboration.InvitationsClient", "ListByProposal", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByProposalSender(req)
	if err != nil {
		result.il.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datacollaboration.InvitationsClient", "ListByProposal", resp, "Failure sending request")
		return
	}

	result.il, err = client.ListByProposalResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datacollaboration.InvitationsClient", "ListByProposal", resp, "Failure responding to request")
	}

	return
}

// ListByProposalPreparer prepares the ListByProposal request.
func (client InvitationsClient) ListByProposalPreparer(ctx context.Context, resourceGroupName string, workspaceName string, proposalName string, skipToken string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"proposalName":      autorest.Encode("path", proposalName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2020-05-04-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(skipToken) > 0 {
		queryParameters["$skipToken"] = autorest.Encode("query", skipToken)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataCollaboration/workspaces/{workspaceName}/proposals/{proposalName}/invitations", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByProposalSender sends the ListByProposal request. The method will close the
// http.Response Body if it receives an error.
func (client InvitationsClient) ListByProposalSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByProposalResponder handles the response to the ListByProposal request. The method always
// closes the http.Response Body.
func (client InvitationsClient) ListByProposalResponder(resp *http.Response) (result InvitationList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByProposalNextResults retrieves the next set of results, if any.
func (client InvitationsClient) listByProposalNextResults(ctx context.Context, lastResults InvitationList) (result InvitationList, err error) {
	req, err := lastResults.invitationListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "datacollaboration.InvitationsClient", "listByProposalNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByProposalSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "datacollaboration.InvitationsClient", "listByProposalNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByProposalResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datacollaboration.InvitationsClient", "listByProposalNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByProposalComplete enumerates all values, automatically crossing page boundaries as required.
func (client InvitationsClient) ListByProposalComplete(ctx context.Context, resourceGroupName string, workspaceName string, proposalName string, skipToken string) (result InvitationListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InvitationsClient.ListByProposal")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByProposal(ctx, resourceGroupName, workspaceName, proposalName, skipToken)
	return
}
