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
	"fmt"
)

// checks if the RotatePasswordRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RotatePasswordRequest{}

// RotatePasswordRequest Rotate password request for the privileged resource
type RotatePasswordRequest struct {
	// The password associated with the privileged resource
	Password string `json:"password"`
	// The version ID of the password secret from the OPA vault
	SecretVersionId      string `json:"secretVersionId"`
	AdditionalProperties map[string]interface{}
}

type _RotatePasswordRequest RotatePasswordRequest

// NewRotatePasswordRequest instantiates a new RotatePasswordRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRotatePasswordRequest(password string, secretVersionId string) *RotatePasswordRequest {
	this := RotatePasswordRequest{}
	this.Password = password
	this.SecretVersionId = secretVersionId
	return &this
}

// NewRotatePasswordRequestWithDefaults instantiates a new RotatePasswordRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRotatePasswordRequestWithDefaults() *RotatePasswordRequest {
	this := RotatePasswordRequest{}
	return &this
}

// GetPassword returns the Password field value
func (o *RotatePasswordRequest) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *RotatePasswordRequest) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *RotatePasswordRequest) SetPassword(v string) {
	o.Password = v
}

// GetSecretVersionId returns the SecretVersionId field value
func (o *RotatePasswordRequest) GetSecretVersionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecretVersionId
}

// GetSecretVersionIdOk returns a tuple with the SecretVersionId field value
// and a boolean to check if the value has been set.
func (o *RotatePasswordRequest) GetSecretVersionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecretVersionId, true
}

// SetSecretVersionId sets field value
func (o *RotatePasswordRequest) SetSecretVersionId(v string) {
	o.SecretVersionId = v
}

func (o RotatePasswordRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RotatePasswordRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["password"] = o.Password
	toSerialize["secretVersionId"] = o.SecretVersionId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RotatePasswordRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"password",
		"secretVersionId",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varRotatePasswordRequest := _RotatePasswordRequest{}

	err = json.Unmarshal(data, &varRotatePasswordRequest)

	if err != nil {
		return err
	}

	*o = RotatePasswordRequest(varRotatePasswordRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "password")
		delete(additionalProperties, "secretVersionId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRotatePasswordRequest struct {
	value *RotatePasswordRequest
	isSet bool
}

func (v NullableRotatePasswordRequest) Get() *RotatePasswordRequest {
	return v.value
}

func (v *NullableRotatePasswordRequest) Set(val *RotatePasswordRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRotatePasswordRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRotatePasswordRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRotatePasswordRequest(val *RotatePasswordRequest) *NullableRotatePasswordRequest {
	return &NullableRotatePasswordRequest{value: val, isSet: true}
}

func (v NullableRotatePasswordRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRotatePasswordRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
