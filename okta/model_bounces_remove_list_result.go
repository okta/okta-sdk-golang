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

// BouncesRemoveListResult struct for BouncesRemoveListResult
type BouncesRemoveListResult struct {
	Errors []BouncesRemoveListError `json:"errors,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BouncesRemoveListResult BouncesRemoveListResult

// NewBouncesRemoveListResult instantiates a new BouncesRemoveListResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBouncesRemoveListResult() *BouncesRemoveListResult {
	this := BouncesRemoveListResult{}
	return &this
}

// NewBouncesRemoveListResultWithDefaults instantiates a new BouncesRemoveListResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBouncesRemoveListResultWithDefaults() *BouncesRemoveListResult {
	this := BouncesRemoveListResult{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *BouncesRemoveListResult) GetErrors() []BouncesRemoveListError {
	if o == nil || o.Errors == nil {
		var ret []BouncesRemoveListError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BouncesRemoveListResult) GetErrorsOk() ([]BouncesRemoveListError, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *BouncesRemoveListResult) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []BouncesRemoveListError and assigns it to the Errors field.
func (o *BouncesRemoveListResult) SetErrors(v []BouncesRemoveListError) {
	o.Errors = v
}

func (o BouncesRemoveListResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BouncesRemoveListResult) UnmarshalJSON(bytes []byte) (err error) {
	varBouncesRemoveListResult := _BouncesRemoveListResult{}

	err = json.Unmarshal(bytes, &varBouncesRemoveListResult)
	if err == nil {
		*o = BouncesRemoveListResult(varBouncesRemoveListResult)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "errors")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableBouncesRemoveListResult struct {
	value *BouncesRemoveListResult
	isSet bool
}

func (v NullableBouncesRemoveListResult) Get() *BouncesRemoveListResult {
	return v.value
}

func (v *NullableBouncesRemoveListResult) Set(val *BouncesRemoveListResult) {
	v.value = val
	v.isSet = true
}

func (v NullableBouncesRemoveListResult) IsSet() bool {
	return v.isSet
}

func (v *NullableBouncesRemoveListResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBouncesRemoveListResult(val *BouncesRemoveListResult) *NullableBouncesRemoveListResult {
	return &NullableBouncesRemoveListResult{value: val, isSet: true}
}

func (v NullableBouncesRemoveListResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBouncesRemoveListResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

