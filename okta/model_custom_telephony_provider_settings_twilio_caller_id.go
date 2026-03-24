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

// checks if the CustomTelephonyProviderSettingsTwilioCallerId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomTelephonyProviderSettingsTwilioCallerId{}

// CustomTelephonyProviderSettingsTwilioCallerId struct for CustomTelephonyProviderSettingsTwilioCallerId
type CustomTelephonyProviderSettingsTwilioCallerId struct {
	// The Twilio caller ID that's used for making calls. You can find this value in your Twilio console.  This method uses Twilio's [Programmable Voice API](https://www.twilio.com/docs/voice).
	TwilioCallerId       *string `json:"twilioCallerId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CustomTelephonyProviderSettingsTwilioCallerId CustomTelephonyProviderSettingsTwilioCallerId

// NewCustomTelephonyProviderSettingsTwilioCallerId instantiates a new CustomTelephonyProviderSettingsTwilioCallerId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomTelephonyProviderSettingsTwilioCallerId() *CustomTelephonyProviderSettingsTwilioCallerId {
	this := CustomTelephonyProviderSettingsTwilioCallerId{}
	return &this
}

// NewCustomTelephonyProviderSettingsTwilioCallerIdWithDefaults instantiates a new CustomTelephonyProviderSettingsTwilioCallerId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomTelephonyProviderSettingsTwilioCallerIdWithDefaults() *CustomTelephonyProviderSettingsTwilioCallerId {
	this := CustomTelephonyProviderSettingsTwilioCallerId{}
	return &this
}

// GetTwilioCallerId returns the TwilioCallerId field value if set, zero value otherwise.
func (o *CustomTelephonyProviderSettingsTwilioCallerId) GetTwilioCallerId() string {
	if o == nil || IsNil(o.TwilioCallerId) {
		var ret string
		return ret
	}
	return *o.TwilioCallerId
}

// GetTwilioCallerIdOk returns a tuple with the TwilioCallerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomTelephonyProviderSettingsTwilioCallerId) GetTwilioCallerIdOk() (*string, bool) {
	if o == nil || IsNil(o.TwilioCallerId) {
		return nil, false
	}
	return o.TwilioCallerId, true
}

// HasTwilioCallerId returns a boolean if a field has been set.
func (o *CustomTelephonyProviderSettingsTwilioCallerId) HasTwilioCallerId() bool {
	if o != nil && !IsNil(o.TwilioCallerId) {
		return true
	}

	return false
}

// SetTwilioCallerId gets a reference to the given string and assigns it to the TwilioCallerId field.
func (o *CustomTelephonyProviderSettingsTwilioCallerId) SetTwilioCallerId(v string) {
	o.TwilioCallerId = &v
}

func (o CustomTelephonyProviderSettingsTwilioCallerId) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomTelephonyProviderSettingsTwilioCallerId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TwilioCallerId) {
		toSerialize["twilioCallerId"] = o.TwilioCallerId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CustomTelephonyProviderSettingsTwilioCallerId) UnmarshalJSON(data []byte) (err error) {
	varCustomTelephonyProviderSettingsTwilioCallerId := _CustomTelephonyProviderSettingsTwilioCallerId{}

	err = json.Unmarshal(data, &varCustomTelephonyProviderSettingsTwilioCallerId)

	if err != nil {
		return err
	}

	*o = CustomTelephonyProviderSettingsTwilioCallerId(varCustomTelephonyProviderSettingsTwilioCallerId)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "twilioCallerId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCustomTelephonyProviderSettingsTwilioCallerId struct {
	value *CustomTelephonyProviderSettingsTwilioCallerId
	isSet bool
}

func (v NullableCustomTelephonyProviderSettingsTwilioCallerId) Get() *CustomTelephonyProviderSettingsTwilioCallerId {
	return v.value
}

func (v *NullableCustomTelephonyProviderSettingsTwilioCallerId) Set(val *CustomTelephonyProviderSettingsTwilioCallerId) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomTelephonyProviderSettingsTwilioCallerId) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomTelephonyProviderSettingsTwilioCallerId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomTelephonyProviderSettingsTwilioCallerId(val *CustomTelephonyProviderSettingsTwilioCallerId) *NullableCustomTelephonyProviderSettingsTwilioCallerId {
	return &NullableCustomTelephonyProviderSettingsTwilioCallerId{value: val, isSet: true}
}

func (v NullableCustomTelephonyProviderSettingsTwilioCallerId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomTelephonyProviderSettingsTwilioCallerId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
