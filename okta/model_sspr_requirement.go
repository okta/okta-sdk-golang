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

API version: 2024.06.1
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// SsprRequirement Describes the initial and secondary authenticator requirements a user needs to reset their password
type SsprRequirement struct {
	Primary *SsprPrimaryRequirement `json:"primary,omitempty"`
	StepUp *SsprStepUpRequirement `json:"stepUp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SsprRequirement SsprRequirement

// NewSsprRequirement instantiates a new SsprRequirement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSsprRequirement() *SsprRequirement {
	this := SsprRequirement{}
	return &this
}

// NewSsprRequirementWithDefaults instantiates a new SsprRequirement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSsprRequirementWithDefaults() *SsprRequirement {
	this := SsprRequirement{}
	return &this
}

// GetPrimary returns the Primary field value if set, zero value otherwise.
func (o *SsprRequirement) GetPrimary() SsprPrimaryRequirement {
	if o == nil || o.Primary == nil {
		var ret SsprPrimaryRequirement
		return ret
	}
	return *o.Primary
}

// GetPrimaryOk returns a tuple with the Primary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsprRequirement) GetPrimaryOk() (*SsprPrimaryRequirement, bool) {
	if o == nil || o.Primary == nil {
		return nil, false
	}
	return o.Primary, true
}

// HasPrimary returns a boolean if a field has been set.
func (o *SsprRequirement) HasPrimary() bool {
	if o != nil && o.Primary != nil {
		return true
	}

	return false
}

// SetPrimary gets a reference to the given SsprPrimaryRequirement and assigns it to the Primary field.
func (o *SsprRequirement) SetPrimary(v SsprPrimaryRequirement) {
	o.Primary = &v
}

// GetStepUp returns the StepUp field value if set, zero value otherwise.
func (o *SsprRequirement) GetStepUp() SsprStepUpRequirement {
	if o == nil || o.StepUp == nil {
		var ret SsprStepUpRequirement
		return ret
	}
	return *o.StepUp
}

// GetStepUpOk returns a tuple with the StepUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsprRequirement) GetStepUpOk() (*SsprStepUpRequirement, bool) {
	if o == nil || o.StepUp == nil {
		return nil, false
	}
	return o.StepUp, true
}

// HasStepUp returns a boolean if a field has been set.
func (o *SsprRequirement) HasStepUp() bool {
	if o != nil && o.StepUp != nil {
		return true
	}

	return false
}

// SetStepUp gets a reference to the given SsprStepUpRequirement and assigns it to the StepUp field.
func (o *SsprRequirement) SetStepUp(v SsprStepUpRequirement) {
	o.StepUp = &v
}

func (o SsprRequirement) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Primary != nil {
		toSerialize["primary"] = o.Primary
	}
	if o.StepUp != nil {
		toSerialize["stepUp"] = o.StepUp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SsprRequirement) UnmarshalJSON(bytes []byte) (err error) {
	varSsprRequirement := _SsprRequirement{}

	err = json.Unmarshal(bytes, &varSsprRequirement)
	if err == nil {
		*o = SsprRequirement(varSsprRequirement)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "primary")
		delete(additionalProperties, "stepUp")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableSsprRequirement struct {
	value *SsprRequirement
	isSet bool
}

func (v NullableSsprRequirement) Get() *SsprRequirement {
	return v.value
}

func (v *NullableSsprRequirement) Set(val *SsprRequirement) {
	v.value = val
	v.isSet = true
}

func (v NullableSsprRequirement) IsSet() bool {
	return v.isSet
}

func (v *NullableSsprRequirement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSsprRequirement(val *SsprRequirement) *NullableSsprRequirement {
	return &NullableSsprRequirement{value: val, isSet: true}
}

func (v NullableSsprRequirement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSsprRequirement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

