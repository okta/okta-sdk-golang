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

API version: 2024.06.1
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// DomainLinks struct for DomainLinks
type DomainLinks struct {
	Self *HrefObjectSelfLink `json:"self,omitempty"`
	Brand *DomainLinksAllOfBrand `json:"brand,omitempty"`
	Certificate *DomainLinksAllOfCertificate `json:"certificate,omitempty"`
	Verify *DomainLinksAllOfVerify `json:"verify,omitempty"`
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
	if o == nil || o.Self == nil {
		var ret HrefObjectSelfLink
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainLinks) GetSelfOk() (*HrefObjectSelfLink, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *DomainLinks) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given HrefObjectSelfLink and assigns it to the Self field.
func (o *DomainLinks) SetSelf(v HrefObjectSelfLink) {
	o.Self = &v
}

// GetBrand returns the Brand field value if set, zero value otherwise.
func (o *DomainLinks) GetBrand() DomainLinksAllOfBrand {
	if o == nil || o.Brand == nil {
		var ret DomainLinksAllOfBrand
		return ret
	}
	return *o.Brand
}

// GetBrandOk returns a tuple with the Brand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainLinks) GetBrandOk() (*DomainLinksAllOfBrand, bool) {
	if o == nil || o.Brand == nil {
		return nil, false
	}
	return o.Brand, true
}

// HasBrand returns a boolean if a field has been set.
func (o *DomainLinks) HasBrand() bool {
	if o != nil && o.Brand != nil {
		return true
	}

	return false
}

// SetBrand gets a reference to the given DomainLinksAllOfBrand and assigns it to the Brand field.
func (o *DomainLinks) SetBrand(v DomainLinksAllOfBrand) {
	o.Brand = &v
}

// GetCertificate returns the Certificate field value if set, zero value otherwise.
func (o *DomainLinks) GetCertificate() DomainLinksAllOfCertificate {
	if o == nil || o.Certificate == nil {
		var ret DomainLinksAllOfCertificate
		return ret
	}
	return *o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainLinks) GetCertificateOk() (*DomainLinksAllOfCertificate, bool) {
	if o == nil || o.Certificate == nil {
		return nil, false
	}
	return o.Certificate, true
}

// HasCertificate returns a boolean if a field has been set.
func (o *DomainLinks) HasCertificate() bool {
	if o != nil && o.Certificate != nil {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given DomainLinksAllOfCertificate and assigns it to the Certificate field.
func (o *DomainLinks) SetCertificate(v DomainLinksAllOfCertificate) {
	o.Certificate = &v
}

// GetVerify returns the Verify field value if set, zero value otherwise.
func (o *DomainLinks) GetVerify() DomainLinksAllOfVerify {
	if o == nil || o.Verify == nil {
		var ret DomainLinksAllOfVerify
		return ret
	}
	return *o.Verify
}

// GetVerifyOk returns a tuple with the Verify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainLinks) GetVerifyOk() (*DomainLinksAllOfVerify, bool) {
	if o == nil || o.Verify == nil {
		return nil, false
	}
	return o.Verify, true
}

// HasVerify returns a boolean if a field has been set.
func (o *DomainLinks) HasVerify() bool {
	if o != nil && o.Verify != nil {
		return true
	}

	return false
}

// SetVerify gets a reference to the given DomainLinksAllOfVerify and assigns it to the Verify field.
func (o *DomainLinks) SetVerify(v DomainLinksAllOfVerify) {
	o.Verify = &v
}

func (o DomainLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Self != nil {
		toSerialize["self"] = o.Self
	}
	if o.Brand != nil {
		toSerialize["brand"] = o.Brand
	}
	if o.Certificate != nil {
		toSerialize["certificate"] = o.Certificate
	}
	if o.Verify != nil {
		toSerialize["verify"] = o.Verify
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DomainLinks) UnmarshalJSON(bytes []byte) (err error) {
	varDomainLinks := _DomainLinks{}

	err = json.Unmarshal(bytes, &varDomainLinks)
	if err == nil {
		*o = DomainLinks(varDomainLinks)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "brand")
		delete(additionalProperties, "certificate")
		delete(additionalProperties, "verify")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
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

