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

API version: 2025.08.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the LogTarget type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogTarget{}

// LogTarget struct for LogTarget
type LogTarget struct {
	// The alternate ID of the target
	AlternateId   *string                 `json:"alternateId,omitempty"`
	ChangeDetails *LogTargetChangeDetails `json:"changeDetails,omitempty"`
	// Further details on the target
	DetailEntry map[string]interface{} `json:"detailEntry,omitempty"`
	// The display name of the target
	DisplayName *string `json:"displayName,omitempty"`
	// The ID of the target
	Id *string `json:"id,omitempty"`
	// The type of target
	Type                 *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LogTarget LogTarget

// NewLogTarget instantiates a new LogTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogTarget() *LogTarget {
	this := LogTarget{}
	return &this
}

// NewLogTargetWithDefaults instantiates a new LogTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogTargetWithDefaults() *LogTarget {
	this := LogTarget{}
	return &this
}

// GetAlternateId returns the AlternateId field value if set, zero value otherwise.
func (o *LogTarget) GetAlternateId() string {
	if o == nil || IsNil(o.AlternateId) {
		var ret string
		return ret
	}
	return *o.AlternateId
}

// GetAlternateIdOk returns a tuple with the AlternateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogTarget) GetAlternateIdOk() (*string, bool) {
	if o == nil || IsNil(o.AlternateId) {
		return nil, false
	}
	return o.AlternateId, true
}

// HasAlternateId returns a boolean if a field has been set.
func (o *LogTarget) HasAlternateId() bool {
	if o != nil && !IsNil(o.AlternateId) {
		return true
	}

	return false
}

// SetAlternateId gets a reference to the given string and assigns it to the AlternateId field.
func (o *LogTarget) SetAlternateId(v string) {
	o.AlternateId = &v
}

// GetChangeDetails returns the ChangeDetails field value if set, zero value otherwise.
func (o *LogTarget) GetChangeDetails() LogTargetChangeDetails {
	if o == nil || IsNil(o.ChangeDetails) {
		var ret LogTargetChangeDetails
		return ret
	}
	return *o.ChangeDetails
}

// GetChangeDetailsOk returns a tuple with the ChangeDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogTarget) GetChangeDetailsOk() (*LogTargetChangeDetails, bool) {
	if o == nil || IsNil(o.ChangeDetails) {
		return nil, false
	}
	return o.ChangeDetails, true
}

// HasChangeDetails returns a boolean if a field has been set.
func (o *LogTarget) HasChangeDetails() bool {
	if o != nil && !IsNil(o.ChangeDetails) {
		return true
	}

	return false
}

// SetChangeDetails gets a reference to the given LogTargetChangeDetails and assigns it to the ChangeDetails field.
func (o *LogTarget) SetChangeDetails(v LogTargetChangeDetails) {
	o.ChangeDetails = &v
}

// GetDetailEntry returns the DetailEntry field value if set, zero value otherwise.
func (o *LogTarget) GetDetailEntry() map[string]interface{} {
	if o == nil || IsNil(o.DetailEntry) {
		var ret map[string]interface{}
		return ret
	}
	return o.DetailEntry
}

// GetDetailEntryOk returns a tuple with the DetailEntry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogTarget) GetDetailEntryOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.DetailEntry) {
		return map[string]interface{}{}, false
	}
	return o.DetailEntry, true
}

// HasDetailEntry returns a boolean if a field has been set.
func (o *LogTarget) HasDetailEntry() bool {
	if o != nil && !IsNil(o.DetailEntry) {
		return true
	}

	return false
}

// SetDetailEntry gets a reference to the given map[string]interface{} and assigns it to the DetailEntry field.
func (o *LogTarget) SetDetailEntry(v map[string]interface{}) {
	o.DetailEntry = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *LogTarget) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogTarget) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *LogTarget) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *LogTarget) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LogTarget) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogTarget) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LogTarget) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LogTarget) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *LogTarget) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogTarget) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *LogTarget) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *LogTarget) SetType(v string) {
	o.Type = &v
}

func (o LogTarget) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogTarget) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AlternateId) {
		toSerialize["alternateId"] = o.AlternateId
	}
	if !IsNil(o.ChangeDetails) {
		toSerialize["changeDetails"] = o.ChangeDetails
	}
	if !IsNil(o.DetailEntry) {
		toSerialize["detailEntry"] = o.DetailEntry
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LogTarget) UnmarshalJSON(data []byte) (err error) {
	varLogTarget := _LogTarget{}

	err = json.Unmarshal(data, &varLogTarget)

	if err != nil {
		return err
	}

	*o = LogTarget(varLogTarget)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "alternateId")
		delete(additionalProperties, "changeDetails")
		delete(additionalProperties, "detailEntry")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "id")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLogTarget struct {
	value *LogTarget
	isSet bool
}

func (v NullableLogTarget) Get() *LogTarget {
	return v.value
}

func (v *NullableLogTarget) Set(val *LogTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableLogTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableLogTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogTarget(val *LogTarget) *NullableLogTarget {
	return &NullableLogTarget{value: val, isSet: true}
}

func (v NullableLogTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
