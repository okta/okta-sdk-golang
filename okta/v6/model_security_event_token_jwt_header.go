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

package okta

import (
	"encoding/json"
	"fmt"
)

// SecurityEventTokenJwtHeader JSON Web Token header for a Security Event Token sent by the SSF Transmitter
type SecurityEventTokenJwtHeader struct {
	// Algorithm used to sign or encrypt the JWT
	Alg string `json:"alg"`
	// Key ID used to sign or encrypt the JWT
	Kid string `json:"kid"`
	// The type of content being signed or encrypted
	Typ string `json:"typ"`
	AdditionalProperties map[string]interface{}
}

type _SecurityEventTokenJwtHeader SecurityEventTokenJwtHeader

// NewSecurityEventTokenJwtHeader instantiates a new SecurityEventTokenJwtHeader object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityEventTokenJwtHeader(alg string, kid string, typ string) *SecurityEventTokenJwtHeader {
	this := SecurityEventTokenJwtHeader{}
	this.Alg = alg
	this.Kid = kid
	this.Typ = typ
	return &this
}

// NewSecurityEventTokenJwtHeaderWithDefaults instantiates a new SecurityEventTokenJwtHeader object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityEventTokenJwtHeaderWithDefaults() *SecurityEventTokenJwtHeader {
	this := SecurityEventTokenJwtHeader{}
	return &this
}

// GetAlg returns the Alg field value
func (o *SecurityEventTokenJwtHeader) GetAlg() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Alg
}

// GetAlgOk returns a tuple with the Alg field value
// and a boolean to check if the value has been set.
func (o *SecurityEventTokenJwtHeader) GetAlgOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Alg, true
}

// SetAlg sets field value
func (o *SecurityEventTokenJwtHeader) SetAlg(v string) {
	o.Alg = v
}

// GetKid returns the Kid field value
func (o *SecurityEventTokenJwtHeader) GetKid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kid
}

// GetKidOk returns a tuple with the Kid field value
// and a boolean to check if the value has been set.
func (o *SecurityEventTokenJwtHeader) GetKidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kid, true
}

// SetKid sets field value
func (o *SecurityEventTokenJwtHeader) SetKid(v string) {
	o.Kid = v
}

// GetTyp returns the Typ field value
func (o *SecurityEventTokenJwtHeader) GetTyp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Typ
}

// GetTypOk returns a tuple with the Typ field value
// and a boolean to check if the value has been set.
func (o *SecurityEventTokenJwtHeader) GetTypOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Typ, true
}

// SetTyp sets field value
func (o *SecurityEventTokenJwtHeader) SetTyp(v string) {
	o.Typ = v
}

func (o SecurityEventTokenJwtHeader) MarshalJSON() ([]byte, error) {
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

func (o *SecurityEventTokenJwtHeader) UnmarshalJSON(bytes []byte) (err error) {
	varSecurityEventTokenJwtHeader := _SecurityEventTokenJwtHeader{}

	err = json.Unmarshal(bytes, &varSecurityEventTokenJwtHeader)
	if err == nil {
		*o = SecurityEventTokenJwtHeader(varSecurityEventTokenJwtHeader)
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

type NullableSecurityEventTokenJwtHeader struct {
	value *SecurityEventTokenJwtHeader
	isSet bool
}

func (v NullableSecurityEventTokenJwtHeader) Get() *SecurityEventTokenJwtHeader {
	return v.value
}

func (v *NullableSecurityEventTokenJwtHeader) Set(val *SecurityEventTokenJwtHeader) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityEventTokenJwtHeader) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityEventTokenJwtHeader) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityEventTokenJwtHeader(val *SecurityEventTokenJwtHeader) *NullableSecurityEventTokenJwtHeader {
	return &NullableSecurityEventTokenJwtHeader{value: val, isSet: true}
}

func (v NullableSecurityEventTokenJwtHeader) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityEventTokenJwtHeader) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

