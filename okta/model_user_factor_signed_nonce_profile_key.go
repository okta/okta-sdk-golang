/*
Okta Admin Management API

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

API version: 2025.08.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the UserFactorSignedNonceProfileKey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserFactorSignedNonceProfileKey{}

// UserFactorSignedNonceProfileKey JSON Web Key (JWK) for signed nonce verification
type UserFactorSignedNonceProfileKey struct {
	// EC curve name (present only for EC keys)
	Crv *string `json:"crv,omitempty"`
	// RSA public exponent (present only for RSA keys)
	E *string `json:"e,omitempty"`
	// Purpose of the key
	JwkType *string `json:"jwkType,omitempty"`
	// Key ID
	Kid *string `json:"kid,omitempty"`
	// Key type
	Kty *string `json:"kty,omitempty"`
	// RSA modulus (present only for RSA keys)
	N *string `json:"n,omitempty"`
	// Key usage
	Use *string `json:"use,omitempty"`
	// EC x-coordinate (present only for EC keys)
	X *string `json:"x,omitempty"`
	// X.509 certificate chain
	X5c []string `json:"x5c,omitempty"`
	// EC y-coordinate (present only for EC keys)
	Y                    *string `json:"y,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserFactorSignedNonceProfileKey UserFactorSignedNonceProfileKey

// NewUserFactorSignedNonceProfileKey instantiates a new UserFactorSignedNonceProfileKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserFactorSignedNonceProfileKey() *UserFactorSignedNonceProfileKey {
	this := UserFactorSignedNonceProfileKey{}
	return &this
}

// NewUserFactorSignedNonceProfileKeyWithDefaults instantiates a new UserFactorSignedNonceProfileKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserFactorSignedNonceProfileKeyWithDefaults() *UserFactorSignedNonceProfileKey {
	this := UserFactorSignedNonceProfileKey{}
	return &this
}

// GetCrv returns the Crv field value if set, zero value otherwise.
func (o *UserFactorSignedNonceProfileKey) GetCrv() string {
	if o == nil || IsNil(o.Crv) {
		var ret string
		return ret
	}
	return *o.Crv
}

// GetCrvOk returns a tuple with the Crv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorSignedNonceProfileKey) GetCrvOk() (*string, bool) {
	if o == nil || IsNil(o.Crv) {
		return nil, false
	}
	return o.Crv, true
}

// HasCrv returns a boolean if a field has been set.
func (o *UserFactorSignedNonceProfileKey) HasCrv() bool {
	if o != nil && !IsNil(o.Crv) {
		return true
	}

	return false
}

// SetCrv gets a reference to the given string and assigns it to the Crv field.
func (o *UserFactorSignedNonceProfileKey) SetCrv(v string) {
	o.Crv = &v
}

// GetE returns the E field value if set, zero value otherwise.
func (o *UserFactorSignedNonceProfileKey) GetE() string {
	if o == nil || IsNil(o.E) {
		var ret string
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorSignedNonceProfileKey) GetEOk() (*string, bool) {
	if o == nil || IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *UserFactorSignedNonceProfileKey) HasE() bool {
	if o != nil && !IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *UserFactorSignedNonceProfileKey) SetE(v string) {
	o.E = &v
}

// GetJwkType returns the JwkType field value if set, zero value otherwise.
func (o *UserFactorSignedNonceProfileKey) GetJwkType() string {
	if o == nil || IsNil(o.JwkType) {
		var ret string
		return ret
	}
	return *o.JwkType
}

// GetJwkTypeOk returns a tuple with the JwkType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorSignedNonceProfileKey) GetJwkTypeOk() (*string, bool) {
	if o == nil || IsNil(o.JwkType) {
		return nil, false
	}
	return o.JwkType, true
}

// HasJwkType returns a boolean if a field has been set.
func (o *UserFactorSignedNonceProfileKey) HasJwkType() bool {
	if o != nil && !IsNil(o.JwkType) {
		return true
	}

	return false
}

// SetJwkType gets a reference to the given string and assigns it to the JwkType field.
func (o *UserFactorSignedNonceProfileKey) SetJwkType(v string) {
	o.JwkType = &v
}

// GetKid returns the Kid field value if set, zero value otherwise.
func (o *UserFactorSignedNonceProfileKey) GetKid() string {
	if o == nil || IsNil(o.Kid) {
		var ret string
		return ret
	}
	return *o.Kid
}

// GetKidOk returns a tuple with the Kid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorSignedNonceProfileKey) GetKidOk() (*string, bool) {
	if o == nil || IsNil(o.Kid) {
		return nil, false
	}
	return o.Kid, true
}

// HasKid returns a boolean if a field has been set.
func (o *UserFactorSignedNonceProfileKey) HasKid() bool {
	if o != nil && !IsNil(o.Kid) {
		return true
	}

	return false
}

// SetKid gets a reference to the given string and assigns it to the Kid field.
func (o *UserFactorSignedNonceProfileKey) SetKid(v string) {
	o.Kid = &v
}

// GetKty returns the Kty field value if set, zero value otherwise.
func (o *UserFactorSignedNonceProfileKey) GetKty() string {
	if o == nil || IsNil(o.Kty) {
		var ret string
		return ret
	}
	return *o.Kty
}

// GetKtyOk returns a tuple with the Kty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorSignedNonceProfileKey) GetKtyOk() (*string, bool) {
	if o == nil || IsNil(o.Kty) {
		return nil, false
	}
	return o.Kty, true
}

// HasKty returns a boolean if a field has been set.
func (o *UserFactorSignedNonceProfileKey) HasKty() bool {
	if o != nil && !IsNil(o.Kty) {
		return true
	}

	return false
}

// SetKty gets a reference to the given string and assigns it to the Kty field.
func (o *UserFactorSignedNonceProfileKey) SetKty(v string) {
	o.Kty = &v
}

// GetN returns the N field value if set, zero value otherwise.
func (o *UserFactorSignedNonceProfileKey) GetN() string {
	if o == nil || IsNil(o.N) {
		var ret string
		return ret
	}
	return *o.N
}

// GetNOk returns a tuple with the N field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorSignedNonceProfileKey) GetNOk() (*string, bool) {
	if o == nil || IsNil(o.N) {
		return nil, false
	}
	return o.N, true
}

// HasN returns a boolean if a field has been set.
func (o *UserFactorSignedNonceProfileKey) HasN() bool {
	if o != nil && !IsNil(o.N) {
		return true
	}

	return false
}

// SetN gets a reference to the given string and assigns it to the N field.
func (o *UserFactorSignedNonceProfileKey) SetN(v string) {
	o.N = &v
}

// GetUse returns the Use field value if set, zero value otherwise.
func (o *UserFactorSignedNonceProfileKey) GetUse() string {
	if o == nil || IsNil(o.Use) {
		var ret string
		return ret
	}
	return *o.Use
}

// GetUseOk returns a tuple with the Use field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorSignedNonceProfileKey) GetUseOk() (*string, bool) {
	if o == nil || IsNil(o.Use) {
		return nil, false
	}
	return o.Use, true
}

// HasUse returns a boolean if a field has been set.
func (o *UserFactorSignedNonceProfileKey) HasUse() bool {
	if o != nil && !IsNil(o.Use) {
		return true
	}

	return false
}

// SetUse gets a reference to the given string and assigns it to the Use field.
func (o *UserFactorSignedNonceProfileKey) SetUse(v string) {
	o.Use = &v
}

// GetX returns the X field value if set, zero value otherwise.
func (o *UserFactorSignedNonceProfileKey) GetX() string {
	if o == nil || IsNil(o.X) {
		var ret string
		return ret
	}
	return *o.X
}

// GetXOk returns a tuple with the X field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorSignedNonceProfileKey) GetXOk() (*string, bool) {
	if o == nil || IsNil(o.X) {
		return nil, false
	}
	return o.X, true
}

// HasX returns a boolean if a field has been set.
func (o *UserFactorSignedNonceProfileKey) HasX() bool {
	if o != nil && !IsNil(o.X) {
		return true
	}

	return false
}

// SetX gets a reference to the given string and assigns it to the X field.
func (o *UserFactorSignedNonceProfileKey) SetX(v string) {
	o.X = &v
}

// GetX5c returns the X5c field value if set, zero value otherwise.
func (o *UserFactorSignedNonceProfileKey) GetX5c() []string {
	if o == nil || IsNil(o.X5c) {
		var ret []string
		return ret
	}
	return o.X5c
}

// GetX5cOk returns a tuple with the X5c field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorSignedNonceProfileKey) GetX5cOk() ([]string, bool) {
	if o == nil || IsNil(o.X5c) {
		return nil, false
	}
	return o.X5c, true
}

// HasX5c returns a boolean if a field has been set.
func (o *UserFactorSignedNonceProfileKey) HasX5c() bool {
	if o != nil && !IsNil(o.X5c) {
		return true
	}

	return false
}

// SetX5c gets a reference to the given []string and assigns it to the X5c field.
func (o *UserFactorSignedNonceProfileKey) SetX5c(v []string) {
	o.X5c = v
}

// GetY returns the Y field value if set, zero value otherwise.
func (o *UserFactorSignedNonceProfileKey) GetY() string {
	if o == nil || IsNil(o.Y) {
		var ret string
		return ret
	}
	return *o.Y
}

// GetYOk returns a tuple with the Y field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorSignedNonceProfileKey) GetYOk() (*string, bool) {
	if o == nil || IsNil(o.Y) {
		return nil, false
	}
	return o.Y, true
}

// HasY returns a boolean if a field has been set.
func (o *UserFactorSignedNonceProfileKey) HasY() bool {
	if o != nil && !IsNil(o.Y) {
		return true
	}

	return false
}

// SetY gets a reference to the given string and assigns it to the Y field.
func (o *UserFactorSignedNonceProfileKey) SetY(v string) {
	o.Y = &v
}

func (o UserFactorSignedNonceProfileKey) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserFactorSignedNonceProfileKey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Crv) {
		toSerialize["crv"] = o.Crv
	}
	if !IsNil(o.E) {
		toSerialize["e"] = o.E
	}
	if !IsNil(o.JwkType) {
		toSerialize["jwkType"] = o.JwkType
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
	if !IsNil(o.Use) {
		toSerialize["use"] = o.Use
	}
	if !IsNil(o.X) {
		toSerialize["x"] = o.X
	}
	if !IsNil(o.X5c) {
		toSerialize["x5c"] = o.X5c
	}
	if !IsNil(o.Y) {
		toSerialize["y"] = o.Y
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserFactorSignedNonceProfileKey) UnmarshalJSON(data []byte) (err error) {
	varUserFactorSignedNonceProfileKey := _UserFactorSignedNonceProfileKey{}

	err = json.Unmarshal(data, &varUserFactorSignedNonceProfileKey)

	if err != nil {
		return err
	}

	*o = UserFactorSignedNonceProfileKey(varUserFactorSignedNonceProfileKey)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "crv")
		delete(additionalProperties, "e")
		delete(additionalProperties, "jwkType")
		delete(additionalProperties, "kid")
		delete(additionalProperties, "kty")
		delete(additionalProperties, "n")
		delete(additionalProperties, "use")
		delete(additionalProperties, "x")
		delete(additionalProperties, "x5c")
		delete(additionalProperties, "y")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserFactorSignedNonceProfileKey struct {
	value *UserFactorSignedNonceProfileKey
	isSet bool
}

func (v NullableUserFactorSignedNonceProfileKey) Get() *UserFactorSignedNonceProfileKey {
	return v.value
}

func (v *NullableUserFactorSignedNonceProfileKey) Set(val *UserFactorSignedNonceProfileKey) {
	v.value = val
	v.isSet = true
}

func (v NullableUserFactorSignedNonceProfileKey) IsSet() bool {
	return v.isSet
}

func (v *NullableUserFactorSignedNonceProfileKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserFactorSignedNonceProfileKey(val *UserFactorSignedNonceProfileKey) *NullableUserFactorSignedNonceProfileKey {
	return &NullableUserFactorSignedNonceProfileKey{value: val, isSet: true}
}

func (v NullableUserFactorSignedNonceProfileKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserFactorSignedNonceProfileKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
