package resourcegraph

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
    "github.com/Azure/go-autorest/autorest/date"
    "github.com/Azure/go-autorest/autorest/to"
    "github.com/Azure/go-autorest/tracing"
    "net/http"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/resourcegraph/mgmt/2019-04-01/resourcegraph"

        // ColumnDataType enumerates the values for column data type.
    type ColumnDataType string

    const (
                // Boolean ...
        Boolean ColumnDataType = "boolean"
                // Integer ...
        Integer ColumnDataType = "integer"
                // Number ...
        Number ColumnDataType = "number"
                // Object ...
        Object ColumnDataType = "object"
                // String ...
        String ColumnDataType = "string"
            )
    // PossibleColumnDataTypeValues returns an array of possible values for the ColumnDataType const type.
    func PossibleColumnDataTypeValues() []ColumnDataType {
        return []ColumnDataType{Boolean,Integer,Number,Object,String}
    }

        // FacetSortOrder enumerates the values for facet sort order.
    type FacetSortOrder string

    const (
                // Asc ...
        Asc FacetSortOrder = "asc"
                // Desc ...
        Desc FacetSortOrder = "desc"
            )
    // PossibleFacetSortOrderValues returns an array of possible values for the FacetSortOrder const type.
    func PossibleFacetSortOrderValues() []FacetSortOrder {
        return []FacetSortOrder{Asc,Desc}
    }

        // ResultFormat enumerates the values for result format.
    type ResultFormat string

    const (
                // ResultFormatObjectArray ...
        ResultFormatObjectArray ResultFormat = "objectArray"
                // ResultFormatTable ...
        ResultFormatTable ResultFormat = "table"
            )
    // PossibleResultFormatValues returns an array of possible values for the ResultFormat const type.
    func PossibleResultFormatValues() []ResultFormat {
        return []ResultFormat{ResultFormatObjectArray,ResultFormatTable}
    }

        // ResultKind enumerates the values for result kind.
    type ResultKind string

    const (
                // Basic ...
        Basic ResultKind = "basic"
            )
    // PossibleResultKindValues returns an array of possible values for the ResultKind const type.
    func PossibleResultKindValues() []ResultKind {
        return []ResultKind{Basic}
    }

        // ResultTruncated enumerates the values for result truncated.
    type ResultTruncated string

    const (
                // False ...
        False ResultTruncated = "false"
                // True ...
        True ResultTruncated = "true"
            )
    // PossibleResultTruncatedValues returns an array of possible values for the ResultTruncated const type.
    func PossibleResultTruncatedValues() []ResultTruncated {
        return []ResultTruncated{False,True}
    }

        // ResultType enumerates the values for result type.
    type ResultType string

    const (
                // ResultTypeFacet ...
        ResultTypeFacet ResultType = "Facet"
                // ResultTypeFacetError ...
        ResultTypeFacetError ResultType = "FacetError"
                // ResultTypeFacetResult ...
        ResultTypeFacetResult ResultType = "FacetResult"
            )
    // PossibleResultTypeValues returns an array of possible values for the ResultType const type.
    func PossibleResultTypeValues() []ResultType {
        return []ResultType{ResultTypeFacet,ResultTypeFacetError,ResultTypeFacetResult}
    }

            // Column query result column descriptor.
            type Column struct {
            // Name - Column name.
            Name *string `json:"name,omitempty"`
            // Type - Column data type. Possible values include: 'String', 'Integer', 'Number', 'Boolean', 'Object'
            Type ColumnDataType `json:"type,omitempty"`
            }

            // Error error details.
            type Error struct {
            // Code - Error code identifying the specific error.
            Code *string `json:"code,omitempty"`
            // Message - A human readable error message.
            Message *string `json:"message,omitempty"`
            // Details - Error details
            Details *[]ErrorDetails `json:"details,omitempty"`
            }

            // ErrorDetails ...
            type ErrorDetails struct {
            // AdditionalProperties - Unmatched properties from the message are deserialized this collection
            AdditionalProperties map[string]interface{} `json:""`
            // Code - Error code identifying the specific error.
            Code *string `json:"code,omitempty"`
            // Message - A human readable error message.
            Message *string `json:"message,omitempty"`
            }

        // MarshalJSON is the custom marshaler for ErrorDetails.
        func (ed ErrorDetails)MarshalJSON() ([]byte, error){
        objectMap := make(map[string]interface{})
                if(ed.Code != nil) {
                objectMap["code"] = ed.Code
                }
                if(ed.Message != nil) {
                objectMap["message"] = ed.Message
                }
                    for k, v := range ed.AdditionalProperties {
                objectMap[k] = v
            }
        return json.Marshal(objectMap)
        }
        // UnmarshalJSON is the custom unmarshaler for ErrorDetails struct.
        func (ed *ErrorDetails) UnmarshalJSON(body []byte) error {
        var m map[string]*json.RawMessage
        err := json.Unmarshal(body, &m)
        if err != nil {
        return err
        }
        for k, v := range  m {
        switch k {
        default:
    if v != nil {
        var additionalProperties interface{}
        err = json.Unmarshal(*v, &additionalProperties)
    if err != nil {
    return err
    }
        if(ed.AdditionalProperties == nil) {
         ed.AdditionalProperties = make(map[string]interface{})
        }
            ed.AdditionalProperties[k] = additionalProperties
    }
                case "code":
    if v != nil {
        var code string
        err = json.Unmarshal(*v, &code)
    if err != nil {
    return err
    }
        ed.Code = &code
    }
                case "message":
    if v != nil {
        var message string
        err = json.Unmarshal(*v, &message)
    if err != nil {
    return err
    }
        ed.Message = &message
    }
            }
        }

        return nil
        }

            // ErrorFieldContract error Field contract.
            type ErrorFieldContract struct {
            // Code - Property level error code.
            Code *string `json:"code,omitempty"`
            // Message - Human-readable representation of property-level error.
            Message *string `json:"message,omitempty"`
            // Target - Property name.
            Target *string `json:"target,omitempty"`
            }

            // ErrorResponse an error response from the API.
            type ErrorResponse struct {
            // Error - Error information.
            Error *Error `json:"error,omitempty"`
            }

