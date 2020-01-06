package faceapi

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
	"github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v1.0/face"
	"github.com/Azure/go-autorest/autorest"
	"github.com/satori/go.uuid"
	"io"
)

// ClientAPI contains the set of methods on the Client type.
type ClientAPI interface {
	DetectWithStream(ctx context.Context, imageParameter io.ReadCloser, returnFaceID *bool, returnFaceLandmarks *bool, returnFaceAttributes []face.AttributeType, recognitionModel face.RecognitionModel, returnRecognitionModel *bool, detectionModel face.DetectionModel) (result face.ListDetectedFace, err error)
	DetectWithURL(ctx context.Context, imageURL face.ImageURL, returnFaceID *bool, returnFaceLandmarks *bool, returnFaceAttributes []face.AttributeType, recognitionModel face.RecognitionModel, returnRecognitionModel *bool, detectionModel face.DetectionModel) (result face.ListDetectedFace, err error)
	FindSimilar(ctx context.Context, body face.FindSimilarRequest) (result face.ListSimilarFace, err error)
	Group(ctx context.Context, body face.GroupRequest) (result face.GroupResult, err error)
	Identify(ctx context.Context, body face.IdentifyRequest) (result face.ListIdentifyResult, err error)
	VerifyFaceToFace(ctx context.Context, body face.VerifyFaceToFaceRequest) (result face.VerifyResult, err error)
	VerifyFaceToPerson(ctx context.Context, body face.VerifyFaceToPersonRequest) (result face.VerifyResult, err error)
}

var _ ClientAPI = (*face.Client)(nil)

// PersonGroupPersonClientAPI contains the set of methods on the PersonGroupPersonClient type.
type PersonGroupPersonClientAPI interface {
	AddFaceFromStream(ctx context.Context, personGroupID string, personID uuid.UUID, imageParameter io.ReadCloser, userData string, targetFace []int32, detectionModel face.DetectionModel) (result face.PersistedFace, err error)
	AddFaceFromURL(ctx context.Context, personGroupID string, personID uuid.UUID, imageURL face.ImageURL, userData string, targetFace []int32, detectionModel face.DetectionModel) (result face.PersistedFace, err error)
	Create(ctx context.Context, personGroupID string, body face.NameAndUserDataContract) (result face.Person, err error)
	Delete(ctx context.Context, personGroupID string, personID uuid.UUID) (result autorest.Response, err error)
	DeleteFace(ctx context.Context, personGroupID string, personID uuid.UUID, persistedFaceID uuid.UUID) (result autorest.Response, err error)
	Get(ctx context.Context, personGroupID string, personID uuid.UUID) (result face.Person, err error)
	GetFace(ctx context.Context, personGroupID string, personID uuid.UUID, persistedFaceID uuid.UUID) (result face.PersistedFace, err error)
	List(ctx context.Context, personGroupID string, start string, top *int32) (result face.ListPerson, err error)
	Update(ctx context.Context, personGroupID string, personID uuid.UUID, body face.NameAndUserDataContract) (result autorest.Response, err error)
	UpdateFace(ctx context.Context, personGroupID string, personID uuid.UUID, persistedFaceID uuid.UUID, body face.UpdateFaceRequest) (result autorest.Response, err error)
}

var _ PersonGroupPersonClientAPI = (*face.PersonGroupPersonClient)(nil)

// PersonGroupClientAPI contains the set of methods on the PersonGroupClient type.
type PersonGroupClientAPI interface {
	Create(ctx context.Context, personGroupID string, body face.MetaDataContract) (result autorest.Response, err error)
	Delete(ctx context.Context, personGroupID string) (result autorest.Response, err error)
	Get(ctx context.Context, personGroupID string, returnRecognitionModel *bool) (result face.PersonGroup, err error)
	GetTrainingStatus(ctx context.Context, personGroupID string) (result face.TrainingStatus, err error)
	List(ctx context.Context, start string, top *int32, returnRecognitionModel *bool) (result face.ListPersonGroup, err error)
	Train(ctx context.Context, personGroupID string) (result autorest.Response, err error)
	Update(ctx context.Context, personGroupID string, body face.NameAndUserDataContract) (result autorest.Response, err error)
}

var _ PersonGroupClientAPI = (*face.PersonGroupClient)(nil)

// ListClientAPI contains the set of methods on the ListClient type.
type ListClientAPI interface {
	AddFaceFromStream(ctx context.Context, faceListID string, imageParameter io.ReadCloser, userData string, targetFace []int32, detectionModel face.DetectionModel) (result face.PersistedFace, err error)
	AddFaceFromURL(ctx context.Context, faceListID string, imageURL face.ImageURL, userData string, targetFace []int32, detectionModel face.DetectionModel) (result face.PersistedFace, err error)
	Create(ctx context.Context, faceListID string, body face.MetaDataContractMandatoryName) (result autorest.Response, err error)
	Delete(ctx context.Context, faceListID string) (result autorest.Response, err error)
	DeleteFace(ctx context.Context, faceListID string, persistedFaceID uuid.UUID) (result autorest.Response, err error)
	Get(ctx context.Context, faceListID string, returnRecognitionModel *bool) (result face.List, err error)
	List(ctx context.Context, returnRecognitionModel *bool) (result face.ListList, err error)
	Update(ctx context.Context, faceListID string, body face.NameAndUserDataContract) (result autorest.Response, err error)
}

var _ ListClientAPI = (*face.ListClient)(nil)

