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

// CapabilitiesObject Defines the configurations for the USER_PROVISIONING feature
type CapabilitiesObject struct {
	Create *CapabilitiesCreateObject `json:"create,omitempty"`
	Update *CapabilitiesUpdateObject `json:"update,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CapabilitiesObject CapabilitiesObject

// NewCapabilitiesObject instantiates a new CapabilitiesObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilitiesObject() *CapabilitiesObject {
	this := CapabilitiesObject{}
	return &this
}

// NewCapabilitiesObjectWithDefaults instantiates a new CapabilitiesObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilitiesObjectWithDefaults() *CapabilitiesObject {
	this := CapabilitiesObject{}
	return &this
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *CapabilitiesObject) GetCreate() CapabilitiesCreateObject {
	if o == nil || o.Create == nil {
		var ret CapabilitiesCreateObject
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitiesObject) GetCreateOk() (*CapabilitiesCreateObject, bool) {
	if o == nil || o.Create == nil {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *CapabilitiesObject) HasCreate() bool {
	if o != nil && o.Create != nil {
		return true
	}

	return false
}

// SetCreate gets a reference to the given CapabilitiesCreateObject and assigns it to the Create field.
func (o *CapabilitiesObject) SetCreate(v CapabilitiesCreateObject) {
	o.Create = &v
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *CapabilitiesObject) GetUpdate() CapabilitiesUpdateObject {
	if o == nil || o.Update == nil {
		var ret CapabilitiesUpdateObject
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitiesObject) GetUpdateOk() (*CapabilitiesUpdateObject, bool) {
	if o == nil || o.Update == nil {
		return nil, false
	}
	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *CapabilitiesObject) HasUpdate() bool {
	if o != nil && o.Update != nil {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given CapabilitiesUpdateObject and assigns it to the Update field.
func (o *CapabilitiesObject) SetUpdate(v CapabilitiesUpdateObject) {
	o.Update = &v
}

func (o CapabilitiesObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Create != nil {
		toSerialize["create"] = o.Create
	}
	if o.Update != nil {
		toSerialize["update"] = o.Update
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CapabilitiesObject) UnmarshalJSON(bytes []byte) (err error) {
	varCapabilitiesObject := _CapabilitiesObject{}

	err = json.Unmarshal(bytes, &varCapabilitiesObject)
	if err == nil {
		*o = CapabilitiesObject(varCapabilitiesObject)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "create")
		delete(additionalProperties, "update")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableCapabilitiesObject struct {
	value *CapabilitiesObject
	isSet bool
}

func (v NullableCapabilitiesObject) Get() *CapabilitiesObject {
	return v.value
}

func (v *NullableCapabilitiesObject) Set(val *CapabilitiesObject) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilitiesObject) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilitiesObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilitiesObject(val *CapabilitiesObject) *NullableCapabilitiesObject {
	return &NullableCapabilitiesObject{value: val, isSet: true}
}

func (v NullableCapabilitiesObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilitiesObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

