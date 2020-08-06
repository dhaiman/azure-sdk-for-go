package personalizer

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
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// LearningMetricsClient is the personalizer Service is an Azure Cognitive Service that makes it easy to target content
// and experiences without complex pre-analysis or cleanup of past data. Given a context and featurized content, the
// Personalizer Service returns which content item to show to users in rewardActionId. As rewards are sent in response
// to the use of rewardActionId, the reinforcement learning algorithm will improve the model and improve performance of
// future rank calls.
type LearningMetricsClient struct {
	BaseClient
}

// NewLearningMetricsClient creates an instance of the LearningMetricsClient client.
func NewLearningMetricsClient(endpoint string) LearningMetricsClient {
	return LearningMetricsClient{New(endpoint)}
}

// Get get learning performance metrics aggregated by time or events.
// Parameters:
// startDate - start date to get the metrics for.
// endDate - end date to get the metrics for.
// aggregationInterval - aggregation time period to aggregate the results for.
// numberOfRecentEvents - number of recent events.
// aggregationEventCount - aggregation event window to aggregate the results for.
func (client LearningMetricsClient) Get(ctx context.Context, startDate *date.Time, endDate *date.Time, aggregationInterval string, numberOfRecentEvents *int64, aggregationEventCount *int64) (result ListListMetric, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LearningMetricsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, startDate, endDate, aggregationInterval, numberOfRecentEvents, aggregationEventCount)
	if err != nil {
		err = autorest.NewErrorWithError(err, "personalizer.LearningMetricsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "personalizer.LearningMetricsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "personalizer.LearningMetricsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client LearningMetricsClient) GetPreparer(ctx context.Context, startDate *date.Time, endDate *date.Time, aggregationInterval string, numberOfRecentEvents *int64, aggregationEventCount *int64) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	queryParameters := map[string]interface{}{}
	if startDate != nil {
		queryParameters["StartDate"] = autorest.Encode("query", *startDate)
	}
	if endDate != nil {
		queryParameters["EndDate"] = autorest.Encode("query", *endDate)
	}
	if len(aggregationInterval) > 0 {
		queryParameters["AggregationInterval"] = autorest.Encode("query", aggregationInterval)
	}
	if numberOfRecentEvents != nil {
		queryParameters["NumberOfRecentEvents"] = autorest.Encode("query", *numberOfRecentEvents)
	}
	if aggregationEventCount != nil {
		queryParameters["AggregationEventCount"] = autorest.Encode("query", *aggregationEventCount)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{Endpoint}/personalizer/v1.0", urlParameters),
		autorest.WithPath("/metrics/learning"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client LearningMetricsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client LearningMetricsClient) GetResponder(resp *http.Response) (result ListListMetric, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
