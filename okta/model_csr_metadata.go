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

// CsrMetadata struct for CsrMetadata
type CsrMetadata struct {
	Subject *CsrMetadataSubject `json:"subject,omitempty"`
	SubjectAltNames *CsrMetadataSubjectAltNames `json:"subjectAltNames,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CsrMetadata CsrMetadata

// NewCsrMetadata instantiates a new CsrMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCsrMetadata() *CsrMetadata {
	this := CsrMetadata{}
	return &this
}

// NewCsrMetadataWithDefaults instantiates a new CsrMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCsrMetadataWithDefaults() *CsrMetadata {
	this := CsrMetadata{}
	return &this
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *CsrMetadata) GetSubject() CsrMetadataSubject {
	if o == nil || o.Subject == nil {
		var ret CsrMetadataSubject
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CsrMetadata) GetSubjectOk() (*CsrMetadataSubject, bool) {
	if o == nil || o.Subject == nil {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *CsrMetadata) HasSubject() bool {
	if o != nil && o.Subject != nil {
		return true
	}

	return false
}

// SetSubject gets a reference to the given CsrMetadataSubject and assigns it to the Subject field.
func (o *CsrMetadata) SetSubject(v CsrMetadataSubject) {
	o.Subject = &v
}

// GetSubjectAltNames returns the SubjectAltNames field value if set, zero value otherwise.
func (o *CsrMetadata) GetSubjectAltNames() CsrMetadataSubjectAltNames {
	if o == nil || o.SubjectAltNames == nil {
		var ret CsrMetadataSubjectAltNames
		return ret
	}
	return *o.SubjectAltNames
}

// GetSubjectAltNamesOk returns a tuple with the SubjectAltNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CsrMetadata) GetSubjectAltNamesOk() (*CsrMetadataSubjectAltNames, bool) {
	if o == nil || o.SubjectAltNames == nil {
		return nil, false
	}
	return o.SubjectAltNames, true
}

// HasSubjectAltNames returns a boolean if a field has been set.
func (o *CsrMetadata) HasSubjectAltNames() bool {
	if o != nil && o.SubjectAltNames != nil {
		return true
	}

	return false
}

// SetSubjectAltNames gets a reference to the given CsrMetadataSubjectAltNames and assigns it to the SubjectAltNames field.
func (o *CsrMetadata) SetSubjectAltNames(v CsrMetadataSubjectAltNames) {
	o.SubjectAltNames = &v
}

func (o CsrMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Subject != nil {
		toSerialize["subject"] = o.Subject
	}
	if o.SubjectAltNames != nil {
		toSerialize["subjectAltNames"] = o.SubjectAltNames
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CsrMetadata) UnmarshalJSON(bytes []byte) (err error) {
	varCsrMetadata := _CsrMetadata{}

	err = json.Unmarshal(bytes, &varCsrMetadata)
	if err == nil {
		*o = CsrMetadata(varCsrMetadata)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "subject")
		delete(additionalProperties, "subjectAltNames")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableCsrMetadata struct {
	value *CsrMetadata
	isSet bool
}

func (v NullableCsrMetadata) Get() *CsrMetadata {
	return v.value
}

func (v *NullableCsrMetadata) Set(val *CsrMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableCsrMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableCsrMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCsrMetadata(val *CsrMetadata) *NullableCsrMetadata {
	return &NullableCsrMetadata{value: val, isSet: true}
}

func (v NullableCsrMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCsrMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

