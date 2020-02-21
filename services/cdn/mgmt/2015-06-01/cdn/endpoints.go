package cdn

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

// EndpointsClient is the cdn Management Client
type EndpointsClient struct {
	BaseClient
}

// NewEndpointsClient creates an instance of the EndpointsClient client.
func NewEndpointsClient(subscriptionID string) EndpointsClient {
	return NewEndpointsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewEndpointsClientWithBaseURI creates an instance of the EndpointsClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewEndpointsClientWithBaseURI(baseURI string, subscriptionID string) EndpointsClient {
	return EndpointsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create sends the create request.
// Parameters:
// endpointName - name of the endpoint within the CDN profile.
// endpointProperties - endpoint properties
// profileName - name of the CDN profile within the resource group.
// resourceGroupName - name of the resource group within the Azure subscription.
func (client EndpointsClient) Create(ctx context.Context, endpointName string, endpointProperties EndpointCreateParameters, profileName string, resourceGroupName string) (result EndpointsCreateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EndpointsClient.Create")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: endpointProperties,
			Constraints: []validation.Constraint{{Target: "endpointProperties.Location", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "endpointProperties.EndpointPropertiesCreateParameters", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "endpointProperties.EndpointPropertiesCreateParameters.Origins", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("cdn.EndpointsClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, endpointName, endpointProperties, profileName, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.EndpointsClient", "Create", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.EndpointsClient", "Create", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client EndpointsClient) CreatePreparer(ctx context.Context, endpointName string, endpointProperties EndpointCreateParameters, profileName string, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"endpointName":      autorest.Encode("path", endpointName),
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}", pathParameters),
		autorest.WithJSON(endpointProperties),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client EndpointsClient) CreateSender(req *http.Request) (future EndpointsCreateFuture, err error) {
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
func (client EndpointsClient) CreateResponder(resp *http.Response) (result Endpoint, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// DeleteIfExists sends the delete if exists request.
// Parameters:
// endpointName - name of the endpoint within the CDN profile.
// profileName - name of the CDN profile within the resource group.
// resourceGroupName - name of the resource group within the Azure subscription.
func (client EndpointsClient) DeleteIfExists(ctx context.Context, endpointName string, profileName string, resourceGroupName string) (result EndpointsDeleteIfExistsFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EndpointsClient.DeleteIfExists")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeleteIfExistsPreparer(ctx, endpointName, profileName, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.EndpointsClient", "DeleteIfExists", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteIfExistsSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.EndpointsClient", "DeleteIfExists", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeleteIfExistsPreparer prepares the DeleteIfExists request.
func (client EndpointsClient) DeleteIfExistsPreparer(ctx context.Context, endpointName string, profileName string, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"endpointName":      autorest.Encode("path", endpointName),
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteIfExistsSender sends the DeleteIfExists request. The method will close the
// http.Response Body if it receives an error.
func (client EndpointsClient) DeleteIfExistsSender(req *http.Request) (future EndpointsDeleteIfExistsFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// DeleteIfExistsResponder handles the response to the DeleteIfExists request. The method always
// closes the http.Response Body.
func (client EndpointsClient) DeleteIfExistsResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get sends the get request.
// Parameters:
// endpointName - name of the endpoint within the CDN profile.
// profileName - name of the CDN profile within the resource group.
// resourceGroupName - name of the resource group within the Azure subscription.
func (client EndpointsClient) Get(ctx context.Context, endpointName string, profileName string, resourceGroupName string) (result Endpoint, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EndpointsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, endpointName, profileName, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.EndpointsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cdn.EndpointsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.EndpointsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client EndpointsClient) GetPreparer(ctx context.Context, endpointName string, profileName string, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"endpointName":      autorest.Encode("path", endpointName),
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client EndpointsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client EndpointsClient) GetResponder(resp *http.Response) (result Endpoint, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByProfile sends the list by profile request.
// Parameters:
// profileName - name of the CDN profile within the resource group.
// resourceGroupName - name of the resource group within the Azure subscription.
func (client EndpointsClient) ListByProfile(ctx context.Context, profileName string, resourceGroupName string) (result EndpointListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EndpointsClient.ListByProfile")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListByProfilePreparer(ctx, profileName, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.EndpointsClient", "ListByProfile", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByProfileSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cdn.EndpointsClient", "ListByProfile", resp, "Failure sending request")
		return
	}

	result, err = client.ListByProfileResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.EndpointsClient", "ListByProfile", resp, "Failure responding to request")
	}

	return
}

// ListByProfilePreparer prepares the ListByProfile request.
func (client EndpointsClient) ListByProfilePreparer(ctx context.Context, profileName string, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByProfileSender sends the ListByProfile request. The method will close the
// http.Response Body if it receives an error.
func (client EndpointsClient) ListByProfileSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByProfileResponder handles the response to the ListByProfile request. The method always
// closes the http.Response Body.
func (client EndpointsClient) ListByProfileResponder(resp *http.Response) (result EndpointListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// LoadContent sends the load content request.
// Parameters:
// endpointName - name of the endpoint within the CDN profile.
// contentFilePaths - the path to the content to be loaded. Path should describe a file.
// profileName - name of the CDN profile within the resource group.
// resourceGroupName - name of the resource group within the Azure subscription.
func (client EndpointsClient) LoadContent(ctx context.Context, endpointName string, contentFilePaths LoadParameters, profileName string, resourceGroupName string) (result EndpointsLoadContentFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EndpointsClient.LoadContent")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: contentFilePaths,
			Constraints: []validation.Constraint{{Target: "contentFilePaths.ContentPaths", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("cdn.EndpointsClient", "LoadContent", err.Error())
	}

	req, err := client.LoadContentPreparer(ctx, endpointName, contentFilePaths, profileName, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.EndpointsClient", "LoadContent", nil, "Failure preparing request")
		return
	}

	result, err = client.LoadContentSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.EndpointsClient", "LoadContent", result.Response(), "Failure sending request")
		return
	}

	return
}

// LoadContentPreparer prepares the LoadContent request.
func (client EndpointsClient) LoadContentPreparer(ctx context.Context, endpointName string, contentFilePaths LoadParameters, profileName string, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"endpointName":      autorest.Encode("path", endpointName),
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/load", pathParameters),
		autorest.WithJSON(contentFilePaths),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// LoadContentSender sends the LoadContent request. The method will close the
// http.Response Body if it receives an error.
func (client EndpointsClient) LoadContentSender(req *http.Request) (future EndpointsLoadContentFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// LoadContentResponder handles the response to the LoadContent request. The method always
// closes the http.Response Body.
func (client EndpointsClient) LoadContentResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// PurgeContent sends the purge content request.
// Parameters:
// endpointName - name of the endpoint within the CDN profile.
// contentFilePaths - the path to the content to be purged. Path can describe a file or directory.
// profileName - name of the CDN profile within the resource group.
// resourceGroupName - name of the resource group within the Azure subscription.
func (client EndpointsClient) PurgeContent(ctx context.Context, endpointName string, contentFilePaths PurgeParameters, profileName string, resourceGroupName string) (result EndpointsPurgeContentFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EndpointsClient.PurgeContent")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: contentFilePaths,
			Constraints: []validation.Constraint{{Target: "contentFilePaths.ContentPaths", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("cdn.EndpointsClient", "PurgeContent", err.Error())
	}

	req, err := client.PurgeContentPreparer(ctx, endpointName, contentFilePaths, profileName, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.EndpointsClient", "PurgeContent", nil, "Failure preparing request")
		return
	}

	result, err = client.PurgeContentSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.EndpointsClient", "PurgeContent", result.Response(), "Failure sending request")
		return
	}

	return
}

// PurgeContentPreparer prepares the PurgeContent request.
func (client EndpointsClient) PurgeContentPreparer(ctx context.Context, endpointName string, contentFilePaths PurgeParameters, profileName string, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"endpointName":      autorest.Encode("path", endpointName),
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/purge", pathParameters),
		autorest.WithJSON(contentFilePaths),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PurgeContentSender sends the PurgeContent request. The method will close the
// http.Response Body if it receives an error.
func (client EndpointsClient) PurgeContentSender(req *http.Request) (future EndpointsPurgeContentFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// PurgeContentResponder handles the response to the PurgeContent request. The method always
// closes the http.Response Body.
func (client EndpointsClient) PurgeContentResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Start sends the start request.
// Parameters:
// endpointName - name of the endpoint within the CDN profile.
// profileName - name of the CDN profile within the resource group.
// resourceGroupName - name of the resource group within the Azure subscription.
func (client EndpointsClient) Start(ctx context.Context, endpointName string, profileName string, resourceGroupName string) (result EndpointsStartFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EndpointsClient.Start")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.StartPreparer(ctx, endpointName, profileName, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.EndpointsClient", "Start", nil, "Failure preparing request")
		return
	}

	result, err = client.StartSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.EndpointsClient", "Start", result.Response(), "Failure sending request")
		return
	}

	return
}

// StartPreparer prepares the Start request.
func (client EndpointsClient) StartPreparer(ctx context.Context, endpointName string, profileName string, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"endpointName":      autorest.Encode("path", endpointName),
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/start", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// StartSender sends the Start request. The method will close the
// http.Response Body if it receives an error.
func (client EndpointsClient) StartSender(req *http.Request) (future EndpointsStartFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// StartResponder handles the response to the Start request. The method always
// closes the http.Response Body.
func (client EndpointsClient) StartResponder(resp *http.Response) (result Endpoint, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Stop sends the stop request.
// Parameters:
// endpointName - name of the endpoint within the CDN profile.
// profileName - name of the CDN profile within the resource group.
// resourceGroupName - name of the resource group within the Azure subscription.
func (client EndpointsClient) Stop(ctx context.Context, endpointName string, profileName string, resourceGroupName string) (result EndpointsStopFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EndpointsClient.Stop")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.StopPreparer(ctx, endpointName, profileName, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.EndpointsClient", "Stop", nil, "Failure preparing request")
		return
	}

	result, err = client.StopSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.EndpointsClient", "Stop", result.Response(), "Failure sending request")
		return
	}

	return
}

// StopPreparer prepares the Stop request.
func (client EndpointsClient) StopPreparer(ctx context.Context, endpointName string, profileName string, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"endpointName":      autorest.Encode("path", endpointName),
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/stop", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// StopSender sends the Stop request. The method will close the
// http.Response Body if it receives an error.
func (client EndpointsClient) StopSender(req *http.Request) (future EndpointsStopFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// StopResponder handles the response to the Stop request. The method always
// closes the http.Response Body.
func (client EndpointsClient) StopResponder(resp *http.Response) (result Endpoint, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update sends the update request.
// Parameters:
// endpointName - name of the endpoint within the CDN profile.
// endpointProperties - endpoint properties
// profileName - name of the CDN profile within the resource group.
// resourceGroupName - name of the resource group within the Azure subscription.
func (client EndpointsClient) Update(ctx context.Context, endpointName string, endpointProperties EndpointUpdateParameters, profileName string, resourceGroupName string) (result EndpointsUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EndpointsClient.Update")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, endpointName, endpointProperties, profileName, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.EndpointsClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.EndpointsClient", "Update", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client EndpointsClient) UpdatePreparer(ctx context.Context, endpointName string, endpointProperties EndpointUpdateParameters, profileName string, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"endpointName":      autorest.Encode("path", endpointName),
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}", pathParameters),
		autorest.WithJSON(endpointProperties),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client EndpointsClient) UpdateSender(req *http.Request) (future EndpointsUpdateFuture, err error) {
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
func (client EndpointsClient) UpdateResponder(resp *http.Response) (result Endpoint, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ValidateCustomDomain sends the validate custom domain request.
// Parameters:
// endpointName - name of the endpoint within the CDN profile.
// customDomainProperties - custom domain to validate.
// profileName - name of the CDN profile within the resource group.
// resourceGroupName - name of the resource group within the Azure subscription.
func (client EndpointsClient) ValidateCustomDomain(ctx context.Context, endpointName string, customDomainProperties ValidateCustomDomainInput, profileName string, resourceGroupName string) (result ValidateCustomDomainOutput, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EndpointsClient.ValidateCustomDomain")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: customDomainProperties,
			Constraints: []validation.Constraint{{Target: "customDomainProperties.HostName", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("cdn.EndpointsClient", "ValidateCustomDomain", err.Error())
	}

	req, err := client.ValidateCustomDomainPreparer(ctx, endpointName, customDomainProperties, profileName, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.EndpointsClient", "ValidateCustomDomain", nil, "Failure preparing request")
		return
	}

	resp, err := client.ValidateCustomDomainSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cdn.EndpointsClient", "ValidateCustomDomain", resp, "Failure sending request")
		return
	}

	result, err = client.ValidateCustomDomainResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.EndpointsClient", "ValidateCustomDomain", resp, "Failure responding to request")
	}

	return
}

// ValidateCustomDomainPreparer prepares the ValidateCustomDomain request.
func (client EndpointsClient) ValidateCustomDomainPreparer(ctx context.Context, endpointName string, customDomainProperties ValidateCustomDomainInput, profileName string, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"endpointName":      autorest.Encode("path", endpointName),
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/validateCustomDomain", pathParameters),
		autorest.WithJSON(customDomainProperties),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ValidateCustomDomainSender sends the ValidateCustomDomain request. The method will close the
// http.Response Body if it receives an error.
func (client EndpointsClient) ValidateCustomDomainSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ValidateCustomDomainResponder handles the response to the ValidateCustomDomain request. The method always
// closes the http.Response Body.
func (client EndpointsClient) ValidateCustomDomainResponder(resp *http.Response) (result ValidateCustomDomainOutput, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
