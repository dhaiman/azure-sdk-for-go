package qnamaker

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
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v4.0/qnamaker"

// ActiveLearningSettingsDTO active Learning settings of the endpoint.
type ActiveLearningSettingsDTO struct {
	// Enable - True/False string providing Active Learning
	Enable *string `json:"enable,omitempty"`
}

// AlterationsDTO collection of words that are synonyms.
type AlterationsDTO struct {
	// Alterations - Words that are synonymous with each other.
	Alterations *[]string `json:"alterations,omitempty"`
}

// ContextDTO context associated with Qna.
type ContextDTO struct {
	// IsContextOnly - To mark if a prompt is relevant only with a previous question or not.
	// true - Do not include this QnA as search result for queries without context
	// false - ignores context and includes this QnA in search result
	IsContextOnly *bool `json:"isContextOnly,omitempty"`
	// Prompts - List of prompts associated with the answer.
	Prompts *[]PromptDTO `json:"prompts,omitempty"`
}

// CreateKbDTO post body schema for CreateKb operation.
type CreateKbDTO struct {
	// Name - Friendly name for the knowledgebase.
	Name *string `json:"name,omitempty"`
	// QnaList - List of Q-A (QnADTO) to be added to the knowledgebase. Q-A Ids are assigned by the service and should be omitted.
	QnaList *[]QnADTO `json:"qnaList,omitempty"`
	// Urls - List of URLs to be used for extracting Q-A.
	Urls *[]string `json:"urls,omitempty"`
	// Files - List of files from which to Extract Q-A.
	Files *[]FileDTO `json:"files,omitempty"`
	// EnableHierarchicalExtraction - Enable hierarchical extraction of Q-A from files and urls. Value to be considered False if this field is not present.
	EnableHierarchicalExtraction *bool `json:"enableHierarchicalExtraction,omitempty"`
	// DefaultAnswerUsedForExtraction - Text string to be used as the answer in any Q-A which has no extracted answer from the document but has a hierarchy. Required when EnableHierarchicalExtraction field is set to True.
	DefaultAnswerUsedForExtraction *string `json:"defaultAnswerUsedForExtraction,omitempty"`
	// Language - Language of the knowledgebase. Please find the list of supported languages <a href="https://aka.ms/qnamaker-languages#languages-supported" target="_blank">here</a>.
	Language *string `json:"language,omitempty"`
}

// CreateKbInputDTO input to create KB.
type CreateKbInputDTO struct {
	// QnaList - List of QNA to be added to the index. Ids are generated by the service and should be omitted.
	QnaList *[]QnADTO `json:"qnaList,omitempty"`
	// Urls - List of URLs to be added to knowledgebase.
	Urls *[]string `json:"urls,omitempty"`
	// Files - List of files to be added to knowledgebase.
	Files *[]FileDTO `json:"files,omitempty"`
}

// DeleteKbContentsDTO PATCH body schema of Delete Operation in UpdateKb
type DeleteKbContentsDTO struct {
	// Ids - List of Qna Ids to be deleted
	Ids *[]int32 `json:"ids,omitempty"`
	// Sources - List of sources to be deleted from knowledgebase.
	Sources *[]string `json:"sources,omitempty"`
}

// EndpointKeysDTO schema for EndpointKeys generate/refresh operations.
type EndpointKeysDTO struct {
	autorest.Response `json:"-"`
	// PrimaryEndpointKey - Primary Access Key.
	PrimaryEndpointKey *string `json:"primaryEndpointKey,omitempty"`
	// SecondaryEndpointKey - Secondary Access Key.
	SecondaryEndpointKey *string `json:"secondaryEndpointKey,omitempty"`
	// InstalledVersion - Current version of runtime.
	InstalledVersion *string `json:"installedVersion,omitempty"`
	// LastStableVersion - Latest version of runtime.
	LastStableVersion *string `json:"lastStableVersion,omitempty"`
	// Language - Language setting of runtime.
	Language *string `json:"language,omitempty"`
}

// EndpointSettingsDTO endpoint settings.
type EndpointSettingsDTO struct {
	autorest.Response `json:"-"`
	// ActiveLearning - Active Learning settings of the endpoint.
	ActiveLearning *EndpointSettingsDTOActiveLearning `json:"activeLearning,omitempty"`
}

