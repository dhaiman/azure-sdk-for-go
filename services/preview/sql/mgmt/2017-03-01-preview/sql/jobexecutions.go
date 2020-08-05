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
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/tracing"
	"github.com/satori/go.uuid"
	"net/http"
)

// JobExecutionsClient is the the Azure SQL Database management API provides a RESTful set of web services that
// interact with Azure SQL Database services to manage your databases. The API enables you to create, retrieve, update,
// and delete databases.
type JobExecutionsClient struct {
	BaseClient
}

// NewJobExecutionsClient creates an instance of the JobExecutionsClient client.
func NewJobExecutionsClient(subscriptionID string) JobExecutionsClient {
	return NewJobExecutionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewJobExecutionsClientWithBaseURI creates an instance of the JobExecutionsClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewJobExecutionsClientWithBaseURI(baseURI string, subscriptionID string) JobExecutionsClient {
	return JobExecutionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Cancel requests cancellation of a job execution.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
// jobAgentName - the name of the job agent.
// jobName - the name of the job.
// jobExecutionID - the id of the job execution to cancel.
func (client JobExecutionsClient) Cancel(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, jobExecutionID uuid.UUID) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobExecutionsClient.Cancel")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CancelPreparer(ctx, resourceGroupName, serverName, jobAgentName, jobName, jobExecutionID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.JobExecutionsClient", "Cancel", nil, "Failure preparing request")
		return
	}

	resp, err := client.CancelSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "sql.JobExecutionsClient", "Cancel", resp, "Failure sending request")
		return
	}

	result, err = client.CancelResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.JobExecutionsClient", "Cancel", resp, "Failure responding to request")
	}

	return
}

