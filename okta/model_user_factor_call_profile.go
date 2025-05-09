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

// UserFactorCallProfile struct for UserFactorCallProfile
type UserFactorCallProfile struct {
	// Extension of the associated `phoneNumber`
	PhoneExtension NullableString `json:"phoneExtension,omitempty"`
	// Phone number of the Factor. You should format phone numbers to use the [E.164 standard](https://www.itu.int/rec/T-REC-E.164/).
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserFactorCallProfile UserFactorCallProfile

// NewUserFactorCallProfile instantiates a new UserFactorCallProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserFactorCallProfile() *UserFactorCallProfile {
	this := UserFactorCallProfile{}
	return &this
}

// NewUserFactorCallProfileWithDefaults instantiates a new UserFactorCallProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserFactorCallProfileWithDefaults() *UserFactorCallProfile {
	this := UserFactorCallProfile{}
	return &this
}

// GetPhoneExtension returns the PhoneExtension field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserFactorCallProfile) GetPhoneExtension() string {
	if o == nil || o.PhoneExtension.Get() == nil {
		var ret string
		return ret
	}
	return *o.PhoneExtension.Get()
}

// GetPhoneExtensionOk returns a tuple with the PhoneExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserFactorCallProfile) GetPhoneExtensionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PhoneExtension.Get(), o.PhoneExtension.IsSet()
}

// HasPhoneExtension returns a boolean if a field has been set.
func (o *UserFactorCallProfile) HasPhoneExtension() bool {
	if o != nil && o.PhoneExtension.IsSet() {
		return true
	}

	return false
}

// SetPhoneExtension gets a reference to the given NullableString and assigns it to the PhoneExtension field.
func (o *UserFactorCallProfile) SetPhoneExtension(v string) {
	o.PhoneExtension.Set(&v)
}
// SetPhoneExtensionNil sets the value for PhoneExtension to be an explicit nil
func (o *UserFactorCallProfile) SetPhoneExtensionNil() {
	o.PhoneExtension.Set(nil)
}

// UnsetPhoneExtension ensures that no value is present for PhoneExtension, not even an explicit nil
func (o *UserFactorCallProfile) UnsetPhoneExtension() {
	o.PhoneExtension.Unset()
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *UserFactorCallProfile) GetPhoneNumber() string {
	if o == nil || o.PhoneNumber == nil {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorCallProfile) GetPhoneNumberOk() (*string, bool) {
	if o == nil || o.PhoneNumber == nil {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *UserFactorCallProfile) HasPhoneNumber() bool {
	if o != nil && o.PhoneNumber != nil {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *UserFactorCallProfile) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

func (o UserFactorCallProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PhoneExtension.IsSet() {
		toSerialize["phoneExtension"] = o.PhoneExtension.Get()
	}
	if o.PhoneNumber != nil {
		toSerialize["phoneNumber"] = o.PhoneNumber
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UserFactorCallProfile) UnmarshalJSON(bytes []byte) (err error) {
	varUserFactorCallProfile := _UserFactorCallProfile{}

	err = json.Unmarshal(bytes, &varUserFactorCallProfile)
	if err == nil {
		*o = UserFactorCallProfile(varUserFactorCallProfile)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "phoneExtension")
		delete(additionalProperties, "phoneNumber")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableUserFactorCallProfile struct {
	value *UserFactorCallProfile
	isSet bool
}

func (v NullableUserFactorCallProfile) Get() *UserFactorCallProfile {
	return v.value
}

func (v *NullableUserFactorCallProfile) Set(val *UserFactorCallProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableUserFactorCallProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableUserFactorCallProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserFactorCallProfile(val *UserFactorCallProfile) *NullableUserFactorCallProfile {
	return &NullableUserFactorCallProfile{value: val, isSet: true}
}

func (v NullableUserFactorCallProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserFactorCallProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

