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

// InlineHookOAuthBasicConfig struct for InlineHookOAuthBasicConfig
type InlineHookOAuthBasicConfig struct {
	AuthType *string `json:"authType,omitempty"`
	ClientId *string `json:"clientId,omitempty"`
	Scope *string `json:"scope,omitempty"`
	TokenUrl *string `json:"tokenUrl,omitempty"`
	AuthScheme *InlineHookChannelConfigAuthScheme `json:"authScheme,omitempty"`
	Headers []InlineHookChannelConfigHeaders `json:"headers,omitempty"`
	Method *string `json:"method,omitempty"`
	Uri *string `json:"uri,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InlineHookOAuthBasicConfig InlineHookOAuthBasicConfig

// NewInlineHookOAuthBasicConfig instantiates a new InlineHookOAuthBasicConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineHookOAuthBasicConfig() *InlineHookOAuthBasicConfig {
	this := InlineHookOAuthBasicConfig{}
	return &this
}

// NewInlineHookOAuthBasicConfigWithDefaults instantiates a new InlineHookOAuthBasicConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineHookOAuthBasicConfigWithDefaults() *InlineHookOAuthBasicConfig {
	this := InlineHookOAuthBasicConfig{}
	return &this
}

// GetAuthType returns the AuthType field value if set, zero value otherwise.
func (o *InlineHookOAuthBasicConfig) GetAuthType() string {
	if o == nil || o.AuthType == nil {
		var ret string
		return ret
	}
	return *o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookOAuthBasicConfig) GetAuthTypeOk() (*string, bool) {
	if o == nil || o.AuthType == nil {
		return nil, false
	}
	return o.AuthType, true
}

// HasAuthType returns a boolean if a field has been set.
func (o *InlineHookOAuthBasicConfig) HasAuthType() bool {
	if o != nil && o.AuthType != nil {
		return true
	}

	return false
}

// SetAuthType gets a reference to the given string and assigns it to the AuthType field.
func (o *InlineHookOAuthBasicConfig) SetAuthType(v string) {
	o.AuthType = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *InlineHookOAuthBasicConfig) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookOAuthBasicConfig) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *InlineHookOAuthBasicConfig) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *InlineHookOAuthBasicConfig) SetClientId(v string) {
	o.ClientId = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *InlineHookOAuthBasicConfig) GetScope() string {
	if o == nil || o.Scope == nil {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookOAuthBasicConfig) GetScopeOk() (*string, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *InlineHookOAuthBasicConfig) HasScope() bool {
	if o != nil && o.Scope != nil {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *InlineHookOAuthBasicConfig) SetScope(v string) {
	o.Scope = &v
}

// GetTokenUrl returns the TokenUrl field value if set, zero value otherwise.
func (o *InlineHookOAuthBasicConfig) GetTokenUrl() string {
	if o == nil || o.TokenUrl == nil {
		var ret string
		return ret
	}
	return *o.TokenUrl
}

// GetTokenUrlOk returns a tuple with the TokenUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookOAuthBasicConfig) GetTokenUrlOk() (*string, bool) {
	if o == nil || o.TokenUrl == nil {
		return nil, false
	}
	return o.TokenUrl, true
}

// HasTokenUrl returns a boolean if a field has been set.
func (o *InlineHookOAuthBasicConfig) HasTokenUrl() bool {
	if o != nil && o.TokenUrl != nil {
		return true
	}

	return false
}

// SetTokenUrl gets a reference to the given string and assigns it to the TokenUrl field.
func (o *InlineHookOAuthBasicConfig) SetTokenUrl(v string) {
	o.TokenUrl = &v
}

// GetAuthScheme returns the AuthScheme field value if set, zero value otherwise.
func (o *InlineHookOAuthBasicConfig) GetAuthScheme() InlineHookChannelConfigAuthScheme {
	if o == nil || o.AuthScheme == nil {
		var ret InlineHookChannelConfigAuthScheme
		return ret
	}
	return *o.AuthScheme
}

// GetAuthSchemeOk returns a tuple with the AuthScheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookOAuthBasicConfig) GetAuthSchemeOk() (*InlineHookChannelConfigAuthScheme, bool) {
	if o == nil || o.AuthScheme == nil {
		return nil, false
	}
	return o.AuthScheme, true
}

// HasAuthScheme returns a boolean if a field has been set.
func (o *InlineHookOAuthBasicConfig) HasAuthScheme() bool {
	if o != nil && o.AuthScheme != nil {
		return true
	}

	return false
}

// SetAuthScheme gets a reference to the given InlineHookChannelConfigAuthScheme and assigns it to the AuthScheme field.
func (o *InlineHookOAuthBasicConfig) SetAuthScheme(v InlineHookChannelConfigAuthScheme) {
	o.AuthScheme = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *InlineHookOAuthBasicConfig) GetHeaders() []InlineHookChannelConfigHeaders {
	if o == nil || o.Headers == nil {
		var ret []InlineHookChannelConfigHeaders
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookOAuthBasicConfig) GetHeadersOk() ([]InlineHookChannelConfigHeaders, bool) {
	if o == nil || o.Headers == nil {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *InlineHookOAuthBasicConfig) HasHeaders() bool {
	if o != nil && o.Headers != nil {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given []InlineHookChannelConfigHeaders and assigns it to the Headers field.
func (o *InlineHookOAuthBasicConfig) SetHeaders(v []InlineHookChannelConfigHeaders) {
	o.Headers = v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *InlineHookOAuthBasicConfig) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookOAuthBasicConfig) GetMethodOk() (*string, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *InlineHookOAuthBasicConfig) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *InlineHookOAuthBasicConfig) SetMethod(v string) {
	o.Method = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *InlineHookOAuthBasicConfig) GetUri() string {
	if o == nil || o.Uri == nil {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookOAuthBasicConfig) GetUriOk() (*string, bool) {
	if o == nil || o.Uri == nil {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *InlineHookOAuthBasicConfig) HasUri() bool {
	if o != nil && o.Uri != nil {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *InlineHookOAuthBasicConfig) SetUri(v string) {
	o.Uri = &v
}

func (o InlineHookOAuthBasicConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthType != nil {
		toSerialize["authType"] = o.AuthType
	}
	if o.ClientId != nil {
		toSerialize["clientId"] = o.ClientId
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	if o.TokenUrl != nil {
		toSerialize["tokenUrl"] = o.TokenUrl
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

func (o *InlineHookOAuthBasicConfig) UnmarshalJSON(bytes []byte) (err error) {
	varInlineHookOAuthBasicConfig := _InlineHookOAuthBasicConfig{}

	err = json.Unmarshal(bytes, &varInlineHookOAuthBasicConfig)
	if err == nil {
		*o = InlineHookOAuthBasicConfig(varInlineHookOAuthBasicConfig)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "authType")
		delete(additionalProperties, "clientId")
		delete(additionalProperties, "scope")
		delete(additionalProperties, "tokenUrl")
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

type NullableInlineHookOAuthBasicConfig struct {
	value *InlineHookOAuthBasicConfig
	isSet bool
}

func (v NullableInlineHookOAuthBasicConfig) Get() *InlineHookOAuthBasicConfig {
	return v.value
}

func (v *NullableInlineHookOAuthBasicConfig) Set(val *InlineHookOAuthBasicConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineHookOAuthBasicConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineHookOAuthBasicConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineHookOAuthBasicConfig(val *InlineHookOAuthBasicConfig) *NullableInlineHookOAuthBasicConfig {
	return &NullableInlineHookOAuthBasicConfig{value: val, isSet: true}
}

func (v NullableInlineHookOAuthBasicConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineHookOAuthBasicConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

