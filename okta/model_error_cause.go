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

// checks if the ErrorCause type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorCause{}

// ErrorCause struct for ErrorCause
type ErrorCause struct {
	ErrorSummary         *string `json:"errorSummary,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ErrorCause ErrorCause

// NewErrorCause instantiates a new ErrorCause object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorCause() *ErrorCause {
	this := ErrorCause{}
	return &this
}

// NewErrorCauseWithDefaults instantiates a new ErrorCause object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorCauseWithDefaults() *ErrorCause {
	this := ErrorCause{}
	return &this
}

// GetErrorSummary returns the ErrorSummary field value if set, zero value otherwise.
func (o *ErrorCause) GetErrorSummary() string {
	if o == nil || IsNil(o.ErrorSummary) {
		var ret string
		return ret
	}
	return *o.ErrorSummary
}

// GetErrorSummaryOk returns a tuple with the ErrorSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorCause) GetErrorSummaryOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorSummary) {
		return nil, false
	}
	return o.ErrorSummary, true
}

// HasErrorSummary returns a boolean if a field has been set.
func (o *ErrorCause) HasErrorSummary() bool {
	if o != nil && !IsNil(o.ErrorSummary) {
		return true
	}

	return false
}

// SetErrorSummary gets a reference to the given string and assigns it to the ErrorSummary field.
func (o *ErrorCause) SetErrorSummary(v string) {
	o.ErrorSummary = &v
}

func (o ErrorCause) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorCause) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ErrorSummary) {
		toSerialize["errorSummary"] = o.ErrorSummary
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ErrorCause) UnmarshalJSON(data []byte) (err error) {
	varErrorCause := _ErrorCause{}

	err = json.Unmarshal(data, &varErrorCause)

	if err != nil {
		return err
	}

	*o = ErrorCause(varErrorCause)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "errorSummary")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableErrorCause struct {
	value *ErrorCause
	isSet bool
}

func (v NullableErrorCause) Get() *ErrorCause {
	return v.value
}

func (v *NullableErrorCause) Set(val *ErrorCause) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorCause) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorCause) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorCause(val *ErrorCause) *NullableErrorCause {
	return &NullableErrorCause{value: val, isSet: true}
}

func (v NullableErrorCause) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorCause) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
