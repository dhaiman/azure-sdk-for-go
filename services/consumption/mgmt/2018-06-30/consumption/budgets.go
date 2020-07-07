package consumption

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

// BudgetsClient is the consumption management client provides access to consumption resources for Azure Enterprise
// Subscriptions.
type BudgetsClient struct {
	BaseClient
}

// NewBudgetsClient creates an instance of the BudgetsClient client.
func NewBudgetsClient(subscriptionID string) BudgetsClient {
	return NewBudgetsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewBudgetsClientWithBaseURI creates an instance of the BudgetsClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewBudgetsClientWithBaseURI(baseURI string, subscriptionID string) BudgetsClient {
	return BudgetsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate the operation to create or update a budget. Update operation requires latest eTag to be set in the
// request mandatorily. You may obtain the latest eTag by performing a get operation. Create operation does not require
// eTag.
// Parameters:
// budgetName - budget Name.
// parameters - parameters supplied to the Create Budget operation.
func (client BudgetsClient) CreateOrUpdate(ctx context.Context, budgetName string, parameters Budget) (result Budget, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BudgetsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.BudgetProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.BudgetProperties.Amount", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.BudgetProperties.TimePeriod", Name: validation.Null, Rule: true,
						Chain: []validation.Constraint{{Target: "parameters.BudgetProperties.TimePeriod.StartDate", Name: validation.Null, Rule: true, Chain: nil}}},
					{Target: "parameters.BudgetProperties.Filters", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.BudgetProperties.Filters.ResourceGroups", Name: validation.Null, Rule: false,
							Chain: []validation.Constraint{{Target: "parameters.BudgetProperties.Filters.ResourceGroups", Name: validation.MaxItems, Rule: 10, Chain: nil},
								{Target: "parameters.BudgetProperties.Filters.ResourceGroups", Name: validation.MinItems, Rule: 0, Chain: nil},
							}},
							{Target: "parameters.BudgetProperties.Filters.Resources", Name: validation.Null, Rule: false,
								Chain: []validation.Constraint{{Target: "parameters.BudgetProperties.Filters.Resources", Name: validation.MaxItems, Rule: 10, Chain: nil},
									{Target: "parameters.BudgetProperties.Filters.Resources", Name: validation.MinItems, Rule: 0, Chain: nil},
								}},
							{Target: "parameters.BudgetProperties.Filters.Meters", Name: validation.Null, Rule: false,
								Chain: []validation.Constraint{{Target: "parameters.BudgetProperties.Filters.Meters", Name: validation.MaxItems, Rule: 10, Chain: nil},
									{Target: "parameters.BudgetProperties.Filters.Meters", Name: validation.MinItems, Rule: 0, Chain: nil},
								}},
						}},
				}}}}}); err != nil {
		return result, validation.NewError("consumption.BudgetsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, budgetName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.BudgetsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "consumption.BudgetsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.BudgetsClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client BudgetsClient) CreateOrUpdatePreparer(ctx context.Context, budgetName string, parameters Budget) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"budgetName":     autorest.Encode("path", budgetName),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Consumption/budgets/{budgetName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client BudgetsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client BudgetsClient) CreateOrUpdateResponder(resp *http.Response) (result Budget, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CreateOrUpdateByResourceGroupName the operation to create or update a budget. Update operation requires latest eTag
// to be set in the request mandatorily. You may obtain the latest eTag by performing a get operation. Create operation
// does not require eTag.
// Parameters:
// resourceGroupName - azure Resource Group Name.
// budgetName - budget Name.
// parameters - parameters supplied to the Create Budget operation.
func (client BudgetsClient) CreateOrUpdateByResourceGroupName(ctx context.Context, resourceGroupName string, budgetName string, parameters Budget) (result Budget, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BudgetsClient.CreateOrUpdateByResourceGroupName")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.BudgetProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.BudgetProperties.Amount", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.BudgetProperties.TimePeriod", Name: validation.Null, Rule: true,
						Chain: []validation.Constraint{{Target: "parameters.BudgetProperties.TimePeriod.StartDate", Name: validation.Null, Rule: true, Chain: nil}}},
					{Target: "parameters.BudgetProperties.Filters", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.BudgetProperties.Filters.ResourceGroups", Name: validation.Null, Rule: false,
							Chain: []validation.Constraint{{Target: "parameters.BudgetProperties.Filters.ResourceGroups", Name: validation.MaxItems, Rule: 10, Chain: nil},
								{Target: "parameters.BudgetProperties.Filters.ResourceGroups", Name: validation.MinItems, Rule: 0, Chain: nil},
							}},
							{Target: "parameters.BudgetProperties.Filters.Resources", Name: validation.Null, Rule: false,
								Chain: []validation.Constraint{{Target: "parameters.BudgetProperties.Filters.Resources", Name: validation.MaxItems, Rule: 10, Chain: nil},
									{Target: "parameters.BudgetProperties.Filters.Resources", Name: validation.MinItems, Rule: 0, Chain: nil},
								}},
							{Target: "parameters.BudgetProperties.Filters.Meters", Name: validation.Null, Rule: false,
								Chain: []validation.Constraint{{Target: "parameters.BudgetProperties.Filters.Meters", Name: validation.MaxItems, Rule: 10, Chain: nil},
									{Target: "parameters.BudgetProperties.Filters.Meters", Name: validation.MinItems, Rule: 0, Chain: nil},
								}},
						}},
				}}}}}); err != nil {
		return result, validation.NewError("consumption.BudgetsClient", "CreateOrUpdateByResourceGroupName", err.Error())
	}

	req, err := client.CreateOrUpdateByResourceGroupNamePreparer(ctx, resourceGroupName, budgetName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.BudgetsClient", "CreateOrUpdateByResourceGroupName", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateByResourceGroupNameSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "consumption.BudgetsClient", "CreateOrUpdateByResourceGroupName", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateByResourceGroupNameResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.BudgetsClient", "CreateOrUpdateByResourceGroupName", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdateByResourceGroupNamePreparer prepares the CreateOrUpdateByResourceGroupName request.
func (client BudgetsClient) CreateOrUpdateByResourceGroupNamePreparer(ctx context.Context, resourceGroupName string, budgetName string, parameters Budget) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"budgetName":        autorest.Encode("path", budgetName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Consumption/budgets/{budgetName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateByResourceGroupNameSender sends the CreateOrUpdateByResourceGroupName request. The method will close the
// http.Response Body if it receives an error.
func (client BudgetsClient) CreateOrUpdateByResourceGroupNameSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateByResourceGroupNameResponder handles the response to the CreateOrUpdateByResourceGroupName request. The method always
// closes the http.Response Body.
func (client BudgetsClient) CreateOrUpdateByResourceGroupNameResponder(resp *http.Response) (result Budget, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete the operation to delete a budget.
// Parameters:
// budgetName - budget Name.
func (client BudgetsClient) Delete(ctx context.Context, budgetName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BudgetsClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, budgetName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.BudgetsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "consumption.BudgetsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.BudgetsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client BudgetsClient) DeletePreparer(ctx context.Context, budgetName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"budgetName":     autorest.Encode("path", budgetName),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Consumption/budgets/{budgetName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client BudgetsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client BudgetsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// DeleteByResourceGroupName the operation to delete a budget.
// Parameters:
// resourceGroupName - azure Resource Group Name.
// budgetName - budget Name.
func (client BudgetsClient) DeleteByResourceGroupName(ctx context.Context, resourceGroupName string, budgetName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BudgetsClient.DeleteByResourceGroupName")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeleteByResourceGroupNamePreparer(ctx, resourceGroupName, budgetName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.BudgetsClient", "DeleteByResourceGroupName", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteByResourceGroupNameSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "consumption.BudgetsClient", "DeleteByResourceGroupName", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteByResourceGroupNameResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.BudgetsClient", "DeleteByResourceGroupName", resp, "Failure responding to request")
	}

	return
}

// DeleteByResourceGroupNamePreparer prepares the DeleteByResourceGroupName request.
func (client BudgetsClient) DeleteByResourceGroupNamePreparer(ctx context.Context, resourceGroupName string, budgetName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"budgetName":        autorest.Encode("path", budgetName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Consumption/budgets/{budgetName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteByResourceGroupNameSender sends the DeleteByResourceGroupName request. The method will close the
// http.Response Body if it receives an error.
func (client BudgetsClient) DeleteByResourceGroupNameSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteByResourceGroupNameResponder handles the response to the DeleteByResourceGroupName request. The method always
// closes the http.Response Body.
func (client BudgetsClient) DeleteByResourceGroupNameResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the budget for a subscription by budget name.
// Parameters:
// budgetName - budget Name.
func (client BudgetsClient) Get(ctx context.Context, budgetName string) (result Budget, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BudgetsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, budgetName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.BudgetsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "consumption.BudgetsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.BudgetsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client BudgetsClient) GetPreparer(ctx context.Context, budgetName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"budgetName":     autorest.Encode("path", budgetName),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Consumption/budgets/{budgetName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client BudgetsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client BudgetsClient) GetResponder(resp *http.Response) (result Budget, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetByResourceGroupName gets the budget for a resource group under a subscription by budget name.
// Parameters:
// resourceGroupName - azure Resource Group Name.
// budgetName - budget Name.
func (client BudgetsClient) GetByResourceGroupName(ctx context.Context, resourceGroupName string, budgetName string) (result Budget, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BudgetsClient.GetByResourceGroupName")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetByResourceGroupNamePreparer(ctx, resourceGroupName, budgetName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.BudgetsClient", "GetByResourceGroupName", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetByResourceGroupNameSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "consumption.BudgetsClient", "GetByResourceGroupName", resp, "Failure sending request")
		return
	}

	result, err = client.GetByResourceGroupNameResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.BudgetsClient", "GetByResourceGroupName", resp, "Failure responding to request")
	}

	return
}

// GetByResourceGroupNamePreparer prepares the GetByResourceGroupName request.
func (client BudgetsClient) GetByResourceGroupNamePreparer(ctx context.Context, resourceGroupName string, budgetName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"budgetName":        autorest.Encode("path", budgetName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Consumption/budgets/{budgetName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetByResourceGroupNameSender sends the GetByResourceGroupName request. The method will close the
// http.Response Body if it receives an error.
func (client BudgetsClient) GetByResourceGroupNameSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetByResourceGroupNameResponder handles the response to the GetByResourceGroupName request. The method always
// closes the http.Response Body.
func (client BudgetsClient) GetByResourceGroupNameResponder(resp *http.Response) (result Budget, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists all budgets for a subscription.
func (client BudgetsClient) List(ctx context.Context) (result BudgetsListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BudgetsClient.List")
		defer func() {
			sc := -1
			if result.blr.Response.Response != nil {
				sc = result.blr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.BudgetsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.blr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "consumption.BudgetsClient", "List", resp, "Failure sending request")
		return
	}

	result.blr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.BudgetsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client BudgetsClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Consumption/budgets", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client BudgetsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client BudgetsClient) ListResponder(resp *http.Response) (result BudgetsListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client BudgetsClient) listNextResults(ctx context.Context, lastResults BudgetsListResult) (result BudgetsListResult, err error) {
	req, err := lastResults.budgetsListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "consumption.BudgetsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "consumption.BudgetsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.BudgetsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client BudgetsClient) ListComplete(ctx context.Context) (result BudgetsListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BudgetsClient.List")
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

// ListByResourceGroupName lists all budgets for a resource group under a subscription.
// Parameters:
// resourceGroupName - azure Resource Group Name.
func (client BudgetsClient) ListByResourceGroupName(ctx context.Context, resourceGroupName string) (result BudgetsListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BudgetsClient.ListByResourceGroupName")
		defer func() {
			sc := -1
			if result.blr.Response.Response != nil {
				sc = result.blr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByResourceGroupNameNextResults
	req, err := client.ListByResourceGroupNamePreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.BudgetsClient", "ListByResourceGroupName", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupNameSender(req)
	if err != nil {
		result.blr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "consumption.BudgetsClient", "ListByResourceGroupName", resp, "Failure sending request")
		return
	}

	result.blr, err = client.ListByResourceGroupNameResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.BudgetsClient", "ListByResourceGroupName", resp, "Failure responding to request")
	}

	return
}

// ListByResourceGroupNamePreparer prepares the ListByResourceGroupName request.
func (client BudgetsClient) ListByResourceGroupNamePreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Consumption/budgets", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupNameSender sends the ListByResourceGroupName request. The method will close the
// http.Response Body if it receives an error.
func (client BudgetsClient) ListByResourceGroupNameSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupNameResponder handles the response to the ListByResourceGroupName request. The method always
// closes the http.Response Body.
func (client BudgetsClient) ListByResourceGroupNameResponder(resp *http.Response) (result BudgetsListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNameNextResults retrieves the next set of results, if any.
func (client BudgetsClient) listByResourceGroupNameNextResults(ctx context.Context, lastResults BudgetsListResult) (result BudgetsListResult, err error) {
	req, err := lastResults.budgetsListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "consumption.BudgetsClient", "listByResourceGroupNameNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupNameSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "consumption.BudgetsClient", "listByResourceGroupNameNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupNameResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.BudgetsClient", "listByResourceGroupNameNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupNameComplete enumerates all values, automatically crossing page boundaries as required.
func (client BudgetsClient) ListByResourceGroupNameComplete(ctx context.Context, resourceGroupName string) (result BudgetsListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BudgetsClient.ListByResourceGroupName")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByResourceGroupName(ctx, resourceGroupName)
	return
}
