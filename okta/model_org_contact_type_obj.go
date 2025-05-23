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

// OrgContactTypeObj struct for OrgContactTypeObj
type OrgContactTypeObj struct {
	ContactType *string `json:"contactType,omitempty"`
	Links *LinksSelf `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrgContactTypeObj OrgContactTypeObj

// NewOrgContactTypeObj instantiates a new OrgContactTypeObj object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgContactTypeObj() *OrgContactTypeObj {
	this := OrgContactTypeObj{}
	return &this
}

// NewOrgContactTypeObjWithDefaults instantiates a new OrgContactTypeObj object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgContactTypeObjWithDefaults() *OrgContactTypeObj {
	this := OrgContactTypeObj{}
	return &this
}

// GetContactType returns the ContactType field value if set, zero value otherwise.
func (o *OrgContactTypeObj) GetContactType() string {
	if o == nil || o.ContactType == nil {
		var ret string
		return ret
	}
	return *o.ContactType
}

// GetContactTypeOk returns a tuple with the ContactType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgContactTypeObj) GetContactTypeOk() (*string, bool) {
	if o == nil || o.ContactType == nil {
		return nil, false
	}
	return o.ContactType, true
}

// HasContactType returns a boolean if a field has been set.
func (o *OrgContactTypeObj) HasContactType() bool {
	if o != nil && o.ContactType != nil {
		return true
	}

	return false
}

// SetContactType gets a reference to the given string and assigns it to the ContactType field.
func (o *OrgContactTypeObj) SetContactType(v string) {
	o.ContactType = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *OrgContactTypeObj) GetLinks() LinksSelf {
	if o == nil || o.Links == nil {
		var ret LinksSelf
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgContactTypeObj) GetLinksOk() (*LinksSelf, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *OrgContactTypeObj) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksSelf and assigns it to the Links field.
func (o *OrgContactTypeObj) SetLinks(v LinksSelf) {
	o.Links = &v
}

func (o OrgContactTypeObj) MarshalJSON() ([]byte, error) {
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

func (o *OrgContactTypeObj) UnmarshalJSON(bytes []byte) (err error) {
	varOrgContactTypeObj := _OrgContactTypeObj{}

	err = json.Unmarshal(bytes, &varOrgContactTypeObj)
	if err == nil {
		*o = OrgContactTypeObj(varOrgContactTypeObj)
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

type NullableOrgContactTypeObj struct {
	value *OrgContactTypeObj
	isSet bool
}

func (v NullableOrgContactTypeObj) Get() *OrgContactTypeObj {
	return v.value
}

func (v *NullableOrgContactTypeObj) Set(val *OrgContactTypeObj) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgContactTypeObj) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgContactTypeObj) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgContactTypeObj(val *OrgContactTypeObj) *NullableOrgContactTypeObj {
	return &NullableOrgContactTypeObj{value: val, isSet: true}
}

func (v NullableOrgContactTypeObj) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgContactTypeObj) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

