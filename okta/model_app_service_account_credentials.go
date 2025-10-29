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

API version: 2025.10.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"fmt"
)

// checks if the AppServiceAccountCredentials type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppServiceAccountCredentials{}

// AppServiceAccountCredentials Credentials for a SaaS app account
type AppServiceAccountCredentials struct {
	// The password associated with the service account
	Password *string `json:"password,omitempty"`
	// The username associated with the service account
	Username             string `json:"username"`
	AdditionalProperties map[string]interface{}
}

type _AppServiceAccountCredentials AppServiceAccountCredentials

// NewAppServiceAccountCredentials instantiates a new AppServiceAccountCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppServiceAccountCredentials(username string) *AppServiceAccountCredentials {
	this := AppServiceAccountCredentials{}
	this.Username = username
	return &this
}

// NewAppServiceAccountCredentialsWithDefaults instantiates a new AppServiceAccountCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppServiceAccountCredentialsWithDefaults() *AppServiceAccountCredentials {
	this := AppServiceAccountCredentials{}
	return &this
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *AppServiceAccountCredentials) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppServiceAccountCredentials) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *AppServiceAccountCredentials) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *AppServiceAccountCredentials) SetPassword(v string) {
	o.Password = &v
}

// GetUsername returns the Username field value
func (o *AppServiceAccountCredentials) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *AppServiceAccountCredentials) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *AppServiceAccountCredentials) SetUsername(v string) {
	o.Username = v
}

func (o AppServiceAccountCredentials) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppServiceAccountCredentials) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	toSerialize["username"] = o.Username

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AppServiceAccountCredentials) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"username",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAppServiceAccountCredentials := _AppServiceAccountCredentials{}

	err = json.Unmarshal(data, &varAppServiceAccountCredentials)

	if err != nil {
		return err
	}

	*o = AppServiceAccountCredentials(varAppServiceAccountCredentials)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "password")
		delete(additionalProperties, "username")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppServiceAccountCredentials struct {
	value *AppServiceAccountCredentials
	isSet bool
}

func (v NullableAppServiceAccountCredentials) Get() *AppServiceAccountCredentials {
	return v.value
}

func (v *NullableAppServiceAccountCredentials) Set(val *AppServiceAccountCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableAppServiceAccountCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableAppServiceAccountCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppServiceAccountCredentials(val *AppServiceAccountCredentials) *NullableAppServiceAccountCredentials {
	return &NullableAppServiceAccountCredentials{value: val, isSet: true}
}

func (v NullableAppServiceAccountCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppServiceAccountCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