// EndpointSettingsDTOActiveLearning active Learning settings of the endpoint.
type EndpointSettingsDTOActiveLearning struct {
	// Enable - True/False string providing Active Learning
	Enable *string `json:"enable,omitempty"`
}

// Error the error object. As per Microsoft One API guidelines -
// https://github.com/Microsoft/api-guidelines/blob/vNext/Guidelines.md#7102-error-condition-responses.
type Error struct {
	// Code - One of a server-defined set of error codes. Possible values include: 'BadArgument', 'Forbidden', 'NotFound', 'KbNotFound', 'Unauthorized', 'Unspecified', 'EndpointKeysError', 'QuotaExceeded', 'QnaRuntimeError', 'SKULimitExceeded', 'OperationNotFound', 'ServiceError', 'ValidationFailure', 'ExtractionFailure'
	Code ErrorCodeType `json:"code,omitempty"`
	// Message - A human-readable representation of the error.
	Message *string `json:"message,omitempty"`
	// Target - The target of the error.
	Target *string `json:"target,omitempty"`
	// Details - An array of details about specific errors that led to this reported error.
	Details *[]Error `json:"details,omitempty"`
	// InnerError - An object containing more specific information than the current object about the error.
	InnerError *InnerErrorModel `json:"innerError,omitempty"`
}

// ErrorResponse error response. As per Microsoft One API guidelines -
// https://github.com/Microsoft/api-guidelines/blob/vNext/Guidelines.md#7102-error-condition-responses.
type ErrorResponse struct {
	// Error - The error object.
	Error *ErrorResponseError `json:"error,omitempty"`
}

// ErrorResponseError the error object.
type ErrorResponseError struct {
	// Code - One of a server-defined set of error codes. Possible values include: 'BadArgument', 'Forbidden', 'NotFound', 'KbNotFound', 'Unauthorized', 'Unspecified', 'EndpointKeysError', 'QuotaExceeded', 'QnaRuntimeError', 'SKULimitExceeded', 'OperationNotFound', 'ServiceError', 'ValidationFailure', 'ExtractionFailure'
	Code ErrorCodeType `json:"code,omitempty"`
	// Message - A human-readable representation of the error.
	Message *string `json:"message,omitempty"`
	// Target - The target of the error.
	Target *string `json:"target,omitempty"`
	// Details - An array of details about specific errors that led to this reported error.
	Details *[]Error `json:"details,omitempty"`
	// InnerError - An object containing more specific information than the current object about the error.
	InnerError *InnerErrorModel `json:"innerError,omitempty"`
}

// FileDTO DTO to hold details of uploaded files.
type FileDTO struct {
	// FileName - File name. Supported file types are ".tsv", ".pdf", ".txt", ".docx", ".xlsx".
	FileName *string `json:"fileName,omitempty"`
	// FileURI - Public URI of the file.
	FileURI *string `json:"fileUri,omitempty"`
}

// InnerErrorModel an object containing more specific information about the error. As per Microsoft One API
// guidelines -
// https://github.com/Microsoft/api-guidelines/blob/vNext/Guidelines.md#7102-error-condition-responses.
type InnerErrorModel struct {
	// Code - A more specific error code than was provided by the containing error.
	Code *string `json:"code,omitempty"`
	// InnerError - An object containing more specific information than the current object about the error.
	InnerError *InnerErrorModel `json:"innerError,omitempty"`
}

