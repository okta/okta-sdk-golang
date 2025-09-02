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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// DevicePostureIdP <x-lifecycle-container><x-lifecycle class=\"ea\"></x-lifecycle></x-lifecycle-container>Device Posture IdP provider
type DevicePostureIdP struct {
	// Indicates whether the device is compliant according to the custom IDP
	Compliant *bool `json:"compliant,omitempty"`
	// Indicates whether the device is managed according to the custom IDP
	Managed *bool `json:"managed,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DevicePostureIdP DevicePostureIdP

// NewDevicePostureIdP instantiates a new DevicePostureIdP object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicePostureIdP() *DevicePostureIdP {
	this := DevicePostureIdP{}
	return &this
}

// NewDevicePostureIdPWithDefaults instantiates a new DevicePostureIdP object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicePostureIdPWithDefaults() *DevicePostureIdP {
	this := DevicePostureIdP{}
	return &this
}

// GetCompliant returns the Compliant field value if set, zero value otherwise.
func (o *DevicePostureIdP) GetCompliant() bool {
	if o == nil || o.Compliant == nil {
		var ret bool
		return ret
	}
	return *o.Compliant
}

// GetCompliantOk returns a tuple with the Compliant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicePostureIdP) GetCompliantOk() (*bool, bool) {
	if o == nil || o.Compliant == nil {
		return nil, false
	}
	return o.Compliant, true
}

// HasCompliant returns a boolean if a field has been set.
func (o *DevicePostureIdP) HasCompliant() bool {
	if o != nil && o.Compliant != nil {
		return true
	}

	return false
}

// SetCompliant gets a reference to the given bool and assigns it to the Compliant field.
func (o *DevicePostureIdP) SetCompliant(v bool) {
	o.Compliant = &v
}

// GetManaged returns the Managed field value if set, zero value otherwise.
func (o *DevicePostureIdP) GetManaged() bool {
	if o == nil || o.Managed == nil {
		var ret bool
		return ret
	}
	return *o.Managed
}

// GetManagedOk returns a tuple with the Managed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicePostureIdP) GetManagedOk() (*bool, bool) {
	if o == nil || o.Managed == nil {
		return nil, false
	}
	return o.Managed, true
}

// HasManaged returns a boolean if a field has been set.
func (o *DevicePostureIdP) HasManaged() bool {
	if o != nil && o.Managed != nil {
		return true
	}

	return false
}

// SetManaged gets a reference to the given bool and assigns it to the Managed field.
func (o *DevicePostureIdP) SetManaged(v bool) {
	o.Managed = &v
}

func (o DevicePostureIdP) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Compliant != nil {
		toSerialize["compliant"] = o.Compliant
	}
	if o.Managed != nil {
		toSerialize["managed"] = o.Managed
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DevicePostureIdP) UnmarshalJSON(bytes []byte) (err error) {
	varDevicePostureIdP := _DevicePostureIdP{}

	err = json.Unmarshal(bytes, &varDevicePostureIdP)
	if err == nil {
		*o = DevicePostureIdP(varDevicePostureIdP)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "compliant")
		delete(additionalProperties, "managed")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableDevicePostureIdP struct {
	value *DevicePostureIdP
	isSet bool
}

func (v NullableDevicePostureIdP) Get() *DevicePostureIdP {
	return v.value
}

func (v *NullableDevicePostureIdP) Set(val *DevicePostureIdP) {
	v.value = val
	v.isSet = true
}

func (v NullableDevicePostureIdP) IsSet() bool {
	return v.isSet
}

func (v *NullableDevicePostureIdP) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevicePostureIdP(val *DevicePostureIdP) *NullableDevicePostureIdP {
	return &NullableDevicePostureIdP{value: val, isSet: true}
}

func (v NullableDevicePostureIdP) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevicePostureIdP) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

