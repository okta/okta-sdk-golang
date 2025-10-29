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
)

// checks if the AuthenticatorEnrollmentPolicyAuthenticatorSettingsConstraints type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthenticatorEnrollmentPolicyAuthenticatorSettingsConstraints{}

// AuthenticatorEnrollmentPolicyAuthenticatorSettingsConstraints Constraints for the authenticator
type AuthenticatorEnrollmentPolicyAuthenticatorSettingsConstraints struct {
	// The list of FIDO2 WebAuthn authenticator groups allowed for enrollment. The authenticators in the group are based on FIDO Alliance Metadata Service that's identified by name or the Authenticator Attestation Global Unique Identifier ([AAGUID](https://support.yubico.com/hc/en-us/articles/360016648959-YubiKey-Hardware-FIDO2-AAGUIDs)) number. These groups are defined in the [WebAuthn authenticator method settings](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/Authenticator/#tag/Authenticator/operation/listAuthenticatorMethods).
	AaguidGroups         []string `json:"aaguidGroups,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthenticatorEnrollmentPolicyAuthenticatorSettingsConstraints AuthenticatorEnrollmentPolicyAuthenticatorSettingsConstraints

// NewAuthenticatorEnrollmentPolicyAuthenticatorSettingsConstraints instantiates a new AuthenticatorEnrollmentPolicyAuthenticatorSettingsConstraints object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorEnrollmentPolicyAuthenticatorSettingsConstraints() *AuthenticatorEnrollmentPolicyAuthenticatorSettingsConstraints {
	this := AuthenticatorEnrollmentPolicyAuthenticatorSettingsConstraints{}
	return &this
}

// NewAuthenticatorEnrollmentPolicyAuthenticatorSettingsConstraintsWithDefaults instantiates a new AuthenticatorEnrollmentPolicyAuthenticatorSettingsConstraints object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorEnrollmentPolicyAuthenticatorSettingsConstraintsWithDefaults() *AuthenticatorEnrollmentPolicyAuthenticatorSettingsConstraints {
	this := AuthenticatorEnrollmentPolicyAuthenticatorSettingsConstraints{}
	return &this
}

// GetAaguidGroups returns the AaguidGroups field value if set, zero value otherwise.
func (o *AuthenticatorEnrollmentPolicyAuthenticatorSettingsConstraints) GetAaguidGroups() []string {
	if o == nil || IsNil(o.AaguidGroups) {
		var ret []string
		return ret
	}
	return o.AaguidGroups
}

// GetAaguidGroupsOk returns a tuple with the AaguidGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorEnrollmentPolicyAuthenticatorSettingsConstraints) GetAaguidGroupsOk() ([]string, bool) {
	if o == nil || IsNil(o.AaguidGroups) {
		return nil, false
	}
	return o.AaguidGroups, true
}

// HasAaguidGroups returns a boolean if a field has been set.
func (o *AuthenticatorEnrollmentPolicyAuthenticatorSettingsConstraints) HasAaguidGroups() bool {
	if o != nil && !IsNil(o.AaguidGroups) {
		return true
	}

	return false
}

// SetAaguidGroups gets a reference to the given []string and assigns it to the AaguidGroups field.
func (o *AuthenticatorEnrollmentPolicyAuthenticatorSettingsConstraints) SetAaguidGroups(v []string) {
	o.AaguidGroups = v
}

func (o AuthenticatorEnrollmentPolicyAuthenticatorSettingsConstraints) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthenticatorEnrollmentPolicyAuthenticatorSettingsConstraints) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AaguidGroups) {
		toSerialize["aaguidGroups"] = o.AaguidGroups
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AuthenticatorEnrollmentPolicyAuthenticatorSettingsConstraints) UnmarshalJSON(data []byte) (err error) {
	varAuthenticatorEnrollmentPolicyAuthenticatorSettingsConstraints := _AuthenticatorEnrollmentPolicyAuthenticatorSettingsConstraints{}

	err = json.Unmarshal(data, &varAuthenticatorEnrollmentPolicyAuthenticatorSettingsConstraints)

	if err != nil {
		return err
	}

	*o = AuthenticatorEnrollmentPolicyAuthenticatorSettingsConstraints(varAuthenticatorEnrollmentPolicyAuthenticatorSettingsConstraints)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "aaguidGroups")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAuthenticatorEnrollmentPolicyAuthenticatorSettingsConstraints struct {
	value *AuthenticatorEnrollmentPolicyAuthenticatorSettingsConstraints
	isSet bool
}

func (v NullableAuthenticatorEnrollmentPolicyAuthenticatorSettingsConstraints) Get() *AuthenticatorEnrollmentPolicyAuthenticatorSettingsConstraints {
	return v.value
}

func (v *NullableAuthenticatorEnrollmentPolicyAuthenticatorSettingsConstraints) Set(val *AuthenticatorEnrollmentPolicyAuthenticatorSettingsConstraints) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorEnrollmentPolicyAuthenticatorSettingsConstraints) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorEnrollmentPolicyAuthenticatorSettingsConstraints) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorEnrollmentPolicyAuthenticatorSettingsConstraints(val *AuthenticatorEnrollmentPolicyAuthenticatorSettingsConstraints) *NullableAuthenticatorEnrollmentPolicyAuthenticatorSettingsConstraints {
	return &NullableAuthenticatorEnrollmentPolicyAuthenticatorSettingsConstraints{value: val, isSet: true}
}

func (v NullableAuthenticatorEnrollmentPolicyAuthenticatorSettingsConstraints) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorEnrollmentPolicyAuthenticatorSettingsConstraints) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
