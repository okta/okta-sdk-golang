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

// ZscalerbyzApplicationSettingsApplication Zscaler app instance properties
type ZscalerbyzApplicationSettingsApplication struct {
	// Your Zscaler domain
	SiteDomain *string `json:"siteDomain,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ZscalerbyzApplicationSettingsApplication ZscalerbyzApplicationSettingsApplication

// NewZscalerbyzApplicationSettingsApplication instantiates a new ZscalerbyzApplicationSettingsApplication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZscalerbyzApplicationSettingsApplication() *ZscalerbyzApplicationSettingsApplication {
	this := ZscalerbyzApplicationSettingsApplication{}
	return &this
}

// NewZscalerbyzApplicationSettingsApplicationWithDefaults instantiates a new ZscalerbyzApplicationSettingsApplication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZscalerbyzApplicationSettingsApplicationWithDefaults() *ZscalerbyzApplicationSettingsApplication {
	this := ZscalerbyzApplicationSettingsApplication{}
	return &this
}

// GetSiteDomain returns the SiteDomain field value if set, zero value otherwise.
func (o *ZscalerbyzApplicationSettingsApplication) GetSiteDomain() string {
	if o == nil || o.SiteDomain == nil {
		var ret string
		return ret
	}
	return *o.SiteDomain
}

// GetSiteDomainOk returns a tuple with the SiteDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZscalerbyzApplicationSettingsApplication) GetSiteDomainOk() (*string, bool) {
	if o == nil || o.SiteDomain == nil {
		return nil, false
	}
	return o.SiteDomain, true
}

// HasSiteDomain returns a boolean if a field has been set.
func (o *ZscalerbyzApplicationSettingsApplication) HasSiteDomain() bool {
	if o != nil && o.SiteDomain != nil {
		return true
	}

	return false
}

// SetSiteDomain gets a reference to the given string and assigns it to the SiteDomain field.
func (o *ZscalerbyzApplicationSettingsApplication) SetSiteDomain(v string) {
	o.SiteDomain = &v
}

func (o ZscalerbyzApplicationSettingsApplication) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SiteDomain != nil {
		toSerialize["siteDomain"] = o.SiteDomain
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ZscalerbyzApplicationSettingsApplication) UnmarshalJSON(bytes []byte) (err error) {
	varZscalerbyzApplicationSettingsApplication := _ZscalerbyzApplicationSettingsApplication{}

	err = json.Unmarshal(bytes, &varZscalerbyzApplicationSettingsApplication)
	if err == nil {
		*o = ZscalerbyzApplicationSettingsApplication(varZscalerbyzApplicationSettingsApplication)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "siteDomain")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableZscalerbyzApplicationSettingsApplication struct {
	value *ZscalerbyzApplicationSettingsApplication
	isSet bool
}

func (v NullableZscalerbyzApplicationSettingsApplication) Get() *ZscalerbyzApplicationSettingsApplication {
	return v.value
}

func (v *NullableZscalerbyzApplicationSettingsApplication) Set(val *ZscalerbyzApplicationSettingsApplication) {
	v.value = val
	v.isSet = true
}

func (v NullableZscalerbyzApplicationSettingsApplication) IsSet() bool {
	return v.isSet
}

func (v *NullableZscalerbyzApplicationSettingsApplication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZscalerbyzApplicationSettingsApplication(val *ZscalerbyzApplicationSettingsApplication) *NullableZscalerbyzApplicationSettingsApplication {
	return &NullableZscalerbyzApplicationSettingsApplication{value: val, isSet: true}
}

func (v NullableZscalerbyzApplicationSettingsApplication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZscalerbyzApplicationSettingsApplication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

