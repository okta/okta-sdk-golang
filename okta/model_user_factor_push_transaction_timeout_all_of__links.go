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

// checks if the UserFactorPushTransactionTimeoutAllOfLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserFactorPushTransactionTimeoutAllOfLinks{}

// UserFactorPushTransactionTimeoutAllOfLinks struct for UserFactorPushTransactionTimeoutAllOfLinks
type UserFactorPushTransactionTimeoutAllOfLinks struct {
	// Verifies the factor resource. See [Verify a factor](/openapi/okta-management/management/tag/UserFactor/#tag/UserFactor/operation/verifyFactor).
	Verify               *HrefObject `json:"verify,omitempty"`
	Factor               *UserFactor `json:"factor,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserFactorPushTransactionTimeoutAllOfLinks UserFactorPushTransactionTimeoutAllOfLinks

// NewUserFactorPushTransactionTimeoutAllOfLinks instantiates a new UserFactorPushTransactionTimeoutAllOfLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserFactorPushTransactionTimeoutAllOfLinks() *UserFactorPushTransactionTimeoutAllOfLinks {
	this := UserFactorPushTransactionTimeoutAllOfLinks{}
	return &this
}

// NewUserFactorPushTransactionTimeoutAllOfLinksWithDefaults instantiates a new UserFactorPushTransactionTimeoutAllOfLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserFactorPushTransactionTimeoutAllOfLinksWithDefaults() *UserFactorPushTransactionTimeoutAllOfLinks {
	this := UserFactorPushTransactionTimeoutAllOfLinks{}
	return &this
}

// GetVerify returns the Verify field value if set, zero value otherwise.
func (o *UserFactorPushTransactionTimeoutAllOfLinks) GetVerify() HrefObject {
	if o == nil || IsNil(o.Verify) {
		var ret HrefObject
		return ret
	}
	return *o.Verify
}

// GetVerifyOk returns a tuple with the Verify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorPushTransactionTimeoutAllOfLinks) GetVerifyOk() (*HrefObject, bool) {
	if o == nil || IsNil(o.Verify) {
		return nil, false
	}
	return o.Verify, true
}

// HasVerify returns a boolean if a field has been set.
func (o *UserFactorPushTransactionTimeoutAllOfLinks) HasVerify() bool {
	if o != nil && !IsNil(o.Verify) {
		return true
	}

	return false
}

// SetVerify gets a reference to the given HrefObject and assigns it to the Verify field.
func (o *UserFactorPushTransactionTimeoutAllOfLinks) SetVerify(v HrefObject) {
	o.Verify = &v
}

// GetFactor returns the Factor field value if set, zero value otherwise.
func (o *UserFactorPushTransactionTimeoutAllOfLinks) GetFactor() UserFactor {
	if o == nil || IsNil(o.Factor) {
		var ret UserFactor
		return ret
	}
	return *o.Factor
}

// GetFactorOk returns a tuple with the Factor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorPushTransactionTimeoutAllOfLinks) GetFactorOk() (*UserFactor, bool) {
	if o == nil || IsNil(o.Factor) {
		return nil, false
	}
	return o.Factor, true
}

// HasFactor returns a boolean if a field has been set.
func (o *UserFactorPushTransactionTimeoutAllOfLinks) HasFactor() bool {
	if o != nil && !IsNil(o.Factor) {
		return true
	}

	return false
}

// SetFactor gets a reference to the given UserFactor and assigns it to the Factor field.
func (o *UserFactorPushTransactionTimeoutAllOfLinks) SetFactor(v UserFactor) {
	o.Factor = &v
}

func (o UserFactorPushTransactionTimeoutAllOfLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserFactorPushTransactionTimeoutAllOfLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Verify) {
		toSerialize["verify"] = o.Verify
	}
	if !IsNil(o.Factor) {
		toSerialize["factor"] = o.Factor
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserFactorPushTransactionTimeoutAllOfLinks) UnmarshalJSON(data []byte) (err error) {
	varUserFactorPushTransactionTimeoutAllOfLinks := _UserFactorPushTransactionTimeoutAllOfLinks{}

	err = json.Unmarshal(data, &varUserFactorPushTransactionTimeoutAllOfLinks)

	if err != nil {
		return err
	}

	*o = UserFactorPushTransactionTimeoutAllOfLinks(varUserFactorPushTransactionTimeoutAllOfLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "verify")
		delete(additionalProperties, "factor")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserFactorPushTransactionTimeoutAllOfLinks struct {
	value *UserFactorPushTransactionTimeoutAllOfLinks
	isSet bool
}

func (v NullableUserFactorPushTransactionTimeoutAllOfLinks) Get() *UserFactorPushTransactionTimeoutAllOfLinks {
	return v.value
}

func (v *NullableUserFactorPushTransactionTimeoutAllOfLinks) Set(val *UserFactorPushTransactionTimeoutAllOfLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableUserFactorPushTransactionTimeoutAllOfLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableUserFactorPushTransactionTimeoutAllOfLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserFactorPushTransactionTimeoutAllOfLinks(val *UserFactorPushTransactionTimeoutAllOfLinks) *NullableUserFactorPushTransactionTimeoutAllOfLinks {
	return &NullableUserFactorPushTransactionTimeoutAllOfLinks{value: val, isSet: true}
}

func (v NullableUserFactorPushTransactionTimeoutAllOfLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserFactorPushTransactionTimeoutAllOfLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
