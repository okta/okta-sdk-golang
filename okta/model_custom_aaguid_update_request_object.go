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

// checks if the CustomAAGUIDUpdateRequestObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomAAGUIDUpdateRequestObject{}

// CustomAAGUIDUpdateRequestObject struct for CustomAAGUIDUpdateRequestObject
type CustomAAGUIDUpdateRequestObject struct {
	// Contains the certificate and information about it
	AttestationRootCertificates  []AttestationRootCertificatesRequestInner `json:"attestationRootCertificates,omitempty"`
	AuthenticatorCharacteristics *AAGUIDAuthenticatorCharacteristics       `json:"authenticatorCharacteristics,omitempty"`
	// The product name associated with this AAGUID.
	Name                 *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CustomAAGUIDUpdateRequestObject CustomAAGUIDUpdateRequestObject

// NewCustomAAGUIDUpdateRequestObject instantiates a new CustomAAGUIDUpdateRequestObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomAAGUIDUpdateRequestObject() *CustomAAGUIDUpdateRequestObject {
	this := CustomAAGUIDUpdateRequestObject{}
	return &this
}

// NewCustomAAGUIDUpdateRequestObjectWithDefaults instantiates a new CustomAAGUIDUpdateRequestObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomAAGUIDUpdateRequestObjectWithDefaults() *CustomAAGUIDUpdateRequestObject {
	this := CustomAAGUIDUpdateRequestObject{}
	return &this
}

// GetAttestationRootCertificates returns the AttestationRootCertificates field value if set, zero value otherwise.
func (o *CustomAAGUIDUpdateRequestObject) GetAttestationRootCertificates() []AttestationRootCertificatesRequestInner {
	if o == nil || IsNil(o.AttestationRootCertificates) {
		var ret []AttestationRootCertificatesRequestInner
		return ret
	}
	return o.AttestationRootCertificates
}

// GetAttestationRootCertificatesOk returns a tuple with the AttestationRootCertificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomAAGUIDUpdateRequestObject) GetAttestationRootCertificatesOk() ([]AttestationRootCertificatesRequestInner, bool) {
	if o == nil || IsNil(o.AttestationRootCertificates) {
		return nil, false
	}
	return o.AttestationRootCertificates, true
}

// HasAttestationRootCertificates returns a boolean if a field has been set.
func (o *CustomAAGUIDUpdateRequestObject) HasAttestationRootCertificates() bool {
	if o != nil && !IsNil(o.AttestationRootCertificates) {
		return true
	}

	return false
}

// SetAttestationRootCertificates gets a reference to the given []AttestationRootCertificatesRequestInner and assigns it to the AttestationRootCertificates field.
func (o *CustomAAGUIDUpdateRequestObject) SetAttestationRootCertificates(v []AttestationRootCertificatesRequestInner) {
	o.AttestationRootCertificates = v
}

// GetAuthenticatorCharacteristics returns the AuthenticatorCharacteristics field value if set, zero value otherwise.
func (o *CustomAAGUIDUpdateRequestObject) GetAuthenticatorCharacteristics() AAGUIDAuthenticatorCharacteristics {
	if o == nil || IsNil(o.AuthenticatorCharacteristics) {
		var ret AAGUIDAuthenticatorCharacteristics
		return ret
	}
	return *o.AuthenticatorCharacteristics
}

// GetAuthenticatorCharacteristicsOk returns a tuple with the AuthenticatorCharacteristics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomAAGUIDUpdateRequestObject) GetAuthenticatorCharacteristicsOk() (*AAGUIDAuthenticatorCharacteristics, bool) {
	if o == nil || IsNil(o.AuthenticatorCharacteristics) {
		return nil, false
	}
	return o.AuthenticatorCharacteristics, true
}

// HasAuthenticatorCharacteristics returns a boolean if a field has been set.
func (o *CustomAAGUIDUpdateRequestObject) HasAuthenticatorCharacteristics() bool {
	if o != nil && !IsNil(o.AuthenticatorCharacteristics) {
		return true
	}

	return false
}

// SetAuthenticatorCharacteristics gets a reference to the given AAGUIDAuthenticatorCharacteristics and assigns it to the AuthenticatorCharacteristics field.
func (o *CustomAAGUIDUpdateRequestObject) SetAuthenticatorCharacteristics(v AAGUIDAuthenticatorCharacteristics) {
	o.AuthenticatorCharacteristics = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CustomAAGUIDUpdateRequestObject) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomAAGUIDUpdateRequestObject) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CustomAAGUIDUpdateRequestObject) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CustomAAGUIDUpdateRequestObject) SetName(v string) {
	o.Name = &v
}

func (o CustomAAGUIDUpdateRequestObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomAAGUIDUpdateRequestObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AttestationRootCertificates) {
		toSerialize["attestationRootCertificates"] = o.AttestationRootCertificates
	}
	if !IsNil(o.AuthenticatorCharacteristics) {
		toSerialize["authenticatorCharacteristics"] = o.AuthenticatorCharacteristics
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CustomAAGUIDUpdateRequestObject) UnmarshalJSON(data []byte) (err error) {
	varCustomAAGUIDUpdateRequestObject := _CustomAAGUIDUpdateRequestObject{}

	err = json.Unmarshal(data, &varCustomAAGUIDUpdateRequestObject)

	if err != nil {
		return err
	}

	*o = CustomAAGUIDUpdateRequestObject(varCustomAAGUIDUpdateRequestObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "attestationRootCertificates")
		delete(additionalProperties, "authenticatorCharacteristics")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCustomAAGUIDUpdateRequestObject struct {
	value *CustomAAGUIDUpdateRequestObject
	isSet bool
}

func (v NullableCustomAAGUIDUpdateRequestObject) Get() *CustomAAGUIDUpdateRequestObject {
	return v.value
}

func (v *NullableCustomAAGUIDUpdateRequestObject) Set(val *CustomAAGUIDUpdateRequestObject) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomAAGUIDUpdateRequestObject) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomAAGUIDUpdateRequestObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomAAGUIDUpdateRequestObject(val *CustomAAGUIDUpdateRequestObject) *NullableCustomAAGUIDUpdateRequestObject {
	return &NullableCustomAAGUIDUpdateRequestObject{value: val, isSet: true}
}

func (v NullableCustomAAGUIDUpdateRequestObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomAAGUIDUpdateRequestObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
