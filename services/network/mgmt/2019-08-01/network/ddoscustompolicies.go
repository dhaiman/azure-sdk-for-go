package network

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

// DdosCustomPoliciesClient is the network Client
type DdosCustomPoliciesClient struct {
	BaseClient
}

// NewDdosCustomPoliciesClient creates an instance of the DdosCustomPoliciesClient client.
func NewDdosCustomPoliciesClient(subscriptionID string) DdosCustomPoliciesClient {
	return NewDdosCustomPoliciesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewDdosCustomPoliciesClientWithBaseURI creates an instance of the DdosCustomPoliciesClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewDdosCustomPoliciesClientWithBaseURI(baseURI string, subscriptionID string) DdosCustomPoliciesClient {
	return DdosCustomPoliciesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a DDoS custom policy.
// Parameters:
// resourceGroupName - the name of the resource group.
// ddosCustomPolicyName - the name of the DDoS custom policy.
// parameters - parameters supplied to the create or update operation.
func (client DdosCustomPoliciesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, ddosCustomPolicyName string, parameters DdosCustomPolicy) (result DdosCustomPoliciesCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DdosCustomPoliciesClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, ddosCustomPolicyName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.DdosCustomPoliciesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.DdosCustomPoliciesClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client DdosCustomPoliciesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, ddosCustomPolicyName string, parameters DdosCustomPolicy) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"ddosCustomPolicyName": autorest.Encode("path", ddosCustomPolicyName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	parameters.Etag = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ddosCustomPolicies/{ddosCustomPolicyName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client DdosCustomPoliciesClient) CreateOrUpdateSender(req *http.Request) (future DdosCustomPoliciesCreateOrUpdateFuture, err error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req, sd...)
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client DdosCustomPoliciesClient) CreateOrUpdateResponder(resp *http.Response) (result DdosCustomPolicy, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the specified DDoS custom policy.
// Parameters:
// resourceGroupName - the name of the resource group.
// ddosCustomPolicyName - the name of the DDoS custom policy.
func (client DdosCustomPoliciesClient) Delete(ctx context.Context, resourceGroupName string, ddosCustomPolicyName string) (result DdosCustomPoliciesDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DdosCustomPoliciesClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, ddosCustomPolicyName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.DdosCustomPoliciesClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.DdosCustomPoliciesClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client DdosCustomPoliciesClient) DeletePreparer(ctx context.Context, resourceGroupName string, ddosCustomPolicyName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"ddosCustomPolicyName": autorest.Encode("path", ddosCustomPolicyName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ddosCustomPolicies/{ddosCustomPolicyName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client DdosCustomPoliciesClient) DeleteSender(req *http.Request) (future DdosCustomPoliciesDeleteFuture, err error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req, sd...)
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client DdosCustomPoliciesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets information about the specified DDoS custom policy.
// Parameters:
// resourceGroupName - the name of the resource group.
// ddosCustomPolicyName - the name of the DDoS custom policy.
func (client DdosCustomPoliciesClient) Get(ctx context.Context, resourceGroupName string, ddosCustomPolicyName string) (result DdosCustomPolicy, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DdosCustomPoliciesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, ddosCustomPolicyName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.DdosCustomPoliciesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.DdosCustomPoliciesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.DdosCustomPoliciesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client DdosCustomPoliciesClient) GetPreparer(ctx context.Context, resourceGroupName string, ddosCustomPolicyName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"ddosCustomPolicyName": autorest.Encode("path", ddosCustomPolicyName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ddosCustomPolicies/{ddosCustomPolicyName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client DdosCustomPoliciesClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client DdosCustomPoliciesClient) GetResponder(resp *http.Response) (result DdosCustomPolicy, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// UpdateTags update a DDoS custom policy tags.
// Parameters:
// resourceGroupName - the name of the resource group.
// ddosCustomPolicyName - the name of the DDoS custom policy.
// parameters - parameters supplied to the update DDoS custom policy resource tags.
func (client DdosCustomPoliciesClient) UpdateTags(ctx context.Context, resourceGroupName string, ddosCustomPolicyName string, parameters TagsObject) (result DdosCustomPoliciesUpdateTagsFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DdosCustomPoliciesClient.UpdateTags")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdateTagsPreparer(ctx, resourceGroupName, ddosCustomPolicyName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.DdosCustomPoliciesClient", "UpdateTags", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateTagsSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.DdosCustomPoliciesClient", "UpdateTags", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdateTagsPreparer prepares the UpdateTags request.
func (client DdosCustomPoliciesClient) UpdateTagsPreparer(ctx context.Context, resourceGroupName string, ddosCustomPolicyName string, parameters TagsObject) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"ddosCustomPolicyName": autorest.Encode("path", ddosCustomPolicyName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ddosCustomPolicies/{ddosCustomPolicyName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateTagsSender sends the UpdateTags request. The method will close the
// http.Response Body if it receives an error.
func (client DdosCustomPoliciesClient) UpdateTagsSender(req *http.Request) (future DdosCustomPoliciesUpdateTagsFuture, err error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req, sd...)
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// UpdateTagsResponder handles the response to the UpdateTags request. The method always
// closes the http.Response Body.
func (client DdosCustomPoliciesClient) UpdateTagsResponder(resp *http.Response) (result DdosCustomPolicy, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
