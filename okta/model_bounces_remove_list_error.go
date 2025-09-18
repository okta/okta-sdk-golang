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

// checks if the BouncesRemoveListError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BouncesRemoveListError{}

// BouncesRemoveListError struct for BouncesRemoveListError
type BouncesRemoveListError struct {
	// An email address with a validation error
	EmailAddress *string `json:"emailAddress,omitempty"`
	// Validation error reason
	Reason               *string `json:"reason,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BouncesRemoveListError BouncesRemoveListError

// NewBouncesRemoveListError instantiates a new BouncesRemoveListError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBouncesRemoveListError() *BouncesRemoveListError {
	this := BouncesRemoveListError{}
	return &this
}

// NewBouncesRemoveListErrorWithDefaults instantiates a new BouncesRemoveListError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBouncesRemoveListErrorWithDefaults() *BouncesRemoveListError {
	this := BouncesRemoveListError{}
	return &this
}

// GetEmailAddress returns the EmailAddress field value if set, zero value otherwise.
func (o *BouncesRemoveListError) GetEmailAddress() string {
	if o == nil || IsNil(o.EmailAddress) {
		var ret string
		return ret
	}
	return *o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BouncesRemoveListError) GetEmailAddressOk() (*string, bool) {
	if o == nil || IsNil(o.EmailAddress) {
		return nil, false
	}
	return o.EmailAddress, true
}

// HasEmailAddress returns a boolean if a field has been set.
func (o *BouncesRemoveListError) HasEmailAddress() bool {
	if o != nil && !IsNil(o.EmailAddress) {
		return true
	}

	return false
}

// SetEmailAddress gets a reference to the given string and assigns it to the EmailAddress field.
func (o *BouncesRemoveListError) SetEmailAddress(v string) {
	o.EmailAddress = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *BouncesRemoveListError) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BouncesRemoveListError) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *BouncesRemoveListError) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *BouncesRemoveListError) SetReason(v string) {
	o.Reason = &v
}

func (o BouncesRemoveListError) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BouncesRemoveListError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EmailAddress) {
		toSerialize["emailAddress"] = o.EmailAddress
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BouncesRemoveListError) UnmarshalJSON(data []byte) (err error) {
	varBouncesRemoveListError := _BouncesRemoveListError{}

	err = json.Unmarshal(data, &varBouncesRemoveListError)

	if err != nil {
		return err
	}

	*o = BouncesRemoveListError(varBouncesRemoveListError)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "emailAddress")
		delete(additionalProperties, "reason")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBouncesRemoveListError struct {
	value *BouncesRemoveListError
	isSet bool
}

func (v NullableBouncesRemoveListError) Get() *BouncesRemoveListError {
	return v.value
}

func (v *NullableBouncesRemoveListError) Set(val *BouncesRemoveListError) {
	v.value = val
	v.isSet = true
}

func (v NullableBouncesRemoveListError) IsSet() bool {
	return v.isSet
}

func (v *NullableBouncesRemoveListError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBouncesRemoveListError(val *BouncesRemoveListError) *NullableBouncesRemoveListError {
	return &NullableBouncesRemoveListError{value: val, isSet: true}
}

func (v NullableBouncesRemoveListError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBouncesRemoveListError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
