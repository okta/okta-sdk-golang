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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the TokenHookResponseError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenHookResponseError{}

// TokenHookResponseError When an error object is returned, it causes Okta to return an OAuth 2.0 error to the requester of the token. In the error response, the value of `error` is `server_error`, and the value of `error_description` is the string that you supplied in the `errorSummary` property of the `error` object that you returned.
type TokenHookResponseError struct {
	// Human-readable summary of the error. If the error object doesn't include the `errorSummary` property defined, the following common default message is returned to the end user: `The callback service returned an error`.
	ErrorSummary         *string `json:"errorSummary,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TokenHookResponseError TokenHookResponseError

// NewTokenHookResponseError instantiates a new TokenHookResponseError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenHookResponseError() *TokenHookResponseError {
	this := TokenHookResponseError{}
	return &this
}

// NewTokenHookResponseErrorWithDefaults instantiates a new TokenHookResponseError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenHookResponseErrorWithDefaults() *TokenHookResponseError {
	this := TokenHookResponseError{}
	return &this
}

// GetErrorSummary returns the ErrorSummary field value if set, zero value otherwise.
func (o *TokenHookResponseError) GetErrorSummary() string {
	if o == nil || IsNil(o.ErrorSummary) {
		var ret string
		return ret
	}
	return *o.ErrorSummary
}

// GetErrorSummaryOk returns a tuple with the ErrorSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenHookResponseError) GetErrorSummaryOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorSummary) {
		return nil, false
	}
	return o.ErrorSummary, true
}

// HasErrorSummary returns a boolean if a field has been set.
func (o *TokenHookResponseError) HasErrorSummary() bool {
	if o != nil && !IsNil(o.ErrorSummary) {
		return true
	}

	return false
}

// SetErrorSummary gets a reference to the given string and assigns it to the ErrorSummary field.
func (o *TokenHookResponseError) SetErrorSummary(v string) {
	o.ErrorSummary = &v
}

func (o TokenHookResponseError) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenHookResponseError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ErrorSummary) {
		toSerialize["errorSummary"] = o.ErrorSummary
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TokenHookResponseError) UnmarshalJSON(data []byte) (err error) {
	varTokenHookResponseError := _TokenHookResponseError{}

	err = json.Unmarshal(data, &varTokenHookResponseError)

	if err != nil {
		return err
	}

	*o = TokenHookResponseError(varTokenHookResponseError)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "errorSummary")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTokenHookResponseError struct {
	value *TokenHookResponseError
	isSet bool
}

func (v NullableTokenHookResponseError) Get() *TokenHookResponseError {
	return v.value
}

func (v *NullableTokenHookResponseError) Set(val *TokenHookResponseError) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenHookResponseError) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenHookResponseError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenHookResponseError(val *TokenHookResponseError) *NullableTokenHookResponseError {
	return &NullableTokenHookResponseError{value: val, isSet: true}
}

func (v NullableTokenHookResponseError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenHookResponseError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
