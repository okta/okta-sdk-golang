/*
Okta Admin Management API

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

// checks if the PreRegistrationError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PreRegistrationError{}

// PreRegistrationError Standard Okta error envelope extended for batch pre-registration. Each entry in `errorCauses[]` identifies the failing device by `serialNumber` so the caller can fix every issue in a single round-trip.
type PreRegistrationError struct {
	// Per-entry validation errors. Each cause identifies the failing device by `serialNumber`.
	ErrorCauses []PreRegistrationErrorCause `json:"errorCauses,omitempty"`
	// An Okta code for this type of error
	ErrorCode *string `json:"errorCode,omitempty"`
	// A unique identifier for this error. This can be used by Okta Support to help with troubleshooting.
	ErrorId *string `json:"errorId,omitempty"`
	// An Okta code for this type of error
	ErrorLink *string `json:"errorLink,omitempty"`
	// A short description of what caused this error. Sometimes this contains dynamically-generated information about your specific error.
	ErrorSummary         *string `json:"errorSummary,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PreRegistrationError PreRegistrationError

// NewPreRegistrationError instantiates a new PreRegistrationError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPreRegistrationError() *PreRegistrationError {
	this := PreRegistrationError{}
	return &this
}

// NewPreRegistrationErrorWithDefaults instantiates a new PreRegistrationError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPreRegistrationErrorWithDefaults() *PreRegistrationError {
	this := PreRegistrationError{}
	return &this
}

// GetErrorCauses returns the ErrorCauses field value if set, zero value otherwise.
func (o *PreRegistrationError) GetErrorCauses() []PreRegistrationErrorCause {
	if o == nil || IsNil(o.ErrorCauses) {
		var ret []PreRegistrationErrorCause
		return ret
	}
	return o.ErrorCauses
}

// GetErrorCausesOk returns a tuple with the ErrorCauses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreRegistrationError) GetErrorCausesOk() ([]PreRegistrationErrorCause, bool) {
	if o == nil || IsNil(o.ErrorCauses) {
		return nil, false
	}
	return o.ErrorCauses, true
}

// HasErrorCauses returns a boolean if a field has been set.
func (o *PreRegistrationError) HasErrorCauses() bool {
	if o != nil && !IsNil(o.ErrorCauses) {
		return true
	}

	return false
}

// SetErrorCauses gets a reference to the given []PreRegistrationErrorCause and assigns it to the ErrorCauses field.
func (o *PreRegistrationError) SetErrorCauses(v []PreRegistrationErrorCause) {
	o.ErrorCauses = v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *PreRegistrationError) GetErrorCode() string {
	if o == nil || IsNil(o.ErrorCode) {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreRegistrationError) GetErrorCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorCode) {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *PreRegistrationError) HasErrorCode() bool {
	if o != nil && !IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *PreRegistrationError) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetErrorId returns the ErrorId field value if set, zero value otherwise.
func (o *PreRegistrationError) GetErrorId() string {
	if o == nil || IsNil(o.ErrorId) {
		var ret string
		return ret
	}
	return *o.ErrorId
}

// GetErrorIdOk returns a tuple with the ErrorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreRegistrationError) GetErrorIdOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorId) {
		return nil, false
	}
	return o.ErrorId, true
}

// HasErrorId returns a boolean if a field has been set.
func (o *PreRegistrationError) HasErrorId() bool {
	if o != nil && !IsNil(o.ErrorId) {
		return true
	}

	return false
}

// SetErrorId gets a reference to the given string and assigns it to the ErrorId field.
func (o *PreRegistrationError) SetErrorId(v string) {
	o.ErrorId = &v
}

// GetErrorLink returns the ErrorLink field value if set, zero value otherwise.
func (o *PreRegistrationError) GetErrorLink() string {
	if o == nil || IsNil(o.ErrorLink) {
		var ret string
		return ret
	}
	return *o.ErrorLink
}

// GetErrorLinkOk returns a tuple with the ErrorLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreRegistrationError) GetErrorLinkOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorLink) {
		return nil, false
	}
	return o.ErrorLink, true
}

// HasErrorLink returns a boolean if a field has been set.
func (o *PreRegistrationError) HasErrorLink() bool {
	if o != nil && !IsNil(o.ErrorLink) {
		return true
	}

	return false
}

// SetErrorLink gets a reference to the given string and assigns it to the ErrorLink field.
func (o *PreRegistrationError) SetErrorLink(v string) {
	o.ErrorLink = &v
}

// GetErrorSummary returns the ErrorSummary field value if set, zero value otherwise.
func (o *PreRegistrationError) GetErrorSummary() string {
	if o == nil || IsNil(o.ErrorSummary) {
		var ret string
		return ret
	}
	return *o.ErrorSummary
}

// GetErrorSummaryOk returns a tuple with the ErrorSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreRegistrationError) GetErrorSummaryOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorSummary) {
		return nil, false
	}
	return o.ErrorSummary, true
}

// HasErrorSummary returns a boolean if a field has been set.
func (o *PreRegistrationError) HasErrorSummary() bool {
	if o != nil && !IsNil(o.ErrorSummary) {
		return true
	}

	return false
}

// SetErrorSummary gets a reference to the given string and assigns it to the ErrorSummary field.
func (o *PreRegistrationError) SetErrorSummary(v string) {
	o.ErrorSummary = &v
}

func (o PreRegistrationError) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PreRegistrationError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ErrorCauses) {
		toSerialize["errorCauses"] = o.ErrorCauses
	}
	if !IsNil(o.ErrorCode) {
		toSerialize["errorCode"] = o.ErrorCode
	}
	if !IsNil(o.ErrorId) {
		toSerialize["errorId"] = o.ErrorId
	}
	if !IsNil(o.ErrorLink) {
		toSerialize["errorLink"] = o.ErrorLink
	}
	if !IsNil(o.ErrorSummary) {
		toSerialize["errorSummary"] = o.ErrorSummary
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PreRegistrationError) UnmarshalJSON(data []byte) (err error) {
	varPreRegistrationError := _PreRegistrationError{}

	err = json.Unmarshal(data, &varPreRegistrationError)

	if err != nil {
		return err
	}

	*o = PreRegistrationError(varPreRegistrationError)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "errorCauses")
		delete(additionalProperties, "errorCode")
		delete(additionalProperties, "errorId")
		delete(additionalProperties, "errorLink")
		delete(additionalProperties, "errorSummary")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePreRegistrationError struct {
	value *PreRegistrationError
	isSet bool
}

func (v NullablePreRegistrationError) Get() *PreRegistrationError {
	return v.value
}

func (v *NullablePreRegistrationError) Set(val *PreRegistrationError) {
	v.value = val
	v.isSet = true
}

func (v NullablePreRegistrationError) IsSet() bool {
	return v.isSet
}

func (v *NullablePreRegistrationError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePreRegistrationError(val *PreRegistrationError) *NullablePreRegistrationError {
	return &NullablePreRegistrationError{value: val, isSet: true}
}

func (v NullablePreRegistrationError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePreRegistrationError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
