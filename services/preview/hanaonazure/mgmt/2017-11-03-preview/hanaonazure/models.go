package hanaonazure

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
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/preview/hanaonazure/mgmt/2017-11-03-preview/hanaonazure"

// HanaHardwareTypeNamesEnum enumerates the values for hana hardware type names enum.
type HanaHardwareTypeNamesEnum string

const (
	// CiscoUCS ...
	CiscoUCS HanaHardwareTypeNamesEnum = "Cisco_UCS"
	// HPE ...
	HPE HanaHardwareTypeNamesEnum = "HPE"
)

// PossibleHanaHardwareTypeNamesEnumValues returns an array of possible values for the HanaHardwareTypeNamesEnum const type.
func PossibleHanaHardwareTypeNamesEnumValues() []HanaHardwareTypeNamesEnum {
	return []HanaHardwareTypeNamesEnum{CiscoUCS, HPE}
}

// HanaInstancePowerStateEnum enumerates the values for hana instance power state enum.
type HanaInstancePowerStateEnum string

const (
	// Restarting ...
	Restarting HanaInstancePowerStateEnum = "restarting"
	// Started ...
	Started HanaInstancePowerStateEnum = "started"
	// Starting ...
	Starting HanaInstancePowerStateEnum = "starting"
	// Stopped ...
	Stopped HanaInstancePowerStateEnum = "stopped"
	// Stopping ...
	Stopping HanaInstancePowerStateEnum = "stopping"
	// Unknown ...
	Unknown HanaInstancePowerStateEnum = "unknown"
)

// PossibleHanaInstancePowerStateEnumValues returns an array of possible values for the HanaInstancePowerStateEnum const type.
func PossibleHanaInstancePowerStateEnumValues() []HanaInstancePowerStateEnum {
	return []HanaInstancePowerStateEnum{Restarting, Started, Starting, Stopped, Stopping, Unknown}
}

// HanaInstanceSizeNamesEnum enumerates the values for hana instance size names enum.
type HanaInstanceSizeNamesEnum string

const (
	// S112 ...
	S112 HanaInstanceSizeNamesEnum = "S112"
	// S144 ...
	S144 HanaInstanceSizeNamesEnum = "S144"
	// S144m ...
	S144m HanaInstanceSizeNamesEnum = "S144m"
	// S192 ...
	S192 HanaInstanceSizeNamesEnum = "S192"
	// S192m ...
	S192m HanaInstanceSizeNamesEnum = "S192m"
	// S192xm ...
	S192xm HanaInstanceSizeNamesEnum = "S192xm"
	// S224 ...
	S224 HanaInstanceSizeNamesEnum = "S224"
	// S224m ...
	S224m HanaInstanceSizeNamesEnum = "S224m"
	// S224om ...
	S224om HanaInstanceSizeNamesEnum = "S224om"
	// S224oo ...
	S224oo HanaInstanceSizeNamesEnum = "S224oo"
	// S224oom ...
	S224oom HanaInstanceSizeNamesEnum = "S224oom"
	// S224ooo ...
	S224ooo HanaInstanceSizeNamesEnum = "S224ooo"
	// S384 ...
	S384 HanaInstanceSizeNamesEnum = "S384"
	// S384m ...
	S384m HanaInstanceSizeNamesEnum = "S384m"
	// S384xm ...
	S384xm HanaInstanceSizeNamesEnum = "S384xm"
	// S384xxm ...
	S384xxm HanaInstanceSizeNamesEnum = "S384xxm"
	// S448 ...
	S448 HanaInstanceSizeNamesEnum = "S448"
	// S448m ...
	S448m HanaInstanceSizeNamesEnum = "S448m"
	// S448om ...
	S448om HanaInstanceSizeNamesEnum = "S448om"
	// S448oo ...
	S448oo HanaInstanceSizeNamesEnum = "S448oo"
	// S448oom ...
	S448oom HanaInstanceSizeNamesEnum = "S448oom"
	// S448ooo ...
	S448ooo HanaInstanceSizeNamesEnum = "S448ooo"
	// S576m ...
	S576m HanaInstanceSizeNamesEnum = "S576m"
	// S576xm ...
	S576xm HanaInstanceSizeNamesEnum = "S576xm"
	// S672 ...
	S672 HanaInstanceSizeNamesEnum = "S672"
	// S672m ...
	S672m HanaInstanceSizeNamesEnum = "S672m"
	// S672om ...
	S672om HanaInstanceSizeNamesEnum = "S672om"
	// S672oo ...
	S672oo HanaInstanceSizeNamesEnum = "S672oo"
	// S672oom ...
	S672oom HanaInstanceSizeNamesEnum = "S672oom"
	// S672ooo ...
	S672ooo HanaInstanceSizeNamesEnum = "S672ooo"
	// S72 ...
	S72 HanaInstanceSizeNamesEnum = "S72"
	// S72m ...
	S72m HanaInstanceSizeNamesEnum = "S72m"
	// S768 ...
	S768 HanaInstanceSizeNamesEnum = "S768"
	// S768m ...
	S768m HanaInstanceSizeNamesEnum = "S768m"
	// S768xm ...
	S768xm HanaInstanceSizeNamesEnum = "S768xm"
	// S896 ...
	S896 HanaInstanceSizeNamesEnum = "S896"
	// S896m ...
	S896m HanaInstanceSizeNamesEnum = "S896m"
	// S896om ...
	S896om HanaInstanceSizeNamesEnum = "S896om"
	// S896oo ...
	S896oo HanaInstanceSizeNamesEnum = "S896oo"
	// S896oom ...
	S896oom HanaInstanceSizeNamesEnum = "S896oom"
	// S896ooo ...
	S896ooo HanaInstanceSizeNamesEnum = "S896ooo"
	// S96 ...
	S96 HanaInstanceSizeNamesEnum = "S96"
	// S960m ...
	S960m HanaInstanceSizeNamesEnum = "S960m"
)

