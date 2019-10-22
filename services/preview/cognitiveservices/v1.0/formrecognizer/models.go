package formrecognizer

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
    "encoding/json"
    "github.com/Azure/go-autorest/autorest"
    "github.com/Azure/go-autorest/autorest/date"
    "github.com/satori/go.uuid"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/preview/cognitiveservices/v1.0/formrecognizer"

        // Status enumerates the values for status.
    type Status string

    const (
                // Failure ...
        Failure Status = "failure"
                // PartialSuccess ...
        PartialSuccess Status = "partialSuccess"
                // Success ...
        Success Status = "success"
            )
    // PossibleStatusValues returns an array of possible values for the Status const type.
    func PossibleStatusValues() []Status {
        return []Status{Failure,PartialSuccess,Success}
    }

        // Status1 enumerates the values for status 1.
    type Status1 string

    const (
                // Created ...
        Created Status1 = "created"
                // Invalid ...
        Invalid Status1 = "invalid"
                // Ready ...
        Ready Status1 = "ready"
            )
    // PossibleStatus1Values returns an array of possible values for the Status1 const type.
    func PossibleStatus1Values() []Status1 {
        return []Status1{Created,Invalid,Ready}
    }

        // Status2 enumerates the values for status 2.
    type Status2 string

    const (
                // Status2Failure ...
        Status2Failure Status2 = "failure"
                // Status2PartialSuccess ...
        Status2PartialSuccess Status2 = "partialSuccess"
                // Status2Success ...
        Status2Success Status2 = "success"
            )
    // PossibleStatus2Values returns an array of possible values for the Status2 const type.
    func PossibleStatus2Values() []Status2 {
        return []Status2{Status2Failure,Status2PartialSuccess,Status2Success}
    }

        // TextOperationStatusCodes enumerates the values for text operation status codes.
    type TextOperationStatusCodes string

    const (
                // Failed ...
        Failed TextOperationStatusCodes = "Failed"
                // NotStarted ...
        NotStarted TextOperationStatusCodes = "Not Started"
                // Running ...
        Running TextOperationStatusCodes = "Running"
                // Succeeded ...
        Succeeded TextOperationStatusCodes = "Succeeded"
            )
    // PossibleTextOperationStatusCodesValues returns an array of possible values for the TextOperationStatusCodes const type.
    func PossibleTextOperationStatusCodesValues() []TextOperationStatusCodes {
        return []TextOperationStatusCodes{Failed,NotStarted,Running,Succeeded}
    }

        // TextRecognitionResultConfidenceClass enumerates the values for text recognition result confidence class.
    type TextRecognitionResultConfidenceClass string

    const (
                // High ...
        High TextRecognitionResultConfidenceClass = "High"
                // Low ...
        Low TextRecognitionResultConfidenceClass = "Low"
            )
    // PossibleTextRecognitionResultConfidenceClassValues returns an array of possible values for the TextRecognitionResultConfidenceClass const type.
    func PossibleTextRecognitionResultConfidenceClassValues() []TextRecognitionResultConfidenceClass {
        return []TextRecognitionResultConfidenceClass{High,Low}
    }

        // TextRecognitionResultDimensionUnit enumerates the values for text recognition result dimension unit.
    type TextRecognitionResultDimensionUnit string

    const (
                // Inch ...
        Inch TextRecognitionResultDimensionUnit = "inch"
                // Pixel ...
        Pixel TextRecognitionResultDimensionUnit = "pixel"
            )
    // PossibleTextRecognitionResultDimensionUnitValues returns an array of possible values for the TextRecognitionResultDimensionUnit const type.
    func PossibleTextRecognitionResultDimensionUnitValues() []TextRecognitionResultDimensionUnit {
        return []TextRecognitionResultDimensionUnit{Inch,Pixel}
    }

        // ValueType enumerates the values for value type.
    type ValueType string

    const (
                // ValueTypeFieldValue ...
        ValueTypeFieldValue ValueType = "fieldValue"
                // ValueTypeNumberValue ...
        ValueTypeNumberValue ValueType = "numberValue"
                // ValueTypeStringValue ...
        ValueTypeStringValue ValueType = "stringValue"
            )
    // PossibleValueTypeValues returns an array of possible values for the ValueType const type.
    func PossibleValueTypeValues() []ValueType {
        return []ValueType{ValueTypeFieldValue,ValueTypeNumberValue,ValueTypeStringValue}
    }

            // AnalyzeResult analyze API call result.
            type AnalyzeResult struct {
            autorest.Response `json:"-"`
            // Status - Status of the analyze operation. Possible values include: 'Status2Success', 'Status2PartialSuccess', 'Status2Failure'
            Status Status2 `json:"status,omitempty"`
            // Pages - Page level information extracted in the analyzed
            // document.
            Pages *[]ExtractedPage `json:"pages,omitempty"`
            // Errors - List of errors reported during the analyze
            // operation.
            Errors *[]FormOperationError `json:"errors,omitempty"`
            }

            // ComputerVisionError details about the API request error.
            type ComputerVisionError struct {
            // Code - The error code.
            Code interface{} `json:"code,omitempty"`
            // Message - A message explaining the error reported by the service.
            Message *string `json:"message,omitempty"`
            // RequestID - A unique request identifier.
            RequestID *string `json:"requestId,omitempty"`
            }

            // ElementReference reference to an OCR word.
            type ElementReference struct {
            Ref *string `json:"$ref,omitempty"`
            }

            // ErrorInformation ...
            type ErrorInformation struct {
            Code *string `json:"code,omitempty"`
            InnerError *InnerError `json:"innerError,omitempty"`
            Message *string `json:"message,omitempty"`
            }

            // ErrorResponse ...
            type ErrorResponse struct {
            Error *ErrorInformation `json:"error,omitempty"`
            }

            // ExtractedKeyValuePair representation of a key-value pair as a list
            // of key and value tokens.
            type ExtractedKeyValuePair struct {
            // Key - List of tokens for the extracted key in a key-value pair.
            Key *[]ExtractedToken `json:"key,omitempty"`
            // Value - List of tokens for the extracted value in a key-value pair.
            Value *[]ExtractedToken `json:"value,omitempty"`
            }

            // ExtractedPage extraction information of a single page in a
            // with a document.
            type ExtractedPage struct {
            // Number - Page number.
            Number *int32 `json:"number,omitempty"`
            // Height - Height of the page (in pixels).
            Height *int32 `json:"height,omitempty"`
            // Width - Width of the page (in pixels).
            Width *int32 `json:"width,omitempty"`
            // ClusterID - Cluster identifier.
            ClusterID *int32 `json:"clusterId,omitempty"`
            // KeyValuePairs - List of Key-Value pairs extracted from the page.
            KeyValuePairs *[]ExtractedKeyValuePair `json:"keyValuePairs,omitempty"`
            // Tables - List of Tables and their information extracted from the page.
            Tables *[]ExtractedTable `json:"tables,omitempty"`
            }

            // ExtractedTable extraction information about a table
            // contained in a page.
            type ExtractedTable struct {
            // ID - Table identifier.
            ID *string `json:"id,omitempty"`
            // Columns - List of columns contained in the table.
            Columns *[]ExtractedTableColumn `json:"columns,omitempty"`
            }

            // ExtractedTableColumn extraction information of a column in
            // a table.
            type ExtractedTableColumn struct {
            // Header - List of extracted tokens for the column header.
            Header *[]ExtractedToken `json:"header,omitempty"`
            // Entries - Extracted text for each cell of a column. Each cell
            // in the column can have a list of one or more tokens.
            Entries *[][]ExtractedToken `json:"entries,omitempty"`
            }

            // ExtractedToken canonical representation of single extracted text.
            type ExtractedToken struct {
            // Text - String value of the extracted text.
            Text *string `json:"text,omitempty"`
            // BoundingBox - Bounding box of the extracted text. Represents the
            // location of the extracted text as a pair of
            // cartesian co-ordinates. The co-ordinate pairs are arranged by
            // top-left, top-right, bottom-right and bottom-left endpoints box
            // with origin reference from the bottom-left of the page.
            BoundingBox *[]float64 `json:"boundingBox,omitempty"`
            // Confidence - A measure of accuracy of the extracted text.
            Confidence *float64 `json:"confidence,omitempty"`
            }

