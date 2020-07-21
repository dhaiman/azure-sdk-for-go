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

// SQLPoolSensitivityLabelsClient is the azure Synapse Analytics Management Client
type SQLPoolSensitivityLabelsClient struct {
	BaseClient
}

// NewSQLPoolSensitivityLabelsClient creates an instance of the SQLPoolSensitivityLabelsClient client.
func NewSQLPoolSensitivityLabelsClient(subscriptionID string) SQLPoolSensitivityLabelsClient {
	return NewSQLPoolSensitivityLabelsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewSQLPoolSensitivityLabelsClientWithBaseURI creates an instance of the SQLPoolSensitivityLabelsClient client using
// a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
func NewSQLPoolSensitivityLabelsClientWithBaseURI(baseURI string, subscriptionID string) SQLPoolSensitivityLabelsClient {
	return SQLPoolSensitivityLabelsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates the sensitivity label of a given column in a Sql pool
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace
// SQLPoolName - SQL pool name
// schemaName - the name of the schema.
// tableName - the name of the table.
// columnName - the name of the column.
// parameters - the column sensitivity label resource.
func (client SQLPoolSensitivityLabelsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, schemaName string, tableName string, columnName string, parameters SensitivityLabel) (result SensitivityLabel, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SQLPoolSensitivityLabelsClient.CreateOrUpdate")
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
		return result, validation.NewError("synapse.SQLPoolSensitivityLabelsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, workspaceName, SQLPoolName, schemaName, tableName, columnName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.SQLPoolSensitivityLabelsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "synapse.SQLPoolSensitivityLabelsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.SQLPoolSensitivityLabelsClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client SQLPoolSensitivityLabelsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, schemaName string, tableName string, columnName string, parameters SensitivityLabel) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"columnName":             autorest.Encode("path", columnName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"schemaName":             autorest.Encode("path", schemaName),
		"sensitivityLabelSource": autorest.Encode("path", "current"),
		"sqlPoolName":            autorest.Encode("path", SQLPoolName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
		"tableName":              autorest.Encode("path", tableName),
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
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/schemas/{schemaName}/tables/{tableName}/columns/{columnName}/sensitivityLabels/{sensitivityLabelSource}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client SQLPoolSensitivityLabelsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client SQLPoolSensitivityLabelsClient) CreateOrUpdateResponder(resp *http.Response) (result SensitivityLabel, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the sensitivity label of a given column in a Sql pool
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace
// SQLPoolName - SQL pool name
// schemaName - the name of the schema.
// tableName - the name of the table.
// columnName - the name of the column.
func (client SQLPoolSensitivityLabelsClient) Delete(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, schemaName string, tableName string, columnName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SQLPoolSensitivityLabelsClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
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
		return result, validation.NewError("synapse.SQLPoolSensitivityLabelsClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, workspaceName, SQLPoolName, schemaName, tableName, columnName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.SQLPoolSensitivityLabelsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "synapse.SQLPoolSensitivityLabelsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.SQLPoolSensitivityLabelsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client SQLPoolSensitivityLabelsClient) DeletePreparer(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, schemaName string, tableName string, columnName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"columnName":             autorest.Encode("path", columnName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"schemaName":             autorest.Encode("path", schemaName),
		"sensitivityLabelSource": autorest.Encode("path", "current"),
		"sqlPoolName":            autorest.Encode("path", SQLPoolName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
		"tableName":              autorest.Encode("path", tableName),
		"workspaceName":          autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2019-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/schemas/{schemaName}/tables/{tableName}/columns/{columnName}/sensitivityLabels/{sensitivityLabelSource}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client SQLPoolSensitivityLabelsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client SQLPoolSensitivityLabelsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// DisableRecommendation disables sensitivity recommendations on a given column
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace
// SQLPoolName - SQL pool name
// schemaName - the name of the schema.
// tableName - the name of the table.
// columnName - the name of the column.
func (client SQLPoolSensitivityLabelsClient) DisableRecommendation(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, schemaName string, tableName string, columnName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SQLPoolSensitivityLabelsClient.DisableRecommendation")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
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
		return result, validation.NewError("synapse.SQLPoolSensitivityLabelsClient", "DisableRecommendation", err.Error())
	}

	req, err := client.DisableRecommendationPreparer(ctx, resourceGroupName, workspaceName, SQLPoolName, schemaName, tableName, columnName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.SQLPoolSensitivityLabelsClient", "DisableRecommendation", nil, "Failure preparing request")
		return
	}

	resp, err := client.DisableRecommendationSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "synapse.SQLPoolSensitivityLabelsClient", "DisableRecommendation", resp, "Failure sending request")
		return
	}

	result, err = client.DisableRecommendationResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.SQLPoolSensitivityLabelsClient", "DisableRecommendation", resp, "Failure responding to request")
	}

	return
}

// DisableRecommendationPreparer prepares the DisableRecommendation request.
func (client SQLPoolSensitivityLabelsClient) DisableRecommendationPreparer(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, schemaName string, tableName string, columnName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"columnName":             autorest.Encode("path", columnName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"schemaName":             autorest.Encode("path", schemaName),
		"sensitivityLabelSource": autorest.Encode("path", "recommended"),
		"sqlPoolName":            autorest.Encode("path", SQLPoolName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
		"tableName":              autorest.Encode("path", tableName),
		"workspaceName":          autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2019-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/schemas/{schemaName}/tables/{tableName}/columns/{columnName}/sensitivityLabels/{sensitivityLabelSource}/disable", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DisableRecommendationSender sends the DisableRecommendation request. The method will close the
// http.Response Body if it receives an error.
func (client SQLPoolSensitivityLabelsClient) DisableRecommendationSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DisableRecommendationResponder handles the response to the DisableRecommendation request. The method always
// closes the http.Response Body.
func (client SQLPoolSensitivityLabelsClient) DisableRecommendationResponder(resp *http.Response) (result autorest.Response, err error) {
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
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace
// SQLPoolName - SQL pool name
// schemaName - the name of the schema.
// tableName - the name of the table.
// columnName - the name of the column.
func (client SQLPoolSensitivityLabelsClient) EnableRecommendation(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, schemaName string, tableName string, columnName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SQLPoolSensitivityLabelsClient.EnableRecommendation")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
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
		return result, validation.NewError("synapse.SQLPoolSensitivityLabelsClient", "EnableRecommendation", err.Error())
	}

	req, err := client.EnableRecommendationPreparer(ctx, resourceGroupName, workspaceName, SQLPoolName, schemaName, tableName, columnName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.SQLPoolSensitivityLabelsClient", "EnableRecommendation", nil, "Failure preparing request")
		return
	}

	resp, err := client.EnableRecommendationSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "synapse.SQLPoolSensitivityLabelsClient", "EnableRecommendation", resp, "Failure sending request")
		return
	}

	result, err = client.EnableRecommendationResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.SQLPoolSensitivityLabelsClient", "EnableRecommendation", resp, "Failure responding to request")
	}

	return
}

// EnableRecommendationPreparer prepares the EnableRecommendation request.
func (client SQLPoolSensitivityLabelsClient) EnableRecommendationPreparer(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, schemaName string, tableName string, columnName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"columnName":             autorest.Encode("path", columnName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"schemaName":             autorest.Encode("path", schemaName),
		"sensitivityLabelSource": autorest.Encode("path", "recommended"),
		"sqlPoolName":            autorest.Encode("path", SQLPoolName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
		"tableName":              autorest.Encode("path", tableName),
		"workspaceName":          autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2019-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/schemas/{schemaName}/tables/{tableName}/columns/{columnName}/sensitivityLabels/{sensitivityLabelSource}/enable", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// EnableRecommendationSender sends the EnableRecommendation request. The method will close the
// http.Response Body if it receives an error.
func (client SQLPoolSensitivityLabelsClient) EnableRecommendationSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// EnableRecommendationResponder handles the response to the EnableRecommendation request. The method always
// closes the http.Response Body.
func (client SQLPoolSensitivityLabelsClient) EnableRecommendationResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// ListCurrent gets SQL pool sensitivity labels.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace
// SQLPoolName - SQL pool name
// filter - an OData filter expression that filters elements in the collection.
func (client SQLPoolSensitivityLabelsClient) ListCurrent(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, filter string) (result SensitivityLabelListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SQLPoolSensitivityLabelsClient.ListCurrent")
		defer func() {
			sc := -1
			if result.sllr.Response.Response != nil {
				sc = result.sllr.Response.Response.StatusCode
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
		return result, validation.NewError("synapse.SQLPoolSensitivityLabelsClient", "ListCurrent", err.Error())
	}

	result.fn = client.listCurrentNextResults
	req, err := client.ListCurrentPreparer(ctx, resourceGroupName, workspaceName, SQLPoolName, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.SQLPoolSensitivityLabelsClient", "ListCurrent", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListCurrentSender(req)
	if err != nil {
		result.sllr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "synapse.SQLPoolSensitivityLabelsClient", "ListCurrent", resp, "Failure sending request")
		return
	}

	result.sllr, err = client.ListCurrentResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.SQLPoolSensitivityLabelsClient", "ListCurrent", resp, "Failure responding to request")
	}

	return
}

// ListCurrentPreparer prepares the ListCurrent request.
func (client SQLPoolSensitivityLabelsClient) ListCurrentPreparer(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"sqlPoolName":       autorest.Encode("path", SQLPoolName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2019-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/currentSensitivityLabels", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListCurrentSender sends the ListCurrent request. The method will close the
// http.Response Body if it receives an error.
func (client SQLPoolSensitivityLabelsClient) ListCurrentSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListCurrentResponder handles the response to the ListCurrent request. The method always
// closes the http.Response Body.
func (client SQLPoolSensitivityLabelsClient) ListCurrentResponder(resp *http.Response) (result SensitivityLabelListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listCurrentNextResults retrieves the next set of results, if any.
func (client SQLPoolSensitivityLabelsClient) listCurrentNextResults(ctx context.Context, lastResults SensitivityLabelListResult) (result SensitivityLabelListResult, err error) {
	req, err := lastResults.sensitivityLabelListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "synapse.SQLPoolSensitivityLabelsClient", "listCurrentNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListCurrentSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "synapse.SQLPoolSensitivityLabelsClient", "listCurrentNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListCurrentResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.SQLPoolSensitivityLabelsClient", "listCurrentNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListCurrentComplete enumerates all values, automatically crossing page boundaries as required.
func (client SQLPoolSensitivityLabelsClient) ListCurrentComplete(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, filter string) (result SensitivityLabelListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SQLPoolSensitivityLabelsClient.ListCurrent")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListCurrent(ctx, resourceGroupName, workspaceName, SQLPoolName, filter)
	return
}

// ListRecommended gets sensitivity labels of a given SQL pool.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace
// SQLPoolName - SQL pool name
// includeDisabledRecommendations - specifies whether to include disabled recommendations or not.
// skipToken - an OData query option to indicate how many elements to skip in the collection.
// filter - an OData filter expression that filters elements in the collection.
func (client SQLPoolSensitivityLabelsClient) ListRecommended(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, includeDisabledRecommendations *bool, skipToken string, filter string) (result SensitivityLabelListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SQLPoolSensitivityLabelsClient.ListRecommended")
		defer func() {
			sc := -1
			if result.sllr.Response.Response != nil {
				sc = result.sllr.Response.Response.StatusCode
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
		return result, validation.NewError("synapse.SQLPoolSensitivityLabelsClient", "ListRecommended", err.Error())
	}

	result.fn = client.listRecommendedNextResults
	req, err := client.ListRecommendedPreparer(ctx, resourceGroupName, workspaceName, SQLPoolName, includeDisabledRecommendations, skipToken, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.SQLPoolSensitivityLabelsClient", "ListRecommended", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListRecommendedSender(req)
	if err != nil {
		result.sllr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "synapse.SQLPoolSensitivityLabelsClient", "ListRecommended", resp, "Failure sending request")
		return
	}

	result.sllr, err = client.ListRecommendedResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.SQLPoolSensitivityLabelsClient", "ListRecommended", resp, "Failure responding to request")
	}

	return
}

// ListRecommendedPreparer prepares the ListRecommended request.
func (client SQLPoolSensitivityLabelsClient) ListRecommendedPreparer(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, includeDisabledRecommendations *bool, skipToken string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"sqlPoolName":       autorest.Encode("path", SQLPoolName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2019-06-01-preview"
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
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/recommendedSensitivityLabels", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListRecommendedSender sends the ListRecommended request. The method will close the
// http.Response Body if it receives an error.
func (client SQLPoolSensitivityLabelsClient) ListRecommendedSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListRecommendedResponder handles the response to the ListRecommended request. The method always
// closes the http.Response Body.
func (client SQLPoolSensitivityLabelsClient) ListRecommendedResponder(resp *http.Response) (result SensitivityLabelListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listRecommendedNextResults retrieves the next set of results, if any.
func (client SQLPoolSensitivityLabelsClient) listRecommendedNextResults(ctx context.Context, lastResults SensitivityLabelListResult) (result SensitivityLabelListResult, err error) {
	req, err := lastResults.sensitivityLabelListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "synapse.SQLPoolSensitivityLabelsClient", "listRecommendedNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListRecommendedSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "synapse.SQLPoolSensitivityLabelsClient", "listRecommendedNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListRecommendedResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.SQLPoolSensitivityLabelsClient", "listRecommendedNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListRecommendedComplete enumerates all values, automatically crossing page boundaries as required.
func (client SQLPoolSensitivityLabelsClient) ListRecommendedComplete(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, includeDisabledRecommendations *bool, skipToken string, filter string) (result SensitivityLabelListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SQLPoolSensitivityLabelsClient.ListRecommended")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListRecommended(ctx, resourceGroupName, workspaceName, SQLPoolName, includeDisabledRecommendations, skipToken, filter)
	return
}