// PossibleHanaInstanceSizeNamesEnumValues returns an array of possible values for the HanaInstanceSizeNamesEnum const type.
func PossibleHanaInstanceSizeNamesEnumValues() []HanaInstanceSizeNamesEnum {
	return []HanaInstanceSizeNamesEnum{S112, S144, S144m, S192, S192m, S192xm, S224, S224m, S224om, S224oo, S224oom, S224ooo, S384, S384m, S384xm, S384xxm, S448, S448m, S448om, S448oo, S448oom, S448ooo, S576m, S576xm, S672, S672m, S672om, S672oo, S672oom, S672ooo, S72, S72m, S768, S768m, S768xm, S896, S896m, S896om, S896oo, S896oom, S896ooo, S96, S960m}
}

// HanaProvisioningStatesEnum enumerates the values for hana provisioning states enum.
type HanaProvisioningStatesEnum string

const (
	// Accepted ...
	Accepted HanaProvisioningStatesEnum = "Accepted"
	// Creating ...
	Creating HanaProvisioningStatesEnum = "Creating"
	// Deleting ...
	Deleting HanaProvisioningStatesEnum = "Deleting"
	// Failed ...
	Failed HanaProvisioningStatesEnum = "Failed"
	// Migrating ...
	Migrating HanaProvisioningStatesEnum = "Migrating"
	// Succeeded ...
	Succeeded HanaProvisioningStatesEnum = "Succeeded"
	// Updating ...
	Updating HanaProvisioningStatesEnum = "Updating"
)

// PossibleHanaProvisioningStatesEnumValues returns an array of possible values for the HanaProvisioningStatesEnum const type.
func PossibleHanaProvisioningStatesEnumValues() []HanaProvisioningStatesEnum {
	return []HanaProvisioningStatesEnum{Accepted, Creating, Deleting, Failed, Migrating, Succeeded, Updating}
}

// Disk specifies the disk information fo the HANA instance
type Disk struct {
	// Name - The disk name.
	Name *string `json:"name,omitempty"`
	// DiskSizeGB - Specifies the size of an empty data disk in gigabytes.
	DiskSizeGB *int32 `json:"diskSizeGB,omitempty"`
	// Lun - READ-ONLY; Specifies the logical unit number of the data disk. This value is used to identify data disks within the VM and therefore must be unique for each data disk attached to a VM.
	Lun *int32 `json:"lun,omitempty"`
}

