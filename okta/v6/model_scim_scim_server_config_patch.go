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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the ScimScimServerConfigPatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScimScimServerConfigPatch{}

// ScimScimServerConfigPatch PATCH operation options
type ScimScimServerConfigPatch struct {
	// Specifies if the PATCH operation is supported
	Supported            *bool `json:"supported,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ScimScimServerConfigPatch ScimScimServerConfigPatch

// NewScimScimServerConfigPatch instantiates a new ScimScimServerConfigPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScimScimServerConfigPatch() *ScimScimServerConfigPatch {
	this := ScimScimServerConfigPatch{}
	var supported bool = false
	this.Supported = &supported
	return &this
}

// NewScimScimServerConfigPatchWithDefaults instantiates a new ScimScimServerConfigPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScimScimServerConfigPatchWithDefaults() *ScimScimServerConfigPatch {
	this := ScimScimServerConfigPatch{}
	var supported bool = false
	this.Supported = &supported
	return &this
}

// GetSupported returns the Supported field value if set, zero value otherwise.
func (o *ScimScimServerConfigPatch) GetSupported() bool {
	if o == nil || IsNil(o.Supported) {
		var ret bool
		return ret
	}
	return *o.Supported
}

// GetSupportedOk returns a tuple with the Supported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimScimServerConfigPatch) GetSupportedOk() (*bool, bool) {
	if o == nil || IsNil(o.Supported) {
		return nil, false
	}
	return o.Supported, true
}

// HasSupported returns a boolean if a field has been set.
func (o *ScimScimServerConfigPatch) HasSupported() bool {
	if o != nil && !IsNil(o.Supported) {
		return true
	}

	return false
}

// SetSupported gets a reference to the given bool and assigns it to the Supported field.
func (o *ScimScimServerConfigPatch) SetSupported(v bool) {
	o.Supported = &v
}

func (o ScimScimServerConfigPatch) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScimScimServerConfigPatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Supported) {
		toSerialize["supported"] = o.Supported
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ScimScimServerConfigPatch) UnmarshalJSON(data []byte) (err error) {
	varScimScimServerConfigPatch := _ScimScimServerConfigPatch{}

	err = json.Unmarshal(data, &varScimScimServerConfigPatch)

	if err != nil {
		return err
	}

	*o = ScimScimServerConfigPatch(varScimScimServerConfigPatch)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "supported")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableScimScimServerConfigPatch struct {
	value *ScimScimServerConfigPatch
	isSet bool
}

func (v NullableScimScimServerConfigPatch) Get() *ScimScimServerConfigPatch {
	return v.value
}

func (v *NullableScimScimServerConfigPatch) Set(val *ScimScimServerConfigPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableScimScimServerConfigPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableScimScimServerConfigPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScimScimServerConfigPatch(val *ScimScimServerConfigPatch) *NullableScimScimServerConfigPatch {
	return &NullableScimScimServerConfigPatch{value: val, isSet: true}
}

func (v NullableScimScimServerConfigPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScimScimServerConfigPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
