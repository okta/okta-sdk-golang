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
)

// checks if the AuthenticatorMethodConstraint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthenticatorMethodConstraint{}

// AuthenticatorMethodConstraint Limits the authenticators that can be used for a given method. Currently, only the `otp` method supports constraints, and Google authenticator (key : 'google_otp') is the only allowed authenticator.
type AuthenticatorMethodConstraint struct {
	AllowedAuthenticators []AuthenticatorIdentity `json:"allowedAuthenticators,omitempty"`
	Method                *string                 `json:"method,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _AuthenticatorMethodConstraint AuthenticatorMethodConstraint

// NewAuthenticatorMethodConstraint instantiates a new AuthenticatorMethodConstraint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorMethodConstraint() *AuthenticatorMethodConstraint {
	this := AuthenticatorMethodConstraint{}
	return &this
}

// NewAuthenticatorMethodConstraintWithDefaults instantiates a new AuthenticatorMethodConstraint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorMethodConstraintWithDefaults() *AuthenticatorMethodConstraint {
	this := AuthenticatorMethodConstraint{}
	return &this
}

// GetAllowedAuthenticators returns the AllowedAuthenticators field value if set, zero value otherwise.
func (o *AuthenticatorMethodConstraint) GetAllowedAuthenticators() []AuthenticatorIdentity {
	if o == nil || IsNil(o.AllowedAuthenticators) {
		var ret []AuthenticatorIdentity
		return ret
	}
	return o.AllowedAuthenticators
}

// GetAllowedAuthenticatorsOk returns a tuple with the AllowedAuthenticators field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorMethodConstraint) GetAllowedAuthenticatorsOk() ([]AuthenticatorIdentity, bool) {
	if o == nil || IsNil(o.AllowedAuthenticators) {
		return nil, false
	}
	return o.AllowedAuthenticators, true
}

// HasAllowedAuthenticators returns a boolean if a field has been set.
func (o *AuthenticatorMethodConstraint) HasAllowedAuthenticators() bool {
	if o != nil && !IsNil(o.AllowedAuthenticators) {
		return true
	}

	return false
}

// SetAllowedAuthenticators gets a reference to the given []AuthenticatorIdentity and assigns it to the AllowedAuthenticators field.
func (o *AuthenticatorMethodConstraint) SetAllowedAuthenticators(v []AuthenticatorIdentity) {
	o.AllowedAuthenticators = v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *AuthenticatorMethodConstraint) GetMethod() string {
	if o == nil || IsNil(o.Method) {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorMethodConstraint) GetMethodOk() (*string, bool) {
	if o == nil || IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *AuthenticatorMethodConstraint) HasMethod() bool {
	if o != nil && !IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *AuthenticatorMethodConstraint) SetMethod(v string) {
	o.Method = &v
}

func (o AuthenticatorMethodConstraint) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthenticatorMethodConstraint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllowedAuthenticators) {
		toSerialize["allowedAuthenticators"] = o.AllowedAuthenticators
	}
	if !IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AuthenticatorMethodConstraint) UnmarshalJSON(data []byte) (err error) {
	varAuthenticatorMethodConstraint := _AuthenticatorMethodConstraint{}

	err = json.Unmarshal(data, &varAuthenticatorMethodConstraint)

	if err != nil {
		return err
	}

	*o = AuthenticatorMethodConstraint(varAuthenticatorMethodConstraint)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "allowedAuthenticators")
		delete(additionalProperties, "method")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAuthenticatorMethodConstraint struct {
	value *AuthenticatorMethodConstraint
	isSet bool
}

func (v NullableAuthenticatorMethodConstraint) Get() *AuthenticatorMethodConstraint {
	return v.value
}

func (v *NullableAuthenticatorMethodConstraint) Set(val *AuthenticatorMethodConstraint) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorMethodConstraint) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorMethodConstraint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorMethodConstraint(val *AuthenticatorMethodConstraint) *NullableAuthenticatorMethodConstraint {
	return &NullableAuthenticatorMethodConstraint{value: val, isSet: true}
}

func (v NullableAuthenticatorMethodConstraint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorMethodConstraint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
