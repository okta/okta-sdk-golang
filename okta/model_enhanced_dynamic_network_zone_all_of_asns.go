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

// checks if the EnhancedDynamicNetworkZoneAllOfAsns type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnhancedDynamicNetworkZoneAllOfAsns{}

// EnhancedDynamicNetworkZoneAllOfAsns The list of ASNs associated with an Enhanced Dynamic Network Zone
type EnhancedDynamicNetworkZoneAllOfAsns struct {
	// An array of ASNs to include for an Enhanced Dynamic Network Zone
	Include []string `json:"include,omitempty"`
	// An array of ASNs to exclude for an Enhanced Dynamic Network Zone
	Exclude              []string `json:"exclude,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EnhancedDynamicNetworkZoneAllOfAsns EnhancedDynamicNetworkZoneAllOfAsns

// NewEnhancedDynamicNetworkZoneAllOfAsns instantiates a new EnhancedDynamicNetworkZoneAllOfAsns object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnhancedDynamicNetworkZoneAllOfAsns() *EnhancedDynamicNetworkZoneAllOfAsns {
	this := EnhancedDynamicNetworkZoneAllOfAsns{}
	return &this
}

// NewEnhancedDynamicNetworkZoneAllOfAsnsWithDefaults instantiates a new EnhancedDynamicNetworkZoneAllOfAsns object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnhancedDynamicNetworkZoneAllOfAsnsWithDefaults() *EnhancedDynamicNetworkZoneAllOfAsns {
	this := EnhancedDynamicNetworkZoneAllOfAsns{}
	return &this
}

// GetInclude returns the Include field value if set, zero value otherwise.
func (o *EnhancedDynamicNetworkZoneAllOfAsns) GetInclude() []string {
	if o == nil || IsNil(o.Include) {
		var ret []string
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnhancedDynamicNetworkZoneAllOfAsns) GetIncludeOk() ([]string, bool) {
	if o == nil || IsNil(o.Include) {
		return nil, false
	}
	return o.Include, true
}

// HasInclude returns a boolean if a field has been set.
func (o *EnhancedDynamicNetworkZoneAllOfAsns) HasInclude() bool {
	if o != nil && !IsNil(o.Include) {
		return true
	}

	return false
}

// SetInclude gets a reference to the given []string and assigns it to the Include field.
func (o *EnhancedDynamicNetworkZoneAllOfAsns) SetInclude(v []string) {
	o.Include = v
}

// GetExclude returns the Exclude field value if set, zero value otherwise.
func (o *EnhancedDynamicNetworkZoneAllOfAsns) GetExclude() []string {
	if o == nil || IsNil(o.Exclude) {
		var ret []string
		return ret
	}
	return o.Exclude
}

// GetExcludeOk returns a tuple with the Exclude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnhancedDynamicNetworkZoneAllOfAsns) GetExcludeOk() ([]string, bool) {
	if o == nil || IsNil(o.Exclude) {
		return nil, false
	}
	return o.Exclude, true
}

// HasExclude returns a boolean if a field has been set.
func (o *EnhancedDynamicNetworkZoneAllOfAsns) HasExclude() bool {
	if o != nil && !IsNil(o.Exclude) {
		return true
	}

	return false
}

// SetExclude gets a reference to the given []string and assigns it to the Exclude field.
func (o *EnhancedDynamicNetworkZoneAllOfAsns) SetExclude(v []string) {
	o.Exclude = v
}

func (o EnhancedDynamicNetworkZoneAllOfAsns) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnhancedDynamicNetworkZoneAllOfAsns) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Include) {
		toSerialize["include"] = o.Include
	}
	if !IsNil(o.Exclude) {
		toSerialize["exclude"] = o.Exclude
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EnhancedDynamicNetworkZoneAllOfAsns) UnmarshalJSON(data []byte) (err error) {
	varEnhancedDynamicNetworkZoneAllOfAsns := _EnhancedDynamicNetworkZoneAllOfAsns{}

	err = json.Unmarshal(data, &varEnhancedDynamicNetworkZoneAllOfAsns)

	if err != nil {
		return err
	}

	*o = EnhancedDynamicNetworkZoneAllOfAsns(varEnhancedDynamicNetworkZoneAllOfAsns)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "include")
		delete(additionalProperties, "exclude")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEnhancedDynamicNetworkZoneAllOfAsns struct {
	value *EnhancedDynamicNetworkZoneAllOfAsns
	isSet bool
}

func (v NullableEnhancedDynamicNetworkZoneAllOfAsns) Get() *EnhancedDynamicNetworkZoneAllOfAsns {
	return v.value
}

func (v *NullableEnhancedDynamicNetworkZoneAllOfAsns) Set(val *EnhancedDynamicNetworkZoneAllOfAsns) {
	v.value = val
	v.isSet = true
}

func (v NullableEnhancedDynamicNetworkZoneAllOfAsns) IsSet() bool {
	return v.isSet
}

func (v *NullableEnhancedDynamicNetworkZoneAllOfAsns) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnhancedDynamicNetworkZoneAllOfAsns(val *EnhancedDynamicNetworkZoneAllOfAsns) *NullableEnhancedDynamicNetworkZoneAllOfAsns {
	return &NullableEnhancedDynamicNetworkZoneAllOfAsns{value: val, isSet: true}
}

func (v NullableEnhancedDynamicNetworkZoneAllOfAsns) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnhancedDynamicNetworkZoneAllOfAsns) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
