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

// TokenResponse struct for TokenResponse
type TokenResponse struct {
	// An access token.
	AccessToken *string `json:"access_token,omitempty"`
	// An opaque device secret. This is returned if the `device_sso` scope is granted.
	DeviceSecret *string `json:"device_secret,omitempty"`
	// The expiration time of the access token in seconds.
	ExpiresIn *int32 `json:"expires_in,omitempty"`
	// An ID token. This is returned if the `openid` scope is granted.
	IdToken *string `json:"id_token,omitempty"`
	// The type of token for token exchange.
	IssuedTokenType *string `json:"issued_token_type,omitempty"`
	// An opaque refresh token. This is returned if the `offline_access` scope is granted.
	RefreshToken *string `json:"refresh_token,omitempty"`
	// The scopes contained in the access token.
	Scope *string `json:"scope,omitempty"`
	// The token type in a `/token` response. The value is generally `Bearer` except for a few instances of token exchange.
	TokenType *string `json:"token_type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TokenResponse TokenResponse

// NewTokenResponse instantiates a new TokenResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenResponse() *TokenResponse {
	this := TokenResponse{}
	return &this
}

// NewTokenResponseWithDefaults instantiates a new TokenResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenResponseWithDefaults() *TokenResponse {
	this := TokenResponse{}
	return &this
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise.
func (o *TokenResponse) GetAccessToken() string {
	if o == nil || o.AccessToken == nil {
		var ret string
		return ret
	}
	return *o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetAccessTokenOk() (*string, bool) {
	if o == nil || o.AccessToken == nil {
		return nil, false
	}
	return o.AccessToken, true
}

// HasAccessToken returns a boolean if a field has been set.
func (o *TokenResponse) HasAccessToken() bool {
	if o != nil && o.AccessToken != nil {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given string and assigns it to the AccessToken field.
func (o *TokenResponse) SetAccessToken(v string) {
	o.AccessToken = &v
}

// GetDeviceSecret returns the DeviceSecret field value if set, zero value otherwise.
func (o *TokenResponse) GetDeviceSecret() string {
	if o == nil || o.DeviceSecret == nil {
		var ret string
		return ret
	}
	return *o.DeviceSecret
}

// GetDeviceSecretOk returns a tuple with the DeviceSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetDeviceSecretOk() (*string, bool) {
	if o == nil || o.DeviceSecret == nil {
		return nil, false
	}
	return o.DeviceSecret, true
}

// HasDeviceSecret returns a boolean if a field has been set.
func (o *TokenResponse) HasDeviceSecret() bool {
	if o != nil && o.DeviceSecret != nil {
		return true
	}

	return false
}

// SetDeviceSecret gets a reference to the given string and assigns it to the DeviceSecret field.
func (o *TokenResponse) SetDeviceSecret(v string) {
	o.DeviceSecret = &v
}

// GetExpiresIn returns the ExpiresIn field value if set, zero value otherwise.
func (o *TokenResponse) GetExpiresIn() int32 {
	if o == nil || o.ExpiresIn == nil {
		var ret int32
		return ret
	}
	return *o.ExpiresIn
}

// GetExpiresInOk returns a tuple with the ExpiresIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetExpiresInOk() (*int32, bool) {
	if o == nil || o.ExpiresIn == nil {
		return nil, false
	}
	return o.ExpiresIn, true
}

// HasExpiresIn returns a boolean if a field has been set.
func (o *TokenResponse) HasExpiresIn() bool {
	if o != nil && o.ExpiresIn != nil {
		return true
	}

	return false
}

// SetExpiresIn gets a reference to the given int32 and assigns it to the ExpiresIn field.
func (o *TokenResponse) SetExpiresIn(v int32) {
	o.ExpiresIn = &v
}

// GetIdToken returns the IdToken field value if set, zero value otherwise.
func (o *TokenResponse) GetIdToken() string {
	if o == nil || o.IdToken == nil {
		var ret string
		return ret
	}
	return *o.IdToken
}

// GetIdTokenOk returns a tuple with the IdToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetIdTokenOk() (*string, bool) {
	if o == nil || o.IdToken == nil {
		return nil, false
	}
	return o.IdToken, true
}

// HasIdToken returns a boolean if a field has been set.
func (o *TokenResponse) HasIdToken() bool {
	if o != nil && o.IdToken != nil {
		return true
	}

	return false
}

// SetIdToken gets a reference to the given string and assigns it to the IdToken field.
func (o *TokenResponse) SetIdToken(v string) {
	o.IdToken = &v
}

// GetIssuedTokenType returns the IssuedTokenType field value if set, zero value otherwise.
func (o *TokenResponse) GetIssuedTokenType() string {
	if o == nil || o.IssuedTokenType == nil {
		var ret string
		return ret
	}
	return *o.IssuedTokenType
}

// GetIssuedTokenTypeOk returns a tuple with the IssuedTokenType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetIssuedTokenTypeOk() (*string, bool) {
	if o == nil || o.IssuedTokenType == nil {
		return nil, false
	}
	return o.IssuedTokenType, true
}

// HasIssuedTokenType returns a boolean if a field has been set.
func (o *TokenResponse) HasIssuedTokenType() bool {
	if o != nil && o.IssuedTokenType != nil {
		return true
	}

	return false
}

// SetIssuedTokenType gets a reference to the given string and assigns it to the IssuedTokenType field.
func (o *TokenResponse) SetIssuedTokenType(v string) {
	o.IssuedTokenType = &v
}

// GetRefreshToken returns the RefreshToken field value if set, zero value otherwise.
func (o *TokenResponse) GetRefreshToken() string {
	if o == nil || o.RefreshToken == nil {
		var ret string
		return ret
	}
	return *o.RefreshToken
}

// GetRefreshTokenOk returns a tuple with the RefreshToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetRefreshTokenOk() (*string, bool) {
	if o == nil || o.RefreshToken == nil {
		return nil, false
	}
	return o.RefreshToken, true
}

// HasRefreshToken returns a boolean if a field has been set.
func (o *TokenResponse) HasRefreshToken() bool {
	if o != nil && o.RefreshToken != nil {
		return true
	}

	return false
}

// SetRefreshToken gets a reference to the given string and assigns it to the RefreshToken field.
func (o *TokenResponse) SetRefreshToken(v string) {
	o.RefreshToken = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *TokenResponse) GetScope() string {
	if o == nil || o.Scope == nil {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetScopeOk() (*string, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *TokenResponse) HasScope() bool {
	if o != nil && o.Scope != nil {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *TokenResponse) SetScope(v string) {
	o.Scope = &v
}

// GetTokenType returns the TokenType field value if set, zero value otherwise.
func (o *TokenResponse) GetTokenType() string {
	if o == nil || o.TokenType == nil {
		var ret string
		return ret
	}
	return *o.TokenType
}

// GetTokenTypeOk returns a tuple with the TokenType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetTokenTypeOk() (*string, bool) {
	if o == nil || o.TokenType == nil {
		return nil, false
	}
	return o.TokenType, true
}

// HasTokenType returns a boolean if a field has been set.
func (o *TokenResponse) HasTokenType() bool {
	if o != nil && o.TokenType != nil {
		return true
	}

	return false
}

// SetTokenType gets a reference to the given string and assigns it to the TokenType field.
func (o *TokenResponse) SetTokenType(v string) {
	o.TokenType = &v
}

func (o TokenResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessToken != nil {
		toSerialize["access_token"] = o.AccessToken
	}
	if o.DeviceSecret != nil {
		toSerialize["device_secret"] = o.DeviceSecret
	}
	if o.ExpiresIn != nil {
		toSerialize["expires_in"] = o.ExpiresIn
	}
	if o.IdToken != nil {
		toSerialize["id_token"] = o.IdToken
	}
	if o.IssuedTokenType != nil {
		toSerialize["issued_token_type"] = o.IssuedTokenType
	}
	if o.RefreshToken != nil {
		toSerialize["refresh_token"] = o.RefreshToken
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	if o.TokenType != nil {
		toSerialize["token_type"] = o.TokenType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TokenResponse) UnmarshalJSON(bytes []byte) (err error) {
	varTokenResponse := _TokenResponse{}

	err = json.Unmarshal(bytes, &varTokenResponse)
	if err == nil {
		*o = TokenResponse(varTokenResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "access_token")
		delete(additionalProperties, "device_secret")
		delete(additionalProperties, "expires_in")
		delete(additionalProperties, "id_token")
		delete(additionalProperties, "issued_token_type")
		delete(additionalProperties, "refresh_token")
		delete(additionalProperties, "scope")
		delete(additionalProperties, "token_type")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableTokenResponse struct {
	value *TokenResponse
	isSet bool
}

func (v NullableTokenResponse) Get() *TokenResponse {
	return v.value
}

func (v *NullableTokenResponse) Set(val *TokenResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenResponse(val *TokenResponse) *NullableTokenResponse {
	return &NullableTokenResponse{value: val, isSet: true}
}

func (v NullableTokenResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

