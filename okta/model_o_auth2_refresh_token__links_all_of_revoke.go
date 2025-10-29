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

// checks if the OAuth2RefreshTokenLinksAllOfRevoke type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OAuth2RefreshTokenLinksAllOfRevoke{}

// OAuth2RefreshTokenLinksAllOfRevoke Link to revoke the refresh Token
type OAuth2RefreshTokenLinksAllOfRevoke struct {
	// Link URI
	Href                 *string                                       `json:"href,omitempty"`
	Hints                *OAuth2RefreshTokenLinksAllOfRevokeAllOfHints `json:"hints,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OAuth2RefreshTokenLinksAllOfRevoke OAuth2RefreshTokenLinksAllOfRevoke

// NewOAuth2RefreshTokenLinksAllOfRevoke instantiates a new OAuth2RefreshTokenLinksAllOfRevoke object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth2RefreshTokenLinksAllOfRevoke() *OAuth2RefreshTokenLinksAllOfRevoke {
	this := OAuth2RefreshTokenLinksAllOfRevoke{}
	return &this
}

// NewOAuth2RefreshTokenLinksAllOfRevokeWithDefaults instantiates a new OAuth2RefreshTokenLinksAllOfRevoke object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth2RefreshTokenLinksAllOfRevokeWithDefaults() *OAuth2RefreshTokenLinksAllOfRevoke {
	this := OAuth2RefreshTokenLinksAllOfRevoke{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *OAuth2RefreshTokenLinksAllOfRevoke) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2RefreshTokenLinksAllOfRevoke) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *OAuth2RefreshTokenLinksAllOfRevoke) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *OAuth2RefreshTokenLinksAllOfRevoke) SetHref(v string) {
	o.Href = &v
}

// GetHints returns the Hints field value if set, zero value otherwise.
func (o *OAuth2RefreshTokenLinksAllOfRevoke) GetHints() OAuth2RefreshTokenLinksAllOfRevokeAllOfHints {
	if o == nil || IsNil(o.Hints) {
		var ret OAuth2RefreshTokenLinksAllOfRevokeAllOfHints
		return ret
	}
	return *o.Hints
}

// GetHintsOk returns a tuple with the Hints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2RefreshTokenLinksAllOfRevoke) GetHintsOk() (*OAuth2RefreshTokenLinksAllOfRevokeAllOfHints, bool) {
	if o == nil || IsNil(o.Hints) {
		return nil, false
	}
	return o.Hints, true
}

// HasHints returns a boolean if a field has been set.
func (o *OAuth2RefreshTokenLinksAllOfRevoke) HasHints() bool {
	if o != nil && !IsNil(o.Hints) {
		return true
	}

	return false
}

// SetHints gets a reference to the given OAuth2RefreshTokenLinksAllOfRevokeAllOfHints and assigns it to the Hints field.
func (o *OAuth2RefreshTokenLinksAllOfRevoke) SetHints(v OAuth2RefreshTokenLinksAllOfRevokeAllOfHints) {
	o.Hints = &v
}

func (o OAuth2RefreshTokenLinksAllOfRevoke) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OAuth2RefreshTokenLinksAllOfRevoke) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !IsNil(o.Hints) {
		toSerialize["hints"] = o.Hints
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OAuth2RefreshTokenLinksAllOfRevoke) UnmarshalJSON(data []byte) (err error) {
	varOAuth2RefreshTokenLinksAllOfRevoke := _OAuth2RefreshTokenLinksAllOfRevoke{}

	err = json.Unmarshal(data, &varOAuth2RefreshTokenLinksAllOfRevoke)

	if err != nil {
		return err
	}

	*o = OAuth2RefreshTokenLinksAllOfRevoke(varOAuth2RefreshTokenLinksAllOfRevoke)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "href")
		delete(additionalProperties, "hints")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOAuth2RefreshTokenLinksAllOfRevoke struct {
	value *OAuth2RefreshTokenLinksAllOfRevoke
	isSet bool
}

func (v NullableOAuth2RefreshTokenLinksAllOfRevoke) Get() *OAuth2RefreshTokenLinksAllOfRevoke {
	return v.value
}

func (v *NullableOAuth2RefreshTokenLinksAllOfRevoke) Set(val *OAuth2RefreshTokenLinksAllOfRevoke) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2RefreshTokenLinksAllOfRevoke) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2RefreshTokenLinksAllOfRevoke) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2RefreshTokenLinksAllOfRevoke(val *OAuth2RefreshTokenLinksAllOfRevoke) *NullableOAuth2RefreshTokenLinksAllOfRevoke {
	return &NullableOAuth2RefreshTokenLinksAllOfRevoke{value: val, isSet: true}
}

func (v NullableOAuth2RefreshTokenLinksAllOfRevoke) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2RefreshTokenLinksAllOfRevoke) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
