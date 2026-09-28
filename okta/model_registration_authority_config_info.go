/*
Okta Admin Management API

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

// checks if the RegistrationAuthorityConfigInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegistrationAuthorityConfigInfo{}

// RegistrationAuthorityConfigInfo Protocol-specific settings for the configuration. It's an empty object for a static or dynamic challenge, and carries the Microsoft Entra ID application settings for a delegated Microsoft Intune challenge.  Okta needs `provider`, `aadClientId`, `aadClientSecret`, and `aadAuthority` to validate Intune SCEP requests, and rejects a request that sends any property not listed here with a `400`. The v1 create operations require all four. A renewal requires none of them: each one falls back to the configuration being cloned, so you send only what you want to change — see `ParallelScepConfig`.
type RegistrationAuthorityConfigInfo struct {
	// The Microsoft Entra ID tenant domain that Okta authenticates against, for example `example.onmicrosoft.com`
	AadAuthority *string `json:"aadAuthority,omitempty"`
	// The client ID of the Microsoft Entra ID application that Okta uses to validate Microsoft Intune SCEP requests
	AadClientId *string `json:"aadClientId,omitempty"`
	// The client secret of the Microsoft Entra ID application. Write-only — it's never returned in a response.
	AadClientSecret *string `json:"aadClientSecret,omitempty"`
	// The mobile device management provider that validates the challenge. Use `INTUNE` for Microsoft Intune.
	Provider *string `json:"provider,omitempty"`
}

// NewRegistrationAuthorityConfigInfo instantiates a new RegistrationAuthorityConfigInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistrationAuthorityConfigInfo() *RegistrationAuthorityConfigInfo {
	this := RegistrationAuthorityConfigInfo{}
	return &this
}

// NewRegistrationAuthorityConfigInfoWithDefaults instantiates a new RegistrationAuthorityConfigInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistrationAuthorityConfigInfoWithDefaults() *RegistrationAuthorityConfigInfo {
	this := RegistrationAuthorityConfigInfo{}
	return &this
}

// GetAadAuthority returns the AadAuthority field value if set, zero value otherwise.
func (o *RegistrationAuthorityConfigInfo) GetAadAuthority() string {
	if o == nil || IsNil(o.AadAuthority) {
		var ret string
		return ret
	}
	return *o.AadAuthority
}

// GetAadAuthorityOk returns a tuple with the AadAuthority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationAuthorityConfigInfo) GetAadAuthorityOk() (*string, bool) {
	if o == nil || IsNil(o.AadAuthority) {
		return nil, false
	}
	return o.AadAuthority, true
}

// HasAadAuthority returns a boolean if a field has been set.
func (o *RegistrationAuthorityConfigInfo) HasAadAuthority() bool {
	if o != nil && !IsNil(o.AadAuthority) {
		return true
	}

	return false
}

// SetAadAuthority gets a reference to the given string and assigns it to the AadAuthority field.
func (o *RegistrationAuthorityConfigInfo) SetAadAuthority(v string) {
	o.AadAuthority = &v
}

// GetAadClientId returns the AadClientId field value if set, zero value otherwise.
func (o *RegistrationAuthorityConfigInfo) GetAadClientId() string {
	if o == nil || IsNil(o.AadClientId) {
		var ret string
		return ret
	}
	return *o.AadClientId
}

// GetAadClientIdOk returns a tuple with the AadClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationAuthorityConfigInfo) GetAadClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.AadClientId) {
		return nil, false
	}
	return o.AadClientId, true
}

// HasAadClientId returns a boolean if a field has been set.
func (o *RegistrationAuthorityConfigInfo) HasAadClientId() bool {
	if o != nil && !IsNil(o.AadClientId) {
		return true
	}

	return false
}

// SetAadClientId gets a reference to the given string and assigns it to the AadClientId field.
func (o *RegistrationAuthorityConfigInfo) SetAadClientId(v string) {
	o.AadClientId = &v
}

// GetAadClientSecret returns the AadClientSecret field value if set, zero value otherwise.
func (o *RegistrationAuthorityConfigInfo) GetAadClientSecret() string {
	if o == nil || IsNil(o.AadClientSecret) {
		var ret string
		return ret
	}
	return *o.AadClientSecret
}

// GetAadClientSecretOk returns a tuple with the AadClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationAuthorityConfigInfo) GetAadClientSecretOk() (*string, bool) {
	if o == nil || IsNil(o.AadClientSecret) {
		return nil, false
	}
	return o.AadClientSecret, true
}

// HasAadClientSecret returns a boolean if a field has been set.
func (o *RegistrationAuthorityConfigInfo) HasAadClientSecret() bool {
	if o != nil && !IsNil(o.AadClientSecret) {
		return true
	}

	return false
}

// SetAadClientSecret gets a reference to the given string and assigns it to the AadClientSecret field.
func (o *RegistrationAuthorityConfigInfo) SetAadClientSecret(v string) {
	o.AadClientSecret = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *RegistrationAuthorityConfigInfo) GetProvider() string {
	if o == nil || IsNil(o.Provider) {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationAuthorityConfigInfo) GetProviderOk() (*string, bool) {
	if o == nil || IsNil(o.Provider) {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *RegistrationAuthorityConfigInfo) HasProvider() bool {
	if o != nil && !IsNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *RegistrationAuthorityConfigInfo) SetProvider(v string) {
	o.Provider = &v
}

func (o RegistrationAuthorityConfigInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegistrationAuthorityConfigInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AadAuthority) {
		toSerialize["aadAuthority"] = o.AadAuthority
	}
	if !IsNil(o.AadClientId) {
		toSerialize["aadClientId"] = o.AadClientId
	}
	if !IsNil(o.AadClientSecret) {
		toSerialize["aadClientSecret"] = o.AadClientSecret
	}
	if !IsNil(o.Provider) {
		toSerialize["provider"] = o.Provider
	}
	return toSerialize, nil
}

type NullableRegistrationAuthorityConfigInfo struct {
	value *RegistrationAuthorityConfigInfo
	isSet bool
}

func (v NullableRegistrationAuthorityConfigInfo) Get() *RegistrationAuthorityConfigInfo {
	return v.value
}

func (v *NullableRegistrationAuthorityConfigInfo) Set(val *RegistrationAuthorityConfigInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistrationAuthorityConfigInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistrationAuthorityConfigInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistrationAuthorityConfigInfo(val *RegistrationAuthorityConfigInfo) *NullableRegistrationAuthorityConfigInfo {
	return &NullableRegistrationAuthorityConfigInfo{value: val, isSet: true}
}

func (v NullableRegistrationAuthorityConfigInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistrationAuthorityConfigInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
