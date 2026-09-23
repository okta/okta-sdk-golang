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

// checks if the ProximityProviderListInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProximityProviderListInner{}

// ProximityProviderListInner struct for ProximityProviderListInner
type ProximityProviderListInner struct {
	// The proximity provider types that are supported in your org
	Type                 *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProximityProviderListInner ProximityProviderListInner

// NewProximityProviderListInner instantiates a new ProximityProviderListInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProximityProviderListInner() *ProximityProviderListInner {
	this := ProximityProviderListInner{}
	return &this
}

// NewProximityProviderListInnerWithDefaults instantiates a new ProximityProviderListInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProximityProviderListInnerWithDefaults() *ProximityProviderListInner {
	this := ProximityProviderListInner{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ProximityProviderListInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProximityProviderListInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ProximityProviderListInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ProximityProviderListInner) SetType(v string) {
	o.Type = &v
}

func (o ProximityProviderListInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProximityProviderListInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProximityProviderListInner) UnmarshalJSON(data []byte) (err error) {
	varProximityProviderListInner := _ProximityProviderListInner{}

	err = json.Unmarshal(data, &varProximityProviderListInner)

	if err != nil {
		return err
	}

	*o = ProximityProviderListInner(varProximityProviderListInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProximityProviderListInner struct {
	value *ProximityProviderListInner
	isSet bool
}

func (v NullableProximityProviderListInner) Get() *ProximityProviderListInner {
	return v.value
}

func (v *NullableProximityProviderListInner) Set(val *ProximityProviderListInner) {
	v.value = val
	v.isSet = true
}

func (v NullableProximityProviderListInner) IsSet() bool {
	return v.isSet
}

func (v *NullableProximityProviderListInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProximityProviderListInner(val *ProximityProviderListInner) *NullableProximityProviderListInner {
	return &NullableProximityProviderListInner{value: val, isSet: true}
}

func (v NullableProximityProviderListInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProximityProviderListInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
