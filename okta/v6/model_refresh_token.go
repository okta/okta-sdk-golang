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

// RefreshToken The refresh token
type RefreshToken struct {
	// The refresh token ID
	Jti *string `json:"jti,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RefreshToken RefreshToken

// NewRefreshToken instantiates a new RefreshToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRefreshToken() *RefreshToken {
	this := RefreshToken{}
	return &this
}

// NewRefreshTokenWithDefaults instantiates a new RefreshToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRefreshTokenWithDefaults() *RefreshToken {
	this := RefreshToken{}
	return &this
}

// GetJti returns the Jti field value if set, zero value otherwise.
func (o *RefreshToken) GetJti() string {
	if o == nil || o.Jti == nil {
		var ret string
		return ret
	}
	return *o.Jti
}

// GetJtiOk returns a tuple with the Jti field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefreshToken) GetJtiOk() (*string, bool) {
	if o == nil || o.Jti == nil {
		return nil, false
	}
	return o.Jti, true
}

// HasJti returns a boolean if a field has been set.
func (o *RefreshToken) HasJti() bool {
	if o != nil && o.Jti != nil {
		return true
	}

	return false
}

// SetJti gets a reference to the given string and assigns it to the Jti field.
func (o *RefreshToken) SetJti(v string) {
	o.Jti = &v
}

func (o RefreshToken) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Jti != nil {
		toSerialize["jti"] = o.Jti
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RefreshToken) UnmarshalJSON(bytes []byte) (err error) {
	varRefreshToken := _RefreshToken{}

	err = json.Unmarshal(bytes, &varRefreshToken)
	if err == nil {
		*o = RefreshToken(varRefreshToken)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "jti")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableRefreshToken struct {
	value *RefreshToken
	isSet bool
}

func (v NullableRefreshToken) Get() *RefreshToken {
	return v.value
}

func (v *NullableRefreshToken) Set(val *RefreshToken) {
	v.value = val
	v.isSet = true
}

func (v NullableRefreshToken) IsSet() bool {
	return v.isSet
}

func (v *NullableRefreshToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefreshToken(val *RefreshToken) *NullableRefreshToken {
	return &NullableRefreshToken{value: val, isSet: true}
}

func (v NullableRefreshToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRefreshToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

