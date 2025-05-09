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
	"time"
)

// SocialAuthToken struct for SocialAuthToken
type SocialAuthToken struct {
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
	Id *string `json:"id,omitempty"`
	Scopes []string `json:"scopes,omitempty"`
	Token *string `json:"token,omitempty"`
	TokenAuthScheme *string `json:"tokenAuthScheme,omitempty"`
	TokenType *string `json:"tokenType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SocialAuthToken SocialAuthToken

// NewSocialAuthToken instantiates a new SocialAuthToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSocialAuthToken() *SocialAuthToken {
	this := SocialAuthToken{}
	return &this
}

// NewSocialAuthTokenWithDefaults instantiates a new SocialAuthToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSocialAuthTokenWithDefaults() *SocialAuthToken {
	this := SocialAuthToken{}
	return &this
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *SocialAuthToken) GetExpiresAt() time.Time {
	if o == nil || o.ExpiresAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SocialAuthToken) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || o.ExpiresAt == nil {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *SocialAuthToken) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt != nil {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *SocialAuthToken) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SocialAuthToken) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SocialAuthToken) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SocialAuthToken) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SocialAuthToken) SetId(v string) {
	o.Id = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *SocialAuthToken) GetScopes() []string {
	if o == nil || o.Scopes == nil {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SocialAuthToken) GetScopesOk() ([]string, bool) {
	if o == nil || o.Scopes == nil {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *SocialAuthToken) HasScopes() bool {
	if o != nil && o.Scopes != nil {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *SocialAuthToken) SetScopes(v []string) {
	o.Scopes = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *SocialAuthToken) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SocialAuthToken) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *SocialAuthToken) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *SocialAuthToken) SetToken(v string) {
	o.Token = &v
}

// GetTokenAuthScheme returns the TokenAuthScheme field value if set, zero value otherwise.
func (o *SocialAuthToken) GetTokenAuthScheme() string {
	if o == nil || o.TokenAuthScheme == nil {
		var ret string
		return ret
	}
	return *o.TokenAuthScheme
}

// GetTokenAuthSchemeOk returns a tuple with the TokenAuthScheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SocialAuthToken) GetTokenAuthSchemeOk() (*string, bool) {
	if o == nil || o.TokenAuthScheme == nil {
		return nil, false
	}
	return o.TokenAuthScheme, true
}

// HasTokenAuthScheme returns a boolean if a field has been set.
func (o *SocialAuthToken) HasTokenAuthScheme() bool {
	if o != nil && o.TokenAuthScheme != nil {
		return true
	}

	return false
}

// SetTokenAuthScheme gets a reference to the given string and assigns it to the TokenAuthScheme field.
func (o *SocialAuthToken) SetTokenAuthScheme(v string) {
	o.TokenAuthScheme = &v
}

// GetTokenType returns the TokenType field value if set, zero value otherwise.
func (o *SocialAuthToken) GetTokenType() string {
	if o == nil || o.TokenType == nil {
		var ret string
		return ret
	}
	return *o.TokenType
}

// GetTokenTypeOk returns a tuple with the TokenType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SocialAuthToken) GetTokenTypeOk() (*string, bool) {
	if o == nil || o.TokenType == nil {
		return nil, false
	}
	return o.TokenType, true
}

// HasTokenType returns a boolean if a field has been set.
func (o *SocialAuthToken) HasTokenType() bool {
	if o != nil && o.TokenType != nil {
		return true
	}

	return false
}

// SetTokenType gets a reference to the given string and assigns it to the TokenType field.
func (o *SocialAuthToken) SetTokenType(v string) {
	o.TokenType = &v
}

func (o SocialAuthToken) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExpiresAt != nil {
		toSerialize["expiresAt"] = o.ExpiresAt
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Scopes != nil {
		toSerialize["scopes"] = o.Scopes
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.TokenAuthScheme != nil {
		toSerialize["tokenAuthScheme"] = o.TokenAuthScheme
	}
	if o.TokenType != nil {
		toSerialize["tokenType"] = o.TokenType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SocialAuthToken) UnmarshalJSON(bytes []byte) (err error) {
	varSocialAuthToken := _SocialAuthToken{}

	err = json.Unmarshal(bytes, &varSocialAuthToken)
	if err == nil {
		*o = SocialAuthToken(varSocialAuthToken)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "expiresAt")
		delete(additionalProperties, "id")
		delete(additionalProperties, "scopes")
		delete(additionalProperties, "token")
		delete(additionalProperties, "tokenAuthScheme")
		delete(additionalProperties, "tokenType")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableSocialAuthToken struct {
	value *SocialAuthToken
	isSet bool
}

func (v NullableSocialAuthToken) Get() *SocialAuthToken {
	return v.value
}

func (v *NullableSocialAuthToken) Set(val *SocialAuthToken) {
	v.value = val
	v.isSet = true
}

func (v NullableSocialAuthToken) IsSet() bool {
	return v.isSet
}

func (v *NullableSocialAuthToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSocialAuthToken(val *SocialAuthToken) *NullableSocialAuthToken {
	return &NullableSocialAuthToken{value: val, isSet: true}
}

func (v NullableSocialAuthToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSocialAuthToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

