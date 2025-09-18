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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the InlineHookOAuthPrivateKeyJwtConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineHookOAuthPrivateKeyJwtConfig{}

// InlineHookOAuthPrivateKeyJwtConfig struct for InlineHookOAuthPrivateKeyJwtConfig
type InlineHookOAuthPrivateKeyJwtConfig struct {
	// Not applicable. Must be `null`.
	AuthScheme NullableString `json:"authScheme,omitempty"`
	// An ID value of the hook key pair generated from the [Hook Keys API](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/HookKey/#tag/HookKey)
	HookKeyId *string `json:"hookKeyId,omitempty"`
	// The method of the Okta inline hook request. Only accepts `POST`.
	Method   *string `json:"method,omitempty"`
	AuthType *string `json:"authType,omitempty"`
	// A publicly exposed string provided by the service that's used to identify the OAuth app and build authorization URLs
	ClientId *string `json:"clientId,omitempty"`
	// Include the scopes that allow you to perform the actions on the hook endpoint that you want to access
	Scope *string `json:"scope,omitempty"`
	// The URI where inline hooks can exchange an authorization code for access and refresh tokens
	TokenUrl *string `json:"tokenUrl,omitempty"`
	// An optional list of key/value pairs for headers that you can send with the request to the external service
	Headers []InlineHookChannelConfigHeaders `json:"headers,omitempty"`
	// The external service endpoint that executes the inline hook handler. It must begin with `https://` and be reachable by Okta. No white space is allowed in the URI.
	Uri                  *string `json:"uri,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InlineHookOAuthPrivateKeyJwtConfig InlineHookOAuthPrivateKeyJwtConfig

// NewInlineHookOAuthPrivateKeyJwtConfig instantiates a new InlineHookOAuthPrivateKeyJwtConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineHookOAuthPrivateKeyJwtConfig() *InlineHookOAuthPrivateKeyJwtConfig {
	this := InlineHookOAuthPrivateKeyJwtConfig{}
	return &this
}

// NewInlineHookOAuthPrivateKeyJwtConfigWithDefaults instantiates a new InlineHookOAuthPrivateKeyJwtConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineHookOAuthPrivateKeyJwtConfigWithDefaults() *InlineHookOAuthPrivateKeyJwtConfig {
	this := InlineHookOAuthPrivateKeyJwtConfig{}
	return &this
}

// GetAuthScheme returns the AuthScheme field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineHookOAuthPrivateKeyJwtConfig) GetAuthScheme() string {
	if o == nil || IsNil(o.AuthScheme.Get()) {
		var ret string
		return ret
	}
	return *o.AuthScheme.Get()
}

// GetAuthSchemeOk returns a tuple with the AuthScheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineHookOAuthPrivateKeyJwtConfig) GetAuthSchemeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthScheme.Get(), o.AuthScheme.IsSet()
}

// HasAuthScheme returns a boolean if a field has been set.
func (o *InlineHookOAuthPrivateKeyJwtConfig) HasAuthScheme() bool {
	if o != nil && o.AuthScheme.IsSet() {
		return true
	}

	return false
}

// SetAuthScheme gets a reference to the given NullableString and assigns it to the AuthScheme field.
func (o *InlineHookOAuthPrivateKeyJwtConfig) SetAuthScheme(v string) {
	o.AuthScheme.Set(&v)
}

// SetAuthSchemeNil sets the value for AuthScheme to be an explicit nil
func (o *InlineHookOAuthPrivateKeyJwtConfig) SetAuthSchemeNil() {
	o.AuthScheme.Set(nil)
}

// UnsetAuthScheme ensures that no value is present for AuthScheme, not even an explicit nil
func (o *InlineHookOAuthPrivateKeyJwtConfig) UnsetAuthScheme() {
	o.AuthScheme.Unset()
}

// GetHookKeyId returns the HookKeyId field value if set, zero value otherwise.
func (o *InlineHookOAuthPrivateKeyJwtConfig) GetHookKeyId() string {
	if o == nil || IsNil(o.HookKeyId) {
		var ret string
		return ret
	}
	return *o.HookKeyId
}

// GetHookKeyIdOk returns a tuple with the HookKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookOAuthPrivateKeyJwtConfig) GetHookKeyIdOk() (*string, bool) {
	if o == nil || IsNil(o.HookKeyId) {
		return nil, false
	}
	return o.HookKeyId, true
}

// HasHookKeyId returns a boolean if a field has been set.
func (o *InlineHookOAuthPrivateKeyJwtConfig) HasHookKeyId() bool {
	if o != nil && !IsNil(o.HookKeyId) {
		return true
	}

	return false
}

// SetHookKeyId gets a reference to the given string and assigns it to the HookKeyId field.
func (o *InlineHookOAuthPrivateKeyJwtConfig) SetHookKeyId(v string) {
	o.HookKeyId = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *InlineHookOAuthPrivateKeyJwtConfig) GetMethod() string {
	if o == nil || IsNil(o.Method) {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookOAuthPrivateKeyJwtConfig) GetMethodOk() (*string, bool) {
	if o == nil || IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *InlineHookOAuthPrivateKeyJwtConfig) HasMethod() bool {
	if o != nil && !IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *InlineHookOAuthPrivateKeyJwtConfig) SetMethod(v string) {
	o.Method = &v
}

// GetAuthType returns the AuthType field value if set, zero value otherwise.
func (o *InlineHookOAuthPrivateKeyJwtConfig) GetAuthType() string {
	if o == nil || IsNil(o.AuthType) {
		var ret string
		return ret
	}
	return *o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookOAuthPrivateKeyJwtConfig) GetAuthTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AuthType) {
		return nil, false
	}
	return o.AuthType, true
}

// HasAuthType returns a boolean if a field has been set.
func (o *InlineHookOAuthPrivateKeyJwtConfig) HasAuthType() bool {
	if o != nil && !IsNil(o.AuthType) {
		return true
	}

	return false
}

// SetAuthType gets a reference to the given string and assigns it to the AuthType field.
func (o *InlineHookOAuthPrivateKeyJwtConfig) SetAuthType(v string) {
	o.AuthType = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *InlineHookOAuthPrivateKeyJwtConfig) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookOAuthPrivateKeyJwtConfig) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *InlineHookOAuthPrivateKeyJwtConfig) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *InlineHookOAuthPrivateKeyJwtConfig) SetClientId(v string) {
	o.ClientId = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *InlineHookOAuthPrivateKeyJwtConfig) GetScope() string {
	if o == nil || IsNil(o.Scope) {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookOAuthPrivateKeyJwtConfig) GetScopeOk() (*string, bool) {
	if o == nil || IsNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *InlineHookOAuthPrivateKeyJwtConfig) HasScope() bool {
	if o != nil && !IsNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *InlineHookOAuthPrivateKeyJwtConfig) SetScope(v string) {
	o.Scope = &v
}

// GetTokenUrl returns the TokenUrl field value if set, zero value otherwise.
func (o *InlineHookOAuthPrivateKeyJwtConfig) GetTokenUrl() string {
	if o == nil || IsNil(o.TokenUrl) {
		var ret string
		return ret
	}
	return *o.TokenUrl
}

// GetTokenUrlOk returns a tuple with the TokenUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookOAuthPrivateKeyJwtConfig) GetTokenUrlOk() (*string, bool) {
	if o == nil || IsNil(o.TokenUrl) {
		return nil, false
	}
	return o.TokenUrl, true
}

// HasTokenUrl returns a boolean if a field has been set.
func (o *InlineHookOAuthPrivateKeyJwtConfig) HasTokenUrl() bool {
	if o != nil && !IsNil(o.TokenUrl) {
		return true
	}

	return false
}

// SetTokenUrl gets a reference to the given string and assigns it to the TokenUrl field.
func (o *InlineHookOAuthPrivateKeyJwtConfig) SetTokenUrl(v string) {
	o.TokenUrl = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *InlineHookOAuthPrivateKeyJwtConfig) GetHeaders() []InlineHookChannelConfigHeaders {
	if o == nil || IsNil(o.Headers) {
		var ret []InlineHookChannelConfigHeaders
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookOAuthPrivateKeyJwtConfig) GetHeadersOk() ([]InlineHookChannelConfigHeaders, bool) {
	if o == nil || IsNil(o.Headers) {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *InlineHookOAuthPrivateKeyJwtConfig) HasHeaders() bool {
	if o != nil && !IsNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given []InlineHookChannelConfigHeaders and assigns it to the Headers field.
func (o *InlineHookOAuthPrivateKeyJwtConfig) SetHeaders(v []InlineHookChannelConfigHeaders) {
	o.Headers = v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *InlineHookOAuthPrivateKeyJwtConfig) GetUri() string {
	if o == nil || IsNil(o.Uri) {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookOAuthPrivateKeyJwtConfig) GetUriOk() (*string, bool) {
	if o == nil || IsNil(o.Uri) {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *InlineHookOAuthPrivateKeyJwtConfig) HasUri() bool {
	if o != nil && !IsNil(o.Uri) {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *InlineHookOAuthPrivateKeyJwtConfig) SetUri(v string) {
	o.Uri = &v
}

func (o InlineHookOAuthPrivateKeyJwtConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineHookOAuthPrivateKeyJwtConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthScheme.IsSet() {
		toSerialize["authScheme"] = o.AuthScheme.Get()
	}
	if !IsNil(o.HookKeyId) {
		toSerialize["hookKeyId"] = o.HookKeyId
	}
	if !IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}
	if !IsNil(o.AuthType) {
		toSerialize["authType"] = o.AuthType
	}
	if !IsNil(o.ClientId) {
		toSerialize["clientId"] = o.ClientId
	}
	if !IsNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}
	if !IsNil(o.TokenUrl) {
		toSerialize["tokenUrl"] = o.TokenUrl
	}
	if !IsNil(o.Headers) {
		toSerialize["headers"] = o.Headers
	}
	if !IsNil(o.Uri) {
		toSerialize["uri"] = o.Uri
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InlineHookOAuthPrivateKeyJwtConfig) UnmarshalJSON(data []byte) (err error) {
	varInlineHookOAuthPrivateKeyJwtConfig := _InlineHookOAuthPrivateKeyJwtConfig{}

	err = json.Unmarshal(data, &varInlineHookOAuthPrivateKeyJwtConfig)

	if err != nil {
		return err
	}

	*o = InlineHookOAuthPrivateKeyJwtConfig(varInlineHookOAuthPrivateKeyJwtConfig)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "authScheme")
		delete(additionalProperties, "hookKeyId")
		delete(additionalProperties, "method")
		delete(additionalProperties, "authType")
		delete(additionalProperties, "clientId")
		delete(additionalProperties, "scope")
		delete(additionalProperties, "tokenUrl")
		delete(additionalProperties, "headers")
		delete(additionalProperties, "uri")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInlineHookOAuthPrivateKeyJwtConfig struct {
	value *InlineHookOAuthPrivateKeyJwtConfig
	isSet bool
}

func (v NullableInlineHookOAuthPrivateKeyJwtConfig) Get() *InlineHookOAuthPrivateKeyJwtConfig {
	return v.value
}

func (v *NullableInlineHookOAuthPrivateKeyJwtConfig) Set(val *InlineHookOAuthPrivateKeyJwtConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineHookOAuthPrivateKeyJwtConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineHookOAuthPrivateKeyJwtConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineHookOAuthPrivateKeyJwtConfig(val *InlineHookOAuthPrivateKeyJwtConfig) *NullableInlineHookOAuthPrivateKeyJwtConfig {
	return &NullableInlineHookOAuthPrivateKeyJwtConfig{value: val, isSet: true}
}

func (v NullableInlineHookOAuthPrivateKeyJwtConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineHookOAuthPrivateKeyJwtConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
