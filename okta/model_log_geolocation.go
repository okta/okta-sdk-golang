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

// checks if the LogGeolocation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogGeolocation{}

// LogGeolocation The latitude and longitude of the geolocation where an action was performed. The object is formatted according to the [ISO 6709](https://www.iso.org/obp/ui/fr/#iso:std:iso:6709:ed-3:v1:en) standard.
type LogGeolocation struct {
	// Latitude which uses two digits for the [integer part](https://www.iso.org/obp/ui/fr/#iso:std:iso:6709:ed-3:v1:en#Latitude)
	Lat *float64 `json:"lat,omitempty"`
	// Longitude which uses three digits for the [integer part](https://www.iso.org/obp/ui/fr/#iso:std:iso:6709:ed-3:v1:en#Longitude)
	Lon                  *float64 `json:"lon,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LogGeolocation LogGeolocation

// NewLogGeolocation instantiates a new LogGeolocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogGeolocation() *LogGeolocation {
	this := LogGeolocation{}
	return &this
}

// NewLogGeolocationWithDefaults instantiates a new LogGeolocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogGeolocationWithDefaults() *LogGeolocation {
	this := LogGeolocation{}
	return &this
}

// GetLat returns the Lat field value if set, zero value otherwise.
func (o *LogGeolocation) GetLat() float64 {
	if o == nil || IsNil(o.Lat) {
		var ret float64
		return ret
	}
	return *o.Lat
}

// GetLatOk returns a tuple with the Lat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogGeolocation) GetLatOk() (*float64, bool) {
	if o == nil || IsNil(o.Lat) {
		return nil, false
	}
	return o.Lat, true
}

// HasLat returns a boolean if a field has been set.
func (o *LogGeolocation) HasLat() bool {
	if o != nil && !IsNil(o.Lat) {
		return true
	}

	return false
}

// SetLat gets a reference to the given float64 and assigns it to the Lat field.
func (o *LogGeolocation) SetLat(v float64) {
	o.Lat = &v
}

// GetLon returns the Lon field value if set, zero value otherwise.
func (o *LogGeolocation) GetLon() float64 {
	if o == nil || IsNil(o.Lon) {
		var ret float64
		return ret
	}
	return *o.Lon
}

// GetLonOk returns a tuple with the Lon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogGeolocation) GetLonOk() (*float64, bool) {
	if o == nil || IsNil(o.Lon) {
		return nil, false
	}
	return o.Lon, true
}

// HasLon returns a boolean if a field has been set.
func (o *LogGeolocation) HasLon() bool {
	if o != nil && !IsNil(o.Lon) {
		return true
	}

	return false
}

// SetLon gets a reference to the given float64 and assigns it to the Lon field.
func (o *LogGeolocation) SetLon(v float64) {
	o.Lon = &v
}

func (o LogGeolocation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogGeolocation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Lat) {
		toSerialize["lat"] = o.Lat
	}
	if !IsNil(o.Lon) {
		toSerialize["lon"] = o.Lon
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LogGeolocation) UnmarshalJSON(data []byte) (err error) {
	varLogGeolocation := _LogGeolocation{}

	err = json.Unmarshal(data, &varLogGeolocation)

	if err != nil {
		return err
	}

	*o = LogGeolocation(varLogGeolocation)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "lat")
		delete(additionalProperties, "lon")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLogGeolocation struct {
	value *LogGeolocation
	isSet bool
}

func (v NullableLogGeolocation) Get() *LogGeolocation {
	return v.value
}

func (v *NullableLogGeolocation) Set(val *LogGeolocation) {
	v.value = val
	v.isSet = true
}

func (v NullableLogGeolocation) IsSet() bool {
	return v.isSet
}

func (v *NullableLogGeolocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogGeolocation(val *LogGeolocation) *NullableLogGeolocation {
	return &NullableLogGeolocation{value: val, isSet: true}
}

func (v NullableLogGeolocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogGeolocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
