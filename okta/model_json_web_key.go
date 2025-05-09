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
	"time"
)

// JsonWebKey struct for JsonWebKey
type JsonWebKey struct {
	// The algorithm used with the Key. Valid value: `RS256`
	Alg *string `json:"alg,omitempty"`
	// Timestamp when the object was created
	Created *time.Time `json:"created,omitempty"`
	// RSA key value (public exponent) for Key binding
	E *string `json:"e,omitempty"`
	// Timestamp when the certificate expires
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
	// Identifies the operation(s) for which the key is intended to be used
	KeyOps []string `json:"key_ops,omitempty"`
	// Unique identifier for the certificate
	Kid *string `json:"kid,omitempty"`
	// Cryptographic algorithm family for the certificate's keypair. Valid value: `RSA`
	Kty *string `json:"kty,omitempty"`
	// Timestamp when the object was last updated
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	// RSA modulus value that is used by both the public and private keys and provides a link between them
	N *string `json:"n,omitempty"`
	// An `ACTIVE` Key is used to sign tokens issued by the authorization server. Supported values: `ACTIVE`, `NEXT`, or `EXPIRED`<br> A `NEXT` Key is the next Key that the authorization server uses to sign tokens when Keys are rotated. The `NEXT` Key might not be listed if it hasn't been generated yet. An `EXPIRED` Key is the previous Key that the authorization server used to sign tokens. The `EXPIRED` Key might not be listed if no Key has expired or the expired Key was deleted.
	Status *string `json:"status,omitempty"`
	// Acceptable use of the certificate. Valid value: `sig`
	Use *string `json:"use,omitempty"`
	// X.509 certificate chain that contains a chain of one or more certificates
	X5c []string `json:"x5c,omitempty"`
	// X.509 certificate SHA-1 thumbprint, which is the base64url-encoded SHA-1 thumbprint (digest) of the DER encoding of an X.509 certificate
	X5t *string `json:"x5t,omitempty"`
	// X.509 certificate SHA-256 thumbprint, which is the base64url-encoded SHA-256 thumbprint (digest) of the DER encoding of an X.509 certificate
	X5tS256 *string `json:"x5t#S256,omitempty"`
	// A URI that refers to a resource for the X.509 public key certificate or certificate chain corresponding to the key used to digitally sign the JWS (JSON Web Signature)
	X5u *string `json:"x5u,omitempty"`
	Links *LinksSelf `json:"_links,omitempty"`
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

// GetAlg returns the Alg field value if set, zero value otherwise.
func (o *JsonWebKey) GetAlg() string {
	if o == nil || o.Alg == nil {
		var ret string
		return ret
	}
	return *o.Alg
}

// GetAlgOk returns a tuple with the Alg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonWebKey) GetAlgOk() (*string, bool) {
	if o == nil || o.Alg == nil {
		return nil, false
	}
	return o.Alg, true
}

// HasAlg returns a boolean if a field has been set.
func (o *JsonWebKey) HasAlg() bool {
	if o != nil && o.Alg != nil {
		return true
	}

	return false
}

