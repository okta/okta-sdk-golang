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
)

// checks if the UserImportRequestDataUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserImportRequestDataUser{}

// UserImportRequestDataUser Provides information on the Okta user profile currently set to be used for the user who is being imported, based on the matching rules and attribute mappings that were applied.
type UserImportRequestDataUser struct {
	// The `data.user.profile` contains the name-value pairs of the attributes in the user profile. If the user has been matched to an existing Okta user, a `data.user.id` object is included, containing the unique identifier of the Okta user profile.  You can change the values of the attributes by means of the `commands` object you return.
	Profile              *map[string]string `json:"profile,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserImportRequestDataUser UserImportRequestDataUser

// NewUserImportRequestDataUser instantiates a new UserImportRequestDataUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserImportRequestDataUser() *UserImportRequestDataUser {
	this := UserImportRequestDataUser{}
	return &this
}

// NewUserImportRequestDataUserWithDefaults instantiates a new UserImportRequestDataUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserImportRequestDataUserWithDefaults() *UserImportRequestDataUser {
	this := UserImportRequestDataUser{}
	return &this
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *UserImportRequestDataUser) GetProfile() map[string]string {
	if o == nil || IsNil(o.Profile) {
		var ret map[string]string
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserImportRequestDataUser) GetProfileOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Profile) {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *UserImportRequestDataUser) HasProfile() bool {
	if o != nil && !IsNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given map[string]string and assigns it to the Profile field.
func (o *UserImportRequestDataUser) SetProfile(v map[string]string) {
	o.Profile = &v
}

func (o UserImportRequestDataUser) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserImportRequestDataUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Profile) {
		toSerialize["profile"] = o.Profile
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserImportRequestDataUser) UnmarshalJSON(data []byte) (err error) {
	varUserImportRequestDataUser := _UserImportRequestDataUser{}

	err = json.Unmarshal(data, &varUserImportRequestDataUser)

	if err != nil {
		return err
	}

	*o = UserImportRequestDataUser(varUserImportRequestDataUser)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "profile")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserImportRequestDataUser struct {
	value *UserImportRequestDataUser
	isSet bool
}

func (v NullableUserImportRequestDataUser) Get() *UserImportRequestDataUser {
	return v.value
}

func (v *NullableUserImportRequestDataUser) Set(val *UserImportRequestDataUser) {
	v.value = val
	v.isSet = true
}

func (v NullableUserImportRequestDataUser) IsSet() bool {
	return v.isSet
}

func (v *NullableUserImportRequestDataUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserImportRequestDataUser(val *UserImportRequestDataUser) *NullableUserImportRequestDataUser {
	return &NullableUserImportRequestDataUser{value: val, isSet: true}
}

func (v NullableUserImportRequestDataUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserImportRequestDataUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
