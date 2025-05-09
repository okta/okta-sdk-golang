/*
Okta Admin Management

Allows customers to easily access the Okta Management APIs

Copyright 2018 - Present Okta, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

API version: 2024.06.1
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// CapabilitiesCreateObject Determines whether Okta assigns a new application account to each user managed by Okta.  Okta doesn't create a new account if it detects that the username specified in Okta already exists in the application. The user's Okta username is assigned by default. 
type CapabilitiesCreateObject struct {
	LifecycleCreate *LifecycleCreateSettingObject `json:"lifecycleCreate,omitempty"`
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
	if o == nil || o.LifecycleCreate == nil {
		var ret LifecycleCreateSettingObject
		return ret
	}
	return *o.LifecycleCreate
}

// GetLifecycleCreateOk returns a tuple with the LifecycleCreate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitiesCreateObject) GetLifecycleCreateOk() (*LifecycleCreateSettingObject, bool) {
	if o == nil || o.LifecycleCreate == nil {
		return nil, false
	}
	return o.LifecycleCreate, true
}

// HasLifecycleCreate returns a boolean if a field has been set.
func (o *CapabilitiesCreateObject) HasLifecycleCreate() bool {
	if o != nil && o.LifecycleCreate != nil {
		return true
	}

	return false
}

// SetLifecycleCreate gets a reference to the given LifecycleCreateSettingObject and assigns it to the LifecycleCreate field.
func (o *CapabilitiesCreateObject) SetLifecycleCreate(v LifecycleCreateSettingObject) {
	o.LifecycleCreate = &v
}

func (o CapabilitiesCreateObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LifecycleCreate != nil {
		toSerialize["lifecycleCreate"] = o.LifecycleCreate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CapabilitiesCreateObject) UnmarshalJSON(bytes []byte) (err error) {
	varCapabilitiesCreateObject := _CapabilitiesCreateObject{}

	err = json.Unmarshal(bytes, &varCapabilitiesCreateObject)
	if err == nil {
		*o = CapabilitiesCreateObject(varCapabilitiesCreateObject)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "lifecycleCreate")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
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

