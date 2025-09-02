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
	"time"
)

// IdPKeyCredential A [JSON Web Key](https://tools.ietf.org/html/rfc7517) for a signature or encryption credential for an IdP
type IdPKeyCredential struct {
	// Timestamp when the object was created
	Created *time.Time `json:"created,omitempty"`
	// The exponent value for the RSA public key
	E *string `json:"e,omitempty"`
	// Timestamp when the object expires
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
	// Unique identifier for the key
	Kid *string `json:"kid,omitempty"`
	// Identifies the cryptographic algorithm family used with the key
	Kty *string `json:"kty,omitempty"`
	// Timestamp when the object was last updated
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	// The modulus value for the RSA public key
	N *string `json:"n,omitempty"`
	// Intended use of the public key
	Use *string `json:"use,omitempty"`
	// Base64-encoded X.509 certificate chain with DER encoding
	X5c []string `json:"x5c,omitempty"`
	// Base64url-encoded SHA-256 thumbprint of the DER encoding of an X.509 certificate
	X5tS256 *string `json:"x5t#S256,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdPKeyCredential IdPKeyCredential

// NewIdPKeyCredential instantiates a new IdPKeyCredential object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdPKeyCredential() *IdPKeyCredential {
	this := IdPKeyCredential{}
	return &this
}

// NewIdPKeyCredentialWithDefaults instantiates a new IdPKeyCredential object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdPKeyCredentialWithDefaults() *IdPKeyCredential {
	this := IdPKeyCredential{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *IdPKeyCredential) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdPKeyCredential) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *IdPKeyCredential) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *IdPKeyCredential) SetCreated(v time.Time) {
	o.Created = &v
}

// GetE returns the E field value if set, zero value otherwise.
func (o *IdPKeyCredential) GetE() string {
	if o == nil || o.E == nil {
		var ret string
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdPKeyCredential) GetEOk() (*string, bool) {
	if o == nil || o.E == nil {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *IdPKeyCredential) HasE() bool {
	if o != nil && o.E != nil {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *IdPKeyCredential) SetE(v string) {
	o.E = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *IdPKeyCredential) GetExpiresAt() time.Time {
	if o == nil || o.ExpiresAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdPKeyCredential) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || o.ExpiresAt == nil {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *IdPKeyCredential) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt != nil {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *IdPKeyCredential) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

// GetKid returns the Kid field value if set, zero value otherwise.
func (o *IdPKeyCredential) GetKid() string {
	if o == nil || o.Kid == nil {
		var ret string
		return ret
	}
	return *o.Kid
}

// GetKidOk returns a tuple with the Kid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdPKeyCredential) GetKidOk() (*string, bool) {
	if o == nil || o.Kid == nil {
		return nil, false
	}
	return o.Kid, true
}

// HasKid returns a boolean if a field has been set.
func (o *IdPKeyCredential) HasKid() bool {
	if o != nil && o.Kid != nil {
		return true
	}

	return false
}

// SetKid gets a reference to the given string and assigns it to the Kid field.
func (o *IdPKeyCredential) SetKid(v string) {
	o.Kid = &v
}

// GetKty returns the Kty field value if set, zero value otherwise.
func (o *IdPKeyCredential) GetKty() string {
	if o == nil || o.Kty == nil {
		var ret string
		return ret
	}
	return *o.Kty
}

// GetKtyOk returns a tuple with the Kty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdPKeyCredential) GetKtyOk() (*string, bool) {
	if o == nil || o.Kty == nil {
		return nil, false
	}
	return o.Kty, true
}

// HasKty returns a boolean if a field has been set.
func (o *IdPKeyCredential) HasKty() bool {
	if o != nil && o.Kty != nil {
		return true
	}

	return false
}

// SetKty gets a reference to the given string and assigns it to the Kty field.
func (o *IdPKeyCredential) SetKty(v string) {
	o.Kty = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *IdPKeyCredential) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdPKeyCredential) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *IdPKeyCredential) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *IdPKeyCredential) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetN returns the N field value if set, zero value otherwise.
func (o *IdPKeyCredential) GetN() string {
	if o == nil || o.N == nil {
		var ret string
		return ret
	}
	return *o.N
}

// GetNOk returns a tuple with the N field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdPKeyCredential) GetNOk() (*string, bool) {
	if o == nil || o.N == nil {
		return nil, false
	}
	return o.N, true
}

// HasN returns a boolean if a field has been set.
func (o *IdPKeyCredential) HasN() bool {
	if o != nil && o.N != nil {
		return true
	}

	return false
}

// SetN gets a reference to the given string and assigns it to the N field.
func (o *IdPKeyCredential) SetN(v string) {
	o.N = &v
}

// GetUse returns the Use field value if set, zero value otherwise.
func (o *IdPKeyCredential) GetUse() string {
	if o == nil || o.Use == nil {
		var ret string
		return ret
	}
	return *o.Use
}

// GetUseOk returns a tuple with the Use field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdPKeyCredential) GetUseOk() (*string, bool) {
	if o == nil || o.Use == nil {
		return nil, false
	}
	return o.Use, true
}

// HasUse returns a boolean if a field has been set.
func (o *IdPKeyCredential) HasUse() bool {
	if o != nil && o.Use != nil {
		return true
	}

	return false
}

// SetUse gets a reference to the given string and assigns it to the Use field.
func (o *IdPKeyCredential) SetUse(v string) {
	o.Use = &v
}

// GetX5c returns the X5c field value if set, zero value otherwise.
func (o *IdPKeyCredential) GetX5c() []string {
	if o == nil || o.X5c == nil {
		var ret []string
		return ret
	}
	return o.X5c
}

// GetX5cOk returns a tuple with the X5c field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdPKeyCredential) GetX5cOk() ([]string, bool) {
	if o == nil || o.X5c == nil {
		return nil, false
	}
	return o.X5c, true
}

// HasX5c returns a boolean if a field has been set.
func (o *IdPKeyCredential) HasX5c() bool {
	if o != nil && o.X5c != nil {
		return true
	}

	return false
}

// SetX5c gets a reference to the given []string and assigns it to the X5c field.
func (o *IdPKeyCredential) SetX5c(v []string) {
	o.X5c = v
}

// GetX5tS256 returns the X5tS256 field value if set, zero value otherwise.
func (o *IdPKeyCredential) GetX5tS256() string {
	if o == nil || o.X5tS256 == nil {
		var ret string
		return ret
	}
	return *o.X5tS256
}

// GetX5tS256Ok returns a tuple with the X5tS256 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdPKeyCredential) GetX5tS256Ok() (*string, bool) {
	if o == nil || o.X5tS256 == nil {
		return nil, false
	}
	return o.X5tS256, true
}

// HasX5tS256 returns a boolean if a field has been set.
func (o *IdPKeyCredential) HasX5tS256() bool {
	if o != nil && o.X5tS256 != nil {
		return true
	}

	return false
}

// SetX5tS256 gets a reference to the given string and assigns it to the X5tS256 field.
func (o *IdPKeyCredential) SetX5tS256(v string) {
	o.X5tS256 = &v
}

func (o IdPKeyCredential) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.E != nil {
		toSerialize["e"] = o.E
	}
	if o.ExpiresAt != nil {
		toSerialize["expiresAt"] = o.ExpiresAt
	}
	if o.Kid != nil {
		toSerialize["kid"] = o.Kid
	}
	if o.Kty != nil {
		toSerialize["kty"] = o.Kty
	}
	if o.LastUpdated != nil {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if o.N != nil {
		toSerialize["n"] = o.N
	}
	if o.Use != nil {
		toSerialize["use"] = o.Use
	}
	if o.X5c != nil {
		toSerialize["x5c"] = o.X5c
	}
	if o.X5tS256 != nil {
		toSerialize["x5t#S256"] = o.X5tS256
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IdPKeyCredential) UnmarshalJSON(bytes []byte) (err error) {
	varIdPKeyCredential := _IdPKeyCredential{}

	err = json.Unmarshal(bytes, &varIdPKeyCredential)
	if err == nil {
		*o = IdPKeyCredential(varIdPKeyCredential)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "created")
		delete(additionalProperties, "e")
		delete(additionalProperties, "expiresAt")
		delete(additionalProperties, "kid")
		delete(additionalProperties, "kty")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "n")
		delete(additionalProperties, "use")
		delete(additionalProperties, "x5c")
		delete(additionalProperties, "x5t#S256")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableIdPKeyCredential struct {
	value *IdPKeyCredential
	isSet bool
}

func (v NullableIdPKeyCredential) Get() *IdPKeyCredential {
	return v.value
}

func (v *NullableIdPKeyCredential) Set(val *IdPKeyCredential) {
	v.value = val
	v.isSet = true
}

func (v NullableIdPKeyCredential) IsSet() bool {
	return v.isSet
}

func (v *NullableIdPKeyCredential) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdPKeyCredential(val *IdPKeyCredential) *NullableIdPKeyCredential {
	return &NullableIdPKeyCredential{value: val, isSet: true}
}

func (v NullableIdPKeyCredential) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdPKeyCredential) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

