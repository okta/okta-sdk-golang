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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// OINSaml20ApplicationSettingsSignOn Contains SAML 2.0 sign-on mode attributes. > **Note:** Set `destinationOverride` to configure any other SAML 2.0 attributes in this section.
type OINSaml20ApplicationSettingsSignOn struct {
	// A list of custom attribute statements for the app's SAML assertion. See [SAML 2.0 Technical Overview](https://docs.oasis-open.org/security/saml/Post2.0/sstc-saml-tech-overview-2.0-cd-02.html).  There are two types of attribute statements: | Type | Description | | ---- | ----------- | | EXPRESSION | Generic attribute statement that can be dynamic and supports [Okta Expression Language](https://developer.okta.com/docs/reference/okta-expression-language/) | | GROUP | Group attribute statement | 
	AttributeStatements []SamlAttributeStatement `json:"attributeStatements,omitempty"`
	// Audience override for CASB configuration. See [CASB config guide](https://help.okta.com/en-us/Content/Topics/Apps/CASB-config-guide.htm).
	AudienceOverride NullableString `json:"audienceOverride,omitempty"`
	// Identifies a specific application resource in an IdP-initiated SSO scenario
	DefaultRelayState NullableString `json:"defaultRelayState,omitempty"`
	// Destination override for CASB configuration. See [CASB config guide](https://help.okta.com/en-us/Content/Topics/Apps/CASB-config-guide.htm).
	DestinationOverride NullableString `json:"destinationOverride,omitempty"`
	// Recipient override for CASB configuration. See [CASB config guide](https://help.okta.com/en-us/Content/Topics/Apps/CASB-config-guide.htm).
	RecipientOverride NullableString `json:"recipientOverride,omitempty"`
	// Determines the SAML app session lifetimes with Okta
	SamlAssertionLifetimeSeconds *int32 `json:"samlAssertionLifetimeSeconds,omitempty"`
	// Assertion Consumer Service (ACS) URL override for CASB configuration. See [CASB config guide](https://help.okta.com/en-us/Content/Topics/Apps/CASB-config-guide.htm).
	SsoAcsUrlOverride NullableString `json:"ssoAcsUrlOverride,omitempty"`
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

// GetAttributeStatements returns the AttributeStatements field value if set, zero value otherwise.
func (o *OINSaml20ApplicationSettingsSignOn) GetAttributeStatements() []SamlAttributeStatement {
	if o == nil || o.AttributeStatements == nil {
		var ret []SamlAttributeStatement
		return ret
	}
	return o.AttributeStatements
}

// GetAttributeStatementsOk returns a tuple with the AttributeStatements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OINSaml20ApplicationSettingsSignOn) GetAttributeStatementsOk() ([]SamlAttributeStatement, bool) {
	if o == nil || o.AttributeStatements == nil {
		return nil, false
	}
	return o.AttributeStatements, true
}

// HasAttributeStatements returns a boolean if a field has been set.
func (o *OINSaml20ApplicationSettingsSignOn) HasAttributeStatements() bool {
	if o != nil && o.AttributeStatements != nil {
		return true
	}

	return false
}

// SetAttributeStatements gets a reference to the given []SamlAttributeStatement and assigns it to the AttributeStatements field.
func (o *OINSaml20ApplicationSettingsSignOn) SetAttributeStatements(v []SamlAttributeStatement) {
	o.AttributeStatements = v
}

// GetAudienceOverride returns the AudienceOverride field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OINSaml20ApplicationSettingsSignOn) GetAudienceOverride() string {
	if o == nil || o.AudienceOverride.Get() == nil {
		var ret string
		return ret
	}
	return *o.AudienceOverride.Get()
}

// GetAudienceOverrideOk returns a tuple with the AudienceOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OINSaml20ApplicationSettingsSignOn) GetAudienceOverrideOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AudienceOverride.Get(), o.AudienceOverride.IsSet()
}

// HasAudienceOverride returns a boolean if a field has been set.
func (o *OINSaml20ApplicationSettingsSignOn) HasAudienceOverride() bool {
	if o != nil && o.AudienceOverride.IsSet() {
		return true
	}

	return false
}

