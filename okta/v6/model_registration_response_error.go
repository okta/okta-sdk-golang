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
)

// RegistrationResponseError For the registration inline hook, the `error` object provides a way of displaying an error message to the end user who is trying to register or update their profile.  * If you're using the Okta Sign-In Widget for Profile Enrollment, only the `errorSummary` messages of the `errorCauses` objects that your external service returns appear as inline errors, given the following:   * You don't customize the error handling behavior of the widget.   * The `location` of `errorSummary` in the `errorCauses` object specifies the request object's user profile attribute. * If you don't return a value for the `errorCauses` object, and deny the user's registration attempt through the `commands` object in your response to Okta, one of the following generic messages appears to the end user:   * \"Registration cannot be completed at this time.\" (SSR)   * \"We found some errors. Please review the form and make corrections.\" (Progressive Enrollment) * If you don't return an `error` object at all and the registration is denied, the following generic message appears to the end user:   * \"Registration denied.\" (SSR)   * \"Profile update denied.\" (Progressive Enrollment)  >**Note:** If you include an error object in your response, no commands are executed and the registration fails. This holds true even if the top-level `errorSummary` and the `errorCauses` objects are omitted.
type RegistrationResponseError struct {
	// Human-readable summary of one or more errors
	ErrorSummary *string `json:"errorSummary,omitempty"`
	ErrorCauses []RegistrationResponseErrorErrorCausesInner `json:"errorCauses,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RegistrationResponseError RegistrationResponseError

// NewRegistrationResponseError instantiates a new RegistrationResponseError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistrationResponseError() *RegistrationResponseError {
	this := RegistrationResponseError{}
	return &this
}

// NewRegistrationResponseErrorWithDefaults instantiates a new RegistrationResponseError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistrationResponseErrorWithDefaults() *RegistrationResponseError {
	this := RegistrationResponseError{}
	return &this
}

// GetErrorSummary returns the ErrorSummary field value if set, zero value otherwise.
func (o *RegistrationResponseError) GetErrorSummary() string {
	if o == nil || o.ErrorSummary == nil {
		var ret string
		return ret
	}
	return *o.ErrorSummary
}

// GetErrorSummaryOk returns a tuple with the ErrorSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationResponseError) GetErrorSummaryOk() (*string, bool) {
	if o == nil || o.ErrorSummary == nil {
		return nil, false
	}
	return o.ErrorSummary, true
}

// HasErrorSummary returns a boolean if a field has been set.
func (o *RegistrationResponseError) HasErrorSummary() bool {
	if o != nil && o.ErrorSummary != nil {
		return true
	}

	return false
}

// SetErrorSummary gets a reference to the given string and assigns it to the ErrorSummary field.
func (o *RegistrationResponseError) SetErrorSummary(v string) {
	o.ErrorSummary = &v
}

// GetErrorCauses returns the ErrorCauses field value if set, zero value otherwise.
func (o *RegistrationResponseError) GetErrorCauses() []RegistrationResponseErrorErrorCausesInner {
	if o == nil || o.ErrorCauses == nil {
		var ret []RegistrationResponseErrorErrorCausesInner
		return ret
	}
	return o.ErrorCauses
}

// GetErrorCausesOk returns a tuple with the ErrorCauses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationResponseError) GetErrorCausesOk() ([]RegistrationResponseErrorErrorCausesInner, bool) {
	if o == nil || o.ErrorCauses == nil {
		return nil, false
	}
	return o.ErrorCauses, true
}

// HasErrorCauses returns a boolean if a field has been set.
func (o *RegistrationResponseError) HasErrorCauses() bool {
	if o != nil && o.ErrorCauses != nil {
		return true
	}

	return false
}

// SetErrorCauses gets a reference to the given []RegistrationResponseErrorErrorCausesInner and assigns it to the ErrorCauses field.
func (o *RegistrationResponseError) SetErrorCauses(v []RegistrationResponseErrorErrorCausesInner) {
	o.ErrorCauses = v
}

func (o RegistrationResponseError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ErrorSummary != nil {
		toSerialize["errorSummary"] = o.ErrorSummary
	}
	if o.ErrorCauses != nil {
		toSerialize["errorCauses"] = o.ErrorCauses
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RegistrationResponseError) UnmarshalJSON(bytes []byte) (err error) {
	varRegistrationResponseError := _RegistrationResponseError{}

	err = json.Unmarshal(bytes, &varRegistrationResponseError)
	if err == nil {
		*o = RegistrationResponseError(varRegistrationResponseError)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "errorSummary")
		delete(additionalProperties, "errorCauses")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableRegistrationResponseError struct {
	value *RegistrationResponseError
	isSet bool
}

func (v NullableRegistrationResponseError) Get() *RegistrationResponseError {
	return v.value
}

func (v *NullableRegistrationResponseError) Set(val *RegistrationResponseError) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistrationResponseError) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistrationResponseError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistrationResponseError(val *RegistrationResponseError) *NullableRegistrationResponseError {
	return &NullableRegistrationResponseError{value: val, isSet: true}
}

func (v NullableRegistrationResponseError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistrationResponseError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

