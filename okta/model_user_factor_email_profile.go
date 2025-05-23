/*
Okta Admin Management

Allows customers to easily access the Okta Management APIs

Copyright 2018 - Present Okta, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

API version: 2024.06.1
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// UserFactorEmailProfile struct for UserFactorEmailProfile
type UserFactorEmailProfile struct {
	// Email address of the user. Must be either the primary or secondary email address associated with the Okta user account.
	Email *string `json:"email,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserFactorEmailProfile UserFactorEmailProfile

// NewUserFactorEmailProfile instantiates a new UserFactorEmailProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserFactorEmailProfile() *UserFactorEmailProfile {
	this := UserFactorEmailProfile{}
	return &this
}

// NewUserFactorEmailProfileWithDefaults instantiates a new UserFactorEmailProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserFactorEmailProfileWithDefaults() *UserFactorEmailProfile {
	this := UserFactorEmailProfile{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UserFactorEmailProfile) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorEmailProfile) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UserFactorEmailProfile) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UserFactorEmailProfile) SetEmail(v string) {
	o.Email = &v
}

func (o UserFactorEmailProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UserFactorEmailProfile) UnmarshalJSON(bytes []byte) (err error) {
	varUserFactorEmailProfile := _UserFactorEmailProfile{}

	err = json.Unmarshal(bytes, &varUserFactorEmailProfile)
	if err == nil {
		*o = UserFactorEmailProfile(varUserFactorEmailProfile)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "email")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableUserFactorEmailProfile struct {
	value *UserFactorEmailProfile
	isSet bool
}

func (v NullableUserFactorEmailProfile) Get() *UserFactorEmailProfile {
	return v.value
}

func (v *NullableUserFactorEmailProfile) Set(val *UserFactorEmailProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableUserFactorEmailProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableUserFactorEmailProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserFactorEmailProfile(val *UserFactorEmailProfile) *NullableUserFactorEmailProfile {
	return &NullableUserFactorEmailProfile{value: val, isSet: true}
}

func (v NullableUserFactorEmailProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserFactorEmailProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

