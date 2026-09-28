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

// checks if the AccountLinkedEnrollmentProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountLinkedEnrollmentProfile{}

// AccountLinkedEnrollmentProfile Profile details for a linked enrollment
type AccountLinkedEnrollmentProfile struct {
	// Authentication method used for the enrollment
	AuthenticationMethod *string `json:"authenticationMethod,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountLinkedEnrollmentProfile AccountLinkedEnrollmentProfile

// NewAccountLinkedEnrollmentProfile instantiates a new AccountLinkedEnrollmentProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountLinkedEnrollmentProfile() *AccountLinkedEnrollmentProfile {
	this := AccountLinkedEnrollmentProfile{}
	return &this
}

// NewAccountLinkedEnrollmentProfileWithDefaults instantiates a new AccountLinkedEnrollmentProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountLinkedEnrollmentProfileWithDefaults() *AccountLinkedEnrollmentProfile {
	this := AccountLinkedEnrollmentProfile{}
	return &this
}

// GetAuthenticationMethod returns the AuthenticationMethod field value if set, zero value otherwise.
func (o *AccountLinkedEnrollmentProfile) GetAuthenticationMethod() string {
	if o == nil || IsNil(o.AuthenticationMethod) {
		var ret string
		return ret
	}
	return *o.AuthenticationMethod
}

// GetAuthenticationMethodOk returns a tuple with the AuthenticationMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountLinkedEnrollmentProfile) GetAuthenticationMethodOk() (*string, bool) {
	if o == nil || IsNil(o.AuthenticationMethod) {
		return nil, false
	}
	return o.AuthenticationMethod, true
}

// HasAuthenticationMethod returns a boolean if a field has been set.
func (o *AccountLinkedEnrollmentProfile) HasAuthenticationMethod() bool {
	if o != nil && !IsNil(o.AuthenticationMethod) {
		return true
	}

	return false
}

// SetAuthenticationMethod gets a reference to the given string and assigns it to the AuthenticationMethod field.
func (o *AccountLinkedEnrollmentProfile) SetAuthenticationMethod(v string) {
	o.AuthenticationMethod = &v
}

func (o AccountLinkedEnrollmentProfile) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountLinkedEnrollmentProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuthenticationMethod) {
		toSerialize["authenticationMethod"] = o.AuthenticationMethod
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountLinkedEnrollmentProfile) UnmarshalJSON(data []byte) (err error) {
	varAccountLinkedEnrollmentProfile := _AccountLinkedEnrollmentProfile{}

	err = json.Unmarshal(data, &varAccountLinkedEnrollmentProfile)

	if err != nil {
		return err
	}

	*o = AccountLinkedEnrollmentProfile(varAccountLinkedEnrollmentProfile)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "authenticationMethod")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountLinkedEnrollmentProfile struct {
	value *AccountLinkedEnrollmentProfile
	isSet bool
}

func (v NullableAccountLinkedEnrollmentProfile) Get() *AccountLinkedEnrollmentProfile {
	return v.value
}

func (v *NullableAccountLinkedEnrollmentProfile) Set(val *AccountLinkedEnrollmentProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountLinkedEnrollmentProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountLinkedEnrollmentProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountLinkedEnrollmentProfile(val *AccountLinkedEnrollmentProfile) *NullableAccountLinkedEnrollmentProfile {
	return &NullableAccountLinkedEnrollmentProfile{value: val, isSet: true}
}

func (v NullableAccountLinkedEnrollmentProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountLinkedEnrollmentProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
