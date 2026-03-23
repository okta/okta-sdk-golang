/*
Okta Admin Management API

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

API version: 2025.08.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"fmt"
)

// checks if the Parameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Parameters{}

// Parameters Attributes used for processing Active Directory or LDAP group membership update
type Parameters struct {
	// The update action to take
	Action string `json:"action"`
	// The attribute that tracks group memberships in Active Directory or LDAP. For Active Directory, use `member`. For LDAP, use the appropriate attribute found in the LDAP server such as, but not limited to, `uniqueMember` or `member`.
	Attribute string `json:"attribute"`
	// List of user IDs whose group memberships to update
	Values               []string `json:"values"`
	AdditionalProperties map[string]interface{}
}

type _Parameters Parameters

// NewParameters instantiates a new Parameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParameters(action string, attribute string, values []string) *Parameters {
	this := Parameters{}
	this.Action = action
	this.Attribute = attribute
	this.Values = values
	return &this
}

// NewParametersWithDefaults instantiates a new Parameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParametersWithDefaults() *Parameters {
	this := Parameters{}
	return &this
}

// GetAction returns the Action field value
func (o *Parameters) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *Parameters) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *Parameters) SetAction(v string) {
	o.Action = v
}

// GetAttribute returns the Attribute field value
func (o *Parameters) GetAttribute() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Attribute
}

// GetAttributeOk returns a tuple with the Attribute field value
// and a boolean to check if the value has been set.
func (o *Parameters) GetAttributeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attribute, true
}

// SetAttribute sets field value
func (o *Parameters) SetAttribute(v string) {
	o.Attribute = v
}

// GetValues returns the Values field value
func (o *Parameters) GetValues() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *Parameters) GetValuesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *Parameters) SetValues(v []string) {
	o.Values = v
}

func (o Parameters) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Parameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["action"] = o.Action
	toSerialize["attribute"] = o.Attribute
	toSerialize["values"] = o.Values

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Parameters) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"action",
		"attribute",
		"values",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varParameters := _Parameters{}

	err = json.Unmarshal(data, &varParameters)

	if err != nil {
		return err
	}

	*o = Parameters(varParameters)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "action")
		delete(additionalProperties, "attribute")
		delete(additionalProperties, "values")
		o.AdditionalProperties = additionalProperties
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
