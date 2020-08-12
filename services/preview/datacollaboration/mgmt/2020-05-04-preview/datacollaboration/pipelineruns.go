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

// PipelineRunsClient is the creates a Microsoft.DataCollaboration management client.
type PipelineRunsClient struct {
	BaseClient
}

// NewPipelineRunsClient creates an instance of the PipelineRunsClient client.
func NewPipelineRunsClient(subscriptionID string) PipelineRunsClient {
	return NewPipelineRunsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewPipelineRunsClientWithBaseURI creates an instance of the PipelineRunsClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewPipelineRunsClientWithBaseURI(baseURI string, subscriptionID string) PipelineRunsClient {
	return PipelineRunsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Cancel request to cancel a pipeline run.
// Parameters:
// resourceGroupName - the resource group name.
// workspaceName - the name of the workspace.
// pipelineName - the name of the pipeline.
// pipelineRunName - the name of the pipeline run.
func (client PipelineRunsClient) Cancel(ctx context.Context, resourceGroupName string, workspaceName string, pipelineName string, pipelineRunName string) (result PipelineRunsCancelFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PipelineRunsClient.Cancel")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CancelPreparer(ctx, resourceGroupName, workspaceName, pipelineName, pipelineRunName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datacollaboration.PipelineRunsClient", "Cancel", nil, "Failure preparing request")
		return
	}

	result, err = client.CancelSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datacollaboration.PipelineRunsClient", "Cancel", result.Response(), "Failure sending request")
		return
	}

	return
}

// CancelPreparer prepares the Cancel request.
func (client PipelineRunsClient) CancelPreparer(ctx context.Context, resourceGroupName string, workspaceName string, pipelineName string, pipelineRunName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"pipelineName":      autorest.Encode("path", pipelineName),
		"pipelineRunName":   autorest.Encode("path", pipelineRunName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2020-05-04-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataCollaboration/workspaces/{workspaceName}/pipelines/{pipelineName}/pipelineRuns/{pipelineRunName}/cancel", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CancelSender sends the Cancel request. The method will close the
// http.Response Body if it receives an error.
func (client PipelineRunsClient) CancelSender(req *http.Request) (future PipelineRunsCancelFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// CancelResponder handles the response to the Cancel request. The method always
// closes the http.Response Body.
func (client PipelineRunsClient) CancelResponder(resp *http.Response) (result PipelineRun, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get get a pipeline run in a pipeline
// Parameters:
// resourceGroupName - the resource group name.
// workspaceName - the name of the workspace.
// pipelineName - the name of the pipeline.
// pipelineRunName - the name of the pipeline run.
func (client PipelineRunsClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, pipelineName string, pipelineRunName string) (result PipelineRun, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PipelineRunsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, workspaceName, pipelineName, pipelineRunName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datacollaboration.PipelineRunsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datacollaboration.PipelineRunsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datacollaboration.PipelineRunsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client PipelineRunsClient) GetPreparer(ctx context.Context, resourceGroupName string, workspaceName string, pipelineName string, pipelineRunName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"pipelineName":      autorest.Encode("path", pipelineName),
		"pipelineRunName":   autorest.Encode("path", pipelineRunName),
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
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataCollaboration/workspaces/{workspaceName}/pipelines/{pipelineName}/pipelineRuns/{pipelineRunName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client PipelineRunsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client PipelineRunsClient) GetResponder(resp *http.Response) (result PipelineRun, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByPipeline list pipelines run of a pipeline
// Parameters:
// resourceGroupName - the resource group name.
// workspaceName - the name of the workspace.
// pipelineName - the name of the pipeline.
// skipToken - continuation token
func (client PipelineRunsClient) ListByPipeline(ctx context.Context, resourceGroupName string, workspaceName string, pipelineName string, skipToken string) (result PipelineRunListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PipelineRunsClient.ListByPipeline")
		defer func() {
			sc := -1
			if result.prl.Response.Response != nil {
				sc = result.prl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByPipelineNextResults
	req, err := client.ListByPipelinePreparer(ctx, resourceGroupName, workspaceName, pipelineName, skipToken)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datacollaboration.PipelineRunsClient", "ListByPipeline", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByPipelineSender(req)
	if err != nil {
		result.prl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datacollaboration.PipelineRunsClient", "ListByPipeline", resp, "Failure sending request")
		return
	}

	result.prl, err = client.ListByPipelineResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datacollaboration.PipelineRunsClient", "ListByPipeline", resp, "Failure responding to request")
	}
	if result.prl.hasNextLink() && result.prl.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListByPipelinePreparer prepares the ListByPipeline request.
func (client PipelineRunsClient) ListByPipelinePreparer(ctx context.Context, resourceGroupName string, workspaceName string, pipelineName string, skipToken string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"pipelineName":      autorest.Encode("path", pipelineName),
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
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataCollaboration/workspaces/{workspaceName}/pipelines/{pipelineName}/pipelineRuns", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByPipelineSender sends the ListByPipeline request. The method will close the
// http.Response Body if it receives an error.
func (client PipelineRunsClient) ListByPipelineSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByPipelineResponder handles the response to the ListByPipeline request. The method always
// closes the http.Response Body.
func (client PipelineRunsClient) ListByPipelineResponder(resp *http.Response) (result PipelineRunList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByPipelineNextResults retrieves the next set of results, if any.
func (client PipelineRunsClient) listByPipelineNextResults(ctx context.Context, lastResults PipelineRunList) (result PipelineRunList, err error) {
	req, err := lastResults.pipelineRunListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "datacollaboration.PipelineRunsClient", "listByPipelineNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByPipelineSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "datacollaboration.PipelineRunsClient", "listByPipelineNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByPipelineResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datacollaboration.PipelineRunsClient", "listByPipelineNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByPipelineComplete enumerates all values, automatically crossing page boundaries as required.
func (client PipelineRunsClient) ListByPipelineComplete(ctx context.Context, resourceGroupName string, workspaceName string, pipelineName string, skipToken string) (result PipelineRunListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PipelineRunsClient.ListByPipeline")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByPipeline(ctx, resourceGroupName, workspaceName, pipelineName, skipToken)
	return
}
