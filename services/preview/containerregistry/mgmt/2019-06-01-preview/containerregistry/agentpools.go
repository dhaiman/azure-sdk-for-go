package containerregistry

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

// AgentPoolsClient is the client for the AgentPools methods of the Containerregistry service.
type AgentPoolsClient struct {
	BaseClient
}

// NewAgentPoolsClient creates an instance of the AgentPoolsClient client.
func NewAgentPoolsClient(subscriptionID string) AgentPoolsClient {
	return NewAgentPoolsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewAgentPoolsClientWithBaseURI creates an instance of the AgentPoolsClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewAgentPoolsClientWithBaseURI(baseURI string, subscriptionID string) AgentPoolsClient {
	return AgentPoolsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create creates an agent pool for a container registry with the specified parameters.
// Parameters:
// resourceGroupName - the name of the resource group to which the container registry belongs.
// registryName - the name of the container registry.
// agentPoolName - the name of the agent pool.
// agentPool - the parameters of an agent pool that needs to scheduled.
func (client AgentPoolsClient) Create(ctx context.Context, resourceGroupName string, registryName string, agentPoolName string, agentPool AgentPool) (result AgentPoolsCreateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AgentPoolsClient.Create")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: registryName,
			Constraints: []validation.Constraint{{Target: "registryName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "registryName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "registryName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}},
		{TargetValue: agentPoolName,
			Constraints: []validation.Constraint{{Target: "agentPoolName", Name: validation.MaxLength, Rule: 20, Chain: nil},
				{Target: "agentPoolName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "agentPoolName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("containerregistry.AgentPoolsClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, resourceGroupName, registryName, agentPoolName, agentPool)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.AgentPoolsClient", "Create", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.AgentPoolsClient", "Create", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client AgentPoolsClient) CreatePreparer(ctx context.Context, resourceGroupName string, registryName string, agentPoolName string, agentPool AgentPool) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"agentPoolName":     autorest.Encode("path", agentPoolName),
		"registryName":      autorest.Encode("path", registryName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/agentPools/{agentPoolName}", pathParameters),
		autorest.WithJSON(agentPool),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client AgentPoolsClient) CreateSender(req *http.Request) (future AgentPoolsCreateFuture, err error) {
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
func (client AgentPoolsClient) CreateResponder(resp *http.Response) (result AgentPool, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a specified agent pool resource.
// Parameters:
// resourceGroupName - the name of the resource group to which the container registry belongs.
// registryName - the name of the container registry.
// agentPoolName - the name of the agent pool.
func (client AgentPoolsClient) Delete(ctx context.Context, resourceGroupName string, registryName string, agentPoolName string) (result AgentPoolsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AgentPoolsClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: registryName,
			Constraints: []validation.Constraint{{Target: "registryName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "registryName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "registryName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}},
		{TargetValue: agentPoolName,
			Constraints: []validation.Constraint{{Target: "agentPoolName", Name: validation.MaxLength, Rule: 20, Chain: nil},
				{Target: "agentPoolName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "agentPoolName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("containerregistry.AgentPoolsClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, registryName, agentPoolName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.AgentPoolsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.AgentPoolsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client AgentPoolsClient) DeletePreparer(ctx context.Context, resourceGroupName string, registryName string, agentPoolName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"agentPoolName":     autorest.Encode("path", agentPoolName),
		"registryName":      autorest.Encode("path", registryName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/agentPools/{agentPoolName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client AgentPoolsClient) DeleteSender(req *http.Request) (future AgentPoolsDeleteFuture, err error) {
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
func (client AgentPoolsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the detailed information for a given agent pool.
// Parameters:
// resourceGroupName - the name of the resource group to which the container registry belongs.
// registryName - the name of the container registry.
// agentPoolName - the name of the agent pool.
func (client AgentPoolsClient) Get(ctx context.Context, resourceGroupName string, registryName string, agentPoolName string) (result AgentPool, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AgentPoolsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: registryName,
			Constraints: []validation.Constraint{{Target: "registryName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "registryName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "registryName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}},
		{TargetValue: agentPoolName,
			Constraints: []validation.Constraint{{Target: "agentPoolName", Name: validation.MaxLength, Rule: 20, Chain: nil},
				{Target: "agentPoolName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "agentPoolName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("containerregistry.AgentPoolsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, registryName, agentPoolName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.AgentPoolsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerregistry.AgentPoolsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.AgentPoolsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client AgentPoolsClient) GetPreparer(ctx context.Context, resourceGroupName string, registryName string, agentPoolName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"agentPoolName":     autorest.Encode("path", agentPoolName),
		"registryName":      autorest.Encode("path", registryName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/agentPools/{agentPoolName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client AgentPoolsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client AgentPoolsClient) GetResponder(resp *http.Response) (result AgentPool, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetQueueStatus gets the count of queued runs for a given agent pool.
// Parameters:
// resourceGroupName - the name of the resource group to which the container registry belongs.
// registryName - the name of the container registry.
// agentPoolName - the name of the agent pool.
func (client AgentPoolsClient) GetQueueStatus(ctx context.Context, resourceGroupName string, registryName string, agentPoolName string) (result AgentPoolQueueStatus, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AgentPoolsClient.GetQueueStatus")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: registryName,
			Constraints: []validation.Constraint{{Target: "registryName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "registryName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "registryName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}},
		{TargetValue: agentPoolName,
			Constraints: []validation.Constraint{{Target: "agentPoolName", Name: validation.MaxLength, Rule: 20, Chain: nil},
				{Target: "agentPoolName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "agentPoolName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("containerregistry.AgentPoolsClient", "GetQueueStatus", err.Error())
	}

	req, err := client.GetQueueStatusPreparer(ctx, resourceGroupName, registryName, agentPoolName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.AgentPoolsClient", "GetQueueStatus", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetQueueStatusSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerregistry.AgentPoolsClient", "GetQueueStatus", resp, "Failure sending request")
		return
	}

	result, err = client.GetQueueStatusResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.AgentPoolsClient", "GetQueueStatus", resp, "Failure responding to request")
	}

	return
}

// GetQueueStatusPreparer prepares the GetQueueStatus request.
func (client AgentPoolsClient) GetQueueStatusPreparer(ctx context.Context, resourceGroupName string, registryName string, agentPoolName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"agentPoolName":     autorest.Encode("path", agentPoolName),
		"registryName":      autorest.Encode("path", registryName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/agentPools/{agentPoolName}/listQueueStatus", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetQueueStatusSender sends the GetQueueStatus request. The method will close the
// http.Response Body if it receives an error.
func (client AgentPoolsClient) GetQueueStatusSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetQueueStatusResponder handles the response to the GetQueueStatus request. The method always
// closes the http.Response Body.
func (client AgentPoolsClient) GetQueueStatusResponder(resp *http.Response) (result AgentPoolQueueStatus, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists all the agent pools for a specified container registry.
// Parameters:
// resourceGroupName - the name of the resource group to which the container registry belongs.
// registryName - the name of the container registry.
func (client AgentPoolsClient) List(ctx context.Context, resourceGroupName string, registryName string) (result AgentPoolListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AgentPoolsClient.List")
		defer func() {
			sc := -1
			if result.aplr.Response.Response != nil {
				sc = result.aplr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: registryName,
			Constraints: []validation.Constraint{{Target: "registryName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "registryName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "registryName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("containerregistry.AgentPoolsClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, registryName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.AgentPoolsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.aplr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerregistry.AgentPoolsClient", "List", resp, "Failure sending request")
		return
	}

	result.aplr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.AgentPoolsClient", "List", resp, "Failure responding to request")
	}
	if result.aplr.hasNextLink() && result.aplr.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListPreparer prepares the List request.
func (client AgentPoolsClient) ListPreparer(ctx context.Context, resourceGroupName string, registryName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"registryName":      autorest.Encode("path", registryName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/agentPools", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client AgentPoolsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client AgentPoolsClient) ListResponder(resp *http.Response) (result AgentPoolListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client AgentPoolsClient) listNextResults(ctx context.Context, lastResults AgentPoolListResult) (result AgentPoolListResult, err error) {
	req, err := lastResults.agentPoolListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "containerregistry.AgentPoolsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "containerregistry.AgentPoolsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.AgentPoolsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client AgentPoolsClient) ListComplete(ctx context.Context, resourceGroupName string, registryName string) (result AgentPoolListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AgentPoolsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, registryName)
	return
}

// Update updates an agent pool with the specified parameters.
// Parameters:
// resourceGroupName - the name of the resource group to which the container registry belongs.
// registryName - the name of the container registry.
// agentPoolName - the name of the agent pool.
// updateParameters - the parameters for updating an agent pool.
func (client AgentPoolsClient) Update(ctx context.Context, resourceGroupName string, registryName string, agentPoolName string, updateParameters AgentPoolUpdateParameters) (result AgentPoolsUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AgentPoolsClient.Update")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: registryName,
			Constraints: []validation.Constraint{{Target: "registryName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "registryName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "registryName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}},
		{TargetValue: agentPoolName,
			Constraints: []validation.Constraint{{Target: "agentPoolName", Name: validation.MaxLength, Rule: 20, Chain: nil},
				{Target: "agentPoolName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "agentPoolName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("containerregistry.AgentPoolsClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, resourceGroupName, registryName, agentPoolName, updateParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.AgentPoolsClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.AgentPoolsClient", "Update", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client AgentPoolsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, registryName string, agentPoolName string, updateParameters AgentPoolUpdateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"agentPoolName":     autorest.Encode("path", agentPoolName),
		"registryName":      autorest.Encode("path", registryName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/agentPools/{agentPoolName}", pathParameters),
		autorest.WithJSON(updateParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client AgentPoolsClient) UpdateSender(req *http.Request) (future AgentPoolsUpdateFuture, err error) {
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
func (client AgentPoolsClient) UpdateResponder(resp *http.Response) (result AgentPool, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
