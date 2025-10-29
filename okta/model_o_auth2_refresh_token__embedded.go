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

// checks if the OAuth2RefreshTokenEmbedded type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OAuth2RefreshTokenEmbedded{}

// OAuth2RefreshTokenEmbedded The embedded resources related to the object if the `expand` query parameter is specified
type OAuth2RefreshTokenEmbedded struct {
	// The scope objects attached to the Token
	Scopes               []OAuth2RefreshTokenScope `json:"scopes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OAuth2RefreshTokenEmbedded OAuth2RefreshTokenEmbedded

// NewOAuth2RefreshTokenEmbedded instantiates a new OAuth2RefreshTokenEmbedded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth2RefreshTokenEmbedded() *OAuth2RefreshTokenEmbedded {
	this := OAuth2RefreshTokenEmbedded{}
	return &this
}

// NewOAuth2RefreshTokenEmbeddedWithDefaults instantiates a new OAuth2RefreshTokenEmbedded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth2RefreshTokenEmbeddedWithDefaults() *OAuth2RefreshTokenEmbedded {
	this := OAuth2RefreshTokenEmbedded{}
	return &this
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *OAuth2RefreshTokenEmbedded) GetScopes() []OAuth2RefreshTokenScope {
	if o == nil || IsNil(o.Scopes) {
		var ret []OAuth2RefreshTokenScope
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2RefreshTokenEmbedded) GetScopesOk() ([]OAuth2RefreshTokenScope, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *OAuth2RefreshTokenEmbedded) HasScopes() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []OAuth2RefreshTokenScope and assigns it to the Scopes field.
func (o *OAuth2RefreshTokenEmbedded) SetScopes(v []OAuth2RefreshTokenScope) {
	o.Scopes = v
}

func (o OAuth2RefreshTokenEmbedded) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OAuth2RefreshTokenEmbedded) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OAuth2RefreshTokenEmbedded) UnmarshalJSON(data []byte) (err error) {
	varOAuth2RefreshTokenEmbedded := _OAuth2RefreshTokenEmbedded{}

	err = json.Unmarshal(data, &varOAuth2RefreshTokenEmbedded)

	if err != nil {
		return err
	}

	*o = OAuth2RefreshTokenEmbedded(varOAuth2RefreshTokenEmbedded)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "scopes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOAuth2RefreshTokenEmbedded struct {
	value *OAuth2RefreshTokenEmbedded
	isSet bool
}

func (v NullableOAuth2RefreshTokenEmbedded) Get() *OAuth2RefreshTokenEmbedded {
	return v.value
}

func (v *NullableOAuth2RefreshTokenEmbedded) Set(val *OAuth2RefreshTokenEmbedded) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2RefreshTokenEmbedded) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2RefreshTokenEmbedded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2RefreshTokenEmbedded(val *OAuth2RefreshTokenEmbedded) *NullableOAuth2RefreshTokenEmbedded {
	return &NullableOAuth2RefreshTokenEmbedded{value: val, isSet: true}
}

func (v NullableOAuth2RefreshTokenEmbedded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2RefreshTokenEmbedded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
