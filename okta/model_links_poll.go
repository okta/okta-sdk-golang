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

// LinksPoll struct for LinksPoll
type LinksPoll struct {
	Poll *LinksPollPoll `json:"poll,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LinksPoll LinksPoll

// NewLinksPoll instantiates a new LinksPoll object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinksPoll() *LinksPoll {
	this := LinksPoll{}
	return &this
}

// NewLinksPollWithDefaults instantiates a new LinksPoll object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinksPollWithDefaults() *LinksPoll {
	this := LinksPoll{}
	return &this
}

// GetPoll returns the Poll field value if set, zero value otherwise.
func (o *LinksPoll) GetPoll() LinksPollPoll {
	if o == nil || o.Poll == nil {
		var ret LinksPollPoll
		return ret
	}
	return *o.Poll
}

// GetPollOk returns a tuple with the Poll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinksPoll) GetPollOk() (*LinksPollPoll, bool) {
	if o == nil || o.Poll == nil {
		return nil, false
	}
	return o.Poll, true
}

// HasPoll returns a boolean if a field has been set.
func (o *LinksPoll) HasPoll() bool {
	if o != nil && o.Poll != nil {
		return true
	}

	return false
}

// SetPoll gets a reference to the given LinksPollPoll and assigns it to the Poll field.
func (o *LinksPoll) SetPoll(v LinksPollPoll) {
	o.Poll = &v
}

func (o LinksPoll) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Poll != nil {
		toSerialize["poll"] = o.Poll
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LinksPoll) UnmarshalJSON(bytes []byte) (err error) {
	varLinksPoll := _LinksPoll{}

	err = json.Unmarshal(bytes, &varLinksPoll)
	if err == nil {
		*o = LinksPoll(varLinksPoll)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "poll")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableLinksPoll struct {
	value *LinksPoll
	isSet bool
}

func (v NullableLinksPoll) Get() *LinksPoll {
	return v.value
}

func (v *NullableLinksPoll) Set(val *LinksPoll) {
	v.value = val
	v.isSet = true
}

func (v NullableLinksPoll) IsSet() bool {
	return v.isSet
}

func (v *NullableLinksPoll) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinksPoll(val *LinksPoll) *NullableLinksPoll {
	return &NullableLinksPoll{value: val, isSet: true}
}

func (v NullableLinksPoll) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinksPoll) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

