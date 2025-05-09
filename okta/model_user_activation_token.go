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

// UserActivationToken struct for UserActivationToken
type UserActivationToken struct {
	ActivationToken *string `json:"activationToken,omitempty"`
	ActivationUrl *string `json:"activationUrl,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserActivationToken UserActivationToken

// NewUserActivationToken instantiates a new UserActivationToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserActivationToken() *UserActivationToken {
	this := UserActivationToken{}
	return &this
}

// NewUserActivationTokenWithDefaults instantiates a new UserActivationToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserActivationTokenWithDefaults() *UserActivationToken {
	this := UserActivationToken{}
	return &this
}

// GetActivationToken returns the ActivationToken field value if set, zero value otherwise.
func (o *UserActivationToken) GetActivationToken() string {
	if o == nil || o.ActivationToken == nil {
		var ret string
		return ret
	}
	return *o.ActivationToken
}

// GetActivationTokenOk returns a tuple with the ActivationToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserActivationToken) GetActivationTokenOk() (*string, bool) {
	if o == nil || o.ActivationToken == nil {
		return nil, false
	}
	return o.ActivationToken, true
}

// HasActivationToken returns a boolean if a field has been set.
func (o *UserActivationToken) HasActivationToken() bool {
	if o != nil && o.ActivationToken != nil {
		return true
	}

	return false
}

// SetActivationToken gets a reference to the given string and assigns it to the ActivationToken field.
func (o *UserActivationToken) SetActivationToken(v string) {
	o.ActivationToken = &v
}

// GetActivationUrl returns the ActivationUrl field value if set, zero value otherwise.
func (o *UserActivationToken) GetActivationUrl() string {
	if o == nil || o.ActivationUrl == nil {
		var ret string
		return ret
	}
	return *o.ActivationUrl
}

// GetActivationUrlOk returns a tuple with the ActivationUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserActivationToken) GetActivationUrlOk() (*string, bool) {
	if o == nil || o.ActivationUrl == nil {
		return nil, false
	}
	return o.ActivationUrl, true
}

// HasActivationUrl returns a boolean if a field has been set.
func (o *UserActivationToken) HasActivationUrl() bool {
	if o != nil && o.ActivationUrl != nil {
		return true
	}

	return false
}

// SetActivationUrl gets a reference to the given string and assigns it to the ActivationUrl field.
func (o *UserActivationToken) SetActivationUrl(v string) {
	o.ActivationUrl = &v
}

func (o UserActivationToken) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ActivationToken != nil {
		toSerialize["activationToken"] = o.ActivationToken
	}
	if o.ActivationUrl != nil {
		toSerialize["activationUrl"] = o.ActivationUrl
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UserActivationToken) UnmarshalJSON(bytes []byte) (err error) {
	varUserActivationToken := _UserActivationToken{}

	err = json.Unmarshal(bytes, &varUserActivationToken)
	if err == nil {
		*o = UserActivationToken(varUserActivationToken)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "activationToken")
		delete(additionalProperties, "activationUrl")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableUserActivationToken struct {
	value *UserActivationToken
	isSet bool
}

func (v NullableUserActivationToken) Get() *UserActivationToken {
	return v.value
}

func (v *NullableUserActivationToken) Set(val *UserActivationToken) {
	v.value = val
	v.isSet = true
}

func (v NullableUserActivationToken) IsSet() bool {
	return v.isSet
}

func (v *NullableUserActivationToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserActivationToken(val *UserActivationToken) *NullableUserActivationToken {
	return &NullableUserActivationToken{value: val, isSet: true}
}

func (v NullableUserActivationToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserActivationToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

