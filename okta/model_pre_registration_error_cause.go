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

// checks if the PreRegistrationErrorCause type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PreRegistrationErrorCause{}

// PreRegistrationErrorCause A single per-entry validation error, extending the standard Okta `ErrorCause` with entry-identifying fields.
type PreRegistrationErrorCause struct {
	ErrorSummary *string `json:"errorSummary,omitempty"`
	// An Okta error code specific to this entry's failure
	ErrorCode *string `json:"errorCode,omitempty"`
	// Zero-based index of the failing entry in the request array. Useful for disambiguation when the same `serialNumber` appears more than once.
	Index *int32 `json:"index,omitempty"`
	// Device platform for pre-registration
	Platform *string `json:"platform,omitempty"`
	// Serial number of the failing entry
	SerialNumber         *string `json:"serialNumber,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PreRegistrationErrorCause PreRegistrationErrorCause

// NewPreRegistrationErrorCause instantiates a new PreRegistrationErrorCause object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPreRegistrationErrorCause() *PreRegistrationErrorCause {
	this := PreRegistrationErrorCause{}
	return &this
}

// NewPreRegistrationErrorCauseWithDefaults instantiates a new PreRegistrationErrorCause object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPreRegistrationErrorCauseWithDefaults() *PreRegistrationErrorCause {
	this := PreRegistrationErrorCause{}
	return &this
}

// GetErrorSummary returns the ErrorSummary field value if set, zero value otherwise.
func (o *PreRegistrationErrorCause) GetErrorSummary() string {
	if o == nil || IsNil(o.ErrorSummary) {
		var ret string
		return ret
	}
	return *o.ErrorSummary
}

// GetErrorSummaryOk returns a tuple with the ErrorSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreRegistrationErrorCause) GetErrorSummaryOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorSummary) {
		return nil, false
	}
	return o.ErrorSummary, true
}

// HasErrorSummary returns a boolean if a field has been set.
func (o *PreRegistrationErrorCause) HasErrorSummary() bool {
	if o != nil && !IsNil(o.ErrorSummary) {
		return true
	}

	return false
}

// SetErrorSummary gets a reference to the given string and assigns it to the ErrorSummary field.
func (o *PreRegistrationErrorCause) SetErrorSummary(v string) {
	o.ErrorSummary = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *PreRegistrationErrorCause) GetErrorCode() string {
	if o == nil || IsNil(o.ErrorCode) {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreRegistrationErrorCause) GetErrorCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorCode) {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *PreRegistrationErrorCause) HasErrorCode() bool {
	if o != nil && !IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *PreRegistrationErrorCause) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *PreRegistrationErrorCause) GetIndex() int32 {
	if o == nil || IsNil(o.Index) {
		var ret int32
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreRegistrationErrorCause) GetIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.Index) {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *PreRegistrationErrorCause) HasIndex() bool {
	if o != nil && !IsNil(o.Index) {
		return true
	}

	return false
}

// SetIndex gets a reference to the given int32 and assigns it to the Index field.
func (o *PreRegistrationErrorCause) SetIndex(v int32) {
	o.Index = &v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *PreRegistrationErrorCause) GetPlatform() string {
	if o == nil || IsNil(o.Platform) {
		var ret string
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreRegistrationErrorCause) GetPlatformOk() (*string, bool) {
	if o == nil || IsNil(o.Platform) {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *PreRegistrationErrorCause) HasPlatform() bool {
	if o != nil && !IsNil(o.Platform) {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given string and assigns it to the Platform field.
func (o *PreRegistrationErrorCause) SetPlatform(v string) {
	o.Platform = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *PreRegistrationErrorCause) GetSerialNumber() string {
	if o == nil || IsNil(o.SerialNumber) {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreRegistrationErrorCause) GetSerialNumberOk() (*string, bool) {
	if o == nil || IsNil(o.SerialNumber) {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *PreRegistrationErrorCause) HasSerialNumber() bool {
	if o != nil && !IsNil(o.SerialNumber) {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *PreRegistrationErrorCause) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

func (o PreRegistrationErrorCause) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PreRegistrationErrorCause) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ErrorSummary) {
		toSerialize["errorSummary"] = o.ErrorSummary
	}
	if !IsNil(o.ErrorCode) {
		toSerialize["errorCode"] = o.ErrorCode
	}
	if !IsNil(o.Index) {
		toSerialize["index"] = o.Index
	}
	if !IsNil(o.Platform) {
		toSerialize["platform"] = o.Platform
	}
	if !IsNil(o.SerialNumber) {
		toSerialize["serialNumber"] = o.SerialNumber
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PreRegistrationErrorCause) UnmarshalJSON(data []byte) (err error) {
	varPreRegistrationErrorCause := _PreRegistrationErrorCause{}

	err = json.Unmarshal(data, &varPreRegistrationErrorCause)

	if err != nil {
		return err
	}

	*o = PreRegistrationErrorCause(varPreRegistrationErrorCause)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "errorSummary")
		delete(additionalProperties, "errorCode")
		delete(additionalProperties, "index")
		delete(additionalProperties, "platform")
		delete(additionalProperties, "serialNumber")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePreRegistrationErrorCause struct {
	value *PreRegistrationErrorCause
	isSet bool
}

func (v NullablePreRegistrationErrorCause) Get() *PreRegistrationErrorCause {
	return v.value
}

func (v *NullablePreRegistrationErrorCause) Set(val *PreRegistrationErrorCause) {
	v.value = val
	v.isSet = true
}

func (v NullablePreRegistrationErrorCause) IsSet() bool {
	return v.isSet
}

func (v *NullablePreRegistrationErrorCause) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePreRegistrationErrorCause(val *PreRegistrationErrorCause) *NullablePreRegistrationErrorCause {
	return &NullablePreRegistrationErrorCause{value: val, isSet: true}
}

func (v NullablePreRegistrationErrorCause) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePreRegistrationErrorCause) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
