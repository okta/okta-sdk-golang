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

// checks if the CustomTelephonyProviderCredentialResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomTelephonyProviderCredentialResponse{}

// CustomTelephonyProviderCredentialResponse struct for CustomTelephonyProviderCredentialResponse
type CustomTelephonyProviderCredentialResponse struct {
	// Indicates whether the provider is enabled and can be used
	Enabled *bool `json:"enabled,omitempty"`
	// ID of the custom telephony provider
	Id *string `json:"id,omitempty"`
	// Indicates whether the provider is the primary telephony provider
	IsPrimaryProvider *bool `json:"isPrimaryProvider,omitempty"`
	// The types of telephony operations (SMS or Voice) that you use with your telephony provider.  `ALL` is the only valid value. It indicates that your provider can handle both SMS messages and voice calls. You're not required to use both types of telephony operations, but your provider can support both.
	ProviderCapability *string `json:"providerCapability,omitempty"`
	// Name of the telephony provider
	ProviderName     *string                          `json:"providerName,omitempty"`
	ProviderSettings *CustomTelephonyProviderSettings `json:"providerSettings,omitempty"`
	// The account string identifier (SID) for your telephony provider account. Your telephony provider gives you this SID.
	ProviderSid          *string `json:"providerSid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CustomTelephonyProviderCredentialResponse CustomTelephonyProviderCredentialResponse

// NewCustomTelephonyProviderCredentialResponse instantiates a new CustomTelephonyProviderCredentialResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomTelephonyProviderCredentialResponse() *CustomTelephonyProviderCredentialResponse {
	this := CustomTelephonyProviderCredentialResponse{}
	return &this
}

// NewCustomTelephonyProviderCredentialResponseWithDefaults instantiates a new CustomTelephonyProviderCredentialResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomTelephonyProviderCredentialResponseWithDefaults() *CustomTelephonyProviderCredentialResponse {
	this := CustomTelephonyProviderCredentialResponse{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CustomTelephonyProviderCredentialResponse) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomTelephonyProviderCredentialResponse) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CustomTelephonyProviderCredentialResponse) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CustomTelephonyProviderCredentialResponse) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CustomTelephonyProviderCredentialResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomTelephonyProviderCredentialResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CustomTelephonyProviderCredentialResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CustomTelephonyProviderCredentialResponse) SetId(v string) {
	o.Id = &v
}

// GetIsPrimaryProvider returns the IsPrimaryProvider field value if set, zero value otherwise.
func (o *CustomTelephonyProviderCredentialResponse) GetIsPrimaryProvider() bool {
	if o == nil || IsNil(o.IsPrimaryProvider) {
		var ret bool
		return ret
	}
	return *o.IsPrimaryProvider
}

// GetIsPrimaryProviderOk returns a tuple with the IsPrimaryProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomTelephonyProviderCredentialResponse) GetIsPrimaryProviderOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPrimaryProvider) {
		return nil, false
	}
	return o.IsPrimaryProvider, true
}

// HasIsPrimaryProvider returns a boolean if a field has been set.
func (o *CustomTelephonyProviderCredentialResponse) HasIsPrimaryProvider() bool {
	if o != nil && !IsNil(o.IsPrimaryProvider) {
		return true
	}

	return false
}

// SetIsPrimaryProvider gets a reference to the given bool and assigns it to the IsPrimaryProvider field.
func (o *CustomTelephonyProviderCredentialResponse) SetIsPrimaryProvider(v bool) {
	o.IsPrimaryProvider = &v
}

// GetProviderCapability returns the ProviderCapability field value if set, zero value otherwise.
func (o *CustomTelephonyProviderCredentialResponse) GetProviderCapability() string {
	if o == nil || IsNil(o.ProviderCapability) {
		var ret string
		return ret
	}
	return *o.ProviderCapability
}

// GetProviderCapabilityOk returns a tuple with the ProviderCapability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomTelephonyProviderCredentialResponse) GetProviderCapabilityOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderCapability) {
		return nil, false
	}
	return o.ProviderCapability, true
}

// HasProviderCapability returns a boolean if a field has been set.
func (o *CustomTelephonyProviderCredentialResponse) HasProviderCapability() bool {
	if o != nil && !IsNil(o.ProviderCapability) {
		return true
	}

	return false
}

// SetProviderCapability gets a reference to the given string and assigns it to the ProviderCapability field.
func (o *CustomTelephonyProviderCredentialResponse) SetProviderCapability(v string) {
	o.ProviderCapability = &v
}

// GetProviderName returns the ProviderName field value if set, zero value otherwise.
func (o *CustomTelephonyProviderCredentialResponse) GetProviderName() string {
	if o == nil || IsNil(o.ProviderName) {
		var ret string
		return ret
	}
	return *o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomTelephonyProviderCredentialResponse) GetProviderNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderName) {
		return nil, false
	}
	return o.ProviderName, true
}

// HasProviderName returns a boolean if a field has been set.
func (o *CustomTelephonyProviderCredentialResponse) HasProviderName() bool {
	if o != nil && !IsNil(o.ProviderName) {
		return true
	}

	return false
}

// SetProviderName gets a reference to the given string and assigns it to the ProviderName field.
func (o *CustomTelephonyProviderCredentialResponse) SetProviderName(v string) {
	o.ProviderName = &v
}

// GetProviderSettings returns the ProviderSettings field value if set, zero value otherwise.
func (o *CustomTelephonyProviderCredentialResponse) GetProviderSettings() CustomTelephonyProviderSettings {
	if o == nil || IsNil(o.ProviderSettings) {
		var ret CustomTelephonyProviderSettings
		return ret
	}
	return *o.ProviderSettings
}

// GetProviderSettingsOk returns a tuple with the ProviderSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomTelephonyProviderCredentialResponse) GetProviderSettingsOk() (*CustomTelephonyProviderSettings, bool) {
	if o == nil || IsNil(o.ProviderSettings) {
		return nil, false
	}
	return o.ProviderSettings, true
}

// HasProviderSettings returns a boolean if a field has been set.
func (o *CustomTelephonyProviderCredentialResponse) HasProviderSettings() bool {
	if o != nil && !IsNil(o.ProviderSettings) {
		return true
	}

	return false
}

// SetProviderSettings gets a reference to the given CustomTelephonyProviderSettings and assigns it to the ProviderSettings field.
func (o *CustomTelephonyProviderCredentialResponse) SetProviderSettings(v CustomTelephonyProviderSettings) {
	o.ProviderSettings = &v
}

// GetProviderSid returns the ProviderSid field value if set, zero value otherwise.
func (o *CustomTelephonyProviderCredentialResponse) GetProviderSid() string {
	if o == nil || IsNil(o.ProviderSid) {
		var ret string
		return ret
	}
	return *o.ProviderSid
}

// GetProviderSidOk returns a tuple with the ProviderSid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomTelephonyProviderCredentialResponse) GetProviderSidOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderSid) {
		return nil, false
	}
	return o.ProviderSid, true
}

// HasProviderSid returns a boolean if a field has been set.
func (o *CustomTelephonyProviderCredentialResponse) HasProviderSid() bool {
	if o != nil && !IsNil(o.ProviderSid) {
		return true
	}

	return false
}

// SetProviderSid gets a reference to the given string and assigns it to the ProviderSid field.
func (o *CustomTelephonyProviderCredentialResponse) SetProviderSid(v string) {
	o.ProviderSid = &v
}

func (o CustomTelephonyProviderCredentialResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomTelephonyProviderCredentialResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IsPrimaryProvider) {
		toSerialize["isPrimaryProvider"] = o.IsPrimaryProvider
	}
	if !IsNil(o.ProviderCapability) {
		toSerialize["providerCapability"] = o.ProviderCapability
	}
	if !IsNil(o.ProviderName) {
		toSerialize["providerName"] = o.ProviderName
	}
	if !IsNil(o.ProviderSettings) {
		toSerialize["providerSettings"] = o.ProviderSettings
	}
	if !IsNil(o.ProviderSid) {
		toSerialize["providerSid"] = o.ProviderSid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CustomTelephonyProviderCredentialResponse) UnmarshalJSON(data []byte) (err error) {
	varCustomTelephonyProviderCredentialResponse := _CustomTelephonyProviderCredentialResponse{}

	err = json.Unmarshal(data, &varCustomTelephonyProviderCredentialResponse)

	if err != nil {
		return err
	}

	*o = CustomTelephonyProviderCredentialResponse(varCustomTelephonyProviderCredentialResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "id")
		delete(additionalProperties, "isPrimaryProvider")
		delete(additionalProperties, "providerCapability")
		delete(additionalProperties, "providerName")
		delete(additionalProperties, "providerSettings")
		delete(additionalProperties, "providerSid")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCustomTelephonyProviderCredentialResponse struct {
	value *CustomTelephonyProviderCredentialResponse
	isSet bool
}

func (v NullableCustomTelephonyProviderCredentialResponse) Get() *CustomTelephonyProviderCredentialResponse {
	return v.value
}

func (v *NullableCustomTelephonyProviderCredentialResponse) Set(val *CustomTelephonyProviderCredentialResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomTelephonyProviderCredentialResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomTelephonyProviderCredentialResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomTelephonyProviderCredentialResponse(val *CustomTelephonyProviderCredentialResponse) *NullableCustomTelephonyProviderCredentialResponse {
	return &NullableCustomTelephonyProviderCredentialResponse{value: val, isSet: true}
}

func (v NullableCustomTelephonyProviderCredentialResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomTelephonyProviderCredentialResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
