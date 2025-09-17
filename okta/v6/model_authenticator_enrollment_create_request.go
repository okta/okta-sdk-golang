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

// checks if the AuthenticatorEnrollmentCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthenticatorEnrollmentCreateRequest{}

// AuthenticatorEnrollmentCreateRequest struct for AuthenticatorEnrollmentCreateRequest
type AuthenticatorEnrollmentCreateRequest struct {
	// Unique identifier of the `phone` authenticator
	AuthenticatorId      string               `json:"authenticatorId"`
	Profile              AuthenticatorProfile `json:"profile"`
	AdditionalProperties map[string]interface{}
}

type _AuthenticatorEnrollmentCreateRequest AuthenticatorEnrollmentCreateRequest

// NewAuthenticatorEnrollmentCreateRequest instantiates a new AuthenticatorEnrollmentCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorEnrollmentCreateRequest(authenticatorId string, profile AuthenticatorProfile) *AuthenticatorEnrollmentCreateRequest {
	this := AuthenticatorEnrollmentCreateRequest{}
	this.AuthenticatorId = authenticatorId
	this.Profile = profile
	return &this
}

// NewAuthenticatorEnrollmentCreateRequestWithDefaults instantiates a new AuthenticatorEnrollmentCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorEnrollmentCreateRequestWithDefaults() *AuthenticatorEnrollmentCreateRequest {
	this := AuthenticatorEnrollmentCreateRequest{}
	return &this
}

// GetAuthenticatorId returns the AuthenticatorId field value
func (o *AuthenticatorEnrollmentCreateRequest) GetAuthenticatorId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthenticatorId
}

// GetAuthenticatorIdOk returns a tuple with the AuthenticatorId field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorEnrollmentCreateRequest) GetAuthenticatorIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthenticatorId, true
}

// SetAuthenticatorId sets field value
func (o *AuthenticatorEnrollmentCreateRequest) SetAuthenticatorId(v string) {
	o.AuthenticatorId = v
}

// GetProfile returns the Profile field value
func (o *AuthenticatorEnrollmentCreateRequest) GetProfile() AuthenticatorProfile {
	if o == nil {
		var ret AuthenticatorProfile
		return ret
	}

	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorEnrollmentCreateRequest) GetProfileOk() (*AuthenticatorProfile, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Profile, true
}

// SetProfile sets field value
func (o *AuthenticatorEnrollmentCreateRequest) SetProfile(v AuthenticatorProfile) {
	o.Profile = v
}

func (o AuthenticatorEnrollmentCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthenticatorEnrollmentCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["authenticatorId"] = o.AuthenticatorId
	toSerialize["profile"] = o.Profile

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AuthenticatorEnrollmentCreateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"authenticatorId",
		"profile",
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

	varAuthenticatorEnrollmentCreateRequest := _AuthenticatorEnrollmentCreateRequest{}

	err = json.Unmarshal(data, &varAuthenticatorEnrollmentCreateRequest)

	if err != nil {
		return err
	}

	*o = AuthenticatorEnrollmentCreateRequest(varAuthenticatorEnrollmentCreateRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "authenticatorId")
		delete(additionalProperties, "profile")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAuthenticatorEnrollmentCreateRequest struct {
	value *AuthenticatorEnrollmentCreateRequest
	isSet bool
}

func (v NullableAuthenticatorEnrollmentCreateRequest) Get() *AuthenticatorEnrollmentCreateRequest {
	return v.value
}

func (v *NullableAuthenticatorEnrollmentCreateRequest) Set(val *AuthenticatorEnrollmentCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorEnrollmentCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorEnrollmentCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorEnrollmentCreateRequest(val *AuthenticatorEnrollmentCreateRequest) *NullableAuthenticatorEnrollmentCreateRequest {
	return &NullableAuthenticatorEnrollmentCreateRequest{value: val, isSet: true}
}

func (v NullableAuthenticatorEnrollmentCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorEnrollmentCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
