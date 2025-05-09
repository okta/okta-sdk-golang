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

// ProvisioningConnectionOauthRequestProfile struct for ProvisioningConnectionOauthRequestProfile
type ProvisioningConnectionOauthRequestProfile struct {
	// OAuth 2.0 is used to authenticate with the app.
	AuthScheme string `json:"authScheme"`
	// Only used for the Okta Org2Org (`okta_org2org`) app. The unique client identifier for the OAuth 2.0 service app from the target org.
	ClientId *string `json:"clientId,omitempty"`
	Settings *Office365ProvisioningSettings `json:"settings,omitempty"`
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
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningConnectionOauthRequestProfile) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *ProvisioningConnectionOauthRequestProfile) HasClientId() bool {
	if o != nil && o.ClientId != nil {
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
	if o == nil || o.Settings == nil {
		var ret Office365ProvisioningSettings
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningConnectionOauthRequestProfile) GetSettingsOk() (*Office365ProvisioningSettings, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *ProvisioningConnectionOauthRequestProfile) HasSettings() bool {
	if o != nil && o.Settings != nil {
		return true
	}

	return false
}

// SetSettings gets a reference to the given Office365ProvisioningSettings and assigns it to the Settings field.
func (o *ProvisioningConnectionOauthRequestProfile) SetSettings(v Office365ProvisioningSettings) {
	o.Settings = &v
}

func (o ProvisioningConnectionOauthRequestProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["authScheme"] = o.AuthScheme
	}
	if o.ClientId != nil {
		toSerialize["clientId"] = o.ClientId
	}
	if o.Settings != nil {
		toSerialize["settings"] = o.Settings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ProvisioningConnectionOauthRequestProfile) UnmarshalJSON(bytes []byte) (err error) {
	varProvisioningConnectionOauthRequestProfile := _ProvisioningConnectionOauthRequestProfile{}

	err = json.Unmarshal(bytes, &varProvisioningConnectionOauthRequestProfile)
	if err == nil {
		*o = ProvisioningConnectionOauthRequestProfile(varProvisioningConnectionOauthRequestProfile)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "authScheme")
		delete(additionalProperties, "clientId")
		delete(additionalProperties, "settings")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
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

