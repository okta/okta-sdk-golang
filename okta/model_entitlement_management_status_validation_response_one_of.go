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

// checks if the EntitlementManagementStatusValidationResponseOneOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitlementManagementStatusValidationResponseOneOf{}

// EntitlementManagementStatusValidationResponseOneOf The validation was successful and the transition to the requested Entitlement Management state can proceed
type EntitlementManagementStatusValidationResponseOneOf struct {
	Valid                bool `json:"valid"`
	AdditionalProperties map[string]interface{}
}

type _EntitlementManagementStatusValidationResponseOneOf EntitlementManagementStatusValidationResponseOneOf

// NewEntitlementManagementStatusValidationResponseOneOf instantiates a new EntitlementManagementStatusValidationResponseOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementManagementStatusValidationResponseOneOf(valid bool) *EntitlementManagementStatusValidationResponseOneOf {
	this := EntitlementManagementStatusValidationResponseOneOf{}
	this.Valid = valid
	return &this
}

// NewEntitlementManagementStatusValidationResponseOneOfWithDefaults instantiates a new EntitlementManagementStatusValidationResponseOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementManagementStatusValidationResponseOneOfWithDefaults() *EntitlementManagementStatusValidationResponseOneOf {
	this := EntitlementManagementStatusValidationResponseOneOf{}
	return &this
}

// GetValid returns the Valid field value
func (o *EntitlementManagementStatusValidationResponseOneOf) GetValid() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Valid
}

// GetValidOk returns a tuple with the Valid field value
// and a boolean to check if the value has been set.
func (o *EntitlementManagementStatusValidationResponseOneOf) GetValidOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Valid, true
}

// SetValid sets field value
func (o *EntitlementManagementStatusValidationResponseOneOf) SetValid(v bool) {
	o.Valid = v
}

func (o EntitlementManagementStatusValidationResponseOneOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitlementManagementStatusValidationResponseOneOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["valid"] = o.Valid

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EntitlementManagementStatusValidationResponseOneOf) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"valid",
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

	varEntitlementManagementStatusValidationResponseOneOf := _EntitlementManagementStatusValidationResponseOneOf{}

	err = json.Unmarshal(data, &varEntitlementManagementStatusValidationResponseOneOf)

	if err != nil {
		return err
	}

	*o = EntitlementManagementStatusValidationResponseOneOf(varEntitlementManagementStatusValidationResponseOneOf)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "valid")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEntitlementManagementStatusValidationResponseOneOf struct {
	value *EntitlementManagementStatusValidationResponseOneOf
	isSet bool
}

func (v NullableEntitlementManagementStatusValidationResponseOneOf) Get() *EntitlementManagementStatusValidationResponseOneOf {
	return v.value
}

func (v *NullableEntitlementManagementStatusValidationResponseOneOf) Set(val *EntitlementManagementStatusValidationResponseOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementManagementStatusValidationResponseOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementManagementStatusValidationResponseOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementManagementStatusValidationResponseOneOf(val *EntitlementManagementStatusValidationResponseOneOf) *NullableEntitlementManagementStatusValidationResponseOneOf {
	return &NullableEntitlementManagementStatusValidationResponseOneOf{value: val, isSet: true}
}

func (v NullableEntitlementManagementStatusValidationResponseOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementManagementStatusValidationResponseOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
