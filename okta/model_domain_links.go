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

// checks if the DomainLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DomainLinks{}

// DomainLinks struct for DomainLinks
type DomainLinks struct {
	Self *HrefObjectSelfLink `json:"self,omitempty"`
	// The associated brand
	Brand *HrefObject `json:"brand,omitempty"`
	// The certificate link references the domain certificate
	Certificate *HrefObject `json:"certificate,omitempty"`
	// The verify link verifies the domain and transitions the domain status to `VERIFIED`
	Verify               *HrefObject `json:"verify,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DomainLinks DomainLinks

// NewDomainLinks instantiates a new DomainLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainLinks() *DomainLinks {
	this := DomainLinks{}
	return &this
}

// NewDomainLinksWithDefaults instantiates a new DomainLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainLinksWithDefaults() *DomainLinks {
	this := DomainLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *DomainLinks) GetSelf() HrefObjectSelfLink {
	if o == nil || IsNil(o.Self) {
		var ret HrefObjectSelfLink
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainLinks) GetSelfOk() (*HrefObjectSelfLink, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *DomainLinks) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given HrefObjectSelfLink and assigns it to the Self field.
func (o *DomainLinks) SetSelf(v HrefObjectSelfLink) {
	o.Self = &v
}

// GetBrand returns the Brand field value if set, zero value otherwise.
func (o *DomainLinks) GetBrand() HrefObject {
	if o == nil || IsNil(o.Brand) {
		var ret HrefObject
		return ret
	}
	return *o.Brand
}

// GetBrandOk returns a tuple with the Brand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainLinks) GetBrandOk() (*HrefObject, bool) {
	if o == nil || IsNil(o.Brand) {
		return nil, false
	}
	return o.Brand, true
}

// HasBrand returns a boolean if a field has been set.
func (o *DomainLinks) HasBrand() bool {
	if o != nil && !IsNil(o.Brand) {
		return true
	}

	return false
}

// SetBrand gets a reference to the given HrefObject and assigns it to the Brand field.
func (o *DomainLinks) SetBrand(v HrefObject) {
	o.Brand = &v
}

// GetCertificate returns the Certificate field value if set, zero value otherwise.
func (o *DomainLinks) GetCertificate() HrefObject {
	if o == nil || IsNil(o.Certificate) {
		var ret HrefObject
		return ret
	}
	return *o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainLinks) GetCertificateOk() (*HrefObject, bool) {
	if o == nil || IsNil(o.Certificate) {
		return nil, false
	}
	return o.Certificate, true
}

// HasCertificate returns a boolean if a field has been set.
func (o *DomainLinks) HasCertificate() bool {
	if o != nil && !IsNil(o.Certificate) {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given HrefObject and assigns it to the Certificate field.
func (o *DomainLinks) SetCertificate(v HrefObject) {
	o.Certificate = &v
}

// GetVerify returns the Verify field value if set, zero value otherwise.
func (o *DomainLinks) GetVerify() HrefObject {
	if o == nil || IsNil(o.Verify) {
		var ret HrefObject
		return ret
	}
	return *o.Verify
}

// GetVerifyOk returns a tuple with the Verify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainLinks) GetVerifyOk() (*HrefObject, bool) {
	if o == nil || IsNil(o.Verify) {
		return nil, false
	}
	return o.Verify, true
}

// HasVerify returns a boolean if a field has been set.
func (o *DomainLinks) HasVerify() bool {
	if o != nil && !IsNil(o.Verify) {
		return true
	}

	return false
}

// SetVerify gets a reference to the given HrefObject and assigns it to the Verify field.
func (o *DomainLinks) SetVerify(v HrefObject) {
	o.Verify = &v
}

func (o DomainLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DomainLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	if !IsNil(o.Brand) {
		toSerialize["brand"] = o.Brand
	}
	if !IsNil(o.Certificate) {
		toSerialize["certificate"] = o.Certificate
	}
	if !IsNil(o.Verify) {
		toSerialize["verify"] = o.Verify
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DomainLinks) UnmarshalJSON(data []byte) (err error) {
	varDomainLinks := _DomainLinks{}

	err = json.Unmarshal(data, &varDomainLinks)

	if err != nil {
		return err
	}

	*o = DomainLinks(varDomainLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "brand")
		delete(additionalProperties, "certificate")
		delete(additionalProperties, "verify")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDomainLinks struct {
	value *DomainLinks
	isSet bool
}

func (v NullableDomainLinks) Get() *DomainLinks {
	return v.value
}

func (v *NullableDomainLinks) Set(val *DomainLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainLinks(val *DomainLinks) *NullableDomainLinks {
	return &NullableDomainLinks{value: val, isSet: true}
}

func (v NullableDomainLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
