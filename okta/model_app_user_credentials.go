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

// AppUserCredentials Specifies a user's credentials for the app. This parameter can be omitted for apps with [sign-on mode](/openapi/okta-management/management/tag/Application/#tag/Application/operation/getApplication!c=200&path=0/signOnMode&t=response) (`signOnMode`) or [authentication schemes](/openapi/okta-management/management/tag/Application/#tag/Application/operation/getApplication!c=200&path=0/credentials/scheme&t=response) (`credentials.scheme`) that don't require credentials. 
type AppUserCredentials struct {
	Password *AppUserPasswordCredential `json:"password,omitempty"`
	// The user's username in the app
	UserName *string `json:"userName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AppUserCredentials AppUserCredentials

// NewAppUserCredentials instantiates a new AppUserCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppUserCredentials() *AppUserCredentials {
	this := AppUserCredentials{}
	return &this
}

// NewAppUserCredentialsWithDefaults instantiates a new AppUserCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppUserCredentialsWithDefaults() *AppUserCredentials {
	this := AppUserCredentials{}
	return &this
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *AppUserCredentials) GetPassword() AppUserPasswordCredential {
	if o == nil || o.Password == nil {
		var ret AppUserPasswordCredential
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppUserCredentials) GetPasswordOk() (*AppUserPasswordCredential, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *AppUserCredentials) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given AppUserPasswordCredential and assigns it to the Password field.
func (o *AppUserCredentials) SetPassword(v AppUserPasswordCredential) {
	o.Password = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *AppUserCredentials) GetUserName() string {
	if o == nil || o.UserName == nil {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppUserCredentials) GetUserNameOk() (*string, bool) {
	if o == nil || o.UserName == nil {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *AppUserCredentials) HasUserName() bool {
	if o != nil && o.UserName != nil {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *AppUserCredentials) SetUserName(v string) {
	o.UserName = &v
}

func (o AppUserCredentials) MarshalJSON() ([]byte, error) {
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

func (o *AppUserCredentials) UnmarshalJSON(bytes []byte) (err error) {
	varAppUserCredentials := _AppUserCredentials{}

	err = json.Unmarshal(bytes, &varAppUserCredentials)
	if err == nil {
		*o = AppUserCredentials(varAppUserCredentials)
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

type NullableAppUserCredentials struct {
	value *AppUserCredentials
	isSet bool
}

func (v NullableAppUserCredentials) Get() *AppUserCredentials {
	return v.value
}

func (v *NullableAppUserCredentials) Set(val *AppUserCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableAppUserCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableAppUserCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppUserCredentials(val *AppUserCredentials) *NullableAppUserCredentials {
	return &NullableAppUserCredentials{value: val, isSet: true}
}

func (v NullableAppUserCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppUserCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

