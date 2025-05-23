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

// EventHookChannelConfig struct for EventHookChannelConfig
type EventHookChannelConfig struct {
	AuthScheme *EventHookChannelConfigAuthScheme `json:"authScheme,omitempty"`
	// Optional list of key/value pairs for headers that can be sent with the request to the external service. For example, `X-Other-Header` is an example of an optional header, with a value of `my-header-value`, that you want Okta to pass to your external service.
	Headers []EventHookChannelConfigHeader `json:"headers,omitempty"`
	// The method of the Okta event hook request
	Method *string `json:"method,omitempty"`
	// The external service endpoint called to execute the event hook handler
	Uri string `json:"uri"`
	AdditionalProperties map[string]interface{}
}

type _EventHookChannelConfig EventHookChannelConfig

// NewEventHookChannelConfig instantiates a new EventHookChannelConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventHookChannelConfig(uri string) *EventHookChannelConfig {
	this := EventHookChannelConfig{}
	this.Uri = uri
	return &this
}

// NewEventHookChannelConfigWithDefaults instantiates a new EventHookChannelConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventHookChannelConfigWithDefaults() *EventHookChannelConfig {
	this := EventHookChannelConfig{}
	return &this
}

// GetAuthScheme returns the AuthScheme field value if set, zero value otherwise.
func (o *EventHookChannelConfig) GetAuthScheme() EventHookChannelConfigAuthScheme {
	if o == nil || o.AuthScheme == nil {
		var ret EventHookChannelConfigAuthScheme
		return ret
	}
	return *o.AuthScheme
}

// GetAuthSchemeOk returns a tuple with the AuthScheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventHookChannelConfig) GetAuthSchemeOk() (*EventHookChannelConfigAuthScheme, bool) {
	if o == nil || o.AuthScheme == nil {
		return nil, false
	}
	return o.AuthScheme, true
}

// HasAuthScheme returns a boolean if a field has been set.
func (o *EventHookChannelConfig) HasAuthScheme() bool {
	if o != nil && o.AuthScheme != nil {
		return true
	}

	return false
}

// SetAuthScheme gets a reference to the given EventHookChannelConfigAuthScheme and assigns it to the AuthScheme field.
func (o *EventHookChannelConfig) SetAuthScheme(v EventHookChannelConfigAuthScheme) {
	o.AuthScheme = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *EventHookChannelConfig) GetHeaders() []EventHookChannelConfigHeader {
	if o == nil || o.Headers == nil {
		var ret []EventHookChannelConfigHeader
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventHookChannelConfig) GetHeadersOk() ([]EventHookChannelConfigHeader, bool) {
	if o == nil || o.Headers == nil {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *EventHookChannelConfig) HasHeaders() bool {
	if o != nil && o.Headers != nil {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given []EventHookChannelConfigHeader and assigns it to the Headers field.
func (o *EventHookChannelConfig) SetHeaders(v []EventHookChannelConfigHeader) {
	o.Headers = v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *EventHookChannelConfig) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventHookChannelConfig) GetMethodOk() (*string, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *EventHookChannelConfig) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *EventHookChannelConfig) SetMethod(v string) {
	o.Method = &v
}

// GetUri returns the Uri field value
func (o *EventHookChannelConfig) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *EventHookChannelConfig) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *EventHookChannelConfig) SetUri(v string) {
	o.Uri = v
}

func (o EventHookChannelConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthScheme != nil {
		toSerialize["authScheme"] = o.AuthScheme
	}
	if o.Headers != nil {
		toSerialize["headers"] = o.Headers
	}
	if o.Method != nil {
		toSerialize["method"] = o.Method
	}
	if true {
		toSerialize["uri"] = o.Uri
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EventHookChannelConfig) UnmarshalJSON(bytes []byte) (err error) {
	varEventHookChannelConfig := _EventHookChannelConfig{}

	err = json.Unmarshal(bytes, &varEventHookChannelConfig)
	if err == nil {
		*o = EventHookChannelConfig(varEventHookChannelConfig)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "authScheme")
		delete(additionalProperties, "headers")
		delete(additionalProperties, "method")
		delete(additionalProperties, "uri")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableEventHookChannelConfig struct {
	value *EventHookChannelConfig
	isSet bool
}

func (v NullableEventHookChannelConfig) Get() *EventHookChannelConfig {
	return v.value
}

func (v *NullableEventHookChannelConfig) Set(val *EventHookChannelConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableEventHookChannelConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableEventHookChannelConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventHookChannelConfig(val *EventHookChannelConfig) *NullableEventHookChannelConfig {
	return &NullableEventHookChannelConfig{value: val, isSet: true}
}

func (v NullableEventHookChannelConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventHookChannelConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

