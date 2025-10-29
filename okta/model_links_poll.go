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

API version: 2025.10.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the LinksPoll type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LinksPoll{}

// LinksPoll struct for LinksPoll
type LinksPoll struct {
	// Polls the factor resource for status information. Always use the `poll` link instead of manually constructing your own URL.
	Poll                 *HrefObject `json:"poll,omitempty"`
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
func (o *LinksPoll) GetPoll() HrefObject {
	if o == nil || IsNil(o.Poll) {
		var ret HrefObject
		return ret
	}
	return *o.Poll
}

// GetPollOk returns a tuple with the Poll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinksPoll) GetPollOk() (*HrefObject, bool) {
	if o == nil || IsNil(o.Poll) {
		return nil, false
	}
	return o.Poll, true
}

// HasPoll returns a boolean if a field has been set.
func (o *LinksPoll) HasPoll() bool {
	if o != nil && !IsNil(o.Poll) {
		return true
	}

	return false
}

// SetPoll gets a reference to the given HrefObject and assigns it to the Poll field.
func (o *LinksPoll) SetPoll(v HrefObject) {
	o.Poll = &v
}

func (o LinksPoll) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LinksPoll) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Poll) {
		toSerialize["poll"] = o.Poll
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LinksPoll) UnmarshalJSON(data []byte) (err error) {
	varLinksPoll := _LinksPoll{}

	err = json.Unmarshal(data, &varLinksPoll)

	if err != nil {
		return err
	}

	*o = LinksPoll(varLinksPoll)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "poll")
		o.AdditionalProperties = additionalProperties
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
