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

// checks if the AppUserCredentials type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppUserCredentials{}

// AppUserCredentials Specifies a user's credentials for the app. This parameter can be omitted for apps with [sign-on mode](/openapi/okta-management/management/tag/Application/#tag/Application/operation/getApplication!c=200&path=0/signOnMode&t=response) (`signOnMode`) or [authentication schemes](/openapi/okta-management/management/tag/Application/#tag/Application/operation/getApplication!c=200&path=0/credentials/scheme&t=response) (`credentials.scheme`) that don't require credentials.
type AppUserCredentials struct {
	Password *AppUserPasswordCredential `json:"password,omitempty"`
	// The user's username in the app  > **Note:** The [userNameTemplate](/openapi/okta-management/management/tag/Application/#tag/Application/operation/createApplication!path=0/credentials/userNameTemplate&t=request) in the application object defines the default username generated when a user is assigned to that app. > If you attempt to assign a username or password to an app with an incompatible [authentication scheme](/openapi/okta-management/management/tag/Application/#tag/Application/operation/createApplication!path=0/credentials/scheme&t=request), the following error is returned: > \"Credentials should not be set on this resource based on the scheme.\"
	UserName             *string `json:"userName,omitempty"`
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
	if o == nil || IsNil(o.Password) {
		var ret AppUserPasswordCredential
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppUserCredentials) GetPasswordOk() (*AppUserPasswordCredential, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *AppUserCredentials) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
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
	if o == nil || IsNil(o.UserName) {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppUserCredentials) GetUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.UserName) {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *AppUserCredentials) HasUserName() bool {
	if o != nil && !IsNil(o.UserName) {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *AppUserCredentials) SetUserName(v string) {
	o.UserName = &v
}

func (o AppUserCredentials) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppUserCredentials) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.UserName) {
		toSerialize["userName"] = o.UserName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AppUserCredentials) UnmarshalJSON(data []byte) (err error) {
	varAppUserCredentials := _AppUserCredentials{}

	err = json.Unmarshal(data, &varAppUserCredentials)

	if err != nil {
		return err
	}

	*o = AppUserCredentials(varAppUserCredentials)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "password")
		delete(additionalProperties, "userName")
		o.AdditionalProperties = additionalProperties
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
