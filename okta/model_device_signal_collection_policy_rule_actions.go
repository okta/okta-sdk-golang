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

API version: 2025.08.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the DeviceSignalCollectionPolicyRuleActions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceSignalCollectionPolicyRuleActions{}

// DeviceSignalCollectionPolicyRuleActions <x-lifecycle-container><x-lifecycle class=\"ea\"></x-lifecycle></x-lifecycle-container>Specifies actions to be taken, or operations that may be allowed, if the rule conditions are satisfied
type DeviceSignalCollectionPolicyRuleActions struct {
	DeviceSignalCollection *DeviceSignalCollectionPolicyRuleDeviceSignalCollection `json:"deviceSignalCollection,omitempty"`
	AdditionalProperties   map[string]interface{}
}

type _DeviceSignalCollectionPolicyRuleActions DeviceSignalCollectionPolicyRuleActions

// NewDeviceSignalCollectionPolicyRuleActions instantiates a new DeviceSignalCollectionPolicyRuleActions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceSignalCollectionPolicyRuleActions() *DeviceSignalCollectionPolicyRuleActions {
	this := DeviceSignalCollectionPolicyRuleActions{}
	return &this
}

// NewDeviceSignalCollectionPolicyRuleActionsWithDefaults instantiates a new DeviceSignalCollectionPolicyRuleActions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceSignalCollectionPolicyRuleActionsWithDefaults() *DeviceSignalCollectionPolicyRuleActions {
	this := DeviceSignalCollectionPolicyRuleActions{}
	return &this
}

// GetDeviceSignalCollection returns the DeviceSignalCollection field value if set, zero value otherwise.
func (o *DeviceSignalCollectionPolicyRuleActions) GetDeviceSignalCollection() DeviceSignalCollectionPolicyRuleDeviceSignalCollection {
	if o == nil || IsNil(o.DeviceSignalCollection) {
		var ret DeviceSignalCollectionPolicyRuleDeviceSignalCollection
		return ret
	}
	return *o.DeviceSignalCollection
}

// GetDeviceSignalCollectionOk returns a tuple with the DeviceSignalCollection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceSignalCollectionPolicyRuleActions) GetDeviceSignalCollectionOk() (*DeviceSignalCollectionPolicyRuleDeviceSignalCollection, bool) {
	if o == nil || IsNil(o.DeviceSignalCollection) {
		return nil, false
	}
	return o.DeviceSignalCollection, true
}

// HasDeviceSignalCollection returns a boolean if a field has been set.
func (o *DeviceSignalCollectionPolicyRuleActions) HasDeviceSignalCollection() bool {
	if o != nil && !IsNil(o.DeviceSignalCollection) {
		return true
	}

	return false
}

// SetDeviceSignalCollection gets a reference to the given DeviceSignalCollectionPolicyRuleDeviceSignalCollection and assigns it to the DeviceSignalCollection field.
func (o *DeviceSignalCollectionPolicyRuleActions) SetDeviceSignalCollection(v DeviceSignalCollectionPolicyRuleDeviceSignalCollection) {
	o.DeviceSignalCollection = &v
}

func (o DeviceSignalCollectionPolicyRuleActions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceSignalCollectionPolicyRuleActions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeviceSignalCollection) {
		toSerialize["deviceSignalCollection"] = o.DeviceSignalCollection
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeviceSignalCollectionPolicyRuleActions) UnmarshalJSON(data []byte) (err error) {
	varDeviceSignalCollectionPolicyRuleActions := _DeviceSignalCollectionPolicyRuleActions{}

	err = json.Unmarshal(data, &varDeviceSignalCollectionPolicyRuleActions)

	if err != nil {
		return err
	}

	*o = DeviceSignalCollectionPolicyRuleActions(varDeviceSignalCollectionPolicyRuleActions)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "deviceSignalCollection")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeviceSignalCollectionPolicyRuleActions struct {
	value *DeviceSignalCollectionPolicyRuleActions
	isSet bool
}

func (v NullableDeviceSignalCollectionPolicyRuleActions) Get() *DeviceSignalCollectionPolicyRuleActions {
	return v.value
}

func (v *NullableDeviceSignalCollectionPolicyRuleActions) Set(val *DeviceSignalCollectionPolicyRuleActions) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceSignalCollectionPolicyRuleActions) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceSignalCollectionPolicyRuleActions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceSignalCollectionPolicyRuleActions(val *DeviceSignalCollectionPolicyRuleActions) *NullableDeviceSignalCollectionPolicyRuleActions {
	return &NullableDeviceSignalCollectionPolicyRuleActions{value: val, isSet: true}
}

func (v NullableDeviceSignalCollectionPolicyRuleActions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceSignalCollectionPolicyRuleActions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
