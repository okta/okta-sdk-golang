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

// SchemasJsonWebKey A [JSON Web Key (JWK)](https://tools.ietf.org/html/rfc7517) is a JSON representation of a cryptographic key. Okta can use these keys to verify the signature of a JWT when provided for the `private_key_jwt` client authentication method or for a signed authorize request object. Okta supports both RSA and Elliptic Curve (EC) keys.
type SchemasJsonWebKey struct {
	Alg *string `json:"alg,omitempty"`
	// The unique identifier of the key
	Kid *string `json:"kid,omitempty"`
	// The type of public key
	Kty *string `json:"kty,omitempty"`
	// The status of the public key
	Status *string `json:"status,omitempty"`
	// The intended use of the public key
	Use *string `json:"use,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SchemasJsonWebKey SchemasJsonWebKey

// NewSchemasJsonWebKey instantiates a new SchemasJsonWebKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemasJsonWebKey() *SchemasJsonWebKey {
	this := SchemasJsonWebKey{}
	return &this
}

// NewSchemasJsonWebKeyWithDefaults instantiates a new SchemasJsonWebKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemasJsonWebKeyWithDefaults() *SchemasJsonWebKey {
	this := SchemasJsonWebKey{}
	return &this
}

// GetAlg returns the Alg field value if set, zero value otherwise.
func (o *SchemasJsonWebKey) GetAlg() string {
	if o == nil || o.Alg == nil {
		var ret string
		return ret
	}
	return *o.Alg
}

// GetAlgOk returns a tuple with the Alg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemasJsonWebKey) GetAlgOk() (*string, bool) {
	if o == nil || o.Alg == nil {
		return nil, false
	}
	return o.Alg, true
}

// HasAlg returns a boolean if a field has been set.
func (o *SchemasJsonWebKey) HasAlg() bool {
	if o != nil && o.Alg != nil {
		return true
	}

	return false
}

// SetAlg gets a reference to the given string and assigns it to the Alg field.
func (o *SchemasJsonWebKey) SetAlg(v string) {
	o.Alg = &v
}

// GetKid returns the Kid field value if set, zero value otherwise.
func (o *SchemasJsonWebKey) GetKid() string {
	if o == nil || o.Kid == nil {
		var ret string
		return ret
	}
	return *o.Kid
}

// GetKidOk returns a tuple with the Kid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemasJsonWebKey) GetKidOk() (*string, bool) {
	if o == nil || o.Kid == nil {
		return nil, false
	}
	return o.Kid, true
}

// HasKid returns a boolean if a field has been set.
func (o *SchemasJsonWebKey) HasKid() bool {
	if o != nil && o.Kid != nil {
		return true
	}

	return false
}

// SetKid gets a reference to the given string and assigns it to the Kid field.
func (o *SchemasJsonWebKey) SetKid(v string) {
	o.Kid = &v
}

// GetKty returns the Kty field value if set, zero value otherwise.
func (o *SchemasJsonWebKey) GetKty() string {
	if o == nil || o.Kty == nil {
		var ret string
		return ret
	}
	return *o.Kty
}

// GetKtyOk returns a tuple with the Kty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemasJsonWebKey) GetKtyOk() (*string, bool) {
	if o == nil || o.Kty == nil {
		return nil, false
	}
	return o.Kty, true
}

// HasKty returns a boolean if a field has been set.
func (o *SchemasJsonWebKey) HasKty() bool {
	if o != nil && o.Kty != nil {
		return true
	}

	return false
}

// SetKty gets a reference to the given string and assigns it to the Kty field.
func (o *SchemasJsonWebKey) SetKty(v string) {
	o.Kty = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SchemasJsonWebKey) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemasJsonWebKey) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SchemasJsonWebKey) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SchemasJsonWebKey) SetStatus(v string) {
	o.Status = &v
}

// GetUse returns the Use field value if set, zero value otherwise.
func (o *SchemasJsonWebKey) GetUse() string {
	if o == nil || o.Use == nil {
		var ret string
		return ret
	}
	return *o.Use
}

// GetUseOk returns a tuple with the Use field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemasJsonWebKey) GetUseOk() (*string, bool) {
	if o == nil || o.Use == nil {
		return nil, false
	}
	return o.Use, true
}

// HasUse returns a boolean if a field has been set.
func (o *SchemasJsonWebKey) HasUse() bool {
	if o != nil && o.Use != nil {
		return true
	}

	return false
}

// SetUse gets a reference to the given string and assigns it to the Use field.
func (o *SchemasJsonWebKey) SetUse(v string) {
	o.Use = &v
}

func (o SchemasJsonWebKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Alg != nil {
		toSerialize["alg"] = o.Alg
	}
	if o.Kid != nil {
		toSerialize["kid"] = o.Kid
	}
	if o.Kty != nil {
		toSerialize["kty"] = o.Kty
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Use != nil {
		toSerialize["use"] = o.Use
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SchemasJsonWebKey) UnmarshalJSON(bytes []byte) (err error) {
	varSchemasJsonWebKey := _SchemasJsonWebKey{}

	err = json.Unmarshal(bytes, &varSchemasJsonWebKey)
	if err == nil {
		*o = SchemasJsonWebKey(varSchemasJsonWebKey)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "alg")
		delete(additionalProperties, "kid")
		delete(additionalProperties, "kty")
		delete(additionalProperties, "status")
		delete(additionalProperties, "use")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableSchemasJsonWebKey struct {
	value *SchemasJsonWebKey
	isSet bool
}

func (v NullableSchemasJsonWebKey) Get() *SchemasJsonWebKey {
	return v.value
}

func (v *NullableSchemasJsonWebKey) Set(val *SchemasJsonWebKey) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemasJsonWebKey) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemasJsonWebKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemasJsonWebKey(val *SchemasJsonWebKey) *NullableSchemasJsonWebKey {
	return &NullableSchemasJsonWebKey{value: val, isSet: true}
}

func (v NullableSchemasJsonWebKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemasJsonWebKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

