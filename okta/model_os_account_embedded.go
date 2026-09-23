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

// checks if the OSAccountEmbedded type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OSAccountEmbedded{}

// OSAccountEmbedded Embedded resources related to the OS account
type OSAccountEmbedded struct {
	// Enrollments linked to this OS account
	AccountLinkedEnrollments []AccountLinkedEnrollment `json:"accountLinkedEnrollments,omitempty"`
	// Users associated with this OS account
	Users                []User `json:"users,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OSAccountEmbedded OSAccountEmbedded

// NewOSAccountEmbedded instantiates a new OSAccountEmbedded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOSAccountEmbedded() *OSAccountEmbedded {
	this := OSAccountEmbedded{}
	return &this
}

// NewOSAccountEmbeddedWithDefaults instantiates a new OSAccountEmbedded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOSAccountEmbeddedWithDefaults() *OSAccountEmbedded {
	this := OSAccountEmbedded{}
	return &this
}

// GetAccountLinkedEnrollments returns the AccountLinkedEnrollments field value if set, zero value otherwise.
func (o *OSAccountEmbedded) GetAccountLinkedEnrollments() []AccountLinkedEnrollment {
	if o == nil || IsNil(o.AccountLinkedEnrollments) {
		var ret []AccountLinkedEnrollment
		return ret
	}
	return o.AccountLinkedEnrollments
}

// GetAccountLinkedEnrollmentsOk returns a tuple with the AccountLinkedEnrollments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSAccountEmbedded) GetAccountLinkedEnrollmentsOk() ([]AccountLinkedEnrollment, bool) {
	if o == nil || IsNil(o.AccountLinkedEnrollments) {
		return nil, false
	}
	return o.AccountLinkedEnrollments, true
}

// HasAccountLinkedEnrollments returns a boolean if a field has been set.
func (o *OSAccountEmbedded) HasAccountLinkedEnrollments() bool {
	if o != nil && !IsNil(o.AccountLinkedEnrollments) {
		return true
	}

	return false
}

// SetAccountLinkedEnrollments gets a reference to the given []AccountLinkedEnrollment and assigns it to the AccountLinkedEnrollments field.
func (o *OSAccountEmbedded) SetAccountLinkedEnrollments(v []AccountLinkedEnrollment) {
	o.AccountLinkedEnrollments = v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *OSAccountEmbedded) GetUsers() []User {
	if o == nil || IsNil(o.Users) {
		var ret []User
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSAccountEmbedded) GetUsersOk() ([]User, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *OSAccountEmbedded) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []User and assigns it to the Users field.
func (o *OSAccountEmbedded) SetUsers(v []User) {
	o.Users = v
}

func (o OSAccountEmbedded) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OSAccountEmbedded) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountLinkedEnrollments) {
		toSerialize["accountLinkedEnrollments"] = o.AccountLinkedEnrollments
	}
	if !IsNil(o.Users) {
		toSerialize["users"] = o.Users
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OSAccountEmbedded) UnmarshalJSON(data []byte) (err error) {
	varOSAccountEmbedded := _OSAccountEmbedded{}

	err = json.Unmarshal(data, &varOSAccountEmbedded)

	if err != nil {
		return err
	}

	*o = OSAccountEmbedded(varOSAccountEmbedded)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "accountLinkedEnrollments")
		delete(additionalProperties, "users")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOSAccountEmbedded struct {
	value *OSAccountEmbedded
	isSet bool
}

func (v NullableOSAccountEmbedded) Get() *OSAccountEmbedded {
	return v.value
}

func (v *NullableOSAccountEmbedded) Set(val *OSAccountEmbedded) {
	v.value = val
	v.isSet = true
}

func (v NullableOSAccountEmbedded) IsSet() bool {
	return v.isSet
}

func (v *NullableOSAccountEmbedded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOSAccountEmbedded(val *OSAccountEmbedded) *NullableOSAccountEmbedded {
	return &NullableOSAccountEmbedded{value: val, isSet: true}
}

func (v NullableOSAccountEmbedded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOSAccountEmbedded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
