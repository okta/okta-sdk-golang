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
	"fmt"
)

// checks if the CertificateAuthorityToScopeMappingResponseMapping type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CertificateAuthorityToScopeMappingResponseMapping{}

// CertificateAuthorityToScopeMappingResponseMapping The Certificate Authority's mapping for the requested scope. Omitted when the Certificate Authority isn't mapped to the requested scope.
type CertificateAuthorityToScopeMappingResponseMapping struct {
	// The unique identifier of the Certificate Authority mapping.
	Id string `json:"id"`
	// The scope (feature) that a Certificate Authority mapping is associated with.
	Scope                string `json:"scope"`
	AdditionalProperties map[string]interface{}
}

type _CertificateAuthorityToScopeMappingResponseMapping CertificateAuthorityToScopeMappingResponseMapping

// NewCertificateAuthorityToScopeMappingResponseMapping instantiates a new CertificateAuthorityToScopeMappingResponseMapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificateAuthorityToScopeMappingResponseMapping(id string, scope string) *CertificateAuthorityToScopeMappingResponseMapping {
	this := CertificateAuthorityToScopeMappingResponseMapping{}
	this.Id = id
	this.Scope = scope
	return &this
}

// NewCertificateAuthorityToScopeMappingResponseMappingWithDefaults instantiates a new CertificateAuthorityToScopeMappingResponseMapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateAuthorityToScopeMappingResponseMappingWithDefaults() *CertificateAuthorityToScopeMappingResponseMapping {
	this := CertificateAuthorityToScopeMappingResponseMapping{}
	return &this
}

// GetId returns the Id field value
func (o *CertificateAuthorityToScopeMappingResponseMapping) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CertificateAuthorityToScopeMappingResponseMapping) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CertificateAuthorityToScopeMappingResponseMapping) SetId(v string) {
	o.Id = v
}

// GetScope returns the Scope field value
func (o *CertificateAuthorityToScopeMappingResponseMapping) GetScope() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *CertificateAuthorityToScopeMappingResponseMapping) GetScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *CertificateAuthorityToScopeMappingResponseMapping) SetScope(v string) {
	o.Scope = v
}

func (o CertificateAuthorityToScopeMappingResponseMapping) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CertificateAuthorityToScopeMappingResponseMapping) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["scope"] = o.Scope

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CertificateAuthorityToScopeMappingResponseMapping) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"scope",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCertificateAuthorityToScopeMappingResponseMapping := _CertificateAuthorityToScopeMappingResponseMapping{}

	err = json.Unmarshal(data, &varCertificateAuthorityToScopeMappingResponseMapping)

	if err != nil {
		return err
	}

	*o = CertificateAuthorityToScopeMappingResponseMapping(varCertificateAuthorityToScopeMappingResponseMapping)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "scope")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCertificateAuthorityToScopeMappingResponseMapping struct {
	value *CertificateAuthorityToScopeMappingResponseMapping
	isSet bool
}

func (v NullableCertificateAuthorityToScopeMappingResponseMapping) Get() *CertificateAuthorityToScopeMappingResponseMapping {
	return v.value
}

func (v *NullableCertificateAuthorityToScopeMappingResponseMapping) Set(val *CertificateAuthorityToScopeMappingResponseMapping) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateAuthorityToScopeMappingResponseMapping) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateAuthorityToScopeMappingResponseMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateAuthorityToScopeMappingResponseMapping(val *CertificateAuthorityToScopeMappingResponseMapping) *NullableCertificateAuthorityToScopeMappingResponseMapping {
	return &NullableCertificateAuthorityToScopeMappingResponseMapping{value: val, isSet: true}
}

func (v NullableCertificateAuthorityToScopeMappingResponseMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateAuthorityToScopeMappingResponseMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
