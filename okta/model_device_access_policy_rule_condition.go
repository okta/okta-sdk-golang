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

API version: 2025.10.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the DeviceAccessPolicyRuleCondition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceAccessPolicyRuleCondition{}

// DeviceAccessPolicyRuleCondition <x-lifecycle class=\"oie\"></x-lifecycle> Specifies the device condition to match on
type DeviceAccessPolicyRuleCondition struct {
	Assurance *DevicePolicyRuleConditionAssurance `json:"assurance,omitempty"`
	// Indicates if the device is managed. A device is considered managed if it's part of a device management system.
	Managed *bool `json:"managed,omitempty"`
	// Indicates if the device is registered. A device is registered if the User enrolls with Okta Verify that's installed on the device. When the `managed` property is passed, you must also include the `registered` property and set it to `true`.
	Registered           *bool `json:"registered,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeviceAccessPolicyRuleCondition DeviceAccessPolicyRuleCondition

// NewDeviceAccessPolicyRuleCondition instantiates a new DeviceAccessPolicyRuleCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceAccessPolicyRuleCondition() *DeviceAccessPolicyRuleCondition {
	this := DeviceAccessPolicyRuleCondition{}
	return &this
}

// NewDeviceAccessPolicyRuleConditionWithDefaults instantiates a new DeviceAccessPolicyRuleCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceAccessPolicyRuleConditionWithDefaults() *DeviceAccessPolicyRuleCondition {
	this := DeviceAccessPolicyRuleCondition{}
	return &this
}

// GetAssurance returns the Assurance field value if set, zero value otherwise.
func (o *DeviceAccessPolicyRuleCondition) GetAssurance() DevicePolicyRuleConditionAssurance {
	if o == nil || IsNil(o.Assurance) {
		var ret DevicePolicyRuleConditionAssurance
		return ret
	}
	return *o.Assurance
}

// GetAssuranceOk returns a tuple with the Assurance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAccessPolicyRuleCondition) GetAssuranceOk() (*DevicePolicyRuleConditionAssurance, bool) {
	if o == nil || IsNil(o.Assurance) {
		return nil, false
	}
	return o.Assurance, true
}

// HasAssurance returns a boolean if a field has been set.
func (o *DeviceAccessPolicyRuleCondition) HasAssurance() bool {
	if o != nil && !IsNil(o.Assurance) {
		return true
	}

	return false
}

// SetAssurance gets a reference to the given DevicePolicyRuleConditionAssurance and assigns it to the Assurance field.
func (o *DeviceAccessPolicyRuleCondition) SetAssurance(v DevicePolicyRuleConditionAssurance) {
	o.Assurance = &v
}

// GetManaged returns the Managed field value if set, zero value otherwise.
func (o *DeviceAccessPolicyRuleCondition) GetManaged() bool {
	if o == nil || IsNil(o.Managed) {
		var ret bool
		return ret
	}
	return *o.Managed
}

// GetManagedOk returns a tuple with the Managed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAccessPolicyRuleCondition) GetManagedOk() (*bool, bool) {
	if o == nil || IsNil(o.Managed) {
		return nil, false
	}
	return o.Managed, true
}

// HasManaged returns a boolean if a field has been set.
func (o *DeviceAccessPolicyRuleCondition) HasManaged() bool {
	if o != nil && !IsNil(o.Managed) {
		return true
	}

	return false
}

// SetManaged gets a reference to the given bool and assigns it to the Managed field.
func (o *DeviceAccessPolicyRuleCondition) SetManaged(v bool) {
	o.Managed = &v
}

// GetRegistered returns the Registered field value if set, zero value otherwise.
func (o *DeviceAccessPolicyRuleCondition) GetRegistered() bool {
	if o == nil || IsNil(o.Registered) {
		var ret bool
		return ret
	}
	return *o.Registered
}

// GetRegisteredOk returns a tuple with the Registered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAccessPolicyRuleCondition) GetRegisteredOk() (*bool, bool) {
	if o == nil || IsNil(o.Registered) {
		return nil, false
	}
	return o.Registered, true
}

// HasRegistered returns a boolean if a field has been set.
func (o *DeviceAccessPolicyRuleCondition) HasRegistered() bool {
	if o != nil && !IsNil(o.Registered) {
		return true
	}

	return false
}

// SetRegistered gets a reference to the given bool and assigns it to the Registered field.
func (o *DeviceAccessPolicyRuleCondition) SetRegistered(v bool) {
	o.Registered = &v
}

func (o DeviceAccessPolicyRuleCondition) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceAccessPolicyRuleCondition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Assurance) {
		toSerialize["assurance"] = o.Assurance
	}
	if !IsNil(o.Managed) {
		toSerialize["managed"] = o.Managed
	}
	if !IsNil(o.Registered) {
		toSerialize["registered"] = o.Registered
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeviceAccessPolicyRuleCondition) UnmarshalJSON(data []byte) (err error) {
	varDeviceAccessPolicyRuleCondition := _DeviceAccessPolicyRuleCondition{}

	err = json.Unmarshal(data, &varDeviceAccessPolicyRuleCondition)

	if err != nil {
		return err
	}

	*o = DeviceAccessPolicyRuleCondition(varDeviceAccessPolicyRuleCondition)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "assurance")
		delete(additionalProperties, "managed")
		delete(additionalProperties, "registered")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeviceAccessPolicyRuleCondition struct {
	value *DeviceAccessPolicyRuleCondition
	isSet bool
}

func (v NullableDeviceAccessPolicyRuleCondition) Get() *DeviceAccessPolicyRuleCondition {
	return v.value
}

func (v *NullableDeviceAccessPolicyRuleCondition) Set(val *DeviceAccessPolicyRuleCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceAccessPolicyRuleCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceAccessPolicyRuleCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceAccessPolicyRuleCondition(val *DeviceAccessPolicyRuleCondition) *NullableDeviceAccessPolicyRuleCondition {
	return &NullableDeviceAccessPolicyRuleCondition{value: val, isSet: true}
}

func (v NullableDeviceAccessPolicyRuleCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceAccessPolicyRuleCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
