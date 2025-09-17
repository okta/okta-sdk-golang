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

// checks if the DeviceSignalCollectionPolicyRuleDeviceSignalCollection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceSignalCollectionPolicyRuleDeviceSignalCollection{}

// DeviceSignalCollectionPolicyRuleDeviceSignalCollection Specifies how device context is collected when a user attempts to sign in
type DeviceSignalCollectionPolicyRuleDeviceSignalCollection struct {
	// Contains the device context provider configuration
	DeviceContextProviders []DeviceContextProvider `json:"deviceContextProviders,omitempty"`
	AdditionalProperties   map[string]interface{}
}

type _DeviceSignalCollectionPolicyRuleDeviceSignalCollection DeviceSignalCollectionPolicyRuleDeviceSignalCollection

// NewDeviceSignalCollectionPolicyRuleDeviceSignalCollection instantiates a new DeviceSignalCollectionPolicyRuleDeviceSignalCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceSignalCollectionPolicyRuleDeviceSignalCollection() *DeviceSignalCollectionPolicyRuleDeviceSignalCollection {
	this := DeviceSignalCollectionPolicyRuleDeviceSignalCollection{}
	return &this
}

// NewDeviceSignalCollectionPolicyRuleDeviceSignalCollectionWithDefaults instantiates a new DeviceSignalCollectionPolicyRuleDeviceSignalCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceSignalCollectionPolicyRuleDeviceSignalCollectionWithDefaults() *DeviceSignalCollectionPolicyRuleDeviceSignalCollection {
	this := DeviceSignalCollectionPolicyRuleDeviceSignalCollection{}
	return &this
}

// GetDeviceContextProviders returns the DeviceContextProviders field value if set, zero value otherwise.
func (o *DeviceSignalCollectionPolicyRuleDeviceSignalCollection) GetDeviceContextProviders() []DeviceContextProvider {
	if o == nil || IsNil(o.DeviceContextProviders) {
		var ret []DeviceContextProvider
		return ret
	}
	return o.DeviceContextProviders
}

// GetDeviceContextProvidersOk returns a tuple with the DeviceContextProviders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceSignalCollectionPolicyRuleDeviceSignalCollection) GetDeviceContextProvidersOk() ([]DeviceContextProvider, bool) {
	if o == nil || IsNil(o.DeviceContextProviders) {
		return nil, false
	}
	return o.DeviceContextProviders, true
}

// HasDeviceContextProviders returns a boolean if a field has been set.
func (o *DeviceSignalCollectionPolicyRuleDeviceSignalCollection) HasDeviceContextProviders() bool {
	if o != nil && !IsNil(o.DeviceContextProviders) {
		return true
	}

	return false
}

// SetDeviceContextProviders gets a reference to the given []DeviceContextProvider and assigns it to the DeviceContextProviders field.
func (o *DeviceSignalCollectionPolicyRuleDeviceSignalCollection) SetDeviceContextProviders(v []DeviceContextProvider) {
	o.DeviceContextProviders = v
}

func (o DeviceSignalCollectionPolicyRuleDeviceSignalCollection) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceSignalCollectionPolicyRuleDeviceSignalCollection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeviceContextProviders) {
		toSerialize["deviceContextProviders"] = o.DeviceContextProviders
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeviceSignalCollectionPolicyRuleDeviceSignalCollection) UnmarshalJSON(data []byte) (err error) {
	varDeviceSignalCollectionPolicyRuleDeviceSignalCollection := _DeviceSignalCollectionPolicyRuleDeviceSignalCollection{}

	err = json.Unmarshal(data, &varDeviceSignalCollectionPolicyRuleDeviceSignalCollection)

	if err != nil {
		return err
	}

	*o = DeviceSignalCollectionPolicyRuleDeviceSignalCollection(varDeviceSignalCollectionPolicyRuleDeviceSignalCollection)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "deviceContextProviders")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeviceSignalCollectionPolicyRuleDeviceSignalCollection struct {
	value *DeviceSignalCollectionPolicyRuleDeviceSignalCollection
	isSet bool
}

func (v NullableDeviceSignalCollectionPolicyRuleDeviceSignalCollection) Get() *DeviceSignalCollectionPolicyRuleDeviceSignalCollection {
	return v.value
}

func (v *NullableDeviceSignalCollectionPolicyRuleDeviceSignalCollection) Set(val *DeviceSignalCollectionPolicyRuleDeviceSignalCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceSignalCollectionPolicyRuleDeviceSignalCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceSignalCollectionPolicyRuleDeviceSignalCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceSignalCollectionPolicyRuleDeviceSignalCollection(val *DeviceSignalCollectionPolicyRuleDeviceSignalCollection) *NullableDeviceSignalCollectionPolicyRuleDeviceSignalCollection {
	return &NullableDeviceSignalCollectionPolicyRuleDeviceSignalCollection{value: val, isSet: true}
}

func (v NullableDeviceSignalCollectionPolicyRuleDeviceSignalCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceSignalCollectionPolicyRuleDeviceSignalCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
