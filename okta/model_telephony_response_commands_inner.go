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

// TelephonyResponseCommandsInner struct for TelephonyResponseCommandsInner
type TelephonyResponseCommandsInner struct {
	// The location where you specify the command. For the Telephony inline hook, there's only one command, `com.okta.telephony.action`.
	Type *string `json:"type,omitempty"`
	// The status of the telephony operation along with optional additional information about the provider, transaction ID and any other transaction metadata.
	Value []TelephonyResponseCommandsInnerValueInner `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TelephonyResponseCommandsInner TelephonyResponseCommandsInner

// NewTelephonyResponseCommandsInner instantiates a new TelephonyResponseCommandsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelephonyResponseCommandsInner() *TelephonyResponseCommandsInner {
	this := TelephonyResponseCommandsInner{}
	return &this
}

// NewTelephonyResponseCommandsInnerWithDefaults instantiates a new TelephonyResponseCommandsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelephonyResponseCommandsInnerWithDefaults() *TelephonyResponseCommandsInner {
	this := TelephonyResponseCommandsInner{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TelephonyResponseCommandsInner) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelephonyResponseCommandsInner) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TelephonyResponseCommandsInner) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TelephonyResponseCommandsInner) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *TelephonyResponseCommandsInner) GetValue() []TelephonyResponseCommandsInnerValueInner {
	if o == nil || o.Value == nil {
		var ret []TelephonyResponseCommandsInnerValueInner
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelephonyResponseCommandsInner) GetValueOk() ([]TelephonyResponseCommandsInnerValueInner, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *TelephonyResponseCommandsInner) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []TelephonyResponseCommandsInnerValueInner and assigns it to the Value field.
func (o *TelephonyResponseCommandsInner) SetValue(v []TelephonyResponseCommandsInnerValueInner) {
	o.Value = v
}

func (o TelephonyResponseCommandsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TelephonyResponseCommandsInner) UnmarshalJSON(bytes []byte) (err error) {
	varTelephonyResponseCommandsInner := _TelephonyResponseCommandsInner{}

	err = json.Unmarshal(bytes, &varTelephonyResponseCommandsInner)
	if err == nil {
		*o = TelephonyResponseCommandsInner(varTelephonyResponseCommandsInner)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableTelephonyResponseCommandsInner struct {
	value *TelephonyResponseCommandsInner
	isSet bool
}

func (v NullableTelephonyResponseCommandsInner) Get() *TelephonyResponseCommandsInner {
	return v.value
}

func (v *NullableTelephonyResponseCommandsInner) Set(val *TelephonyResponseCommandsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableTelephonyResponseCommandsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableTelephonyResponseCommandsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelephonyResponseCommandsInner(val *TelephonyResponseCommandsInner) *NullableTelephonyResponseCommandsInner {
	return &NullableTelephonyResponseCommandsInner{value: val, isSet: true}
}

func (v NullableTelephonyResponseCommandsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelephonyResponseCommandsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

