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

// checks if the Sso type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Sso{}

// Sso Supported SSO protocol configurations. You must configure at least one protocol: `oidc` or `saml`
type Sso struct {
	Oidc                 *Oidc `json:"oidc,omitempty"`
	Saml                 *Saml `json:"saml,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Sso Sso

// NewSso instantiates a new Sso object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSso() *Sso {
	this := Sso{}
	return &this
}

// NewSsoWithDefaults instantiates a new Sso object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSsoWithDefaults() *Sso {
	this := Sso{}
	return &this
}

// GetOidc returns the Oidc field value if set, zero value otherwise.
func (o *Sso) GetOidc() Oidc {
	if o == nil || IsNil(o.Oidc) {
		var ret Oidc
		return ret
	}
	return *o.Oidc
}

// GetOidcOk returns a tuple with the Oidc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sso) GetOidcOk() (*Oidc, bool) {
	if o == nil || IsNil(o.Oidc) {
		return nil, false
	}
	return o.Oidc, true
}

// HasOidc returns a boolean if a field has been set.
func (o *Sso) HasOidc() bool {
	if o != nil && !IsNil(o.Oidc) {
		return true
	}

	return false
}

// SetOidc gets a reference to the given Oidc and assigns it to the Oidc field.
func (o *Sso) SetOidc(v Oidc) {
	o.Oidc = &v
}

// GetSaml returns the Saml field value if set, zero value otherwise.
func (o *Sso) GetSaml() Saml {
	if o == nil || IsNil(o.Saml) {
		var ret Saml
		return ret
	}
	return *o.Saml
}

// GetSamlOk returns a tuple with the Saml field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sso) GetSamlOk() (*Saml, bool) {
	if o == nil || IsNil(o.Saml) {
		return nil, false
	}
	return o.Saml, true
}

// HasSaml returns a boolean if a field has been set.
func (o *Sso) HasSaml() bool {
	if o != nil && !IsNil(o.Saml) {
		return true
	}

	return false
}

// SetSaml gets a reference to the given Saml and assigns it to the Saml field.
func (o *Sso) SetSaml(v Saml) {
	o.Saml = &v
}

func (o Sso) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Sso) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Oidc) {
		toSerialize["oidc"] = o.Oidc
	}
	if !IsNil(o.Saml) {
		toSerialize["saml"] = o.Saml
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Sso) UnmarshalJSON(data []byte) (err error) {
	varSso := _Sso{}

	err = json.Unmarshal(data, &varSso)

	if err != nil {
		return err
	}

	*o = Sso(varSso)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "oidc")
		delete(additionalProperties, "saml")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSso struct {
	value *Sso
	isSet bool
}

func (v NullableSso) Get() *Sso {
	return v.value
}

func (v *NullableSso) Set(val *Sso) {
	v.value = val
	v.isSet = true
}

func (v NullableSso) IsSet() bool {
	return v.isSet
}

func (v *NullableSso) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSso(val *Sso) *NullableSso {
	return &NullableSso{value: val, isSet: true}
}

func (v NullableSso) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSso) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
