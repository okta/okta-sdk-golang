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

// checks if the CustomTelephonyProviderSettingsTelesignServiceSms type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomTelephonyProviderSettingsTelesignServiceSms{}

// CustomTelephonyProviderSettingsTelesignServiceSms struct for CustomTelephonyProviderSettingsTelesignServiceSms
type CustomTelephonyProviderSettingsTelesignServiceSms struct {
	// The Telesign service identifier used for sending SMS messages. You can find this value in your Telesign console.  The `telesignVerifyService` method uses Telesign's [Verify API](https://developer.telesign.com/enterprise/docs/verify-api-overview). And the `telesignMessagingService` method uses Telesign's [SMS](https://developer.telesign.com/enterprise/docs/voice-overview) and [Messaging](https://developer.telesign.com/enterprise/docs/messaging-overview) APIs.
	TelesignService      *string `json:"telesignService,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CustomTelephonyProviderSettingsTelesignServiceSms CustomTelephonyProviderSettingsTelesignServiceSms

// NewCustomTelephonyProviderSettingsTelesignServiceSms instantiates a new CustomTelephonyProviderSettingsTelesignServiceSms object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomTelephonyProviderSettingsTelesignServiceSms() *CustomTelephonyProviderSettingsTelesignServiceSms {
	this := CustomTelephonyProviderSettingsTelesignServiceSms{}
	return &this
}

// NewCustomTelephonyProviderSettingsTelesignServiceSmsWithDefaults instantiates a new CustomTelephonyProviderSettingsTelesignServiceSms object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomTelephonyProviderSettingsTelesignServiceSmsWithDefaults() *CustomTelephonyProviderSettingsTelesignServiceSms {
	this := CustomTelephonyProviderSettingsTelesignServiceSms{}
	return &this
}

// GetTelesignService returns the TelesignService field value if set, zero value otherwise.
func (o *CustomTelephonyProviderSettingsTelesignServiceSms) GetTelesignService() string {
	if o == nil || IsNil(o.TelesignService) {
		var ret string
		return ret
	}
	return *o.TelesignService
}

// GetTelesignServiceOk returns a tuple with the TelesignService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomTelephonyProviderSettingsTelesignServiceSms) GetTelesignServiceOk() (*string, bool) {
	if o == nil || IsNil(o.TelesignService) {
		return nil, false
	}
	return o.TelesignService, true
}

// HasTelesignService returns a boolean if a field has been set.
func (o *CustomTelephonyProviderSettingsTelesignServiceSms) HasTelesignService() bool {
	if o != nil && !IsNil(o.TelesignService) {
		return true
	}

	return false
}

// SetTelesignService gets a reference to the given string and assigns it to the TelesignService field.
func (o *CustomTelephonyProviderSettingsTelesignServiceSms) SetTelesignService(v string) {
	o.TelesignService = &v
}

func (o CustomTelephonyProviderSettingsTelesignServiceSms) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomTelephonyProviderSettingsTelesignServiceSms) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TelesignService) {
		toSerialize["telesignService"] = o.TelesignService
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CustomTelephonyProviderSettingsTelesignServiceSms) UnmarshalJSON(data []byte) (err error) {
	varCustomTelephonyProviderSettingsTelesignServiceSms := _CustomTelephonyProviderSettingsTelesignServiceSms{}

	err = json.Unmarshal(data, &varCustomTelephonyProviderSettingsTelesignServiceSms)

	if err != nil {
		return err
	}

	*o = CustomTelephonyProviderSettingsTelesignServiceSms(varCustomTelephonyProviderSettingsTelesignServiceSms)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "telesignService")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCustomTelephonyProviderSettingsTelesignServiceSms struct {
	value *CustomTelephonyProviderSettingsTelesignServiceSms
	isSet bool
}

func (v NullableCustomTelephonyProviderSettingsTelesignServiceSms) Get() *CustomTelephonyProviderSettingsTelesignServiceSms {
	return v.value
}

func (v *NullableCustomTelephonyProviderSettingsTelesignServiceSms) Set(val *CustomTelephonyProviderSettingsTelesignServiceSms) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomTelephonyProviderSettingsTelesignServiceSms) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomTelephonyProviderSettingsTelesignServiceSms) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomTelephonyProviderSettingsTelesignServiceSms(val *CustomTelephonyProviderSettingsTelesignServiceSms) *NullableCustomTelephonyProviderSettingsTelesignServiceSms {
	return &NullableCustomTelephonyProviderSettingsTelesignServiceSms{value: val, isSet: true}
}

func (v NullableCustomTelephonyProviderSettingsTelesignServiceSms) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomTelephonyProviderSettingsTelesignServiceSms) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
