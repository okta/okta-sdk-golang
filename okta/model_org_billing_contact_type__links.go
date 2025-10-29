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

// checks if the OrgBillingContactTypeLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrgBillingContactTypeLinks{}

// OrgBillingContactTypeLinks Specifies link relations (see [Web Linking](https://www.rfc-editor.org/rfc/rfc8288)) available for the org billing contact type object using the [JSON Hypertext Application Language](https://datatracker.ietf.org/doc/html/draft-kelly-json-hal-06) specification
type OrgBillingContactTypeLinks struct {
	// Link to the org billing [contact type user](/openapi/okta-management/management/tag/OrgSettingContact/#tag/OrgSettingContact/operation/getOrgContactUser) resource
	Billing              *HrefObject `json:"billing,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrgBillingContactTypeLinks OrgBillingContactTypeLinks

// NewOrgBillingContactTypeLinks instantiates a new OrgBillingContactTypeLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgBillingContactTypeLinks() *OrgBillingContactTypeLinks {
	this := OrgBillingContactTypeLinks{}
	return &this
}

// NewOrgBillingContactTypeLinksWithDefaults instantiates a new OrgBillingContactTypeLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgBillingContactTypeLinksWithDefaults() *OrgBillingContactTypeLinks {
	this := OrgBillingContactTypeLinks{}
	return &this
}

// GetBilling returns the Billing field value if set, zero value otherwise.
func (o *OrgBillingContactTypeLinks) GetBilling() HrefObject {
	if o == nil || IsNil(o.Billing) {
		var ret HrefObject
		return ret
	}
	return *o.Billing
}

// GetBillingOk returns a tuple with the Billing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgBillingContactTypeLinks) GetBillingOk() (*HrefObject, bool) {
	if o == nil || IsNil(o.Billing) {
		return nil, false
	}
	return o.Billing, true
}

// HasBilling returns a boolean if a field has been set.
func (o *OrgBillingContactTypeLinks) HasBilling() bool {
	if o != nil && !IsNil(o.Billing) {
		return true
	}

	return false
}

// SetBilling gets a reference to the given HrefObject and assigns it to the Billing field.
func (o *OrgBillingContactTypeLinks) SetBilling(v HrefObject) {
	o.Billing = &v
}

func (o OrgBillingContactTypeLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrgBillingContactTypeLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Billing) {
		toSerialize["billing"] = o.Billing
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrgBillingContactTypeLinks) UnmarshalJSON(data []byte) (err error) {
	varOrgBillingContactTypeLinks := _OrgBillingContactTypeLinks{}

	err = json.Unmarshal(data, &varOrgBillingContactTypeLinks)

	if err != nil {
		return err
	}

	*o = OrgBillingContactTypeLinks(varOrgBillingContactTypeLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "billing")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrgBillingContactTypeLinks struct {
	value *OrgBillingContactTypeLinks
	isSet bool
}

func (v NullableOrgBillingContactTypeLinks) Get() *OrgBillingContactTypeLinks {
	return v.value
}

func (v *NullableOrgBillingContactTypeLinks) Set(val *OrgBillingContactTypeLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgBillingContactTypeLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgBillingContactTypeLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgBillingContactTypeLinks(val *OrgBillingContactTypeLinks) *NullableOrgBillingContactTypeLinks {
	return &NullableOrgBillingContactTypeLinks{value: val, isSet: true}
}

func (v NullableOrgBillingContactTypeLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgBillingContactTypeLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
