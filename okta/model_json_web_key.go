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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"time"
)

// checks if the JsonWebKey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JsonWebKey{}

// JsonWebKey struct for JsonWebKey
type JsonWebKey struct {
	// Timestamp when the object was created
	Created *time.Time `json:"created,omitempty"`
	// RSA key value (public exponent) for Key binding
	E *string `json:"e,omitempty"`
	// Timestamp when the certificate expires
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
	// Unique identifier for the certificate
	Kid *string `json:"kid,omitempty"`
	// Cryptographic algorithm family for the certificate's keypair. Valid value: `RSA`
	Kty *string `json:"kty,omitempty"`
	// Timestamp when the object was last updated
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	// RSA modulus value that is used by both the public and private keys and provides a link between them
	N *string `json:"n,omitempty"`
	// Acceptable use of the certificate. Valid value: `sig`
	Use *string `json:"use,omitempty"`
	// X.509 certificate chain that contains a chain of one or more certificates
	X5c []string `json:"x5c,omitempty"`
	// X.509 certificate SHA-256 thumbprint, which is the base64url-encoded SHA-256 thumbprint (digest) of the DER encoding of an X.509 certificate
	X5tS256              *string `json:"x5t#S256,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _JsonWebKey JsonWebKey

// NewJsonWebKey instantiates a new JsonWebKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJsonWebKey() *JsonWebKey {
	this := JsonWebKey{}
	return &this
}

// NewJsonWebKeyWithDefaults instantiates a new JsonWebKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJsonWebKeyWithDefaults() *JsonWebKey {
	this := JsonWebKey{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *JsonWebKey) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonWebKey) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *JsonWebKey) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *JsonWebKey) SetCreated(v time.Time) {
	o.Created = &v
}

// GetE returns the E field value if set, zero value otherwise.
func (o *JsonWebKey) GetE() string {
	if o == nil || IsNil(o.E) {
		var ret string
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonWebKey) GetEOk() (*string, bool) {
	if o == nil || IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *JsonWebKey) HasE() bool {
	if o != nil && !IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *JsonWebKey) SetE(v string) {
	o.E = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *JsonWebKey) GetExpiresAt() time.Time {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonWebKey) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *JsonWebKey) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *JsonWebKey) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

// GetKid returns the Kid field value if set, zero value otherwise.
func (o *JsonWebKey) GetKid() string {
	if o == nil || IsNil(o.Kid) {
		var ret string
		return ret
	}
	return *o.Kid
}

// GetKidOk returns a tuple with the Kid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonWebKey) GetKidOk() (*string, bool) {
	if o == nil || IsNil(o.Kid) {
		return nil, false
	}
	return o.Kid, true
}

// HasKid returns a boolean if a field has been set.
func (o *JsonWebKey) HasKid() bool {
	if o != nil && !IsNil(o.Kid) {
		return true
	}

	return false
}

// SetKid gets a reference to the given string and assigns it to the Kid field.
func (o *JsonWebKey) SetKid(v string) {
	o.Kid = &v
}

// GetKty returns the Kty field value if set, zero value otherwise.
func (o *JsonWebKey) GetKty() string {
	if o == nil || IsNil(o.Kty) {
		var ret string
		return ret
	}
	return *o.Kty
}

// GetKtyOk returns a tuple with the Kty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonWebKey) GetKtyOk() (*string, bool) {
	if o == nil || IsNil(o.Kty) {
		return nil, false
	}
	return o.Kty, true
}

// HasKty returns a boolean if a field has been set.
func (o *JsonWebKey) HasKty() bool {
	if o != nil && !IsNil(o.Kty) {
		return true
	}

	return false
}

// SetKty gets a reference to the given string and assigns it to the Kty field.
func (o *JsonWebKey) SetKty(v string) {
	o.Kty = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *JsonWebKey) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonWebKey) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *JsonWebKey) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *JsonWebKey) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetN returns the N field value if set, zero value otherwise.
func (o *JsonWebKey) GetN() string {
	if o == nil || IsNil(o.N) {
		var ret string
		return ret
	}
	return *o.N
}

// GetNOk returns a tuple with the N field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonWebKey) GetNOk() (*string, bool) {
	if o == nil || IsNil(o.N) {
		return nil, false
	}
	return o.N, true
}

// HasN returns a boolean if a field has been set.
func (o *JsonWebKey) HasN() bool {
	if o != nil && !IsNil(o.N) {
		return true
	}

	return false
}

// SetN gets a reference to the given string and assigns it to the N field.
func (o *JsonWebKey) SetN(v string) {
	o.N = &v
}

// GetUse returns the Use field value if set, zero value otherwise.
func (o *JsonWebKey) GetUse() string {
	if o == nil || IsNil(o.Use) {
		var ret string
		return ret
	}
	return *o.Use
}

// GetUseOk returns a tuple with the Use field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonWebKey) GetUseOk() (*string, bool) {
	if o == nil || IsNil(o.Use) {
		return nil, false
	}
	return o.Use, true
}

// HasUse returns a boolean if a field has been set.
func (o *JsonWebKey) HasUse() bool {
	if o != nil && !IsNil(o.Use) {
		return true
	}

	return false
}

// SetUse gets a reference to the given string and assigns it to the Use field.
func (o *JsonWebKey) SetUse(v string) {
	o.Use = &v
}

// GetX5c returns the X5c field value if set, zero value otherwise.
func (o *JsonWebKey) GetX5c() []string {
	if o == nil || IsNil(o.X5c) {
		var ret []string
		return ret
	}
	return o.X5c
}

// GetX5cOk returns a tuple with the X5c field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonWebKey) GetX5cOk() ([]string, bool) {
	if o == nil || IsNil(o.X5c) {
		return nil, false
	}
	return o.X5c, true
}

// HasX5c returns a boolean if a field has been set.
func (o *JsonWebKey) HasX5c() bool {
	if o != nil && !IsNil(o.X5c) {
		return true
	}

	return false
}

// SetX5c gets a reference to the given []string and assigns it to the X5c field.
func (o *JsonWebKey) SetX5c(v []string) {
	o.X5c = v
}

// GetX5tS256 returns the X5tS256 field value if set, zero value otherwise.
func (o *JsonWebKey) GetX5tS256() string {
	if o == nil || IsNil(o.X5tS256) {
		var ret string
		return ret
	}
	return *o.X5tS256
}

// GetX5tS256Ok returns a tuple with the X5tS256 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonWebKey) GetX5tS256Ok() (*string, bool) {
	if o == nil || IsNil(o.X5tS256) {
		return nil, false
	}
	return o.X5tS256, true
}

// HasX5tS256 returns a boolean if a field has been set.
func (o *JsonWebKey) HasX5tS256() bool {
	if o != nil && !IsNil(o.X5tS256) {
		return true
	}

	return false
}

// SetX5tS256 gets a reference to the given string and assigns it to the X5tS256 field.
func (o *JsonWebKey) SetX5tS256(v string) {
	o.X5tS256 = &v
}

func (o JsonWebKey) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JsonWebKey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.E) {
		toSerialize["e"] = o.E
	}
	if !IsNil(o.ExpiresAt) {
		toSerialize["expiresAt"] = o.ExpiresAt
	}
	if !IsNil(o.Kid) {
		toSerialize["kid"] = o.Kid
	}
	if !IsNil(o.Kty) {
		toSerialize["kty"] = o.Kty
	}
	if !IsNil(o.LastUpdated) {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if !IsNil(o.N) {
		toSerialize["n"] = o.N
	}
	if !IsNil(o.Use) {
		toSerialize["use"] = o.Use
	}
	if !IsNil(o.X5c) {
		toSerialize["x5c"] = o.X5c
	}
	if !IsNil(o.X5tS256) {
		toSerialize["x5t#S256"] = o.X5tS256
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *JsonWebKey) UnmarshalJSON(data []byte) (err error) {
	varJsonWebKey := _JsonWebKey{}

	err = json.Unmarshal(data, &varJsonWebKey)

	if err != nil {
		return err
	}

	*o = JsonWebKey(varJsonWebKey)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
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
	}

	return err
}

type NullableJsonWebKey struct {
	value *JsonWebKey
	isSet bool
}

func (v NullableJsonWebKey) Get() *JsonWebKey {
	return v.value
}

func (v *NullableJsonWebKey) Set(val *JsonWebKey) {
	v.value = val
	v.isSet = true
}

func (v NullableJsonWebKey) IsSet() bool {
	return v.isSet
}

func (v *NullableJsonWebKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJsonWebKey(val *JsonWebKey) *NullableJsonWebKey {
	return &NullableJsonWebKey{value: val, isSet: true}
}

func (v NullableJsonWebKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJsonWebKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