// Display detailed HANA operation information
type Display struct {
	// Provider - READ-ONLY; The localized friendly form of the resource provider name. This form is also expected to include the publisher/company responsible. Use Title Casing. Begin with "Microsoft" for 1st party services.
	Provider *string `json:"provider,omitempty"`
	// Resource - READ-ONLY; The localized friendly form of the resource type related to this action/operation. This form should match the public documentation for the resource provider. Use Title Casing. For examples, refer to the “name” section.
	Resource *string `json:"resource,omitempty"`
	// Operation - READ-ONLY; The localized friendly name for the operation as shown to the user. This name should be concise (to fit in drop downs), but clear (self-documenting). Use Title Casing and include the entity/resource to which it applies.
	Operation *string `json:"operation,omitempty"`
	// Description - READ-ONLY; The localized friendly description for the operation as shown to the user. This description should be thorough, yet concise. It will be used in tool-tips and detailed views.
	Description *string `json:"description,omitempty"`
	// Origin - READ-ONLY; The intended executor of the operation; governs the display of the operation in the RBAC UX and the audit logs UX. Default value is 'user,system'
	Origin *string `json:"origin,omitempty"`
}

// ErrorResponse describes the format of Error response.
type ErrorResponse struct {
	// Code - Error code
	Code *string `json:"code,omitempty"`
	// Message - Error message indicating why the operation failed.
	Message *string `json:"message,omitempty"`
}

// HanaInstance HANA instance info on Azure (ARM properties and HANA properties)
type HanaInstance struct {
	autorest.Response `json:"-"`
	// HanaInstanceProperties - HANA instance properties
	*HanaInstanceProperties `json:"properties,omitempty"`
	// ID - READ-ONLY; Resource ID
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Resource name
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Resource type
	Type *string `json:"type,omitempty"`
	// Location - Resource location
	Location *string `json:"location,omitempty"`
	// Tags - READ-ONLY; Resource tags
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for HanaInstance.
func (hi HanaInstance) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if hi.HanaInstanceProperties != nil {
		objectMap["properties"] = hi.HanaInstanceProperties
	}
	if hi.Location != nil {
		objectMap["location"] = hi.Location
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for HanaInstance struct.
func (hi *HanaInstance) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var hanaInstanceProperties HanaInstanceProperties
				err = json.Unmarshal(*v, &hanaInstanceProperties)
				if err != nil {
					return err
				}
				hi.HanaInstanceProperties = &hanaInstanceProperties
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				hi.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				hi.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				hi.Type = &typeVar
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				hi.Location = &location
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				hi.Tags = tags
			}
		}
	}

	return nil
}

// HanaInstanceProperties describes the properties of a HANA instance.
type HanaInstanceProperties struct {
	// HardwareProfile - Specifies the hardware settings for the HANA instance.
	HardwareProfile *HardwareProfile `json:"hardwareProfile,omitempty"`
	// StorageProfile - Specifies the storage settings for the HANA instance disks.
	StorageProfile *StorageProfile `json:"storageProfile,omitempty"`
	// OsProfile - Specifies the operating system settings for the HANA instance.
	OsProfile *OSProfile `json:"osProfile,omitempty"`
	// NetworkProfile - Specifies the network settings for the HANA instance.
	NetworkProfile *NetworkProfile `json:"networkProfile,omitempty"`
	// HanaInstanceID - READ-ONLY; Specifies the HANA instance unique ID.
	HanaInstanceID *string `json:"hanaInstanceId,omitempty"`
	// PowerState - READ-ONLY; Resource power state. Possible values include: 'Starting', 'Started', 'Stopping', 'Stopped', 'Restarting', 'Unknown'
	PowerState HanaInstancePowerStateEnum `json:"powerState,omitempty"`
	// ProximityPlacementGroup - READ-ONLY; Resource proximity placement group
	ProximityPlacementGroup *string `json:"proximityPlacementGroup,omitempty"`
	// HwRevision - READ-ONLY; Hardware revision of a HANA instance
	HwRevision *string `json:"hwRevision,omitempty"`
	// PartnerNodeID - ARM ID of another HanaInstance that will share a network with this HanaInstance
	PartnerNodeID *string `json:"partnerNodeId,omitempty"`
	// ProvisioningState - READ-ONLY; State of provisioning of the HanaInstance. Possible values include: 'Accepted', 'Creating', 'Updating', 'Failed', 'Succeeded', 'Deleting', 'Migrating'
	ProvisioningState HanaProvisioningStatesEnum `json:"provisioningState,omitempty"`
}

