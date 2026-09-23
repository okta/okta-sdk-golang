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

// checks if the EntitlementManagementStatusValidationResponseOneOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitlementManagementStatusValidationResponseOneOf1{}

// EntitlementManagementStatusValidationResponseOneOf1 The validation failed and the transition to the requested Entitlement Management state cannot proceed
type EntitlementManagementStatusValidationResponseOneOf1 struct {
	Valid bool `json:"valid"`
	// The machine-readable reason code for blocking the transition to the Entitlement Management state
	ReasonCode string `json:"reasonCode"`
	// Human-readable description of the validation failure.
	Message              string `json:"message"`
	AdditionalProperties map[string]interface{}
}

type _EntitlementManagementStatusValidationResponseOneOf1 EntitlementManagementStatusValidationResponseOneOf1

// NewEntitlementManagementStatusValidationResponseOneOf1 instantiates a new EntitlementManagementStatusValidationResponseOneOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementManagementStatusValidationResponseOneOf1(valid bool, reasonCode string, message string) *EntitlementManagementStatusValidationResponseOneOf1 {
	this := EntitlementManagementStatusValidationResponseOneOf1{}
	this.Valid = valid
	this.ReasonCode = reasonCode
	this.Message = message
	return &this
}

// NewEntitlementManagementStatusValidationResponseOneOf1WithDefaults instantiates a new EntitlementManagementStatusValidationResponseOneOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementManagementStatusValidationResponseOneOf1WithDefaults() *EntitlementManagementStatusValidationResponseOneOf1 {
	this := EntitlementManagementStatusValidationResponseOneOf1{}
	return &this
}

// GetValid returns the Valid field value
func (o *EntitlementManagementStatusValidationResponseOneOf1) GetValid() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Valid
}

// GetValidOk returns a tuple with the Valid field value
// and a boolean to check if the value has been set.
func (o *EntitlementManagementStatusValidationResponseOneOf1) GetValidOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Valid, true
}

// SetValid sets field value
func (o *EntitlementManagementStatusValidationResponseOneOf1) SetValid(v bool) {
	o.Valid = v
}

// GetReasonCode returns the ReasonCode field value
func (o *EntitlementManagementStatusValidationResponseOneOf1) GetReasonCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReasonCode
}

// GetReasonCodeOk returns a tuple with the ReasonCode field value
// and a boolean to check if the value has been set.
func (o *EntitlementManagementStatusValidationResponseOneOf1) GetReasonCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReasonCode, true
}

// SetReasonCode sets field value
func (o *EntitlementManagementStatusValidationResponseOneOf1) SetReasonCode(v string) {
	o.ReasonCode = v
}

// GetMessage returns the Message field value
func (o *EntitlementManagementStatusValidationResponseOneOf1) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *EntitlementManagementStatusValidationResponseOneOf1) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *EntitlementManagementStatusValidationResponseOneOf1) SetMessage(v string) {
	o.Message = v
}

func (o EntitlementManagementStatusValidationResponseOneOf1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitlementManagementStatusValidationResponseOneOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["valid"] = o.Valid
	toSerialize["reasonCode"] = o.ReasonCode
	toSerialize["message"] = o.Message

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EntitlementManagementStatusValidationResponseOneOf1) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"valid",
		"reasonCode",
		"message",
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

	varEntitlementManagementStatusValidationResponseOneOf1 := _EntitlementManagementStatusValidationResponseOneOf1{}

	err = json.Unmarshal(data, &varEntitlementManagementStatusValidationResponseOneOf1)

	if err != nil {
		return err
	}

	*o = EntitlementManagementStatusValidationResponseOneOf1(varEntitlementManagementStatusValidationResponseOneOf1)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "valid")
		delete(additionalProperties, "reasonCode")
		delete(additionalProperties, "message")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEntitlementManagementStatusValidationResponseOneOf1 struct {
	value *EntitlementManagementStatusValidationResponseOneOf1
	isSet bool
}

func (v NullableEntitlementManagementStatusValidationResponseOneOf1) Get() *EntitlementManagementStatusValidationResponseOneOf1 {
	return v.value
}

func (v *NullableEntitlementManagementStatusValidationResponseOneOf1) Set(val *EntitlementManagementStatusValidationResponseOneOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementManagementStatusValidationResponseOneOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementManagementStatusValidationResponseOneOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementManagementStatusValidationResponseOneOf1(val *EntitlementManagementStatusValidationResponseOneOf1) *NullableEntitlementManagementStatusValidationResponseOneOf1 {
	return &NullableEntitlementManagementStatusValidationResponseOneOf1{value: val, isSet: true}
}

func (v NullableEntitlementManagementStatusValidationResponseOneOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementManagementStatusValidationResponseOneOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
