package securityinsight

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

// AlertRuleKind enumerates the values for alert rule kind.
type AlertRuleKind string

const (
	// Fusion ...
	Fusion AlertRuleKind = "Fusion"
	// MicrosoftSecurityIncidentCreation ...
	MicrosoftSecurityIncidentCreation AlertRuleKind = "MicrosoftSecurityIncidentCreation"
	// Scheduled ...
	Scheduled AlertRuleKind = "Scheduled"
)

// PossibleAlertRuleKindValues returns an array of possible values for the AlertRuleKind const type.
func PossibleAlertRuleKindValues() []AlertRuleKind {
	return []AlertRuleKind{Fusion, MicrosoftSecurityIncidentCreation, Scheduled}
}

// AlertSeverity enumerates the values for alert severity.
type AlertSeverity string

const (
	// High High severity
	High AlertSeverity = "High"
	// Informational Informational severity
	Informational AlertSeverity = "Informational"
	// Low Low severity
	Low AlertSeverity = "Low"
	// Medium Medium severity
	Medium AlertSeverity = "Medium"
)

// PossibleAlertSeverityValues returns an array of possible values for the AlertSeverity const type.
func PossibleAlertSeverityValues() []AlertSeverity {
	return []AlertSeverity{High, Informational, Low, Medium}
}

// AttackTactic enumerates the values for attack tactic.
type AttackTactic string

const (
	// Collection ...
	Collection AttackTactic = "Collection"
	// CommandAndControl ...
	CommandAndControl AttackTactic = "CommandAndControl"
	// CredentialAccess ...
	CredentialAccess AttackTactic = "CredentialAccess"
	// DefenseEvasion ...
	DefenseEvasion AttackTactic = "DefenseEvasion"
	// Discovery ...
	Discovery AttackTactic = "Discovery"
	// Execution ...
	Execution AttackTactic = "Execution"
	// Exfiltration ...
	Exfiltration AttackTactic = "Exfiltration"
	// Impact ...
	Impact AttackTactic = "Impact"
	// InitialAccess ...
	InitialAccess AttackTactic = "InitialAccess"
	// LateralMovement ...
	LateralMovement AttackTactic = "LateralMovement"
	// Persistence ...
	Persistence AttackTactic = "Persistence"
	// PrivilegeEscalation ...
	PrivilegeEscalation AttackTactic = "PrivilegeEscalation"
)

// PossibleAttackTacticValues returns an array of possible values for the AttackTactic const type.
func PossibleAttackTacticValues() []AttackTactic {
	return []AttackTactic{Collection, CommandAndControl, CredentialAccess, DefenseEvasion, Discovery, Execution, Exfiltration, Impact, InitialAccess, LateralMovement, Persistence, PrivilegeEscalation}
}

// CaseSeverity enumerates the values for case severity.
type CaseSeverity string

const (
	// CaseSeverityCritical Critical severity
	CaseSeverityCritical CaseSeverity = "Critical"
	// CaseSeverityHigh High severity
	CaseSeverityHigh CaseSeverity = "High"
	// CaseSeverityInformational Informational severity
	CaseSeverityInformational CaseSeverity = "Informational"
	// CaseSeverityLow Low severity
	CaseSeverityLow CaseSeverity = "Low"
	// CaseSeverityMedium Medium severity
	CaseSeverityMedium CaseSeverity = "Medium"
)

// PossibleCaseSeverityValues returns an array of possible values for the CaseSeverity const type.
func PossibleCaseSeverityValues() []CaseSeverity {
	return []CaseSeverity{CaseSeverityCritical, CaseSeverityHigh, CaseSeverityInformational, CaseSeverityLow, CaseSeverityMedium}
}

// DataConnectorKind enumerates the values for data connector kind.
type DataConnectorKind string

const (
	// DataConnectorKindAmazonWebServicesCloudTrail ...
	DataConnectorKindAmazonWebServicesCloudTrail DataConnectorKind = "AmazonWebServicesCloudTrail"
	// DataConnectorKindAzureActiveDirectory ...
	DataConnectorKindAzureActiveDirectory DataConnectorKind = "AzureActiveDirectory"
	// DataConnectorKindAzureAdvancedThreatProtection ...
	DataConnectorKindAzureAdvancedThreatProtection DataConnectorKind = "AzureAdvancedThreatProtection"
	// DataConnectorKindAzureSecurityCenter ...
	DataConnectorKindAzureSecurityCenter DataConnectorKind = "AzureSecurityCenter"
	// DataConnectorKindMicrosoftCloudAppSecurity ...
	DataConnectorKindMicrosoftCloudAppSecurity DataConnectorKind = "MicrosoftCloudAppSecurity"
	// DataConnectorKindMicrosoftDefenderAdvancedThreatProtection ...
	DataConnectorKindMicrosoftDefenderAdvancedThreatProtection DataConnectorKind = "MicrosoftDefenderAdvancedThreatProtection"
	// DataConnectorKindOffice365 ...
	DataConnectorKindOffice365 DataConnectorKind = "Office365"
	// DataConnectorKindThreatIntelligence ...
	DataConnectorKindThreatIntelligence DataConnectorKind = "ThreatIntelligence"
)

