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

// EntityRiskPolicyRuleAllOfConditionsEntityRisk The risk score level of the entity risk policy rule
type EntityRiskPolicyRuleAllOfConditionsEntityRisk struct {
	Level *string `json:"level,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EntityRiskPolicyRuleAllOfConditionsEntityRisk EntityRiskPolicyRuleAllOfConditionsEntityRisk

// NewEntityRiskPolicyRuleAllOfConditionsEntityRisk instantiates a new EntityRiskPolicyRuleAllOfConditionsEntityRisk object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntityRiskPolicyRuleAllOfConditionsEntityRisk() *EntityRiskPolicyRuleAllOfConditionsEntityRisk {
	this := EntityRiskPolicyRuleAllOfConditionsEntityRisk{}
	return &this
}

// NewEntityRiskPolicyRuleAllOfConditionsEntityRiskWithDefaults instantiates a new EntityRiskPolicyRuleAllOfConditionsEntityRisk object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntityRiskPolicyRuleAllOfConditionsEntityRiskWithDefaults() *EntityRiskPolicyRuleAllOfConditionsEntityRisk {
	this := EntityRiskPolicyRuleAllOfConditionsEntityRisk{}
	return &this
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *EntityRiskPolicyRuleAllOfConditionsEntityRisk) GetLevel() string {
	if o == nil || o.Level == nil {
		var ret string
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityRiskPolicyRuleAllOfConditionsEntityRisk) GetLevelOk() (*string, bool) {
	if o == nil || o.Level == nil {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *EntityRiskPolicyRuleAllOfConditionsEntityRisk) HasLevel() bool {
	if o != nil && o.Level != nil {
		return true
	}

	return false
}

// SetLevel gets a reference to the given string and assigns it to the Level field.
func (o *EntityRiskPolicyRuleAllOfConditionsEntityRisk) SetLevel(v string) {
	o.Level = &v
}

func (o EntityRiskPolicyRuleAllOfConditionsEntityRisk) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Level != nil {
		toSerialize["level"] = o.Level
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EntityRiskPolicyRuleAllOfConditionsEntityRisk) UnmarshalJSON(bytes []byte) (err error) {
	varEntityRiskPolicyRuleAllOfConditionsEntityRisk := _EntityRiskPolicyRuleAllOfConditionsEntityRisk{}

	err = json.Unmarshal(bytes, &varEntityRiskPolicyRuleAllOfConditionsEntityRisk)
	if err == nil {
		*o = EntityRiskPolicyRuleAllOfConditionsEntityRisk(varEntityRiskPolicyRuleAllOfConditionsEntityRisk)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "level")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableEntityRiskPolicyRuleAllOfConditionsEntityRisk struct {
	value *EntityRiskPolicyRuleAllOfConditionsEntityRisk
	isSet bool
}

func (v NullableEntityRiskPolicyRuleAllOfConditionsEntityRisk) Get() *EntityRiskPolicyRuleAllOfConditionsEntityRisk {
	return v.value
}

func (v *NullableEntityRiskPolicyRuleAllOfConditionsEntityRisk) Set(val *EntityRiskPolicyRuleAllOfConditionsEntityRisk) {
	v.value = val
	v.isSet = true
}

func (v NullableEntityRiskPolicyRuleAllOfConditionsEntityRisk) IsSet() bool {
	return v.isSet
}

func (v *NullableEntityRiskPolicyRuleAllOfConditionsEntityRisk) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntityRiskPolicyRuleAllOfConditionsEntityRisk(val *EntityRiskPolicyRuleAllOfConditionsEntityRisk) *NullableEntityRiskPolicyRuleAllOfConditionsEntityRisk {
	return &NullableEntityRiskPolicyRuleAllOfConditionsEntityRisk{value: val, isSet: true}
}

func (v NullableEntityRiskPolicyRuleAllOfConditionsEntityRisk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntityRiskPolicyRuleAllOfConditionsEntityRisk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

