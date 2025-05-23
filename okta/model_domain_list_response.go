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

// DomainListResponse Defines a list of domains with a subset of the properties for each domain.
type DomainListResponse struct {
	// Each element of the array defines an individual domain.
	Domains []DomainResponse `json:"domains,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DomainListResponse DomainListResponse

// NewDomainListResponse instantiates a new DomainListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainListResponse() *DomainListResponse {
	this := DomainListResponse{}
	return &this
}

// NewDomainListResponseWithDefaults instantiates a new DomainListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainListResponseWithDefaults() *DomainListResponse {
	this := DomainListResponse{}
	return &this
}

// GetDomains returns the Domains field value if set, zero value otherwise.
func (o *DomainListResponse) GetDomains() []DomainResponse {
	if o == nil || o.Domains == nil {
		var ret []DomainResponse
		return ret
	}
	return o.Domains
}

// GetDomainsOk returns a tuple with the Domains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainListResponse) GetDomainsOk() ([]DomainResponse, bool) {
	if o == nil || o.Domains == nil {
		return nil, false
	}
	return o.Domains, true
}

// HasDomains returns a boolean if a field has been set.
func (o *DomainListResponse) HasDomains() bool {
	if o != nil && o.Domains != nil {
		return true
	}

	return false
}

// SetDomains gets a reference to the given []DomainResponse and assigns it to the Domains field.
func (o *DomainListResponse) SetDomains(v []DomainResponse) {
	o.Domains = v
}

func (o DomainListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Domains != nil {
		toSerialize["domains"] = o.Domains
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DomainListResponse) UnmarshalJSON(bytes []byte) (err error) {
	varDomainListResponse := _DomainListResponse{}

	err = json.Unmarshal(bytes, &varDomainListResponse)
	if err == nil {
		*o = DomainListResponse(varDomainListResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "domains")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableDomainListResponse struct {
	value *DomainListResponse
	isSet bool
}

func (v NullableDomainListResponse) Get() *DomainListResponse {
	return v.value
}

func (v *NullableDomainListResponse) Set(val *DomainListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainListResponse(val *DomainListResponse) *NullableDomainListResponse {
	return &NullableDomainListResponse{value: val, isSet: true}
}

func (v NullableDomainListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

