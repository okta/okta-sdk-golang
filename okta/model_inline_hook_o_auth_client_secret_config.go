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

API version: 2025.08.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the InlineHookOAuthClientSecretConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineHookOAuthClientSecretConfig{}

// InlineHookOAuthClientSecretConfig struct for InlineHookOAuthClientSecretConfig
type InlineHookOAuthClientSecretConfig struct {
	// Not applicable. Must be `null`.
	AuthScheme NullableString `json:"authScheme,omitempty"`
	AuthType   *string        `json:"authType,omitempty"`
	// A publicly exposed string provided by the service that's used to identify the OAuth app and build authorization URLs
	ClientId *string `json:"clientId,omitempty"`
	// Include the scopes that allow you to perform the actions on the hook endpoint that you want to access
	Scope *string `json:"scope,omitempty"`
	// The URI where inline hooks can exchange an authorization code for access and refresh tokens
	TokenUrl *string `json:"tokenUrl,omitempty"`
	// An optional list of key/value pairs for headers that you can send with the request to the external service
	Headers []InlineHookChannelConfigHeaders `json:"headers,omitempty"`
	// The method of the Okta inline hook request
	Method *string `json:"method,omitempty"`
	// The external service endpoint that executes the inline hook handler. It must begin with `https://` and be reachable by Okta. No white space is allowed in the URI.
	Uri                  *string `json:"uri,omitempty"`
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

// GetAuthScheme returns the AuthScheme field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineHookOAuthClientSecretConfig) GetAuthScheme() string {
	if o == nil || IsNil(o.AuthScheme.Get()) {
		var ret string
		return ret
	}
	return *o.AuthScheme.Get()
}

// GetAuthSchemeOk returns a tuple with the AuthScheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineHookOAuthClientSecretConfig) GetAuthSchemeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthScheme.Get(), o.AuthScheme.IsSet()
}

// HasAuthScheme returns a boolean if a field has been set.
func (o *InlineHookOAuthClientSecretConfig) HasAuthScheme() bool {
	if o != nil && o.AuthScheme.IsSet() {
		return true
	}

	return false
}

// SetAuthScheme gets a reference to the given NullableString and assigns it to the AuthScheme field.
func (o *InlineHookOAuthClientSecretConfig) SetAuthScheme(v string) {
	o.AuthScheme.Set(&v)
}

// SetAuthSchemeNil sets the value for AuthScheme to be an explicit nil
func (o *InlineHookOAuthClientSecretConfig) SetAuthSchemeNil() {
	o.AuthScheme.Set(nil)
}

// UnsetAuthScheme ensures that no value is present for AuthScheme, not even an explicit nil
func (o *InlineHookOAuthClientSecretConfig) UnsetAuthScheme() {
	o.AuthScheme.Unset()
}

// GetAuthType returns the AuthType field value if set, zero value otherwise.
func (o *InlineHookOAuthClientSecretConfig) GetAuthType() string {
	if o == nil || IsNil(o.AuthType) {
		var ret string
		return ret
	}
	return *o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookOAuthClientSecretConfig) GetAuthTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AuthType) {
		return nil, false
	}
	return o.AuthType, true
}

// HasAuthType returns a boolean if a field has been set.
func (o *InlineHookOAuthClientSecretConfig) HasAuthType() bool {
	if o != nil && !IsNil(o.AuthType) {
		return true
	}

	return false
}

// SetAuthType gets a reference to the given string and assigns it to the AuthType field.
func (o *InlineHookOAuthClientSecretConfig) SetAuthType(v string) {
	o.AuthType = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *InlineHookOAuthClientSecretConfig) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookOAuthClientSecretConfig) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *InlineHookOAuthClientSecretConfig) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *InlineHookOAuthClientSecretConfig) SetClientId(v string) {
	o.ClientId = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *InlineHookOAuthClientSecretConfig) GetScope() string {
	if o == nil || IsNil(o.Scope) {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookOAuthClientSecretConfig) GetScopeOk() (*string, bool) {
	if o == nil || IsNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *InlineHookOAuthClientSecretConfig) HasScope() bool {
	if o != nil && !IsNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *InlineHookOAuthClientSecretConfig) SetScope(v string) {
	o.Scope = &v
}

// GetTokenUrl returns the TokenUrl field value if set, zero value otherwise.
func (o *InlineHookOAuthClientSecretConfig) GetTokenUrl() string {
	if o == nil || IsNil(o.TokenUrl) {
		var ret string
		return ret
	}
	return *o.TokenUrl
}

// GetTokenUrlOk returns a tuple with the TokenUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookOAuthClientSecretConfig) GetTokenUrlOk() (*string, bool) {
	if o == nil || IsNil(o.TokenUrl) {
		return nil, false
	}
	return o.TokenUrl, true
}

// HasTokenUrl returns a boolean if a field has been set.
func (o *InlineHookOAuthClientSecretConfig) HasTokenUrl() bool {
	if o != nil && !IsNil(o.TokenUrl) {
		return true
	}

	return false
}

// SetTokenUrl gets a reference to the given string and assigns it to the TokenUrl field.
func (o *InlineHookOAuthClientSecretConfig) SetTokenUrl(v string) {
	o.TokenUrl = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *InlineHookOAuthClientSecretConfig) GetHeaders() []InlineHookChannelConfigHeaders {
	if o == nil || IsNil(o.Headers) {
		var ret []InlineHookChannelConfigHeaders
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookOAuthClientSecretConfig) GetHeadersOk() ([]InlineHookChannelConfigHeaders, bool) {
	if o == nil || IsNil(o.Headers) {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *InlineHookOAuthClientSecretConfig) HasHeaders() bool {
	if o != nil && !IsNil(o.Headers) {
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
	if o == nil || IsNil(o.Method) {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookOAuthClientSecretConfig) GetMethodOk() (*string, bool) {
	if o == nil || IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *InlineHookOAuthClientSecretConfig) HasMethod() bool {
	if o != nil && !IsNil(o.Method) {
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
	if o == nil || IsNil(o.Uri) {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookOAuthClientSecretConfig) GetUriOk() (*string, bool) {
	if o == nil || IsNil(o.Uri) {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *InlineHookOAuthClientSecretConfig) HasUri() bool {
	if o != nil && !IsNil(o.Uri) {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *InlineHookOAuthClientSecretConfig) SetUri(v string) {
	o.Uri = &v
}

func (o InlineHookOAuthClientSecretConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineHookOAuthClientSecretConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthScheme.IsSet() {
		toSerialize["authScheme"] = o.AuthScheme.Get()
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

func (o *InlineHookOAuthClientSecretConfig) UnmarshalJSON(data []byte) (err error) {
	varInlineHookOAuthClientSecretConfig := _InlineHookOAuthClientSecretConfig{}

	err = json.Unmarshal(data, &varInlineHookOAuthClientSecretConfig)

	if err != nil {
		return err
	}

	*o = InlineHookOAuthClientSecretConfig(varInlineHookOAuthClientSecretConfig)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "authScheme")
		delete(additionalProperties, "authType")
		delete(additionalProperties, "clientId")
		delete(additionalProperties, "scope")
		delete(additionalProperties, "tokenUrl")
		delete(additionalProperties, "headers")
		delete(additionalProperties, "method")
		delete(additionalProperties, "uri")
		o.AdditionalProperties = additionalProperties
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
