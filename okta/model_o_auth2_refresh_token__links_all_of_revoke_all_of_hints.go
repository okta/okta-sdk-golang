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

// OAuth2RefreshTokenLinksAllOfRevokeAllOfHints struct for OAuth2RefreshTokenLinksAllOfRevokeAllOfHints
type OAuth2RefreshTokenLinksAllOfRevokeAllOfHints struct {
	Allow []string `json:"allow,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OAuth2RefreshTokenLinksAllOfRevokeAllOfHints OAuth2RefreshTokenLinksAllOfRevokeAllOfHints

// NewOAuth2RefreshTokenLinksAllOfRevokeAllOfHints instantiates a new OAuth2RefreshTokenLinksAllOfRevokeAllOfHints object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth2RefreshTokenLinksAllOfRevokeAllOfHints() *OAuth2RefreshTokenLinksAllOfRevokeAllOfHints {
	this := OAuth2RefreshTokenLinksAllOfRevokeAllOfHints{}
	return &this
}

// NewOAuth2RefreshTokenLinksAllOfRevokeAllOfHintsWithDefaults instantiates a new OAuth2RefreshTokenLinksAllOfRevokeAllOfHints object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth2RefreshTokenLinksAllOfRevokeAllOfHintsWithDefaults() *OAuth2RefreshTokenLinksAllOfRevokeAllOfHints {
	this := OAuth2RefreshTokenLinksAllOfRevokeAllOfHints{}
	return &this
}

// GetAllow returns the Allow field value if set, zero value otherwise.
func (o *OAuth2RefreshTokenLinksAllOfRevokeAllOfHints) GetAllow() []string {
	if o == nil || o.Allow == nil {
		var ret []string
		return ret
	}
	return o.Allow
}

// GetAllowOk returns a tuple with the Allow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2RefreshTokenLinksAllOfRevokeAllOfHints) GetAllowOk() ([]string, bool) {
	if o == nil || o.Allow == nil {
		return nil, false
	}
	return o.Allow, true
}

// HasAllow returns a boolean if a field has been set.
func (o *OAuth2RefreshTokenLinksAllOfRevokeAllOfHints) HasAllow() bool {
	if o != nil && o.Allow != nil {
		return true
	}

	return false
}

// SetAllow gets a reference to the given []string and assigns it to the Allow field.
func (o *OAuth2RefreshTokenLinksAllOfRevokeAllOfHints) SetAllow(v []string) {
	o.Allow = v
}

func (o OAuth2RefreshTokenLinksAllOfRevokeAllOfHints) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Allow != nil {
		toSerialize["allow"] = o.Allow
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OAuth2RefreshTokenLinksAllOfRevokeAllOfHints) UnmarshalJSON(bytes []byte) (err error) {
	varOAuth2RefreshTokenLinksAllOfRevokeAllOfHints := _OAuth2RefreshTokenLinksAllOfRevokeAllOfHints{}

	err = json.Unmarshal(bytes, &varOAuth2RefreshTokenLinksAllOfRevokeAllOfHints)
	if err == nil {
		*o = OAuth2RefreshTokenLinksAllOfRevokeAllOfHints(varOAuth2RefreshTokenLinksAllOfRevokeAllOfHints)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "allow")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableOAuth2RefreshTokenLinksAllOfRevokeAllOfHints struct {
	value *OAuth2RefreshTokenLinksAllOfRevokeAllOfHints
	isSet bool
}

func (v NullableOAuth2RefreshTokenLinksAllOfRevokeAllOfHints) Get() *OAuth2RefreshTokenLinksAllOfRevokeAllOfHints {
	return v.value
}

func (v *NullableOAuth2RefreshTokenLinksAllOfRevokeAllOfHints) Set(val *OAuth2RefreshTokenLinksAllOfRevokeAllOfHints) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2RefreshTokenLinksAllOfRevokeAllOfHints) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2RefreshTokenLinksAllOfRevokeAllOfHints) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2RefreshTokenLinksAllOfRevokeAllOfHints(val *OAuth2RefreshTokenLinksAllOfRevokeAllOfHints) *NullableOAuth2RefreshTokenLinksAllOfRevokeAllOfHints {
	return &NullableOAuth2RefreshTokenLinksAllOfRevokeAllOfHints{value: val, isSet: true}
}

func (v NullableOAuth2RefreshTokenLinksAllOfRevokeAllOfHints) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2RefreshTokenLinksAllOfRevokeAllOfHints) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

