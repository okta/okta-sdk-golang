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

// checks if the CertificateAuthorityToScopeMapping type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CertificateAuthorityToScopeMapping{}

// CertificateAuthorityToScopeMapping A persisted Certificate Authority to scope mapping. Returned as items in the `PUT /device-identity/api/v1/certificate-authority-mappings` response, representing the active mappings resulting from the reconcile.
type CertificateAuthorityToScopeMapping struct {
	// The `id` of the Trusted Certificate Authority that's mapped to the scope.
	CaId string `json:"caId"`
	// The unique identifier of the Certificate Authority mapping.
	Id string `json:"id"`
	// The scope (feature) that a Certificate Authority mapping is associated with.
	Scope                string `json:"scope"`
	AdditionalProperties map[string]interface{}
}

type _CertificateAuthorityToScopeMapping CertificateAuthorityToScopeMapping

// NewCertificateAuthorityToScopeMapping instantiates a new CertificateAuthorityToScopeMapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificateAuthorityToScopeMapping(caId string, id string, scope string) *CertificateAuthorityToScopeMapping {
	this := CertificateAuthorityToScopeMapping{}
	this.CaId = caId
	this.Id = id
	this.Scope = scope
	return &this
}

// NewCertificateAuthorityToScopeMappingWithDefaults instantiates a new CertificateAuthorityToScopeMapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateAuthorityToScopeMappingWithDefaults() *CertificateAuthorityToScopeMapping {
	this := CertificateAuthorityToScopeMapping{}
	return &this
}

// GetCaId returns the CaId field value
func (o *CertificateAuthorityToScopeMapping) GetCaId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CaId
}

// GetCaIdOk returns a tuple with the CaId field value
// and a boolean to check if the value has been set.
func (o *CertificateAuthorityToScopeMapping) GetCaIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CaId, true
}

// SetCaId sets field value
func (o *CertificateAuthorityToScopeMapping) SetCaId(v string) {
	o.CaId = v
}

// GetId returns the Id field value
func (o *CertificateAuthorityToScopeMapping) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CertificateAuthorityToScopeMapping) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CertificateAuthorityToScopeMapping) SetId(v string) {
	o.Id = v
}

// GetScope returns the Scope field value
func (o *CertificateAuthorityToScopeMapping) GetScope() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *CertificateAuthorityToScopeMapping) GetScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *CertificateAuthorityToScopeMapping) SetScope(v string) {
	o.Scope = v
}

func (o CertificateAuthorityToScopeMapping) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CertificateAuthorityToScopeMapping) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["caId"] = o.CaId
	toSerialize["id"] = o.Id
	toSerialize["scope"] = o.Scope

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CertificateAuthorityToScopeMapping) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"caId",
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

	varCertificateAuthorityToScopeMapping := _CertificateAuthorityToScopeMapping{}

	err = json.Unmarshal(data, &varCertificateAuthorityToScopeMapping)

	if err != nil {
		return err
	}

	*o = CertificateAuthorityToScopeMapping(varCertificateAuthorityToScopeMapping)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "caId")
		delete(additionalProperties, "id")
		delete(additionalProperties, "scope")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCertificateAuthorityToScopeMapping struct {
	value *CertificateAuthorityToScopeMapping
	isSet bool
}

func (v NullableCertificateAuthorityToScopeMapping) Get() *CertificateAuthorityToScopeMapping {
	return v.value
}

func (v *NullableCertificateAuthorityToScopeMapping) Set(val *CertificateAuthorityToScopeMapping) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateAuthorityToScopeMapping) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateAuthorityToScopeMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateAuthorityToScopeMapping(val *CertificateAuthorityToScopeMapping) *NullableCertificateAuthorityToScopeMapping {
	return &NullableCertificateAuthorityToScopeMapping{value: val, isSet: true}
}

func (v NullableCertificateAuthorityToScopeMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateAuthorityToScopeMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
