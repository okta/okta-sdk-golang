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

// UserFactorPushTransactionWaitingNMCAllOfLinks struct for UserFactorPushTransactionWaitingNMCAllOfLinks
type UserFactorPushTransactionWaitingNMCAllOfLinks struct {
	// Polls the factor resource for status information. Always use the `poll` link instead of manually constructing your own URL.
	Poll *HrefObject `json:"poll,omitempty"`
	// Cancels a `push` factor challenge with a `WAITING` status
	Cancel *HrefObject `json:"cancel,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserFactorPushTransactionWaitingNMCAllOfLinks UserFactorPushTransactionWaitingNMCAllOfLinks

// NewUserFactorPushTransactionWaitingNMCAllOfLinks instantiates a new UserFactorPushTransactionWaitingNMCAllOfLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserFactorPushTransactionWaitingNMCAllOfLinks() *UserFactorPushTransactionWaitingNMCAllOfLinks {
	this := UserFactorPushTransactionWaitingNMCAllOfLinks{}
	return &this
}

// NewUserFactorPushTransactionWaitingNMCAllOfLinksWithDefaults instantiates a new UserFactorPushTransactionWaitingNMCAllOfLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserFactorPushTransactionWaitingNMCAllOfLinksWithDefaults() *UserFactorPushTransactionWaitingNMCAllOfLinks {
	this := UserFactorPushTransactionWaitingNMCAllOfLinks{}
	return &this
}

// GetPoll returns the Poll field value if set, zero value otherwise.
func (o *UserFactorPushTransactionWaitingNMCAllOfLinks) GetPoll() HrefObject {
	if o == nil || o.Poll == nil {
		var ret HrefObject
		return ret
	}
	return *o.Poll
}

// GetPollOk returns a tuple with the Poll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorPushTransactionWaitingNMCAllOfLinks) GetPollOk() (*HrefObject, bool) {
	if o == nil || o.Poll == nil {
		return nil, false
	}
	return o.Poll, true
}

// HasPoll returns a boolean if a field has been set.
func (o *UserFactorPushTransactionWaitingNMCAllOfLinks) HasPoll() bool {
	if o != nil && o.Poll != nil {
		return true
	}

	return false
}

// SetPoll gets a reference to the given HrefObject and assigns it to the Poll field.
func (o *UserFactorPushTransactionWaitingNMCAllOfLinks) SetPoll(v HrefObject) {
	o.Poll = &v
}

// GetCancel returns the Cancel field value if set, zero value otherwise.
func (o *UserFactorPushTransactionWaitingNMCAllOfLinks) GetCancel() HrefObject {
	if o == nil || o.Cancel == nil {
		var ret HrefObject
		return ret
	}
	return *o.Cancel
}

// GetCancelOk returns a tuple with the Cancel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorPushTransactionWaitingNMCAllOfLinks) GetCancelOk() (*HrefObject, bool) {
	if o == nil || o.Cancel == nil {
		return nil, false
	}
	return o.Cancel, true
}

// HasCancel returns a boolean if a field has been set.
func (o *UserFactorPushTransactionWaitingNMCAllOfLinks) HasCancel() bool {
	if o != nil && o.Cancel != nil {
		return true
	}

	return false
}

// SetCancel gets a reference to the given HrefObject and assigns it to the Cancel field.
func (o *UserFactorPushTransactionWaitingNMCAllOfLinks) SetCancel(v HrefObject) {
	o.Cancel = &v
}

func (o UserFactorPushTransactionWaitingNMCAllOfLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Poll != nil {
		toSerialize["poll"] = o.Poll
	}
	if o.Cancel != nil {
		toSerialize["cancel"] = o.Cancel
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UserFactorPushTransactionWaitingNMCAllOfLinks) UnmarshalJSON(bytes []byte) (err error) {
	varUserFactorPushTransactionWaitingNMCAllOfLinks := _UserFactorPushTransactionWaitingNMCAllOfLinks{}

	err = json.Unmarshal(bytes, &varUserFactorPushTransactionWaitingNMCAllOfLinks)
	if err == nil {
		*o = UserFactorPushTransactionWaitingNMCAllOfLinks(varUserFactorPushTransactionWaitingNMCAllOfLinks)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "poll")
		delete(additionalProperties, "cancel")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableUserFactorPushTransactionWaitingNMCAllOfLinks struct {
	value *UserFactorPushTransactionWaitingNMCAllOfLinks
	isSet bool
}

func (v NullableUserFactorPushTransactionWaitingNMCAllOfLinks) Get() *UserFactorPushTransactionWaitingNMCAllOfLinks {
	return v.value
}

func (v *NullableUserFactorPushTransactionWaitingNMCAllOfLinks) Set(val *UserFactorPushTransactionWaitingNMCAllOfLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableUserFactorPushTransactionWaitingNMCAllOfLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableUserFactorPushTransactionWaitingNMCAllOfLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserFactorPushTransactionWaitingNMCAllOfLinks(val *UserFactorPushTransactionWaitingNMCAllOfLinks) *NullableUserFactorPushTransactionWaitingNMCAllOfLinks {
	return &NullableUserFactorPushTransactionWaitingNMCAllOfLinks{value: val, isSet: true}
}

func (v NullableUserFactorPushTransactionWaitingNMCAllOfLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserFactorPushTransactionWaitingNMCAllOfLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

