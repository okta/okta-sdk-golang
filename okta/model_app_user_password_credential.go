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

// AppUserPasswordCredential The user's password. This is a write-only property. An empty `password` object is returned to indicate that a password value exists.
type AppUserPasswordCredential struct {
	// Password value
	Value *string `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AppUserPasswordCredential AppUserPasswordCredential

// NewAppUserPasswordCredential instantiates a new AppUserPasswordCredential object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppUserPasswordCredential() *AppUserPasswordCredential {
	this := AppUserPasswordCredential{}
	return &this
}

// NewAppUserPasswordCredentialWithDefaults instantiates a new AppUserPasswordCredential object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppUserPasswordCredentialWithDefaults() *AppUserPasswordCredential {
	this := AppUserPasswordCredential{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *AppUserPasswordCredential) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppUserPasswordCredential) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *AppUserPasswordCredential) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *AppUserPasswordCredential) SetValue(v string) {
	o.Value = &v
}

func (o AppUserPasswordCredential) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AppUserPasswordCredential) UnmarshalJSON(bytes []byte) (err error) {
	varAppUserPasswordCredential := _AppUserPasswordCredential{}

	err = json.Unmarshal(bytes, &varAppUserPasswordCredential)
	if err == nil {
		*o = AppUserPasswordCredential(varAppUserPasswordCredential)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableAppUserPasswordCredential struct {
	value *AppUserPasswordCredential
	isSet bool
}

func (v NullableAppUserPasswordCredential) Get() *AppUserPasswordCredential {
	return v.value
}

func (v *NullableAppUserPasswordCredential) Set(val *AppUserPasswordCredential) {
	v.value = val
	v.isSet = true
}

func (v NullableAppUserPasswordCredential) IsSet() bool {
	return v.isSet
}

func (v *NullableAppUserPasswordCredential) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppUserPasswordCredential(val *AppUserPasswordCredential) *NullableAppUserPasswordCredential {
	return &NullableAppUserPasswordCredential{value: val, isSet: true}
}

func (v NullableAppUserPasswordCredential) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppUserPasswordCredential) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

