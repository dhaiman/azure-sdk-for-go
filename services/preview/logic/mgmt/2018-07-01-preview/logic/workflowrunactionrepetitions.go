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
    "github.com/Azure/go-autorest/autorest"
    "github.com/Azure/go-autorest/autorest/azure"
    "net/http"
    "context"
    "github.com/Azure/go-autorest/tracing"
)

// WorkflowRunActionRepetitionsClient is the REST API for Azure Logic Apps.
type WorkflowRunActionRepetitionsClient struct {
    BaseClient
}
// NewWorkflowRunActionRepetitionsClient creates an instance of the WorkflowRunActionRepetitionsClient client.
func NewWorkflowRunActionRepetitionsClient(subscriptionID string) WorkflowRunActionRepetitionsClient {
    return NewWorkflowRunActionRepetitionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWorkflowRunActionRepetitionsClientWithBaseURI creates an instance of the WorkflowRunActionRepetitionsClient
// client using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI
// (sovereign clouds, Azure stack).
    func NewWorkflowRunActionRepetitionsClientWithBaseURI(baseURI string, subscriptionID string) WorkflowRunActionRepetitionsClient {
        return WorkflowRunActionRepetitionsClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// Get get a workflow run action repetition.
    // Parameters:
        // resourceGroupName - the resource group name.
        // workflowName - the workflow name.
        // runName - the workflow run name.
        // actionName - the workflow action name.
        // repetitionName - the workflow repetition.
func (client WorkflowRunActionRepetitionsClient) Get(ctx context.Context, resourceGroupName string, workflowName string, runName string, actionName string, repetitionName string) (result WorkflowRunActionRepetitionDefinition, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/WorkflowRunActionRepetitionsClient.Get")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.GetPreparer(ctx, resourceGroupName, workflowName, runName, actionName, repetitionName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "logic.WorkflowRunActionRepetitionsClient", "Get", nil , "Failure preparing request")
    return
    }

            resp, err := client.GetSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "logic.WorkflowRunActionRepetitionsClient", "Get", resp, "Failure sending request")
            return
            }

            result, err = client.GetResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "logic.WorkflowRunActionRepetitionsClient", "Get", resp, "Failure responding to request")
            }

    return
    }

    // GetPreparer prepares the Get request.
    func (client WorkflowRunActionRepetitionsClient) GetPreparer(ctx context.Context, resourceGroupName string, workflowName string, runName string, actionName string, repetitionName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "actionName": autorest.Encode("path",actionName),
            "repetitionName": autorest.Encode("path",repetitionName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "runName": autorest.Encode("path",runName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            "workflowName": autorest.Encode("path",workflowName),
            }

                        const APIVersion = "2018-07-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs/{runName}/actions/{actionName}/repetitions/{repetitionName}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client WorkflowRunActionRepetitionsClient) GetSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client WorkflowRunActionRepetitionsClient) GetResponder(resp *http.Response) (result WorkflowRunActionRepetitionDefinition, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// List get all of a workflow run action repetitions.
    // Parameters:
        // resourceGroupName - the resource group name.
        // workflowName - the workflow name.
        // runName - the workflow run name.
        // actionName - the workflow action name.
func (client WorkflowRunActionRepetitionsClient) List(ctx context.Context, resourceGroupName string, workflowName string, runName string, actionName string) (result WorkflowRunActionRepetitionDefinitionCollection, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/WorkflowRunActionRepetitionsClient.List")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.ListPreparer(ctx, resourceGroupName, workflowName, runName, actionName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "logic.WorkflowRunActionRepetitionsClient", "List", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "logic.WorkflowRunActionRepetitionsClient", "List", resp, "Failure sending request")
            return
            }

            result, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "logic.WorkflowRunActionRepetitionsClient", "List", resp, "Failure responding to request")
            }

    return
    }

    // ListPreparer prepares the List request.
    func (client WorkflowRunActionRepetitionsClient) ListPreparer(ctx context.Context, resourceGroupName string, workflowName string, runName string, actionName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "actionName": autorest.Encode("path",actionName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "runName": autorest.Encode("path",runName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            "workflowName": autorest.Encode("path",workflowName),
            }

                        const APIVersion = "2018-07-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs/{runName}/actions/{actionName}/repetitions",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListSender sends the List request. The method will close the
    // http.Response Body if it receives an error.
    func (client WorkflowRunActionRepetitionsClient) ListSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client WorkflowRunActionRepetitionsClient) ListResponder(resp *http.Response) (result WorkflowRunActionRepetitionDefinitionCollection, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// ListExpressionTraces lists a workflow run expression trace.
    // Parameters:
        // resourceGroupName - the resource group name.
        // workflowName - the workflow name.
        // runName - the workflow run name.
        // actionName - the workflow action name.
        // repetitionName - the workflow repetition.
func (client WorkflowRunActionRepetitionsClient) ListExpressionTraces(ctx context.Context, resourceGroupName string, workflowName string, runName string, actionName string, repetitionName string) (result ExpressionTraces, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/WorkflowRunActionRepetitionsClient.ListExpressionTraces")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.ListExpressionTracesPreparer(ctx, resourceGroupName, workflowName, runName, actionName, repetitionName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "logic.WorkflowRunActionRepetitionsClient", "ListExpressionTraces", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListExpressionTracesSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "logic.WorkflowRunActionRepetitionsClient", "ListExpressionTraces", resp, "Failure sending request")
            return
            }

            result, err = client.ListExpressionTracesResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "logic.WorkflowRunActionRepetitionsClient", "ListExpressionTraces", resp, "Failure responding to request")
            }

    return
    }

    // ListExpressionTracesPreparer prepares the ListExpressionTraces request.
    func (client WorkflowRunActionRepetitionsClient) ListExpressionTracesPreparer(ctx context.Context, resourceGroupName string, workflowName string, runName string, actionName string, repetitionName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "actionName": autorest.Encode("path",actionName),
            "repetitionName": autorest.Encode("path",repetitionName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "runName": autorest.Encode("path",runName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            "workflowName": autorest.Encode("path",workflowName),
            }

                        const APIVersion = "2018-07-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsPost(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs/{runName}/actions/{actionName}/repetitions/{repetitionName}/listExpressionTraces",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListExpressionTracesSender sends the ListExpressionTraces request. The method will close the
    // http.Response Body if it receives an error.
    func (client WorkflowRunActionRepetitionsClient) ListExpressionTracesSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

// ListExpressionTracesResponder handles the response to the ListExpressionTraces request. The method always
// closes the http.Response Body.
func (client WorkflowRunActionRepetitionsClient) ListExpressionTracesResponder(resp *http.Response) (result ExpressionTraces, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