// SetAlg gets a reference to the given string and assigns it to the Alg field.
func (o *JsonWebKey) SetAlg(v string) {
	o.Alg = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *JsonWebKey) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonWebKey) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *JsonWebKey) HasCreated() bool {
	if o != nil && o.Created != nil {
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
	if o == nil || o.E == nil {
		var ret string
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonWebKey) GetEOk() (*string, bool) {
	if o == nil || o.E == nil {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *JsonWebKey) HasE() bool {
	if o != nil && o.E != nil {
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
	if o == nil || o.ExpiresAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonWebKey) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || o.ExpiresAt == nil {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *JsonWebKey) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt != nil {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *JsonWebKey) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

// GetKeyOps returns the KeyOps field value if set, zero value otherwise.
func (o *JsonWebKey) GetKeyOps() []string {
	if o == nil || o.KeyOps == nil {
		var ret []string
		return ret
	}
	return o.KeyOps
}

// GetKeyOpsOk returns a tuple with the KeyOps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonWebKey) GetKeyOpsOk() ([]string, bool) {
	if o == nil || o.KeyOps == nil {
		return nil, false
	}
	return o.KeyOps, true
}

// HasKeyOps returns a boolean if a field has been set.
func (o *JsonWebKey) HasKeyOps() bool {
	if o != nil && o.KeyOps != nil {
		return true
	}

	return false
}

// SetKeyOps gets a reference to the given []string and assigns it to the KeyOps field.
func (o *JsonWebKey) SetKeyOps(v []string) {
	o.KeyOps = v
}

// GetKid returns the Kid field value if set, zero value otherwise.
func (o *JsonWebKey) GetKid() string {
	if o == nil || o.Kid == nil {
		var ret string
		return ret
	}
	return *o.Kid
}

// GetKidOk returns a tuple with the Kid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonWebKey) GetKidOk() (*string, bool) {
	if o == nil || o.Kid == nil {
		return nil, false
	}
	return o.Kid, true
}

// HasKid returns a boolean if a field has been set.
func (o *JsonWebKey) HasKid() bool {
	if o != nil && o.Kid != nil {
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
	if o == nil || o.Kty == nil {
		var ret string
		return ret
	}
	return *o.Kty
}

// GetKtyOk returns a tuple with the Kty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonWebKey) GetKtyOk() (*string, bool) {
	if o == nil || o.Kty == nil {
		return nil, false
	}
	return o.Kty, true
}

// HasKty returns a boolean if a field has been set.
func (o *JsonWebKey) HasKty() bool {
	if o != nil && o.Kty != nil {
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
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonWebKey) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *JsonWebKey) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
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
	if o == nil || o.N == nil {
		var ret string
		return ret
	}
	return *o.N
}

// GetNOk returns a tuple with the N field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonWebKey) GetNOk() (*string, bool) {
	if o == nil || o.N == nil {
		return nil, false
	}
	return o.N, true
}

// HasN returns a boolean if a field has been set.
func (o *JsonWebKey) HasN() bool {
	if o != nil && o.N != nil {
		return true
	}

	return false
}

// SetN gets a reference to the given string and assigns it to the N field.
func (o *JsonWebKey) SetN(v string) {
	o.N = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *JsonWebKey) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonWebKey) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *JsonWebKey) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *JsonWebKey) SetStatus(v string) {
	o.Status = &v
}

// GetUse returns the Use field value if set, zero value otherwise.
func (o *JsonWebKey) GetUse() string {
	if o == nil || o.Use == nil {
		var ret string
		return ret
	}
	return *o.Use
}

// GetUseOk returns a tuple with the Use field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonWebKey) GetUseOk() (*string, bool) {
	if o == nil || o.Use == nil {
		return nil, false
	}
	return o.Use, true
}

// HasUse returns a boolean if a field has been set.
func (o *JsonWebKey) HasUse() bool {
	if o != nil && o.Use != nil {
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
	if o == nil || o.X5c == nil {
		var ret []string
		return ret
	}
	return o.X5c
}

// GetX5cOk returns a tuple with the X5c field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonWebKey) GetX5cOk() ([]string, bool) {
	if o == nil || o.X5c == nil {
		return nil, false
	}
	return o.X5c, true
}

// HasX5c returns a boolean if a field has been set.
func (o *JsonWebKey) HasX5c() bool {
	if o != nil && o.X5c != nil {
		return true
	}

	return false
}

// SetX5c gets a reference to the given []string and assigns it to the X5c field.
func (o *JsonWebKey) SetX5c(v []string) {
	o.X5c = v
}

// GetX5t returns the X5t field value if set, zero value otherwise.
func (o *JsonWebKey) GetX5t() string {
	if o == nil || o.X5t == nil {
		var ret string
		return ret
	}
	return *o.X5t
}

// GetX5tOk returns a tuple with the X5t field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonWebKey) GetX5tOk() (*string, bool) {
	if o == nil || o.X5t == nil {
		return nil, false
	}
	return o.X5t, true
}

