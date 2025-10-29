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

// checks if the OAuth2RefreshTokenScopeLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OAuth2RefreshTokenScopeLinks{}

// OAuth2RefreshTokenScopeLinks Specifies link relations (see [Web Linking](https://www.rfc-editor.org/rfc/rfc8288)) available for the current status of an application using the [JSON Hypertext Application Language](https://datatracker.ietf.org/doc/html/draft-kelly-json-hal-06) specification. This object is used for dynamic discovery of related resources and lifecycle operations.
type OAuth2RefreshTokenScopeLinks struct {
	// Link to Scope resource
	Scope                *OfflineAccessScopeResourceHrefObject `json:"scope,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OAuth2RefreshTokenScopeLinks OAuth2RefreshTokenScopeLinks

// NewOAuth2RefreshTokenScopeLinks instantiates a new OAuth2RefreshTokenScopeLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth2RefreshTokenScopeLinks() *OAuth2RefreshTokenScopeLinks {
	this := OAuth2RefreshTokenScopeLinks{}
	return &this
}

// NewOAuth2RefreshTokenScopeLinksWithDefaults instantiates a new OAuth2RefreshTokenScopeLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth2RefreshTokenScopeLinksWithDefaults() *OAuth2RefreshTokenScopeLinks {
	this := OAuth2RefreshTokenScopeLinks{}
	return &this
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *OAuth2RefreshTokenScopeLinks) GetScope() OfflineAccessScopeResourceHrefObject {
	if o == nil || IsNil(o.Scope) {
		var ret OfflineAccessScopeResourceHrefObject
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2RefreshTokenScopeLinks) GetScopeOk() (*OfflineAccessScopeResourceHrefObject, bool) {
	if o == nil || IsNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *OAuth2RefreshTokenScopeLinks) HasScope() bool {
	if o != nil && !IsNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given OfflineAccessScopeResourceHrefObject and assigns it to the Scope field.
func (o *OAuth2RefreshTokenScopeLinks) SetScope(v OfflineAccessScopeResourceHrefObject) {
	o.Scope = &v
}

func (o OAuth2RefreshTokenScopeLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OAuth2RefreshTokenScopeLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OAuth2RefreshTokenScopeLinks) UnmarshalJSON(data []byte) (err error) {
	varOAuth2RefreshTokenScopeLinks := _OAuth2RefreshTokenScopeLinks{}

	err = json.Unmarshal(data, &varOAuth2RefreshTokenScopeLinks)

	if err != nil {
		return err
	}

	*o = OAuth2RefreshTokenScopeLinks(varOAuth2RefreshTokenScopeLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "scope")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOAuth2RefreshTokenScopeLinks struct {
	value *OAuth2RefreshTokenScopeLinks
	isSet bool
}

func (v NullableOAuth2RefreshTokenScopeLinks) Get() *OAuth2RefreshTokenScopeLinks {
	return v.value
}

func (v *NullableOAuth2RefreshTokenScopeLinks) Set(val *OAuth2RefreshTokenScopeLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2RefreshTokenScopeLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2RefreshTokenScopeLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2RefreshTokenScopeLinks(val *OAuth2RefreshTokenScopeLinks) *NullableOAuth2RefreshTokenScopeLinks {
	return &NullableOAuth2RefreshTokenScopeLinks{value: val, isSet: true}
}

func (v NullableOAuth2RefreshTokenScopeLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2RefreshTokenScopeLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
