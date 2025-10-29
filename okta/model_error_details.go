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

// checks if the ErrorDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorDetails{}

// ErrorDetails Details about an error that occurred during the operation
type ErrorDetails struct {
	// The error code
	Code *string `json:"code,omitempty"`
	// The error message
	Message              *string `json:"message,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ErrorDetails ErrorDetails

// NewErrorDetails instantiates a new ErrorDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorDetails() *ErrorDetails {
	this := ErrorDetails{}
	return &this
}

// NewErrorDetailsWithDefaults instantiates a new ErrorDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorDetailsWithDefaults() *ErrorDetails {
	this := ErrorDetails{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ErrorDetails) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorDetails) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ErrorDetails) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ErrorDetails) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ErrorDetails) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorDetails) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ErrorDetails) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ErrorDetails) SetMessage(v string) {
	o.Message = &v
}

func (o ErrorDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ErrorDetails) UnmarshalJSON(data []byte) (err error) {
	varErrorDetails := _ErrorDetails{}

	err = json.Unmarshal(data, &varErrorDetails)

	if err != nil {
		return err
	}

	*o = ErrorDetails(varErrorDetails)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "message")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableErrorDetails struct {
	value *ErrorDetails
	isSet bool
}

func (v NullableErrorDetails) Get() *ErrorDetails {
	return v.value
}

func (v *NullableErrorDetails) Set(val *ErrorDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorDetails(val *ErrorDetails) *NullableErrorDetails {
	return &NullableErrorDetails{value: val, isSet: true}
}

func (v NullableErrorDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
