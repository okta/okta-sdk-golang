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

// InlineHookOAuthClientSecretConfig struct for InlineHookOAuthClientSecretConfig
type InlineHookOAuthClientSecretConfig struct {
	ClientSecret *string `json:"clientSecret,omitempty"`
	AuthScheme *InlineHookChannelConfigAuthScheme `json:"authScheme,omitempty"`
	Headers []InlineHookChannelConfigHeaders `json:"headers,omitempty"`
	Method *string `json:"method,omitempty"`
	Uri *string `json:"uri,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InlineHookOAuthClientSecretConfig InlineHookOAuthClientSecretConfig

// NewInlineHookOAuthClientSecretConfig instantiates a new InlineHookOAuthClientSecretConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineHookOAuthClientSecretConfig() *InlineHookOAuthClientSecretConfig {
	this := InlineHookOAuthClientSecretConfig{}
	return &this
}

// NewInlineHookOAuthClientSecretConfigWithDefaults instantiates a new InlineHookOAuthClientSecretConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineHookOAuthClientSecretConfigWithDefaults() *InlineHookOAuthClientSecretConfig {
	this := InlineHookOAuthClientSecretConfig{}
	return &this
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *InlineHookOAuthClientSecretConfig) GetClientSecret() string {
	if o == nil || o.ClientSecret == nil {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookOAuthClientSecretConfig) GetClientSecretOk() (*string, bool) {
	if o == nil || o.ClientSecret == nil {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *InlineHookOAuthClientSecretConfig) HasClientSecret() bool {
	if o != nil && o.ClientSecret != nil {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *InlineHookOAuthClientSecretConfig) SetClientSecret(v string) {
	o.ClientSecret = &v
}

// GetAuthScheme returns the AuthScheme field value if set, zero value otherwise.
func (o *InlineHookOAuthClientSecretConfig) GetAuthScheme() InlineHookChannelConfigAuthScheme {
	if o == nil || o.AuthScheme == nil {
		var ret InlineHookChannelConfigAuthScheme
		return ret
	}
	return *o.AuthScheme
}

// GetAuthSchemeOk returns a tuple with the AuthScheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookOAuthClientSecretConfig) GetAuthSchemeOk() (*InlineHookChannelConfigAuthScheme, bool) {
	if o == nil || o.AuthScheme == nil {
		return nil, false
	}
	return o.AuthScheme, true
}

// HasAuthScheme returns a boolean if a field has been set.
func (o *InlineHookOAuthClientSecretConfig) HasAuthScheme() bool {
	if o != nil && o.AuthScheme != nil {
		return true
	}

	return false
}

// SetAuthScheme gets a reference to the given InlineHookChannelConfigAuthScheme and assigns it to the AuthScheme field.
func (o *InlineHookOAuthClientSecretConfig) SetAuthScheme(v InlineHookChannelConfigAuthScheme) {
	o.AuthScheme = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *InlineHookOAuthClientSecretConfig) GetHeaders() []InlineHookChannelConfigHeaders {
	if o == nil || o.Headers == nil {
		var ret []InlineHookChannelConfigHeaders
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookOAuthClientSecretConfig) GetHeadersOk() ([]InlineHookChannelConfigHeaders, bool) {
	if o == nil || o.Headers == nil {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *InlineHookOAuthClientSecretConfig) HasHeaders() bool {
	if o != nil && o.Headers != nil {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given []InlineHookChannelConfigHeaders and assigns it to the Headers field.
func (o *InlineHookOAuthClientSecretConfig) SetHeaders(v []InlineHookChannelConfigHeaders) {
	o.Headers = v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *InlineHookOAuthClientSecretConfig) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookOAuthClientSecretConfig) GetMethodOk() (*string, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *InlineHookOAuthClientSecretConfig) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *InlineHookOAuthClientSecretConfig) SetMethod(v string) {
	o.Method = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *InlineHookOAuthClientSecretConfig) GetUri() string {
	if o == nil || o.Uri == nil {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookOAuthClientSecretConfig) GetUriOk() (*string, bool) {
	if o == nil || o.Uri == nil {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *InlineHookOAuthClientSecretConfig) HasUri() bool {
	if o != nil && o.Uri != nil {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *InlineHookOAuthClientSecretConfig) SetUri(v string) {
	o.Uri = &v
}

func (o InlineHookOAuthClientSecretConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientSecret != nil {
		toSerialize["clientSecret"] = o.ClientSecret
	}
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

func (o *InlineHookOAuthClientSecretConfig) UnmarshalJSON(bytes []byte) (err error) {
	varInlineHookOAuthClientSecretConfig := _InlineHookOAuthClientSecretConfig{}

	err = json.Unmarshal(bytes, &varInlineHookOAuthClientSecretConfig)
	if err == nil {
		*o = InlineHookOAuthClientSecretConfig(varInlineHookOAuthClientSecretConfig)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "clientSecret")
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

type NullableInlineHookOAuthClientSecretConfig struct {
	value *InlineHookOAuthClientSecretConfig
	isSet bool
}

func (v NullableInlineHookOAuthClientSecretConfig) Get() *InlineHookOAuthClientSecretConfig {
	return v.value
}

func (v *NullableInlineHookOAuthClientSecretConfig) Set(val *InlineHookOAuthClientSecretConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineHookOAuthClientSecretConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineHookOAuthClientSecretConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineHookOAuthClientSecretConfig(val *InlineHookOAuthClientSecretConfig) *NullableInlineHookOAuthClientSecretConfig {
	return &NullableInlineHookOAuthClientSecretConfig{value: val, isSet: true}
}

func (v NullableInlineHookOAuthClientSecretConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineHookOAuthClientSecretConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

