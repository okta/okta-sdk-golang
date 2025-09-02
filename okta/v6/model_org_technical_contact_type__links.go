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

// OrgTechnicalContactTypeLinks Specifies link relations (see [Web Linking](https://www.rfc-editor.org/rfc/rfc8288)) available for the org technical Contact Type object using the [JSON Hypertext Application Language](https://datatracker.ietf.org/doc/html/draft-kelly-json-hal-06) specification
type OrgTechnicalContactTypeLinks struct {
	// Link to the org technical [Contact Type User](/openapi/okta-management/management/tag/OrgSettingContact/#tag/OrgSettingContact/operation/getOrgContactUser) resource
	Technical *HrefObject `json:"technical,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrgTechnicalContactTypeLinks OrgTechnicalContactTypeLinks

// NewOrgTechnicalContactTypeLinks instantiates a new OrgTechnicalContactTypeLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgTechnicalContactTypeLinks() *OrgTechnicalContactTypeLinks {
	this := OrgTechnicalContactTypeLinks{}
	return &this
}

// NewOrgTechnicalContactTypeLinksWithDefaults instantiates a new OrgTechnicalContactTypeLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgTechnicalContactTypeLinksWithDefaults() *OrgTechnicalContactTypeLinks {
	this := OrgTechnicalContactTypeLinks{}
	return &this
}

// GetTechnical returns the Technical field value if set, zero value otherwise.
func (o *OrgTechnicalContactTypeLinks) GetTechnical() HrefObject {
	if o == nil || o.Technical == nil {
		var ret HrefObject
		return ret
	}
	return *o.Technical
}

// GetTechnicalOk returns a tuple with the Technical field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgTechnicalContactTypeLinks) GetTechnicalOk() (*HrefObject, bool) {
	if o == nil || o.Technical == nil {
		return nil, false
	}
	return o.Technical, true
}

// HasTechnical returns a boolean if a field has been set.
func (o *OrgTechnicalContactTypeLinks) HasTechnical() bool {
	if o != nil && o.Technical != nil {
		return true
	}

	return false
}

// SetTechnical gets a reference to the given HrefObject and assigns it to the Technical field.
func (o *OrgTechnicalContactTypeLinks) SetTechnical(v HrefObject) {
	o.Technical = &v
}

func (o OrgTechnicalContactTypeLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Technical != nil {
		toSerialize["technical"] = o.Technical
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OrgTechnicalContactTypeLinks) UnmarshalJSON(bytes []byte) (err error) {
	varOrgTechnicalContactTypeLinks := _OrgTechnicalContactTypeLinks{}

	err = json.Unmarshal(bytes, &varOrgTechnicalContactTypeLinks)
	if err == nil {
		*o = OrgTechnicalContactTypeLinks(varOrgTechnicalContactTypeLinks)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "technical")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableOrgTechnicalContactTypeLinks struct {
	value *OrgTechnicalContactTypeLinks
	isSet bool
}

func (v NullableOrgTechnicalContactTypeLinks) Get() *OrgTechnicalContactTypeLinks {
	return v.value
}

func (v *NullableOrgTechnicalContactTypeLinks) Set(val *OrgTechnicalContactTypeLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgTechnicalContactTypeLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgTechnicalContactTypeLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgTechnicalContactTypeLinks(val *OrgTechnicalContactTypeLinks) *NullableOrgTechnicalContactTypeLinks {
	return &NullableOrgTechnicalContactTypeLinks{value: val, isSet: true}
}

func (v NullableOrgTechnicalContactTypeLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgTechnicalContactTypeLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

