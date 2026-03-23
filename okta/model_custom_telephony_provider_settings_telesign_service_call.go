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

// checks if the CustomTelephonyProviderSettingsTelesignServiceCall type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomTelephonyProviderSettingsTelesignServiceCall{}

// CustomTelephonyProviderSettingsTelesignServiceCall struct for CustomTelephonyProviderSettingsTelesignServiceCall
type CustomTelephonyProviderSettingsTelesignServiceCall struct {
	// The Telesign service identifier used for sending making calls. You can find this value in your Telesign console.  The `telesignVerifyService` method uses Telesign's [Verify API](https://developer.telesign.com/enterprise/docs/verify-api-overview). And the `telesignVoiceService` method uses Telesign's [Voice API](https://developer.telesign.com/enterprise/docs/voice-overview).
	TelesignService      *string `json:"telesignService,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CustomTelephonyProviderSettingsTelesignServiceCall CustomTelephonyProviderSettingsTelesignServiceCall

// NewCustomTelephonyProviderSettingsTelesignServiceCall instantiates a new CustomTelephonyProviderSettingsTelesignServiceCall object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomTelephonyProviderSettingsTelesignServiceCall() *CustomTelephonyProviderSettingsTelesignServiceCall {
	this := CustomTelephonyProviderSettingsTelesignServiceCall{}
	return &this
}

// NewCustomTelephonyProviderSettingsTelesignServiceCallWithDefaults instantiates a new CustomTelephonyProviderSettingsTelesignServiceCall object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomTelephonyProviderSettingsTelesignServiceCallWithDefaults() *CustomTelephonyProviderSettingsTelesignServiceCall {
	this := CustomTelephonyProviderSettingsTelesignServiceCall{}
	return &this
}

// GetTelesignService returns the TelesignService field value if set, zero value otherwise.
func (o *CustomTelephonyProviderSettingsTelesignServiceCall) GetTelesignService() string {
	if o == nil || IsNil(o.TelesignService) {
		var ret string
		return ret
	}
	return *o.TelesignService
}

// GetTelesignServiceOk returns a tuple with the TelesignService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomTelephonyProviderSettingsTelesignServiceCall) GetTelesignServiceOk() (*string, bool) {
	if o == nil || IsNil(o.TelesignService) {
		return nil, false
	}
	return o.TelesignService, true
}

// HasTelesignService returns a boolean if a field has been set.
func (o *CustomTelephonyProviderSettingsTelesignServiceCall) HasTelesignService() bool {
	if o != nil && !IsNil(o.TelesignService) {
		return true
	}

	return false
}

// SetTelesignService gets a reference to the given string and assigns it to the TelesignService field.
func (o *CustomTelephonyProviderSettingsTelesignServiceCall) SetTelesignService(v string) {
	o.TelesignService = &v
}

func (o CustomTelephonyProviderSettingsTelesignServiceCall) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomTelephonyProviderSettingsTelesignServiceCall) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TelesignService) {
		toSerialize["telesignService"] = o.TelesignService
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CustomTelephonyProviderSettingsTelesignServiceCall) UnmarshalJSON(data []byte) (err error) {
	varCustomTelephonyProviderSettingsTelesignServiceCall := _CustomTelephonyProviderSettingsTelesignServiceCall{}

	err = json.Unmarshal(data, &varCustomTelephonyProviderSettingsTelesignServiceCall)

	if err != nil {
		return err
	}

	*o = CustomTelephonyProviderSettingsTelesignServiceCall(varCustomTelephonyProviderSettingsTelesignServiceCall)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "telesignService")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCustomTelephonyProviderSettingsTelesignServiceCall struct {
	value *CustomTelephonyProviderSettingsTelesignServiceCall
	isSet bool
}

func (v NullableCustomTelephonyProviderSettingsTelesignServiceCall) Get() *CustomTelephonyProviderSettingsTelesignServiceCall {
	return v.value
}

func (v *NullableCustomTelephonyProviderSettingsTelesignServiceCall) Set(val *CustomTelephonyProviderSettingsTelesignServiceCall) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomTelephonyProviderSettingsTelesignServiceCall) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomTelephonyProviderSettingsTelesignServiceCall) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomTelephonyProviderSettingsTelesignServiceCall(val *CustomTelephonyProviderSettingsTelesignServiceCall) *NullableCustomTelephonyProviderSettingsTelesignServiceCall {
	return &NullableCustomTelephonyProviderSettingsTelesignServiceCall{value: val, isSet: true}
}

func (v NullableCustomTelephonyProviderSettingsTelesignServiceCall) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomTelephonyProviderSettingsTelesignServiceCall) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
