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

API version: 2025.10.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the BouncesRemoveListResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BouncesRemoveListResult{}

// BouncesRemoveListResult struct for BouncesRemoveListResult
type BouncesRemoveListResult struct {
	// A list of emails that wasn't added to the email-bounced remove list and the error reason
	Errors               []BouncesRemoveListError `json:"errors,omitempty"`
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
	if o == nil || IsNil(o.Errors) {
		var ret []BouncesRemoveListError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BouncesRemoveListResult) GetErrorsOk() ([]BouncesRemoveListError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *BouncesRemoveListResult) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []BouncesRemoveListError and assigns it to the Errors field.
func (o *BouncesRemoveListResult) SetErrors(v []BouncesRemoveListError) {
	o.Errors = v
}

func (o BouncesRemoveListResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BouncesRemoveListResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BouncesRemoveListResult) UnmarshalJSON(data []byte) (err error) {
	varBouncesRemoveListResult := _BouncesRemoveListResult{}

	err = json.Unmarshal(data, &varBouncesRemoveListResult)

	if err != nil {
		return err
	}

	*o = BouncesRemoveListResult(varBouncesRemoveListResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "errors")
		o.AdditionalProperties = additionalProperties
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
