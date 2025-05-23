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

// ApplicationLicensing struct for ApplicationLicensing
type ApplicationLicensing struct {
	// Number of licenses purchased for the app
	SeatCount *int32 `json:"seatCount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplicationLicensing ApplicationLicensing

// NewApplicationLicensing instantiates a new ApplicationLicensing object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationLicensing() *ApplicationLicensing {
	this := ApplicationLicensing{}
	return &this
}

// NewApplicationLicensingWithDefaults instantiates a new ApplicationLicensing object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationLicensingWithDefaults() *ApplicationLicensing {
	this := ApplicationLicensing{}
	return &this
}

// GetSeatCount returns the SeatCount field value if set, zero value otherwise.
func (o *ApplicationLicensing) GetSeatCount() int32 {
	if o == nil || o.SeatCount == nil {
		var ret int32
		return ret
	}
	return *o.SeatCount
}

// GetSeatCountOk returns a tuple with the SeatCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationLicensing) GetSeatCountOk() (*int32, bool) {
	if o == nil || o.SeatCount == nil {
		return nil, false
	}
	return o.SeatCount, true
}

// HasSeatCount returns a boolean if a field has been set.
func (o *ApplicationLicensing) HasSeatCount() bool {
	if o != nil && o.SeatCount != nil {
		return true
	}

	return false
}

// SetSeatCount gets a reference to the given int32 and assigns it to the SeatCount field.
func (o *ApplicationLicensing) SetSeatCount(v int32) {
	o.SeatCount = &v
}

func (o ApplicationLicensing) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SeatCount != nil {
		toSerialize["seatCount"] = o.SeatCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ApplicationLicensing) UnmarshalJSON(bytes []byte) (err error) {
	varApplicationLicensing := _ApplicationLicensing{}

	err = json.Unmarshal(bytes, &varApplicationLicensing)
	if err == nil {
		*o = ApplicationLicensing(varApplicationLicensing)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "seatCount")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableApplicationLicensing struct {
	value *ApplicationLicensing
	isSet bool
}

func (v NullableApplicationLicensing) Get() *ApplicationLicensing {
	return v.value
}

func (v *NullableApplicationLicensing) Set(val *ApplicationLicensing) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationLicensing) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationLicensing) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationLicensing(val *ApplicationLicensing) *NullableApplicationLicensing {
	return &NullableApplicationLicensing{value: val, isSet: true}
}

func (v NullableApplicationLicensing) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationLicensing) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

