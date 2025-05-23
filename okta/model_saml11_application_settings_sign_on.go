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

// Saml11ApplicationSettingsSignOn SAML 1.1 sign-on mode attributes
type Saml11ApplicationSettingsSignOn struct {
	// The intended audience of the SAML assertion. This is usually the Entity ID of your application.
	AudienceOverride *string `json:"audienceOverride,omitempty"`
	// The URL of the resource to direct users after they successfully sign in to the SP using SAML. See the SP documentation to check if you need to specify a RelayState. In most instances, you can leave this field blank.
	DefaultRelayState *string `json:"defaultRelayState,omitempty"`
	// The location where the application can present the SAML assertion. This is usually the Single Sign-On (SSO) URL.
	RecipientOverride *string `json:"recipientOverride,omitempty"`
	// Assertion Consumer Services (ACS) URL value for the Service Provider (SP). This URL is always used for Identity Provider (IdP) initiated sign-on requests.
	SsoAcsUrlOverride *string `json:"ssoAcsUrlOverride,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Saml11ApplicationSettingsSignOn Saml11ApplicationSettingsSignOn

// NewSaml11ApplicationSettingsSignOn instantiates a new Saml11ApplicationSettingsSignOn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSaml11ApplicationSettingsSignOn() *Saml11ApplicationSettingsSignOn {
	this := Saml11ApplicationSettingsSignOn{}
	return &this
}

// NewSaml11ApplicationSettingsSignOnWithDefaults instantiates a new Saml11ApplicationSettingsSignOn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSaml11ApplicationSettingsSignOnWithDefaults() *Saml11ApplicationSettingsSignOn {
	this := Saml11ApplicationSettingsSignOn{}
	return &this
}

// GetAudienceOverride returns the AudienceOverride field value if set, zero value otherwise.
func (o *Saml11ApplicationSettingsSignOn) GetAudienceOverride() string {
	if o == nil || o.AudienceOverride == nil {
		var ret string
		return ret
	}
	return *o.AudienceOverride
}

// GetAudienceOverrideOk returns a tuple with the AudienceOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Saml11ApplicationSettingsSignOn) GetAudienceOverrideOk() (*string, bool) {
	if o == nil || o.AudienceOverride == nil {
		return nil, false
	}
	return o.AudienceOverride, true
}

// HasAudienceOverride returns a boolean if a field has been set.
func (o *Saml11ApplicationSettingsSignOn) HasAudienceOverride() bool {
	if o != nil && o.AudienceOverride != nil {
		return true
	}

	return false
}

// SetAudienceOverride gets a reference to the given string and assigns it to the AudienceOverride field.
func (o *Saml11ApplicationSettingsSignOn) SetAudienceOverride(v string) {
	o.AudienceOverride = &v
}

// GetDefaultRelayState returns the DefaultRelayState field value if set, zero value otherwise.
func (o *Saml11ApplicationSettingsSignOn) GetDefaultRelayState() string {
	if o == nil || o.DefaultRelayState == nil {
		var ret string
		return ret
	}
	return *o.DefaultRelayState
}

// GetDefaultRelayStateOk returns a tuple with the DefaultRelayState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Saml11ApplicationSettingsSignOn) GetDefaultRelayStateOk() (*string, bool) {
	if o == nil || o.DefaultRelayState == nil {
		return nil, false
	}
	return o.DefaultRelayState, true
}

// HasDefaultRelayState returns a boolean if a field has been set.
func (o *Saml11ApplicationSettingsSignOn) HasDefaultRelayState() bool {
	if o != nil && o.DefaultRelayState != nil {
		return true
	}

	return false
}

// SetDefaultRelayState gets a reference to the given string and assigns it to the DefaultRelayState field.
func (o *Saml11ApplicationSettingsSignOn) SetDefaultRelayState(v string) {
	o.DefaultRelayState = &v
}

// GetRecipientOverride returns the RecipientOverride field value if set, zero value otherwise.
func (o *Saml11ApplicationSettingsSignOn) GetRecipientOverride() string {
	if o == nil || o.RecipientOverride == nil {
		var ret string
		return ret
	}
	return *o.RecipientOverride
}

// GetRecipientOverrideOk returns a tuple with the RecipientOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Saml11ApplicationSettingsSignOn) GetRecipientOverrideOk() (*string, bool) {
	if o == nil || o.RecipientOverride == nil {
		return nil, false
	}
	return o.RecipientOverride, true
}

// HasRecipientOverride returns a boolean if a field has been set.
func (o *Saml11ApplicationSettingsSignOn) HasRecipientOverride() bool {
	if o != nil && o.RecipientOverride != nil {
		return true
	}

	return false
}

// SetRecipientOverride gets a reference to the given string and assigns it to the RecipientOverride field.
func (o *Saml11ApplicationSettingsSignOn) SetRecipientOverride(v string) {
	o.RecipientOverride = &v
}

// GetSsoAcsUrlOverride returns the SsoAcsUrlOverride field value if set, zero value otherwise.
func (o *Saml11ApplicationSettingsSignOn) GetSsoAcsUrlOverride() string {
	if o == nil || o.SsoAcsUrlOverride == nil {
		var ret string
		return ret
	}
	return *o.SsoAcsUrlOverride
}

// GetSsoAcsUrlOverrideOk returns a tuple with the SsoAcsUrlOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Saml11ApplicationSettingsSignOn) GetSsoAcsUrlOverrideOk() (*string, bool) {
	if o == nil || o.SsoAcsUrlOverride == nil {
		return nil, false
	}
	return o.SsoAcsUrlOverride, true
}

// HasSsoAcsUrlOverride returns a boolean if a field has been set.
func (o *Saml11ApplicationSettingsSignOn) HasSsoAcsUrlOverride() bool {
	if o != nil && o.SsoAcsUrlOverride != nil {
		return true
	}

	return false
}

// SetSsoAcsUrlOverride gets a reference to the given string and assigns it to the SsoAcsUrlOverride field.
func (o *Saml11ApplicationSettingsSignOn) SetSsoAcsUrlOverride(v string) {
	o.SsoAcsUrlOverride = &v
}

func (o Saml11ApplicationSettingsSignOn) MarshalJSON() ([]byte, error) {
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

func (o *Saml11ApplicationSettingsSignOn) UnmarshalJSON(bytes []byte) (err error) {
	varSaml11ApplicationSettingsSignOn := _Saml11ApplicationSettingsSignOn{}

	err = json.Unmarshal(bytes, &varSaml11ApplicationSettingsSignOn)
	if err == nil {
		*o = Saml11ApplicationSettingsSignOn(varSaml11ApplicationSettingsSignOn)
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

type NullableSaml11ApplicationSettingsSignOn struct {
	value *Saml11ApplicationSettingsSignOn
	isSet bool
}

func (v NullableSaml11ApplicationSettingsSignOn) Get() *Saml11ApplicationSettingsSignOn {
	return v.value
}

func (v *NullableSaml11ApplicationSettingsSignOn) Set(val *Saml11ApplicationSettingsSignOn) {
	v.value = val
	v.isSet = true
}

func (v NullableSaml11ApplicationSettingsSignOn) IsSet() bool {
	return v.isSet
}

func (v *NullableSaml11ApplicationSettingsSignOn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSaml11ApplicationSettingsSignOn(val *Saml11ApplicationSettingsSignOn) *NullableSaml11ApplicationSettingsSignOn {
	return &NullableSaml11ApplicationSettingsSignOn{value: val, isSet: true}
}

func (v NullableSaml11ApplicationSettingsSignOn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSaml11ApplicationSettingsSignOn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

