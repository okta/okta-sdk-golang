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

// checks if the EnhancedDynamicNetworkZoneAllOfIpServiceCategories type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnhancedDynamicNetworkZoneAllOfIpServiceCategories{}

// EnhancedDynamicNetworkZoneAllOfIpServiceCategories IP services, such as a proxy or VPN, to include or exclude for an Enhanced Dynamic Network Zone
type EnhancedDynamicNetworkZoneAllOfIpServiceCategories struct {
	// IP services to include for an Enhanced Dynamic Network Zone
	Include []string `json:"include,omitempty"`
	// IP services to exclude for an Enhanced Dynamic Network Zone
	Exclude              []string `json:"exclude,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EnhancedDynamicNetworkZoneAllOfIpServiceCategories EnhancedDynamicNetworkZoneAllOfIpServiceCategories

// NewEnhancedDynamicNetworkZoneAllOfIpServiceCategories instantiates a new EnhancedDynamicNetworkZoneAllOfIpServiceCategories object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnhancedDynamicNetworkZoneAllOfIpServiceCategories() *EnhancedDynamicNetworkZoneAllOfIpServiceCategories {
	this := EnhancedDynamicNetworkZoneAllOfIpServiceCategories{}
	return &this
}

// NewEnhancedDynamicNetworkZoneAllOfIpServiceCategoriesWithDefaults instantiates a new EnhancedDynamicNetworkZoneAllOfIpServiceCategories object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnhancedDynamicNetworkZoneAllOfIpServiceCategoriesWithDefaults() *EnhancedDynamicNetworkZoneAllOfIpServiceCategories {
	this := EnhancedDynamicNetworkZoneAllOfIpServiceCategories{}
	return &this
}

// GetInclude returns the Include field value if set, zero value otherwise.
func (o *EnhancedDynamicNetworkZoneAllOfIpServiceCategories) GetInclude() []string {
	if o == nil || IsNil(o.Include) {
		var ret []string
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnhancedDynamicNetworkZoneAllOfIpServiceCategories) GetIncludeOk() ([]string, bool) {
	if o == nil || IsNil(o.Include) {
		return nil, false
	}
	return o.Include, true
}

// HasInclude returns a boolean if a field has been set.
func (o *EnhancedDynamicNetworkZoneAllOfIpServiceCategories) HasInclude() bool {
	if o != nil && !IsNil(o.Include) {
		return true
	}

	return false
}

// SetInclude gets a reference to the given []string and assigns it to the Include field.
func (o *EnhancedDynamicNetworkZoneAllOfIpServiceCategories) SetInclude(v []string) {
	o.Include = v
}

// GetExclude returns the Exclude field value if set, zero value otherwise.
func (o *EnhancedDynamicNetworkZoneAllOfIpServiceCategories) GetExclude() []string {
	if o == nil || IsNil(o.Exclude) {
		var ret []string
		return ret
	}
	return o.Exclude
}

// GetExcludeOk returns a tuple with the Exclude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnhancedDynamicNetworkZoneAllOfIpServiceCategories) GetExcludeOk() ([]string, bool) {
	if o == nil || IsNil(o.Exclude) {
		return nil, false
	}
	return o.Exclude, true
}

// HasExclude returns a boolean if a field has been set.
func (o *EnhancedDynamicNetworkZoneAllOfIpServiceCategories) HasExclude() bool {
	if o != nil && !IsNil(o.Exclude) {
		return true
	}

	return false
}

// SetExclude gets a reference to the given []string and assigns it to the Exclude field.
func (o *EnhancedDynamicNetworkZoneAllOfIpServiceCategories) SetExclude(v []string) {
	o.Exclude = v
}

func (o EnhancedDynamicNetworkZoneAllOfIpServiceCategories) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnhancedDynamicNetworkZoneAllOfIpServiceCategories) ToMap() (map[string]interface{}, error) {
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

func (o *EnhancedDynamicNetworkZoneAllOfIpServiceCategories) UnmarshalJSON(data []byte) (err error) {
	varEnhancedDynamicNetworkZoneAllOfIpServiceCategories := _EnhancedDynamicNetworkZoneAllOfIpServiceCategories{}

	err = json.Unmarshal(data, &varEnhancedDynamicNetworkZoneAllOfIpServiceCategories)

	if err != nil {
		return err
	}

	*o = EnhancedDynamicNetworkZoneAllOfIpServiceCategories(varEnhancedDynamicNetworkZoneAllOfIpServiceCategories)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "include")
		delete(additionalProperties, "exclude")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEnhancedDynamicNetworkZoneAllOfIpServiceCategories struct {
	value *EnhancedDynamicNetworkZoneAllOfIpServiceCategories
	isSet bool
}

func (v NullableEnhancedDynamicNetworkZoneAllOfIpServiceCategories) Get() *EnhancedDynamicNetworkZoneAllOfIpServiceCategories {
	return v.value
}

func (v *NullableEnhancedDynamicNetworkZoneAllOfIpServiceCategories) Set(val *EnhancedDynamicNetworkZoneAllOfIpServiceCategories) {
	v.value = val
	v.isSet = true
}

func (v NullableEnhancedDynamicNetworkZoneAllOfIpServiceCategories) IsSet() bool {
	return v.isSet
}

func (v *NullableEnhancedDynamicNetworkZoneAllOfIpServiceCategories) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnhancedDynamicNetworkZoneAllOfIpServiceCategories(val *EnhancedDynamicNetworkZoneAllOfIpServiceCategories) *NullableEnhancedDynamicNetworkZoneAllOfIpServiceCategories {
	return &NullableEnhancedDynamicNetworkZoneAllOfIpServiceCategories{value: val, isSet: true}
}

func (v NullableEnhancedDynamicNetworkZoneAllOfIpServiceCategories) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnhancedDynamicNetworkZoneAllOfIpServiceCategories) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
