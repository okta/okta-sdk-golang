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

// CsrMetadataSubject struct for CsrMetadataSubject
type CsrMetadataSubject struct {
	CommonName *string `json:"commonName,omitempty"`
	CountryName *string `json:"countryName,omitempty"`
	LocalityName *string `json:"localityName,omitempty"`
	OrganizationalUnitName *string `json:"organizationalUnitName,omitempty"`
	OrganizationName *string `json:"organizationName,omitempty"`
	StateOrProvinceName *string `json:"stateOrProvinceName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CsrMetadataSubject CsrMetadataSubject

// NewCsrMetadataSubject instantiates a new CsrMetadataSubject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCsrMetadataSubject() *CsrMetadataSubject {
	this := CsrMetadataSubject{}
	return &this
}

// NewCsrMetadataSubjectWithDefaults instantiates a new CsrMetadataSubject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCsrMetadataSubjectWithDefaults() *CsrMetadataSubject {
	this := CsrMetadataSubject{}
	return &this
}

// GetCommonName returns the CommonName field value if set, zero value otherwise.
func (o *CsrMetadataSubject) GetCommonName() string {
	if o == nil || o.CommonName == nil {
		var ret string
		return ret
	}
	return *o.CommonName
}

// GetCommonNameOk returns a tuple with the CommonName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CsrMetadataSubject) GetCommonNameOk() (*string, bool) {
	if o == nil || o.CommonName == nil {
		return nil, false
	}
	return o.CommonName, true
}

// HasCommonName returns a boolean if a field has been set.
func (o *CsrMetadataSubject) HasCommonName() bool {
	if o != nil && o.CommonName != nil {
		return true
	}

	return false
}

// SetCommonName gets a reference to the given string and assigns it to the CommonName field.
func (o *CsrMetadataSubject) SetCommonName(v string) {
	o.CommonName = &v
}

// GetCountryName returns the CountryName field value if set, zero value otherwise.
func (o *CsrMetadataSubject) GetCountryName() string {
	if o == nil || o.CountryName == nil {
		var ret string
		return ret
	}
	return *o.CountryName
}

// GetCountryNameOk returns a tuple with the CountryName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CsrMetadataSubject) GetCountryNameOk() (*string, bool) {
	if o == nil || o.CountryName == nil {
		return nil, false
	}
	return o.CountryName, true
}

// HasCountryName returns a boolean if a field has been set.
func (o *CsrMetadataSubject) HasCountryName() bool {
	if o != nil && o.CountryName != nil {
		return true
	}

	return false
}

// SetCountryName gets a reference to the given string and assigns it to the CountryName field.
func (o *CsrMetadataSubject) SetCountryName(v string) {
	o.CountryName = &v
}

// GetLocalityName returns the LocalityName field value if set, zero value otherwise.
func (o *CsrMetadataSubject) GetLocalityName() string {
	if o == nil || o.LocalityName == nil {
		var ret string
		return ret
	}
	return *o.LocalityName
}

// GetLocalityNameOk returns a tuple with the LocalityName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CsrMetadataSubject) GetLocalityNameOk() (*string, bool) {
	if o == nil || o.LocalityName == nil {
		return nil, false
	}
	return o.LocalityName, true
}

// HasLocalityName returns a boolean if a field has been set.
func (o *CsrMetadataSubject) HasLocalityName() bool {
	if o != nil && o.LocalityName != nil {
		return true
	}

	return false
}

// SetLocalityName gets a reference to the given string and assigns it to the LocalityName field.
func (o *CsrMetadataSubject) SetLocalityName(v string) {
	o.LocalityName = &v
}

// GetOrganizationalUnitName returns the OrganizationalUnitName field value if set, zero value otherwise.
func (o *CsrMetadataSubject) GetOrganizationalUnitName() string {
	if o == nil || o.OrganizationalUnitName == nil {
		var ret string
		return ret
	}
	return *o.OrganizationalUnitName
}

// GetOrganizationalUnitNameOk returns a tuple with the OrganizationalUnitName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CsrMetadataSubject) GetOrganizationalUnitNameOk() (*string, bool) {
	if o == nil || o.OrganizationalUnitName == nil {
		return nil, false
	}
	return o.OrganizationalUnitName, true
}

// HasOrganizationalUnitName returns a boolean if a field has been set.
func (o *CsrMetadataSubject) HasOrganizationalUnitName() bool {
	if o != nil && o.OrganizationalUnitName != nil {
		return true
	}

	return false
}

// SetOrganizationalUnitName gets a reference to the given string and assigns it to the OrganizationalUnitName field.
func (o *CsrMetadataSubject) SetOrganizationalUnitName(v string) {
	o.OrganizationalUnitName = &v
}

// GetOrganizationName returns the OrganizationName field value if set, zero value otherwise.
func (o *CsrMetadataSubject) GetOrganizationName() string {
	if o == nil || o.OrganizationName == nil {
		var ret string
		return ret
	}
	return *o.OrganizationName
}

// GetOrganizationNameOk returns a tuple with the OrganizationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CsrMetadataSubject) GetOrganizationNameOk() (*string, bool) {
	if o == nil || o.OrganizationName == nil {
		return nil, false
	}
	return o.OrganizationName, true
}

// HasOrganizationName returns a boolean if a field has been set.
func (o *CsrMetadataSubject) HasOrganizationName() bool {
	if o != nil && o.OrganizationName != nil {
		return true
	}

	return false
}

// SetOrganizationName gets a reference to the given string and assigns it to the OrganizationName field.
func (o *CsrMetadataSubject) SetOrganizationName(v string) {
	o.OrganizationName = &v
}

// GetStateOrProvinceName returns the StateOrProvinceName field value if set, zero value otherwise.
func (o *CsrMetadataSubject) GetStateOrProvinceName() string {
	if o == nil || o.StateOrProvinceName == nil {
		var ret string
		return ret
	}
	return *o.StateOrProvinceName
}

// GetStateOrProvinceNameOk returns a tuple with the StateOrProvinceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CsrMetadataSubject) GetStateOrProvinceNameOk() (*string, bool) {
	if o == nil || o.StateOrProvinceName == nil {
		return nil, false
	}
	return o.StateOrProvinceName, true
}

// HasStateOrProvinceName returns a boolean if a field has been set.
func (o *CsrMetadataSubject) HasStateOrProvinceName() bool {
	if o != nil && o.StateOrProvinceName != nil {
		return true
	}

	return false
}

// SetStateOrProvinceName gets a reference to the given string and assigns it to the StateOrProvinceName field.
func (o *CsrMetadataSubject) SetStateOrProvinceName(v string) {
	o.StateOrProvinceName = &v
}

func (o CsrMetadataSubject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CommonName != nil {
		toSerialize["commonName"] = o.CommonName
	}
	if o.CountryName != nil {
		toSerialize["countryName"] = o.CountryName
	}
	if o.LocalityName != nil {
		toSerialize["localityName"] = o.LocalityName
	}
	if o.OrganizationalUnitName != nil {
		toSerialize["organizationalUnitName"] = o.OrganizationalUnitName
	}
	if o.OrganizationName != nil {
		toSerialize["organizationName"] = o.OrganizationName
	}
	if o.StateOrProvinceName != nil {
		toSerialize["stateOrProvinceName"] = o.StateOrProvinceName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CsrMetadataSubject) UnmarshalJSON(bytes []byte) (err error) {
	varCsrMetadataSubject := _CsrMetadataSubject{}

	err = json.Unmarshal(bytes, &varCsrMetadataSubject)
	if err == nil {
		*o = CsrMetadataSubject(varCsrMetadataSubject)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "commonName")
		delete(additionalProperties, "countryName")
		delete(additionalProperties, "localityName")
		delete(additionalProperties, "organizationalUnitName")
		delete(additionalProperties, "organizationName")
		delete(additionalProperties, "stateOrProvinceName")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableCsrMetadataSubject struct {
	value *CsrMetadataSubject
	isSet bool
}

func (v NullableCsrMetadataSubject) Get() *CsrMetadataSubject {
	return v.value
}

func (v *NullableCsrMetadataSubject) Set(val *CsrMetadataSubject) {
	v.value = val
	v.isSet = true
}

func (v NullableCsrMetadataSubject) IsSet() bool {
	return v.isSet
}

func (v *NullableCsrMetadataSubject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCsrMetadataSubject(val *CsrMetadataSubject) *NullableCsrMetadataSubject {
	return &NullableCsrMetadataSubject{value: val, isSet: true}
}

func (v NullableCsrMetadataSubject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCsrMetadataSubject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

