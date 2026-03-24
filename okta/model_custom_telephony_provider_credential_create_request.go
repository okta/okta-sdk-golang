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

// checks if the CustomTelephonyProviderCredentialCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomTelephonyProviderCredentialCreateRequest{}

// CustomTelephonyProviderCredentialCreateRequest Create custom telephony provider credentials
type CustomTelephonyProviderCredentialCreateRequest struct {
	// The authentication token that's used to authenticate requests to the telephony provider. Your telephony provider gives you this token.
	ProviderAuthToken *string `json:"providerAuthToken,omitempty"`
	// The types of telephony operations (SMS or Voice) that you use with your telephony provider.  `ALL` is the only valid value. It indicates that your provider can handle both SMS messages and voice calls. You're not required to use both types of telephony operations, but your provider can support both.
	ProviderCapability *string `json:"providerCapability,omitempty"`
	// The name of the telephony provider
	ProviderName     *string                          `json:"providerName,omitempty"`
	ProviderSettings *CustomTelephonyProviderSettings `json:"providerSettings,omitempty"`
	// The account string identifier (SID) for your telephony provider account. Your telephony provider gives you this SID.
	ProviderSid          *string `json:"providerSid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CustomTelephonyProviderCredentialCreateRequest CustomTelephonyProviderCredentialCreateRequest

// NewCustomTelephonyProviderCredentialCreateRequest instantiates a new CustomTelephonyProviderCredentialCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomTelephonyProviderCredentialCreateRequest() *CustomTelephonyProviderCredentialCreateRequest {
	this := CustomTelephonyProviderCredentialCreateRequest{}
	return &this
}

// NewCustomTelephonyProviderCredentialCreateRequestWithDefaults instantiates a new CustomTelephonyProviderCredentialCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomTelephonyProviderCredentialCreateRequestWithDefaults() *CustomTelephonyProviderCredentialCreateRequest {
	this := CustomTelephonyProviderCredentialCreateRequest{}
	return &this
}

// GetProviderAuthToken returns the ProviderAuthToken field value if set, zero value otherwise.
func (o *CustomTelephonyProviderCredentialCreateRequest) GetProviderAuthToken() string {
	if o == nil || IsNil(o.ProviderAuthToken) {
		var ret string
		return ret
	}
	return *o.ProviderAuthToken
}

// GetProviderAuthTokenOk returns a tuple with the ProviderAuthToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomTelephonyProviderCredentialCreateRequest) GetProviderAuthTokenOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderAuthToken) {
		return nil, false
	}
	return o.ProviderAuthToken, true
}

// HasProviderAuthToken returns a boolean if a field has been set.
func (o *CustomTelephonyProviderCredentialCreateRequest) HasProviderAuthToken() bool {
	if o != nil && !IsNil(o.ProviderAuthToken) {
		return true
	}

	return false
}

// SetProviderAuthToken gets a reference to the given string and assigns it to the ProviderAuthToken field.
func (o *CustomTelephonyProviderCredentialCreateRequest) SetProviderAuthToken(v string) {
	o.ProviderAuthToken = &v
}

// GetProviderCapability returns the ProviderCapability field value if set, zero value otherwise.
func (o *CustomTelephonyProviderCredentialCreateRequest) GetProviderCapability() string {
	if o == nil || IsNil(o.ProviderCapability) {
		var ret string
		return ret
	}
	return *o.ProviderCapability
}

// GetProviderCapabilityOk returns a tuple with the ProviderCapability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomTelephonyProviderCredentialCreateRequest) GetProviderCapabilityOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderCapability) {
		return nil, false
	}
	return o.ProviderCapability, true
}

// HasProviderCapability returns a boolean if a field has been set.
func (o *CustomTelephonyProviderCredentialCreateRequest) HasProviderCapability() bool {
	if o != nil && !IsNil(o.ProviderCapability) {
		return true
	}

	return false
}

// SetProviderCapability gets a reference to the given string and assigns it to the ProviderCapability field.
func (o *CustomTelephonyProviderCredentialCreateRequest) SetProviderCapability(v string) {
	o.ProviderCapability = &v
}

// GetProviderName returns the ProviderName field value if set, zero value otherwise.
func (o *CustomTelephonyProviderCredentialCreateRequest) GetProviderName() string {
	if o == nil || IsNil(o.ProviderName) {
		var ret string
		return ret
	}
	return *o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomTelephonyProviderCredentialCreateRequest) GetProviderNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderName) {
		return nil, false
	}
	return o.ProviderName, true
}

// HasProviderName returns a boolean if a field has been set.
func (o *CustomTelephonyProviderCredentialCreateRequest) HasProviderName() bool {
	if o != nil && !IsNil(o.ProviderName) {
		return true
	}

	return false
}

// SetProviderName gets a reference to the given string and assigns it to the ProviderName field.
func (o *CustomTelephonyProviderCredentialCreateRequest) SetProviderName(v string) {
	o.ProviderName = &v
}

// GetProviderSettings returns the ProviderSettings field value if set, zero value otherwise.
func (o *CustomTelephonyProviderCredentialCreateRequest) GetProviderSettings() CustomTelephonyProviderSettings {
	if o == nil || IsNil(o.ProviderSettings) {
		var ret CustomTelephonyProviderSettings
		return ret
	}
	return *o.ProviderSettings
}

// GetProviderSettingsOk returns a tuple with the ProviderSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomTelephonyProviderCredentialCreateRequest) GetProviderSettingsOk() (*CustomTelephonyProviderSettings, bool) {
	if o == nil || IsNil(o.ProviderSettings) {
		return nil, false
	}
	return o.ProviderSettings, true
}

// HasProviderSettings returns a boolean if a field has been set.
func (o *CustomTelephonyProviderCredentialCreateRequest) HasProviderSettings() bool {
	if o != nil && !IsNil(o.ProviderSettings) {
		return true
	}

	return false
}

// SetProviderSettings gets a reference to the given CustomTelephonyProviderSettings and assigns it to the ProviderSettings field.
func (o *CustomTelephonyProviderCredentialCreateRequest) SetProviderSettings(v CustomTelephonyProviderSettings) {
	o.ProviderSettings = &v
}

// GetProviderSid returns the ProviderSid field value if set, zero value otherwise.
func (o *CustomTelephonyProviderCredentialCreateRequest) GetProviderSid() string {
	if o == nil || IsNil(o.ProviderSid) {
		var ret string
		return ret
	}
	return *o.ProviderSid
}

// GetProviderSidOk returns a tuple with the ProviderSid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomTelephonyProviderCredentialCreateRequest) GetProviderSidOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderSid) {
		return nil, false
	}
	return o.ProviderSid, true
}

// HasProviderSid returns a boolean if a field has been set.
func (o *CustomTelephonyProviderCredentialCreateRequest) HasProviderSid() bool {
	if o != nil && !IsNil(o.ProviderSid) {
		return true
	}

	return false
}

// SetProviderSid gets a reference to the given string and assigns it to the ProviderSid field.
func (o *CustomTelephonyProviderCredentialCreateRequest) SetProviderSid(v string) {
	o.ProviderSid = &v
}

func (o CustomTelephonyProviderCredentialCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomTelephonyProviderCredentialCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProviderAuthToken) {
		toSerialize["providerAuthToken"] = o.ProviderAuthToken
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

func (o *CustomTelephonyProviderCredentialCreateRequest) UnmarshalJSON(data []byte) (err error) {
	varCustomTelephonyProviderCredentialCreateRequest := _CustomTelephonyProviderCredentialCreateRequest{}

	err = json.Unmarshal(data, &varCustomTelephonyProviderCredentialCreateRequest)

	if err != nil {
		return err
	}

	*o = CustomTelephonyProviderCredentialCreateRequest(varCustomTelephonyProviderCredentialCreateRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "providerAuthToken")
		delete(additionalProperties, "providerCapability")
		delete(additionalProperties, "providerName")
		delete(additionalProperties, "providerSettings")
		delete(additionalProperties, "providerSid")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCustomTelephonyProviderCredentialCreateRequest struct {
	value *CustomTelephonyProviderCredentialCreateRequest
	isSet bool
}

func (v NullableCustomTelephonyProviderCredentialCreateRequest) Get() *CustomTelephonyProviderCredentialCreateRequest {
	return v.value
}

func (v *NullableCustomTelephonyProviderCredentialCreateRequest) Set(val *CustomTelephonyProviderCredentialCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomTelephonyProviderCredentialCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomTelephonyProviderCredentialCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomTelephonyProviderCredentialCreateRequest(val *CustomTelephonyProviderCredentialCreateRequest) *NullableCustomTelephonyProviderCredentialCreateRequest {
	return &NullableCustomTelephonyProviderCredentialCreateRequest{value: val, isSet: true}
}

func (v NullableCustomTelephonyProviderCredentialCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomTelephonyProviderCredentialCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
