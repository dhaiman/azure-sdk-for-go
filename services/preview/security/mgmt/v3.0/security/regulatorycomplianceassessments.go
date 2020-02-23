package security

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

// RegulatoryComplianceAssessmentsClient is the API spec for Microsoft.Security (Azure Security Center) resource
// provider
type RegulatoryComplianceAssessmentsClient struct {
	BaseClient
}

// NewRegulatoryComplianceAssessmentsClient creates an instance of the RegulatoryComplianceAssessmentsClient client.
func NewRegulatoryComplianceAssessmentsClient(subscriptionID string, ascLocation string) RegulatoryComplianceAssessmentsClient {
	return NewRegulatoryComplianceAssessmentsClientWithBaseURI(DefaultBaseURI, subscriptionID, ascLocation)
}

// NewRegulatoryComplianceAssessmentsClientWithBaseURI creates an instance of the RegulatoryComplianceAssessmentsClient
// client using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI
// (sovereign clouds, Azure stack).
func NewRegulatoryComplianceAssessmentsClientWithBaseURI(baseURI string, subscriptionID string, ascLocation string) RegulatoryComplianceAssessmentsClient {
	return RegulatoryComplianceAssessmentsClient{NewWithBaseURI(baseURI, subscriptionID, ascLocation)}
}

// Get supported regulatory compliance details and state for selected assessment
// Parameters:
// regulatoryComplianceStandardName - name of the regulatory compliance standard object
// regulatoryComplianceControlName - name of the regulatory compliance control object
// regulatoryComplianceAssessmentName - name of the regulatory compliance assessment object
func (client RegulatoryComplianceAssessmentsClient) Get(ctx context.Context, regulatoryComplianceStandardName string, regulatoryComplianceControlName string, regulatoryComplianceAssessmentName string) (result RegulatoryComplianceAssessment, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RegulatoryComplianceAssessmentsClient.Get")
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
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("security.RegulatoryComplianceAssessmentsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, regulatoryComplianceStandardName, regulatoryComplianceControlName, regulatoryComplianceAssessmentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.RegulatoryComplianceAssessmentsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "security.RegulatoryComplianceAssessmentsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.RegulatoryComplianceAssessmentsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client RegulatoryComplianceAssessmentsClient) GetPreparer(ctx context.Context, regulatoryComplianceStandardName string, regulatoryComplianceControlName string, regulatoryComplianceAssessmentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"regulatoryComplianceAssessmentName": autorest.Encode("path", regulatoryComplianceAssessmentName),
		"regulatoryComplianceControlName":    autorest.Encode("path", regulatoryComplianceControlName),
		"regulatoryComplianceStandardName":   autorest.Encode("path", regulatoryComplianceStandardName),
		"subscriptionId":                     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Security/regulatoryComplianceStandards/{regulatoryComplianceStandardName}/regulatoryComplianceControls/{regulatoryComplianceControlName}/regulatoryComplianceAssessments/{regulatoryComplianceAssessmentName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client RegulatoryComplianceAssessmentsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client RegulatoryComplianceAssessmentsClient) GetResponder(resp *http.Response) (result RegulatoryComplianceAssessment, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List details and state of assessments mapped to selected regulatory compliance control
// Parameters:
// regulatoryComplianceStandardName - name of the regulatory compliance standard object
// regulatoryComplianceControlName - name of the regulatory compliance control object
// filter - oData filter. Optional.
func (client RegulatoryComplianceAssessmentsClient) List(ctx context.Context, regulatoryComplianceStandardName string, regulatoryComplianceControlName string, filter string) (result RegulatoryComplianceAssessmentListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RegulatoryComplianceAssessmentsClient.List")
		defer func() {
			sc := -1
			if result.rcal.Response.Response != nil {
				sc = result.rcal.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("security.RegulatoryComplianceAssessmentsClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, regulatoryComplianceStandardName, regulatoryComplianceControlName, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.RegulatoryComplianceAssessmentsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.rcal.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "security.RegulatoryComplianceAssessmentsClient", "List", resp, "Failure sending request")
		return
	}

	result.rcal, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.RegulatoryComplianceAssessmentsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client RegulatoryComplianceAssessmentsClient) ListPreparer(ctx context.Context, regulatoryComplianceStandardName string, regulatoryComplianceControlName string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"regulatoryComplianceControlName":  autorest.Encode("path", regulatoryComplianceControlName),
		"regulatoryComplianceStandardName": autorest.Encode("path", regulatoryComplianceStandardName),
		"subscriptionId":                   autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Security/regulatoryComplianceStandards/{regulatoryComplianceStandardName}/regulatoryComplianceControls/{regulatoryComplianceControlName}/regulatoryComplianceAssessments", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client RegulatoryComplianceAssessmentsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client RegulatoryComplianceAssessmentsClient) ListResponder(resp *http.Response) (result RegulatoryComplianceAssessmentList, err error) {
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
func (client RegulatoryComplianceAssessmentsClient) listNextResults(ctx context.Context, lastResults RegulatoryComplianceAssessmentList) (result RegulatoryComplianceAssessmentList, err error) {
	req, err := lastResults.regulatoryComplianceAssessmentListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "security.RegulatoryComplianceAssessmentsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "security.RegulatoryComplianceAssessmentsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.RegulatoryComplianceAssessmentsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client RegulatoryComplianceAssessmentsClient) ListComplete(ctx context.Context, regulatoryComplianceStandardName string, regulatoryComplianceControlName string, filter string) (result RegulatoryComplianceAssessmentListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RegulatoryComplianceAssessmentsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, regulatoryComplianceStandardName, regulatoryComplianceControlName, filter)
	return
}
