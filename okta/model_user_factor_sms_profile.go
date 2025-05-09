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

// UserFactorSMSProfile struct for UserFactorSMSProfile
type UserFactorSMSProfile struct {
	// Phone number of the Factor. You should format phone numbers to use the [E.164 standard](https://www.itu.int/rec/T-REC-E.164/).
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserFactorSMSProfile UserFactorSMSProfile

// NewUserFactorSMSProfile instantiates a new UserFactorSMSProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserFactorSMSProfile() *UserFactorSMSProfile {
	this := UserFactorSMSProfile{}
	return &this
}

// NewUserFactorSMSProfileWithDefaults instantiates a new UserFactorSMSProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserFactorSMSProfileWithDefaults() *UserFactorSMSProfile {
	this := UserFactorSMSProfile{}
	return &this
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *UserFactorSMSProfile) GetPhoneNumber() string {
	if o == nil || o.PhoneNumber == nil {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorSMSProfile) GetPhoneNumberOk() (*string, bool) {
	if o == nil || o.PhoneNumber == nil {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *UserFactorSMSProfile) HasPhoneNumber() bool {
	if o != nil && o.PhoneNumber != nil {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *UserFactorSMSProfile) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

func (o UserFactorSMSProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PhoneNumber != nil {
		toSerialize["phoneNumber"] = o.PhoneNumber
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UserFactorSMSProfile) UnmarshalJSON(bytes []byte) (err error) {
	varUserFactorSMSProfile := _UserFactorSMSProfile{}

	err = json.Unmarshal(bytes, &varUserFactorSMSProfile)
	if err == nil {
		*o = UserFactorSMSProfile(varUserFactorSMSProfile)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "phoneNumber")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableUserFactorSMSProfile struct {
	value *UserFactorSMSProfile
	isSet bool
}

func (v NullableUserFactorSMSProfile) Get() *UserFactorSMSProfile {
	return v.value
}

func (v *NullableUserFactorSMSProfile) Set(val *UserFactorSMSProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableUserFactorSMSProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableUserFactorSMSProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserFactorSMSProfile(val *UserFactorSMSProfile) *NullableUserFactorSMSProfile {
	return &NullableUserFactorSMSProfile{value: val, isSet: true}
}

func (v NullableUserFactorSMSProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserFactorSMSProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

