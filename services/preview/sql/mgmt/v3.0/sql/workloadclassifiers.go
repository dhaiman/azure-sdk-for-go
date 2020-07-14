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

// WorkloadClassifiersClient is the the Azure SQL Database management API provides a RESTful set of web services that
// interact with Azure SQL Database services to manage your databases. The API enables you to create, retrieve, update,
// and delete databases.
type WorkloadClassifiersClient struct {
	BaseClient
}

// NewWorkloadClassifiersClient creates an instance of the WorkloadClassifiersClient client.
func NewWorkloadClassifiersClient(subscriptionID string) WorkloadClassifiersClient {
	return NewWorkloadClassifiersClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWorkloadClassifiersClientWithBaseURI creates an instance of the WorkloadClassifiersClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewWorkloadClassifiersClientWithBaseURI(baseURI string, subscriptionID string) WorkloadClassifiersClient {
	return WorkloadClassifiersClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a workload classifier.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
// databaseName - the name of the database.
// workloadGroupName - the name of the workload group from which to receive the classifier from.
// workloadClassifierName - the name of the workload classifier to create/update.
// parameters - the properties of the workload classifier.
func (client WorkloadClassifiersClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, databaseName string, workloadGroupName string, workloadClassifierName string, parameters WorkloadClassifier) (result WorkloadClassifiersCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WorkloadClassifiersClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.WorkloadClassifierProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.WorkloadClassifierProperties.MemberName", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("sql.WorkloadClassifiersClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, serverName, databaseName, workloadGroupName, workloadClassifierName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.WorkloadClassifiersClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.WorkloadClassifiersClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client WorkloadClassifiersClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, serverName string, databaseName string, workloadGroupName string, workloadClassifierName string, parameters WorkloadClassifier) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"databaseName":           autorest.Encode("path", databaseName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"serverName":             autorest.Encode("path", serverName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
		"workloadClassifierName": autorest.Encode("path", workloadClassifierName),
		"workloadGroupName":      autorest.Encode("path", workloadGroupName),
	}

	const APIVersion = "2019-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/workloadGroups/{workloadGroupName}/workloadClassifiers/{workloadClassifierName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client WorkloadClassifiersClient) CreateOrUpdateSender(req *http.Request) (future WorkloadClassifiersCreateOrUpdateFuture, err error) {
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
func (client WorkloadClassifiersClient) CreateOrUpdateResponder(resp *http.Response) (result WorkloadClassifier, err error) {
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
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
// databaseName - the name of the database.
// workloadGroupName - the name of the workload group from which to receive the classifier from.
// workloadClassifierName - the name of the workload classifier to delete.
func (client WorkloadClassifiersClient) Delete(ctx context.Context, resourceGroupName string, serverName string, databaseName string, workloadGroupName string, workloadClassifierName string) (result WorkloadClassifiersDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WorkloadClassifiersClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, serverName, databaseName, workloadGroupName, workloadClassifierName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.WorkloadClassifiersClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.WorkloadClassifiersClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client WorkloadClassifiersClient) DeletePreparer(ctx context.Context, resourceGroupName string, serverName string, databaseName string, workloadGroupName string, workloadClassifierName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"databaseName":           autorest.Encode("path", databaseName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"serverName":             autorest.Encode("path", serverName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
		"workloadClassifierName": autorest.Encode("path", workloadClassifierName),
		"workloadGroupName":      autorest.Encode("path", workloadGroupName),
	}

	const APIVersion = "2019-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/workloadGroups/{workloadGroupName}/workloadClassifiers/{workloadClassifierName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client WorkloadClassifiersClient) DeleteSender(req *http.Request) (future WorkloadClassifiersDeleteFuture, err error) {
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
func (client WorkloadClassifiersClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets a workload classifier
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
// databaseName - the name of the database.
// workloadGroupName - the name of the workload group from which to receive the classifier from.
// workloadClassifierName - the name of the workload classifier.
func (client WorkloadClassifiersClient) Get(ctx context.Context, resourceGroupName string, serverName string, databaseName string, workloadGroupName string, workloadClassifierName string) (result WorkloadClassifier, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WorkloadClassifiersClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, serverName, databaseName, workloadGroupName, workloadClassifierName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.WorkloadClassifiersClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.WorkloadClassifiersClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.WorkloadClassifiersClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client WorkloadClassifiersClient) GetPreparer(ctx context.Context, resourceGroupName string, serverName string, databaseName string, workloadGroupName string, workloadClassifierName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"databaseName":           autorest.Encode("path", databaseName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"serverName":             autorest.Encode("path", serverName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
		"workloadClassifierName": autorest.Encode("path", workloadClassifierName),
		"workloadGroupName":      autorest.Encode("path", workloadGroupName),
	}

	const APIVersion = "2019-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/workloadGroups/{workloadGroupName}/workloadClassifiers/{workloadClassifierName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client WorkloadClassifiersClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client WorkloadClassifiersClient) GetResponder(resp *http.Response) (result WorkloadClassifier, err error) {
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
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
// databaseName - the name of the database.
// workloadGroupName - the name of the workload group from which to receive the classifiers from.
func (client WorkloadClassifiersClient) ListByWorkloadGroup(ctx context.Context, resourceGroupName string, serverName string, databaseName string, workloadGroupName string) (result WorkloadClassifierListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WorkloadClassifiersClient.ListByWorkloadGroup")
		defer func() {
			sc := -1
			if result.wclr.Response.Response != nil {
				sc = result.wclr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByWorkloadGroupNextResults
	req, err := client.ListByWorkloadGroupPreparer(ctx, resourceGroupName, serverName, databaseName, workloadGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.WorkloadClassifiersClient", "ListByWorkloadGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByWorkloadGroupSender(req)
	if err != nil {
		result.wclr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.WorkloadClassifiersClient", "ListByWorkloadGroup", resp, "Failure sending request")
		return
	}

	result.wclr, err = client.ListByWorkloadGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.WorkloadClassifiersClient", "ListByWorkloadGroup", resp, "Failure responding to request")
	}

	return
}

// ListByWorkloadGroupPreparer prepares the ListByWorkloadGroup request.
func (client WorkloadClassifiersClient) ListByWorkloadGroupPreparer(ctx context.Context, resourceGroupName string, serverName string, databaseName string, workloadGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"databaseName":      autorest.Encode("path", databaseName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workloadGroupName": autorest.Encode("path", workloadGroupName),
	}

	const APIVersion = "2019-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/workloadGroups/{workloadGroupName}/workloadClassifiers", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByWorkloadGroupSender sends the ListByWorkloadGroup request. The method will close the
// http.Response Body if it receives an error.
func (client WorkloadClassifiersClient) ListByWorkloadGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByWorkloadGroupResponder handles the response to the ListByWorkloadGroup request. The method always
// closes the http.Response Body.
func (client WorkloadClassifiersClient) ListByWorkloadGroupResponder(resp *http.Response) (result WorkloadClassifierListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByWorkloadGroupNextResults retrieves the next set of results, if any.
func (client WorkloadClassifiersClient) listByWorkloadGroupNextResults(ctx context.Context, lastResults WorkloadClassifierListResult) (result WorkloadClassifierListResult, err error) {
	req, err := lastResults.workloadClassifierListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "sql.WorkloadClassifiersClient", "listByWorkloadGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByWorkloadGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "sql.WorkloadClassifiersClient", "listByWorkloadGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByWorkloadGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.WorkloadClassifiersClient", "listByWorkloadGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByWorkloadGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client WorkloadClassifiersClient) ListByWorkloadGroupComplete(ctx context.Context, resourceGroupName string, serverName string, databaseName string, workloadGroupName string) (result WorkloadClassifierListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WorkloadClassifiersClient.ListByWorkloadGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByWorkloadGroup(ctx, resourceGroupName, serverName, databaseName, workloadGroupName)
	return
}
