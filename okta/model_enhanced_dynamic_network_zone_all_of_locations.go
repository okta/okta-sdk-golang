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

// EnhancedDynamicNetworkZoneAllOfLocations <div class=\"x-lifecycle-container\"><x-lifecycle class=\"ea\"></x-lifecycle></div>The list of geolocations to include or exclude for an Enhanced Dynamic Network Zone
type EnhancedDynamicNetworkZoneAllOfLocations struct {
	// An array of geolocations to include for an Enhanced Dynamic Network Zone
	Include []NetworkZoneLocation `json:"include,omitempty"`
	// An array of geolocations to exclude for an Enhanced Dynamic Network Zone
	Exclude []NetworkZoneLocation `json:"exclude,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EnhancedDynamicNetworkZoneAllOfLocations EnhancedDynamicNetworkZoneAllOfLocations

// NewEnhancedDynamicNetworkZoneAllOfLocations instantiates a new EnhancedDynamicNetworkZoneAllOfLocations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnhancedDynamicNetworkZoneAllOfLocations() *EnhancedDynamicNetworkZoneAllOfLocations {
	this := EnhancedDynamicNetworkZoneAllOfLocations{}
	return &this
}

// NewEnhancedDynamicNetworkZoneAllOfLocationsWithDefaults instantiates a new EnhancedDynamicNetworkZoneAllOfLocations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnhancedDynamicNetworkZoneAllOfLocationsWithDefaults() *EnhancedDynamicNetworkZoneAllOfLocations {
	this := EnhancedDynamicNetworkZoneAllOfLocations{}
	return &this
}

// GetInclude returns the Include field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnhancedDynamicNetworkZoneAllOfLocations) GetInclude() []NetworkZoneLocation {
	if o == nil {
		var ret []NetworkZoneLocation
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnhancedDynamicNetworkZoneAllOfLocations) GetIncludeOk() ([]NetworkZoneLocation, bool) {
	if o == nil || o.Include == nil {
		return nil, false
	}
	return o.Include, true
}

// HasInclude returns a boolean if a field has been set.
func (o *EnhancedDynamicNetworkZoneAllOfLocations) HasInclude() bool {
	if o != nil && o.Include != nil {
		return true
	}

	return false
}

// SetInclude gets a reference to the given []NetworkZoneLocation and assigns it to the Include field.
func (o *EnhancedDynamicNetworkZoneAllOfLocations) SetInclude(v []NetworkZoneLocation) {
	o.Include = v
}

// GetExclude returns the Exclude field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnhancedDynamicNetworkZoneAllOfLocations) GetExclude() []NetworkZoneLocation {
	if o == nil {
		var ret []NetworkZoneLocation
		return ret
	}
	return o.Exclude
}

// GetExcludeOk returns a tuple with the Exclude field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnhancedDynamicNetworkZoneAllOfLocations) GetExcludeOk() ([]NetworkZoneLocation, bool) {
	if o == nil || o.Exclude == nil {
		return nil, false
	}
	return o.Exclude, true
}

// HasExclude returns a boolean if a field has been set.
func (o *EnhancedDynamicNetworkZoneAllOfLocations) HasExclude() bool {
	if o != nil && o.Exclude != nil {
		return true
	}

	return false
}

// SetExclude gets a reference to the given []NetworkZoneLocation and assigns it to the Exclude field.
func (o *EnhancedDynamicNetworkZoneAllOfLocations) SetExclude(v []NetworkZoneLocation) {
	o.Exclude = v
}

func (o EnhancedDynamicNetworkZoneAllOfLocations) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Include != nil {
		toSerialize["include"] = o.Include
	}
	if o.Exclude != nil {
		toSerialize["exclude"] = o.Exclude
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EnhancedDynamicNetworkZoneAllOfLocations) UnmarshalJSON(bytes []byte) (err error) {
	varEnhancedDynamicNetworkZoneAllOfLocations := _EnhancedDynamicNetworkZoneAllOfLocations{}

	err = json.Unmarshal(bytes, &varEnhancedDynamicNetworkZoneAllOfLocations)
	if err == nil {
		*o = EnhancedDynamicNetworkZoneAllOfLocations(varEnhancedDynamicNetworkZoneAllOfLocations)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "include")
		delete(additionalProperties, "exclude")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableEnhancedDynamicNetworkZoneAllOfLocations struct {
	value *EnhancedDynamicNetworkZoneAllOfLocations
	isSet bool
}

func (v NullableEnhancedDynamicNetworkZoneAllOfLocations) Get() *EnhancedDynamicNetworkZoneAllOfLocations {
	return v.value
}

func (v *NullableEnhancedDynamicNetworkZoneAllOfLocations) Set(val *EnhancedDynamicNetworkZoneAllOfLocations) {
	v.value = val
	v.isSet = true
}

func (v NullableEnhancedDynamicNetworkZoneAllOfLocations) IsSet() bool {
	return v.isSet
}

func (v *NullableEnhancedDynamicNetworkZoneAllOfLocations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnhancedDynamicNetworkZoneAllOfLocations(val *EnhancedDynamicNetworkZoneAllOfLocations) *NullableEnhancedDynamicNetworkZoneAllOfLocations {
	return &NullableEnhancedDynamicNetworkZoneAllOfLocations{value: val, isSet: true}
}

func (v NullableEnhancedDynamicNetworkZoneAllOfLocations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnhancedDynamicNetworkZoneAllOfLocations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

