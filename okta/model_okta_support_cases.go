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

// checks if the OktaSupportCases type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OktaSupportCases{}

// OktaSupportCases struct for OktaSupportCases
type OktaSupportCases struct {
	SupportCases         []OktaSupportCase `json:"supportCases,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OktaSupportCases OktaSupportCases

// NewOktaSupportCases instantiates a new OktaSupportCases object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOktaSupportCases() *OktaSupportCases {
	this := OktaSupportCases{}
	return &this
}

// NewOktaSupportCasesWithDefaults instantiates a new OktaSupportCases object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOktaSupportCasesWithDefaults() *OktaSupportCases {
	this := OktaSupportCases{}
	return &this
}

// GetSupportCases returns the SupportCases field value if set, zero value otherwise.
func (o *OktaSupportCases) GetSupportCases() []OktaSupportCase {
	if o == nil || IsNil(o.SupportCases) {
		var ret []OktaSupportCase
		return ret
	}
	return o.SupportCases
}

// GetSupportCasesOk returns a tuple with the SupportCases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaSupportCases) GetSupportCasesOk() ([]OktaSupportCase, bool) {
	if o == nil || IsNil(o.SupportCases) {
		return nil, false
	}
	return o.SupportCases, true
}

// HasSupportCases returns a boolean if a field has been set.
func (o *OktaSupportCases) HasSupportCases() bool {
	if o != nil && !IsNil(o.SupportCases) {
		return true
	}

	return false
}

// SetSupportCases gets a reference to the given []OktaSupportCase and assigns it to the SupportCases field.
func (o *OktaSupportCases) SetSupportCases(v []OktaSupportCase) {
	o.SupportCases = v
}

func (o OktaSupportCases) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OktaSupportCases) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SupportCases) {
		toSerialize["supportCases"] = o.SupportCases
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OktaSupportCases) UnmarshalJSON(data []byte) (err error) {
	varOktaSupportCases := _OktaSupportCases{}

	err = json.Unmarshal(data, &varOktaSupportCases)

	if err != nil {
		return err
	}

	*o = OktaSupportCases(varOktaSupportCases)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "supportCases")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOktaSupportCases struct {
	value *OktaSupportCases
	isSet bool
}

func (v NullableOktaSupportCases) Get() *OktaSupportCases {
	return v.value
}

func (v *NullableOktaSupportCases) Set(val *OktaSupportCases) {
	v.value = val
	v.isSet = true
}

func (v NullableOktaSupportCases) IsSet() bool {
	return v.isSet
}

func (v *NullableOktaSupportCases) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOktaSupportCases(val *OktaSupportCases) *NullableOktaSupportCases {
	return &NullableOktaSupportCases{value: val, isSet: true}
}

func (v NullableOktaSupportCases) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOktaSupportCases) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
