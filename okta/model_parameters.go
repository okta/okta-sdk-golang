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

// Parameters Attributes used for processing AD Group membership update
type Parameters struct {
	// The update action to take
	Action *string `json:"action,omitempty"`
	// The attribute that tracks group memberships in AD. This should be `member` for AD.
	Attribute *string `json:"attribute,omitempty"`
	// List of user IDs whose group memberships to update
	Values []string `json:"values,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Parameters Parameters

// NewParameters instantiates a new Parameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParameters() *Parameters {
	this := Parameters{}
	return &this
}

// NewParametersWithDefaults instantiates a new Parameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParametersWithDefaults() *Parameters {
	this := Parameters{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *Parameters) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Parameters) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *Parameters) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *Parameters) SetAction(v string) {
	o.Action = &v
}

// GetAttribute returns the Attribute field value if set, zero value otherwise.
func (o *Parameters) GetAttribute() string {
	if o == nil || o.Attribute == nil {
		var ret string
		return ret
	}
	return *o.Attribute
}

// GetAttributeOk returns a tuple with the Attribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Parameters) GetAttributeOk() (*string, bool) {
	if o == nil || o.Attribute == nil {
		return nil, false
	}
	return o.Attribute, true
}

// HasAttribute returns a boolean if a field has been set.
func (o *Parameters) HasAttribute() bool {
	if o != nil && o.Attribute != nil {
		return true
	}

	return false
}

// SetAttribute gets a reference to the given string and assigns it to the Attribute field.
func (o *Parameters) SetAttribute(v string) {
	o.Attribute = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *Parameters) GetValues() []string {
	if o == nil || o.Values == nil {
		var ret []string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Parameters) GetValuesOk() ([]string, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *Parameters) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given []string and assigns it to the Values field.
func (o *Parameters) SetValues(v []string) {
	o.Values = v
}

func (o Parameters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}
	if o.Attribute != nil {
		toSerialize["attribute"] = o.Attribute
	}
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Parameters) UnmarshalJSON(bytes []byte) (err error) {
	varParameters := _Parameters{}

	err = json.Unmarshal(bytes, &varParameters)
	if err == nil {
		*o = Parameters(varParameters)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "action")
		delete(additionalProperties, "attribute")
		delete(additionalProperties, "values")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableParameters struct {
	value *Parameters
	isSet bool
}

func (v NullableParameters) Get() *Parameters {
	return v.value
}

func (v *NullableParameters) Set(val *Parameters) {
	v.value = val
	v.isSet = true
}

func (v NullableParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParameters(val *Parameters) *NullableParameters {
	return &NullableParameters{value: val, isSet: true}
}

func (v NullableParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

