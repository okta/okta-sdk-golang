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

// checks if the CustomTelephonyProviderSettingsTwilioMessagingService type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomTelephonyProviderSettingsTwilioMessagingService{}

// CustomTelephonyProviderSettingsTwilioMessagingService struct for CustomTelephonyProviderSettingsTwilioMessagingService
type CustomTelephonyProviderSettingsTwilioMessagingService struct {
	// The Twilio Messaging Service SID used for sending SMS messages. You can find this value in your Twilio console.  This method uses Twilio's [Programmable Messaging API](https://www.twilio.com/docs/messaging).
	TwilioMessageSid     *string `json:"twilioMessageSid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CustomTelephonyProviderSettingsTwilioMessagingService CustomTelephonyProviderSettingsTwilioMessagingService

// NewCustomTelephonyProviderSettingsTwilioMessagingService instantiates a new CustomTelephonyProviderSettingsTwilioMessagingService object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomTelephonyProviderSettingsTwilioMessagingService() *CustomTelephonyProviderSettingsTwilioMessagingService {
	this := CustomTelephonyProviderSettingsTwilioMessagingService{}
	return &this
}

// NewCustomTelephonyProviderSettingsTwilioMessagingServiceWithDefaults instantiates a new CustomTelephonyProviderSettingsTwilioMessagingService object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomTelephonyProviderSettingsTwilioMessagingServiceWithDefaults() *CustomTelephonyProviderSettingsTwilioMessagingService {
	this := CustomTelephonyProviderSettingsTwilioMessagingService{}
	return &this
}

// GetTwilioMessageSid returns the TwilioMessageSid field value if set, zero value otherwise.
func (o *CustomTelephonyProviderSettingsTwilioMessagingService) GetTwilioMessageSid() string {
	if o == nil || IsNil(o.TwilioMessageSid) {
		var ret string
		return ret
	}
	return *o.TwilioMessageSid
}

// GetTwilioMessageSidOk returns a tuple with the TwilioMessageSid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomTelephonyProviderSettingsTwilioMessagingService) GetTwilioMessageSidOk() (*string, bool) {
	if o == nil || IsNil(o.TwilioMessageSid) {
		return nil, false
	}
	return o.TwilioMessageSid, true
}

// HasTwilioMessageSid returns a boolean if a field has been set.
func (o *CustomTelephonyProviderSettingsTwilioMessagingService) HasTwilioMessageSid() bool {
	if o != nil && !IsNil(o.TwilioMessageSid) {
		return true
	}

	return false
}

// SetTwilioMessageSid gets a reference to the given string and assigns it to the TwilioMessageSid field.
func (o *CustomTelephonyProviderSettingsTwilioMessagingService) SetTwilioMessageSid(v string) {
	o.TwilioMessageSid = &v
}

func (o CustomTelephonyProviderSettingsTwilioMessagingService) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomTelephonyProviderSettingsTwilioMessagingService) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TwilioMessageSid) {
		toSerialize["twilioMessageSid"] = o.TwilioMessageSid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CustomTelephonyProviderSettingsTwilioMessagingService) UnmarshalJSON(data []byte) (err error) {
	varCustomTelephonyProviderSettingsTwilioMessagingService := _CustomTelephonyProviderSettingsTwilioMessagingService{}

	err = json.Unmarshal(data, &varCustomTelephonyProviderSettingsTwilioMessagingService)

	if err != nil {
		return err
	}

	*o = CustomTelephonyProviderSettingsTwilioMessagingService(varCustomTelephonyProviderSettingsTwilioMessagingService)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "twilioMessageSid")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCustomTelephonyProviderSettingsTwilioMessagingService struct {
	value *CustomTelephonyProviderSettingsTwilioMessagingService
	isSet bool
}

func (v NullableCustomTelephonyProviderSettingsTwilioMessagingService) Get() *CustomTelephonyProviderSettingsTwilioMessagingService {
	return v.value
}

func (v *NullableCustomTelephonyProviderSettingsTwilioMessagingService) Set(val *CustomTelephonyProviderSettingsTwilioMessagingService) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomTelephonyProviderSettingsTwilioMessagingService) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomTelephonyProviderSettingsTwilioMessagingService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomTelephonyProviderSettingsTwilioMessagingService(val *CustomTelephonyProviderSettingsTwilioMessagingService) *NullableCustomTelephonyProviderSettingsTwilioMessagingService {
	return &NullableCustomTelephonyProviderSettingsTwilioMessagingService{value: val, isSet: true}
}

func (v NullableCustomTelephonyProviderSettingsTwilioMessagingService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomTelephonyProviderSettingsTwilioMessagingService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