// KnowledgebaseDTO response schema for CreateKb operation.
type KnowledgebaseDTO struct {
	autorest.Response `json:"-"`
	// ID - Unique id that identifies a knowledgebase.
	ID *string `json:"id,omitempty"`
	// HostName - URL host name at which the knowledgebase is hosted.
	HostName *string `json:"hostName,omitempty"`
	// LastAccessedTimestamp - Time stamp at which the knowledgebase was last accessed (UTC).
	LastAccessedTimestamp *string `json:"lastAccessedTimestamp,omitempty"`
	// LastChangedTimestamp - Time stamp at which the knowledgebase was last modified (UTC).
	LastChangedTimestamp *string `json:"lastChangedTimestamp,omitempty"`
	// LastPublishedTimestamp - Time stamp at which the knowledgebase was last published (UTC).
	LastPublishedTimestamp *string `json:"lastPublishedTimestamp,omitempty"`
	// Name - Friendly name of the knowledgebase.
	Name *string `json:"name,omitempty"`
	// UserID - User who created / owns the knowledgebase.
	UserID *string `json:"userId,omitempty"`
	// Urls - URL sources from which Q-A were extracted and added to the knowledgebase.
	Urls *[]string `json:"urls,omitempty"`
	// Sources - Custom sources from which Q-A were extracted or explicitly added to the knowledgebase.
	Sources *[]string `json:"sources,omitempty"`
}

// KnowledgebasesDTO collection of knowledgebases owned by a user.
type KnowledgebasesDTO struct {
	autorest.Response `json:"-"`
	// Knowledgebases - Collection of knowledgebase records.
	Knowledgebases *[]KnowledgebaseDTO `json:"knowledgebases,omitempty"`
}

// MetadataDTO name - value pair of metadata.
type MetadataDTO struct {
	// Name - Metadata name.
	Name *string `json:"name,omitempty"`
	// Value - Metadata value.
	Value *string `json:"value,omitempty"`
}

// Operation record to track long running operation.
type Operation struct {
	autorest.Response `json:"-"`
	// OperationState - Operation state. Possible values include: 'Failed', 'NotStarted', 'Running', 'Succeeded'
	OperationState OperationStateType `json:"operationState,omitempty"`
	// CreatedTimestamp - Timestamp when the operation was created.
	CreatedTimestamp *string `json:"createdTimestamp,omitempty"`
	// LastActionTimestamp - Timestamp when the current state was entered.
	LastActionTimestamp *string `json:"lastActionTimestamp,omitempty"`
	// ResourceLocation - Relative URI to the target resource location for completed resources.
	ResourceLocation *string `json:"resourceLocation,omitempty"`
	// UserID - User Id
	UserID *string `json:"userId,omitempty"`
	// OperationID - Operation Id.
	OperationID *string `json:"operationId,omitempty"`
	// ErrorResponse - Error details in case of failures.
	ErrorResponse *ErrorResponse `json:"errorResponse,omitempty"`
}

// PromptDTO prompt for an answer.
type PromptDTO struct {
	// DisplayOrder - Index of the prompt - used in ordering of the prompts
	DisplayOrder *int32 `json:"displayOrder,omitempty"`
	// QnaID - Qna id corresponding to the prompt - if QnaId is present, QnADTO object is ignored.
	QnaID *int32 `json:"qnaId,omitempty"`
	// Qna - QnADTO - Either QnaId or QnADTO needs to be present in a PromptDTO object
	Qna *PromptDTOQna `json:"qna,omitempty"`
	// DisplayText - Text displayed to represent a follow up question prompt
	DisplayText *string `json:"displayText,omitempty"`
}

// PromptDTOQna qnADTO - Either QnaId or QnADTO needs to be present in a PromptDTO object
type PromptDTOQna struct {
	// ID - Unique id for the Q-A.
	ID *int32 `json:"id,omitempty"`
	// Answer - Answer text
	Answer *string `json:"answer,omitempty"`
	// Source - Source from which Q-A was indexed. eg. https://docs.microsoft.com/en-us/azure/cognitive-services/QnAMaker/FAQs
	Source *string `json:"source,omitempty"`
	// Questions - List of questions associated with the answer.
	Questions *[]string `json:"questions,omitempty"`
	// Metadata - List of metadata associated with the answer.
	Metadata *[]MetadataDTO `json:"metadata,omitempty"`
	// Context - Context of a QnA
	Context *QnADTOContext `json:"context,omitempty"`
}

// QnADocumentsDTO list of QnADTO
type QnADocumentsDTO struct {
	autorest.Response `json:"-"`
	// QnaDocuments - List of answers.
	QnaDocuments *[]QnADTO `json:"qnaDocuments,omitempty"`
}

