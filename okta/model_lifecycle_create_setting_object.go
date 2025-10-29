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

// checks if the LifecycleCreateSettingObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LifecycleCreateSettingObject{}

// LifecycleCreateSettingObject Determines whether to update a user in the app when a user in Okta is updated
type LifecycleCreateSettingObject struct {
	// Setting status
	Status               *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LifecycleCreateSettingObject LifecycleCreateSettingObject

// NewLifecycleCreateSettingObject instantiates a new LifecycleCreateSettingObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLifecycleCreateSettingObject() *LifecycleCreateSettingObject {
	this := LifecycleCreateSettingObject{}
	var status string = "DISABLED"
	this.Status = &status
	return &this
}

// NewLifecycleCreateSettingObjectWithDefaults instantiates a new LifecycleCreateSettingObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLifecycleCreateSettingObjectWithDefaults() *LifecycleCreateSettingObject {
	this := LifecycleCreateSettingObject{}
	var status string = "DISABLED"
	this.Status = &status
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *LifecycleCreateSettingObject) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LifecycleCreateSettingObject) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *LifecycleCreateSettingObject) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *LifecycleCreateSettingObject) SetStatus(v string) {
	o.Status = &v
}

func (o LifecycleCreateSettingObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LifecycleCreateSettingObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LifecycleCreateSettingObject) UnmarshalJSON(data []byte) (err error) {
	varLifecycleCreateSettingObject := _LifecycleCreateSettingObject{}

	err = json.Unmarshal(data, &varLifecycleCreateSettingObject)

	if err != nil {
		return err
	}

	*o = LifecycleCreateSettingObject(varLifecycleCreateSettingObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLifecycleCreateSettingObject struct {
	value *LifecycleCreateSettingObject
	isSet bool
}

func (v NullableLifecycleCreateSettingObject) Get() *LifecycleCreateSettingObject {
	return v.value
}

func (v *NullableLifecycleCreateSettingObject) Set(val *LifecycleCreateSettingObject) {
	v.value = val
	v.isSet = true
}

func (v NullableLifecycleCreateSettingObject) IsSet() bool {
	return v.isSet
}

func (v *NullableLifecycleCreateSettingObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLifecycleCreateSettingObject(val *LifecycleCreateSettingObject) *NullableLifecycleCreateSettingObject {
	return &NullableLifecycleCreateSettingObject{value: val, isSet: true}
}

func (v NullableLifecycleCreateSettingObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLifecycleCreateSettingObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
