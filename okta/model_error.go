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

// Error struct for Error
type Error struct {
	ErrorCauses []ErrorCause `json:"errorCauses,omitempty"`
	// An Okta code for this type of error
	ErrorCode *string `json:"errorCode,omitempty"`
	// A unique identifier for this error. This can be used by Okta Support to help with troubleshooting.
	ErrorId *string `json:"errorId,omitempty"`
	// An Okta code for this type of error
	ErrorLink *string `json:"errorLink,omitempty"`
	// A short description of what caused this error. Sometimes this contains dynamically-generated information about your specific error.
	ErrorSummary *string `json:"errorSummary,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Error Error

// NewError instantiates a new Error object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewError() *Error {
	this := Error{}
	return &this
}

// NewErrorWithDefaults instantiates a new Error object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorWithDefaults() *Error {
	this := Error{}
	return &this
}

// GetErrorCauses returns the ErrorCauses field value if set, zero value otherwise.
func (o *Error) GetErrorCauses() []ErrorCause {
	if o == nil || o.ErrorCauses == nil {
		var ret []ErrorCause
		return ret
	}
	return o.ErrorCauses
}

// GetErrorCausesOk returns a tuple with the ErrorCauses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Error) GetErrorCausesOk() ([]ErrorCause, bool) {
	if o == nil || o.ErrorCauses == nil {
		return nil, false
	}
	return o.ErrorCauses, true
}

// HasErrorCauses returns a boolean if a field has been set.
func (o *Error) HasErrorCauses() bool {
	if o != nil && o.ErrorCauses != nil {
		return true
	}

	return false
}

// SetErrorCauses gets a reference to the given []ErrorCause and assigns it to the ErrorCauses field.
func (o *Error) SetErrorCauses(v []ErrorCause) {
	o.ErrorCauses = v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *Error) GetErrorCode() string {
	if o == nil || o.ErrorCode == nil {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Error) GetErrorCodeOk() (*string, bool) {
	if o == nil || o.ErrorCode == nil {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *Error) HasErrorCode() bool {
	if o != nil && o.ErrorCode != nil {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *Error) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetErrorId returns the ErrorId field value if set, zero value otherwise.
func (o *Error) GetErrorId() string {
	if o == nil || o.ErrorId == nil {
		var ret string
		return ret
	}
	return *o.ErrorId
}

// GetErrorIdOk returns a tuple with the ErrorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Error) GetErrorIdOk() (*string, bool) {
	if o == nil || o.ErrorId == nil {
		return nil, false
	}
	return o.ErrorId, true
}

// HasErrorId returns a boolean if a field has been set.
func (o *Error) HasErrorId() bool {
	if o != nil && o.ErrorId != nil {
		return true
	}

	return false
}

// SetErrorId gets a reference to the given string and assigns it to the ErrorId field.
func (o *Error) SetErrorId(v string) {
	o.ErrorId = &v
}

// GetErrorLink returns the ErrorLink field value if set, zero value otherwise.
func (o *Error) GetErrorLink() string {
	if o == nil || o.ErrorLink == nil {
		var ret string
		return ret
	}
	return *o.ErrorLink
}

// GetErrorLinkOk returns a tuple with the ErrorLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Error) GetErrorLinkOk() (*string, bool) {
	if o == nil || o.ErrorLink == nil {
		return nil, false
	}
	return o.ErrorLink, true
}

// HasErrorLink returns a boolean if a field has been set.
func (o *Error) HasErrorLink() bool {
	if o != nil && o.ErrorLink != nil {
		return true
	}

	return false
}

// SetErrorLink gets a reference to the given string and assigns it to the ErrorLink field.
func (o *Error) SetErrorLink(v string) {
	o.ErrorLink = &v
}

// GetErrorSummary returns the ErrorSummary field value if set, zero value otherwise.
func (o *Error) GetErrorSummary() string {
	if o == nil || o.ErrorSummary == nil {
		var ret string
		return ret
	}
	return *o.ErrorSummary
}

// GetErrorSummaryOk returns a tuple with the ErrorSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Error) GetErrorSummaryOk() (*string, bool) {
	if o == nil || o.ErrorSummary == nil {
		return nil, false
	}
	return o.ErrorSummary, true
}

// HasErrorSummary returns a boolean if a field has been set.
func (o *Error) HasErrorSummary() bool {
	if o != nil && o.ErrorSummary != nil {
		return true
	}

	return false
}

// SetErrorSummary gets a reference to the given string and assigns it to the ErrorSummary field.
func (o *Error) SetErrorSummary(v string) {
	o.ErrorSummary = &v
}

func (o Error) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ErrorCauses != nil {
		toSerialize["errorCauses"] = o.ErrorCauses
	}
	if o.ErrorCode != nil {
		toSerialize["errorCode"] = o.ErrorCode
	}
	if o.ErrorId != nil {
		toSerialize["errorId"] = o.ErrorId
	}
	if o.ErrorLink != nil {
		toSerialize["errorLink"] = o.ErrorLink
	}
	if o.ErrorSummary != nil {
		toSerialize["errorSummary"] = o.ErrorSummary
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Error) UnmarshalJSON(bytes []byte) (err error) {
	varError := _Error{}

	err = json.Unmarshal(bytes, &varError)
	if err == nil {
		*o = Error(varError)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "errorCauses")
		delete(additionalProperties, "errorCode")
		delete(additionalProperties, "errorId")
		delete(additionalProperties, "errorLink")
		delete(additionalProperties, "errorSummary")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableError struct {
	value *Error
	isSet bool
}

func (v NullableError) Get() *Error {
	return v.value
}

func (v *NullableError) Set(val *Error) {
	v.value = val
	v.isSet = true
}

func (v NullableError) IsSet() bool {
	return v.isSet
}

func (v *NullableError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableError(val *Error) *NullableError {
	return &NullableError{value: val, isSet: true}
}

func (v NullableError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

