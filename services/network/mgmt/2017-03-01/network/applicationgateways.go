package network

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

// ApplicationGatewaysClient is the network Client
type ApplicationGatewaysClient struct {
	BaseClient
}

// NewApplicationGatewaysClient creates an instance of the ApplicationGatewaysClient client.
func NewApplicationGatewaysClient(subscriptionID string) ApplicationGatewaysClient {
	return NewApplicationGatewaysClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewApplicationGatewaysClientWithBaseURI creates an instance of the ApplicationGatewaysClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewApplicationGatewaysClientWithBaseURI(baseURI string, subscriptionID string) ApplicationGatewaysClient {
	return ApplicationGatewaysClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// BackendHealth gets the backend health of the specified application gateway in a resource group.
// Parameters:
// resourceGroupName - the name of the resource group.
// applicationGatewayName - the name of the application gateway.
// expand - expands BackendAddressPool and BackendHttpSettings referenced in backend health.
func (client ApplicationGatewaysClient) BackendHealth(ctx context.Context, resourceGroupName string, applicationGatewayName string, expand string) (result ApplicationGatewaysBackendHealthFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplicationGatewaysClient.BackendHealth")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.BackendHealthPreparer(ctx, resourceGroupName, applicationGatewayName, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ApplicationGatewaysClient", "BackendHealth", nil, "Failure preparing request")
		return
	}

	result, err = client.BackendHealthSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ApplicationGatewaysClient", "BackendHealth", result.Response(), "Failure sending request")
		return
	}

	return
}

