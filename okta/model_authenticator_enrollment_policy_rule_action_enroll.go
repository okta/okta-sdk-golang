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

// checks if the AuthenticatorEnrollmentPolicyRuleActionEnroll type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthenticatorEnrollmentPolicyRuleActionEnroll{}

// AuthenticatorEnrollmentPolicyRuleActionEnroll Specifies whether the user is to be enrolled the first time they `LOGIN`, the next time they are in the `CHALLENGE` process, or `NEVER`
type AuthenticatorEnrollmentPolicyRuleActionEnroll struct {
	Self                 *string `json:"self,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthenticatorEnrollmentPolicyRuleActionEnroll AuthenticatorEnrollmentPolicyRuleActionEnroll

// NewAuthenticatorEnrollmentPolicyRuleActionEnroll instantiates a new AuthenticatorEnrollmentPolicyRuleActionEnroll object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorEnrollmentPolicyRuleActionEnroll() *AuthenticatorEnrollmentPolicyRuleActionEnroll {
	this := AuthenticatorEnrollmentPolicyRuleActionEnroll{}
	return &this
}

// NewAuthenticatorEnrollmentPolicyRuleActionEnrollWithDefaults instantiates a new AuthenticatorEnrollmentPolicyRuleActionEnroll object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorEnrollmentPolicyRuleActionEnrollWithDefaults() *AuthenticatorEnrollmentPolicyRuleActionEnroll {
	this := AuthenticatorEnrollmentPolicyRuleActionEnroll{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *AuthenticatorEnrollmentPolicyRuleActionEnroll) GetSelf() string {
	if o == nil || IsNil(o.Self) {
		var ret string
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorEnrollmentPolicyRuleActionEnroll) GetSelfOk() (*string, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *AuthenticatorEnrollmentPolicyRuleActionEnroll) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given string and assigns it to the Self field.
func (o *AuthenticatorEnrollmentPolicyRuleActionEnroll) SetSelf(v string) {
	o.Self = &v
}

func (o AuthenticatorEnrollmentPolicyRuleActionEnroll) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthenticatorEnrollmentPolicyRuleActionEnroll) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AuthenticatorEnrollmentPolicyRuleActionEnroll) UnmarshalJSON(data []byte) (err error) {
	varAuthenticatorEnrollmentPolicyRuleActionEnroll := _AuthenticatorEnrollmentPolicyRuleActionEnroll{}

	err = json.Unmarshal(data, &varAuthenticatorEnrollmentPolicyRuleActionEnroll)

	if err != nil {
		return err
	}

	*o = AuthenticatorEnrollmentPolicyRuleActionEnroll(varAuthenticatorEnrollmentPolicyRuleActionEnroll)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "self")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAuthenticatorEnrollmentPolicyRuleActionEnroll struct {
	value *AuthenticatorEnrollmentPolicyRuleActionEnroll
	isSet bool
}

func (v NullableAuthenticatorEnrollmentPolicyRuleActionEnroll) Get() *AuthenticatorEnrollmentPolicyRuleActionEnroll {
	return v.value
}

func (v *NullableAuthenticatorEnrollmentPolicyRuleActionEnroll) Set(val *AuthenticatorEnrollmentPolicyRuleActionEnroll) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorEnrollmentPolicyRuleActionEnroll) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorEnrollmentPolicyRuleActionEnroll) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorEnrollmentPolicyRuleActionEnroll(val *AuthenticatorEnrollmentPolicyRuleActionEnroll) *NullableAuthenticatorEnrollmentPolicyRuleActionEnroll {
	return &NullableAuthenticatorEnrollmentPolicyRuleActionEnroll{value: val, isSet: true}
}

func (v NullableAuthenticatorEnrollmentPolicyRuleActionEnroll) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorEnrollmentPolicyRuleActionEnroll) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
