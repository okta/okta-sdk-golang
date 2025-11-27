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
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the STSVaultSecretConnectionCreatable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &STSVaultSecretConnectionCreatable{}

// STSVaultSecretConnectionCreatable Create an STS connection for a vaulted secret
type STSVaultSecretConnectionCreatable struct {
	// Type of connection authentication method
	ConnectionType string `json:"connectionType"`
	// The authentication protocol type used for the connection
	ProtocolType *string                                 `json:"protocolType,omitempty"`
	Secret       STSVaultSecretConnectionCreatableSecret `json:"secret"`
}

type _STSVaultSecretConnectionCreatable STSVaultSecretConnectionCreatable

// NewSTSVaultSecretConnectionCreatable instantiates a new STSVaultSecretConnectionCreatable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSTSVaultSecretConnectionCreatable(connectionType string, secret STSVaultSecretConnectionCreatableSecret) *STSVaultSecretConnectionCreatable {
	this := STSVaultSecretConnectionCreatable{}
	this.ConnectionType = connectionType
	this.Secret = secret
	return &this
}

// NewSTSVaultSecretConnectionCreatableWithDefaults instantiates a new STSVaultSecretConnectionCreatable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSTSVaultSecretConnectionCreatableWithDefaults() *STSVaultSecretConnectionCreatable {
	this := STSVaultSecretConnectionCreatable{}
	return &this
}

// GetConnectionType returns the ConnectionType field value
func (o *STSVaultSecretConnectionCreatable) GetConnectionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectionType
}

// GetConnectionTypeOk returns a tuple with the ConnectionType field value
// and a boolean to check if the value has been set.
func (o *STSVaultSecretConnectionCreatable) GetConnectionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionType, true
}

// SetConnectionType sets field value
func (o *STSVaultSecretConnectionCreatable) SetConnectionType(v string) {
	o.ConnectionType = v
}

// GetProtocolType returns the ProtocolType field value if set, zero value otherwise.
func (o *STSVaultSecretConnectionCreatable) GetProtocolType() string {
	if o == nil || IsNil(o.ProtocolType) {
		var ret string
		return ret
	}
	return *o.ProtocolType
}

// GetProtocolTypeOk returns a tuple with the ProtocolType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *STSVaultSecretConnectionCreatable) GetProtocolTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ProtocolType) {
		return nil, false
	}
	return o.ProtocolType, true
}

// HasProtocolType returns a boolean if a field has been set.
func (o *STSVaultSecretConnectionCreatable) HasProtocolType() bool {
	if o != nil && !IsNil(o.ProtocolType) {
		return true
	}

	return false
}

// SetProtocolType gets a reference to the given string and assigns it to the ProtocolType field.
func (o *STSVaultSecretConnectionCreatable) SetProtocolType(v string) {
	o.ProtocolType = &v
}

// GetSecret returns the Secret field value
func (o *STSVaultSecretConnectionCreatable) GetSecret() STSVaultSecretConnectionCreatableSecret {
	if o == nil {
		var ret STSVaultSecretConnectionCreatableSecret
		return ret
	}

	return o.Secret
}

// GetSecretOk returns a tuple with the Secret field value
// and a boolean to check if the value has been set.
func (o *STSVaultSecretConnectionCreatable) GetSecretOk() (*STSVaultSecretConnectionCreatableSecret, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Secret, true
}

// SetSecret sets field value
func (o *STSVaultSecretConnectionCreatable) SetSecret(v STSVaultSecretConnectionCreatableSecret) {
	o.Secret = v
}

func (o STSVaultSecretConnectionCreatable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o STSVaultSecretConnectionCreatable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["connectionType"] = o.ConnectionType
	if !IsNil(o.ProtocolType) {
		toSerialize["protocolType"] = o.ProtocolType
	}
	toSerialize["secret"] = o.Secret
	return toSerialize, nil
}

func (o *STSVaultSecretConnectionCreatable) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"connectionType",
		"secret",
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

	varSTSVaultSecretConnectionCreatable := _STSVaultSecretConnectionCreatable{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSTSVaultSecretConnectionCreatable)

	if err != nil {
		return err
	}

	*o = STSVaultSecretConnectionCreatable(varSTSVaultSecretConnectionCreatable)

	return err
}

type NullableSTSVaultSecretConnectionCreatable struct {
	value *STSVaultSecretConnectionCreatable
	isSet bool
}

func (v NullableSTSVaultSecretConnectionCreatable) Get() *STSVaultSecretConnectionCreatable {
	return v.value
}

func (v *NullableSTSVaultSecretConnectionCreatable) Set(val *STSVaultSecretConnectionCreatable) {
	v.value = val
	v.isSet = true
}

func (v NullableSTSVaultSecretConnectionCreatable) IsSet() bool {
	return v.isSet
}

func (v *NullableSTSVaultSecretConnectionCreatable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSTSVaultSecretConnectionCreatable(val *STSVaultSecretConnectionCreatable) *NullableSTSVaultSecretConnectionCreatable {
	return &NullableSTSVaultSecretConnectionCreatable{value: val, isSet: true}
}

func (v NullableSTSVaultSecretConnectionCreatable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSTSVaultSecretConnectionCreatable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