// SetAudienceOverride gets a reference to the given NullableString and assigns it to the AudienceOverride field.
func (o *OINSaml20ApplicationSettingsSignOn) SetAudienceOverride(v string) {
	o.AudienceOverride.Set(&v)
}
// SetAudienceOverrideNil sets the value for AudienceOverride to be an explicit nil
func (o *OINSaml20ApplicationSettingsSignOn) SetAudienceOverrideNil() {
	o.AudienceOverride.Set(nil)
}

// UnsetAudienceOverride ensures that no value is present for AudienceOverride, not even an explicit nil
func (o *OINSaml20ApplicationSettingsSignOn) UnsetAudienceOverride() {
	o.AudienceOverride.Unset()
}

// GetDefaultRelayState returns the DefaultRelayState field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OINSaml20ApplicationSettingsSignOn) GetDefaultRelayState() string {
	if o == nil || o.DefaultRelayState.Get() == nil {
		var ret string
		return ret
	}
	return *o.DefaultRelayState.Get()
}

// GetDefaultRelayStateOk returns a tuple with the DefaultRelayState field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OINSaml20ApplicationSettingsSignOn) GetDefaultRelayStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultRelayState.Get(), o.DefaultRelayState.IsSet()
}

// HasDefaultRelayState returns a boolean if a field has been set.
func (o *OINSaml20ApplicationSettingsSignOn) HasDefaultRelayState() bool {
	if o != nil && o.DefaultRelayState.IsSet() {
		return true
	}

	return false
}

// SetDefaultRelayState gets a reference to the given NullableString and assigns it to the DefaultRelayState field.
func (o *OINSaml20ApplicationSettingsSignOn) SetDefaultRelayState(v string) {
	o.DefaultRelayState.Set(&v)
}
// SetDefaultRelayStateNil sets the value for DefaultRelayState to be an explicit nil
func (o *OINSaml20ApplicationSettingsSignOn) SetDefaultRelayStateNil() {
	o.DefaultRelayState.Set(nil)
}

// UnsetDefaultRelayState ensures that no value is present for DefaultRelayState, not even an explicit nil
func (o *OINSaml20ApplicationSettingsSignOn) UnsetDefaultRelayState() {
	o.DefaultRelayState.Unset()
}

// GetDestinationOverride returns the DestinationOverride field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OINSaml20ApplicationSettingsSignOn) GetDestinationOverride() string {
	if o == nil || o.DestinationOverride.Get() == nil {
		var ret string
		return ret
	}
	return *o.DestinationOverride.Get()
}

// GetDestinationOverrideOk returns a tuple with the DestinationOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OINSaml20ApplicationSettingsSignOn) GetDestinationOverrideOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DestinationOverride.Get(), o.DestinationOverride.IsSet()
}

// HasDestinationOverride returns a boolean if a field has been set.
func (o *OINSaml20ApplicationSettingsSignOn) HasDestinationOverride() bool {
	if o != nil && o.DestinationOverride.IsSet() {
		return true
	}

	return false
}

// SetDestinationOverride gets a reference to the given NullableString and assigns it to the DestinationOverride field.
func (o *OINSaml20ApplicationSettingsSignOn) SetDestinationOverride(v string) {
	o.DestinationOverride.Set(&v)
}
// SetDestinationOverrideNil sets the value for DestinationOverride to be an explicit nil
func (o *OINSaml20ApplicationSettingsSignOn) SetDestinationOverrideNil() {
	o.DestinationOverride.Set(nil)
}

// UnsetDestinationOverride ensures that no value is present for DestinationOverride, not even an explicit nil
func (o *OINSaml20ApplicationSettingsSignOn) UnsetDestinationOverride() {
	o.DestinationOverride.Unset()
}

// GetRecipientOverride returns the RecipientOverride field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OINSaml20ApplicationSettingsSignOn) GetRecipientOverride() string {
	if o == nil || o.RecipientOverride.Get() == nil {
		var ret string
		return ret
	}
	return *o.RecipientOverride.Get()
}

