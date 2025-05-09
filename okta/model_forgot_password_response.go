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

// ForgotPasswordResponse struct for ForgotPasswordResponse
type ForgotPasswordResponse struct {
	ResetPasswordUrl *string `json:"resetPasswordUrl,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ForgotPasswordResponse ForgotPasswordResponse

// NewForgotPasswordResponse instantiates a new ForgotPasswordResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewForgotPasswordResponse() *ForgotPasswordResponse {
	this := ForgotPasswordResponse{}
	return &this
}

// NewForgotPasswordResponseWithDefaults instantiates a new ForgotPasswordResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewForgotPasswordResponseWithDefaults() *ForgotPasswordResponse {
	this := ForgotPasswordResponse{}
	return &this
}

// GetResetPasswordUrl returns the ResetPasswordUrl field value if set, zero value otherwise.
func (o *ForgotPasswordResponse) GetResetPasswordUrl() string {
	if o == nil || o.ResetPasswordUrl == nil {
		var ret string
		return ret
	}
	return *o.ResetPasswordUrl
}

// GetResetPasswordUrlOk returns a tuple with the ResetPasswordUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForgotPasswordResponse) GetResetPasswordUrlOk() (*string, bool) {
	if o == nil || o.ResetPasswordUrl == nil {
		return nil, false
	}
	return o.ResetPasswordUrl, true
}

// HasResetPasswordUrl returns a boolean if a field has been set.
func (o *ForgotPasswordResponse) HasResetPasswordUrl() bool {
	if o != nil && o.ResetPasswordUrl != nil {
		return true
	}

	return false
}

// SetResetPasswordUrl gets a reference to the given string and assigns it to the ResetPasswordUrl field.
func (o *ForgotPasswordResponse) SetResetPasswordUrl(v string) {
	o.ResetPasswordUrl = &v
}

func (o ForgotPasswordResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResetPasswordUrl != nil {
		toSerialize["resetPasswordUrl"] = o.ResetPasswordUrl
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ForgotPasswordResponse) UnmarshalJSON(bytes []byte) (err error) {
	varForgotPasswordResponse := _ForgotPasswordResponse{}

	err = json.Unmarshal(bytes, &varForgotPasswordResponse)
	if err == nil {
		*o = ForgotPasswordResponse(varForgotPasswordResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "resetPasswordUrl")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableForgotPasswordResponse struct {
	value *ForgotPasswordResponse
	isSet bool
}

func (v NullableForgotPasswordResponse) Get() *ForgotPasswordResponse {
	return v.value
}

func (v *NullableForgotPasswordResponse) Set(val *ForgotPasswordResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableForgotPasswordResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableForgotPasswordResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableForgotPasswordResponse(val *ForgotPasswordResponse) *NullableForgotPasswordResponse {
	return &NullableForgotPasswordResponse{value: val, isSet: true}
}

func (v NullableForgotPasswordResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableForgotPasswordResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

