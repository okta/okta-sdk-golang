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

// checks if the JwkUse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JwkUse{}

// JwkUse struct for JwkUse
type JwkUse struct {
	// Purpose of the certificate. The only supported value is `sig`.
	Use                  *string `json:"use,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _JwkUse JwkUse

// NewJwkUse instantiates a new JwkUse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJwkUse() *JwkUse {
	this := JwkUse{}
	return &this
}

// NewJwkUseWithDefaults instantiates a new JwkUse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJwkUseWithDefaults() *JwkUse {
	this := JwkUse{}
	return &this
}

// GetUse returns the Use field value if set, zero value otherwise.
func (o *JwkUse) GetUse() string {
	if o == nil || IsNil(o.Use) {
		var ret string
		return ret
	}
	return *o.Use
}

// GetUseOk returns a tuple with the Use field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JwkUse) GetUseOk() (*string, bool) {
	if o == nil || IsNil(o.Use) {
		return nil, false
	}
	return o.Use, true
}

// HasUse returns a boolean if a field has been set.
func (o *JwkUse) HasUse() bool {
	if o != nil && !IsNil(o.Use) {
		return true
	}

	return false
}

// SetUse gets a reference to the given string and assigns it to the Use field.
func (o *JwkUse) SetUse(v string) {
	o.Use = &v
}

func (o JwkUse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JwkUse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Use) {
		toSerialize["use"] = o.Use
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *JwkUse) UnmarshalJSON(data []byte) (err error) {
	varJwkUse := _JwkUse{}

	err = json.Unmarshal(data, &varJwkUse)

	if err != nil {
		return err
	}

	*o = JwkUse(varJwkUse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "use")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableJwkUse struct {
	value *JwkUse
	isSet bool
}

func (v NullableJwkUse) Get() *JwkUse {
	return v.value
}

func (v *NullableJwkUse) Set(val *JwkUse) {
	v.value = val
	v.isSet = true
}

func (v NullableJwkUse) IsSet() bool {
	return v.isSet
}

func (v *NullableJwkUse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJwkUse(val *JwkUse) *NullableJwkUse {
	return &NullableJwkUse{value: val, isSet: true}
}

func (v NullableJwkUse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJwkUse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
