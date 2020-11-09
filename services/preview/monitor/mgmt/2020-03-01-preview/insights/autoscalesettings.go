package insights

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

// AutoscaleSettingsClient is the monitor Management Client
type AutoscaleSettingsClient struct {
	BaseClient
}

// NewAutoscaleSettingsClient creates an instance of the AutoscaleSettingsClient client.
func NewAutoscaleSettingsClient(subscriptionID string) AutoscaleSettingsClient {
	return NewAutoscaleSettingsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewAutoscaleSettingsClientWithBaseURI creates an instance of the AutoscaleSettingsClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewAutoscaleSettingsClientWithBaseURI(baseURI string, subscriptionID string) AutoscaleSettingsClient {
	return AutoscaleSettingsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates an autoscale setting.
// Parameters:
// resourceGroupName - the name of the resource group.
// autoscaleSettingName - the autoscale setting name.
// parameters - parameters supplied to the operation.
func (client AutoscaleSettingsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, autoscaleSettingName string, parameters AutoscaleSettingResource) (result AutoscaleSettingResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AutoscaleSettingsClient.CreateOrUpdate")
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
			Constraints: []validation.Constraint{{Target: "parameters.AutoscaleSetting", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "parameters.AutoscaleSetting.Profiles", Name: validation.Null, Rule: true,
					Chain: []validation.Constraint{{Target: "parameters.AutoscaleSetting.Profiles", Name: validation.MaxItems, Rule: 20, Chain: nil}}},
				}}}}}); err != nil {
		return result, validation.NewError("insights.AutoscaleSettingsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, autoscaleSettingName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.AutoscaleSettingsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "insights.AutoscaleSettingsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.AutoscaleSettingsClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client AutoscaleSettingsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, autoscaleSettingName string, parameters AutoscaleSettingResource) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"autoscaleSettingName": autorest.Encode("path", autoscaleSettingName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/microsoft.insights/autoscalesettings/{autoscaleSettingName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client AutoscaleSettingsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client AutoscaleSettingsClient) CreateOrUpdateResponder(resp *http.Response) (result AutoscaleSettingResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes and autoscale setting
// Parameters:
// resourceGroupName - the name of the resource group.
// autoscaleSettingName - the autoscale setting name.
func (client AutoscaleSettingsClient) Delete(ctx context.Context, resourceGroupName string, autoscaleSettingName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AutoscaleSettingsClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, autoscaleSettingName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.AutoscaleSettingsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "insights.AutoscaleSettingsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.AutoscaleSettingsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client AutoscaleSettingsClient) DeletePreparer(ctx context.Context, resourceGroupName string, autoscaleSettingName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"autoscaleSettingName": autorest.Encode("path", autoscaleSettingName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/microsoft.insights/autoscalesettings/{autoscaleSettingName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client AutoscaleSettingsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client AutoscaleSettingsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets an autoscale setting
// Parameters:
// resourceGroupName - the name of the resource group.
// autoscaleSettingName - the autoscale setting name.
func (client AutoscaleSettingsClient) Get(ctx context.Context, resourceGroupName string, autoscaleSettingName string) (result AutoscaleSettingResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AutoscaleSettingsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, autoscaleSettingName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.AutoscaleSettingsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "insights.AutoscaleSettingsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.AutoscaleSettingsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client AutoscaleSettingsClient) GetPreparer(ctx context.Context, resourceGroupName string, autoscaleSettingName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"autoscaleSettingName": autorest.Encode("path", autoscaleSettingName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/microsoft.insights/autoscalesettings/{autoscaleSettingName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client AutoscaleSettingsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client AutoscaleSettingsClient) GetResponder(resp *http.Response) (result AutoscaleSettingResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroup lists the autoscale settings for a resource group
// Parameters:
// resourceGroupName - the name of the resource group.
func (client AutoscaleSettingsClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (result AutoscaleSettingResourceCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AutoscaleSettingsClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.asrc.Response.Response != nil {
				sc = result.asrc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.AutoscaleSettingsClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.asrc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "insights.AutoscaleSettingsClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.asrc, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.AutoscaleSettingsClient", "ListByResourceGroup", resp, "Failure responding to request")
	}
	if result.asrc.hasNextLink() && result.asrc.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client AutoscaleSettingsClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/microsoft.insights/autoscalesettings", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client AutoscaleSettingsClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client AutoscaleSettingsClient) ListByResourceGroupResponder(resp *http.Response) (result AutoscaleSettingResourceCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client AutoscaleSettingsClient) listByResourceGroupNextResults(ctx context.Context, lastResults AutoscaleSettingResourceCollection) (result AutoscaleSettingResourceCollection, err error) {
	req, err := lastResults.autoscaleSettingResourceCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "insights.AutoscaleSettingsClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "insights.AutoscaleSettingsClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.AutoscaleSettingsClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client AutoscaleSettingsClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result AutoscaleSettingResourceCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AutoscaleSettingsClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName)
	return
}

// ListBySubscription lists the autoscale settings for a subscription
func (client AutoscaleSettingsClient) ListBySubscription(ctx context.Context) (result AutoscaleSettingResourceCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AutoscaleSettingsClient.ListBySubscription")
		defer func() {
			sc := -1
			if result.asrc.Response.Response != nil {
				sc = result.asrc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listBySubscriptionNextResults
	req, err := client.ListBySubscriptionPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.AutoscaleSettingsClient", "ListBySubscription", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.asrc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "insights.AutoscaleSettingsClient", "ListBySubscription", resp, "Failure sending request")
		return
	}

	result.asrc, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.AutoscaleSettingsClient", "ListBySubscription", resp, "Failure responding to request")
	}
	if result.asrc.hasNextLink() && result.asrc.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListBySubscriptionPreparer prepares the ListBySubscription request.
func (client AutoscaleSettingsClient) ListBySubscriptionPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/microsoft.insights/autoscalesettings", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListBySubscriptionSender sends the ListBySubscription request. The method will close the
// http.Response Body if it receives an error.
func (client AutoscaleSettingsClient) ListBySubscriptionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListBySubscriptionResponder handles the response to the ListBySubscription request. The method always
// closes the http.Response Body.
func (client AutoscaleSettingsClient) ListBySubscriptionResponder(resp *http.Response) (result AutoscaleSettingResourceCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listBySubscriptionNextResults retrieves the next set of results, if any.
func (client AutoscaleSettingsClient) listBySubscriptionNextResults(ctx context.Context, lastResults AutoscaleSettingResourceCollection) (result AutoscaleSettingResourceCollection, err error) {
	req, err := lastResults.autoscaleSettingResourceCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "insights.AutoscaleSettingsClient", "listBySubscriptionNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "insights.AutoscaleSettingsClient", "listBySubscriptionNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.AutoscaleSettingsClient", "listBySubscriptionNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListBySubscriptionComplete enumerates all values, automatically crossing page boundaries as required.
func (client AutoscaleSettingsClient) ListBySubscriptionComplete(ctx context.Context) (result AutoscaleSettingResourceCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AutoscaleSettingsClient.ListBySubscription")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListBySubscription(ctx)
	return
}

// Update updates an existing AutoscaleSettingsResource. To update other fields use the CreateOrUpdate method.
// Parameters:
// resourceGroupName - the name of the resource group.
// autoscaleSettingName - the autoscale setting name.
// autoscaleSettingResource - parameters supplied to the operation.
func (client AutoscaleSettingsClient) Update(ctx context.Context, resourceGroupName string, autoscaleSettingName string, autoscaleSettingResource AutoscaleSettingResourcePatch) (result AutoscaleSettingResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AutoscaleSettingsClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, resourceGroupName, autoscaleSettingName, autoscaleSettingResource)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.AutoscaleSettingsClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "insights.AutoscaleSettingsClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.AutoscaleSettingsClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client AutoscaleSettingsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, autoscaleSettingName string, autoscaleSettingResource AutoscaleSettingResourcePatch) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"autoscaleSettingName": autorest.Encode("path", autoscaleSettingName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/microsoft.insights/autoscalesettings/{autoscaleSettingName}", pathParameters),
		autorest.WithJSON(autoscaleSettingResource),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client AutoscaleSettingsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client AutoscaleSettingsClient) UpdateResponder(resp *http.Response) (result AutoscaleSettingResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
