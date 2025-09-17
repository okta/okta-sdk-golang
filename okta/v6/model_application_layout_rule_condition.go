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

// checks if the ApplicationLayoutRuleCondition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationLayoutRuleCondition{}

// ApplicationLayoutRuleCondition struct for ApplicationLayoutRuleCondition
type ApplicationLayoutRuleCondition struct {
	Schema               map[string]interface{} `json:"schema,omitempty"`
	Scope                *string                `json:"scope,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplicationLayoutRuleCondition ApplicationLayoutRuleCondition

// NewApplicationLayoutRuleCondition instantiates a new ApplicationLayoutRuleCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationLayoutRuleCondition() *ApplicationLayoutRuleCondition {
	this := ApplicationLayoutRuleCondition{}
	return &this
}

// NewApplicationLayoutRuleConditionWithDefaults instantiates a new ApplicationLayoutRuleCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationLayoutRuleConditionWithDefaults() *ApplicationLayoutRuleCondition {
	this := ApplicationLayoutRuleCondition{}
	return &this
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *ApplicationLayoutRuleCondition) GetSchema() map[string]interface{} {
	if o == nil || IsNil(o.Schema) {
		var ret map[string]interface{}
		return ret
	}
	return o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationLayoutRuleCondition) GetSchemaOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Schema) {
		return map[string]interface{}{}, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *ApplicationLayoutRuleCondition) HasSchema() bool {
	if o != nil && !IsNil(o.Schema) {
		return true
	}

	return false
}

// SetSchema gets a reference to the given map[string]interface{} and assigns it to the Schema field.
func (o *ApplicationLayoutRuleCondition) SetSchema(v map[string]interface{}) {
	o.Schema = v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *ApplicationLayoutRuleCondition) GetScope() string {
	if o == nil || IsNil(o.Scope) {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationLayoutRuleCondition) GetScopeOk() (*string, bool) {
	if o == nil || IsNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *ApplicationLayoutRuleCondition) HasScope() bool {
	if o != nil && !IsNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *ApplicationLayoutRuleCondition) SetScope(v string) {
	o.Scope = &v
}

func (o ApplicationLayoutRuleCondition) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationLayoutRuleCondition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Schema) {
		toSerialize["schema"] = o.Schema
	}
	if !IsNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApplicationLayoutRuleCondition) UnmarshalJSON(data []byte) (err error) {
	varApplicationLayoutRuleCondition := _ApplicationLayoutRuleCondition{}

	err = json.Unmarshal(data, &varApplicationLayoutRuleCondition)

	if err != nil {
		return err
	}

	*o = ApplicationLayoutRuleCondition(varApplicationLayoutRuleCondition)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "schema")
		delete(additionalProperties, "scope")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApplicationLayoutRuleCondition struct {
	value *ApplicationLayoutRuleCondition
	isSet bool
}

func (v NullableApplicationLayoutRuleCondition) Get() *ApplicationLayoutRuleCondition {
	return v.value
}

func (v *NullableApplicationLayoutRuleCondition) Set(val *ApplicationLayoutRuleCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationLayoutRuleCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationLayoutRuleCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationLayoutRuleCondition(val *ApplicationLayoutRuleCondition) *NullableApplicationLayoutRuleCondition {
	return &NullableApplicationLayoutRuleCondition{value: val, isSet: true}
}

func (v NullableApplicationLayoutRuleCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationLayoutRuleCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