// QnADTO q-A object.
type QnADTO struct {
	// ID - Unique id for the Q-A.
	ID *int32 `json:"id,omitempty"`
	// Answer - Answer text
	Answer *string `json:"answer,omitempty"`
	// Source - Source from which Q-A was indexed. eg. https://docs.microsoft.com/en-us/azure/cognitive-services/QnAMaker/FAQs
	Source *string `json:"source,omitempty"`
	// Questions - List of questions associated with the answer.
	Questions *[]string `json:"questions,omitempty"`
	// Metadata - List of metadata associated with the answer.
	Metadata *[]MetadataDTO `json:"metadata,omitempty"`
	// Context - Context of a QnA
	Context *QnADTOContext `json:"context,omitempty"`
}

// QnADTOContext context of a QnA
type QnADTOContext struct {
	// IsContextOnly - To mark if a prompt is relevant only with a previous question or not.
	// true - Do not include this QnA as search result for queries without context
	// false - ignores context and includes this QnA in search result
	IsContextOnly *bool `json:"isContextOnly,omitempty"`
	// Prompts - List of prompts associated with the answer.
	Prompts *[]PromptDTO `json:"prompts,omitempty"`
}

// ReplaceKbDTO post body schema for Replace KB operation.
type ReplaceKbDTO struct {
	// QnAList - List of Q-A (QnADTO) to be added to the knowledgebase. Q-A Ids are assigned by the service and should be omitted.
	QnAList *[]QnADTO `json:"qnAList,omitempty"`
}

// UpdateContextDTO update Body schema to represent context to be updated
type UpdateContextDTO struct {
	// PromptsToDelete - List of prompts associated with qna to be deleted
	PromptsToDelete *[]int32 `json:"promptsToDelete,omitempty"`
	// PromptsToAdd - List of prompts to be added to the qna.
	PromptsToAdd *[]PromptDTO `json:"promptsToAdd,omitempty"`
	// IsContextOnly - To mark if a prompt is relevant only with a previous question or not.
	// true - Do not include this QnA as search result for queries without context
	// false - ignores context and includes this QnA in search result
	IsContextOnly *bool `json:"isContextOnly,omitempty"`
}

// UpdateKbContentsDTO PATCH body schema for Update operation in Update Kb
type UpdateKbContentsDTO struct {
	// Name - Friendly name for the knowledgebase.
	Name *string `json:"name,omitempty"`
	// QnaList - List of Q-A (UpdateQnaDTO) to be added to the knowledgebase.
	QnaList *[]UpdateQnaDTO `json:"qnaList,omitempty"`
	// Urls - List of existing URLs to be refreshed. The content will be extracted again and re-indexed.
	Urls *[]string `json:"urls,omitempty"`
}

// UpdateKbOperationDTO contains list of QnAs to be updated
type UpdateKbOperationDTO struct {
	// Add - An instance of CreateKbInputDTO for add operation
	Add *UpdateKbOperationDTOAdd `json:"add,omitempty"`
	// Delete - An instance of DeleteKbContentsDTO for delete Operation
	Delete *UpdateKbOperationDTODelete `json:"delete,omitempty"`
	// Update - An instance of UpdateKbContentsDTO for Update Operation
	Update *UpdateKbOperationDTOUpdate `json:"update,omitempty"`
	// EnableHierarchicalExtraction - Enable hierarchical extraction of Q-A from files and urls. The value set during KB creation will be used if this field is not present.
	EnableHierarchicalExtraction *bool `json:"enableHierarchicalExtraction,omitempty"`
	// DefaultAnswerUsedForExtraction - Text string to be used as the answer in any Q-A which has no extracted answer from the document but has a hierarchy. Required when EnableHierarchicalExtraction field is set to True.
	DefaultAnswerUsedForExtraction *string `json:"defaultAnswerUsedForExtraction,omitempty"`
}

// UpdateKbOperationDTOAdd an instance of CreateKbInputDTO for add operation
type UpdateKbOperationDTOAdd struct {
	// QnaList - List of QNA to be added to the index. Ids are generated by the service and should be omitted.
	QnaList *[]QnADTO `json:"qnaList,omitempty"`
	// Urls - List of URLs to be added to knowledgebase.
	Urls *[]string `json:"urls,omitempty"`
	// Files - List of files to be added to knowledgebase.
	Files *[]FileDTO `json:"files,omitempty"`
}