// CancelPreparer prepares the Cancel request.
func (client JobExecutionsClient) CancelPreparer(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, jobExecutionID uuid.UUID) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobAgentName":      autorest.Encode("path", jobAgentName),
		"jobExecutionId":    autorest.Encode("path", jobExecutionID),
		"jobName":           autorest.Encode("path", jobName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/jobAgents/{jobAgentName}/jobs/{jobName}/executions/{jobExecutionId}/cancel", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CancelSender sends the Cancel request. The method will close the
// http.Response Body if it receives an error.
func (client JobExecutionsClient) CancelSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CancelResponder handles the response to the Cancel request. The method always
// closes the http.Response Body.
func (client JobExecutionsClient) CancelResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Create starts an elastic job execution.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
// jobAgentName - the name of the job agent.
// jobName - the name of the job to get.
func (client JobExecutionsClient) Create(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string) (result JobExecutionsCreateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobExecutionsClient.Create")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreatePreparer(ctx, resourceGroupName, serverName, jobAgentName, jobName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.JobExecutionsClient", "Create", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.JobExecutionsClient", "Create", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client JobExecutionsClient) CreatePreparer(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobAgentName":      autorest.Encode("path", jobAgentName),
		"jobName":           autorest.Encode("path", jobName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/jobAgents/{jobAgentName}/jobs/{jobName}/start", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client JobExecutionsClient) CreateSender(req *http.Request) (future JobExecutionsCreateFuture, err error) {
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
func (client JobExecutionsClient) CreateResponder(resp *http.Response) (result JobExecution, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CreateOrUpdate creates or updates a job execution.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
// jobAgentName - the name of the job agent.
// jobName - the name of the job to get.
// jobExecutionID - the job execution id to create the job execution under.
func (client JobExecutionsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, jobExecutionID uuid.UUID) (result JobExecutionsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobExecutionsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, serverName, jobAgentName, jobName, jobExecutionID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.JobExecutionsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.JobExecutionsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client JobExecutionsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, jobExecutionID uuid.UUID) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobAgentName":      autorest.Encode("path", jobAgentName),
		"jobExecutionId":    autorest.Encode("path", jobExecutionID),
		"jobName":           autorest.Encode("path", jobName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/jobAgents/{jobAgentName}/jobs/{jobName}/executions/{jobExecutionId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client JobExecutionsClient) CreateOrUpdateSender(req *http.Request) (future JobExecutionsCreateOrUpdateFuture, err error) {
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
func (client JobExecutionsClient) CreateOrUpdateResponder(resp *http.Response) (result JobExecution, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get gets a job execution.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
// jobAgentName - the name of the job agent.
// jobName - the name of the job.
// jobExecutionID - the id of the job execution
func (client JobExecutionsClient) Get(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, jobExecutionID uuid.UUID) (result JobExecution, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobExecutionsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, serverName, jobAgentName, jobName, jobExecutionID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.JobExecutionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.JobExecutionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.JobExecutionsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client JobExecutionsClient) GetPreparer(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, jobExecutionID uuid.UUID) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobAgentName":      autorest.Encode("path", jobAgentName),
		"jobExecutionId":    autorest.Encode("path", jobExecutionID),
		"jobName":           autorest.Encode("path", jobName),
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
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/jobAgents/{jobAgentName}/jobs/{jobName}/executions/{jobExecutionId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client JobExecutionsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client JobExecutionsClient) GetResponder(resp *http.Response) (result JobExecution, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByAgent lists all executions in a job agent.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
// jobAgentName - the name of the job agent.
// createTimeMin - if specified, only job executions created at or after the specified time are included.
// createTimeMax - if specified, only job executions created before the specified time are included.
// endTimeMin - if specified, only job executions completed at or after the specified time are included.
// endTimeMax - if specified, only job executions completed before the specified time are included.
// isActive - if specified, only active or only completed job executions are included.
// skip - the number of elements in the collection to skip.
// top - the number of elements to return from the collection.
func (client JobExecutionsClient) ListByAgent(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, createTimeMin *date.Time, createTimeMax *date.Time, endTimeMin *date.Time, endTimeMax *date.Time, isActive *bool, skip *int32, top *int32) (result JobExecutionListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobExecutionsClient.ListByAgent")
		defer func() {
			sc := -1
			if result.jelr.Response.Response != nil {
				sc = result.jelr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByAgentNextResults
	req, err := client.ListByAgentPreparer(ctx, resourceGroupName, serverName, jobAgentName, createTimeMin, createTimeMax, endTimeMin, endTimeMax, isActive, skip, top)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.JobExecutionsClient", "ListByAgent", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByAgentSender(req)
	if err != nil {
		result.jelr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.JobExecutionsClient", "ListByAgent", resp, "Failure sending request")
		return
	}

	result.jelr, err = client.ListByAgentResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.JobExecutionsClient", "ListByAgent", resp, "Failure responding to request")
	}
	if result.jelr.hasNextLink() && result.jelr.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListByAgentPreparer prepares the ListByAgent request.
func (client JobExecutionsClient) ListByAgentPreparer(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, createTimeMin *date.Time, createTimeMax *date.Time, endTimeMin *date.Time, endTimeMax *date.Time, isActive *bool, skip *int32, top *int32) (*http.Request, error) {
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
	if createTimeMin != nil {
		queryParameters["createTimeMin"] = autorest.Encode("query", *createTimeMin)
	}
	if createTimeMax != nil {
		queryParameters["createTimeMax"] = autorest.Encode("query", *createTimeMax)
	}
	if endTimeMin != nil {
		queryParameters["endTimeMin"] = autorest.Encode("query", *endTimeMin)
	}
	if endTimeMax != nil {
		queryParameters["endTimeMax"] = autorest.Encode("query", *endTimeMax)
	}
	if isActive != nil {
		queryParameters["isActive"] = autorest.Encode("query", *isActive)
	}
	if skip != nil {
		queryParameters["$skip"] = autorest.Encode("query", *skip)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/jobAgents/{jobAgentName}/executions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByAgentSender sends the ListByAgent request. The method will close the
// http.Response Body if it receives an error.
func (client JobExecutionsClient) ListByAgentSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByAgentResponder handles the response to the ListByAgent request. The method always
// closes the http.Response Body.
func (client JobExecutionsClient) ListByAgentResponder(resp *http.Response) (result JobExecutionListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByAgentNextResults retrieves the next set of results, if any.
func (client JobExecutionsClient) listByAgentNextResults(ctx context.Context, lastResults JobExecutionListResult) (result JobExecutionListResult, err error) {
	req, err := lastResults.jobExecutionListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "sql.JobExecutionsClient", "listByAgentNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByAgentSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "sql.JobExecutionsClient", "listByAgentNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByAgentResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.JobExecutionsClient", "listByAgentNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByAgentComplete enumerates all values, automatically crossing page boundaries as required.
func (client JobExecutionsClient) ListByAgentComplete(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, createTimeMin *date.Time, createTimeMax *date.Time, endTimeMin *date.Time, endTimeMax *date.Time, isActive *bool, skip *int32, top *int32) (result JobExecutionListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobExecutionsClient.ListByAgent")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByAgent(ctx, resourceGroupName, serverName, jobAgentName, createTimeMin, createTimeMax, endTimeMin, endTimeMax, isActive, skip, top)
	return
}

// ListByJob lists a job's executions.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
// jobAgentName - the name of the job agent.
// jobName - the name of the job to get.
// createTimeMin - if specified, only job executions created at or after the specified time are included.
// createTimeMax - if specified, only job executions created before the specified time are included.
// endTimeMin - if specified, only job executions completed at or after the specified time are included.
// endTimeMax - if specified, only job executions completed before the specified time are included.
// isActive - if specified, only active or only completed job executions are included.
// skip - the number of elements in the collection to skip.
// top - the number of elements to return from the collection.
func (client JobExecutionsClient) ListByJob(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, createTimeMin *date.Time, createTimeMax *date.Time, endTimeMin *date.Time, endTimeMax *date.Time, isActive *bool, skip *int32, top *int32) (result JobExecutionListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobExecutionsClient.ListByJob")
		defer func() {
			sc := -1
			if result.jelr.Response.Response != nil {
				sc = result.jelr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByJobNextResults
	req, err := client.ListByJobPreparer(ctx, resourceGroupName, serverName, jobAgentName, jobName, createTimeMin, createTimeMax, endTimeMin, endTimeMax, isActive, skip, top)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.JobExecutionsClient", "ListByJob", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByJobSender(req)
	if err != nil {
		result.jelr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.JobExecutionsClient", "ListByJob", resp, "Failure sending request")
		return
	}

	result.jelr, err = client.ListByJobResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.JobExecutionsClient", "ListByJob", resp, "Failure responding to request")
	}
	if result.jelr.hasNextLink() && result.jelr.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListByJobPreparer prepares the ListByJob request.
func (client JobExecutionsClient) ListByJobPreparer(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, createTimeMin *date.Time, createTimeMax *date.Time, endTimeMin *date.Time, endTimeMax *date.Time, isActive *bool, skip *int32, top *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobAgentName":      autorest.Encode("path", jobAgentName),
		"jobName":           autorest.Encode("path", jobName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if createTimeMin != nil {
		queryParameters["createTimeMin"] = autorest.Encode("query", *createTimeMin)
	}
	if createTimeMax != nil {
		queryParameters["createTimeMax"] = autorest.Encode("query", *createTimeMax)
	}
	if endTimeMin != nil {
		queryParameters["endTimeMin"] = autorest.Encode("query", *endTimeMin)
	}
	if endTimeMax != nil {
		queryParameters["endTimeMax"] = autorest.Encode("query", *endTimeMax)
	}
	if isActive != nil {
		queryParameters["isActive"] = autorest.Encode("query", *isActive)
	}
	if skip != nil {
		queryParameters["$skip"] = autorest.Encode("query", *skip)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/jobAgents/{jobAgentName}/jobs/{jobName}/executions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByJobSender sends the ListByJob request. The method will close the
// http.Response Body if it receives an error.
func (client JobExecutionsClient) ListByJobSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByJobResponder handles the response to the ListByJob request. The method always
// closes the http.Response Body.
func (client JobExecutionsClient) ListByJobResponder(resp *http.Response) (result JobExecutionListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByJobNextResults retrieves the next set of results, if any.
func (client JobExecutionsClient) listByJobNextResults(ctx context.Context, lastResults JobExecutionListResult) (result JobExecutionListResult, err error) {
	req, err := lastResults.jobExecutionListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "sql.JobExecutionsClient", "listByJobNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByJobSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "sql.JobExecutionsClient", "listByJobNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByJobResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.JobExecutionsClient", "listByJobNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByJobComplete enumerates all values, automatically crossing page boundaries as required.
func (client JobExecutionsClient) ListByJobComplete(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, createTimeMin *date.Time, createTimeMax *date.Time, endTimeMin *date.Time, endTimeMax *date.Time, isActive *bool, skip *int32, top *int32) (result JobExecutionListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobExecutionsClient.ListByJob")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByJob(ctx, resourceGroupName, serverName, jobAgentName, jobName, createTimeMin, createTimeMax, endTimeMin, endTimeMax, isActive, skip, top)
	return
}
