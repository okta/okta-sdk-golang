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

// PlatformConditionEvaluatorPlatform struct for PlatformConditionEvaluatorPlatform
type PlatformConditionEvaluatorPlatform struct {
	Os *PlatformConditionEvaluatorPlatformOperatingSystem `json:"os,omitempty"`
	Type *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PlatformConditionEvaluatorPlatform PlatformConditionEvaluatorPlatform

// NewPlatformConditionEvaluatorPlatform instantiates a new PlatformConditionEvaluatorPlatform object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlatformConditionEvaluatorPlatform() *PlatformConditionEvaluatorPlatform {
	this := PlatformConditionEvaluatorPlatform{}
	return &this
}

// NewPlatformConditionEvaluatorPlatformWithDefaults instantiates a new PlatformConditionEvaluatorPlatform object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlatformConditionEvaluatorPlatformWithDefaults() *PlatformConditionEvaluatorPlatform {
	this := PlatformConditionEvaluatorPlatform{}
	return &this
}

// GetOs returns the Os field value if set, zero value otherwise.
func (o *PlatformConditionEvaluatorPlatform) GetOs() PlatformConditionEvaluatorPlatformOperatingSystem {
	if o == nil || o.Os == nil {
		var ret PlatformConditionEvaluatorPlatformOperatingSystem
		return ret
	}
	return *o.Os
}

// GetOsOk returns a tuple with the Os field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformConditionEvaluatorPlatform) GetOsOk() (*PlatformConditionEvaluatorPlatformOperatingSystem, bool) {
	if o == nil || o.Os == nil {
		return nil, false
	}
	return o.Os, true
}

// HasOs returns a boolean if a field has been set.
func (o *PlatformConditionEvaluatorPlatform) HasOs() bool {
	if o != nil && o.Os != nil {
		return true
	}

	return false
}

// SetOs gets a reference to the given PlatformConditionEvaluatorPlatformOperatingSystem and assigns it to the Os field.
func (o *PlatformConditionEvaluatorPlatform) SetOs(v PlatformConditionEvaluatorPlatformOperatingSystem) {
	o.Os = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PlatformConditionEvaluatorPlatform) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformConditionEvaluatorPlatform) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PlatformConditionEvaluatorPlatform) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PlatformConditionEvaluatorPlatform) SetType(v string) {
	o.Type = &v
}

func (o PlatformConditionEvaluatorPlatform) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Os != nil {
		toSerialize["os"] = o.Os
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PlatformConditionEvaluatorPlatform) UnmarshalJSON(bytes []byte) (err error) {
	varPlatformConditionEvaluatorPlatform := _PlatformConditionEvaluatorPlatform{}

	err = json.Unmarshal(bytes, &varPlatformConditionEvaluatorPlatform)
	if err == nil {
		*o = PlatformConditionEvaluatorPlatform(varPlatformConditionEvaluatorPlatform)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "os")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullablePlatformConditionEvaluatorPlatform struct {
	value *PlatformConditionEvaluatorPlatform
	isSet bool
}

func (v NullablePlatformConditionEvaluatorPlatform) Get() *PlatformConditionEvaluatorPlatform {
	return v.value
}

func (v *NullablePlatformConditionEvaluatorPlatform) Set(val *PlatformConditionEvaluatorPlatform) {
	v.value = val
	v.isSet = true
}

func (v NullablePlatformConditionEvaluatorPlatform) IsSet() bool {
	return v.isSet
}

func (v *NullablePlatformConditionEvaluatorPlatform) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlatformConditionEvaluatorPlatform(val *PlatformConditionEvaluatorPlatform) *NullablePlatformConditionEvaluatorPlatform {
	return &NullablePlatformConditionEvaluatorPlatform{value: val, isSet: true}
}

func (v NullablePlatformConditionEvaluatorPlatform) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlatformConditionEvaluatorPlatform) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