// UpdateKbOperationDTODelete an instance of DeleteKbContentsDTO for delete Operation
type UpdateKbOperationDTODelete struct {
	// Ids - List of Qna Ids to be deleted
	Ids *[]int32 `json:"ids,omitempty"`
	// Sources - List of sources to be deleted from knowledgebase.
	Sources *[]string `json:"sources,omitempty"`
}

// UpdateKbOperationDTOUpdate an instance of UpdateKbContentsDTO for Update Operation
type UpdateKbOperationDTOUpdate struct {
	// Name - Friendly name for the knowledgebase.
	Name *string `json:"name,omitempty"`
	// QnaList - List of Q-A (UpdateQnaDTO) to be added to the knowledgebase.
	QnaList *[]UpdateQnaDTO `json:"qnaList,omitempty"`
	// Urls - List of existing URLs to be refreshed. The content will be extracted again and re-indexed.
	Urls *[]string `json:"urls,omitempty"`
}

// UpdateMetadataDTO PATCH Body schema to represent list of Metadata to be updated
type UpdateMetadataDTO struct {
	// Delete - List of Metadata associated with answer to be deleted
	Delete *[]MetadataDTO `json:"delete,omitempty"`
	// Add - List of metadata associated with answer to be added
	Add *[]MetadataDTO `json:"add,omitempty"`
}

// UpdateQnaDTO PATCH Body schema for Update Qna List
type UpdateQnaDTO struct {
	// ID - Unique id for the Q-A
	ID *int32 `json:"id,omitempty"`
	// Answer - Answer text
	Answer *string `json:"answer,omitempty"`
	// Source - Source from which Q-A was indexed. eg. https://docs.microsoft.com/en-us/azure/cognitive-services/QnAMaker/FAQs
	Source *string `json:"source,omitempty"`
	// Questions - List of questions associated with the answer.
	Questions *UpdateQnaDTOQuestions `json:"questions,omitempty"`
	// Metadata - List of metadata associated with the answer to be updated
	Metadata *UpdateQnaDTOMetadata `json:"metadata,omitempty"`
	// Context - Context associated with Qna to be updated.
	Context *UpdateQnaDTOContext `json:"context,omitempty"`
}

// UpdateQnaDTOContext context associated with Qna to be updated.
type UpdateQnaDTOContext struct {
	// PromptsToDelete - List of prompts associated with qna to be deleted
	PromptsToDelete *[]int32 `json:"promptsToDelete,omitempty"`
	// PromptsToAdd - List of prompts to be added to the qna.
	PromptsToAdd *[]PromptDTO `json:"promptsToAdd,omitempty"`
	// IsContextOnly - To mark if a prompt is relevant only with a previous question or not.
	// true - Do not include this QnA as search result for queries without context
	// false - ignores context and includes this QnA in search result
	IsContextOnly *bool `json:"isContextOnly,omitempty"`
}

// UpdateQnaDTOMetadata list of metadata associated with the answer to be updated
type UpdateQnaDTOMetadata struct {
	// Delete - List of Metadata associated with answer to be deleted
	Delete *[]MetadataDTO `json:"delete,omitempty"`
	// Add - List of metadata associated with answer to be added
	Add *[]MetadataDTO `json:"add,omitempty"`
}

// UpdateQnaDTOQuestions list of questions associated with the answer.
type UpdateQnaDTOQuestions struct {
	// Add - List of questions to be added
	Add *[]string `json:"add,omitempty"`
	// Delete - List of questions to be deleted.
	Delete *[]string `json:"delete,omitempty"`
}

// UpdateQuestionsDTO PATCH Body schema for Update Kb which contains list of questions to be added and deleted
type UpdateQuestionsDTO struct {
	// Add - List of questions to be added
	Add *[]string `json:"add,omitempty"`
	// Delete - List of questions to be deleted.
	Delete *[]string `json:"delete,omitempty"`
}

// WordAlterationsDTO collection of word alterations.
type WordAlterationsDTO struct {
	autorest.Response `json:"-"`
	// WordAlterations - Collection of word alterations.
	WordAlterations *[]AlterationsDTO `json:"wordAlterations,omitempty"`
}
