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

// checks if the AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll{}

// AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll Enrollment requirements for the authenticator
type AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll struct {
	// Requirements for the user-initiated enrollment
	Self                 *string                                   `json:"self,omitempty"`
	GracePeriod          *EnrollmentPolicyAuthenticatorGracePeriod `json:"gracePeriod,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll

// NewAuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll instantiates a new AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll() *AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll {
	this := AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll{}
	var self string = "NOT_ALLOWED"
	this.Self = &self
	return &this
}

// NewAuthenticatorEnrollmentPolicyAuthenticatorSettingsEnrollWithDefaults instantiates a new AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorEnrollmentPolicyAuthenticatorSettingsEnrollWithDefaults() *AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll {
	this := AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll{}
	var self string = "NOT_ALLOWED"
	this.Self = &self
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll) GetSelf() string {
	if o == nil || IsNil(o.Self) {
		var ret string
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll) GetSelfOk() (*string, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given string and assigns it to the Self field.
func (o *AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll) SetSelf(v string) {
	o.Self = &v
}

// GetGracePeriod returns the GracePeriod field value if set, zero value otherwise.
func (o *AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll) GetGracePeriod() EnrollmentPolicyAuthenticatorGracePeriod {
	if o == nil || IsNil(o.GracePeriod) {
		var ret EnrollmentPolicyAuthenticatorGracePeriod
		return ret
	}
	return *o.GracePeriod
}

// GetGracePeriodOk returns a tuple with the GracePeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll) GetGracePeriodOk() (*EnrollmentPolicyAuthenticatorGracePeriod, bool) {
	if o == nil || IsNil(o.GracePeriod) {
		return nil, false
	}
	return o.GracePeriod, true
}

// HasGracePeriod returns a boolean if a field has been set.
func (o *AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll) HasGracePeriod() bool {
	if o != nil && !IsNil(o.GracePeriod) {
		return true
	}

	return false
}

// SetGracePeriod gets a reference to the given EnrollmentPolicyAuthenticatorGracePeriod and assigns it to the GracePeriod field.
func (o *AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll) SetGracePeriod(v EnrollmentPolicyAuthenticatorGracePeriod) {
	o.GracePeriod = &v
}

func (o AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	if !IsNil(o.GracePeriod) {
		toSerialize["gracePeriod"] = o.GracePeriod
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll) UnmarshalJSON(data []byte) (err error) {
	varAuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll := _AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll{}

	err = json.Unmarshal(data, &varAuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll)

	if err != nil {
		return err
	}

	*o = AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll(varAuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "gracePeriod")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll struct {
	value *AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll
	isSet bool
}

func (v NullableAuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll) Get() *AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll {
	return v.value
}

func (v *NullableAuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll) Set(val *AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll(val *AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll) *NullableAuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll {
	return &NullableAuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll{value: val, isSet: true}
}

func (v NullableAuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
