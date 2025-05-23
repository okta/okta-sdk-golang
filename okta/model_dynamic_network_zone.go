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

// DynamicNetworkZone struct for DynamicNetworkZone
type DynamicNetworkZone struct {
	NetworkZone
	// An array of ASNs for a Network Zone
	Asns []string `json:"asns,omitempty"`
	// The proxy type used for a Dynamic Network Zone
	ProxyType *string `json:"proxyType,omitempty"`
	// An array of geolocations for a Dynamic Network Zone
	Locations []NetworkZoneLocation `json:"locations,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DynamicNetworkZone DynamicNetworkZone

// NewDynamicNetworkZone instantiates a new DynamicNetworkZone object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDynamicNetworkZone(name string, type_ string) *DynamicNetworkZone {
	this := DynamicNetworkZone{}
	this.Name = name
	this.Type = type_
	return &this
}

// NewDynamicNetworkZoneWithDefaults instantiates a new DynamicNetworkZone object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDynamicNetworkZoneWithDefaults() *DynamicNetworkZone {
	this := DynamicNetworkZone{}
	return &this
}

// GetAsns returns the Asns field value if set, zero value otherwise.
func (o *DynamicNetworkZone) GetAsns() []string {
	if o == nil || o.Asns == nil {
		var ret []string
		return ret
	}
	return o.Asns
}

// GetAsnsOk returns a tuple with the Asns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicNetworkZone) GetAsnsOk() ([]string, bool) {
	if o == nil || o.Asns == nil {
		return nil, false
	}
	return o.Asns, true
}

// HasAsns returns a boolean if a field has been set.
func (o *DynamicNetworkZone) HasAsns() bool {
	if o != nil && o.Asns != nil {
		return true
	}

	return false
}

// SetAsns gets a reference to the given []string and assigns it to the Asns field.
func (o *DynamicNetworkZone) SetAsns(v []string) {
	o.Asns = v
}

// GetProxyType returns the ProxyType field value if set, zero value otherwise.
func (o *DynamicNetworkZone) GetProxyType() string {
	if o == nil || o.ProxyType == nil {
		var ret string
		return ret
	}
	return *o.ProxyType
}

// GetProxyTypeOk returns a tuple with the ProxyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicNetworkZone) GetProxyTypeOk() (*string, bool) {
	if o == nil || o.ProxyType == nil {
		return nil, false
	}
	return o.ProxyType, true
}

// HasProxyType returns a boolean if a field has been set.
func (o *DynamicNetworkZone) HasProxyType() bool {
	if o != nil && o.ProxyType != nil {
		return true
	}

	return false
}

// SetProxyType gets a reference to the given string and assigns it to the ProxyType field.
func (o *DynamicNetworkZone) SetProxyType(v string) {
	o.ProxyType = &v
}

// GetLocations returns the Locations field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DynamicNetworkZone) GetLocations() []NetworkZoneLocation {
	if o == nil {
		var ret []NetworkZoneLocation
		return ret
	}
	return o.Locations
}

// GetLocationsOk returns a tuple with the Locations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DynamicNetworkZone) GetLocationsOk() ([]NetworkZoneLocation, bool) {
	if o == nil || o.Locations == nil {
		return nil, false
	}
	return o.Locations, true
}

// HasLocations returns a boolean if a field has been set.
func (o *DynamicNetworkZone) HasLocations() bool {
	if o != nil && o.Locations != nil {
		return true
	}

	return false
}

// SetLocations gets a reference to the given []NetworkZoneLocation and assigns it to the Locations field.
func (o *DynamicNetworkZone) SetLocations(v []NetworkZoneLocation) {
	o.Locations = v
}

func (o DynamicNetworkZone) MarshalJSON() ([]byte, error) {
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
	if o.ProxyType != nil {
		toSerialize["proxyType"] = o.ProxyType
	}
	if o.Locations != nil {
		toSerialize["locations"] = o.Locations
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DynamicNetworkZone) UnmarshalJSON(bytes []byte) (err error) {
	type DynamicNetworkZoneWithoutEmbeddedStruct struct {
		// An array of ASNs for a Network Zone
		Asns []string `json:"asns,omitempty"`
		// The proxy type used for a Dynamic Network Zone
		ProxyType *string `json:"proxyType,omitempty"`
		// An array of geolocations for a Dynamic Network Zone
		Locations []NetworkZoneLocation `json:"locations,omitempty"`
	}

	varDynamicNetworkZoneWithoutEmbeddedStruct := DynamicNetworkZoneWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varDynamicNetworkZoneWithoutEmbeddedStruct)
	if err == nil {
		varDynamicNetworkZone := _DynamicNetworkZone{}
		varDynamicNetworkZone.Asns = varDynamicNetworkZoneWithoutEmbeddedStruct.Asns
		varDynamicNetworkZone.ProxyType = varDynamicNetworkZoneWithoutEmbeddedStruct.ProxyType
		varDynamicNetworkZone.Locations = varDynamicNetworkZoneWithoutEmbeddedStruct.Locations
		*o = DynamicNetworkZone(varDynamicNetworkZone)
	} else {
		return err
	}

	varDynamicNetworkZone := _DynamicNetworkZone{}

	err = json.Unmarshal(bytes, &varDynamicNetworkZone)
	if err == nil {
		o.NetworkZone = varDynamicNetworkZone.NetworkZone
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "asns")
		delete(additionalProperties, "proxyType")
		delete(additionalProperties, "locations")

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

type NullableDynamicNetworkZone struct {
	value *DynamicNetworkZone
	isSet bool
}

func (v NullableDynamicNetworkZone) Get() *DynamicNetworkZone {
	return v.value
}

func (v *NullableDynamicNetworkZone) Set(val *DynamicNetworkZone) {
	v.value = val
	v.isSet = true
}

func (v NullableDynamicNetworkZone) IsSet() bool {
	return v.isSet
}

func (v *NullableDynamicNetworkZone) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDynamicNetworkZone(val *DynamicNetworkZone) *NullableDynamicNetworkZone {
	return &NullableDynamicNetworkZone{value: val, isSet: true}
}

func (v NullableDynamicNetworkZone) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDynamicNetworkZone) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

