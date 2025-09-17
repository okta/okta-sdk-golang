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

// checks if the UserFactorPushTransactionRejectedAllOfLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserFactorPushTransactionRejectedAllOfLinks{}

// UserFactorPushTransactionRejectedAllOfLinks struct for UserFactorPushTransactionRejectedAllOfLinks
type UserFactorPushTransactionRejectedAllOfLinks struct {
	// Verifies the factor resource. See [Verify a factor](/openapi/okta-management/management/tag/UserFactor/#tag/UserFactor/operation/verifyFactor).
	Verify               *HrefObject `json:"verify,omitempty"`
	Factor               *UserFactor `json:"factor,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserFactorPushTransactionRejectedAllOfLinks UserFactorPushTransactionRejectedAllOfLinks

// NewUserFactorPushTransactionRejectedAllOfLinks instantiates a new UserFactorPushTransactionRejectedAllOfLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserFactorPushTransactionRejectedAllOfLinks() *UserFactorPushTransactionRejectedAllOfLinks {
	this := UserFactorPushTransactionRejectedAllOfLinks{}
	return &this
}

// NewUserFactorPushTransactionRejectedAllOfLinksWithDefaults instantiates a new UserFactorPushTransactionRejectedAllOfLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserFactorPushTransactionRejectedAllOfLinksWithDefaults() *UserFactorPushTransactionRejectedAllOfLinks {
	this := UserFactorPushTransactionRejectedAllOfLinks{}
	return &this
}

// GetVerify returns the Verify field value if set, zero value otherwise.
func (o *UserFactorPushTransactionRejectedAllOfLinks) GetVerify() HrefObject {
	if o == nil || IsNil(o.Verify) {
		var ret HrefObject
		return ret
	}
	return *o.Verify
}

// GetVerifyOk returns a tuple with the Verify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorPushTransactionRejectedAllOfLinks) GetVerifyOk() (*HrefObject, bool) {
	if o == nil || IsNil(o.Verify) {
		return nil, false
	}
	return o.Verify, true
}

// HasVerify returns a boolean if a field has been set.
func (o *UserFactorPushTransactionRejectedAllOfLinks) HasVerify() bool {
	if o != nil && !IsNil(o.Verify) {
		return true
	}

	return false
}

// SetVerify gets a reference to the given HrefObject and assigns it to the Verify field.
func (o *UserFactorPushTransactionRejectedAllOfLinks) SetVerify(v HrefObject) {
	o.Verify = &v
}

// GetFactor returns the Factor field value if set, zero value otherwise.
func (o *UserFactorPushTransactionRejectedAllOfLinks) GetFactor() UserFactor {
	if o == nil || IsNil(o.Factor) {
		var ret UserFactor
		return ret
	}
	return *o.Factor
}

// GetFactorOk returns a tuple with the Factor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorPushTransactionRejectedAllOfLinks) GetFactorOk() (*UserFactor, bool) {
	if o == nil || IsNil(o.Factor) {
		return nil, false
	}
	return o.Factor, true
}

// HasFactor returns a boolean if a field has been set.
func (o *UserFactorPushTransactionRejectedAllOfLinks) HasFactor() bool {
	if o != nil && !IsNil(o.Factor) {
		return true
	}

	return false
}

// SetFactor gets a reference to the given UserFactor and assigns it to the Factor field.
func (o *UserFactorPushTransactionRejectedAllOfLinks) SetFactor(v UserFactor) {
	o.Factor = &v
}

func (o UserFactorPushTransactionRejectedAllOfLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserFactorPushTransactionRejectedAllOfLinks) ToMap() (map[string]interface{}, error) {
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

func (o *UserFactorPushTransactionRejectedAllOfLinks) UnmarshalJSON(data []byte) (err error) {
	varUserFactorPushTransactionRejectedAllOfLinks := _UserFactorPushTransactionRejectedAllOfLinks{}

	err = json.Unmarshal(data, &varUserFactorPushTransactionRejectedAllOfLinks)

	if err != nil {
		return err
	}

	*o = UserFactorPushTransactionRejectedAllOfLinks(varUserFactorPushTransactionRejectedAllOfLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "verify")
		delete(additionalProperties, "factor")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserFactorPushTransactionRejectedAllOfLinks struct {
	value *UserFactorPushTransactionRejectedAllOfLinks
	isSet bool
}

func (v NullableUserFactorPushTransactionRejectedAllOfLinks) Get() *UserFactorPushTransactionRejectedAllOfLinks {
	return v.value
}

func (v *NullableUserFactorPushTransactionRejectedAllOfLinks) Set(val *UserFactorPushTransactionRejectedAllOfLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableUserFactorPushTransactionRejectedAllOfLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableUserFactorPushTransactionRejectedAllOfLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserFactorPushTransactionRejectedAllOfLinks(val *UserFactorPushTransactionRejectedAllOfLinks) *NullableUserFactorPushTransactionRejectedAllOfLinks {
	return &NullableUserFactorPushTransactionRejectedAllOfLinks{value: val, isSet: true}
}

func (v NullableUserFactorPushTransactionRejectedAllOfLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserFactorPushTransactionRejectedAllOfLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
