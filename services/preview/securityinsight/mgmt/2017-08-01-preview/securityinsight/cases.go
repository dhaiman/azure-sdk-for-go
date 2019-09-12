package securityinsight

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

// CasesClient is the API spec for Microsoft.SecurityInsights (Azure Security Insights) resource provider
type CasesClient struct {
	BaseClient
}

// NewCasesClient creates an instance of the CasesClient client.
func NewCasesClient(subscriptionID string) CasesClient {
	return NewCasesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewCasesClientWithBaseURI creates an instance of the CasesClient client.
func NewCasesClientWithBaseURI(baseURI string, subscriptionID string) CasesClient {
	return CasesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates the case.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// operationalInsightsResourceProvider - the namespace of workspaces resource provider-
// Microsoft.OperationalInsights.
// workspaceName - the name of the workspace.
// caseID - case ID
// caseParameter - the case
func (client CasesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, caseID string, caseParameter Case) (result Case, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CasesClient.CreateOrUpdate")
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
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: caseParameter,
			Constraints: []validation.Constraint{{Target: "caseParameter.CaseProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "caseParameter.CaseProperties.Owner", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "caseParameter.CaseProperties.Owner.ObjectID", Name: validation.Null, Rule: true, Chain: nil}}},
					{Target: "caseParameter.CaseProperties.StartTimeUtc", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "caseParameter.CaseProperties.Title", Name: validation.Null, Rule: true, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("securityinsight.CasesClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, operationalInsightsResourceProvider, workspaceName, caseID, caseParameter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.CasesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "securityinsight.CasesClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.CasesClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client CasesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, caseID string, caseParameter Case) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"caseId":                              autorest.Encode("path", caseID),
		"operationalInsightsResourceProvider": autorest.Encode("path", operationalInsightsResourceProvider),
		"resourceGroupName":                   autorest.Encode("path", resourceGroupName),
		"subscriptionId":                      autorest.Encode("path", client.SubscriptionID),
		"workspaceName":                       autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2019-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{operationalInsightsResourceProvider}/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/cases/{caseId}", pathParameters),
		autorest.WithJSON(caseParameter),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client CasesClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client CasesClient) CreateOrUpdateResponder(resp *http.Response) (result Case, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete the case
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// operationalInsightsResourceProvider - the namespace of workspaces resource provider-
// Microsoft.OperationalInsights.
// workspaceName - the name of the workspace.
// caseID - case ID
func (client CasesClient) Delete(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, caseID string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CasesClient.Delete")
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
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("securityinsight.CasesClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, operationalInsightsResourceProvider, workspaceName, caseID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.CasesClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "securityinsight.CasesClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.CasesClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client CasesClient) DeletePreparer(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, caseID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"caseId":                              autorest.Encode("path", caseID),
		"operationalInsightsResourceProvider": autorest.Encode("path", operationalInsightsResourceProvider),
		"resourceGroupName":                   autorest.Encode("path", resourceGroupName),
		"subscriptionId":                      autorest.Encode("path", client.SubscriptionID),
		"workspaceName":                       autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2019-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{operationalInsightsResourceProvider}/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/cases/{caseId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client CasesClient) DeleteSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client CasesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets a case.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// operationalInsightsResourceProvider - the namespace of workspaces resource provider-
// Microsoft.OperationalInsights.
// workspaceName - the name of the workspace.
// caseID - case ID
func (client CasesClient) Get(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, caseID string) (result Case, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CasesClient.Get")
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
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("securityinsight.CasesClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, operationalInsightsResourceProvider, workspaceName, caseID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.CasesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "securityinsight.CasesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.CasesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client CasesClient) GetPreparer(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, caseID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"caseId":                              autorest.Encode("path", caseID),
		"operationalInsightsResourceProvider": autorest.Encode("path", operationalInsightsResourceProvider),
		"resourceGroupName":                   autorest.Encode("path", resourceGroupName),
		"subscriptionId":                      autorest.Encode("path", client.SubscriptionID),
		"workspaceName":                       autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2019-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{operationalInsightsResourceProvider}/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/cases/{caseId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client CasesClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client CasesClient) GetResponder(resp *http.Response) (result Case, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetComment gets a case comment
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// operationalInsightsResourceProvider - the namespace of workspaces resource provider-
// Microsoft.OperationalInsights.
// workspaceName - the name of the workspace.
// caseID - case ID
// caseCommentID - case comment ID
func (client CasesClient) GetComment(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, caseID string, caseCommentID string) (result CaseComment, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CasesClient.GetComment")
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
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("securityinsight.CasesClient", "GetComment", err.Error())
	}

	req, err := client.GetCommentPreparer(ctx, resourceGroupName, operationalInsightsResourceProvider, workspaceName, caseID, caseCommentID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.CasesClient", "GetComment", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetCommentSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "securityinsight.CasesClient", "GetComment", resp, "Failure sending request")
		return
	}

	result, err = client.GetCommentResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.CasesClient", "GetComment", resp, "Failure responding to request")
	}

	return
}

// GetCommentPreparer prepares the GetComment request.
func (client CasesClient) GetCommentPreparer(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, caseID string, caseCommentID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"caseCommentId":                       autorest.Encode("path", caseCommentID),
		"caseId":                              autorest.Encode("path", caseID),
		"operationalInsightsResourceProvider": autorest.Encode("path", operationalInsightsResourceProvider),
		"resourceGroupName":                   autorest.Encode("path", resourceGroupName),
		"subscriptionId":                      autorest.Encode("path", client.SubscriptionID),
		"workspaceName":                       autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2019-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{operationalInsightsResourceProvider}/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/cases/{caseId}/comments/{caseCommentId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetCommentSender sends the GetComment request. The method will close the
// http.Response Body if it receives an error.
func (client CasesClient) GetCommentSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetCommentResponder handles the response to the GetComment request. The method always
// closes the http.Response Body.
func (client CasesClient) GetCommentResponder(resp *http.Response) (result CaseComment, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets all cases.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// operationalInsightsResourceProvider - the namespace of workspaces resource provider-
// Microsoft.OperationalInsights.
// workspaceName - the name of the workspace.
// filter - filters the results, based on a Boolean condition. Optional.
// orderby - sorts the results. Optional.
// top - returns only the first n results. Optional.
// skipToken - skiptoken is only used if a previous operation returned a partial result. If a previous response
// contains a nextLink element, the value of the nextLink element will include a skiptoken parameter that
// specifies a starting point to use for subsequent calls. Optional.
func (client CasesClient) List(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, filter string, orderby string, top *int32, skipToken string) (result CaseListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CasesClient.List")
		defer func() {
			sc := -1
			if result.cl.Response.Response != nil {
				sc = result.cl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("securityinsight.CasesClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, operationalInsightsResourceProvider, workspaceName, filter, orderby, top, skipToken)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.CasesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.cl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "securityinsight.CasesClient", "List", resp, "Failure sending request")
		return
	}

	result.cl, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.CasesClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client CasesClient) ListPreparer(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, filter string, orderby string, top *int32, skipToken string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"operationalInsightsResourceProvider": autorest.Encode("path", operationalInsightsResourceProvider),
		"resourceGroupName":                   autorest.Encode("path", resourceGroupName),
		"subscriptionId":                      autorest.Encode("path", client.SubscriptionID),
		"workspaceName":                       autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2019-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if len(orderby) > 0 {
		queryParameters["$orderby"] = autorest.Encode("query", orderby)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(skipToken) > 0 {
		queryParameters["$skipToken"] = autorest.Encode("query", skipToken)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{operationalInsightsResourceProvider}/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/cases", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client CasesClient) ListSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client CasesClient) ListResponder(resp *http.Response) (result CaseList, err error) {
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
func (client CasesClient) listNextResults(ctx context.Context, lastResults CaseList) (result CaseList, err error) {
	req, err := lastResults.caseListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "securityinsight.CasesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "securityinsight.CasesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.CasesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client CasesClient) ListComplete(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, filter string, orderby string, top *int32, skipToken string) (result CaseListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CasesClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, operationalInsightsResourceProvider, workspaceName, filter, orderby, top, skipToken)
	return
}
