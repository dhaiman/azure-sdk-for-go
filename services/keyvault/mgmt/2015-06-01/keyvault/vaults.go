package keyvault

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

// VaultsClient is the the Azure management API provides a RESTful set of web services that interact with Azure Key
// Vault.
type VaultsClient struct {
    BaseClient
}
// NewVaultsClient creates an instance of the VaultsClient client.
func NewVaultsClient(subscriptionID string) VaultsClient {
    return NewVaultsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewVaultsClientWithBaseURI creates an instance of the VaultsClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
    func NewVaultsClientWithBaseURI(baseURI string, subscriptionID string) VaultsClient {
        return VaultsClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// CreateOrUpdate create or update a key vault in the specified subscription.
    // Parameters:
        // resourceGroupName - the name of the Resource Group to which the server belongs.
        // vaultName - name of the vault
        // parameters - parameters to create or update the vault
func (client VaultsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, vaultName string, parameters VaultCreateOrUpdateParameters) (result Vault, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/VaultsClient.CreateOrUpdate")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: vaultName,
             Constraints: []validation.Constraint{	{Target: "vaultName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9-]{3,24}$`, Chain: nil }}},
            { TargetValue: parameters,
             Constraints: []validation.Constraint{	{Target: "parameters.Location", Name: validation.Null, Rule: true, Chain: nil },
            	{Target: "parameters.Properties", Name: validation.Null, Rule: true ,
            Chain: []validation.Constraint{	{Target: "parameters.Properties.TenantID", Name: validation.Null, Rule: true, Chain: nil },
            	{Target: "parameters.Properties.Sku", Name: validation.Null, Rule: true ,
            Chain: []validation.Constraint{	{Target: "parameters.Properties.Sku.Family", Name: validation.Null, Rule: true, Chain: nil },
            }},
            	{Target: "parameters.Properties.AccessPolicies", Name: validation.Null, Rule: true ,
            Chain: []validation.Constraint{	{Target: "parameters.Properties.AccessPolicies", Name: validation.MaxItems, Rule: 16, Chain: nil },
            }},
            }}}}}); err != nil {
            return result, validation.NewError("keyvault.VaultsClient", "CreateOrUpdate", err.Error())
            }

                req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, vaultName, parameters)
    if err != nil {
    err = autorest.NewErrorWithError(err, "keyvault.VaultsClient", "CreateOrUpdate", nil , "Failure preparing request")
    return
    }

            resp, err := client.CreateOrUpdateSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "keyvault.VaultsClient", "CreateOrUpdate", resp, "Failure sending request")
            return
            }

            result, err = client.CreateOrUpdateResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "keyvault.VaultsClient", "CreateOrUpdate", resp, "Failure responding to request")
            }

    return
    }

    // CreateOrUpdatePreparer prepares the CreateOrUpdate request.
    func (client VaultsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, vaultName string, parameters VaultCreateOrUpdateParameters) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            "vaultName": autorest.Encode("path",vaultName),
            }

                        const APIVersion = "2015-06-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPut(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}",pathParameters),
    autorest.WithJSON(parameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
    // http.Response Body if it receives an error.
    func (client VaultsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client VaultsClient) CreateOrUpdateResponder(resp *http.Response) (result Vault, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusCreated),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// Delete deletes the specified Azure key vault.
    // Parameters:
        // resourceGroupName - the name of the Resource Group to which the vault belongs.
        // vaultName - the name of the vault to delete
func (client VaultsClient) Delete(ctx context.Context, resourceGroupName string, vaultName string) (result autorest.Response, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/VaultsClient.Delete")
        defer func() {
            sc := -1
            if result.Response != nil {
                sc = result.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.DeletePreparer(ctx, resourceGroupName, vaultName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "keyvault.VaultsClient", "Delete", nil , "Failure preparing request")
    return
    }

            resp, err := client.DeleteSender(req)
            if err != nil {
            result.Response = resp
            err = autorest.NewErrorWithError(err, "keyvault.VaultsClient", "Delete", resp, "Failure sending request")
            return
            }

            result, err = client.DeleteResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "keyvault.VaultsClient", "Delete", resp, "Failure responding to request")
            }

    return
    }

    // DeletePreparer prepares the Delete request.
    func (client VaultsClient) DeletePreparer(ctx context.Context, resourceGroupName string, vaultName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            "vaultName": autorest.Encode("path",vaultName),
            }

                        const APIVersion = "2015-06-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsDelete(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // DeleteSender sends the Delete request. The method will close the
    // http.Response Body if it receives an error.
    func (client VaultsClient) DeleteSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client VaultsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusNoContent),
    autorest.ByClosing())
    result.Response = resp
        return
    }

// Get gets the specified Azure key vault.
    // Parameters:
        // resourceGroupName - the name of the Resource Group to which the vault belongs.
        // vaultName - the name of the vault.
