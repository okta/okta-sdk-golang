/*
Okta Admin Management

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

API version: 2025.10.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the SupportedMethods type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SupportedMethods{}

// SupportedMethods The supported methods of an authenticator
type SupportedMethods struct {
	Settings *SupportedMethodsSettings `json:"settings,omitempty"`
	Status   *string                   `json:"status,omitempty"`
	// The type of authenticator method
	Type                 *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SupportedMethods SupportedMethods

// NewSupportedMethods instantiates a new SupportedMethods object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSupportedMethods() *SupportedMethods {
	this := SupportedMethods{}
	return &this
}

// NewSupportedMethodsWithDefaults instantiates a new SupportedMethods object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSupportedMethodsWithDefaults() *SupportedMethods {
	this := SupportedMethods{}
	return &this
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *SupportedMethods) GetSettings() SupportedMethodsSettings {
	if o == nil || IsNil(o.Settings) {
		var ret SupportedMethodsSettings
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedMethods) GetSettingsOk() (*SupportedMethodsSettings, bool) {
	if o == nil || IsNil(o.Settings) {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *SupportedMethods) HasSettings() bool {
	if o != nil && !IsNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given SupportedMethodsSettings and assigns it to the Settings field.
func (o *SupportedMethods) SetSettings(v SupportedMethodsSettings) {
	o.Settings = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SupportedMethods) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedMethods) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SupportedMethods) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SupportedMethods) SetStatus(v string) {
	o.Status = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SupportedMethods) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedMethods) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SupportedMethods) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SupportedMethods) SetType(v string) {
	o.Type = &v
}

func (o SupportedMethods) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SupportedMethods) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Settings) {
		toSerialize["settings"] = o.Settings
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SupportedMethods) UnmarshalJSON(data []byte) (err error) {
	varSupportedMethods := _SupportedMethods{}

	err = json.Unmarshal(data, &varSupportedMethods)

	if err != nil {
		return err
	}

	*o = SupportedMethods(varSupportedMethods)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "settings")
		delete(additionalProperties, "status")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSupportedMethods struct {
	value *SupportedMethods
	isSet bool
}

func (v NullableSupportedMethods) Get() *SupportedMethods {
	return v.value
}

func (v *NullableSupportedMethods) Set(val *SupportedMethods) {
	v.value = val
	v.isSet = true
}

func (v NullableSupportedMethods) IsSet() bool {
	return v.isSet
}

func (v *NullableSupportedMethods) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupportedMethods(val *SupportedMethods) *NullableSupportedMethods {
	return &NullableSupportedMethods{value: val, isSet: true}
}

func (v NullableSupportedMethods) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSupportedMethods) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
