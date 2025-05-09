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

// SecurityEventTokenRequestJwtHeader JSON Web Token header for a Security Event Token
type SecurityEventTokenRequestJwtHeader struct {
	// Algorithm used to sign or encrypt the JWT
	Alg string `json:"alg"`
	// Key ID used to sign or encrypt the JWT
	Kid string `json:"kid"`
	// The type of content being signed or encrypted
	Typ string `json:"typ"`
	AdditionalProperties map[string]interface{}
}

type _SecurityEventTokenRequestJwtHeader SecurityEventTokenRequestJwtHeader

// NewSecurityEventTokenRequestJwtHeader instantiates a new SecurityEventTokenRequestJwtHeader object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityEventTokenRequestJwtHeader(alg string, kid string, typ string) *SecurityEventTokenRequestJwtHeader {
	this := SecurityEventTokenRequestJwtHeader{}
	this.Alg = alg
	this.Kid = kid
	this.Typ = typ
	return &this
}

// NewSecurityEventTokenRequestJwtHeaderWithDefaults instantiates a new SecurityEventTokenRequestJwtHeader object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityEventTokenRequestJwtHeaderWithDefaults() *SecurityEventTokenRequestJwtHeader {
	this := SecurityEventTokenRequestJwtHeader{}
	return &this
}

// GetAlg returns the Alg field value
func (o *SecurityEventTokenRequestJwtHeader) GetAlg() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Alg
}

// GetAlgOk returns a tuple with the Alg field value
// and a boolean to check if the value has been set.
func (o *SecurityEventTokenRequestJwtHeader) GetAlgOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Alg, true
}

// SetAlg sets field value
func (o *SecurityEventTokenRequestJwtHeader) SetAlg(v string) {
	o.Alg = v
}

// GetKid returns the Kid field value
func (o *SecurityEventTokenRequestJwtHeader) GetKid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kid
}

// GetKidOk returns a tuple with the Kid field value
// and a boolean to check if the value has been set.
func (o *SecurityEventTokenRequestJwtHeader) GetKidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kid, true
}

// SetKid sets field value
func (o *SecurityEventTokenRequestJwtHeader) SetKid(v string) {
	o.Kid = v
}

// GetTyp returns the Typ field value
func (o *SecurityEventTokenRequestJwtHeader) GetTyp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Typ
}

// GetTypOk returns a tuple with the Typ field value
// and a boolean to check if the value has been set.
func (o *SecurityEventTokenRequestJwtHeader) GetTypOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Typ, true
}

// SetTyp sets field value
func (o *SecurityEventTokenRequestJwtHeader) SetTyp(v string) {
	o.Typ = v
}

func (o SecurityEventTokenRequestJwtHeader) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["alg"] = o.Alg
	}
	if true {
		toSerialize["kid"] = o.Kid
	}
	if true {
		toSerialize["typ"] = o.Typ
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SecurityEventTokenRequestJwtHeader) UnmarshalJSON(bytes []byte) (err error) {
	varSecurityEventTokenRequestJwtHeader := _SecurityEventTokenRequestJwtHeader{}

	err = json.Unmarshal(bytes, &varSecurityEventTokenRequestJwtHeader)
	if err == nil {
		*o = SecurityEventTokenRequestJwtHeader(varSecurityEventTokenRequestJwtHeader)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "alg")
		delete(additionalProperties, "kid")
		delete(additionalProperties, "typ")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableSecurityEventTokenRequestJwtHeader struct {
	value *SecurityEventTokenRequestJwtHeader
	isSet bool
}

func (v NullableSecurityEventTokenRequestJwtHeader) Get() *SecurityEventTokenRequestJwtHeader {
	return v.value
}

func (v *NullableSecurityEventTokenRequestJwtHeader) Set(val *SecurityEventTokenRequestJwtHeader) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityEventTokenRequestJwtHeader) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityEventTokenRequestJwtHeader) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityEventTokenRequestJwtHeader(val *SecurityEventTokenRequestJwtHeader) *NullableSecurityEventTokenRequestJwtHeader {
	return &NullableSecurityEventTokenRequestJwtHeader{value: val, isSet: true}
}

func (v NullableSecurityEventTokenRequestJwtHeader) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityEventTokenRequestJwtHeader) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

