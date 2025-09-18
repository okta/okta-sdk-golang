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

// checks if the Conditions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Conditions{}

// Conditions Conditions of applying realm assignment
type Conditions struct {
	Expression *Expression `json:"expression,omitempty"`
	// ID of the profile source
	ProfileSourceId      *string `json:"profileSourceId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Conditions Conditions

// NewConditions instantiates a new Conditions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConditions() *Conditions {
	this := Conditions{}
	return &this
}

// NewConditionsWithDefaults instantiates a new Conditions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConditionsWithDefaults() *Conditions {
	this := Conditions{}
	return &this
}

// GetExpression returns the Expression field value if set, zero value otherwise.
func (o *Conditions) GetExpression() Expression {
	if o == nil || IsNil(o.Expression) {
		var ret Expression
		return ret
	}
	return *o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Conditions) GetExpressionOk() (*Expression, bool) {
	if o == nil || IsNil(o.Expression) {
		return nil, false
	}
	return o.Expression, true
}

// HasExpression returns a boolean if a field has been set.
func (o *Conditions) HasExpression() bool {
	if o != nil && !IsNil(o.Expression) {
		return true
	}

	return false
}

// SetExpression gets a reference to the given Expression and assigns it to the Expression field.
func (o *Conditions) SetExpression(v Expression) {
	o.Expression = &v
}

// GetProfileSourceId returns the ProfileSourceId field value if set, zero value otherwise.
func (o *Conditions) GetProfileSourceId() string {
	if o == nil || IsNil(o.ProfileSourceId) {
		var ret string
		return ret
	}
	return *o.ProfileSourceId
}

// GetProfileSourceIdOk returns a tuple with the ProfileSourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Conditions) GetProfileSourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProfileSourceId) {
		return nil, false
	}
	return o.ProfileSourceId, true
}

// HasProfileSourceId returns a boolean if a field has been set.
func (o *Conditions) HasProfileSourceId() bool {
	if o != nil && !IsNil(o.ProfileSourceId) {
		return true
	}

	return false
}

// SetProfileSourceId gets a reference to the given string and assigns it to the ProfileSourceId field.
func (o *Conditions) SetProfileSourceId(v string) {
	o.ProfileSourceId = &v
}

func (o Conditions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Conditions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Expression) {
		toSerialize["expression"] = o.Expression
	}
	if !IsNil(o.ProfileSourceId) {
		toSerialize["profileSourceId"] = o.ProfileSourceId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Conditions) UnmarshalJSON(data []byte) (err error) {
	varConditions := _Conditions{}

	err = json.Unmarshal(data, &varConditions)

	if err != nil {
		return err
	}

	*o = Conditions(varConditions)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "expression")
		delete(additionalProperties, "profileSourceId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConditions struct {
	value *Conditions
	isSet bool
}

func (v NullableConditions) Get() *Conditions {
	return v.value
}

func (v *NullableConditions) Set(val *Conditions) {
	v.value = val
	v.isSet = true
}

func (v NullableConditions) IsSet() bool {
	return v.isSet
}

func (v *NullableConditions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConditions(val *Conditions) *NullableConditions {
	return &NullableConditions{value: val, isSet: true}
}

func (v NullableConditions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConditions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
