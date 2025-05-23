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

// SimulateResultRules struct for SimulateResultRules
type SimulateResultRules struct {
	// List of all conditions involved for this rule evaluation
	Conditions []SimulateResultConditions `json:"conditions,omitempty"`
	// The unique ID number of the policy rule
	Id *string `json:"id,omitempty"`
	// The name of the policy rule
	Name *string `json:"name,omitempty"`
	// The result of this entity evaluation
	Status *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SimulateResultRules SimulateResultRules

// NewSimulateResultRules instantiates a new SimulateResultRules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimulateResultRules() *SimulateResultRules {
	this := SimulateResultRules{}
	return &this
}

// NewSimulateResultRulesWithDefaults instantiates a new SimulateResultRules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimulateResultRulesWithDefaults() *SimulateResultRules {
	this := SimulateResultRules{}
	return &this
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *SimulateResultRules) GetConditions() []SimulateResultConditions {
	if o == nil || o.Conditions == nil {
		var ret []SimulateResultConditions
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimulateResultRules) GetConditionsOk() ([]SimulateResultConditions, bool) {
	if o == nil || o.Conditions == nil {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *SimulateResultRules) HasConditions() bool {
	if o != nil && o.Conditions != nil {
		return true
	}

	return false
}

// SetConditions gets a reference to the given []SimulateResultConditions and assigns it to the Conditions field.
func (o *SimulateResultRules) SetConditions(v []SimulateResultConditions) {
	o.Conditions = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SimulateResultRules) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimulateResultRules) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SimulateResultRules) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SimulateResultRules) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SimulateResultRules) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimulateResultRules) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SimulateResultRules) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SimulateResultRules) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SimulateResultRules) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimulateResultRules) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SimulateResultRules) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SimulateResultRules) SetStatus(v string) {
	o.Status = &v
}

func (o SimulateResultRules) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Conditions != nil {
		toSerialize["conditions"] = o.Conditions
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SimulateResultRules) UnmarshalJSON(bytes []byte) (err error) {
	varSimulateResultRules := _SimulateResultRules{}

	err = json.Unmarshal(bytes, &varSimulateResultRules)
	if err == nil {
		*o = SimulateResultRules(varSimulateResultRules)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "conditions")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableSimulateResultRules struct {
	value *SimulateResultRules
	isSet bool
}

func (v NullableSimulateResultRules) Get() *SimulateResultRules {
	return v.value
}

func (v *NullableSimulateResultRules) Set(val *SimulateResultRules) {
	v.value = val
	v.isSet = true
}

func (v NullableSimulateResultRules) IsSet() bool {
	return v.isSet
}

func (v *NullableSimulateResultRules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimulateResultRules(val *SimulateResultRules) *NullableSimulateResultRules {
	return &NullableSimulateResultRules{value: val, isSet: true}
}

func (v NullableSimulateResultRules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimulateResultRules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