// PossibleDataConnectorKindValues returns an array of possible values for the DataConnectorKind const type.
func PossibleDataConnectorKindValues() []DataConnectorKind {
	return []DataConnectorKind{DataConnectorKindAmazonWebServicesCloudTrail, DataConnectorKindAzureActiveDirectory, DataConnectorKindAzureAdvancedThreatProtection, DataConnectorKindAzureSecurityCenter, DataConnectorKindMicrosoftCloudAppSecurity, DataConnectorKindMicrosoftDefenderAdvancedThreatProtection, DataConnectorKindOffice365, DataConnectorKindThreatIntelligence}
}

// DataTypeState enumerates the values for data type state.
type DataTypeState string

const (
	// Disabled ...
	Disabled DataTypeState = "Disabled"
	// Enabled ...
	Enabled DataTypeState = "Enabled"
)

// PossibleDataTypeStateValues returns an array of possible values for the DataTypeState const type.
func PossibleDataTypeStateValues() []DataTypeState {
	return []DataTypeState{Disabled, Enabled}
}

// IncidentClassification enumerates the values for incident classification.
type IncidentClassification string

const (
	// BenignPositive Incident was benign positive
	BenignPositive IncidentClassification = "BenignPositive"
	// FalsePositive Incident was false positive
	FalsePositive IncidentClassification = "FalsePositive"
	// TruePositive Incident was true positive
	TruePositive IncidentClassification = "TruePositive"
	// Undetermined Incident classification was undetermined
	Undetermined IncidentClassification = "Undetermined"
)

// PossibleIncidentClassificationValues returns an array of possible values for the IncidentClassification const type.
func PossibleIncidentClassificationValues() []IncidentClassification {
	return []IncidentClassification{BenignPositive, FalsePositive, TruePositive, Undetermined}
}

// IncidentClassificationReason enumerates the values for incident classification reason.
type IncidentClassificationReason string

const (
	// InaccurateData Classification reason was inaccurate data
	InaccurateData IncidentClassificationReason = "InaccurateData"
	// IncorrectAlertLogic Classification reason was incorrect alert logic
	IncorrectAlertLogic IncidentClassificationReason = "IncorrectAlertLogic"
	// SuspiciousActivity Classification reason was suspicious activity
	SuspiciousActivity IncidentClassificationReason = "SuspiciousActivity"
	// SuspiciousButExpected Classification reason was suspicious but expected
	SuspiciousButExpected IncidentClassificationReason = "SuspiciousButExpected"
)

// PossibleIncidentClassificationReasonValues returns an array of possible values for the IncidentClassificationReason const type.
func PossibleIncidentClassificationReasonValues() []IncidentClassificationReason {
	return []IncidentClassificationReason{InaccurateData, IncorrectAlertLogic, SuspiciousActivity, SuspiciousButExpected}
}

// IncidentLabelType enumerates the values for incident label type.
type IncidentLabelType string

const (
	// AutoAssigned Label automatically created by the system
	AutoAssigned IncidentLabelType = "AutoAssigned"
	// User Label manually created by a user
	User IncidentLabelType = "User"
)

// PossibleIncidentLabelTypeValues returns an array of possible values for the IncidentLabelType const type.
func PossibleIncidentLabelTypeValues() []IncidentLabelType {
	return []IncidentLabelType{AutoAssigned, User}
}

// IncidentSeverity enumerates the values for incident severity.
type IncidentSeverity string

const (
	// IncidentSeverityHigh High severity
	IncidentSeverityHigh IncidentSeverity = "High"
	// IncidentSeverityInformational Informational severity
	IncidentSeverityInformational IncidentSeverity = "Informational"
	// IncidentSeverityLow Low severity
	IncidentSeverityLow IncidentSeverity = "Low"
	// IncidentSeverityMedium Medium severity
	IncidentSeverityMedium IncidentSeverity = "Medium"
)

// PossibleIncidentSeverityValues returns an array of possible values for the IncidentSeverity const type.
func PossibleIncidentSeverityValues() []IncidentSeverity {
	return []IncidentSeverity{IncidentSeverityHigh, IncidentSeverityInformational, IncidentSeverityLow, IncidentSeverityMedium}
}

// IncidentStatus enumerates the values for incident status.
type IncidentStatus string

