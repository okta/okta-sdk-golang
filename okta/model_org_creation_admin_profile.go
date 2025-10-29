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
	"fmt"
)

// checks if the OrgCreationAdminProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrgCreationAdminProfile{}

// OrgCreationAdminProfile Specifies the profile attributes for the first super admin user. The minimal set of required attributes are `email`, `firstName`, `lastName`, and `login`. See [profile](/openapi/okta-management/management/tag/User/#tag/User/operation/getUser!c=200&path=profile&t=response) for additional profile attributes.
type OrgCreationAdminProfile struct {
	// Given name of the User (`givenName`)
	FirstName NullableString `json:"firstName"`
	// The family name of the User (`familyName`)
	LastName NullableString `json:"lastName"`
	// The primary email address of the User. For validation, see [RFC 5322 Section 3.2.3](https://datatracker.ietf.org/doc/html/rfc5322#section-3.2.3).
	Email string `json:"email"`
	// The unique identifier for the User (`username`)
	Login                string `json:"login"`
	AdditionalProperties map[string]interface{}
}

type _OrgCreationAdminProfile OrgCreationAdminProfile

// NewOrgCreationAdminProfile instantiates a new OrgCreationAdminProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgCreationAdminProfile(firstName NullableString, lastName NullableString, email string, login string) *OrgCreationAdminProfile {
	this := OrgCreationAdminProfile{}
	this.FirstName = firstName
	this.LastName = lastName
	this.Email = email
	this.Login = login
	return &this
}

// NewOrgCreationAdminProfileWithDefaults instantiates a new OrgCreationAdminProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgCreationAdminProfileWithDefaults() *OrgCreationAdminProfile {
	this := OrgCreationAdminProfile{}
	return &this
}

// GetFirstName returns the FirstName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *OrgCreationAdminProfile) GetFirstName() string {
	if o == nil || o.FirstName.Get() == nil {
		var ret string
		return ret
	}

	return *o.FirstName.Get()
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrgCreationAdminProfile) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FirstName.Get(), o.FirstName.IsSet()
}

// SetFirstName sets field value
func (o *OrgCreationAdminProfile) SetFirstName(v string) {
	o.FirstName.Set(&v)
}

// GetLastName returns the LastName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *OrgCreationAdminProfile) GetLastName() string {
	if o == nil || o.LastName.Get() == nil {
		var ret string
		return ret
	}

	return *o.LastName.Get()
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrgCreationAdminProfile) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastName.Get(), o.LastName.IsSet()
}

// SetLastName sets field value
func (o *OrgCreationAdminProfile) SetLastName(v string) {
	o.LastName.Set(&v)
}

// GetEmail returns the Email field value
func (o *OrgCreationAdminProfile) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *OrgCreationAdminProfile) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *OrgCreationAdminProfile) SetEmail(v string) {
	o.Email = v
}

// GetLogin returns the Login field value
func (o *OrgCreationAdminProfile) GetLogin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Login
}

// GetLoginOk returns a tuple with the Login field value
// and a boolean to check if the value has been set.
func (o *OrgCreationAdminProfile) GetLoginOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Login, true
}

// SetLogin sets field value
func (o *OrgCreationAdminProfile) SetLogin(v string) {
	o.Login = v
}

func (o OrgCreationAdminProfile) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrgCreationAdminProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["firstName"] = o.FirstName.Get()
	toSerialize["lastName"] = o.LastName.Get()
	toSerialize["email"] = o.Email
	toSerialize["login"] = o.Login

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrgCreationAdminProfile) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"firstName",
		"lastName",
		"email",
		"login",
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

	varOrgCreationAdminProfile := _OrgCreationAdminProfile{}

	err = json.Unmarshal(data, &varOrgCreationAdminProfile)

	if err != nil {
		return err
	}

	*o = OrgCreationAdminProfile(varOrgCreationAdminProfile)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "firstName")
		delete(additionalProperties, "lastName")
		delete(additionalProperties, "email")
		delete(additionalProperties, "login")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrgCreationAdminProfile struct {
	value *OrgCreationAdminProfile
	isSet bool
}

func (v NullableOrgCreationAdminProfile) Get() *OrgCreationAdminProfile {
	return v.value
}

func (v *NullableOrgCreationAdminProfile) Set(val *OrgCreationAdminProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgCreationAdminProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgCreationAdminProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgCreationAdminProfile(val *OrgCreationAdminProfile) *NullableOrgCreationAdminProfile {
	return &NullableOrgCreationAdminProfile{value: val, isSet: true}
}

func (v NullableOrgCreationAdminProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgCreationAdminProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
