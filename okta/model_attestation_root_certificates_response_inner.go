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

// checks if the AttestationRootCertificatesResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AttestationRootCertificatesResponseInner{}

// AttestationRootCertificatesResponseInner struct for AttestationRootCertificatesResponseInner
type AttestationRootCertificatesResponseInner struct {
	// X.509 certificate chain
	X5c *string `json:"x5c,omitempty"`
	// SHA-256 hash (thumbprint) of the X.509 certificate
	X5tS256 *string `json:"x5t#S256,omitempty"`
	// Issuer of certificate
	Iss *string `json:"iss,omitempty"`
	// Expiry date of certificate
	Exp                  *string `json:"exp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AttestationRootCertificatesResponseInner AttestationRootCertificatesResponseInner

// NewAttestationRootCertificatesResponseInner instantiates a new AttestationRootCertificatesResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttestationRootCertificatesResponseInner() *AttestationRootCertificatesResponseInner {
	this := AttestationRootCertificatesResponseInner{}
	return &this
}

// NewAttestationRootCertificatesResponseInnerWithDefaults instantiates a new AttestationRootCertificatesResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttestationRootCertificatesResponseInnerWithDefaults() *AttestationRootCertificatesResponseInner {
	this := AttestationRootCertificatesResponseInner{}
	return &this
}

// GetX5c returns the X5c field value if set, zero value otherwise.
func (o *AttestationRootCertificatesResponseInner) GetX5c() string {
	if o == nil || IsNil(o.X5c) {
		var ret string
		return ret
	}
	return *o.X5c
}

// GetX5cOk returns a tuple with the X5c field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttestationRootCertificatesResponseInner) GetX5cOk() (*string, bool) {
	if o == nil || IsNil(o.X5c) {
		return nil, false
	}
	return o.X5c, true
}

// HasX5c returns a boolean if a field has been set.
func (o *AttestationRootCertificatesResponseInner) HasX5c() bool {
	if o != nil && !IsNil(o.X5c) {
		return true
	}

	return false
}

// SetX5c gets a reference to the given string and assigns it to the X5c field.
func (o *AttestationRootCertificatesResponseInner) SetX5c(v string) {
	o.X5c = &v
}

// GetX5tS256 returns the X5tS256 field value if set, zero value otherwise.
func (o *AttestationRootCertificatesResponseInner) GetX5tS256() string {
	if o == nil || IsNil(o.X5tS256) {
		var ret string
		return ret
	}
	return *o.X5tS256
}

// GetX5tS256Ok returns a tuple with the X5tS256 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttestationRootCertificatesResponseInner) GetX5tS256Ok() (*string, bool) {
	if o == nil || IsNil(o.X5tS256) {
		return nil, false
	}
	return o.X5tS256, true
}

// HasX5tS256 returns a boolean if a field has been set.
func (o *AttestationRootCertificatesResponseInner) HasX5tS256() bool {
	if o != nil && !IsNil(o.X5tS256) {
		return true
	}

	return false
}

// SetX5tS256 gets a reference to the given string and assigns it to the X5tS256 field.
func (o *AttestationRootCertificatesResponseInner) SetX5tS256(v string) {
	o.X5tS256 = &v
}

// GetIss returns the Iss field value if set, zero value otherwise.
func (o *AttestationRootCertificatesResponseInner) GetIss() string {
	if o == nil || IsNil(o.Iss) {
		var ret string
		return ret
	}
	return *o.Iss
}

// GetIssOk returns a tuple with the Iss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttestationRootCertificatesResponseInner) GetIssOk() (*string, bool) {
	if o == nil || IsNil(o.Iss) {
		return nil, false
	}
	return o.Iss, true
}

// HasIss returns a boolean if a field has been set.
func (o *AttestationRootCertificatesResponseInner) HasIss() bool {
	if o != nil && !IsNil(o.Iss) {
		return true
	}

	return false
}

// SetIss gets a reference to the given string and assigns it to the Iss field.
func (o *AttestationRootCertificatesResponseInner) SetIss(v string) {
	o.Iss = &v
}

// GetExp returns the Exp field value if set, zero value otherwise.
func (o *AttestationRootCertificatesResponseInner) GetExp() string {
	if o == nil || IsNil(o.Exp) {
		var ret string
		return ret
	}
	return *o.Exp
}

// GetExpOk returns a tuple with the Exp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttestationRootCertificatesResponseInner) GetExpOk() (*string, bool) {
	if o == nil || IsNil(o.Exp) {
		return nil, false
	}
	return o.Exp, true
}

// HasExp returns a boolean if a field has been set.
func (o *AttestationRootCertificatesResponseInner) HasExp() bool {
	if o != nil && !IsNil(o.Exp) {
		return true
	}

	return false
}

// SetExp gets a reference to the given string and assigns it to the Exp field.
func (o *AttestationRootCertificatesResponseInner) SetExp(v string) {
	o.Exp = &v
}

func (o AttestationRootCertificatesResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AttestationRootCertificatesResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.X5c) {
		toSerialize["x5c"] = o.X5c
	}
	if !IsNil(o.X5tS256) {
		toSerialize["x5t#S256"] = o.X5tS256
	}
	if !IsNil(o.Iss) {
		toSerialize["iss"] = o.Iss
	}
	if !IsNil(o.Exp) {
		toSerialize["exp"] = o.Exp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AttestationRootCertificatesResponseInner) UnmarshalJSON(data []byte) (err error) {
	varAttestationRootCertificatesResponseInner := _AttestationRootCertificatesResponseInner{}

	err = json.Unmarshal(data, &varAttestationRootCertificatesResponseInner)

	if err != nil {
		return err
	}

	*o = AttestationRootCertificatesResponseInner(varAttestationRootCertificatesResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "x5c")
		delete(additionalProperties, "x5t#S256")
		delete(additionalProperties, "iss")
		delete(additionalProperties, "exp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAttestationRootCertificatesResponseInner struct {
	value *AttestationRootCertificatesResponseInner
	isSet bool
}

func (v NullableAttestationRootCertificatesResponseInner) Get() *AttestationRootCertificatesResponseInner {
	return v.value
}

func (v *NullableAttestationRootCertificatesResponseInner) Set(val *AttestationRootCertificatesResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAttestationRootCertificatesResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAttestationRootCertificatesResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttestationRootCertificatesResponseInner(val *AttestationRootCertificatesResponseInner) *NullableAttestationRootCertificatesResponseInner {
	return &NullableAttestationRootCertificatesResponseInner{value: val, isSet: true}
}

func (v NullableAttestationRootCertificatesResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttestationRootCertificatesResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
