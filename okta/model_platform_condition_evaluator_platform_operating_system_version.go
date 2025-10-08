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

API version: 2025.08.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the PlatformConditionEvaluatorPlatformOperatingSystemVersion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlatformConditionEvaluatorPlatformOperatingSystemVersion{}

// PlatformConditionEvaluatorPlatformOperatingSystemVersion struct for PlatformConditionEvaluatorPlatformOperatingSystemVersion
type PlatformConditionEvaluatorPlatformOperatingSystemVersion struct {
	MatchType            *string `json:"matchType,omitempty"`
	Value                *string `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PlatformConditionEvaluatorPlatformOperatingSystemVersion PlatformConditionEvaluatorPlatformOperatingSystemVersion

// NewPlatformConditionEvaluatorPlatformOperatingSystemVersion instantiates a new PlatformConditionEvaluatorPlatformOperatingSystemVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlatformConditionEvaluatorPlatformOperatingSystemVersion() *PlatformConditionEvaluatorPlatformOperatingSystemVersion {
	this := PlatformConditionEvaluatorPlatformOperatingSystemVersion{}
	return &this
}

// NewPlatformConditionEvaluatorPlatformOperatingSystemVersionWithDefaults instantiates a new PlatformConditionEvaluatorPlatformOperatingSystemVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlatformConditionEvaluatorPlatformOperatingSystemVersionWithDefaults() *PlatformConditionEvaluatorPlatformOperatingSystemVersion {
	this := PlatformConditionEvaluatorPlatformOperatingSystemVersion{}
	return &this
}

// GetMatchType returns the MatchType field value if set, zero value otherwise.
func (o *PlatformConditionEvaluatorPlatformOperatingSystemVersion) GetMatchType() string {
	if o == nil || IsNil(o.MatchType) {
		var ret string
		return ret
	}
	return *o.MatchType
}

// GetMatchTypeOk returns a tuple with the MatchType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformConditionEvaluatorPlatformOperatingSystemVersion) GetMatchTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MatchType) {
		return nil, false
	}
	return o.MatchType, true
}

// HasMatchType returns a boolean if a field has been set.
func (o *PlatformConditionEvaluatorPlatformOperatingSystemVersion) HasMatchType() bool {
	if o != nil && !IsNil(o.MatchType) {
		return true
	}

	return false
}

// SetMatchType gets a reference to the given string and assigns it to the MatchType field.
func (o *PlatformConditionEvaluatorPlatformOperatingSystemVersion) SetMatchType(v string) {
	o.MatchType = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *PlatformConditionEvaluatorPlatformOperatingSystemVersion) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformConditionEvaluatorPlatformOperatingSystemVersion) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *PlatformConditionEvaluatorPlatformOperatingSystemVersion) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *PlatformConditionEvaluatorPlatformOperatingSystemVersion) SetValue(v string) {
	o.Value = &v
}

func (o PlatformConditionEvaluatorPlatformOperatingSystemVersion) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlatformConditionEvaluatorPlatformOperatingSystemVersion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MatchType) {
		toSerialize["matchType"] = o.MatchType
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PlatformConditionEvaluatorPlatformOperatingSystemVersion) UnmarshalJSON(data []byte) (err error) {
	varPlatformConditionEvaluatorPlatformOperatingSystemVersion := _PlatformConditionEvaluatorPlatformOperatingSystemVersion{}

	err = json.Unmarshal(data, &varPlatformConditionEvaluatorPlatformOperatingSystemVersion)

	if err != nil {
		return err
	}

	*o = PlatformConditionEvaluatorPlatformOperatingSystemVersion(varPlatformConditionEvaluatorPlatformOperatingSystemVersion)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "matchType")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePlatformConditionEvaluatorPlatformOperatingSystemVersion struct {
	value *PlatformConditionEvaluatorPlatformOperatingSystemVersion
	isSet bool
}

func (v NullablePlatformConditionEvaluatorPlatformOperatingSystemVersion) Get() *PlatformConditionEvaluatorPlatformOperatingSystemVersion {
	return v.value
}

func (v *NullablePlatformConditionEvaluatorPlatformOperatingSystemVersion) Set(val *PlatformConditionEvaluatorPlatformOperatingSystemVersion) {
	v.value = val
	v.isSet = true
}

func (v NullablePlatformConditionEvaluatorPlatformOperatingSystemVersion) IsSet() bool {
	return v.isSet
}

func (v *NullablePlatformConditionEvaluatorPlatformOperatingSystemVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlatformConditionEvaluatorPlatformOperatingSystemVersion(val *PlatformConditionEvaluatorPlatformOperatingSystemVersion) *NullablePlatformConditionEvaluatorPlatformOperatingSystemVersion {
	return &NullablePlatformConditionEvaluatorPlatformOperatingSystemVersion{value: val, isSet: true}
}

func (v NullablePlatformConditionEvaluatorPlatformOperatingSystemVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlatformConditionEvaluatorPlatformOperatingSystemVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
