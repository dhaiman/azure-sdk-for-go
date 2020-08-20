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

// WebApplicationFirewallPoliciesClient is the network Client
type WebApplicationFirewallPoliciesClient struct {
	BaseClient
}

// NewWebApplicationFirewallPoliciesClient creates an instance of the WebApplicationFirewallPoliciesClient client.
func NewWebApplicationFirewallPoliciesClient(subscriptionID string) WebApplicationFirewallPoliciesClient {
	return NewWebApplicationFirewallPoliciesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWebApplicationFirewallPoliciesClientWithBaseURI creates an instance of the WebApplicationFirewallPoliciesClient
// client using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI
// (sovereign clouds, Azure stack).
func NewWebApplicationFirewallPoliciesClientWithBaseURI(baseURI string, subscriptionID string) WebApplicationFirewallPoliciesClient {
	return WebApplicationFirewallPoliciesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or update policy with specified rule set name within a resource group.
// Parameters:
// resourceGroupName - the name of the resource group.
// policyName - the name of the policy.
// parameters - policy to be created.
func (client WebApplicationFirewallPoliciesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, policyName string, parameters WebApplicationFirewallPolicy) (result WebApplicationFirewallPolicy, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WebApplicationFirewallPoliciesClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: policyName,
			Constraints: []validation.Constraint{{Target: "policyName", Name: validation.MaxLength, Rule: 128, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.WebApplicationFirewallPolicyPropertiesFormat", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.WebApplicationFirewallPolicyPropertiesFormat.PolicySettings", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "parameters.WebApplicationFirewallPolicyPropertiesFormat.PolicySettings.MaxRequestBodySizeInKb", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.WebApplicationFirewallPolicyPropertiesFormat.PolicySettings.MaxRequestBodySizeInKb", Name: validation.InclusiveMaximum, Rule: int64(128), Chain: nil},
							{Target: "parameters.WebApplicationFirewallPolicyPropertiesFormat.PolicySettings.MaxRequestBodySizeInKb", Name: validation.InclusiveMinimum, Rule: int64(8), Chain: nil},
						}},
						{Target: "parameters.WebApplicationFirewallPolicyPropertiesFormat.PolicySettings.FileUploadLimitInMb", Name: validation.Null, Rule: false,
							Chain: []validation.Constraint{{Target: "parameters.WebApplicationFirewallPolicyPropertiesFormat.PolicySettings.FileUploadLimitInMb", Name: validation.InclusiveMinimum, Rule: int64(0), Chain: nil}}},
					}},
					{Target: "parameters.WebApplicationFirewallPolicyPropertiesFormat.ManagedRules", Name: validation.Null, Rule: true,
						Chain: []validation.Constraint{{Target: "parameters.WebApplicationFirewallPolicyPropertiesFormat.ManagedRules.ManagedRuleSets", Name: validation.Null, Rule: true, Chain: nil}}},
				}}}}}); err != nil {
		return result, validation.NewError("network.WebApplicationFirewallPoliciesClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, policyName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.WebApplicationFirewallPoliciesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.WebApplicationFirewallPoliciesClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.WebApplicationFirewallPoliciesClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client WebApplicationFirewallPoliciesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, policyName string, parameters WebApplicationFirewallPolicy) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"policyName":        autorest.Encode("path", policyName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	parameters.Etag = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ApplicationGatewayWebApplicationFirewallPolicies/{policyName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client WebApplicationFirewallPoliciesClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client WebApplicationFirewallPoliciesClient) CreateOrUpdateResponder(resp *http.Response) (result WebApplicationFirewallPolicy, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes Policy.
// Parameters:
// resourceGroupName - the name of the resource group.
// policyName - the name of the policy.
func (client WebApplicationFirewallPoliciesClient) Delete(ctx context.Context, resourceGroupName string, policyName string) (result WebApplicationFirewallPoliciesDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WebApplicationFirewallPoliciesClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: policyName,
			Constraints: []validation.Constraint{{Target: "policyName", Name: validation.MaxLength, Rule: 128, Chain: nil}}}}); err != nil {
		return result, validation.NewError("network.WebApplicationFirewallPoliciesClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, policyName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.WebApplicationFirewallPoliciesClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.WebApplicationFirewallPoliciesClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client WebApplicationFirewallPoliciesClient) DeletePreparer(ctx context.Context, resourceGroupName string, policyName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"policyName":        autorest.Encode("path", policyName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ApplicationGatewayWebApplicationFirewallPolicies/{policyName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client WebApplicationFirewallPoliciesClient) DeleteSender(req *http.Request) (future WebApplicationFirewallPoliciesDeleteFuture, err error) {
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
func (client WebApplicationFirewallPoliciesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get retrieve protection policy with specified name within a resource group.
// Parameters:
// resourceGroupName - the name of the resource group.
// policyName - the name of the policy.
func (client WebApplicationFirewallPoliciesClient) Get(ctx context.Context, resourceGroupName string, policyName string) (result WebApplicationFirewallPolicy, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WebApplicationFirewallPoliciesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: policyName,
			Constraints: []validation.Constraint{{Target: "policyName", Name: validation.MaxLength, Rule: 128, Chain: nil}}}}); err != nil {
		return result, validation.NewError("network.WebApplicationFirewallPoliciesClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, policyName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.WebApplicationFirewallPoliciesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.WebApplicationFirewallPoliciesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.WebApplicationFirewallPoliciesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client WebApplicationFirewallPoliciesClient) GetPreparer(ctx context.Context, resourceGroupName string, policyName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"policyName":        autorest.Encode("path", policyName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ApplicationGatewayWebApplicationFirewallPolicies/{policyName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client WebApplicationFirewallPoliciesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client WebApplicationFirewallPoliciesClient) GetResponder(resp *http.Response) (result WebApplicationFirewallPolicy, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists all of the protection policies within a resource group.
// Parameters:
// resourceGroupName - the name of the resource group.
func (client WebApplicationFirewallPoliciesClient) List(ctx context.Context, resourceGroupName string) (result WebApplicationFirewallPolicyListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WebApplicationFirewallPoliciesClient.List")
		defer func() {
			sc := -1
			if result.wafplr.Response.Response != nil {
				sc = result.wafplr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.WebApplicationFirewallPoliciesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.wafplr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.WebApplicationFirewallPoliciesClient", "List", resp, "Failure sending request")
		return
	}

	result.wafplr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.WebApplicationFirewallPoliciesClient", "List", resp, "Failure responding to request")
	}
	if result.wafplr.hasNextLink() && result.wafplr.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListPreparer prepares the List request.
func (client WebApplicationFirewallPoliciesClient) ListPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ApplicationGatewayWebApplicationFirewallPolicies", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client WebApplicationFirewallPoliciesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client WebApplicationFirewallPoliciesClient) ListResponder(resp *http.Response) (result WebApplicationFirewallPolicyListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client WebApplicationFirewallPoliciesClient) listNextResults(ctx context.Context, lastResults WebApplicationFirewallPolicyListResult) (result WebApplicationFirewallPolicyListResult, err error) {
	req, err := lastResults.webApplicationFirewallPolicyListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.WebApplicationFirewallPoliciesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.WebApplicationFirewallPoliciesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.WebApplicationFirewallPoliciesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client WebApplicationFirewallPoliciesClient) ListComplete(ctx context.Context, resourceGroupName string) (result WebApplicationFirewallPolicyListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WebApplicationFirewallPoliciesClient.List")
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

// ListAll gets all the WAF policies in a subscription.
func (client WebApplicationFirewallPoliciesClient) ListAll(ctx context.Context) (result WebApplicationFirewallPolicyListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WebApplicationFirewallPoliciesClient.ListAll")
		defer func() {
			sc := -1
			if result.wafplr.Response.Response != nil {
				sc = result.wafplr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listAllNextResults
	req, err := client.ListAllPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.WebApplicationFirewallPoliciesClient", "ListAll", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListAllSender(req)
	if err != nil {
		result.wafplr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.WebApplicationFirewallPoliciesClient", "ListAll", resp, "Failure sending request")
		return
	}

	result.wafplr, err = client.ListAllResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.WebApplicationFirewallPoliciesClient", "ListAll", resp, "Failure responding to request")
	}
	if result.wafplr.hasNextLink() && result.wafplr.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListAllPreparer prepares the ListAll request.
func (client WebApplicationFirewallPoliciesClient) ListAllPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Network/ApplicationGatewayWebApplicationFirewallPolicies", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListAllSender sends the ListAll request. The method will close the
// http.Response Body if it receives an error.
func (client WebApplicationFirewallPoliciesClient) ListAllSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListAllResponder handles the response to the ListAll request. The method always
// closes the http.Response Body.
func (client WebApplicationFirewallPoliciesClient) ListAllResponder(resp *http.Response) (result WebApplicationFirewallPolicyListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listAllNextResults retrieves the next set of results, if any.
func (client WebApplicationFirewallPoliciesClient) listAllNextResults(ctx context.Context, lastResults WebApplicationFirewallPolicyListResult) (result WebApplicationFirewallPolicyListResult, err error) {
	req, err := lastResults.webApplicationFirewallPolicyListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.WebApplicationFirewallPoliciesClient", "listAllNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListAllSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.WebApplicationFirewallPoliciesClient", "listAllNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListAllResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.WebApplicationFirewallPoliciesClient", "listAllNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListAllComplete enumerates all values, automatically crossing page boundaries as required.
func (client WebApplicationFirewallPoliciesClient) ListAllComplete(ctx context.Context) (result WebApplicationFirewallPolicyListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WebApplicationFirewallPoliciesClient.ListAll")
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
