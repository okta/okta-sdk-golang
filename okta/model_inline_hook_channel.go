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

// checks if the InlineHookChannel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineHookChannel{}

// InlineHookChannel struct for InlineHookChannel
type InlineHookChannel struct {
	Type *string `json:"type,omitempty"`
	// Version of the inline hook type. The currently supported version is `1.0.0`.
	Version              *string `json:"version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InlineHookChannel InlineHookChannel

// NewInlineHookChannel instantiates a new InlineHookChannel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineHookChannel() *InlineHookChannel {
	this := InlineHookChannel{}
	return &this
}

// NewInlineHookChannelWithDefaults instantiates a new InlineHookChannel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineHookChannelWithDefaults() *InlineHookChannel {
	this := InlineHookChannel{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InlineHookChannel) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookChannel) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InlineHookChannel) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InlineHookChannel) SetType(v string) {
	o.Type = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *InlineHookChannel) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookChannel) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *InlineHookChannel) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *InlineHookChannel) SetVersion(v string) {
	o.Version = &v
}

func (o InlineHookChannel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineHookChannel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InlineHookChannel) UnmarshalJSON(data []byte) (err error) {
	varInlineHookChannel := _InlineHookChannel{}

	err = json.Unmarshal(data, &varInlineHookChannel)

	if err != nil {
		return err
	}

	*o = InlineHookChannel(varInlineHookChannel)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInlineHookChannel struct {
	value *InlineHookChannel
	isSet bool
}

func (v NullableInlineHookChannel) Get() *InlineHookChannel {
	return v.value
}

func (v *NullableInlineHookChannel) Set(val *InlineHookChannel) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineHookChannel) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineHookChannel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineHookChannel(val *InlineHookChannel) *NullableInlineHookChannel {
	return &NullableInlineHookChannel{value: val, isSet: true}
}

func (v NullableInlineHookChannel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineHookChannel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
