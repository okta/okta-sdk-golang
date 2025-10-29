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

// checks if the InlineHookRequestObjectUrl type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineHookRequestObjectUrl{}

// InlineHookRequestObjectUrl The URL of the API endpoint
type InlineHookRequestObjectUrl struct {
	// The URL value of the API endpoint
	Value                *string `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InlineHookRequestObjectUrl InlineHookRequestObjectUrl

// NewInlineHookRequestObjectUrl instantiates a new InlineHookRequestObjectUrl object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineHookRequestObjectUrl() *InlineHookRequestObjectUrl {
	this := InlineHookRequestObjectUrl{}
	return &this
}

// NewInlineHookRequestObjectUrlWithDefaults instantiates a new InlineHookRequestObjectUrl object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineHookRequestObjectUrlWithDefaults() *InlineHookRequestObjectUrl {
	this := InlineHookRequestObjectUrl{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *InlineHookRequestObjectUrl) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookRequestObjectUrl) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *InlineHookRequestObjectUrl) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *InlineHookRequestObjectUrl) SetValue(v string) {
	o.Value = &v
}

func (o InlineHookRequestObjectUrl) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineHookRequestObjectUrl) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InlineHookRequestObjectUrl) UnmarshalJSON(data []byte) (err error) {
	varInlineHookRequestObjectUrl := _InlineHookRequestObjectUrl{}

	err = json.Unmarshal(data, &varInlineHookRequestObjectUrl)

	if err != nil {
		return err
	}

	*o = InlineHookRequestObjectUrl(varInlineHookRequestObjectUrl)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInlineHookRequestObjectUrl struct {
	value *InlineHookRequestObjectUrl
	isSet bool
}

func (v NullableInlineHookRequestObjectUrl) Get() *InlineHookRequestObjectUrl {
	return v.value
}

func (v *NullableInlineHookRequestObjectUrl) Set(val *InlineHookRequestObjectUrl) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineHookRequestObjectUrl) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineHookRequestObjectUrl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineHookRequestObjectUrl(val *InlineHookRequestObjectUrl) *NullableInlineHookRequestObjectUrl {
	return &NullableInlineHookRequestObjectUrl{value: val, isSet: true}
}

func (v NullableInlineHookRequestObjectUrl) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineHookRequestObjectUrl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
