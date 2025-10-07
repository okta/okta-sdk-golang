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

// checks if the UploadYubikeyOtpTokenSeedRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UploadYubikeyOtpTokenSeedRequest{}

// UploadYubikeyOtpTokenSeedRequest struct for UploadYubikeyOtpTokenSeedRequest
type UploadYubikeyOtpTokenSeedRequest struct {
	// The unique identifier assigned to each YubiKey device
	SerialNumber *string `json:"serialNumber,omitempty"`
	// The YubiKey's public ID
	PublicId *string `json:"publicId,omitempty"`
	// The YubiKey's private ID
	PrivateId *string `json:"privateId,omitempty"`
	// The cryptographic key used in the AES (Advanced Encryption Standard) algorithm to encrypt and decrypt the YubiKey OTP
	AesKey               *string `json:"aesKey,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UploadYubikeyOtpTokenSeedRequest UploadYubikeyOtpTokenSeedRequest

// NewUploadYubikeyOtpTokenSeedRequest instantiates a new UploadYubikeyOtpTokenSeedRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUploadYubikeyOtpTokenSeedRequest() *UploadYubikeyOtpTokenSeedRequest {
	this := UploadYubikeyOtpTokenSeedRequest{}
	return &this
}

// NewUploadYubikeyOtpTokenSeedRequestWithDefaults instantiates a new UploadYubikeyOtpTokenSeedRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUploadYubikeyOtpTokenSeedRequestWithDefaults() *UploadYubikeyOtpTokenSeedRequest {
	this := UploadYubikeyOtpTokenSeedRequest{}
	return &this
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *UploadYubikeyOtpTokenSeedRequest) GetSerialNumber() string {
	if o == nil || IsNil(o.SerialNumber) {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadYubikeyOtpTokenSeedRequest) GetSerialNumberOk() (*string, bool) {
	if o == nil || IsNil(o.SerialNumber) {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *UploadYubikeyOtpTokenSeedRequest) HasSerialNumber() bool {
	if o != nil && !IsNil(o.SerialNumber) {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *UploadYubikeyOtpTokenSeedRequest) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetPublicId returns the PublicId field value if set, zero value otherwise.
func (o *UploadYubikeyOtpTokenSeedRequest) GetPublicId() string {
	if o == nil || IsNil(o.PublicId) {
		var ret string
		return ret
	}
	return *o.PublicId
}

// GetPublicIdOk returns a tuple with the PublicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadYubikeyOtpTokenSeedRequest) GetPublicIdOk() (*string, bool) {
	if o == nil || IsNil(o.PublicId) {
		return nil, false
	}
	return o.PublicId, true
}

// HasPublicId returns a boolean if a field has been set.
func (o *UploadYubikeyOtpTokenSeedRequest) HasPublicId() bool {
	if o != nil && !IsNil(o.PublicId) {
		return true
	}

	return false
}

// SetPublicId gets a reference to the given string and assigns it to the PublicId field.
func (o *UploadYubikeyOtpTokenSeedRequest) SetPublicId(v string) {
	o.PublicId = &v
}

// GetPrivateId returns the PrivateId field value if set, zero value otherwise.
func (o *UploadYubikeyOtpTokenSeedRequest) GetPrivateId() string {
	if o == nil || IsNil(o.PrivateId) {
		var ret string
		return ret
	}
	return *o.PrivateId
}

// GetPrivateIdOk returns a tuple with the PrivateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadYubikeyOtpTokenSeedRequest) GetPrivateIdOk() (*string, bool) {
	if o == nil || IsNil(o.PrivateId) {
		return nil, false
	}
	return o.PrivateId, true
}

// HasPrivateId returns a boolean if a field has been set.
func (o *UploadYubikeyOtpTokenSeedRequest) HasPrivateId() bool {
	if o != nil && !IsNil(o.PrivateId) {
		return true
	}

	return false
}

// SetPrivateId gets a reference to the given string and assigns it to the PrivateId field.
func (o *UploadYubikeyOtpTokenSeedRequest) SetPrivateId(v string) {
	o.PrivateId = &v
}

// GetAesKey returns the AesKey field value if set, zero value otherwise.
func (o *UploadYubikeyOtpTokenSeedRequest) GetAesKey() string {
	if o == nil || IsNil(o.AesKey) {
		var ret string
		return ret
	}
	return *o.AesKey
}

// GetAesKeyOk returns a tuple with the AesKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadYubikeyOtpTokenSeedRequest) GetAesKeyOk() (*string, bool) {
	if o == nil || IsNil(o.AesKey) {
		return nil, false
	}
	return o.AesKey, true
}

// HasAesKey returns a boolean if a field has been set.
func (o *UploadYubikeyOtpTokenSeedRequest) HasAesKey() bool {
	if o != nil && !IsNil(o.AesKey) {
		return true
	}

	return false
}

// SetAesKey gets a reference to the given string and assigns it to the AesKey field.
func (o *UploadYubikeyOtpTokenSeedRequest) SetAesKey(v string) {
	o.AesKey = &v
}

func (o UploadYubikeyOtpTokenSeedRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UploadYubikeyOtpTokenSeedRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SerialNumber) {
		toSerialize["serialNumber"] = o.SerialNumber
	}
	if !IsNil(o.PublicId) {
		toSerialize["publicId"] = o.PublicId
	}
	if !IsNil(o.PrivateId) {
		toSerialize["privateId"] = o.PrivateId
	}
	if !IsNil(o.AesKey) {
		toSerialize["aesKey"] = o.AesKey
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UploadYubikeyOtpTokenSeedRequest) UnmarshalJSON(data []byte) (err error) {
	varUploadYubikeyOtpTokenSeedRequest := _UploadYubikeyOtpTokenSeedRequest{}

	err = json.Unmarshal(data, &varUploadYubikeyOtpTokenSeedRequest)

	if err != nil {
		return err
	}

	*o = UploadYubikeyOtpTokenSeedRequest(varUploadYubikeyOtpTokenSeedRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "serialNumber")
		delete(additionalProperties, "publicId")
		delete(additionalProperties, "privateId")
		delete(additionalProperties, "aesKey")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUploadYubikeyOtpTokenSeedRequest struct {
	value *UploadYubikeyOtpTokenSeedRequest
	isSet bool
}

func (v NullableUploadYubikeyOtpTokenSeedRequest) Get() *UploadYubikeyOtpTokenSeedRequest {
	return v.value
}

func (v *NullableUploadYubikeyOtpTokenSeedRequest) Set(val *UploadYubikeyOtpTokenSeedRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUploadYubikeyOtpTokenSeedRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUploadYubikeyOtpTokenSeedRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUploadYubikeyOtpTokenSeedRequest(val *UploadYubikeyOtpTokenSeedRequest) *NullableUploadYubikeyOtpTokenSeedRequest {
	return &NullableUploadYubikeyOtpTokenSeedRequest{value: val, isSet: true}
}

func (v NullableUploadYubikeyOtpTokenSeedRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUploadYubikeyOtpTokenSeedRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
