package policy

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

// DefinitionsClient is the to manage and control access to your resources, you can define customized policies and
// assign them at a scope.
type DefinitionsClient struct {
	BaseClient
}

// NewDefinitionsClient creates an instance of the DefinitionsClient client.
func NewDefinitionsClient(subscriptionID string) DefinitionsClient {
	return NewDefinitionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewDefinitionsClientWithBaseURI creates an instance of the DefinitionsClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewDefinitionsClientWithBaseURI(baseURI string, subscriptionID string) DefinitionsClient {
	return DefinitionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate this operation creates or updates a policy definition in the given subscription with the given name.
// Parameters:
// policyDefinitionName - the name of the policy definition to create.
// parameters - the policy definition properties.
func (client DefinitionsClient) CreateOrUpdate(ctx context.Context, policyDefinitionName string, parameters Definition) (result Definition, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DefinitionsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, policyDefinitionName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client DefinitionsClient) CreateOrUpdatePreparer(ctx context.Context, policyDefinitionName string, parameters Definition) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"policyDefinitionName": autorest.Encode("path", policyDefinitionName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	parameters.ID = nil
	parameters.Name = nil
	parameters.Type = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policyDefinitions/{policyDefinitionName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client DefinitionsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client DefinitionsClient) CreateOrUpdateResponder(resp *http.Response) (result Definition, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CreateOrUpdateAtManagementGroup this operation creates or updates a policy definition in the given management group
// with the given name.
// Parameters:
// policyDefinitionName - the name of the policy definition to create.
// parameters - the policy definition properties.
// managementGroupID - the ID of the management group.
func (client DefinitionsClient) CreateOrUpdateAtManagementGroup(ctx context.Context, policyDefinitionName string, parameters Definition, managementGroupID string) (result Definition, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DefinitionsClient.CreateOrUpdateAtManagementGroup")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdateAtManagementGroupPreparer(ctx, policyDefinitionName, parameters, managementGroupID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "CreateOrUpdateAtManagementGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateAtManagementGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "CreateOrUpdateAtManagementGroup", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateAtManagementGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "CreateOrUpdateAtManagementGroup", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdateAtManagementGroupPreparer prepares the CreateOrUpdateAtManagementGroup request.
func (client DefinitionsClient) CreateOrUpdateAtManagementGroupPreparer(ctx context.Context, policyDefinitionName string, parameters Definition, managementGroupID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"managementGroupId":    autorest.Encode("path", managementGroupID),
		"policyDefinitionName": autorest.Encode("path", policyDefinitionName),
	}

	const APIVersion = "2018-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	parameters.ID = nil
	parameters.Name = nil
	parameters.Type = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Management/managementgroups/{managementGroupId}/providers/Microsoft.Authorization/policyDefinitions/{policyDefinitionName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateAtManagementGroupSender sends the CreateOrUpdateAtManagementGroup request. The method will close the
// http.Response Body if it receives an error.
func (client DefinitionsClient) CreateOrUpdateAtManagementGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateOrUpdateAtManagementGroupResponder handles the response to the CreateOrUpdateAtManagementGroup request. The method always
// closes the http.Response Body.
func (client DefinitionsClient) CreateOrUpdateAtManagementGroupResponder(resp *http.Response) (result Definition, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete this operation deletes the policy definition in the given subscription with the given name.
// Parameters:
// policyDefinitionName - the name of the policy definition to delete.
func (client DefinitionsClient) Delete(ctx context.Context, policyDefinitionName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DefinitionsClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, policyDefinitionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client DefinitionsClient) DeletePreparer(ctx context.Context, policyDefinitionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"policyDefinitionName": autorest.Encode("path", policyDefinitionName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policyDefinitions/{policyDefinitionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client DefinitionsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client DefinitionsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// DeleteAtManagementGroup this operation deletes the policy definition in the given management group with the given
// name.
// Parameters:
// policyDefinitionName - the name of the policy definition to delete.
// managementGroupID - the ID of the management group.
func (client DefinitionsClient) DeleteAtManagementGroup(ctx context.Context, policyDefinitionName string, managementGroupID string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DefinitionsClient.DeleteAtManagementGroup")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeleteAtManagementGroupPreparer(ctx, policyDefinitionName, managementGroupID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "DeleteAtManagementGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteAtManagementGroupSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "DeleteAtManagementGroup", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteAtManagementGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "DeleteAtManagementGroup", resp, "Failure responding to request")
	}

	return
}

// DeleteAtManagementGroupPreparer prepares the DeleteAtManagementGroup request.
func (client DefinitionsClient) DeleteAtManagementGroupPreparer(ctx context.Context, policyDefinitionName string, managementGroupID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"managementGroupId":    autorest.Encode("path", managementGroupID),
		"policyDefinitionName": autorest.Encode("path", policyDefinitionName),
	}

	const APIVersion = "2018-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Management/managementgroups/{managementGroupId}/providers/Microsoft.Authorization/policyDefinitions/{policyDefinitionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteAtManagementGroupSender sends the DeleteAtManagementGroup request. The method will close the
// http.Response Body if it receives an error.
func (client DefinitionsClient) DeleteAtManagementGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteAtManagementGroupResponder handles the response to the DeleteAtManagementGroup request. The method always
// closes the http.Response Body.
func (client DefinitionsClient) DeleteAtManagementGroupResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get this operation retrieves the policy definition in the given subscription with the given name.
// Parameters:
// policyDefinitionName - the name of the policy definition to get.
func (client DefinitionsClient) Get(ctx context.Context, policyDefinitionName string) (result Definition, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DefinitionsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, policyDefinitionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client DefinitionsClient) GetPreparer(ctx context.Context, policyDefinitionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"policyDefinitionName": autorest.Encode("path", policyDefinitionName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policyDefinitions/{policyDefinitionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client DefinitionsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client DefinitionsClient) GetResponder(resp *http.Response) (result Definition, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetAtManagementGroup this operation retrieves the policy definition in the given management group with the given
// name.
// Parameters:
// policyDefinitionName - the name of the policy definition to get.
// managementGroupID - the ID of the management group.
func (client DefinitionsClient) GetAtManagementGroup(ctx context.Context, policyDefinitionName string, managementGroupID string) (result Definition, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DefinitionsClient.GetAtManagementGroup")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetAtManagementGroupPreparer(ctx, policyDefinitionName, managementGroupID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "GetAtManagementGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetAtManagementGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "GetAtManagementGroup", resp, "Failure sending request")
		return
	}

	result, err = client.GetAtManagementGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "GetAtManagementGroup", resp, "Failure responding to request")
	}

	return
}

// GetAtManagementGroupPreparer prepares the GetAtManagementGroup request.
func (client DefinitionsClient) GetAtManagementGroupPreparer(ctx context.Context, policyDefinitionName string, managementGroupID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"managementGroupId":    autorest.Encode("path", managementGroupID),
		"policyDefinitionName": autorest.Encode("path", policyDefinitionName),
	}

	const APIVersion = "2018-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Management/managementgroups/{managementGroupId}/providers/Microsoft.Authorization/policyDefinitions/{policyDefinitionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetAtManagementGroupSender sends the GetAtManagementGroup request. The method will close the
// http.Response Body if it receives an error.
func (client DefinitionsClient) GetAtManagementGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetAtManagementGroupResponder handles the response to the GetAtManagementGroup request. The method always
// closes the http.Response Body.
func (client DefinitionsClient) GetAtManagementGroupResponder(resp *http.Response) (result Definition, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetBuiltIn this operation retrieves the built-in policy definition with the given name.
// Parameters:
// policyDefinitionName - the name of the built-in policy definition to get.
func (client DefinitionsClient) GetBuiltIn(ctx context.Context, policyDefinitionName string) (result Definition, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DefinitionsClient.GetBuiltIn")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetBuiltInPreparer(ctx, policyDefinitionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "GetBuiltIn", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetBuiltInSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "GetBuiltIn", resp, "Failure sending request")
		return
	}

	result, err = client.GetBuiltInResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "GetBuiltIn", resp, "Failure responding to request")
	}

	return
}

// GetBuiltInPreparer prepares the GetBuiltIn request.
func (client DefinitionsClient) GetBuiltInPreparer(ctx context.Context, policyDefinitionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"policyDefinitionName": autorest.Encode("path", policyDefinitionName),
	}

	const APIVersion = "2018-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Authorization/policyDefinitions/{policyDefinitionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetBuiltInSender sends the GetBuiltIn request. The method will close the
// http.Response Body if it receives an error.
func (client DefinitionsClient) GetBuiltInSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetBuiltInResponder handles the response to the GetBuiltIn request. The method always
// closes the http.Response Body.
func (client DefinitionsClient) GetBuiltInResponder(resp *http.Response) (result Definition, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List this operation retrieves a list of all the policy definitions in a given subscription.
func (client DefinitionsClient) List(ctx context.Context) (result DefinitionListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DefinitionsClient.List")
		defer func() {
			sc := -1
			if result.dlr.Response.Response != nil {
				sc = result.dlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.dlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "List", resp, "Failure sending request")
		return
	}

	result.dlr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client DefinitionsClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policyDefinitions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client DefinitionsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client DefinitionsClient) ListResponder(resp *http.Response) (result DefinitionListResult, err error) {
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
func (client DefinitionsClient) listNextResults(ctx context.Context, lastResults DefinitionListResult) (result DefinitionListResult, err error) {
	req, err := lastResults.definitionListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "policy.DefinitionsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "policy.DefinitionsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client DefinitionsClient) ListComplete(ctx context.Context) (result DefinitionListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DefinitionsClient.List")
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

// ListBuiltIn this operation retrieves a list of all the built-in policy definitions.
func (client DefinitionsClient) ListBuiltIn(ctx context.Context) (result DefinitionListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DefinitionsClient.ListBuiltIn")
		defer func() {
			sc := -1
			if result.dlr.Response.Response != nil {
				sc = result.dlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listBuiltInNextResults
	req, err := client.ListBuiltInPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "ListBuiltIn", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBuiltInSender(req)
	if err != nil {
		result.dlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "ListBuiltIn", resp, "Failure sending request")
		return
	}

	result.dlr, err = client.ListBuiltInResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "ListBuiltIn", resp, "Failure responding to request")
	}

	return
}

// ListBuiltInPreparer prepares the ListBuiltIn request.
func (client DefinitionsClient) ListBuiltInPreparer(ctx context.Context) (*http.Request, error) {
	const APIVersion = "2018-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Authorization/policyDefinitions"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListBuiltInSender sends the ListBuiltIn request. The method will close the
// http.Response Body if it receives an error.
func (client DefinitionsClient) ListBuiltInSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListBuiltInResponder handles the response to the ListBuiltIn request. The method always
// closes the http.Response Body.
func (client DefinitionsClient) ListBuiltInResponder(resp *http.Response) (result DefinitionListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listBuiltInNextResults retrieves the next set of results, if any.
func (client DefinitionsClient) listBuiltInNextResults(ctx context.Context, lastResults DefinitionListResult) (result DefinitionListResult, err error) {
	req, err := lastResults.definitionListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "policy.DefinitionsClient", "listBuiltInNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListBuiltInSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "policy.DefinitionsClient", "listBuiltInNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListBuiltInResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "listBuiltInNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListBuiltInComplete enumerates all values, automatically crossing page boundaries as required.
func (client DefinitionsClient) ListBuiltInComplete(ctx context.Context) (result DefinitionListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DefinitionsClient.ListBuiltIn")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListBuiltIn(ctx)
	return
}

// ListByManagementGroup this operation retrieves a list of all the policy definitions in a given management group.
// Parameters:
// managementGroupID - the ID of the management group.
func (client DefinitionsClient) ListByManagementGroup(ctx context.Context, managementGroupID string) (result DefinitionListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DefinitionsClient.ListByManagementGroup")
		defer func() {
			sc := -1
			if result.dlr.Response.Response != nil {
				sc = result.dlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByManagementGroupNextResults
	req, err := client.ListByManagementGroupPreparer(ctx, managementGroupID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "ListByManagementGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByManagementGroupSender(req)
	if err != nil {
		result.dlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "ListByManagementGroup", resp, "Failure sending request")
		return
	}

	result.dlr, err = client.ListByManagementGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "ListByManagementGroup", resp, "Failure responding to request")
	}

	return
}

// ListByManagementGroupPreparer prepares the ListByManagementGroup request.
func (client DefinitionsClient) ListByManagementGroupPreparer(ctx context.Context, managementGroupID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"managementGroupId": autorest.Encode("path", managementGroupID),
	}

	const APIVersion = "2018-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Management/managementgroups/{managementGroupId}/providers/Microsoft.Authorization/policyDefinitions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByManagementGroupSender sends the ListByManagementGroup request. The method will close the
// http.Response Body if it receives an error.
func (client DefinitionsClient) ListByManagementGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListByManagementGroupResponder handles the response to the ListByManagementGroup request. The method always
// closes the http.Response Body.
func (client DefinitionsClient) ListByManagementGroupResponder(resp *http.Response) (result DefinitionListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByManagementGroupNextResults retrieves the next set of results, if any.
func (client DefinitionsClient) listByManagementGroupNextResults(ctx context.Context, lastResults DefinitionListResult) (result DefinitionListResult, err error) {
	req, err := lastResults.definitionListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "policy.DefinitionsClient", "listByManagementGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByManagementGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "policy.DefinitionsClient", "listByManagementGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByManagementGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "listByManagementGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByManagementGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client DefinitionsClient) ListByManagementGroupComplete(ctx context.Context, managementGroupID string) (result DefinitionListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DefinitionsClient.ListByManagementGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByManagementGroup(ctx, managementGroupID)
	return
}
