package machinelearningservices

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

// PrivateLinkResourcesClient is the these APIs allow end users to operate on Azure Machine Learning Workspace
// resources.
type PrivateLinkResourcesClient struct {
	BaseClient
}

// NewPrivateLinkResourcesClient creates an instance of the PrivateLinkResourcesClient client.
func NewPrivateLinkResourcesClient(subscriptionID string) PrivateLinkResourcesClient {
	return NewPrivateLinkResourcesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewPrivateLinkResourcesClientWithBaseURI creates an instance of the PrivateLinkResourcesClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewPrivateLinkResourcesClientWithBaseURI(baseURI string, subscriptionID string) PrivateLinkResourcesClient {
	return PrivateLinkResourcesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// ListByWorkspace gets the private link resources that need to be created for a workspace.
// Parameters:
// resourceGroupName - name of the resource group in which workspace is located.
// workspaceName - name of Azure Machine Learning workspace.
func (client PrivateLinkResourcesClient) ListByWorkspace(ctx context.Context, resourceGroupName string, workspaceName string) (result PrivateLinkResourceListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateLinkResourcesClient.ListByWorkspace")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListByWorkspacePreparer(ctx, resourceGroupName, workspaceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "machinelearningservices.PrivateLinkResourcesClient", "ListByWorkspace", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByWorkspaceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "machinelearningservices.PrivateLinkResourcesClient", "ListByWorkspace", resp, "Failure sending request")
		return
	}

	result, err = client.ListByWorkspaceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "machinelearningservices.PrivateLinkResourcesClient", "ListByWorkspace", resp, "Failure responding to request")
	}

	return
}

// ListByWorkspacePreparer prepares the ListByWorkspace request.
func (client PrivateLinkResourcesClient) ListByWorkspacePreparer(ctx context.Context, resourceGroupName string, workspaceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2020-02-18-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/privateLinkResources", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByWorkspaceSender sends the ListByWorkspace request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateLinkResourcesClient) ListByWorkspaceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByWorkspaceResponder handles the response to the ListByWorkspace request. The method always
// closes the http.Response Body.
func (client PrivateLinkResourcesClient) ListByWorkspaceResponder(resp *http.Response) (result PrivateLinkResourceListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
