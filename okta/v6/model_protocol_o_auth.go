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

// ProtocolOAuth Protocol settings for authentication using the [OAuth 2.0 Authorization Code flow](https://tools.ietf.org/html/rfc6749#section-4.1)
type ProtocolOAuth struct {
	Credentials *OAuthCredentials `json:"credentials,omitempty"`
	Endpoints *OAuthEndpoints `json:"endpoints,omitempty"`
	// IdP-defined permission bundles to request delegated access from the user. > **Note:** The [identity provider type](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/IdentityProvider/#tag/IdentityProvider/operation/createIdentityProvider!path=type&t=request) table lists the scopes that are supported for each IdP.
	Scopes []string `json:"scopes,omitempty"`
	// OAuth 2.0 Authorization Code flow
	Type *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProtocolOAuth ProtocolOAuth

// NewProtocolOAuth instantiates a new ProtocolOAuth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProtocolOAuth() *ProtocolOAuth {
	this := ProtocolOAuth{}
	return &this
}

// NewProtocolOAuthWithDefaults instantiates a new ProtocolOAuth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProtocolOAuthWithDefaults() *ProtocolOAuth {
	this := ProtocolOAuth{}
	return &this
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *ProtocolOAuth) GetCredentials() OAuthCredentials {
	if o == nil || o.Credentials == nil {
		var ret OAuthCredentials
		return ret
	}
	return *o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtocolOAuth) GetCredentialsOk() (*OAuthCredentials, bool) {
	if o == nil || o.Credentials == nil {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *ProtocolOAuth) HasCredentials() bool {
	if o != nil && o.Credentials != nil {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given OAuthCredentials and assigns it to the Credentials field.
func (o *ProtocolOAuth) SetCredentials(v OAuthCredentials) {
	o.Credentials = &v
}

// GetEndpoints returns the Endpoints field value if set, zero value otherwise.
func (o *ProtocolOAuth) GetEndpoints() OAuthEndpoints {
	if o == nil || o.Endpoints == nil {
		var ret OAuthEndpoints
		return ret
	}
	return *o.Endpoints
}

// GetEndpointsOk returns a tuple with the Endpoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtocolOAuth) GetEndpointsOk() (*OAuthEndpoints, bool) {
	if o == nil || o.Endpoints == nil {
		return nil, false
	}
	return o.Endpoints, true
}

// HasEndpoints returns a boolean if a field has been set.
func (o *ProtocolOAuth) HasEndpoints() bool {
	if o != nil && o.Endpoints != nil {
		return true
	}

	return false
}

// SetEndpoints gets a reference to the given OAuthEndpoints and assigns it to the Endpoints field.
func (o *ProtocolOAuth) SetEndpoints(v OAuthEndpoints) {
	o.Endpoints = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *ProtocolOAuth) GetScopes() []string {
	if o == nil || o.Scopes == nil {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtocolOAuth) GetScopesOk() ([]string, bool) {
	if o == nil || o.Scopes == nil {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *ProtocolOAuth) HasScopes() bool {
	if o != nil && o.Scopes != nil {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *ProtocolOAuth) SetScopes(v []string) {
	o.Scopes = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ProtocolOAuth) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtocolOAuth) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ProtocolOAuth) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ProtocolOAuth) SetType(v string) {
	o.Type = &v
}

func (o ProtocolOAuth) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Credentials != nil {
		toSerialize["credentials"] = o.Credentials
	}
	if o.Endpoints != nil {
		toSerialize["endpoints"] = o.Endpoints
	}
	if o.Scopes != nil {
		toSerialize["scopes"] = o.Scopes
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ProtocolOAuth) UnmarshalJSON(bytes []byte) (err error) {
	varProtocolOAuth := _ProtocolOAuth{}

	err = json.Unmarshal(bytes, &varProtocolOAuth)
	if err == nil {
		*o = ProtocolOAuth(varProtocolOAuth)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "credentials")
		delete(additionalProperties, "endpoints")
		delete(additionalProperties, "scopes")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableProtocolOAuth struct {
	value *ProtocolOAuth
	isSet bool
}

func (v NullableProtocolOAuth) Get() *ProtocolOAuth {
	return v.value
}

func (v *NullableProtocolOAuth) Set(val *ProtocolOAuth) {
	v.value = val
	v.isSet = true
}

func (v NullableProtocolOAuth) IsSet() bool {
	return v.isSet
}

func (v *NullableProtocolOAuth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProtocolOAuth(val *ProtocolOAuth) *NullableProtocolOAuth {
	return &NullableProtocolOAuth{value: val, isSet: true}
}

func (v NullableProtocolOAuth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProtocolOAuth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

