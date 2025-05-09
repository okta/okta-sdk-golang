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

// EmailTestAddresses struct for EmailTestAddresses
type EmailTestAddresses struct {
	// Email address that sends test emails
	From string `json:"from"`
	// Email address that receives test emails
	To string `json:"to"`
	AdditionalProperties map[string]interface{}
}

type _EmailTestAddresses EmailTestAddresses

// NewEmailTestAddresses instantiates a new EmailTestAddresses object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailTestAddresses(from string, to string) *EmailTestAddresses {
	this := EmailTestAddresses{}
	this.From = from
	this.To = to
	return &this
}

// NewEmailTestAddressesWithDefaults instantiates a new EmailTestAddresses object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailTestAddressesWithDefaults() *EmailTestAddresses {
	this := EmailTestAddresses{}
	return &this
}

// GetFrom returns the From field value
func (o *EmailTestAddresses) GetFrom() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *EmailTestAddresses) GetFromOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value
func (o *EmailTestAddresses) SetFrom(v string) {
	o.From = v
}

// GetTo returns the To field value
func (o *EmailTestAddresses) GetTo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *EmailTestAddresses) GetToOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value
func (o *EmailTestAddresses) SetTo(v string) {
	o.To = v
}

func (o EmailTestAddresses) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["from"] = o.From
	}
	if true {
		toSerialize["to"] = o.To
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EmailTestAddresses) UnmarshalJSON(bytes []byte) (err error) {
	varEmailTestAddresses := _EmailTestAddresses{}

	err = json.Unmarshal(bytes, &varEmailTestAddresses)
	if err == nil {
		*o = EmailTestAddresses(varEmailTestAddresses)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "from")
		delete(additionalProperties, "to")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableEmailTestAddresses struct {
	value *EmailTestAddresses
	isSet bool
}

func (v NullableEmailTestAddresses) Get() *EmailTestAddresses {
	return v.value
}

func (v *NullableEmailTestAddresses) Set(val *EmailTestAddresses) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailTestAddresses) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailTestAddresses) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailTestAddresses(val *EmailTestAddresses) *NullableEmailTestAddresses {
	return &NullableEmailTestAddresses{value: val, isSet: true}
}

func (v NullableEmailTestAddresses) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailTestAddresses) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

