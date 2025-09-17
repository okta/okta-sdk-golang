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
	"time"
)

// checks if the Csr type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Csr{}

// Csr struct for Csr
type Csr struct {
	// Timestamp when the object was created
	Created              *time.Time `json:"created,omitempty"`
	Csr                  *string    `json:"csr,omitempty"`
	Id                   *string    `json:"id,omitempty"`
	Kty                  *string    `json:"kty,omitempty"`
	Links                *CSRLinks  `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Csr Csr

// NewCsr instantiates a new Csr object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCsr() *Csr {
	this := Csr{}
	return &this
}

// NewCsrWithDefaults instantiates a new Csr object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCsrWithDefaults() *Csr {
	this := Csr{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *Csr) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Csr) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *Csr) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *Csr) SetCreated(v time.Time) {
	o.Created = &v
}

// GetCsr returns the Csr field value if set, zero value otherwise.
func (o *Csr) GetCsr() string {
	if o == nil || IsNil(o.Csr) {
		var ret string
		return ret
	}
	return *o.Csr
}

// GetCsrOk returns a tuple with the Csr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Csr) GetCsrOk() (*string, bool) {
	if o == nil || IsNil(o.Csr) {
		return nil, false
	}
	return o.Csr, true
}

// HasCsr returns a boolean if a field has been set.
func (o *Csr) HasCsr() bool {
	if o != nil && !IsNil(o.Csr) {
		return true
	}

	return false
}

// SetCsr gets a reference to the given string and assigns it to the Csr field.
func (o *Csr) SetCsr(v string) {
	o.Csr = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Csr) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Csr) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Csr) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Csr) SetId(v string) {
	o.Id = &v
}

// GetKty returns the Kty field value if set, zero value otherwise.
func (o *Csr) GetKty() string {
	if o == nil || IsNil(o.Kty) {
		var ret string
		return ret
	}
	return *o.Kty
}

// GetKtyOk returns a tuple with the Kty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Csr) GetKtyOk() (*string, bool) {
	if o == nil || IsNil(o.Kty) {
		return nil, false
	}
	return o.Kty, true
}

// HasKty returns a boolean if a field has been set.
func (o *Csr) HasKty() bool {
	if o != nil && !IsNil(o.Kty) {
		return true
	}

	return false
}

// SetKty gets a reference to the given string and assigns it to the Kty field.
func (o *Csr) SetKty(v string) {
	o.Kty = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Csr) GetLinks() CSRLinks {
	if o == nil || IsNil(o.Links) {
		var ret CSRLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Csr) GetLinksOk() (*CSRLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Csr) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given CSRLinks and assigns it to the Links field.
func (o *Csr) SetLinks(v CSRLinks) {
	o.Links = &v
}

func (o Csr) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Csr) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Csr) {
		toSerialize["csr"] = o.Csr
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Kty) {
		toSerialize["kty"] = o.Kty
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Csr) UnmarshalJSON(data []byte) (err error) {
	varCsr := _Csr{}

	err = json.Unmarshal(data, &varCsr)

	if err != nil {
		return err
	}

	*o = Csr(varCsr)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "created")
		delete(additionalProperties, "csr")
		delete(additionalProperties, "id")
		delete(additionalProperties, "kty")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCsr struct {
	value *Csr
	isSet bool
}

func (v NullableCsr) Get() *Csr {
	return v.value
}

func (v *NullableCsr) Set(val *Csr) {
	v.value = val
	v.isSet = true
}

func (v NullableCsr) IsSet() bool {
	return v.isSet
}

func (v *NullableCsr) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCsr(val *Csr) *NullableCsr {
	return &NullableCsr{value: val, isSet: true}
}

func (v NullableCsr) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCsr) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
