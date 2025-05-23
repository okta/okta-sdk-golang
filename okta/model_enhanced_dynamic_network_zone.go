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
	"reflect"
	"strings"
)

// EnhancedDynamicNetworkZone struct for EnhancedDynamicNetworkZone
type EnhancedDynamicNetworkZone struct {
	NetworkZone
	Asns *EnhancedDynamicNetworkZoneAllOfAsns `json:"asns,omitempty"`
	Locations *EnhancedDynamicNetworkZoneAllOfLocations `json:"locations,omitempty"`
	IpServiceCategories *EnhancedDynamicNetworkZoneAllOfIpServiceCategories `json:"ipServiceCategories,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EnhancedDynamicNetworkZone EnhancedDynamicNetworkZone

// NewEnhancedDynamicNetworkZone instantiates a new EnhancedDynamicNetworkZone object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnhancedDynamicNetworkZone(name string, type_ string) *EnhancedDynamicNetworkZone {
	this := EnhancedDynamicNetworkZone{}
	this.Name = name
	this.Type = type_
	return &this
}

// NewEnhancedDynamicNetworkZoneWithDefaults instantiates a new EnhancedDynamicNetworkZone object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnhancedDynamicNetworkZoneWithDefaults() *EnhancedDynamicNetworkZone {
	this := EnhancedDynamicNetworkZone{}
	return &this
}

// GetAsns returns the Asns field value if set, zero value otherwise.
func (o *EnhancedDynamicNetworkZone) GetAsns() EnhancedDynamicNetworkZoneAllOfAsns {
	if o == nil || o.Asns == nil {
		var ret EnhancedDynamicNetworkZoneAllOfAsns
		return ret
	}
	return *o.Asns
}

// GetAsnsOk returns a tuple with the Asns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnhancedDynamicNetworkZone) GetAsnsOk() (*EnhancedDynamicNetworkZoneAllOfAsns, bool) {
	if o == nil || o.Asns == nil {
		return nil, false
	}
	return o.Asns, true
}

// HasAsns returns a boolean if a field has been set.
func (o *EnhancedDynamicNetworkZone) HasAsns() bool {
	if o != nil && o.Asns != nil {
		return true
	}

	return false
}

// SetAsns gets a reference to the given EnhancedDynamicNetworkZoneAllOfAsns and assigns it to the Asns field.
func (o *EnhancedDynamicNetworkZone) SetAsns(v EnhancedDynamicNetworkZoneAllOfAsns) {
	o.Asns = &v
}

// GetLocations returns the Locations field value if set, zero value otherwise.
func (o *EnhancedDynamicNetworkZone) GetLocations() EnhancedDynamicNetworkZoneAllOfLocations {
	if o == nil || o.Locations == nil {
		var ret EnhancedDynamicNetworkZoneAllOfLocations
		return ret
	}
	return *o.Locations
}

// GetLocationsOk returns a tuple with the Locations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnhancedDynamicNetworkZone) GetLocationsOk() (*EnhancedDynamicNetworkZoneAllOfLocations, bool) {
	if o == nil || o.Locations == nil {
		return nil, false
	}
	return o.Locations, true
}

// HasLocations returns a boolean if a field has been set.
func (o *EnhancedDynamicNetworkZone) HasLocations() bool {
	if o != nil && o.Locations != nil {
		return true
	}

	return false
}

// SetLocations gets a reference to the given EnhancedDynamicNetworkZoneAllOfLocations and assigns it to the Locations field.
func (o *EnhancedDynamicNetworkZone) SetLocations(v EnhancedDynamicNetworkZoneAllOfLocations) {
	o.Locations = &v
}

// GetIpServiceCategories returns the IpServiceCategories field value if set, zero value otherwise.
func (o *EnhancedDynamicNetworkZone) GetIpServiceCategories() EnhancedDynamicNetworkZoneAllOfIpServiceCategories {
	if o == nil || o.IpServiceCategories == nil {
		var ret EnhancedDynamicNetworkZoneAllOfIpServiceCategories
		return ret
	}
	return *o.IpServiceCategories
}

// GetIpServiceCategoriesOk returns a tuple with the IpServiceCategories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnhancedDynamicNetworkZone) GetIpServiceCategoriesOk() (*EnhancedDynamicNetworkZoneAllOfIpServiceCategories, bool) {
	if o == nil || o.IpServiceCategories == nil {
		return nil, false
	}
	return o.IpServiceCategories, true
}

// HasIpServiceCategories returns a boolean if a field has been set.
func (o *EnhancedDynamicNetworkZone) HasIpServiceCategories() bool {
	if o != nil && o.IpServiceCategories != nil {
		return true
	}

	return false
}

// SetIpServiceCategories gets a reference to the given EnhancedDynamicNetworkZoneAllOfIpServiceCategories and assigns it to the IpServiceCategories field.
func (o *EnhancedDynamicNetworkZone) SetIpServiceCategories(v EnhancedDynamicNetworkZoneAllOfIpServiceCategories) {
	o.IpServiceCategories = &v
}

func (o EnhancedDynamicNetworkZone) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedNetworkZone, errNetworkZone := json.Marshal(o.NetworkZone)
	if errNetworkZone != nil {
		return []byte{}, errNetworkZone
	}
	errNetworkZone = json.Unmarshal([]byte(serializedNetworkZone), &toSerialize)
	if errNetworkZone != nil {
		return []byte{}, errNetworkZone
	}
	if o.Asns != nil {
		toSerialize["asns"] = o.Asns
	}
	if o.Locations != nil {
		toSerialize["locations"] = o.Locations
	}
	if o.IpServiceCategories != nil {
		toSerialize["ipServiceCategories"] = o.IpServiceCategories
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EnhancedDynamicNetworkZone) UnmarshalJSON(bytes []byte) (err error) {
	type EnhancedDynamicNetworkZoneWithoutEmbeddedStruct struct {
		Asns *EnhancedDynamicNetworkZoneAllOfAsns `json:"asns,omitempty"`
		Locations *EnhancedDynamicNetworkZoneAllOfLocations `json:"locations,omitempty"`
		IpServiceCategories *EnhancedDynamicNetworkZoneAllOfIpServiceCategories `json:"ipServiceCategories,omitempty"`
	}

	varEnhancedDynamicNetworkZoneWithoutEmbeddedStruct := EnhancedDynamicNetworkZoneWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varEnhancedDynamicNetworkZoneWithoutEmbeddedStruct)
	if err == nil {
		varEnhancedDynamicNetworkZone := _EnhancedDynamicNetworkZone{}
		varEnhancedDynamicNetworkZone.Asns = varEnhancedDynamicNetworkZoneWithoutEmbeddedStruct.Asns
		varEnhancedDynamicNetworkZone.Locations = varEnhancedDynamicNetworkZoneWithoutEmbeddedStruct.Locations
		varEnhancedDynamicNetworkZone.IpServiceCategories = varEnhancedDynamicNetworkZoneWithoutEmbeddedStruct.IpServiceCategories
		*o = EnhancedDynamicNetworkZone(varEnhancedDynamicNetworkZone)
	} else {
		return err
	}

	varEnhancedDynamicNetworkZone := _EnhancedDynamicNetworkZone{}

	err = json.Unmarshal(bytes, &varEnhancedDynamicNetworkZone)
	if err == nil {
		o.NetworkZone = varEnhancedDynamicNetworkZone.NetworkZone
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "asns")
		delete(additionalProperties, "locations")
		delete(additionalProperties, "ipServiceCategories")

		// remove fields from embedded structs
		reflectNetworkZone := reflect.ValueOf(o.NetworkZone)
		for i := 0; i < reflectNetworkZone.Type().NumField(); i++ {
			t := reflectNetworkZone.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableEnhancedDynamicNetworkZone struct {
	value *EnhancedDynamicNetworkZone
	isSet bool
}

func (v NullableEnhancedDynamicNetworkZone) Get() *EnhancedDynamicNetworkZone {
	return v.value
}

func (v *NullableEnhancedDynamicNetworkZone) Set(val *EnhancedDynamicNetworkZone) {
	v.value = val
	v.isSet = true
}

func (v NullableEnhancedDynamicNetworkZone) IsSet() bool {
	return v.isSet
}

func (v *NullableEnhancedDynamicNetworkZone) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnhancedDynamicNetworkZone(val *EnhancedDynamicNetworkZone) *NullableEnhancedDynamicNetworkZone {
	return &NullableEnhancedDynamicNetworkZone{value: val, isSet: true}
}

func (v NullableEnhancedDynamicNetworkZone) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnhancedDynamicNetworkZone) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

