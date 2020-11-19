package reservations

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
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/preview/reservations/mgmt/2018-06-01/reservations"

// AppliedReservationList ...
type AppliedReservationList struct {
	Value *[]string `json:"value,omitempty"`
	// NextLink - Url to get the next page of reservations
	NextLink *string `json:"nextLink,omitempty"`
}

// AppliedReservations ...
type AppliedReservations struct {
	autorest.Response `json:"-"`
	// ID - READ-ONLY; Identifier of the applied reservations
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Name of resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Type of resource. "Microsoft.Capacity/AppliedReservations"
	Type                           *string `json:"type,omitempty"`
	*AppliedReservationsProperties `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for AppliedReservations.
func (ar AppliedReservations) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if ar.AppliedReservationsProperties != nil {
		objectMap["properties"] = ar.AppliedReservationsProperties
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for AppliedReservations struct.
func (ar *AppliedReservations) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				ar.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				ar.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				ar.Type = &typeVar
			}
		case "properties":
			if v != nil {
				var appliedReservationsProperties AppliedReservationsProperties
				err = json.Unmarshal(*v, &appliedReservationsProperties)
				if err != nil {
					return err
				}
				ar.AppliedReservationsProperties = &appliedReservationsProperties
			}
		}
	}

	return nil
}

// AppliedReservationsProperties ...
type AppliedReservationsProperties struct {
	ReservationOrderIds *AppliedReservationList `json:"reservationOrderIds,omitempty"`
}

// Catalog ...
type Catalog struct {
	// ResourceType - READ-ONLY; The type of resource the SKU applies to.
	ResourceType *string `json:"resourceType,omitempty"`
	// Name - READ-ONLY; The name of SKU
	Name *string `json:"name,omitempty"`
	// Terms - READ-ONLY; Available reservation terms for this resource
	Terms *[]ReservationTerm `json:"terms,omitempty"`
	// Locations - READ-ONLY
	Locations *[]string `json:"locations,omitempty"`
	// SkuProperties - READ-ONLY
	SkuProperties *[]SkuProperty `json:"skuProperties,omitempty"`
	// Restrictions - READ-ONLY
	Restrictions *[]SkuRestriction `json:"restrictions,omitempty"`
}

// Error ...
type Error struct {
	Error *ExtendedErrorInfo `json:"error,omitempty"`
}

// ExtendedErrorInfo ...
type ExtendedErrorInfo struct {
	// Code - Possible values include: 'NotSpecified', 'InternalServerError', 'ServerTimeout', 'AuthorizationFailed', 'BadRequest', 'ClientCertificateThumbprintNotSet', 'InvalidRequestContent', 'OperationFailed', 'HTTPMethodNotSupported', 'InvalidRequestURI', 'MissingTenantID', 'InvalidTenantID', 'InvalidReservationOrderID', 'InvalidReservationID', 'ReservationIDNotInReservationOrder', 'ReservationOrderNotFound', 'InvalidSubscriptionID', 'InvalidAccessToken', 'InvalidLocationID', 'UnauthenticatedRequestsThrottled', 'InvalidHealthCheckType', 'Forbidden', 'BillingScopeIDCannotBeChanged', 'AppliedScopesNotAssociatedWithCommerceAccount', 'PatchValuesSameAsExisting', 'RoleAssignmentCreationFailed', 'ReservationOrderCreationFailed', 'ReservationOrderNotEnabled', 'CapacityUpdateScopesFailed', 'UnsupportedReservationTerm', 'ReservationOrderIDAlreadyExists', 'RiskCheckFailed', 'CreateQuoteFailed', 'ActivateQuoteFailed', 'NonsupportedAccountID', 'PaymentInstrumentNotFound', 'MissingAppliedScopesForSingle', 'NoValidReservationsToReRate', 'ReRateOnlyAllowedForEA', 'OperationCannotBePerformedInCurrentState', 'InvalidSingleAppliedScopesCount', 'InvalidFulfillmentRequestParameters', 'NotSupportedCountry', 'InvalidRefundQuantity', 'PurchaseError', 'BillingCustomerInputError', 'BillingPaymentInstrumentSoftError', 'BillingPaymentInstrumentHardError', 'BillingTransientError', 'BillingError', 'FulfillmentConfigurationError', 'FulfillmentOutOfStockError', 'FulfillmentTransientError', 'FulfillmentError', 'CalculatePriceFailed'
	Code    ErrorResponseCode `json:"code,omitempty"`
	Message *string           `json:"message,omitempty"`
}

// ExtendedStatusInfo ...
type ExtendedStatusInfo struct {
	// StatusCode - Possible values include: 'StatusCodeNone', 'StatusCodePending', 'StatusCodeActive', 'StatusCodePurchaseError', 'StatusCodePaymentInstrumentError', 'StatusCodeSplit', 'StatusCodeMerged', 'StatusCodeExpired', 'StatusCodeSucceeded'
	StatusCode StatusCode `json:"statusCode,omitempty"`
	// Message - The message giving detailed information about the status code.
	Message *string `json:"message,omitempty"`
}

// List ...
type List struct {
	autorest.Response `json:"-"`
	Value             *[]Response `json:"value,omitempty"`
	// NextLink - Url to get the next page of reservations.
	NextLink *string `json:"nextLink,omitempty"`
}

// ListCatalog ...
type ListCatalog struct {
	autorest.Response `json:"-"`
	Value             *[]Catalog `json:"value,omitempty"`
}

// ListIterator provides access to a complete listing of Response values.
type ListIterator struct {
	i    int
	page ListPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *ListIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ListIterator.NextWithContext")
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
func (iter *ListIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter ListIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter ListIterator) Response() List {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter ListIterator) Value() Response {
	if !iter.page.NotDone() {
		return Response{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the ListIterator type.
func NewListIterator(page ListPage) ListIterator {
	return ListIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (l List) IsEmpty() bool {
	return l.Value == nil || len(*l.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (l List) hasNextLink() bool {
	return l.NextLink != nil && len(*l.NextLink) != 0
}

// listPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (l List) listPreparer(ctx context.Context) (*http.Request, error) {
	if !l.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(l.NextLink)))
}

// ListPage contains a page of Response values.
type ListPage struct {
	fn func(context.Context, List) (List, error)
	l  List
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *ListPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ListPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.l)
		if err != nil {
			return err
		}
		page.l = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *ListPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page ListPage) NotDone() bool {
	return !page.l.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page ListPage) Response() List {
	return page.l
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page ListPage) Values() []Response {
	if page.l.IsEmpty() {
		return nil
	}
	return *page.l.Value
}

// Creates a new instance of the ListPage type.
func NewListPage(cur List, getNextPage func(context.Context, List) (List, error)) ListPage {
	return ListPage{
		fn: getNextPage,
		l:  cur,
	}
}

// ListResponse ...
type ListResponse struct {
	autorest.Response `json:"-"`
	Value             *[]Response `json:"value,omitempty"`
}

// MergeProperties ...
type MergeProperties struct {
	// Sources - Format of the resource id should be /providers/Microsoft.Capacity/reservationOrders/{reservationOrderId}/reservations/{reservationId}
	Sources *[]string `json:"sources,omitempty"`
}

// MergePropertiesType ...
type MergePropertiesType struct {
	// MergeDestination - Reservation Resource Id Created due to the merge. Format of the resource Id is /providers/Microsoft.Capacity/reservationOrders/{reservationOrderId}/reservations/{reservationId}
	MergeDestination *string `json:"mergeDestination,omitempty"`
	// MergeSources - Resource Ids of the Source Reservation's merged to form this Reservation. Format of the resource Id is /providers/Microsoft.Capacity/reservationOrders/{reservationOrderId}/reservations/{reservationId}
	MergeSources *[]string `json:"mergeSources,omitempty"`
}

// MergeRequest ...
type MergeRequest struct {
	*MergeProperties `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for MergeRequest.
func (mr MergeRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if mr.MergeProperties != nil {
		objectMap["properties"] = mr.MergeProperties
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for MergeRequest struct.
func (mr *MergeRequest) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var mergeProperties MergeProperties
				err = json.Unmarshal(*v, &mergeProperties)
				if err != nil {
					return err
				}
				mr.MergeProperties = &mergeProperties
			}
		}
	}

	return nil
}

