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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"fmt"
)

// CaepCredentialChangeEventReasonUser struct for CaepCredentialChangeEventReasonUser
type CaepCredentialChangeEventReasonUser struct {
	// The event reason in English
	En string `json:"en"`
	AdditionalProperties map[string]interface{}
}

type _CaepCredentialChangeEventReasonUser CaepCredentialChangeEventReasonUser

// NewCaepCredentialChangeEventReasonUser instantiates a new CaepCredentialChangeEventReasonUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCaepCredentialChangeEventReasonUser(en string) *CaepCredentialChangeEventReasonUser {
	this := CaepCredentialChangeEventReasonUser{}
	this.En = en
	return &this
}

// NewCaepCredentialChangeEventReasonUserWithDefaults instantiates a new CaepCredentialChangeEventReasonUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCaepCredentialChangeEventReasonUserWithDefaults() *CaepCredentialChangeEventReasonUser {
	this := CaepCredentialChangeEventReasonUser{}
	return &this
}

// GetEn returns the En field value
func (o *CaepCredentialChangeEventReasonUser) GetEn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.En
}

// GetEnOk returns a tuple with the En field value
// and a boolean to check if the value has been set.
func (o *CaepCredentialChangeEventReasonUser) GetEnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.En, true
}

// SetEn sets field value
func (o *CaepCredentialChangeEventReasonUser) SetEn(v string) {
	o.En = v
}

func (o CaepCredentialChangeEventReasonUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["en"] = o.En
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CaepCredentialChangeEventReasonUser) UnmarshalJSON(bytes []byte) (err error) {
	varCaepCredentialChangeEventReasonUser := _CaepCredentialChangeEventReasonUser{}

	err = json.Unmarshal(bytes, &varCaepCredentialChangeEventReasonUser)
	if err == nil {
		*o = CaepCredentialChangeEventReasonUser(varCaepCredentialChangeEventReasonUser)
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

type NullableCaepCredentialChangeEventReasonUser struct {
	value *CaepCredentialChangeEventReasonUser
	isSet bool
}

func (v NullableCaepCredentialChangeEventReasonUser) Get() *CaepCredentialChangeEventReasonUser {
	return v.value
}

func (v *NullableCaepCredentialChangeEventReasonUser) Set(val *CaepCredentialChangeEventReasonUser) {
	v.value = val
	v.isSet = true
}

func (v NullableCaepCredentialChangeEventReasonUser) IsSet() bool {
	return v.isSet
}

func (v *NullableCaepCredentialChangeEventReasonUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCaepCredentialChangeEventReasonUser(val *CaepCredentialChangeEventReasonUser) *NullableCaepCredentialChangeEventReasonUser {
	return &NullableCaepCredentialChangeEventReasonUser{value: val, isSet: true}
}

func (v NullableCaepCredentialChangeEventReasonUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCaepCredentialChangeEventReasonUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

