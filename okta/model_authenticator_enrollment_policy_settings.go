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

// checks if the AuthenticatorEnrollmentPolicySettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthenticatorEnrollmentPolicySettings{}

// AuthenticatorEnrollmentPolicySettings Specifies the policy level settings  > **Note:** In Identity Engine, the Multifactor (MFA) Enrollment policy name has changed to authenticator enrollment policy. The policy type of `MFA_ENROLL` remains unchanged. However, the `settings` data is updated for authenticators. Policy `settings` are included only for those authenticators that are enabled.
type AuthenticatorEnrollmentPolicySettings struct {
	// List of authenticator policy settings  <x-lifecycle class=\"oie\"></x-lifecycle> For orgs with the Authenticator enrollment policy feature enabled, the new default authenticator enrollment policy created by Okta contains the `authenticators` property in the policy settings. Existing default authenticator enrollment policies from a migrated Classic Engine org remain unchanged. The policies still use the `factors` property in their settings. The `authenticators` parameter allows you to configure all available authenticators, including authentication and recovery. The `factors` parameter only allows you to configure multifactor authentication.
	Authenticators []AuthenticatorEnrollmentPolicyAuthenticatorSettings `json:"authenticators,omitempty"`
	// Type of policy configuration object  <x-lifecycle class=\"oie\"></x-lifecycle> The `type` property in the policy `settings` is only applicable to the authenticator enrollment policy available in Identity Engine.
	Type                 *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthenticatorEnrollmentPolicySettings AuthenticatorEnrollmentPolicySettings

// NewAuthenticatorEnrollmentPolicySettings instantiates a new AuthenticatorEnrollmentPolicySettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorEnrollmentPolicySettings() *AuthenticatorEnrollmentPolicySettings {
	this := AuthenticatorEnrollmentPolicySettings{}
	var type_ string = "FACTORS"
	this.Type = &type_
	return &this
}

// NewAuthenticatorEnrollmentPolicySettingsWithDefaults instantiates a new AuthenticatorEnrollmentPolicySettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorEnrollmentPolicySettingsWithDefaults() *AuthenticatorEnrollmentPolicySettings {
	this := AuthenticatorEnrollmentPolicySettings{}
	var type_ string = "FACTORS"
	this.Type = &type_
	return &this
}

// GetAuthenticators returns the Authenticators field value if set, zero value otherwise.
func (o *AuthenticatorEnrollmentPolicySettings) GetAuthenticators() []AuthenticatorEnrollmentPolicyAuthenticatorSettings {
	if o == nil || IsNil(o.Authenticators) {
		var ret []AuthenticatorEnrollmentPolicyAuthenticatorSettings
		return ret
	}
	return o.Authenticators
}

// GetAuthenticatorsOk returns a tuple with the Authenticators field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorEnrollmentPolicySettings) GetAuthenticatorsOk() ([]AuthenticatorEnrollmentPolicyAuthenticatorSettings, bool) {
	if o == nil || IsNil(o.Authenticators) {
		return nil, false
	}
	return o.Authenticators, true
}

// HasAuthenticators returns a boolean if a field has been set.
func (o *AuthenticatorEnrollmentPolicySettings) HasAuthenticators() bool {
	if o != nil && !IsNil(o.Authenticators) {
		return true
	}

	return false
}

// SetAuthenticators gets a reference to the given []AuthenticatorEnrollmentPolicyAuthenticatorSettings and assigns it to the Authenticators field.
func (o *AuthenticatorEnrollmentPolicySettings) SetAuthenticators(v []AuthenticatorEnrollmentPolicyAuthenticatorSettings) {
	o.Authenticators = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AuthenticatorEnrollmentPolicySettings) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorEnrollmentPolicySettings) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AuthenticatorEnrollmentPolicySettings) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AuthenticatorEnrollmentPolicySettings) SetType(v string) {
	o.Type = &v
}

func (o AuthenticatorEnrollmentPolicySettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthenticatorEnrollmentPolicySettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Authenticators) {
		toSerialize["authenticators"] = o.Authenticators
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AuthenticatorEnrollmentPolicySettings) UnmarshalJSON(data []byte) (err error) {
	varAuthenticatorEnrollmentPolicySettings := _AuthenticatorEnrollmentPolicySettings{}

	err = json.Unmarshal(data, &varAuthenticatorEnrollmentPolicySettings)

	if err != nil {
		return err
	}

	*o = AuthenticatorEnrollmentPolicySettings(varAuthenticatorEnrollmentPolicySettings)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "authenticators")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAuthenticatorEnrollmentPolicySettings struct {
	value *AuthenticatorEnrollmentPolicySettings
	isSet bool
}

func (v NullableAuthenticatorEnrollmentPolicySettings) Get() *AuthenticatorEnrollmentPolicySettings {
	return v.value
}

func (v *NullableAuthenticatorEnrollmentPolicySettings) Set(val *AuthenticatorEnrollmentPolicySettings) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorEnrollmentPolicySettings) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorEnrollmentPolicySettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorEnrollmentPolicySettings(val *AuthenticatorEnrollmentPolicySettings) *NullableAuthenticatorEnrollmentPolicySettings {
	return &NullableAuthenticatorEnrollmentPolicySettings{value: val, isSet: true}
}

func (v NullableAuthenticatorEnrollmentPolicySettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorEnrollmentPolicySettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
