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

// CaepDeviceComplianceChangeEventReasonUser struct for CaepDeviceComplianceChangeEventReasonUser
type CaepDeviceComplianceChangeEventReasonUser struct {
	// The event reason in English
	En string `json:"en"`
	AdditionalProperties map[string]interface{}
}

type _CaepDeviceComplianceChangeEventReasonUser CaepDeviceComplianceChangeEventReasonUser

// NewCaepDeviceComplianceChangeEventReasonUser instantiates a new CaepDeviceComplianceChangeEventReasonUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCaepDeviceComplianceChangeEventReasonUser(en string) *CaepDeviceComplianceChangeEventReasonUser {
	this := CaepDeviceComplianceChangeEventReasonUser{}
	this.En = en
	return &this
}

// NewCaepDeviceComplianceChangeEventReasonUserWithDefaults instantiates a new CaepDeviceComplianceChangeEventReasonUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCaepDeviceComplianceChangeEventReasonUserWithDefaults() *CaepDeviceComplianceChangeEventReasonUser {
	this := CaepDeviceComplianceChangeEventReasonUser{}
	return &this
}

// GetEn returns the En field value
func (o *CaepDeviceComplianceChangeEventReasonUser) GetEn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.En
}

// GetEnOk returns a tuple with the En field value
// and a boolean to check if the value has been set.
func (o *CaepDeviceComplianceChangeEventReasonUser) GetEnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.En, true
}

// SetEn sets field value
func (o *CaepDeviceComplianceChangeEventReasonUser) SetEn(v string) {
	o.En = v
}

func (o CaepDeviceComplianceChangeEventReasonUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["en"] = o.En
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CaepDeviceComplianceChangeEventReasonUser) UnmarshalJSON(bytes []byte) (err error) {
	varCaepDeviceComplianceChangeEventReasonUser := _CaepDeviceComplianceChangeEventReasonUser{}

	err = json.Unmarshal(bytes, &varCaepDeviceComplianceChangeEventReasonUser)
	if err == nil {
		*o = CaepDeviceComplianceChangeEventReasonUser(varCaepDeviceComplianceChangeEventReasonUser)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "en")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableCaepDeviceComplianceChangeEventReasonUser struct {
	value *CaepDeviceComplianceChangeEventReasonUser
	isSet bool
}

func (v NullableCaepDeviceComplianceChangeEventReasonUser) Get() *CaepDeviceComplianceChangeEventReasonUser {
	return v.value
}

func (v *NullableCaepDeviceComplianceChangeEventReasonUser) Set(val *CaepDeviceComplianceChangeEventReasonUser) {
	v.value = val
	v.isSet = true
}

func (v NullableCaepDeviceComplianceChangeEventReasonUser) IsSet() bool {
	return v.isSet
}

func (v *NullableCaepDeviceComplianceChangeEventReasonUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCaepDeviceComplianceChangeEventReasonUser(val *CaepDeviceComplianceChangeEventReasonUser) *NullableCaepDeviceComplianceChangeEventReasonUser {
	return &NullableCaepDeviceComplianceChangeEventReasonUser{value: val, isSet: true}
}

func (v NullableCaepDeviceComplianceChangeEventReasonUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCaepDeviceComplianceChangeEventReasonUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

