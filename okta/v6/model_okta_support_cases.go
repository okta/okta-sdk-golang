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
)

// OktaSupportCases struct for OktaSupportCases
type OktaSupportCases struct {
	SupportCases []OktaSupportCase `json:"supportCases,omitempty"`
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
	if o == nil || o.SupportCases == nil {
		var ret []OktaSupportCase
		return ret
	}
	return o.SupportCases
}

// GetSupportCasesOk returns a tuple with the SupportCases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaSupportCases) GetSupportCasesOk() ([]OktaSupportCase, bool) {
	if o == nil || o.SupportCases == nil {
		return nil, false
	}
	return o.SupportCases, true
}

// HasSupportCases returns a boolean if a field has been set.
func (o *OktaSupportCases) HasSupportCases() bool {
	if o != nil && o.SupportCases != nil {
		return true
	}

	return false
}

// SetSupportCases gets a reference to the given []OktaSupportCase and assigns it to the SupportCases field.
func (o *OktaSupportCases) SetSupportCases(v []OktaSupportCase) {
	o.SupportCases = v
}

func (o OktaSupportCases) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SupportCases != nil {
		toSerialize["supportCases"] = o.SupportCases
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OktaSupportCases) UnmarshalJSON(bytes []byte) (err error) {
	varOktaSupportCases := _OktaSupportCases{}

	err = json.Unmarshal(bytes, &varOktaSupportCases)
	if err == nil {
		*o = OktaSupportCases(varOktaSupportCases)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "supportCases")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
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

