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

// PlatformConditionEvaluatorPlatformOperatingSystemVersion struct for PlatformConditionEvaluatorPlatformOperatingSystemVersion
type PlatformConditionEvaluatorPlatformOperatingSystemVersion struct {
	MatchType *string `json:"matchType,omitempty"`
	Value *string `json:"value,omitempty"`
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
	if o == nil || o.MatchType == nil {
		var ret string
		return ret
	}
	return *o.MatchType
}

// GetMatchTypeOk returns a tuple with the MatchType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformConditionEvaluatorPlatformOperatingSystemVersion) GetMatchTypeOk() (*string, bool) {
	if o == nil || o.MatchType == nil {
		return nil, false
	}
	return o.MatchType, true
}

// HasMatchType returns a boolean if a field has been set.
func (o *PlatformConditionEvaluatorPlatformOperatingSystemVersion) HasMatchType() bool {
	if o != nil && o.MatchType != nil {
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
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformConditionEvaluatorPlatformOperatingSystemVersion) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *PlatformConditionEvaluatorPlatformOperatingSystemVersion) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *PlatformConditionEvaluatorPlatformOperatingSystemVersion) SetValue(v string) {
	o.Value = &v
}

func (o PlatformConditionEvaluatorPlatformOperatingSystemVersion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MatchType != nil {
		toSerialize["matchType"] = o.MatchType
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PlatformConditionEvaluatorPlatformOperatingSystemVersion) UnmarshalJSON(bytes []byte) (err error) {
	varPlatformConditionEvaluatorPlatformOperatingSystemVersion := _PlatformConditionEvaluatorPlatformOperatingSystemVersion{}

	err = json.Unmarshal(bytes, &varPlatformConditionEvaluatorPlatformOperatingSystemVersion)
	if err == nil {
		*o = PlatformConditionEvaluatorPlatformOperatingSystemVersion(varPlatformConditionEvaluatorPlatformOperatingSystemVersion)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "matchType")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
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

