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

// InlineHookCreate An inline hook object that specifies the details of the inline hook
type InlineHookCreate struct {
	Channel *InlineHookChannelCreate `json:"channel,omitempty"`
	// The display name of the inline hook
	Name *string `json:"name,omitempty"`
	// One of the inline hook types
	Type *string `json:"type,omitempty"`
	// Version of the inline hook type. The currently supported version is `1.0.0`.
	Version *string `json:"version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InlineHookCreate InlineHookCreate

// NewInlineHookCreate instantiates a new InlineHookCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineHookCreate() *InlineHookCreate {
	this := InlineHookCreate{}
	return &this
}

// NewInlineHookCreateWithDefaults instantiates a new InlineHookCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineHookCreateWithDefaults() *InlineHookCreate {
	this := InlineHookCreate{}
	return &this
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *InlineHookCreate) GetChannel() InlineHookChannelCreate {
	if o == nil || o.Channel == nil {
		var ret InlineHookChannelCreate
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookCreate) GetChannelOk() (*InlineHookChannelCreate, bool) {
	if o == nil || o.Channel == nil {
		return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *InlineHookCreate) HasChannel() bool {
	if o != nil && o.Channel != nil {
		return true
	}

	return false
}

// SetChannel gets a reference to the given InlineHookChannelCreate and assigns it to the Channel field.
func (o *InlineHookCreate) SetChannel(v InlineHookChannelCreate) {
	o.Channel = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineHookCreate) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookCreate) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineHookCreate) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineHookCreate) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InlineHookCreate) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookCreate) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InlineHookCreate) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InlineHookCreate) SetType(v string) {
	o.Type = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *InlineHookCreate) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookCreate) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *InlineHookCreate) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *InlineHookCreate) SetVersion(v string) {
	o.Version = &v
}

func (o InlineHookCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Channel != nil {
		toSerialize["channel"] = o.Channel
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *InlineHookCreate) UnmarshalJSON(bytes []byte) (err error) {
	varInlineHookCreate := _InlineHookCreate{}

	err = json.Unmarshal(bytes, &varInlineHookCreate)
	if err == nil {
		*o = InlineHookCreate(varInlineHookCreate)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "channel")
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		delete(additionalProperties, "version")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableInlineHookCreate struct {
	value *InlineHookCreate
	isSet bool
}

func (v NullableInlineHookCreate) Get() *InlineHookCreate {
	return v.value
}

func (v *NullableInlineHookCreate) Set(val *InlineHookCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineHookCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineHookCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineHookCreate(val *InlineHookCreate) *NullableInlineHookCreate {
	return &NullableInlineHookCreate{value: val, isSet: true}
}

func (v NullableInlineHookCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineHookCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

