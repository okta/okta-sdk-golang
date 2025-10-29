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

// checks if the UserFactorPushTransactionWaitingNoNMCAllOfLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserFactorPushTransactionWaitingNoNMCAllOfLinks{}

// UserFactorPushTransactionWaitingNoNMCAllOfLinks struct for UserFactorPushTransactionWaitingNoNMCAllOfLinks
type UserFactorPushTransactionWaitingNoNMCAllOfLinks struct {
	// Polls the factor resource for status information. Always use the `poll` link instead of manually constructing your own URL.
	Poll *HrefObject `json:"poll,omitempty"`
	// Cancels a `push` factor challenge with a `WAITING` status
	Cancel               *HrefObject `json:"cancel,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserFactorPushTransactionWaitingNoNMCAllOfLinks UserFactorPushTransactionWaitingNoNMCAllOfLinks

// NewUserFactorPushTransactionWaitingNoNMCAllOfLinks instantiates a new UserFactorPushTransactionWaitingNoNMCAllOfLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserFactorPushTransactionWaitingNoNMCAllOfLinks() *UserFactorPushTransactionWaitingNoNMCAllOfLinks {
	this := UserFactorPushTransactionWaitingNoNMCAllOfLinks{}
	return &this
}

// NewUserFactorPushTransactionWaitingNoNMCAllOfLinksWithDefaults instantiates a new UserFactorPushTransactionWaitingNoNMCAllOfLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserFactorPushTransactionWaitingNoNMCAllOfLinksWithDefaults() *UserFactorPushTransactionWaitingNoNMCAllOfLinks {
	this := UserFactorPushTransactionWaitingNoNMCAllOfLinks{}
	return &this
}

// GetPoll returns the Poll field value if set, zero value otherwise.
func (o *UserFactorPushTransactionWaitingNoNMCAllOfLinks) GetPoll() HrefObject {
	if o == nil || IsNil(o.Poll) {
		var ret HrefObject
		return ret
	}
	return *o.Poll
}

// GetPollOk returns a tuple with the Poll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorPushTransactionWaitingNoNMCAllOfLinks) GetPollOk() (*HrefObject, bool) {
	if o == nil || IsNil(o.Poll) {
		return nil, false
	}
	return o.Poll, true
}

// HasPoll returns a boolean if a field has been set.
func (o *UserFactorPushTransactionWaitingNoNMCAllOfLinks) HasPoll() bool {
	if o != nil && !IsNil(o.Poll) {
		return true
	}

	return false
}

// SetPoll gets a reference to the given HrefObject and assigns it to the Poll field.
func (o *UserFactorPushTransactionWaitingNoNMCAllOfLinks) SetPoll(v HrefObject) {
	o.Poll = &v
}

// GetCancel returns the Cancel field value if set, zero value otherwise.
func (o *UserFactorPushTransactionWaitingNoNMCAllOfLinks) GetCancel() HrefObject {
	if o == nil || IsNil(o.Cancel) {
		var ret HrefObject
		return ret
	}
	return *o.Cancel
}

// GetCancelOk returns a tuple with the Cancel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorPushTransactionWaitingNoNMCAllOfLinks) GetCancelOk() (*HrefObject, bool) {
	if o == nil || IsNil(o.Cancel) {
		return nil, false
	}
	return o.Cancel, true
}

// HasCancel returns a boolean if a field has been set.
func (o *UserFactorPushTransactionWaitingNoNMCAllOfLinks) HasCancel() bool {
	if o != nil && !IsNil(o.Cancel) {
		return true
	}

	return false
}

// SetCancel gets a reference to the given HrefObject and assigns it to the Cancel field.
func (o *UserFactorPushTransactionWaitingNoNMCAllOfLinks) SetCancel(v HrefObject) {
	o.Cancel = &v
}

func (o UserFactorPushTransactionWaitingNoNMCAllOfLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserFactorPushTransactionWaitingNoNMCAllOfLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Poll) {
		toSerialize["poll"] = o.Poll
	}
	if !IsNil(o.Cancel) {
		toSerialize["cancel"] = o.Cancel
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserFactorPushTransactionWaitingNoNMCAllOfLinks) UnmarshalJSON(data []byte) (err error) {
	varUserFactorPushTransactionWaitingNoNMCAllOfLinks := _UserFactorPushTransactionWaitingNoNMCAllOfLinks{}

	err = json.Unmarshal(data, &varUserFactorPushTransactionWaitingNoNMCAllOfLinks)

	if err != nil {
		return err
	}

	*o = UserFactorPushTransactionWaitingNoNMCAllOfLinks(varUserFactorPushTransactionWaitingNoNMCAllOfLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "poll")
		delete(additionalProperties, "cancel")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserFactorPushTransactionWaitingNoNMCAllOfLinks struct {
	value *UserFactorPushTransactionWaitingNoNMCAllOfLinks
	isSet bool
}

func (v NullableUserFactorPushTransactionWaitingNoNMCAllOfLinks) Get() *UserFactorPushTransactionWaitingNoNMCAllOfLinks {
	return v.value
}

func (v *NullableUserFactorPushTransactionWaitingNoNMCAllOfLinks) Set(val *UserFactorPushTransactionWaitingNoNMCAllOfLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableUserFactorPushTransactionWaitingNoNMCAllOfLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableUserFactorPushTransactionWaitingNoNMCAllOfLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserFactorPushTransactionWaitingNoNMCAllOfLinks(val *UserFactorPushTransactionWaitingNoNMCAllOfLinks) *NullableUserFactorPushTransactionWaitingNoNMCAllOfLinks {
	return &NullableUserFactorPushTransactionWaitingNoNMCAllOfLinks{value: val, isSet: true}
}

func (v NullableUserFactorPushTransactionWaitingNoNMCAllOfLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserFactorPushTransactionWaitingNoNMCAllOfLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
