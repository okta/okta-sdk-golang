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
	"fmt"
)

// checks if the IdPCertificateCredential type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdPCertificateCredential{}

// IdPCertificateCredential struct for IdPCertificateCredential
type IdPCertificateCredential struct {
	// Base64-encoded X.509 certificate chain with DER encoding
	X5c                  []string `json:"x5c"`
	AdditionalProperties map[string]interface{}
}

type _IdPCertificateCredential IdPCertificateCredential

// NewIdPCertificateCredential instantiates a new IdPCertificateCredential object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdPCertificateCredential(x5c []string) *IdPCertificateCredential {
	this := IdPCertificateCredential{}
	this.X5c = x5c
	return &this
}

// NewIdPCertificateCredentialWithDefaults instantiates a new IdPCertificateCredential object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdPCertificateCredentialWithDefaults() *IdPCertificateCredential {
	this := IdPCertificateCredential{}
	return &this
}

// GetX5c returns the X5c field value
func (o *IdPCertificateCredential) GetX5c() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.X5c
}

// GetX5cOk returns a tuple with the X5c field value
// and a boolean to check if the value has been set.
func (o *IdPCertificateCredential) GetX5cOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.X5c, true
}

// SetX5c sets field value
func (o *IdPCertificateCredential) SetX5c(v []string) {
	o.X5c = v
}

func (o IdPCertificateCredential) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdPCertificateCredential) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["x5c"] = o.X5c

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IdPCertificateCredential) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"x5c",
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

	varIdPCertificateCredential := _IdPCertificateCredential{}

	err = json.Unmarshal(data, &varIdPCertificateCredential)

	if err != nil {
		return err
	}

	*o = IdPCertificateCredential(varIdPCertificateCredential)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "x5c")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdPCertificateCredential struct {
	value *IdPCertificateCredential
	isSet bool
}

func (v NullableIdPCertificateCredential) Get() *IdPCertificateCredential {
	return v.value
}

func (v *NullableIdPCertificateCredential) Set(val *IdPCertificateCredential) {
	v.value = val
	v.isSet = true
}

func (v NullableIdPCertificateCredential) IsSet() bool {
	return v.isSet
}

func (v *NullableIdPCertificateCredential) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdPCertificateCredential(val *IdPCertificateCredential) *NullableIdPCertificateCredential {
	return &NullableIdPCertificateCredential{value: val, isSet: true}
}

func (v NullableIdPCertificateCredential) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdPCertificateCredential) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
