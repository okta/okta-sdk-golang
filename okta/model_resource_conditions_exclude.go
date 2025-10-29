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

// checks if the ResourceConditionsExclude type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceConditionsExclude{}

// ResourceConditionsExclude Specific resources to exclude
type ResourceConditionsExclude struct {
	// List of specific resources to exclude in ORN format
	OktaORN              []string `json:"okta:ORN,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResourceConditionsExclude ResourceConditionsExclude

// NewResourceConditionsExclude instantiates a new ResourceConditionsExclude object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceConditionsExclude() *ResourceConditionsExclude {
	this := ResourceConditionsExclude{}
	return &this
}

// NewResourceConditionsExcludeWithDefaults instantiates a new ResourceConditionsExclude object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceConditionsExcludeWithDefaults() *ResourceConditionsExclude {
	this := ResourceConditionsExclude{}
	return &this
}

// GetOktaORN returns the OktaORN field value if set, zero value otherwise.
func (o *ResourceConditionsExclude) GetOktaORN() []string {
	if o == nil || IsNil(o.OktaORN) {
		var ret []string
		return ret
	}
	return o.OktaORN
}

// GetOktaORNOk returns a tuple with the OktaORN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceConditionsExclude) GetOktaORNOk() ([]string, bool) {
	if o == nil || IsNil(o.OktaORN) {
		return nil, false
	}
	return o.OktaORN, true
}

// HasOktaORN returns a boolean if a field has been set.
func (o *ResourceConditionsExclude) HasOktaORN() bool {
	if o != nil && !IsNil(o.OktaORN) {
		return true
	}

	return false
}

// SetOktaORN gets a reference to the given []string and assigns it to the OktaORN field.
func (o *ResourceConditionsExclude) SetOktaORN(v []string) {
	o.OktaORN = v
}

func (o ResourceConditionsExclude) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceConditionsExclude) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OktaORN) {
		toSerialize["okta:ORN"] = o.OktaORN
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResourceConditionsExclude) UnmarshalJSON(data []byte) (err error) {
	varResourceConditionsExclude := _ResourceConditionsExclude{}

	err = json.Unmarshal(data, &varResourceConditionsExclude)

	if err != nil {
		return err
	}

	*o = ResourceConditionsExclude(varResourceConditionsExclude)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "okta:ORN")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResourceConditionsExclude struct {
	value *ResourceConditionsExclude
	isSet bool
}

func (v NullableResourceConditionsExclude) Get() *ResourceConditionsExclude {
	return v.value
}

func (v *NullableResourceConditionsExclude) Set(val *ResourceConditionsExclude) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceConditionsExclude) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceConditionsExclude) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceConditionsExclude(val *ResourceConditionsExclude) *NullableResourceConditionsExclude {
	return &NullableResourceConditionsExclude{value: val, isSet: true}
}

func (v NullableResourceConditionsExclude) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceConditionsExclude) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
