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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"fmt"
)

// checks if the CaepCredentialChangeEventReasonAdmin type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CaepCredentialChangeEventReasonAdmin{}

// CaepCredentialChangeEventReasonAdmin struct for CaepCredentialChangeEventReasonAdmin
type CaepCredentialChangeEventReasonAdmin struct {
	// The event reason in English
	En                   string `json:"en"`
	AdditionalProperties map[string]interface{}
}

type _CaepCredentialChangeEventReasonAdmin CaepCredentialChangeEventReasonAdmin

// NewCaepCredentialChangeEventReasonAdmin instantiates a new CaepCredentialChangeEventReasonAdmin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCaepCredentialChangeEventReasonAdmin(en string) *CaepCredentialChangeEventReasonAdmin {
	this := CaepCredentialChangeEventReasonAdmin{}
	this.En = en
	return &this
}

// NewCaepCredentialChangeEventReasonAdminWithDefaults instantiates a new CaepCredentialChangeEventReasonAdmin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCaepCredentialChangeEventReasonAdminWithDefaults() *CaepCredentialChangeEventReasonAdmin {
	this := CaepCredentialChangeEventReasonAdmin{}
	return &this
}

// GetEn returns the En field value
func (o *CaepCredentialChangeEventReasonAdmin) GetEn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.En
}

// GetEnOk returns a tuple with the En field value
// and a boolean to check if the value has been set.
func (o *CaepCredentialChangeEventReasonAdmin) GetEnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.En, true
}

// SetEn sets field value
func (o *CaepCredentialChangeEventReasonAdmin) SetEn(v string) {
	o.En = v
}

func (o CaepCredentialChangeEventReasonAdmin) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CaepCredentialChangeEventReasonAdmin) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["en"] = o.En

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CaepCredentialChangeEventReasonAdmin) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"en",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCaepCredentialChangeEventReasonAdmin := _CaepCredentialChangeEventReasonAdmin{}

	err = json.Unmarshal(data, &varCaepCredentialChangeEventReasonAdmin)

	if err != nil {
		return err
	}

	*o = CaepCredentialChangeEventReasonAdmin(varCaepCredentialChangeEventReasonAdmin)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "en")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCaepCredentialChangeEventReasonAdmin struct {
	value *CaepCredentialChangeEventReasonAdmin
	isSet bool
}

func (v NullableCaepCredentialChangeEventReasonAdmin) Get() *CaepCredentialChangeEventReasonAdmin {
	return v.value
}

func (v *NullableCaepCredentialChangeEventReasonAdmin) Set(val *CaepCredentialChangeEventReasonAdmin) {
	v.value = val
	v.isSet = true
}

func (v NullableCaepCredentialChangeEventReasonAdmin) IsSet() bool {
	return v.isSet
}

func (v *NullableCaepCredentialChangeEventReasonAdmin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCaepCredentialChangeEventReasonAdmin(val *CaepCredentialChangeEventReasonAdmin) *NullableCaepCredentialChangeEventReasonAdmin {
	return &NullableCaepCredentialChangeEventReasonAdmin{value: val, isSet: true}
}

func (v NullableCaepCredentialChangeEventReasonAdmin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCaepCredentialChangeEventReasonAdmin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
