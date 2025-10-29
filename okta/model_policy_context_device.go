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

// checks if the PolicyContextDevice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PolicyContextDevice{}

// PolicyContextDevice struct for PolicyContextDevice
type PolicyContextDevice struct {
	// The platform of the device, for example, IOS.
	Platform *string `json:"platform,omitempty"`
	// If the device is registered
	Registered *bool `json:"registered,omitempty"`
	// If the device is managed
	Managed *bool `json:"managed,omitempty"`
	// The device assurance policy ID for the simulation
	AssuranceId          *string `json:"assuranceId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PolicyContextDevice PolicyContextDevice

// NewPolicyContextDevice instantiates a new PolicyContextDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyContextDevice() *PolicyContextDevice {
	this := PolicyContextDevice{}
	return &this
}

// NewPolicyContextDeviceWithDefaults instantiates a new PolicyContextDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyContextDeviceWithDefaults() *PolicyContextDevice {
	this := PolicyContextDevice{}
	return &this
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *PolicyContextDevice) GetPlatform() string {
	if o == nil || IsNil(o.Platform) {
		var ret string
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyContextDevice) GetPlatformOk() (*string, bool) {
	if o == nil || IsNil(o.Platform) {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *PolicyContextDevice) HasPlatform() bool {
	if o != nil && !IsNil(o.Platform) {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given string and assigns it to the Platform field.
func (o *PolicyContextDevice) SetPlatform(v string) {
	o.Platform = &v
}

// GetRegistered returns the Registered field value if set, zero value otherwise.
func (o *PolicyContextDevice) GetRegistered() bool {
	if o == nil || IsNil(o.Registered) {
		var ret bool
		return ret
	}
	return *o.Registered
}

// GetRegisteredOk returns a tuple with the Registered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyContextDevice) GetRegisteredOk() (*bool, bool) {
	if o == nil || IsNil(o.Registered) {
		return nil, false
	}
	return o.Registered, true
}

// HasRegistered returns a boolean if a field has been set.
func (o *PolicyContextDevice) HasRegistered() bool {
	if o != nil && !IsNil(o.Registered) {
		return true
	}

	return false
}

// SetRegistered gets a reference to the given bool and assigns it to the Registered field.
func (o *PolicyContextDevice) SetRegistered(v bool) {
	o.Registered = &v
}

// GetManaged returns the Managed field value if set, zero value otherwise.
func (o *PolicyContextDevice) GetManaged() bool {
	if o == nil || IsNil(o.Managed) {
		var ret bool
		return ret
	}
	return *o.Managed
}

// GetManagedOk returns a tuple with the Managed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyContextDevice) GetManagedOk() (*bool, bool) {
	if o == nil || IsNil(o.Managed) {
		return nil, false
	}
	return o.Managed, true
}

// HasManaged returns a boolean if a field has been set.
func (o *PolicyContextDevice) HasManaged() bool {
	if o != nil && !IsNil(o.Managed) {
		return true
	}

	return false
}

// SetManaged gets a reference to the given bool and assigns it to the Managed field.
func (o *PolicyContextDevice) SetManaged(v bool) {
	o.Managed = &v
}

// GetAssuranceId returns the AssuranceId field value if set, zero value otherwise.
func (o *PolicyContextDevice) GetAssuranceId() string {
	if o == nil || IsNil(o.AssuranceId) {
		var ret string
		return ret
	}
	return *o.AssuranceId
}

// GetAssuranceIdOk returns a tuple with the AssuranceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyContextDevice) GetAssuranceIdOk() (*string, bool) {
	if o == nil || IsNil(o.AssuranceId) {
		return nil, false
	}
	return o.AssuranceId, true
}

// HasAssuranceId returns a boolean if a field has been set.
func (o *PolicyContextDevice) HasAssuranceId() bool {
	if o != nil && !IsNil(o.AssuranceId) {
		return true
	}

	return false
}

// SetAssuranceId gets a reference to the given string and assigns it to the AssuranceId field.
func (o *PolicyContextDevice) SetAssuranceId(v string) {
	o.AssuranceId = &v
}

func (o PolicyContextDevice) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PolicyContextDevice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Platform) {
		toSerialize["platform"] = o.Platform
	}
	if !IsNil(o.Registered) {
		toSerialize["registered"] = o.Registered
	}
	if !IsNil(o.Managed) {
		toSerialize["managed"] = o.Managed
	}
	if !IsNil(o.AssuranceId) {
		toSerialize["assuranceId"] = o.AssuranceId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PolicyContextDevice) UnmarshalJSON(data []byte) (err error) {
	varPolicyContextDevice := _PolicyContextDevice{}

	err = json.Unmarshal(data, &varPolicyContextDevice)

	if err != nil {
		return err
	}

	*o = PolicyContextDevice(varPolicyContextDevice)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "platform")
		delete(additionalProperties, "registered")
		delete(additionalProperties, "managed")
		delete(additionalProperties, "assuranceId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePolicyContextDevice struct {
	value *PolicyContextDevice
	isSet bool
}

func (v NullablePolicyContextDevice) Get() *PolicyContextDevice {
	return v.value
}

func (v *NullablePolicyContextDevice) Set(val *PolicyContextDevice) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyContextDevice) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyContextDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyContextDevice(val *PolicyContextDevice) *NullablePolicyContextDevice {
	return &NullablePolicyContextDevice{value: val, isSet: true}
}

func (v NullablePolicyContextDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyContextDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
