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

// TrendMicroApexOneServiceApplicationSettingsApplication Trend Micro Apex One as a Service app instance properties
type TrendMicroApexOneServiceApplicationSettingsApplication struct {
	// Base Trend Micro Apex One Service URL
	BaseURL string `json:"baseURL"`
	AdditionalProperties map[string]interface{}
}

type _TrendMicroApexOneServiceApplicationSettingsApplication TrendMicroApexOneServiceApplicationSettingsApplication

// NewTrendMicroApexOneServiceApplicationSettingsApplication instantiates a new TrendMicroApexOneServiceApplicationSettingsApplication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrendMicroApexOneServiceApplicationSettingsApplication(baseURL string) *TrendMicroApexOneServiceApplicationSettingsApplication {
	this := TrendMicroApexOneServiceApplicationSettingsApplication{}
	this.BaseURL = baseURL
	return &this
}

// NewTrendMicroApexOneServiceApplicationSettingsApplicationWithDefaults instantiates a new TrendMicroApexOneServiceApplicationSettingsApplication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrendMicroApexOneServiceApplicationSettingsApplicationWithDefaults() *TrendMicroApexOneServiceApplicationSettingsApplication {
	this := TrendMicroApexOneServiceApplicationSettingsApplication{}
	return &this
}

// GetBaseURL returns the BaseURL field value
func (o *TrendMicroApexOneServiceApplicationSettingsApplication) GetBaseURL() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BaseURL
}

// GetBaseURLOk returns a tuple with the BaseURL field value
// and a boolean to check if the value has been set.
func (o *TrendMicroApexOneServiceApplicationSettingsApplication) GetBaseURLOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseURL, true
}

// SetBaseURL sets field value
func (o *TrendMicroApexOneServiceApplicationSettingsApplication) SetBaseURL(v string) {
	o.BaseURL = v
}

func (o TrendMicroApexOneServiceApplicationSettingsApplication) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["baseURL"] = o.BaseURL
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TrendMicroApexOneServiceApplicationSettingsApplication) UnmarshalJSON(bytes []byte) (err error) {
	varTrendMicroApexOneServiceApplicationSettingsApplication := _TrendMicroApexOneServiceApplicationSettingsApplication{}

	err = json.Unmarshal(bytes, &varTrendMicroApexOneServiceApplicationSettingsApplication)
	if err == nil {
		*o = TrendMicroApexOneServiceApplicationSettingsApplication(varTrendMicroApexOneServiceApplicationSettingsApplication)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "baseURL")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableTrendMicroApexOneServiceApplicationSettingsApplication struct {
	value *TrendMicroApexOneServiceApplicationSettingsApplication
	isSet bool
}

func (v NullableTrendMicroApexOneServiceApplicationSettingsApplication) Get() *TrendMicroApexOneServiceApplicationSettingsApplication {
	return v.value
}

func (v *NullableTrendMicroApexOneServiceApplicationSettingsApplication) Set(val *TrendMicroApexOneServiceApplicationSettingsApplication) {
	v.value = val
	v.isSet = true
}

func (v NullableTrendMicroApexOneServiceApplicationSettingsApplication) IsSet() bool {
	return v.isSet
}

func (v *NullableTrendMicroApexOneServiceApplicationSettingsApplication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrendMicroApexOneServiceApplicationSettingsApplication(val *TrendMicroApexOneServiceApplicationSettingsApplication) *NullableTrendMicroApexOneServiceApplicationSettingsApplication {
	return &NullableTrendMicroApexOneServiceApplicationSettingsApplication{value: val, isSet: true}
}

func (v NullableTrendMicroApexOneServiceApplicationSettingsApplication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrendMicroApexOneServiceApplicationSettingsApplication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

