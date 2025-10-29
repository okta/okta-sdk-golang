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

// checks if the AuthenticatorKeyOktaVerifyAllOfSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthenticatorKeyOktaVerifyAllOfSettings{}

// AuthenticatorKeyOktaVerifyAllOfSettings struct for AuthenticatorKeyOktaVerifyAllOfSettings
type AuthenticatorKeyOktaVerifyAllOfSettings struct {
	ChannelBinding *ChannelBinding `json:"channelBinding,omitempty"`
	Compliance     *Compliance     `json:"compliance,omitempty"`
	// User verification setting. Possible values `DISCOURAGED` (the authenticator isn't asked to perform user verification, but may do so at its discretion), `PREFERRED` (the client uses an authenticator capable of user verification if possible), or `REQUIRED`(the client uses only an authenticator capable of user verification)
	UserVerification *string `json:"userVerification,omitempty"`
	// The application instance ID
	AppInstanceId        *string `json:"appInstanceId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthenticatorKeyOktaVerifyAllOfSettings AuthenticatorKeyOktaVerifyAllOfSettings

// NewAuthenticatorKeyOktaVerifyAllOfSettings instantiates a new AuthenticatorKeyOktaVerifyAllOfSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorKeyOktaVerifyAllOfSettings() *AuthenticatorKeyOktaVerifyAllOfSettings {
	this := AuthenticatorKeyOktaVerifyAllOfSettings{}
	return &this
}

// NewAuthenticatorKeyOktaVerifyAllOfSettingsWithDefaults instantiates a new AuthenticatorKeyOktaVerifyAllOfSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorKeyOktaVerifyAllOfSettingsWithDefaults() *AuthenticatorKeyOktaVerifyAllOfSettings {
	this := AuthenticatorKeyOktaVerifyAllOfSettings{}
	return &this
}

// GetChannelBinding returns the ChannelBinding field value if set, zero value otherwise.
func (o *AuthenticatorKeyOktaVerifyAllOfSettings) GetChannelBinding() ChannelBinding {
	if o == nil || IsNil(o.ChannelBinding) {
		var ret ChannelBinding
		return ret
	}
	return *o.ChannelBinding
}

// GetChannelBindingOk returns a tuple with the ChannelBinding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorKeyOktaVerifyAllOfSettings) GetChannelBindingOk() (*ChannelBinding, bool) {
	if o == nil || IsNil(o.ChannelBinding) {
		return nil, false
	}
	return o.ChannelBinding, true
}

// HasChannelBinding returns a boolean if a field has been set.
func (o *AuthenticatorKeyOktaVerifyAllOfSettings) HasChannelBinding() bool {
	if o != nil && !IsNil(o.ChannelBinding) {
		return true
	}

	return false
}

// SetChannelBinding gets a reference to the given ChannelBinding and assigns it to the ChannelBinding field.
func (o *AuthenticatorKeyOktaVerifyAllOfSettings) SetChannelBinding(v ChannelBinding) {
	o.ChannelBinding = &v
}

// GetCompliance returns the Compliance field value if set, zero value otherwise.
func (o *AuthenticatorKeyOktaVerifyAllOfSettings) GetCompliance() Compliance {
	if o == nil || IsNil(o.Compliance) {
		var ret Compliance
		return ret
	}
	return *o.Compliance
}

// GetComplianceOk returns a tuple with the Compliance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorKeyOktaVerifyAllOfSettings) GetComplianceOk() (*Compliance, bool) {
	if o == nil || IsNil(o.Compliance) {
		return nil, false
	}
	return o.Compliance, true
}

// HasCompliance returns a boolean if a field has been set.
func (o *AuthenticatorKeyOktaVerifyAllOfSettings) HasCompliance() bool {
	if o != nil && !IsNil(o.Compliance) {
		return true
	}

	return false
}

// SetCompliance gets a reference to the given Compliance and assigns it to the Compliance field.
func (o *AuthenticatorKeyOktaVerifyAllOfSettings) SetCompliance(v Compliance) {
	o.Compliance = &v
}

// GetUserVerification returns the UserVerification field value if set, zero value otherwise.
func (o *AuthenticatorKeyOktaVerifyAllOfSettings) GetUserVerification() string {
	if o == nil || IsNil(o.UserVerification) {
		var ret string
		return ret
	}
	return *o.UserVerification
}

// GetUserVerificationOk returns a tuple with the UserVerification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorKeyOktaVerifyAllOfSettings) GetUserVerificationOk() (*string, bool) {
	if o == nil || IsNil(o.UserVerification) {
		return nil, false
	}
	return o.UserVerification, true
}

// HasUserVerification returns a boolean if a field has been set.
func (o *AuthenticatorKeyOktaVerifyAllOfSettings) HasUserVerification() bool {
	if o != nil && !IsNil(o.UserVerification) {
		return true
	}

	return false
}

// SetUserVerification gets a reference to the given string and assigns it to the UserVerification field.
func (o *AuthenticatorKeyOktaVerifyAllOfSettings) SetUserVerification(v string) {
	o.UserVerification = &v
}

// GetAppInstanceId returns the AppInstanceId field value if set, zero value otherwise.
func (o *AuthenticatorKeyOktaVerifyAllOfSettings) GetAppInstanceId() string {
	if o == nil || IsNil(o.AppInstanceId) {
		var ret string
		return ret
	}
	return *o.AppInstanceId
}

// GetAppInstanceIdOk returns a tuple with the AppInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorKeyOktaVerifyAllOfSettings) GetAppInstanceIdOk() (*string, bool) {
	if o == nil || IsNil(o.AppInstanceId) {
		return nil, false
	}
	return o.AppInstanceId, true
}

// HasAppInstanceId returns a boolean if a field has been set.
func (o *AuthenticatorKeyOktaVerifyAllOfSettings) HasAppInstanceId() bool {
	if o != nil && !IsNil(o.AppInstanceId) {
		return true
	}

	return false
}

// SetAppInstanceId gets a reference to the given string and assigns it to the AppInstanceId field.
func (o *AuthenticatorKeyOktaVerifyAllOfSettings) SetAppInstanceId(v string) {
	o.AppInstanceId = &v
}

func (o AuthenticatorKeyOktaVerifyAllOfSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthenticatorKeyOktaVerifyAllOfSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ChannelBinding) {
		toSerialize["channelBinding"] = o.ChannelBinding
	}
	if !IsNil(o.Compliance) {
		toSerialize["compliance"] = o.Compliance
	}
	if !IsNil(o.UserVerification) {
		toSerialize["userVerification"] = o.UserVerification
	}
	if !IsNil(o.AppInstanceId) {
		toSerialize["appInstanceId"] = o.AppInstanceId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AuthenticatorKeyOktaVerifyAllOfSettings) UnmarshalJSON(data []byte) (err error) {
	varAuthenticatorKeyOktaVerifyAllOfSettings := _AuthenticatorKeyOktaVerifyAllOfSettings{}

	err = json.Unmarshal(data, &varAuthenticatorKeyOktaVerifyAllOfSettings)

	if err != nil {
		return err
	}

	*o = AuthenticatorKeyOktaVerifyAllOfSettings(varAuthenticatorKeyOktaVerifyAllOfSettings)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "channelBinding")
		delete(additionalProperties, "compliance")
		delete(additionalProperties, "userVerification")
		delete(additionalProperties, "appInstanceId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAuthenticatorKeyOktaVerifyAllOfSettings struct {
	value *AuthenticatorKeyOktaVerifyAllOfSettings
	isSet bool
}

func (v NullableAuthenticatorKeyOktaVerifyAllOfSettings) Get() *AuthenticatorKeyOktaVerifyAllOfSettings {
	return v.value
}

func (v *NullableAuthenticatorKeyOktaVerifyAllOfSettings) Set(val *AuthenticatorKeyOktaVerifyAllOfSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorKeyOktaVerifyAllOfSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorKeyOktaVerifyAllOfSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorKeyOktaVerifyAllOfSettings(val *AuthenticatorKeyOktaVerifyAllOfSettings) *NullableAuthenticatorKeyOktaVerifyAllOfSettings {
	return &NullableAuthenticatorKeyOktaVerifyAllOfSettings{value: val, isSet: true}
}

func (v NullableAuthenticatorKeyOktaVerifyAllOfSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorKeyOktaVerifyAllOfSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
