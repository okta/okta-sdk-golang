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

// UserFactorPushTransactionTimeoutAllOfLinks struct for UserFactorPushTransactionTimeoutAllOfLinks
type UserFactorPushTransactionTimeoutAllOfLinks struct {
	// Verifies the factor resource. See [Verify a factor](/openapi/okta-management/management/tag/UserFactor/#tag/UserFactor/operation/verifyFactor).
	Verify *HrefObject `json:"verify,omitempty"`
	Factor *UserFactor `json:"factor,omitempty"`
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
	if o == nil || o.Verify == nil {
		var ret HrefObject
		return ret
	}
	return *o.Verify
}

// GetVerifyOk returns a tuple with the Verify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorPushTransactionTimeoutAllOfLinks) GetVerifyOk() (*HrefObject, bool) {
	if o == nil || o.Verify == nil {
		return nil, false
	}
	return o.Verify, true
}

// HasVerify returns a boolean if a field has been set.
func (o *UserFactorPushTransactionTimeoutAllOfLinks) HasVerify() bool {
	if o != nil && o.Verify != nil {
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
	if o == nil || o.Factor == nil {
		var ret UserFactor
		return ret
	}
	return *o.Factor
}

// GetFactorOk returns a tuple with the Factor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorPushTransactionTimeoutAllOfLinks) GetFactorOk() (*UserFactor, bool) {
	if o == nil || o.Factor == nil {
		return nil, false
	}
	return o.Factor, true
}

// HasFactor returns a boolean if a field has been set.
func (o *UserFactorPushTransactionTimeoutAllOfLinks) HasFactor() bool {
	if o != nil && o.Factor != nil {
		return true
	}

	return false
}

// SetFactor gets a reference to the given UserFactor and assigns it to the Factor field.
func (o *UserFactorPushTransactionTimeoutAllOfLinks) SetFactor(v UserFactor) {
	o.Factor = &v
}

func (o UserFactorPushTransactionTimeoutAllOfLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Verify != nil {
		toSerialize["verify"] = o.Verify
	}
	if o.Factor != nil {
		toSerialize["factor"] = o.Factor
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UserFactorPushTransactionTimeoutAllOfLinks) UnmarshalJSON(bytes []byte) (err error) {
	varUserFactorPushTransactionTimeoutAllOfLinks := _UserFactorPushTransactionTimeoutAllOfLinks{}

	err = json.Unmarshal(bytes, &varUserFactorPushTransactionTimeoutAllOfLinks)
	if err == nil {
		*o = UserFactorPushTransactionTimeoutAllOfLinks(varUserFactorPushTransactionTimeoutAllOfLinks)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "verify")
		delete(additionalProperties, "factor")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
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

