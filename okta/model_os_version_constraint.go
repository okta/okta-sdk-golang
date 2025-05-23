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

// OSVersionConstraint struct for OSVersionConstraint
type OSVersionConstraint struct {
	DynamicVersionRequirement *OSVersionConstraintDynamicVersionRequirement `json:"dynamicVersionRequirement,omitempty"`
	// Indicates the Windows major version
	MajorVersionConstraint string `json:"majorVersionConstraint"`
	// The Windows device version must be equal to or newer than the specified version
	Minimum *string `json:"minimum,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OSVersionConstraint OSVersionConstraint

// NewOSVersionConstraint instantiates a new OSVersionConstraint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOSVersionConstraint(majorVersionConstraint string) *OSVersionConstraint {
	this := OSVersionConstraint{}
	this.MajorVersionConstraint = majorVersionConstraint
	return &this
}

// NewOSVersionConstraintWithDefaults instantiates a new OSVersionConstraint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOSVersionConstraintWithDefaults() *OSVersionConstraint {
	this := OSVersionConstraint{}
	return &this
}

// GetDynamicVersionRequirement returns the DynamicVersionRequirement field value if set, zero value otherwise.
func (o *OSVersionConstraint) GetDynamicVersionRequirement() OSVersionConstraintDynamicVersionRequirement {
	if o == nil || o.DynamicVersionRequirement == nil {
		var ret OSVersionConstraintDynamicVersionRequirement
		return ret
	}
	return *o.DynamicVersionRequirement
}

// GetDynamicVersionRequirementOk returns a tuple with the DynamicVersionRequirement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSVersionConstraint) GetDynamicVersionRequirementOk() (*OSVersionConstraintDynamicVersionRequirement, bool) {
	if o == nil || o.DynamicVersionRequirement == nil {
		return nil, false
	}
	return o.DynamicVersionRequirement, true
}

// HasDynamicVersionRequirement returns a boolean if a field has been set.
func (o *OSVersionConstraint) HasDynamicVersionRequirement() bool {
	if o != nil && o.DynamicVersionRequirement != nil {
		return true
	}

	return false
}

// SetDynamicVersionRequirement gets a reference to the given OSVersionConstraintDynamicVersionRequirement and assigns it to the DynamicVersionRequirement field.
func (o *OSVersionConstraint) SetDynamicVersionRequirement(v OSVersionConstraintDynamicVersionRequirement) {
	o.DynamicVersionRequirement = &v
}

// GetMajorVersionConstraint returns the MajorVersionConstraint field value
func (o *OSVersionConstraint) GetMajorVersionConstraint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MajorVersionConstraint
}

// GetMajorVersionConstraintOk returns a tuple with the MajorVersionConstraint field value
// and a boolean to check if the value has been set.
func (o *OSVersionConstraint) GetMajorVersionConstraintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MajorVersionConstraint, true
}

// SetMajorVersionConstraint sets field value
func (o *OSVersionConstraint) SetMajorVersionConstraint(v string) {
	o.MajorVersionConstraint = v
}

// GetMinimum returns the Minimum field value if set, zero value otherwise.
func (o *OSVersionConstraint) GetMinimum() string {
	if o == nil || o.Minimum == nil {
		var ret string
		return ret
	}
	return *o.Minimum
}

// GetMinimumOk returns a tuple with the Minimum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSVersionConstraint) GetMinimumOk() (*string, bool) {
	if o == nil || o.Minimum == nil {
		return nil, false
	}
	return o.Minimum, true
}

// HasMinimum returns a boolean if a field has been set.
func (o *OSVersionConstraint) HasMinimum() bool {
	if o != nil && o.Minimum != nil {
		return true
	}

	return false
}

// SetMinimum gets a reference to the given string and assigns it to the Minimum field.
func (o *OSVersionConstraint) SetMinimum(v string) {
	o.Minimum = &v
}

func (o OSVersionConstraint) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DynamicVersionRequirement != nil {
		toSerialize["dynamicVersionRequirement"] = o.DynamicVersionRequirement
	}
	if true {
		toSerialize["majorVersionConstraint"] = o.MajorVersionConstraint
	}
	if o.Minimum != nil {
		toSerialize["minimum"] = o.Minimum
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OSVersionConstraint) UnmarshalJSON(bytes []byte) (err error) {
	varOSVersionConstraint := _OSVersionConstraint{}

	err = json.Unmarshal(bytes, &varOSVersionConstraint)
	if err == nil {
		*o = OSVersionConstraint(varOSVersionConstraint)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "dynamicVersionRequirement")
		delete(additionalProperties, "majorVersionConstraint")
		delete(additionalProperties, "minimum")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableOSVersionConstraint struct {
	value *OSVersionConstraint
	isSet bool
}

func (v NullableOSVersionConstraint) Get() *OSVersionConstraint {
	return v.value
}

func (v *NullableOSVersionConstraint) Set(val *OSVersionConstraint) {
	v.value = val
	v.isSet = true
}

func (v NullableOSVersionConstraint) IsSet() bool {
	return v.isSet
}

func (v *NullableOSVersionConstraint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOSVersionConstraint(val *OSVersionConstraint) *NullableOSVersionConstraint {
	return &NullableOSVersionConstraint{value: val, isSet: true}
}

func (v NullableOSVersionConstraint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOSVersionConstraint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

