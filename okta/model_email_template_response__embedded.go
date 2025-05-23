/*
Okta Admin Management

Allows customers to easily access the Okta Management APIs

Copyright 2018 - Present Okta, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

API version: 2024.06.1
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// EmailTemplateResponseEmbedded struct for EmailTemplateResponseEmbedded
type EmailTemplateResponseEmbedded struct {
	Settings *EmailSettingsResponse `json:"settings,omitempty"`
	CustomizationCount *int32 `json:"customizationCount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EmailTemplateResponseEmbedded EmailTemplateResponseEmbedded

// NewEmailTemplateResponseEmbedded instantiates a new EmailTemplateResponseEmbedded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailTemplateResponseEmbedded() *EmailTemplateResponseEmbedded {
	this := EmailTemplateResponseEmbedded{}
	return &this
}

// NewEmailTemplateResponseEmbeddedWithDefaults instantiates a new EmailTemplateResponseEmbedded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailTemplateResponseEmbeddedWithDefaults() *EmailTemplateResponseEmbedded {
	this := EmailTemplateResponseEmbedded{}
	return &this
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *EmailTemplateResponseEmbedded) GetSettings() EmailSettingsResponse {
	if o == nil || o.Settings == nil {
		var ret EmailSettingsResponse
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailTemplateResponseEmbedded) GetSettingsOk() (*EmailSettingsResponse, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *EmailTemplateResponseEmbedded) HasSettings() bool {
	if o != nil && o.Settings != nil {
		return true
	}

	return false
}

// SetSettings gets a reference to the given EmailSettingsResponse and assigns it to the Settings field.
func (o *EmailTemplateResponseEmbedded) SetSettings(v EmailSettingsResponse) {
	o.Settings = &v
}

// GetCustomizationCount returns the CustomizationCount field value if set, zero value otherwise.
func (o *EmailTemplateResponseEmbedded) GetCustomizationCount() int32 {
	if o == nil || o.CustomizationCount == nil {
		var ret int32
		return ret
	}
	return *o.CustomizationCount
}

// GetCustomizationCountOk returns a tuple with the CustomizationCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailTemplateResponseEmbedded) GetCustomizationCountOk() (*int32, bool) {
	if o == nil || o.CustomizationCount == nil {
		return nil, false
	}
	return o.CustomizationCount, true
}

// HasCustomizationCount returns a boolean if a field has been set.
func (o *EmailTemplateResponseEmbedded) HasCustomizationCount() bool {
	if o != nil && o.CustomizationCount != nil {
		return true
	}

	return false
}

// SetCustomizationCount gets a reference to the given int32 and assigns it to the CustomizationCount field.
func (o *EmailTemplateResponseEmbedded) SetCustomizationCount(v int32) {
	o.CustomizationCount = &v
}

func (o EmailTemplateResponseEmbedded) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Settings != nil {
		toSerialize["settings"] = o.Settings
	}
	if o.CustomizationCount != nil {
		toSerialize["customizationCount"] = o.CustomizationCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EmailTemplateResponseEmbedded) UnmarshalJSON(bytes []byte) (err error) {
	varEmailTemplateResponseEmbedded := _EmailTemplateResponseEmbedded{}

	err = json.Unmarshal(bytes, &varEmailTemplateResponseEmbedded)
	if err == nil {
		*o = EmailTemplateResponseEmbedded(varEmailTemplateResponseEmbedded)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "settings")
		delete(additionalProperties, "customizationCount")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableEmailTemplateResponseEmbedded struct {
	value *EmailTemplateResponseEmbedded
	isSet bool
}

func (v NullableEmailTemplateResponseEmbedded) Get() *EmailTemplateResponseEmbedded {
	return v.value
}

func (v *NullableEmailTemplateResponseEmbedded) Set(val *EmailTemplateResponseEmbedded) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailTemplateResponseEmbedded) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailTemplateResponseEmbedded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailTemplateResponseEmbedded(val *EmailTemplateResponseEmbedded) *NullableEmailTemplateResponseEmbedded {
	return &NullableEmailTemplateResponseEmbedded{value: val, isSet: true}
}

func (v NullableEmailTemplateResponseEmbedded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailTemplateResponseEmbedded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

