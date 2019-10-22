package logic

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

// WorkflowRunOperationsClient is the REST API for Azure Logic Apps.
type WorkflowRunOperationsClient struct {
	BaseClient
}

// NewWorkflowRunOperationsClient creates an instance of the WorkflowRunOperationsClient client.
func NewWorkflowRunOperationsClient(subscriptionID string) WorkflowRunOperationsClient {
	return NewWorkflowRunOperationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWorkflowRunOperationsClientWithBaseURI creates an instance of the WorkflowRunOperationsClient client.
func NewWorkflowRunOperationsClientWithBaseURI(baseURI string, subscriptionID string) WorkflowRunOperationsClient {
	return WorkflowRunOperationsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get gets an operation for a run.
// Parameters:
// resourceGroupName - the resource group name.
// workflowName - the workflow name.
// runName - the workflow run name.
// operationID - the workflow operation id.
func (client WorkflowRunOperationsClient) Get(ctx context.Context, resourceGroupName string, workflowName string, runName string, operationID string) (result WorkflowRun, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WorkflowRunOperationsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, workflowName, runName, operationID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.WorkflowRunOperationsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "logic.WorkflowRunOperationsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.WorkflowRunOperationsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client WorkflowRunOperationsClient) GetPreparer(ctx context.Context, resourceGroupName string, workflowName string, runName string, operationID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"operationId":       autorest.Encode("path", operationID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"runName":           autorest.Encode("path", runName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workflowName":      autorest.Encode("path", workflowName),
	}

	const APIVersion = "2019-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs/{runName}/operations/{operationId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowRunOperationsClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client WorkflowRunOperationsClient) GetResponder(resp *http.Response) (result WorkflowRun, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
