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

// checks if the CustomAAGUIDResponseObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomAAGUIDResponseObject{}

// CustomAAGUIDResponseObject struct for CustomAAGUIDResponseObject
type CustomAAGUIDResponseObject struct {
	// A unique 128-bit identifier that's assigned to a specific model of security key or authenticator
	Aaguid                       *string                                    `json:"aaguid,omitempty"`
	AttestationRootCertificates  []AttestationRootCertificatesResponseInner `json:"attestationRootCertificates,omitempty"`
	AuthenticatorCharacteristics *AAGUIDAuthenticatorCharacteristics        `json:"authenticatorCharacteristics,omitempty"`
	// The product name associated with the AAGUID
	Name                 *string    `json:"name,omitempty"`
	Links                *LinksSelf `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CustomAAGUIDResponseObject CustomAAGUIDResponseObject

// NewCustomAAGUIDResponseObject instantiates a new CustomAAGUIDResponseObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomAAGUIDResponseObject() *CustomAAGUIDResponseObject {
	this := CustomAAGUIDResponseObject{}
	return &this
}

// NewCustomAAGUIDResponseObjectWithDefaults instantiates a new CustomAAGUIDResponseObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomAAGUIDResponseObjectWithDefaults() *CustomAAGUIDResponseObject {
	this := CustomAAGUIDResponseObject{}
	return &this
}

// GetAaguid returns the Aaguid field value if set, zero value otherwise.
func (o *CustomAAGUIDResponseObject) GetAaguid() string {
	if o == nil || IsNil(o.Aaguid) {
		var ret string
		return ret
	}
	return *o.Aaguid
}

// GetAaguidOk returns a tuple with the Aaguid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomAAGUIDResponseObject) GetAaguidOk() (*string, bool) {
	if o == nil || IsNil(o.Aaguid) {
		return nil, false
	}
	return o.Aaguid, true
}

// HasAaguid returns a boolean if a field has been set.
func (o *CustomAAGUIDResponseObject) HasAaguid() bool {
	if o != nil && !IsNil(o.Aaguid) {
		return true
	}

	return false
}

// SetAaguid gets a reference to the given string and assigns it to the Aaguid field.
func (o *CustomAAGUIDResponseObject) SetAaguid(v string) {
	o.Aaguid = &v
}

// GetAttestationRootCertificates returns the AttestationRootCertificates field value if set, zero value otherwise.
func (o *CustomAAGUIDResponseObject) GetAttestationRootCertificates() []AttestationRootCertificatesResponseInner {
	if o == nil || IsNil(o.AttestationRootCertificates) {
		var ret []AttestationRootCertificatesResponseInner
		return ret
	}
	return o.AttestationRootCertificates
}

// GetAttestationRootCertificatesOk returns a tuple with the AttestationRootCertificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomAAGUIDResponseObject) GetAttestationRootCertificatesOk() ([]AttestationRootCertificatesResponseInner, bool) {
	if o == nil || IsNil(o.AttestationRootCertificates) {
		return nil, false
	}
	return o.AttestationRootCertificates, true
}

// HasAttestationRootCertificates returns a boolean if a field has been set.
func (o *CustomAAGUIDResponseObject) HasAttestationRootCertificates() bool {
	if o != nil && !IsNil(o.AttestationRootCertificates) {
		return true
	}

	return false
}

// SetAttestationRootCertificates gets a reference to the given []AttestationRootCertificatesResponseInner and assigns it to the AttestationRootCertificates field.
func (o *CustomAAGUIDResponseObject) SetAttestationRootCertificates(v []AttestationRootCertificatesResponseInner) {
	o.AttestationRootCertificates = v
}

// GetAuthenticatorCharacteristics returns the AuthenticatorCharacteristics field value if set, zero value otherwise.
func (o *CustomAAGUIDResponseObject) GetAuthenticatorCharacteristics() AAGUIDAuthenticatorCharacteristics {
	if o == nil || IsNil(o.AuthenticatorCharacteristics) {
		var ret AAGUIDAuthenticatorCharacteristics
		return ret
	}
	return *o.AuthenticatorCharacteristics
}

// GetAuthenticatorCharacteristicsOk returns a tuple with the AuthenticatorCharacteristics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomAAGUIDResponseObject) GetAuthenticatorCharacteristicsOk() (*AAGUIDAuthenticatorCharacteristics, bool) {
	if o == nil || IsNil(o.AuthenticatorCharacteristics) {
		return nil, false
	}
	return o.AuthenticatorCharacteristics, true
}

// HasAuthenticatorCharacteristics returns a boolean if a field has been set.
func (o *CustomAAGUIDResponseObject) HasAuthenticatorCharacteristics() bool {
	if o != nil && !IsNil(o.AuthenticatorCharacteristics) {
		return true
	}

	return false
}

// SetAuthenticatorCharacteristics gets a reference to the given AAGUIDAuthenticatorCharacteristics and assigns it to the AuthenticatorCharacteristics field.
func (o *CustomAAGUIDResponseObject) SetAuthenticatorCharacteristics(v AAGUIDAuthenticatorCharacteristics) {
	o.AuthenticatorCharacteristics = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CustomAAGUIDResponseObject) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomAAGUIDResponseObject) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CustomAAGUIDResponseObject) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CustomAAGUIDResponseObject) SetName(v string) {
	o.Name = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CustomAAGUIDResponseObject) GetLinks() LinksSelf {
	if o == nil || IsNil(o.Links) {
		var ret LinksSelf
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomAAGUIDResponseObject) GetLinksOk() (*LinksSelf, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CustomAAGUIDResponseObject) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksSelf and assigns it to the Links field.
func (o *CustomAAGUIDResponseObject) SetLinks(v LinksSelf) {
	o.Links = &v
}

func (o CustomAAGUIDResponseObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomAAGUIDResponseObject) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CustomAAGUIDResponseObject) UnmarshalJSON(data []byte) (err error) {
	varCustomAAGUIDResponseObject := _CustomAAGUIDResponseObject{}

	err = json.Unmarshal(data, &varCustomAAGUIDResponseObject)

	if err != nil {
		return err
	}

	*o = CustomAAGUIDResponseObject(varCustomAAGUIDResponseObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "aaguid")
		delete(additionalProperties, "attestationRootCertificates")
		delete(additionalProperties, "authenticatorCharacteristics")
		delete(additionalProperties, "name")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCustomAAGUIDResponseObject struct {
	value *CustomAAGUIDResponseObject
	isSet bool
}

func (v NullableCustomAAGUIDResponseObject) Get() *CustomAAGUIDResponseObject {
	return v.value
}

func (v *NullableCustomAAGUIDResponseObject) Set(val *CustomAAGUIDResponseObject) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomAAGUIDResponseObject) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomAAGUIDResponseObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomAAGUIDResponseObject(val *CustomAAGUIDResponseObject) *NullableCustomAAGUIDResponseObject {
	return &NullableCustomAAGUIDResponseObject{value: val, isSet: true}
}

func (v NullableCustomAAGUIDResponseObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomAAGUIDResponseObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
