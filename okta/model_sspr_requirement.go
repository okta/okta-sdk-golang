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

// checks if the SsprRequirement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SsprRequirement{}

// SsprRequirement <x-lifecycle class=\"oie\"></x-lifecycle> Describes the initial and secondary authenticator requirements a user needs to reset their password
type SsprRequirement struct {
	// Determines which authentication requirements a user needs to perform self-service operations. `AUTH_POLICY` defers conditions and authentication requirements to the [Okta account management policy](https://developer.okta.com/docs/guides/okta-account-management-policy/main/). `LEGACY` refers to the requirements described by this rule.
	AccessControl        *string                 `json:"accessControl,omitempty"`
	Primary              *SsprPrimaryRequirement `json:"primary,omitempty"`
	StepUp               *SsprStepUpRequirement  `json:"stepUp,omitempty"`
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

// GetAccessControl returns the AccessControl field value if set, zero value otherwise.
func (o *SsprRequirement) GetAccessControl() string {
	if o == nil || IsNil(o.AccessControl) {
		var ret string
		return ret
	}
	return *o.AccessControl
}

// GetAccessControlOk returns a tuple with the AccessControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsprRequirement) GetAccessControlOk() (*string, bool) {
	if o == nil || IsNil(o.AccessControl) {
		return nil, false
	}
	return o.AccessControl, true
}

// HasAccessControl returns a boolean if a field has been set.
func (o *SsprRequirement) HasAccessControl() bool {
	if o != nil && !IsNil(o.AccessControl) {
		return true
	}

	return false
}

// SetAccessControl gets a reference to the given string and assigns it to the AccessControl field.
func (o *SsprRequirement) SetAccessControl(v string) {
	o.AccessControl = &v
}

// GetPrimary returns the Primary field value if set, zero value otherwise.
func (o *SsprRequirement) GetPrimary() SsprPrimaryRequirement {
	if o == nil || IsNil(o.Primary) {
		var ret SsprPrimaryRequirement
		return ret
	}
	return *o.Primary
}

// GetPrimaryOk returns a tuple with the Primary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsprRequirement) GetPrimaryOk() (*SsprPrimaryRequirement, bool) {
	if o == nil || IsNil(o.Primary) {
		return nil, false
	}
	return o.Primary, true
}

// HasPrimary returns a boolean if a field has been set.
func (o *SsprRequirement) HasPrimary() bool {
	if o != nil && !IsNil(o.Primary) {
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
	if o == nil || IsNil(o.StepUp) {
		var ret SsprStepUpRequirement
		return ret
	}
	return *o.StepUp
}

// GetStepUpOk returns a tuple with the StepUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsprRequirement) GetStepUpOk() (*SsprStepUpRequirement, bool) {
	if o == nil || IsNil(o.StepUp) {
		return nil, false
	}
	return o.StepUp, true
}

// HasStepUp returns a boolean if a field has been set.
func (o *SsprRequirement) HasStepUp() bool {
	if o != nil && !IsNil(o.StepUp) {
		return true
	}

	return false
}

// SetStepUp gets a reference to the given SsprStepUpRequirement and assigns it to the StepUp field.
func (o *SsprRequirement) SetStepUp(v SsprStepUpRequirement) {
	o.StepUp = &v
}

func (o SsprRequirement) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SsprRequirement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessControl) {
		toSerialize["accessControl"] = o.AccessControl
	}
	if !IsNil(o.Primary) {
		toSerialize["primary"] = o.Primary
	}
	if !IsNil(o.StepUp) {
		toSerialize["stepUp"] = o.StepUp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SsprRequirement) UnmarshalJSON(data []byte) (err error) {
	varSsprRequirement := _SsprRequirement{}

	err = json.Unmarshal(data, &varSsprRequirement)

	if err != nil {
		return err
	}

	*o = SsprRequirement(varSsprRequirement)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "accessControl")
		delete(additionalProperties, "primary")
		delete(additionalProperties, "stepUp")
		o.AdditionalProperties = additionalProperties
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
