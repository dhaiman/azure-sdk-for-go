package advisor

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
	"github.com/satori/go.uuid"
	"net/http"
)

// RecommendationsClient is the REST APIs for Azure Advisor
type RecommendationsClient struct {
	BaseClient
}

// NewRecommendationsClient creates an instance of the RecommendationsClient client.
func NewRecommendationsClient(subscriptionID string) RecommendationsClient {
	return NewRecommendationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewRecommendationsClientWithBaseURI creates an instance of the RecommendationsClient client.
func NewRecommendationsClientWithBaseURI(baseURI string, subscriptionID string) RecommendationsClient {
	return RecommendationsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Generate initiates the recommendation generation or computation process for a subscription. This operation is
// asynchronous. The generated recommendations are stored in a cache in the Advisor service.
func (client RecommendationsClient) Generate(ctx context.Context) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RecommendationsClient.Generate")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GeneratePreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "advisor.RecommendationsClient", "Generate", nil, "Failure preparing request")
		return
	}

	resp, err := client.GenerateSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "advisor.RecommendationsClient", "Generate", resp, "Failure sending request")
		return
	}

	result, err = client.GenerateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "advisor.RecommendationsClient", "Generate", resp, "Failure responding to request")
	}

	return
}

// GeneratePreparer prepares the Generate request.
func (client RecommendationsClient) GeneratePreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-19"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Advisor/generateRecommendations", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GenerateSender sends the Generate request. The method will close the
// http.Response Body if it receives an error.
func (client RecommendationsClient) GenerateSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GenerateResponder handles the response to the Generate request. The method always
// closes the http.Response Body.
func (client RecommendationsClient) GenerateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get obtains details of a cached recommendation.
// Parameters:
// resourceURI - the fully qualified Azure Resource Manager identifier of the resource to which the
// recommendation applies.
// recommendationID - the recommendation ID.
func (client RecommendationsClient) Get(ctx context.Context, resourceURI string, recommendationID string) (result ResourceRecommendationBase, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RecommendationsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceURI, recommendationID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "advisor.RecommendationsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "advisor.RecommendationsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "advisor.RecommendationsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client RecommendationsClient) GetPreparer(ctx context.Context, resourceURI string, recommendationID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"recommendationId": autorest.Encode("path", recommendationID),
		"resourceUri":      autorest.Encode("path", resourceURI),
	}

	const APIVersion = "2017-04-19"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{resourceUri}/providers/Microsoft.Advisor/recommendations/{recommendationId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client RecommendationsClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client RecommendationsClient) GetResponder(resp *http.Response) (result ResourceRecommendationBase, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetGenerateStatus retrieves the status of the recommendation computation or generation process. Invoke this API
// after calling the generation recommendation. The URI of this API is returned in the Location field of the response
// header.
// Parameters:
// operationID - the operation ID, which can be found from the Location field in the generate recommendation
// response header.
func (client RecommendationsClient) GetGenerateStatus(ctx context.Context, operationID uuid.UUID) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RecommendationsClient.GetGenerateStatus")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetGenerateStatusPreparer(ctx, operationID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "advisor.RecommendationsClient", "GetGenerateStatus", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetGenerateStatusSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "advisor.RecommendationsClient", "GetGenerateStatus", resp, "Failure sending request")
		return
	}

	result, err = client.GetGenerateStatusResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "advisor.RecommendationsClient", "GetGenerateStatus", resp, "Failure responding to request")
	}

	return
}

// GetGenerateStatusPreparer prepares the GetGenerateStatus request.
func (client RecommendationsClient) GetGenerateStatusPreparer(ctx context.Context, operationID uuid.UUID) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"operationId":    autorest.Encode("path", operationID),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-19"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Advisor/generateRecommendations/{operationId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetGenerateStatusSender sends the GetGenerateStatus request. The method will close the
// http.Response Body if it receives an error.
func (client RecommendationsClient) GetGenerateStatusSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetGenerateStatusResponder handles the response to the GetGenerateStatus request. The method always
// closes the http.Response Body.
func (client RecommendationsClient) GetGenerateStatusResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// List obtains cached recommendations for a subscription. The recommendations are generated or computed by invoking
// generateRecommendations.
// Parameters:
// filter - the filter to apply to the recommendations.
// top - the number of recommendations per page if a paged version of this API is being used.
// skipToken - the page-continuation token to use with a paged version of this API.
func (client RecommendationsClient) List(ctx context.Context, filter string, top *int32, skipToken string) (result ResourceRecommendationBaseListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RecommendationsClient.List")
		defer func() {
			sc := -1
			if result.rrblr.Response.Response != nil {
				sc = result.rrblr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, filter, top, skipToken)
	if err != nil {
		err = autorest.NewErrorWithError(err, "advisor.RecommendationsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.rrblr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "advisor.RecommendationsClient", "List", resp, "Failure sending request")
		return
	}

	result.rrblr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "advisor.RecommendationsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client RecommendationsClient) ListPreparer(ctx context.Context, filter string, top *int32, skipToken string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-19"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
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
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Advisor/recommendations", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client RecommendationsClient) ListSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client RecommendationsClient) ListResponder(resp *http.Response) (result ResourceRecommendationBaseListResult, err error) {
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
func (client RecommendationsClient) listNextResults(ctx context.Context, lastResults ResourceRecommendationBaseListResult) (result ResourceRecommendationBaseListResult, err error) {
	req, err := lastResults.resourceRecommendationBaseListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "advisor.RecommendationsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "advisor.RecommendationsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "advisor.RecommendationsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client RecommendationsClient) ListComplete(ctx context.Context, filter string, top *int32, skipToken string) (result ResourceRecommendationBaseListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RecommendationsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, filter, top, skipToken)
	return
}
