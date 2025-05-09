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

// InlineHookChannelConfig struct for InlineHookChannelConfig
type InlineHookChannelConfig struct {
	AuthScheme *InlineHookChannelConfigAuthScheme `json:"authScheme,omitempty"`
	Headers []InlineHookChannelConfigHeaders `json:"headers,omitempty"`
	Method *string `json:"method,omitempty"`
	Uri *string `json:"uri,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InlineHookChannelConfig InlineHookChannelConfig

// NewInlineHookChannelConfig instantiates a new InlineHookChannelConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineHookChannelConfig() *InlineHookChannelConfig {
	this := InlineHookChannelConfig{}
	return &this
}

// NewInlineHookChannelConfigWithDefaults instantiates a new InlineHookChannelConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineHookChannelConfigWithDefaults() *InlineHookChannelConfig {
	this := InlineHookChannelConfig{}
	return &this
}

// GetAuthScheme returns the AuthScheme field value if set, zero value otherwise.
func (o *InlineHookChannelConfig) GetAuthScheme() InlineHookChannelConfigAuthScheme {
	if o == nil || o.AuthScheme == nil {
		var ret InlineHookChannelConfigAuthScheme
		return ret
	}
	return *o.AuthScheme
}

// GetAuthSchemeOk returns a tuple with the AuthScheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookChannelConfig) GetAuthSchemeOk() (*InlineHookChannelConfigAuthScheme, bool) {
	if o == nil || o.AuthScheme == nil {
		return nil, false
	}
	return o.AuthScheme, true
}

// HasAuthScheme returns a boolean if a field has been set.
func (o *InlineHookChannelConfig) HasAuthScheme() bool {
	if o != nil && o.AuthScheme != nil {
		return true
	}

	return false
}

// SetAuthScheme gets a reference to the given InlineHookChannelConfigAuthScheme and assigns it to the AuthScheme field.
func (o *InlineHookChannelConfig) SetAuthScheme(v InlineHookChannelConfigAuthScheme) {
	o.AuthScheme = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *InlineHookChannelConfig) GetHeaders() []InlineHookChannelConfigHeaders {
	if o == nil || o.Headers == nil {
		var ret []InlineHookChannelConfigHeaders
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookChannelConfig) GetHeadersOk() ([]InlineHookChannelConfigHeaders, bool) {
	if o == nil || o.Headers == nil {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *InlineHookChannelConfig) HasHeaders() bool {
	if o != nil && o.Headers != nil {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given []InlineHookChannelConfigHeaders and assigns it to the Headers field.
func (o *InlineHookChannelConfig) SetHeaders(v []InlineHookChannelConfigHeaders) {
	o.Headers = v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *InlineHookChannelConfig) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookChannelConfig) GetMethodOk() (*string, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *InlineHookChannelConfig) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *InlineHookChannelConfig) SetMethod(v string) {
	o.Method = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *InlineHookChannelConfig) GetUri() string {
	if o == nil || o.Uri == nil {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookChannelConfig) GetUriOk() (*string, bool) {
	if o == nil || o.Uri == nil {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *InlineHookChannelConfig) HasUri() bool {
	if o != nil && o.Uri != nil {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *InlineHookChannelConfig) SetUri(v string) {
	o.Uri = &v
}

func (o InlineHookChannelConfig) MarshalJSON() ([]byte, error) {
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
	if o.Uri != nil {
		toSerialize["uri"] = o.Uri
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *InlineHookChannelConfig) UnmarshalJSON(bytes []byte) (err error) {
	varInlineHookChannelConfig := _InlineHookChannelConfig{}

	err = json.Unmarshal(bytes, &varInlineHookChannelConfig)
	if err == nil {
		*o = InlineHookChannelConfig(varInlineHookChannelConfig)
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

type NullableInlineHookChannelConfig struct {
	value *InlineHookChannelConfig
	isSet bool
}

func (v NullableInlineHookChannelConfig) Get() *InlineHookChannelConfig {
	return v.value
}

func (v *NullableInlineHookChannelConfig) Set(val *InlineHookChannelConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineHookChannelConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineHookChannelConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineHookChannelConfig(val *InlineHookChannelConfig) *NullableInlineHookChannelConfig {
	return &NullableInlineHookChannelConfig{value: val, isSet: true}
}

func (v NullableInlineHookChannelConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineHookChannelConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

