/*
Okta Admin Management API

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

// checks if the OidcIssuer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OidcIssuer{}

// OidcIssuer OIDC issuer identifier (issuer URL / expected `iss` value) for the external IdP
type OidcIssuer struct {
	// Issuer URL identifying the external IdP (the expected `iss` value in OIDC tokens)
	Url                  *string `json:"url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OidcIssuer OidcIssuer

// NewOidcIssuer instantiates a new OidcIssuer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOidcIssuer() *OidcIssuer {
	this := OidcIssuer{}
	return &this
}

// NewOidcIssuerWithDefaults instantiates a new OidcIssuer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOidcIssuerWithDefaults() *OidcIssuer {
	this := OidcIssuer{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *OidcIssuer) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OidcIssuer) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *OidcIssuer) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *OidcIssuer) SetUrl(v string) {
	o.Url = &v
}

func (o OidcIssuer) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OidcIssuer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OidcIssuer) UnmarshalJSON(data []byte) (err error) {
	varOidcIssuer := _OidcIssuer{}

	err = json.Unmarshal(data, &varOidcIssuer)

	if err != nil {
		return err
	}

	*o = OidcIssuer(varOidcIssuer)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOidcIssuer struct {
	value *OidcIssuer
	isSet bool
}

func (v NullableOidcIssuer) Get() *OidcIssuer {
	return v.value
}

func (v *NullableOidcIssuer) Set(val *OidcIssuer) {
	v.value = val
	v.isSet = true
}

func (v NullableOidcIssuer) IsSet() bool {
	return v.isSet
}

func (v *NullableOidcIssuer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOidcIssuer(val *OidcIssuer) *NullableOidcIssuer {
	return &NullableOidcIssuer{value: val, isSet: true}
}

func (v NullableOidcIssuer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOidcIssuer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
