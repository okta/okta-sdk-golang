/*
Okta Admin Management

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

API version: 2025.10.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the ImportScheduleObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImportScheduleObject{}

// ImportScheduleObject Import schedule configuration
type ImportScheduleObject struct {
	// Determines the full import schedule
	FullImport *ImportScheduleSettings `json:"fullImport,omitempty"`
	// Determines the incremental import schedule
	IncrementalImport *ImportScheduleSettings `json:"incrementalImport,omitempty"`
	// Setting status
	Status               *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ImportScheduleObject ImportScheduleObject

// NewImportScheduleObject instantiates a new ImportScheduleObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportScheduleObject() *ImportScheduleObject {
	this := ImportScheduleObject{}
	return &this
}

// NewImportScheduleObjectWithDefaults instantiates a new ImportScheduleObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportScheduleObjectWithDefaults() *ImportScheduleObject {
	this := ImportScheduleObject{}
	return &this
}

// GetFullImport returns the FullImport field value if set, zero value otherwise.
func (o *ImportScheduleObject) GetFullImport() ImportScheduleSettings {
	if o == nil || IsNil(o.FullImport) {
		var ret ImportScheduleSettings
		return ret
	}
	return *o.FullImport
}

// GetFullImportOk returns a tuple with the FullImport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportScheduleObject) GetFullImportOk() (*ImportScheduleSettings, bool) {
	if o == nil || IsNil(o.FullImport) {
		return nil, false
	}
	return o.FullImport, true
}

// HasFullImport returns a boolean if a field has been set.
func (o *ImportScheduleObject) HasFullImport() bool {
	if o != nil && !IsNil(o.FullImport) {
		return true
	}

	return false
}

// SetFullImport gets a reference to the given ImportScheduleSettings and assigns it to the FullImport field.
func (o *ImportScheduleObject) SetFullImport(v ImportScheduleSettings) {
	o.FullImport = &v
}

// GetIncrementalImport returns the IncrementalImport field value if set, zero value otherwise.
func (o *ImportScheduleObject) GetIncrementalImport() ImportScheduleSettings {
	if o == nil || IsNil(o.IncrementalImport) {
		var ret ImportScheduleSettings
		return ret
	}
	return *o.IncrementalImport
}

// GetIncrementalImportOk returns a tuple with the IncrementalImport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportScheduleObject) GetIncrementalImportOk() (*ImportScheduleSettings, bool) {
	if o == nil || IsNil(o.IncrementalImport) {
		return nil, false
	}
	return o.IncrementalImport, true
}

// HasIncrementalImport returns a boolean if a field has been set.
func (o *ImportScheduleObject) HasIncrementalImport() bool {
	if o != nil && !IsNil(o.IncrementalImport) {
		return true
	}

	return false
}

// SetIncrementalImport gets a reference to the given ImportScheduleSettings and assigns it to the IncrementalImport field.
func (o *ImportScheduleObject) SetIncrementalImport(v ImportScheduleSettings) {
	o.IncrementalImport = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ImportScheduleObject) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportScheduleObject) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ImportScheduleObject) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ImportScheduleObject) SetStatus(v string) {
	o.Status = &v
}

func (o ImportScheduleObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImportScheduleObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FullImport) {
		toSerialize["fullImport"] = o.FullImport
	}
	if !IsNil(o.IncrementalImport) {
		toSerialize["incrementalImport"] = o.IncrementalImport
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ImportScheduleObject) UnmarshalJSON(data []byte) (err error) {
	varImportScheduleObject := _ImportScheduleObject{}

	err = json.Unmarshal(data, &varImportScheduleObject)

	if err != nil {
		return err
	}

	*o = ImportScheduleObject(varImportScheduleObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "fullImport")
		delete(additionalProperties, "incrementalImport")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableImportScheduleObject struct {
	value *ImportScheduleObject
	isSet bool
}

func (v NullableImportScheduleObject) Get() *ImportScheduleObject {
	return v.value
}

func (v *NullableImportScheduleObject) Set(val *ImportScheduleObject) {
	v.value = val
	v.isSet = true
}

func (v NullableImportScheduleObject) IsSet() bool {
	return v.isSet
}

func (v *NullableImportScheduleObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportScheduleObject(val *ImportScheduleObject) *NullableImportScheduleObject {
	return &NullableImportScheduleObject{value: val, isSet: true}
}

func (v NullableImportScheduleObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportScheduleObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
