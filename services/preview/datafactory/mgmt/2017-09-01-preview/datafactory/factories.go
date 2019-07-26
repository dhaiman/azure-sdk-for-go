package datafactory

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

// FactoriesClient is the the Azure Data Factory V2 management API provides a RESTful set of web services that interact
// with Azure Data Factory V2 services.
type FactoriesClient struct {
    BaseClient
}
// NewFactoriesClient creates an instance of the FactoriesClient client.
func NewFactoriesClient(subscriptionID string) FactoriesClient {
    return NewFactoriesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewFactoriesClientWithBaseURI creates an instance of the FactoriesClient client.
    func NewFactoriesClientWithBaseURI(baseURI string, subscriptionID string) FactoriesClient {
        return FactoriesClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// CancelPipelineRun cancel a pipeline run by its run ID.
    // Parameters:
        // resourceGroupName - the resource group name.
        // factoryName - the factory name.
        // runID - the pipeline run identifier.
func (client FactoriesClient) CancelPipelineRun(ctx context.Context, resourceGroupName string, factoryName string, runID string) (result autorest.Response, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/FactoriesClient.CancelPipelineRun")
        defer func() {
            sc := -1
            if result.Response != nil {
                sc = result.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: resourceGroupName,
             Constraints: []validation.Constraint{	{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil },
            	{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil },
            	{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil }}},
            { TargetValue: factoryName,
             Constraints: []validation.Constraint{	{Target: "factoryName", Name: validation.MaxLength, Rule: 63, Chain: nil },
            	{Target: "factoryName", Name: validation.MinLength, Rule: 3, Chain: nil },
            	{Target: "factoryName", Name: validation.Pattern, Rule: `^[A-Za-z0-9]+(?:-[A-Za-z0-9]+)*$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("datafactory.FactoriesClient", "CancelPipelineRun", err.Error())
            }

                req, err := client.CancelPipelineRunPreparer(ctx, resourceGroupName, factoryName, runID)
    if err != nil {
    err = autorest.NewErrorWithError(err, "datafactory.FactoriesClient", "CancelPipelineRun", nil , "Failure preparing request")
    return
    }

            resp, err := client.CancelPipelineRunSender(req)
            if err != nil {
            result.Response = resp
            err = autorest.NewErrorWithError(err, "datafactory.FactoriesClient", "CancelPipelineRun", resp, "Failure sending request")
            return
            }

            result, err = client.CancelPipelineRunResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "datafactory.FactoriesClient", "CancelPipelineRun", resp, "Failure responding to request")
            }

    return
    }

    // CancelPipelineRunPreparer prepares the CancelPipelineRun request.
    func (client FactoriesClient) CancelPipelineRunPreparer(ctx context.Context, resourceGroupName string, factoryName string, runID string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "factoryName": autorest.Encode("path",factoryName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "runId": autorest.Encode("path",runID),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2017-09-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsPost(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/cancelpipelinerun/{runId}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CancelPipelineRunSender sends the CancelPipelineRun request. The method will close the
    // http.Response Body if it receives an error.
    func (client FactoriesClient) CancelPipelineRunSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// CancelPipelineRunResponder handles the response to the CancelPipelineRun request. The method always
// closes the http.Response Body.
func (client FactoriesClient) CancelPipelineRunResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByClosing())
    result.Response = resp
        return
    }

// ConfigureFactoryRepo updates a factory's repo information.
    // Parameters:
        // locationID - the location identifier.
        // factoryRepoUpdate - update factory repo request definition.
func (client FactoriesClient) ConfigureFactoryRepo(ctx context.Context, locationID string, factoryRepoUpdate FactoryRepoUpdate) (result Factory, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/FactoriesClient.ConfigureFactoryRepo")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.ConfigureFactoryRepoPreparer(ctx, locationID, factoryRepoUpdate)
    if err != nil {
    err = autorest.NewErrorWithError(err, "datafactory.FactoriesClient", "ConfigureFactoryRepo", nil , "Failure preparing request")
    return
    }

            resp, err := client.ConfigureFactoryRepoSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "datafactory.FactoriesClient", "ConfigureFactoryRepo", resp, "Failure sending request")
            return
            }

            result, err = client.ConfigureFactoryRepoResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "datafactory.FactoriesClient", "ConfigureFactoryRepo", resp, "Failure responding to request")
            }

    return
    }

    // ConfigureFactoryRepoPreparer prepares the ConfigureFactoryRepo request.
    func (client FactoriesClient) ConfigureFactoryRepoPreparer(ctx context.Context, locationID string, factoryRepoUpdate FactoryRepoUpdate) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "locationId": autorest.Encode("path",locationID),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2017-09-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPost(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.DataFactory/locations/{locationId}/configureFactoryRepo",pathParameters),
    autorest.WithJSON(factoryRepoUpdate),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ConfigureFactoryRepoSender sends the ConfigureFactoryRepo request. The method will close the
    // http.Response Body if it receives an error.
    func (client FactoriesClient) ConfigureFactoryRepoSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// ConfigureFactoryRepoResponder handles the response to the ConfigureFactoryRepo request. The method always
// closes the http.Response Body.
func (client FactoriesClient) ConfigureFactoryRepoResponder(resp *http.Response) (result Factory, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// CreateOrUpdate creates or updates a factory.
    // Parameters:
        // resourceGroupName - the resource group name.
        // factoryName - the factory name.
        // factory - factory resource definition.
func (client FactoriesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, factoryName string, factory Factory) (result Factory, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/FactoriesClient.CreateOrUpdate")
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
            { TargetValue: factoryName,
             Constraints: []validation.Constraint{	{Target: "factoryName", Name: validation.MaxLength, Rule: 63, Chain: nil },
            	{Target: "factoryName", Name: validation.MinLength, Rule: 3, Chain: nil },
            	{Target: "factoryName", Name: validation.Pattern, Rule: `^[A-Za-z0-9]+(?:-[A-Za-z0-9]+)*$`, Chain: nil }}},
            { TargetValue: factory,
             Constraints: []validation.Constraint{	{Target: "factory.Identity", Name: validation.Null, Rule: false ,
            Chain: []validation.Constraint{	{Target: "factory.Identity.Type", Name: validation.Null, Rule: true, Chain: nil },
            }}}}}); err != nil {
            return result, validation.NewError("datafactory.FactoriesClient", "CreateOrUpdate", err.Error())
            }

                req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, factoryName, factory)
    if err != nil {
    err = autorest.NewErrorWithError(err, "datafactory.FactoriesClient", "CreateOrUpdate", nil , "Failure preparing request")
    return
    }

            resp, err := client.CreateOrUpdateSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "datafactory.FactoriesClient", "CreateOrUpdate", resp, "Failure sending request")
            return
            }

            result, err = client.CreateOrUpdateResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "datafactory.FactoriesClient", "CreateOrUpdate", resp, "Failure responding to request")
            }

    return
    }

    // CreateOrUpdatePreparer prepares the CreateOrUpdate request.
    func (client FactoriesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, factoryName string, factory Factory) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "factoryName": autorest.Encode("path",factoryName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2017-09-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPut(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}",pathParameters),
    autorest.WithJSON(factory),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
    // http.Response Body if it receives an error.
    func (client FactoriesClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client FactoriesClient) CreateOrUpdateResponder(resp *http.Response) (result Factory, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// Delete deletes a factory.
    // Parameters:
        // resourceGroupName - the resource group name.
        // factoryName - the factory name.
func (client FactoriesClient) Delete(ctx context.Context, resourceGroupName string, factoryName string) (result autorest.Response, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/FactoriesClient.Delete")
        defer func() {
            sc := -1
            if result.Response != nil {
                sc = result.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: resourceGroupName,
             Constraints: []validation.Constraint{	{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil },
            	{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil },
            	{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil }}},
            { TargetValue: factoryName,
             Constraints: []validation.Constraint{	{Target: "factoryName", Name: validation.MaxLength, Rule: 63, Chain: nil },
            	{Target: "factoryName", Name: validation.MinLength, Rule: 3, Chain: nil },
            	{Target: "factoryName", Name: validation.Pattern, Rule: `^[A-Za-z0-9]+(?:-[A-Za-z0-9]+)*$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("datafactory.FactoriesClient", "Delete", err.Error())
            }

                req, err := client.DeletePreparer(ctx, resourceGroupName, factoryName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "datafactory.FactoriesClient", "Delete", nil , "Failure preparing request")
    return
    }

            resp, err := client.DeleteSender(req)
            if err != nil {
            result.Response = resp
            err = autorest.NewErrorWithError(err, "datafactory.FactoriesClient", "Delete", resp, "Failure sending request")
            return
            }

            result, err = client.DeleteResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "datafactory.FactoriesClient", "Delete", resp, "Failure responding to request")
            }

    return
    }

    // DeletePreparer prepares the Delete request.
    func (client FactoriesClient) DeletePreparer(ctx context.Context, resourceGroupName string, factoryName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "factoryName": autorest.Encode("path",factoryName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2017-09-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsDelete(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // DeleteSender sends the Delete request. The method will close the
    // http.Response Body if it receives an error.
    func (client FactoriesClient) DeleteSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client FactoriesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusNoContent),
    autorest.ByClosing())
    result.Response = resp
        return
    }

// Get gets a factory.
    // Parameters:
        // resourceGroupName - the resource group name.
        // factoryName - the factory name.
func (client FactoriesClient) Get(ctx context.Context, resourceGroupName string, factoryName string) (result Factory, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/FactoriesClient.Get")
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
            { TargetValue: factoryName,
             Constraints: []validation.Constraint{	{Target: "factoryName", Name: validation.MaxLength, Rule: 63, Chain: nil },
            	{Target: "factoryName", Name: validation.MinLength, Rule: 3, Chain: nil },
            	{Target: "factoryName", Name: validation.Pattern, Rule: `^[A-Za-z0-9]+(?:-[A-Za-z0-9]+)*$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("datafactory.FactoriesClient", "Get", err.Error())
            }

                req, err := client.GetPreparer(ctx, resourceGroupName, factoryName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "datafactory.FactoriesClient", "Get", nil , "Failure preparing request")
    return
    }

            resp, err := client.GetSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "datafactory.FactoriesClient", "Get", resp, "Failure sending request")
            return
            }

            result, err = client.GetResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "datafactory.FactoriesClient", "Get", resp, "Failure responding to request")
            }

    return
    }

    // GetPreparer prepares the Get request.
    func (client FactoriesClient) GetPreparer(ctx context.Context, resourceGroupName string, factoryName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "factoryName": autorest.Encode("path",factoryName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2017-09-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client FactoriesClient) GetSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client FactoriesClient) GetResponder(resp *http.Response) (result Factory, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// List lists factories under the specified subscription.
func (client FactoriesClient) List(ctx context.Context) (result FactoryListResponsePage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/FactoriesClient.List")
        defer func() {
            sc := -1
            if result.flr.Response.Response != nil {
                sc = result.flr.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
                result.fn = client.listNextResults
    req, err := client.ListPreparer(ctx)
    if err != nil {
    err = autorest.NewErrorWithError(err, "datafactory.FactoriesClient", "List", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListSender(req)
            if err != nil {
            result.flr.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "datafactory.FactoriesClient", "List", resp, "Failure sending request")
            return
            }

            result.flr, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "datafactory.FactoriesClient", "List", resp, "Failure responding to request")
            }

    return
    }

    // ListPreparer prepares the List request.
    func (client FactoriesClient) ListPreparer(ctx context.Context) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2017-09-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.DataFactory/factories",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListSender sends the List request. The method will close the
    // http.Response Body if it receives an error.
    func (client FactoriesClient) ListSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client FactoriesClient) ListResponder(resp *http.Response) (result FactoryListResponse, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

            // listNextResults retrieves the next set of results, if any.
            func (client FactoriesClient) listNextResults(ctx context.Context, lastResults FactoryListResponse) (result FactoryListResponse, err error) {
            req, err := lastResults.factoryListResponsePreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "datafactory.FactoriesClient", "listNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "datafactory.FactoriesClient", "listNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "datafactory.FactoriesClient", "listNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

    // ListComplete enumerates all values, automatically crossing page boundaries as required.
    func (client FactoriesClient) ListComplete(ctx context.Context) (result FactoryListResponseIterator, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/FactoriesClient.List")
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

// ListByResourceGroup lists factories.
    // Parameters:
        // resourceGroupName - the resource group name.
func (client FactoriesClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (result FactoryListResponsePage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/FactoriesClient.ListByResourceGroup")
        defer func() {
            sc := -1
            if result.flr.Response.Response != nil {
                sc = result.flr.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: resourceGroupName,
             Constraints: []validation.Constraint{	{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil },
            	{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil },
            	{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("datafactory.FactoriesClient", "ListByResourceGroup", err.Error())
            }

                        result.fn = client.listByResourceGroupNextResults
    req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "datafactory.FactoriesClient", "ListByResourceGroup", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListByResourceGroupSender(req)
            if err != nil {
            result.flr.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "datafactory.FactoriesClient", "ListByResourceGroup", resp, "Failure sending request")
            return
            }

            result.flr, err = client.ListByResourceGroupResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "datafactory.FactoriesClient", "ListByResourceGroup", resp, "Failure responding to request")
            }

    return
    }

    // ListByResourceGroupPreparer prepares the ListByResourceGroup request.
    func (client FactoriesClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2017-09-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
    // http.Response Body if it receives an error.
    func (client FactoriesClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client FactoriesClient) ListByResourceGroupResponder(resp *http.Response) (result FactoryListResponse, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

            // listByResourceGroupNextResults retrieves the next set of results, if any.
            func (client FactoriesClient) listByResourceGroupNextResults(ctx context.Context, lastResults FactoryListResponse) (result FactoryListResponse, err error) {
            req, err := lastResults.factoryListResponsePreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "datafactory.FactoriesClient", "listByResourceGroupNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListByResourceGroupSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "datafactory.FactoriesClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListByResourceGroupResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "datafactory.FactoriesClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

    // ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
    func (client FactoriesClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result FactoryListResponseIterator, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/FactoriesClient.ListByResourceGroup")
            defer func() {
                sc := -1
                if result.Response().Response.Response != nil {
                    sc = result.page.Response().Response.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
     }
        result.page, err = client.ListByResourceGroup(ctx, resourceGroupName)
                return
        }

// Update updates a factory.
    // Parameters:
        // resourceGroupName - the resource group name.
        // factoryName - the factory name.
        // factoryUpdateParameters - the parameters for updating a factory.
func (client FactoriesClient) Update(ctx context.Context, resourceGroupName string, factoryName string, factoryUpdateParameters FactoryUpdateParameters) (result Factory, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/FactoriesClient.Update")
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
            { TargetValue: factoryName,
             Constraints: []validation.Constraint{	{Target: "factoryName", Name: validation.MaxLength, Rule: 63, Chain: nil },
            	{Target: "factoryName", Name: validation.MinLength, Rule: 3, Chain: nil },
            	{Target: "factoryName", Name: validation.Pattern, Rule: `^[A-Za-z0-9]+(?:-[A-Za-z0-9]+)*$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("datafactory.FactoriesClient", "Update", err.Error())
            }

                req, err := client.UpdatePreparer(ctx, resourceGroupName, factoryName, factoryUpdateParameters)
    if err != nil {
    err = autorest.NewErrorWithError(err, "datafactory.FactoriesClient", "Update", nil , "Failure preparing request")
    return
    }

            resp, err := client.UpdateSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "datafactory.FactoriesClient", "Update", resp, "Failure sending request")
            return
            }

            result, err = client.UpdateResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "datafactory.FactoriesClient", "Update", resp, "Failure responding to request")
            }

    return
    }

    // UpdatePreparer prepares the Update request.
    func (client FactoriesClient) UpdatePreparer(ctx context.Context, resourceGroupName string, factoryName string, factoryUpdateParameters FactoryUpdateParameters) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "factoryName": autorest.Encode("path",factoryName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2017-09-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPatch(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}",pathParameters),
    autorest.WithJSON(factoryUpdateParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // UpdateSender sends the Update request. The method will close the
    // http.Response Body if it receives an error.
    func (client FactoriesClient) UpdateSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client FactoriesClient) UpdateResponder(resp *http.Response) (result Factory, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

