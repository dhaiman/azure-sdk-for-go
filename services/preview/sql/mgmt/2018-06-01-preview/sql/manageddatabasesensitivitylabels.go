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

// ManagedDatabaseSensitivityLabelsClient is the the Azure SQL Database management API provides a RESTful set of web
// services that interact with Azure SQL Database services to manage your databases. The API enables you to create,
// retrieve, update, and delete databases.
type ManagedDatabaseSensitivityLabelsClient struct {
	BaseClient
}

// NewManagedDatabaseSensitivityLabelsClient creates an instance of the ManagedDatabaseSensitivityLabelsClient client.
func NewManagedDatabaseSensitivityLabelsClient(subscriptionID string) ManagedDatabaseSensitivityLabelsClient {
	return NewManagedDatabaseSensitivityLabelsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewManagedDatabaseSensitivityLabelsClientWithBaseURI creates an instance of the
// ManagedDatabaseSensitivityLabelsClient client using a custom endpoint.  Use this when interacting with an Azure
// cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewManagedDatabaseSensitivityLabelsClientWithBaseURI(baseURI string, subscriptionID string) ManagedDatabaseSensitivityLabelsClient {
	return ManagedDatabaseSensitivityLabelsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates the sensitivity label of a given column
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// managedInstanceName - the name of the managed instance.
// databaseName - the name of the database.
// schemaName - the name of the schema.
// tableName - the name of the table.
// columnName - the name of the column.
// parameters - the column sensitivity label resource.
func (client ManagedDatabaseSensitivityLabelsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, schemaName string, tableName string, columnName string, parameters SensitivityLabel) (result SensitivityLabel, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ManagedDatabaseSensitivityLabelsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, managedInstanceName, databaseName, schemaName, tableName, columnName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedDatabaseSensitivityLabelsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.ManagedDatabaseSensitivityLabelsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedDatabaseSensitivityLabelsClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ManagedDatabaseSensitivityLabelsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, schemaName string, tableName string, columnName string, parameters SensitivityLabel) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"columnName":             autorest.Encode("path", columnName),
		"databaseName":           autorest.Encode("path", databaseName),
		"managedInstanceName":    autorest.Encode("path", managedInstanceName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"schemaName":             autorest.Encode("path", schemaName),
		"sensitivityLabelSource": autorest.Encode("path", "current"),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
		"tableName":              autorest.Encode("path", tableName),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/databases/{databaseName}/schemas/{schemaName}/tables/{tableName}/columns/{columnName}/sensitivityLabels/{sensitivityLabelSource}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ManagedDatabaseSensitivityLabelsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ManagedDatabaseSensitivityLabelsClient) CreateOrUpdateResponder(resp *http.Response) (result SensitivityLabel, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the sensitivity label of a given column
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// managedInstanceName - the name of the managed instance.
// databaseName - the name of the database.
// schemaName - the name of the schema.
// tableName - the name of the table.
// columnName - the name of the column.
func (client ManagedDatabaseSensitivityLabelsClient) Delete(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, schemaName string, tableName string, columnName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ManagedDatabaseSensitivityLabelsClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, managedInstanceName, databaseName, schemaName, tableName, columnName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedDatabaseSensitivityLabelsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "sql.ManagedDatabaseSensitivityLabelsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedDatabaseSensitivityLabelsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ManagedDatabaseSensitivityLabelsClient) DeletePreparer(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, schemaName string, tableName string, columnName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"columnName":             autorest.Encode("path", columnName),
		"databaseName":           autorest.Encode("path", databaseName),
		"managedInstanceName":    autorest.Encode("path", managedInstanceName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"schemaName":             autorest.Encode("path", schemaName),
		"sensitivityLabelSource": autorest.Encode("path", "current"),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
		"tableName":              autorest.Encode("path", tableName),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/databases/{databaseName}/schemas/{schemaName}/tables/{tableName}/columns/{columnName}/sensitivityLabels/{sensitivityLabelSource}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ManagedDatabaseSensitivityLabelsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ManagedDatabaseSensitivityLabelsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// DisableRecommendation disables sensitivity recommendations on a given column
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// managedInstanceName - the name of the managed instance.
// databaseName - the name of the database.
// schemaName - the name of the schema.
// tableName - the name of the table.
// columnName - the name of the column.
func (client ManagedDatabaseSensitivityLabelsClient) DisableRecommendation(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, schemaName string, tableName string, columnName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ManagedDatabaseSensitivityLabelsClient.DisableRecommendation")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DisableRecommendationPreparer(ctx, resourceGroupName, managedInstanceName, databaseName, schemaName, tableName, columnName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedDatabaseSensitivityLabelsClient", "DisableRecommendation", nil, "Failure preparing request")
		return
	}

	resp, err := client.DisableRecommendationSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "sql.ManagedDatabaseSensitivityLabelsClient", "DisableRecommendation", resp, "Failure sending request")
		return
	}

	result, err = client.DisableRecommendationResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedDatabaseSensitivityLabelsClient", "DisableRecommendation", resp, "Failure responding to request")
	}

	return
}

