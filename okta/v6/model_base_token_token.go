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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// BaseTokenToken The token
type BaseTokenToken struct {
	Lifetime *BaseTokenTokenLifetime `json:"lifetime,omitempty"`
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
	if o == nil || o.Lifetime == nil {
		var ret BaseTokenTokenLifetime
		return ret
	}
	return *o.Lifetime
}

// GetLifetimeOk returns a tuple with the Lifetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseTokenToken) GetLifetimeOk() (*BaseTokenTokenLifetime, bool) {
	if o == nil || o.Lifetime == nil {
		return nil, false
	}
	return o.Lifetime, true
}

// HasLifetime returns a boolean if a field has been set.
func (o *BaseTokenToken) HasLifetime() bool {
	if o != nil && o.Lifetime != nil {
		return true
	}

	return false
}

// SetLifetime gets a reference to the given BaseTokenTokenLifetime and assigns it to the Lifetime field.
func (o *BaseTokenToken) SetLifetime(v BaseTokenTokenLifetime) {
	o.Lifetime = &v
}

func (o BaseTokenToken) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Lifetime != nil {
		toSerialize["lifetime"] = o.Lifetime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BaseTokenToken) UnmarshalJSON(bytes []byte) (err error) {
	varBaseTokenToken := _BaseTokenToken{}

	err = json.Unmarshal(bytes, &varBaseTokenToken)
	if err == nil {
		*o = BaseTokenToken(varBaseTokenToken)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "lifetime")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
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

