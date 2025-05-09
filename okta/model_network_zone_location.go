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

// NetworkZoneLocation struct for NetworkZoneLocation
type NetworkZoneLocation struct {
	// The two-character ISO-3166-1 country code. Don't use continent codes since they are treated as generic codes for undesignated countries. <br>For example: `US`
	Country *string `json:"country,omitempty"`
	// (Optional) The ISO-3166-2 region code appended to the country code (`countryCode-regionCode`), or `null` if empty. Don't use continent codes since they are treated as generic codes for undesignated regions. <br>For example: `CA` (for `US-CA` country and region code)
	Region *string `json:"region,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NetworkZoneLocation NetworkZoneLocation

// NewNetworkZoneLocation instantiates a new NetworkZoneLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkZoneLocation() *NetworkZoneLocation {
	this := NetworkZoneLocation{}
	return &this
}

// NewNetworkZoneLocationWithDefaults instantiates a new NetworkZoneLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkZoneLocationWithDefaults() *NetworkZoneLocation {
	this := NetworkZoneLocation{}
	return &this
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *NetworkZoneLocation) GetCountry() string {
	if o == nil || o.Country == nil {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkZoneLocation) GetCountryOk() (*string, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *NetworkZoneLocation) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *NetworkZoneLocation) SetCountry(v string) {
	o.Country = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *NetworkZoneLocation) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkZoneLocation) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *NetworkZoneLocation) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *NetworkZoneLocation) SetRegion(v string) {
	o.Region = &v
}

func (o NetworkZoneLocation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Country != nil {
		toSerialize["country"] = o.Country
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NetworkZoneLocation) UnmarshalJSON(bytes []byte) (err error) {
	varNetworkZoneLocation := _NetworkZoneLocation{}

	err = json.Unmarshal(bytes, &varNetworkZoneLocation)
	if err == nil {
		*o = NetworkZoneLocation(varNetworkZoneLocation)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "country")
		delete(additionalProperties, "region")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableNetworkZoneLocation struct {
	value *NetworkZoneLocation
	isSet bool
}

func (v NullableNetworkZoneLocation) Get() *NetworkZoneLocation {
	return v.value
}

func (v *NullableNetworkZoneLocation) Set(val *NetworkZoneLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkZoneLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkZoneLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkZoneLocation(val *NetworkZoneLocation) *NullableNetworkZoneLocation {
	return &NullableNetworkZoneLocation{value: val, isSet: true}
}

func (v NullableNetworkZoneLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkZoneLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

