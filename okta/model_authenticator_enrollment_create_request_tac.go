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

// checks if the AuthenticatorEnrollmentCreateRequestTac type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthenticatorEnrollmentCreateRequestTac{}

// AuthenticatorEnrollmentCreateRequestTac struct for AuthenticatorEnrollmentCreateRequestTac
type AuthenticatorEnrollmentCreateRequestTac struct {
	// Unique identifier of the TAC authenticator
	AuthenticatorId      string                          `json:"authenticatorId"`
	Profile              *AuthenticatorProfileTacRequest `json:"profile,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthenticatorEnrollmentCreateRequestTac AuthenticatorEnrollmentCreateRequestTac

// NewAuthenticatorEnrollmentCreateRequestTac instantiates a new AuthenticatorEnrollmentCreateRequestTac object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorEnrollmentCreateRequestTac(authenticatorId string) *AuthenticatorEnrollmentCreateRequestTac {
	this := AuthenticatorEnrollmentCreateRequestTac{}
	this.AuthenticatorId = authenticatorId
	return &this
}

// NewAuthenticatorEnrollmentCreateRequestTacWithDefaults instantiates a new AuthenticatorEnrollmentCreateRequestTac object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorEnrollmentCreateRequestTacWithDefaults() *AuthenticatorEnrollmentCreateRequestTac {
	this := AuthenticatorEnrollmentCreateRequestTac{}
	return &this
}

// GetAuthenticatorId returns the AuthenticatorId field value
func (o *AuthenticatorEnrollmentCreateRequestTac) GetAuthenticatorId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthenticatorId
}

// GetAuthenticatorIdOk returns a tuple with the AuthenticatorId field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorEnrollmentCreateRequestTac) GetAuthenticatorIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthenticatorId, true
}

// SetAuthenticatorId sets field value
func (o *AuthenticatorEnrollmentCreateRequestTac) SetAuthenticatorId(v string) {
	o.AuthenticatorId = v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *AuthenticatorEnrollmentCreateRequestTac) GetProfile() AuthenticatorProfileTacRequest {
	if o == nil || IsNil(o.Profile) {
		var ret AuthenticatorProfileTacRequest
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorEnrollmentCreateRequestTac) GetProfileOk() (*AuthenticatorProfileTacRequest, bool) {
	if o == nil || IsNil(o.Profile) {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *AuthenticatorEnrollmentCreateRequestTac) HasProfile() bool {
	if o != nil && !IsNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given AuthenticatorProfileTacRequest and assigns it to the Profile field.
func (o *AuthenticatorEnrollmentCreateRequestTac) SetProfile(v AuthenticatorProfileTacRequest) {
	o.Profile = &v
}

func (o AuthenticatorEnrollmentCreateRequestTac) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthenticatorEnrollmentCreateRequestTac) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["authenticatorId"] = o.AuthenticatorId
	if !IsNil(o.Profile) {
		toSerialize["profile"] = o.Profile
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AuthenticatorEnrollmentCreateRequestTac) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"authenticatorId",
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

	varAuthenticatorEnrollmentCreateRequestTac := _AuthenticatorEnrollmentCreateRequestTac{}

	err = json.Unmarshal(data, &varAuthenticatorEnrollmentCreateRequestTac)

	if err != nil {
		return err
	}

	*o = AuthenticatorEnrollmentCreateRequestTac(varAuthenticatorEnrollmentCreateRequestTac)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "authenticatorId")
		delete(additionalProperties, "profile")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAuthenticatorEnrollmentCreateRequestTac struct {
	value *AuthenticatorEnrollmentCreateRequestTac
	isSet bool
}

func (v NullableAuthenticatorEnrollmentCreateRequestTac) Get() *AuthenticatorEnrollmentCreateRequestTac {
	return v.value
}

func (v *NullableAuthenticatorEnrollmentCreateRequestTac) Set(val *AuthenticatorEnrollmentCreateRequestTac) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorEnrollmentCreateRequestTac) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorEnrollmentCreateRequestTac) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorEnrollmentCreateRequestTac(val *AuthenticatorEnrollmentCreateRequestTac) *NullableAuthenticatorEnrollmentCreateRequestTac {
	return &NullableAuthenticatorEnrollmentCreateRequestTac{value: val, isSet: true}
}

func (v NullableAuthenticatorEnrollmentCreateRequestTac) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorEnrollmentCreateRequestTac) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
