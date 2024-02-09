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

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
package okta

import (
	"encoding/json"
)

// OAuth2RefreshTokenLinksAllOfRevoke Link to revoke the refresh Token
type OAuth2RefreshTokenLinksAllOfRevoke struct {
	Hints *OAuth2RefreshTokenLinksAllOfRevokeAllOfHints `json:"hints,omitempty"`
	// Link URI
	Href string `json:"href"`
	// Link name
	Title *string `json:"title,omitempty"`
	// The media type of the link. If omitted, it is implicitly `application/json`.
	Type                 *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OAuth2RefreshTokenLinksAllOfRevoke OAuth2RefreshTokenLinksAllOfRevoke

// NewOAuth2RefreshTokenLinksAllOfRevoke instantiates a new OAuth2RefreshTokenLinksAllOfRevoke object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth2RefreshTokenLinksAllOfRevoke(href string) *OAuth2RefreshTokenLinksAllOfRevoke {
	this := OAuth2RefreshTokenLinksAllOfRevoke{}
	this.Href = href
	return &this
}

// NewOAuth2RefreshTokenLinksAllOfRevokeWithDefaults instantiates a new OAuth2RefreshTokenLinksAllOfRevoke object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth2RefreshTokenLinksAllOfRevokeWithDefaults() *OAuth2RefreshTokenLinksAllOfRevoke {
	this := OAuth2RefreshTokenLinksAllOfRevoke{}
	return &this
}

// GetHints returns the Hints field value if set, zero value otherwise.
func (o *OAuth2RefreshTokenLinksAllOfRevoke) GetHints() OAuth2RefreshTokenLinksAllOfRevokeAllOfHints {
	if o == nil || o.Hints == nil {
		var ret OAuth2RefreshTokenLinksAllOfRevokeAllOfHints
		return ret
	}
	return *o.Hints
}

// GetHintsOk returns a tuple with the Hints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2RefreshTokenLinksAllOfRevoke) GetHintsOk() (*OAuth2RefreshTokenLinksAllOfRevokeAllOfHints, bool) {
	if o == nil || o.Hints == nil {
		return nil, false
	}
	return o.Hints, true
}

// HasHints returns a boolean if a field has been set.
func (o *OAuth2RefreshTokenLinksAllOfRevoke) HasHints() bool {
	if o != nil && o.Hints != nil {
		return true
	}

	return false
}

// SetHints gets a reference to the given OAuth2RefreshTokenLinksAllOfRevokeAllOfHints and assigns it to the Hints field.
func (o *OAuth2RefreshTokenLinksAllOfRevoke) SetHints(v OAuth2RefreshTokenLinksAllOfRevokeAllOfHints) {
	o.Hints = &v
}

// GetHref returns the Href field value
func (o *OAuth2RefreshTokenLinksAllOfRevoke) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *OAuth2RefreshTokenLinksAllOfRevoke) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *OAuth2RefreshTokenLinksAllOfRevoke) SetHref(v string) {
	o.Href = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *OAuth2RefreshTokenLinksAllOfRevoke) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2RefreshTokenLinksAllOfRevoke) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *OAuth2RefreshTokenLinksAllOfRevoke) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *OAuth2RefreshTokenLinksAllOfRevoke) SetTitle(v string) {
	o.Title = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *OAuth2RefreshTokenLinksAllOfRevoke) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2RefreshTokenLinksAllOfRevoke) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *OAuth2RefreshTokenLinksAllOfRevoke) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *OAuth2RefreshTokenLinksAllOfRevoke) SetType(v string) {
	o.Type = &v
}

func (o OAuth2RefreshTokenLinksAllOfRevoke) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Hints != nil {
		toSerialize["hints"] = o.Hints
	}
	if true {
		toSerialize["href"] = o.Href
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OAuth2RefreshTokenLinksAllOfRevoke) UnmarshalJSON(bytes []byte) (err error) {
	varOAuth2RefreshTokenLinksAllOfRevoke := _OAuth2RefreshTokenLinksAllOfRevoke{}

	err = json.Unmarshal(bytes, &varOAuth2RefreshTokenLinksAllOfRevoke)
	if err == nil {
		*o = OAuth2RefreshTokenLinksAllOfRevoke(varOAuth2RefreshTokenLinksAllOfRevoke)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "hints")
		delete(additionalProperties, "href")
		delete(additionalProperties, "title")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
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