// BasicFacet a facet containing additional statistics on the response of a query. Can be either FacetResult or
// FacetError.
        type BasicFacet interface {
            AsFacetResult () (*FacetResult, bool)
            AsFacetError () (*FacetError, bool)
        AsFacet () (*Facet, bool)
        }

        // Facet a facet containing additional statistics on the response of a query. Can be either FacetResult or
        // FacetError.
        type Facet struct {
        // Expression - Facet expression, same as in the corresponding facet request.
        Expression *string `json:"expression,omitempty"`
        // ResultType - Possible values include: 'ResultTypeFacet', 'ResultTypeFacetResult', 'ResultTypeFacetError'
        ResultType ResultType `json:"resultType,omitempty"`
        }

        func unmarshalBasicFacet(body []byte) (BasicFacet, error){
        var m map[string]interface{}
        err := json.Unmarshal(body, &m)
        if err != nil {
        return nil, err
        }

        switch m["resultType"] {
            case string(ResultTypeFacetResult):
            var fr FacetResult
            err := json.Unmarshal(body, &fr)
            return fr, err
            case string(ResultTypeFacetError):
            var fe FacetError
            err := json.Unmarshal(body, &fe)
            return fe, err
            default:
        var f Facet
        err := json.Unmarshal(body, &f)
        return f, err
        }
        }
        func unmarshalBasicFacetArray(body []byte) ([]BasicFacet, error){
        var rawMessages []*json.RawMessage
        err := json.Unmarshal(body, &rawMessages)
        if err != nil {
        return nil, err
        }

        fArray := make([]BasicFacet, len(rawMessages))

        for index, rawMessage := range rawMessages {
        f, err := unmarshalBasicFacet(*rawMessage)
        if err != nil {
        return nil, err
        }
        fArray[index] = f
        }
        return fArray, nil
        }

        // MarshalJSON is the custom marshaler for Facet.
        func (f Facet)MarshalJSON() ([]byte, error){
            f.ResultType = ResultTypeFacet
            objectMap := make(map[string]interface{})
                if(f.Expression != nil) {
                objectMap["expression"] = f.Expression
                }
                if(f.ResultType != "") {
                objectMap["resultType"] = f.ResultType
                }
                return json.Marshal(objectMap)
        }

            // AsFacetResult is the BasicFacet implementation for Facet.
            func (f Facet) AsFacetResult() (*FacetResult, bool) {
                return nil, false
            }

            // AsFacetError is the BasicFacet implementation for Facet.
            func (f Facet) AsFacetError() (*FacetError, bool) {
                return nil, false
            }

            // AsFacet is the BasicFacet implementation for Facet.
            func (f Facet) AsFacet() (*Facet, bool) {
                return &f, true
            }

                // AsBasicFacet is the BasicFacet implementation for Facet.
                func(f Facet) AsBasicFacet()(BasicFacet, bool) {
                    return &f, true
                }


            // FacetError a facet whose execution resulted in an error.
            type FacetError struct {
            // Errors - An array containing detected facet errors with details.
            Errors *[]ErrorDetails `json:"errors,omitempty"`
            // Expression - Facet expression, same as in the corresponding facet request.
            Expression *string `json:"expression,omitempty"`
            // ResultType - Possible values include: 'ResultTypeFacet', 'ResultTypeFacetResult', 'ResultTypeFacetError'
            ResultType ResultType `json:"resultType,omitempty"`
            }

        // MarshalJSON is the custom marshaler for FacetError.
        func (fe FacetError)MarshalJSON() ([]byte, error){
            fe.ResultType = ResultTypeFacetError
            objectMap := make(map[string]interface{})
                if(fe.Errors != nil) {
                objectMap["errors"] = fe.Errors
                }
                if(fe.Expression != nil) {
                objectMap["expression"] = fe.Expression
                }
                if(fe.ResultType != "") {
                objectMap["resultType"] = fe.ResultType
                }
                return json.Marshal(objectMap)
        }

            // AsFacetResult is the BasicFacet implementation for FacetError.
            func (fe FacetError) AsFacetResult() (*FacetResult, bool) {
                return nil, false
            }

            // AsFacetError is the BasicFacet implementation for FacetError.
            func (fe FacetError) AsFacetError() (*FacetError, bool) {
                return &fe, true
            }

            // AsFacet is the BasicFacet implementation for FacetError.
            func (fe FacetError) AsFacet() (*Facet, bool) {
                return nil, false
            }

                // AsBasicFacet is the BasicFacet implementation for FacetError.
                func(fe FacetError) AsBasicFacet()(BasicFacet, bool) {
                    return &fe, true
                }


            // FacetRequest a request to compute additional statistics (facets) over the query results.
            type FacetRequest struct {
            // Expression - The column or list of columns to summarize by
            Expression *string `json:"expression,omitempty"`
            // Options - The options for facet evaluation
            Options *FacetRequestOptions `json:"options,omitempty"`
            }

            // FacetRequestOptions the options for facet evaluation
            type FacetRequestOptions struct {
            // SortBy - The column name or query expression to sort on. Defaults to count if not present.
            SortBy *string `json:"sortBy,omitempty"`
            // SortOrder - The sorting order by the selected column (count by default). Possible values include: 'Asc', 'Desc'
            SortOrder FacetSortOrder `json:"sortOrder,omitempty"`
            // Filter - Specifies the filter condition for the 'where' clause which will be run on main query's result, just before the actual faceting.
            Filter *string `json:"filter,omitempty"`
            // Top - The maximum number of facet rows that should be returned.
            Top *int32 `json:"$top,omitempty"`
            }

            // FacetResult successfully executed facet containing additional statistics on the response of a query.
            type FacetResult struct {
            // TotalRecords - Number of total records in the facet results.
            TotalRecords *int64 `json:"totalRecords,omitempty"`
            // Count - Number of records returned in the facet response.
            Count *int32 `json:"count,omitempty"`
            // Data - A table containing the desired facets. Only present if the facet is valid.
            Data interface{} `json:"data,omitempty"`
            // Expression - Facet expression, same as in the corresponding facet request.
            Expression *string `json:"expression,omitempty"`
            // ResultType - Possible values include: 'ResultTypeFacet', 'ResultTypeFacetResult', 'ResultTypeFacetError'
            ResultType ResultType `json:"resultType,omitempty"`
            }

        // MarshalJSON is the custom marshaler for FacetResult.
        func (fr FacetResult)MarshalJSON() ([]byte, error){
            fr.ResultType = ResultTypeFacetResult
            objectMap := make(map[string]interface{})
                if(fr.TotalRecords != nil) {
                objectMap["totalRecords"] = fr.TotalRecords
                }
                if(fr.Count != nil) {
                objectMap["count"] = fr.Count
                }
                if(fr.Data != nil) {
                objectMap["data"] = fr.Data
                }
                if(fr.Expression != nil) {
                objectMap["expression"] = fr.Expression
                }
                if(fr.ResultType != "") {
                objectMap["resultType"] = fr.ResultType
                }
                return json.Marshal(objectMap)
        }

            // AsFacetResult is the BasicFacet implementation for FacetResult.
            func (fr FacetResult) AsFacetResult() (*FacetResult, bool) {
                return &fr, true
            }

            // AsFacetError is the BasicFacet implementation for FacetResult.
            func (fr FacetResult) AsFacetError() (*FacetError, bool) {
                return nil, false
            }

            // AsFacet is the BasicFacet implementation for FacetResult.
            func (fr FacetResult) AsFacet() (*Facet, bool) {
                return nil, false
            }

                // AsBasicFacet is the BasicFacet implementation for FacetResult.
                func(fr FacetResult) AsBasicFacet()(BasicFacet, bool) {
                    return &fr, true
                }


            // GraphQueryError error message body that will indicate why the operation failed.
            type GraphQueryError struct {
            // Code - Service-defined error code. This code serves as a sub-status for the HTTP error code specified in the response.
            Code *string `json:"code,omitempty"`
            // Message - Human-readable representation of the error.
            Message *string `json:"message,omitempty"`
            // Details - The list of invalid fields send in request, in case of validation error.
            Details *[]ErrorFieldContract `json:"details,omitempty"`
            }

            // GraphQueryListResult graph query list result.
            type GraphQueryListResult struct {
            autorest.Response `json:"-"`
            // NextLink - URL to fetch the next set of queries.
            NextLink *string `json:"nextLink,omitempty"`
            // Value - READ-ONLY; An array of graph queries.
            Value *[]GraphQueryResource `json:"value,omitempty"`
            }

            // GraphQueryListResultIterator provides access to a complete listing of GraphQueryResource values.
            type GraphQueryListResultIterator struct {
                i int
                page GraphQueryListResultPage
            }
        // NextWithContext advances to the next value.  If there was an error making
        // the request the iterator does not advance and the error is returned.
        func (iter * GraphQueryListResultIterator) NextWithContext(ctx context.Context) (err error) {
        if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/GraphQueryListResultIterator.NextWithContext")
        defer func() {
        sc := -1
        if iter.Response().Response.Response != nil {
        sc = iter.Response().Response.Response.StatusCode
        }
        tracing.EndSpan(ctx, sc, err)
        }()
        }
        iter.i++
        if iter.i < len(iter. page.Values()) {
        return nil
        }
        err = iter.page.NextWithContext(ctx)
        if err != nil {
        iter. i--
        return err
        }
        iter.i = 0
        return nil
        }
        // Next advances to the next value.  If there was an error making
        // the request the iterator does not advance and the error is returned.
        // Deprecated: Use NextWithContext() instead.
        func (iter * GraphQueryListResultIterator) Next() error {
        return iter.NextWithContext(context.Background())
        }
        // NotDone returns true if the enumeration should be started or is not yet complete.
        func (iter GraphQueryListResultIterator) NotDone() bool {
        return iter.page.NotDone() && iter.i < len(iter. page.Values())
        }
        // Response returns the raw server response from the last page request.
        func (iter GraphQueryListResultIterator) Response() GraphQueryListResult {
        return iter.page.Response()
        }
        // Value returns the current value or a zero-initialized value if the
        // iterator has advanced beyond the end of the collection.
        func (iter GraphQueryListResultIterator) Value() GraphQueryResource {
        if !iter.page.NotDone() {
        return GraphQueryResource{}
        }
        return iter.page.Values()[iter.i]
        }
        // Creates a new instance of the GraphQueryListResultIterator type.
        func NewGraphQueryListResultIterator (page GraphQueryListResultPage) GraphQueryListResultIterator {
            return GraphQueryListResultIterator{page: page}
        }


                // IsEmpty returns true if the ListResult contains no values.
                func (gqlr GraphQueryListResult) IsEmpty() bool {
                return gqlr.Value == nil || len(*gqlr.Value) == 0
                }

                    // graphQueryListResultPreparer prepares a request to retrieve the next set of results.
                    // It returns nil if no more results exist.
                    func (gqlr GraphQueryListResult) graphQueryListResultPreparer(ctx context.Context) (*http.Request, error) {
                    if gqlr.NextLink == nil || len(to.String(gqlr.NextLink)) < 1 {
                    return nil, nil
                    }
                    return autorest.Prepare((&http.Request{}).WithContext(ctx),
                    autorest.AsJSON(),
                    autorest.AsGet(),
                    autorest.WithBaseURL(to.String( gqlr.NextLink)));
                    }

            // GraphQueryListResultPage contains a page of GraphQueryResource values.
            type GraphQueryListResultPage struct {
                fn func(context.Context, GraphQueryListResult) (GraphQueryListResult, error)
                gqlr GraphQueryListResult
            }

        // NextWithContext advances to the next page of values.  If there was an error making
        // the request the page does not advance and the error is returned.
        func (page * GraphQueryListResultPage) NextWithContext(ctx context.Context) (err error) {
        if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/GraphQueryListResultPage.NextWithContext")
        defer func() {
        sc := -1
        if page.Response().Response.Response != nil {
        sc = page.Response().Response.Response.StatusCode
        }
        tracing.EndSpan(ctx, sc, err)
        }()
        }
        next, err := page.fn(ctx, page.gqlr)
        if err != nil {
        return err
        }
        page.gqlr = next
        return nil
        }

        // Next advances to the next page of values.  If there was an error making
        // the request the page does not advance and the error is returned.
        // Deprecated: Use NextWithContext() instead.
        func (page * GraphQueryListResultPage) Next() error {
        return page.NextWithContext(context.Background())
        }
        // NotDone returns true if the page enumeration should be started or is not yet complete.
        func (page GraphQueryListResultPage) NotDone() bool {
        return !page.gqlr.IsEmpty()
        }
        // Response returns the raw server response from the last page request.
        func (page GraphQueryListResultPage) Response() GraphQueryListResult {
        return page.gqlr
        }
        // Values returns the slice of values for the current page or nil if there are no values.
        func (page GraphQueryListResultPage) Values() []GraphQueryResource {
        if page.gqlr.IsEmpty() {
        return nil
        }
        return *page.gqlr.Value
        }
        // Creates a new instance of the GraphQueryListResultPage type.
        func NewGraphQueryListResultPage (getNextPage func(context.Context, GraphQueryListResult) (GraphQueryListResult, error)) GraphQueryListResultPage {
            return GraphQueryListResultPage{fn: getNextPage}
        }

            // GraphQueryProperties properties that contain a graph query.
            type GraphQueryProperties struct {
            // TimeModified - READ-ONLY; Date and time in UTC of the last modification that was made to this graph query definition.
            TimeModified *date.Time `json:"timeModified,omitempty"`
            // Description - The description of a graph query.
            Description *string `json:"description,omitempty"`
            // Query - KQL query that will be graph.
            Query *string `json:"query,omitempty"`
            // ResultKind - READ-ONLY; Enum indicating a type of graph query. Possible values include: 'Basic'
            ResultKind ResultKind `json:"resultKind,omitempty"`
            }

            // GraphQueryPropertiesUpdateParameters properties that contain a workbook for PATCH operation.
            type GraphQueryPropertiesUpdateParameters struct {
            // Description - The description of a graph query.
            Description *string `json:"description,omitempty"`
            // Query - KQL query that will be graph.
            Query *string `json:"query,omitempty"`
            }

            // GraphQueryResource graph Query entity definition.
            type GraphQueryResource struct {
            autorest.Response `json:"-"`
            // GraphQueryProperties - Metadata describing a graph query for an Azure resource.
            *GraphQueryProperties `json:"properties,omitempty"`
            // ID - READ-ONLY; Azure resource Id
            ID *string `json:"id,omitempty"`
            // Name - READ-ONLY; Azure resource name. This is GUID value. The display name should be assigned within properties field.
            Name *string `json:"name,omitempty"`
            // Type - READ-ONLY; Azure resource type
            Type *string `json:"type,omitempty"`
            // ETag - This will be used to handle Optimistic Concurrency. If not present, it will always overwrite the existing resource without checking conflict.
            ETag *string `json:"eTag,omitempty"`
            // Tags - Resource tags
            Tags map[string]*string `json:"tags"`
            }

        // MarshalJSON is the custom marshaler for GraphQueryResource.
        func (gqr GraphQueryResource)MarshalJSON() ([]byte, error){
        objectMap := make(map[string]interface{})
                if(gqr.GraphQueryProperties != nil) {
                objectMap["properties"] = gqr.GraphQueryProperties
                }
                if(gqr.ETag != nil) {
                objectMap["eTag"] = gqr.ETag
                }
                if(gqr.Tags != nil) {
                objectMap["tags"] = gqr.Tags
                }
                return json.Marshal(objectMap)
        }
        // UnmarshalJSON is the custom unmarshaler for GraphQueryResource struct.
        func (gqr *GraphQueryResource) UnmarshalJSON(body []byte) error {
        var m map[string]*json.RawMessage
        err := json.Unmarshal(body, &m)
        if err != nil {
        return err
        }
        for k, v := range  m {
        switch k {
                case "properties":
    if v != nil {
        var graphQueryProperties GraphQueryProperties
        err = json.Unmarshal(*v, &graphQueryProperties)
    if err != nil {
    return err
    }
        gqr.GraphQueryProperties = &graphQueryProperties
    }
                case "id":
    if v != nil {
        var ID string
        err = json.Unmarshal(*v, &ID)
    if err != nil {
    return err
    }
        gqr.ID = &ID
    }
                case "name":
    if v != nil {
        var name string
        err = json.Unmarshal(*v, &name)
    if err != nil {
    return err
    }
        gqr.Name = &name
    }
                case "type":
    if v != nil {
        var typeVar string
        err = json.Unmarshal(*v, &typeVar)
    if err != nil {
    return err
    }
        gqr.Type = &typeVar
    }
                case "eTag":
    if v != nil {
        var eTag string
        err = json.Unmarshal(*v, &eTag)
    if err != nil {
    return err
    }
        gqr.ETag = &eTag
    }
                case "tags":
    if v != nil {
        var tags map[string]*string
        err = json.Unmarshal(*v, &tags)
    if err != nil {
    return err
    }
        gqr.Tags = tags
    }
            }
        }

        return nil
        }

            // GraphQueryUpdateParameters the parameters that can be provided when updating workbook properties
            // properties.
            type GraphQueryUpdateParameters struct {
            // Tags - Resource tags
            Tags map[string]*string `json:"tags"`
            // ETag - This will be used to handle Optimistic Concurrency. If not present, it will always overwrite the existing resource without checking conflict.
            ETag *string `json:"eTag,omitempty"`
            // GraphQueryPropertiesUpdateParameters - Metadata describing a graph query for an Azure resource.
            *GraphQueryPropertiesUpdateParameters `json:"properties,omitempty"`
            }

        // MarshalJSON is the custom marshaler for GraphQueryUpdateParameters.
        func (gqup GraphQueryUpdateParameters)MarshalJSON() ([]byte, error){
        objectMap := make(map[string]interface{})
                if(gqup.Tags != nil) {
                objectMap["tags"] = gqup.Tags
                }
                if(gqup.ETag != nil) {
                objectMap["eTag"] = gqup.ETag
                }
                if(gqup.GraphQueryPropertiesUpdateParameters != nil) {
                objectMap["properties"] = gqup.GraphQueryPropertiesUpdateParameters
                }
                return json.Marshal(objectMap)
        }
        // UnmarshalJSON is the custom unmarshaler for GraphQueryUpdateParameters struct.
        func (gqup *GraphQueryUpdateParameters) UnmarshalJSON(body []byte) error {
        var m map[string]*json.RawMessage
        err := json.Unmarshal(body, &m)
        if err != nil {
        return err
        }
        for k, v := range  m {
        switch k {
                case "tags":
    if v != nil {
        var tags map[string]*string
        err = json.Unmarshal(*v, &tags)
    if err != nil {
    return err
    }
        gqup.Tags = tags
    }
                case "eTag":
    if v != nil {
        var eTag string
        err = json.Unmarshal(*v, &eTag)
    if err != nil {
    return err
    }
        gqup.ETag = &eTag
    }
                case "properties":
    if v != nil {
        var graphQueryPropertiesUpdateParameters GraphQueryPropertiesUpdateParameters
        err = json.Unmarshal(*v, &graphQueryPropertiesUpdateParameters)
    if err != nil {
    return err
    }
        gqup.GraphQueryPropertiesUpdateParameters = &graphQueryPropertiesUpdateParameters
    }
            }
        }

        return nil
        }

            // Operation resource Graph REST API operation definition.
            type Operation struct {
            // Name - Operation name: {provider}/{resource}/{operation}
            Name *string `json:"name,omitempty"`
            // Display - Display metadata associated with the operation.
            Display *OperationDisplay `json:"display,omitempty"`
            // Origin - The origin of operations.
            Origin *string `json:"origin,omitempty"`
            }

            // OperationDisplay display metadata associated with the operation.
            type OperationDisplay struct {
            // Provider - Service provider: Microsoft Resource Graph.
            Provider *string `json:"provider,omitempty"`
            // Resource - Resource on which the operation is performed etc.
            Resource *string `json:"resource,omitempty"`
            // Operation - Type of operation: get, read, delete, etc.
            Operation *string `json:"operation,omitempty"`
            // Description - Description for the operation.
            Description *string `json:"description,omitempty"`
            }

            // OperationListResult result of the request to list Resource Graph operations. It contains a list of
            // operations and a URL link to get the next set of results.
            type OperationListResult struct {
            autorest.Response `json:"-"`
            // Value - List of Resource Graph operations supported by the Resource Graph resource provider.
            Value *[]Operation `json:"value,omitempty"`
            }

            // QueryRequest describes a query to be executed.
            type QueryRequest struct {
            // Subscriptions - Azure subscriptions against which to execute the query.
            Subscriptions *[]string `json:"subscriptions,omitempty"`
            // Query - The resources query.
            Query *string `json:"query,omitempty"`
            // Options - The query evaluation options
            Options *QueryRequestOptions `json:"options,omitempty"`
            // Facets - An array of facet requests to be computed against the query result.
            Facets *[]FacetRequest `json:"facets,omitempty"`
            }

            // QueryRequestOptions the options for query evaluation
            type QueryRequestOptions struct {
            // SkipToken - Continuation token for pagination, capturing the next page size and offset, as well as the context of the query.
            SkipToken *string `json:"$skipToken,omitempty"`
            // Top - The maximum number of rows that the query should return. Overrides the page size when ```$skipToken``` property is present.
            Top *int32 `json:"$top,omitempty"`
            // Skip - The number of rows to skip from the beginning of the results. Overrides the next page offset when ```$skipToken``` property is present.
            Skip *int32 `json:"$skip,omitempty"`
            // ResultFormat - Defines in which format query result returned. Possible values include: 'ResultFormatTable', 'ResultFormatObjectArray'
            ResultFormat ResultFormat `json:"resultFormat,omitempty"`
            }

            // QueryResponse query result.
            type QueryResponse struct {
            autorest.Response `json:"-"`
            // TotalRecords - Number of total records matching the query.
            TotalRecords *int64 `json:"totalRecords,omitempty"`
            // Count - Number of records returned in the current response. In the case of paging, this is the number of records in the current page.
            Count *int64 `json:"count,omitempty"`
            // ResultTruncated - Indicates whether the query results are truncated. Possible values include: 'True', 'False'
            ResultTruncated ResultTruncated `json:"resultTruncated,omitempty"`
            // SkipToken - When present, the value can be passed to a subsequent query call (together with the same query and subscriptions used in the current request) to retrieve the next page of data.
            SkipToken *string `json:"$skipToken,omitempty"`
            // Data - Query output in tabular format.
            Data interface{} `json:"data,omitempty"`
            // Facets - Query facets.
            Facets *[]BasicFacet `json:"facets,omitempty"`
            }
        // UnmarshalJSON is the custom unmarshaler for QueryResponse struct.
        func (qr *QueryResponse) UnmarshalJSON(body []byte) error {
        var m map[string]*json.RawMessage
        err := json.Unmarshal(body, &m)
        if err != nil {
        return err
        }
        for k, v := range  m {
        switch k {
                case "totalRecords":
    if v != nil {
        var totalRecords int64
        err = json.Unmarshal(*v, &totalRecords)
    if err != nil {
    return err
    }
        qr.TotalRecords = &totalRecords
    }
                case "count":
    if v != nil {
        var count int64
        err = json.Unmarshal(*v, &count)
    if err != nil {
    return err
    }
        qr.Count = &count
    }
                case "resultTruncated":
    if v != nil {
        var resultTruncated ResultTruncated
        err = json.Unmarshal(*v, &resultTruncated)
    if err != nil {
    return err
    }
        qr.ResultTruncated = resultTruncated
    }
                case "$skipToken":
    if v != nil {
        var skipToken string
        err = json.Unmarshal(*v, &skipToken)
    if err != nil {
    return err
    }
        qr.SkipToken = &skipToken
    }
                case "data":
    if v != nil {
        var data interface{}
        err = json.Unmarshal(*v, &data)
    if err != nil {
    return err
    }
        qr.Data = data
    }
                case "facets":
    if v != nil {
        facets, err := unmarshalBasicFacetArray(*v)
    if err != nil {
    return err
    }
        qr.Facets = &facets
    }
            }
        }

        return nil
        }

            // Resource an azure resource object
            type Resource struct {
            // ID - READ-ONLY; Azure resource Id
            ID *string `json:"id,omitempty"`
            // Name - READ-ONLY; Azure resource name. This is GUID value. The display name should be assigned within properties field.
            Name *string `json:"name,omitempty"`
            // Type - READ-ONLY; Azure resource type
            Type *string `json:"type,omitempty"`
            // ETag - This will be used to handle Optimistic Concurrency. If not present, it will always overwrite the existing resource without checking conflict.
            ETag *string `json:"eTag,omitempty"`
            // Tags - Resource tags
            Tags map[string]*string `json:"tags"`
            }

        // MarshalJSON is the custom marshaler for Resource.
        func (r Resource)MarshalJSON() ([]byte, error){
        objectMap := make(map[string]interface{})
                if(r.ETag != nil) {
                objectMap["eTag"] = r.ETag
                }
                if(r.Tags != nil) {
                objectMap["tags"] = r.Tags
                }
                return json.Marshal(objectMap)
        }

            // Table query output in tabular format.
            type Table struct {
            // Columns - Query result column descriptors.
            Columns *[]Column `json:"columns,omitempty"`
            // Rows - Query result rows.
            Rows *[][]interface{} `json:"rows,omitempty"`
            }

