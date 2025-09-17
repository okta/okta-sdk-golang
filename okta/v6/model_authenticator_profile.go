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
	"fmt"
)

// checks if the AuthenticatorProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthenticatorProfile{}

// AuthenticatorProfile Defines the authenticator specific parameters
type AuthenticatorProfile struct {
	// The phone number for a `call` or `sms` authenticator enrollment.
	PhoneNumber          string `json:"phoneNumber"`
	AdditionalProperties map[string]interface{}
}

type _AuthenticatorProfile AuthenticatorProfile

// NewAuthenticatorProfile instantiates a new AuthenticatorProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorProfile(phoneNumber string) *AuthenticatorProfile {
	this := AuthenticatorProfile{}
	this.PhoneNumber = phoneNumber
	return &this
}

// NewAuthenticatorProfileWithDefaults instantiates a new AuthenticatorProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorProfileWithDefaults() *AuthenticatorProfile {
	this := AuthenticatorProfile{}
	return &this
}

// GetPhoneNumber returns the PhoneNumber field value
func (o *AuthenticatorProfile) GetPhoneNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorProfile) GetPhoneNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PhoneNumber, true
}

// SetPhoneNumber sets field value
func (o *AuthenticatorProfile) SetPhoneNumber(v string) {
	o.PhoneNumber = v
}

func (o AuthenticatorProfile) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthenticatorProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["phoneNumber"] = o.PhoneNumber

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AuthenticatorProfile) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"phoneNumber",
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

	varAuthenticatorProfile := _AuthenticatorProfile{}

	err = json.Unmarshal(data, &varAuthenticatorProfile)

	if err != nil {
		return err
	}

	*o = AuthenticatorProfile(varAuthenticatorProfile)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "phoneNumber")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAuthenticatorProfile struct {
	value *AuthenticatorProfile
	isSet bool
}

func (v NullableAuthenticatorProfile) Get() *AuthenticatorProfile {
	return v.value
}

func (v *NullableAuthenticatorProfile) Set(val *AuthenticatorProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorProfile(val *AuthenticatorProfile) *NullableAuthenticatorProfile {
	return &NullableAuthenticatorProfile{value: val, isSet: true}
}

func (v NullableAuthenticatorProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
