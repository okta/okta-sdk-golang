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

// TokenPayLoadData struct for TokenPayLoadData
type TokenPayLoadData struct {
	Context *TokenPayLoadDataContext `json:"context,omitempty"`
	// Provides information on the properties of the ID token that Okta has generated, including the existing claims that it contains
	Identity *BaseToken `json:"identity,omitempty"`
	Access *TokenPayLoadDataAccess `json:"access,omitempty"`
	RefreshToken *RefreshToken `json:"refresh_token,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TokenPayLoadData TokenPayLoadData

// NewTokenPayLoadData instantiates a new TokenPayLoadData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenPayLoadData() *TokenPayLoadData {
	this := TokenPayLoadData{}
	return &this
}

// NewTokenPayLoadDataWithDefaults instantiates a new TokenPayLoadData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenPayLoadDataWithDefaults() *TokenPayLoadData {
	this := TokenPayLoadData{}
	return &this
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *TokenPayLoadData) GetContext() TokenPayLoadDataContext {
	if o == nil || o.Context == nil {
		var ret TokenPayLoadDataContext
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenPayLoadData) GetContextOk() (*TokenPayLoadDataContext, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *TokenPayLoadData) HasContext() bool {
	if o != nil && o.Context != nil {
		return true
	}

	return false
}

// SetContext gets a reference to the given TokenPayLoadDataContext and assigns it to the Context field.
func (o *TokenPayLoadData) SetContext(v TokenPayLoadDataContext) {
	o.Context = &v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *TokenPayLoadData) GetIdentity() BaseToken {
	if o == nil || o.Identity == nil {
		var ret BaseToken
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenPayLoadData) GetIdentityOk() (*BaseToken, bool) {
	if o == nil || o.Identity == nil {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *TokenPayLoadData) HasIdentity() bool {
	if o != nil && o.Identity != nil {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given BaseToken and assigns it to the Identity field.
func (o *TokenPayLoadData) SetIdentity(v BaseToken) {
	o.Identity = &v
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *TokenPayLoadData) GetAccess() TokenPayLoadDataAccess {
	if o == nil || o.Access == nil {
		var ret TokenPayLoadDataAccess
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenPayLoadData) GetAccessOk() (*TokenPayLoadDataAccess, bool) {
	if o == nil || o.Access == nil {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *TokenPayLoadData) HasAccess() bool {
	if o != nil && o.Access != nil {
		return true
	}

	return false
}

// SetAccess gets a reference to the given TokenPayLoadDataAccess and assigns it to the Access field.
func (o *TokenPayLoadData) SetAccess(v TokenPayLoadDataAccess) {
	o.Access = &v
}

// GetRefreshToken returns the RefreshToken field value if set, zero value otherwise.
func (o *TokenPayLoadData) GetRefreshToken() RefreshToken {
	if o == nil || o.RefreshToken == nil {
		var ret RefreshToken
		return ret
	}
	return *o.RefreshToken
}

// GetRefreshTokenOk returns a tuple with the RefreshToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenPayLoadData) GetRefreshTokenOk() (*RefreshToken, bool) {
	if o == nil || o.RefreshToken == nil {
		return nil, false
	}
	return o.RefreshToken, true
}

// HasRefreshToken returns a boolean if a field has been set.
func (o *TokenPayLoadData) HasRefreshToken() bool {
	if o != nil && o.RefreshToken != nil {
		return true
	}

	return false
}

// SetRefreshToken gets a reference to the given RefreshToken and assigns it to the RefreshToken field.
func (o *TokenPayLoadData) SetRefreshToken(v RefreshToken) {
	o.RefreshToken = &v
}

func (o TokenPayLoadData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Context != nil {
		toSerialize["context"] = o.Context
	}
	if o.Identity != nil {
		toSerialize["identity"] = o.Identity
	}
	if o.Access != nil {
		toSerialize["access"] = o.Access
	}
	if o.RefreshToken != nil {
		toSerialize["refresh_token"] = o.RefreshToken
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TokenPayLoadData) UnmarshalJSON(bytes []byte) (err error) {
	varTokenPayLoadData := _TokenPayLoadData{}

	err = json.Unmarshal(bytes, &varTokenPayLoadData)
	if err == nil {
		*o = TokenPayLoadData(varTokenPayLoadData)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "context")
		delete(additionalProperties, "identity")
		delete(additionalProperties, "access")
		delete(additionalProperties, "refresh_token")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableTokenPayLoadData struct {
	value *TokenPayLoadData
	isSet bool
}

func (v NullableTokenPayLoadData) Get() *TokenPayLoadData {
	return v.value
}

func (v *NullableTokenPayLoadData) Set(val *TokenPayLoadData) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenPayLoadData) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenPayLoadData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenPayLoadData(val *TokenPayLoadData) *NullableTokenPayLoadData {
	return &NullableTokenPayLoadData{value: val, isSet: true}
}

func (v NullableTokenPayLoadData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenPayLoadData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

