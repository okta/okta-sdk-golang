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
)

// Embedded The Public Key Details are defined in the `_embedded` property of the Key object.
type Embedded struct {
	// Algorithm used in the key
	Alg *string `json:"alg,omitempty"`
	// RSA key value (exponent) for key binding
	E *string `json:"e,omitempty"`
	// Unique identifier for the certificate
	Kid *string `json:"kid,omitempty"`
	// Cryptographic algorithm family for the certificate's keypair
	Kty *string `json:"kty,omitempty"`
	// RSA key value (modulus) for key binding
	N *string `json:"n,omitempty"`
	// Acceptable use of the certificate
	Use NullableString `json:"use,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Embedded Embedded

// NewEmbedded instantiates a new Embedded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmbedded() *Embedded {
	this := Embedded{}
	return &this
}

// NewEmbeddedWithDefaults instantiates a new Embedded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmbeddedWithDefaults() *Embedded {
	this := Embedded{}
	return &this
}

// GetAlg returns the Alg field value if set, zero value otherwise.
func (o *Embedded) GetAlg() string {
	if o == nil || o.Alg == nil {
		var ret string
		return ret
	}
	return *o.Alg
}

// GetAlgOk returns a tuple with the Alg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Embedded) GetAlgOk() (*string, bool) {
	if o == nil || o.Alg == nil {
		return nil, false
	}
	return o.Alg, true
}

// HasAlg returns a boolean if a field has been set.
func (o *Embedded) HasAlg() bool {
	if o != nil && o.Alg != nil {
		return true
	}

	return false
}

// SetAlg gets a reference to the given string and assigns it to the Alg field.
func (o *Embedded) SetAlg(v string) {
	o.Alg = &v
}

// GetE returns the E field value if set, zero value otherwise.
func (o *Embedded) GetE() string {
	if o == nil || o.E == nil {
		var ret string
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Embedded) GetEOk() (*string, bool) {
	if o == nil || o.E == nil {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *Embedded) HasE() bool {
	if o != nil && o.E != nil {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *Embedded) SetE(v string) {
	o.E = &v
}

// GetKid returns the Kid field value if set, zero value otherwise.
func (o *Embedded) GetKid() string {
	if o == nil || o.Kid == nil {
		var ret string
		return ret
	}
	return *o.Kid
}

// GetKidOk returns a tuple with the Kid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Embedded) GetKidOk() (*string, bool) {
	if o == nil || o.Kid == nil {
		return nil, false
	}
	return o.Kid, true
}

// HasKid returns a boolean if a field has been set.
func (o *Embedded) HasKid() bool {
	if o != nil && o.Kid != nil {
		return true
	}

	return false
}

// SetKid gets a reference to the given string and assigns it to the Kid field.
func (o *Embedded) SetKid(v string) {
	o.Kid = &v
}

// GetKty returns the Kty field value if set, zero value otherwise.
func (o *Embedded) GetKty() string {
	if o == nil || o.Kty == nil {
		var ret string
		return ret
	}
	return *o.Kty
}

// GetKtyOk returns a tuple with the Kty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Embedded) GetKtyOk() (*string, bool) {
	if o == nil || o.Kty == nil {
		return nil, false
	}
	return o.Kty, true
}

// HasKty returns a boolean if a field has been set.
func (o *Embedded) HasKty() bool {
	if o != nil && o.Kty != nil {
		return true
	}

	return false
}

// SetKty gets a reference to the given string and assigns it to the Kty field.
func (o *Embedded) SetKty(v string) {
	o.Kty = &v
}

// GetN returns the N field value if set, zero value otherwise.
func (o *Embedded) GetN() string {
	if o == nil || o.N == nil {
		var ret string
		return ret
	}
	return *o.N
}

// GetNOk returns a tuple with the N field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Embedded) GetNOk() (*string, bool) {
	if o == nil || o.N == nil {
		return nil, false
	}
	return o.N, true
}

// HasN returns a boolean if a field has been set.
func (o *Embedded) HasN() bool {
	if o != nil && o.N != nil {
		return true
	}

	return false
}

// SetN gets a reference to the given string and assigns it to the N field.
func (o *Embedded) SetN(v string) {
	o.N = &v
}

// GetUse returns the Use field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Embedded) GetUse() string {
	if o == nil || o.Use.Get() == nil {
		var ret string
		return ret
	}
	return *o.Use.Get()
}

// GetUseOk returns a tuple with the Use field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Embedded) GetUseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Use.Get(), o.Use.IsSet()
}

// HasUse returns a boolean if a field has been set.
func (o *Embedded) HasUse() bool {
	if o != nil && o.Use.IsSet() {
		return true
	}

	return false
}

// SetUse gets a reference to the given NullableString and assigns it to the Use field.
func (o *Embedded) SetUse(v string) {
	o.Use.Set(&v)
}
// SetUseNil sets the value for Use to be an explicit nil
func (o *Embedded) SetUseNil() {
	o.Use.Set(nil)
}

// UnsetUse ensures that no value is present for Use, not even an explicit nil
func (o *Embedded) UnsetUse() {
	o.Use.Unset()
}

func (o Embedded) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Alg != nil {
		toSerialize["alg"] = o.Alg
	}
	if o.E != nil {
		toSerialize["e"] = o.E
	}
	if o.Kid != nil {
		toSerialize["kid"] = o.Kid
	}
	if o.Kty != nil {
		toSerialize["kty"] = o.Kty
	}
	if o.N != nil {
		toSerialize["n"] = o.N
	}
	if o.Use.IsSet() {
		toSerialize["use"] = o.Use.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Embedded) UnmarshalJSON(bytes []byte) (err error) {
	varEmbedded := _Embedded{}

	err = json.Unmarshal(bytes, &varEmbedded)
	if err == nil {
		*o = Embedded(varEmbedded)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "alg")
		delete(additionalProperties, "e")
		delete(additionalProperties, "kid")
		delete(additionalProperties, "kty")
		delete(additionalProperties, "n")
		delete(additionalProperties, "use")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableEmbedded struct {
	value *Embedded
	isSet bool
}

func (v NullableEmbedded) Get() *Embedded {
	return v.value
}

func (v *NullableEmbedded) Set(val *Embedded) {
	v.value = val
	v.isSet = true
}

func (v NullableEmbedded) IsSet() bool {
	return v.isSet
}

func (v *NullableEmbedded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmbedded(val *Embedded) *NullableEmbedded {
	return &NullableEmbedded{value: val, isSet: true}
}

func (v NullableEmbedded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmbedded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

