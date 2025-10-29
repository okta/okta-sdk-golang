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

API version: 2025.10.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"time"
)

// checks if the AuthenticatorProfileTacResponsePost type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthenticatorProfileTacResponsePost{}

// AuthenticatorProfileTacResponsePost Defines the authenticator specific parameters
type AuthenticatorProfileTacResponsePost struct {
	// The time when the TAC enrollment expires in the UTC timezone
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
	// Determines whether an enrollment can be used more than once
	MultiUse *bool `json:"multiUse,omitempty"`
	// A temporary access code used for authentication. It can be used one or more times and is valid for a defined period specified by the `ttl` property. The `tac` is returned in the response when the enrollment is created. It is not returned when the enrollment is retrieved. Issuing a new TAC invalidates any existing TAC for this user.
	Tac                  *string `json:"tac,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthenticatorProfileTacResponsePost AuthenticatorProfileTacResponsePost

// NewAuthenticatorProfileTacResponsePost instantiates a new AuthenticatorProfileTacResponsePost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorProfileTacResponsePost() *AuthenticatorProfileTacResponsePost {
	this := AuthenticatorProfileTacResponsePost{}
	return &this
}

// NewAuthenticatorProfileTacResponsePostWithDefaults instantiates a new AuthenticatorProfileTacResponsePost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorProfileTacResponsePostWithDefaults() *AuthenticatorProfileTacResponsePost {
	this := AuthenticatorProfileTacResponsePost{}
	return &this
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *AuthenticatorProfileTacResponsePost) GetExpiresAt() time.Time {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorProfileTacResponsePost) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *AuthenticatorProfileTacResponsePost) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *AuthenticatorProfileTacResponsePost) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

// GetMultiUse returns the MultiUse field value if set, zero value otherwise.
func (o *AuthenticatorProfileTacResponsePost) GetMultiUse() bool {
	if o == nil || IsNil(o.MultiUse) {
		var ret bool
		return ret
	}
	return *o.MultiUse
}

// GetMultiUseOk returns a tuple with the MultiUse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorProfileTacResponsePost) GetMultiUseOk() (*bool, bool) {
	if o == nil || IsNil(o.MultiUse) {
		return nil, false
	}
	return o.MultiUse, true
}

// HasMultiUse returns a boolean if a field has been set.
func (o *AuthenticatorProfileTacResponsePost) HasMultiUse() bool {
	if o != nil && !IsNil(o.MultiUse) {
		return true
	}

	return false
}

// SetMultiUse gets a reference to the given bool and assigns it to the MultiUse field.
func (o *AuthenticatorProfileTacResponsePost) SetMultiUse(v bool) {
	o.MultiUse = &v
}

// GetTac returns the Tac field value if set, zero value otherwise.
func (o *AuthenticatorProfileTacResponsePost) GetTac() string {
	if o == nil || IsNil(o.Tac) {
		var ret string
		return ret
	}
	return *o.Tac
}

// GetTacOk returns a tuple with the Tac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorProfileTacResponsePost) GetTacOk() (*string, bool) {
	if o == nil || IsNil(o.Tac) {
		return nil, false
	}
	return o.Tac, true
}

// HasTac returns a boolean if a field has been set.
func (o *AuthenticatorProfileTacResponsePost) HasTac() bool {
	if o != nil && !IsNil(o.Tac) {
		return true
	}

	return false
}

// SetTac gets a reference to the given string and assigns it to the Tac field.
func (o *AuthenticatorProfileTacResponsePost) SetTac(v string) {
	o.Tac = &v
}

func (o AuthenticatorProfileTacResponsePost) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthenticatorProfileTacResponsePost) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExpiresAt) {
		toSerialize["expiresAt"] = o.ExpiresAt
	}
	if !IsNil(o.MultiUse) {
		toSerialize["multiUse"] = o.MultiUse
	}
	if !IsNil(o.Tac) {
		toSerialize["tac"] = o.Tac
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AuthenticatorProfileTacResponsePost) UnmarshalJSON(data []byte) (err error) {
	varAuthenticatorProfileTacResponsePost := _AuthenticatorProfileTacResponsePost{}

	err = json.Unmarshal(data, &varAuthenticatorProfileTacResponsePost)

	if err != nil {
		return err
	}

	*o = AuthenticatorProfileTacResponsePost(varAuthenticatorProfileTacResponsePost)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "expiresAt")
		delete(additionalProperties, "multiUse")
		delete(additionalProperties, "tac")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAuthenticatorProfileTacResponsePost struct {
	value *AuthenticatorProfileTacResponsePost
	isSet bool
}

func (v NullableAuthenticatorProfileTacResponsePost) Get() *AuthenticatorProfileTacResponsePost {
	return v.value
}

func (v *NullableAuthenticatorProfileTacResponsePost) Set(val *AuthenticatorProfileTacResponsePost) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorProfileTacResponsePost) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorProfileTacResponsePost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorProfileTacResponsePost(val *AuthenticatorProfileTacResponsePost) *NullableAuthenticatorProfileTacResponsePost {
	return &NullableAuthenticatorProfileTacResponsePost{value: val, isSet: true}
}

func (v NullableAuthenticatorProfileTacResponsePost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorProfileTacResponsePost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
