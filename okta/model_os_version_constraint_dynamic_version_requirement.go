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

// checks if the OSVersionConstraintDynamicVersionRequirement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OSVersionConstraintDynamicVersionRequirement{}

// OSVersionConstraintDynamicVersionRequirement Contains the necessary properties for a dynamic Windows version requirement
type OSVersionConstraintDynamicVersionRequirement struct {
	// Indicates the type of the dynamic Windows version requirement
	Type *string `json:"type,omitempty"`
	// Indicates the distance from the latest Windows major version
	DistanceFromLatestMajor *int32 `json:"distanceFromLatestMajor,omitempty"`
	// Indicates whether the policy requires Windows devices to be on the latest security patch
	LatestSecurityPatch  *bool `json:"latestSecurityPatch,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OSVersionConstraintDynamicVersionRequirement OSVersionConstraintDynamicVersionRequirement

// NewOSVersionConstraintDynamicVersionRequirement instantiates a new OSVersionConstraintDynamicVersionRequirement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOSVersionConstraintDynamicVersionRequirement() *OSVersionConstraintDynamicVersionRequirement {
	this := OSVersionConstraintDynamicVersionRequirement{}
	return &this
}

// NewOSVersionConstraintDynamicVersionRequirementWithDefaults instantiates a new OSVersionConstraintDynamicVersionRequirement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOSVersionConstraintDynamicVersionRequirementWithDefaults() *OSVersionConstraintDynamicVersionRequirement {
	this := OSVersionConstraintDynamicVersionRequirement{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *OSVersionConstraintDynamicVersionRequirement) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSVersionConstraintDynamicVersionRequirement) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *OSVersionConstraintDynamicVersionRequirement) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *OSVersionConstraintDynamicVersionRequirement) SetType(v string) {
	o.Type = &v
}

// GetDistanceFromLatestMajor returns the DistanceFromLatestMajor field value if set, zero value otherwise.
func (o *OSVersionConstraintDynamicVersionRequirement) GetDistanceFromLatestMajor() int32 {
	if o == nil || IsNil(o.DistanceFromLatestMajor) {
		var ret int32
		return ret
	}
	return *o.DistanceFromLatestMajor
}

// GetDistanceFromLatestMajorOk returns a tuple with the DistanceFromLatestMajor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSVersionConstraintDynamicVersionRequirement) GetDistanceFromLatestMajorOk() (*int32, bool) {
	if o == nil || IsNil(o.DistanceFromLatestMajor) {
		return nil, false
	}
	return o.DistanceFromLatestMajor, true
}

// HasDistanceFromLatestMajor returns a boolean if a field has been set.
func (o *OSVersionConstraintDynamicVersionRequirement) HasDistanceFromLatestMajor() bool {
	if o != nil && !IsNil(o.DistanceFromLatestMajor) {
		return true
	}

	return false
}

// SetDistanceFromLatestMajor gets a reference to the given int32 and assigns it to the DistanceFromLatestMajor field.
func (o *OSVersionConstraintDynamicVersionRequirement) SetDistanceFromLatestMajor(v int32) {
	o.DistanceFromLatestMajor = &v
}

// GetLatestSecurityPatch returns the LatestSecurityPatch field value if set, zero value otherwise.
func (o *OSVersionConstraintDynamicVersionRequirement) GetLatestSecurityPatch() bool {
	if o == nil || IsNil(o.LatestSecurityPatch) {
		var ret bool
		return ret
	}
	return *o.LatestSecurityPatch
}

// GetLatestSecurityPatchOk returns a tuple with the LatestSecurityPatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSVersionConstraintDynamicVersionRequirement) GetLatestSecurityPatchOk() (*bool, bool) {
	if o == nil || IsNil(o.LatestSecurityPatch) {
		return nil, false
	}
	return o.LatestSecurityPatch, true
}

// HasLatestSecurityPatch returns a boolean if a field has been set.
func (o *OSVersionConstraintDynamicVersionRequirement) HasLatestSecurityPatch() bool {
	if o != nil && !IsNil(o.LatestSecurityPatch) {
		return true
	}

	return false
}

// SetLatestSecurityPatch gets a reference to the given bool and assigns it to the LatestSecurityPatch field.
func (o *OSVersionConstraintDynamicVersionRequirement) SetLatestSecurityPatch(v bool) {
	o.LatestSecurityPatch = &v
}

func (o OSVersionConstraintDynamicVersionRequirement) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OSVersionConstraintDynamicVersionRequirement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.DistanceFromLatestMajor) {
		toSerialize["distanceFromLatestMajor"] = o.DistanceFromLatestMajor
	}
	if !IsNil(o.LatestSecurityPatch) {
		toSerialize["latestSecurityPatch"] = o.LatestSecurityPatch
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OSVersionConstraintDynamicVersionRequirement) UnmarshalJSON(data []byte) (err error) {
	varOSVersionConstraintDynamicVersionRequirement := _OSVersionConstraintDynamicVersionRequirement{}

	err = json.Unmarshal(data, &varOSVersionConstraintDynamicVersionRequirement)

	if err != nil {
		return err
	}

	*o = OSVersionConstraintDynamicVersionRequirement(varOSVersionConstraintDynamicVersionRequirement)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "distanceFromLatestMajor")
		delete(additionalProperties, "latestSecurityPatch")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOSVersionConstraintDynamicVersionRequirement struct {
	value *OSVersionConstraintDynamicVersionRequirement
	isSet bool
}

func (v NullableOSVersionConstraintDynamicVersionRequirement) Get() *OSVersionConstraintDynamicVersionRequirement {
	return v.value
}

func (v *NullableOSVersionConstraintDynamicVersionRequirement) Set(val *OSVersionConstraintDynamicVersionRequirement) {
	v.value = val
	v.isSet = true
}

func (v NullableOSVersionConstraintDynamicVersionRequirement) IsSet() bool {
	return v.isSet
}

func (v *NullableOSVersionConstraintDynamicVersionRequirement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOSVersionConstraintDynamicVersionRequirement(val *OSVersionConstraintDynamicVersionRequirement) *NullableOSVersionConstraintDynamicVersionRequirement {
	return &NullableOSVersionConstraintDynamicVersionRequirement{value: val, isSet: true}
}

func (v NullableOSVersionConstraintDynamicVersionRequirement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOSVersionConstraintDynamicVersionRequirement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