// OperationDisplay ...
type OperationDisplay struct {
	Provider    *string `json:"provider,omitempty"`
	Resource    *string `json:"resource,omitempty"`
	Operation   *string `json:"operation,omitempty"`
	Description *string `json:"description,omitempty"`
}

// OperationList ...
type OperationList struct {
	autorest.Response `json:"-"`
	Value             *[]OperationResponse `json:"value,omitempty"`
	// NextLink - Url to get the next page of items.
	NextLink *string `json:"nextLink,omitempty"`
}

// OperationListIterator provides access to a complete listing of OperationResponse values.
type OperationListIterator struct {
	i    int
	page OperationListPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *OperationListIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationListIterator.NextWithContext")
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
func (iter *OperationListIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter OperationListIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter OperationListIterator) Response() OperationList {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter OperationListIterator) Value() OperationResponse {
	if !iter.page.NotDone() {
		return OperationResponse{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the OperationListIterator type.
func NewOperationListIterator(page OperationListPage) OperationListIterator {
	return OperationListIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (ol OperationList) IsEmpty() bool {
	return ol.Value == nil || len(*ol.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (ol OperationList) hasNextLink() bool {
	return ol.NextLink != nil && len(*ol.NextLink) != 0
}

// operationListPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (ol OperationList) operationListPreparer(ctx context.Context) (*http.Request, error) {
	if !ol.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(ol.NextLink)))
}

// OperationListPage contains a page of OperationResponse values.
type OperationListPage struct {
	fn func(context.Context, OperationList) (OperationList, error)
	ol OperationList
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *OperationListPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationListPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.ol)
		if err != nil {
			return err
		}
		page.ol = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *OperationListPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page OperationListPage) NotDone() bool {
	return !page.ol.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page OperationListPage) Response() OperationList {
	return page.ol
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page OperationListPage) Values() []OperationResponse {
	if page.ol.IsEmpty() {
		return nil
	}
	return *page.ol.Value
}

// Creates a new instance of the OperationListPage type.
func NewOperationListPage(cur OperationList, getNextPage func(context.Context, OperationList) (OperationList, error)) OperationListPage {
	return OperationListPage{
		fn: getNextPage,
		ol: cur,
	}
}

// OperationResponse ...
type OperationResponse struct {
	Name    *string           `json:"name,omitempty"`
	Display *OperationDisplay `json:"display,omitempty"`
	Origin  *string           `json:"origin,omitempty"`
}

// OrderList ...
type OrderList struct {
	autorest.Response `json:"-"`
	Value             *[]OrderResponse `json:"value,omitempty"`
	// NextLink - Url to get the next page of reservationOrders.
	NextLink *string `json:"nextLink,omitempty"`
}

// OrderListIterator provides access to a complete listing of OrderResponse values.
type OrderListIterator struct {
	i    int
	page OrderListPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *OrderListIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OrderListIterator.NextWithContext")
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
func (iter *OrderListIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter OrderListIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter OrderListIterator) Response() OrderList {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter OrderListIterator) Value() OrderResponse {
	if !iter.page.NotDone() {
		return OrderResponse{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the OrderListIterator type.
func NewOrderListIterator(page OrderListPage) OrderListIterator {
	return OrderListIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (ol OrderList) IsEmpty() bool {
	return ol.Value == nil || len(*ol.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (ol OrderList) hasNextLink() bool {
	return ol.NextLink != nil && len(*ol.NextLink) != 0
}

// orderListPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (ol OrderList) orderListPreparer(ctx context.Context) (*http.Request, error) {
	if !ol.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(ol.NextLink)))
}

// OrderListPage contains a page of OrderResponse values.
type OrderListPage struct {
	fn func(context.Context, OrderList) (OrderList, error)
	ol OrderList
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *OrderListPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OrderListPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.ol)
		if err != nil {
			return err
		}
		page.ol = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *OrderListPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page OrderListPage) NotDone() bool {
	return !page.ol.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page OrderListPage) Response() OrderList {
	return page.ol
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page OrderListPage) Values() []OrderResponse {
	if page.ol.IsEmpty() {
		return nil
	}
	return *page.ol.Value
}

// Creates a new instance of the OrderListPage type.
func NewOrderListPage(cur OrderList, getNextPage func(context.Context, OrderList) (OrderList, error)) OrderListPage {
	return OrderListPage{
		fn: getNextPage,
		ol: cur,
	}
}

// OrderProperties ...
type OrderProperties struct {
	// DisplayName - Friendly name for user to easily identified the reservation.
	DisplayName *string `json:"displayName,omitempty"`
	// RequestDateTime - This is the DateTime when the reservation was initially requested for purchase.
	RequestDateTime *date.Time `json:"requestDateTime,omitempty"`
	// CreatedDateTime - This is the DateTime when the reservation was created.
	CreatedDateTime *date.Time `json:"createdDateTime,omitempty"`
	// ExpiryDate - This is the date when the Reservation will expire.
	ExpiryDate *date.Date `json:"expiryDate,omitempty"`
	// OriginalQuantity - Total Quantity of the SKUs purchased in the Reservation.
	OriginalQuantity *int32 `json:"originalQuantity,omitempty"`
	// Term - Possible values include: 'P1Y', 'P3Y'
	Term ReservationTerm `json:"term,omitempty"`
	// ProvisioningState - Current state of the reservation.
	ProvisioningState    *string     `json:"provisioningState,omitempty"`
	ReservationsProperty *[]Response `json:"reservations,omitempty"`
}

// OrderResponse ...
type OrderResponse struct {
	autorest.Response `json:"-"`
	Etag              *int32 `json:"etag,omitempty"`
	// ID - READ-ONLY; Identifier of the reservation
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Name of the reservation
	Name             *string `json:"name,omitempty"`
	*OrderProperties `json:"properties,omitempty"`
	// Type - READ-ONLY; Type of resource. "Microsoft.Capacity/reservations"
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for OrderResponse.
func (or OrderResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if or.Etag != nil {
		objectMap["etag"] = or.Etag
	}
	if or.OrderProperties != nil {
		objectMap["properties"] = or.OrderProperties
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for OrderResponse struct.
func (or *OrderResponse) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "etag":
			if v != nil {
				var etag int32
				err = json.Unmarshal(*v, &etag)
				if err != nil {
					return err
				}
				or.Etag = &etag
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				or.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				or.Name = &name
			}
		case "properties":
			if v != nil {
				var orderProperties OrderProperties
				err = json.Unmarshal(*v, &orderProperties)
				if err != nil {
					return err
				}
				or.OrderProperties = &orderProperties
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				or.Type = &typeVar
			}
		}
	}

	return nil
}

// Patch ...
type Patch struct {
	*PatchProperties `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for Patch.
func (p Patch) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if p.PatchProperties != nil {
		objectMap["properties"] = p.PatchProperties
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for Patch struct.
func (p *Patch) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var patchProperties PatchProperties
				err = json.Unmarshal(*v, &patchProperties)
				if err != nil {
					return err
				}
				p.PatchProperties = &patchProperties
			}
		}
	}

	return nil
}

// PatchProperties ...
type PatchProperties struct {
	// AppliedScopeType - Possible values include: 'Single', 'Shared'
	AppliedScopeType AppliedScopeType `json:"appliedScopeType,omitempty"`
	AppliedScopes    *[]string        `json:"appliedScopes,omitempty"`
	// InstanceFlexibility - Possible values include: 'On', 'Off', 'NotSupported'
	InstanceFlexibility InstanceFlexibility `json:"instanceFlexibility,omitempty"`
	// Name - Name of the Reservation
	Name *string `json:"name,omitempty"`
}

// Properties ...
type Properties struct {
	// ReservedResourceType - Possible values include: 'VirtualMachines', 'SQLDatabases', 'SuseLinux', 'CosmosDb', 'RedHat'
	ReservedResourceType ReservedResourceType `json:"reservedResourceType,omitempty"`
	// InstanceFlexibility - Possible values include: 'On', 'Off', 'NotSupported'
	InstanceFlexibility InstanceFlexibility `json:"instanceFlexibility,omitempty"`
	// DisplayName - Friendly name for user to easily identify the reservation
	DisplayName   *string   `json:"displayName,omitempty"`
	AppliedScopes *[]string `json:"appliedScopes,omitempty"`
	// AppliedScopeType - Possible values include: 'Single', 'Shared'
	AppliedScopeType AppliedScopeType `json:"appliedScopeType,omitempty"`
	// Quantity - Quantity of the SKUs that are part of the Reservation.
	Quantity *int32 `json:"quantity,omitempty"`
	// ProvisioningState - Current state of the reservation.
	ProvisioningState *string `json:"provisioningState,omitempty"`
	// EffectiveDateTime - DateTime of the Reservation starting when this version is effective from.
	EffectiveDateTime *date.Time `json:"effectiveDateTime,omitempty"`
	// LastUpdatedDateTime - READ-ONLY; DateTime of the last time the Reservation was updated.
	LastUpdatedDateTime *date.Time `json:"lastUpdatedDateTime,omitempty"`
	// ExpiryDate - This is the date when the Reservation will expire.
	ExpiryDate *date.Date `json:"expiryDate,omitempty"`
	// SkuDescription - Description of the SKU in english.
	SkuDescription     *string              `json:"skuDescription,omitempty"`
	ExtendedStatusInfo *ExtendedStatusInfo  `json:"extendedStatusInfo,omitempty"`
	SplitProperties    *SplitPropertiesType `json:"splitProperties,omitempty"`
	MergeProperties    *MergePropertiesType `json:"mergeProperties,omitempty"`
}

// MarshalJSON is the custom marshaler for Properties.
func (p Properties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if p.ReservedResourceType != "" {
		objectMap["reservedResourceType"] = p.ReservedResourceType
	}
	if p.InstanceFlexibility != "" {
		objectMap["instanceFlexibility"] = p.InstanceFlexibility
	}
	if p.DisplayName != nil {
		objectMap["displayName"] = p.DisplayName
	}
	if p.AppliedScopes != nil {
		objectMap["appliedScopes"] = p.AppliedScopes
	}
	if p.AppliedScopeType != "" {
		objectMap["appliedScopeType"] = p.AppliedScopeType
	}
	if p.Quantity != nil {
		objectMap["quantity"] = p.Quantity
	}
	if p.ProvisioningState != nil {
		objectMap["provisioningState"] = p.ProvisioningState
	}
	if p.EffectiveDateTime != nil {
		objectMap["effectiveDateTime"] = p.EffectiveDateTime
	}
	if p.ExpiryDate != nil {
		objectMap["expiryDate"] = p.ExpiryDate
	}
	if p.SkuDescription != nil {
		objectMap["skuDescription"] = p.SkuDescription
	}
	if p.ExtendedStatusInfo != nil {
		objectMap["extendedStatusInfo"] = p.ExtendedStatusInfo
	}
	if p.SplitProperties != nil {
		objectMap["splitProperties"] = p.SplitProperties
	}
	if p.MergeProperties != nil {
		objectMap["mergeProperties"] = p.MergeProperties
	}
	return json.Marshal(objectMap)
}

// ReservationMergeFuture an abstraction for monitoring and retrieving the results of a long-running
// operation.
type ReservationMergeFuture struct {
	azure.Future
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future *ReservationMergeFuture) Result(client Client) (lr ListResponse, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservations.ReservationMergeFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("reservations.ReservationMergeFuture")
		return
	}
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if lr.Response.Response, err = future.GetResult(sender); err == nil && lr.Response.Response.StatusCode != http.StatusNoContent {
		lr, err = client.MergeResponder(lr.Response.Response)
		if err != nil {
			err = autorest.NewErrorWithError(err, "reservations.ReservationMergeFuture", "Result", lr.Response.Response, "Failure responding to request")
		}
	}
	return
}

// ReservationUpdateFuture an abstraction for monitoring and retrieving the results of a long-running
// operation.
type ReservationUpdateFuture struct {
	azure.Future
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future *ReservationUpdateFuture) Result(client Client) (r Response, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservations.ReservationUpdateFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("reservations.ReservationUpdateFuture")
		return
	}
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if r.Response.Response, err = future.GetResult(sender); err == nil && r.Response.Response.StatusCode != http.StatusNoContent {
		r, err = client.UpdateResponder(r.Response.Response)
		if err != nil {
			err = autorest.NewErrorWithError(err, "reservations.ReservationUpdateFuture", "Result", r.Response.Response, "Failure responding to request")
		}
	}
	return
}

// Response ...
type Response struct {
	autorest.Response `json:"-"`
	// Location - READ-ONLY; The Azure Region where the reserved resource lives.
	Location *string `json:"location,omitempty"`
	Etag     *int32  `json:"etag,omitempty"`
	// ID - READ-ONLY; Identifier of the reservation
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Name of the reservation
	Name       *string     `json:"name,omitempty"`
	Sku        *SkuName    `json:"sku,omitempty"`
	Properties *Properties `json:"properties,omitempty"`
	// Type - READ-ONLY; Type of resource. "Microsoft.Capacity/reservationOrders/reservations"
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for Response.
func (r Response) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if r.Etag != nil {
		objectMap["etag"] = r.Etag
	}
	if r.Sku != nil {
		objectMap["sku"] = r.Sku
	}
	if r.Properties != nil {
		objectMap["properties"] = r.Properties
	}
	return json.Marshal(objectMap)
}

// SkuName ...
type SkuName struct {
	Name *string `json:"name,omitempty"`
}

// SkuProperty ...
type SkuProperty struct {
	// Name - An invariant to describe the feature.
	Name *string `json:"name,omitempty"`
	// Value - An invariant if the feature is measured by quantity.
	Value *string `json:"value,omitempty"`
}

// SkuRestriction ...
type SkuRestriction struct {
	// Type - The type of restrictions.
	Type *string `json:"type,omitempty"`
	// Values - The value of restrictions. If the restriction type is set to location. This would be different locations where the SKU is restricted.
	Values *[]string `json:"values,omitempty"`
	// ReasonCode - The reason for restriction.
	ReasonCode *string `json:"reasonCode,omitempty"`
}

// SplitFuture an abstraction for monitoring and retrieving the results of a long-running operation.
type SplitFuture struct {
	azure.Future
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future *SplitFuture) Result(client Client) (lr ListResponse, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservations.SplitFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("reservations.SplitFuture")
		return
	}
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if lr.Response.Response, err = future.GetResult(sender); err == nil && lr.Response.Response.StatusCode != http.StatusNoContent {
		lr, err = client.SplitResponder(lr.Response.Response)
		if err != nil {
			err = autorest.NewErrorWithError(err, "reservations.SplitFuture", "Result", lr.Response.Response, "Failure responding to request")
		}
	}
	return
}

// SplitProperties ...
type SplitProperties struct {
	// Quantities - List of the quantities in the new reservations to create.
	Quantities *[]int32 `json:"quantities,omitempty"`
	// ReservationID - Resource id of the reservation to be split. Format of the resource id should be /providers/Microsoft.Capacity/reservationOrders/{reservationOrderId}/reservations/{reservationId}
	ReservationID *string `json:"reservationId,omitempty"`
}

// SplitPropertiesType ...
type SplitPropertiesType struct {
	// SplitDestinations - List of destination Resource Id that are created due to split. Format of the resource Id is /providers/Microsoft.Capacity/reservationOrders/{reservationOrderId}/reservations/{reservationId}
	SplitDestinations *[]string `json:"splitDestinations,omitempty"`
	// SplitSource - Resource Id of the Reservation from which this is split. Format of the resource Id is /providers/Microsoft.Capacity/reservationOrders/{reservationOrderId}/reservations/{reservationId}
	SplitSource *string `json:"splitSource,omitempty"`
}

// SplitRequest ...
type SplitRequest struct {
	*SplitProperties `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for SplitRequest.
func (sr SplitRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if sr.SplitProperties != nil {
		objectMap["properties"] = sr.SplitProperties
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for SplitRequest struct.
func (sr *SplitRequest) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var splitProperties SplitProperties
				err = json.Unmarshal(*v, &splitProperties)
				if err != nil {
					return err
				}
				sr.SplitProperties = &splitProperties
			}
		}
	}

	return nil
}
