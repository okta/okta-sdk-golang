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

// MDMEnrollmentPolicyRuleCondition struct for MDMEnrollmentPolicyRuleCondition
type MDMEnrollmentPolicyRuleCondition struct {
	BlockNonSafeAndroid *bool `json:"blockNonSafeAndroid,omitempty"`
	Enrollment *string `json:"enrollment,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MDMEnrollmentPolicyRuleCondition MDMEnrollmentPolicyRuleCondition

// NewMDMEnrollmentPolicyRuleCondition instantiates a new MDMEnrollmentPolicyRuleCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMDMEnrollmentPolicyRuleCondition() *MDMEnrollmentPolicyRuleCondition {
	this := MDMEnrollmentPolicyRuleCondition{}
	return &this
}

// NewMDMEnrollmentPolicyRuleConditionWithDefaults instantiates a new MDMEnrollmentPolicyRuleCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMDMEnrollmentPolicyRuleConditionWithDefaults() *MDMEnrollmentPolicyRuleCondition {
	this := MDMEnrollmentPolicyRuleCondition{}
	return &this
}

// GetBlockNonSafeAndroid returns the BlockNonSafeAndroid field value if set, zero value otherwise.
func (o *MDMEnrollmentPolicyRuleCondition) GetBlockNonSafeAndroid() bool {
	if o == nil || o.BlockNonSafeAndroid == nil {
		var ret bool
		return ret
	}
	return *o.BlockNonSafeAndroid
}

// GetBlockNonSafeAndroidOk returns a tuple with the BlockNonSafeAndroid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MDMEnrollmentPolicyRuleCondition) GetBlockNonSafeAndroidOk() (*bool, bool) {
	if o == nil || o.BlockNonSafeAndroid == nil {
		return nil, false
	}
	return o.BlockNonSafeAndroid, true
}

// HasBlockNonSafeAndroid returns a boolean if a field has been set.
func (o *MDMEnrollmentPolicyRuleCondition) HasBlockNonSafeAndroid() bool {
	if o != nil && o.BlockNonSafeAndroid != nil {
		return true
	}

	return false
}

// SetBlockNonSafeAndroid gets a reference to the given bool and assigns it to the BlockNonSafeAndroid field.
func (o *MDMEnrollmentPolicyRuleCondition) SetBlockNonSafeAndroid(v bool) {
	o.BlockNonSafeAndroid = &v
}

// GetEnrollment returns the Enrollment field value if set, zero value otherwise.
func (o *MDMEnrollmentPolicyRuleCondition) GetEnrollment() string {
	if o == nil || o.Enrollment == nil {
		var ret string
		return ret
	}
	return *o.Enrollment
}

// GetEnrollmentOk returns a tuple with the Enrollment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MDMEnrollmentPolicyRuleCondition) GetEnrollmentOk() (*string, bool) {
	if o == nil || o.Enrollment == nil {
		return nil, false
	}
	return o.Enrollment, true
}

// HasEnrollment returns a boolean if a field has been set.
func (o *MDMEnrollmentPolicyRuleCondition) HasEnrollment() bool {
	if o != nil && o.Enrollment != nil {
		return true
	}

	return false
}

// SetEnrollment gets a reference to the given string and assigns it to the Enrollment field.
func (o *MDMEnrollmentPolicyRuleCondition) SetEnrollment(v string) {
	o.Enrollment = &v
}

func (o MDMEnrollmentPolicyRuleCondition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BlockNonSafeAndroid != nil {
		toSerialize["blockNonSafeAndroid"] = o.BlockNonSafeAndroid
	}
	if o.Enrollment != nil {
		toSerialize["enrollment"] = o.Enrollment
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MDMEnrollmentPolicyRuleCondition) UnmarshalJSON(bytes []byte) (err error) {
	varMDMEnrollmentPolicyRuleCondition := _MDMEnrollmentPolicyRuleCondition{}

	err = json.Unmarshal(bytes, &varMDMEnrollmentPolicyRuleCondition)
	if err == nil {
		*o = MDMEnrollmentPolicyRuleCondition(varMDMEnrollmentPolicyRuleCondition)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "blockNonSafeAndroid")
		delete(additionalProperties, "enrollment")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableMDMEnrollmentPolicyRuleCondition struct {
	value *MDMEnrollmentPolicyRuleCondition
	isSet bool
}

func (v NullableMDMEnrollmentPolicyRuleCondition) Get() *MDMEnrollmentPolicyRuleCondition {
	return v.value
}

func (v *NullableMDMEnrollmentPolicyRuleCondition) Set(val *MDMEnrollmentPolicyRuleCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableMDMEnrollmentPolicyRuleCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableMDMEnrollmentPolicyRuleCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMDMEnrollmentPolicyRuleCondition(val *MDMEnrollmentPolicyRuleCondition) *NullableMDMEnrollmentPolicyRuleCondition {
	return &NullableMDMEnrollmentPolicyRuleCondition{value: val, isSet: true}
}

func (v NullableMDMEnrollmentPolicyRuleCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMDMEnrollmentPolicyRuleCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

