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

// checks if the OAuth2ClientLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OAuth2ClientLinks{}

// OAuth2ClientLinks struct for OAuth2ClientLinks
type OAuth2ClientLinks struct {
	Self *HrefObjectSelfLink `json:"self,omitempty"`
	// Link to the grant resources
	Grants *GrantResourcesHrefObject `json:"grants,omitempty"`
	// Link to the token resources
	Tokens               *TokenResourcesHrefObject `json:"tokens,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OAuth2ClientLinks OAuth2ClientLinks

// NewOAuth2ClientLinks instantiates a new OAuth2ClientLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth2ClientLinks() *OAuth2ClientLinks {
	this := OAuth2ClientLinks{}
	return &this
}

// NewOAuth2ClientLinksWithDefaults instantiates a new OAuth2ClientLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth2ClientLinksWithDefaults() *OAuth2ClientLinks {
	this := OAuth2ClientLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *OAuth2ClientLinks) GetSelf() HrefObjectSelfLink {
	if o == nil || IsNil(o.Self) {
		var ret HrefObjectSelfLink
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ClientLinks) GetSelfOk() (*HrefObjectSelfLink, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *OAuth2ClientLinks) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given HrefObjectSelfLink and assigns it to the Self field.
func (o *OAuth2ClientLinks) SetSelf(v HrefObjectSelfLink) {
	o.Self = &v
}

// GetGrants returns the Grants field value if set, zero value otherwise.
func (o *OAuth2ClientLinks) GetGrants() GrantResourcesHrefObject {
	if o == nil || IsNil(o.Grants) {
		var ret GrantResourcesHrefObject
		return ret
	}
	return *o.Grants
}

// GetGrantsOk returns a tuple with the Grants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ClientLinks) GetGrantsOk() (*GrantResourcesHrefObject, bool) {
	if o == nil || IsNil(o.Grants) {
		return nil, false
	}
	return o.Grants, true
}

// HasGrants returns a boolean if a field has been set.
func (o *OAuth2ClientLinks) HasGrants() bool {
	if o != nil && !IsNil(o.Grants) {
		return true
	}

	return false
}

// SetGrants gets a reference to the given GrantResourcesHrefObject and assigns it to the Grants field.
func (o *OAuth2ClientLinks) SetGrants(v GrantResourcesHrefObject) {
	o.Grants = &v
}

// GetTokens returns the Tokens field value if set, zero value otherwise.
func (o *OAuth2ClientLinks) GetTokens() TokenResourcesHrefObject {
	if o == nil || IsNil(o.Tokens) {
		var ret TokenResourcesHrefObject
		return ret
	}
	return *o.Tokens
}

// GetTokensOk returns a tuple with the Tokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ClientLinks) GetTokensOk() (*TokenResourcesHrefObject, bool) {
	if o == nil || IsNil(o.Tokens) {
		return nil, false
	}
	return o.Tokens, true
}

// HasTokens returns a boolean if a field has been set.
func (o *OAuth2ClientLinks) HasTokens() bool {
	if o != nil && !IsNil(o.Tokens) {
		return true
	}

	return false
}

// SetTokens gets a reference to the given TokenResourcesHrefObject and assigns it to the Tokens field.
func (o *OAuth2ClientLinks) SetTokens(v TokenResourcesHrefObject) {
	o.Tokens = &v
}

func (o OAuth2ClientLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OAuth2ClientLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	if !IsNil(o.Grants) {
		toSerialize["grants"] = o.Grants
	}
	if !IsNil(o.Tokens) {
		toSerialize["tokens"] = o.Tokens
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OAuth2ClientLinks) UnmarshalJSON(data []byte) (err error) {
	varOAuth2ClientLinks := _OAuth2ClientLinks{}

	err = json.Unmarshal(data, &varOAuth2ClientLinks)

	if err != nil {
		return err
	}

	*o = OAuth2ClientLinks(varOAuth2ClientLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "grants")
		delete(additionalProperties, "tokens")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOAuth2ClientLinks struct {
	value *OAuth2ClientLinks
	isSet bool
}

func (v NullableOAuth2ClientLinks) Get() *OAuth2ClientLinks {
	return v.value
}

func (v *NullableOAuth2ClientLinks) Set(val *OAuth2ClientLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2ClientLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2ClientLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2ClientLinks(val *OAuth2ClientLinks) *NullableOAuth2ClientLinks {
	return &NullableOAuth2ClientLinks{value: val, isSet: true}
}

func (v NullableOAuth2ClientLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2ClientLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