const (
	// IncidentStatusActive An active incident which is being handled
	IncidentStatusActive IncidentStatus = "Active"
	// IncidentStatusClosed A non-active incident
	IncidentStatusClosed IncidentStatus = "Closed"
	// IncidentStatusNew An active incident which isn't being handled currently
	IncidentStatusNew IncidentStatus = "New"
)

// PossibleIncidentStatusValues returns an array of possible values for the IncidentStatus const type.
func PossibleIncidentStatusValues() []IncidentStatus {
	return []IncidentStatus{IncidentStatusActive, IncidentStatusClosed, IncidentStatusNew}
}

// Kind enumerates the values for kind.
type Kind string

const (
	// KindAlertRule ...
	KindAlertRule Kind = "AlertRule"
	// KindFusion ...
	KindFusion Kind = "Fusion"
	// KindMicrosoftSecurityIncidentCreation ...
	KindMicrosoftSecurityIncidentCreation Kind = "MicrosoftSecurityIncidentCreation"
	// KindScheduled ...
	KindScheduled Kind = "Scheduled"
)

// PossibleKindValues returns an array of possible values for the Kind const type.
func PossibleKindValues() []Kind {
	return []Kind{KindAlertRule, KindFusion, KindMicrosoftSecurityIncidentCreation, KindScheduled}
}

// KindBasicAlertRuleTemplate enumerates the values for kind basic alert rule template.
type KindBasicAlertRuleTemplate string

const (
	// KindBasicAlertRuleTemplateKindAlertRuleTemplate ...
	KindBasicAlertRuleTemplateKindAlertRuleTemplate KindBasicAlertRuleTemplate = "AlertRuleTemplate"
	// KindBasicAlertRuleTemplateKindFusion ...
	KindBasicAlertRuleTemplateKindFusion KindBasicAlertRuleTemplate = "Fusion"
	// KindBasicAlertRuleTemplateKindMicrosoftSecurityIncidentCreation ...
	KindBasicAlertRuleTemplateKindMicrosoftSecurityIncidentCreation KindBasicAlertRuleTemplate = "MicrosoftSecurityIncidentCreation"
	// KindBasicAlertRuleTemplateKindScheduled ...
	KindBasicAlertRuleTemplateKindScheduled KindBasicAlertRuleTemplate = "Scheduled"
)

// PossibleKindBasicAlertRuleTemplateValues returns an array of possible values for the KindBasicAlertRuleTemplate const type.
func PossibleKindBasicAlertRuleTemplateValues() []KindBasicAlertRuleTemplate {
	return []KindBasicAlertRuleTemplate{KindBasicAlertRuleTemplateKindAlertRuleTemplate, KindBasicAlertRuleTemplateKindFusion, KindBasicAlertRuleTemplateKindMicrosoftSecurityIncidentCreation, KindBasicAlertRuleTemplateKindScheduled}
}

// KindBasicDataConnector enumerates the values for kind basic data connector.
type KindBasicDataConnector string

const (
	// KindAmazonWebServicesCloudTrail ...
	KindAmazonWebServicesCloudTrail KindBasicDataConnector = "AmazonWebServicesCloudTrail"
	// KindAzureActiveDirectory ...
	KindAzureActiveDirectory KindBasicDataConnector = "AzureActiveDirectory"
	// KindAzureAdvancedThreatProtection ...
	KindAzureAdvancedThreatProtection KindBasicDataConnector = "AzureAdvancedThreatProtection"
	// KindAzureSecurityCenter ...
	KindAzureSecurityCenter KindBasicDataConnector = "AzureSecurityCenter"
	// KindDataConnector ...
	KindDataConnector KindBasicDataConnector = "DataConnector"
	// KindMicrosoftCloudAppSecurity ...
	KindMicrosoftCloudAppSecurity KindBasicDataConnector = "MicrosoftCloudAppSecurity"
	// KindMicrosoftDefenderAdvancedThreatProtection ...
	KindMicrosoftDefenderAdvancedThreatProtection KindBasicDataConnector = "MicrosoftDefenderAdvancedThreatProtection"
	// KindOffice365 ...
	KindOffice365 KindBasicDataConnector = "Office365"
	// KindThreatIntelligence ...
	KindThreatIntelligence KindBasicDataConnector = "ThreatIntelligence"
)

// PossibleKindBasicDataConnectorValues returns an array of possible values for the KindBasicDataConnector const type.
func PossibleKindBasicDataConnectorValues() []KindBasicDataConnector {
	return []KindBasicDataConnector{KindAmazonWebServicesCloudTrail, KindAzureActiveDirectory, KindAzureAdvancedThreatProtection, KindAzureSecurityCenter, KindDataConnector, KindMicrosoftCloudAppSecurity, KindMicrosoftDefenderAdvancedThreatProtection, KindOffice365, KindThreatIntelligence}
}

// KindBasicSettings enumerates the values for kind basic settings.
type KindBasicSettings string

