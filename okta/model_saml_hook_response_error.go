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

// checks if the SAMLHookResponseError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SAMLHookResponseError{}

// SAMLHookResponseError An object to return an error. Returning an error causes Okta to record a failure event in the Okta System Log. The string supplied in the `errorSummary` property is recorded in the System Log event. > **Note:** If the error object doesn't include the defined `errorSummary` property, the following common default message is returned to the end user: `The callback service returned an error`.  > **Note:** If a response to a SAML inline hook request isn't received from your external service within three seconds, a timeout occurs. In this scenario, the Okta SAML inline hook process continues, and the user is created.
type SAMLHookResponseError struct {
	// A human-readable summary of the error
	ErrorSummary         *string `json:"errorSummary,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SAMLHookResponseError SAMLHookResponseError

// NewSAMLHookResponseError instantiates a new SAMLHookResponseError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSAMLHookResponseError() *SAMLHookResponseError {
	this := SAMLHookResponseError{}
	return &this
}

// NewSAMLHookResponseErrorWithDefaults instantiates a new SAMLHookResponseError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSAMLHookResponseErrorWithDefaults() *SAMLHookResponseError {
	this := SAMLHookResponseError{}
	return &this
}

// GetErrorSummary returns the ErrorSummary field value if set, zero value otherwise.
func (o *SAMLHookResponseError) GetErrorSummary() string {
	if o == nil || IsNil(o.ErrorSummary) {
		var ret string
		return ret
	}
	return *o.ErrorSummary
}

// GetErrorSummaryOk returns a tuple with the ErrorSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLHookResponseError) GetErrorSummaryOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorSummary) {
		return nil, false
	}
	return o.ErrorSummary, true
}

// HasErrorSummary returns a boolean if a field has been set.
func (o *SAMLHookResponseError) HasErrorSummary() bool {
	if o != nil && !IsNil(o.ErrorSummary) {
		return true
	}

	return false
}

// SetErrorSummary gets a reference to the given string and assigns it to the ErrorSummary field.
func (o *SAMLHookResponseError) SetErrorSummary(v string) {
	o.ErrorSummary = &v
}

func (o SAMLHookResponseError) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SAMLHookResponseError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ErrorSummary) {
		toSerialize["errorSummary"] = o.ErrorSummary
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SAMLHookResponseError) UnmarshalJSON(data []byte) (err error) {
	varSAMLHookResponseError := _SAMLHookResponseError{}

	err = json.Unmarshal(data, &varSAMLHookResponseError)

	if err != nil {
		return err
	}

	*o = SAMLHookResponseError(varSAMLHookResponseError)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "errorSummary")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSAMLHookResponseError struct {
	value *SAMLHookResponseError
	isSet bool
}

func (v NullableSAMLHookResponseError) Get() *SAMLHookResponseError {
	return v.value
}

func (v *NullableSAMLHookResponseError) Set(val *SAMLHookResponseError) {
	v.value = val
	v.isSet = true
}

func (v NullableSAMLHookResponseError) IsSet() bool {
	return v.isSet
}

func (v *NullableSAMLHookResponseError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSAMLHookResponseError(val *SAMLHookResponseError) *NullableSAMLHookResponseError {
	return &NullableSAMLHookResponseError{value: val, isSet: true}
}

func (v NullableSAMLHookResponseError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSAMLHookResponseError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
