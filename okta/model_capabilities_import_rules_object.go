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

// checks if the CapabilitiesImportRulesObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CapabilitiesImportRulesObject{}

// CapabilitiesImportRulesObject Defines user import rules
type CapabilitiesImportRulesObject struct {
	UserCreateAndMatch   *CapabilitiesImportRulesUserCreateAndMatchObject `json:"userCreateAndMatch,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CapabilitiesImportRulesObject CapabilitiesImportRulesObject

// NewCapabilitiesImportRulesObject instantiates a new CapabilitiesImportRulesObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilitiesImportRulesObject() *CapabilitiesImportRulesObject {
	this := CapabilitiesImportRulesObject{}
	return &this
}

// NewCapabilitiesImportRulesObjectWithDefaults instantiates a new CapabilitiesImportRulesObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilitiesImportRulesObjectWithDefaults() *CapabilitiesImportRulesObject {
	this := CapabilitiesImportRulesObject{}
	return &this
}

// GetUserCreateAndMatch returns the UserCreateAndMatch field value if set, zero value otherwise.
func (o *CapabilitiesImportRulesObject) GetUserCreateAndMatch() CapabilitiesImportRulesUserCreateAndMatchObject {
	if o == nil || IsNil(o.UserCreateAndMatch) {
		var ret CapabilitiesImportRulesUserCreateAndMatchObject
		return ret
	}
	return *o.UserCreateAndMatch
}

// GetUserCreateAndMatchOk returns a tuple with the UserCreateAndMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitiesImportRulesObject) GetUserCreateAndMatchOk() (*CapabilitiesImportRulesUserCreateAndMatchObject, bool) {
	if o == nil || IsNil(o.UserCreateAndMatch) {
		return nil, false
	}
	return o.UserCreateAndMatch, true
}

// HasUserCreateAndMatch returns a boolean if a field has been set.
func (o *CapabilitiesImportRulesObject) HasUserCreateAndMatch() bool {
	if o != nil && !IsNil(o.UserCreateAndMatch) {
		return true
	}

	return false
}

// SetUserCreateAndMatch gets a reference to the given CapabilitiesImportRulesUserCreateAndMatchObject and assigns it to the UserCreateAndMatch field.
func (o *CapabilitiesImportRulesObject) SetUserCreateAndMatch(v CapabilitiesImportRulesUserCreateAndMatchObject) {
	o.UserCreateAndMatch = &v
}

func (o CapabilitiesImportRulesObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CapabilitiesImportRulesObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserCreateAndMatch) {
		toSerialize["userCreateAndMatch"] = o.UserCreateAndMatch
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CapabilitiesImportRulesObject) UnmarshalJSON(data []byte) (err error) {
	varCapabilitiesImportRulesObject := _CapabilitiesImportRulesObject{}

	err = json.Unmarshal(data, &varCapabilitiesImportRulesObject)

	if err != nil {
		return err
	}

	*o = CapabilitiesImportRulesObject(varCapabilitiesImportRulesObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "userCreateAndMatch")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCapabilitiesImportRulesObject struct {
	value *CapabilitiesImportRulesObject
	isSet bool
}

func (v NullableCapabilitiesImportRulesObject) Get() *CapabilitiesImportRulesObject {
	return v.value
}

func (v *NullableCapabilitiesImportRulesObject) Set(val *CapabilitiesImportRulesObject) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilitiesImportRulesObject) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilitiesImportRulesObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilitiesImportRulesObject(val *CapabilitiesImportRulesObject) *NullableCapabilitiesImportRulesObject {
	return &NullableCapabilitiesImportRulesObject{value: val, isSet: true}
}

func (v NullableCapabilitiesImportRulesObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilitiesImportRulesObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
