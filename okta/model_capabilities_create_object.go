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

// checks if the CapabilitiesCreateObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CapabilitiesCreateObject{}

// CapabilitiesCreateObject Determines whether Okta assigns a new app account to each user managed by Okta.  Okta doesn't create a new account if it detects that the username specified in Okta already exists in the app. The user's Okta username is assigned by default.
type CapabilitiesCreateObject struct {
	LifecycleCreate      *LifecycleCreateSettingObject `json:"lifecycleCreate,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CapabilitiesCreateObject CapabilitiesCreateObject

// NewCapabilitiesCreateObject instantiates a new CapabilitiesCreateObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilitiesCreateObject() *CapabilitiesCreateObject {
	this := CapabilitiesCreateObject{}
	return &this
}

// NewCapabilitiesCreateObjectWithDefaults instantiates a new CapabilitiesCreateObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilitiesCreateObjectWithDefaults() *CapabilitiesCreateObject {
	this := CapabilitiesCreateObject{}
	return &this
}

// GetLifecycleCreate returns the LifecycleCreate field value if set, zero value otherwise.
func (o *CapabilitiesCreateObject) GetLifecycleCreate() LifecycleCreateSettingObject {
	if o == nil || IsNil(o.LifecycleCreate) {
		var ret LifecycleCreateSettingObject
		return ret
	}
	return *o.LifecycleCreate
}

// GetLifecycleCreateOk returns a tuple with the LifecycleCreate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitiesCreateObject) GetLifecycleCreateOk() (*LifecycleCreateSettingObject, bool) {
	if o == nil || IsNil(o.LifecycleCreate) {
		return nil, false
	}
	return o.LifecycleCreate, true
}

// HasLifecycleCreate returns a boolean if a field has been set.
func (o *CapabilitiesCreateObject) HasLifecycleCreate() bool {
	if o != nil && !IsNil(o.LifecycleCreate) {
		return true
	}

	return false
}

// SetLifecycleCreate gets a reference to the given LifecycleCreateSettingObject and assigns it to the LifecycleCreate field.
func (o *CapabilitiesCreateObject) SetLifecycleCreate(v LifecycleCreateSettingObject) {
	o.LifecycleCreate = &v
}

func (o CapabilitiesCreateObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CapabilitiesCreateObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LifecycleCreate) {
		toSerialize["lifecycleCreate"] = o.LifecycleCreate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CapabilitiesCreateObject) UnmarshalJSON(data []byte) (err error) {
	varCapabilitiesCreateObject := _CapabilitiesCreateObject{}

	err = json.Unmarshal(data, &varCapabilitiesCreateObject)

	if err != nil {
		return err
	}

	*o = CapabilitiesCreateObject(varCapabilitiesCreateObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "lifecycleCreate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCapabilitiesCreateObject struct {
	value *CapabilitiesCreateObject
	isSet bool
}

func (v NullableCapabilitiesCreateObject) Get() *CapabilitiesCreateObject {
	return v.value
}

func (v *NullableCapabilitiesCreateObject) Set(val *CapabilitiesCreateObject) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilitiesCreateObject) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilitiesCreateObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilitiesCreateObject(val *CapabilitiesCreateObject) *NullableCapabilitiesCreateObject {
	return &NullableCapabilitiesCreateObject{value: val, isSet: true}
}

func (v NullableCapabilitiesCreateObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilitiesCreateObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