// DisableRecommendationPreparer prepares the DisableRecommendation request.
func (client ManagedDatabaseSensitivityLabelsClient) DisableRecommendationPreparer(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, schemaName string, tableName string, columnName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"columnName":             autorest.Encode("path", columnName),
		"databaseName":           autorest.Encode("path", databaseName),
		"managedInstanceName":    autorest.Encode("path", managedInstanceName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"schemaName":             autorest.Encode("path", schemaName),
		"sensitivityLabelSource": autorest.Encode("path", "recommended"),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
		"tableName":              autorest.Encode("path", tableName),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/databases/{databaseName}/schemas/{schemaName}/tables/{tableName}/columns/{columnName}/sensitivityLabels/{sensitivityLabelSource}/disable", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DisableRecommendationSender sends the DisableRecommendation request. The method will close the
// http.Response Body if it receives an error.
func (client ManagedDatabaseSensitivityLabelsClient) DisableRecommendationSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DisableRecommendationResponder handles the response to the DisableRecommendation request. The method always
// closes the http.Response Body.
func (client ManagedDatabaseSensitivityLabelsClient) DisableRecommendationResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// EnableRecommendation enables sensitivity recommendations on a given column (recommendations are enabled by default
// on all columns)
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// managedInstanceName - the name of the managed instance.
// databaseName - the name of the database.
// schemaName - the name of the schema.
// tableName - the name of the table.
// columnName - the name of the column.
func (client ManagedDatabaseSensitivityLabelsClient) EnableRecommendation(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, schemaName string, tableName string, columnName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ManagedDatabaseSensitivityLabelsClient.EnableRecommendation")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.EnableRecommendationPreparer(ctx, resourceGroupName, managedInstanceName, databaseName, schemaName, tableName, columnName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedDatabaseSensitivityLabelsClient", "EnableRecommendation", nil, "Failure preparing request")
		return
	}

	resp, err := client.EnableRecommendationSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "sql.ManagedDatabaseSensitivityLabelsClient", "EnableRecommendation", resp, "Failure sending request")
		return
	}

	result, err = client.EnableRecommendationResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedDatabaseSensitivityLabelsClient", "EnableRecommendation", resp, "Failure responding to request")
	}

	return
}

