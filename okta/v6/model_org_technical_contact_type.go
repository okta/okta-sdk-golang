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

// OrgTechnicalContactType Org technical contact
type OrgTechnicalContactType struct {
	// Type of contact
	ContactType *string `json:"contactType,omitempty"`
	Links *OrgTechnicalContactTypeLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrgTechnicalContactType OrgTechnicalContactType

// NewOrgTechnicalContactType instantiates a new OrgTechnicalContactType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgTechnicalContactType() *OrgTechnicalContactType {
	this := OrgTechnicalContactType{}
	return &this
}

// NewOrgTechnicalContactTypeWithDefaults instantiates a new OrgTechnicalContactType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgTechnicalContactTypeWithDefaults() *OrgTechnicalContactType {
	this := OrgTechnicalContactType{}
	return &this
}

// GetContactType returns the ContactType field value if set, zero value otherwise.
func (o *OrgTechnicalContactType) GetContactType() string {
	if o == nil || o.ContactType == nil {
		var ret string
		return ret
	}
	return *o.ContactType
}

// GetContactTypeOk returns a tuple with the ContactType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgTechnicalContactType) GetContactTypeOk() (*string, bool) {
	if o == nil || o.ContactType == nil {
		return nil, false
	}
	return o.ContactType, true
}

// HasContactType returns a boolean if a field has been set.
func (o *OrgTechnicalContactType) HasContactType() bool {
	if o != nil && o.ContactType != nil {
		return true
	}

	return false
}

// SetContactType gets a reference to the given string and assigns it to the ContactType field.
func (o *OrgTechnicalContactType) SetContactType(v string) {
	o.ContactType = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *OrgTechnicalContactType) GetLinks() OrgTechnicalContactTypeLinks {
	if o == nil || o.Links == nil {
		var ret OrgTechnicalContactTypeLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgTechnicalContactType) GetLinksOk() (*OrgTechnicalContactTypeLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *OrgTechnicalContactType) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given OrgTechnicalContactTypeLinks and assigns it to the Links field.
func (o *OrgTechnicalContactType) SetLinks(v OrgTechnicalContactTypeLinks) {
	o.Links = &v
}

func (o OrgTechnicalContactType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ContactType != nil {
		toSerialize["contactType"] = o.ContactType
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OrgTechnicalContactType) UnmarshalJSON(bytes []byte) (err error) {
	varOrgTechnicalContactType := _OrgTechnicalContactType{}

	err = json.Unmarshal(bytes, &varOrgTechnicalContactType)
	if err == nil {
		*o = OrgTechnicalContactType(varOrgTechnicalContactType)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "contactType")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableOrgTechnicalContactType struct {
	value *OrgTechnicalContactType
	isSet bool
}

func (v NullableOrgTechnicalContactType) Get() *OrgTechnicalContactType {
	return v.value
}

func (v *NullableOrgTechnicalContactType) Set(val *OrgTechnicalContactType) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgTechnicalContactType) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgTechnicalContactType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgTechnicalContactType(val *OrgTechnicalContactType) *NullableOrgTechnicalContactType {
	return &NullableOrgTechnicalContactType{value: val, isSet: true}
}

func (v NullableOrgTechnicalContactType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgTechnicalContactType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

