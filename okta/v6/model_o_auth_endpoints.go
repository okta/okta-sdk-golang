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

// checks if the OAuthEndpoints type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OAuthEndpoints{}

// OAuthEndpoints The `OAUTH2` and `OIDC` protocols support the `authorization` and `token` endpoints. Also, the `OIDC` protocol supports the `userInfo` and `jwks` endpoints.  The IdP Authorization Server (AS) endpoints are currently defined as part of the [IdP provider]((https://developer.okta.com/docs/api/openapi/okta-management/management/tag/IdentityProvider/#tag/IdentityProvider/operation/createIdentityProvider!path=type&t=request)) and are read-only.
type OAuthEndpoints struct {
	Authorization        *OAuthAuthorizationEndpoint `json:"authorization,omitempty"`
	Jwks                 *OidcJwksEndpoint           `json:"jwks,omitempty"`
	Slo                  *OidcSloEndpoint            `json:"slo,omitempty"`
	Token                *OAuthTokenEndpoint         `json:"token,omitempty"`
	UserInfo             *OidcUserInfoEndpoint       `json:"userInfo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OAuthEndpoints OAuthEndpoints

// NewOAuthEndpoints instantiates a new OAuthEndpoints object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuthEndpoints() *OAuthEndpoints {
	this := OAuthEndpoints{}
	return &this
}

// NewOAuthEndpointsWithDefaults instantiates a new OAuthEndpoints object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuthEndpointsWithDefaults() *OAuthEndpoints {
	this := OAuthEndpoints{}
	return &this
}

// GetAuthorization returns the Authorization field value if set, zero value otherwise.
func (o *OAuthEndpoints) GetAuthorization() OAuthAuthorizationEndpoint {
	if o == nil || IsNil(o.Authorization) {
		var ret OAuthAuthorizationEndpoint
		return ret
	}
	return *o.Authorization
}

// GetAuthorizationOk returns a tuple with the Authorization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthEndpoints) GetAuthorizationOk() (*OAuthAuthorizationEndpoint, bool) {
	if o == nil || IsNil(o.Authorization) {
		return nil, false
	}
	return o.Authorization, true
}

// HasAuthorization returns a boolean if a field has been set.
func (o *OAuthEndpoints) HasAuthorization() bool {
	if o != nil && !IsNil(o.Authorization) {
		return true
	}

	return false
}

// SetAuthorization gets a reference to the given OAuthAuthorizationEndpoint and assigns it to the Authorization field.
func (o *OAuthEndpoints) SetAuthorization(v OAuthAuthorizationEndpoint) {
	o.Authorization = &v
}

// GetJwks returns the Jwks field value if set, zero value otherwise.
func (o *OAuthEndpoints) GetJwks() OidcJwksEndpoint {
	if o == nil || IsNil(o.Jwks) {
		var ret OidcJwksEndpoint
		return ret
	}
	return *o.Jwks
}

// GetJwksOk returns a tuple with the Jwks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthEndpoints) GetJwksOk() (*OidcJwksEndpoint, bool) {
	if o == nil || IsNil(o.Jwks) {
		return nil, false
	}
	return o.Jwks, true
}

// HasJwks returns a boolean if a field has been set.
func (o *OAuthEndpoints) HasJwks() bool {
	if o != nil && !IsNil(o.Jwks) {
		return true
	}

	return false
}

// SetJwks gets a reference to the given OidcJwksEndpoint and assigns it to the Jwks field.
func (o *OAuthEndpoints) SetJwks(v OidcJwksEndpoint) {
	o.Jwks = &v
}

// GetSlo returns the Slo field value if set, zero value otherwise.
func (o *OAuthEndpoints) GetSlo() OidcSloEndpoint {
	if o == nil || IsNil(o.Slo) {
		var ret OidcSloEndpoint
		return ret
	}
	return *o.Slo
}

// GetSloOk returns a tuple with the Slo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthEndpoints) GetSloOk() (*OidcSloEndpoint, bool) {
	if o == nil || IsNil(o.Slo) {
		return nil, false
	}
	return o.Slo, true
}

// HasSlo returns a boolean if a field has been set.
func (o *OAuthEndpoints) HasSlo() bool {
	if o != nil && !IsNil(o.Slo) {
		return true
	}

	return false
}

// SetSlo gets a reference to the given OidcSloEndpoint and assigns it to the Slo field.
func (o *OAuthEndpoints) SetSlo(v OidcSloEndpoint) {
	o.Slo = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *OAuthEndpoints) GetToken() OAuthTokenEndpoint {
	if o == nil || IsNil(o.Token) {
		var ret OAuthTokenEndpoint
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthEndpoints) GetTokenOk() (*OAuthTokenEndpoint, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *OAuthEndpoints) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given OAuthTokenEndpoint and assigns it to the Token field.
func (o *OAuthEndpoints) SetToken(v OAuthTokenEndpoint) {
	o.Token = &v
}

// GetUserInfo returns the UserInfo field value if set, zero value otherwise.
func (o *OAuthEndpoints) GetUserInfo() OidcUserInfoEndpoint {
	if o == nil || IsNil(o.UserInfo) {
		var ret OidcUserInfoEndpoint
		return ret
	}
	return *o.UserInfo
}

// GetUserInfoOk returns a tuple with the UserInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthEndpoints) GetUserInfoOk() (*OidcUserInfoEndpoint, bool) {
	if o == nil || IsNil(o.UserInfo) {
		return nil, false
	}
	return o.UserInfo, true
}

// HasUserInfo returns a boolean if a field has been set.
func (o *OAuthEndpoints) HasUserInfo() bool {
	if o != nil && !IsNil(o.UserInfo) {
		return true
	}

	return false
}

// SetUserInfo gets a reference to the given OidcUserInfoEndpoint and assigns it to the UserInfo field.
func (o *OAuthEndpoints) SetUserInfo(v OidcUserInfoEndpoint) {
	o.UserInfo = &v
}

func (o OAuthEndpoints) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OAuthEndpoints) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Authorization) {
		toSerialize["authorization"] = o.Authorization
	}
	if !IsNil(o.Jwks) {
		toSerialize["jwks"] = o.Jwks
	}
	if !IsNil(o.Slo) {
		toSerialize["slo"] = o.Slo
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.UserInfo) {
		toSerialize["userInfo"] = o.UserInfo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OAuthEndpoints) UnmarshalJSON(data []byte) (err error) {
	varOAuthEndpoints := _OAuthEndpoints{}

	err = json.Unmarshal(data, &varOAuthEndpoints)

	if err != nil {
		return err
	}

	*o = OAuthEndpoints(varOAuthEndpoints)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "authorization")
		delete(additionalProperties, "jwks")
		delete(additionalProperties, "slo")
		delete(additionalProperties, "token")
		delete(additionalProperties, "userInfo")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOAuthEndpoints struct {
	value *OAuthEndpoints
	isSet bool
}

func (v NullableOAuthEndpoints) Get() *OAuthEndpoints {
	return v.value
}

func (v *NullableOAuthEndpoints) Set(val *OAuthEndpoints) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuthEndpoints) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuthEndpoints) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuthEndpoints(val *OAuthEndpoints) *NullableOAuthEndpoints {
	return &NullableOAuthEndpoints{value: val, isSet: true}
}

func (v NullableOAuthEndpoints) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuthEndpoints) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
