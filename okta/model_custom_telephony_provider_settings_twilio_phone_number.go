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

// checks if the CustomTelephonyProviderSettingsTwilioPhoneNumber type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomTelephonyProviderSettingsTwilioPhoneNumber{}

// CustomTelephonyProviderSettingsTwilioPhoneNumber struct for CustomTelephonyProviderSettingsTwilioPhoneNumber
type CustomTelephonyProviderSettingsTwilioPhoneNumber struct {
	// The Twilio phone number that's used for sending SMS messages or voice calls. You can find this value in your Twilio console.  This method uses Twilio's [Programmable Messaging API](https://www.twilio.com/docs/messaging).
	TwilioPhoneNumber    *string `json:"twilioPhoneNumber,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CustomTelephonyProviderSettingsTwilioPhoneNumber CustomTelephonyProviderSettingsTwilioPhoneNumber

// NewCustomTelephonyProviderSettingsTwilioPhoneNumber instantiates a new CustomTelephonyProviderSettingsTwilioPhoneNumber object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomTelephonyProviderSettingsTwilioPhoneNumber() *CustomTelephonyProviderSettingsTwilioPhoneNumber {
	this := CustomTelephonyProviderSettingsTwilioPhoneNumber{}
	return &this
}

// NewCustomTelephonyProviderSettingsTwilioPhoneNumberWithDefaults instantiates a new CustomTelephonyProviderSettingsTwilioPhoneNumber object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomTelephonyProviderSettingsTwilioPhoneNumberWithDefaults() *CustomTelephonyProviderSettingsTwilioPhoneNumber {
	this := CustomTelephonyProviderSettingsTwilioPhoneNumber{}
	return &this
}

// GetTwilioPhoneNumber returns the TwilioPhoneNumber field value if set, zero value otherwise.
func (o *CustomTelephonyProviderSettingsTwilioPhoneNumber) GetTwilioPhoneNumber() string {
	if o == nil || IsNil(o.TwilioPhoneNumber) {
		var ret string
		return ret
	}
	return *o.TwilioPhoneNumber
}

// GetTwilioPhoneNumberOk returns a tuple with the TwilioPhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomTelephonyProviderSettingsTwilioPhoneNumber) GetTwilioPhoneNumberOk() (*string, bool) {
	if o == nil || IsNil(o.TwilioPhoneNumber) {
		return nil, false
	}
	return o.TwilioPhoneNumber, true
}

// HasTwilioPhoneNumber returns a boolean if a field has been set.
func (o *CustomTelephonyProviderSettingsTwilioPhoneNumber) HasTwilioPhoneNumber() bool {
	if o != nil && !IsNil(o.TwilioPhoneNumber) {
		return true
	}

	return false
}

// SetTwilioPhoneNumber gets a reference to the given string and assigns it to the TwilioPhoneNumber field.
func (o *CustomTelephonyProviderSettingsTwilioPhoneNumber) SetTwilioPhoneNumber(v string) {
	o.TwilioPhoneNumber = &v
}

func (o CustomTelephonyProviderSettingsTwilioPhoneNumber) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomTelephonyProviderSettingsTwilioPhoneNumber) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TwilioPhoneNumber) {
		toSerialize["twilioPhoneNumber"] = o.TwilioPhoneNumber
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CustomTelephonyProviderSettingsTwilioPhoneNumber) UnmarshalJSON(data []byte) (err error) {
	varCustomTelephonyProviderSettingsTwilioPhoneNumber := _CustomTelephonyProviderSettingsTwilioPhoneNumber{}

	err = json.Unmarshal(data, &varCustomTelephonyProviderSettingsTwilioPhoneNumber)

	if err != nil {
		return err
	}

	*o = CustomTelephonyProviderSettingsTwilioPhoneNumber(varCustomTelephonyProviderSettingsTwilioPhoneNumber)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "twilioPhoneNumber")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCustomTelephonyProviderSettingsTwilioPhoneNumber struct {
	value *CustomTelephonyProviderSettingsTwilioPhoneNumber
	isSet bool
}

func (v NullableCustomTelephonyProviderSettingsTwilioPhoneNumber) Get() *CustomTelephonyProviderSettingsTwilioPhoneNumber {
	return v.value
}

func (v *NullableCustomTelephonyProviderSettingsTwilioPhoneNumber) Set(val *CustomTelephonyProviderSettingsTwilioPhoneNumber) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomTelephonyProviderSettingsTwilioPhoneNumber) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomTelephonyProviderSettingsTwilioPhoneNumber) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomTelephonyProviderSettingsTwilioPhoneNumber(val *CustomTelephonyProviderSettingsTwilioPhoneNumber) *NullableCustomTelephonyProviderSettingsTwilioPhoneNumber {
	return &NullableCustomTelephonyProviderSettingsTwilioPhoneNumber{value: val, isSet: true}
}

func (v NullableCustomTelephonyProviderSettingsTwilioPhoneNumber) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomTelephonyProviderSettingsTwilioPhoneNumber) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
