// +build go1.9

// Copyright 2019 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package eventgrid

import original "github.com/Azure/azure-sdk-for-go/services/eventgrid/2018-01-01/eventgrid"

type MediaJobErrorCategory = original.MediaJobErrorCategory

const (
	Configuration MediaJobErrorCategory = original.Configuration
	Content       MediaJobErrorCategory = original.Content
	Download      MediaJobErrorCategory = original.Download
	Service       MediaJobErrorCategory = original.Service
	Upload        MediaJobErrorCategory = original.Upload
)

type MediaJobErrorCode = original.MediaJobErrorCode

const (
	ConfigurationUnsupported MediaJobErrorCode = original.ConfigurationUnsupported
	ContentMalformed         MediaJobErrorCode = original.ContentMalformed
	ContentUnsupported       MediaJobErrorCode = original.ContentUnsupported
	DownloadNotAccessible    MediaJobErrorCode = original.DownloadNotAccessible
	DownloadTransientError   MediaJobErrorCode = original.DownloadTransientError
	ServiceError             MediaJobErrorCode = original.ServiceError
	ServiceTransientError    MediaJobErrorCode = original.ServiceTransientError
	UploadNotAccessible      MediaJobErrorCode = original.UploadNotAccessible
	UploadTransientError     MediaJobErrorCode = original.UploadTransientError
)

type MediaJobRetry = original.MediaJobRetry

const (
	DoNotRetry MediaJobRetry = original.DoNotRetry
	MayRetry   MediaJobRetry = original.MayRetry
)

type MediaJobState = original.MediaJobState

const (
	Canceled   MediaJobState = original.Canceled
	Canceling  MediaJobState = original.Canceling
	Error      MediaJobState = original.Error
	Finished   MediaJobState = original.Finished
	Processing MediaJobState = original.Processing
	Queued     MediaJobState = original.Queued
	Scheduled  MediaJobState = original.Scheduled
)

type OdataType = original.OdataType

const (
	OdataTypeMediaJobOutput               OdataType = original.OdataTypeMediaJobOutput
	OdataTypeMicrosoftMediaJobOutputAsset OdataType = original.OdataTypeMicrosoftMediaJobOutputAsset
)

