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
	"fmt"
)

// checks if the OAuth2Settings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OAuth2Settings{}

// OAuth2Settings OAuth 2.0 configuration used for authType `OAUTH2`
type OAuth2Settings struct {
	// The URL to the authorization server's authorization endpoint
	AuthorizeEndpoint string `json:"authorizeEndpoint"`
	// The OAuth 2.0 client identifier
	ClientId string `json:"clientId"`
	// The OAuth 2.0 client secret
	ClientSecret string `json:"clientSecret"`
	// List of OAuth 2.0 scopes
	Scopes []string `json:"scopes,omitempty"`
	// The URL to the authorization server's token endpoint
	TokenEndpoint        string `json:"tokenEndpoint"`
	AdditionalProperties map[string]interface{}
}

type _OAuth2Settings OAuth2Settings

// NewOAuth2Settings instantiates a new OAuth2Settings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth2Settings(authorizeEndpoint string, clientId string, clientSecret string, tokenEndpoint string) *OAuth2Settings {
	this := OAuth2Settings{}
	this.AuthorizeEndpoint = authorizeEndpoint
	this.ClientId = clientId
	this.ClientSecret = clientSecret
	this.TokenEndpoint = tokenEndpoint
	return &this
}

// NewOAuth2SettingsWithDefaults instantiates a new OAuth2Settings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth2SettingsWithDefaults() *OAuth2Settings {
	this := OAuth2Settings{}
	return &this
}

// GetAuthorizeEndpoint returns the AuthorizeEndpoint field value
func (o *OAuth2Settings) GetAuthorizeEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthorizeEndpoint
}

// GetAuthorizeEndpointOk returns a tuple with the AuthorizeEndpoint field value
// and a boolean to check if the value has been set.
func (o *OAuth2Settings) GetAuthorizeEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorizeEndpoint, true
}

// SetAuthorizeEndpoint sets field value
func (o *OAuth2Settings) SetAuthorizeEndpoint(v string) {
	o.AuthorizeEndpoint = v
}

// GetClientId returns the ClientId field value
func (o *OAuth2Settings) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *OAuth2Settings) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *OAuth2Settings) SetClientId(v string) {
	o.ClientId = v
}

// GetClientSecret returns the ClientSecret field value
func (o *OAuth2Settings) GetClientSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value
// and a boolean to check if the value has been set.
func (o *OAuth2Settings) GetClientSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientSecret, true
}

// SetClientSecret sets field value
func (o *OAuth2Settings) SetClientSecret(v string) {
	o.ClientSecret = v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *OAuth2Settings) GetScopes() []string {
	if o == nil || IsNil(o.Scopes) {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2Settings) GetScopesOk() ([]string, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *OAuth2Settings) HasScopes() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *OAuth2Settings) SetScopes(v []string) {
	o.Scopes = v
}

// GetTokenEndpoint returns the TokenEndpoint field value
func (o *OAuth2Settings) GetTokenEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenEndpoint
}

// GetTokenEndpointOk returns a tuple with the TokenEndpoint field value
// and a boolean to check if the value has been set.
func (o *OAuth2Settings) GetTokenEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenEndpoint, true
}

// SetTokenEndpoint sets field value
func (o *OAuth2Settings) SetTokenEndpoint(v string) {
	o.TokenEndpoint = v
}

func (o OAuth2Settings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OAuth2Settings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["authorizeEndpoint"] = o.AuthorizeEndpoint
	toSerialize["clientId"] = o.ClientId
	toSerialize["clientSecret"] = o.ClientSecret
	if !IsNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}
	toSerialize["tokenEndpoint"] = o.TokenEndpoint

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OAuth2Settings) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"authorizeEndpoint",
		"clientId",
		"clientSecret",
		"tokenEndpoint",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varOAuth2Settings := _OAuth2Settings{}

	err = json.Unmarshal(data, &varOAuth2Settings)

	if err != nil {
		return err
	}

	*o = OAuth2Settings(varOAuth2Settings)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "authorizeEndpoint")
		delete(additionalProperties, "clientId")
		delete(additionalProperties, "clientSecret")
		delete(additionalProperties, "scopes")
		delete(additionalProperties, "tokenEndpoint")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOAuth2Settings struct {
	value *OAuth2Settings
	isSet bool
}

func (v NullableOAuth2Settings) Get() *OAuth2Settings {
	return v.value
}

func (v *NullableOAuth2Settings) Set(val *OAuth2Settings) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2Settings) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2Settings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2Settings(val *OAuth2Settings) *NullableOAuth2Settings {
	return &NullableOAuth2Settings{value: val, isSet: true}
}

func (v NullableOAuth2Settings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2Settings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
