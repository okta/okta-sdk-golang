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

// PasswordImportRequestDataContextCredential struct for PasswordImportRequestDataContextCredential
type PasswordImportRequestDataContextCredential struct {
	// The `username` that the end user supplied when attempting to sign in to Okta.
	Username *string `json:"username,omitempty"`
	// The `password` that the end user supplied when attempting to sign in to Okta.
	Password *string `json:"password,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PasswordImportRequestDataContextCredential PasswordImportRequestDataContextCredential

// NewPasswordImportRequestDataContextCredential instantiates a new PasswordImportRequestDataContextCredential object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordImportRequestDataContextCredential() *PasswordImportRequestDataContextCredential {
	this := PasswordImportRequestDataContextCredential{}
	return &this
}

// NewPasswordImportRequestDataContextCredentialWithDefaults instantiates a new PasswordImportRequestDataContextCredential object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordImportRequestDataContextCredentialWithDefaults() *PasswordImportRequestDataContextCredential {
	this := PasswordImportRequestDataContextCredential{}
	return &this
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *PasswordImportRequestDataContextCredential) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordImportRequestDataContextCredential) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *PasswordImportRequestDataContextCredential) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *PasswordImportRequestDataContextCredential) SetUsername(v string) {
	o.Username = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *PasswordImportRequestDataContextCredential) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordImportRequestDataContextCredential) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *PasswordImportRequestDataContextCredential) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *PasswordImportRequestDataContextCredential) SetPassword(v string) {
	o.Password = &v
}

func (o PasswordImportRequestDataContextCredential) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PasswordImportRequestDataContextCredential) UnmarshalJSON(bytes []byte) (err error) {
	varPasswordImportRequestDataContextCredential := _PasswordImportRequestDataContextCredential{}

	err = json.Unmarshal(bytes, &varPasswordImportRequestDataContextCredential)
	if err == nil {
		*o = PasswordImportRequestDataContextCredential(varPasswordImportRequestDataContextCredential)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "username")
		delete(additionalProperties, "password")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullablePasswordImportRequestDataContextCredential struct {
	value *PasswordImportRequestDataContextCredential
	isSet bool
}

func (v NullablePasswordImportRequestDataContextCredential) Get() *PasswordImportRequestDataContextCredential {
	return v.value
}

func (v *NullablePasswordImportRequestDataContextCredential) Set(val *PasswordImportRequestDataContextCredential) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordImportRequestDataContextCredential) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordImportRequestDataContextCredential) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordImportRequestDataContextCredential(val *PasswordImportRequestDataContextCredential) *NullablePasswordImportRequestDataContextCredential {
	return &NullablePasswordImportRequestDataContextCredential{value: val, isSet: true}
}

func (v NullablePasswordImportRequestDataContextCredential) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordImportRequestDataContextCredential) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

