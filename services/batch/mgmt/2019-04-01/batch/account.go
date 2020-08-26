package batch

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

// AccountClient is the client for the Account methods of the Batch service.
type AccountClient struct {
	BaseClient
}

// NewAccountClient creates an instance of the AccountClient client.
func NewAccountClient(subscriptionID string) AccountClient {
	return NewAccountClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewAccountClientWithBaseURI creates an instance of the AccountClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewAccountClientWithBaseURI(baseURI string, subscriptionID string) AccountClient {
	return AccountClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create creates a new Batch account with the specified parameters. Existing accounts cannot be updated with this API
// and should instead be updated with the Update Batch Account API.
// Parameters:
// resourceGroupName - the name of the resource group that contains the Batch account.
// accountName - a name for the Batch account which must be unique within the region. Batch account names must
// be between 3 and 24 characters in length and must use only numbers and lowercase letters. This name is used
// as part of the DNS name that is used to access the Batch service in the region in which the account is
// created. For example: http://accountname.region.batch.azure.com/.
// parameters - additional parameters for account creation.
func (client AccountClient) Create(ctx context.Context, resourceGroupName string, accountName string, parameters AccountCreateParameters) (result AccountCreateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AccountClient.Create")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.Location", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.AccountCreateProperties", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "parameters.AccountCreateProperties.AutoStorage", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.AccountCreateProperties.AutoStorage.StorageAccountID", Name: validation.Null, Rule: true, Chain: nil}}},
						{Target: "parameters.AccountCreateProperties.KeyVaultReference", Name: validation.Null, Rule: false,
							Chain: []validation.Constraint{{Target: "parameters.AccountCreateProperties.KeyVaultReference.ID", Name: validation.Null, Rule: true, Chain: nil},
								{Target: "parameters.AccountCreateProperties.KeyVaultReference.URL", Name: validation.Null, Rule: true, Chain: nil},
							}},
					}}}}}); err != nil {
		return result, validation.NewError("batch.AccountClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, resourceGroupName, accountName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.AccountClient", "Create", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.AccountClient", "Create", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client AccountClient) CreatePreparer(ctx context.Context, resourceGroupName string, accountName string, parameters AccountCreateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client AccountClient) CreateSender(req *http.Request) (future AccountCreateFuture, err error) {
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
func (client AccountClient) CreateResponder(resp *http.Response) (result Account, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the specified Batch account.
// Parameters:
// resourceGroupName - the name of the resource group that contains the Batch account.
// accountName - the name of the Batch account.
func (client AccountClient) Delete(ctx context.Context, resourceGroupName string, accountName string) (result AccountDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AccountClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("batch.AccountClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, accountName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.AccountClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.AccountClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client AccountClient) DeletePreparer(ctx context.Context, resourceGroupName string, accountName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client AccountClient) DeleteSender(req *http.Request) (future AccountDeleteFuture, err error) {
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
func (client AccountClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets information about the specified Batch account.
// Parameters:
// resourceGroupName - the name of the resource group that contains the Batch account.
// accountName - the name of the Batch account.
func (client AccountClient) Get(ctx context.Context, resourceGroupName string, accountName string) (result Account, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AccountClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("batch.AccountClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, accountName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.AccountClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "batch.AccountClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.AccountClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client AccountClient) GetPreparer(ctx context.Context, resourceGroupName string, accountName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client AccountClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client AccountClient) GetResponder(resp *http.Response) (result Account, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetKeys this operation applies only to Batch accounts created with a poolAllocationMode of 'BatchService'. If the
// Batch account was created with a poolAllocationMode of 'UserSubscription', clients cannot use access to keys to
// authenticate, and must use Azure Active Directory instead. In this case, getting the keys will fail.
// Parameters:
// resourceGroupName - the name of the resource group that contains the Batch account.
// accountName - the name of the Batch account.
func (client AccountClient) GetKeys(ctx context.Context, resourceGroupName string, accountName string) (result AccountKeys, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AccountClient.GetKeys")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("batch.AccountClient", "GetKeys", err.Error())
	}

	req, err := client.GetKeysPreparer(ctx, resourceGroupName, accountName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.AccountClient", "GetKeys", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetKeysSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "batch.AccountClient", "GetKeys", resp, "Failure sending request")
		return
	}

	result, err = client.GetKeysResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.AccountClient", "GetKeys", resp, "Failure responding to request")
	}

	return
}

// GetKeysPreparer prepares the GetKeys request.
func (client AccountClient) GetKeysPreparer(ctx context.Context, resourceGroupName string, accountName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/listKeys", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetKeysSender sends the GetKeys request. The method will close the
// http.Response Body if it receives an error.
func (client AccountClient) GetKeysSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetKeysResponder handles the response to the GetKeys request. The method always
// closes the http.Response Body.
func (client AccountClient) GetKeysResponder(resp *http.Response) (result AccountKeys, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets information about the Batch accounts associated with the subscription.
func (client AccountClient) List(ctx context.Context) (result AccountListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AccountClient.List")
		defer func() {
			sc := -1
			if result.alr.Response.Response != nil {
				sc = result.alr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.AccountClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.alr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "batch.AccountClient", "List", resp, "Failure sending request")
		return
	}

	result.alr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.AccountClient", "List", resp, "Failure responding to request")
	}
	if result.alr.hasNextLink() && result.alr.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListPreparer prepares the List request.
func (client AccountClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Batch/batchAccounts", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client AccountClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client AccountClient) ListResponder(resp *http.Response) (result AccountListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client AccountClient) listNextResults(ctx context.Context, lastResults AccountListResult) (result AccountListResult, err error) {
	req, err := lastResults.accountListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "batch.AccountClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "batch.AccountClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.AccountClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client AccountClient) ListComplete(ctx context.Context) (result AccountListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AccountClient.List")
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

// ListByResourceGroup gets information about the Batch accounts associated with the specified resource group.
// Parameters:
// resourceGroupName - the name of the resource group that contains the Batch account.
func (client AccountClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (result AccountListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AccountClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.alr.Response.Response != nil {
				sc = result.alr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.AccountClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.alr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "batch.AccountClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.alr, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.AccountClient", "ListByResourceGroup", resp, "Failure responding to request")
	}
	if result.alr.hasNextLink() && result.alr.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client AccountClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client AccountClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client AccountClient) ListByResourceGroupResponder(resp *http.Response) (result AccountListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client AccountClient) listByResourceGroupNextResults(ctx context.Context, lastResults AccountListResult) (result AccountListResult, err error) {
	req, err := lastResults.accountListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "batch.AccountClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "batch.AccountClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.AccountClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client AccountClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result AccountListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AccountClient.ListByResourceGroup")
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

// RegenerateKey regenerates the specified account key for the Batch account.
// Parameters:
// resourceGroupName - the name of the resource group that contains the Batch account.
// accountName - the name of the Batch account.
// parameters - the type of key to regenerate.
func (client AccountClient) RegenerateKey(ctx context.Context, resourceGroupName string, accountName string, parameters AccountRegenerateKeyParameters) (result AccountKeys, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AccountClient.RegenerateKey")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("batch.AccountClient", "RegenerateKey", err.Error())
	}

	req, err := client.RegenerateKeyPreparer(ctx, resourceGroupName, accountName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.AccountClient", "RegenerateKey", nil, "Failure preparing request")
		return
	}

	resp, err := client.RegenerateKeySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "batch.AccountClient", "RegenerateKey", resp, "Failure sending request")
		return
	}

	result, err = client.RegenerateKeyResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.AccountClient", "RegenerateKey", resp, "Failure responding to request")
	}

	return
}

// RegenerateKeyPreparer prepares the RegenerateKey request.
func (client AccountClient) RegenerateKeyPreparer(ctx context.Context, resourceGroupName string, accountName string, parameters AccountRegenerateKeyParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/regenerateKeys", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RegenerateKeySender sends the RegenerateKey request. The method will close the
// http.Response Body if it receives an error.
func (client AccountClient) RegenerateKeySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// RegenerateKeyResponder handles the response to the RegenerateKey request. The method always
// closes the http.Response Body.
func (client AccountClient) RegenerateKeyResponder(resp *http.Response) (result AccountKeys, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// SynchronizeAutoStorageKeys synchronizes access keys for the auto-storage account configured for the specified Batch
// account.
// Parameters:
// resourceGroupName - the name of the resource group that contains the Batch account.
// accountName - the name of the Batch account.
func (client AccountClient) SynchronizeAutoStorageKeys(ctx context.Context, resourceGroupName string, accountName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AccountClient.SynchronizeAutoStorageKeys")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("batch.AccountClient", "SynchronizeAutoStorageKeys", err.Error())
	}

	req, err := client.SynchronizeAutoStorageKeysPreparer(ctx, resourceGroupName, accountName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.AccountClient", "SynchronizeAutoStorageKeys", nil, "Failure preparing request")
		return
	}

	resp, err := client.SynchronizeAutoStorageKeysSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "batch.AccountClient", "SynchronizeAutoStorageKeys", resp, "Failure sending request")
		return
	}

	result, err = client.SynchronizeAutoStorageKeysResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.AccountClient", "SynchronizeAutoStorageKeys", resp, "Failure responding to request")
	}

	return
}

// SynchronizeAutoStorageKeysPreparer prepares the SynchronizeAutoStorageKeys request.
func (client AccountClient) SynchronizeAutoStorageKeysPreparer(ctx context.Context, resourceGroupName string, accountName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/syncAutoStorageKeys", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// SynchronizeAutoStorageKeysSender sends the SynchronizeAutoStorageKeys request. The method will close the
// http.Response Body if it receives an error.
func (client AccountClient) SynchronizeAutoStorageKeysSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// SynchronizeAutoStorageKeysResponder handles the response to the SynchronizeAutoStorageKeys request. The method always
// closes the http.Response Body.
func (client AccountClient) SynchronizeAutoStorageKeysResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Update updates the properties of an existing Batch account.
// Parameters:
// resourceGroupName - the name of the resource group that contains the Batch account.
// accountName - the name of the Batch account.
// parameters - additional parameters for account update.
func (client AccountClient) Update(ctx context.Context, resourceGroupName string, accountName string, parameters AccountUpdateParameters) (result Account, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AccountClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("batch.AccountClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, resourceGroupName, accountName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.AccountClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "batch.AccountClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.AccountClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client AccountClient) UpdatePreparer(ctx context.Context, resourceGroupName string, accountName string, parameters AccountUpdateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client AccountClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client AccountClient) UpdateResponder(resp *http.Response) (result Account, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
