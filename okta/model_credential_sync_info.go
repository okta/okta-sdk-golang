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
	"time"
)

// checks if the CredentialSyncInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CredentialSyncInfo{}

// CredentialSyncInfo struct for CredentialSyncInfo
type CredentialSyncInfo struct {
	// The error code for the type of error
	ErrorCode *string `json:"errorCode,omitempty"`
	// A short description of the error
	ErrorReason *string `json:"errorReason,omitempty"`
	// The version ID of the password secret from the OPA vault.
	SecretVersionId *string `json:"secretVersionId,omitempty"`
	// Current credential sync status of the privileged resource
	SyncState *string `json:"syncState,omitempty"`
	// Timestamp when the credential was changed
	SyncTime             *time.Time `json:"syncTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CredentialSyncInfo CredentialSyncInfo

// NewCredentialSyncInfo instantiates a new CredentialSyncInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredentialSyncInfo() *CredentialSyncInfo {
	this := CredentialSyncInfo{}
	return &this
}

// NewCredentialSyncInfoWithDefaults instantiates a new CredentialSyncInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredentialSyncInfoWithDefaults() *CredentialSyncInfo {
	this := CredentialSyncInfo{}
	return &this
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *CredentialSyncInfo) GetErrorCode() string {
	if o == nil || IsNil(o.ErrorCode) {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialSyncInfo) GetErrorCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorCode) {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *CredentialSyncInfo) HasErrorCode() bool {
	if o != nil && !IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *CredentialSyncInfo) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetErrorReason returns the ErrorReason field value if set, zero value otherwise.
func (o *CredentialSyncInfo) GetErrorReason() string {
	if o == nil || IsNil(o.ErrorReason) {
		var ret string
		return ret
	}
	return *o.ErrorReason
}

// GetErrorReasonOk returns a tuple with the ErrorReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialSyncInfo) GetErrorReasonOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorReason) {
		return nil, false
	}
	return o.ErrorReason, true
}

// HasErrorReason returns a boolean if a field has been set.
func (o *CredentialSyncInfo) HasErrorReason() bool {
	if o != nil && !IsNil(o.ErrorReason) {
		return true
	}

	return false
}

// SetErrorReason gets a reference to the given string and assigns it to the ErrorReason field.
func (o *CredentialSyncInfo) SetErrorReason(v string) {
	o.ErrorReason = &v
}

// GetSecretVersionId returns the SecretVersionId field value if set, zero value otherwise.
func (o *CredentialSyncInfo) GetSecretVersionId() string {
	if o == nil || IsNil(o.SecretVersionId) {
		var ret string
		return ret
	}
	return *o.SecretVersionId
}

// GetSecretVersionIdOk returns a tuple with the SecretVersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialSyncInfo) GetSecretVersionIdOk() (*string, bool) {
	if o == nil || IsNil(o.SecretVersionId) {
		return nil, false
	}
	return o.SecretVersionId, true
}

// HasSecretVersionId returns a boolean if a field has been set.
func (o *CredentialSyncInfo) HasSecretVersionId() bool {
	if o != nil && !IsNil(o.SecretVersionId) {
		return true
	}

	return false
}

// SetSecretVersionId gets a reference to the given string and assigns it to the SecretVersionId field.
func (o *CredentialSyncInfo) SetSecretVersionId(v string) {
	o.SecretVersionId = &v
}

// GetSyncState returns the SyncState field value if set, zero value otherwise.
func (o *CredentialSyncInfo) GetSyncState() string {
	if o == nil || IsNil(o.SyncState) {
		var ret string
		return ret
	}
	return *o.SyncState
}

// GetSyncStateOk returns a tuple with the SyncState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialSyncInfo) GetSyncStateOk() (*string, bool) {
	if o == nil || IsNil(o.SyncState) {
		return nil, false
	}
	return o.SyncState, true
}

// HasSyncState returns a boolean if a field has been set.
func (o *CredentialSyncInfo) HasSyncState() bool {
	if o != nil && !IsNil(o.SyncState) {
		return true
	}

	return false
}

// SetSyncState gets a reference to the given string and assigns it to the SyncState field.
func (o *CredentialSyncInfo) SetSyncState(v string) {
	o.SyncState = &v
}

// GetSyncTime returns the SyncTime field value if set, zero value otherwise.
func (o *CredentialSyncInfo) GetSyncTime() time.Time {
	if o == nil || IsNil(o.SyncTime) {
		var ret time.Time
		return ret
	}
	return *o.SyncTime
}

// GetSyncTimeOk returns a tuple with the SyncTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialSyncInfo) GetSyncTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.SyncTime) {
		return nil, false
	}
	return o.SyncTime, true
}

// HasSyncTime returns a boolean if a field has been set.
func (o *CredentialSyncInfo) HasSyncTime() bool {
	if o != nil && !IsNil(o.SyncTime) {
		return true
	}

	return false
}

// SetSyncTime gets a reference to the given time.Time and assigns it to the SyncTime field.
func (o *CredentialSyncInfo) SetSyncTime(v time.Time) {
	o.SyncTime = &v
}

func (o CredentialSyncInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CredentialSyncInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ErrorCode) {
		toSerialize["errorCode"] = o.ErrorCode
	}
	if !IsNil(o.ErrorReason) {
		toSerialize["errorReason"] = o.ErrorReason
	}
	if !IsNil(o.SecretVersionId) {
		toSerialize["secretVersionId"] = o.SecretVersionId
	}
	if !IsNil(o.SyncState) {
		toSerialize["syncState"] = o.SyncState
	}
	if !IsNil(o.SyncTime) {
		toSerialize["syncTime"] = o.SyncTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CredentialSyncInfo) UnmarshalJSON(data []byte) (err error) {
	varCredentialSyncInfo := _CredentialSyncInfo{}

	err = json.Unmarshal(data, &varCredentialSyncInfo)

	if err != nil {
		return err
	}

	*o = CredentialSyncInfo(varCredentialSyncInfo)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "errorCode")
		delete(additionalProperties, "errorReason")
		delete(additionalProperties, "secretVersionId")
		delete(additionalProperties, "syncState")
		delete(additionalProperties, "syncTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCredentialSyncInfo struct {
	value *CredentialSyncInfo
	isSet bool
}

func (v NullableCredentialSyncInfo) Get() *CredentialSyncInfo {
	return v.value
}

func (v *NullableCredentialSyncInfo) Set(val *CredentialSyncInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableCredentialSyncInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableCredentialSyncInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredentialSyncInfo(val *CredentialSyncInfo) *NullableCredentialSyncInfo {
	return &NullableCredentialSyncInfo{value: val, isSet: true}
}

func (v NullableCredentialSyncInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredentialSyncInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
