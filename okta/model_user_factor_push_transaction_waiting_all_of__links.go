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

// UserFactorPushTransactionWaitingAllOfLinks struct for UserFactorPushTransactionWaitingAllOfLinks
type UserFactorPushTransactionWaitingAllOfLinks struct {
	Poll *LinksPollPoll `json:"poll,omitempty"`
	Cancel *LinksCancelCancel `json:"cancel,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserFactorPushTransactionWaitingAllOfLinks UserFactorPushTransactionWaitingAllOfLinks

// NewUserFactorPushTransactionWaitingAllOfLinks instantiates a new UserFactorPushTransactionWaitingAllOfLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserFactorPushTransactionWaitingAllOfLinks() *UserFactorPushTransactionWaitingAllOfLinks {
	this := UserFactorPushTransactionWaitingAllOfLinks{}
	return &this
}

// NewUserFactorPushTransactionWaitingAllOfLinksWithDefaults instantiates a new UserFactorPushTransactionWaitingAllOfLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserFactorPushTransactionWaitingAllOfLinksWithDefaults() *UserFactorPushTransactionWaitingAllOfLinks {
	this := UserFactorPushTransactionWaitingAllOfLinks{}
	return &this
}

// GetPoll returns the Poll field value if set, zero value otherwise.
func (o *UserFactorPushTransactionWaitingAllOfLinks) GetPoll() LinksPollPoll {
	if o == nil || o.Poll == nil {
		var ret LinksPollPoll
		return ret
	}
	return *o.Poll
}

// GetPollOk returns a tuple with the Poll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorPushTransactionWaitingAllOfLinks) GetPollOk() (*LinksPollPoll, bool) {
	if o == nil || o.Poll == nil {
		return nil, false
	}
	return o.Poll, true
}

// HasPoll returns a boolean if a field has been set.
func (o *UserFactorPushTransactionWaitingAllOfLinks) HasPoll() bool {
	if o != nil && o.Poll != nil {
		return true
	}

	return false
}

// SetPoll gets a reference to the given LinksPollPoll and assigns it to the Poll field.
func (o *UserFactorPushTransactionWaitingAllOfLinks) SetPoll(v LinksPollPoll) {
	o.Poll = &v
}

// GetCancel returns the Cancel field value if set, zero value otherwise.
func (o *UserFactorPushTransactionWaitingAllOfLinks) GetCancel() LinksCancelCancel {
	if o == nil || o.Cancel == nil {
		var ret LinksCancelCancel
		return ret
	}
	return *o.Cancel
}

// GetCancelOk returns a tuple with the Cancel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorPushTransactionWaitingAllOfLinks) GetCancelOk() (*LinksCancelCancel, bool) {
	if o == nil || o.Cancel == nil {
		return nil, false
	}
	return o.Cancel, true
}

// HasCancel returns a boolean if a field has been set.
func (o *UserFactorPushTransactionWaitingAllOfLinks) HasCancel() bool {
	if o != nil && o.Cancel != nil {
		return true
	}

	return false
}

// SetCancel gets a reference to the given LinksCancelCancel and assigns it to the Cancel field.
func (o *UserFactorPushTransactionWaitingAllOfLinks) SetCancel(v LinksCancelCancel) {
	o.Cancel = &v
}

func (o UserFactorPushTransactionWaitingAllOfLinks) MarshalJSON() ([]byte, error) {
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

func (o *UserFactorPushTransactionWaitingAllOfLinks) UnmarshalJSON(bytes []byte) (err error) {
	varUserFactorPushTransactionWaitingAllOfLinks := _UserFactorPushTransactionWaitingAllOfLinks{}

	err = json.Unmarshal(bytes, &varUserFactorPushTransactionWaitingAllOfLinks)
	if err == nil {
		*o = UserFactorPushTransactionWaitingAllOfLinks(varUserFactorPushTransactionWaitingAllOfLinks)
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

type NullableUserFactorPushTransactionWaitingAllOfLinks struct {
	value *UserFactorPushTransactionWaitingAllOfLinks
	isSet bool
}

func (v NullableUserFactorPushTransactionWaitingAllOfLinks) Get() *UserFactorPushTransactionWaitingAllOfLinks {
	return v.value
}

func (v *NullableUserFactorPushTransactionWaitingAllOfLinks) Set(val *UserFactorPushTransactionWaitingAllOfLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableUserFactorPushTransactionWaitingAllOfLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableUserFactorPushTransactionWaitingAllOfLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserFactorPushTransactionWaitingAllOfLinks(val *UserFactorPushTransactionWaitingAllOfLinks) *NullableUserFactorPushTransactionWaitingAllOfLinks {
	return &NullableUserFactorPushTransactionWaitingAllOfLinks{value: val, isSet: true}
}

func (v NullableUserFactorPushTransactionWaitingAllOfLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserFactorPushTransactionWaitingAllOfLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

