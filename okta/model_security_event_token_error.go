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

// SecurityEventTokenError Error object thrown when parsing the Security Event Token
type SecurityEventTokenError struct {
	// Describes the error > **Note:** SET claim fields with underscores (snake case) are presented in camelcase. For example, `previous_status` appears as `previousStatus`. 
	Description *string `json:"description,omitempty"`
	// A code that describes the category of the error
	Err *string `json:"err,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SecurityEventTokenError SecurityEventTokenError

// NewSecurityEventTokenError instantiates a new SecurityEventTokenError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityEventTokenError() *SecurityEventTokenError {
	this := SecurityEventTokenError{}
	return &this
}

// NewSecurityEventTokenErrorWithDefaults instantiates a new SecurityEventTokenError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityEventTokenErrorWithDefaults() *SecurityEventTokenError {
	this := SecurityEventTokenError{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SecurityEventTokenError) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityEventTokenError) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SecurityEventTokenError) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SecurityEventTokenError) SetDescription(v string) {
	o.Description = &v
}

// GetErr returns the Err field value if set, zero value otherwise.
func (o *SecurityEventTokenError) GetErr() string {
	if o == nil || o.Err == nil {
		var ret string
		return ret
	}
	return *o.Err
}

// GetErrOk returns a tuple with the Err field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityEventTokenError) GetErrOk() (*string, bool) {
	if o == nil || o.Err == nil {
		return nil, false
	}
	return o.Err, true
}

// HasErr returns a boolean if a field has been set.
func (o *SecurityEventTokenError) HasErr() bool {
	if o != nil && o.Err != nil {
		return true
	}

	return false
}

// SetErr gets a reference to the given string and assigns it to the Err field.
func (o *SecurityEventTokenError) SetErr(v string) {
	o.Err = &v
}

func (o SecurityEventTokenError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Err != nil {
		toSerialize["err"] = o.Err
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SecurityEventTokenError) UnmarshalJSON(bytes []byte) (err error) {
	varSecurityEventTokenError := _SecurityEventTokenError{}

	err = json.Unmarshal(bytes, &varSecurityEventTokenError)
	if err == nil {
		*o = SecurityEventTokenError(varSecurityEventTokenError)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "err")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableSecurityEventTokenError struct {
	value *SecurityEventTokenError
	isSet bool
}

func (v NullableSecurityEventTokenError) Get() *SecurityEventTokenError {
	return v.value
}

func (v *NullableSecurityEventTokenError) Set(val *SecurityEventTokenError) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityEventTokenError) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityEventTokenError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityEventTokenError(val *SecurityEventTokenError) *NullableSecurityEventTokenError {
	return &NullableSecurityEventTokenError{value: val, isSet: true}
}

func (v NullableSecurityEventTokenError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityEventTokenError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

