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

// checks if the AuthorizationServerJsonWebKey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthorizationServerJsonWebKey{}

// AuthorizationServerJsonWebKey struct for AuthorizationServerJsonWebKey
type AuthorizationServerJsonWebKey struct {
	// The algorithm used with the Key. Valid value: `RS256`
	Alg *string `json:"alg,omitempty"`
	// RSA key value (public exponent) for Key binding
	E *string `json:"e,omitempty"`
	// Unique identifier for the key
	Kid *string `json:"kid,omitempty"`
	// Cryptographic algorithm family for the certificate's keypair. Valid value: `RSA`
	Kty *string `json:"kty,omitempty"`
	// RSA modulus value that is used by both the public and private keys and provides a link between them
	N *string `json:"n,omitempty"`
	// An `ACTIVE` Key is used to sign tokens issued by the authorization server. Supported values: `ACTIVE`, `NEXT`, or `EXPIRED`<br> A `NEXT` Key is the next Key that the authorization server uses to sign tokens when Keys are rotated. The `NEXT` Key might not be listed if it hasn't been generated. An `EXPIRED` Key is the previous Key that the authorization server used to sign tokens. The `EXPIRED` Key might not be listed if no Key has expired or the expired Key was deleted.
	Status *string `json:"status,omitempty"`
	// Acceptable use of the key. Valid value: `sig`
	Use                  *string    `json:"use,omitempty"`
	Links                *LinksSelf `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthorizationServerJsonWebKey AuthorizationServerJsonWebKey

// NewAuthorizationServerJsonWebKey instantiates a new AuthorizationServerJsonWebKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizationServerJsonWebKey() *AuthorizationServerJsonWebKey {
	this := AuthorizationServerJsonWebKey{}
	return &this
}

// NewAuthorizationServerJsonWebKeyWithDefaults instantiates a new AuthorizationServerJsonWebKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizationServerJsonWebKeyWithDefaults() *AuthorizationServerJsonWebKey {
	this := AuthorizationServerJsonWebKey{}
	return &this
}

// GetAlg returns the Alg field value if set, zero value otherwise.
func (o *AuthorizationServerJsonWebKey) GetAlg() string {
	if o == nil || IsNil(o.Alg) {
		var ret string
		return ret
	}
	return *o.Alg
}

// GetAlgOk returns a tuple with the Alg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServerJsonWebKey) GetAlgOk() (*string, bool) {
	if o == nil || IsNil(o.Alg) {
		return nil, false
	}
	return o.Alg, true
}

// HasAlg returns a boolean if a field has been set.
func (o *AuthorizationServerJsonWebKey) HasAlg() bool {
	if o != nil && !IsNil(o.Alg) {
		return true
	}

	return false
}

// SetAlg gets a reference to the given string and assigns it to the Alg field.
func (o *AuthorizationServerJsonWebKey) SetAlg(v string) {
	o.Alg = &v
}

// GetE returns the E field value if set, zero value otherwise.
func (o *AuthorizationServerJsonWebKey) GetE() string {
	if o == nil || IsNil(o.E) {
		var ret string
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServerJsonWebKey) GetEOk() (*string, bool) {
	if o == nil || IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *AuthorizationServerJsonWebKey) HasE() bool {
	if o != nil && !IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *AuthorizationServerJsonWebKey) SetE(v string) {
	o.E = &v
}

// GetKid returns the Kid field value if set, zero value otherwise.
func (o *AuthorizationServerJsonWebKey) GetKid() string {
	if o == nil || IsNil(o.Kid) {
		var ret string
		return ret
	}
	return *o.Kid
}

// GetKidOk returns a tuple with the Kid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServerJsonWebKey) GetKidOk() (*string, bool) {
	if o == nil || IsNil(o.Kid) {
		return nil, false
	}
	return o.Kid, true
}

// HasKid returns a boolean if a field has been set.
func (o *AuthorizationServerJsonWebKey) HasKid() bool {
	if o != nil && !IsNil(o.Kid) {
		return true
	}

	return false
}

// SetKid gets a reference to the given string and assigns it to the Kid field.
func (o *AuthorizationServerJsonWebKey) SetKid(v string) {
	o.Kid = &v
}

// GetKty returns the Kty field value if set, zero value otherwise.
func (o *AuthorizationServerJsonWebKey) GetKty() string {
	if o == nil || IsNil(o.Kty) {
		var ret string
		return ret
	}
	return *o.Kty
}

// GetKtyOk returns a tuple with the Kty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServerJsonWebKey) GetKtyOk() (*string, bool) {
	if o == nil || IsNil(o.Kty) {
		return nil, false
	}
	return o.Kty, true
}

// HasKty returns a boolean if a field has been set.
func (o *AuthorizationServerJsonWebKey) HasKty() bool {
	if o != nil && !IsNil(o.Kty) {
		return true
	}

	return false
}

// SetKty gets a reference to the given string and assigns it to the Kty field.
func (o *AuthorizationServerJsonWebKey) SetKty(v string) {
	o.Kty = &v
}

// GetN returns the N field value if set, zero value otherwise.
func (o *AuthorizationServerJsonWebKey) GetN() string {
	if o == nil || IsNil(o.N) {
		var ret string
		return ret
	}
	return *o.N
}

// GetNOk returns a tuple with the N field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServerJsonWebKey) GetNOk() (*string, bool) {
	if o == nil || IsNil(o.N) {
		return nil, false
	}
	return o.N, true
}

// HasN returns a boolean if a field has been set.
func (o *AuthorizationServerJsonWebKey) HasN() bool {
	if o != nil && !IsNil(o.N) {
		return true
	}

	return false
}

// SetN gets a reference to the given string and assigns it to the N field.
func (o *AuthorizationServerJsonWebKey) SetN(v string) {
	o.N = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AuthorizationServerJsonWebKey) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServerJsonWebKey) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AuthorizationServerJsonWebKey) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AuthorizationServerJsonWebKey) SetStatus(v string) {
	o.Status = &v
}

// GetUse returns the Use field value if set, zero value otherwise.
func (o *AuthorizationServerJsonWebKey) GetUse() string {
	if o == nil || IsNil(o.Use) {
		var ret string
		return ret
	}
	return *o.Use
}

// GetUseOk returns a tuple with the Use field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServerJsonWebKey) GetUseOk() (*string, bool) {
	if o == nil || IsNil(o.Use) {
		return nil, false
	}
	return o.Use, true
}

// HasUse returns a boolean if a field has been set.
func (o *AuthorizationServerJsonWebKey) HasUse() bool {
	if o != nil && !IsNil(o.Use) {
		return true
	}

	return false
}

// SetUse gets a reference to the given string and assigns it to the Use field.
func (o *AuthorizationServerJsonWebKey) SetUse(v string) {
	o.Use = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AuthorizationServerJsonWebKey) GetLinks() LinksSelf {
	if o == nil || IsNil(o.Links) {
		var ret LinksSelf
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServerJsonWebKey) GetLinksOk() (*LinksSelf, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AuthorizationServerJsonWebKey) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksSelf and assigns it to the Links field.
func (o *AuthorizationServerJsonWebKey) SetLinks(v LinksSelf) {
	o.Links = &v
}

func (o AuthorizationServerJsonWebKey) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthorizationServerJsonWebKey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Alg) {
		toSerialize["alg"] = o.Alg
	}
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
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AuthorizationServerJsonWebKey) UnmarshalJSON(data []byte) (err error) {
	varAuthorizationServerJsonWebKey := _AuthorizationServerJsonWebKey{}

	err = json.Unmarshal(data, &varAuthorizationServerJsonWebKey)

	if err != nil {
		return err
	}

	*o = AuthorizationServerJsonWebKey(varAuthorizationServerJsonWebKey)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "alg")
		delete(additionalProperties, "e")
		delete(additionalProperties, "kid")
		delete(additionalProperties, "kty")
		delete(additionalProperties, "n")
		delete(additionalProperties, "status")
		delete(additionalProperties, "use")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAuthorizationServerJsonWebKey struct {
	value *AuthorizationServerJsonWebKey
	isSet bool
}

func (v NullableAuthorizationServerJsonWebKey) Get() *AuthorizationServerJsonWebKey {
	return v.value
}

func (v *NullableAuthorizationServerJsonWebKey) Set(val *AuthorizationServerJsonWebKey) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizationServerJsonWebKey) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizationServerJsonWebKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizationServerJsonWebKey(val *AuthorizationServerJsonWebKey) *NullableAuthorizationServerJsonWebKey {
	return &NullableAuthorizationServerJsonWebKey{value: val, isSet: true}
}

func (v NullableAuthorizationServerJsonWebKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizationServerJsonWebKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
