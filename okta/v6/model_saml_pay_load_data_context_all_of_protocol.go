/*
Okta Admin Management

Allows customers to easily access the Okta Management APIs

Copyright 2018 - Present Okta, Inc.

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
)

// SAMLPayLoadDataContextAllOfProtocol Details of the assertion protocol being used
type SAMLPayLoadDataContextAllOfProtocol struct {
	// The type of authentication protocol being used for the assertion
	Type *string `json:"type,omitempty"`
	Issuer *SAMLPayLoadDataContextAllOfProtocolIssuer `json:"issuer,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SAMLPayLoadDataContextAllOfProtocol SAMLPayLoadDataContextAllOfProtocol

// NewSAMLPayLoadDataContextAllOfProtocol instantiates a new SAMLPayLoadDataContextAllOfProtocol object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSAMLPayLoadDataContextAllOfProtocol() *SAMLPayLoadDataContextAllOfProtocol {
	this := SAMLPayLoadDataContextAllOfProtocol{}
	return &this
}

// NewSAMLPayLoadDataContextAllOfProtocolWithDefaults instantiates a new SAMLPayLoadDataContextAllOfProtocol object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSAMLPayLoadDataContextAllOfProtocolWithDefaults() *SAMLPayLoadDataContextAllOfProtocol {
	this := SAMLPayLoadDataContextAllOfProtocol{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SAMLPayLoadDataContextAllOfProtocol) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLPayLoadDataContextAllOfProtocol) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SAMLPayLoadDataContextAllOfProtocol) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SAMLPayLoadDataContextAllOfProtocol) SetType(v string) {
	o.Type = &v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *SAMLPayLoadDataContextAllOfProtocol) GetIssuer() SAMLPayLoadDataContextAllOfProtocolIssuer {
	if o == nil || o.Issuer == nil {
		var ret SAMLPayLoadDataContextAllOfProtocolIssuer
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLPayLoadDataContextAllOfProtocol) GetIssuerOk() (*SAMLPayLoadDataContextAllOfProtocolIssuer, bool) {
	if o == nil || o.Issuer == nil {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *SAMLPayLoadDataContextAllOfProtocol) HasIssuer() bool {
	if o != nil && o.Issuer != nil {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given SAMLPayLoadDataContextAllOfProtocolIssuer and assigns it to the Issuer field.
func (o *SAMLPayLoadDataContextAllOfProtocol) SetIssuer(v SAMLPayLoadDataContextAllOfProtocolIssuer) {
	o.Issuer = &v
}

func (o SAMLPayLoadDataContextAllOfProtocol) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Issuer != nil {
		toSerialize["issuer"] = o.Issuer
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SAMLPayLoadDataContextAllOfProtocol) UnmarshalJSON(bytes []byte) (err error) {
	varSAMLPayLoadDataContextAllOfProtocol := _SAMLPayLoadDataContextAllOfProtocol{}

	err = json.Unmarshal(bytes, &varSAMLPayLoadDataContextAllOfProtocol)
	if err == nil {
		*o = SAMLPayLoadDataContextAllOfProtocol(varSAMLPayLoadDataContextAllOfProtocol)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "issuer")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableSAMLPayLoadDataContextAllOfProtocol struct {
	value *SAMLPayLoadDataContextAllOfProtocol
	isSet bool
}

func (v NullableSAMLPayLoadDataContextAllOfProtocol) Get() *SAMLPayLoadDataContextAllOfProtocol {
	return v.value
}

func (v *NullableSAMLPayLoadDataContextAllOfProtocol) Set(val *SAMLPayLoadDataContextAllOfProtocol) {
	v.value = val
	v.isSet = true
}

func (v NullableSAMLPayLoadDataContextAllOfProtocol) IsSet() bool {
	return v.isSet
}

func (v *NullableSAMLPayLoadDataContextAllOfProtocol) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSAMLPayLoadDataContextAllOfProtocol(val *SAMLPayLoadDataContextAllOfProtocol) *NullableSAMLPayLoadDataContextAllOfProtocol {
	return &NullableSAMLPayLoadDataContextAllOfProtocol{value: val, isSet: true}
}

func (v NullableSAMLPayLoadDataContextAllOfProtocol) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSAMLPayLoadDataContextAllOfProtocol) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