// GetRecipientOverrideOk returns a tuple with the RecipientOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OINSaml20ApplicationSettingsSignOn) GetRecipientOverrideOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RecipientOverride.Get(), o.RecipientOverride.IsSet()
}

// HasRecipientOverride returns a boolean if a field has been set.
func (o *OINSaml20ApplicationSettingsSignOn) HasRecipientOverride() bool {
	if o != nil && o.RecipientOverride.IsSet() {
		return true
	}

	return false
}

// SetRecipientOverride gets a reference to the given NullableString and assigns it to the RecipientOverride field.
func (o *OINSaml20ApplicationSettingsSignOn) SetRecipientOverride(v string) {
	o.RecipientOverride.Set(&v)
}
// SetRecipientOverrideNil sets the value for RecipientOverride to be an explicit nil
func (o *OINSaml20ApplicationSettingsSignOn) SetRecipientOverrideNil() {
	o.RecipientOverride.Set(nil)
}

// UnsetRecipientOverride ensures that no value is present for RecipientOverride, not even an explicit nil
func (o *OINSaml20ApplicationSettingsSignOn) UnsetRecipientOverride() {
	o.RecipientOverride.Unset()
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

// GetSsoAcsUrlOverride returns the SsoAcsUrlOverride field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OINSaml20ApplicationSettingsSignOn) GetSsoAcsUrlOverride() string {
	if o == nil || o.SsoAcsUrlOverride.Get() == nil {
		var ret string
		return ret
	}
	return *o.SsoAcsUrlOverride.Get()
}

// GetSsoAcsUrlOverrideOk returns a tuple with the SsoAcsUrlOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OINSaml20ApplicationSettingsSignOn) GetSsoAcsUrlOverrideOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SsoAcsUrlOverride.Get(), o.SsoAcsUrlOverride.IsSet()
}

// HasSsoAcsUrlOverride returns a boolean if a field has been set.
func (o *OINSaml20ApplicationSettingsSignOn) HasSsoAcsUrlOverride() bool {
	if o != nil && o.SsoAcsUrlOverride.IsSet() {
		return true
	}

	return false
}

// SetSsoAcsUrlOverride gets a reference to the given NullableString and assigns it to the SsoAcsUrlOverride field.
func (o *OINSaml20ApplicationSettingsSignOn) SetSsoAcsUrlOverride(v string) {
	o.SsoAcsUrlOverride.Set(&v)
}
// SetSsoAcsUrlOverrideNil sets the value for SsoAcsUrlOverride to be an explicit nil
func (o *OINSaml20ApplicationSettingsSignOn) SetSsoAcsUrlOverrideNil() {
	o.SsoAcsUrlOverride.Set(nil)
}

// UnsetSsoAcsUrlOverride ensures that no value is present for SsoAcsUrlOverride, not even an explicit nil
func (o *OINSaml20ApplicationSettingsSignOn) UnsetSsoAcsUrlOverride() {
	o.SsoAcsUrlOverride.Unset()
}

func (o OINSaml20ApplicationSettingsSignOn) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AttributeStatements != nil {
		toSerialize["attributeStatements"] = o.AttributeStatements
	}
	if o.AudienceOverride.IsSet() {
		toSerialize["audienceOverride"] = o.AudienceOverride.Get()
	}
	if o.DefaultRelayState.IsSet() {
		toSerialize["defaultRelayState"] = o.DefaultRelayState.Get()
	}
	if o.DestinationOverride.IsSet() {
		toSerialize["destinationOverride"] = o.DestinationOverride.Get()
	}
	if o.RecipientOverride.IsSet() {
		toSerialize["recipientOverride"] = o.RecipientOverride.Get()
	}
	if o.SamlAssertionLifetimeSeconds != nil {
		toSerialize["samlAssertionLifetimeSeconds"] = o.SamlAssertionLifetimeSeconds
	}
	if o.SsoAcsUrlOverride.IsSet() {
		toSerialize["ssoAcsUrlOverride"] = o.SsoAcsUrlOverride.Get()
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
		delete(additionalProperties, "attributeStatements")
		delete(additionalProperties, "audienceOverride")
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