// HanaInstancesCreateFuture an abstraction for monitoring and retrieving the results of a long-running
// operation.
type HanaInstancesCreateFuture struct {
	azure.Future
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future *HanaInstancesCreateFuture) Result(client HanaInstancesClient) (hi HanaInstance, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hanaonazure.HanaInstancesCreateFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("hanaonazure.HanaInstancesCreateFuture")
		return
	}
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if hi.Response.Response, err = future.GetResult(sender); err == nil && hi.Response.Response.StatusCode != http.StatusNoContent {
		hi, err = client.CreateResponder(hi.Response.Response)
		if err != nil {
			err = autorest.NewErrorWithError(err, "hanaonazure.HanaInstancesCreateFuture", "Result", hi.Response.Response, "Failure responding to request")
		}
	}
	return
}

// HanaInstancesDeleteFuture an abstraction for monitoring and retrieving the results of a long-running
// operation.
type HanaInstancesDeleteFuture struct {
	azure.Future
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future *HanaInstancesDeleteFuture) Result(client HanaInstancesClient) (ar autorest.Response, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hanaonazure.HanaInstancesDeleteFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("hanaonazure.HanaInstancesDeleteFuture")
		return
	}
	ar.Response = future.Response()
	return
}

// HanaInstancesListResult the response from the List HANA Instances operation.
type HanaInstancesListResult struct {
	autorest.Response `json:"-"`
	// Value - The list of SAP HANA on Azure instances.
	Value *[]HanaInstance `json:"value,omitempty"`
	// NextLink - The URL to get the next set of HANA instances.
	NextLink *string `json:"nextLink,omitempty"`
}

// HanaInstancesListResultIterator provides access to a complete listing of HanaInstance values.
type HanaInstancesListResultIterator struct {
	i    int
	page HanaInstancesListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *HanaInstancesListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/HanaInstancesListResultIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *HanaInstancesListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter HanaInstancesListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter HanaInstancesListResultIterator) Response() HanaInstancesListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter HanaInstancesListResultIterator) Value() HanaInstance {
	if !iter.page.NotDone() {
		return HanaInstance{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the HanaInstancesListResultIterator type.
func NewHanaInstancesListResultIterator(page HanaInstancesListResultPage) HanaInstancesListResultIterator {
	return HanaInstancesListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (hilr HanaInstancesListResult) IsEmpty() bool {
	return hilr.Value == nil || len(*hilr.Value) == 0
}

// hanaInstancesListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (hilr HanaInstancesListResult) hanaInstancesListResultPreparer(ctx context.Context) (*http.Request, error) {
	if hilr.NextLink == nil || len(to.String(hilr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(hilr.NextLink)))
}

// HanaInstancesListResultPage contains a page of HanaInstance values.
type HanaInstancesListResultPage struct {
	fn   func(context.Context, HanaInstancesListResult) (HanaInstancesListResult, error)
	hilr HanaInstancesListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *HanaInstancesListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/HanaInstancesListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	next, err := page.fn(ctx, page.hilr)
	if err != nil {
		return err
	}
	page.hilr = next
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *HanaInstancesListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page HanaInstancesListResultPage) NotDone() bool {
	return !page.hilr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page HanaInstancesListResultPage) Response() HanaInstancesListResult {
	return page.hilr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page HanaInstancesListResultPage) Values() []HanaInstance {
	if page.hilr.IsEmpty() {
		return nil
	}
	return *page.hilr.Value
}

// Creates a new instance of the HanaInstancesListResultPage type.
func NewHanaInstancesListResultPage(getNextPage func(context.Context, HanaInstancesListResult) (HanaInstancesListResult, error)) HanaInstancesListResultPage {
	return HanaInstancesListResultPage{fn: getNextPage}
}

// HanaInstancesRestartFuture an abstraction for monitoring and retrieving the results of a long-running
// operation.
type HanaInstancesRestartFuture struct {
	azure.Future
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future *HanaInstancesRestartFuture) Result(client HanaInstancesClient) (ar autorest.Response, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hanaonazure.HanaInstancesRestartFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("hanaonazure.HanaInstancesRestartFuture")
		return
	}
	ar.Response = future.Response()
	return
}

// HanaInstancesShutdownFuture an abstraction for monitoring and retrieving the results of a long-running
// operation.
type HanaInstancesShutdownFuture struct {
	azure.Future
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future *HanaInstancesShutdownFuture) Result(client HanaInstancesClient) (ar autorest.Response, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hanaonazure.HanaInstancesShutdownFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("hanaonazure.HanaInstancesShutdownFuture")
		return
	}
	ar.Response = future.Response()
	return
}

// HanaInstancesStartFuture an abstraction for monitoring and retrieving the results of a long-running
// operation.
type HanaInstancesStartFuture struct {
	azure.Future
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future *HanaInstancesStartFuture) Result(client HanaInstancesClient) (ar autorest.Response, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hanaonazure.HanaInstancesStartFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("hanaonazure.HanaInstancesStartFuture")
		return
	}
	ar.Response = future.Response()
	return
}

