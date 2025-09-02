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

// InlineObjectResultsInner struct for InlineObjectResultsInner
type InlineObjectResultsInner struct {
	// Domain for your org
	Domain *string `json:"domain,omitempty"`
	// Domain failover message
	Message *string `json:"message,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InlineObjectResultsInner InlineObjectResultsInner

// NewInlineObjectResultsInner instantiates a new InlineObjectResultsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObjectResultsInner() *InlineObjectResultsInner {
	this := InlineObjectResultsInner{}
	return &this
}

// NewInlineObjectResultsInnerWithDefaults instantiates a new InlineObjectResultsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObjectResultsInnerWithDefaults() *InlineObjectResultsInner {
	this := InlineObjectResultsInner{}
	return &this
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *InlineObjectResultsInner) GetDomain() string {
	if o == nil || o.Domain == nil {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObjectResultsInner) GetDomainOk() (*string, bool) {
	if o == nil || o.Domain == nil {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *InlineObjectResultsInner) HasDomain() bool {
	if o != nil && o.Domain != nil {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *InlineObjectResultsInner) SetDomain(v string) {
	o.Domain = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *InlineObjectResultsInner) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObjectResultsInner) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *InlineObjectResultsInner) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *InlineObjectResultsInner) SetMessage(v string) {
	o.Message = &v
}

func (o InlineObjectResultsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Domain != nil {
		toSerialize["domain"] = o.Domain
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *InlineObjectResultsInner) UnmarshalJSON(bytes []byte) (err error) {
	varInlineObjectResultsInner := _InlineObjectResultsInner{}

	err = json.Unmarshal(bytes, &varInlineObjectResultsInner)
	if err == nil {
		*o = InlineObjectResultsInner(varInlineObjectResultsInner)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "domain")
		delete(additionalProperties, "message")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableInlineObjectResultsInner struct {
	value *InlineObjectResultsInner
	isSet bool
}

func (v NullableInlineObjectResultsInner) Get() *InlineObjectResultsInner {
	return v.value
}

func (v *NullableInlineObjectResultsInner) Set(val *InlineObjectResultsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObjectResultsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObjectResultsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObjectResultsInner(val *InlineObjectResultsInner) *NullableInlineObjectResultsInner {
	return &NullableInlineObjectResultsInner{value: val, isSet: true}
}

func (v NullableInlineObjectResultsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObjectResultsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

