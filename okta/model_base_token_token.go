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

// checks if the BaseTokenToken type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseTokenToken{}

// BaseTokenToken The token
type BaseTokenToken struct {
	Lifetime             *BaseTokenTokenLifetime `json:"lifetime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BaseTokenToken BaseTokenToken

// NewBaseTokenToken instantiates a new BaseTokenToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseTokenToken() *BaseTokenToken {
	this := BaseTokenToken{}
	return &this
}

// NewBaseTokenTokenWithDefaults instantiates a new BaseTokenToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseTokenTokenWithDefaults() *BaseTokenToken {
	this := BaseTokenToken{}
	return &this
}

// GetLifetime returns the Lifetime field value if set, zero value otherwise.
func (o *BaseTokenToken) GetLifetime() BaseTokenTokenLifetime {
	if o == nil || IsNil(o.Lifetime) {
		var ret BaseTokenTokenLifetime
		return ret
	}
	return *o.Lifetime
}

// GetLifetimeOk returns a tuple with the Lifetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseTokenToken) GetLifetimeOk() (*BaseTokenTokenLifetime, bool) {
	if o == nil || IsNil(o.Lifetime) {
		return nil, false
	}
	return o.Lifetime, true
}

// HasLifetime returns a boolean if a field has been set.
func (o *BaseTokenToken) HasLifetime() bool {
	if o != nil && !IsNil(o.Lifetime) {
		return true
	}

	return false
}

// SetLifetime gets a reference to the given BaseTokenTokenLifetime and assigns it to the Lifetime field.
func (o *BaseTokenToken) SetLifetime(v BaseTokenTokenLifetime) {
	o.Lifetime = &v
}

func (o BaseTokenToken) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaseTokenToken) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Lifetime) {
		toSerialize["lifetime"] = o.Lifetime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BaseTokenToken) UnmarshalJSON(data []byte) (err error) {
	varBaseTokenToken := _BaseTokenToken{}

	err = json.Unmarshal(data, &varBaseTokenToken)

	if err != nil {
		return err
	}

	*o = BaseTokenToken(varBaseTokenToken)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "lifetime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBaseTokenToken struct {
	value *BaseTokenToken
	isSet bool
}

func (v NullableBaseTokenToken) Get() *BaseTokenToken {
	return v.value
}

func (v *NullableBaseTokenToken) Set(val *BaseTokenToken) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseTokenToken) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseTokenToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseTokenToken(val *BaseTokenToken) *NullableBaseTokenToken {
	return &NullableBaseTokenToken{value: val, isSet: true}
}

func (v NullableBaseTokenToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseTokenToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
