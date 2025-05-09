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

// LogGeographicalContext struct for LogGeographicalContext
type LogGeographicalContext struct {
	City *string `json:"city,omitempty"`
	Country *string `json:"country,omitempty"`
	Geolocation *LogGeolocation `json:"geolocation,omitempty"`
	PostalCode *string `json:"postalCode,omitempty"`
	State *string `json:"state,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LogGeographicalContext LogGeographicalContext

// NewLogGeographicalContext instantiates a new LogGeographicalContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogGeographicalContext() *LogGeographicalContext {
	this := LogGeographicalContext{}
	return &this
}

// NewLogGeographicalContextWithDefaults instantiates a new LogGeographicalContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogGeographicalContextWithDefaults() *LogGeographicalContext {
	this := LogGeographicalContext{}
	return &this
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *LogGeographicalContext) GetCity() string {
	if o == nil || o.City == nil {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogGeographicalContext) GetCityOk() (*string, bool) {
	if o == nil || o.City == nil {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *LogGeographicalContext) HasCity() bool {
	if o != nil && o.City != nil {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *LogGeographicalContext) SetCity(v string) {
	o.City = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *LogGeographicalContext) GetCountry() string {
	if o == nil || o.Country == nil {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogGeographicalContext) GetCountryOk() (*string, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *LogGeographicalContext) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *LogGeographicalContext) SetCountry(v string) {
	o.Country = &v
}

// GetGeolocation returns the Geolocation field value if set, zero value otherwise.
func (o *LogGeographicalContext) GetGeolocation() LogGeolocation {
	if o == nil || o.Geolocation == nil {
		var ret LogGeolocation
		return ret
	}
	return *o.Geolocation
}

// GetGeolocationOk returns a tuple with the Geolocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogGeographicalContext) GetGeolocationOk() (*LogGeolocation, bool) {
	if o == nil || o.Geolocation == nil {
		return nil, false
	}
	return o.Geolocation, true
}

// HasGeolocation returns a boolean if a field has been set.
func (o *LogGeographicalContext) HasGeolocation() bool {
	if o != nil && o.Geolocation != nil {
		return true
	}

	return false
}

// SetGeolocation gets a reference to the given LogGeolocation and assigns it to the Geolocation field.
func (o *LogGeographicalContext) SetGeolocation(v LogGeolocation) {
	o.Geolocation = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *LogGeographicalContext) GetPostalCode() string {
	if o == nil || o.PostalCode == nil {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogGeographicalContext) GetPostalCodeOk() (*string, bool) {
	if o == nil || o.PostalCode == nil {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *LogGeographicalContext) HasPostalCode() bool {
	if o != nil && o.PostalCode != nil {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *LogGeographicalContext) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *LogGeographicalContext) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogGeographicalContext) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *LogGeographicalContext) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *LogGeographicalContext) SetState(v string) {
	o.State = &v
}

func (o LogGeographicalContext) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.City != nil {
		toSerialize["city"] = o.City
	}
	if o.Country != nil {
		toSerialize["country"] = o.Country
	}
	if o.Geolocation != nil {
		toSerialize["geolocation"] = o.Geolocation
	}
	if o.PostalCode != nil {
		toSerialize["postalCode"] = o.PostalCode
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LogGeographicalContext) UnmarshalJSON(bytes []byte) (err error) {
	varLogGeographicalContext := _LogGeographicalContext{}

	err = json.Unmarshal(bytes, &varLogGeographicalContext)
	if err == nil {
		*o = LogGeographicalContext(varLogGeographicalContext)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "city")
		delete(additionalProperties, "country")
		delete(additionalProperties, "geolocation")
		delete(additionalProperties, "postalCode")
		delete(additionalProperties, "state")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableLogGeographicalContext struct {
	value *LogGeographicalContext
	isSet bool
}

func (v NullableLogGeographicalContext) Get() *LogGeographicalContext {
	return v.value
}

func (v *NullableLogGeographicalContext) Set(val *LogGeographicalContext) {
	v.value = val
	v.isSet = true
}

func (v NullableLogGeographicalContext) IsSet() bool {
	return v.isSet
}

func (v *NullableLogGeographicalContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogGeographicalContext(val *LogGeographicalContext) *NullableLogGeographicalContext {
	return &NullableLogGeographicalContext{value: val, isSet: true}
}

func (v NullableLogGeographicalContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogGeographicalContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