const (
	// KindSettings ...
	KindSettings KindBasicSettings = "Settings"
	// KindToggleSettings ...
	KindToggleSettings KindBasicSettings = "ToggleSettings"
	// KindUebaSettings ...
	KindUebaSettings KindBasicSettings = "UebaSettings"
)

// PossibleKindBasicSettingsValues returns an array of possible values for the KindBasicSettings const type.
func PossibleKindBasicSettingsValues() []KindBasicSettings {
	return []KindBasicSettings{KindSettings, KindToggleSettings, KindUebaSettings}
}

// LicenseStatus enumerates the values for license status.
type LicenseStatus string

const (
	// LicenseStatusDisabled ...
	LicenseStatusDisabled LicenseStatus = "Disabled"
	// LicenseStatusEnabled ...
	LicenseStatusEnabled LicenseStatus = "Enabled"
)

// PossibleLicenseStatusValues returns an array of possible values for the LicenseStatus const type.
func PossibleLicenseStatusValues() []LicenseStatus {
	return []LicenseStatus{LicenseStatusDisabled, LicenseStatusEnabled}
}

// MicrosoftSecurityProductName enumerates the values for microsoft security product name.
type MicrosoftSecurityProductName string

const (
	// AzureActiveDirectoryIdentityProtection ...
	AzureActiveDirectoryIdentityProtection MicrosoftSecurityProductName = "Azure Active Directory Identity Protection"
	// AzureAdvancedThreatProtection ...
	AzureAdvancedThreatProtection MicrosoftSecurityProductName = "Azure Advanced Threat Protection"
	// AzureSecurityCenter ...
	AzureSecurityCenter MicrosoftSecurityProductName = "Azure Security Center"
	// AzureSecurityCenterforIoT ...
	AzureSecurityCenterforIoT MicrosoftSecurityProductName = "Azure Security Center for IoT"
	// MicrosoftCloudAppSecurity ...
	MicrosoftCloudAppSecurity MicrosoftSecurityProductName = "Microsoft Cloud App Security"
)

// PossibleMicrosoftSecurityProductNameValues returns an array of possible values for the MicrosoftSecurityProductName const type.
func PossibleMicrosoftSecurityProductNameValues() []MicrosoftSecurityProductName {
	return []MicrosoftSecurityProductName{AzureActiveDirectoryIdentityProtection, AzureAdvancedThreatProtection, AzureSecurityCenter, AzureSecurityCenterforIoT, MicrosoftCloudAppSecurity}
}

// SettingKind enumerates the values for setting kind.
type SettingKind string

const (
	// SettingKindToggleSettings ...
	SettingKindToggleSettings SettingKind = "ToggleSettings"
	// SettingKindUebaSettings ...
	SettingKindUebaSettings SettingKind = "UebaSettings"
)

// PossibleSettingKindValues returns an array of possible values for the SettingKind const type.
func PossibleSettingKindValues() []SettingKind {
	return []SettingKind{SettingKindToggleSettings, SettingKindUebaSettings}
}

// StatusInMcas enumerates the values for status in mcas.
type StatusInMcas string

const (
	// StatusInMcasDisabled ...
	StatusInMcasDisabled StatusInMcas = "Disabled"
	// StatusInMcasEnabled ...
	StatusInMcasEnabled StatusInMcas = "Enabled"
)

// PossibleStatusInMcasValues returns an array of possible values for the StatusInMcas const type.
func PossibleStatusInMcasValues() []StatusInMcas {
	return []StatusInMcas{StatusInMcasDisabled, StatusInMcasEnabled}
}

// TemplateStatus enumerates the values for template status.
type TemplateStatus string

const (
	// Available Alert rule template is available.
	Available TemplateStatus = "Available"
	// Installed Alert rule template installed. and can not use more then once
	Installed TemplateStatus = "Installed"
	// NotAvailable Alert rule template is not available
	NotAvailable TemplateStatus = "NotAvailable"
)

// PossibleTemplateStatusValues returns an array of possible values for the TemplateStatus const type.
func PossibleTemplateStatusValues() []TemplateStatus {
	return []TemplateStatus{Available, Installed, NotAvailable}
}

// TriggerOperator enumerates the values for trigger operator.
type TriggerOperator string

const (
	// Equal ...
	Equal TriggerOperator = "Equal"
	// GreaterThan ...
	GreaterThan TriggerOperator = "GreaterThan"
	// LessThan ...
	LessThan TriggerOperator = "LessThan"
	// NotEqual ...
	NotEqual TriggerOperator = "NotEqual"
)

// PossibleTriggerOperatorValues returns an array of possible values for the TriggerOperator const type.
func PossibleTriggerOperatorValues() []TriggerOperator {
	return []TriggerOperator{Equal, GreaterThan, LessThan, NotEqual}
}
