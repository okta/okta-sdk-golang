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

// OSVersion Specifies the OS requirement for the policy.  There are two types of OS requirements:  * **Static**: A specific OS version requirement that doesn't change until you update the policy. A static OS requirement is specified with the `osVersion.minimum` property. * **Dynamic**: An OS version requirement that is relative to the latest major OS release and security patch. A dynamic OS requirement is specified with the `osVersion.dynamicVersionRequirement` property. > **Note:** Dynamic OS requirements are available only if the **Dynamic OS version compliance** [self-service EA](/openapi/okta-management/guides/release-lifecycle/#early-access-ea) feature is enabled. You can't specify both `osVersion.minimum` and `osVersion.dynamicVersionRequirement` properties at the same time. 
type OSVersion struct {
	DynamicVersionRequirement *OSVersionDynamicVersionRequirement `json:"dynamicVersionRequirement,omitempty"`
	// The device version must be equal to or newer than the specified version string (maximum of three components for iOS and macOS, and maximum of four components for Android)
	Minimum *string `json:"minimum,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OSVersion OSVersion

// NewOSVersion instantiates a new OSVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOSVersion() *OSVersion {
	this := OSVersion{}
	return &this
}

// NewOSVersionWithDefaults instantiates a new OSVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOSVersionWithDefaults() *OSVersion {
	this := OSVersion{}
	return &this
}

// GetDynamicVersionRequirement returns the DynamicVersionRequirement field value if set, zero value otherwise.
func (o *OSVersion) GetDynamicVersionRequirement() OSVersionDynamicVersionRequirement {
	if o == nil || o.DynamicVersionRequirement == nil {
		var ret OSVersionDynamicVersionRequirement
		return ret
	}
	return *o.DynamicVersionRequirement
}

// GetDynamicVersionRequirementOk returns a tuple with the DynamicVersionRequirement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSVersion) GetDynamicVersionRequirementOk() (*OSVersionDynamicVersionRequirement, bool) {
	if o == nil || o.DynamicVersionRequirement == nil {
		return nil, false
	}
	return o.DynamicVersionRequirement, true
}

// HasDynamicVersionRequirement returns a boolean if a field has been set.
func (o *OSVersion) HasDynamicVersionRequirement() bool {
	if o != nil && o.DynamicVersionRequirement != nil {
		return true
	}

	return false
}

// SetDynamicVersionRequirement gets a reference to the given OSVersionDynamicVersionRequirement and assigns it to the DynamicVersionRequirement field.
func (o *OSVersion) SetDynamicVersionRequirement(v OSVersionDynamicVersionRequirement) {
	o.DynamicVersionRequirement = &v
}

// GetMinimum returns the Minimum field value if set, zero value otherwise.
func (o *OSVersion) GetMinimum() string {
	if o == nil || o.Minimum == nil {
		var ret string
		return ret
	}
	return *o.Minimum
}

// GetMinimumOk returns a tuple with the Minimum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSVersion) GetMinimumOk() (*string, bool) {
	if o == nil || o.Minimum == nil {
		return nil, false
	}
	return o.Minimum, true
}

// HasMinimum returns a boolean if a field has been set.
func (o *OSVersion) HasMinimum() bool {
	if o != nil && o.Minimum != nil {
		return true
	}

	return false
}

// SetMinimum gets a reference to the given string and assigns it to the Minimum field.
func (o *OSVersion) SetMinimum(v string) {
	o.Minimum = &v
}

func (o OSVersion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DynamicVersionRequirement != nil {
		toSerialize["dynamicVersionRequirement"] = o.DynamicVersionRequirement
	}
	if o.Minimum != nil {
		toSerialize["minimum"] = o.Minimum
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OSVersion) UnmarshalJSON(bytes []byte) (err error) {
	varOSVersion := _OSVersion{}

	err = json.Unmarshal(bytes, &varOSVersion)
	if err == nil {
		*o = OSVersion(varOSVersion)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "dynamicVersionRequirement")
		delete(additionalProperties, "minimum")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableOSVersion struct {
	value *OSVersion
	isSet bool
}

func (v NullableOSVersion) Get() *OSVersion {
	return v.value
}

func (v *NullableOSVersion) Set(val *OSVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableOSVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableOSVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOSVersion(val *OSVersion) *NullableOSVersion {
	return &NullableOSVersion{value: val, isSet: true}
}

func (v NullableOSVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOSVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