// HardwareProfile specifies the hardware settings for the HANA instance.
type HardwareProfile struct {
	// HardwareType - READ-ONLY; Name of the hardware type (vendor and/or their product name). Possible values include: 'CiscoUCS', 'HPE'
	HardwareType HanaHardwareTypeNamesEnum `json:"hardwareType,omitempty"`
	// HanaInstanceSize - READ-ONLY; Specifies the HANA instance SKU. Possible values include: 'S72m', 'S144m', 'S72', 'S144', 'S192', 'S192m', 'S192xm', 'S96', 'S112', 'S224', 'S224m', 'S224om', 'S224oo', 'S224oom', 'S224ooo', 'S384', 'S384m', 'S384xm', 'S384xxm', 'S448', 'S448m', 'S448om', 'S448oo', 'S448oom', 'S448ooo', 'S576m', 'S576xm', 'S672', 'S672m', 'S672om', 'S672oo', 'S672oom', 'S672ooo', 'S768', 'S768m', 'S768xm', 'S896', 'S896m', 'S896om', 'S896oo', 'S896oom', 'S896ooo', 'S960m'
	HanaInstanceSize HanaInstanceSizeNamesEnum `json:"hanaInstanceSize,omitempty"`
}

// IPAddress specifies the IP address of the network interface.
type IPAddress struct {
	// IPAddress - Specifies the IP address of the network interface.
	IPAddress *string `json:"ipAddress,omitempty"`
}

// MonitoringDetails details needed to monitor a Hana Instance
type MonitoringDetails struct {
	// HanaSubnet - ARM ID of an Azure Subnet with access to the HANA instance.
	HanaSubnet *string `json:"hanaSubnet,omitempty"`
	// HanaHostname - Hostname of the HANA Instance blade.
	HanaHostname *string `json:"hanaHostname,omitempty"`
	// HanaDbName - Name of the database itself.
	HanaDbName *string `json:"hanaDbName,omitempty"`
	// HanaDbSQLPort - The port number of the tenant DB. Used to connect to the DB.
	HanaDbSQLPort *int32 `json:"hanaDbSqlPort,omitempty"`
	// HanaDbUsername - Username for the HANA database to login to for monitoring
	HanaDbUsername *string `json:"hanaDbUsername,omitempty"`
	// HanaDbPassword - Password for the HANA database to login for monitoring
	HanaDbPassword *string `json:"hanaDbPassword,omitempty"`
}

// NetworkProfile specifies the network settings for the HANA instance disks.
type NetworkProfile struct {
	// NetworkInterfaces - Specifies the network interfaces for the HANA instance.
	NetworkInterfaces *[]IPAddress `json:"networkInterfaces,omitempty"`
	// CircuitID - READ-ONLY; Specifies the circuit id for connecting to express route.
	CircuitID *string `json:"circuitId,omitempty"`
}

// Operation HANA operation information
type Operation struct {
	// Name - READ-ONLY; The name of the operation being performed on this particular object. This name should match the action name that appears in RBAC / the event service.
	Name *string `json:"name,omitempty"`
	// Display - Displayed HANA operation information
	Display *Display `json:"display,omitempty"`
}

// OperationList list of HANA operations
type OperationList struct {
	autorest.Response `json:"-"`
	// Value - List of HANA operations
	Value *[]Operation `json:"value,omitempty"`
}

// OSProfile specifies the operating system settings for the HANA instance.
type OSProfile struct {
	// ComputerName - Specifies the host OS name of the HANA instance.
	ComputerName *string `json:"computerName,omitempty"`
	// OsType - READ-ONLY; This property allows you to specify the type of the OS.
	OsType *string `json:"osType,omitempty"`
	// Version - READ-ONLY; Specifies version of operating system.
	Version *string `json:"version,omitempty"`
	// SSHPublicKey - Specifies the SSH public key used to access the operating system.
	SSHPublicKey *string `json:"sshPublicKey,omitempty"`
}