// BackendHealthPreparer prepares the BackendHealth request.
func (client ApplicationGatewaysClient) BackendHealthPreparer(ctx context.Context, resourceGroupName string, applicationGatewayName string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationGatewayName": autorest.Encode("path", applicationGatewayName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationGateways/{applicationGatewayName}/backendhealth", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// BackendHealthSender sends the BackendHealth request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationGatewaysClient) BackendHealthSender(req *http.Request) (future ApplicationGatewaysBackendHealthFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// BackendHealthResponder handles the response to the BackendHealth request. The method always
// closes the http.Response Body.
func (client ApplicationGatewaysClient) BackendHealthResponder(resp *http.Response) (result ApplicationGatewayBackendHealth, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CreateOrUpdate creates or updates the specified application gateway.
// Parameters:
// resourceGroupName - the name of the resource group.
// applicationGatewayName - the name of the application gateway.
// parameters - parameters supplied to the create or update application gateway operation.
func (client ApplicationGatewaysClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, applicationGatewayName string, parameters ApplicationGateway) (result ApplicationGatewaysCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplicationGatewaysClient.CreateOrUpdate")
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
			Constraints: []validation.Constraint{{Target: "parameters.ApplicationGatewayPropertiesFormat", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.ApplicationGatewayPropertiesFormat.WebApplicationFirewallConfiguration", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "parameters.ApplicationGatewayPropertiesFormat.WebApplicationFirewallConfiguration.Enabled", Name: validation.Null, Rule: true, Chain: nil},
						{Target: "parameters.ApplicationGatewayPropertiesFormat.WebApplicationFirewallConfiguration.RuleSetType", Name: validation.Null, Rule: true, Chain: nil},
						{Target: "parameters.ApplicationGatewayPropertiesFormat.WebApplicationFirewallConfiguration.RuleSetVersion", Name: validation.Null, Rule: true, Chain: nil},
					}},
				}}}}}); err != nil {
		return result, validation.NewError("network.ApplicationGatewaysClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, applicationGatewayName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ApplicationGatewaysClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ApplicationGatewaysClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ApplicationGatewaysClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, applicationGatewayName string, parameters ApplicationGateway) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationGatewayName": autorest.Encode("path", applicationGatewayName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationGateways/{applicationGatewayName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationGatewaysClient) CreateOrUpdateSender(req *http.Request) (future ApplicationGatewaysCreateOrUpdateFuture, err error) {
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
func (client ApplicationGatewaysClient) CreateOrUpdateResponder(resp *http.Response) (result ApplicationGateway, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the specified application gateway.
// Parameters:
// resourceGroupName - the name of the resource group.
// applicationGatewayName - the name of the application gateway.
func (client ApplicationGatewaysClient) Delete(ctx context.Context, resourceGroupName string, applicationGatewayName string) (result ApplicationGatewaysDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplicationGatewaysClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, applicationGatewayName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ApplicationGatewaysClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ApplicationGatewaysClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ApplicationGatewaysClient) DeletePreparer(ctx context.Context, resourceGroupName string, applicationGatewayName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationGatewayName": autorest.Encode("path", applicationGatewayName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationGateways/{applicationGatewayName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationGatewaysClient) DeleteSender(req *http.Request) (future ApplicationGatewaysDeleteFuture, err error) {
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
func (client ApplicationGatewaysClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the specified application gateway.
// Parameters:
// resourceGroupName - the name of the resource group.
// applicationGatewayName - the name of the application gateway.
func (client ApplicationGatewaysClient) Get(ctx context.Context, resourceGroupName string, applicationGatewayName string) (result ApplicationGateway, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplicationGatewaysClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, applicationGatewayName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ApplicationGatewaysClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.ApplicationGatewaysClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ApplicationGatewaysClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ApplicationGatewaysClient) GetPreparer(ctx context.Context, resourceGroupName string, applicationGatewayName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationGatewayName": autorest.Encode("path", applicationGatewayName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationGateways/{applicationGatewayName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationGatewaysClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ApplicationGatewaysClient) GetResponder(resp *http.Response) (result ApplicationGateway, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists all application gateways in a resource group.
// Parameters:
// resourceGroupName - the name of the resource group.
func (client ApplicationGatewaysClient) List(ctx context.Context, resourceGroupName string) (result ApplicationGatewayListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplicationGatewaysClient.List")
		defer func() {
			sc := -1
			if result.aglr.Response.Response != nil {
				sc = result.aglr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ApplicationGatewaysClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.aglr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.ApplicationGatewaysClient", "List", resp, "Failure sending request")
		return
	}

	result.aglr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ApplicationGatewaysClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ApplicationGatewaysClient) ListPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationGateways", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationGatewaysClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ApplicationGatewaysClient) ListResponder(resp *http.Response) (result ApplicationGatewayListResult, err error) {
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
func (client ApplicationGatewaysClient) listNextResults(ctx context.Context, lastResults ApplicationGatewayListResult) (result ApplicationGatewayListResult, err error) {
	req, err := lastResults.applicationGatewayListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.ApplicationGatewaysClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.ApplicationGatewaysClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ApplicationGatewaysClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client ApplicationGatewaysClient) ListComplete(ctx context.Context, resourceGroupName string) (result ApplicationGatewayListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplicationGatewaysClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName)
	return
}

// ListAll gets all the application gateways in a subscription.
func (client ApplicationGatewaysClient) ListAll(ctx context.Context) (result ApplicationGatewayListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplicationGatewaysClient.ListAll")
		defer func() {
			sc := -1
			if result.aglr.Response.Response != nil {
				sc = result.aglr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listAllNextResults
	req, err := client.ListAllPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ApplicationGatewaysClient", "ListAll", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListAllSender(req)
	if err != nil {
		result.aglr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.ApplicationGatewaysClient", "ListAll", resp, "Failure sending request")
		return
	}

	result.aglr, err = client.ListAllResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ApplicationGatewaysClient", "ListAll", resp, "Failure responding to request")
	}

	return
}

// ListAllPreparer prepares the ListAll request.
func (client ApplicationGatewaysClient) ListAllPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Network/applicationGateways", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListAllSender sends the ListAll request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationGatewaysClient) ListAllSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListAllResponder handles the response to the ListAll request. The method always
// closes the http.Response Body.
func (client ApplicationGatewaysClient) ListAllResponder(resp *http.Response) (result ApplicationGatewayListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listAllNextResults retrieves the next set of results, if any.
func (client ApplicationGatewaysClient) listAllNextResults(ctx context.Context, lastResults ApplicationGatewayListResult) (result ApplicationGatewayListResult, err error) {
	req, err := lastResults.applicationGatewayListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.ApplicationGatewaysClient", "listAllNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListAllSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.ApplicationGatewaysClient", "listAllNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListAllResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ApplicationGatewaysClient", "listAllNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListAllComplete enumerates all values, automatically crossing page boundaries as required.
func (client ApplicationGatewaysClient) ListAllComplete(ctx context.Context) (result ApplicationGatewayListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplicationGatewaysClient.ListAll")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListAll(ctx)
	return
}

// ListAvailableWafRuleSets lists all available web application firewall rule sets.
func (client ApplicationGatewaysClient) ListAvailableWafRuleSets(ctx context.Context) (result ApplicationGatewayAvailableWafRuleSetsResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplicationGatewaysClient.ListAvailableWafRuleSets")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListAvailableWafRuleSetsPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ApplicationGatewaysClient", "ListAvailableWafRuleSets", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListAvailableWafRuleSetsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.ApplicationGatewaysClient", "ListAvailableWafRuleSets", resp, "Failure sending request")
		return
	}

	result, err = client.ListAvailableWafRuleSetsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ApplicationGatewaysClient", "ListAvailableWafRuleSets", resp, "Failure responding to request")
	}

	return
}

// ListAvailableWafRuleSetsPreparer prepares the ListAvailableWafRuleSets request.
func (client ApplicationGatewaysClient) ListAvailableWafRuleSetsPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Network/applicationGatewayAvailableWafRuleSets", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListAvailableWafRuleSetsSender sends the ListAvailableWafRuleSets request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationGatewaysClient) ListAvailableWafRuleSetsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListAvailableWafRuleSetsResponder handles the response to the ListAvailableWafRuleSets request. The method always
// closes the http.Response Body.
func (client ApplicationGatewaysClient) ListAvailableWafRuleSetsResponder(resp *http.Response) (result ApplicationGatewayAvailableWafRuleSetsResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Start starts the specified application gateway.
// Parameters:
// resourceGroupName - the name of the resource group.
// applicationGatewayName - the name of the application gateway.
func (client ApplicationGatewaysClient) Start(ctx context.Context, resourceGroupName string, applicationGatewayName string) (result ApplicationGatewaysStartFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplicationGatewaysClient.Start")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.StartPreparer(ctx, resourceGroupName, applicationGatewayName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ApplicationGatewaysClient", "Start", nil, "Failure preparing request")
		return
	}

	result, err = client.StartSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ApplicationGatewaysClient", "Start", result.Response(), "Failure sending request")
		return
	}

	return
}

// StartPreparer prepares the Start request.
func (client ApplicationGatewaysClient) StartPreparer(ctx context.Context, resourceGroupName string, applicationGatewayName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationGatewayName": autorest.Encode("path", applicationGatewayName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationGateways/{applicationGatewayName}/start", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// StartSender sends the Start request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationGatewaysClient) StartSender(req *http.Request) (future ApplicationGatewaysStartFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// StartResponder handles the response to the Start request. The method always
// closes the http.Response Body.
func (client ApplicationGatewaysClient) StartResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Stop stops the specified application gateway in a resource group.
// Parameters:
// resourceGroupName - the name of the resource group.
// applicationGatewayName - the name of the application gateway.
func (client ApplicationGatewaysClient) Stop(ctx context.Context, resourceGroupName string, applicationGatewayName string) (result ApplicationGatewaysStopFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplicationGatewaysClient.Stop")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.StopPreparer(ctx, resourceGroupName, applicationGatewayName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ApplicationGatewaysClient", "Stop", nil, "Failure preparing request")
		return
	}

	result, err = client.StopSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ApplicationGatewaysClient", "Stop", result.Response(), "Failure sending request")
		return
	}

	return
}

// StopPreparer prepares the Stop request.
func (client ApplicationGatewaysClient) StopPreparer(ctx context.Context, resourceGroupName string, applicationGatewayName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationGatewayName": autorest.Encode("path", applicationGatewayName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationGateways/{applicationGatewayName}/stop", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// StopSender sends the Stop request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationGatewaysClient) StopSender(req *http.Request) (future ApplicationGatewaysStopFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// StopResponder handles the response to the Stop request. The method always
// closes the http.Response Body.
func (client ApplicationGatewaysClient) StopResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}
