/*
Okta Admin Management API

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

// checks if the MacOSAccountProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MacOSAccountProfile{}

// MacOSAccountProfile struct for MacOSAccountProfile
type MacOSAccountProfile struct {
	// Unique identifier for the macOS account
	AccountUUID *string `json:"accountUUID,omitempty"`
	// Full name of the account user
	FullName *string `json:"fullName,omitempty"`
	// Username of the account
	Username             *string `json:"username,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MacOSAccountProfile MacOSAccountProfile

// NewMacOSAccountProfile instantiates a new MacOSAccountProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMacOSAccountProfile() *MacOSAccountProfile {
	this := MacOSAccountProfile{}
	return &this
}

// NewMacOSAccountProfileWithDefaults instantiates a new MacOSAccountProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMacOSAccountProfileWithDefaults() *MacOSAccountProfile {
	this := MacOSAccountProfile{}
	return &this
}

// GetAccountUUID returns the AccountUUID field value if set, zero value otherwise.
func (o *MacOSAccountProfile) GetAccountUUID() string {
	if o == nil || IsNil(o.AccountUUID) {
		var ret string
		return ret
	}
	return *o.AccountUUID
}

// GetAccountUUIDOk returns a tuple with the AccountUUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MacOSAccountProfile) GetAccountUUIDOk() (*string, bool) {
	if o == nil || IsNil(o.AccountUUID) {
		return nil, false
	}
	return o.AccountUUID, true
}

// HasAccountUUID returns a boolean if a field has been set.
func (o *MacOSAccountProfile) HasAccountUUID() bool {
	if o != nil && !IsNil(o.AccountUUID) {
		return true
	}

	return false
}

// SetAccountUUID gets a reference to the given string and assigns it to the AccountUUID field.
func (o *MacOSAccountProfile) SetAccountUUID(v string) {
	o.AccountUUID = &v
}

// GetFullName returns the FullName field value if set, zero value otherwise.
func (o *MacOSAccountProfile) GetFullName() string {
	if o == nil || IsNil(o.FullName) {
		var ret string
		return ret
	}
	return *o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MacOSAccountProfile) GetFullNameOk() (*string, bool) {
	if o == nil || IsNil(o.FullName) {
		return nil, false
	}
	return o.FullName, true
}

// HasFullName returns a boolean if a field has been set.
func (o *MacOSAccountProfile) HasFullName() bool {
	if o != nil && !IsNil(o.FullName) {
		return true
	}

	return false
}

// SetFullName gets a reference to the given string and assigns it to the FullName field.
func (o *MacOSAccountProfile) SetFullName(v string) {
	o.FullName = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *MacOSAccountProfile) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MacOSAccountProfile) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *MacOSAccountProfile) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *MacOSAccountProfile) SetUsername(v string) {
	o.Username = &v
}

func (o MacOSAccountProfile) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MacOSAccountProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountUUID) {
		toSerialize["accountUUID"] = o.AccountUUID
	}
	if !IsNil(o.FullName) {
		toSerialize["fullName"] = o.FullName
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MacOSAccountProfile) UnmarshalJSON(data []byte) (err error) {
	varMacOSAccountProfile := _MacOSAccountProfile{}

	err = json.Unmarshal(data, &varMacOSAccountProfile)

	if err != nil {
		return err
	}

	*o = MacOSAccountProfile(varMacOSAccountProfile)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "accountUUID")
		delete(additionalProperties, "fullName")
		delete(additionalProperties, "username")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMacOSAccountProfile struct {
	value *MacOSAccountProfile
	isSet bool
}

func (v NullableMacOSAccountProfile) Get() *MacOSAccountProfile {
	return v.value
}

func (v *NullableMacOSAccountProfile) Set(val *MacOSAccountProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableMacOSAccountProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableMacOSAccountProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMacOSAccountProfile(val *MacOSAccountProfile) *NullableMacOSAccountProfile {
	return &NullableMacOSAccountProfile{value: val, isSet: true}
}

func (v NullableMacOSAccountProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMacOSAccountProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
