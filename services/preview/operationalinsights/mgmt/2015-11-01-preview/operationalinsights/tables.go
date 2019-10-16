package operationalinsights

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
    "github.com/Azure/go-autorest/autorest/validation"
)

// TablesClient is the operational Insights Client
type TablesClient struct {
    BaseClient
}
// NewTablesClient creates an instance of the TablesClient client.
func NewTablesClient() TablesClient {
    return NewTablesClientWithBaseURI(DefaultBaseURI, )
}

// NewTablesClientWithBaseURI creates an instance of the TablesClient client.
    func NewTablesClientWithBaseURI(baseURI string, ) TablesClient {
        return TablesClient{ NewWithBaseURI(baseURI, )}
    }

// CreateOrUpdate sets a Log Analytics workspace table properties.
    // Parameters:
        // subscriptionID - gets subscription credentials which uniquely identify Microsoft Azure subscription. The
        // subscription ID forms part of the URI for every service call.
        // resourceGroupName - the name of the resource group to get. The name is case insensitive.
        // workspaceName - the name of the workspace.
        // tableName - the name of the table.
        // parameters - the parameters required to set table properties.
func (client TablesClient) CreateOrUpdate(ctx context.Context, subscriptionID string, resourceGroupName string, workspaceName string, tableName string, parameters Table) (result Table, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/TablesClient.CreateOrUpdate")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: resourceGroupName,
             Constraints: []validation.Constraint{	{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil },
            	{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil },
            	{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil }}},
            { TargetValue: workspaceName,
             Constraints: []validation.Constraint{	{Target: "workspaceName", Name: validation.MaxLength, Rule: 63, Chain: nil },
            	{Target: "workspaceName", Name: validation.MinLength, Rule: 4, Chain: nil },
            	{Target: "workspaceName", Name: validation.Pattern, Rule: `^[A-Za-z0-9][A-Za-z0-9-]+[A-Za-z0-9]$`, Chain: nil }}},
            { TargetValue: parameters,
             Constraints: []validation.Constraint{	{Target: "parameters.TableProperties", Name: validation.Null, Rule: false ,
            Chain: []validation.Constraint{	{Target: "parameters.TableProperties.RetentionInDays", Name: validation.Null, Rule: false ,
            Chain: []validation.Constraint{	{Target: "parameters.TableProperties.RetentionInDays", Name: validation.InclusiveMaximum, Rule: int64(730), Chain: nil },
            	{Target: "parameters.TableProperties.RetentionInDays", Name: validation.InclusiveMinimum, Rule: 30, Chain: nil },
            }},
            }}}}}); err != nil {
            return result, validation.NewError("operationalinsights.TablesClient", "CreateOrUpdate", err.Error())
            }

                req, err := client.CreateOrUpdatePreparer(ctx, subscriptionID, resourceGroupName, workspaceName, tableName, parameters)
    if err != nil {
    err = autorest.NewErrorWithError(err, "operationalinsights.TablesClient", "CreateOrUpdate", nil , "Failure preparing request")
    return
    }

            resp, err := client.CreateOrUpdateSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "operationalinsights.TablesClient", "CreateOrUpdate", resp, "Failure sending request")
            return
            }

            result, err = client.CreateOrUpdateResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "operationalinsights.TablesClient", "CreateOrUpdate", resp, "Failure responding to request")
            }

    return
    }

    // CreateOrUpdatePreparer prepares the CreateOrUpdate request.
    func (client TablesClient) CreateOrUpdatePreparer(ctx context.Context, subscriptionID string, resourceGroupName string, workspaceName string, tableName string, parameters Table) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",subscriptionID),
            "tableName": autorest.Encode("path",tableName),
            "workspaceName": autorest.Encode("path",workspaceName),
            }

                        const APIVersion = "2015-11-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

                    parameters.ID = nil
                parameters.Name = nil
    preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPut(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/tables/{tableName}",pathParameters),
    autorest.WithJSON(parameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
    // http.Response Body if it receives an error.
    func (client TablesClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client TablesClient) CreateOrUpdateResponder(resp *http.Response) (result Table, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// Get gets a Log Analytics workspace table.
    // Parameters:
        // subscriptionID - gets subscription credentials which uniquely identify Microsoft Azure subscription. The
        // subscription ID forms part of the URI for every service call.
        // resourceGroupName - the name of the resource group to get. The name is case insensitive.
        // workspaceName - the name of the workspace.
        // tableName - the name of the table.
func (client TablesClient) Get(ctx context.Context, subscriptionID string, resourceGroupName string, workspaceName string, tableName string) (result Table, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/TablesClient.Get")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: resourceGroupName,
             Constraints: []validation.Constraint{	{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil },
            	{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil },
            	{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil }}},
            { TargetValue: workspaceName,
             Constraints: []validation.Constraint{	{Target: "workspaceName", Name: validation.MaxLength, Rule: 63, Chain: nil },
            	{Target: "workspaceName", Name: validation.MinLength, Rule: 4, Chain: nil },
            	{Target: "workspaceName", Name: validation.Pattern, Rule: `^[A-Za-z0-9][A-Za-z0-9-]+[A-Za-z0-9]$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("operationalinsights.TablesClient", "Get", err.Error())
            }

                req, err := client.GetPreparer(ctx, subscriptionID, resourceGroupName, workspaceName, tableName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "operationalinsights.TablesClient", "Get", nil , "Failure preparing request")
    return
    }

            resp, err := client.GetSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "operationalinsights.TablesClient", "Get", resp, "Failure sending request")
            return
            }

            result, err = client.GetResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "operationalinsights.TablesClient", "Get", resp, "Failure responding to request")
            }

    return
    }

    // GetPreparer prepares the Get request.
    func (client TablesClient) GetPreparer(ctx context.Context, subscriptionID string, resourceGroupName string, workspaceName string, tableName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",subscriptionID),
            "tableName": autorest.Encode("path",tableName),
            "workspaceName": autorest.Encode("path",workspaceName),
            }

                        const APIVersion = "2015-11-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/tables/{tableName}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client TablesClient) GetSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client TablesClient) GetResponder(resp *http.Response) (result Table, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

