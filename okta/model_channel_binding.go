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

// ChannelBinding struct for ChannelBinding
type ChannelBinding struct {
	Required *string `json:"required,omitempty"`
	Style *string `json:"style,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ChannelBinding ChannelBinding

// NewChannelBinding instantiates a new ChannelBinding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelBinding() *ChannelBinding {
	this := ChannelBinding{}
	return &this
}

// NewChannelBindingWithDefaults instantiates a new ChannelBinding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelBindingWithDefaults() *ChannelBinding {
	this := ChannelBinding{}
	return &this
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *ChannelBinding) GetRequired() string {
	if o == nil || o.Required == nil {
		var ret string
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelBinding) GetRequiredOk() (*string, bool) {
	if o == nil || o.Required == nil {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *ChannelBinding) HasRequired() bool {
	if o != nil && o.Required != nil {
		return true
	}

	return false
}

// SetRequired gets a reference to the given string and assigns it to the Required field.
func (o *ChannelBinding) SetRequired(v string) {
	o.Required = &v
}

// GetStyle returns the Style field value if set, zero value otherwise.
func (o *ChannelBinding) GetStyle() string {
	if o == nil || o.Style == nil {
		var ret string
		return ret
	}
	return *o.Style
}

// GetStyleOk returns a tuple with the Style field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelBinding) GetStyleOk() (*string, bool) {
	if o == nil || o.Style == nil {
		return nil, false
	}
	return o.Style, true
}

// HasStyle returns a boolean if a field has been set.
func (o *ChannelBinding) HasStyle() bool {
	if o != nil && o.Style != nil {
		return true
	}

	return false
}

// SetStyle gets a reference to the given string and assigns it to the Style field.
func (o *ChannelBinding) SetStyle(v string) {
	o.Style = &v
}

func (o ChannelBinding) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Required != nil {
		toSerialize["required"] = o.Required
	}
	if o.Style != nil {
		toSerialize["style"] = o.Style
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ChannelBinding) UnmarshalJSON(bytes []byte) (err error) {
	varChannelBinding := _ChannelBinding{}

	err = json.Unmarshal(bytes, &varChannelBinding)
	if err == nil {
		*o = ChannelBinding(varChannelBinding)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "required")
		delete(additionalProperties, "style")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableChannelBinding struct {
	value *ChannelBinding
	isSet bool
}

func (v NullableChannelBinding) Get() *ChannelBinding {
	return v.value
}

func (v *NullableChannelBinding) Set(val *ChannelBinding) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelBinding) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelBinding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelBinding(val *ChannelBinding) *NullableChannelBinding {
	return &NullableChannelBinding{value: val, isSet: true}
}

func (v NullableChannelBinding) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelBinding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

