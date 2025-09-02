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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// UserFactorTokenHOTPProfile struct for UserFactorTokenHOTPProfile
type UserFactorTokenHOTPProfile struct {
	// Unique secret key used to generate the OTP
	SharedSecret *string `json:"sharedSecret,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserFactorTokenHOTPProfile UserFactorTokenHOTPProfile

// NewUserFactorTokenHOTPProfile instantiates a new UserFactorTokenHOTPProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserFactorTokenHOTPProfile() *UserFactorTokenHOTPProfile {
	this := UserFactorTokenHOTPProfile{}
	return &this
}

// NewUserFactorTokenHOTPProfileWithDefaults instantiates a new UserFactorTokenHOTPProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserFactorTokenHOTPProfileWithDefaults() *UserFactorTokenHOTPProfile {
	this := UserFactorTokenHOTPProfile{}
	return &this
}

// GetSharedSecret returns the SharedSecret field value if set, zero value otherwise.
func (o *UserFactorTokenHOTPProfile) GetSharedSecret() string {
	if o == nil || o.SharedSecret == nil {
		var ret string
		return ret
	}
	return *o.SharedSecret
}

// GetSharedSecretOk returns a tuple with the SharedSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorTokenHOTPProfile) GetSharedSecretOk() (*string, bool) {
	if o == nil || o.SharedSecret == nil {
		return nil, false
	}
	return o.SharedSecret, true
}

// HasSharedSecret returns a boolean if a field has been set.
func (o *UserFactorTokenHOTPProfile) HasSharedSecret() bool {
	if o != nil && o.SharedSecret != nil {
		return true
	}

	return false
}

// SetSharedSecret gets a reference to the given string and assigns it to the SharedSecret field.
func (o *UserFactorTokenHOTPProfile) SetSharedSecret(v string) {
	o.SharedSecret = &v
}

func (o UserFactorTokenHOTPProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SharedSecret != nil {
		toSerialize["sharedSecret"] = o.SharedSecret
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UserFactorTokenHOTPProfile) UnmarshalJSON(bytes []byte) (err error) {
	varUserFactorTokenHOTPProfile := _UserFactorTokenHOTPProfile{}

	err = json.Unmarshal(bytes, &varUserFactorTokenHOTPProfile)
	if err == nil {
		*o = UserFactorTokenHOTPProfile(varUserFactorTokenHOTPProfile)
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

type NullableUserFactorTokenHOTPProfile struct {
	value *UserFactorTokenHOTPProfile
	isSet bool
}

func (v NullableUserFactorTokenHOTPProfile) Get() *UserFactorTokenHOTPProfile {
	return v.value
}

func (v *NullableUserFactorTokenHOTPProfile) Set(val *UserFactorTokenHOTPProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableUserFactorTokenHOTPProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableUserFactorTokenHOTPProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserFactorTokenHOTPProfile(val *UserFactorTokenHOTPProfile) *NullableUserFactorTokenHOTPProfile {
	return &NullableUserFactorTokenHOTPProfile{value: val, isSet: true}
}

func (v NullableUserFactorTokenHOTPProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserFactorTokenHOTPProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

