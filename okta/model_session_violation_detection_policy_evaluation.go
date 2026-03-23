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
)

// checks if the SessionViolationDetectionPolicyEvaluation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SessionViolationDetectionPolicyEvaluation{}

// SessionViolationDetectionPolicyEvaluation <x-lifecycle-container><x-lifecycle class=\"oie\"></x-lifecycle></x-lifecycle-container>Used to control evaluation of the session sign-on policies
type SessionViolationDetectionPolicyEvaluation struct {
	// When true, the sign-on policies of the session are evaluated
	Enabled              *bool `json:"enabled,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SessionViolationDetectionPolicyEvaluation SessionViolationDetectionPolicyEvaluation

// NewSessionViolationDetectionPolicyEvaluation instantiates a new SessionViolationDetectionPolicyEvaluation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionViolationDetectionPolicyEvaluation() *SessionViolationDetectionPolicyEvaluation {
	this := SessionViolationDetectionPolicyEvaluation{}
	return &this
}

// NewSessionViolationDetectionPolicyEvaluationWithDefaults instantiates a new SessionViolationDetectionPolicyEvaluation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionViolationDetectionPolicyEvaluationWithDefaults() *SessionViolationDetectionPolicyEvaluation {
	this := SessionViolationDetectionPolicyEvaluation{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SessionViolationDetectionPolicyEvaluation) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionViolationDetectionPolicyEvaluation) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SessionViolationDetectionPolicyEvaluation) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SessionViolationDetectionPolicyEvaluation) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o SessionViolationDetectionPolicyEvaluation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SessionViolationDetectionPolicyEvaluation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SessionViolationDetectionPolicyEvaluation) UnmarshalJSON(data []byte) (err error) {
	varSessionViolationDetectionPolicyEvaluation := _SessionViolationDetectionPolicyEvaluation{}

	err = json.Unmarshal(data, &varSessionViolationDetectionPolicyEvaluation)

	if err != nil {
		return err
	}

	*o = SessionViolationDetectionPolicyEvaluation(varSessionViolationDetectionPolicyEvaluation)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "enabled")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSessionViolationDetectionPolicyEvaluation struct {
	value *SessionViolationDetectionPolicyEvaluation
	isSet bool
}

func (v NullableSessionViolationDetectionPolicyEvaluation) Get() *SessionViolationDetectionPolicyEvaluation {
	return v.value
}

func (v *NullableSessionViolationDetectionPolicyEvaluation) Set(val *SessionViolationDetectionPolicyEvaluation) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionViolationDetectionPolicyEvaluation) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionViolationDetectionPolicyEvaluation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionViolationDetectionPolicyEvaluation(val *SessionViolationDetectionPolicyEvaluation) *NullableSessionViolationDetectionPolicyEvaluation {
	return &NullableSessionViolationDetectionPolicyEvaluation{value: val, isSet: true}
}

func (v NullableSessionViolationDetectionPolicyEvaluation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionViolationDetectionPolicyEvaluation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
