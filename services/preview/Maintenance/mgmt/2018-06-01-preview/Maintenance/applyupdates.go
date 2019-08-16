package maintenance

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

// ApplyUpdatesClient is the azure Maintenance Management Client
type ApplyUpdatesClient struct {
	BaseClient
}

// NewApplyUpdatesClient creates an instance of the ApplyUpdatesClient client.
func NewApplyUpdatesClient(subscriptionID string) ApplyUpdatesClient {
	return NewApplyUpdatesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewApplyUpdatesClientWithBaseURI creates an instance of the ApplyUpdatesClient client.
func NewApplyUpdatesClientWithBaseURI(baseURI string, subscriptionID string) ApplyUpdatesClient {
	return ApplyUpdatesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate apply maintenance updates to resource
// Parameters:
// resourceGroupName - resource group name
// providerName - resource provider name
// resourceType - resource type
// resourceName - resource identifier
func (client ApplyUpdatesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, providerName string, resourceType string, resourceName string) (result ApplyUpdate, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplyUpdatesClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, providerName, resourceType, resourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "maintenance.ApplyUpdatesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "maintenance.ApplyUpdatesClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "maintenance.ApplyUpdatesClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ApplyUpdatesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, providerName string, resourceType string, resourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"providerName":      autorest.Encode("path", providerName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"resourceType":      autorest.Encode("path", resourceType),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{providerName}/{resourceType}/{resourceName}/providers/Microsoft.Maintenance/applyUpdates/default", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ApplyUpdatesClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ApplyUpdatesClient) CreateOrUpdateResponder(resp *http.Response) (result ApplyUpdate, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CreateOrUpdateParent apply maintenance updates to resource with parent
// Parameters:
// resourceGroupName - resource group name
// resourceParentType - resource parent type
// resourceParentName - resource parent identifier
// providerName - resource provider name
// resourceType - resource type
// resourceName - resource identifier
func (client ApplyUpdatesClient) CreateOrUpdateParent(ctx context.Context, resourceGroupName string, resourceParentType string, resourceParentName string, providerName string, resourceType string, resourceName string) (result ApplyUpdate, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplyUpdatesClient.CreateOrUpdateParent")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdateParentPreparer(ctx, resourceGroupName, resourceParentType, resourceParentName, providerName, resourceType, resourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "maintenance.ApplyUpdatesClient", "CreateOrUpdateParent", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateParentSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "maintenance.ApplyUpdatesClient", "CreateOrUpdateParent", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateParentResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "maintenance.ApplyUpdatesClient", "CreateOrUpdateParent", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdateParentPreparer prepares the CreateOrUpdateParent request.
func (client ApplyUpdatesClient) CreateOrUpdateParentPreparer(ctx context.Context, resourceGroupName string, resourceParentType string, resourceParentName string, providerName string, resourceType string, resourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"providerName":       autorest.Encode("path", providerName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"resourceName":       autorest.Encode("path", resourceName),
		"resourceParentName": autorest.Encode("path", resourceParentName),
		"resourceParentType": autorest.Encode("path", resourceParentType),
		"resourceType":       autorest.Encode("path", resourceType),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{providerName}/{resourceParentType}/{resourceParentName}/{resourceType}/{resourceName}/providers/Microsoft.Maintenance/applyUpdates/default", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateParentSender sends the CreateOrUpdateParent request. The method will close the
// http.Response Body if it receives an error.
func (client ApplyUpdatesClient) CreateOrUpdateParentSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// CreateOrUpdateParentResponder handles the response to the CreateOrUpdateParent request. The method always
// closes the http.Response Body.
func (client ApplyUpdatesClient) CreateOrUpdateParentResponder(resp *http.Response) (result ApplyUpdate, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get track maintenance updates to resource
// Parameters:
// resourceGroupName - resource group name
// providerName - resource provider name
// resourceType - resource type
// resourceName - resource identifier
// applyUpdateName - applyUpdate Id
func (client ApplyUpdatesClient) Get(ctx context.Context, resourceGroupName string, providerName string, resourceType string, resourceName string, applyUpdateName string) (result ApplyUpdate, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplyUpdatesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, providerName, resourceType, resourceName, applyUpdateName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "maintenance.ApplyUpdatesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "maintenance.ApplyUpdatesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "maintenance.ApplyUpdatesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ApplyUpdatesClient) GetPreparer(ctx context.Context, resourceGroupName string, providerName string, resourceType string, resourceName string, applyUpdateName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applyUpdateName":   autorest.Encode("path", applyUpdateName),
		"providerName":      autorest.Encode("path", providerName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"resourceType":      autorest.Encode("path", resourceType),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{providerName}/{resourceType}/{resourceName}/providers/Microsoft.Maintenance/applyUpdates/{applyUpdateName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ApplyUpdatesClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ApplyUpdatesClient) GetResponder(resp *http.Response) (result ApplyUpdate, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetParent track maintenance updates to resource with parent
// Parameters:
// resourceGroupName - resource group name
// resourceParentType - resource parent type
// resourceParentName - resource parent identifier
// providerName - resource provider name
// resourceType - resource type
// resourceName - resource identifier
// applyUpdateName - applyUpdate Id
func (client ApplyUpdatesClient) GetParent(ctx context.Context, resourceGroupName string, resourceParentType string, resourceParentName string, providerName string, resourceType string, resourceName string, applyUpdateName string) (result ApplyUpdate, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplyUpdatesClient.GetParent")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetParentPreparer(ctx, resourceGroupName, resourceParentType, resourceParentName, providerName, resourceType, resourceName, applyUpdateName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "maintenance.ApplyUpdatesClient", "GetParent", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetParentSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "maintenance.ApplyUpdatesClient", "GetParent", resp, "Failure sending request")
		return
	}

	result, err = client.GetParentResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "maintenance.ApplyUpdatesClient", "GetParent", resp, "Failure responding to request")
	}

	return
}

// GetParentPreparer prepares the GetParent request.
func (client ApplyUpdatesClient) GetParentPreparer(ctx context.Context, resourceGroupName string, resourceParentType string, resourceParentName string, providerName string, resourceType string, resourceName string, applyUpdateName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applyUpdateName":    autorest.Encode("path", applyUpdateName),
		"providerName":       autorest.Encode("path", providerName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"resourceName":       autorest.Encode("path", resourceName),
		"resourceParentName": autorest.Encode("path", resourceParentName),
		"resourceParentType": autorest.Encode("path", resourceParentType),
		"resourceType":       autorest.Encode("path", resourceType),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{providerName}/{resourceParentType}/{resourceParentName}/{resourceType}/{resourceName}/providers/Microsoft.Maintenance/applyUpdates/{applyUpdateName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetParentSender sends the GetParent request. The method will close the
// http.Response Body if it receives an error.
func (client ApplyUpdatesClient) GetParentSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetParentResponder handles the response to the GetParent request. The method always
// closes the http.Response Body.
func (client ApplyUpdatesClient) GetParentResponder(resp *http.Response) (result ApplyUpdate, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