func (client VaultsClient) Get(ctx context.Context, resourceGroupName string, vaultName string) (result Vault, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/VaultsClient.Get")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.GetPreparer(ctx, resourceGroupName, vaultName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "keyvault.VaultsClient", "Get", nil , "Failure preparing request")
    return
    }

            resp, err := client.GetSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "keyvault.VaultsClient", "Get", resp, "Failure sending request")
            return
            }

            result, err = client.GetResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "keyvault.VaultsClient", "Get", resp, "Failure responding to request")
            }

    return
    }

    // GetPreparer prepares the Get request.
    func (client VaultsClient) GetPreparer(ctx context.Context, resourceGroupName string, vaultName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            "vaultName": autorest.Encode("path",vaultName),
            }

                        const APIVersion = "2015-06-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client VaultsClient) GetSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client VaultsClient) GetResponder(resp *http.Response) (result Vault, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// List the List operation gets information about the vaults associated with the subscription.
    // Parameters:
        // top - maximum number of results to return.
func (client VaultsClient) List(ctx context.Context, top *int32) (result ResourceListResultPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/VaultsClient.List")
        defer func() {
            sc := -1
            if result.rlr.Response.Response != nil {
                sc = result.rlr.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
                result.fn = client.listNextResults
    req, err := client.ListPreparer(ctx, top)
    if err != nil {
    err = autorest.NewErrorWithError(err, "keyvault.VaultsClient", "List", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListSender(req)
            if err != nil {
            result.rlr.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "keyvault.VaultsClient", "List", resp, "Failure sending request")
            return
            }

            result.rlr, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "keyvault.VaultsClient", "List", resp, "Failure responding to request")
            }

    return
    }

    // ListPreparer prepares the List request.
    func (client VaultsClient) ListPreparer(ctx context.Context, top *int32) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2015-11-01"
        queryParameters := map[string]interface{} {
        "$filter": autorest.Encode("query", "resourceType eq 'Microsoft.KeyVault/vaults'"),
        "api-version": APIVersion,
        }
            if top != nil {
            queryParameters["$top"] = autorest.Encode("query",*top)
            }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resources",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListSender sends the List request. The method will close the
    // http.Response Body if it receives an error.
    func (client VaultsClient) ListSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client VaultsClient) ListResponder(resp *http.Response) (result ResourceListResult, err error) {
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
            func (client VaultsClient) listNextResults(ctx context.Context, lastResults ResourceListResult) (result ResourceListResult, err error) {
            req, err := lastResults.resourceListResultPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "keyvault.VaultsClient", "listNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "keyvault.VaultsClient", "listNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "keyvault.VaultsClient", "listNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

    // ListComplete enumerates all values, automatically crossing page boundaries as required.
    func (client VaultsClient) ListComplete(ctx context.Context, top *int32) (result ResourceListResultIterator, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/VaultsClient.List")
            defer func() {
                sc := -1
                if result.Response().Response.Response != nil {
                    sc = result.page.Response().Response.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
     }
        result.page, err = client.List(ctx, top)
                return
        }

// ListByResourceGroup the List operation gets information about the vaults associated with the subscription and within
// the specified resource group.
    // Parameters:
        // resourceGroupName - the name of the Resource Group to which the vault belongs.
        // top - maximum number of results to return.
func (client VaultsClient) ListByResourceGroup(ctx context.Context, resourceGroupName string, top *int32) (result VaultListResultPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/VaultsClient.ListByResourceGroup")
        defer func() {
            sc := -1
            if result.vlr.Response.Response != nil {
                sc = result.vlr.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
                result.fn = client.listByResourceGroupNextResults
    req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName, top)
    if err != nil {
    err = autorest.NewErrorWithError(err, "keyvault.VaultsClient", "ListByResourceGroup", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListByResourceGroupSender(req)
            if err != nil {
            result.vlr.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "keyvault.VaultsClient", "ListByResourceGroup", resp, "Failure sending request")
            return
            }

            result.vlr, err = client.ListByResourceGroupResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "keyvault.VaultsClient", "ListByResourceGroup", resp, "Failure responding to request")
            }

    return
    }

    // ListByResourceGroupPreparer prepares the ListByResourceGroup request.
    func (client VaultsClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string, top *int32) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2015-06-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }
            if top != nil {
            queryParameters["$top"] = autorest.Encode("query",*top)
            }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
    // http.Response Body if it receives an error.
    func (client VaultsClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client VaultsClient) ListByResourceGroupResponder(resp *http.Response) (result VaultListResult, err error) {
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
            func (client VaultsClient) listByResourceGroupNextResults(ctx context.Context, lastResults VaultListResult) (result VaultListResult, err error) {
            req, err := lastResults.vaultListResultPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "keyvault.VaultsClient", "listByResourceGroupNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListByResourceGroupSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "keyvault.VaultsClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListByResourceGroupResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "keyvault.VaultsClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

    // ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
    func (client VaultsClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, top *int32) (result VaultListResultIterator, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/VaultsClient.ListByResourceGroup")
            defer func() {
                sc := -1
                if result.Response().Response.Response != nil {
                    sc = result.page.Response().Response.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
     }
        result.page, err = client.ListByResourceGroup(ctx, resourceGroupName, top)
                return
        }

