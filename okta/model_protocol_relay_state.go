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

// ProtocolRelayState struct for ProtocolRelayState
type ProtocolRelayState struct {
	Format *string `json:"format,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProtocolRelayState ProtocolRelayState

// NewProtocolRelayState instantiates a new ProtocolRelayState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProtocolRelayState() *ProtocolRelayState {
	this := ProtocolRelayState{}
	return &this
}

// NewProtocolRelayStateWithDefaults instantiates a new ProtocolRelayState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProtocolRelayStateWithDefaults() *ProtocolRelayState {
	this := ProtocolRelayState{}
	return &this
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *ProtocolRelayState) GetFormat() string {
	if o == nil || o.Format == nil {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtocolRelayState) GetFormatOk() (*string, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *ProtocolRelayState) HasFormat() bool {
	if o != nil && o.Format != nil {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *ProtocolRelayState) SetFormat(v string) {
	o.Format = &v
}

func (o ProtocolRelayState) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Format != nil {
		toSerialize["format"] = o.Format
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ProtocolRelayState) UnmarshalJSON(bytes []byte) (err error) {
	varProtocolRelayState := _ProtocolRelayState{}

	err = json.Unmarshal(bytes, &varProtocolRelayState)
	if err == nil {
		*o = ProtocolRelayState(varProtocolRelayState)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "format")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableProtocolRelayState struct {
	value *ProtocolRelayState
	isSet bool
}

func (v NullableProtocolRelayState) Get() *ProtocolRelayState {
	return v.value
}

func (v *NullableProtocolRelayState) Set(val *ProtocolRelayState) {
	v.value = val
	v.isSet = true
}

func (v NullableProtocolRelayState) IsSet() bool {
	return v.isSet
}

func (v *NullableProtocolRelayState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProtocolRelayState(val *ProtocolRelayState) *NullableProtocolRelayState {
	return &NullableProtocolRelayState{value: val, isSet: true}
}

func (v NullableProtocolRelayState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProtocolRelayState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

