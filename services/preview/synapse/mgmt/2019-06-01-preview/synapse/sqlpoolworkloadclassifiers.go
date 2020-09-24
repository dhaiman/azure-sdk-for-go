package synapse

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

// SQLPoolWorkloadClassifiersClient is the azure Synapse Analytics Management Client
type SQLPoolWorkloadClassifiersClient struct {
	BaseClient
}

// NewSQLPoolWorkloadClassifiersClient creates an instance of the SQLPoolWorkloadClassifiersClient client.
func NewSQLPoolWorkloadClassifiersClient(subscriptionID string) SQLPoolWorkloadClassifiersClient {
	return NewSQLPoolWorkloadClassifiersClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewSQLPoolWorkloadClassifiersClientWithBaseURI creates an instance of the SQLPoolWorkloadClassifiersClient client
// using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
func NewSQLPoolWorkloadClassifiersClientWithBaseURI(baseURI string, subscriptionID string) SQLPoolWorkloadClassifiersClient {
	return SQLPoolWorkloadClassifiersClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a workload classifier.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace
// SQLPoolName - SQL pool name
// workloadGroupName - the name of the workload group from which to receive the classifier from.
// workloadClassifierName - the name of the workload classifier to create/update.
// parameters - the properties of the workload classifier.
func (client SQLPoolWorkloadClassifiersClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, workloadGroupName string, workloadClassifierName string, parameters WorkloadClassifier) (result SQLPoolWorkloadClassifiersCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SQLPoolWorkloadClassifiersClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.WorkloadClassifierProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.WorkloadClassifierProperties.MemberName", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("synapse.SQLPoolWorkloadClassifiersClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, workspaceName, SQLPoolName, workloadGroupName, workloadClassifierName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.SQLPoolWorkloadClassifiersClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.SQLPoolWorkloadClassifiersClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client SQLPoolWorkloadClassifiersClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, workloadGroupName string, workloadClassifierName string, parameters WorkloadClassifier) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"sqlPoolName":            autorest.Encode("path", SQLPoolName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
		"workloadClassifierName": autorest.Encode("path", workloadClassifierName),
		"workloadGroupName":      autorest.Encode("path", workloadGroupName),
		"workspaceName":          autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2019-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/workloadGroups/{workloadGroupName}/workloadClassifiers/{workloadClassifierName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client SQLPoolWorkloadClassifiersClient) CreateOrUpdateSender(req *http.Request) (future SQLPoolWorkloadClassifiersCreateOrUpdateFuture, err error) {
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
func (client SQLPoolWorkloadClassifiersClient) CreateOrUpdateResponder(resp *http.Response) (result WorkloadClassifier, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a workload classifier.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace
// SQLPoolName - SQL pool name
// workloadGroupName - the name of the workload group from which to receive the classifier from.
// workloadClassifierName - the name of the workload classifier to delete.
func (client SQLPoolWorkloadClassifiersClient) Delete(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, workloadGroupName string, workloadClassifierName string) (result SQLPoolWorkloadClassifiersDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SQLPoolWorkloadClassifiersClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("synapse.SQLPoolWorkloadClassifiersClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, workspaceName, SQLPoolName, workloadGroupName, workloadClassifierName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.SQLPoolWorkloadClassifiersClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.SQLPoolWorkloadClassifiersClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client SQLPoolWorkloadClassifiersClient) DeletePreparer(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, workloadGroupName string, workloadClassifierName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"sqlPoolName":            autorest.Encode("path", SQLPoolName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
		"workloadClassifierName": autorest.Encode("path", workloadClassifierName),
		"workloadGroupName":      autorest.Encode("path", workloadGroupName),
		"workspaceName":          autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2019-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/workloadGroups/{workloadGroupName}/workloadClassifiers/{workloadClassifierName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client SQLPoolWorkloadClassifiersClient) DeleteSender(req *http.Request) (future SQLPoolWorkloadClassifiersDeleteFuture, err error) {
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
func (client SQLPoolWorkloadClassifiersClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets a workload classifier
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace
// SQLPoolName - SQL pool name
// workloadGroupName - the name of the workload group from which to receive the classifier from.
// workloadClassifierName - the name of the workload classifier.
func (client SQLPoolWorkloadClassifiersClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, workloadGroupName string, workloadClassifierName string) (result WorkloadClassifier, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SQLPoolWorkloadClassifiersClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("synapse.SQLPoolWorkloadClassifiersClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, workspaceName, SQLPoolName, workloadGroupName, workloadClassifierName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.SQLPoolWorkloadClassifiersClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "synapse.SQLPoolWorkloadClassifiersClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.SQLPoolWorkloadClassifiersClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client SQLPoolWorkloadClassifiersClient) GetPreparer(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, workloadGroupName string, workloadClassifierName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"sqlPoolName":            autorest.Encode("path", SQLPoolName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
		"workloadClassifierName": autorest.Encode("path", workloadClassifierName),
		"workloadGroupName":      autorest.Encode("path", workloadGroupName),
		"workspaceName":          autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2019-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/workloadGroups/{workloadGroupName}/workloadClassifiers/{workloadClassifierName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client SQLPoolWorkloadClassifiersClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client SQLPoolWorkloadClassifiersClient) GetResponder(resp *http.Response) (result WorkloadClassifier, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByWorkloadGroup gets the list of workload classifiers for a workload group
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace
// SQLPoolName - SQL pool name
// workloadGroupName - the name of the workload group from which to receive the classifiers from.
func (client SQLPoolWorkloadClassifiersClient) ListByWorkloadGroup(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, workloadGroupName string) (result WorkloadClassifierListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SQLPoolWorkloadClassifiersClient.ListByWorkloadGroup")
		defer func() {
			sc := -1
			if result.wclr.Response.Response != nil {
				sc = result.wclr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("synapse.SQLPoolWorkloadClassifiersClient", "ListByWorkloadGroup", err.Error())
	}

	result.fn = client.listByWorkloadGroupNextResults
	req, err := client.ListByWorkloadGroupPreparer(ctx, resourceGroupName, workspaceName, SQLPoolName, workloadGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.SQLPoolWorkloadClassifiersClient", "ListByWorkloadGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByWorkloadGroupSender(req)
	if err != nil {
		result.wclr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "synapse.SQLPoolWorkloadClassifiersClient", "ListByWorkloadGroup", resp, "Failure sending request")
		return
	}

	result.wclr, err = client.ListByWorkloadGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.SQLPoolWorkloadClassifiersClient", "ListByWorkloadGroup", resp, "Failure responding to request")
	}
	if result.wclr.hasNextLink() && result.wclr.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListByWorkloadGroupPreparer prepares the ListByWorkloadGroup request.
func (client SQLPoolWorkloadClassifiersClient) ListByWorkloadGroupPreparer(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, workloadGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"sqlPoolName":       autorest.Encode("path", SQLPoolName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workloadGroupName": autorest.Encode("path", workloadGroupName),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2019-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/workloadGroups/{workloadGroupName}/workloadClassifiers", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByWorkloadGroupSender sends the ListByWorkloadGroup request. The method will close the
// http.Response Body if it receives an error.
func (client SQLPoolWorkloadClassifiersClient) ListByWorkloadGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByWorkloadGroupResponder handles the response to the ListByWorkloadGroup request. The method always
// closes the http.Response Body.
func (client SQLPoolWorkloadClassifiersClient) ListByWorkloadGroupResponder(resp *http.Response) (result WorkloadClassifierListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByWorkloadGroupNextResults retrieves the next set of results, if any.
func (client SQLPoolWorkloadClassifiersClient) listByWorkloadGroupNextResults(ctx context.Context, lastResults WorkloadClassifierListResult) (result WorkloadClassifierListResult, err error) {
	req, err := lastResults.workloadClassifierListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "synapse.SQLPoolWorkloadClassifiersClient", "listByWorkloadGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByWorkloadGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "synapse.SQLPoolWorkloadClassifiersClient", "listByWorkloadGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByWorkloadGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.SQLPoolWorkloadClassifiersClient", "listByWorkloadGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByWorkloadGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client SQLPoolWorkloadClassifiersClient) ListByWorkloadGroupComplete(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, workloadGroupName string) (result WorkloadClassifierListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SQLPoolWorkloadClassifiersClient.ListByWorkloadGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByWorkloadGroup(ctx, resourceGroupName, workspaceName, SQLPoolName, workloadGroupName)
	return
}
