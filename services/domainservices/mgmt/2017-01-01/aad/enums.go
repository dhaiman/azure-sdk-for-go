package aad

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

// ExternalAccess enumerates the values for external access.
type ExternalAccess string

const (
	// Disabled ...
	Disabled ExternalAccess = "Disabled"
	// Enabled ...
	Enabled ExternalAccess = "Enabled"
)

// PossibleExternalAccessValues returns an array of possible values for the ExternalAccess const type.
func PossibleExternalAccessValues() []ExternalAccess {
	return []ExternalAccess{Disabled, Enabled}
}

// FilteredSync enumerates the values for filtered sync.
type FilteredSync string

const (
	// FilteredSyncDisabled ...
	FilteredSyncDisabled FilteredSync = "Disabled"
	// FilteredSyncEnabled ...
	FilteredSyncEnabled FilteredSync = "Enabled"
)

// PossibleFilteredSyncValues returns an array of possible values for the FilteredSync const type.
func PossibleFilteredSyncValues() []FilteredSync {
	return []FilteredSync{FilteredSyncDisabled, FilteredSyncEnabled}
}

// Ldaps enumerates the values for ldaps.
type Ldaps string

const (
	// LdapsDisabled ...
	LdapsDisabled Ldaps = "Disabled"
	// LdapsEnabled ...
	LdapsEnabled Ldaps = "Enabled"
)

// PossibleLdapsValues returns an array of possible values for the Ldaps const type.
func PossibleLdapsValues() []Ldaps {
	return []Ldaps{LdapsDisabled, LdapsEnabled}
}

// NotifyDcAdmins enumerates the values for notify dc admins.
type NotifyDcAdmins string

const (
	// NotifyDcAdminsDisabled ...
	NotifyDcAdminsDisabled NotifyDcAdmins = "Disabled"
	// NotifyDcAdminsEnabled ...
	NotifyDcAdminsEnabled NotifyDcAdmins = "Enabled"
)

// PossibleNotifyDcAdminsValues returns an array of possible values for the NotifyDcAdmins const type.
func PossibleNotifyDcAdminsValues() []NotifyDcAdmins {
	return []NotifyDcAdmins{NotifyDcAdminsDisabled, NotifyDcAdminsEnabled}
}

// NotifyGlobalAdmins enumerates the values for notify global admins.
type NotifyGlobalAdmins string

const (
	// NotifyGlobalAdminsDisabled ...
	NotifyGlobalAdminsDisabled NotifyGlobalAdmins = "Disabled"
	// NotifyGlobalAdminsEnabled ...
	NotifyGlobalAdminsEnabled NotifyGlobalAdmins = "Enabled"
)

// PossibleNotifyGlobalAdminsValues returns an array of possible values for the NotifyGlobalAdmins const type.
func PossibleNotifyGlobalAdminsValues() []NotifyGlobalAdmins {
	return []NotifyGlobalAdmins{NotifyGlobalAdminsDisabled, NotifyGlobalAdminsEnabled}
}

// NtlmV1 enumerates the values for ntlm v1.
type NtlmV1 string

const (
	// NtlmV1Disabled ...
	NtlmV1Disabled NtlmV1 = "Disabled"
	// NtlmV1Enabled ...
	NtlmV1Enabled NtlmV1 = "Enabled"
)

// PossibleNtlmV1Values returns an array of possible values for the NtlmV1 const type.
func PossibleNtlmV1Values() []NtlmV1 {
	return []NtlmV1{NtlmV1Disabled, NtlmV1Enabled}
}

// SyncNtlmPasswords enumerates the values for sync ntlm passwords.
type SyncNtlmPasswords string

const (
	// SyncNtlmPasswordsDisabled ...
	SyncNtlmPasswordsDisabled SyncNtlmPasswords = "Disabled"
	// SyncNtlmPasswordsEnabled ...
	SyncNtlmPasswordsEnabled SyncNtlmPasswords = "Enabled"
)

// PossibleSyncNtlmPasswordsValues returns an array of possible values for the SyncNtlmPasswords const type.
func PossibleSyncNtlmPasswordsValues() []SyncNtlmPasswords {
	return []SyncNtlmPasswords{SyncNtlmPasswordsDisabled, SyncNtlmPasswordsEnabled}
}

// TLSV1 enumerates the values for tlsv1.
type TLSV1 string

const (
	// TLSV1Disabled ...
	TLSV1Disabled TLSV1 = "Disabled"
	// TLSV1Enabled ...
	TLSV1Enabled TLSV1 = "Enabled"
)

// PossibleTLSV1Values returns an array of possible values for the TLSV1 const type.
func PossibleTLSV1Values() []TLSV1 {
	return []TLSV1{TLSV1Disabled, TLSV1Enabled}
}