// Resource the resource model definition.
type Resource struct {
	// ID - READ-ONLY; Resource ID
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Resource name
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Resource type
	Type *string `json:"type,omitempty"`
	// Location - Resource location
	Location *string `json:"location,omitempty"`
	// Tags - READ-ONLY; Resource tags
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for Resource.
func (r Resource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if r.Location != nil {
		objectMap["location"] = r.Location
	}
	return json.Marshal(objectMap)
}

// SapMonitor SAP monitor info on Azure (ARM properties and SAP monitor properties)
type SapMonitor struct {
	autorest.Response `json:"-"`
	// SapMonitorProperties - SAP monitor properties
	*SapMonitorProperties `json:"properties,omitempty"`
	// ID - READ-ONLY; Resource ID
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Resource name
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Resource type
	Type *string `json:"type,omitempty"`
	// Location - Resource location
	Location *string `json:"location,omitempty"`
	// Tags - READ-ONLY; Resource tags
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for SapMonitor.
func (sm SapMonitor) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if sm.SapMonitorProperties != nil {
		objectMap["properties"] = sm.SapMonitorProperties
	}
	if sm.Location != nil {
		objectMap["location"] = sm.Location
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for SapMonitor struct.
func (sm *SapMonitor) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var sapMonitorProperties SapMonitorProperties
				err = json.Unmarshal(*v, &sapMonitorProperties)
				if err != nil {
					return err
				}
				sm.SapMonitorProperties = &sapMonitorProperties
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				sm.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				sm.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				sm.Type = &typeVar
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				sm.Location = &location
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				sm.Tags = tags
			}
		}
	}

	return nil
}

// SapMonitorListResult the response from the List SAP monitors operation.
type SapMonitorListResult struct {
	autorest.Response `json:"-"`
	// Value - The list of SAP monitors.
	Value *[]SapMonitor `json:"value,omitempty"`
	// NextLink - The URL to get the next set of SAP monitors.
	NextLink *string `json:"nextLink,omitempty"`
}

