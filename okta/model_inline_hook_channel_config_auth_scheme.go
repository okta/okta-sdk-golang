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

// InlineHookChannelConfigAuthScheme struct for InlineHookChannelConfigAuthScheme
type InlineHookChannelConfigAuthScheme struct {
	Key *string `json:"key,omitempty"`
	Type *string `json:"type,omitempty"`
	Value *string `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InlineHookChannelConfigAuthScheme InlineHookChannelConfigAuthScheme

// NewInlineHookChannelConfigAuthScheme instantiates a new InlineHookChannelConfigAuthScheme object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineHookChannelConfigAuthScheme() *InlineHookChannelConfigAuthScheme {
	this := InlineHookChannelConfigAuthScheme{}
	return &this
}

// NewInlineHookChannelConfigAuthSchemeWithDefaults instantiates a new InlineHookChannelConfigAuthScheme object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineHookChannelConfigAuthSchemeWithDefaults() *InlineHookChannelConfigAuthScheme {
	this := InlineHookChannelConfigAuthScheme{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *InlineHookChannelConfigAuthScheme) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookChannelConfigAuthScheme) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *InlineHookChannelConfigAuthScheme) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *InlineHookChannelConfigAuthScheme) SetKey(v string) {
	o.Key = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InlineHookChannelConfigAuthScheme) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookChannelConfigAuthScheme) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InlineHookChannelConfigAuthScheme) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InlineHookChannelConfigAuthScheme) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *InlineHookChannelConfigAuthScheme) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookChannelConfigAuthScheme) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *InlineHookChannelConfigAuthScheme) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *InlineHookChannelConfigAuthScheme) SetValue(v string) {
	o.Value = &v
}

func (o InlineHookChannelConfigAuthScheme) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
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

func (o *InlineHookChannelConfigAuthScheme) UnmarshalJSON(bytes []byte) (err error) {
	varInlineHookChannelConfigAuthScheme := _InlineHookChannelConfigAuthScheme{}

	err = json.Unmarshal(bytes, &varInlineHookChannelConfigAuthScheme)
	if err == nil {
		*o = InlineHookChannelConfigAuthScheme(varInlineHookChannelConfigAuthScheme)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "key")
		delete(additionalProperties, "type")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableInlineHookChannelConfigAuthScheme struct {
	value *InlineHookChannelConfigAuthScheme
	isSet bool
}

func (v NullableInlineHookChannelConfigAuthScheme) Get() *InlineHookChannelConfigAuthScheme {
	return v.value
}

func (v *NullableInlineHookChannelConfigAuthScheme) Set(val *InlineHookChannelConfigAuthScheme) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineHookChannelConfigAuthScheme) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineHookChannelConfigAuthScheme) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineHookChannelConfigAuthScheme(val *InlineHookChannelConfigAuthScheme) *NullableInlineHookChannelConfigAuthScheme {
	return &NullableInlineHookChannelConfigAuthScheme{value: val, isSet: true}
}

func (v NullableInlineHookChannelConfigAuthScheme) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineHookChannelConfigAuthScheme) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

