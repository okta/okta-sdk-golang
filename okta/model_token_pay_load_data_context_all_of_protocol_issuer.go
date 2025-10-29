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

// checks if the TokenPayLoadDataContextAllOfProtocolIssuer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenPayLoadDataContextAllOfProtocolIssuer{}

// TokenPayLoadDataContextAllOfProtocolIssuer The authorization server's issuer identifier
type TokenPayLoadDataContextAllOfProtocolIssuer struct {
	// The authorization server's issuer identifier
	Uri                  *string `json:"uri,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TokenPayLoadDataContextAllOfProtocolIssuer TokenPayLoadDataContextAllOfProtocolIssuer

// NewTokenPayLoadDataContextAllOfProtocolIssuer instantiates a new TokenPayLoadDataContextAllOfProtocolIssuer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenPayLoadDataContextAllOfProtocolIssuer() *TokenPayLoadDataContextAllOfProtocolIssuer {
	this := TokenPayLoadDataContextAllOfProtocolIssuer{}
	return &this
}

// NewTokenPayLoadDataContextAllOfProtocolIssuerWithDefaults instantiates a new TokenPayLoadDataContextAllOfProtocolIssuer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenPayLoadDataContextAllOfProtocolIssuerWithDefaults() *TokenPayLoadDataContextAllOfProtocolIssuer {
	this := TokenPayLoadDataContextAllOfProtocolIssuer{}
	return &this
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *TokenPayLoadDataContextAllOfProtocolIssuer) GetUri() string {
	if o == nil || IsNil(o.Uri) {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenPayLoadDataContextAllOfProtocolIssuer) GetUriOk() (*string, bool) {
	if o == nil || IsNil(o.Uri) {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *TokenPayLoadDataContextAllOfProtocolIssuer) HasUri() bool {
	if o != nil && !IsNil(o.Uri) {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *TokenPayLoadDataContextAllOfProtocolIssuer) SetUri(v string) {
	o.Uri = &v
}

func (o TokenPayLoadDataContextAllOfProtocolIssuer) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenPayLoadDataContextAllOfProtocolIssuer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Uri) {
		toSerialize["uri"] = o.Uri
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TokenPayLoadDataContextAllOfProtocolIssuer) UnmarshalJSON(data []byte) (err error) {
	varTokenPayLoadDataContextAllOfProtocolIssuer := _TokenPayLoadDataContextAllOfProtocolIssuer{}

	err = json.Unmarshal(data, &varTokenPayLoadDataContextAllOfProtocolIssuer)

	if err != nil {
		return err
	}

	*o = TokenPayLoadDataContextAllOfProtocolIssuer(varTokenPayLoadDataContextAllOfProtocolIssuer)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "uri")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTokenPayLoadDataContextAllOfProtocolIssuer struct {
	value *TokenPayLoadDataContextAllOfProtocolIssuer
	isSet bool
}

func (v NullableTokenPayLoadDataContextAllOfProtocolIssuer) Get() *TokenPayLoadDataContextAllOfProtocolIssuer {
	return v.value
}

func (v *NullableTokenPayLoadDataContextAllOfProtocolIssuer) Set(val *TokenPayLoadDataContextAllOfProtocolIssuer) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenPayLoadDataContextAllOfProtocolIssuer) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenPayLoadDataContextAllOfProtocolIssuer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenPayLoadDataContextAllOfProtocolIssuer(val *TokenPayLoadDataContextAllOfProtocolIssuer) *NullableTokenPayLoadDataContextAllOfProtocolIssuer {
	return &NullableTokenPayLoadDataContextAllOfProtocolIssuer{value: val, isSet: true}
}

func (v NullableTokenPayLoadDataContextAllOfProtocolIssuer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenPayLoadDataContextAllOfProtocolIssuer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
