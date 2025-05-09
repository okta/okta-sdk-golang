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

// PrivilegedResourceCredentials Credentials for the privileged account
type PrivilegedResourceCredentials struct {
	// The password associated with the privileged resource
	Password *string `json:"password,omitempty"`
	// The username associated with the privileged resource
	UserName *string `json:"userName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PrivilegedResourceCredentials PrivilegedResourceCredentials

// NewPrivilegedResourceCredentials instantiates a new PrivilegedResourceCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrivilegedResourceCredentials() *PrivilegedResourceCredentials {
	this := PrivilegedResourceCredentials{}
	return &this
}

// NewPrivilegedResourceCredentialsWithDefaults instantiates a new PrivilegedResourceCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrivilegedResourceCredentialsWithDefaults() *PrivilegedResourceCredentials {
	this := PrivilegedResourceCredentials{}
	return &this
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *PrivilegedResourceCredentials) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivilegedResourceCredentials) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *PrivilegedResourceCredentials) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *PrivilegedResourceCredentials) SetPassword(v string) {
	o.Password = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *PrivilegedResourceCredentials) GetUserName() string {
	if o == nil || o.UserName == nil {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivilegedResourceCredentials) GetUserNameOk() (*string, bool) {
	if o == nil || o.UserName == nil {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *PrivilegedResourceCredentials) HasUserName() bool {
	if o != nil && o.UserName != nil {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *PrivilegedResourceCredentials) SetUserName(v string) {
	o.UserName = &v
}

func (o PrivilegedResourceCredentials) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.UserName != nil {
		toSerialize["userName"] = o.UserName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PrivilegedResourceCredentials) UnmarshalJSON(bytes []byte) (err error) {
	varPrivilegedResourceCredentials := _PrivilegedResourceCredentials{}

	err = json.Unmarshal(bytes, &varPrivilegedResourceCredentials)
	if err == nil {
		*o = PrivilegedResourceCredentials(varPrivilegedResourceCredentials)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "password")
		delete(additionalProperties, "userName")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullablePrivilegedResourceCredentials struct {
	value *PrivilegedResourceCredentials
	isSet bool
}

func (v NullablePrivilegedResourceCredentials) Get() *PrivilegedResourceCredentials {
	return v.value
}

func (v *NullablePrivilegedResourceCredentials) Set(val *PrivilegedResourceCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullablePrivilegedResourceCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullablePrivilegedResourceCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrivilegedResourceCredentials(val *PrivilegedResourceCredentials) *NullablePrivilegedResourceCredentials {
	return &NullablePrivilegedResourceCredentials{value: val, isSet: true}
}

func (v NullablePrivilegedResourceCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrivilegedResourceCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

