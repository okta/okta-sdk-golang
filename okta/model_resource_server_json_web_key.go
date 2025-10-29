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

// checks if the ResourceServerJsonWebKey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceServerJsonWebKey{}

// ResourceServerJsonWebKey A [JSON Web Key (JWK)](https://tools.ietf.org/html/rfc7517) is a JSON representation of a cryptographic key. Okta can use the active key to encrypt the access token minted by the authorization server. Okta supports only RSA keys with 'use: enc'.
type ResourceServerJsonWebKey struct {
	// The key exponent of a RSA key
	E *string `json:"e,omitempty"`
	// The unique identifier of the key
	Kid *string `json:"kid,omitempty"`
	// The type of public key
	Kty *string `json:"kty,omitempty"`
	// The modulus of the RSA key
	N *string `json:"n,omitempty"`
	// The status of the public key
	Status *string `json:"status,omitempty"`
	// The intended use of the public key
	Use                  *string `json:"use,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResourceServerJsonWebKey ResourceServerJsonWebKey

// NewResourceServerJsonWebKey instantiates a new ResourceServerJsonWebKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceServerJsonWebKey() *ResourceServerJsonWebKey {
	this := ResourceServerJsonWebKey{}
	return &this
}

// NewResourceServerJsonWebKeyWithDefaults instantiates a new ResourceServerJsonWebKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceServerJsonWebKeyWithDefaults() *ResourceServerJsonWebKey {
	this := ResourceServerJsonWebKey{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *ResourceServerJsonWebKey) GetE() string {
	if o == nil || IsNil(o.E) {
		var ret string
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceServerJsonWebKey) GetEOk() (*string, bool) {
	if o == nil || IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *ResourceServerJsonWebKey) HasE() bool {
	if o != nil && !IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *ResourceServerJsonWebKey) SetE(v string) {
	o.E = &v
}

// GetKid returns the Kid field value if set, zero value otherwise.
func (o *ResourceServerJsonWebKey) GetKid() string {
	if o == nil || IsNil(o.Kid) {
		var ret string
		return ret
	}
	return *o.Kid
}

// GetKidOk returns a tuple with the Kid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceServerJsonWebKey) GetKidOk() (*string, bool) {
	if o == nil || IsNil(o.Kid) {
		return nil, false
	}
	return o.Kid, true
}

// HasKid returns a boolean if a field has been set.
func (o *ResourceServerJsonWebKey) HasKid() bool {
	if o != nil && !IsNil(o.Kid) {
		return true
	}

	return false
}

// SetKid gets a reference to the given string and assigns it to the Kid field.
func (o *ResourceServerJsonWebKey) SetKid(v string) {
	o.Kid = &v
}

// GetKty returns the Kty field value if set, zero value otherwise.
func (o *ResourceServerJsonWebKey) GetKty() string {
	if o == nil || IsNil(o.Kty) {
		var ret string
		return ret
	}
	return *o.Kty
}

// GetKtyOk returns a tuple with the Kty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceServerJsonWebKey) GetKtyOk() (*string, bool) {
	if o == nil || IsNil(o.Kty) {
		return nil, false
	}
	return o.Kty, true
}

// HasKty returns a boolean if a field has been set.
func (o *ResourceServerJsonWebKey) HasKty() bool {
	if o != nil && !IsNil(o.Kty) {
		return true
	}

	return false
}

// SetKty gets a reference to the given string and assigns it to the Kty field.
func (o *ResourceServerJsonWebKey) SetKty(v string) {
	o.Kty = &v
}

// GetN returns the N field value if set, zero value otherwise.
func (o *ResourceServerJsonWebKey) GetN() string {
	if o == nil || IsNil(o.N) {
		var ret string
		return ret
	}
	return *o.N
}

// GetNOk returns a tuple with the N field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceServerJsonWebKey) GetNOk() (*string, bool) {
	if o == nil || IsNil(o.N) {
		return nil, false
	}
	return o.N, true
}

// HasN returns a boolean if a field has been set.
func (o *ResourceServerJsonWebKey) HasN() bool {
	if o != nil && !IsNil(o.N) {
		return true
	}

	return false
}

// SetN gets a reference to the given string and assigns it to the N field.
func (o *ResourceServerJsonWebKey) SetN(v string) {
	o.N = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ResourceServerJsonWebKey) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceServerJsonWebKey) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ResourceServerJsonWebKey) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ResourceServerJsonWebKey) SetStatus(v string) {
	o.Status = &v
}

// GetUse returns the Use field value if set, zero value otherwise.
func (o *ResourceServerJsonWebKey) GetUse() string {
	if o == nil || IsNil(o.Use) {
		var ret string
		return ret
	}
	return *o.Use
}

// GetUseOk returns a tuple with the Use field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceServerJsonWebKey) GetUseOk() (*string, bool) {
	if o == nil || IsNil(o.Use) {
		return nil, false
	}
	return o.Use, true
}

// HasUse returns a boolean if a field has been set.
func (o *ResourceServerJsonWebKey) HasUse() bool {
	if o != nil && !IsNil(o.Use) {
		return true
	}

	return false
}

// SetUse gets a reference to the given string and assigns it to the Use field.
func (o *ResourceServerJsonWebKey) SetUse(v string) {
	o.Use = &v
}

func (o ResourceServerJsonWebKey) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceServerJsonWebKey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.E) {
		toSerialize["e"] = o.E
	}
	if !IsNil(o.Kid) {
		toSerialize["kid"] = o.Kid
	}
	if !IsNil(o.Kty) {
		toSerialize["kty"] = o.Kty
	}
	if !IsNil(o.N) {
		toSerialize["n"] = o.N
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Use) {
		toSerialize["use"] = o.Use
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResourceServerJsonWebKey) UnmarshalJSON(data []byte) (err error) {
	varResourceServerJsonWebKey := _ResourceServerJsonWebKey{}

	err = json.Unmarshal(data, &varResourceServerJsonWebKey)

	if err != nil {
		return err
	}

	*o = ResourceServerJsonWebKey(varResourceServerJsonWebKey)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "e")
		delete(additionalProperties, "kid")
		delete(additionalProperties, "kty")
		delete(additionalProperties, "n")
		delete(additionalProperties, "status")
		delete(additionalProperties, "use")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResourceServerJsonWebKey struct {
	value *ResourceServerJsonWebKey
	isSet bool
}

func (v NullableResourceServerJsonWebKey) Get() *ResourceServerJsonWebKey {
	return v.value
}

func (v *NullableResourceServerJsonWebKey) Set(val *ResourceServerJsonWebKey) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceServerJsonWebKey) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceServerJsonWebKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceServerJsonWebKey(val *ResourceServerJsonWebKey) *NullableResourceServerJsonWebKey {
	return &NullableResourceServerJsonWebKey{value: val, isSet: true}
}

func (v NullableResourceServerJsonWebKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceServerJsonWebKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
