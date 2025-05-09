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

// OSVersionDynamicVersionRequirement <div class=\"x-lifecycle-container\"><x-lifecycle class=\"ea\"></x-lifecycle></div>Contains the necessary properties for a dynamic version requirement
type OSVersionDynamicVersionRequirement struct {
	// Indicates the type of the dynamic OS version requirement
	Type *string `json:"type,omitempty"`
	// Indicates the distance from the latest major version
	DistanceFromLatestMajor *int32 `json:"distanceFromLatestMajor,omitempty"`
	// Indicates whether the device needs to be on the latest security patch
	LatestSecurityPatch *bool `json:"latestSecurityPatch,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OSVersionDynamicVersionRequirement OSVersionDynamicVersionRequirement

// NewOSVersionDynamicVersionRequirement instantiates a new OSVersionDynamicVersionRequirement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOSVersionDynamicVersionRequirement() *OSVersionDynamicVersionRequirement {
	this := OSVersionDynamicVersionRequirement{}
	return &this
}

// NewOSVersionDynamicVersionRequirementWithDefaults instantiates a new OSVersionDynamicVersionRequirement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOSVersionDynamicVersionRequirementWithDefaults() *OSVersionDynamicVersionRequirement {
	this := OSVersionDynamicVersionRequirement{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *OSVersionDynamicVersionRequirement) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSVersionDynamicVersionRequirement) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *OSVersionDynamicVersionRequirement) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *OSVersionDynamicVersionRequirement) SetType(v string) {
	o.Type = &v
}

// GetDistanceFromLatestMajor returns the DistanceFromLatestMajor field value if set, zero value otherwise.
func (o *OSVersionDynamicVersionRequirement) GetDistanceFromLatestMajor() int32 {
	if o == nil || o.DistanceFromLatestMajor == nil {
		var ret int32
		return ret
	}
	return *o.DistanceFromLatestMajor
}

// GetDistanceFromLatestMajorOk returns a tuple with the DistanceFromLatestMajor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSVersionDynamicVersionRequirement) GetDistanceFromLatestMajorOk() (*int32, bool) {
	if o == nil || o.DistanceFromLatestMajor == nil {
		return nil, false
	}
	return o.DistanceFromLatestMajor, true
}

// HasDistanceFromLatestMajor returns a boolean if a field has been set.
func (o *OSVersionDynamicVersionRequirement) HasDistanceFromLatestMajor() bool {
	if o != nil && o.DistanceFromLatestMajor != nil {
		return true
	}

	return false
}

// SetDistanceFromLatestMajor gets a reference to the given int32 and assigns it to the DistanceFromLatestMajor field.
func (o *OSVersionDynamicVersionRequirement) SetDistanceFromLatestMajor(v int32) {
	o.DistanceFromLatestMajor = &v
}

// GetLatestSecurityPatch returns the LatestSecurityPatch field value if set, zero value otherwise.
func (o *OSVersionDynamicVersionRequirement) GetLatestSecurityPatch() bool {
	if o == nil || o.LatestSecurityPatch == nil {
		var ret bool
		return ret
	}
	return *o.LatestSecurityPatch
}

// GetLatestSecurityPatchOk returns a tuple with the LatestSecurityPatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSVersionDynamicVersionRequirement) GetLatestSecurityPatchOk() (*bool, bool) {
	if o == nil || o.LatestSecurityPatch == nil {
		return nil, false
	}
	return o.LatestSecurityPatch, true
}

// HasLatestSecurityPatch returns a boolean if a field has been set.
func (o *OSVersionDynamicVersionRequirement) HasLatestSecurityPatch() bool {
	if o != nil && o.LatestSecurityPatch != nil {
		return true
	}

	return false
}

// SetLatestSecurityPatch gets a reference to the given bool and assigns it to the LatestSecurityPatch field.
func (o *OSVersionDynamicVersionRequirement) SetLatestSecurityPatch(v bool) {
	o.LatestSecurityPatch = &v
}

func (o OSVersionDynamicVersionRequirement) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.DistanceFromLatestMajor != nil {
		toSerialize["distanceFromLatestMajor"] = o.DistanceFromLatestMajor
	}
	if o.LatestSecurityPatch != nil {
		toSerialize["latestSecurityPatch"] = o.LatestSecurityPatch
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OSVersionDynamicVersionRequirement) UnmarshalJSON(bytes []byte) (err error) {
	varOSVersionDynamicVersionRequirement := _OSVersionDynamicVersionRequirement{}

	err = json.Unmarshal(bytes, &varOSVersionDynamicVersionRequirement)
	if err == nil {
		*o = OSVersionDynamicVersionRequirement(varOSVersionDynamicVersionRequirement)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "distanceFromLatestMajor")
		delete(additionalProperties, "latestSecurityPatch")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableOSVersionDynamicVersionRequirement struct {
	value *OSVersionDynamicVersionRequirement
	isSet bool
}

func (v NullableOSVersionDynamicVersionRequirement) Get() *OSVersionDynamicVersionRequirement {
	return v.value
}

func (v *NullableOSVersionDynamicVersionRequirement) Set(val *OSVersionDynamicVersionRequirement) {
	v.value = val
	v.isSet = true
}

func (v NullableOSVersionDynamicVersionRequirement) IsSet() bool {
	return v.isSet
}

func (v *NullableOSVersionDynamicVersionRequirement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOSVersionDynamicVersionRequirement(val *OSVersionDynamicVersionRequirement) *NullableOSVersionDynamicVersionRequirement {
	return &NullableOSVersionDynamicVersionRequirement{value: val, isSet: true}
}

func (v NullableOSVersionDynamicVersionRequirement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOSVersionDynamicVersionRequirement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

