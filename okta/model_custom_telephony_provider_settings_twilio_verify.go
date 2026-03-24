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

// checks if the CustomTelephonyProviderSettingsTwilioVerify type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomTelephonyProviderSettingsTwilioVerify{}

// CustomTelephonyProviderSettingsTwilioVerify struct for CustomTelephonyProviderSettingsTwilioVerify
type CustomTelephonyProviderSettingsTwilioVerify struct {
	// The Twilio Verify Service SID used for sending verification messages or calls. You can find this value in your Twilio console.  This method uses Twilio's [Verify API](https://www.twilio.com/docs/verify/api).
	TwilioVerifySid      *string `json:"twilioVerifySid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CustomTelephonyProviderSettingsTwilioVerify CustomTelephonyProviderSettingsTwilioVerify

// NewCustomTelephonyProviderSettingsTwilioVerify instantiates a new CustomTelephonyProviderSettingsTwilioVerify object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomTelephonyProviderSettingsTwilioVerify() *CustomTelephonyProviderSettingsTwilioVerify {
	this := CustomTelephonyProviderSettingsTwilioVerify{}
	return &this
}

// NewCustomTelephonyProviderSettingsTwilioVerifyWithDefaults instantiates a new CustomTelephonyProviderSettingsTwilioVerify object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomTelephonyProviderSettingsTwilioVerifyWithDefaults() *CustomTelephonyProviderSettingsTwilioVerify {
	this := CustomTelephonyProviderSettingsTwilioVerify{}
	return &this
}

// GetTwilioVerifySid returns the TwilioVerifySid field value if set, zero value otherwise.
func (o *CustomTelephonyProviderSettingsTwilioVerify) GetTwilioVerifySid() string {
	if o == nil || IsNil(o.TwilioVerifySid) {
		var ret string
		return ret
	}
	return *o.TwilioVerifySid
}

// GetTwilioVerifySidOk returns a tuple with the TwilioVerifySid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomTelephonyProviderSettingsTwilioVerify) GetTwilioVerifySidOk() (*string, bool) {
	if o == nil || IsNil(o.TwilioVerifySid) {
		return nil, false
	}
	return o.TwilioVerifySid, true
}

// HasTwilioVerifySid returns a boolean if a field has been set.
func (o *CustomTelephonyProviderSettingsTwilioVerify) HasTwilioVerifySid() bool {
	if o != nil && !IsNil(o.TwilioVerifySid) {
		return true
	}

	return false
}

// SetTwilioVerifySid gets a reference to the given string and assigns it to the TwilioVerifySid field.
func (o *CustomTelephonyProviderSettingsTwilioVerify) SetTwilioVerifySid(v string) {
	o.TwilioVerifySid = &v
}

func (o CustomTelephonyProviderSettingsTwilioVerify) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomTelephonyProviderSettingsTwilioVerify) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TwilioVerifySid) {
		toSerialize["twilioVerifySid"] = o.TwilioVerifySid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CustomTelephonyProviderSettingsTwilioVerify) UnmarshalJSON(data []byte) (err error) {
	varCustomTelephonyProviderSettingsTwilioVerify := _CustomTelephonyProviderSettingsTwilioVerify{}

	err = json.Unmarshal(data, &varCustomTelephonyProviderSettingsTwilioVerify)

	if err != nil {
		return err
	}

	*o = CustomTelephonyProviderSettingsTwilioVerify(varCustomTelephonyProviderSettingsTwilioVerify)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "twilioVerifySid")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCustomTelephonyProviderSettingsTwilioVerify struct {
	value *CustomTelephonyProviderSettingsTwilioVerify
	isSet bool
}

func (v NullableCustomTelephonyProviderSettingsTwilioVerify) Get() *CustomTelephonyProviderSettingsTwilioVerify {
	return v.value
}

func (v *NullableCustomTelephonyProviderSettingsTwilioVerify) Set(val *CustomTelephonyProviderSettingsTwilioVerify) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomTelephonyProviderSettingsTwilioVerify) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomTelephonyProviderSettingsTwilioVerify) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomTelephonyProviderSettingsTwilioVerify(val *CustomTelephonyProviderSettingsTwilioVerify) *NullableCustomTelephonyProviderSettingsTwilioVerify {
	return &NullableCustomTelephonyProviderSettingsTwilioVerify{value: val, isSet: true}
}

func (v NullableCustomTelephonyProviderSettingsTwilioVerify) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomTelephonyProviderSettingsTwilioVerify) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