// HasX5t returns a boolean if a field has been set.
func (o *JsonWebKey) HasX5t() bool {
	if o != nil && o.X5t != nil {
		return true
	}

	return false
}

// SetX5t gets a reference to the given string and assigns it to the X5t field.
func (o *JsonWebKey) SetX5t(v string) {
	o.X5t = &v
}

// GetX5tS256 returns the X5tS256 field value if set, zero value otherwise.
func (o *JsonWebKey) GetX5tS256() string {
	if o == nil || o.X5tS256 == nil {
		var ret string
		return ret
	}
	return *o.X5tS256
}

// GetX5tS256Ok returns a tuple with the X5tS256 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonWebKey) GetX5tS256Ok() (*string, bool) {
	if o == nil || o.X5tS256 == nil {
		return nil, false
	}
	return o.X5tS256, true
}

// HasX5tS256 returns a boolean if a field has been set.
func (o *JsonWebKey) HasX5tS256() bool {
	if o != nil && o.X5tS256 != nil {
		return true
	}

	return false
}

// SetX5tS256 gets a reference to the given string and assigns it to the X5tS256 field.
func (o *JsonWebKey) SetX5tS256(v string) {
	o.X5tS256 = &v
}

// GetX5u returns the X5u field value if set, zero value otherwise.
func (o *JsonWebKey) GetX5u() string {
	if o == nil || o.X5u == nil {
		var ret string
		return ret
	}
	return *o.X5u
}

// GetX5uOk returns a tuple with the X5u field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonWebKey) GetX5uOk() (*string, bool) {
	if o == nil || o.X5u == nil {
		return nil, false
	}
	return o.X5u, true
}

// HasX5u returns a boolean if a field has been set.
func (o *JsonWebKey) HasX5u() bool {
	if o != nil && o.X5u != nil {
		return true
	}

	return false
}

// SetX5u gets a reference to the given string and assigns it to the X5u field.
func (o *JsonWebKey) SetX5u(v string) {
	o.X5u = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *JsonWebKey) GetLinks() LinksSelf {
	if o == nil || o.Links == nil {
		var ret LinksSelf
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonWebKey) GetLinksOk() (*LinksSelf, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *JsonWebKey) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksSelf and assigns it to the Links field.
func (o *JsonWebKey) SetLinks(v LinksSelf) {
	o.Links = &v
}

func (o JsonWebKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Alg != nil {
		toSerialize["alg"] = o.Alg
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.E != nil {
		toSerialize["e"] = o.E
	}
	if o.ExpiresAt != nil {
		toSerialize["expiresAt"] = o.ExpiresAt
	}
	if o.KeyOps != nil {
		toSerialize["key_ops"] = o.KeyOps
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
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Use != nil {
		toSerialize["use"] = o.Use
	}
	if o.X5c != nil {
		toSerialize["x5c"] = o.X5c
	}
	if o.X5t != nil {
		toSerialize["x5t"] = o.X5t
	}
	if o.X5tS256 != nil {
		toSerialize["x5t#S256"] = o.X5tS256
	}
	if o.X5u != nil {
		toSerialize["x5u"] = o.X5u
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *JsonWebKey) UnmarshalJSON(bytes []byte) (err error) {
	varJsonWebKey := _JsonWebKey{}

	err = json.Unmarshal(bytes, &varJsonWebKey)
	if err == nil {
		*o = JsonWebKey(varJsonWebKey)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "alg")
		delete(additionalProperties, "created")
		delete(additionalProperties, "e")
		delete(additionalProperties, "expiresAt")
		delete(additionalProperties, "key_ops")
		delete(additionalProperties, "kid")
		delete(additionalProperties, "kty")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "n")
		delete(additionalProperties, "status")
		delete(additionalProperties, "use")
		delete(additionalProperties, "x5c")
		delete(additionalProperties, "x5t")
		delete(additionalProperties, "x5t#S256")
		delete(additionalProperties, "x5u")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
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

