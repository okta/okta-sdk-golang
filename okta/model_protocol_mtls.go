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

// checks if the ProtocolMtls type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProtocolMtls{}

// ProtocolMtls Protocol settings for the [MTLS Protocol](https://tools.ietf.org/html/rfc5246#section-7.4.4)
type ProtocolMtls struct {
	Credentials *MtlsCredentials `json:"credentials,omitempty"`
	Endpoints   *MtlsEndpoints   `json:"endpoints,omitempty"`
	// Mutual TLS
	Type                 *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProtocolMtls ProtocolMtls

// NewProtocolMtls instantiates a new ProtocolMtls object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProtocolMtls() *ProtocolMtls {
	this := ProtocolMtls{}
	return &this
}

// NewProtocolMtlsWithDefaults instantiates a new ProtocolMtls object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProtocolMtlsWithDefaults() *ProtocolMtls {
	this := ProtocolMtls{}
	return &this
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *ProtocolMtls) GetCredentials() MtlsCredentials {
	if o == nil || IsNil(o.Credentials) {
		var ret MtlsCredentials
		return ret
	}
	return *o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtocolMtls) GetCredentialsOk() (*MtlsCredentials, bool) {
	if o == nil || IsNil(o.Credentials) {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *ProtocolMtls) HasCredentials() bool {
	if o != nil && !IsNil(o.Credentials) {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given MtlsCredentials and assigns it to the Credentials field.
func (o *ProtocolMtls) SetCredentials(v MtlsCredentials) {
	o.Credentials = &v
}

// GetEndpoints returns the Endpoints field value if set, zero value otherwise.
func (o *ProtocolMtls) GetEndpoints() MtlsEndpoints {
	if o == nil || IsNil(o.Endpoints) {
		var ret MtlsEndpoints
		return ret
	}
	return *o.Endpoints
}

// GetEndpointsOk returns a tuple with the Endpoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtocolMtls) GetEndpointsOk() (*MtlsEndpoints, bool) {
	if o == nil || IsNil(o.Endpoints) {
		return nil, false
	}
	return o.Endpoints, true
}

// HasEndpoints returns a boolean if a field has been set.
func (o *ProtocolMtls) HasEndpoints() bool {
	if o != nil && !IsNil(o.Endpoints) {
		return true
	}

	return false
}

// SetEndpoints gets a reference to the given MtlsEndpoints and assigns it to the Endpoints field.
func (o *ProtocolMtls) SetEndpoints(v MtlsEndpoints) {
	o.Endpoints = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ProtocolMtls) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtocolMtls) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ProtocolMtls) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ProtocolMtls) SetType(v string) {
	o.Type = &v
}

func (o ProtocolMtls) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProtocolMtls) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Credentials) {
		toSerialize["credentials"] = o.Credentials
	}
	if !IsNil(o.Endpoints) {
		toSerialize["endpoints"] = o.Endpoints
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProtocolMtls) UnmarshalJSON(data []byte) (err error) {
	varProtocolMtls := _ProtocolMtls{}

	err = json.Unmarshal(data, &varProtocolMtls)

	if err != nil {
		return err
	}

	*o = ProtocolMtls(varProtocolMtls)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "credentials")
		delete(additionalProperties, "endpoints")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProtocolMtls struct {
	value *ProtocolMtls
	isSet bool
}

func (v NullableProtocolMtls) Get() *ProtocolMtls {
	return v.value
}

func (v *NullableProtocolMtls) Set(val *ProtocolMtls) {
	v.value = val
	v.isSet = true
}

func (v NullableProtocolMtls) IsSet() bool {
	return v.isSet
}

func (v *NullableProtocolMtls) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProtocolMtls(val *ProtocolMtls) *NullableProtocolMtls {
	return &NullableProtocolMtls{value: val, isSet: true}
}

func (v NullableProtocolMtls) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProtocolMtls) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