// LargePersonGroupPersonClientAPI contains the set of methods on the LargePersonGroupPersonClient type.
type LargePersonGroupPersonClientAPI interface {
	AddFaceFromStream(ctx context.Context, largePersonGroupID string, personID uuid.UUID, imageParameter io.ReadCloser, userData string, targetFace []int32, detectionModel face.DetectionModel) (result face.PersistedFace, err error)
	AddFaceFromURL(ctx context.Context, largePersonGroupID string, personID uuid.UUID, imageURL face.ImageURL, userData string, targetFace []int32, detectionModel face.DetectionModel) (result face.PersistedFace, err error)
	Create(ctx context.Context, largePersonGroupID string, body face.NameAndUserDataContract) (result face.Person, err error)
	Delete(ctx context.Context, largePersonGroupID string, personID uuid.UUID) (result autorest.Response, err error)
	DeleteFace(ctx context.Context, largePersonGroupID string, personID uuid.UUID, persistedFaceID uuid.UUID) (result autorest.Response, err error)
	Get(ctx context.Context, largePersonGroupID string, personID uuid.UUID) (result face.Person, err error)
	GetFace(ctx context.Context, largePersonGroupID string, personID uuid.UUID, persistedFaceID uuid.UUID) (result face.PersistedFace, err error)
	List(ctx context.Context, largePersonGroupID string, start string, top *int32) (result face.ListPerson, err error)
	Update(ctx context.Context, largePersonGroupID string, personID uuid.UUID, body face.NameAndUserDataContract) (result autorest.Response, err error)
	UpdateFace(ctx context.Context, largePersonGroupID string, personID uuid.UUID, persistedFaceID uuid.UUID, body face.UpdateFaceRequest) (result autorest.Response, err error)
}

var _ LargePersonGroupPersonClientAPI = (*face.LargePersonGroupPersonClient)(nil)

// LargePersonGroupClientAPI contains the set of methods on the LargePersonGroupClient type.
type LargePersonGroupClientAPI interface {
	Create(ctx context.Context, largePersonGroupID string, body face.MetaDataContract) (result autorest.Response, err error)
	Delete(ctx context.Context, largePersonGroupID string) (result autorest.Response, err error)
	Get(ctx context.Context, largePersonGroupID string, returnRecognitionModel *bool) (result face.LargePersonGroup, err error)
	GetTrainingStatus(ctx context.Context, largePersonGroupID string) (result face.TrainingStatus, err error)
	List(ctx context.Context, start string, top *int32, returnRecognitionModel *bool) (result face.ListLargePersonGroup, err error)
	Train(ctx context.Context, largePersonGroupID string) (result autorest.Response, err error)
	Update(ctx context.Context, largePersonGroupID string, body face.NameAndUserDataContract) (result autorest.Response, err error)
}

var _ LargePersonGroupClientAPI = (*face.LargePersonGroupClient)(nil)

// LargeFaceListClientAPI contains the set of methods on the LargeFaceListClient type.
type LargeFaceListClientAPI interface {
	AddFaceFromStream(ctx context.Context, largeFaceListID string, imageParameter io.ReadCloser, userData string, targetFace []int32, detectionModel face.DetectionModel) (result face.PersistedFace, err error)
	AddFaceFromURL(ctx context.Context, largeFaceListID string, imageURL face.ImageURL, userData string, targetFace []int32, detectionModel face.DetectionModel) (result face.PersistedFace, err error)
	Create(ctx context.Context, largeFaceListID string, body face.MetaDataContractMandatoryName) (result autorest.Response, err error)
	Delete(ctx context.Context, largeFaceListID string) (result autorest.Response, err error)
	DeleteFace(ctx context.Context, largeFaceListID string, persistedFaceID uuid.UUID) (result autorest.Response, err error)
	Get(ctx context.Context, largeFaceListID string, returnRecognitionModel *bool) (result face.LargeFaceList, err error)
	GetFace(ctx context.Context, largeFaceListID string, persistedFaceID uuid.UUID) (result face.PersistedFace, err error)
	GetTrainingStatus(ctx context.Context, largeFaceListID string) (result face.TrainingStatus, err error)
	List(ctx context.Context, returnRecognitionModel *bool) (result face.ListLargeFaceList, err error)
	ListFaces(ctx context.Context, largeFaceListID string, start string, top *int32) (result face.ListPersistedFace, err error)
	Train(ctx context.Context, largeFaceListID string) (result autorest.Response, err error)
	Update(ctx context.Context, largeFaceListID string, body face.NameAndUserDataContract) (result autorest.Response, err error)
	UpdateFace(ctx context.Context, largeFaceListID string, persistedFaceID uuid.UUID, body face.UpdateFaceRequest) (result autorest.Response, err error)
}

var _ LargeFaceListClientAPI = (*face.LargeFaceListClient)(nil)

// SnapshotClientAPI contains the set of methods on the SnapshotClient type.
type SnapshotClientAPI interface {
	Apply(ctx context.Context, snapshotID uuid.UUID, body face.ApplySnapshotRequest) (result autorest.Response, err error)
	Delete(ctx context.Context, snapshotID uuid.UUID) (result autorest.Response, err error)
	Get(ctx context.Context, snapshotID uuid.UUID) (result face.Snapshot, err error)
	GetOperationStatus(ctx context.Context, operationID uuid.UUID) (result face.OperationStatus, err error)
	List(ctx context.Context, typeParameter face.SnapshotObjectType, applyScope []uuid.UUID) (result face.ListSnapshot, err error)
	Take(ctx context.Context, body face.TakeSnapshotRequest) (result autorest.Response, err error)
	Update(ctx context.Context, snapshotID uuid.UUID, body face.UpdateSnapshotRequest) (result autorest.Response, err error)
}

var _ SnapshotClientAPI = (*face.SnapshotClient)(nil)
