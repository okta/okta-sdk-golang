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
	"time"
)

// UserFactorActivatePush Activation requests have a short lifetime and expire if activation isn't completed before the indicated timestamp. If the activation expires, use the returned `activate` link to restart the process.
type UserFactorActivatePush struct {
	// Timestamp when the Factor verification attempt expires
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
	// Result of a Factor verification
	FactorResult *string `json:"factorResult,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserFactorActivatePush UserFactorActivatePush

// NewUserFactorActivatePush instantiates a new UserFactorActivatePush object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserFactorActivatePush() *UserFactorActivatePush {
	this := UserFactorActivatePush{}
	return &this
}

// NewUserFactorActivatePushWithDefaults instantiates a new UserFactorActivatePush object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserFactorActivatePushWithDefaults() *UserFactorActivatePush {
	this := UserFactorActivatePush{}
	return &this
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *UserFactorActivatePush) GetExpiresAt() time.Time {
	if o == nil || o.ExpiresAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorActivatePush) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || o.ExpiresAt == nil {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *UserFactorActivatePush) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt != nil {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *UserFactorActivatePush) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

// GetFactorResult returns the FactorResult field value if set, zero value otherwise.
func (o *UserFactorActivatePush) GetFactorResult() string {
	if o == nil || o.FactorResult == nil {
		var ret string
		return ret
	}
	return *o.FactorResult
}

// GetFactorResultOk returns a tuple with the FactorResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorActivatePush) GetFactorResultOk() (*string, bool) {
	if o == nil || o.FactorResult == nil {
		return nil, false
	}
	return o.FactorResult, true
}

// HasFactorResult returns a boolean if a field has been set.
func (o *UserFactorActivatePush) HasFactorResult() bool {
	if o != nil && o.FactorResult != nil {
		return true
	}

	return false
}

// SetFactorResult gets a reference to the given string and assigns it to the FactorResult field.
func (o *UserFactorActivatePush) SetFactorResult(v string) {
	o.FactorResult = &v
}

func (o UserFactorActivatePush) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExpiresAt != nil {
		toSerialize["expiresAt"] = o.ExpiresAt
	}
	if o.FactorResult != nil {
		toSerialize["factorResult"] = o.FactorResult
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UserFactorActivatePush) UnmarshalJSON(bytes []byte) (err error) {
	varUserFactorActivatePush := _UserFactorActivatePush{}

	err = json.Unmarshal(bytes, &varUserFactorActivatePush)
	if err == nil {
		*o = UserFactorActivatePush(varUserFactorActivatePush)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "expiresAt")
		delete(additionalProperties, "factorResult")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableUserFactorActivatePush struct {
	value *UserFactorActivatePush
	isSet bool
}

func (v NullableUserFactorActivatePush) Get() *UserFactorActivatePush {
	return v.value
}

func (v *NullableUserFactorActivatePush) Set(val *UserFactorActivatePush) {
	v.value = val
	v.isSet = true
}

func (v NullableUserFactorActivatePush) IsSet() bool {
	return v.isSet
}

func (v *NullableUserFactorActivatePush) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserFactorActivatePush(val *UserFactorActivatePush) *NullableUserFactorActivatePush {
	return &NullableUserFactorActivatePush{value: val, isSet: true}
}

func (v NullableUserFactorActivatePush) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserFactorActivatePush) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

