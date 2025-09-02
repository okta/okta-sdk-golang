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
	"time"
)

// IdPCsr Defines a CSR for a signature or decryption credential for an IdP
type IdPCsr struct {
	// Timestamp when the object was created
	Created *time.Time `json:"created,omitempty"`
	// Base64-encoded CSR in DER format
	Csr *string `json:"csr,omitempty"`
	// Unique identifier for the CSR
	Id *string `json:"id,omitempty"`
	// Cryptographic algorithm family for the CSR's keypair
	Kty *string `json:"kty,omitempty"`
	Links *IdPCsrLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdPCsr IdPCsr

// NewIdPCsr instantiates a new IdPCsr object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdPCsr() *IdPCsr {
	this := IdPCsr{}
	return &this
}

// NewIdPCsrWithDefaults instantiates a new IdPCsr object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdPCsrWithDefaults() *IdPCsr {
	this := IdPCsr{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *IdPCsr) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdPCsr) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *IdPCsr) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *IdPCsr) SetCreated(v time.Time) {
	o.Created = &v
}

// GetCsr returns the Csr field value if set, zero value otherwise.
func (o *IdPCsr) GetCsr() string {
	if o == nil || o.Csr == nil {
		var ret string
		return ret
	}
	return *o.Csr
}

// GetCsrOk returns a tuple with the Csr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdPCsr) GetCsrOk() (*string, bool) {
	if o == nil || o.Csr == nil {
		return nil, false
	}
	return o.Csr, true
}

// HasCsr returns a boolean if a field has been set.
func (o *IdPCsr) HasCsr() bool {
	if o != nil && o.Csr != nil {
		return true
	}

	return false
}

// SetCsr gets a reference to the given string and assigns it to the Csr field.
func (o *IdPCsr) SetCsr(v string) {
	o.Csr = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IdPCsr) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdPCsr) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IdPCsr) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IdPCsr) SetId(v string) {
	o.Id = &v
}

// GetKty returns the Kty field value if set, zero value otherwise.
func (o *IdPCsr) GetKty() string {
	if o == nil || o.Kty == nil {
		var ret string
		return ret
	}
	return *o.Kty
}

// GetKtyOk returns a tuple with the Kty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdPCsr) GetKtyOk() (*string, bool) {
	if o == nil || o.Kty == nil {
		return nil, false
	}
	return o.Kty, true
}

// HasKty returns a boolean if a field has been set.
func (o *IdPCsr) HasKty() bool {
	if o != nil && o.Kty != nil {
		return true
	}

	return false
}

// SetKty gets a reference to the given string and assigns it to the Kty field.
func (o *IdPCsr) SetKty(v string) {
	o.Kty = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *IdPCsr) GetLinks() IdPCsrLinks {
	if o == nil || o.Links == nil {
		var ret IdPCsrLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdPCsr) GetLinksOk() (*IdPCsrLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *IdPCsr) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given IdPCsrLinks and assigns it to the Links field.
func (o *IdPCsr) SetLinks(v IdPCsrLinks) {
	o.Links = &v
}

func (o IdPCsr) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.Csr != nil {
		toSerialize["csr"] = o.Csr
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Kty != nil {
		toSerialize["kty"] = o.Kty
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IdPCsr) UnmarshalJSON(bytes []byte) (err error) {
	varIdPCsr := _IdPCsr{}

	err = json.Unmarshal(bytes, &varIdPCsr)
	if err == nil {
		*o = IdPCsr(varIdPCsr)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "created")
		delete(additionalProperties, "csr")
		delete(additionalProperties, "id")
		delete(additionalProperties, "kty")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableIdPCsr struct {
	value *IdPCsr
	isSet bool
}

func (v NullableIdPCsr) Get() *IdPCsr {
	return v.value
}

func (v *NullableIdPCsr) Set(val *IdPCsr) {
	v.value = val
	v.isSet = true
}

func (v NullableIdPCsr) IsSet() bool {
	return v.isSet
}

func (v *NullableIdPCsr) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdPCsr(val *IdPCsr) *NullableIdPCsr {
	return &NullableIdPCsr{value: val, isSet: true}
}

func (v NullableIdPCsr) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdPCsr) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

