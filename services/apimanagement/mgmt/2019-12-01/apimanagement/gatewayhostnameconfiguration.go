package apimanagement

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

// GatewayHostnameConfigurationClient is the apiManagement Client
type GatewayHostnameConfigurationClient struct {
	BaseClient
}

// NewGatewayHostnameConfigurationClient creates an instance of the GatewayHostnameConfigurationClient client.
func NewGatewayHostnameConfigurationClient(subscriptionID string) GatewayHostnameConfigurationClient {
	return NewGatewayHostnameConfigurationClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewGatewayHostnameConfigurationClientWithBaseURI creates an instance of the GatewayHostnameConfigurationClient
// client using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI
// (sovereign clouds, Azure stack).
func NewGatewayHostnameConfigurationClientWithBaseURI(baseURI string, subscriptionID string) GatewayHostnameConfigurationClient {
	return GatewayHostnameConfigurationClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates of updates hostname configuration for a Gateway.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// gatewayID - gateway entity identifier. Must be unique in the current API Management service instance. Must
// not have value 'managed'
// hcID - gateway hostname configuration identifier. Must be unique in the scope of parent Gateway entity.
func (client GatewayHostnameConfigurationClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, gatewayID string, hcID string, parameters GatewayHostnameConfigurationContract) (result GatewayHostnameConfigurationContract, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/GatewayHostnameConfigurationClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: gatewayID,
			Constraints: []validation.Constraint{{Target: "gatewayID", Name: validation.MaxLength, Rule: 80, Chain: nil},
				{Target: "gatewayID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: hcID,
			Constraints: []validation.Constraint{{Target: "hcID", Name: validation.MaxLength, Rule: 80, Chain: nil},
				{Target: "hcID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.GatewayHostnameConfigurationClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, serviceName, gatewayID, hcID, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.GatewayHostnameConfigurationClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.GatewayHostnameConfigurationClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.GatewayHostnameConfigurationClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client GatewayHostnameConfigurationClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, serviceName string, gatewayID string, hcID string, parameters GatewayHostnameConfigurationContract) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"gatewayId":         autorest.Encode("path", gatewayID),
		"hcId":              autorest.Encode("path", hcID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/gateways/{gatewayId}/hostnameConfigurations/{hcId}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client GatewayHostnameConfigurationClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client GatewayHostnameConfigurationClient) CreateOrUpdateResponder(resp *http.Response) (result GatewayHostnameConfigurationContract, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the specified hostname configuration from the specified Gateway.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// gatewayID - gateway entity identifier. Must be unique in the current API Management service instance. Must
// not have value 'managed'
// hcID - gateway hostname configuration identifier. Must be unique in the scope of parent Gateway entity.
func (client GatewayHostnameConfigurationClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, gatewayID string, hcID string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/GatewayHostnameConfigurationClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: gatewayID,
			Constraints: []validation.Constraint{{Target: "gatewayID", Name: validation.MaxLength, Rule: 80, Chain: nil},
				{Target: "gatewayID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: hcID,
			Constraints: []validation.Constraint{{Target: "hcID", Name: validation.MaxLength, Rule: 80, Chain: nil},
				{Target: "hcID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.GatewayHostnameConfigurationClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, serviceName, gatewayID, hcID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.GatewayHostnameConfigurationClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "apimanagement.GatewayHostnameConfigurationClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.GatewayHostnameConfigurationClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client GatewayHostnameConfigurationClient) DeletePreparer(ctx context.Context, resourceGroupName string, serviceName string, gatewayID string, hcID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"gatewayId":         autorest.Encode("path", gatewayID),
		"hcId":              autorest.Encode("path", hcID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/gateways/{gatewayId}/hostnameConfigurations/{hcId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client GatewayHostnameConfigurationClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client GatewayHostnameConfigurationClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the details of the Gateway hostname configuration specified by its identifier.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// gatewayID - gateway entity identifier. Must be unique in the current API Management service instance. Must
// not have value 'managed'
// hcID - gateway hostname configuration identifier. Must be unique in the scope of parent Gateway entity.
func (client GatewayHostnameConfigurationClient) Get(ctx context.Context, resourceGroupName string, serviceName string, gatewayID string, hcID string) (result GatewayHostnameConfigurationContract, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/GatewayHostnameConfigurationClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: gatewayID,
			Constraints: []validation.Constraint{{Target: "gatewayID", Name: validation.MaxLength, Rule: 80, Chain: nil},
				{Target: "gatewayID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: hcID,
			Constraints: []validation.Constraint{{Target: "hcID", Name: validation.MaxLength, Rule: 80, Chain: nil},
				{Target: "hcID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.GatewayHostnameConfigurationClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, serviceName, gatewayID, hcID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.GatewayHostnameConfigurationClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.GatewayHostnameConfigurationClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.GatewayHostnameConfigurationClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client GatewayHostnameConfigurationClient) GetPreparer(ctx context.Context, resourceGroupName string, serviceName string, gatewayID string, hcID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"gatewayId":         autorest.Encode("path", gatewayID),
		"hcId":              autorest.Encode("path", hcID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/gateways/{gatewayId}/hostnameConfigurations/{hcId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client GatewayHostnameConfigurationClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client GatewayHostnameConfigurationClient) GetResponder(resp *http.Response) (result GatewayHostnameConfigurationContract, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetEntityTag checks that hostname configuration entity specified by identifier exists for specified Gateway entity.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// gatewayID - gateway entity identifier. Must be unique in the current API Management service instance. Must
// not have value 'managed'
// hcID - gateway hostname configuration identifier. Must be unique in the scope of parent Gateway entity.
func (client GatewayHostnameConfigurationClient) GetEntityTag(ctx context.Context, resourceGroupName string, serviceName string, gatewayID string, hcID string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/GatewayHostnameConfigurationClient.GetEntityTag")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: gatewayID,
			Constraints: []validation.Constraint{{Target: "gatewayID", Name: validation.MaxLength, Rule: 80, Chain: nil},
				{Target: "gatewayID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: hcID,
			Constraints: []validation.Constraint{{Target: "hcID", Name: validation.MaxLength, Rule: 80, Chain: nil},
				{Target: "hcID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.GatewayHostnameConfigurationClient", "GetEntityTag", err.Error())
	}

	req, err := client.GetEntityTagPreparer(ctx, resourceGroupName, serviceName, gatewayID, hcID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.GatewayHostnameConfigurationClient", "GetEntityTag", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetEntityTagSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "apimanagement.GatewayHostnameConfigurationClient", "GetEntityTag", resp, "Failure sending request")
		return
	}

	result, err = client.GetEntityTagResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.GatewayHostnameConfigurationClient", "GetEntityTag", resp, "Failure responding to request")
	}

	return
}

// GetEntityTagPreparer prepares the GetEntityTag request.
func (client GatewayHostnameConfigurationClient) GetEntityTagPreparer(ctx context.Context, resourceGroupName string, serviceName string, gatewayID string, hcID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"gatewayId":         autorest.Encode("path", gatewayID),
		"hcId":              autorest.Encode("path", hcID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsHead(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/gateways/{gatewayId}/hostnameConfigurations/{hcId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetEntityTagSender sends the GetEntityTag request. The method will close the
// http.Response Body if it receives an error.
func (client GatewayHostnameConfigurationClient) GetEntityTagSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetEntityTagResponder handles the response to the GetEntityTag request. The method always
// closes the http.Response Body.
func (client GatewayHostnameConfigurationClient) GetEntityTagResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// ListByService lists the collection of hostname configurations for the specified gateway.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// gatewayID - gateway entity identifier. Must be unique in the current API Management service instance. Must
// not have value 'managed'
// top - number of records to return.
// skip - number of records to skip.
func (client GatewayHostnameConfigurationClient) ListByService(ctx context.Context, resourceGroupName string, serviceName string, gatewayID string, top *int32, skip *int32) (result GatewayHostnameConfigurationCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/GatewayHostnameConfigurationClient.ListByService")
		defer func() {
			sc := -1
			if result.ghcc.Response.Response != nil {
				sc = result.ghcc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: gatewayID,
			Constraints: []validation.Constraint{{Target: "gatewayID", Name: validation.MaxLength, Rule: 80, Chain: nil},
				{Target: "gatewayID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMinimum, Rule: int64(1), Chain: nil}}}}},
		{TargetValue: skip,
			Constraints: []validation.Constraint{{Target: "skip", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "skip", Name: validation.InclusiveMinimum, Rule: int64(0), Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("apimanagement.GatewayHostnameConfigurationClient", "ListByService", err.Error())
	}

	result.fn = client.listByServiceNextResults
	req, err := client.ListByServicePreparer(ctx, resourceGroupName, serviceName, gatewayID, top, skip)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.GatewayHostnameConfigurationClient", "ListByService", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByServiceSender(req)
	if err != nil {
		result.ghcc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.GatewayHostnameConfigurationClient", "ListByService", resp, "Failure sending request")
		return
	}

	result.ghcc, err = client.ListByServiceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.GatewayHostnameConfigurationClient", "ListByService", resp, "Failure responding to request")
	}

	return
}

// ListByServicePreparer prepares the ListByService request.
func (client GatewayHostnameConfigurationClient) ListByServicePreparer(ctx context.Context, resourceGroupName string, serviceName string, gatewayID string, top *int32, skip *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"gatewayId":         autorest.Encode("path", gatewayID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if skip != nil {
		queryParameters["$skip"] = autorest.Encode("query", *skip)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/gateways/{gatewayId}/hostnameConfigurations", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByServiceSender sends the ListByService request. The method will close the
// http.Response Body if it receives an error.
func (client GatewayHostnameConfigurationClient) ListByServiceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByServiceResponder handles the response to the ListByService request. The method always
// closes the http.Response Body.
func (client GatewayHostnameConfigurationClient) ListByServiceResponder(resp *http.Response) (result GatewayHostnameConfigurationCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByServiceNextResults retrieves the next set of results, if any.
func (client GatewayHostnameConfigurationClient) listByServiceNextResults(ctx context.Context, lastResults GatewayHostnameConfigurationCollection) (result GatewayHostnameConfigurationCollection, err error) {
	req, err := lastResults.gatewayHostnameConfigurationCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "apimanagement.GatewayHostnameConfigurationClient", "listByServiceNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByServiceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "apimanagement.GatewayHostnameConfigurationClient", "listByServiceNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByServiceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.GatewayHostnameConfigurationClient", "listByServiceNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByServiceComplete enumerates all values, automatically crossing page boundaries as required.
func (client GatewayHostnameConfigurationClient) ListByServiceComplete(ctx context.Context, resourceGroupName string, serviceName string, gatewayID string, top *int32, skip *int32) (result GatewayHostnameConfigurationCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/GatewayHostnameConfigurationClient.ListByService")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByService(ctx, resourceGroupName, serviceName, gatewayID, top, skip)
	return
}
