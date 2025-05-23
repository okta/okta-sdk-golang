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

// OINSaml11ApplicationSettingsSignOn Contains SAML 1.1 sign-on mode attributes
type OINSaml11ApplicationSettingsSignOn struct {
	// Audience override for CASB configuration. See [CASB config guide](https://help.okta.com/en-us/Content/Topics/Apps/CASB-config-guide.htm).
	AudienceOverride *string `json:"audienceOverride,omitempty"`
	// Identifies a specific application resource in an IdP-initiated SSO scenario
	DefaultRelayState *string `json:"defaultRelayState,omitempty"`
	// Recipient override for CASB configuration. See [CASB config guide](https://help.okta.com/en-us/Content/Topics/Apps/CASB-config-guide.htm).
	RecipientOverride *string `json:"recipientOverride,omitempty"`
	// Assertion Consumer Service (ACS) URL override for CASB configuration. See [CASB config guide](https://help.okta.com/en-us/Content/Topics/Apps/CASB-config-guide.htm).
	SsoAcsUrlOverride *string `json:"ssoAcsUrlOverride,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OINSaml11ApplicationSettingsSignOn OINSaml11ApplicationSettingsSignOn

// NewOINSaml11ApplicationSettingsSignOn instantiates a new OINSaml11ApplicationSettingsSignOn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOINSaml11ApplicationSettingsSignOn() *OINSaml11ApplicationSettingsSignOn {
	this := OINSaml11ApplicationSettingsSignOn{}
	return &this
}

// NewOINSaml11ApplicationSettingsSignOnWithDefaults instantiates a new OINSaml11ApplicationSettingsSignOn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOINSaml11ApplicationSettingsSignOnWithDefaults() *OINSaml11ApplicationSettingsSignOn {
	this := OINSaml11ApplicationSettingsSignOn{}
	return &this
}

// GetAudienceOverride returns the AudienceOverride field value if set, zero value otherwise.
func (o *OINSaml11ApplicationSettingsSignOn) GetAudienceOverride() string {
	if o == nil || o.AudienceOverride == nil {
		var ret string
		return ret
	}
	return *o.AudienceOverride
}

// GetAudienceOverrideOk returns a tuple with the AudienceOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OINSaml11ApplicationSettingsSignOn) GetAudienceOverrideOk() (*string, bool) {
	if o == nil || o.AudienceOverride == nil {
		return nil, false
	}
	return o.AudienceOverride, true
}

// HasAudienceOverride returns a boolean if a field has been set.
func (o *OINSaml11ApplicationSettingsSignOn) HasAudienceOverride() bool {
	if o != nil && o.AudienceOverride != nil {
		return true
	}

	return false
}

// SetAudienceOverride gets a reference to the given string and assigns it to the AudienceOverride field.
func (o *OINSaml11ApplicationSettingsSignOn) SetAudienceOverride(v string) {
	o.AudienceOverride = &v
}

// GetDefaultRelayState returns the DefaultRelayState field value if set, zero value otherwise.
func (o *OINSaml11ApplicationSettingsSignOn) GetDefaultRelayState() string {
	if o == nil || o.DefaultRelayState == nil {
		var ret string
		return ret
	}
	return *o.DefaultRelayState
}

// GetDefaultRelayStateOk returns a tuple with the DefaultRelayState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OINSaml11ApplicationSettingsSignOn) GetDefaultRelayStateOk() (*string, bool) {
	if o == nil || o.DefaultRelayState == nil {
		return nil, false
	}
	return o.DefaultRelayState, true
}

// HasDefaultRelayState returns a boolean if a field has been set.
func (o *OINSaml11ApplicationSettingsSignOn) HasDefaultRelayState() bool {
	if o != nil && o.DefaultRelayState != nil {
		return true
	}

	return false
}

// SetDefaultRelayState gets a reference to the given string and assigns it to the DefaultRelayState field.
func (o *OINSaml11ApplicationSettingsSignOn) SetDefaultRelayState(v string) {
	o.DefaultRelayState = &v
}

// GetRecipientOverride returns the RecipientOverride field value if set, zero value otherwise.
func (o *OINSaml11ApplicationSettingsSignOn) GetRecipientOverride() string {
	if o == nil || o.RecipientOverride == nil {
		var ret string
		return ret
	}
	return *o.RecipientOverride
}

// GetRecipientOverrideOk returns a tuple with the RecipientOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OINSaml11ApplicationSettingsSignOn) GetRecipientOverrideOk() (*string, bool) {
	if o == nil || o.RecipientOverride == nil {
		return nil, false
	}
	return o.RecipientOverride, true
}

// HasRecipientOverride returns a boolean if a field has been set.
func (o *OINSaml11ApplicationSettingsSignOn) HasRecipientOverride() bool {
	if o != nil && o.RecipientOverride != nil {
		return true
	}

	return false
}

// SetRecipientOverride gets a reference to the given string and assigns it to the RecipientOverride field.
func (o *OINSaml11ApplicationSettingsSignOn) SetRecipientOverride(v string) {
	o.RecipientOverride = &v
}

// GetSsoAcsUrlOverride returns the SsoAcsUrlOverride field value if set, zero value otherwise.
func (o *OINSaml11ApplicationSettingsSignOn) GetSsoAcsUrlOverride() string {
	if o == nil || o.SsoAcsUrlOverride == nil {
		var ret string
		return ret
	}
	return *o.SsoAcsUrlOverride
}

// GetSsoAcsUrlOverrideOk returns a tuple with the SsoAcsUrlOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OINSaml11ApplicationSettingsSignOn) GetSsoAcsUrlOverrideOk() (*string, bool) {
	if o == nil || o.SsoAcsUrlOverride == nil {
		return nil, false
	}
	return o.SsoAcsUrlOverride, true
}

// HasSsoAcsUrlOverride returns a boolean if a field has been set.
func (o *OINSaml11ApplicationSettingsSignOn) HasSsoAcsUrlOverride() bool {
	if o != nil && o.SsoAcsUrlOverride != nil {
		return true
	}

	return false
}

// SetSsoAcsUrlOverride gets a reference to the given string and assigns it to the SsoAcsUrlOverride field.
func (o *OINSaml11ApplicationSettingsSignOn) SetSsoAcsUrlOverride(v string) {
	o.SsoAcsUrlOverride = &v
}

func (o OINSaml11ApplicationSettingsSignOn) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AudienceOverride != nil {
		toSerialize["audienceOverride"] = o.AudienceOverride
	}
	if o.DefaultRelayState != nil {
		toSerialize["defaultRelayState"] = o.DefaultRelayState
	}
	if o.RecipientOverride != nil {
		toSerialize["recipientOverride"] = o.RecipientOverride
	}
	if o.SsoAcsUrlOverride != nil {
		toSerialize["ssoAcsUrlOverride"] = o.SsoAcsUrlOverride
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OINSaml11ApplicationSettingsSignOn) UnmarshalJSON(bytes []byte) (err error) {
	varOINSaml11ApplicationSettingsSignOn := _OINSaml11ApplicationSettingsSignOn{}

	err = json.Unmarshal(bytes, &varOINSaml11ApplicationSettingsSignOn)
	if err == nil {
		*o = OINSaml11ApplicationSettingsSignOn(varOINSaml11ApplicationSettingsSignOn)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "audienceOverride")
		delete(additionalProperties, "defaultRelayState")
		delete(additionalProperties, "recipientOverride")
		delete(additionalProperties, "ssoAcsUrlOverride")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableOINSaml11ApplicationSettingsSignOn struct {
	value *OINSaml11ApplicationSettingsSignOn
	isSet bool
}

func (v NullableOINSaml11ApplicationSettingsSignOn) Get() *OINSaml11ApplicationSettingsSignOn {
	return v.value
}

func (v *NullableOINSaml11ApplicationSettingsSignOn) Set(val *OINSaml11ApplicationSettingsSignOn) {
	v.value = val
	v.isSet = true
}

func (v NullableOINSaml11ApplicationSettingsSignOn) IsSet() bool {
	return v.isSet
}

func (v *NullableOINSaml11ApplicationSettingsSignOn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOINSaml11ApplicationSettingsSignOn(val *OINSaml11ApplicationSettingsSignOn) *NullableOINSaml11ApplicationSettingsSignOn {
	return &NullableOINSaml11ApplicationSettingsSignOn{value: val, isSet: true}
}

func (v NullableOINSaml11ApplicationSettingsSignOn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOINSaml11ApplicationSettingsSignOn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

