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
	"fmt"
)

// checks if the SessionViolationDetectionRiskScoreCondition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SessionViolationDetectionRiskScoreCondition{}

// SessionViolationDetectionRiskScoreCondition <x-lifecycle-container><x-lifecycle class=\"oie\"></x-lifecycle></x-lifecycle-container>Specifies the minimum risk score level required to trigger the rule
type SessionViolationDetectionRiskScoreCondition struct {
	// The minimum risk level to match.
	MinRiskLevel         string `json:"minRiskLevel"`
	AdditionalProperties map[string]interface{}
}

type _SessionViolationDetectionRiskScoreCondition SessionViolationDetectionRiskScoreCondition

// NewSessionViolationDetectionRiskScoreCondition instantiates a new SessionViolationDetectionRiskScoreCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionViolationDetectionRiskScoreCondition(minRiskLevel string) *SessionViolationDetectionRiskScoreCondition {
	this := SessionViolationDetectionRiskScoreCondition{}
	this.MinRiskLevel = minRiskLevel
	return &this
}

// NewSessionViolationDetectionRiskScoreConditionWithDefaults instantiates a new SessionViolationDetectionRiskScoreCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionViolationDetectionRiskScoreConditionWithDefaults() *SessionViolationDetectionRiskScoreCondition {
	this := SessionViolationDetectionRiskScoreCondition{}
	return &this
}

// GetMinRiskLevel returns the MinRiskLevel field value
func (o *SessionViolationDetectionRiskScoreCondition) GetMinRiskLevel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MinRiskLevel
}

// GetMinRiskLevelOk returns a tuple with the MinRiskLevel field value
// and a boolean to check if the value has been set.
func (o *SessionViolationDetectionRiskScoreCondition) GetMinRiskLevelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinRiskLevel, true
}

// SetMinRiskLevel sets field value
func (o *SessionViolationDetectionRiskScoreCondition) SetMinRiskLevel(v string) {
	o.MinRiskLevel = v
}

func (o SessionViolationDetectionRiskScoreCondition) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SessionViolationDetectionRiskScoreCondition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["minRiskLevel"] = o.MinRiskLevel

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SessionViolationDetectionRiskScoreCondition) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"minRiskLevel",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSessionViolationDetectionRiskScoreCondition := _SessionViolationDetectionRiskScoreCondition{}

	err = json.Unmarshal(data, &varSessionViolationDetectionRiskScoreCondition)

	if err != nil {
		return err
	}

	*o = SessionViolationDetectionRiskScoreCondition(varSessionViolationDetectionRiskScoreCondition)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "minRiskLevel")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSessionViolationDetectionRiskScoreCondition struct {
	value *SessionViolationDetectionRiskScoreCondition
	isSet bool
}

func (v NullableSessionViolationDetectionRiskScoreCondition) Get() *SessionViolationDetectionRiskScoreCondition {
	return v.value
}

func (v *NullableSessionViolationDetectionRiskScoreCondition) Set(val *SessionViolationDetectionRiskScoreCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionViolationDetectionRiskScoreCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionViolationDetectionRiskScoreCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionViolationDetectionRiskScoreCondition(val *SessionViolationDetectionRiskScoreCondition) *NullableSessionViolationDetectionRiskScoreCondition {
	return &NullableSessionViolationDetectionRiskScoreCondition{value: val, isSet: true}
}

func (v NullableSessionViolationDetectionRiskScoreCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionViolationDetectionRiskScoreCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
