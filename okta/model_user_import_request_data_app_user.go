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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the UserImportRequestDataAppUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserImportRequestDataAppUser{}

// UserImportRequestDataAppUser The app user profile being imported
type UserImportRequestDataAppUser struct {
	// Provides the name-value pairs of the attributes contained in the app user profile of the user who is being imported. You can change the values of attributes in the user's app profile by means of the `commands` object you return. If you change attributes in the app profile, they then flow through to the Okta user profile, based on matching and mapping rules.
	Profile              *map[string]string `json:"profile,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserImportRequestDataAppUser UserImportRequestDataAppUser

// NewUserImportRequestDataAppUser instantiates a new UserImportRequestDataAppUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserImportRequestDataAppUser() *UserImportRequestDataAppUser {
	this := UserImportRequestDataAppUser{}
	return &this
}

// NewUserImportRequestDataAppUserWithDefaults instantiates a new UserImportRequestDataAppUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserImportRequestDataAppUserWithDefaults() *UserImportRequestDataAppUser {
	this := UserImportRequestDataAppUser{}
	return &this
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *UserImportRequestDataAppUser) GetProfile() map[string]string {
	if o == nil || IsNil(o.Profile) {
		var ret map[string]string
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserImportRequestDataAppUser) GetProfileOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Profile) {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *UserImportRequestDataAppUser) HasProfile() bool {
	if o != nil && !IsNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given map[string]string and assigns it to the Profile field.
func (o *UserImportRequestDataAppUser) SetProfile(v map[string]string) {
	o.Profile = &v
}

func (o UserImportRequestDataAppUser) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserImportRequestDataAppUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Profile) {
		toSerialize["profile"] = o.Profile
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserImportRequestDataAppUser) UnmarshalJSON(data []byte) (err error) {
	varUserImportRequestDataAppUser := _UserImportRequestDataAppUser{}

	err = json.Unmarshal(data, &varUserImportRequestDataAppUser)

	if err != nil {
		return err
	}

	*o = UserImportRequestDataAppUser(varUserImportRequestDataAppUser)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "profile")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserImportRequestDataAppUser struct {
	value *UserImportRequestDataAppUser
	isSet bool
}

func (v NullableUserImportRequestDataAppUser) Get() *UserImportRequestDataAppUser {
	return v.value
}

func (v *NullableUserImportRequestDataAppUser) Set(val *UserImportRequestDataAppUser) {
	v.value = val
	v.isSet = true
}

func (v NullableUserImportRequestDataAppUser) IsSet() bool {
	return v.isSet
}

func (v *NullableUserImportRequestDataAppUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserImportRequestDataAppUser(val *UserImportRequestDataAppUser) *NullableUserImportRequestDataAppUser {
	return &NullableUserImportRequestDataAppUser{value: val, isSet: true}
}

func (v NullableUserImportRequestDataAppUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserImportRequestDataAppUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
