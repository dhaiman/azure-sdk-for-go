package devops

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

// PipelineTemplateDefinitionsClient is the azure DevOps Resource Provider
type PipelineTemplateDefinitionsClient struct {
	BaseClient
}

// NewPipelineTemplateDefinitionsClient creates an instance of the PipelineTemplateDefinitionsClient client.
func NewPipelineTemplateDefinitionsClient(subscriptionID string) PipelineTemplateDefinitionsClient {
	return NewPipelineTemplateDefinitionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewPipelineTemplateDefinitionsClientWithBaseURI creates an instance of the PipelineTemplateDefinitionsClient client
// using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
func NewPipelineTemplateDefinitionsClientWithBaseURI(baseURI string, subscriptionID string) PipelineTemplateDefinitionsClient {
	return PipelineTemplateDefinitionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List lists all pipeline templates which can be used to configure a Pipeline.
func (client PipelineTemplateDefinitionsClient) List(ctx context.Context) (result PipelineTemplateDefinitionListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PipelineTemplateDefinitionsClient.List")
		defer func() {
			sc := -1
			if result.ptdlr.Response.Response != nil {
				sc = result.ptdlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devops.PipelineTemplateDefinitionsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.ptdlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "devops.PipelineTemplateDefinitionsClient", "List", resp, "Failure sending request")
		return
	}

	result.ptdlr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devops.PipelineTemplateDefinitionsClient", "List", resp, "Failure responding to request")
	}
	if result.ptdlr.hasNextLink() && result.ptdlr.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListPreparer prepares the List request.
func (client PipelineTemplateDefinitionsClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	const APIVersion = "2020-07-13-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.DevOps/pipelineTemplateDefinitions"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client PipelineTemplateDefinitionsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client PipelineTemplateDefinitionsClient) ListResponder(resp *http.Response) (result PipelineTemplateDefinitionListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client PipelineTemplateDefinitionsClient) listNextResults(ctx context.Context, lastResults PipelineTemplateDefinitionListResult) (result PipelineTemplateDefinitionListResult, err error) {
	req, err := lastResults.pipelineTemplateDefinitionListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "devops.PipelineTemplateDefinitionsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "devops.PipelineTemplateDefinitionsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devops.PipelineTemplateDefinitionsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client PipelineTemplateDefinitionsClient) ListComplete(ctx context.Context) (result PipelineTemplateDefinitionListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PipelineTemplateDefinitionsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx)
	return
}
