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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the SimulateResultPoliciesItems type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SimulateResultPoliciesItems{}

// SimulateResultPoliciesItems struct for SimulateResultPoliciesItems
type SimulateResultPoliciesItems struct {
	// List of all conditions involved for this policy evaluation
	Conditions []SimulateResultConditions `json:"conditions,omitempty"`
	// ID of the specified policy type
	Id *string `json:"id,omitempty"`
	// Policy name
	Name  *string               `json:"name,omitempty"`
	Rules []SimulateResultRules `json:"rules,omitempty"`
	// The result of this entity evaluation
	Status               *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SimulateResultPoliciesItems SimulateResultPoliciesItems

// NewSimulateResultPoliciesItems instantiates a new SimulateResultPoliciesItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimulateResultPoliciesItems() *SimulateResultPoliciesItems {
	this := SimulateResultPoliciesItems{}
	return &this
}

// NewSimulateResultPoliciesItemsWithDefaults instantiates a new SimulateResultPoliciesItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimulateResultPoliciesItemsWithDefaults() *SimulateResultPoliciesItems {
	this := SimulateResultPoliciesItems{}
	return &this
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *SimulateResultPoliciesItems) GetConditions() []SimulateResultConditions {
	if o == nil || IsNil(o.Conditions) {
		var ret []SimulateResultConditions
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimulateResultPoliciesItems) GetConditionsOk() ([]SimulateResultConditions, bool) {
	if o == nil || IsNil(o.Conditions) {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *SimulateResultPoliciesItems) HasConditions() bool {
	if o != nil && !IsNil(o.Conditions) {
		return true
	}

	return false
}

// SetConditions gets a reference to the given []SimulateResultConditions and assigns it to the Conditions field.
func (o *SimulateResultPoliciesItems) SetConditions(v []SimulateResultConditions) {
	o.Conditions = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SimulateResultPoliciesItems) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimulateResultPoliciesItems) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SimulateResultPoliciesItems) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SimulateResultPoliciesItems) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SimulateResultPoliciesItems) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimulateResultPoliciesItems) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SimulateResultPoliciesItems) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SimulateResultPoliciesItems) SetName(v string) {
	o.Name = &v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *SimulateResultPoliciesItems) GetRules() []SimulateResultRules {
	if o == nil || IsNil(o.Rules) {
		var ret []SimulateResultRules
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimulateResultPoliciesItems) GetRulesOk() ([]SimulateResultRules, bool) {
	if o == nil || IsNil(o.Rules) {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *SimulateResultPoliciesItems) HasRules() bool {
	if o != nil && !IsNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []SimulateResultRules and assigns it to the Rules field.
func (o *SimulateResultPoliciesItems) SetRules(v []SimulateResultRules) {
	o.Rules = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SimulateResultPoliciesItems) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimulateResultPoliciesItems) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SimulateResultPoliciesItems) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SimulateResultPoliciesItems) SetStatus(v string) {
	o.Status = &v
}

func (o SimulateResultPoliciesItems) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SimulateResultPoliciesItems) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Conditions) {
		toSerialize["conditions"] = o.Conditions
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SimulateResultPoliciesItems) UnmarshalJSON(data []byte) (err error) {
	varSimulateResultPoliciesItems := _SimulateResultPoliciesItems{}

	err = json.Unmarshal(data, &varSimulateResultPoliciesItems)

	if err != nil {
		return err
	}

	*o = SimulateResultPoliciesItems(varSimulateResultPoliciesItems)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "conditions")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "rules")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSimulateResultPoliciesItems struct {
	value *SimulateResultPoliciesItems
	isSet bool
}

func (v NullableSimulateResultPoliciesItems) Get() *SimulateResultPoliciesItems {
	return v.value
}

func (v *NullableSimulateResultPoliciesItems) Set(val *SimulateResultPoliciesItems) {
	v.value = val
	v.isSet = true
}

func (v NullableSimulateResultPoliciesItems) IsSet() bool {
	return v.isSet
}

func (v *NullableSimulateResultPoliciesItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimulateResultPoliciesItems(val *SimulateResultPoliciesItems) *NullableSimulateResultPoliciesItems {
	return &NullableSimulateResultPoliciesItems{value: val, isSet: true}
}

func (v NullableSimulateResultPoliciesItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimulateResultPoliciesItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
