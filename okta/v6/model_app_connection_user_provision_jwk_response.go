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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"fmt"
)

// AppConnectionUserProvisionJWKResponse struct for AppConnectionUserProvisionJWKResponse
type AppConnectionUserProvisionJWKResponse struct {
	Jwks AppConnectionUserProvisionJWKList `json:"jwks"`
	AdditionalProperties map[string]interface{}
}

type _AppConnectionUserProvisionJWKResponse AppConnectionUserProvisionJWKResponse

// NewAppConnectionUserProvisionJWKResponse instantiates a new AppConnectionUserProvisionJWKResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppConnectionUserProvisionJWKResponse(jwks AppConnectionUserProvisionJWKList) *AppConnectionUserProvisionJWKResponse {
	this := AppConnectionUserProvisionJWKResponse{}
	this.Jwks = jwks
	return &this
}

// NewAppConnectionUserProvisionJWKResponseWithDefaults instantiates a new AppConnectionUserProvisionJWKResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppConnectionUserProvisionJWKResponseWithDefaults() *AppConnectionUserProvisionJWKResponse {
	this := AppConnectionUserProvisionJWKResponse{}
	return &this
}

// GetJwks returns the Jwks field value
func (o *AppConnectionUserProvisionJWKResponse) GetJwks() AppConnectionUserProvisionJWKList {
	if o == nil {
		var ret AppConnectionUserProvisionJWKList
		return ret
	}

	return o.Jwks
}

// GetJwksOk returns a tuple with the Jwks field value
// and a boolean to check if the value has been set.
func (o *AppConnectionUserProvisionJWKResponse) GetJwksOk() (*AppConnectionUserProvisionJWKList, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Jwks, true
}

// SetJwks sets field value
func (o *AppConnectionUserProvisionJWKResponse) SetJwks(v AppConnectionUserProvisionJWKList) {
	o.Jwks = v
}

func (o AppConnectionUserProvisionJWKResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["jwks"] = o.Jwks
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AppConnectionUserProvisionJWKResponse) UnmarshalJSON(bytes []byte) (err error) {
	varAppConnectionUserProvisionJWKResponse := _AppConnectionUserProvisionJWKResponse{}

	err = json.Unmarshal(bytes, &varAppConnectionUserProvisionJWKResponse)
	if err == nil {
		*o = AppConnectionUserProvisionJWKResponse(varAppConnectionUserProvisionJWKResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "jwks")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableAppConnectionUserProvisionJWKResponse struct {
	value *AppConnectionUserProvisionJWKResponse
	isSet bool
}

func (v NullableAppConnectionUserProvisionJWKResponse) Get() *AppConnectionUserProvisionJWKResponse {
	return v.value
}

func (v *NullableAppConnectionUserProvisionJWKResponse) Set(val *AppConnectionUserProvisionJWKResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAppConnectionUserProvisionJWKResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAppConnectionUserProvisionJWKResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppConnectionUserProvisionJWKResponse(val *AppConnectionUserProvisionJWKResponse) *NullableAppConnectionUserProvisionJWKResponse {
	return &NullableAppConnectionUserProvisionJWKResponse{value: val, isSet: true}
}

func (v NullableAppConnectionUserProvisionJWKResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppConnectionUserProvisionJWKResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

