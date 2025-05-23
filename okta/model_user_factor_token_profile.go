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

// UserFactorTokenProfile struct for UserFactorTokenProfile
type UserFactorTokenProfile struct {
	// ID for the Factor credential
	CredentialId *string `json:"credentialId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserFactorTokenProfile UserFactorTokenProfile

// NewUserFactorTokenProfile instantiates a new UserFactorTokenProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserFactorTokenProfile() *UserFactorTokenProfile {
	this := UserFactorTokenProfile{}
	return &this
}

// NewUserFactorTokenProfileWithDefaults instantiates a new UserFactorTokenProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserFactorTokenProfileWithDefaults() *UserFactorTokenProfile {
	this := UserFactorTokenProfile{}
	return &this
}

// GetCredentialId returns the CredentialId field value if set, zero value otherwise.
func (o *UserFactorTokenProfile) GetCredentialId() string {
	if o == nil || o.CredentialId == nil {
		var ret string
		return ret
	}
	return *o.CredentialId
}

// GetCredentialIdOk returns a tuple with the CredentialId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorTokenProfile) GetCredentialIdOk() (*string, bool) {
	if o == nil || o.CredentialId == nil {
		return nil, false
	}
	return o.CredentialId, true
}

// HasCredentialId returns a boolean if a field has been set.
func (o *UserFactorTokenProfile) HasCredentialId() bool {
	if o != nil && o.CredentialId != nil {
		return true
	}

	return false
}

// SetCredentialId gets a reference to the given string and assigns it to the CredentialId field.
func (o *UserFactorTokenProfile) SetCredentialId(v string) {
	o.CredentialId = &v
}

func (o UserFactorTokenProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CredentialId != nil {
		toSerialize["credentialId"] = o.CredentialId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UserFactorTokenProfile) UnmarshalJSON(bytes []byte) (err error) {
	varUserFactorTokenProfile := _UserFactorTokenProfile{}

	err = json.Unmarshal(bytes, &varUserFactorTokenProfile)
	if err == nil {
		*o = UserFactorTokenProfile(varUserFactorTokenProfile)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "credentialId")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableUserFactorTokenProfile struct {
	value *UserFactorTokenProfile
	isSet bool
}

func (v NullableUserFactorTokenProfile) Get() *UserFactorTokenProfile {
	return v.value
}

func (v *NullableUserFactorTokenProfile) Set(val *UserFactorTokenProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableUserFactorTokenProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableUserFactorTokenProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserFactorTokenProfile(val *UserFactorTokenProfile) *NullableUserFactorTokenProfile {
	return &NullableUserFactorTokenProfile{value: val, isSet: true}
}

func (v NullableUserFactorTokenProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserFactorTokenProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

