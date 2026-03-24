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

// checks if the CustomTelephonyProviderCredentialUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomTelephonyProviderCredentialUpdateRequest{}

// CustomTelephonyProviderCredentialUpdateRequest Update custom telephony provider credentials
type CustomTelephonyProviderCredentialUpdateRequest struct {
	// ID of the custom telephony provider
	Id *string `json:"id,omitempty"`
	// The authentication token that's used to authenticate requests to the telephony provider. Your telephony provider gives you this token.
	ProviderAuthToken *string                          `json:"providerAuthToken,omitempty"`
	ProviderSettings  *CustomTelephonyProviderSettings `json:"providerSettings,omitempty"`
	// The account string identifier (SID) for your telephony provider account. Your telephony provider gives you this SID.
	ProviderSid          *string `json:"providerSid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CustomTelephonyProviderCredentialUpdateRequest CustomTelephonyProviderCredentialUpdateRequest

// NewCustomTelephonyProviderCredentialUpdateRequest instantiates a new CustomTelephonyProviderCredentialUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomTelephonyProviderCredentialUpdateRequest() *CustomTelephonyProviderCredentialUpdateRequest {
	this := CustomTelephonyProviderCredentialUpdateRequest{}
	return &this
}

// NewCustomTelephonyProviderCredentialUpdateRequestWithDefaults instantiates a new CustomTelephonyProviderCredentialUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomTelephonyProviderCredentialUpdateRequestWithDefaults() *CustomTelephonyProviderCredentialUpdateRequest {
	this := CustomTelephonyProviderCredentialUpdateRequest{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CustomTelephonyProviderCredentialUpdateRequest) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomTelephonyProviderCredentialUpdateRequest) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CustomTelephonyProviderCredentialUpdateRequest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CustomTelephonyProviderCredentialUpdateRequest) SetId(v string) {
	o.Id = &v
}

// GetProviderAuthToken returns the ProviderAuthToken field value if set, zero value otherwise.
func (o *CustomTelephonyProviderCredentialUpdateRequest) GetProviderAuthToken() string {
	if o == nil || IsNil(o.ProviderAuthToken) {
		var ret string
		return ret
	}
	return *o.ProviderAuthToken
}

// GetProviderAuthTokenOk returns a tuple with the ProviderAuthToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomTelephonyProviderCredentialUpdateRequest) GetProviderAuthTokenOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderAuthToken) {
		return nil, false
	}
	return o.ProviderAuthToken, true
}

// HasProviderAuthToken returns a boolean if a field has been set.
func (o *CustomTelephonyProviderCredentialUpdateRequest) HasProviderAuthToken() bool {
	if o != nil && !IsNil(o.ProviderAuthToken) {
		return true
	}

	return false
}

// SetProviderAuthToken gets a reference to the given string and assigns it to the ProviderAuthToken field.
func (o *CustomTelephonyProviderCredentialUpdateRequest) SetProviderAuthToken(v string) {
	o.ProviderAuthToken = &v
}

// GetProviderSettings returns the ProviderSettings field value if set, zero value otherwise.
func (o *CustomTelephonyProviderCredentialUpdateRequest) GetProviderSettings() CustomTelephonyProviderSettings {
	if o == nil || IsNil(o.ProviderSettings) {
		var ret CustomTelephonyProviderSettings
		return ret
	}
	return *o.ProviderSettings
}

// GetProviderSettingsOk returns a tuple with the ProviderSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomTelephonyProviderCredentialUpdateRequest) GetProviderSettingsOk() (*CustomTelephonyProviderSettings, bool) {
	if o == nil || IsNil(o.ProviderSettings) {
		return nil, false
	}
	return o.ProviderSettings, true
}

// HasProviderSettings returns a boolean if a field has been set.
func (o *CustomTelephonyProviderCredentialUpdateRequest) HasProviderSettings() bool {
	if o != nil && !IsNil(o.ProviderSettings) {
		return true
	}

	return false
}

// SetProviderSettings gets a reference to the given CustomTelephonyProviderSettings and assigns it to the ProviderSettings field.
func (o *CustomTelephonyProviderCredentialUpdateRequest) SetProviderSettings(v CustomTelephonyProviderSettings) {
	o.ProviderSettings = &v
}

// GetProviderSid returns the ProviderSid field value if set, zero value otherwise.
func (o *CustomTelephonyProviderCredentialUpdateRequest) GetProviderSid() string {
	if o == nil || IsNil(o.ProviderSid) {
		var ret string
		return ret
	}
	return *o.ProviderSid
}

// GetProviderSidOk returns a tuple with the ProviderSid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomTelephonyProviderCredentialUpdateRequest) GetProviderSidOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderSid) {
		return nil, false
	}
	return o.ProviderSid, true
}

// HasProviderSid returns a boolean if a field has been set.
func (o *CustomTelephonyProviderCredentialUpdateRequest) HasProviderSid() bool {
	if o != nil && !IsNil(o.ProviderSid) {
		return true
	}

	return false
}

// SetProviderSid gets a reference to the given string and assigns it to the ProviderSid field.
func (o *CustomTelephonyProviderCredentialUpdateRequest) SetProviderSid(v string) {
	o.ProviderSid = &v
}

func (o CustomTelephonyProviderCredentialUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomTelephonyProviderCredentialUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ProviderAuthToken) {
		toSerialize["providerAuthToken"] = o.ProviderAuthToken
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

func (o *CustomTelephonyProviderCredentialUpdateRequest) UnmarshalJSON(data []byte) (err error) {
	varCustomTelephonyProviderCredentialUpdateRequest := _CustomTelephonyProviderCredentialUpdateRequest{}

	err = json.Unmarshal(data, &varCustomTelephonyProviderCredentialUpdateRequest)

	if err != nil {
		return err
	}

	*o = CustomTelephonyProviderCredentialUpdateRequest(varCustomTelephonyProviderCredentialUpdateRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "providerAuthToken")
		delete(additionalProperties, "providerSettings")
		delete(additionalProperties, "providerSid")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCustomTelephonyProviderCredentialUpdateRequest struct {
	value *CustomTelephonyProviderCredentialUpdateRequest
	isSet bool
}

func (v NullableCustomTelephonyProviderCredentialUpdateRequest) Get() *CustomTelephonyProviderCredentialUpdateRequest {
	return v.value
}

func (v *NullableCustomTelephonyProviderCredentialUpdateRequest) Set(val *CustomTelephonyProviderCredentialUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomTelephonyProviderCredentialUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomTelephonyProviderCredentialUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomTelephonyProviderCredentialUpdateRequest(val *CustomTelephonyProviderCredentialUpdateRequest) *NullableCustomTelephonyProviderCredentialUpdateRequest {
	return &NullableCustomTelephonyProviderCredentialUpdateRequest{value: val, isSet: true}
}

func (v NullableCustomTelephonyProviderCredentialUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomTelephonyProviderCredentialUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
