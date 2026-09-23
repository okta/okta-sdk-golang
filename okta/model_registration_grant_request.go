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
	"fmt"
)

// checks if the RegistrationGrantRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegistrationGrantRequest{}

// RegistrationGrantRequest struct for RegistrationGrantRequest
type RegistrationGrantRequest struct {
	// Type of registration grant
	GrantType string `json:"grantType"`
	// ID of the registration secret (from `POST /device-identity/api/v1/registration-secrets`)
	RegistrationSecretId string `json:"registrationSecretId"`
	AdditionalProperties map[string]interface{}
}

type _RegistrationGrantRequest RegistrationGrantRequest

// NewRegistrationGrantRequest instantiates a new RegistrationGrantRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistrationGrantRequest(grantType string, registrationSecretId string) *RegistrationGrantRequest {
	this := RegistrationGrantRequest{}
	this.GrantType = grantType
	this.RegistrationSecretId = registrationSecretId
	return &this
}

// NewRegistrationGrantRequestWithDefaults instantiates a new RegistrationGrantRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistrationGrantRequestWithDefaults() *RegistrationGrantRequest {
	this := RegistrationGrantRequest{}
	return &this
}

// GetGrantType returns the GrantType field value
func (o *RegistrationGrantRequest) GetGrantType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GrantType
}

// GetGrantTypeOk returns a tuple with the GrantType field value
// and a boolean to check if the value has been set.
func (o *RegistrationGrantRequest) GetGrantTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GrantType, true
}

// SetGrantType sets field value
func (o *RegistrationGrantRequest) SetGrantType(v string) {
	o.GrantType = v
}

// GetRegistrationSecretId returns the RegistrationSecretId field value
func (o *RegistrationGrantRequest) GetRegistrationSecretId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RegistrationSecretId
}

// GetRegistrationSecretIdOk returns a tuple with the RegistrationSecretId field value
// and a boolean to check if the value has been set.
func (o *RegistrationGrantRequest) GetRegistrationSecretIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RegistrationSecretId, true
}

// SetRegistrationSecretId sets field value
func (o *RegistrationGrantRequest) SetRegistrationSecretId(v string) {
	o.RegistrationSecretId = v
}

func (o RegistrationGrantRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegistrationGrantRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["grantType"] = o.GrantType
	toSerialize["registrationSecretId"] = o.RegistrationSecretId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RegistrationGrantRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"grantType",
		"registrationSecretId",
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

	varRegistrationGrantRequest := _RegistrationGrantRequest{}

	err = json.Unmarshal(data, &varRegistrationGrantRequest)

	if err != nil {
		return err
	}

	*o = RegistrationGrantRequest(varRegistrationGrantRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "grantType")
		delete(additionalProperties, "registrationSecretId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRegistrationGrantRequest struct {
	value *RegistrationGrantRequest
	isSet bool
}

func (v NullableRegistrationGrantRequest) Get() *RegistrationGrantRequest {
	return v.value
}

func (v *NullableRegistrationGrantRequest) Set(val *RegistrationGrantRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistrationGrantRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistrationGrantRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistrationGrantRequest(val *RegistrationGrantRequest) *NullableRegistrationGrantRequest {
	return &NullableRegistrationGrantRequest{value: val, isSet: true}
}

func (v NullableRegistrationGrantRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistrationGrantRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
