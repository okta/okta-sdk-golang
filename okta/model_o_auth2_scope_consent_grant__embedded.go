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
)

// checks if the OAuth2ScopeConsentGrantEmbedded type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OAuth2ScopeConsentGrantEmbedded{}

// OAuth2ScopeConsentGrantEmbedded Embedded resources related to the Grant
type OAuth2ScopeConsentGrantEmbedded struct {
	Scope                *OAuth2ScopeConsentGrantEmbeddedScope `json:"scope,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OAuth2ScopeConsentGrantEmbedded OAuth2ScopeConsentGrantEmbedded

// NewOAuth2ScopeConsentGrantEmbedded instantiates a new OAuth2ScopeConsentGrantEmbedded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth2ScopeConsentGrantEmbedded() *OAuth2ScopeConsentGrantEmbedded {
	this := OAuth2ScopeConsentGrantEmbedded{}
	return &this
}

// NewOAuth2ScopeConsentGrantEmbeddedWithDefaults instantiates a new OAuth2ScopeConsentGrantEmbedded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth2ScopeConsentGrantEmbeddedWithDefaults() *OAuth2ScopeConsentGrantEmbedded {
	this := OAuth2ScopeConsentGrantEmbedded{}
	return &this
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *OAuth2ScopeConsentGrantEmbedded) GetScope() OAuth2ScopeConsentGrantEmbeddedScope {
	if o == nil || IsNil(o.Scope) {
		var ret OAuth2ScopeConsentGrantEmbeddedScope
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ScopeConsentGrantEmbedded) GetScopeOk() (*OAuth2ScopeConsentGrantEmbeddedScope, bool) {
	if o == nil || IsNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *OAuth2ScopeConsentGrantEmbedded) HasScope() bool {
	if o != nil && !IsNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given OAuth2ScopeConsentGrantEmbeddedScope and assigns it to the Scope field.
func (o *OAuth2ScopeConsentGrantEmbedded) SetScope(v OAuth2ScopeConsentGrantEmbeddedScope) {
	o.Scope = &v
}

func (o OAuth2ScopeConsentGrantEmbedded) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OAuth2ScopeConsentGrantEmbedded) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OAuth2ScopeConsentGrantEmbedded) UnmarshalJSON(data []byte) (err error) {
	varOAuth2ScopeConsentGrantEmbedded := _OAuth2ScopeConsentGrantEmbedded{}

	err = json.Unmarshal(data, &varOAuth2ScopeConsentGrantEmbedded)

	if err != nil {
		return err
	}

	*o = OAuth2ScopeConsentGrantEmbedded(varOAuth2ScopeConsentGrantEmbedded)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "scope")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOAuth2ScopeConsentGrantEmbedded struct {
	value *OAuth2ScopeConsentGrantEmbedded
	isSet bool
}

func (v NullableOAuth2ScopeConsentGrantEmbedded) Get() *OAuth2ScopeConsentGrantEmbedded {
	return v.value
}

func (v *NullableOAuth2ScopeConsentGrantEmbedded) Set(val *OAuth2ScopeConsentGrantEmbedded) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2ScopeConsentGrantEmbedded) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2ScopeConsentGrantEmbedded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2ScopeConsentGrantEmbedded(val *OAuth2ScopeConsentGrantEmbedded) *NullableOAuth2ScopeConsentGrantEmbedded {
	return &NullableOAuth2ScopeConsentGrantEmbedded{value: val, isSet: true}
}

func (v NullableOAuth2ScopeConsentGrantEmbedded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2ScopeConsentGrantEmbedded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