// SapMonitorListResultIterator provides access to a complete listing of SapMonitor values.
type SapMonitorListResultIterator struct {
	i    int
	page SapMonitorListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *SapMonitorListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SapMonitorListResultIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *SapMonitorListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter SapMonitorListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter SapMonitorListResultIterator) Response() SapMonitorListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter SapMonitorListResultIterator) Value() SapMonitor {
	if !iter.page.NotDone() {
		return SapMonitor{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the SapMonitorListResultIterator type.
func NewSapMonitorListResultIterator(page SapMonitorListResultPage) SapMonitorListResultIterator {
	return SapMonitorListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (smlr SapMonitorListResult) IsEmpty() bool {
	return smlr.Value == nil || len(*smlr.Value) == 0
}

// sapMonitorListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (smlr SapMonitorListResult) sapMonitorListResultPreparer(ctx context.Context) (*http.Request, error) {
	if smlr.NextLink == nil || len(to.String(smlr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(smlr.NextLink)))
}

// SapMonitorListResultPage contains a page of SapMonitor values.
type SapMonitorListResultPage struct {
	fn   func(context.Context, SapMonitorListResult) (SapMonitorListResult, error)
	smlr SapMonitorListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *SapMonitorListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SapMonitorListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	next, err := page.fn(ctx, page.smlr)
	if err != nil {
		return err
	}
	page.smlr = next
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *SapMonitorListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page SapMonitorListResultPage) NotDone() bool {
	return !page.smlr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page SapMonitorListResultPage) Response() SapMonitorListResult {
	return page.smlr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page SapMonitorListResultPage) Values() []SapMonitor {
	if page.smlr.IsEmpty() {
		return nil
	}
	return *page.smlr.Value
}

// Creates a new instance of the SapMonitorListResultPage type.
func NewSapMonitorListResultPage(getNextPage func(context.Context, SapMonitorListResult) (SapMonitorListResult, error)) SapMonitorListResultPage {
	return SapMonitorListResultPage{fn: getNextPage}
}

// SapMonitorProperties describes the properties of a SAP monitor.
type SapMonitorProperties struct {
	// HanaSubnet - Specifies the SAP monitor unique ID.
	HanaSubnet *string `json:"hanaSubnet,omitempty"`
	// HanaHostname - Hostname of the HANA instance.
	HanaHostname *string `json:"hanaHostname,omitempty"`
	// HanaDbName - Database name of the HANA instance.
	HanaDbName *string `json:"hanaDbName,omitempty"`
	// HanaDbSQLPort - Database port of the HANA instance.
	HanaDbSQLPort *int32 `json:"hanaDbSqlPort,omitempty"`
	// HanaDbUsername - Database username of the HANA instance.
	HanaDbUsername *string `json:"hanaDbUsername,omitempty"`
	// HanaDbPassword - Database password of the HANA instance.
	HanaDbPassword *string `json:"hanaDbPassword,omitempty"`
	// HanaDbPasswordKeyVaultURL - KeyVault URL link to the password for the HANA database.
	HanaDbPasswordKeyVaultURL *string `json:"hanaDbPasswordKeyVaultUrl,omitempty"`
	// HanaDbCredentialsMsiID - MSI ID passed by customer which has access to customer's KeyVault and to be assigned to the Collector VM.
	HanaDbCredentialsMsiID *string `json:"hanaDbCredentialsMsiId,omitempty"`
	// KeyVaultID - Key Vault ID containing customer's HANA credentials.
	KeyVaultID *string `json:"keyVaultId,omitempty"`
	// ProvisioningState - READ-ONLY; State of provisioning of the HanaInstance. Possible values include: 'Accepted', 'Creating', 'Updating', 'Failed', 'Succeeded', 'Deleting', 'Migrating'
	ProvisioningState HanaProvisioningStatesEnum `json:"provisioningState,omitempty"`
	// ManagedResourceGroupName - READ-ONLY; The name of the resource group the SAP Monitor resources get deployed into.
	ManagedResourceGroupName *string `json:"managedResourceGroupName,omitempty"`
	// LogAnalyticsWorkspaceArmID - The ARM ID of the Log Analytics Workspace that is used for monitoring
	LogAnalyticsWorkspaceArmID *string `json:"logAnalyticsWorkspaceArmId,omitempty"`
	// EnableCustomerAnalytics - The value indicating whether to send analytics to Microsoft
	EnableCustomerAnalytics *bool `json:"enableCustomerAnalytics,omitempty"`
	// LogAnalyticsWorkspaceID - The workspace ID of the log analytics workspace to be used for monitoring
	LogAnalyticsWorkspaceID *string `json:"logAnalyticsWorkspaceId,omitempty"`
	// LogAnalyticsWorkspaceSharedKey - The shared key of the log analytics workspace that is used for monitoring
	LogAnalyticsWorkspaceSharedKey *string `json:"logAnalyticsWorkspaceSharedKey,omitempty"`
}

// SapMonitorsCreateFuture an abstraction for monitoring and retrieving the results of a long-running
// operation.
type SapMonitorsCreateFuture struct {
	azure.Future
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future *SapMonitorsCreateFuture) Result(client SapMonitorsClient) (sm SapMonitor, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hanaonazure.SapMonitorsCreateFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("hanaonazure.SapMonitorsCreateFuture")
		return
	}
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if sm.Response.Response, err = future.GetResult(sender); err == nil && sm.Response.Response.StatusCode != http.StatusNoContent {
		sm, err = client.CreateResponder(sm.Response.Response)
		if err != nil {
			err = autorest.NewErrorWithError(err, "hanaonazure.SapMonitorsCreateFuture", "Result", sm.Response.Response, "Failure responding to request")
		}
	}
	return
}

// SapMonitorsDeleteFuture an abstraction for monitoring and retrieving the results of a long-running
// operation.
type SapMonitorsDeleteFuture struct {
	azure.Future
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future *SapMonitorsDeleteFuture) Result(client SapMonitorsClient) (ar autorest.Response, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hanaonazure.SapMonitorsDeleteFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("hanaonazure.SapMonitorsDeleteFuture")
		return
	}
	ar.Response = future.Response()
	return
}

// StorageProfile specifies the storage settings for the HANA instance disks.
type StorageProfile struct {
	// NfsIPAddress - READ-ONLY; IP Address to connect to storage.
	NfsIPAddress *string `json:"nfsIpAddress,omitempty"`
	// OsDisks - Specifies information about the operating system disk used by the hana instance.
	OsDisks *[]Disk `json:"osDisks,omitempty"`
}

// Tags tags field of the HANA instance.
type Tags struct {
	// Tags - Tags field of the HANA instance.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for Tags.
func (t Tags) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if t.Tags != nil {
		objectMap["tags"] = t.Tags
	}
	return json.Marshal(objectMap)
}
