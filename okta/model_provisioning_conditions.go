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

// checks if the ProvisioningConditions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProvisioningConditions{}

// ProvisioningConditions Conditional behaviors for an IdP user during authentication
type ProvisioningConditions struct {
	Deprovisioned        *ProvisioningDeprovisionedCondition `json:"deprovisioned,omitempty"`
	Suspended            *ProvisioningSuspendedCondition     `json:"suspended,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProvisioningConditions ProvisioningConditions

// NewProvisioningConditions instantiates a new ProvisioningConditions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvisioningConditions() *ProvisioningConditions {
	this := ProvisioningConditions{}
	return &this
}

// NewProvisioningConditionsWithDefaults instantiates a new ProvisioningConditions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvisioningConditionsWithDefaults() *ProvisioningConditions {
	this := ProvisioningConditions{}
	return &this
}

// GetDeprovisioned returns the Deprovisioned field value if set, zero value otherwise.
func (o *ProvisioningConditions) GetDeprovisioned() ProvisioningDeprovisionedCondition {
	if o == nil || IsNil(o.Deprovisioned) {
		var ret ProvisioningDeprovisionedCondition
		return ret
	}
	return *o.Deprovisioned
}

// GetDeprovisionedOk returns a tuple with the Deprovisioned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningConditions) GetDeprovisionedOk() (*ProvisioningDeprovisionedCondition, bool) {
	if o == nil || IsNil(o.Deprovisioned) {
		return nil, false
	}
	return o.Deprovisioned, true
}

// HasDeprovisioned returns a boolean if a field has been set.
func (o *ProvisioningConditions) HasDeprovisioned() bool {
	if o != nil && !IsNil(o.Deprovisioned) {
		return true
	}

	return false
}

// SetDeprovisioned gets a reference to the given ProvisioningDeprovisionedCondition and assigns it to the Deprovisioned field.
func (o *ProvisioningConditions) SetDeprovisioned(v ProvisioningDeprovisionedCondition) {
	o.Deprovisioned = &v
}

// GetSuspended returns the Suspended field value if set, zero value otherwise.
func (o *ProvisioningConditions) GetSuspended() ProvisioningSuspendedCondition {
	if o == nil || IsNil(o.Suspended) {
		var ret ProvisioningSuspendedCondition
		return ret
	}
	return *o.Suspended
}

// GetSuspendedOk returns a tuple with the Suspended field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningConditions) GetSuspendedOk() (*ProvisioningSuspendedCondition, bool) {
	if o == nil || IsNil(o.Suspended) {
		return nil, false
	}
	return o.Suspended, true
}

// HasSuspended returns a boolean if a field has been set.
func (o *ProvisioningConditions) HasSuspended() bool {
	if o != nil && !IsNil(o.Suspended) {
		return true
	}

	return false
}

// SetSuspended gets a reference to the given ProvisioningSuspendedCondition and assigns it to the Suspended field.
func (o *ProvisioningConditions) SetSuspended(v ProvisioningSuspendedCondition) {
	o.Suspended = &v
}

func (o ProvisioningConditions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProvisioningConditions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Deprovisioned) {
		toSerialize["deprovisioned"] = o.Deprovisioned
	}
	if !IsNil(o.Suspended) {
		toSerialize["suspended"] = o.Suspended
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProvisioningConditions) UnmarshalJSON(data []byte) (err error) {
	varProvisioningConditions := _ProvisioningConditions{}

	err = json.Unmarshal(data, &varProvisioningConditions)

	if err != nil {
		return err
	}

	*o = ProvisioningConditions(varProvisioningConditions)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "deprovisioned")
		delete(additionalProperties, "suspended")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProvisioningConditions struct {
	value *ProvisioningConditions
	isSet bool
}

func (v NullableProvisioningConditions) Get() *ProvisioningConditions {
	return v.value
}

func (v *NullableProvisioningConditions) Set(val *ProvisioningConditions) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisioningConditions) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisioningConditions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisioningConditions(val *ProvisioningConditions) *NullableProvisioningConditions {
	return &NullableProvisioningConditions{value: val, isSet: true}
}

func (v NullableProvisioningConditions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisioningConditions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
