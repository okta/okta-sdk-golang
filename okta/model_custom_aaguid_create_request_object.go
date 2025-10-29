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

API version: 2025.08.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the CustomAAGUIDCreateRequestObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomAAGUIDCreateRequestObject{}

// CustomAAGUIDCreateRequestObject struct for CustomAAGUIDCreateRequestObject
type CustomAAGUIDCreateRequestObject struct {
	// An Authenticator Attestation Global Unique Identifier (AAGUID) is a 128-bit identifier indicating the model.
	Aaguid *string `json:"aaguid,omitempty"`
	// Contains the certificate and information about it
	AttestationRootCertificates  []AttestationRootCertificatesRequestInner `json:"attestationRootCertificates,omitempty"`
	AuthenticatorCharacteristics *AAGUIDAuthenticatorCharacteristics       `json:"authenticatorCharacteristics,omitempty"`
	AdditionalProperties         map[string]interface{}
}

type _CustomAAGUIDCreateRequestObject CustomAAGUIDCreateRequestObject

// NewCustomAAGUIDCreateRequestObject instantiates a new CustomAAGUIDCreateRequestObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomAAGUIDCreateRequestObject() *CustomAAGUIDCreateRequestObject {
	this := CustomAAGUIDCreateRequestObject{}
	return &this
}

// NewCustomAAGUIDCreateRequestObjectWithDefaults instantiates a new CustomAAGUIDCreateRequestObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomAAGUIDCreateRequestObjectWithDefaults() *CustomAAGUIDCreateRequestObject {
	this := CustomAAGUIDCreateRequestObject{}
	return &this
}

// GetAaguid returns the Aaguid field value if set, zero value otherwise.
func (o *CustomAAGUIDCreateRequestObject) GetAaguid() string {
	if o == nil || IsNil(o.Aaguid) {
		var ret string
		return ret
	}
	return *o.Aaguid
}

// GetAaguidOk returns a tuple with the Aaguid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomAAGUIDCreateRequestObject) GetAaguidOk() (*string, bool) {
	if o == nil || IsNil(o.Aaguid) {
		return nil, false
	}
	return o.Aaguid, true
}

// HasAaguid returns a boolean if a field has been set.
func (o *CustomAAGUIDCreateRequestObject) HasAaguid() bool {
	if o != nil && !IsNil(o.Aaguid) {
		return true
	}

	return false
}

// SetAaguid gets a reference to the given string and assigns it to the Aaguid field.
func (o *CustomAAGUIDCreateRequestObject) SetAaguid(v string) {
	o.Aaguid = &v
}

// GetAttestationRootCertificates returns the AttestationRootCertificates field value if set, zero value otherwise.
func (o *CustomAAGUIDCreateRequestObject) GetAttestationRootCertificates() []AttestationRootCertificatesRequestInner {
	if o == nil || IsNil(o.AttestationRootCertificates) {
		var ret []AttestationRootCertificatesRequestInner
		return ret
	}
	return o.AttestationRootCertificates
}

// GetAttestationRootCertificatesOk returns a tuple with the AttestationRootCertificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomAAGUIDCreateRequestObject) GetAttestationRootCertificatesOk() ([]AttestationRootCertificatesRequestInner, bool) {
	if o == nil || IsNil(o.AttestationRootCertificates) {
		return nil, false
	}
	return o.AttestationRootCertificates, true
}

// HasAttestationRootCertificates returns a boolean if a field has been set.
func (o *CustomAAGUIDCreateRequestObject) HasAttestationRootCertificates() bool {
	if o != nil && !IsNil(o.AttestationRootCertificates) {
		return true
	}

	return false
}

// SetAttestationRootCertificates gets a reference to the given []AttestationRootCertificatesRequestInner and assigns it to the AttestationRootCertificates field.
func (o *CustomAAGUIDCreateRequestObject) SetAttestationRootCertificates(v []AttestationRootCertificatesRequestInner) {
	o.AttestationRootCertificates = v
}

// GetAuthenticatorCharacteristics returns the AuthenticatorCharacteristics field value if set, zero value otherwise.
func (o *CustomAAGUIDCreateRequestObject) GetAuthenticatorCharacteristics() AAGUIDAuthenticatorCharacteristics {
	if o == nil || IsNil(o.AuthenticatorCharacteristics) {
		var ret AAGUIDAuthenticatorCharacteristics
		return ret
	}
	return *o.AuthenticatorCharacteristics
}

// GetAuthenticatorCharacteristicsOk returns a tuple with the AuthenticatorCharacteristics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomAAGUIDCreateRequestObject) GetAuthenticatorCharacteristicsOk() (*AAGUIDAuthenticatorCharacteristics, bool) {
	if o == nil || IsNil(o.AuthenticatorCharacteristics) {
		return nil, false
	}
	return o.AuthenticatorCharacteristics, true
}

// HasAuthenticatorCharacteristics returns a boolean if a field has been set.
func (o *CustomAAGUIDCreateRequestObject) HasAuthenticatorCharacteristics() bool {
	if o != nil && !IsNil(o.AuthenticatorCharacteristics) {
		return true
	}

	return false
}

// SetAuthenticatorCharacteristics gets a reference to the given AAGUIDAuthenticatorCharacteristics and assigns it to the AuthenticatorCharacteristics field.
func (o *CustomAAGUIDCreateRequestObject) SetAuthenticatorCharacteristics(v AAGUIDAuthenticatorCharacteristics) {
	o.AuthenticatorCharacteristics = &v
}

func (o CustomAAGUIDCreateRequestObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomAAGUIDCreateRequestObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Aaguid) {
		toSerialize["aaguid"] = o.Aaguid
	}
	if !IsNil(o.AttestationRootCertificates) {
		toSerialize["attestationRootCertificates"] = o.AttestationRootCertificates
	}
	if !IsNil(o.AuthenticatorCharacteristics) {
		toSerialize["authenticatorCharacteristics"] = o.AuthenticatorCharacteristics
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CustomAAGUIDCreateRequestObject) UnmarshalJSON(data []byte) (err error) {
	varCustomAAGUIDCreateRequestObject := _CustomAAGUIDCreateRequestObject{}

	err = json.Unmarshal(data, &varCustomAAGUIDCreateRequestObject)

	if err != nil {
		return err
	}

	*o = CustomAAGUIDCreateRequestObject(varCustomAAGUIDCreateRequestObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "aaguid")
		delete(additionalProperties, "attestationRootCertificates")
		delete(additionalProperties, "authenticatorCharacteristics")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCustomAAGUIDCreateRequestObject struct {
	value *CustomAAGUIDCreateRequestObject
	isSet bool
}

func (v NullableCustomAAGUIDCreateRequestObject) Get() *CustomAAGUIDCreateRequestObject {
	return v.value
}

func (v *NullableCustomAAGUIDCreateRequestObject) Set(val *CustomAAGUIDCreateRequestObject) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomAAGUIDCreateRequestObject) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomAAGUIDCreateRequestObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomAAGUIDCreateRequestObject(val *CustomAAGUIDCreateRequestObject) *NullableCustomAAGUIDCreateRequestObject {
	return &NullableCustomAAGUIDCreateRequestObject{value: val, isSet: true}
}

func (v NullableCustomAAGUIDCreateRequestObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomAAGUIDCreateRequestObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
