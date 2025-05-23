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

// SamlApplicationSettingsApplication struct for SamlApplicationSettingsApplication
type SamlApplicationSettingsApplication struct {
	AcsUrl *string `json:"acsUrl,omitempty"`
	AudRestriction *string `json:"audRestriction,omitempty"`
	BaseUrl *string `json:"baseUrl,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SamlApplicationSettingsApplication SamlApplicationSettingsApplication

// NewSamlApplicationSettingsApplication instantiates a new SamlApplicationSettingsApplication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSamlApplicationSettingsApplication() *SamlApplicationSettingsApplication {
	this := SamlApplicationSettingsApplication{}
	return &this
}

// NewSamlApplicationSettingsApplicationWithDefaults instantiates a new SamlApplicationSettingsApplication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSamlApplicationSettingsApplicationWithDefaults() *SamlApplicationSettingsApplication {
	this := SamlApplicationSettingsApplication{}
	return &this
}

// GetAcsUrl returns the AcsUrl field value if set, zero value otherwise.
func (o *SamlApplicationSettingsApplication) GetAcsUrl() string {
	if o == nil || o.AcsUrl == nil {
		var ret string
		return ret
	}
	return *o.AcsUrl
}

// GetAcsUrlOk returns a tuple with the AcsUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlApplicationSettingsApplication) GetAcsUrlOk() (*string, bool) {
	if o == nil || o.AcsUrl == nil {
		return nil, false
	}
	return o.AcsUrl, true
}

// HasAcsUrl returns a boolean if a field has been set.
func (o *SamlApplicationSettingsApplication) HasAcsUrl() bool {
	if o != nil && o.AcsUrl != nil {
		return true
	}

	return false
}

// SetAcsUrl gets a reference to the given string and assigns it to the AcsUrl field.
func (o *SamlApplicationSettingsApplication) SetAcsUrl(v string) {
	o.AcsUrl = &v
}

// GetAudRestriction returns the AudRestriction field value if set, zero value otherwise.
func (o *SamlApplicationSettingsApplication) GetAudRestriction() string {
	if o == nil || o.AudRestriction == nil {
		var ret string
		return ret
	}
	return *o.AudRestriction
}

// GetAudRestrictionOk returns a tuple with the AudRestriction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlApplicationSettingsApplication) GetAudRestrictionOk() (*string, bool) {
	if o == nil || o.AudRestriction == nil {
		return nil, false
	}
	return o.AudRestriction, true
}

// HasAudRestriction returns a boolean if a field has been set.
func (o *SamlApplicationSettingsApplication) HasAudRestriction() bool {
	if o != nil && o.AudRestriction != nil {
		return true
	}

	return false
}

// SetAudRestriction gets a reference to the given string and assigns it to the AudRestriction field.
func (o *SamlApplicationSettingsApplication) SetAudRestriction(v string) {
	o.AudRestriction = &v
}

// GetBaseUrl returns the BaseUrl field value if set, zero value otherwise.
func (o *SamlApplicationSettingsApplication) GetBaseUrl() string {
	if o == nil || o.BaseUrl == nil {
		var ret string
		return ret
	}
	return *o.BaseUrl
}

// GetBaseUrlOk returns a tuple with the BaseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlApplicationSettingsApplication) GetBaseUrlOk() (*string, bool) {
	if o == nil || o.BaseUrl == nil {
		return nil, false
	}
	return o.BaseUrl, true
}

// HasBaseUrl returns a boolean if a field has been set.
func (o *SamlApplicationSettingsApplication) HasBaseUrl() bool {
	if o != nil && o.BaseUrl != nil {
		return true
	}

	return false
}

// SetBaseUrl gets a reference to the given string and assigns it to the BaseUrl field.
func (o *SamlApplicationSettingsApplication) SetBaseUrl(v string) {
	o.BaseUrl = &v
}

func (o SamlApplicationSettingsApplication) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AcsUrl != nil {
		toSerialize["acsUrl"] = o.AcsUrl
	}
	if o.AudRestriction != nil {
		toSerialize["audRestriction"] = o.AudRestriction
	}
	if o.BaseUrl != nil {
		toSerialize["baseUrl"] = o.BaseUrl
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SamlApplicationSettingsApplication) UnmarshalJSON(bytes []byte) (err error) {
	varSamlApplicationSettingsApplication := _SamlApplicationSettingsApplication{}

	err = json.Unmarshal(bytes, &varSamlApplicationSettingsApplication)
	if err == nil {
		*o = SamlApplicationSettingsApplication(varSamlApplicationSettingsApplication)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "acsUrl")
		delete(additionalProperties, "audRestriction")
		delete(additionalProperties, "baseUrl")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableSamlApplicationSettingsApplication struct {
	value *SamlApplicationSettingsApplication
	isSet bool
}

func (v NullableSamlApplicationSettingsApplication) Get() *SamlApplicationSettingsApplication {
	return v.value
}

func (v *NullableSamlApplicationSettingsApplication) Set(val *SamlApplicationSettingsApplication) {
	v.value = val
	v.isSet = true
}

func (v NullableSamlApplicationSettingsApplication) IsSet() bool {
	return v.isSet
}

func (v *NullableSamlApplicationSettingsApplication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSamlApplicationSettingsApplication(val *SamlApplicationSettingsApplication) *NullableSamlApplicationSettingsApplication {
	return &NullableSamlApplicationSettingsApplication{value: val, isSet: true}
}

func (v NullableSamlApplicationSettingsApplication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSamlApplicationSettingsApplication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

