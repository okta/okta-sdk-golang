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

// checks if the SamlAssertionEncryption type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SamlAssertionEncryption{}

// SamlAssertionEncryption Determines if the app supports encrypted assertions
type SamlAssertionEncryption struct {
	// Indicates whether Okta encrypts the assertions that it sends to the Service Provider
	Enabled *bool `json:"enabled,omitempty"`
	// The encryption algorithm used to encrypt the SAML assertion
	EncryptionAlgorithm *string `json:"encryptionAlgorithm,omitempty"`
	// The key transport algorithm used to encrypt the SAML assertion
	KeyTransportAlgorithm *string `json:"keyTransportAlgorithm,omitempty"`
	// A list that contains exactly one x509 encoded certificate which Okta uses to encrypt the SAML assertion with
	X5c                  []string `json:"x5c,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SamlAssertionEncryption SamlAssertionEncryption

// NewSamlAssertionEncryption instantiates a new SamlAssertionEncryption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSamlAssertionEncryption() *SamlAssertionEncryption {
	this := SamlAssertionEncryption{}
	return &this
}

// NewSamlAssertionEncryptionWithDefaults instantiates a new SamlAssertionEncryption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSamlAssertionEncryptionWithDefaults() *SamlAssertionEncryption {
	this := SamlAssertionEncryption{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SamlAssertionEncryption) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlAssertionEncryption) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SamlAssertionEncryption) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SamlAssertionEncryption) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetEncryptionAlgorithm returns the EncryptionAlgorithm field value if set, zero value otherwise.
func (o *SamlAssertionEncryption) GetEncryptionAlgorithm() string {
	if o == nil || IsNil(o.EncryptionAlgorithm) {
		var ret string
		return ret
	}
	return *o.EncryptionAlgorithm
}

// GetEncryptionAlgorithmOk returns a tuple with the EncryptionAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlAssertionEncryption) GetEncryptionAlgorithmOk() (*string, bool) {
	if o == nil || IsNil(o.EncryptionAlgorithm) {
		return nil, false
	}
	return o.EncryptionAlgorithm, true
}

// HasEncryptionAlgorithm returns a boolean if a field has been set.
func (o *SamlAssertionEncryption) HasEncryptionAlgorithm() bool {
	if o != nil && !IsNil(o.EncryptionAlgorithm) {
		return true
	}

	return false
}

// SetEncryptionAlgorithm gets a reference to the given string and assigns it to the EncryptionAlgorithm field.
func (o *SamlAssertionEncryption) SetEncryptionAlgorithm(v string) {
	o.EncryptionAlgorithm = &v
}

// GetKeyTransportAlgorithm returns the KeyTransportAlgorithm field value if set, zero value otherwise.
func (o *SamlAssertionEncryption) GetKeyTransportAlgorithm() string {
	if o == nil || IsNil(o.KeyTransportAlgorithm) {
		var ret string
		return ret
	}
	return *o.KeyTransportAlgorithm
}

// GetKeyTransportAlgorithmOk returns a tuple with the KeyTransportAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlAssertionEncryption) GetKeyTransportAlgorithmOk() (*string, bool) {
	if o == nil || IsNil(o.KeyTransportAlgorithm) {
		return nil, false
	}
	return o.KeyTransportAlgorithm, true
}

// HasKeyTransportAlgorithm returns a boolean if a field has been set.
func (o *SamlAssertionEncryption) HasKeyTransportAlgorithm() bool {
	if o != nil && !IsNil(o.KeyTransportAlgorithm) {
		return true
	}

	return false
}

// SetKeyTransportAlgorithm gets a reference to the given string and assigns it to the KeyTransportAlgorithm field.
func (o *SamlAssertionEncryption) SetKeyTransportAlgorithm(v string) {
	o.KeyTransportAlgorithm = &v
}

// GetX5c returns the X5c field value if set, zero value otherwise.
func (o *SamlAssertionEncryption) GetX5c() []string {
	if o == nil || IsNil(o.X5c) {
		var ret []string
		return ret
	}
	return o.X5c
}

// GetX5cOk returns a tuple with the X5c field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlAssertionEncryption) GetX5cOk() ([]string, bool) {
	if o == nil || IsNil(o.X5c) {
		return nil, false
	}
	return o.X5c, true
}

// HasX5c returns a boolean if a field has been set.
func (o *SamlAssertionEncryption) HasX5c() bool {
	if o != nil && !IsNil(o.X5c) {
		return true
	}

	return false
}

// SetX5c gets a reference to the given []string and assigns it to the X5c field.
func (o *SamlAssertionEncryption) SetX5c(v []string) {
	o.X5c = v
}

func (o SamlAssertionEncryption) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SamlAssertionEncryption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.EncryptionAlgorithm) {
		toSerialize["encryptionAlgorithm"] = o.EncryptionAlgorithm
	}
	if !IsNil(o.KeyTransportAlgorithm) {
		toSerialize["keyTransportAlgorithm"] = o.KeyTransportAlgorithm
	}
	if !IsNil(o.X5c) {
		toSerialize["x5c"] = o.X5c
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SamlAssertionEncryption) UnmarshalJSON(data []byte) (err error) {
	varSamlAssertionEncryption := _SamlAssertionEncryption{}

	err = json.Unmarshal(data, &varSamlAssertionEncryption)

	if err != nil {
		return err
	}

	*o = SamlAssertionEncryption(varSamlAssertionEncryption)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "encryptionAlgorithm")
		delete(additionalProperties, "keyTransportAlgorithm")
		delete(additionalProperties, "x5c")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSamlAssertionEncryption struct {
	value *SamlAssertionEncryption
	isSet bool
}

func (v NullableSamlAssertionEncryption) Get() *SamlAssertionEncryption {
	return v.value
}

func (v *NullableSamlAssertionEncryption) Set(val *SamlAssertionEncryption) {
	v.value = val
	v.isSet = true
}

func (v NullableSamlAssertionEncryption) IsSet() bool {
	return v.isSet
}

func (v *NullableSamlAssertionEncryption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSamlAssertionEncryption(val *SamlAssertionEncryption) *NullableSamlAssertionEncryption {
	return &NullableSamlAssertionEncryption{value: val, isSet: true}
}

func (v NullableSamlAssertionEncryption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSamlAssertionEncryption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
