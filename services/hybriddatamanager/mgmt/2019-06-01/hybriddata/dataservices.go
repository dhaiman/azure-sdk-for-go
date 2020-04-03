package hybriddata

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

// DataServicesClient is the client for the DataServices methods of the Hybriddata service.
type DataServicesClient struct {
	BaseClient
}

// NewDataServicesClient creates an instance of the DataServicesClient client.
func NewDataServicesClient(subscriptionID string) DataServicesClient {
	return NewDataServicesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewDataServicesClientWithBaseURI creates an instance of the DataServicesClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewDataServicesClientWithBaseURI(baseURI string, subscriptionID string) DataServicesClient {
	return DataServicesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get gets the data service that matches the data service name given.
// Parameters:
// dataServiceName - the name of the data service that is being queried.
// resourceGroupName - the Resource Group Name
// dataManagerName - the name of the DataManager Resource within the specified resource group. DataManager
// names must be between 3 and 24 characters in length and use any alphanumeric and underscore only
func (client DataServicesClient) Get(ctx context.Context, dataServiceName string, resourceGroupName string, dataManagerName string) (result DataService, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DataServicesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: dataManagerName,
			Constraints: []validation.Constraint{{Target: "dataManagerName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "dataManagerName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "dataManagerName", Name: validation.Pattern, Rule: `^[-\w\.]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("hybriddata.DataServicesClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, dataServiceName, resourceGroupName, dataManagerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybriddata.DataServicesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "hybriddata.DataServicesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybriddata.DataServicesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client DataServicesClient) GetPreparer(ctx context.Context, dataServiceName string, resourceGroupName string, dataManagerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"dataManagerName":   autorest.Encode("path", dataManagerName),
		"dataServiceName":   autorest.Encode("path", dataServiceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataServices/{dataServiceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client DataServicesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client DataServicesClient) GetResponder(resp *http.Response) (result DataService, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByDataManager this method gets all the data services.
// Parameters:
// resourceGroupName - the Resource Group Name
// dataManagerName - the name of the DataManager Resource within the specified resource group. DataManager
// names must be between 3 and 24 characters in length and use any alphanumeric and underscore only
func (client DataServicesClient) ListByDataManager(ctx context.Context, resourceGroupName string, dataManagerName string) (result DataServiceListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DataServicesClient.ListByDataManager")
		defer func() {
			sc := -1
			if result.dsl.Response.Response != nil {
				sc = result.dsl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: dataManagerName,
			Constraints: []validation.Constraint{{Target: "dataManagerName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "dataManagerName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "dataManagerName", Name: validation.Pattern, Rule: `^[-\w\.]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("hybriddata.DataServicesClient", "ListByDataManager", err.Error())
	}

	result.fn = client.listByDataManagerNextResults
	req, err := client.ListByDataManagerPreparer(ctx, resourceGroupName, dataManagerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybriddata.DataServicesClient", "ListByDataManager", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByDataManagerSender(req)
	if err != nil {
		result.dsl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "hybriddata.DataServicesClient", "ListByDataManager", resp, "Failure sending request")
		return
	}

	result.dsl, err = client.ListByDataManagerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybriddata.DataServicesClient", "ListByDataManager", resp, "Failure responding to request")
	}

	return
}

// ListByDataManagerPreparer prepares the ListByDataManager request.
func (client DataServicesClient) ListByDataManagerPreparer(ctx context.Context, resourceGroupName string, dataManagerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"dataManagerName":   autorest.Encode("path", dataManagerName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataServices", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByDataManagerSender sends the ListByDataManager request. The method will close the
// http.Response Body if it receives an error.
func (client DataServicesClient) ListByDataManagerSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByDataManagerResponder handles the response to the ListByDataManager request. The method always
// closes the http.Response Body.
func (client DataServicesClient) ListByDataManagerResponder(resp *http.Response) (result DataServiceList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByDataManagerNextResults retrieves the next set of results, if any.
func (client DataServicesClient) listByDataManagerNextResults(ctx context.Context, lastResults DataServiceList) (result DataServiceList, err error) {
	req, err := lastResults.dataServiceListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "hybriddata.DataServicesClient", "listByDataManagerNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByDataManagerSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "hybriddata.DataServicesClient", "listByDataManagerNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByDataManagerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybriddata.DataServicesClient", "listByDataManagerNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByDataManagerComplete enumerates all values, automatically crossing page boundaries as required.
func (client DataServicesClient) ListByDataManagerComplete(ctx context.Context, resourceGroupName string, dataManagerName string) (result DataServiceListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DataServicesClient.ListByDataManager")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByDataManager(ctx, resourceGroupName, dataManagerName)
	return
}
