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

// checks if the CustomTelephonyProviderSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomTelephonyProviderSettings{}

// CustomTelephonyProviderSettings Settings for custom telephony provider.  These settings vary based on the telephony provider and the type of telephony operation (SMS or Voice). For `sms` and `call`, you can select one method per telephony operation (`sms` and `call`) for sending messages or voice calls.  > **Note:** Configure your telephony provider settings before selecting the methods for sending SMS messages or making voice calls. For example, if you select Twilio as your telephony provider, and you want to send SMS messages using Twilio's Verify Service, ensure that you have the Verify Service set up in your Twilio account. You can then use the `twilioVerifySid` field under `sms` to provide the necessary SID.
type CustomTelephonyProviderSettings struct {
	Call                 *CustomTelephonyProviderSettingsCall `json:"call,omitempty"`
	Sms                  *CustomTelephonyProviderSettingsSms  `json:"sms,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CustomTelephonyProviderSettings CustomTelephonyProviderSettings

// NewCustomTelephonyProviderSettings instantiates a new CustomTelephonyProviderSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomTelephonyProviderSettings() *CustomTelephonyProviderSettings {
	this := CustomTelephonyProviderSettings{}
	return &this
}

// NewCustomTelephonyProviderSettingsWithDefaults instantiates a new CustomTelephonyProviderSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomTelephonyProviderSettingsWithDefaults() *CustomTelephonyProviderSettings {
	this := CustomTelephonyProviderSettings{}
	return &this
}

// GetCall returns the Call field value if set, zero value otherwise.
func (o *CustomTelephonyProviderSettings) GetCall() CustomTelephonyProviderSettingsCall {
	if o == nil || IsNil(o.Call) {
		var ret CustomTelephonyProviderSettingsCall
		return ret
	}
	return *o.Call
}

// GetCallOk returns a tuple with the Call field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomTelephonyProviderSettings) GetCallOk() (*CustomTelephonyProviderSettingsCall, bool) {
	if o == nil || IsNil(o.Call) {
		return nil, false
	}
	return o.Call, true
}

// HasCall returns a boolean if a field has been set.
func (o *CustomTelephonyProviderSettings) HasCall() bool {
	if o != nil && !IsNil(o.Call) {
		return true
	}

	return false
}

// SetCall gets a reference to the given CustomTelephonyProviderSettingsCall and assigns it to the Call field.
func (o *CustomTelephonyProviderSettings) SetCall(v CustomTelephonyProviderSettingsCall) {
	o.Call = &v
}

// GetSms returns the Sms field value if set, zero value otherwise.
func (o *CustomTelephonyProviderSettings) GetSms() CustomTelephonyProviderSettingsSms {
	if o == nil || IsNil(o.Sms) {
		var ret CustomTelephonyProviderSettingsSms
		return ret
	}
	return *o.Sms
}

// GetSmsOk returns a tuple with the Sms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomTelephonyProviderSettings) GetSmsOk() (*CustomTelephonyProviderSettingsSms, bool) {
	if o == nil || IsNil(o.Sms) {
		return nil, false
	}
	return o.Sms, true
}

// HasSms returns a boolean if a field has been set.
func (o *CustomTelephonyProviderSettings) HasSms() bool {
	if o != nil && !IsNil(o.Sms) {
		return true
	}

	return false
}

// SetSms gets a reference to the given CustomTelephonyProviderSettingsSms and assigns it to the Sms field.
func (o *CustomTelephonyProviderSettings) SetSms(v CustomTelephonyProviderSettingsSms) {
	o.Sms = &v
}

func (o CustomTelephonyProviderSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomTelephonyProviderSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Call) {
		toSerialize["call"] = o.Call
	}
	if !IsNil(o.Sms) {
		toSerialize["sms"] = o.Sms
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CustomTelephonyProviderSettings) UnmarshalJSON(data []byte) (err error) {
	varCustomTelephonyProviderSettings := _CustomTelephonyProviderSettings{}

	err = json.Unmarshal(data, &varCustomTelephonyProviderSettings)

	if err != nil {
		return err
	}

	*o = CustomTelephonyProviderSettings(varCustomTelephonyProviderSettings)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "call")
		delete(additionalProperties, "sms")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCustomTelephonyProviderSettings struct {
	value *CustomTelephonyProviderSettings
	isSet bool
}

func (v NullableCustomTelephonyProviderSettings) Get() *CustomTelephonyProviderSettings {
	return v.value
}

func (v *NullableCustomTelephonyProviderSettings) Set(val *CustomTelephonyProviderSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomTelephonyProviderSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomTelephonyProviderSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomTelephonyProviderSettings(val *CustomTelephonyProviderSettings) *NullableCustomTelephonyProviderSettings {
	return &NullableCustomTelephonyProviderSettings{value: val, isSet: true}
}

func (v NullableCustomTelephonyProviderSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomTelephonyProviderSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
