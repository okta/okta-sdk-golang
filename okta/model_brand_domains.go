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

// checks if the BrandDomains type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BrandDomains{}

// BrandDomains Defines a list of domains with a subset of the properties for each domain
type BrandDomains struct {
	// Each element of the array defines an individual domain
	Domains              []DomainResponse `json:"domains,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BrandDomains BrandDomains

// NewBrandDomains instantiates a new BrandDomains object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrandDomains() *BrandDomains {
	this := BrandDomains{}
	return &this
}

// NewBrandDomainsWithDefaults instantiates a new BrandDomains object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrandDomainsWithDefaults() *BrandDomains {
	this := BrandDomains{}
	return &this
}

// GetDomains returns the Domains field value if set, zero value otherwise.
func (o *BrandDomains) GetDomains() []DomainResponse {
	if o == nil || IsNil(o.Domains) {
		var ret []DomainResponse
		return ret
	}
	return o.Domains
}

// GetDomainsOk returns a tuple with the Domains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandDomains) GetDomainsOk() ([]DomainResponse, bool) {
	if o == nil || IsNil(o.Domains) {
		return nil, false
	}
	return o.Domains, true
}

// HasDomains returns a boolean if a field has been set.
func (o *BrandDomains) HasDomains() bool {
	if o != nil && !IsNil(o.Domains) {
		return true
	}

	return false
}

// SetDomains gets a reference to the given []DomainResponse and assigns it to the Domains field.
func (o *BrandDomains) SetDomains(v []DomainResponse) {
	o.Domains = v
}

func (o BrandDomains) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BrandDomains) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Domains) {
		toSerialize["domains"] = o.Domains
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BrandDomains) UnmarshalJSON(data []byte) (err error) {
	varBrandDomains := _BrandDomains{}

	err = json.Unmarshal(data, &varBrandDomains)

	if err != nil {
		return err
	}

	*o = BrandDomains(varBrandDomains)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "domains")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBrandDomains struct {
	value *BrandDomains
	isSet bool
}

func (v NullableBrandDomains) Get() *BrandDomains {
	return v.value
}

func (v *NullableBrandDomains) Set(val *BrandDomains) {
	v.value = val
	v.isSet = true
}

func (v NullableBrandDomains) IsSet() bool {
	return v.isSet
}

func (v *NullableBrandDomains) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrandDomains(val *BrandDomains) *NullableBrandDomains {
	return &NullableBrandDomains{value: val, isSet: true}
}

func (v NullableBrandDomains) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBrandDomains) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
