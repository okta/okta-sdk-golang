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

// checks if the BaseTokenTokenLifetime type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseTokenTokenLifetime{}

// BaseTokenTokenLifetime Lifetime of the token
type BaseTokenTokenLifetime struct {
	// Time in seconds until the token expires
	Expiration           *int32 `json:"expiration,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BaseTokenTokenLifetime BaseTokenTokenLifetime

// NewBaseTokenTokenLifetime instantiates a new BaseTokenTokenLifetime object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseTokenTokenLifetime() *BaseTokenTokenLifetime {
	this := BaseTokenTokenLifetime{}
	return &this
}

// NewBaseTokenTokenLifetimeWithDefaults instantiates a new BaseTokenTokenLifetime object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseTokenTokenLifetimeWithDefaults() *BaseTokenTokenLifetime {
	this := BaseTokenTokenLifetime{}
	return &this
}

// GetExpiration returns the Expiration field value if set, zero value otherwise.
func (o *BaseTokenTokenLifetime) GetExpiration() int32 {
	if o == nil || IsNil(o.Expiration) {
		var ret int32
		return ret
	}
	return *o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseTokenTokenLifetime) GetExpirationOk() (*int32, bool) {
	if o == nil || IsNil(o.Expiration) {
		return nil, false
	}
	return o.Expiration, true
}

// HasExpiration returns a boolean if a field has been set.
func (o *BaseTokenTokenLifetime) HasExpiration() bool {
	if o != nil && !IsNil(o.Expiration) {
		return true
	}

	return false
}

// SetExpiration gets a reference to the given int32 and assigns it to the Expiration field.
func (o *BaseTokenTokenLifetime) SetExpiration(v int32) {
	o.Expiration = &v
}

func (o BaseTokenTokenLifetime) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaseTokenTokenLifetime) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Expiration) {
		toSerialize["expiration"] = o.Expiration
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BaseTokenTokenLifetime) UnmarshalJSON(data []byte) (err error) {
	varBaseTokenTokenLifetime := _BaseTokenTokenLifetime{}

	err = json.Unmarshal(data, &varBaseTokenTokenLifetime)

	if err != nil {
		return err
	}

	*o = BaseTokenTokenLifetime(varBaseTokenTokenLifetime)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "expiration")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBaseTokenTokenLifetime struct {
	value *BaseTokenTokenLifetime
	isSet bool
}

func (v NullableBaseTokenTokenLifetime) Get() *BaseTokenTokenLifetime {
	return v.value
}

func (v *NullableBaseTokenTokenLifetime) Set(val *BaseTokenTokenLifetime) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseTokenTokenLifetime) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseTokenTokenLifetime) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseTokenTokenLifetime(val *BaseTokenTokenLifetime) *NullableBaseTokenTokenLifetime {
	return &NullableBaseTokenTokenLifetime{value: val, isSet: true}
}

func (v NullableBaseTokenTokenLifetime) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseTokenTokenLifetime) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
