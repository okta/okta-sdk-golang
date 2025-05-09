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

// LinksSend struct for LinksSend
type LinksSend struct {
	Send *LinksSendSend `json:"send,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LinksSend LinksSend

// NewLinksSend instantiates a new LinksSend object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinksSend() *LinksSend {
	this := LinksSend{}
	return &this
}

// NewLinksSendWithDefaults instantiates a new LinksSend object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinksSendWithDefaults() *LinksSend {
	this := LinksSend{}
	return &this
}

// GetSend returns the Send field value if set, zero value otherwise.
func (o *LinksSend) GetSend() LinksSendSend {
	if o == nil || o.Send == nil {
		var ret LinksSendSend
		return ret
	}
	return *o.Send
}

// GetSendOk returns a tuple with the Send field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinksSend) GetSendOk() (*LinksSendSend, bool) {
	if o == nil || o.Send == nil {
		return nil, false
	}
	return o.Send, true
}

// HasSend returns a boolean if a field has been set.
func (o *LinksSend) HasSend() bool {
	if o != nil && o.Send != nil {
		return true
	}

	return false
}

// SetSend gets a reference to the given LinksSendSend and assigns it to the Send field.
func (o *LinksSend) SetSend(v LinksSendSend) {
	o.Send = &v
}

func (o LinksSend) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Send != nil {
		toSerialize["send"] = o.Send
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LinksSend) UnmarshalJSON(bytes []byte) (err error) {
	varLinksSend := _LinksSend{}

	err = json.Unmarshal(bytes, &varLinksSend)
	if err == nil {
		*o = LinksSend(varLinksSend)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "send")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableLinksSend struct {
	value *LinksSend
	isSet bool
}

func (v NullableLinksSend) Get() *LinksSend {
	return v.value
}

func (v *NullableLinksSend) Set(val *LinksSend) {
	v.value = val
	v.isSet = true
}

func (v NullableLinksSend) IsSet() bool {
	return v.isSet
}

func (v *NullableLinksSend) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinksSend(val *LinksSend) *NullableLinksSend {
	return &NullableLinksSend{value: val, isSet: true}
}

func (v NullableLinksSend) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinksSend) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

