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
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// SetDefinitionsClient is the to manage and control access to your resources, you can define customized policies and
// assign them at a scope.
type SetDefinitionsClient struct {
	BaseClient
}

// NewSetDefinitionsClient creates an instance of the SetDefinitionsClient client.
func NewSetDefinitionsClient(subscriptionID string) SetDefinitionsClient {
	return NewSetDefinitionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewSetDefinitionsClientWithBaseURI creates an instance of the SetDefinitionsClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewSetDefinitionsClientWithBaseURI(baseURI string, subscriptionID string) SetDefinitionsClient {
	return SetDefinitionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a policy set definition.
// Parameters:
// policySetDefinitionName - the name of the policy set definition to create.
// parameters - the policy set definition properties.
func (client SetDefinitionsClient) CreateOrUpdate(ctx context.Context, policySetDefinitionName string, parameters SetDefinition) (result SetDefinition, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SetDefinitionsClient.CreateOrUpdate")
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
			Constraints: []validation.Constraint{{Target: "parameters.SetDefinitionProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.SetDefinitionProperties.PolicyDefinitions", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("policy.SetDefinitionsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, policySetDefinitionName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client SetDefinitionsClient) CreateOrUpdatePreparer(ctx context.Context, policySetDefinitionName string, parameters SetDefinition) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"policySetDefinitionName": autorest.Encode("path", policySetDefinitionName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-06-01-preview"
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
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policySetDefinitions/{policySetDefinitionName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client SetDefinitionsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client SetDefinitionsClient) CreateOrUpdateResponder(resp *http.Response) (result SetDefinition, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CreateOrUpdateAtManagementGroup creates or updates a policy set definition at management group level.
// Parameters:
// policySetDefinitionName - the name of the policy set definition to create.
// parameters - the policy set definition properties.
// managementGroupID - the ID of the management group.
func (client SetDefinitionsClient) CreateOrUpdateAtManagementGroup(ctx context.Context, policySetDefinitionName string, parameters SetDefinition, managementGroupID string) (result SetDefinition, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SetDefinitionsClient.CreateOrUpdateAtManagementGroup")
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
			Constraints: []validation.Constraint{{Target: "parameters.SetDefinitionProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.SetDefinitionProperties.PolicyDefinitions", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("policy.SetDefinitionsClient", "CreateOrUpdateAtManagementGroup", err.Error())
	}

	req, err := client.CreateOrUpdateAtManagementGroupPreparer(ctx, policySetDefinitionName, parameters, managementGroupID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "CreateOrUpdateAtManagementGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateAtManagementGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "CreateOrUpdateAtManagementGroup", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateAtManagementGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "CreateOrUpdateAtManagementGroup", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdateAtManagementGroupPreparer prepares the CreateOrUpdateAtManagementGroup request.
func (client SetDefinitionsClient) CreateOrUpdateAtManagementGroupPreparer(ctx context.Context, policySetDefinitionName string, parameters SetDefinition, managementGroupID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"managementGroupId":       autorest.Encode("path", managementGroupID),
		"policySetDefinitionName": autorest.Encode("path", policySetDefinitionName),
	}

	const APIVersion = "2017-06-01-preview"
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
		autorest.WithPathParameters("/providers/Microsoft.Management/managementgroups/{managementGroupId}/providers/Microsoft.Authorization/policySetDefinitions/{policySetDefinitionName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateAtManagementGroupSender sends the CreateOrUpdateAtManagementGroup request. The method will close the
// http.Response Body if it receives an error.
func (client SetDefinitionsClient) CreateOrUpdateAtManagementGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateOrUpdateAtManagementGroupResponder handles the response to the CreateOrUpdateAtManagementGroup request. The method always
// closes the http.Response Body.
func (client SetDefinitionsClient) CreateOrUpdateAtManagementGroupResponder(resp *http.Response) (result SetDefinition, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a policy set definition.
// Parameters:
// policySetDefinitionName - the name of the policy set definition to delete.
func (client SetDefinitionsClient) Delete(ctx context.Context, policySetDefinitionName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SetDefinitionsClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, policySetDefinitionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client SetDefinitionsClient) DeletePreparer(ctx context.Context, policySetDefinitionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"policySetDefinitionName": autorest.Encode("path", policySetDefinitionName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policySetDefinitions/{policySetDefinitionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client SetDefinitionsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client SetDefinitionsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// DeleteAtManagementGroup deletes a policy set definition at management group level.
// Parameters:
// policySetDefinitionName - the name of the policy set definition to delete.
// managementGroupID - the ID of the management group.
func (client SetDefinitionsClient) DeleteAtManagementGroup(ctx context.Context, policySetDefinitionName string, managementGroupID string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SetDefinitionsClient.DeleteAtManagementGroup")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeleteAtManagementGroupPreparer(ctx, policySetDefinitionName, managementGroupID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "DeleteAtManagementGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteAtManagementGroupSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "DeleteAtManagementGroup", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteAtManagementGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "DeleteAtManagementGroup", resp, "Failure responding to request")
	}

	return
}

// DeleteAtManagementGroupPreparer prepares the DeleteAtManagementGroup request.
func (client SetDefinitionsClient) DeleteAtManagementGroupPreparer(ctx context.Context, policySetDefinitionName string, managementGroupID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"managementGroupId":       autorest.Encode("path", managementGroupID),
		"policySetDefinitionName": autorest.Encode("path", policySetDefinitionName),
	}

	const APIVersion = "2017-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Management/managementgroups/{managementGroupId}/providers/Microsoft.Authorization/policySetDefinitions/{policySetDefinitionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteAtManagementGroupSender sends the DeleteAtManagementGroup request. The method will close the
// http.Response Body if it receives an error.
func (client SetDefinitionsClient) DeleteAtManagementGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteAtManagementGroupResponder handles the response to the DeleteAtManagementGroup request. The method always
// closes the http.Response Body.
func (client SetDefinitionsClient) DeleteAtManagementGroupResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the policy set definition.
// Parameters:
// policySetDefinitionName - the name of the policy set definition to get.
func (client SetDefinitionsClient) Get(ctx context.Context, policySetDefinitionName string) (result SetDefinition, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SetDefinitionsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, policySetDefinitionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client SetDefinitionsClient) GetPreparer(ctx context.Context, policySetDefinitionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"policySetDefinitionName": autorest.Encode("path", policySetDefinitionName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policySetDefinitions/{policySetDefinitionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client SetDefinitionsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client SetDefinitionsClient) GetResponder(resp *http.Response) (result SetDefinition, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetAtManagementGroup gets the policy set definition at management group level.
// Parameters:
// policySetDefinitionName - the name of the policy set definition to get.
// managementGroupID - the ID of the management group.
func (client SetDefinitionsClient) GetAtManagementGroup(ctx context.Context, policySetDefinitionName string, managementGroupID string) (result SetDefinition, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SetDefinitionsClient.GetAtManagementGroup")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetAtManagementGroupPreparer(ctx, policySetDefinitionName, managementGroupID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "GetAtManagementGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetAtManagementGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "GetAtManagementGroup", resp, "Failure sending request")
		return
	}

	result, err = client.GetAtManagementGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "GetAtManagementGroup", resp, "Failure responding to request")
	}

	return
}

// GetAtManagementGroupPreparer prepares the GetAtManagementGroup request.
func (client SetDefinitionsClient) GetAtManagementGroupPreparer(ctx context.Context, policySetDefinitionName string, managementGroupID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"managementGroupId":       autorest.Encode("path", managementGroupID),
		"policySetDefinitionName": autorest.Encode("path", policySetDefinitionName),
	}

	const APIVersion = "2017-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Management/managementgroups/{managementGroupId}/providers/Microsoft.Authorization/policySetDefinitions/{policySetDefinitionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetAtManagementGroupSender sends the GetAtManagementGroup request. The method will close the
// http.Response Body if it receives an error.
func (client SetDefinitionsClient) GetAtManagementGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetAtManagementGroupResponder handles the response to the GetAtManagementGroup request. The method always
// closes the http.Response Body.
func (client SetDefinitionsClient) GetAtManagementGroupResponder(resp *http.Response) (result SetDefinition, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetBuiltIn gets the built in policy set definition.
// Parameters:
// policySetDefinitionName - the name of the policy set definition to get.
func (client SetDefinitionsClient) GetBuiltIn(ctx context.Context, policySetDefinitionName string) (result SetDefinition, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SetDefinitionsClient.GetBuiltIn")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetBuiltInPreparer(ctx, policySetDefinitionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "GetBuiltIn", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetBuiltInSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "GetBuiltIn", resp, "Failure sending request")
		return
	}

	result, err = client.GetBuiltInResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "GetBuiltIn", resp, "Failure responding to request")
	}

	return
}

// GetBuiltInPreparer prepares the GetBuiltIn request.
func (client SetDefinitionsClient) GetBuiltInPreparer(ctx context.Context, policySetDefinitionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"policySetDefinitionName": autorest.Encode("path", policySetDefinitionName),
	}

	const APIVersion = "2017-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Authorization/policySetDefinitions/{policySetDefinitionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetBuiltInSender sends the GetBuiltIn request. The method will close the
// http.Response Body if it receives an error.
func (client SetDefinitionsClient) GetBuiltInSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetBuiltInResponder handles the response to the GetBuiltIn request. The method always
// closes the http.Response Body.
func (client SetDefinitionsClient) GetBuiltInResponder(resp *http.Response) (result SetDefinition, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets all the policy set definitions for a subscription.
func (client SetDefinitionsClient) List(ctx context.Context) (result SetDefinitionListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SetDefinitionsClient.List")
		defer func() {
			sc := -1
			if result.sdlr.Response.Response != nil {
				sc = result.sdlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.sdlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "List", resp, "Failure sending request")
		return
	}

	result.sdlr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client SetDefinitionsClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policySetDefinitions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client SetDefinitionsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client SetDefinitionsClient) ListResponder(resp *http.Response) (result SetDefinitionListResult, err error) {
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
func (client SetDefinitionsClient) listNextResults(ctx context.Context, lastResults SetDefinitionListResult) (result SetDefinitionListResult, err error) {
	req, err := lastResults.setDefinitionListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client SetDefinitionsClient) ListComplete(ctx context.Context) (result SetDefinitionListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SetDefinitionsClient.List")
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

// ListBuiltIn gets all the built in policy set definitions.
func (client SetDefinitionsClient) ListBuiltIn(ctx context.Context) (result SetDefinitionListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SetDefinitionsClient.ListBuiltIn")
		defer func() {
			sc := -1
			if result.sdlr.Response.Response != nil {
				sc = result.sdlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listBuiltInNextResults
	req, err := client.ListBuiltInPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "ListBuiltIn", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBuiltInSender(req)
	if err != nil {
		result.sdlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "ListBuiltIn", resp, "Failure sending request")
		return
	}

	result.sdlr, err = client.ListBuiltInResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "ListBuiltIn", resp, "Failure responding to request")
	}

	return
}

// ListBuiltInPreparer prepares the ListBuiltIn request.
func (client SetDefinitionsClient) ListBuiltInPreparer(ctx context.Context) (*http.Request, error) {
	const APIVersion = "2017-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Authorization/policySetDefinitions"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListBuiltInSender sends the ListBuiltIn request. The method will close the
// http.Response Body if it receives an error.
func (client SetDefinitionsClient) ListBuiltInSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListBuiltInResponder handles the response to the ListBuiltIn request. The method always
// closes the http.Response Body.
func (client SetDefinitionsClient) ListBuiltInResponder(resp *http.Response) (result SetDefinitionListResult, err error) {
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
func (client SetDefinitionsClient) listBuiltInNextResults(ctx context.Context, lastResults SetDefinitionListResult) (result SetDefinitionListResult, err error) {
	req, err := lastResults.setDefinitionListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "listBuiltInNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListBuiltInSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "listBuiltInNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListBuiltInResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "listBuiltInNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListBuiltInComplete enumerates all values, automatically crossing page boundaries as required.
func (client SetDefinitionsClient) ListBuiltInComplete(ctx context.Context) (result SetDefinitionListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SetDefinitionsClient.ListBuiltIn")
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

// ListByManagementGroup gets all the policy set definitions for a subscription at management group.
// Parameters:
// managementGroupID - the ID of the management group.
func (client SetDefinitionsClient) ListByManagementGroup(ctx context.Context, managementGroupID string) (result SetDefinitionListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SetDefinitionsClient.ListByManagementGroup")
		defer func() {
			sc := -1
			if result.sdlr.Response.Response != nil {
				sc = result.sdlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByManagementGroupNextResults
	req, err := client.ListByManagementGroupPreparer(ctx, managementGroupID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "ListByManagementGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByManagementGroupSender(req)
	if err != nil {
		result.sdlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "ListByManagementGroup", resp, "Failure sending request")
		return
	}

	result.sdlr, err = client.ListByManagementGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "ListByManagementGroup", resp, "Failure responding to request")
	}

	return
}

// ListByManagementGroupPreparer prepares the ListByManagementGroup request.
func (client SetDefinitionsClient) ListByManagementGroupPreparer(ctx context.Context, managementGroupID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"managementGroupId": autorest.Encode("path", managementGroupID),
	}

	const APIVersion = "2017-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Management/managementgroups/{managementGroupId}/providers/Microsoft.Authorization/policySetDefinitions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByManagementGroupSender sends the ListByManagementGroup request. The method will close the
// http.Response Body if it receives an error.
func (client SetDefinitionsClient) ListByManagementGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListByManagementGroupResponder handles the response to the ListByManagementGroup request. The method always
// closes the http.Response Body.
func (client SetDefinitionsClient) ListByManagementGroupResponder(resp *http.Response) (result SetDefinitionListResult, err error) {
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
func (client SetDefinitionsClient) listByManagementGroupNextResults(ctx context.Context, lastResults SetDefinitionListResult) (result SetDefinitionListResult, err error) {
	req, err := lastResults.setDefinitionListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "listByManagementGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByManagementGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "listByManagementGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByManagementGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.SetDefinitionsClient", "listByManagementGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByManagementGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client SetDefinitionsClient) ListByManagementGroupComplete(ctx context.Context, managementGroupID string) (result SetDefinitionListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SetDefinitionsClient.ListByManagementGroup")
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
