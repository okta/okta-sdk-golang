/*
Okta Admin Management API

Allows customers to easily access the Okta Management APIs

Copyright 2025 - Present Okta, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

API version: 2025.08.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the ImportRoadblockAffectedApp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImportRoadblockAffectedApp{}

// ImportRoadblockAffectedApp An app whose assignments the import would remove
type ImportRoadblockAffectedApp struct {
	// ID of the app instance that would lose assignments
	AppInstanceId *string `json:"appInstanceId,omitempty"`
	// Label of the app instance that would lose assignments
	AppInstanceLabel *string `json:"appInstanceLabel,omitempty"`
	// Number of assignments to this app used as the denominator for `unassignmentPercentage`
	Assignments *int32 `json:"assignments,omitempty"`
	// The app-level percentage threshold that this app crossed
	Threshold *int32 `json:"threshold,omitempty"`
	// The share of assignments to this app that the import would remove (`unassignments` / `assignments` × 100)
	UnassignmentPercentage *float32 `json:"unassignmentPercentage,omitempty"`
	// Number of assignments to this app that the import would remove
	Unassignments        *int32 `json:"unassignments,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ImportRoadblockAffectedApp ImportRoadblockAffectedApp

// NewImportRoadblockAffectedApp instantiates a new ImportRoadblockAffectedApp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportRoadblockAffectedApp() *ImportRoadblockAffectedApp {
	this := ImportRoadblockAffectedApp{}
	return &this
}

// NewImportRoadblockAffectedAppWithDefaults instantiates a new ImportRoadblockAffectedApp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportRoadblockAffectedAppWithDefaults() *ImportRoadblockAffectedApp {
	this := ImportRoadblockAffectedApp{}
	return &this
}

// GetAppInstanceId returns the AppInstanceId field value if set, zero value otherwise.
func (o *ImportRoadblockAffectedApp) GetAppInstanceId() string {
	if o == nil || IsNil(o.AppInstanceId) {
		var ret string
		return ret
	}
	return *o.AppInstanceId
}

// GetAppInstanceIdOk returns a tuple with the AppInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportRoadblockAffectedApp) GetAppInstanceIdOk() (*string, bool) {
	if o == nil || IsNil(o.AppInstanceId) {
		return nil, false
	}
	return o.AppInstanceId, true
}

// HasAppInstanceId returns a boolean if a field has been set.
func (o *ImportRoadblockAffectedApp) HasAppInstanceId() bool {
	if o != nil && !IsNil(o.AppInstanceId) {
		return true
	}

	return false
}

// SetAppInstanceId gets a reference to the given string and assigns it to the AppInstanceId field.
func (o *ImportRoadblockAffectedApp) SetAppInstanceId(v string) {
	o.AppInstanceId = &v
}

// GetAppInstanceLabel returns the AppInstanceLabel field value if set, zero value otherwise.
func (o *ImportRoadblockAffectedApp) GetAppInstanceLabel() string {
	if o == nil || IsNil(o.AppInstanceLabel) {
		var ret string
		return ret
	}
	return *o.AppInstanceLabel
}

// GetAppInstanceLabelOk returns a tuple with the AppInstanceLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportRoadblockAffectedApp) GetAppInstanceLabelOk() (*string, bool) {
	if o == nil || IsNil(o.AppInstanceLabel) {
		return nil, false
	}
	return o.AppInstanceLabel, true
}

// HasAppInstanceLabel returns a boolean if a field has been set.
func (o *ImportRoadblockAffectedApp) HasAppInstanceLabel() bool {
	if o != nil && !IsNil(o.AppInstanceLabel) {
		return true
	}

	return false
}

// SetAppInstanceLabel gets a reference to the given string and assigns it to the AppInstanceLabel field.
func (o *ImportRoadblockAffectedApp) SetAppInstanceLabel(v string) {
	o.AppInstanceLabel = &v
}

// GetAssignments returns the Assignments field value if set, zero value otherwise.
func (o *ImportRoadblockAffectedApp) GetAssignments() int32 {
	if o == nil || IsNil(o.Assignments) {
		var ret int32
		return ret
	}
	return *o.Assignments
}

// GetAssignmentsOk returns a tuple with the Assignments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportRoadblockAffectedApp) GetAssignmentsOk() (*int32, bool) {
	if o == nil || IsNil(o.Assignments) {
		return nil, false
	}
	return o.Assignments, true
}

// HasAssignments returns a boolean if a field has been set.
func (o *ImportRoadblockAffectedApp) HasAssignments() bool {
	if o != nil && !IsNil(o.Assignments) {
		return true
	}

	return false
}

// SetAssignments gets a reference to the given int32 and assigns it to the Assignments field.
func (o *ImportRoadblockAffectedApp) SetAssignments(v int32) {
	o.Assignments = &v
}

// GetThreshold returns the Threshold field value if set, zero value otherwise.
func (o *ImportRoadblockAffectedApp) GetThreshold() int32 {
	if o == nil || IsNil(o.Threshold) {
		var ret int32
		return ret
	}
	return *o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportRoadblockAffectedApp) GetThresholdOk() (*int32, bool) {
	if o == nil || IsNil(o.Threshold) {
		return nil, false
	}
	return o.Threshold, true
}

// HasThreshold returns a boolean if a field has been set.
func (o *ImportRoadblockAffectedApp) HasThreshold() bool {
	if o != nil && !IsNil(o.Threshold) {
		return true
	}

	return false
}

// SetThreshold gets a reference to the given int32 and assigns it to the Threshold field.
func (o *ImportRoadblockAffectedApp) SetThreshold(v int32) {
	o.Threshold = &v
}

// GetUnassignmentPercentage returns the UnassignmentPercentage field value if set, zero value otherwise.
func (o *ImportRoadblockAffectedApp) GetUnassignmentPercentage() float32 {
	if o == nil || IsNil(o.UnassignmentPercentage) {
		var ret float32
		return ret
	}
	return *o.UnassignmentPercentage
}

// GetUnassignmentPercentageOk returns a tuple with the UnassignmentPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportRoadblockAffectedApp) GetUnassignmentPercentageOk() (*float32, bool) {
	if o == nil || IsNil(o.UnassignmentPercentage) {
		return nil, false
	}
	return o.UnassignmentPercentage, true
}

// HasUnassignmentPercentage returns a boolean if a field has been set.
func (o *ImportRoadblockAffectedApp) HasUnassignmentPercentage() bool {
	if o != nil && !IsNil(o.UnassignmentPercentage) {
		return true
	}

	return false
}

// SetUnassignmentPercentage gets a reference to the given float32 and assigns it to the UnassignmentPercentage field.
func (o *ImportRoadblockAffectedApp) SetUnassignmentPercentage(v float32) {
	o.UnassignmentPercentage = &v
}

// GetUnassignments returns the Unassignments field value if set, zero value otherwise.
func (o *ImportRoadblockAffectedApp) GetUnassignments() int32 {
	if o == nil || IsNil(o.Unassignments) {
		var ret int32
		return ret
	}
	return *o.Unassignments
}

// GetUnassignmentsOk returns a tuple with the Unassignments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportRoadblockAffectedApp) GetUnassignmentsOk() (*int32, bool) {
	if o == nil || IsNil(o.Unassignments) {
		return nil, false
	}
	return o.Unassignments, true
}

// HasUnassignments returns a boolean if a field has been set.
func (o *ImportRoadblockAffectedApp) HasUnassignments() bool {
	if o != nil && !IsNil(o.Unassignments) {
		return true
	}

	return false
}

// SetUnassignments gets a reference to the given int32 and assigns it to the Unassignments field.
func (o *ImportRoadblockAffectedApp) SetUnassignments(v int32) {
	o.Unassignments = &v
}

func (o ImportRoadblockAffectedApp) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImportRoadblockAffectedApp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppInstanceId) {
		toSerialize["appInstanceId"] = o.AppInstanceId
	}
	if !IsNil(o.AppInstanceLabel) {
		toSerialize["appInstanceLabel"] = o.AppInstanceLabel
	}
	if !IsNil(o.Assignments) {
		toSerialize["assignments"] = o.Assignments
	}
	if !IsNil(o.Threshold) {
		toSerialize["threshold"] = o.Threshold
	}
	if !IsNil(o.UnassignmentPercentage) {
		toSerialize["unassignmentPercentage"] = o.UnassignmentPercentage
	}
	if !IsNil(o.Unassignments) {
		toSerialize["unassignments"] = o.Unassignments
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ImportRoadblockAffectedApp) UnmarshalJSON(data []byte) (err error) {
	varImportRoadblockAffectedApp := _ImportRoadblockAffectedApp{}

	err = json.Unmarshal(data, &varImportRoadblockAffectedApp)

	if err != nil {
		return err
	}

	*o = ImportRoadblockAffectedApp(varImportRoadblockAffectedApp)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "appInstanceId")
		delete(additionalProperties, "appInstanceLabel")
		delete(additionalProperties, "assignments")
		delete(additionalProperties, "threshold")
		delete(additionalProperties, "unassignmentPercentage")
		delete(additionalProperties, "unassignments")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableImportRoadblockAffectedApp struct {
	value *ImportRoadblockAffectedApp
	isSet bool
}

func (v NullableImportRoadblockAffectedApp) Get() *ImportRoadblockAffectedApp {
	return v.value
}

func (v *NullableImportRoadblockAffectedApp) Set(val *ImportRoadblockAffectedApp) {
	v.value = val
	v.isSet = true
}

func (v NullableImportRoadblockAffectedApp) IsSet() bool {
	return v.isSet
}

func (v *NullableImportRoadblockAffectedApp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportRoadblockAffectedApp(val *ImportRoadblockAffectedApp) *NullableImportRoadblockAffectedApp {
	return &NullableImportRoadblockAffectedApp{value: val, isSet: true}
}

func (v NullableImportRoadblockAffectedApp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportRoadblockAffectedApp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
