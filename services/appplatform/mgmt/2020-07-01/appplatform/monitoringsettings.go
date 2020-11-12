package appplatform

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

// MonitoringSettingsClient is the REST API for Azure Spring Cloud
type MonitoringSettingsClient struct {
	BaseClient
}

// NewMonitoringSettingsClient creates an instance of the MonitoringSettingsClient client.
func NewMonitoringSettingsClient(subscriptionID string) MonitoringSettingsClient {
	return NewMonitoringSettingsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewMonitoringSettingsClientWithBaseURI creates an instance of the MonitoringSettingsClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewMonitoringSettingsClientWithBaseURI(baseURI string, subscriptionID string) MonitoringSettingsClient {
	return MonitoringSettingsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get get the Monitoring Setting and its properties.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serviceName - the name of the Service resource.
func (client MonitoringSettingsClient) Get(ctx context.Context, resourceGroupName string, serviceName string) (result MonitoringSettingResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MonitoringSettingsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, serviceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.MonitoringSettingsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "appplatform.MonitoringSettingsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.MonitoringSettingsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client MonitoringSettingsClient) GetPreparer(ctx context.Context, resourceGroupName string, serviceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/monitoringSettings/default", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client MonitoringSettingsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client MonitoringSettingsClient) GetResponder(resp *http.Response) (result MonitoringSettingResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// UpdatePatch update the Monitoring Setting.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serviceName - the name of the Service resource.
// monitoringSettingResource - parameters for the update operation
func (client MonitoringSettingsClient) UpdatePatch(ctx context.Context, resourceGroupName string, serviceName string, monitoringSettingResource MonitoringSettingResource) (result MonitoringSettingsUpdatePatchFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MonitoringSettingsClient.UpdatePatch")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePatchPreparer(ctx, resourceGroupName, serviceName, monitoringSettingResource)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.MonitoringSettingsClient", "UpdatePatch", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdatePatchSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.MonitoringSettingsClient", "UpdatePatch", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdatePatchPreparer prepares the UpdatePatch request.
func (client MonitoringSettingsClient) UpdatePatchPreparer(ctx context.Context, resourceGroupName string, serviceName string, monitoringSettingResource MonitoringSettingResource) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/monitoringSettings/default", pathParameters),
		autorest.WithJSON(monitoringSettingResource),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdatePatchSender sends the UpdatePatch request. The method will close the
// http.Response Body if it receives an error.
func (client MonitoringSettingsClient) UpdatePatchSender(req *http.Request) (future MonitoringSettingsUpdatePatchFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// UpdatePatchResponder handles the response to the UpdatePatch request. The method always
// closes the http.Response Body.
func (client MonitoringSettingsClient) UpdatePatchResponder(resp *http.Response) (result MonitoringSettingResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// UpdatePut update the Monitoring Setting.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serviceName - the name of the Service resource.
// monitoringSettingResource - parameters for the update operation
func (client MonitoringSettingsClient) UpdatePut(ctx context.Context, resourceGroupName string, serviceName string, monitoringSettingResource MonitoringSettingResource) (result MonitoringSettingsUpdatePutFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MonitoringSettingsClient.UpdatePut")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: monitoringSettingResource,
			Constraints: []validation.Constraint{{Target: "monitoringSettingResource.Properties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "monitoringSettingResource.Properties.AppInsightsSamplingRate", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "monitoringSettingResource.Properties.AppInsightsSamplingRate", Name: validation.InclusiveMaximum, Rule: int64(100), Chain: nil},
						{Target: "monitoringSettingResource.Properties.AppInsightsSamplingRate", Name: validation.InclusiveMinimum, Rule: int64(0), Chain: nil},
					}},
				}}}}}); err != nil {
		return result, validation.NewError("appplatform.MonitoringSettingsClient", "UpdatePut", err.Error())
	}

	req, err := client.UpdatePutPreparer(ctx, resourceGroupName, serviceName, monitoringSettingResource)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.MonitoringSettingsClient", "UpdatePut", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdatePutSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.MonitoringSettingsClient", "UpdatePut", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdatePutPreparer prepares the UpdatePut request.
func (client MonitoringSettingsClient) UpdatePutPreparer(ctx context.Context, resourceGroupName string, serviceName string, monitoringSettingResource MonitoringSettingResource) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/monitoringSettings/default", pathParameters),
		autorest.WithJSON(monitoringSettingResource),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdatePutSender sends the UpdatePut request. The method will close the
// http.Response Body if it receives an error.
func (client MonitoringSettingsClient) UpdatePutSender(req *http.Request) (future MonitoringSettingsUpdatePutFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// UpdatePutResponder handles the response to the UpdatePut request. The method always
// closes the http.Response Body.
func (client MonitoringSettingsClient) UpdatePutResponder(resp *http.Response) (result MonitoringSettingResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
