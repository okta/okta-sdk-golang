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

// checks if the UserType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserType{}

// UserType The user type that determines the schema for the user's profile. The `type` property is a map that identifies the [User Types](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/UserType/#tag/UserType).  Currently it contains a single element, `id`. It can be specified when creating a new user, and ca be updated by an admin on a full replace of an existing user (but not a partial update).
type UserType struct {
	// The ID of the user type
	Id                   *string `json:"id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserType UserType

// NewUserType instantiates a new UserType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserType() *UserType {
	this := UserType{}
	return &this
}

// NewUserTypeWithDefaults instantiates a new UserType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserTypeWithDefaults() *UserType {
	this := UserType{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UserType) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserType) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UserType) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UserType) SetId(v string) {
	o.Id = &v
}

func (o UserType) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserType) UnmarshalJSON(data []byte) (err error) {
	varUserType := _UserType{}

	err = json.Unmarshal(data, &varUserType)

	if err != nil {
		return err
	}

	*o = UserType(varUserType)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserType struct {
	value *UserType
	isSet bool
}

func (v NullableUserType) Get() *UserType {
	return v.value
}

func (v *NullableUserType) Set(val *UserType) {
	v.value = val
	v.isSet = true
}

func (v NullableUserType) IsSet() bool {
	return v.isSet
}

func (v *NullableUserType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserType(val *UserType) *NullableUserType {
	return &NullableUserType{value: val, isSet: true}
}

func (v NullableUserType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