// BasicFieldValue base class representing a recognized field value.
        type BasicFieldValue interface {
            AsStringValue () (*StringValue, bool)
            AsNumberValue () (*NumberValue, bool)
        AsFieldValue () (*FieldValue, bool)
        }

        // FieldValue base class representing a recognized field value.
        type FieldValue struct {
        // Text - OCR text content of the recognized field.
        Text *string `json:"text,omitempty"`
        // Elements - List of references to OCR words comprising the recognized field value.
        Elements *[]ElementReference `json:"elements,omitempty"`
        // ValueType - Possible values include: 'ValueTypeFieldValue', 'ValueTypeStringValue', 'ValueTypeNumberValue'
        ValueType ValueType `json:"valueType,omitempty"`
        }

        func unmarshalBasicFieldValue(body []byte) (BasicFieldValue, error){
        var m map[string]interface{}
        err := json.Unmarshal(body, &m)
        if err != nil {
        return nil, err
        }

        switch m["valueType"] {
            case string(ValueTypeStringValue):
            var sv StringValue
            err := json.Unmarshal(body, &sv)
            return sv, err
            case string(ValueTypeNumberValue):
            var nv NumberValue
            err := json.Unmarshal(body, &nv)
            return nv, err
            default:
        var fv FieldValue
        err := json.Unmarshal(body, &fv)
        return fv, err
        }
        }
        func unmarshalBasicFieldValueArray(body []byte) ([]BasicFieldValue, error){
        var rawMessages []*json.RawMessage
        err := json.Unmarshal(body, &rawMessages)
        if err != nil {
        return nil, err
        }

        fvArray := make([]BasicFieldValue, len(rawMessages))

        for index, rawMessage := range rawMessages {
        fv, err := unmarshalBasicFieldValue(*rawMessage)
        if err != nil {
        return nil, err
        }
        fvArray[index] = fv
        }
        return fvArray, nil
        }

        // MarshalJSON is the custom marshaler for FieldValue.
        func (fv FieldValue)MarshalJSON() ([]byte, error){
            fv.ValueType = ValueTypeFieldValue
            objectMap := make(map[string]interface{})
                if(fv.Text != nil) {
                objectMap["text"] = fv.Text
                }
                if(fv.Elements != nil) {
                objectMap["elements"] = fv.Elements
                }
                if(fv.ValueType != "") {
                objectMap["valueType"] = fv.ValueType
                }
                return json.Marshal(objectMap)
        }

            // AsStringValue is the BasicFieldValue implementation for FieldValue.
            func (fv FieldValue) AsStringValue() (*StringValue, bool) {
                return nil, false
            }

            // AsNumberValue is the BasicFieldValue implementation for FieldValue.
            func (fv FieldValue) AsNumberValue() (*NumberValue, bool) {
                return nil, false
            }

            // AsFieldValue is the BasicFieldValue implementation for FieldValue.
            func (fv FieldValue) AsFieldValue() (*FieldValue, bool) {
                return &fv, true
            }

                // AsBasicFieldValue is the BasicFieldValue implementation for FieldValue.
                func(fv FieldValue) AsBasicFieldValue()(BasicFieldValue, bool) {
                    return &fv, true
                }


            // FormDocumentReport ...
            type FormDocumentReport struct {
            // DocumentName - Reference to the data that the report is for.
            DocumentName *string `json:"documentName,omitempty"`
            // Pages - Total number of pages trained on.
            Pages *int32 `json:"pages,omitempty"`
            // Errors - List of errors per page.
            Errors *[]string `json:"errors,omitempty"`
            // Status - Status of the training operation. Possible values include: 'Success', 'PartialSuccess', 'Failure'
            Status Status `json:"status,omitempty"`
            }

            // FormOperationError error reported during an operation.
            type FormOperationError struct {
            // ErrorMessage - Message reported during the train operation.
            ErrorMessage *string `json:"errorMessage,omitempty"`
            }

            // ImageURL ...
            type ImageURL struct {
            // URL - Publicly reachable URL of an image.
            URL *string `json:"url,omitempty"`
            }

            // InnerError ...
            type InnerError struct {
            RequestID *string `json:"requestId,omitempty"`
            }

            // KeysResult result of an operation to get
            // the keys extracted by a model.
            type KeysResult struct {
            autorest.Response `json:"-"`
            // Clusters - Object mapping ClusterIds to Key lists.
            Clusters map[string][]string `json:"clusters"`
            }

        // MarshalJSON is the custom marshaler for KeysResult.
        func (kr KeysResult)MarshalJSON() ([]byte, error){
        objectMap := make(map[string]interface{})
                if(kr.Clusters != nil) {
                objectMap["clusters"] = kr.Clusters
                }
                return json.Marshal(objectMap)
        }

            // Line an object representing a recognized text line.
            type Line struct {
            // BoundingBox - Bounding box of a recognized line.
            BoundingBox *[]int32 `json:"boundingBox,omitempty"`
            // Text - The text content of the line.
            Text *string `json:"text,omitempty"`
            // Words - List of words in the text line.
            Words *[]Word `json:"words,omitempty"`
            }

            // ModelResult result of a model status query operation.
            type ModelResult struct {
            autorest.Response `json:"-"`
            // ModelID - Get or set model identifier.
            ModelID *uuid.UUID `json:"modelId,omitempty"`
            // Status - Get or set the status of model. Possible values include: 'Created', 'Ready', 'Invalid'
            Status Status1 `json:"status,omitempty"`
            // CreatedDateTime - Get or set the created date time of the model.
            CreatedDateTime *date.Time `json:"createdDateTime,omitempty"`
            // LastUpdatedDateTime - Get or set the model last updated datetime.
            LastUpdatedDateTime *date.Time `json:"lastUpdatedDateTime,omitempty"`
            }

            // ModelsResult result of query operation to fetch multiple models.
            type ModelsResult struct {
            autorest.Response `json:"-"`
            // ModelsProperty - Collection of models.
            ModelsProperty *[]ModelResult `json:"models,omitempty"`
            }

            // NumberValue recognized numeric field value.
            type NumberValue struct {
            // Value - Numeric value of the recognized field.
            Value *float64 `json:"value,omitempty"`
            // Text - OCR text content of the recognized field.
            Text *string `json:"text,omitempty"`
            // Elements - List of references to OCR words comprising the recognized field value.
            Elements *[]ElementReference `json:"elements,omitempty"`
            // ValueType - Possible values include: 'ValueTypeFieldValue', 'ValueTypeStringValue', 'ValueTypeNumberValue'
            ValueType ValueType `json:"valueType,omitempty"`
            }

        // MarshalJSON is the custom marshaler for NumberValue.
        func (nv NumberValue)MarshalJSON() ([]byte, error){
            nv.ValueType = ValueTypeNumberValue
            objectMap := make(map[string]interface{})
                if(nv.Value != nil) {
                objectMap["value"] = nv.Value
                }
                if(nv.Text != nil) {
                objectMap["text"] = nv.Text
                }
                if(nv.Elements != nil) {
                objectMap["elements"] = nv.Elements
                }
                if(nv.ValueType != "") {
                objectMap["valueType"] = nv.ValueType
                }
                return json.Marshal(objectMap)
        }

            // AsStringValue is the BasicFieldValue implementation for NumberValue.
            func (nv NumberValue) AsStringValue() (*StringValue, bool) {
                return nil, false
            }

            // AsNumberValue is the BasicFieldValue implementation for NumberValue.
            func (nv NumberValue) AsNumberValue() (*NumberValue, bool) {
                return &nv, true
            }

            // AsFieldValue is the BasicFieldValue implementation for NumberValue.
            func (nv NumberValue) AsFieldValue() (*FieldValue, bool) {
                return nil, false
            }

                // AsBasicFieldValue is the BasicFieldValue implementation for NumberValue.
                func(nv NumberValue) AsBasicFieldValue()(BasicFieldValue, bool) {
                    return &nv, true
                }


            // ReadReceiptResult analysis result of the 'Batch Read Receipt' operation.
            type ReadReceiptResult struct {
            autorest.Response `json:"-"`
            // Status - Status of the read operation. Possible values include: 'NotStarted', 'Running', 'Failed', 'Succeeded'
            Status TextOperationStatusCodes `json:"status,omitempty"`
            // RecognitionResults - Text recognition result of the 'Batch Read Receipt' operation.
            RecognitionResults *[]TextRecognitionResult `json:"recognitionResults,omitempty"`
            // UnderstandingResults - Semantic understanding result of the 'Batch Read Receipt' operation.
            UnderstandingResults *[]UnderstandingResult `json:"understandingResults,omitempty"`
            }

            // StringValue recognized string field value.
            type StringValue struct {
            // Value - String value of the recognized field.
            Value *string `json:"value,omitempty"`
            // Text - OCR text content of the recognized field.
            Text *string `json:"text,omitempty"`
            // Elements - List of references to OCR words comprising the recognized field value.
            Elements *[]ElementReference `json:"elements,omitempty"`
            // ValueType - Possible values include: 'ValueTypeFieldValue', 'ValueTypeStringValue', 'ValueTypeNumberValue'
            ValueType ValueType `json:"valueType,omitempty"`
            }

        // MarshalJSON is the custom marshaler for StringValue.
        func (sv StringValue)MarshalJSON() ([]byte, error){
            sv.ValueType = ValueTypeStringValue
            objectMap := make(map[string]interface{})
                if(sv.Value != nil) {
                objectMap["value"] = sv.Value
                }
                if(sv.Text != nil) {
                objectMap["text"] = sv.Text
                }
                if(sv.Elements != nil) {
                objectMap["elements"] = sv.Elements
                }
                if(sv.ValueType != "") {
                objectMap["valueType"] = sv.ValueType
                }
                return json.Marshal(objectMap)
        }

            // AsStringValue is the BasicFieldValue implementation for StringValue.
            func (sv StringValue) AsStringValue() (*StringValue, bool) {
                return &sv, true
            }

            // AsNumberValue is the BasicFieldValue implementation for StringValue.
            func (sv StringValue) AsNumberValue() (*NumberValue, bool) {
                return nil, false
            }

            // AsFieldValue is the BasicFieldValue implementation for StringValue.
            func (sv StringValue) AsFieldValue() (*FieldValue, bool) {
                return nil, false
            }

                // AsBasicFieldValue is the BasicFieldValue implementation for StringValue.
                func(sv StringValue) AsBasicFieldValue()(BasicFieldValue, bool) {
                    return &sv, true
                }


            // TextRecognitionResult an object representing a recognized text region
            type TextRecognitionResult struct {
            // Page - The 1-based page number of the recognition result.
            Page *int32 `json:"page,omitempty"`
            // ClockwiseOrientation - The orientation of the image in degrees in the clockwise direction. Range between [0, 360).
            ClockwiseOrientation *float64 `json:"clockwiseOrientation,omitempty"`
            // Width - The width of the image in pixels or the PDF in inches.
            Width *float64 `json:"width,omitempty"`
            // Height - The height of the image in pixels or the PDF in inches.
            Height *float64 `json:"height,omitempty"`
            // Unit - The unit used in the Width, Height and BoundingBox. For images, the unit is 'pixel'. For PDF, the unit is 'inch'. Possible values include: 'Pixel', 'Inch'
            Unit TextRecognitionResultDimensionUnit `json:"unit,omitempty"`
            // Lines - A list of recognized text lines.
            Lines *[]Line `json:"lines,omitempty"`
            }

            // TrainRequest contract to initiate a train request.
            type TrainRequest struct {
            // Source - Get or set source path.
            Source *string `json:"source,omitempty"`
            // SourceFilter - Get or set filter to further search the
            // source path for content.
            SourceFilter *TrainSourceFilter `json:"sourceFilter,omitempty"`
            }

            // TrainResult response of the Train API call.
            type TrainResult struct {
            autorest.Response `json:"-"`
            // ModelID - Identifier of the model.
            ModelID *uuid.UUID `json:"modelId,omitempty"`
            // TrainingDocuments - List of documents used to train the model and the
            // train operation error reported by each.
            TrainingDocuments *[]FormDocumentReport `json:"trainingDocuments,omitempty"`
            // Errors - Errors returned during the training operation.
            Errors *[]FormOperationError `json:"errors,omitempty"`
            }

            // TrainSourceFilter filters to be applied when traversing a data source.
            type TrainSourceFilter struct {
            // Prefix - A case-sensitive prefix string to filter content
            // under the source location. For e.g., when using a Azure Blob
            // Uri use the prefix to restrict subfolders for content.
            Prefix *string `json:"prefix,omitempty"`
            // IncludeSubFolders - A flag to indicate if sub folders within the set of
            // prefix folders will also need to be included when searching
            // for content to be preprocessed.
            IncludeSubFolders *bool `json:"includeSubFolders,omitempty"`
            }

            // UnderstandingResult a set of extracted fields corresponding to a semantic object, such as a receipt, in
            // the input document.
            type UnderstandingResult struct {
            // Pages - List of pages where the document is found.
            Pages *[]int32 `json:"pages,omitempty"`
            // Fields - Dictionary of recognized field values.
            Fields map[string]*FieldValue `json:"fields"`
            }

        // MarshalJSON is the custom marshaler for UnderstandingResult.
        func (ur UnderstandingResult)MarshalJSON() ([]byte, error){
        objectMap := make(map[string]interface{})
                if(ur.Pages != nil) {
                objectMap["pages"] = ur.Pages
                }
                if(ur.Fields != nil) {
                objectMap["fields"] = ur.Fields
                }
                return json.Marshal(objectMap)
        }

            // Word an object representing a recognized word.
            type Word struct {
            // BoundingBox - Bounding box of a recognized word.
            BoundingBox *[]int32 `json:"boundingBox,omitempty"`
            // Text - The text content of the word.
            Text *string `json:"text,omitempty"`
            // Confidence - Qualitative confidence measure. Possible values include: 'High', 'Low'
            Confidence TextRecognitionResultConfidenceClass `json:"confidence,omitempty"`
            }

