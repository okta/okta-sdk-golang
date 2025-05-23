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

// ResendUserFactor struct for ResendUserFactor
type ResendUserFactor struct {
	// Type of the Factor
	FactorType *string `json:"factorType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResendUserFactor ResendUserFactor

// NewResendUserFactor instantiates a new ResendUserFactor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResendUserFactor() *ResendUserFactor {
	this := ResendUserFactor{}
	return &this
}

// NewResendUserFactorWithDefaults instantiates a new ResendUserFactor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResendUserFactorWithDefaults() *ResendUserFactor {
	this := ResendUserFactor{}
	return &this
}

// GetFactorType returns the FactorType field value if set, zero value otherwise.
func (o *ResendUserFactor) GetFactorType() string {
	if o == nil || o.FactorType == nil {
		var ret string
		return ret
	}
	return *o.FactorType
}

// GetFactorTypeOk returns a tuple with the FactorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResendUserFactor) GetFactorTypeOk() (*string, bool) {
	if o == nil || o.FactorType == nil {
		return nil, false
	}
	return o.FactorType, true
}

// HasFactorType returns a boolean if a field has been set.
func (o *ResendUserFactor) HasFactorType() bool {
	if o != nil && o.FactorType != nil {
		return true
	}

	return false
}

// SetFactorType gets a reference to the given string and assigns it to the FactorType field.
func (o *ResendUserFactor) SetFactorType(v string) {
	o.FactorType = &v
}

func (o ResendUserFactor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FactorType != nil {
		toSerialize["factorType"] = o.FactorType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ResendUserFactor) UnmarshalJSON(bytes []byte) (err error) {
	varResendUserFactor := _ResendUserFactor{}

	err = json.Unmarshal(bytes, &varResendUserFactor)
	if err == nil {
		*o = ResendUserFactor(varResendUserFactor)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "factorType")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableResendUserFactor struct {
	value *ResendUserFactor
	isSet bool
}

func (v NullableResendUserFactor) Get() *ResendUserFactor {
	return v.value
}

func (v *NullableResendUserFactor) Set(val *ResendUserFactor) {
	v.value = val
	v.isSet = true
}

func (v NullableResendUserFactor) IsSet() bool {
	return v.isSet
}

func (v *NullableResendUserFactor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResendUserFactor(val *ResendUserFactor) *NullableResendUserFactor {
	return &NullableResendUserFactor{value: val, isSet: true}
}

func (v NullableResendUserFactor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResendUserFactor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

