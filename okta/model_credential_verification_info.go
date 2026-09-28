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
	"time"
)

// checks if the CredentialVerificationInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CredentialVerificationInfo{}

// CredentialVerificationInfo Credential verification information
type CredentialVerificationInfo struct {
	// Timestamp that indicates when the latest verification job reached a terminal state (`COMPLETED` or `FAILED`). Absent while `verificationState` is set to `NOT_INITIALIZED` or `IN_PROGRESS`.
	Completed *time.Time `json:"completed,omitempty"`
	// Indicates whether the credentials were verified successfully. `true` if the credentials are valid against the downstream application and `false` if they are invalid. Present only when `verificationState` is `COMPLETED`.
	CredentialValid NullableBool `json:"credentialValid,omitempty"`
	// Indicates why credential validation didn't succeed. Present when `verificationState` is `FAILED`, or when `verificationState` is `COMPLETED` and `credentialValid` is `false`.
	ErrorCode NullableString `json:"errorCode,omitempty"`
	// Description of the error
	ErrorReason NullableString `json:"errorReason,omitempty"`
	// Unique identifier of the most recent verification job initiated for this privileged resource
	RequestId NullableString `json:"requestId,omitempty"`
	// Current state of the latest credential verification job for the privileged resource
	VerificationState    *string `json:"verificationState,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CredentialVerificationInfo CredentialVerificationInfo

// NewCredentialVerificationInfo instantiates a new CredentialVerificationInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredentialVerificationInfo() *CredentialVerificationInfo {
	this := CredentialVerificationInfo{}
	return &this
}

// NewCredentialVerificationInfoWithDefaults instantiates a new CredentialVerificationInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredentialVerificationInfoWithDefaults() *CredentialVerificationInfo {
	this := CredentialVerificationInfo{}
	return &this
}

// GetCompleted returns the Completed field value if set, zero value otherwise.
func (o *CredentialVerificationInfo) GetCompleted() time.Time {
	if o == nil || IsNil(o.Completed) {
		var ret time.Time
		return ret
	}
	return *o.Completed
}

// GetCompletedOk returns a tuple with the Completed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialVerificationInfo) GetCompletedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Completed) {
		return nil, false
	}
	return o.Completed, true
}

// HasCompleted returns a boolean if a field has been set.
func (o *CredentialVerificationInfo) HasCompleted() bool {
	if o != nil && !IsNil(o.Completed) {
		return true
	}

	return false
}

// SetCompleted gets a reference to the given time.Time and assigns it to the Completed field.
func (o *CredentialVerificationInfo) SetCompleted(v time.Time) {
	o.Completed = &v
}

// GetCredentialValid returns the CredentialValid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CredentialVerificationInfo) GetCredentialValid() bool {
	if o == nil || IsNil(o.CredentialValid.Get()) {
		var ret bool
		return ret
	}
	return *o.CredentialValid.Get()
}

// GetCredentialValidOk returns a tuple with the CredentialValid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CredentialVerificationInfo) GetCredentialValidOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.CredentialValid.Get(), o.CredentialValid.IsSet()
}

// HasCredentialValid returns a boolean if a field has been set.
func (o *CredentialVerificationInfo) HasCredentialValid() bool {
	if o != nil && o.CredentialValid.IsSet() {
		return true
	}

	return false
}

// SetCredentialValid gets a reference to the given NullableBool and assigns it to the CredentialValid field.
func (o *CredentialVerificationInfo) SetCredentialValid(v bool) {
	o.CredentialValid.Set(&v)
}

// SetCredentialValidNil sets the value for CredentialValid to be an explicit nil
func (o *CredentialVerificationInfo) SetCredentialValidNil() {
	o.CredentialValid.Set(nil)
}

// UnsetCredentialValid ensures that no value is present for CredentialValid, not even an explicit nil
func (o *CredentialVerificationInfo) UnsetCredentialValid() {
	o.CredentialValid.Unset()
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CredentialVerificationInfo) GetErrorCode() string {
	if o == nil || IsNil(o.ErrorCode.Get()) {
		var ret string
		return ret
	}
	return *o.ErrorCode.Get()
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CredentialVerificationInfo) GetErrorCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorCode.Get(), o.ErrorCode.IsSet()
}

// HasErrorCode returns a boolean if a field has been set.
func (o *CredentialVerificationInfo) HasErrorCode() bool {
	if o != nil && o.ErrorCode.IsSet() {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given NullableString and assigns it to the ErrorCode field.
func (o *CredentialVerificationInfo) SetErrorCode(v string) {
	o.ErrorCode.Set(&v)
}

// SetErrorCodeNil sets the value for ErrorCode to be an explicit nil
func (o *CredentialVerificationInfo) SetErrorCodeNil() {
	o.ErrorCode.Set(nil)
}

// UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
func (o *CredentialVerificationInfo) UnsetErrorCode() {
	o.ErrorCode.Unset()
}

// GetErrorReason returns the ErrorReason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CredentialVerificationInfo) GetErrorReason() string {
	if o == nil || IsNil(o.ErrorReason.Get()) {
		var ret string
		return ret
	}
	return *o.ErrorReason.Get()
}

// GetErrorReasonOk returns a tuple with the ErrorReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CredentialVerificationInfo) GetErrorReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorReason.Get(), o.ErrorReason.IsSet()
}

// HasErrorReason returns a boolean if a field has been set.
func (o *CredentialVerificationInfo) HasErrorReason() bool {
	if o != nil && o.ErrorReason.IsSet() {
		return true
	}

	return false
}

// SetErrorReason gets a reference to the given NullableString and assigns it to the ErrorReason field.
func (o *CredentialVerificationInfo) SetErrorReason(v string) {
	o.ErrorReason.Set(&v)
}

// SetErrorReasonNil sets the value for ErrorReason to be an explicit nil
func (o *CredentialVerificationInfo) SetErrorReasonNil() {
	o.ErrorReason.Set(nil)
}

// UnsetErrorReason ensures that no value is present for ErrorReason, not even an explicit nil
func (o *CredentialVerificationInfo) UnsetErrorReason() {
	o.ErrorReason.Unset()
}

// GetRequestId returns the RequestId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CredentialVerificationInfo) GetRequestId() string {
	if o == nil || IsNil(o.RequestId.Get()) {
		var ret string
		return ret
	}
	return *o.RequestId.Get()
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CredentialVerificationInfo) GetRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestId.Get(), o.RequestId.IsSet()
}

// HasRequestId returns a boolean if a field has been set.
func (o *CredentialVerificationInfo) HasRequestId() bool {
	if o != nil && o.RequestId.IsSet() {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given NullableString and assigns it to the RequestId field.
func (o *CredentialVerificationInfo) SetRequestId(v string) {
	o.RequestId.Set(&v)
}

// SetRequestIdNil sets the value for RequestId to be an explicit nil
func (o *CredentialVerificationInfo) SetRequestIdNil() {
	o.RequestId.Set(nil)
}

// UnsetRequestId ensures that no value is present for RequestId, not even an explicit nil
func (o *CredentialVerificationInfo) UnsetRequestId() {
	o.RequestId.Unset()
}

// GetVerificationState returns the VerificationState field value if set, zero value otherwise.
func (o *CredentialVerificationInfo) GetVerificationState() string {
	if o == nil || IsNil(o.VerificationState) {
		var ret string
		return ret
	}
	return *o.VerificationState
}

// GetVerificationStateOk returns a tuple with the VerificationState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialVerificationInfo) GetVerificationStateOk() (*string, bool) {
	if o == nil || IsNil(o.VerificationState) {
		return nil, false
	}
	return o.VerificationState, true
}

// HasVerificationState returns a boolean if a field has been set.
func (o *CredentialVerificationInfo) HasVerificationState() bool {
	if o != nil && !IsNil(o.VerificationState) {
		return true
	}

	return false
}

// SetVerificationState gets a reference to the given string and assigns it to the VerificationState field.
func (o *CredentialVerificationInfo) SetVerificationState(v string) {
	o.VerificationState = &v
}

func (o CredentialVerificationInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CredentialVerificationInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Completed) {
		toSerialize["completed"] = o.Completed
	}
	if o.CredentialValid.IsSet() {
		toSerialize["credentialValid"] = o.CredentialValid.Get()
	}
	if o.ErrorCode.IsSet() {
		toSerialize["errorCode"] = o.ErrorCode.Get()
	}
	if o.ErrorReason.IsSet() {
		toSerialize["errorReason"] = o.ErrorReason.Get()
	}
	if o.RequestId.IsSet() {
		toSerialize["requestId"] = o.RequestId.Get()
	}
	if !IsNil(o.VerificationState) {
		toSerialize["verificationState"] = o.VerificationState
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CredentialVerificationInfo) UnmarshalJSON(data []byte) (err error) {
	varCredentialVerificationInfo := _CredentialVerificationInfo{}

	err = json.Unmarshal(data, &varCredentialVerificationInfo)

	if err != nil {
		return err
	}

	*o = CredentialVerificationInfo(varCredentialVerificationInfo)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "completed")
		delete(additionalProperties, "credentialValid")
		delete(additionalProperties, "errorCode")
		delete(additionalProperties, "errorReason")
		delete(additionalProperties, "requestId")
		delete(additionalProperties, "verificationState")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCredentialVerificationInfo struct {
	value *CredentialVerificationInfo
	isSet bool
}

func (v NullableCredentialVerificationInfo) Get() *CredentialVerificationInfo {
	return v.value
}

func (v *NullableCredentialVerificationInfo) Set(val *CredentialVerificationInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableCredentialVerificationInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableCredentialVerificationInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredentialVerificationInfo(val *CredentialVerificationInfo) *NullableCredentialVerificationInfo {
	return &NullableCredentialVerificationInfo{value: val, isSet: true}
}

func (v NullableCredentialVerificationInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredentialVerificationInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
