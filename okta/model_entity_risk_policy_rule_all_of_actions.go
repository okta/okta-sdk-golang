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

// EntityRiskPolicyRuleAllOfActions The action to take based on the risk event
type EntityRiskPolicyRuleAllOfActions struct {
	EntityRisk *EntityRiskPolicyRuleAllOfActionsEntityRisk `json:"entityRisk,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EntityRiskPolicyRuleAllOfActions EntityRiskPolicyRuleAllOfActions

// NewEntityRiskPolicyRuleAllOfActions instantiates a new EntityRiskPolicyRuleAllOfActions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntityRiskPolicyRuleAllOfActions() *EntityRiskPolicyRuleAllOfActions {
	this := EntityRiskPolicyRuleAllOfActions{}
	return &this
}

// NewEntityRiskPolicyRuleAllOfActionsWithDefaults instantiates a new EntityRiskPolicyRuleAllOfActions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntityRiskPolicyRuleAllOfActionsWithDefaults() *EntityRiskPolicyRuleAllOfActions {
	this := EntityRiskPolicyRuleAllOfActions{}
	return &this
}

// GetEntityRisk returns the EntityRisk field value if set, zero value otherwise.
func (o *EntityRiskPolicyRuleAllOfActions) GetEntityRisk() EntityRiskPolicyRuleAllOfActionsEntityRisk {
	if o == nil || o.EntityRisk == nil {
		var ret EntityRiskPolicyRuleAllOfActionsEntityRisk
		return ret
	}
	return *o.EntityRisk
}

// GetEntityRiskOk returns a tuple with the EntityRisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityRiskPolicyRuleAllOfActions) GetEntityRiskOk() (*EntityRiskPolicyRuleAllOfActionsEntityRisk, bool) {
	if o == nil || o.EntityRisk == nil {
		return nil, false
	}
	return o.EntityRisk, true
}

// HasEntityRisk returns a boolean if a field has been set.
func (o *EntityRiskPolicyRuleAllOfActions) HasEntityRisk() bool {
	if o != nil && o.EntityRisk != nil {
		return true
	}

	return false
}

// SetEntityRisk gets a reference to the given EntityRiskPolicyRuleAllOfActionsEntityRisk and assigns it to the EntityRisk field.
func (o *EntityRiskPolicyRuleAllOfActions) SetEntityRisk(v EntityRiskPolicyRuleAllOfActionsEntityRisk) {
	o.EntityRisk = &v
}

func (o EntityRiskPolicyRuleAllOfActions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EntityRisk != nil {
		toSerialize["entityRisk"] = o.EntityRisk
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EntityRiskPolicyRuleAllOfActions) UnmarshalJSON(bytes []byte) (err error) {
	varEntityRiskPolicyRuleAllOfActions := _EntityRiskPolicyRuleAllOfActions{}

	err = json.Unmarshal(bytes, &varEntityRiskPolicyRuleAllOfActions)
	if err == nil {
		*o = EntityRiskPolicyRuleAllOfActions(varEntityRiskPolicyRuleAllOfActions)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "entityRisk")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableEntityRiskPolicyRuleAllOfActions struct {
	value *EntityRiskPolicyRuleAllOfActions
	isSet bool
}

func (v NullableEntityRiskPolicyRuleAllOfActions) Get() *EntityRiskPolicyRuleAllOfActions {
	return v.value
}

func (v *NullableEntityRiskPolicyRuleAllOfActions) Set(val *EntityRiskPolicyRuleAllOfActions) {
	v.value = val
	v.isSet = true
}

func (v NullableEntityRiskPolicyRuleAllOfActions) IsSet() bool {
	return v.isSet
}

func (v *NullableEntityRiskPolicyRuleAllOfActions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntityRiskPolicyRuleAllOfActions(val *EntityRiskPolicyRuleAllOfActions) *NullableEntityRiskPolicyRuleAllOfActions {
	return &NullableEntityRiskPolicyRuleAllOfActions{value: val, isSet: true}
}

func (v NullableEntityRiskPolicyRuleAllOfActions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntityRiskPolicyRuleAllOfActions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

