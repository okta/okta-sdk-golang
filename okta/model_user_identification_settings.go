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

// checks if the UserIdentificationSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserIdentificationSettings{}

// UserIdentificationSettings <x-lifecycle-container><x-lifecycle class=\"oie\"></x-lifecycle></x-lifecycle-container>Specifies the user identification settings available to users when signing in
type UserIdentificationSettings struct {
	SecurityMethods      *UserIdentificationSecurityMethods `json:"securityMethods,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserIdentificationSettings UserIdentificationSettings

// NewUserIdentificationSettings instantiates a new UserIdentificationSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserIdentificationSettings() *UserIdentificationSettings {
	this := UserIdentificationSettings{}
	return &this
}

// NewUserIdentificationSettingsWithDefaults instantiates a new UserIdentificationSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserIdentificationSettingsWithDefaults() *UserIdentificationSettings {
	this := UserIdentificationSettings{}
	return &this
}

// GetSecurityMethods returns the SecurityMethods field value if set, zero value otherwise.
func (o *UserIdentificationSettings) GetSecurityMethods() UserIdentificationSecurityMethods {
	if o == nil || IsNil(o.SecurityMethods) {
		var ret UserIdentificationSecurityMethods
		return ret
	}
	return *o.SecurityMethods
}

// GetSecurityMethodsOk returns a tuple with the SecurityMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserIdentificationSettings) GetSecurityMethodsOk() (*UserIdentificationSecurityMethods, bool) {
	if o == nil || IsNil(o.SecurityMethods) {
		return nil, false
	}
	return o.SecurityMethods, true
}

// HasSecurityMethods returns a boolean if a field has been set.
func (o *UserIdentificationSettings) HasSecurityMethods() bool {
	if o != nil && !IsNil(o.SecurityMethods) {
		return true
	}

	return false
}

// SetSecurityMethods gets a reference to the given UserIdentificationSecurityMethods and assigns it to the SecurityMethods field.
func (o *UserIdentificationSettings) SetSecurityMethods(v UserIdentificationSecurityMethods) {
	o.SecurityMethods = &v
}

func (o UserIdentificationSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserIdentificationSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SecurityMethods) {
		toSerialize["securityMethods"] = o.SecurityMethods
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserIdentificationSettings) UnmarshalJSON(data []byte) (err error) {
	varUserIdentificationSettings := _UserIdentificationSettings{}

	err = json.Unmarshal(data, &varUserIdentificationSettings)

	if err != nil {
		return err
	}

	*o = UserIdentificationSettings(varUserIdentificationSettings)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "securityMethods")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserIdentificationSettings struct {
	value *UserIdentificationSettings
	isSet bool
}

func (v NullableUserIdentificationSettings) Get() *UserIdentificationSettings {
	return v.value
}

func (v *NullableUserIdentificationSettings) Set(val *UserIdentificationSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableUserIdentificationSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableUserIdentificationSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserIdentificationSettings(val *UserIdentificationSettings) *NullableUserIdentificationSettings {
	return &NullableUserIdentificationSettings{value: val, isSet: true}
}

func (v NullableUserIdentificationSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserIdentificationSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
