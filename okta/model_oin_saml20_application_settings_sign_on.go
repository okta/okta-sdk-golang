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

// OINSaml20ApplicationSettingsSignOn Contains SAML 2.0 sign-on mode attributes. > **Note:** Set `destinationOverride` to configure any other SAML 2.0 attributes in this section.
type OINSaml20ApplicationSettingsSignOn struct {
	// Audience override for CASB configuration. See [CASB config guide](https://help.okta.com/en-us/Content/Topics/Apps/CASB-config-guide.htm).
	AudienceOverride *string `json:"audienceOverride,omitempty"`
	ConfiguredAttributeStatements []SamlAttributeStatement `json:"configuredAttributeStatements,omitempty"`
	// Identifies a specific application resource in an IdP-initiated SSO scenario
	DefaultRelayState *string `json:"defaultRelayState,omitempty"`
	// Destination override for CASB configuration. See [CASB config guide](https://help.okta.com/en-us/Content/Topics/Apps/CASB-config-guide.htm).
	DestinationOverride *string `json:"destinationOverride,omitempty"`
	// Recipient override for CASB configuration. See [CASB config guide](https://help.okta.com/en-us/Content/Topics/Apps/CASB-config-guide.htm).
	RecipientOverride *string `json:"recipientOverride,omitempty"`
	// Determines the SAML app session lifetimes with Okta
	SamlAssertionLifetimeSeconds *int32 `json:"samlAssertionLifetimeSeconds,omitempty"`
	// Assertion Consumer Service (ACS) URL override for CASB configuration. See [CASB config guide](https://help.okta.com/en-us/Content/Topics/Apps/CASB-config-guide.htm).
	SsoAcsUrlOverride *string `json:"ssoAcsUrlOverride,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OINSaml20ApplicationSettingsSignOn OINSaml20ApplicationSettingsSignOn

// NewOINSaml20ApplicationSettingsSignOn instantiates a new OINSaml20ApplicationSettingsSignOn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOINSaml20ApplicationSettingsSignOn() *OINSaml20ApplicationSettingsSignOn {
	this := OINSaml20ApplicationSettingsSignOn{}
	return &this
}

// NewOINSaml20ApplicationSettingsSignOnWithDefaults instantiates a new OINSaml20ApplicationSettingsSignOn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOINSaml20ApplicationSettingsSignOnWithDefaults() *OINSaml20ApplicationSettingsSignOn {
	this := OINSaml20ApplicationSettingsSignOn{}
	return &this
}

// GetAudienceOverride returns the AudienceOverride field value if set, zero value otherwise.
func (o *OINSaml20ApplicationSettingsSignOn) GetAudienceOverride() string {
	if o == nil || o.AudienceOverride == nil {
		var ret string
		return ret
	}
	return *o.AudienceOverride
}

// GetAudienceOverrideOk returns a tuple with the AudienceOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OINSaml20ApplicationSettingsSignOn) GetAudienceOverrideOk() (*string, bool) {
	if o == nil || o.AudienceOverride == nil {
		return nil, false
	}
	return o.AudienceOverride, true
}

// HasAudienceOverride returns a boolean if a field has been set.
func (o *OINSaml20ApplicationSettingsSignOn) HasAudienceOverride() bool {
	if o != nil && o.AudienceOverride != nil {
		return true
	}

	return false
}

// SetAudienceOverride gets a reference to the given string and assigns it to the AudienceOverride field.
func (o *OINSaml20ApplicationSettingsSignOn) SetAudienceOverride(v string) {
	o.AudienceOverride = &v
}

// GetConfiguredAttributeStatements returns the ConfiguredAttributeStatements field value if set, zero value otherwise.
func (o *OINSaml20ApplicationSettingsSignOn) GetConfiguredAttributeStatements() []SamlAttributeStatement {
	if o == nil || o.ConfiguredAttributeStatements == nil {
		var ret []SamlAttributeStatement
		return ret
	}
	return o.ConfiguredAttributeStatements
}

// GetConfiguredAttributeStatementsOk returns a tuple with the ConfiguredAttributeStatements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OINSaml20ApplicationSettingsSignOn) GetConfiguredAttributeStatementsOk() ([]SamlAttributeStatement, bool) {
	if o == nil || o.ConfiguredAttributeStatements == nil {
		return nil, false
	}
	return o.ConfiguredAttributeStatements, true
}

// HasConfiguredAttributeStatements returns a boolean if a field has been set.
func (o *OINSaml20ApplicationSettingsSignOn) HasConfiguredAttributeStatements() bool {
	if o != nil && o.ConfiguredAttributeStatements != nil {
		return true
	}

	return false
}

// SetConfiguredAttributeStatements gets a reference to the given []SamlAttributeStatement and assigns it to the ConfiguredAttributeStatements field.
func (o *OINSaml20ApplicationSettingsSignOn) SetConfiguredAttributeStatements(v []SamlAttributeStatement) {
	o.ConfiguredAttributeStatements = v
}

// GetDefaultRelayState returns the DefaultRelayState field value if set, zero value otherwise.
func (o *OINSaml20ApplicationSettingsSignOn) GetDefaultRelayState() string {
	if o == nil || o.DefaultRelayState == nil {
		var ret string
		return ret
	}
	return *o.DefaultRelayState
}

// GetDefaultRelayStateOk returns a tuple with the DefaultRelayState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OINSaml20ApplicationSettingsSignOn) GetDefaultRelayStateOk() (*string, bool) {
	if o == nil || o.DefaultRelayState == nil {
		return nil, false
	}
	return o.DefaultRelayState, true
}

// HasDefaultRelayState returns a boolean if a field has been set.
func (o *OINSaml20ApplicationSettingsSignOn) HasDefaultRelayState() bool {
	if o != nil && o.DefaultRelayState != nil {
		return true
	}

	return false
}

// SetDefaultRelayState gets a reference to the given string and assigns it to the DefaultRelayState field.
func (o *OINSaml20ApplicationSettingsSignOn) SetDefaultRelayState(v string) {
	o.DefaultRelayState = &v
}

// GetDestinationOverride returns the DestinationOverride field value if set, zero value otherwise.
func (o *OINSaml20ApplicationSettingsSignOn) GetDestinationOverride() string {
	if o == nil || o.DestinationOverride == nil {
		var ret string
		return ret
	}
	return *o.DestinationOverride
}

// GetDestinationOverrideOk returns a tuple with the DestinationOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OINSaml20ApplicationSettingsSignOn) GetDestinationOverrideOk() (*string, bool) {
	if o == nil || o.DestinationOverride == nil {
		return nil, false
	}
	return o.DestinationOverride, true
}

// HasDestinationOverride returns a boolean if a field has been set.
func (o *OINSaml20ApplicationSettingsSignOn) HasDestinationOverride() bool {
	if o != nil && o.DestinationOverride != nil {
		return true
	}

	return false
}

// SetDestinationOverride gets a reference to the given string and assigns it to the DestinationOverride field.
func (o *OINSaml20ApplicationSettingsSignOn) SetDestinationOverride(v string) {
	o.DestinationOverride = &v
}

// GetRecipientOverride returns the RecipientOverride field value if set, zero value otherwise.
func (o *OINSaml20ApplicationSettingsSignOn) GetRecipientOverride() string {
	if o == nil || o.RecipientOverride == nil {
		var ret string
		return ret
	}
	return *o.RecipientOverride
}

// GetRecipientOverrideOk returns a tuple with the RecipientOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OINSaml20ApplicationSettingsSignOn) GetRecipientOverrideOk() (*string, bool) {
	if o == nil || o.RecipientOverride == nil {
		return nil, false
	}
	return o.RecipientOverride, true
}

// HasRecipientOverride returns a boolean if a field has been set.
func (o *OINSaml20ApplicationSettingsSignOn) HasRecipientOverride() bool {
	if o != nil && o.RecipientOverride != nil {
		return true
	}

	return false
}

// SetRecipientOverride gets a reference to the given string and assigns it to the RecipientOverride field.
func (o *OINSaml20ApplicationSettingsSignOn) SetRecipientOverride(v string) {
	o.RecipientOverride = &v
}

// GetSamlAssertionLifetimeSeconds returns the SamlAssertionLifetimeSeconds field value if set, zero value otherwise.
func (o *OINSaml20ApplicationSettingsSignOn) GetSamlAssertionLifetimeSeconds() int32 {
	if o == nil || o.SamlAssertionLifetimeSeconds == nil {
		var ret int32
		return ret
	}
	return *o.SamlAssertionLifetimeSeconds
}

// GetSamlAssertionLifetimeSecondsOk returns a tuple with the SamlAssertionLifetimeSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OINSaml20ApplicationSettingsSignOn) GetSamlAssertionLifetimeSecondsOk() (*int32, bool) {
	if o == nil || o.SamlAssertionLifetimeSeconds == nil {
		return nil, false
	}
	return o.SamlAssertionLifetimeSeconds, true
}

// HasSamlAssertionLifetimeSeconds returns a boolean if a field has been set.
func (o *OINSaml20ApplicationSettingsSignOn) HasSamlAssertionLifetimeSeconds() bool {
	if o != nil && o.SamlAssertionLifetimeSeconds != nil {
		return true
	}

	return false
}

// SetSamlAssertionLifetimeSeconds gets a reference to the given int32 and assigns it to the SamlAssertionLifetimeSeconds field.
func (o *OINSaml20ApplicationSettingsSignOn) SetSamlAssertionLifetimeSeconds(v int32) {
	o.SamlAssertionLifetimeSeconds = &v
}

// GetSsoAcsUrlOverride returns the SsoAcsUrlOverride field value if set, zero value otherwise.
func (o *OINSaml20ApplicationSettingsSignOn) GetSsoAcsUrlOverride() string {
	if o == nil || o.SsoAcsUrlOverride == nil {
		var ret string
		return ret
	}
	return *o.SsoAcsUrlOverride
}

// GetSsoAcsUrlOverrideOk returns a tuple with the SsoAcsUrlOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OINSaml20ApplicationSettingsSignOn) GetSsoAcsUrlOverrideOk() (*string, bool) {
	if o == nil || o.SsoAcsUrlOverride == nil {
		return nil, false
	}
	return o.SsoAcsUrlOverride, true
}

// HasSsoAcsUrlOverride returns a boolean if a field has been set.
func (o *OINSaml20ApplicationSettingsSignOn) HasSsoAcsUrlOverride() bool {
	if o != nil && o.SsoAcsUrlOverride != nil {
		return true
	}

	return false
}

// SetSsoAcsUrlOverride gets a reference to the given string and assigns it to the SsoAcsUrlOverride field.
func (o *OINSaml20ApplicationSettingsSignOn) SetSsoAcsUrlOverride(v string) {
	o.SsoAcsUrlOverride = &v
}

func (o OINSaml20ApplicationSettingsSignOn) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AudienceOverride != nil {
		toSerialize["audienceOverride"] = o.AudienceOverride
	}
	if o.ConfiguredAttributeStatements != nil {
		toSerialize["configuredAttributeStatements"] = o.ConfiguredAttributeStatements
	}
	if o.DefaultRelayState != nil {
		toSerialize["defaultRelayState"] = o.DefaultRelayState
	}
	if o.DestinationOverride != nil {
		toSerialize["destinationOverride"] = o.DestinationOverride
	}
	if o.RecipientOverride != nil {
		toSerialize["recipientOverride"] = o.RecipientOverride
	}
	if o.SamlAssertionLifetimeSeconds != nil {
		toSerialize["samlAssertionLifetimeSeconds"] = o.SamlAssertionLifetimeSeconds
	}
	if o.SsoAcsUrlOverride != nil {
		toSerialize["ssoAcsUrlOverride"] = o.SsoAcsUrlOverride
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OINSaml20ApplicationSettingsSignOn) UnmarshalJSON(bytes []byte) (err error) {
	varOINSaml20ApplicationSettingsSignOn := _OINSaml20ApplicationSettingsSignOn{}

	err = json.Unmarshal(bytes, &varOINSaml20ApplicationSettingsSignOn)
	if err == nil {
		*o = OINSaml20ApplicationSettingsSignOn(varOINSaml20ApplicationSettingsSignOn)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "audienceOverride")
		delete(additionalProperties, "configuredAttributeStatements")
		delete(additionalProperties, "defaultRelayState")
		delete(additionalProperties, "destinationOverride")
		delete(additionalProperties, "recipientOverride")
		delete(additionalProperties, "samlAssertionLifetimeSeconds")
		delete(additionalProperties, "ssoAcsUrlOverride")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableOINSaml20ApplicationSettingsSignOn struct {
	value *OINSaml20ApplicationSettingsSignOn
	isSet bool
}

func (v NullableOINSaml20ApplicationSettingsSignOn) Get() *OINSaml20ApplicationSettingsSignOn {
	return v.value
}

func (v *NullableOINSaml20ApplicationSettingsSignOn) Set(val *OINSaml20ApplicationSettingsSignOn) {
	v.value = val
	v.isSet = true
}

func (v NullableOINSaml20ApplicationSettingsSignOn) IsSet() bool {
	return v.isSet
}

func (v *NullableOINSaml20ApplicationSettingsSignOn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOINSaml20ApplicationSettingsSignOn(val *OINSaml20ApplicationSettingsSignOn) *NullableOINSaml20ApplicationSettingsSignOn {
	return &NullableOINSaml20ApplicationSettingsSignOn{value: val, isSet: true}
}

func (v NullableOINSaml20ApplicationSettingsSignOn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOINSaml20ApplicationSettingsSignOn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

