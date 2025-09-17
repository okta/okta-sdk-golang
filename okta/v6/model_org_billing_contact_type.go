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
)

// checks if the OrgBillingContactType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrgBillingContactType{}

// OrgBillingContactType Org billing contact
type OrgBillingContactType struct {
	// Type of contact
	ContactType          *string                     `json:"contactType,omitempty"`
	Links                *OrgBillingContactTypeLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrgBillingContactType OrgBillingContactType

// NewOrgBillingContactType instantiates a new OrgBillingContactType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgBillingContactType() *OrgBillingContactType {
	this := OrgBillingContactType{}
	return &this
}

// NewOrgBillingContactTypeWithDefaults instantiates a new OrgBillingContactType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgBillingContactTypeWithDefaults() *OrgBillingContactType {
	this := OrgBillingContactType{}
	return &this
}

// GetContactType returns the ContactType field value if set, zero value otherwise.
func (o *OrgBillingContactType) GetContactType() string {
	if o == nil || IsNil(o.ContactType) {
		var ret string
		return ret
	}
	return *o.ContactType
}

// GetContactTypeOk returns a tuple with the ContactType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgBillingContactType) GetContactTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ContactType) {
		return nil, false
	}
	return o.ContactType, true
}

// HasContactType returns a boolean if a field has been set.
func (o *OrgBillingContactType) HasContactType() bool {
	if o != nil && !IsNil(o.ContactType) {
		return true
	}

	return false
}

// SetContactType gets a reference to the given string and assigns it to the ContactType field.
func (o *OrgBillingContactType) SetContactType(v string) {
	o.ContactType = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *OrgBillingContactType) GetLinks() OrgBillingContactTypeLinks {
	if o == nil || IsNil(o.Links) {
		var ret OrgBillingContactTypeLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgBillingContactType) GetLinksOk() (*OrgBillingContactTypeLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *OrgBillingContactType) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given OrgBillingContactTypeLinks and assigns it to the Links field.
func (o *OrgBillingContactType) SetLinks(v OrgBillingContactTypeLinks) {
	o.Links = &v
}

func (o OrgBillingContactType) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrgBillingContactType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContactType) {
		toSerialize["contactType"] = o.ContactType
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrgBillingContactType) UnmarshalJSON(data []byte) (err error) {
	varOrgBillingContactType := _OrgBillingContactType{}

	err = json.Unmarshal(data, &varOrgBillingContactType)

	if err != nil {
		return err
	}

	*o = OrgBillingContactType(varOrgBillingContactType)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "contactType")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrgBillingContactType struct {
	value *OrgBillingContactType
	isSet bool
}

func (v NullableOrgBillingContactType) Get() *OrgBillingContactType {
	return v.value
}

func (v *NullableOrgBillingContactType) Set(val *OrgBillingContactType) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgBillingContactType) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgBillingContactType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgBillingContactType(val *OrgBillingContactType) *NullableOrgBillingContactType {
	return &NullableOrgBillingContactType{value: val, isSet: true}
}

func (v NullableOrgBillingContactType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgBillingContactType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
