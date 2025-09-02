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

// TokenPayLoadDataAccess Provides information on the properties of the access token that Okta has generated, including the existing claims that it contains
type TokenPayLoadDataAccess struct {
	// Claims included in the token. Consists of name-value pairs for each included claim. For descriptions of the claims that you can include, see the Okta [OpenID Connect and OAuth 2.0 API reference](/openapi/okta-oauth/guides/overview/#claims).
	Claims map[string]interface{} `json:"claims,omitempty"`
	Token *BaseTokenToken `json:"token,omitempty"`
	// The scopes contained in the token. For descriptions of the scopes that you can include, see the Okta [OpenID Connect and OAuth 2.0 API reference](/openapi/okta-oauth/guides/overview/#scopes).
	Scopes map[string]interface{} `json:"scopes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TokenPayLoadDataAccess TokenPayLoadDataAccess

// NewTokenPayLoadDataAccess instantiates a new TokenPayLoadDataAccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenPayLoadDataAccess() *TokenPayLoadDataAccess {
	this := TokenPayLoadDataAccess{}
	return &this
}

// NewTokenPayLoadDataAccessWithDefaults instantiates a new TokenPayLoadDataAccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenPayLoadDataAccessWithDefaults() *TokenPayLoadDataAccess {
	this := TokenPayLoadDataAccess{}
	return &this
}

// GetClaims returns the Claims field value if set, zero value otherwise.
func (o *TokenPayLoadDataAccess) GetClaims() map[string]interface{} {
	if o == nil || o.Claims == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Claims
}

// GetClaimsOk returns a tuple with the Claims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenPayLoadDataAccess) GetClaimsOk() (map[string]interface{}, bool) {
	if o == nil || o.Claims == nil {
		return nil, false
	}
	return o.Claims, true
}

// HasClaims returns a boolean if a field has been set.
func (o *TokenPayLoadDataAccess) HasClaims() bool {
	if o != nil && o.Claims != nil {
		return true
	}

	return false
}

// SetClaims gets a reference to the given map[string]interface{} and assigns it to the Claims field.
func (o *TokenPayLoadDataAccess) SetClaims(v map[string]interface{}) {
	o.Claims = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *TokenPayLoadDataAccess) GetToken() BaseTokenToken {
	if o == nil || o.Token == nil {
		var ret BaseTokenToken
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenPayLoadDataAccess) GetTokenOk() (*BaseTokenToken, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *TokenPayLoadDataAccess) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given BaseTokenToken and assigns it to the Token field.
func (o *TokenPayLoadDataAccess) SetToken(v BaseTokenToken) {
	o.Token = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *TokenPayLoadDataAccess) GetScopes() map[string]interface{} {
	if o == nil || o.Scopes == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenPayLoadDataAccess) GetScopesOk() (map[string]interface{}, bool) {
	if o == nil || o.Scopes == nil {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *TokenPayLoadDataAccess) HasScopes() bool {
	if o != nil && o.Scopes != nil {
		return true
	}

	return false
}

// SetScopes gets a reference to the given map[string]interface{} and assigns it to the Scopes field.
func (o *TokenPayLoadDataAccess) SetScopes(v map[string]interface{}) {
	o.Scopes = v
}

func (o TokenPayLoadDataAccess) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Claims != nil {
		toSerialize["claims"] = o.Claims
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.Scopes != nil {
		toSerialize["scopes"] = o.Scopes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TokenPayLoadDataAccess) UnmarshalJSON(bytes []byte) (err error) {
	varTokenPayLoadDataAccess := _TokenPayLoadDataAccess{}

	err = json.Unmarshal(bytes, &varTokenPayLoadDataAccess)
	if err == nil {
		*o = TokenPayLoadDataAccess(varTokenPayLoadDataAccess)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "claims")
		delete(additionalProperties, "token")
		delete(additionalProperties, "scopes")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableTokenPayLoadDataAccess struct {
	value *TokenPayLoadDataAccess
	isSet bool
}

func (v NullableTokenPayLoadDataAccess) Get() *TokenPayLoadDataAccess {
	return v.value
}

func (v *NullableTokenPayLoadDataAccess) Set(val *TokenPayLoadDataAccess) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenPayLoadDataAccess) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenPayLoadDataAccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenPayLoadDataAccess(val *TokenPayLoadDataAccess) *NullableTokenPayLoadDataAccess {
	return &NullableTokenPayLoadDataAccess{value: val, isSet: true}
}

func (v NullableTokenPayLoadDataAccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenPayLoadDataAccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

