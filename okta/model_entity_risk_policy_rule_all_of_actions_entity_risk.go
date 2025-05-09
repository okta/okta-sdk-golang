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

// EntityRiskPolicyRuleAllOfActionsEntityRisk The object that contains the `actions` array
type EntityRiskPolicyRuleAllOfActionsEntityRisk struct {
	// The `entityRisk` object's `actions` array can be empty or contain one of two `action` object value pairs. This object determines the specific response to a risk event.
	Actions []EntityRiskPolicyRuleActionsObject `json:"actions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EntityRiskPolicyRuleAllOfActionsEntityRisk EntityRiskPolicyRuleAllOfActionsEntityRisk

// NewEntityRiskPolicyRuleAllOfActionsEntityRisk instantiates a new EntityRiskPolicyRuleAllOfActionsEntityRisk object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntityRiskPolicyRuleAllOfActionsEntityRisk() *EntityRiskPolicyRuleAllOfActionsEntityRisk {
	this := EntityRiskPolicyRuleAllOfActionsEntityRisk{}
	return &this
}

// NewEntityRiskPolicyRuleAllOfActionsEntityRiskWithDefaults instantiates a new EntityRiskPolicyRuleAllOfActionsEntityRisk object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntityRiskPolicyRuleAllOfActionsEntityRiskWithDefaults() *EntityRiskPolicyRuleAllOfActionsEntityRisk {
	this := EntityRiskPolicyRuleAllOfActionsEntityRisk{}
	return &this
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *EntityRiskPolicyRuleAllOfActionsEntityRisk) GetActions() []EntityRiskPolicyRuleActionsObject {
	if o == nil || o.Actions == nil {
		var ret []EntityRiskPolicyRuleActionsObject
		return ret
	}
	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityRiskPolicyRuleAllOfActionsEntityRisk) GetActionsOk() ([]EntityRiskPolicyRuleActionsObject, bool) {
	if o == nil || o.Actions == nil {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *EntityRiskPolicyRuleAllOfActionsEntityRisk) HasActions() bool {
	if o != nil && o.Actions != nil {
		return true
	}

	return false
}

// SetActions gets a reference to the given []EntityRiskPolicyRuleActionsObject and assigns it to the Actions field.
func (o *EntityRiskPolicyRuleAllOfActionsEntityRisk) SetActions(v []EntityRiskPolicyRuleActionsObject) {
	o.Actions = v
}

func (o EntityRiskPolicyRuleAllOfActionsEntityRisk) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Actions != nil {
		toSerialize["actions"] = o.Actions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EntityRiskPolicyRuleAllOfActionsEntityRisk) UnmarshalJSON(bytes []byte) (err error) {
	varEntityRiskPolicyRuleAllOfActionsEntityRisk := _EntityRiskPolicyRuleAllOfActionsEntityRisk{}

	err = json.Unmarshal(bytes, &varEntityRiskPolicyRuleAllOfActionsEntityRisk)
	if err == nil {
		*o = EntityRiskPolicyRuleAllOfActionsEntityRisk(varEntityRiskPolicyRuleAllOfActionsEntityRisk)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "actions")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableEntityRiskPolicyRuleAllOfActionsEntityRisk struct {
	value *EntityRiskPolicyRuleAllOfActionsEntityRisk
	isSet bool
}

func (v NullableEntityRiskPolicyRuleAllOfActionsEntityRisk) Get() *EntityRiskPolicyRuleAllOfActionsEntityRisk {
	return v.value
}

func (v *NullableEntityRiskPolicyRuleAllOfActionsEntityRisk) Set(val *EntityRiskPolicyRuleAllOfActionsEntityRisk) {
	v.value = val
	v.isSet = true
}

func (v NullableEntityRiskPolicyRuleAllOfActionsEntityRisk) IsSet() bool {
	return v.isSet
}

func (v *NullableEntityRiskPolicyRuleAllOfActionsEntityRisk) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntityRiskPolicyRuleAllOfActionsEntityRisk(val *EntityRiskPolicyRuleAllOfActionsEntityRisk) *NullableEntityRiskPolicyRuleAllOfActionsEntityRisk {
	return &NullableEntityRiskPolicyRuleAllOfActionsEntityRisk{value: val, isSet: true}
}

func (v NullableEntityRiskPolicyRuleAllOfActionsEntityRisk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntityRiskPolicyRuleAllOfActionsEntityRisk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

