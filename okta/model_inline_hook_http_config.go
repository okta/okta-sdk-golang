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

// checks if the InlineHookHttpConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineHookHttpConfig{}

// InlineHookHttpConfig struct for InlineHookHttpConfig
type InlineHookHttpConfig struct {
	AuthScheme NullableInlineHookChannelConfigAuthSchemeResponse `json:"authScheme,omitempty"`
	// An optional list of key/value pairs for headers that you can send with the request to the external service
	Headers []InlineHookChannelConfigHeaders `json:"headers,omitempty"`
	// The method of the Okta inline hook request
	Method *string `json:"method,omitempty"`
	// The external service endpoint that executes the inline hook handler. It must begin with `https://` and be reachable by Okta. No white space is allowed in the URI.
	Uri                  *string `json:"uri,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InlineHookHttpConfig InlineHookHttpConfig

// NewInlineHookHttpConfig instantiates a new InlineHookHttpConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineHookHttpConfig() *InlineHookHttpConfig {
	this := InlineHookHttpConfig{}
	return &this
}

// NewInlineHookHttpConfigWithDefaults instantiates a new InlineHookHttpConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineHookHttpConfigWithDefaults() *InlineHookHttpConfig {
	this := InlineHookHttpConfig{}
	return &this
}

// GetAuthScheme returns the AuthScheme field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineHookHttpConfig) GetAuthScheme() InlineHookChannelConfigAuthSchemeResponse {
	if o == nil || IsNil(o.AuthScheme.Get()) {
		var ret InlineHookChannelConfigAuthSchemeResponse
		return ret
	}
	return *o.AuthScheme.Get()
}

// GetAuthSchemeOk returns a tuple with the AuthScheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineHookHttpConfig) GetAuthSchemeOk() (*InlineHookChannelConfigAuthSchemeResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthScheme.Get(), o.AuthScheme.IsSet()
}

// HasAuthScheme returns a boolean if a field has been set.
func (o *InlineHookHttpConfig) HasAuthScheme() bool {
	if o != nil && o.AuthScheme.IsSet() {
		return true
	}

	return false
}

// SetAuthScheme gets a reference to the given NullableInlineHookChannelConfigAuthSchemeResponse and assigns it to the AuthScheme field.
func (o *InlineHookHttpConfig) SetAuthScheme(v InlineHookChannelConfigAuthSchemeResponse) {
	o.AuthScheme.Set(&v)
}

// SetAuthSchemeNil sets the value for AuthScheme to be an explicit nil
func (o *InlineHookHttpConfig) SetAuthSchemeNil() {
	o.AuthScheme.Set(nil)
}

// UnsetAuthScheme ensures that no value is present for AuthScheme, not even an explicit nil
func (o *InlineHookHttpConfig) UnsetAuthScheme() {
	o.AuthScheme.Unset()
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *InlineHookHttpConfig) GetHeaders() []InlineHookChannelConfigHeaders {
	if o == nil || IsNil(o.Headers) {
		var ret []InlineHookChannelConfigHeaders
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookHttpConfig) GetHeadersOk() ([]InlineHookChannelConfigHeaders, bool) {
	if o == nil || IsNil(o.Headers) {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *InlineHookHttpConfig) HasHeaders() bool {
	if o != nil && !IsNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given []InlineHookChannelConfigHeaders and assigns it to the Headers field.
func (o *InlineHookHttpConfig) SetHeaders(v []InlineHookChannelConfigHeaders) {
	o.Headers = v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *InlineHookHttpConfig) GetMethod() string {
	if o == nil || IsNil(o.Method) {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookHttpConfig) GetMethodOk() (*string, bool) {
	if o == nil || IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *InlineHookHttpConfig) HasMethod() bool {
	if o != nil && !IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *InlineHookHttpConfig) SetMethod(v string) {
	o.Method = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *InlineHookHttpConfig) GetUri() string {
	if o == nil || IsNil(o.Uri) {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookHttpConfig) GetUriOk() (*string, bool) {
	if o == nil || IsNil(o.Uri) {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *InlineHookHttpConfig) HasUri() bool {
	if o != nil && !IsNil(o.Uri) {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *InlineHookHttpConfig) SetUri(v string) {
	o.Uri = &v
}

func (o InlineHookHttpConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineHookHttpConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthScheme.IsSet() {
		toSerialize["authScheme"] = o.AuthScheme.Get()
	}
	if !IsNil(o.Headers) {
		toSerialize["headers"] = o.Headers
	}
	if !IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}
	if !IsNil(o.Uri) {
		toSerialize["uri"] = o.Uri
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InlineHookHttpConfig) UnmarshalJSON(data []byte) (err error) {
	varInlineHookHttpConfig := _InlineHookHttpConfig{}

	err = json.Unmarshal(data, &varInlineHookHttpConfig)

	if err != nil {
		return err
	}

	*o = InlineHookHttpConfig(varInlineHookHttpConfig)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "authScheme")
		delete(additionalProperties, "headers")
		delete(additionalProperties, "method")
		delete(additionalProperties, "uri")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInlineHookHttpConfig struct {
	value *InlineHookHttpConfig
	isSet bool
}

func (v NullableInlineHookHttpConfig) Get() *InlineHookHttpConfig {
	return v.value
}

func (v *NullableInlineHookHttpConfig) Set(val *InlineHookHttpConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineHookHttpConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineHookHttpConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineHookHttpConfig(val *InlineHookHttpConfig) *NullableInlineHookHttpConfig {
	return &NullableInlineHookHttpConfig{value: val, isSet: true}
}

func (v NullableInlineHookHttpConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineHookHttpConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
