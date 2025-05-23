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

// InlineHookRequestObjectRequestUrl The URL of the API endpoint
type InlineHookRequestObjectRequestUrl struct {
	// The URL value of the API endpoint
	Value *string `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InlineHookRequestObjectRequestUrl InlineHookRequestObjectRequestUrl

// NewInlineHookRequestObjectRequestUrl instantiates a new InlineHookRequestObjectRequestUrl object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineHookRequestObjectRequestUrl() *InlineHookRequestObjectRequestUrl {
	this := InlineHookRequestObjectRequestUrl{}
	return &this
}

// NewInlineHookRequestObjectRequestUrlWithDefaults instantiates a new InlineHookRequestObjectRequestUrl object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineHookRequestObjectRequestUrlWithDefaults() *InlineHookRequestObjectRequestUrl {
	this := InlineHookRequestObjectRequestUrl{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *InlineHookRequestObjectRequestUrl) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookRequestObjectRequestUrl) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *InlineHookRequestObjectRequestUrl) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *InlineHookRequestObjectRequestUrl) SetValue(v string) {
	o.Value = &v
}

func (o InlineHookRequestObjectRequestUrl) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *InlineHookRequestObjectRequestUrl) UnmarshalJSON(bytes []byte) (err error) {
	varInlineHookRequestObjectRequestUrl := _InlineHookRequestObjectRequestUrl{}

	err = json.Unmarshal(bytes, &varInlineHookRequestObjectRequestUrl)
	if err == nil {
		*o = InlineHookRequestObjectRequestUrl(varInlineHookRequestObjectRequestUrl)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableInlineHookRequestObjectRequestUrl struct {
	value *InlineHookRequestObjectRequestUrl
	isSet bool
}

func (v NullableInlineHookRequestObjectRequestUrl) Get() *InlineHookRequestObjectRequestUrl {
	return v.value
}

func (v *NullableInlineHookRequestObjectRequestUrl) Set(val *InlineHookRequestObjectRequestUrl) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineHookRequestObjectRequestUrl) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineHookRequestObjectRequestUrl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineHookRequestObjectRequestUrl(val *InlineHookRequestObjectRequestUrl) *NullableInlineHookRequestObjectRequestUrl {
	return &NullableInlineHookRequestObjectRequestUrl{value: val, isSet: true}
}

func (v NullableInlineHookRequestObjectRequestUrl) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineHookRequestObjectRequestUrl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