type AppConfigurationKeyValueDeletedEventData = original.AppConfigurationKeyValueDeletedEventData
type AppConfigurationKeyValueModifiedEventData = original.AppConfigurationKeyValueModifiedEventData
type BaseClient = original.BaseClient
type BasicMediaJobOutput = original.BasicMediaJobOutput
type ContainerRegistryArtifactEventData = original.ContainerRegistryArtifactEventData
type ContainerRegistryArtifactEventTarget = original.ContainerRegistryArtifactEventTarget
type ContainerRegistryChartDeletedEventData = original.ContainerRegistryChartDeletedEventData
type ContainerRegistryChartPushedEventData = original.ContainerRegistryChartPushedEventData
type ContainerRegistryEventActor = original.ContainerRegistryEventActor
type ContainerRegistryEventData = original.ContainerRegistryEventData
type ContainerRegistryEventRequest = original.ContainerRegistryEventRequest
type ContainerRegistryEventSource = original.ContainerRegistryEventSource
type ContainerRegistryEventTarget = original.ContainerRegistryEventTarget
type ContainerRegistryImageDeletedEventData = original.ContainerRegistryImageDeletedEventData
type ContainerRegistryImagePushedEventData = original.ContainerRegistryImagePushedEventData
type DeviceConnectionStateEventInfo = original.DeviceConnectionStateEventInfo
type DeviceConnectionStateEventProperties = original.DeviceConnectionStateEventProperties
type DeviceLifeCycleEventProperties = original.DeviceLifeCycleEventProperties
type DeviceTelemetryEventProperties = original.DeviceTelemetryEventProperties
type DeviceTwinInfo = original.DeviceTwinInfo
type DeviceTwinInfoProperties = original.DeviceTwinInfoProperties
type DeviceTwinInfoX509Thumbprint = original.DeviceTwinInfoX509Thumbprint
type DeviceTwinMetadata = original.DeviceTwinMetadata
type DeviceTwinProperties = original.DeviceTwinProperties
type Event = original.Event
type EventHubCaptureFileCreatedEventData = original.EventHubCaptureFileCreatedEventData
type IotHubDeviceConnectedEventData = original.IotHubDeviceConnectedEventData
type IotHubDeviceCreatedEventData = original.IotHubDeviceCreatedEventData
type IotHubDeviceDeletedEventData = original.IotHubDeviceDeletedEventData
type IotHubDeviceDisconnectedEventData = original.IotHubDeviceDisconnectedEventData
type IotHubDeviceTelemetryEventData = original.IotHubDeviceTelemetryEventData
type KeyVaultCertificateExpiredEventData = original.KeyVaultCertificateExpiredEventData
type KeyVaultCertificateNearExpiryEventData = original.KeyVaultCertificateNearExpiryEventData
type KeyVaultCertificateNewVersionCreatedEventData = original.KeyVaultCertificateNewVersionCreatedEventData
type KeyVaultKeyExpiredEventData = original.KeyVaultKeyExpiredEventData
type KeyVaultKeyNearExpiryEventData = original.KeyVaultKeyNearExpiryEventData
type KeyVaultKeyNewVersionCreatedEventData = original.KeyVaultKeyNewVersionCreatedEventData
type KeyVaultSecretExpiredEventData = original.KeyVaultSecretExpiredEventData
type KeyVaultSecretNearExpiryEventData = original.KeyVaultSecretNearExpiryEventData
type KeyVaultSecretNewVersionCreatedEventData = original.KeyVaultSecretNewVersionCreatedEventData
type MapsGeofenceEnteredEventData = original.MapsGeofenceEnteredEventData
type MapsGeofenceEventProperties = original.MapsGeofenceEventProperties
type MapsGeofenceExitedEventData = original.MapsGeofenceExitedEventData
type MapsGeofenceGeometry = original.MapsGeofenceGeometry
type MapsGeofenceResultEventData = original.MapsGeofenceResultEventData
type MediaJobCanceledEventData = original.MediaJobCanceledEventData
type MediaJobCancelingEventData = original.MediaJobCancelingEventData
type MediaJobError = original.MediaJobError
type MediaJobErrorDetail = original.MediaJobErrorDetail
type MediaJobErroredEventData = original.MediaJobErroredEventData
type MediaJobFinishedEventData = original.MediaJobFinishedEventData
type MediaJobOutput = original.MediaJobOutput
type MediaJobOutputAsset = original.MediaJobOutputAsset
type MediaJobOutputCanceledEventData = original.MediaJobOutputCanceledEventData
type MediaJobOutputCancelingEventData = original.MediaJobOutputCancelingEventData
type MediaJobOutputErroredEventData = original.MediaJobOutputErroredEventData
type MediaJobOutputFinishedEventData = original.MediaJobOutputFinishedEventData
type MediaJobOutputProcessingEventData = original.MediaJobOutputProcessingEventData
type MediaJobOutputProgressEventData = original.MediaJobOutputProgressEventData
type MediaJobOutputScheduledEventData = original.MediaJobOutputScheduledEventData
type MediaJobOutputStateChangeEventData = original.MediaJobOutputStateChangeEventData
type MediaJobProcessingEventData = original.MediaJobProcessingEventData
type MediaJobScheduledEventData = original.MediaJobScheduledEventData
type MediaJobStateChangeEventData = original.MediaJobStateChangeEventData
type MediaLiveEventConnectionRejectedEventData = original.MediaLiveEventConnectionRejectedEventData
type MediaLiveEventEncoderConnectedEventData = original.MediaLiveEventEncoderConnectedEventData
type MediaLiveEventEncoderDisconnectedEventData = original.MediaLiveEventEncoderDisconnectedEventData
type MediaLiveEventIncomingDataChunkDroppedEventData = original.MediaLiveEventIncomingDataChunkDroppedEventData
type MediaLiveEventIncomingStreamReceivedEventData = original.MediaLiveEventIncomingStreamReceivedEventData
type MediaLiveEventIncomingStreamsOutOfSyncEventData = original.MediaLiveEventIncomingStreamsOutOfSyncEventData
type MediaLiveEventIncomingVideoStreamsOutOfSyncEventData = original.MediaLiveEventIncomingVideoStreamsOutOfSyncEventData
type MediaLiveEventIngestHeartbeatEventData = original.MediaLiveEventIngestHeartbeatEventData
type MediaLiveEventTrackDiscontinuityDetectedEventData = original.MediaLiveEventTrackDiscontinuityDetectedEventData
type ResourceActionCancelData = original.ResourceActionCancelData
type ResourceActionFailureData = original.ResourceActionFailureData
type ResourceActionSuccessData = original.ResourceActionSuccessData
type ResourceDeleteCancelData = original.ResourceDeleteCancelData
type ResourceDeleteFailureData = original.ResourceDeleteFailureData
type ResourceDeleteSuccessData = original.ResourceDeleteSuccessData
type ResourceWriteCancelData = original.ResourceWriteCancelData
type ResourceWriteFailureData = original.ResourceWriteFailureData
type ResourceWriteSuccessData = original.ResourceWriteSuccessData
type ServiceBusActiveMessagesAvailableWithNoListenersEventData = original.ServiceBusActiveMessagesAvailableWithNoListenersEventData
type ServiceBusDeadletterMessagesAvailableWithNoListenersEventData = original.ServiceBusDeadletterMessagesAvailableWithNoListenersEventData
type SignalRServiceClientConnectionConnectedEventData = original.SignalRServiceClientConnectionConnectedEventData
type SignalRServiceClientConnectionDisconnectedEventData = original.SignalRServiceClientConnectionDisconnectedEventData
type StorageBlobCreatedEventData = original.StorageBlobCreatedEventData
type StorageBlobDeletedEventData = original.StorageBlobDeletedEventData
type StorageBlobRenamedEventData = original.StorageBlobRenamedEventData
type StorageDirectoryCreatedEventData = original.StorageDirectoryCreatedEventData
type StorageDirectoryDeletedEventData = original.StorageDirectoryDeletedEventData
type StorageDirectoryRenamedEventData = original.StorageDirectoryRenamedEventData
type SubscriptionDeletedEventData = original.SubscriptionDeletedEventData
type SubscriptionValidationEventData = original.SubscriptionValidationEventData
type SubscriptionValidationResponse = original.SubscriptionValidationResponse

func New() BaseClient {
	return original.New()
}
func NewWithoutDefaults() BaseClient {
	return original.NewWithoutDefaults()
}
func PossibleMediaJobErrorCategoryValues() []MediaJobErrorCategory {
	return original.PossibleMediaJobErrorCategoryValues()
}
func PossibleMediaJobErrorCodeValues() []MediaJobErrorCode {
	return original.PossibleMediaJobErrorCodeValues()
}
func PossibleMediaJobRetryValues() []MediaJobRetry {
	return original.PossibleMediaJobRetryValues()
}
func PossibleMediaJobStateValues() []MediaJobState {
	return original.PossibleMediaJobStateValues()
}
func PossibleOdataTypeValues() []OdataType {
	return original.PossibleOdataTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
