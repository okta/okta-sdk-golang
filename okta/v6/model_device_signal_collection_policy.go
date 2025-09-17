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
	"fmt"
	"reflect"
	"strings"
)

// checks if the DeviceSignalCollectionPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceSignalCollectionPolicy{}

// DeviceSignalCollectionPolicy struct for DeviceSignalCollectionPolicy
type DeviceSignalCollectionPolicy struct {
	Policy
	// Policy conditions aren't supported. Conditions are applied at the rule level for this policy type.
	Conditions           NullableString             `json:"conditions,omitempty"`
	Embedded             *AccessPolicyAllOfEmbedded `json:"_embedded,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeviceSignalCollectionPolicy DeviceSignalCollectionPolicy

// NewDeviceSignalCollectionPolicy instantiates a new DeviceSignalCollectionPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceSignalCollectionPolicy(name string, type_ string) *DeviceSignalCollectionPolicy {
	this := DeviceSignalCollectionPolicy{}
	this.Name = name
	var system bool = false
	this.System = &system
	this.Type = type_
	return &this
}

// NewDeviceSignalCollectionPolicyWithDefaults instantiates a new DeviceSignalCollectionPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceSignalCollectionPolicyWithDefaults() *DeviceSignalCollectionPolicy {
	this := DeviceSignalCollectionPolicy{}
	return &this
}

// GetConditions returns the Conditions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceSignalCollectionPolicy) GetConditions() string {
	if o == nil || IsNil(o.Conditions.Get()) {
		var ret string
		return ret
	}
	return *o.Conditions.Get()
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceSignalCollectionPolicy) GetConditionsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Conditions.Get(), o.Conditions.IsSet()
}

// HasConditions returns a boolean if a field has been set.
func (o *DeviceSignalCollectionPolicy) HasConditions() bool {
	if o != nil && o.Conditions.IsSet() {
		return true
	}

	return false
}

// SetConditions gets a reference to the given NullableString and assigns it to the Conditions field.
func (o *DeviceSignalCollectionPolicy) SetConditions(v string) {
	o.Conditions.Set(&v)
}

// SetConditionsNil sets the value for Conditions to be an explicit nil
func (o *DeviceSignalCollectionPolicy) SetConditionsNil() {
	o.Conditions.Set(nil)
}

// UnsetConditions ensures that no value is present for Conditions, not even an explicit nil
func (o *DeviceSignalCollectionPolicy) UnsetConditions() {
	o.Conditions.Unset()
}

// GetEmbedded returns the Embedded field value if set, zero value otherwise.
func (o *DeviceSignalCollectionPolicy) GetEmbedded() AccessPolicyAllOfEmbedded {
	if o == nil || IsNil(o.Embedded) {
		var ret AccessPolicyAllOfEmbedded
		return ret
	}
	return *o.Embedded
}

// GetEmbeddedOk returns a tuple with the Embedded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceSignalCollectionPolicy) GetEmbeddedOk() (*AccessPolicyAllOfEmbedded, bool) {
	if o == nil || IsNil(o.Embedded) {
		return nil, false
	}
	return o.Embedded, true
}

// HasEmbedded returns a boolean if a field has been set.
func (o *DeviceSignalCollectionPolicy) HasEmbedded() bool {
	if o != nil && !IsNil(o.Embedded) {
		return true
	}

	return false
}

// SetEmbedded gets a reference to the given AccessPolicyAllOfEmbedded and assigns it to the Embedded field.
func (o *DeviceSignalCollectionPolicy) SetEmbedded(v AccessPolicyAllOfEmbedded) {
	o.Embedded = &v
}

func (o DeviceSignalCollectionPolicy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceSignalCollectionPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicy, errPolicy := json.Marshal(o.Policy)
	if errPolicy != nil {
		return map[string]interface{}{}, errPolicy
	}
	errPolicy = json.Unmarshal([]byte(serializedPolicy), &toSerialize)
	if errPolicy != nil {
		return map[string]interface{}{}, errPolicy
	}
	if o.Conditions.IsSet() {
		toSerialize["conditions"] = o.Conditions.Get()
	}
	if !IsNil(o.Embedded) {
		toSerialize["_embedded"] = o.Embedded
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeviceSignalCollectionPolicy) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"type",
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

	type DeviceSignalCollectionPolicyWithoutEmbeddedStruct struct {
		// Policy conditions aren't supported. Conditions are applied at the rule level for this policy type.
		Conditions NullableString             `json:"conditions,omitempty"`
		Embedded   *AccessPolicyAllOfEmbedded `json:"_embedded,omitempty"`
	}

	varDeviceSignalCollectionPolicyWithoutEmbeddedStruct := DeviceSignalCollectionPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varDeviceSignalCollectionPolicyWithoutEmbeddedStruct)
	if err == nil {
		varDeviceSignalCollectionPolicy := _DeviceSignalCollectionPolicy{}
		varDeviceSignalCollectionPolicy.Conditions = varDeviceSignalCollectionPolicyWithoutEmbeddedStruct.Conditions
		varDeviceSignalCollectionPolicy.Embedded = varDeviceSignalCollectionPolicyWithoutEmbeddedStruct.Embedded
		*o = DeviceSignalCollectionPolicy(varDeviceSignalCollectionPolicy)
	} else {
		return err
	}

	varDeviceSignalCollectionPolicy := _DeviceSignalCollectionPolicy{}

	err = json.Unmarshal(data, &varDeviceSignalCollectionPolicy)
	if err == nil {
		o.Policy = varDeviceSignalCollectionPolicy.Policy
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "conditions")
		delete(additionalProperties, "_embedded")

		// remove fields from embedded structs
		reflectPolicy := reflect.ValueOf(o.Policy)
		for i := 0; i < reflectPolicy.Type().NumField(); i++ {
			t := reflectPolicy.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeviceSignalCollectionPolicy struct {
	value *DeviceSignalCollectionPolicy
	isSet bool
}

func (v NullableDeviceSignalCollectionPolicy) Get() *DeviceSignalCollectionPolicy {
	return v.value
}

func (v *NullableDeviceSignalCollectionPolicy) Set(val *DeviceSignalCollectionPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceSignalCollectionPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceSignalCollectionPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceSignalCollectionPolicy(val *DeviceSignalCollectionPolicy) *NullableDeviceSignalCollectionPolicy {
	return &NullableDeviceSignalCollectionPolicy{value: val, isSet: true}
}

func (v NullableDeviceSignalCollectionPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceSignalCollectionPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