// EnableRecommendationPreparer prepares the EnableRecommendation request.
func (client ManagedDatabaseSensitivityLabelsClient) EnableRecommendationPreparer(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, schemaName string, tableName string, columnName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"columnName":             autorest.Encode("path", columnName),
		"databaseName":           autorest.Encode("path", databaseName),
		"managedInstanceName":    autorest.Encode("path", managedInstanceName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"schemaName":             autorest.Encode("path", schemaName),
		"sensitivityLabelSource": autorest.Encode("path", "recommended"),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
		"tableName":              autorest.Encode("path", tableName),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/databases/{databaseName}/schemas/{schemaName}/tables/{tableName}/columns/{columnName}/sensitivityLabels/{sensitivityLabelSource}/enable", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// EnableRecommendationSender sends the EnableRecommendation request. The method will close the
// http.Response Body if it receives an error.
func (client ManagedDatabaseSensitivityLabelsClient) EnableRecommendationSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// EnableRecommendationResponder handles the response to the EnableRecommendation request. The method always
// closes the http.Response Body.
func (client ManagedDatabaseSensitivityLabelsClient) EnableRecommendationResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the sensitivity label of a given column
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// managedInstanceName - the name of the managed instance.
// databaseName - the name of the database.
// schemaName - the name of the schema.
// tableName - the name of the table.
// columnName - the name of the column.
// sensitivityLabelSource - the source of the sensitivity label.
func (client ManagedDatabaseSensitivityLabelsClient) Get(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, schemaName string, tableName string, columnName string, sensitivityLabelSource SensitivityLabelSource) (result SensitivityLabel, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ManagedDatabaseSensitivityLabelsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, managedInstanceName, databaseName, schemaName, tableName, columnName, sensitivityLabelSource)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedDatabaseSensitivityLabelsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.ManagedDatabaseSensitivityLabelsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedDatabaseSensitivityLabelsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ManagedDatabaseSensitivityLabelsClient) GetPreparer(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, schemaName string, tableName string, columnName string, sensitivityLabelSource SensitivityLabelSource) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"columnName":             autorest.Encode("path", columnName),
		"databaseName":           autorest.Encode("path", databaseName),
		"managedInstanceName":    autorest.Encode("path", managedInstanceName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"schemaName":             autorest.Encode("path", schemaName),
		"sensitivityLabelSource": autorest.Encode("path", sensitivityLabelSource),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
		"tableName":              autorest.Encode("path", tableName),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/databases/{databaseName}/schemas/{schemaName}/tables/{tableName}/columns/{columnName}/sensitivityLabels/{sensitivityLabelSource}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ManagedDatabaseSensitivityLabelsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ManagedDatabaseSensitivityLabelsClient) GetResponder(resp *http.Response) (result SensitivityLabel, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListCurrentByDatabase gets the sensitivity labels of a given database
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// managedInstanceName - the name of the managed instance.
// databaseName - the name of the database.
// filter - an OData filter expression that filters elements in the collection.
func (client ManagedDatabaseSensitivityLabelsClient) ListCurrentByDatabase(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, filter string) (result SensitivityLabelListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ManagedDatabaseSensitivityLabelsClient.ListCurrentByDatabase")
		defer func() {
			sc := -1
			if result.sllr.Response.Response != nil {
				sc = result.sllr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listCurrentByDatabaseNextResults
	req, err := client.ListCurrentByDatabasePreparer(ctx, resourceGroupName, managedInstanceName, databaseName, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedDatabaseSensitivityLabelsClient", "ListCurrentByDatabase", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListCurrentByDatabaseSender(req)
	if err != nil {
		result.sllr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.ManagedDatabaseSensitivityLabelsClient", "ListCurrentByDatabase", resp, "Failure sending request")
		return
	}

	result.sllr, err = client.ListCurrentByDatabaseResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedDatabaseSensitivityLabelsClient", "ListCurrentByDatabase", resp, "Failure responding to request")
	}

	return
}

// ListCurrentByDatabasePreparer prepares the ListCurrentByDatabase request.
func (client ManagedDatabaseSensitivityLabelsClient) ListCurrentByDatabasePreparer(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"databaseName":        autorest.Encode("path", databaseName),
		"managedInstanceName": autorest.Encode("path", managedInstanceName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/databases/{databaseName}/currentSensitivityLabels", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListCurrentByDatabaseSender sends the ListCurrentByDatabase request. The method will close the
// http.Response Body if it receives an error.
func (client ManagedDatabaseSensitivityLabelsClient) ListCurrentByDatabaseSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListCurrentByDatabaseResponder handles the response to the ListCurrentByDatabase request. The method always
// closes the http.Response Body.
func (client ManagedDatabaseSensitivityLabelsClient) ListCurrentByDatabaseResponder(resp *http.Response) (result SensitivityLabelListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listCurrentByDatabaseNextResults retrieves the next set of results, if any.
func (client ManagedDatabaseSensitivityLabelsClient) listCurrentByDatabaseNextResults(ctx context.Context, lastResults SensitivityLabelListResult) (result SensitivityLabelListResult, err error) {
	req, err := lastResults.sensitivityLabelListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "sql.ManagedDatabaseSensitivityLabelsClient", "listCurrentByDatabaseNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListCurrentByDatabaseSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "sql.ManagedDatabaseSensitivityLabelsClient", "listCurrentByDatabaseNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListCurrentByDatabaseResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedDatabaseSensitivityLabelsClient", "listCurrentByDatabaseNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListCurrentByDatabaseComplete enumerates all values, automatically crossing page boundaries as required.
func (client ManagedDatabaseSensitivityLabelsClient) ListCurrentByDatabaseComplete(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, filter string) (result SensitivityLabelListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ManagedDatabaseSensitivityLabelsClient.ListCurrentByDatabase")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListCurrentByDatabase(ctx, resourceGroupName, managedInstanceName, databaseName, filter)
	return
}

// ListRecommendedByDatabase gets the sensitivity labels of a given database
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// managedInstanceName - the name of the managed instance.
// databaseName - the name of the database.
// includeDisabledRecommendations - specifies whether to include disabled recommendations or not.
// filter - an OData filter expression that filters elements in the collection.
func (client ManagedDatabaseSensitivityLabelsClient) ListRecommendedByDatabase(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, includeDisabledRecommendations *bool, skipToken string, filter string) (result SensitivityLabelListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ManagedDatabaseSensitivityLabelsClient.ListRecommendedByDatabase")
		defer func() {
			sc := -1
			if result.sllr.Response.Response != nil {
				sc = result.sllr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listRecommendedByDatabaseNextResults
	req, err := client.ListRecommendedByDatabasePreparer(ctx, resourceGroupName, managedInstanceName, databaseName, includeDisabledRecommendations, skipToken, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedDatabaseSensitivityLabelsClient", "ListRecommendedByDatabase", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListRecommendedByDatabaseSender(req)
	if err != nil {
		result.sllr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.ManagedDatabaseSensitivityLabelsClient", "ListRecommendedByDatabase", resp, "Failure sending request")
		return
	}

	result.sllr, err = client.ListRecommendedByDatabaseResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedDatabaseSensitivityLabelsClient", "ListRecommendedByDatabase", resp, "Failure responding to request")
	}

	return
}

// ListRecommendedByDatabasePreparer prepares the ListRecommendedByDatabase request.
func (client ManagedDatabaseSensitivityLabelsClient) ListRecommendedByDatabasePreparer(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, includeDisabledRecommendations *bool, skipToken string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"databaseName":        autorest.Encode("path", databaseName),
		"managedInstanceName": autorest.Encode("path", managedInstanceName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if includeDisabledRecommendations != nil {
		queryParameters["includeDisabledRecommendations"] = autorest.Encode("query", *includeDisabledRecommendations)
	}
	if len(skipToken) > 0 {
		queryParameters["$skipToken"] = autorest.Encode("query", skipToken)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/databases/{databaseName}/recommendedSensitivityLabels", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListRecommendedByDatabaseSender sends the ListRecommendedByDatabase request. The method will close the
// http.Response Body if it receives an error.
func (client ManagedDatabaseSensitivityLabelsClient) ListRecommendedByDatabaseSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListRecommendedByDatabaseResponder handles the response to the ListRecommendedByDatabase request. The method always
// closes the http.Response Body.
func (client ManagedDatabaseSensitivityLabelsClient) ListRecommendedByDatabaseResponder(resp *http.Response) (result SensitivityLabelListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listRecommendedByDatabaseNextResults retrieves the next set of results, if any.
func (client ManagedDatabaseSensitivityLabelsClient) listRecommendedByDatabaseNextResults(ctx context.Context, lastResults SensitivityLabelListResult) (result SensitivityLabelListResult, err error) {
	req, err := lastResults.sensitivityLabelListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "sql.ManagedDatabaseSensitivityLabelsClient", "listRecommendedByDatabaseNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListRecommendedByDatabaseSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "sql.ManagedDatabaseSensitivityLabelsClient", "listRecommendedByDatabaseNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListRecommendedByDatabaseResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedDatabaseSensitivityLabelsClient", "listRecommendedByDatabaseNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListRecommendedByDatabaseComplete enumerates all values, automatically crossing page boundaries as required.
func (client ManagedDatabaseSensitivityLabelsClient) ListRecommendedByDatabaseComplete(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, includeDisabledRecommendations *bool, skipToken string, filter string) (result SensitivityLabelListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ManagedDatabaseSensitivityLabelsClient.ListRecommendedByDatabase")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListRecommendedByDatabase(ctx, resourceGroupName, managedInstanceName, databaseName, includeDisabledRecommendations, skipToken, filter)
	return
}
