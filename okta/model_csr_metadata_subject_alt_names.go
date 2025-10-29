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

// checks if the CsrMetadataSubjectAltNames type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CsrMetadataSubjectAltNames{}

// CsrMetadataSubjectAltNames struct for CsrMetadataSubjectAltNames
type CsrMetadataSubjectAltNames struct {
	// DNS names of the subject
	DnsNames             []string `json:"dnsNames,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CsrMetadataSubjectAltNames CsrMetadataSubjectAltNames

// NewCsrMetadataSubjectAltNames instantiates a new CsrMetadataSubjectAltNames object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCsrMetadataSubjectAltNames() *CsrMetadataSubjectAltNames {
	this := CsrMetadataSubjectAltNames{}
	return &this
}

// NewCsrMetadataSubjectAltNamesWithDefaults instantiates a new CsrMetadataSubjectAltNames object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCsrMetadataSubjectAltNamesWithDefaults() *CsrMetadataSubjectAltNames {
	this := CsrMetadataSubjectAltNames{}
	return &this
}

// GetDnsNames returns the DnsNames field value if set, zero value otherwise.
func (o *CsrMetadataSubjectAltNames) GetDnsNames() []string {
	if o == nil || IsNil(o.DnsNames) {
		var ret []string
		return ret
	}
	return o.DnsNames
}

// GetDnsNamesOk returns a tuple with the DnsNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CsrMetadataSubjectAltNames) GetDnsNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.DnsNames) {
		return nil, false
	}
	return o.DnsNames, true
}

// HasDnsNames returns a boolean if a field has been set.
func (o *CsrMetadataSubjectAltNames) HasDnsNames() bool {
	if o != nil && !IsNil(o.DnsNames) {
		return true
	}

	return false
}

// SetDnsNames gets a reference to the given []string and assigns it to the DnsNames field.
func (o *CsrMetadataSubjectAltNames) SetDnsNames(v []string) {
	o.DnsNames = v
}

func (o CsrMetadataSubjectAltNames) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CsrMetadataSubjectAltNames) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DnsNames) {
		toSerialize["dnsNames"] = o.DnsNames
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CsrMetadataSubjectAltNames) UnmarshalJSON(data []byte) (err error) {
	varCsrMetadataSubjectAltNames := _CsrMetadataSubjectAltNames{}

	err = json.Unmarshal(data, &varCsrMetadataSubjectAltNames)

	if err != nil {
		return err
	}

	*o = CsrMetadataSubjectAltNames(varCsrMetadataSubjectAltNames)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "dnsNames")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCsrMetadataSubjectAltNames struct {
	value *CsrMetadataSubjectAltNames
	isSet bool
}

func (v NullableCsrMetadataSubjectAltNames) Get() *CsrMetadataSubjectAltNames {
	return v.value
}

func (v *NullableCsrMetadataSubjectAltNames) Set(val *CsrMetadataSubjectAltNames) {
	v.value = val
	v.isSet = true
}

func (v NullableCsrMetadataSubjectAltNames) IsSet() bool {
	return v.isSet
}

func (v *NullableCsrMetadataSubjectAltNames) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCsrMetadataSubjectAltNames(val *CsrMetadataSubjectAltNames) *NullableCsrMetadataSubjectAltNames {
	return &NullableCsrMetadataSubjectAltNames{value: val, isSet: true}
}

func (v NullableCsrMetadataSubjectAltNames) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCsrMetadataSubjectAltNames) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
