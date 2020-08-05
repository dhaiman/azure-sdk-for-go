package sql

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

// JobAgentsClient is the the Azure SQL Database management API provides a RESTful set of web services that interact
// with Azure SQL Database services to manage your databases. The API enables you to create, retrieve, update, and
// delete databases.
type JobAgentsClient struct {
	BaseClient
}

// NewJobAgentsClient creates an instance of the JobAgentsClient client.
func NewJobAgentsClient(subscriptionID string) JobAgentsClient {
	return NewJobAgentsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewJobAgentsClientWithBaseURI creates an instance of the JobAgentsClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewJobAgentsClientWithBaseURI(baseURI string, subscriptionID string) JobAgentsClient {
	return JobAgentsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a job agent.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
// jobAgentName - the name of the job agent to be created or updated.
// parameters - the requested job agent resource state.
func (client JobAgentsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, parameters JobAgent) (result JobAgentsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobAgentsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.Sku", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.Sku.Name", Name: validation.Null, Rule: true, Chain: nil}}},
				{Target: "parameters.JobAgentProperties", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "parameters.JobAgentProperties.DatabaseID", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("sql.JobAgentsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, serverName, jobAgentName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.JobAgentsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.JobAgentsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client JobAgentsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, parameters JobAgent) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobAgentName":      autorest.Encode("path", jobAgentName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/jobAgents/{jobAgentName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client JobAgentsClient) CreateOrUpdateSender(req *http.Request) (future JobAgentsCreateOrUpdateFuture, err error) {
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
func (client JobAgentsClient) CreateOrUpdateResponder(resp *http.Response) (result JobAgent, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a job agent.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
// jobAgentName - the name of the job agent to be deleted.
func (client JobAgentsClient) Delete(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string) (result JobAgentsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobAgentsClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, serverName, jobAgentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.JobAgentsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.JobAgentsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client JobAgentsClient) DeletePreparer(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobAgentName":      autorest.Encode("path", jobAgentName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/jobAgents/{jobAgentName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client JobAgentsClient) DeleteSender(req *http.Request) (future JobAgentsDeleteFuture, err error) {
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
func (client JobAgentsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets a job agent.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
// jobAgentName - the name of the job agent to be retrieved.
func (client JobAgentsClient) Get(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string) (result JobAgent, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobAgentsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, serverName, jobAgentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.JobAgentsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.JobAgentsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.JobAgentsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client JobAgentsClient) GetPreparer(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobAgentName":      autorest.Encode("path", jobAgentName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/jobAgents/{jobAgentName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client JobAgentsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client JobAgentsClient) GetResponder(resp *http.Response) (result JobAgent, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByServer gets a list of job agents in a server.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
func (client JobAgentsClient) ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result JobAgentListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobAgentsClient.ListByServer")
		defer func() {
			sc := -1
			if result.jalr.Response.Response != nil {
				sc = result.jalr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByServerNextResults
	req, err := client.ListByServerPreparer(ctx, resourceGroupName, serverName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.JobAgentsClient", "ListByServer", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByServerSender(req)
	if err != nil {
		result.jalr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.JobAgentsClient", "ListByServer", resp, "Failure sending request")
		return
	}

	result.jalr, err = client.ListByServerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.JobAgentsClient", "ListByServer", resp, "Failure responding to request")
	}
	if result.jalr.hasNextLink() && result.jalr.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListByServerPreparer prepares the ListByServer request.
func (client JobAgentsClient) ListByServerPreparer(ctx context.Context, resourceGroupName string, serverName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/jobAgents", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByServerSender sends the ListByServer request. The method will close the
// http.Response Body if it receives an error.
func (client JobAgentsClient) ListByServerSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByServerResponder handles the response to the ListByServer request. The method always
// closes the http.Response Body.
func (client JobAgentsClient) ListByServerResponder(resp *http.Response) (result JobAgentListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByServerNextResults retrieves the next set of results, if any.
func (client JobAgentsClient) listByServerNextResults(ctx context.Context, lastResults JobAgentListResult) (result JobAgentListResult, err error) {
	req, err := lastResults.jobAgentListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "sql.JobAgentsClient", "listByServerNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByServerSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "sql.JobAgentsClient", "listByServerNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByServerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.JobAgentsClient", "listByServerNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByServerComplete enumerates all values, automatically crossing page boundaries as required.
func (client JobAgentsClient) ListByServerComplete(ctx context.Context, resourceGroupName string, serverName string) (result JobAgentListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobAgentsClient.ListByServer")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByServer(ctx, resourceGroupName, serverName)
	return
}

// Update updates a job agent.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
// jobAgentName - the name of the job agent to be updated.
// parameters - the update to the job agent.
func (client JobAgentsClient) Update(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, parameters JobAgentUpdate) (result JobAgentsUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobAgentsClient.Update")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, resourceGroupName, serverName, jobAgentName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.JobAgentsClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.JobAgentsClient", "Update", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client JobAgentsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, parameters JobAgentUpdate) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobAgentName":      autorest.Encode("path", jobAgentName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/jobAgents/{jobAgentName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client JobAgentsClient) UpdateSender(req *http.Request) (future JobAgentsUpdateFuture, err error) {
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
func (client JobAgentsClient) UpdateResponder(resp *http.Response) (result JobAgent, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
