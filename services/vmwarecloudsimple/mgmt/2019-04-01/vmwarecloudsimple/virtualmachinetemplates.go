package vmwarecloudsimple

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
    "github.com/Azure/go-autorest/autorest"
    "github.com/Azure/go-autorest/autorest/azure"
    "net/http"
    "context"
    "github.com/Azure/go-autorest/tracing"
)

// VirtualMachineTemplatesClient is the description of the new service
type VirtualMachineTemplatesClient struct {
    BaseClient
}
// NewVirtualMachineTemplatesClient creates an instance of the VirtualMachineTemplatesClient client.
func NewVirtualMachineTemplatesClient(subscriptionID string, referer string) VirtualMachineTemplatesClient {
    return NewVirtualMachineTemplatesClientWithBaseURI(DefaultBaseURI, subscriptionID, referer)
}

// NewVirtualMachineTemplatesClientWithBaseURI creates an instance of the VirtualMachineTemplatesClient client.
    func NewVirtualMachineTemplatesClientWithBaseURI(baseURI string, subscriptionID string, referer string) VirtualMachineTemplatesClient {
        return VirtualMachineTemplatesClient{ NewWithBaseURI(baseURI, subscriptionID, referer)}
    }

// Get returns virtual machine templates by its name
    // Parameters:
        // regionID - the region Id (westus, eastus)
        // pcName - the private cloud name
        // virtualMachineTemplateName - virtual machine template id (vsphereId)
func (client VirtualMachineTemplatesClient) Get(ctx context.Context, regionID string, pcName string, virtualMachineTemplateName string) (result VirtualMachineTemplate, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/VirtualMachineTemplatesClient.Get")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.GetPreparer(ctx, regionID, pcName, virtualMachineTemplateName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "vmwarecloudsimple.VirtualMachineTemplatesClient", "Get", nil , "Failure preparing request")
    return
    }

            resp, err := client.GetSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "vmwarecloudsimple.VirtualMachineTemplatesClient", "Get", resp, "Failure sending request")
            return
            }

            result, err = client.GetResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "vmwarecloudsimple.VirtualMachineTemplatesClient", "Get", resp, "Failure responding to request")
            }

    return
    }

    // GetPreparer prepares the Get request.
    func (client VirtualMachineTemplatesClient) GetPreparer(ctx context.Context, regionID string, pcName string, virtualMachineTemplateName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "pcName": autorest.Encode("path",pcName),
            "regionId": autorest.Encode("path",regionID),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            "virtualMachineTemplateName": autorest.Encode("path",virtualMachineTemplateName),
            }

                        const APIVersion = "2019-04-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/locations/{regionId}/privateClouds/{pcName}/virtualMachineTemplates/{virtualMachineTemplateName}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client VirtualMachineTemplatesClient) GetSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client VirtualMachineTemplatesClient) GetResponder(resp *http.Response) (result VirtualMachineTemplate, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// List returns list of virtual machine templates in region for private cloud
    // Parameters:
        // pcName - the private cloud name
        // regionID - the region Id (westus, eastus)
        // resourcePoolName - resource pool used to derive vSphere cluster which contains VM templates
func (client VirtualMachineTemplatesClient) List(ctx context.Context, pcName string, regionID string, resourcePoolName string) (result VirtualMachineTemplateListResponsePage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/VirtualMachineTemplatesClient.List")
        defer func() {
            sc := -1
            if result.vmtlr.Response.Response != nil {
                sc = result.vmtlr.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
                result.fn = client.listNextResults
    req, err := client.ListPreparer(ctx, pcName, regionID, resourcePoolName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "vmwarecloudsimple.VirtualMachineTemplatesClient", "List", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListSender(req)
            if err != nil {
            result.vmtlr.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "vmwarecloudsimple.VirtualMachineTemplatesClient", "List", resp, "Failure sending request")
            return
            }

            result.vmtlr, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "vmwarecloudsimple.VirtualMachineTemplatesClient", "List", resp, "Failure responding to request")
            }

    return
    }

    // ListPreparer prepares the List request.
    func (client VirtualMachineTemplatesClient) ListPreparer(ctx context.Context, pcName string, regionID string, resourcePoolName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "pcName": autorest.Encode("path",pcName),
            "regionId": autorest.Encode("path",regionID),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2019-04-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        "resourcePoolName": autorest.Encode("query",resourcePoolName),
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/locations/{regionId}/privateClouds/{pcName}/virtualMachineTemplates",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListSender sends the List request. The method will close the
    // http.Response Body if it receives an error.
    func (client VirtualMachineTemplatesClient) ListSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client VirtualMachineTemplatesClient) ListResponder(resp *http.Response) (result VirtualMachineTemplateListResponse, err error) {
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
            func (client VirtualMachineTemplatesClient) listNextResults(ctx context.Context, lastResults VirtualMachineTemplateListResponse) (result VirtualMachineTemplateListResponse, err error) {
            req, err := lastResults.virtualMachineTemplateListResponsePreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "vmwarecloudsimple.VirtualMachineTemplatesClient", "listNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "vmwarecloudsimple.VirtualMachineTemplatesClient", "listNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "vmwarecloudsimple.VirtualMachineTemplatesClient", "listNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

    // ListComplete enumerates all values, automatically crossing page boundaries as required.
    func (client VirtualMachineTemplatesClient) ListComplete(ctx context.Context, pcName string, regionID string, resourcePoolName string) (result VirtualMachineTemplateListResponseIterator, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/VirtualMachineTemplatesClient.List")
            defer func() {
                sc := -1
                if result.Response().Response.Response != nil {
                    sc = result.page.Response().Response.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
     }
        result.page, err = client.List(ctx, pcName, regionID, resourcePoolName)
                return
        }

