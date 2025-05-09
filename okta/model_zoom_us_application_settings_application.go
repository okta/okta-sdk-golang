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

// ZoomUsApplicationSettingsApplication Zoom app instance properties
type ZoomUsApplicationSettingsApplication struct {
	// Your Zoom subdomain
	SubDomain string `json:"subDomain"`
	AdditionalProperties map[string]interface{}
}

type _ZoomUsApplicationSettingsApplication ZoomUsApplicationSettingsApplication

// NewZoomUsApplicationSettingsApplication instantiates a new ZoomUsApplicationSettingsApplication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZoomUsApplicationSettingsApplication(subDomain string) *ZoomUsApplicationSettingsApplication {
	this := ZoomUsApplicationSettingsApplication{}
	this.SubDomain = subDomain
	return &this
}

// NewZoomUsApplicationSettingsApplicationWithDefaults instantiates a new ZoomUsApplicationSettingsApplication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZoomUsApplicationSettingsApplicationWithDefaults() *ZoomUsApplicationSettingsApplication {
	this := ZoomUsApplicationSettingsApplication{}
	return &this
}

// GetSubDomain returns the SubDomain field value
func (o *ZoomUsApplicationSettingsApplication) GetSubDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubDomain
}

// GetSubDomainOk returns a tuple with the SubDomain field value
// and a boolean to check if the value has been set.
func (o *ZoomUsApplicationSettingsApplication) GetSubDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubDomain, true
}

// SetSubDomain sets field value
func (o *ZoomUsApplicationSettingsApplication) SetSubDomain(v string) {
	o.SubDomain = v
}

func (o ZoomUsApplicationSettingsApplication) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["subDomain"] = o.SubDomain
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ZoomUsApplicationSettingsApplication) UnmarshalJSON(bytes []byte) (err error) {
	varZoomUsApplicationSettingsApplication := _ZoomUsApplicationSettingsApplication{}

	err = json.Unmarshal(bytes, &varZoomUsApplicationSettingsApplication)
	if err == nil {
		*o = ZoomUsApplicationSettingsApplication(varZoomUsApplicationSettingsApplication)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "subDomain")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableZoomUsApplicationSettingsApplication struct {
	value *ZoomUsApplicationSettingsApplication
	isSet bool
}

func (v NullableZoomUsApplicationSettingsApplication) Get() *ZoomUsApplicationSettingsApplication {
	return v.value
}

func (v *NullableZoomUsApplicationSettingsApplication) Set(val *ZoomUsApplicationSettingsApplication) {
	v.value = val
	v.isSet = true
}

func (v NullableZoomUsApplicationSettingsApplication) IsSet() bool {
	return v.isSet
}

func (v *NullableZoomUsApplicationSettingsApplication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZoomUsApplicationSettingsApplication(val *ZoomUsApplicationSettingsApplication) *NullableZoomUsApplicationSettingsApplication {
	return &NullableZoomUsApplicationSettingsApplication{value: val, isSet: true}
}

func (v NullableZoomUsApplicationSettingsApplication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZoomUsApplicationSettingsApplication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

