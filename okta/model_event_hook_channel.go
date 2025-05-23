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

// EventHookChannel struct for EventHookChannel
type EventHookChannel struct {
	Config EventHookChannelConfig `json:"config"`
	// The channel type. Currently supports `HTTP`.
	Type string `json:"type"`
	// Version of the channel. Currently the only supported version is `1.0.0``.
	Version string `json:"version"`
	AdditionalProperties map[string]interface{}
}

type _EventHookChannel EventHookChannel

// NewEventHookChannel instantiates a new EventHookChannel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventHookChannel(config EventHookChannelConfig, type_ string, version string) *EventHookChannel {
	this := EventHookChannel{}
	this.Config = config
	this.Type = type_
	this.Version = version
	return &this
}

// NewEventHookChannelWithDefaults instantiates a new EventHookChannel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventHookChannelWithDefaults() *EventHookChannel {
	this := EventHookChannel{}
	return &this
}

// GetConfig returns the Config field value
func (o *EventHookChannel) GetConfig() EventHookChannelConfig {
	if o == nil {
		var ret EventHookChannelConfig
		return ret
	}

	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *EventHookChannel) GetConfigOk() (*EventHookChannelConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Config, true
}

// SetConfig sets field value
func (o *EventHookChannel) SetConfig(v EventHookChannelConfig) {
	o.Config = v
}

// GetType returns the Type field value
func (o *EventHookChannel) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EventHookChannel) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *EventHookChannel) SetType(v string) {
	o.Type = v
}

// GetVersion returns the Version field value
func (o *EventHookChannel) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *EventHookChannel) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *EventHookChannel) SetVersion(v string) {
	o.Version = v
}

func (o EventHookChannel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["config"] = o.Config
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EventHookChannel) UnmarshalJSON(bytes []byte) (err error) {
	varEventHookChannel := _EventHookChannel{}

	err = json.Unmarshal(bytes, &varEventHookChannel)
	if err == nil {
		*o = EventHookChannel(varEventHookChannel)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "config")
		delete(additionalProperties, "type")
		delete(additionalProperties, "version")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableEventHookChannel struct {
	value *EventHookChannel
	isSet bool
}

func (v NullableEventHookChannel) Get() *EventHookChannel {
	return v.value
}

func (v *NullableEventHookChannel) Set(val *EventHookChannel) {
	v.value = val
	v.isSet = true
}

func (v NullableEventHookChannel) IsSet() bool {
	return v.isSet
}

func (v *NullableEventHookChannel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventHookChannel(val *EventHookChannel) *NullableEventHookChannel {
	return &NullableEventHookChannel{value: val, isSet: true}
}

func (v NullableEventHookChannel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventHookChannel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

