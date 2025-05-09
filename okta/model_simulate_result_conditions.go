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

// SimulateResultConditions struct for SimulateResultConditions
type SimulateResultConditions struct {
	// The result of this entity evaluation
	Status *string `json:"status,omitempty"`
	// The type of condition
	Type *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SimulateResultConditions SimulateResultConditions

// NewSimulateResultConditions instantiates a new SimulateResultConditions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimulateResultConditions() *SimulateResultConditions {
	this := SimulateResultConditions{}
	return &this
}

// NewSimulateResultConditionsWithDefaults instantiates a new SimulateResultConditions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimulateResultConditionsWithDefaults() *SimulateResultConditions {
	this := SimulateResultConditions{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SimulateResultConditions) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimulateResultConditions) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SimulateResultConditions) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SimulateResultConditions) SetStatus(v string) {
	o.Status = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SimulateResultConditions) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimulateResultConditions) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SimulateResultConditions) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SimulateResultConditions) SetType(v string) {
	o.Type = &v
}

func (o SimulateResultConditions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SimulateResultConditions) UnmarshalJSON(bytes []byte) (err error) {
	varSimulateResultConditions := _SimulateResultConditions{}

	err = json.Unmarshal(bytes, &varSimulateResultConditions)
	if err == nil {
		*o = SimulateResultConditions(varSimulateResultConditions)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "status")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableSimulateResultConditions struct {
	value *SimulateResultConditions
	isSet bool
}

func (v NullableSimulateResultConditions) Get() *SimulateResultConditions {
	return v.value
}

func (v *NullableSimulateResultConditions) Set(val *SimulateResultConditions) {
	v.value = val
	v.isSet = true
}

func (v NullableSimulateResultConditions) IsSet() bool {
	return v.isSet
}

func (v *NullableSimulateResultConditions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimulateResultConditions(val *SimulateResultConditions) *NullableSimulateResultConditions {
	return &NullableSimulateResultConditions{value: val, isSet: true}
}

func (v NullableSimulateResultConditions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimulateResultConditions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

