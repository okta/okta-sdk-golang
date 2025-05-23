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

// UserFactorPushTransaction struct for UserFactorPushTransaction
type UserFactorPushTransaction struct {
	// Result of the verification transaction
	FactorResult *string `json:"factorResult,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserFactorPushTransaction UserFactorPushTransaction

// NewUserFactorPushTransaction instantiates a new UserFactorPushTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserFactorPushTransaction() *UserFactorPushTransaction {
	this := UserFactorPushTransaction{}
	return &this
}

// NewUserFactorPushTransactionWithDefaults instantiates a new UserFactorPushTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserFactorPushTransactionWithDefaults() *UserFactorPushTransaction {
	this := UserFactorPushTransaction{}
	return &this
}

// GetFactorResult returns the FactorResult field value if set, zero value otherwise.
func (o *UserFactorPushTransaction) GetFactorResult() string {
	if o == nil || o.FactorResult == nil {
		var ret string
		return ret
	}
	return *o.FactorResult
}

// GetFactorResultOk returns a tuple with the FactorResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorPushTransaction) GetFactorResultOk() (*string, bool) {
	if o == nil || o.FactorResult == nil {
		return nil, false
	}
	return o.FactorResult, true
}

// HasFactorResult returns a boolean if a field has been set.
func (o *UserFactorPushTransaction) HasFactorResult() bool {
	if o != nil && o.FactorResult != nil {
		return true
	}

	return false
}

// SetFactorResult gets a reference to the given string and assigns it to the FactorResult field.
func (o *UserFactorPushTransaction) SetFactorResult(v string) {
	o.FactorResult = &v
}

func (o UserFactorPushTransaction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FactorResult != nil {
		toSerialize["factorResult"] = o.FactorResult
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UserFactorPushTransaction) UnmarshalJSON(bytes []byte) (err error) {
	varUserFactorPushTransaction := _UserFactorPushTransaction{}

	err = json.Unmarshal(bytes, &varUserFactorPushTransaction)
	if err == nil {
		*o = UserFactorPushTransaction(varUserFactorPushTransaction)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "factorResult")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableUserFactorPushTransaction struct {
	value *UserFactorPushTransaction
	isSet bool
}

func (v NullableUserFactorPushTransaction) Get() *UserFactorPushTransaction {
	return v.value
}

func (v *NullableUserFactorPushTransaction) Set(val *UserFactorPushTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableUserFactorPushTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableUserFactorPushTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserFactorPushTransaction(val *UserFactorPushTransaction) *NullableUserFactorPushTransaction {
	return &NullableUserFactorPushTransaction{value: val, isSet: true}
}

func (v NullableUserFactorPushTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserFactorPushTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

