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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the InlineHookChannelConfigHeaders type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineHookChannelConfigHeaders{}

// InlineHookChannelConfigHeaders struct for InlineHookChannelConfigHeaders
type InlineHookChannelConfigHeaders struct {
	// The optional field or header name
	Key *string `json:"key,omitempty"`
	// The value for the key
	Value                *string `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InlineHookChannelConfigHeaders InlineHookChannelConfigHeaders

// NewInlineHookChannelConfigHeaders instantiates a new InlineHookChannelConfigHeaders object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineHookChannelConfigHeaders() *InlineHookChannelConfigHeaders {
	this := InlineHookChannelConfigHeaders{}
	return &this
}

// NewInlineHookChannelConfigHeadersWithDefaults instantiates a new InlineHookChannelConfigHeaders object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineHookChannelConfigHeadersWithDefaults() *InlineHookChannelConfigHeaders {
	this := InlineHookChannelConfigHeaders{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *InlineHookChannelConfigHeaders) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookChannelConfigHeaders) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *InlineHookChannelConfigHeaders) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *InlineHookChannelConfigHeaders) SetKey(v string) {
	o.Key = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *InlineHookChannelConfigHeaders) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookChannelConfigHeaders) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *InlineHookChannelConfigHeaders) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *InlineHookChannelConfigHeaders) SetValue(v string) {
	o.Value = &v
}

func (o InlineHookChannelConfigHeaders) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineHookChannelConfigHeaders) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InlineHookChannelConfigHeaders) UnmarshalJSON(data []byte) (err error) {
	varInlineHookChannelConfigHeaders := _InlineHookChannelConfigHeaders{}

	err = json.Unmarshal(data, &varInlineHookChannelConfigHeaders)

	if err != nil {
		return err
	}

	*o = InlineHookChannelConfigHeaders(varInlineHookChannelConfigHeaders)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "key")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInlineHookChannelConfigHeaders struct {
	value *InlineHookChannelConfigHeaders
	isSet bool
}

func (v NullableInlineHookChannelConfigHeaders) Get() *InlineHookChannelConfigHeaders {
	return v.value
}

func (v *NullableInlineHookChannelConfigHeaders) Set(val *InlineHookChannelConfigHeaders) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineHookChannelConfigHeaders) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineHookChannelConfigHeaders) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineHookChannelConfigHeaders(val *InlineHookChannelConfigHeaders) *NullableInlineHookChannelConfigHeaders {
	return &NullableInlineHookChannelConfigHeaders{value: val, isSet: true}
}

func (v NullableInlineHookChannelConfigHeaders) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineHookChannelConfigHeaders) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
