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

// JobStepsClient is the the Azure SQL Database management API provides a RESTful set of web services that interact
// with Azure SQL Database services to manage your databases. The API enables you to create, retrieve, update, and
// delete databases.
type JobStepsClient struct {
	BaseClient
}

// NewJobStepsClient creates an instance of the JobStepsClient client.
func NewJobStepsClient(subscriptionID string) JobStepsClient {
	return NewJobStepsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewJobStepsClientWithBaseURI creates an instance of the JobStepsClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewJobStepsClientWithBaseURI(baseURI string, subscriptionID string) JobStepsClient {
	return JobStepsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a job step. This will implicitly create a new job version.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
// jobAgentName - the name of the job agent.
// jobName - the name of the job.
// stepName - the name of the job step.
// parameters - the requested state of the job step.
func (client JobStepsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, stepName string, parameters JobStep) (result JobStep, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobStepsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.JobStepProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.JobStepProperties.TargetGroup", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.JobStepProperties.Credential", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.JobStepProperties.Action", Name: validation.Null, Rule: true,
						Chain: []validation.Constraint{{Target: "parameters.JobStepProperties.Action.Value", Name: validation.Null, Rule: true, Chain: nil}}},
					{Target: "parameters.JobStepProperties.Output", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.JobStepProperties.Output.ServerName", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "parameters.JobStepProperties.Output.DatabaseName", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "parameters.JobStepProperties.Output.TableName", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "parameters.JobStepProperties.Output.Credential", Name: validation.Null, Rule: true, Chain: nil},
						}},
				}}}}}); err != nil {
		return result, validation.NewError("sql.JobStepsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, serverName, jobAgentName, jobName, stepName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.JobStepsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.JobStepsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.JobStepsClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client JobStepsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, stepName string, parameters JobStep) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobAgentName":      autorest.Encode("path", jobAgentName),
		"jobName":           autorest.Encode("path", jobName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"stepName":          autorest.Encode("path", stepName),
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
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/jobAgents/{jobAgentName}/jobs/{jobName}/steps/{stepName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client JobStepsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client JobStepsClient) CreateOrUpdateResponder(resp *http.Response) (result JobStep, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a job step. This will implicitly create a new job version.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
// jobAgentName - the name of the job agent.
// jobName - the name of the job.
// stepName - the name of the job step to delete.
func (client JobStepsClient) Delete(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, stepName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobStepsClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, serverName, jobAgentName, jobName, stepName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.JobStepsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "sql.JobStepsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.JobStepsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client JobStepsClient) DeletePreparer(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, stepName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobAgentName":      autorest.Encode("path", jobAgentName),
		"jobName":           autorest.Encode("path", jobName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"stepName":          autorest.Encode("path", stepName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/jobAgents/{jobAgentName}/jobs/{jobName}/steps/{stepName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client JobStepsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client JobStepsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets a job step in a job's current version.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
// jobAgentName - the name of the job agent.
// jobName - the name of the job.
// stepName - the name of the job step.
func (client JobStepsClient) Get(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, stepName string) (result JobStep, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobStepsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, serverName, jobAgentName, jobName, stepName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.JobStepsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.JobStepsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.JobStepsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client JobStepsClient) GetPreparer(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, stepName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobAgentName":      autorest.Encode("path", jobAgentName),
		"jobName":           autorest.Encode("path", jobName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"stepName":          autorest.Encode("path", stepName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/jobAgents/{jobAgentName}/jobs/{jobName}/steps/{stepName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client JobStepsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client JobStepsClient) GetResponder(resp *http.Response) (result JobStep, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetByVersion gets the specified version of a job step.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
// jobAgentName - the name of the job agent.
// jobName - the name of the job.
// jobVersion - the version of the job to get.
// stepName - the name of the job step.
func (client JobStepsClient) GetByVersion(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, jobVersion int32, stepName string) (result JobStep, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobStepsClient.GetByVersion")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetByVersionPreparer(ctx, resourceGroupName, serverName, jobAgentName, jobName, jobVersion, stepName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.JobStepsClient", "GetByVersion", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetByVersionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.JobStepsClient", "GetByVersion", resp, "Failure sending request")
		return
	}

	result, err = client.GetByVersionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.JobStepsClient", "GetByVersion", resp, "Failure responding to request")
	}

	return
}

// GetByVersionPreparer prepares the GetByVersion request.
func (client JobStepsClient) GetByVersionPreparer(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, jobVersion int32, stepName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobAgentName":      autorest.Encode("path", jobAgentName),
		"jobName":           autorest.Encode("path", jobName),
		"jobVersion":        autorest.Encode("path", jobVersion),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"stepName":          autorest.Encode("path", stepName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/jobAgents/{jobAgentName}/jobs/{jobName}/versions/{jobVersion}/steps/{stepName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetByVersionSender sends the GetByVersion request. The method will close the
// http.Response Body if it receives an error.
func (client JobStepsClient) GetByVersionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetByVersionResponder handles the response to the GetByVersion request. The method always
// closes the http.Response Body.
func (client JobStepsClient) GetByVersionResponder(resp *http.Response) (result JobStep, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByJob gets all job steps for a job's current version.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
// jobAgentName - the name of the job agent.
// jobName - the name of the job to get.
func (client JobStepsClient) ListByJob(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string) (result JobStepListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobStepsClient.ListByJob")
		defer func() {
			sc := -1
			if result.jslr.Response.Response != nil {
				sc = result.jslr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByJobNextResults
	req, err := client.ListByJobPreparer(ctx, resourceGroupName, serverName, jobAgentName, jobName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.JobStepsClient", "ListByJob", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByJobSender(req)
	if err != nil {
		result.jslr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.JobStepsClient", "ListByJob", resp, "Failure sending request")
		return
	}

	result.jslr, err = client.ListByJobResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.JobStepsClient", "ListByJob", resp, "Failure responding to request")
	}

	return
}

// ListByJobPreparer prepares the ListByJob request.
func (client JobStepsClient) ListByJobPreparer(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string) (*http.Request, error) {
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
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/jobAgents/{jobAgentName}/jobs/{jobName}/steps", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByJobSender sends the ListByJob request. The method will close the
// http.Response Body if it receives an error.
func (client JobStepsClient) ListByJobSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByJobResponder handles the response to the ListByJob request. The method always
// closes the http.Response Body.
func (client JobStepsClient) ListByJobResponder(resp *http.Response) (result JobStepListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByJobNextResults retrieves the next set of results, if any.
func (client JobStepsClient) listByJobNextResults(ctx context.Context, lastResults JobStepListResult) (result JobStepListResult, err error) {
	req, err := lastResults.jobStepListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "sql.JobStepsClient", "listByJobNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByJobSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "sql.JobStepsClient", "listByJobNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByJobResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.JobStepsClient", "listByJobNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByJobComplete enumerates all values, automatically crossing page boundaries as required.
func (client JobStepsClient) ListByJobComplete(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string) (result JobStepListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobStepsClient.ListByJob")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByJob(ctx, resourceGroupName, serverName, jobAgentName, jobName)
	return
}

// ListByVersion gets all job steps in the specified job version.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
// jobAgentName - the name of the job agent.
// jobName - the name of the job to get.
// jobVersion - the version of the job to get.
func (client JobStepsClient) ListByVersion(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, jobVersion int32) (result JobStepListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobStepsClient.ListByVersion")
		defer func() {
			sc := -1
			if result.jslr.Response.Response != nil {
				sc = result.jslr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByVersionNextResults
	req, err := client.ListByVersionPreparer(ctx, resourceGroupName, serverName, jobAgentName, jobName, jobVersion)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.JobStepsClient", "ListByVersion", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByVersionSender(req)
	if err != nil {
		result.jslr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.JobStepsClient", "ListByVersion", resp, "Failure sending request")
		return
	}

	result.jslr, err = client.ListByVersionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.JobStepsClient", "ListByVersion", resp, "Failure responding to request")
	}

	return
}

// ListByVersionPreparer prepares the ListByVersion request.
func (client JobStepsClient) ListByVersionPreparer(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, jobVersion int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobAgentName":      autorest.Encode("path", jobAgentName),
		"jobName":           autorest.Encode("path", jobName),
		"jobVersion":        autorest.Encode("path", jobVersion),
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
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/jobAgents/{jobAgentName}/jobs/{jobName}/versions/{jobVersion}/steps", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByVersionSender sends the ListByVersion request. The method will close the
// http.Response Body if it receives an error.
func (client JobStepsClient) ListByVersionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByVersionResponder handles the response to the ListByVersion request. The method always
// closes the http.Response Body.
func (client JobStepsClient) ListByVersionResponder(resp *http.Response) (result JobStepListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByVersionNextResults retrieves the next set of results, if any.
func (client JobStepsClient) listByVersionNextResults(ctx context.Context, lastResults JobStepListResult) (result JobStepListResult, err error) {
	req, err := lastResults.jobStepListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "sql.JobStepsClient", "listByVersionNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByVersionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "sql.JobStepsClient", "listByVersionNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByVersionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.JobStepsClient", "listByVersionNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByVersionComplete enumerates all values, automatically crossing page boundaries as required.
func (client JobStepsClient) ListByVersionComplete(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, jobVersion int32) (result JobStepListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/JobStepsClient.ListByVersion")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByVersion(ctx, resourceGroupName, serverName, jobAgentName, jobName, jobVersion)
	return
}
