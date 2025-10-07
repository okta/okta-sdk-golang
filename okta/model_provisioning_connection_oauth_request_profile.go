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
	"fmt"
)

// checks if the ProvisioningConnectionOauthRequestProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProvisioningConnectionOauthRequestProfile{}

// ProvisioningConnectionOauthRequestProfile struct for ProvisioningConnectionOauthRequestProfile
type ProvisioningConnectionOauthRequestProfile struct {
	// OAuth 2.0 is used to authenticate with the app.
	AuthScheme string `json:"authScheme"`
	// Only used for the Okta Org2Org (`okta_org2org`) app. The unique client identifier for the OAuth 2.0 service app from the target org.
	ClientId             *string                                  `json:"clientId,omitempty"`
	Settings             *Office365ProvisioningSettings           `json:"settings,omitempty"`
	Signing              *Org2OrgProvisioningOAuthSigningSettings `json:"signing,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProvisioningConnectionOauthRequestProfile ProvisioningConnectionOauthRequestProfile

// NewProvisioningConnectionOauthRequestProfile instantiates a new ProvisioningConnectionOauthRequestProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvisioningConnectionOauthRequestProfile(authScheme string) *ProvisioningConnectionOauthRequestProfile {
	this := ProvisioningConnectionOauthRequestProfile{}
	this.AuthScheme = authScheme
	return &this
}

// NewProvisioningConnectionOauthRequestProfileWithDefaults instantiates a new ProvisioningConnectionOauthRequestProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvisioningConnectionOauthRequestProfileWithDefaults() *ProvisioningConnectionOauthRequestProfile {
	this := ProvisioningConnectionOauthRequestProfile{}
	return &this
}

// GetAuthScheme returns the AuthScheme field value
func (o *ProvisioningConnectionOauthRequestProfile) GetAuthScheme() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthScheme
}

// GetAuthSchemeOk returns a tuple with the AuthScheme field value
// and a boolean to check if the value has been set.
func (o *ProvisioningConnectionOauthRequestProfile) GetAuthSchemeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthScheme, true
}

// SetAuthScheme sets field value
func (o *ProvisioningConnectionOauthRequestProfile) SetAuthScheme(v string) {
	o.AuthScheme = v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *ProvisioningConnectionOauthRequestProfile) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningConnectionOauthRequestProfile) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *ProvisioningConnectionOauthRequestProfile) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *ProvisioningConnectionOauthRequestProfile) SetClientId(v string) {
	o.ClientId = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *ProvisioningConnectionOauthRequestProfile) GetSettings() Office365ProvisioningSettings {
	if o == nil || IsNil(o.Settings) {
		var ret Office365ProvisioningSettings
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningConnectionOauthRequestProfile) GetSettingsOk() (*Office365ProvisioningSettings, bool) {
	if o == nil || IsNil(o.Settings) {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *ProvisioningConnectionOauthRequestProfile) HasSettings() bool {
	if o != nil && !IsNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given Office365ProvisioningSettings and assigns it to the Settings field.
func (o *ProvisioningConnectionOauthRequestProfile) SetSettings(v Office365ProvisioningSettings) {
	o.Settings = &v
}

// GetSigning returns the Signing field value if set, zero value otherwise.
func (o *ProvisioningConnectionOauthRequestProfile) GetSigning() Org2OrgProvisioningOAuthSigningSettings {
	if o == nil || IsNil(o.Signing) {
		var ret Org2OrgProvisioningOAuthSigningSettings
		return ret
	}
	return *o.Signing
}

// GetSigningOk returns a tuple with the Signing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningConnectionOauthRequestProfile) GetSigningOk() (*Org2OrgProvisioningOAuthSigningSettings, bool) {
	if o == nil || IsNil(o.Signing) {
		return nil, false
	}
	return o.Signing, true
}

// HasSigning returns a boolean if a field has been set.
func (o *ProvisioningConnectionOauthRequestProfile) HasSigning() bool {
	if o != nil && !IsNil(o.Signing) {
		return true
	}

	return false
}

// SetSigning gets a reference to the given Org2OrgProvisioningOAuthSigningSettings and assigns it to the Signing field.
func (o *ProvisioningConnectionOauthRequestProfile) SetSigning(v Org2OrgProvisioningOAuthSigningSettings) {
	o.Signing = &v
}

func (o ProvisioningConnectionOauthRequestProfile) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProvisioningConnectionOauthRequestProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["authScheme"] = o.AuthScheme
	if !IsNil(o.ClientId) {
		toSerialize["clientId"] = o.ClientId
	}
	if !IsNil(o.Settings) {
		toSerialize["settings"] = o.Settings
	}
	if !IsNil(o.Signing) {
		toSerialize["signing"] = o.Signing
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProvisioningConnectionOauthRequestProfile) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"authScheme",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varProvisioningConnectionOauthRequestProfile := _ProvisioningConnectionOauthRequestProfile{}

	err = json.Unmarshal(data, &varProvisioningConnectionOauthRequestProfile)

	if err != nil {
		return err
	}

	*o = ProvisioningConnectionOauthRequestProfile(varProvisioningConnectionOauthRequestProfile)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "authScheme")
		delete(additionalProperties, "clientId")
		delete(additionalProperties, "settings")
		delete(additionalProperties, "signing")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProvisioningConnectionOauthRequestProfile struct {
	value *ProvisioningConnectionOauthRequestProfile
	isSet bool
}

func (v NullableProvisioningConnectionOauthRequestProfile) Get() *ProvisioningConnectionOauthRequestProfile {
	return v.value
}

func (v *NullableProvisioningConnectionOauthRequestProfile) Set(val *ProvisioningConnectionOauthRequestProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisioningConnectionOauthRequestProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisioningConnectionOauthRequestProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisioningConnectionOauthRequestProfile(val *ProvisioningConnectionOauthRequestProfile) *NullableProvisioningConnectionOauthRequestProfile {
	return &NullableProvisioningConnectionOauthRequestProfile{value: val, isSet: true}
}

func (v NullableProvisioningConnectionOauthRequestProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisioningConnectionOauthRequestProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
