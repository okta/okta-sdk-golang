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

// EnhancedDynamicNetworkZoneAllOfAsns <div class=\"x-lifecycle-container\"><x-lifecycle class=\"ea\"></x-lifecycle></div>The list of ASNs associated with an Enhanced Dynamic Network Zone
type EnhancedDynamicNetworkZoneAllOfAsns struct {
	// An array of ASNs to include for an Enhanced Dynamic Network Zone
	Include []string `json:"include,omitempty"`
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
	if o == nil || o.Include == nil {
		var ret []string
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnhancedDynamicNetworkZoneAllOfAsns) GetIncludeOk() ([]string, bool) {
	if o == nil || o.Include == nil {
		return nil, false
	}
	return o.Include, true
}

// HasInclude returns a boolean if a field has been set.
func (o *EnhancedDynamicNetworkZoneAllOfAsns) HasInclude() bool {
	if o != nil && o.Include != nil {
		return true
	}

	return false
}

// SetInclude gets a reference to the given []string and assigns it to the Include field.
func (o *EnhancedDynamicNetworkZoneAllOfAsns) SetInclude(v []string) {
	o.Include = v
}

func (o EnhancedDynamicNetworkZoneAllOfAsns) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Include != nil {
		toSerialize["include"] = o.Include
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EnhancedDynamicNetworkZoneAllOfAsns) UnmarshalJSON(bytes []byte) (err error) {
	varEnhancedDynamicNetworkZoneAllOfAsns := _EnhancedDynamicNetworkZoneAllOfAsns{}

	err = json.Unmarshal(bytes, &varEnhancedDynamicNetworkZoneAllOfAsns)
	if err == nil {
		*o = EnhancedDynamicNetworkZoneAllOfAsns(varEnhancedDynamicNetworkZoneAllOfAsns)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "include")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
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

