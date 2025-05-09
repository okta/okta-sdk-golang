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

// UserFactorCustomHOTPProfile struct for UserFactorCustomHOTPProfile
type UserFactorCustomHOTPProfile struct {
	// Unique secret key used to generate the OTP
	SharedSecret *string `json:"sharedSecret,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserFactorCustomHOTPProfile UserFactorCustomHOTPProfile

// NewUserFactorCustomHOTPProfile instantiates a new UserFactorCustomHOTPProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserFactorCustomHOTPProfile() *UserFactorCustomHOTPProfile {
	this := UserFactorCustomHOTPProfile{}
	return &this
}

// NewUserFactorCustomHOTPProfileWithDefaults instantiates a new UserFactorCustomHOTPProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserFactorCustomHOTPProfileWithDefaults() *UserFactorCustomHOTPProfile {
	this := UserFactorCustomHOTPProfile{}
	return &this
}

// GetSharedSecret returns the SharedSecret field value if set, zero value otherwise.
func (o *UserFactorCustomHOTPProfile) GetSharedSecret() string {
	if o == nil || o.SharedSecret == nil {
		var ret string
		return ret
	}
	return *o.SharedSecret
}

// GetSharedSecretOk returns a tuple with the SharedSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorCustomHOTPProfile) GetSharedSecretOk() (*string, bool) {
	if o == nil || o.SharedSecret == nil {
		return nil, false
	}
	return o.SharedSecret, true
}

// HasSharedSecret returns a boolean if a field has been set.
func (o *UserFactorCustomHOTPProfile) HasSharedSecret() bool {
	if o != nil && o.SharedSecret != nil {
		return true
	}

	return false
}

// SetSharedSecret gets a reference to the given string and assigns it to the SharedSecret field.
func (o *UserFactorCustomHOTPProfile) SetSharedSecret(v string) {
	o.SharedSecret = &v
}

func (o UserFactorCustomHOTPProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SharedSecret != nil {
		toSerialize["sharedSecret"] = o.SharedSecret
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UserFactorCustomHOTPProfile) UnmarshalJSON(bytes []byte) (err error) {
	varUserFactorCustomHOTPProfile := _UserFactorCustomHOTPProfile{}

	err = json.Unmarshal(bytes, &varUserFactorCustomHOTPProfile)
	if err == nil {
		*o = UserFactorCustomHOTPProfile(varUserFactorCustomHOTPProfile)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "sharedSecret")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableUserFactorCustomHOTPProfile struct {
	value *UserFactorCustomHOTPProfile
	isSet bool
}

func (v NullableUserFactorCustomHOTPProfile) Get() *UserFactorCustomHOTPProfile {
	return v.value
}

func (v *NullableUserFactorCustomHOTPProfile) Set(val *UserFactorCustomHOTPProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableUserFactorCustomHOTPProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableUserFactorCustomHOTPProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserFactorCustomHOTPProfile(val *UserFactorCustomHOTPProfile) *NullableUserFactorCustomHOTPProfile {
	return &NullableUserFactorCustomHOTPProfile{value: val, isSet: true}
}

func (v NullableUserFactorCustomHOTPProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserFactorCustomHOTPProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

