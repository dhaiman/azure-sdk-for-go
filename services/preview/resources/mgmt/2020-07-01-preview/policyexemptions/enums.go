package policy

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

// CreatedByType enumerates the values for created by type.
type CreatedByType string

const (
	// Application ...
	Application CreatedByType = "Application"
	// Key ...
	Key CreatedByType = "Key"
	// ManagedIdentity ...
	ManagedIdentity CreatedByType = "ManagedIdentity"
	// User ...
	User CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns an array of possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{Application, Key, ManagedIdentity, User}
}

// ExemptionCategory enumerates the values for exemption category.
type ExemptionCategory string

const (
	// Mitigated This category of exemptions usually means the mitigation actions have been applied to the
	// scope.
	Mitigated ExemptionCategory = "Mitigated"
	// Waiver This category of exemptions usually means the scope is not applicable for the policy.
	Waiver ExemptionCategory = "Waiver"
)

// PossibleExemptionCategoryValues returns an array of possible values for the ExemptionCategory const type.
func PossibleExemptionCategoryValues() []ExemptionCategory {
	return []ExemptionCategory{Mitigated, Waiver}
}
