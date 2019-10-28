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
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient is the the Azure SQL Database management API
// provides a RESTful set of web services that interact with Azure SQL Database services to manage your databases. The
// API enables you to create, retrieve, update, and delete databases.
type ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient struct {
	BaseClient
}

// NewManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient creates an instance of the
// ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient client.
func NewManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient(subscriptionID string) ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient {
	return NewManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClientWithBaseURI creates an instance of the
// ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient client.
func NewManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClientWithBaseURI(baseURI string, subscriptionID string) ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient {
	return ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate sets a database's long term retention policy.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// managedInstanceName - the name of the managed instance.
// parameters - the long term retention policy info.
func (client ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, managedInstanceName string, restorableDroppedDatabaseID string, parameters ManagedBackupShortTermRetentionPolicy) (result ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, managedInstanceName, restorableDroppedDatabaseID, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, managedInstanceName string, restorableDroppedDatabaseID string, parameters ManagedBackupShortTermRetentionPolicy) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"managedInstanceName":         autorest.Encode("path", managedInstanceName),
		"policyName":                  autorest.Encode("path", "default"),
		"resourceGroupName":           autorest.Encode("path", resourceGroupName),
		"restorableDroppedDatabaseId": autorest.Encode("path", restorableDroppedDatabaseID),
		"subscriptionId":              autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/restorableDroppedDatabases/{restorableDroppedDatabaseId}/backupShortTermRetentionPolicies/{policyName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient) CreateOrUpdateSender(req *http.Request) (future ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesCreateOrUpdateFuture, err error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req, sd...)
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient) CreateOrUpdateResponder(resp *http.Response) (result ManagedBackupShortTermRetentionPolicy, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get gets a dropped database's short term retention policy.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// managedInstanceName - the name of the managed instance.
func (client ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient) Get(ctx context.Context, resourceGroupName string, managedInstanceName string, restorableDroppedDatabaseID string) (result ManagedBackupShortTermRetentionPolicy, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, managedInstanceName, restorableDroppedDatabaseID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient) GetPreparer(ctx context.Context, resourceGroupName string, managedInstanceName string, restorableDroppedDatabaseID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"managedInstanceName":         autorest.Encode("path", managedInstanceName),
		"policyName":                  autorest.Encode("path", "default"),
		"resourceGroupName":           autorest.Encode("path", resourceGroupName),
		"restorableDroppedDatabaseId": autorest.Encode("path", restorableDroppedDatabaseID),
		"subscriptionId":              autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/restorableDroppedDatabases/{restorableDroppedDatabaseId}/backupShortTermRetentionPolicies/{policyName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient) GetResponder(resp *http.Response) (result ManagedBackupShortTermRetentionPolicy, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByRestorableDroppedDatabase gets a dropped database's short term retention policy list.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// managedInstanceName - the name of the managed instance.
func (client ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient) ListByRestorableDroppedDatabase(ctx context.Context, resourceGroupName string, managedInstanceName string, restorableDroppedDatabaseID string) (result ManagedBackupShortTermRetentionPolicyListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient.ListByRestorableDroppedDatabase")
		defer func() {
			sc := -1
			if result.mbstrplr.Response.Response != nil {
				sc = result.mbstrplr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByRestorableDroppedDatabaseNextResults
	req, err := client.ListByRestorableDroppedDatabasePreparer(ctx, resourceGroupName, managedInstanceName, restorableDroppedDatabaseID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient", "ListByRestorableDroppedDatabase", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByRestorableDroppedDatabaseSender(req)
	if err != nil {
		result.mbstrplr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient", "ListByRestorableDroppedDatabase", resp, "Failure sending request")
		return
	}

	result.mbstrplr, err = client.ListByRestorableDroppedDatabaseResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient", "ListByRestorableDroppedDatabase", resp, "Failure responding to request")
	}

	return
}

// ListByRestorableDroppedDatabasePreparer prepares the ListByRestorableDroppedDatabase request.
func (client ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient) ListByRestorableDroppedDatabasePreparer(ctx context.Context, resourceGroupName string, managedInstanceName string, restorableDroppedDatabaseID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"managedInstanceName":         autorest.Encode("path", managedInstanceName),
		"resourceGroupName":           autorest.Encode("path", resourceGroupName),
		"restorableDroppedDatabaseId": autorest.Encode("path", restorableDroppedDatabaseID),
		"subscriptionId":              autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/restorableDroppedDatabases/{restorableDroppedDatabaseId}/backupShortTermRetentionPolicies", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByRestorableDroppedDatabaseSender sends the ListByRestorableDroppedDatabase request. The method will close the
// http.Response Body if it receives an error.
func (client ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient) ListByRestorableDroppedDatabaseSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListByRestorableDroppedDatabaseResponder handles the response to the ListByRestorableDroppedDatabase request. The method always
// closes the http.Response Body.
func (client ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient) ListByRestorableDroppedDatabaseResponder(resp *http.Response) (result ManagedBackupShortTermRetentionPolicyListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByRestorableDroppedDatabaseNextResults retrieves the next set of results, if any.
func (client ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient) listByRestorableDroppedDatabaseNextResults(ctx context.Context, lastResults ManagedBackupShortTermRetentionPolicyListResult) (result ManagedBackupShortTermRetentionPolicyListResult, err error) {
	req, err := lastResults.managedBackupShortTermRetentionPolicyListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "sql.ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient", "listByRestorableDroppedDatabaseNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByRestorableDroppedDatabaseSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "sql.ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient", "listByRestorableDroppedDatabaseNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByRestorableDroppedDatabaseResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient", "listByRestorableDroppedDatabaseNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByRestorableDroppedDatabaseComplete enumerates all values, automatically crossing page boundaries as required.
func (client ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient) ListByRestorableDroppedDatabaseComplete(ctx context.Context, resourceGroupName string, managedInstanceName string, restorableDroppedDatabaseID string) (result ManagedBackupShortTermRetentionPolicyListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient.ListByRestorableDroppedDatabase")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByRestorableDroppedDatabase(ctx, resourceGroupName, managedInstanceName, restorableDroppedDatabaseID)
	return
}

// Update sets a database's long term retention policy.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// managedInstanceName - the name of the managed instance.
// parameters - the long term retention policy info.
func (client ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient) Update(ctx context.Context, resourceGroupName string, managedInstanceName string, restorableDroppedDatabaseID string, parameters ManagedBackupShortTermRetentionPolicy) (result ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient.Update")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, resourceGroupName, managedInstanceName, restorableDroppedDatabaseID, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient", "Update", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient) UpdatePreparer(ctx context.Context, resourceGroupName string, managedInstanceName string, restorableDroppedDatabaseID string, parameters ManagedBackupShortTermRetentionPolicy) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"managedInstanceName":         autorest.Encode("path", managedInstanceName),
		"policyName":                  autorest.Encode("path", "default"),
		"resourceGroupName":           autorest.Encode("path", resourceGroupName),
		"restorableDroppedDatabaseId": autorest.Encode("path", restorableDroppedDatabaseID),
		"subscriptionId":              autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/restorableDroppedDatabases/{restorableDroppedDatabaseId}/backupShortTermRetentionPolicies/{policyName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient) UpdateSender(req *http.Request) (future ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesUpdateFuture, err error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req, sd...)
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient) UpdateResponder(resp *http.Response) (result ManagedBackupShortTermRetentionPolicy, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
