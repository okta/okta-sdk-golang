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

// checks if the Error409 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Error409{}

// Error409 Conflict error object
type Error409 struct {
	// Another request has already been received for the settings for this email template
	ErrorCauses []string `json:"errorCauses,omitempty"`
	// E0000254
	ErrorCode *string `json:"errorCode,omitempty"`
	// sampleH3iLB6bpBcbnV9E09Fy
	ErrorId *string `json:"errorId,omitempty"`
	// E0000254
	ErrorLink *string `json:"errorLink,omitempty"`
	// Another request has already been received for the settings for this email template
	ErrorSummary         *string `json:"errorSummary,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Error409 Error409

// NewError409 instantiates a new Error409 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewError409() *Error409 {
	this := Error409{}
	return &this
}

// NewError409WithDefaults instantiates a new Error409 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewError409WithDefaults() *Error409 {
	this := Error409{}
	return &this
}

// GetErrorCauses returns the ErrorCauses field value if set, zero value otherwise.
func (o *Error409) GetErrorCauses() []string {
	if o == nil || IsNil(o.ErrorCauses) {
		var ret []string
		return ret
	}
	return o.ErrorCauses
}

// GetErrorCausesOk returns a tuple with the ErrorCauses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Error409) GetErrorCausesOk() ([]string, bool) {
	if o == nil || IsNil(o.ErrorCauses) {
		return nil, false
	}
	return o.ErrorCauses, true
}

// HasErrorCauses returns a boolean if a field has been set.
func (o *Error409) HasErrorCauses() bool {
	if o != nil && !IsNil(o.ErrorCauses) {
		return true
	}

	return false
}

// SetErrorCauses gets a reference to the given []string and assigns it to the ErrorCauses field.
func (o *Error409) SetErrorCauses(v []string) {
	o.ErrorCauses = v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *Error409) GetErrorCode() string {
	if o == nil || IsNil(o.ErrorCode) {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Error409) GetErrorCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorCode) {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *Error409) HasErrorCode() bool {
	if o != nil && !IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *Error409) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetErrorId returns the ErrorId field value if set, zero value otherwise.
func (o *Error409) GetErrorId() string {
	if o == nil || IsNil(o.ErrorId) {
		var ret string
		return ret
	}
	return *o.ErrorId
}

// GetErrorIdOk returns a tuple with the ErrorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Error409) GetErrorIdOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorId) {
		return nil, false
	}
	return o.ErrorId, true
}

// HasErrorId returns a boolean if a field has been set.
func (o *Error409) HasErrorId() bool {
	if o != nil && !IsNil(o.ErrorId) {
		return true
	}

	return false
}

// SetErrorId gets a reference to the given string and assigns it to the ErrorId field.
func (o *Error409) SetErrorId(v string) {
	o.ErrorId = &v
}

// GetErrorLink returns the ErrorLink field value if set, zero value otherwise.
func (o *Error409) GetErrorLink() string {
	if o == nil || IsNil(o.ErrorLink) {
		var ret string
		return ret
	}
	return *o.ErrorLink
}

// GetErrorLinkOk returns a tuple with the ErrorLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Error409) GetErrorLinkOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorLink) {
		return nil, false
	}
	return o.ErrorLink, true
}

// HasErrorLink returns a boolean if a field has been set.
func (o *Error409) HasErrorLink() bool {
	if o != nil && !IsNil(o.ErrorLink) {
		return true
	}

	return false
}

// SetErrorLink gets a reference to the given string and assigns it to the ErrorLink field.
func (o *Error409) SetErrorLink(v string) {
	o.ErrorLink = &v
}

// GetErrorSummary returns the ErrorSummary field value if set, zero value otherwise.
func (o *Error409) GetErrorSummary() string {
	if o == nil || IsNil(o.ErrorSummary) {
		var ret string
		return ret
	}
	return *o.ErrorSummary
}

// GetErrorSummaryOk returns a tuple with the ErrorSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Error409) GetErrorSummaryOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorSummary) {
		return nil, false
	}
	return o.ErrorSummary, true
}

// HasErrorSummary returns a boolean if a field has been set.
func (o *Error409) HasErrorSummary() bool {
	if o != nil && !IsNil(o.ErrorSummary) {
		return true
	}

	return false
}

// SetErrorSummary gets a reference to the given string and assigns it to the ErrorSummary field.
func (o *Error409) SetErrorSummary(v string) {
	o.ErrorSummary = &v
}

func (o Error409) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Error409) ToMap() (map[string]interface{}, error) {
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

func (o *Error409) UnmarshalJSON(data []byte) (err error) {
	varError409 := _Error409{}

	err = json.Unmarshal(data, &varError409)

	if err != nil {
		return err
	}

	*o = Error409(varError409)

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

type NullableError409 struct {
	value *Error409
	isSet bool
}

func (v NullableError409) Get() *Error409 {
	return v.value
}

func (v *NullableError409) Set(val *Error409) {
	v.value = val
	v.isSet = true
}

func (v NullableError409) IsSet() bool {
	return v.isSet
}

func (v *NullableError409) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableError409(val *Error409) *NullableError409 {
	return &NullableError409{value: val, isSet: true}
}

func (v NullableError409) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableError409) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
